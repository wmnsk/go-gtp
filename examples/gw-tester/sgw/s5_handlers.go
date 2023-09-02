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

func (s *sgw) handleCreateSessionResponse(s5cConn *gtpv2.Conn, pgwAddr net.Addr, msg message.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), pgwAddr)
	if s.mc != nil {
		s.mc.messagesReceived.WithLabelValues(pgwAddr.String(), msg.MessageTypeName()).Inc()
	}

	s5Session, err := s5cConn.GetSessionByTEID(msg.TEID(), pgwAddr)
	if err != nil {
		return err
	}

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	csRspFromPGW := msg.(*message.CreateSessionResponse)

	// check Cause value first.
	if causeIE := csRspFromPGW.Cause; causeIE != nil {
		cause, err := causeIE.Cause()
		if err != nil {
			return err
		}
		if cause != gtpv2.CauseRequestAccepted {
			s5cConn.RemoveSession(s5Session)
			// this is not such a fatal error worth stopping the whole program.
			// in the real case it is better to take some action based on the Cause, though.
			return &gtpv2.CauseNotOKError{
				MsgType: csRspFromPGW.MessageTypeName(),
				Cause:   cause,
				Msg:     fmt.Sprintf("subscriber: %s", s5Session.IMSI),
			}
		}
	} else {
		s5cConn.RemoveSession(s5Session)
		return &gtpv2.RequiredIEMissingError{
			Type: ie.Cause,
		}
	}

	bearer := s5Session.GetDefaultBearer()
	// retrieve values that P-GW gave.
	if paaIE := csRspFromPGW.PAA; paaIE != nil {
		ip, err := paaIE.IPAddress()
		if err != nil {
			return err
		}
		bearer.SubscriberIP = ip
	} else {
		s5cConn.RemoveSession(s5Session)
		return &gtpv2.RequiredIEMissingError{Type: ie.PDNAddressAllocation}
	}

	if fteidcIE := csRspFromPGW.PGWS5S8FTEIDC; fteidcIE != nil {
		it, err := fteidcIE.InterfaceType()
		if err != nil {
			return err
		}
		teid, err := fteidcIE.TEID()
		if err != nil {
			return err
		}
		s5Session.AddTEID(it, teid)
	} else {
		s5cConn.RemoveSession(s5Session)
		return &gtpv2.RequiredIEMissingError{Type: ie.FullyQualifiedTEID}
	}

	if brCtxIE := csRspFromPGW.BearerContextsCreated; brCtxIE != nil {
		for _, childIE := range brCtxIE[0].ChildIEs {
			switch childIE.Type {
			case ie.Cause:
				cause, err := childIE.Cause()
				if err != nil {
					return err
				}
				if cause != gtpv2.CauseRequestAccepted {
					s5cConn.RemoveSession(s5Session)
					return &gtpv2.CauseNotOKError{
						MsgType: csRspFromPGW.MessageTypeName(),
						Cause:   cause,
						Msg:     fmt.Sprintf("subscriber: %s", s5Session.IMSI),
					}
				}
			case ie.EPSBearerID:
				ebi, err := childIE.EPSBearerID()
				if err != nil {
					return err
				}
				bearer.EBI = ebi
			case ie.FullyQualifiedTEID:
				if err := s.handleFTEIDU(childIE, s5Session, bearer); err != nil {
					return err
				}
			case ie.ChargingID:
				cid, err := childIE.ChargingID()
				if err != nil {
					return err
				}
				bearer.ChargingID = cid
			}
		}
	} else {
		s5cConn.RemoveSession(s5Session)
		return &gtpv2.RequiredIEMissingError{Type: ie.BearerContext}
	}

	if err := s5Session.Activate(); err != nil {
		s5cConn.RemoveSession(s5Session)
		return err
	}

	s11Session, err := s.s11Conn.GetSessionByIMSI(s5Session.IMSI)
	if err != nil {
		return err
	}

	if err := gtpv2.PassMessageTo(s11Session, csRspFromPGW, 5*time.Second); err != nil {
		return err
	}

	return nil
}

func (s *sgw) handleDeleteSessionResponse(s5cConn *gtpv2.Conn, pgwAddr net.Addr, msg message.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), pgwAddr)
	if s.mc != nil {
		s.mc.messagesReceived.WithLabelValues(pgwAddr.String(), msg.MessageTypeName()).Inc()
	}

	s5Session, err := s5cConn.GetSessionByTEID(msg.TEID(), pgwAddr)
	if err != nil {
		return err
	}

	s11Session, err := s.s11Conn.GetSessionByIMSI(s5Session.IMSI)
	if err != nil {
		return err
	}

	if err := gtpv2.PassMessageTo(s11Session, msg, 5*time.Second); err != nil {
		return err
	}

	// even the cause indicates failure, session should be removed locally.
	log.Printf("Session deleted for Subscriber: %s", s5Session.IMSI)
	s5cConn.RemoveSession(s5Session)
	return nil
}

