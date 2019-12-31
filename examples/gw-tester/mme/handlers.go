// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net"

	v2 "github.com/wmnsk/go-gtp/v2"
	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
)

func (m *mme) handleCreateSessionResponse(c *v2.Conn, sgwAddr net.Addr, msg messages.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), sgwAddr)

	// find the session associated with TEID
	session, err := c.GetSessionByTEID(msg.TEID(), sgwAddr)
	if err != nil {
		c.RemoveSession(session)
		return err
	}
	bearer := session.GetDefaultBearer()

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	csRspFromSGW := msg.(*messages.CreateSessionResponse)

	// check Cause value first.
	if ie := csRspFromSGW.Cause; ie != nil {
		cause, err := ie.Cause()
		if err != nil {
			return err
		}
		if cause != v2.CauseRequestAccepted {
			c.RemoveSession(session)
			return &v2.CauseNotOKError{
				MsgType: csRspFromSGW.MessageTypeName(),
				Cause:   cause,
				Msg:     fmt.Sprintf("subscriber: %s", session.IMSI),
			}
		}
	} else {
		return &v2.RequiredIEMissingError{Type: msg.MessageType()}
	}

	if ie := csRspFromSGW.PAA; ie != nil {
		bearer.SubscriberIP, err = ie.IPAddress()
		if err != nil {
			return err
		}
	}
	if ie := csRspFromSGW.SenderFTEIDC; ie != nil {
		teid, err := ie.TEID()
		if err != nil {
			return err
		}
		session.AddTEID(v2.IFTypeS11S4SGWGTPC, teid)
	} else {
		return &v2.RequiredIEMissingError{Type: ies.FullyQualifiedTEID}
	}

	s11sgwTEID, err := session.GetTEID(v2.IFTypeS11S4SGWGTPC)
	if err != nil {
		c.RemoveSession(session)
		return err
	}
	s11mmeTEID, err := session.GetTEID(v2.IFTypeS11MMEGTPC)
	if err != nil {
		c.RemoveSession(session)
		return err
	}

	if brCtxIE := csRspFromSGW.BearerContextsCreated; brCtxIE != nil {
		for _, ie := range brCtxIE.ChildIEs {
			switch ie.Type {
			case ies.EPSBearerID:
				bearer.EBI, err = ie.EPSBearerID()
				if err != nil {
					return err
				}
			case ies.FullyQualifiedTEID:
				if ie.Instance() != 0 {
					continue
				}
				it, err := ie.InterfaceType()
				if err != nil {
					return err
				}
				teid, err := ie.TEID()
				if err != nil {
					return err
				}
				session.AddTEID(it, teid)
			}
		}
	} else {
		return &v2.RequiredIEMissingError{Type: ies.BearerContext}
	}

	if err := session.Activate(); err != nil {
		c.RemoveSession(session)
		return err
	}

	log.Printf(
		"Session created with S-GW for Subscriber: %s;\n\tS11 S-GW: %s, TEID->: %#x, TEID<-: %#x",
		session.Subscriber.IMSI, sgwAddr, s11sgwTEID, s11mmeTEID,
	)
	m.created <- struct{}{}
	return nil
}

func (m *mme) handleModifyBearerResponse(c *v2.Conn, sgwAddr net.Addr, msg messages.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), sgwAddr)

	session, err := c.GetSessionByTEID(msg.TEID(), sgwAddr)
	if err != nil {
		return err
	}

	mbRspFromSGW := msg.(*messages.ModifyBearerResponse)
	if ie := mbRspFromSGW.Cause; ie != nil {
		cause, err := ie.Cause()
		if err != nil {
			return err
		}
		if cause != v2.CauseRequestAccepted {
			return &v2.CauseNotOKError{
				MsgType: msg.MessageTypeName(),
				Cause:   cause,
				Msg:     fmt.Sprintf("subscriber: %s", session.IMSI),
			}
		}
	} else {
		return &v2.RequiredIEMissingError{Type: ies.Cause}
	}

	if brCtxIE := mbRspFromSGW.BearerContextsModified; brCtxIE != nil {
		for _, ie := range brCtxIE.ChildIEs {
			switch ie.Type {
			case ies.FullyQualifiedTEID:
				if ie.Instance() != 0 {
					continue
				}
				it, err := ie.InterfaceType()
				if err != nil {
					return err
				}
				teid, err := ie.TEID()
				if err != nil {
					return err
				}
				session.AddTEID(it, teid)

				m.sgw.s1uIP, err = ie.IPAddress()
				if err != nil {
					return err
				}
			}
		}
	} else {
		return &v2.RequiredIEMissingError{Type: ies.BearerContext}
	}

	log.Printf("Bearer modified with S-GW for Subscriber: %s", session.IMSI)
	m.modified <- struct{}{}
	return nil
}

func (m *mme) handleDeleteSessionResponse(c *v2.Conn, sgwAddr net.Addr, msg messages.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), sgwAddr)

	session, err := c.GetSessionByTEID(msg.TEID(), sgwAddr)
	if err != nil {
		return err
	}

	c.RemoveSession(session)
	log.Printf("Session deleted with S-GW for Subscriber: %s", session.IMSI)
	return nil
}
