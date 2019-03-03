// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"net"
	"strings"

	v1 "github.com/wmnsk/go-gtp/v1"

	"github.com/pkg/errors"
	v2 "github.com/wmnsk/go-gtp/v2"
	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
)

// getSubscriberIP is to get IP address to be assigned to the subscriber.
//
// In the real case, P-GW may ask AAA and PCRF retrieve required information for subscriber,
// but here, to keep the example simple, this just returns subscriber's IP address defined in
// the map "subIPMap".
func getSubscriberIP(sub *v2.Subscriber) (string, error) {
	subIPMap := map[string]string{
		"123451234567891": "10.10.10.1",
		"123451234567892": "10.10.10.2",
		"123451234567893": "10.10.10.3",
		"123451234567894": "10.10.10.4",
		"123451234567895": "10.10.10.5",
	}

	if ip, ok := subIPMap[sub.IMSI]; ok {
		return ip, nil
	}
	return "", fmt.Errorf("Subscriber %s not found", sub.IMSI)
}

var (
	loggerCh = make(chan string)
	errCh    = make(chan error)

	uConn *v1.UPlaneConn
)

func handleCreateSessionRequest(c *v2.Conn, sgwAddr net.Addr, msg messages.Message) error {
	loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), sgwAddr)

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	csReqFromSGW := msg.(*messages.CreateSessionRequest)

	// keep session information retrieved from the message.
	session := v2.NewSession(sgwAddr, &v2.Subscriber{Location: &v2.Location{}})
	bearer := session.GetDefaultBearer()
	if ie := csReqFromSGW.IMSI; ie != nil {
		session.IMSI = ie.IMSI()
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.IMSI}
	}
	if ie := csReqFromSGW.MSISDN; ie != nil {
		session.MSISDN = ie.MSISDN()
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.MSISDN}
	}
	if ie := csReqFromSGW.MEI; ie != nil {
		session.IMEI = ie.MobileEquipmentIdentity()
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.MobileEquipmentIdentity}
	}
	if ie := csReqFromSGW.APN; ie != nil {
		bearer.APN = ie.AccessPointName()
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.AccessPointName}
	}
	if ie := csReqFromSGW.ServingNetwork; ie != nil {
		session.MCC = ie.MCC()
		session.MNC = ie.MNC()
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.ServingNetwork}
	}
	if ie := csReqFromSGW.RATType; ie != nil {
		session.RATType = ie.RATType()
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.RATType}
	}
	if ie := csReqFromSGW.SenderFTEIDC; ie != nil {
		session.AddTEID(v2.IFTypeS5S8SGWGTPC, ie.TEID())
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.FullyQualifiedTEID}
	}

	var teidOut uint32
	if brCtxIE := csReqFromSGW.BearerContextsToBeCreated; brCtxIE != nil {
		for _, ie := range brCtxIE.ChildIEs {
			switch ie.Type {
			case ies.EPSBearerID:
				bearer.EBI = ie.EPSBearerID()
			case ies.FullyQualifiedTEID:
				session.AddTEID(ie.InterfaceType(), ie.TEID())
				teidOut = ie.TEID()
			}
		}
	} else {
		return &v2.ErrRequiredIEMissing{Type: ies.BearerContext}
	}

	var err error
	bearer.SubscriberIP, err = getSubscriberIP(session.Subscriber)
	if err != nil {
		return err
	}

	cIP := strings.Split(c.LocalAddr().String(), ":")[0]
	uIP := strings.Split(*s5u, ":")[0]
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
	session.AddTEID(v2.IFTypeS5S8PGWGTPC, s5cFTEID.TEID())
	session.AddTEID(v2.IFTypeS5S8PGWGTPU, s5uFTEID.TEID())

	if err := c.RespondTo(sgwAddr, csReqFromSGW, csRspFromPGW); err != nil {
		return err
	}

	s5pgwTEID, err := session.GetTEID(v2.IFTypeS5S8PGWGTPC)
	if err != nil {
		return err
	}

	// don't forget to activate and add session created to the session list
	if err := session.Activate(); err != nil {
		return err
	}
	c.AddSession(session)

	if uConn == nil {
		laddr, err := net.ResolveUDPAddr("udp", *s5u)
		if err != nil {
			return err
		}
		uConn, err = v1.ListenAndServeUPlane(laddr, 0, errCh)
		if err != nil {
			return err
		}
	}
	loggerCh <- fmt.Sprintf("Started listening on %s", uConn.LocalAddr())

	go func() {
		buf := make([]byte, 1500)
		for {
			n, raddr, _, err := uConn.ReadFromGTP(buf)
			if err != nil {
				return
			}

			rsp := make([]byte, n)
			// swap IP
			copy(rsp[12:16], buf[16:20])
			copy(rsp[16:20], buf[12:16])
			// update message type and checksum
			copy(rsp, buf[:n])
			rsp[20] = 0
			rsp[22] = 0x9b

			if _, err := uConn.WriteToGTP(teidOut, rsp, raddr); err != nil {
				return
			}
		}
	}()

	loggerCh <- fmt.Sprintf("Session created with S-GW for subscriber: %s;\n\tS5C S-GW: %s, TEID->: %#x, TEID<-: %#x",
		session.Subscriber.IMSI, sgwAddr, s5sgwTEID, s5pgwTEID,
	)
	return nil
}

func handleDeleteSessionRequest(c *v2.Conn, sgwAddr net.Addr, msg messages.Message) error {
	loggerCh <- fmt.Sprintf("Received %s from %s", msg.MessageTypeName(), sgwAddr)

	// assert type to refer to the struct field specific to the message.
	// in general, no need to check if it can be type-asserted, as long as the MessageType is
	// specified correctly in AddHandler().
	session, err := c.GetSessionByTEID(msg.TEID())
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
		loggerCh <- errors.Wrap(err, "Error").Error()
		return nil
	}
	dsr := messages.NewDeleteSessionResponse(
		teid, 0,
		ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
	)
	if err := c.RespondTo(sgwAddr, msg, dsr); err != nil {
		return err
	}

	loggerCh <- fmt.Sprintf("Session deleted for Subscriber: %s", session.IMSI)
	c.RemoveSession(session)
	return nil
}
