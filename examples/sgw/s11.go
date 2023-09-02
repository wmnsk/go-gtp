// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Command sgw is a dead simple implementation of S-GW only with GTP-related features.
package main

import (
	"fmt"
	"net"
	"time"

	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
)

func handleCreateSessionRequest(s11Conn *gtpv2.Conn, mmeAddr net.Addr, msg message.Message) error {
	sgw.loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), mmeAddr)

	s11Session := gtpv2.NewSession(mmeAddr, &gtpv2.Subscriber{Location: &gtpv2.Location{}})
	s11Bearer := s11Session.GetDefaultBearer()

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	csReqFromMME := msg.(*message.CreateSessionRequest)

	var pgwAddrString string
	if teidIE := csReqFromMME.PGWS5S8FTEIDC; teidIE != nil {
		ip, err := teidIE.IPAddress()
		if err != nil {
			return err
		}
		pgwAddrString = ip + gtpv2.GTPCPort

		teid, err := teidIE.TEID()
		if err != nil {
			return err
		}
		s11Session.AddTEID(gtpv2.IFTypeS5S8PGWGTPC, teid)
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.FullyQualifiedTEID}
	}
	if senderIE := csReqFromMME.SenderFTEIDC; senderIE != nil {
		teid, err := senderIE.TEID()
		if err != nil {
			return err
		}
		s11Session.AddTEID(gtpv2.IFTypeS11MMEGTPC, teid)
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.FullyQualifiedTEID}
	}

	laddr, err := net.ResolveUDPAddr("udp", *s5c+gtpv2.GTPCPort)
	if err != nil {
		return err
	}
	raddr, err := net.ResolveUDPAddr("udp", pgwAddrString)
	if err != nil {
		return err
	}

	// keep session information retrieved from the message.
	// XXX - should return error if required IE is missing.
	if imsiIE := csReqFromMME.IMSI; imsiIE != nil {
		imsi, err := imsiIE.IMSI()
		if err != nil {
			return err
		}

		// remove previous session for the same subscriber if exists.
		sess, err := s11Conn.GetSessionByIMSI(imsi)
		if err != nil {
			switch err.(type) {
			case *gtpv2.UnknownIMSIError:
				// whole new session. just ignore.
			default:
				return fmt.Errorf("got something unexpected: %w", err)
			}
		} else {
			s11Conn.RemoveSession(sess)
		}

		s11Session.IMSI = imsi
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.IMSI}
	}
	if msisdnIE := csReqFromMME.MSISDN; msisdnIE != nil {
		s11Session.MSISDN, err = msisdnIE.MSISDN()
		if err != nil {
			return err
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.MSISDN}
	}
	if meiIE := csReqFromMME.MEI; meiIE != nil {
		s11Session.IMEI, err = meiIE.MobileEquipmentIdentity()
		if err != nil {
			return err
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.MobileEquipmentIdentity}
	}
	if apnIE := csReqFromMME.APN; apnIE != nil {
		s11Bearer.APN, err = apnIE.AccessPointName()
		if err != nil {
			return err
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.AccessPointName}
	}
	if netIE := csReqFromMME.ServingNetwork; netIE != nil {
		s11Session.MCC, err = netIE.MCC()
		if err != nil {
			return err
		}
		s11Session.MNC, err = netIE.MNC()
		if err != nil {
			return err
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.ServingNetwork}
	}
	if ratIE := csReqFromMME.RATType; ratIE != nil {
		s11Session.RATType, err = ratIE.RATType()
		if err != nil {
			return err
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.RATType}
	}

	s11IP, _, err := net.SplitHostPort(*s11 + gtpv2.GTPCPort)
	if err != nil {
		return fmt.Errorf("failed to get IP for S11: %w", err)
	}
	senderFTEID := s11Conn.NewSenderFTEID(s11IP, "")
	s11sgwTEID := senderFTEID.MustTEID()
	s11Conn.RegisterSession(s11sgwTEID, s11Session)

	s5cIP := laddr.IP.String()
	s5uIP, _, err := net.SplitHostPort(*s5u + gtpv2.GTPCPort)
	if err != nil {
		return err
	}
	s5cFTEID := sgw.s5cConn.NewSenderFTEID(s5cIP, "")
	s5uFTEID := sgw.s5uConn.NewFTEID(gtpv2.IFTypeS5S8SGWGTPU, s5uIP, "").WithInstance(2)

	s5Session, seq, err := sgw.s5cConn.CreateSession(
		raddr,
		csReqFromMME.IMSI, csReqFromMME.MSISDN, csReqFromMME.MEI, csReqFromMME.ServingNetwork,
		csReqFromMME.RATType, csReqFromMME.IndicationFlags, s5cFTEID, csReqFromMME.PGWS5S8FTEIDC,
		csReqFromMME.APN, csReqFromMME.SelectionMode, csReqFromMME.PDNType, csReqFromMME.PAA,
		csReqFromMME.APNRestriction, csReqFromMME.AMBR, csReqFromMME.ULI,
		ie.NewBearerContext(
			ie.NewEPSBearerID(5),
			s5uFTEID,
			ie.NewBearerQoS(1, 2, 1, 0xff, 0, 0, 0, 0),
		),
		csReqFromMME.MMEFQCSID,
		ie.NewFullyQualifiedCSID(s5uIP, 1).WithInstance(1),
	)
	if err != nil {
		return err
	}
	s5Session.AddTEID(s5uFTEID.MustInterfaceType(), s5uFTEID.MustTEID())
	sgw.s5cConn.RegisterSession(s5cFTEID.MustTEID(), s5Session)

	sgw.loggerCh <- fmt.Sprintf("Sent Create Session Request to %s for %s", pgwAddrString, s5Session.IMSI)

	var csRspFromSGW *message.CreateSessionResponse
	s11mmeTEID, err := s11Session.GetTEID(gtpv2.IFTypeS11MMEGTPC)
	if err != nil {
		s11Conn.RemoveSession(s11Session)
		return err
	}
	incomingMsg, err := s11Session.WaitMessage(seq, 5*time.Second)
	if err != nil {
		csRspFromSGW = message.NewCreateSessionResponse(
			s11mmeTEID, 0,
			ie.NewCause(gtpv2.CauseNoResourcesAvailable, 0, 0, 0, nil),
		)

		if err := s11Conn.RespondTo(mmeAddr, csReqFromMME, csRspFromSGW); err != nil {
			s11Conn.RemoveSession(s11Session)
			return err
		}
		sgw.loggerCh <- fmt.Sprintf(
			"Sent %s with failure code: %d, target subscriber: %s",
			csRspFromSGW.MessageTypeName(), gtpv2.CausePGWNotResponding, s11Session.IMSI,
		)
		s11Conn.RemoveSession(s11Session)
		return err
	}

	var csRspFromPGW *message.CreateSessionResponse
	switch m := incomingMsg.(type) {
	case *message.CreateSessionResponse:
		// move forward
		csRspFromPGW = m
	default:
		s11Conn.RemoveSession(s11Session)
		return &gtpv2.UnexpectedTypeError{Msg: incomingMsg}
	}
	// if everything in CreateSessionResponse seems OK, relay it to MME.
	s1uIP, _, err := net.SplitHostPort(*s1u + gtpv2.GTPCPort)
	if err != nil {
		return fmt.Errorf("failed to get IP for S1-U: %w", err)
	}
	s1usgwFTEID := sgw.s1uConn.NewFTEID(gtpv2.IFTypeS1USGWGTPU, s1uIP, "")
	csRspFromSGW = csRspFromPGW
	csRspFromSGW.SenderFTEIDC = senderFTEID
	csRspFromSGW.SGWFQCSID = ie.NewFullyQualifiedCSID(s1uIP, 1).WithInstance(1)
	csRspFromSGW.BearerContextsCreated[0].Add(s1usgwFTEID)
	csRspFromSGW.BearerContextsCreated[0].Remove(ie.ChargingID, 0)
	csRspFromSGW.SetTEID(s11mmeTEID)
	csRspFromSGW.SetLength()

	if err := s11Conn.RespondTo(mmeAddr, csReqFromMME, csRspFromSGW); err != nil {
		s11Conn.RemoveSession(s11Session)
		return err
	}
	s11Session.AddTEID(senderFTEID.MustInterfaceType(), s11sgwTEID)
	s11Session.AddTEID(s1usgwFTEID.MustInterfaceType(), s1usgwFTEID.MustTEID())

	s5cpgwTEID, err := s5Session.GetTEID(gtpv2.IFTypeS5S8PGWGTPC)
	if err != nil {
		s11Conn.RemoveSession(s11Session)
		return err
	}
	s5csgwTEID, err := s5Session.GetTEID(gtpv2.IFTypeS5S8SGWGTPC)
	if err != nil {
		s11Conn.RemoveSession(s11Session)
		return err
	}

	if err := s11Session.Activate(); err != nil {
		sgw.loggerCh <- fmt.Sprintf("Error: %v", err)
		s11Conn.RemoveSession(s11Session)
		return err
	}

	sgw.loggerCh <- fmt.Sprintf(
		"Session created with MME and P-GW for Subscriber: %s;\n\tS11 MME:  %s, TEID->: %#x, TEID<-: %#x\n\tS5C P-GW: %s, TEID->: %#x, TEID<-: %#x",
		s5Session.Subscriber.IMSI, mmeAddr, s11mmeTEID, s11sgwTEID, pgwAddrString, s5cpgwTEID, s5csgwTEID,
	)
	return nil
}

