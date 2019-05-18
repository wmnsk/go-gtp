// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Command sgw is a dead simple implementation of S-GW only with GTP-related features.
package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/pkg/errors"
	v2 "github.com/wmnsk/go-gtp/v2"
	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
)

func handleCreateSessionRequest(s11Conn *v2.Conn, mmeAddr net.Addr, msg messages.Message) error {
	loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), mmeAddr)

	s11Session := v2.NewSession(mmeAddr, &v2.Subscriber{Location: &v2.Location{}})
	s11Bearer := s11Session.GetDefaultBearer()

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	csReqFromMME := msg.(*messages.CreateSessionRequest)

	var pgwAddrString string
	if ie := csReqFromMME.PGWS5S8FTEIDC; ie != nil {
		pgwAddrString = ie.IPAddress() + ":2123"
		s11Session.AddTEID(v2.IFTypeS5S8PGWGTPC, ie.TEID())
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.FullyQualifiedTEID}
	}
	if ie := csReqFromMME.SenderFTEIDC; ie != nil {
		s11Session.AddTEID(v2.IFTypeS11MMEGTPC, ie.TEID())
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.FullyQualifiedTEID}
	}

	laddr, err := net.ResolveUDPAddr("udp", *s5c)
	if err != nil {
		return err
	}
	raddr, err := net.ResolveUDPAddr("udp", pgwAddrString)
	if err != nil {
		return err
	}
	if s5cConn == nil {
		s5cConn, err = v2.Dial(laddr, raddr, 0, errCh)
		if err != nil {
			return err
		}
	}

	// keep session information retrieved from the message.
	// XXX - should return error if required IE is missing.
	if ie := csReqFromMME.IMSI; ie != nil {
		s11Session.IMSI = ie.IMSI()
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.IMSI}
	}
	if ie := csReqFromMME.MSISDN; ie != nil {
		s11Session.MSISDN = ie.MSISDN()
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.MSISDN}
	}
	if ie := csReqFromMME.MEI; ie != nil {
		s11Session.IMEI = ie.MobileEquipmentIdentity()
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.MobileEquipmentIdentity}
	}
	if ie := csReqFromMME.APN; ie != nil {
		s11Bearer.APN = ie.AccessPointName()
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.AccessPointName}
	}
	if ie := csReqFromMME.ServingNetwork; ie != nil {
		s11Session.MCC = ie.MCC()
		s11Session.MNC = ie.MNC()
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.ServingNetwork}
	}
	if ie := csReqFromMME.RATType; ie != nil {
		s11Session.RATType = ie.RATType()
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.RATType}
	}
	s11Conn.AddSession(s11Session)

	// register handlers in s11Conn before sending CreateSession to P-GW.
	createCh := make(chan *messages.CreateSessionResponse)
	s5cConn.AddHandler(
		messages.MsgTypeCreateSessionResponse,
		func(s5cConn *v2.Conn, pgwAddr net.Addr, msg messages.Message) error {
			loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), pgwAddr)

			s5Session, err := s5cConn.GetSessionByTEID(msg.TEID())
			if err != nil {
				return err
			}

			// assert type to refer to the struct field specific to the message.
			// in general, no need to check if it can be type-asserted, as long as the MessageType is
			// specified correctly in AddHandler().
			csRspFromPGW := msg.(*messages.CreateSessionResponse)

			// check Cause value first.
			if ie := csRspFromPGW.Cause; ie != nil {
				if cause := ie.Cause(); cause != v2.CauseRequestAccepted {
					s5cConn.RemoveSession(s5Session)
					// this is not such a fatal error worth stopping the whole program.
					// in the real case it is better to take some action based on the Cause, though.
					return &v2.ErrCauseNotOK{
						MsgType: csRspFromPGW.MessageTypeName(),
						Cause:   cause,
						Msg:     fmt.Sprintf("subscriber: %s", s5Session.IMSI),
					}
				}
			} else {
				s5cConn.RemoveSession(s5Session)
				return &v2.ErrRequiredIEMissing{
					Type: ies.Cause,
				}
			}

			bearer := s5Session.GetDefaultBearer()
			// retrieve values that P-GW gave.
			if ie := csRspFromPGW.PAA; ie != nil {
				bearer.SubscriberIP = ie.IPAddress()
			} else {
				s5cConn.RemoveSession(s5Session)
				return &v2.ErrRequiredIEMissing{Type: ies.PDNAddressAllocation}
			}
			if ie := csRspFromPGW.PGWS5S8FTEIDC; ie != nil {
				s5Session.AddTEID(ie.InterfaceType(), ie.TEID())
			} else {
				s5cConn.RemoveSession(s5Session)
				return &v2.ErrRequiredIEMissing{Type: ies.FullyQualifiedTEID}
			}

			if brCtxIE := csRspFromPGW.BearerContextsCreated; brCtxIE != nil {
				for _, ie := range brCtxIE.ChildIEs {
					switch ie.Type {
					case ies.Cause:
						if cause := ie.Cause(); cause != v2.CauseRequestAccepted {
							s5cConn.RemoveSession(s5Session)
							return &v2.ErrCauseNotOK{
								MsgType: csRspFromPGW.MessageTypeName(),
								Cause:   cause,
								Msg:     fmt.Sprintf("subscriber: %s", s5Session.IMSI),
							}
						}
					case ies.EPSBearerID:
						bearer.EBI = ie.EPSBearerID()
					case ies.FullyQualifiedTEID:
						if err := handleFTEIDU(ie, s5Session, bearer); err != nil {
							return err
						}
					case ies.ChargingID:
						bearer.ChargingID = ie.ChargingID()
					}
				}
			} else {
				s5cConn.RemoveSession(s5Session)
				return &v2.ErrRequiredIEMissing{Type: ies.BearerContext}
			}

			if err := s5Session.Activate(); err != nil {
				s5cConn.RemoveSession(s5Session)
				return err
			}

			go func() {
				createCh <- csRspFromPGW
			}()
			return nil
		},
	)
	s5cConn.AddHandler(
		messages.MsgTypeDeleteSessionResponse,
		func(s5cConn *v2.Conn, pgwAddr net.Addr, msg messages.Message) error {
			loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), pgwAddr)

			session, err := s5cConn.GetSessionByTEID(msg.TEID())
			if err != nil {
				// this is not such a fatal error worth stopping the whole program.
				loggerCh <- errors.Wrap(err, "Error").Error()
				return nil
			}

			loggerCh <- fmt.Sprintf("Session deleted with P-GW for Subscriber: %s", session.IMSI)
			s5cConn.RemoveSession(session)
			delCh <- struct{}{}
			return nil
		},
	)

	s5cIP := strings.Split(laddr.IP.String(), ":")[0]
	s5cFTEID := s5cConn.NewFTEID(v2.IFTypeS5S8SGWGTPC, s5cIP, "")
	s5uFTEID := s5cConn.NewFTEID(v2.IFTypeS5S8SGWGTPU, s5cIP, "").WithInstance(2)

	s5Session, err := s5cConn.CreateSession(
		raddr,
		csReqFromMME.IMSI, csReqFromMME.MSISDN, csReqFromMME.MEI, csReqFromMME.ServingNetwork,
		csReqFromMME.RATType, csReqFromMME.IndicationFlags, s5cFTEID, csReqFromMME.PGWS5S8FTEIDC,
		csReqFromMME.APN, csReqFromMME.SelectionMode, csReqFromMME.PDNType, csReqFromMME.PAA,
		csReqFromMME.APNRestriction, csReqFromMME.AMBR, csReqFromMME.ULI,
		ies.NewBearerContext(
			ies.NewEPSBearerID(5),
			s5uFTEID,
			ies.NewBearerQoS(1, 2, 1, 0xff, 0, 0, 0, 0),
		),
		csReqFromMME.MMEFQCSID,
		ies.NewFullyQualifiedCSID(s5cIP, 1).WithInstance(1),
	)
	if err != nil {
		return err
	}
	s5Session.AddTEID(s5uFTEID.InterfaceType(), s5uFTEID.TEID())
	s5cConn.AddSession(s5Session)

	loggerCh <- fmt.Sprintf("Sent Create Session Request to %s for %s", pgwAddrString, s5Session.IMSI)

	doneCh := make(chan struct{})
	failCh := make(chan error)
	go func() {
		var csRspFromSGW *messages.CreateSessionResponse
		s11mmeTEID, err := s11Session.GetTEID(v2.IFTypeS11MMEGTPC)
		if err != nil {
			failCh <- err
			return
		}
		select {
		case csRspFromPGW := <-createCh:
			// if everything in CreateSessionResponse seems OK, relay it to MME.
			s11IP := strings.Split(*s11, ":")[0]
			senderFTEID := s11Conn.NewFTEID(v2.IFTypeS11S4SGWGTPC, s11IP, "")
			s1usgwFTEID := s11Conn.NewFTEID(v2.IFTypeS1USGWGTPU, s11IP, "")
			csRspFromSGW = csRspFromPGW
			csRspFromSGW.SenderFTEIDC = senderFTEID
			csRspFromSGW.SGWFQCSID = ies.NewFullyQualifiedCSID(laddr.IP.String(), 1).WithInstance(1)
			csRspFromSGW.BearerContextsCreated.Add(s1usgwFTEID)
			csRspFromSGW.BearerContextsCreated.Remove(ies.ChargingID, 0)
			csRspFromSGW.SetTEID(s11mmeTEID)
			csRspFromSGW.SetLength()

			if err := s11Conn.RespondTo(mmeAddr, csReqFromMME, csRspFromSGW); err != nil {
				failCh <- err
				return
			}
			s11Session.AddTEID(senderFTEID.InterfaceType(), senderFTEID.TEID())
			s11Session.AddTEID(s1usgwFTEID.InterfaceType(), s1usgwFTEID.TEID())
		case <-time.After(5 * time.Second):
			csRspFromSGW = messages.NewCreateSessionResponse(
				s11mmeTEID, 0,
				ies.NewCause(v2.CauseNoResourcesAvailable, 0, 0, 0, nil),
			)

			if err := s11Conn.RespondTo(mmeAddr, csReqFromMME, csRspFromSGW); err != nil {
				failCh <- err
				return
			}
			loggerCh <- fmt.Sprintf("Sent %s with failure code: %d, target subscriber: %s", csRspFromSGW.MessageTypeName(), v2.CausePGWNotResponding, s11Session.IMSI)
			failCh <- v2.ErrTimeout
			return
		}

		s11sgwTEID, err := s11Session.GetTEID(v2.IFTypeS11S4SGWGTPC)
		if err != nil {
			failCh <- err
			return
		}
		s5cpgwTEID, err := s5Session.GetTEID(v2.IFTypeS5S8PGWGTPC)
		if err != nil {
			failCh <- err
			return
		}
		s5csgwTEID, err := s5Session.GetTEID(v2.IFTypeS5S8SGWGTPC)
		if err != nil {
			failCh <- err
			return
		}
		loggerCh <- fmt.Sprintf(
			"Session created with MME and P-GW for Subscriber: %s;\n\tS11 MME:  %s, TEID->: %#x, TEID<-: %#x\n\tS5C P-GW: %s, TEID->: %#x, TEID<-: %#x",
			s5Session.Subscriber.IMSI, mmeAddr, s11mmeTEID, s11sgwTEID, pgwAddrString, s5cpgwTEID, s5csgwTEID,
		)
		doneCh <- struct{}{}
	}()

	select {
	case <-doneCh:
		if s11Session.Activate(); err != nil {
			loggerCh <- errors.Wrap(err, "Error").Error()
			s11Conn.RemoveSession(s11Session)
			return nil
		}
		return nil
	case err := <-failCh:
		s11Conn.RemoveSession(s11Session)
		return err
	case <-time.After(10 * time.Second):
		s11Conn.RemoveSession(s11Session)
		return v2.ErrTimeout
	}
}

