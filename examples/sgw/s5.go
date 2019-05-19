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
		// this is not such a fatal error worth stopping the whole program.
		sgw.loggerCh <- errors.Wrap(err, "Error").Error()
		return nil
	}

	s11Session, err := sgw.s11Conn.GetSessionByIMSI(s5Session.IMSI)
	if err != nil {
		return err
	}

	if err := v2.PassMessageTo(s11Session, msg, 5*time.Second); err != nil {
		return err
	}

	// even the cause indicates failure, session should be removed locally.
	sgw.loggerCh <- fmt.Sprintf("Session deleted with P-GW for Subscriber: %s", s5Session.IMSI)
	s5cConn.RemoveSession(s5Session)
	return nil
}
