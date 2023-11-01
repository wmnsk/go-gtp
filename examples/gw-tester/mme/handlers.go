// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net"

	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
)

func (m *mme) handleCreateSessionResponse(c *gtpv2.Conn, sgwAddr net.Addr, msg message.Message) error {
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
		if cause != gtpv2.CauseRequestAccepted {
			c.RemoveSession(session)
			return &gtpv2.CauseNotOKError{
				MsgType: csRspFromSGW.MessageTypeName(),
				Cause:   cause,
				Msg:     fmt.Sprintf("subscriber: %s", session.IMSI),
			}
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: msg.MessageType()}
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
		session.AddTEID(gtpv2.IFTypeS11S4SGWGTPC, teid)
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.FullyQualifiedTEID}
	}

	s11sgwTEID, err := session.GetTEID(gtpv2.IFTypeS11S4SGWGTPC)
	if err != nil {
		c.RemoveSession(session)
		return err
	}
	s11mmeTEID, err := session.GetTEID(gtpv2.IFTypeS11MMEGTPC)
	if err != nil {
		c.RemoveSession(session)
		return err
	}

	if brCtxIE := csRspFromSGW.BearerContextsCreated; brCtxIE != nil {
		for _, childIE := range brCtxIE[0].ChildIEs {
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
				
				m.sgw.s1uIP, err = childIE.IPAddress()
				if err != nil {
					return err
				}

			}
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.BearerContext}
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

func (m *mme) handleModifyBearerResponse(c *gtpv2.Conn, sgwAddr net.Addr, msg message.Message) error {
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
		if cause != gtpv2.CauseRequestAccepted {
			return &gtpv2.CauseNotOKError{
				MsgType: msg.MessageTypeName(),
				Cause:   cause,
				Msg:     fmt.Sprintf("subscriber: %s", session.IMSI),
			}
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.Cause}
	}

	if brCtxIE := mbRspFromSGW.BearerContextsModified; brCtxIE != nil {
		for _, childIE := range brCtxIE[0].ChildIEs {
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

				// m.sgw.s1uIP, err = childIE.IPAddress()
				// if err != nil {
				// 	return err
				// }
			}
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.BearerContext}
	}

	log.Printf("Bearer modified with S-GW for Subscriber: %s", session.IMSI)
	m.modified <- struct{}{}
	return nil
}

func (m *mme) handleDeleteSessionResponse(c *gtpv2.Conn, sgwAddr net.Addr, msg message.Message) error {
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