func handleModifyBearerRequest(s11Conn *v2.Conn, mmeAddr net.Addr, msg messages.Message) error {
	loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), mmeAddr)

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	s11Session, err := s11Conn.GetSessionByTEID(msg.TEID())
	if err != nil {
		return err
	}
	s5cSession, err := s5cConn.GetSessionByIMSI(s11Session.IMSI)
	if err != nil {
		return err
	}
	s1uBearer := s11Session.GetDefaultBearer()
	s5uBearer := s5cSession.GetDefaultBearer()

	mbReqFromMME := msg.(*messages.ModifyBearerRequest)
	if brCtxIE := mbReqFromMME.BearerContextsToBeModified; brCtxIE != nil {
		for _, ie := range brCtxIE.ChildIEs {
			switch ie.Type {
			case ies.Indication:
				// do nothing in this example.
				// S-GW should change its beahavior based on indication flags like;
				//  - pass Modify Bearer Request to P-GW if handover is indicated.
				//  - XXX...
			case ies.FullyQualifiedTEID:
				if err := handleFTEIDU(ie, s11Session, s1uBearer); err != nil {
					return err
				}
			}
		}
	}

	s11mmeTEID, err := s11Session.GetTEID(v2.IFTypeS11MMEGTPC)
	if err != nil {
		return err
	}
	s1usgwTEID, err := s11Session.GetTEID(v2.IFTypeS1USGWGTPU)
	if err != nil {
		return err
	}
	s5usgwTEID, err := s5cSession.GetTEID(v2.IFTypeS5S8SGWGTPU)
	if err != nil {
		return err
	}
	s1uConn.RelayTo(s5uConn, s1usgwTEID, s5uBearer.OutgoingTEID(), s5uBearer.RemoteAddress())
	s5uConn.RelayTo(s1uConn, s5usgwTEID, s1uBearer.OutgoingTEID(), s1uBearer.RemoteAddress())

	log.Printf("S1=>S5: %s, %08x, %08x", s5uBearer.RemoteAddress(), s1usgwTEID, s5uBearer.OutgoingTEID())
	log.Printf("S5=>S1: %s, %08x, %08x", s1uBearer.RemoteAddress(), s5usgwTEID, s1uBearer.OutgoingTEID())

	s1uIP, _, err := net.SplitHostPort(*s1u)
	if err != nil {
		return err
	}
	mbRspFromSGW := messages.NewModifyBearerResponse(
		s11mmeTEID, 0,
		ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
		ies.NewBearerContext(
			ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
			ies.NewEPSBearerID(s1uBearer.EBI),
			ies.NewFullyQualifiedTEID(v2.IFTypeS1USGWGTPU, s1usgwTEID, s1uIP, ""),
		),
	)

	if err := s11Conn.RespondTo(mmeAddr, msg, mbRspFromSGW); err != nil {
		return err
	}

	loggerCh <- fmt.Sprintf(
		"Started listening on U-Plane for Subscriber: %s;\n\tS1-U: %s\n\tS5-U: %s",
		s11Session.IMSI, *s1u, *s5u,
	)
	return nil
}

