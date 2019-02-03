// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Command sgw is a dead simple implementation of S-GW only with GTP-related features.
//
// S-GW follows the steps below if there's no unexpected events in the middle.
//
// 1. Start listening on S11 interface.
//
// 2. If MME connects to S-GW with Create Session Request, S-GW sends Create Session Request
// to P-GW whose IP is specified by MME with F-TEID IE.
//
// 3. Wait for Create Session Response coming from P-GW with Cause="request accepted", and
// other IEs required are properly set.
//
// 4. Respond to MME with Create Session Response. Here the C-Plane Session is considered to
// be created properly.
//
// 5. If MME sends Modify Bearer Request with eNB information inside, set incoming TEID to
// Bearer and start listening on U-Plane.
//
// 6. If some U-Plane message comes from eNB/P-GW, relay it to P-GW/eNB with TEID and IP
// properly set as told while exchanging the C-Plane signals.
package main

import (
	"flag"
	"log"
	"net"
	"time"

	v1 "github.com/wmnsk/go-gtp/gtp/v1"

	v2 "github.com/wmnsk/go-gtp/gtp/v2"
	"github.com/wmnsk/go-gtp/gtp/v2/messages"
)

// command-line arguments
var (
	s11 = flag.String("s11", "127.0.0.112:2123", "local IP:Port on S11 interface.")
	s5c = flag.String("s5c", "127.0.0.51:2123", "local IP:Port on S5-C interface.")
	s1u = flag.String("s1u", "127.0.0.2:2152", "local IP:Port on S1-U interface.")
	s5u = flag.String("s5u", "127.0.0.3:2152", "local IP:Port on S5-U interface.")
)

var (
	delCh    = make(chan struct{})
	loggerCh = make(chan string)
	errCh    = make(chan error)
)

func main() {
	flag.Parse()
	log.SetPrefix("[S-GW] ")

	// start listening on the specified IP:Port.
	s11laddr, err := net.ResolveUDPAddr("udp", *s11)
	if err != nil {
		log.Fatal(err)
	}
	s1uladdr, err := net.ResolveUDPAddr("udp", *s1u)
	if err != nil {
		log.Fatal(err)
	}
	s5uladdr, err := net.ResolveUDPAddr("udp", *s5u)
	if err != nil {
		log.Fatal(err)
	}

	s11Conn, err := v2.ListenAndServe(s11laddr, 0, errCh)
	if err != nil {
		log.Fatal(err)
	}
	defer s11Conn.Close()
	log.Printf("Started serving on %s", s11Conn.LocalAddr())

	// register handlers for ALL the messages you expect remote endpoint to send.
	s11Conn.AddHandlers(map[uint8]v2.HandlerFunc{
		messages.MsgTypeCreateSessionRequest: handleCreateSessionRequest,
		messages.MsgTypeModifyBearerRequest:  handleModifyBearerRequest,
		messages.MsgTypeDeleteSessionRequest: handleDeleteSessionRequest,
	})

	// let relay start working here.
	// this just drops packets until TEID and peer information is registered.
	s1uConn, err := v1.ListenAndServeUPlane(s1uladdr, 0, errCh)
	if err != nil {
		log.Fatal(err)
	}
	s5uConn, err := v1.ListenAndServeUPlane(s5uladdr, 0, errCh)
	if err != nil {
		log.Fatal(err)
	}

	relay = v1.NewRelay(s1uConn, s5uConn)
	go relay.Run()
	defer relay.Close()

	// wait for events(logs, errors, timers).
	for {
		select {
		case str := <-loggerCh:
			log.Println(str)
		case err := <-errCh:
			log.Printf("Warning: %s", err)
		case <-time.After(10 * time.Second):
			var activeIMSIs []string
			for _, sess := range s11Conn.Sessions {
				if !sess.IsActive() {
					continue
				}
				activeIMSIs = append(activeIMSIs, sess.IMSI)
			}
			if len(activeIMSIs) == 0 {
				continue
			}

			log.Println("Active Subcribers:")
			for _, imsi := range activeIMSIs {
				log.Printf("\t%s", imsi)
			}
			activeIMSIs = nil
		}
	}
}
