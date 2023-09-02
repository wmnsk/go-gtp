// Copyright 2019-2023 go-gtp authors. All rights reserved.
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
	"context"
	"flag"
	"log"
	"net"
	"time"

	"github.com/wmnsk/go-gtp/gtpv1"
	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/message"
)

// command-line arguments and global variables
var (
	s11 = flag.String("s11", "127.0.0.112", "local IP on S11 interface.")
	s5c = flag.String("s5c", "127.0.0.51", "local IP on S5-C interface.")
	s1u = flag.String("s1u", "127.0.0.2", "local IP on S1-U interface.")
	s5u = flag.String("s5u", "127.0.0.3", "local IP on S5-U interface.")

	sgw *sGateway
)

type sGateway struct {
	s11Conn, s5cConn *gtpv2.Conn
	s1uConn, s5uConn *gtpv1.UPlaneConn

	loggerCh chan string
	errCh    chan error
}

func newSGW(s11, s5c, s1u, s5u net.Addr) (*sGateway, error) {
	s := &sGateway{
		loggerCh: make(chan string),
		errCh:    make(chan error),
	}

	ctx := context.Background()

	var err error
	s.s11Conn = gtpv2.NewConn(s11, gtpv2.IFTypeS11S4SGWGTPC, 0)
	go func() {
		if err = s.s11Conn.ListenAndServe(ctx); err != nil {
			log.Println(err)
			return
		}
	}()
	log.Printf("Started serving on %s", s11)

	s.s5cConn = gtpv2.NewConn(s5c, gtpv2.IFTypeS5S8SGWGTPC, 0)
	go func() {
		if err = s.s5cConn.ListenAndServe(ctx); err != nil {
			log.Println(err)
			return
		}
	}()
	log.Printf("Started serving on %s", s5c)

	s.s1uConn = gtpv1.NewUPlaneConn(s1u)
	go func() {
		if err = s.s1uConn.ListenAndServe(ctx); err != nil {
			log.Println(err)
			return
		}
	}()

	go func() {
		s.s5uConn = gtpv1.NewUPlaneConn(s5u)
		if err = s.s5uConn.ListenAndServe(ctx); err != nil {
			log.Println(err)
			return
		}
	}()

	return s, nil
}

func (s *sGateway) run() error {
	defer func() {
		s.s11Conn.Close()
		s.s5cConn.Close()
	}()

	// wait for events(logs, errors, timers).
	for {
		select {
		case str := <-s.loggerCh:
			log.Println(str)
		case err := <-s.errCh:
			log.Printf("Warning: %v", err)
		case <-time.After(10 * time.Second):
			var activeIMSIs []string
			for _, sess := range s.s11Conn.Sessions() {
				if !sess.IsActive() {
					continue
				}
				activeIMSIs = append(activeIMSIs, sess.IMSI)
			}
			if len(activeIMSIs) == 0 {
				continue
			}

			log.Println("Active Subscribers:")
			for _, imsi := range activeIMSIs {
				log.Printf("\t%s", imsi)
			}
			activeIMSIs = nil
		}
	}
}

func main() {
	flag.Parse()
	log.SetPrefix("[S-GW] ")

	// resolve specified IP:Port as net.UDPAddr.
	s11, err := net.ResolveUDPAddr("udp", *s11+gtpv2.GTPCPort)
	if err != nil {
		log.Println(err)
		return
	}
	s5c, err := net.ResolveUDPAddr("udp", *s5c+gtpv2.GTPCPort)
	if err != nil {
		log.Println(err)
		return
	}
	s1u, err := net.ResolveUDPAddr("udp", *s1u+gtpv2.GTPUPort)
	if err != nil {
		log.Println(err)
		return
	}
	s5u, err := net.ResolveUDPAddr("udp", *s5u+gtpv2.GTPUPort)
	if err != nil {
		log.Println(err)
		return
	}

	sgw, err = newSGW(s11, s5c, s1u, s5u)
	if err != nil {
		log.Println(err)
		return
	}

	// register handlers for ALL the message you expect remote endpoint to send.
	sgw.s11Conn.AddHandlers(map[uint8]gtpv2.HandlerFunc{
		message.MsgTypeCreateSessionRequest: handleCreateSessionRequest,
		message.MsgTypeModifyBearerRequest:  handleModifyBearerRequest,
		message.MsgTypeDeleteSessionRequest: handleDeleteSessionRequest,
		message.MsgTypeDeleteBearerResponse: handleDeleteBearerResponse,
	})
	sgw.s5cConn.AddHandlers(map[uint8]gtpv2.HandlerFunc{
		message.MsgTypeCreateSessionResponse: handleCreateSessionResponse,
		message.MsgTypeDeleteSessionResponse: handleDeleteSessionResponse,
		message.MsgTypeDeleteBearerRequest:   handleDeleteBearerRequest,
	})

	log.Fatal(sgw.run())
}