func handleDeleteSessionRequest(s11Conn *v2.Conn, mmeAddr net.Addr, msg messages.Message) error {
	loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), mmeAddr)

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	s11Session, err := s11Conn.GetSessionByTEID(msg.TEID())
	if err != nil {
		return err
	}

	s5Session, err := s5cConn.GetSessionByIMSI(s11Session.IMSI)
	if err != nil {
		return err
	}

	s5cpgwTEID, err := s5Session.GetTEID(v2.IFTypeS5S8PGWGTPC)
	if err != nil {
		return err
	}

	if err := s5cConn.DeleteSession(
		s5cpgwTEID,
		ies.NewEPSBearerID(s5Session.GetDefaultBearer().EBI),
	); err != nil {
		return err
	}

	// wait for response from P-GW for 5 seconds
	doneCh := make(chan struct{})
	failCh := make(chan error)
	go func(delCh chan struct{}) {
		select {
		case <-delCh:
			break
		case <-time.After(5 * time.Second):
			loggerCh <- "No Delete Session Response from P-GW, purging session with MME anyway"
		}
		// respond to MME with DeleteSessionResponse.
		s11mmeTEID, err := s11Session.GetTEID(v2.IFTypeS11MMEGTPC)
		if err != nil {
			failCh <- err
			return
		}
		dsRspFromSGW := messages.NewDeleteSessionResponse(s11mmeTEID, s11Session.Sequence)
		if err := s11Conn.RespondTo(mmeAddr, msg, dsRspFromSGW); err != nil {
			failCh <- err
			return
		}

		loggerCh <- fmt.Sprintf("Session deleted with MME for Subscriber: %s", s11Session.IMSI)
		s11Conn.RemoveSession(s11Session)
		doneCh <- struct{}{}
	}(delCh)
	select {
	case <-doneCh:
		return nil
	case err := <-failCh:
		return err
	}
}

func handleFTEIDU(ie *ies.IE, session *v2.Session, bearer *v2.Bearer) error {
	if ie.Type != ies.FullyQualifiedTEID {
		return v2.ErrUnexpectedType
	}

	addr, err := net.ResolveUDPAddr("udp", ie.IPAddress()+":2152")
	if err != nil {
		return err
	}
	bearer.SetRemoteAddress(addr)
	bearer.SetOutgoingTEID(ie.TEID())

	session.AddTEID(ie.InterfaceType(), ie.TEID())
	return nil
}
