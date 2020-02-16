// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"log"
	"net"
	"strings"

	"github.com/pkg/errors"
	v2 "github.com/wmnsk/go-gtp/v2"
	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
)

func (p *pgw) handleCreateSessionRequest(c *v2.Conn, sgwAddr net.Addr, msg messages.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), sgwAddr)
	if p.mc != nil {
		p.mc.messagesReceived.WithLabelValues(sgwAddr.String(), msg.MessageTypeName()).Inc()
	}

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	csReqFromSGW := msg.(*messages.CreateSessionRequest)

	// keep session information retrieved from the message.
	session := v2.NewSession(sgwAddr, &v2.Subscriber{Location: &v2.Location{}})
	bearer := session.GetDefaultBearer()
	var err error
	if ie := csReqFromSGW.IMSI; ie != nil {
		imsi, err := ie.IMSI()
		if err != nil {
			return err
		}
		session.IMSI = imsi

		// remove previous session for the same subscriber if exists.
		sess, err := c.GetSessionByIMSI(imsi)
		if err != nil {
			switch err.(type) {
			case *v2.UnknownIMSIError:
				// whole new session. just ignore.
			default:
				return errors.Wrap(err, "got something unexpected")
			}
		} else {
			c.RemoveSession(sess)
		}
	} else {
		return &v2.RequiredIEMissingError{Type: ies.IMSI}
	}
	if ie := csReqFromSGW.MSISDN; ie != nil {
		session.MSISDN, err = ie.MSISDN()
		if err != nil {
			return err
		}
	} else {
		return &v2.RequiredIEMissingError{Type: ies.MSISDN}
	}
	if ie := csReqFromSGW.MEI; ie != nil {
		session.IMEI, err = ie.MobileEquipmentIdentity()
		if err != nil {
			return err
		}
	} else {
		return &v2.RequiredIEMissingError{Type: ies.MobileEquipmentIdentity}
	}
	if ie := csReqFromSGW.APN; ie != nil {
		bearer.APN, err = ie.AccessPointName()
		if err != nil {
			return err
		}
	} else {
		return &v2.RequiredIEMissingError{Type: ies.AccessPointName}
	}
	if ie := csReqFromSGW.ServingNetwork; ie != nil {
		session.MCC, err = ie.MCC()
		if err != nil {
			return err
		}
		session.MNC, err = ie.MNC()
		if err != nil {
			return err
		}
	} else {
		return &v2.RequiredIEMissingError{Type: ies.ServingNetwork}
	}
	if ie := csReqFromSGW.RATType; ie != nil {
		session.RATType, err = ie.RATType()
		if err != nil {
			return err
		}
	} else {
		return &v2.RequiredIEMissingError{Type: ies.RATType}
	}
	if ie := csReqFromSGW.SenderFTEIDC; ie != nil {
		teid, err := ie.TEID()
		if err != nil {
			return err
		}
		session.AddTEID(v2.IFTypeS5S8SGWGTPC, teid)
	} else {
		return &v2.RequiredIEMissingError{Type: ies.FullyQualifiedTEID}
	}

	var s5sgwuIP string
	var oteiU uint32
	if brCtxIE := csReqFromSGW.BearerContextsToBeCreated; brCtxIE != nil {
		for _, ie := range brCtxIE.ChildIEs {
			switch ie.Type {
			case ies.EPSBearerID:
				bearer.EBI, err = ie.EPSBearerID()
				if err != nil {
					return err
				}
			case ies.FullyQualifiedTEID:
				it, err := ie.InterfaceType()
				if err != nil {
					return err
				}
				oteiU, err = ie.TEID()
				if err != nil {
					return err
				}
				session.AddTEID(it, oteiU)

				s5sgwuIP, err = ie.IPAddress()
				if err != nil {
					return err
				}
			}
		}
	} else {
		return &v2.RequiredIEMissingError{Type: ies.BearerContext}
	}

	if ie := csReqFromSGW.PAA; ie != nil {
		bearer.SubscriberIP, err = ie.IPAddress()
		if err != nil {
			return err
		}
	} else {
		return &v2.RequiredIEMissingError{Type: ies.PDNAddressAllocation}
	}

	cIP := strings.Split(c.LocalAddr().String(), ":")[0]
	uIP := strings.Split(p.s5u, ":")[0]
	s5cFTEID := c.NewFTEID(v2.IFTypeS5S8PGWGTPC, cIP, "").WithInstance(1)
	s5uFTEID := c.NewFTEID(v2.IFTypeS5S8PGWGTPU, uIP, "").WithInstance(2)
	s5sgwTEID, err := session.GetTEID(v2.IFTypeS5S8SGWGTPC)
	if err != nil {
		return err
	}
	csRspFromPGW := messages.NewCreateSessionResponse(
		s5sgwTEID, 0,
		ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
		s5cFTEID,
		ies.NewPDNAddressAllocation(bearer.SubscriberIP),
		ies.NewAPNRestriction(v2.APNRestrictionPublic2),
		ies.NewBearerContext(
			ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
			ies.NewEPSBearerID(bearer.EBI),
			s5uFTEID,
			ies.NewChargingID(bearer.ChargingID),
		),
	)
	if csReqFromSGW.SGWFQCSID != nil {
		csRspFromPGW.PGWFQCSID = ies.NewFullyQualifiedCSID(cIP, 1)
	}
	session.AddTEID(v2.IFTypeS5S8PGWGTPC, s5cFTEID.MustTEID())
	session.AddTEID(v2.IFTypeS5S8PGWGTPU, s5uFTEID.MustTEID())

	if err := c.RespondTo(sgwAddr, csReqFromSGW, csRspFromPGW); err != nil {
		return err
	}
	if p.mc != nil {
		p.mc.messagesSent.WithLabelValues(sgwAddr.String(), csRspFromPGW.MessageTypeName()).Inc()
	}

	s5pgwTEID, err := session.GetTEID(v2.IFTypeS5S8PGWGTPC)
	if err != nil {
		return err
	}

	// don't forget to activate and add session created to the session list
	if err := session.Activate(); err != nil {
		return err
	}
	c.RegisterSession(s5pgwTEID, session)

	if err := p.setupUPlane(net.ParseIP(s5sgwuIP), net.ParseIP(bearer.SubscriberIP), oteiU, s5uFTEID.MustTEID()); err != nil {
		return err
	}

	log.Printf("Session created with S-GW for subscriber: %s;\n\tS5C S-GW: %s, TEID->: %#x, TEID<-: %#x",
		session.Subscriber.IMSI, sgwAddr, s5sgwTEID, s5pgwTEID,
	)
	return nil
}

