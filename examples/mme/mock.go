// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/pkg/errors"

	v1 "github.com/wmnsk/go-gtp/v1"
	v2 "github.com/wmnsk/go-gtp/v2"
	"github.com/wmnsk/go-gtp/v2/ies"
)

// getPGWIP is to get P-GW's IP address according to APN.
//
// DNS should be used in the real case, but here, to keep the example simple,
// this function just returns IP address hard-coded in apnIPMap.
func getPGWIP(apn string) (string, error) {
	apnIPMap := map[string]string{
		"some-apn-1.example": "127.0.0.52",
		"some-apn-2.example": "127.0.0.53",
	}

	if ip, ok := apnIPMap[apn]; ok {
		return ip, nil
	}
	return "", fmt.Errorf("got unknown APN: %s", apn)
}

// dispatch sends subscribers to attachCh, which will be handled in handleAttach().
func dispatch(subs []*v2.Subscriber) {
	for _, sub := range subs {
		// wait for 0-255ms before sending request (just for a little bit of reality)
		/*
			u8buf := make([]byte, 1)
			rand.Read(u8buf)
			time.Sleep(time.Duration(u8buf[0]) * time.Millisecond)
		*/
		time.Sleep(100 * time.Millisecond)

		attachCh <- sub
	}
}

// handleAttach is to start the session creation on S11.
// in the real case this should be called after the procedure on S1AP/NAS has been done.
func handleAttach(raddr net.Addr, c *v2.Conn, sub *v2.Subscriber, br *v2.Bearer) error {
	// remove previous session for the same subscriber if exists.
	sess, err := c.GetSessionByIMSI(sub.IMSI)
	if err != nil {
		switch err.(type) {
		case *v2.ErrUnknownIMSI:
			// whole new session. just ignore.
		default:
			return errors.Wrap(err, "got something unexpected")
		}
	} else {
		teid, err := sess.GetTEID(v2.IFTypeS11S4SGWGTPC)
		if err != nil {
			return v2.ErrTEIDNotFound
		}
		// send Delete Session Request to cleanup sessions in S/P-GW.
		if _, err := c.DeleteSession(teid, sess.PeerAddr); err != nil {
			return errors.Wrap(err, "got something unexpected")
		}
		c.RemoveSession(sess)
	}

	pgwAddr, err := getPGWIP(br.APN)
	if err != nil {
		return err
	}

	var pci, pvi uint8
	if br.PCI {
		pci = 1
	}
	if br.PVI {
		pvi = 1
	}
	localIP := strings.Split(c.LocalAddr().String(), ":")[0]
	session, _, err := c.CreateSession(
		raddr,
		ies.NewIMSI(sub.IMSI),
		ies.NewMSISDN(sub.MSISDN),
		ies.NewMobileEquipmentIdentity(sub.IMEI),
		ies.NewUserLocationInformation(
			0, 0, 0, 1, 1, 0, 0, 0,
			sub.MCC, sub.MCC, sub.LAC, sub.CI, sub.SAI, sub.RAI, sub.TAI, sub.ECI, sub.MeNBI, sub.EMeNBI,
		),
		ies.NewRATType(sub.RATType),
		ies.NewIndicationFromOctets(0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00),
		c.NewFTEID(v2.IFTypeS11MMEGTPC, localIP, ""),
		c.NewFTEID(v2.IFTypeS5S8PGWGTPC, pgwAddr, "").WithInstance(1),
		ies.NewAccessPointName(br.APN),
		ies.NewSelectionMode(v2.SelectionModeMSorNetworkProvidedAPNSubscribedVerified),
		ies.NewPDNType(v2.PDNTypeIPv4),
		ies.NewPDNAddressAllocation("0.0.0.0"),
		ies.NewAPNRestriction(v2.APNRestrictionNoExistingContextsorRestriction),
		ies.NewAggregateMaximumBitRate(0, 0),
		ies.NewBearerContext(
			ies.NewEPSBearerID(br.EBI),
			ies.NewBearerQoS(pci, br.PL, pvi, br.QCI, br.MBRUL, br.MBRDL, br.GBRUL, br.GBRDL),
		),
		ies.NewFullyQualifiedCSID(localIP, 1),
		ies.NewServingNetwork(sub.MCC, sub.MNC),
		ies.NewUETimeZone(9*time.Hour, 0),
	)
	if err != nil {
		return err
	}

	c.AddSession(session)
	return nil
}

var (
	uConn   *v1.UPlaneConn
	payload = []byte{ // ICMP Echo to 8.8.8.8 over IP(src will be replaced), checksum is invalid.
		// IP
		0x45, 0x00, 0x00, 0x54, 0x00, 0x01, 0x40, 0x00, 0x3f, 0x01, 0x00, 0x00, 0xde, 0xad, 0xbe, 0xef,
		0x08, 0x08, 0x08, 0x08,
		// ICMP
		0x08, 0x00, 0x93, 0x6a, 0x00, 0x01, 0x00, 0x01, 0xdf, 0xd5, 0x2c, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x99, 0xea, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x11, 0x12, 0x13,
		0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20, 0x21, 0x22, 0x23,
		0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30, 0x31, 0x32, 0x33,
		0x34, 0x35, 0x36, 0x37,
	}
)

type mockUEeNB struct {
	laddr, raddr net.Addr

	subscriberIP string
	teidOut      uint32
	payload      []byte
}

func (m mockUEeNB) run(errCh chan error) {
	if uConn == nil {
		// Listen on eNB S1-U interface.
		enbUPlaneAddr, err := net.ResolveUDPAddr("udp", *s1enb)
		if err != nil {
			log.Fatal(err)
		}
		m.laddr = enbUPlaneAddr

		uConn, err = v1.ListenAndServeUPlane(m.laddr, 0, errCh)
		if err != nil {
			errCh <- err
			return
		}
	}

	go func(teid uint32, payload []byte, raddr net.Addr) {
		for {
			copy(payload[12:16], net.ParseIP(m.subscriberIP).To4())
			if _, err := uConn.WriteToGTP(teid, m.payload, raddr); err != nil {
				errCh <- err
				return
			}
			time.Sleep(3 * time.Second)
		}
	}(m.teidOut, m.payload, m.raddr)

	go once.Do(func() {
		buf := make([]byte, 1500)
		for {
			if uConn == nil {
				errCh <- errors.New("uConn conn is not open")
				return
			}

			n, raddr, _, err := uConn.ReadFromGTP(buf)
			if err != nil {
				errCh <- err
				return
			}
			loggerCh <- fmt.Sprintf("Received from %s: %x", raddr, buf[:n])
		}
	})
}
