// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Command mme is a reference implementation of MME with go-gtp.
//
// MME follows the steps below if there's no unexpected events in the middle.
// Note that the  S1 and DNS procedures is just mocked to make it work in
// standalone manner.
//
// 1. Exchange Echo to S-GW address specified in command-line argument.
//
// 2. Start dispatching subscribers by sending Create Session Request to S-GW.
// APN is handled with getPGWIP(), which is hard-coded.
//
// 3. Wait for Create Session Response coming from S-GW with Cause="request accepted".
//
// 4. Create mocked UE and eNB with the required values set as told by S-GW, start
// listening on the interface specified with s1enb flag,  and send Modify Bearer Request
// to S-GW.
//
// 5. Wait for Modify Bearer Response coming from S-GW with Cause="request accepted".
//
// 6. Start sending payload(ICMP Echo Request) encapsulated with GTPv1-U Header, and printing
// the payload of encapsulated packets received.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/wmnsk/go-gtp/gtpv1"
	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
)

// command-line flags.
var (
	s11mme = flag.String("s11mme", "127.0.0.111", "local IP on S11 interface.")
	s11sgw = flag.String("s11sgw", "127.0.0.112", "S-GW's IP on S11 interface.")
	s1enb  = flag.String("s1enb", "127.0.0.1", "local IP on S1-U of pseudo eNB.")
)

// variables globally shared.
var (
	attachCh  = make(chan *gtpv2.Subscriber)
	createdCh = make(chan string)
	loggerCh  = make(chan string)
	errCh     = make(chan error)

	once  = sync.Once{}
	delWG = sync.WaitGroup{}
)