func (s *sgw) handleDeleteBearerRequest(s5cConn *gtpv2.Conn, pgwAddr net.Addr, msg message.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), pgwAddr)
	if s.mc != nil {
		s.mc.messagesReceived.WithLabelValues(pgwAddr.String(), msg.MessageTypeName()).Inc()
	}

	s5Session, err := s5cConn.GetSessionByTEID(msg.TEID(), pgwAddr)
	if err != nil {
		return err
	}

	s11Session, err := s.s11Conn.GetSessionByIMSI(s5Session.IMSI)
	if err != nil {
		return err
	}

	s5cpgwTEID, err := s5Session.GetTEID(gtpv2.IFTypeS5S8PGWGTPC)
	if err != nil {
		return err
	}

	s11mmeTEID, err := s11Session.GetTEID(gtpv2.IFTypeS11MMEGTPC)
	if err != nil {
		return err
	}

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	dbReqFromPGW := msg.(*message.DeleteBearerRequest)

	var dbRspFromSGW *message.DeleteBearerResponse
	var ebi *ie.IE

	if ie := dbReqFromPGW.LinkedEBI; ie != nil {
		ebi = ie
	}

	if e := dbReqFromPGW.EBIs; e != nil {
		ebiIE := e[0]
		// shouldn't be both.
		if ebi != nil {
			dbRspFromSGW = message.NewDeleteBearerResponse(
				s5cpgwTEID, 0,
				ie.NewCause(gtpv2.CauseContextNotFound, 0, 0, 0, ebiIE),
			)
			if err := s5cConn.RespondTo(pgwAddr, dbReqFromPGW, dbRspFromSGW); err != nil {
				return err
			}
			return fmt.Errorf("%T from %s had both Linked EBI and EBIs IE", dbReqFromPGW, pgwAddr)
		}
		ebi = ebiIE
	}

	if ebi == nil {
		dbRspFromSGW = message.NewDeleteBearerResponse(
			s5cpgwTEID, 0, ie.NewCause(gtpv2.CauseMandatoryIEMissing,
				0, 0, 0, ie.NewEPSBearerID(0),
			),
		)
		if err := s5cConn.RespondTo(pgwAddr, dbReqFromPGW, dbRspFromSGW); err != nil {
			return err
		}
		if s.mc != nil {
			s.mc.messagesSent.WithLabelValues(pgwAddr.String(), dbRspFromSGW.MessageTypeName()).Inc()
		}
		return err
	}

	// check if bearer associated with EBI exists or not.
	_, err = s5Session.LookupBearerByEBI(ebi.MustEPSBearerID())
	if err != nil {
		dbRspFromSGW = message.NewDeleteBearerResponse(
			s5cpgwTEID, 0,
			ie.NewCause(gtpv2.CauseContextNotFound, 0, 0, 0, nil),
		)
		if err := s5cConn.RespondTo(pgwAddr, dbReqFromPGW, dbRspFromSGW); err != nil {
			return err
		}
		if s.mc != nil {
			s.mc.messagesSent.WithLabelValues(pgwAddr.String(), dbRspFromSGW.MessageTypeName()).Inc()
		}
		return err
	}

	// forward to MME
	seq, err := s.s11Conn.DeleteBearer(s11mmeTEID, s11Session, ebi)
	if err != nil {
		return err
	}

	// wait for response from MME for 5 seconds
	incomingMessage, err := s5Session.WaitMessage(seq, 5*time.Second)
	if err != nil {
		dbRspFromSGW = message.NewDeleteBearerResponse(
			s5cpgwTEID, 0,
			ie.NewCause(gtpv2.CauseNoResourcesAvailable, 0, 0, 0, nil),
		)

		if err := s5cConn.RespondTo(pgwAddr, dbReqFromPGW, dbRspFromSGW); err != nil {
			return err
		}
		if s.mc != nil {
			s.mc.messagesSent.WithLabelValues(pgwAddr.String(), dbRspFromSGW.MessageTypeName()).Inc()
		}

		// remove anyway, as P-GW no longer keeps bearer locally
		s5Session.RemoveBearerByEBI(ebi.MustEPSBearerID())
		s11Session.RemoveBearerByEBI(ebi.MustEPSBearerID())
		return err
	}

	switch m := incomingMessage.(type) {
	case *message.DeleteBearerResponse:
		// move forward
		dbRspFromSGW = m
	default:
		return &gtpv2.UnexpectedTypeError{Msg: incomingMessage}
	}

	dbRspFromSGW.SetTEID(s5cpgwTEID)
	if err := s5cConn.RespondTo(pgwAddr, msg, dbRspFromSGW); err != nil {
		return err
	}
	if s.mc != nil {
		s.mc.messagesSent.WithLabelValues(pgwAddr.String(), dbRspFromSGW.MessageTypeName()).Inc()
	}

	s5Session.RemoveBearerByEBI(ebi.MustEPSBearerID())
	s11Session.RemoveBearerByEBI(ebi.MustEPSBearerID())

	return nil
}
