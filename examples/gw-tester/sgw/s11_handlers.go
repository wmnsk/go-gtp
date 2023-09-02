// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
)

func (s *sgw) handleCreateSessionRequest(s11Conn *gtpv2.Conn, mmeAddr net.Addr, msg message.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), mmeAddr)
	if s.mc != nil {
		s.mc.messagesReceived.WithLabelValues(mmeAddr.String(), msg.MessageTypeName()).Inc()
	}

	s11Session := gtpv2.NewSession(mmeAddr, &gtpv2.Subscriber{Location: &gtpv2.Location{}})
	s11Bearer := s11Session.GetDefaultBearer()

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	csReqFromMME := msg.(*message.CreateSessionRequest)

	var pgwAddrString string
	if fteidcIE := csReqFromMME.PGWS5S8FTEIDC; fteidcIE != nil {
		ip, err := fteidcIE.IPAddress()
		if err != nil {
			return err
		}
		pgwAddrString = ip + gtpv2.GTPCPort

		teid, err := fteidcIE.TEID()
		if err != nil {
			return err
		}
		s11Session.AddTEID(gtpv2.IFTypeS5S8PGWGTPC, teid)
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.FullyQualifiedTEID}
	}

	if fteidcIE := csReqFromMME.SenderFTEIDC; fteidcIE != nil {
		teid, err := fteidcIE.TEID()
		if err != nil {
			return err
		}
		s11Session.AddTEID(gtpv2.IFTypeS11MMEGTPC, teid)
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.FullyQualifiedTEID}
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
	s11sgwFTEID := s11Conn.NewSenderFTEID(s.s11IP, "")
	s11sgwTEID := s11sgwFTEID.MustTEID()
	s11Conn.RegisterSession(s11sgwTEID, s11Session)

	s5cFTEID := s.s5cConn.NewSenderFTEID(s.s5cIP, "")
	s5uFTEID := s.s5uConn.NewFTEID(gtpv2.IFTypeS5S8SGWGTPU, s.s5uIP, "").WithInstance(2)

	s5Session, seq, err := s.s5cConn.CreateSession(
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
		ie.NewFullyQualifiedCSID(s.s5uIP, 1).WithInstance(1),
	)
	if err != nil {
		return err
	}
	s5Session.AddTEID(s5uFTEID.MustInterfaceType(), s5uFTEID.MustTEID())

	log.Printf("Sent Create Session Request to %s for %s", pgwAddrString, s5Session.IMSI)
	if s.mc != nil {
		s.mc.messagesSent.WithLabelValues(mmeAddr.String(), "Create Session Request").Inc()
	}

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
		log.Printf(
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

		bearer := s11Session.GetDefaultBearer()
		if ie := csRspFromPGW.PAA; ie != nil {
			bearer.SubscriberIP, err = ie.IPAddress()
			if err != nil {
				return err
			}
		}
	default:
		s11Conn.RemoveSession(s11Session)
		return &gtpv2.UnexpectedTypeError{Msg: incomingMsg}
	}

	// if everything in CreateSessionResponse seems OK, relay it to MME.
	s1usgwFTEID := s.s1uConn.NewFTEID(gtpv2.IFTypeS1USGWGTPU, s.s1uIP, "")
	csRspFromSGW = csRspFromPGW
	csRspFromSGW.SenderFTEIDC = s11sgwFTEID
	csRspFromSGW.SGWFQCSID = ie.NewFullyQualifiedCSID(s.s1uIP, 1).WithInstance(1)
	csRspFromSGW.BearerContextsCreated[0].Add(s1usgwFTEID)
	csRspFromSGW.BearerContextsCreated[0].Remove(ie.ChargingID, 0)
	csRspFromSGW.SetTEID(s11mmeTEID)
	csRspFromSGW.SetLength()

	s11Session.AddTEID(s11sgwFTEID.MustInterfaceType(), s11sgwTEID)
	s11Session.AddTEID(s1usgwFTEID.MustInterfaceType(), s1usgwFTEID.MustTEID())

	if err := s11Conn.RespondTo(mmeAddr, csReqFromMME, csRspFromSGW); err != nil {
		s11Conn.RemoveSession(s11Session)
		return err
	}
	if s.mc != nil {
		s.mc.messagesSent.WithLabelValues(mmeAddr.String(), csRspFromSGW.MessageTypeName()).Inc()
	}

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
		s11Conn.RemoveSession(s11Session)
		return err
	}

	log.Printf(
		"Session created with MME and P-GW for Subscriber: %s;\n\tS11 MME:  %s, TEID->: %#x, TEID<-: %#x\n\tS5C P-GW: %s, TEID->: %#x, TEID<-: %#x",
		s5Session.Subscriber.IMSI, mmeAddr, s11mmeTEID, s11sgwTEID, pgwAddrString, s5cpgwTEID, s5csgwTEID,
	)
	return nil
}