func handleModifyBearerRequest(s11Conn *gtpv2.Conn, mmeAddr net.Addr, msg message.Message) error {
	sgw.loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), mmeAddr)

	s11Session, err := s11Conn.GetSessionByTEID(msg.TEID(), mmeAddr)
	if err != nil {
		return err
	}
	s5cSession, err := sgw.s5cConn.GetSessionByIMSI(s11Session.IMSI)
	if err != nil {
		return err
	}
	s1uBearer := s11Session.GetDefaultBearer()
	s5uBearer := s5cSession.GetDefaultBearer()

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	mbReqFromMME := msg.(*message.ModifyBearerRequest)
	if brCtxIE := mbReqFromMME.BearerContextsToBeModified; brCtxIE != nil {
		for _, childIE := range brCtxIE[0].ChildIEs {
			switch childIE.Type {
			case ie.Indication:
				// do nothing in this example.
				// S-GW should change its beahavior based on indication flags like;
				//  - pass Modify Bearer Request to P-GW if handover is indicated.
				//  - XXX...
			case ie.FullyQualifiedTEID:
				if err := handleFTEIDU(childIE, s11Session, s1uBearer); err != nil {
					return err
				}
			}
		}
	}

	s11mmeTEID, err := s11Session.GetTEID(gtpv2.IFTypeS11MMEGTPC)
	if err != nil {
		return err
	}
	s1usgwTEID, err := s11Session.GetTEID(gtpv2.IFTypeS1USGWGTPU)
	if err != nil {
		return err
	}
	s5usgwTEID, err := s5cSession.GetTEID(gtpv2.IFTypeS5S8SGWGTPU)
	if err != nil {
		return err
	}
	if err := sgw.s1uConn.RelayTo(
		sgw.s5uConn, s1usgwTEID, s5uBearer.OutgoingTEID(), s5uBearer.RemoteAddress(),
	); err != nil {
		return err
	}
	if err := sgw.s5uConn.RelayTo(
		sgw.s1uConn, s5usgwTEID, s1uBearer.OutgoingTEID(), s1uBearer.RemoteAddress(),
	); err != nil {
		return err
	}

	s1uIP, _, err := net.SplitHostPort(*s1u + gtpv2.GTPCPort)
	if err != nil {
		return err
	}
	mbRspFromSGW := message.NewModifyBearerResponse(
		s11mmeTEID, 0,
		ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
		ie.NewBearerContext(
			ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
			ie.NewEPSBearerID(s1uBearer.EBI),
			ie.NewFullyQualifiedTEID(gtpv2.IFTypeS1USGWGTPU, s1usgwTEID, s1uIP, ""),
		),
	)

	if err := s11Conn.RespondTo(mmeAddr, msg, mbRspFromSGW); err != nil {
		return err
	}

	sgw.loggerCh <- fmt.Sprintf(
		"Started listening on U-Plane for Subscriber: %s;\n\tS1-U: %s\n\tS5-U: %s",
		s11Session.IMSI, *s1u, *s5u,
	)
	return nil
}