func main() {
	flag.Parse()
	log.SetPrefix("[MME] ")

	laddr, err := net.ResolveUDPAddr("udp", *s11mme+gtpv2.GTPCPort)
	if err != nil {
		log.Println(err)
		return
	}
	raddr, err := net.ResolveUDPAddr("udp", *s11sgw+gtpv2.GTPCPort)
	if err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup *Conn first to check if the remote endpoint is awaken.
	s11Conn, err := gtpv2.Dial(ctx, laddr, raddr, gtpv2.IFTypeS11MMEGTPC, 0)
	if err != nil {
		log.Println(err)
		return
	}
	defer s11Conn.Close()
	log.Printf("Connection established with %s", raddr.String())

	// register handlers for ALL the message you expect remote endpoint to send.
	// by default, Echo and VersionNotsupported is handled without explicit declaration.
	s11Conn.AddHandlers(map[uint8]gtpv2.HandlerFunc{
		message.MsgTypeCreateSessionResponse: handleCreateSessionResponse,
		message.MsgTypeModifyBearerResponse:  handleModifyBearerResponse,
		message.MsgTypeDeleteSessionResponse: handleDeleteSessionResponse,
	})

	// Listen on eNB S1-U interface.
	enbUPlaneAddr, err := net.ResolveUDPAddr("udp", *s1enb+gtpv2.GTPUPort)
	if err != nil {
		log.Println(err)
		return
	}

	uConn = gtpv1.NewUPlaneConn(enbUPlaneAddr)
	defer uConn.Close()

	go func() {
		if err = uConn.ListenAndServe(ctx); err != nil {
			log.Println(err)
			return
		}
	}()
	log.Printf("Started listening on %s", enbUPlaneAddr)

	// here you should wait for UEs to come attaching to your network.
	// in this example, the following five subscribers are to be attached.
	// working as worker-dispatcher is preferable in the real case
	go dispatch([]*gtpv2.Subscriber{{
		IMSI: "123451234567891", MSISDN: "8130900000001", IMEI: "123456780000011",
		Location: &gtpv2.Location{MCC: "123", MNC: "45", RATType: gtpv2.RATTypeEUTRAN, TAI: 0x0001, ECI: 0x00000101},
	}, {
		IMSI: "123451234567892", MSISDN: "8130900000002", IMEI: "123456780000012",
		Location: &gtpv2.Location{MCC: "123", MNC: "45", RATType: gtpv2.RATTypeEUTRAN, TAI: 0x0002, ECI: 0x00000202},
	}, {
		IMSI: "123451234567893", MSISDN: "8130900000003", IMEI: "123456780000013",
		Location: &gtpv2.Location{MCC: "123", MNC: "45", RATType: gtpv2.RATTypeEUTRAN, TAI: 0x0003, ECI: 0x00000303},
	}, {
		IMSI: "123451234567894", MSISDN: "8130900000004", IMEI: "123456780000014",
		Location: &gtpv2.Location{MCC: "123", MNC: "45", RATType: gtpv2.RATTypeEUTRAN, TAI: 0x0004, ECI: 0x00000404},
	}, {
		IMSI: "123451234567895", MSISDN: "8130900000005", IMEI: "123456780000015",
		Location: &gtpv2.Location{MCC: "123", MNC: "45", RATType: gtpv2.RATTypeEUTRAN, TAI: 0x0005, ECI: 0x00000505},
	}})

	bearer := gtpv2.NewBearer(5, "", &gtpv2.QoSProfile{
		PL: 2, QCI: 255, MBRUL: 0xffffffff, MBRDL: 0xffffffff, GBRUL: 0xffffffff, GBRDL: 0xffffffff,
	})
	for {
		select {
		// print logs coming from handlers working background
		case str := <-loggerCh:
			log.Println(str)
		// print errors coming from handlers working background
		// it's better to switch over the error to distinguish fatal ones to others.
		case err := <-errCh:
			log.Printf("Warning: %s", err)
		// handle attach requests
		case sub := <-attachCh:
			log.Printf("Started creating session for subscriber: %s", sub.IMSI)
			go func() {
				bearer.APN = "some-apn-1.example"
				if sub.TAI%2 == 0 {
					bearer.APN = "some-apn-2.example"
				}
				if err := handleAttach(raddr, s11Conn, sub, bearer); err != nil {
					errCh <- err
					return
				}
			}()
		case imsi := <-createdCh:
			go func() {
				sess, err := s11Conn.GetSessionByIMSI(imsi)
				if err != nil {
					errCh <- err
					return
				}

				enbIP := strings.Split(*s1enb, ":")[0]
				enbFTEID := uConn.NewFTEID(gtpv2.IFTypeS1UeNodeBGTPU, enbIP, "")
				teid, err := sess.GetTEID(gtpv2.IFTypeS11S4SGWGTPC)
				if err != nil {
					errCh <- err
				}

				if _, err := s11Conn.ModifyBearer(
					teid, sess,
					ie.NewIndicationFromOctets(0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00),
					ie.NewBearerContext(ie.NewEPSBearerID(sess.GetDefaultBearer().EBI), enbFTEID),
				); err != nil {
					errCh <- err
					return
				}
				it, err := enbFTEID.InterfaceType()
				if err != nil {
					errCh <- err
					return
				}
				enbTEID, err := enbFTEID.TEID()
				if err != nil {
					errCh <- err
					return
				}
				sess.AddTEID(it, enbTEID)

				loggerCh <- fmt.Sprintf("Sent Modify Bearer Request for %s", imsi)
			}()
		// delete all the sessions after 30 seconds
		case <-time.After(30 * time.Second):
			for _, sess := range s11Conn.Sessions() {
				teid, err := sess.GetTEID(gtpv2.IFTypeS11S4SGWGTPC)
				if err != nil {
					errCh <- gtpv2.ErrTEIDNotFound
					return
				}
				if _, err := s11Conn.DeleteSession(teid, sess); err != nil {
					log.Printf("Warning: %s", err)
				}
				delWG.Add(1)
				log.Printf("Sent Delete Session Request for %s", sess.IMSI)
			}

			// invoke goroutine to let the logger work
			go func() {
				delWG.Wait()
				log.Fatal("Inactivity timer expired, exitting...")
			}()
		}
	}
}