func (s *sgw) handleModifyBearerRequest(s11Conn *gtpv2.Conn, mmeAddr net.Addr, msg message.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), mmeAddr)
	if s.mc != nil {
		s.mc.messagesReceived.WithLabelValues(mmeAddr.String(), msg.MessageTypeName()).Inc()
	}

	s11Session, err := s11Conn.GetSessionByTEID(msg.TEID(), mmeAddr)
	if err != nil {
		return err
	}
	s5cSession, err := s.s5cConn.GetSessionByIMSI(s11Session.IMSI)
	if err != nil {
		return err
	}
	s1uBearer := s11Session.GetDefaultBearer()
	s5uBearer := s5cSession.GetDefaultBearer()

	var enbIP string
	mbReqFromMME := msg.(*message.ModifyBearerRequest)
	if brCtxIE := mbReqFromMME.BearerContextsToBeModified; brCtxIE != nil {
		for _, childIE := range brCtxIE[0].ChildIEs {
			switch childIE.Type {
			case ie.Indication:
				// do nothing in this implementation.
			case ie.FullyQualifiedTEID:
				if err := s.handleFTEIDU(childIE, s11Session, s1uBearer); err != nil {
					return err
				}
				enbIP, err = childIE.IPAddress()
				if err != nil {
					return err
				}
			}
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.BearerContext}
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
	pgwIP, _, err := net.SplitHostPort(s5uBearer.RemoteAddress().String())
	if err != nil {
		return err
	}

	if s.useKernelGTP {
		if err := s.s1uConn.AddTunnelOverride(
			net.ParseIP(enbIP), net.ParseIP(s1uBearer.SubscriberIP), s1uBearer.OutgoingTEID(), s1usgwTEID,
		); err != nil {
			return err
		}
		if err := s.s5uConn.AddTunnelOverride(
			net.ParseIP(pgwIP), net.ParseIP(s5uBearer.SubscriberIP), s5uBearer.OutgoingTEID(), s5usgwTEID,
		); err != nil {
			return err
		}
	} else {
		if err := s.s1uConn.RelayTo(
			s.s5uConn, s1usgwTEID, s5uBearer.OutgoingTEID(), s5uBearer.RemoteAddress(),
		); err != nil {
			return err
		}
		if err := s.s5uConn.RelayTo(
			s.s1uConn, s5usgwTEID, s1uBearer.OutgoingTEID(), s1uBearer.RemoteAddress(),
		); err != nil {
			return err
		}
	}

	mbRspFromSGW := message.NewModifyBearerResponse(
		s11mmeTEID, 0,
		ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
		ie.NewBearerContext(
			ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
			ie.NewEPSBearerID(s1uBearer.EBI),
			ie.NewFullyQualifiedTEID(gtpv2.IFTypeS1USGWGTPU, s1usgwTEID, s.s1uIP, ""),
		),
	)

	if err := s11Conn.RespondTo(mmeAddr, msg, mbRspFromSGW); err != nil {
		return err
	}
	if s.mc != nil {
		s.mc.messagesSent.WithLabelValues(mmeAddr.String(), mbRspFromSGW.MessageTypeName()).Inc()
	}

	log.Printf(
		"Started listening on U-Plane for Subscriber: %s;\n\tS1-U: %s\n\tS5-U: %s",
		s11Session.IMSI, s.s1uConn.LocalAddr(), s.s5uConn.LocalAddr(),
	)
	return nil
}

func (s *sgw) handleDeleteSessionRequest(s11Conn *gtpv2.Conn, mmeAddr net.Addr, msg message.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), mmeAddr)
	if s.mc != nil {
		s.mc.messagesReceived.WithLabelValues(mmeAddr.String(), msg.MessageTypeName()).Inc()
	}

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	dsReqFromMME := msg.(*message.DeleteSessionRequest)

	s11Session, err := s11Conn.GetSessionByTEID(msg.TEID(), mmeAddr)
	if err != nil {
		return err
	}

	s5Session, err := s.s5cConn.GetSessionByIMSI(s11Session.IMSI)
	if err != nil {
		return err
	}

	s5cpgwTEID, err := s5Session.GetTEID(gtpv2.IFTypeS5S8PGWGTPC)
	if err != nil {
		return err
	}

	seq, err := s.s5cConn.DeleteSession(
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

	incomingMessage, err := s11Session.WaitMessage(seq, 5*time.Second)
	if err != nil {
		dsRspFromSGW = message.NewDeleteSessionResponse(
			s11mmeTEID, 0,
			ie.NewCause(gtpv2.CausePGWNotResponding, 0, 0, 0, nil),
		)

		if err := s11Conn.RespondTo(mmeAddr, dsReqFromMME, dsRspFromSGW); err != nil {
			return err
		}
		log.Printf(
			"Sent %s with failure code: %d, target subscriber: %s",
			dsRspFromSGW.MessageTypeName(), gtpv2.CausePGWNotResponding, s11Session.IMSI,
		)
		if s.mc != nil {
			s.mc.messagesSent.WithLabelValues(mmeAddr.String(), dsRspFromSGW.MessageTypeName()).Inc()
		}
		return err
	}

	// use the cause as it is.
	switch m := incomingMessage.(type) {
	case *message.DeleteSessionResponse:
		// move forward
		dsRspFromSGW = m
	default:
		return &gtpv2.UnexpectedTypeError{Msg: incomingMessage}
	}

	dsRspFromSGW.SetTEID(s11mmeTEID)
	if err := s11Conn.RespondTo(mmeAddr, msg, dsRspFromSGW); err != nil {
		return err
	}

	log.Printf("Session deleted for Subscriber: %s", s11Session.IMSI)
	if s.mc != nil {
		s.mc.messagesSent.WithLabelValues(mmeAddr.String(), dsRspFromSGW.MessageTypeName()).Inc()
	}

	s11Conn.RemoveSession(s11Session)
	return nil
}

func (s *sgw) handleDeleteBearerResponse(s11Conn *gtpv2.Conn, mmeAddr net.Addr, msg message.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), mmeAddr)
	if s.mc != nil {
		s.mc.messagesReceived.WithLabelValues(mmeAddr.String(), msg.MessageTypeName()).Inc()
	}

	s11Session, err := s11Conn.GetSessionByTEID(msg.TEID(), mmeAddr)
	if err != nil {
		return err
	}

	s5Session, err := s.s5cConn.GetSessionByIMSI(s11Session.IMSI)
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

func (s *sgw) handleFTEIDU(fteiduIE *ie.IE, session *gtpv2.Session, bearer *gtpv2.Bearer) error {
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
