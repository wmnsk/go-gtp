// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
	"log"
	"net"

	v2 "github.com/wmnsk/go-gtp/gtpv2"
)

func (m *mme) handleCreateSessionResponse(c *v2.Conn, sgwAddr net.Addr, msg message.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), sgwAddr)
	if m.mc != nil {
		m.mc.messagesReceived.WithLabelValues(sgwAddr.String(), msg.MessageTypeName()).Inc()
	}

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
	csRspFromSGW := msg.(*message.CreateSessionResponse)

	// check Cause value first.
	if causeIE := csRspFromSGW.Cause; causeIE != nil {
		cause, err := causeIE.Cause()
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

	if paaIE := csRspFromSGW.PAA; paaIE != nil {
		bearer.SubscriberIP, err = paaIE.IPAddress()
		if err != nil {
			return err
		}
	}
	if fteidcIE := csRspFromSGW.SenderFTEIDC; fteidcIE != nil {
		teid, err := fteidcIE.TEID()
		if err != nil {
			return err
		}
		session.AddTEID(v2.IFTypeS11S4SGWGTPC, teid)
	} else {
		return &v2.RequiredIEMissingError{Type: ie.FullyQualifiedTEID}
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
		for _, childIE := range brCtxIE.ChildIEs {
			switch childIE.Type {
			case ie.EPSBearerID:
				bearer.EBI, err = childIE.EPSBearerID()
				if err != nil {
					return err
				}
			case ie.FullyQualifiedTEID:
				if childIE.Instance() != 0 {
					continue
				}
				it, err := childIE.InterfaceType()
				if err != nil {
					return err
				}
				teid, err := childIE.TEID()
				if err != nil {
					return err
				}
				session.AddTEID(it, teid)
			}
		}
	} else {
		return &v2.RequiredIEMissingError{Type: ie.BearerContext}
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

func (m *mme) handleModifyBearerResponse(c *v2.Conn, sgwAddr net.Addr, msg message.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), sgwAddr)
	if m.mc != nil {
		m.mc.messagesReceived.WithLabelValues(sgwAddr.String(), msg.MessageTypeName()).Inc()
	}

	session, err := c.GetSessionByTEID(msg.TEID(), sgwAddr)
	if err != nil {
		return err
	}

	mbRspFromSGW := msg.(*message.ModifyBearerResponse)
	if causeIE := mbRspFromSGW.Cause; causeIE != nil {
		cause, err := causeIE.Cause()
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
		return &v2.RequiredIEMissingError{Type: ie.Cause}
	}

	if brCtxIE := mbRspFromSGW.BearerContextsModified; brCtxIE != nil {
		for _, childIE := range brCtxIE.ChildIEs {
			switch childIE.Type {
			case ie.FullyQualifiedTEID:
				if childIE.Instance() != 0 {
					continue
				}
				it, err := childIE.InterfaceType()
				if err != nil {
					return err
				}
				teid, err := childIE.TEID()
				if err != nil {
					return err
				}
				session.AddTEID(it, teid)

				m.sgw.s1uIP, err = childIE.IPAddress()
				if err != nil {
					return err
				}
			}
		}
	} else {
		return &v2.RequiredIEMissingError{Type: ie.BearerContext}
	}

	log.Printf("Bearer modified with S-GW for Subscriber: %s", session.IMSI)
	m.modified <- struct{}{}
	return nil
}

func (m *mme) handleDeleteSessionResponse(c *v2.Conn, sgwAddr net.Addr, msg message.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), sgwAddr)
	if m.mc != nil {
		m.mc.messagesReceived.WithLabelValues(sgwAddr.String(), msg.MessageTypeName()).Inc()
	}

	session, err := c.GetSessionByTEID(msg.TEID(), sgwAddr)
	if err != nil {
		return err
	}

	c.RemoveSession(session)
	log.Printf("Session deleted with S-GW for Subscriber: %s", session.IMSI)
	return nil
}