func (p *pgw) handleDeleteSessionRequest(c *v2.Conn, sgwAddr net.Addr, msg messages.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), sgwAddr)
	if p.mc != nil {
		p.mc.messagesReceived.WithLabelValues(sgwAddr.String(), msg.MessageTypeName()).Inc()
	}

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	session, err := c.GetSessionByTEID(msg.TEID(), sgwAddr)
	if err != nil {
		dsr := messages.NewDeleteSessionResponse(
			0, 0,
			ies.NewCause(v2.CauseIMSIIMEINotKnown, 0, 0, 0, nil),
		)
		if err := c.RespondTo(sgwAddr, msg, dsr); err != nil {
			return err
		}

		return err
	}

	// respond to S-GW with DeleteSessionResponse.
	teid, err := session.GetTEID(v2.IFTypeS5S8SGWGTPC)
	if err != nil {
		log.Println(err)
		return nil
	}
	dsr := messages.NewDeleteSessionResponse(
		teid, 0,
		ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
	)
	if err := c.RespondTo(sgwAddr, msg, dsr); err != nil {
		return err
	}
	if p.mc != nil {
		p.mc.messagesSent.WithLabelValues(sgwAddr.String(), dsr.MessageTypeName()).Inc()
	}

	log.Printf("Session deleted for Subscriber: %s", session.IMSI)
	c.RemoveSession(session)
	return nil
}