func handleDeleteSessionRequest(s11Conn *gtpv2.Conn, mmeAddr net.Addr, msg message.Message) error {
	sgw.loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), mmeAddr)

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	dsReqFromMME := msg.(*message.DeleteSessionRequest)

	s11Session, err := s11Conn.GetSessionByTEID(msg.TEID(), mmeAddr)
	if err != nil {
		return err
	}

	s5Session, err := sgw.s5cConn.GetSessionByIMSI(s11Session.IMSI)
	if err != nil {
		return err
	}

	s5cpgwTEID, err := s5Session.GetTEID(gtpv2.IFTypeS5S8PGWGTPC)
	if err != nil {
		return err
	}

	seq, err := sgw.s5cConn.DeleteSession(
		s5cpgwTEID, s5Session,
		ie.NewEPSBearerID(s5Session.GetDefaultBearer().EBI),
	)
	if err != nil {
		return err
	}

	var dsRspFromSGW *message.DeleteSessionResponse
	s11mmeTEID, err := s11Session.GetTEID(gtpv2.IFTypeS11MMEGTPC)
	if err != nil {
		return err
	}

	incomingMsg, err := s11Session.WaitMessage(seq, 5*time.Second)
	if err != nil {
		dsRspFromSGW = message.NewDeleteSessionResponse(
			s11mmeTEID, 0,
			ie.NewCause(gtpv2.CausePGWNotResponding, 0, 0, 0, nil),
		)

		if err := s11Conn.RespondTo(mmeAddr, dsReqFromMME, dsRspFromSGW); err != nil {
			return err
		}
		sgw.loggerCh <- fmt.Sprintf(
			"Sent %s with failure code: %d, target subscriber: %s",
			dsRspFromSGW.MessageTypeName(), gtpv2.CausePGWNotResponding, s11Session.IMSI,
		)
		return err
	}

	// use the cause as it is.
	switch m := incomingMsg.(type) {
	case *message.DeleteSessionResponse:
		// move forward
		dsRspFromSGW = m
	default:
		return &gtpv2.UnexpectedTypeError{Msg: incomingMsg}
	}

	dsRspFromSGW.SetTEID(s11mmeTEID)
	if err := s11Conn.RespondTo(mmeAddr, msg, dsRspFromSGW); err != nil {
		return err
	}

	sgw.loggerCh <- fmt.Sprintf("Session deleted for Subscriber: %s", s11Session.IMSI)
	s11Conn.RemoveSession(s11Session)

	return nil
}

