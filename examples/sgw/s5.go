package main

import (
	"fmt"
	"net"
	"time"

	"github.com/pkg/errors"
	v2 "github.com/wmnsk/go-gtp/v2"
	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
)

func handleCreateSessionResponse(s5cConn *v2.Conn, pgwAddr net.Addr, msg messages.Message) error {
	sgw.loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), pgwAddr)

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
		cause, err := ie.Cause()
		if err != nil {
			return err
		}
		if cause != v2.CauseRequestAccepted {
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
		ip, err := ie.IPAddress()
		if err != nil {
			return err
		}
		bearer.SubscriberIP = ip
	} else {
		s5cConn.RemoveSession(s5Session)
		return &v2.ErrRequiredIEMissing{Type: ies.PDNAddressAllocation}
	}
	if ie := csRspFromPGW.PGWS5S8FTEIDC; ie != nil {
		it, err := ie.InterfaceType()
		if err != nil {
			return err
		}
		teid, err := ie.TEID()
		if err != nil {
			return err
		}
		s5Session.AddTEID(it, teid)
	} else {
		s5cConn.RemoveSession(s5Session)
		return &v2.ErrRequiredIEMissing{Type: ies.FullyQualifiedTEID}
	}

	if brCtxIE := csRspFromPGW.BearerContextsCreated; brCtxIE != nil {
		for _, ie := range brCtxIE.ChildIEs {
			switch ie.Type {
			case ies.Cause:
				cause, err := ie.Cause()
				if err != nil {
					return err
				}
				if cause != v2.CauseRequestAccepted {
					s5cConn.RemoveSession(s5Session)
					return &v2.ErrCauseNotOK{
						MsgType: csRspFromPGW.MessageTypeName(),
						Cause:   cause,
						Msg:     fmt.Sprintf("subscriber: %s", s5Session.IMSI),
					}
				}
			case ies.EPSBearerID:
				ebi, err := ie.EPSBearerID()
				if err != nil {
					return err
				}
				bearer.EBI = ebi
			case ies.FullyQualifiedTEID:
				if err := handleFTEIDU(ie, s5Session, bearer); err != nil {
					return err
				}
			case ies.ChargingID:
				cid, err := ie.ChargingID()
				if err != nil {
					return err
				}
				bearer.ChargingID = cid
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

	s11Session, err := sgw.s11Conn.GetSessionByIMSI(s5Session.IMSI)
	if err != nil {
		return err
	}

	if err := v2.PassMessageTo(s11Session, csRspFromPGW, 5*time.Second); err != nil {
		return err
	}

	return nil
}

func handleDeleteSessionResponse(s5cConn *v2.Conn, pgwAddr net.Addr, msg messages.Message) error {
	sgw.loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), pgwAddr)

	s5Session, err := s5cConn.GetSessionByTEID(msg.TEID())
	if err != nil {
		return err
	}

	s11Session, err := sgw.s11Conn.GetSessionByIMSI(s5Session.IMSI)
	if err != nil {
		return err
	}

	if err := v2.PassMessageTo(s11Session, msg, 5*time.Second); err != nil {
		return err
	}

	// even the cause indicates failure, session should be removed locally.
	sgw.loggerCh <- fmt.Sprintf("Session deleted for Subscriber: %s", s5Session.IMSI)
	s5cConn.RemoveSession(s5Session)
	return nil
}

func handleDeleteBearerRequest(s5cConn *v2.Conn, pgwAddr net.Addr, msg messages.Message) error {
	sgw.loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), pgwAddr)

	s5Session, err := s5cConn.GetSessionByTEID(msg.TEID())
	if err != nil {
		return err
	}

	s11Session, err := sgw.s11Conn.GetSessionByIMSI(s5Session.IMSI)
	if err != nil {
		return err
	}

	s5cpgwTEID, err := s5Session.GetTEID(v2.IFTypeS5S8PGWGTPC)
	if err != nil {
		return err
	}

	s11mmeTEID, err := s11Session.GetTEID(v2.IFTypeS11MMEGTPC)
	if err != nil {
		return err
	}

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	dbReqFromPGW := msg.(*messages.DeleteBearerRequest)

	var dbRspFromSGW *messages.DeleteBearerResponse
	var ebi *ies.IE
	if ie := dbReqFromPGW.LinkedEBI; ie != nil {
		ebi = ie
	}
	if ie := dbReqFromPGW.EBI; ie != nil {
		// shouldn't be both.
		if ebi != nil {
			dbRspFromSGW = messages.NewDeleteBearerResponse(
				s5cpgwTEID, 0,
				ies.NewCause(v2.CauseContextNotFound, 0, 0, 0, ie),
			)
			if err := s5cConn.RespondTo(pgwAddr, dbReqFromPGW, dbRspFromSGW); err != nil {
				return err
			}
			return errors.Errorf(
				"%T from %s had both Linked EBI and EBIs IE",
				dbReqFromPGW, pgwAddr,
			)
		}
		ebi = ie
	}

	if ebi == nil {
		dbRspFromSGW = messages.NewDeleteBearerResponse(
			s5cpgwTEID, 0, ies.NewCause(v2.CauseMandatoryIEMissing,
				0, 0, 0, ies.NewEPSBearerID(0),
			),
		)
		if err := s5cConn.RespondTo(pgwAddr, dbReqFromPGW, dbRspFromSGW); err != nil {
			return err
		}
		return err
	}

	// check if bearer associated with EBI exists or not.
	_, err = s5Session.LookupBearerByEBI(ebi.MustEPSBearerID())
	if err != nil {
		dbRspFromSGW = messages.NewDeleteBearerResponse(
			s5cpgwTEID, 0,
			ies.NewCause(v2.CauseContextNotFound, 0, 0, 0, nil),
		)
		if err := s5cConn.RespondTo(pgwAddr, dbReqFromPGW, dbRspFromSGW); err != nil {
			return err
		}
		return err
	}

	// forward to MME
	seq, err := sgw.s11Conn.DeleteBearer(s11mmeTEID, ebi)
	if err != nil {
		return err
	}

	// wait for response from MME for 5 seconds
	message, err := s5Session.WaitMessage(seq, 5*time.Second)
	if err != nil {
		dbRspFromSGW = messages.NewDeleteBearerResponse(
			s5cpgwTEID, 0,
			ies.NewCause(v2.CauseNoResourcesAvailable, 0, 0, 0, nil),
		)

		if err := s5cConn.RespondTo(pgwAddr, dbReqFromPGW, dbRspFromSGW); err != nil {
			return err
		}
		// remove anyway, as P-GW no longer keeps bearer locally
		s5Session.RemoveBearerByEBI(ebi.MustEPSBearerID())
		s11Session.RemoveBearerByEBI(ebi.MustEPSBearerID())
		return err
	}

	switch m := message.(type) {
	case *messages.DeleteBearerResponse:
		// move forward
		dbRspFromSGW = m
	default:
		return &v2.ErrUnexpectedType{Msg: message}
	}

	dbRspFromSGW.SetTEID(s5cpgwTEID)
	if err := s5cConn.RespondTo(pgwAddr, msg, dbRspFromSGW); err != nil {
		return err
	}

	s5Session.RemoveBearerByEBI(ebi.MustEPSBearerID())
	s11Session.RemoveBearerByEBI(ebi.MustEPSBearerID())

	return nil
}
