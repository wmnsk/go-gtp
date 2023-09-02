// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
)

func (p *pgw) handleCreateSessionRequest(c *gtpv2.Conn, sgwAddr net.Addr, msg message.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), sgwAddr)
	if p.mc != nil {
		p.mc.messagesReceived.WithLabelValues(sgwAddr.String(), msg.MessageTypeName()).Inc()
	}

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	csReqFromSGW := msg.(*message.CreateSessionRequest)

	// keep session information retrieved from the message.
	session := gtpv2.NewSession(sgwAddr, &gtpv2.Subscriber{Location: &gtpv2.Location{}})
	bearer := session.GetDefaultBearer()
	var err error
	if imsiIE := csReqFromSGW.IMSI; imsiIE != nil {
		imsi, err := imsiIE.IMSI()
		if err != nil {
			return err
		}
		session.IMSI = imsi

		// remove previous session for the same subscriber if exists.
		sess, err := c.GetSessionByIMSI(imsi)
		if err != nil {
			switch err.(type) {
			case *gtpv2.UnknownIMSIError:
				// whole new session. just ignore.
			default:
				return fmt.Errorf("got something unexpected: %w", err)
			}
		} else {
			c.RemoveSession(sess)
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.IMSI}
	}
	if msisdnIE := csReqFromSGW.MSISDN; msisdnIE != nil {
		session.MSISDN, err = msisdnIE.MSISDN()
		if err != nil {
			return err
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.MSISDN}
	}
	if meiIE := csReqFromSGW.MEI; meiIE != nil {
		session.IMEI, err = meiIE.MobileEquipmentIdentity()
		if err != nil {
			return err
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.MobileEquipmentIdentity}
	}
	if apnIE := csReqFromSGW.APN; apnIE != nil {
		bearer.APN, err = apnIE.AccessPointName()
		if err != nil {
			return err
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.AccessPointName}
	}
	if netIE := csReqFromSGW.ServingNetwork; netIE != nil {
		session.MCC, err = netIE.MCC()
		if err != nil {
			return err
		}
		session.MNC, err = netIE.MNC()
		if err != nil {
			return err
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.ServingNetwork}
	}
	if ratIE := csReqFromSGW.RATType; ratIE != nil {
		session.RATType, err = ratIE.RATType()
		if err != nil {
			return err
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.RATType}
	}
	if fteidcIE := csReqFromSGW.SenderFTEIDC; fteidcIE != nil {
		teid, err := fteidcIE.TEID()
		if err != nil {
			return err
		}
		session.AddTEID(gtpv2.IFTypeS5S8SGWGTPC, teid)
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.FullyQualifiedTEID}
	}

	var s5sgwuIP string
	var oteiU uint32
	if brCtxIE := csReqFromSGW.BearerContextsToBeCreated; brCtxIE != nil {
		for _, childIE := range brCtxIE[0].ChildIEs {
			switch childIE.Type {
			case ie.EPSBearerID:
				bearer.EBI, err = childIE.EPSBearerID()
				if err != nil {
					return err
				}
			case ie.FullyQualifiedTEID:
				it, err := childIE.InterfaceType()
				if err != nil {
					return err
				}
				oteiU, err = childIE.TEID()
				if err != nil {
					return err
				}
				session.AddTEID(it, oteiU)

				s5sgwuIP, err = childIE.IPAddress()
				if err != nil {
					return err
				}
			}
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.BearerContext}
	}

	if paaIE := csReqFromSGW.PAA; paaIE != nil {
		bearer.SubscriberIP, err = paaIE.IPAddress()
		if err != nil {
			return err
		}
	} else {
		return &gtpv2.RequiredIEMissingError{Type: ie.PDNAddressAllocation}
	}

	cIP := strings.Split(c.LocalAddr().String(), ":")[0]
	uIP := strings.Split(p.s5u, ":")[0]
	s5cFTEID := c.NewSenderFTEID(cIP, "").WithInstance(1)
	s5uFTEID := p.uConn.NewFTEID(gtpv2.IFTypeS5S8PGWGTPU, uIP, "").WithInstance(2)
	s5sgwTEID, err := session.GetTEID(gtpv2.IFTypeS5S8SGWGTPC)
	if err != nil {
		return err
	}
	csRspFromPGW := message.NewCreateSessionResponse(
		s5sgwTEID, 0,
		ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
		s5cFTEID,
		ie.NewPDNAddressAllocation(bearer.SubscriberIP),
		ie.NewAPNRestriction(gtpv2.APNRestrictionPublic2),
		ie.NewBearerContext(
			ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
			ie.NewEPSBearerID(bearer.EBI),
			s5uFTEID,
			ie.NewChargingID(bearer.ChargingID),
		),
	)
	if csReqFromSGW.SGWFQCSID != nil {
		csRspFromPGW.PGWFQCSID = ie.NewFullyQualifiedCSID(cIP, 1)
	}
	session.AddTEID(gtpv2.IFTypeS5S8PGWGTPC, s5cFTEID.MustTEID())
	session.AddTEID(gtpv2.IFTypeS5S8PGWGTPU, s5uFTEID.MustTEID())

	if err := c.RespondTo(sgwAddr, csReqFromSGW, csRspFromPGW); err != nil {
		return err
	}
	if p.mc != nil {
		p.mc.messagesSent.WithLabelValues(sgwAddr.String(), csRspFromPGW.MessageTypeName()).Inc()
	}

	s5pgwTEID, err := session.GetTEID(gtpv2.IFTypeS5S8PGWGTPC)
	if err != nil {
		return err
	}

	// don't forget to activate and add session created to the session list
	if err := session.Activate(); err != nil {
		return err
	}
	c.RegisterSession(s5pgwTEID, session)

	if p.useKernelGTP {
		if err := p.setupUPlane(net.ParseIP(s5sgwuIP), net.ParseIP(bearer.SubscriberIP), oteiU, s5uFTEID.MustTEID()); err != nil {
			return err
		}
	}

	log.Printf("Session created with S-GW for subscriber: %s;\n\tS5C S-GW: %s, TEID->: %#x, TEID<-: %#x",
		session.Subscriber.IMSI, sgwAddr, s5sgwTEID, s5pgwTEID,
	)
	return nil
}

func (p *pgw) handleDeleteSessionRequest(c *gtpv2.Conn, sgwAddr net.Addr, msg message.Message) error {
	log.Printf("Received %s from %s", msg.MessageTypeName(), sgwAddr)
	if p.mc != nil {
		p.mc.messagesReceived.WithLabelValues(sgwAddr.String(), msg.MessageTypeName()).Inc()
	}

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	session, err := c.GetSessionByTEID(msg.TEID(), sgwAddr)
	if err != nil {
		dsr := message.NewDeleteSessionResponse(
			0, 0,
			ie.NewCause(gtpv2.CauseIMSIIMEINotKnown, 0, 0, 0, nil),
		)
		if err := c.RespondTo(sgwAddr, msg, dsr); err != nil {
			return err
		}

		return err
	}

	// respond to S-GW with DeleteSessionResponse.
	teid, err := session.GetTEID(gtpv2.IFTypeS5S8SGWGTPC)
	if err != nil {
		log.Println(err)
		return nil
	}
	dsr := message.NewDeleteSessionResponse(
		teid, 0,
		ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
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