func handleDeleteBearerResponse(s11Conn *gtpv2.Conn, mmeAddr net.Addr, msg message.Message) error {
	sgw.loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), mmeAddr)

	s11Session, err := s11Conn.GetSessionByTEID(msg.TEID(), mmeAddr)
	if err != nil {
		return err
	}

	s5Session, err := sgw.s5cConn.GetSessionByIMSI(s11Session.IMSI)
	if err != nil {
		return err
	}

	if err := gtpv2.PassMessageTo(s5Session, msg, 5*time.Second); err != nil {
		return err
	}

	// remove bearer in handleDeleteBearerRequest instead of doing here,
	// as Delete Bearer Request does not necessarily have EBI.
	return nil
}

func handleFTEIDU(fteiduIE *ie.IE, session *gtpv2.Session, bearer *gtpv2.Bearer) error {
	if fteiduIE.Type != ie.FullyQualifiedTEID {
		return &gtpv2.UnexpectedIEError{IEType: fteiduIE.Type}
	}

	ip, err := fteiduIE.IPAddress()
	if err != nil {
		return err
	}
	addr, err := net.ResolveUDPAddr("udp", ip+gtpv2.GTPUPort)
	if err != nil {
		return err
	}
	bearer.SetRemoteAddress(addr)

	teid, err := fteiduIE.TEID()
	if err != nil {
		return err
	}
	bearer.SetOutgoingTEID(teid)

	it, err := fteiduIE.InterfaceType()
	if err != nil {
		return err
	}
	session.AddTEID(it, teid)
	return nil
}
