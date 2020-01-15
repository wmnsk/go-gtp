// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/wmnsk/go-gtp/examples/gw-tester/s1mme"
	v2 "github.com/wmnsk/go-gtp/v2"
	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
)

// Session represents a subscriber.
type Session struct {
	IMSI   string
	MSISDN string
	IMEISV string

	SrcIP string

	itei uint32
}

type mme struct {
	s1mmeListener net.Listener
	s11Addr       net.Addr
	s11IP         string
	s11Conn       *v2.Conn

	created  chan struct{}
	modified chan struct{}

	apn      string
	mcc, mnc string

	enb struct {
		mcc   string
		mnc   string
		tai   uint16
		eci   uint32
		s1uIP string
	}

	sgw struct {
		s11IP string
		s1uIP string
	}

	pgw struct {
		s5cIP string
	}

	errCh chan error
}

func newMME(cfg *Config) (*mme, error) {
	m := &mme{
		mcc: cfg.MCC,
		mnc: cfg.MNC,
		apn: cfg.APN,

		created:  make(chan struct{}, 1),
		modified: make(chan struct{}, 1),

		errCh: make(chan error, 1),
	}
	m.sgw.s11IP = cfg.SgwS11
	m.pgw.s5cIP = cfg.PgwS5C

	// setup S11 conn
	var err error
	m.s11Addr, err = net.ResolveUDPAddr("udp", cfg.LocalAddrs.S11Addr)
	if err != nil {
		return nil, err
	}
	m.s11IP, _, err = net.SplitHostPort(cfg.LocalAddrs.S11Addr)
	if err != nil {
		return nil, err
	}

	// setup gRPC server
	m.s1mmeListener, err = net.Listen("tcp", cfg.LocalAddrs.S1CAddr)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *mme) run(ctx context.Context) error {
	fatalCh := make(chan error, 1)

	srv := grpc.NewServer()
	s1mme.RegisterAttacherServer(srv, m)
	go func() {
		if err := srv.Serve(m.s1mmeListener); err != nil {
			fatalCh <- errors.Errorf("error on serving gRPC: %s", err)
			return
		}
	}()

	m.s11Conn = v2.NewConn(m.s11Addr, 0)
	go func() {
		if err := m.s11Conn.ListenAndServe(ctx); err != nil {
			log.Println(err)
			return
		}
	}()
	log.Printf("Started running MME\n  S1MME: %s, S11: %s", m.s1mmeListener.Addr(), m.s11Addr)

	m.s11Conn.AddHandlers(map[uint8]v2.HandlerFunc{
		messages.MsgTypeCreateSessionResponse: m.handleCreateSessionResponse,
		messages.MsgTypeModifyBearerResponse:  m.handleModifyBearerResponse,
		messages.MsgTypeDeleteSessionResponse: m.handleDeleteSessionResponse,
	})

	for {
		select {
		case <-ctx.Done():
			// srv.Serve returns when lis is closed
			if err := m.s1mmeListener.Close(); err != nil {
				return err
			}
			return nil
		case err := <-fatalCh:
			return err
		}
	}
}

func (m *mme) reload(cfg *Config) error {
	// TODO: implement
	return nil
}

// Attach is called by eNB by gRPC.
func (m *mme) Attach(ctx context.Context, req *s1mme.AttachRequest) (*s1mme.AttachResponse, error) {
	sess := &Session{
		IMSI:   req.Imsi,
		MSISDN: req.Msisdn,
		IMEISV: req.Imeisv,
		SrcIP:  req.SrcIp,
		itei:   req.ITei,
	}

	var err error
	m.enb.s1uIP, _, err = net.SplitHostPort(req.S1UAddr)
	if err != nil {
		return nil, err
	}

	errCh := make(chan error, 1)
	rspCh := make(chan *s1mme.AttachResponse)
	go func() {
		m.enb.mcc = req.Location.Mcc
		m.enb.mnc = req.Location.Mnc
		m.enb.tai = uint16(req.Location.Tai)
		m.enb.eci = req.Location.Eci

		session, err := m.CreateSession(sess)
		if err != nil {
			errCh <- err
			return
		}
		log.Printf("Sent Create Session Request for %s", session.IMSI)

		select {
		case <-m.created:
			// go forward
		case <-time.After(5 * time.Second):
			errCh <- errors.Errorf("timed out: %s", session.IMSI)
		}

		if _, err = m.ModifyBearer(session, sess); err != nil {
			errCh <- err
			return
		}
		log.Printf("Sent Modify Bearer Request for %s", session.IMSI)

		select {
		case <-m.modified:
			// go forward
		case <-time.After(5 * time.Second):
			errCh <- errors.Errorf("timed out: %s", session.IMSI)
		}

		s1teid, err := session.GetTEID(v2.IFTypeS1USGWGTPU)
		if err != nil {
			errCh <- err
			return
		}

		rspCh <- &s1mme.AttachResponse{
			Cause:   s1mme.Cause_SUCCESS,
			SgwAddr: m.sgw.s1uIP + ":2152",
			OTei:    s1teid,
		}
	}()

	select {
	case err := <-errCh:
		return nil, err
	case rsp := <-rspCh:
		return rsp, nil
	}
}

// Detach is called by eNB by gRPC.
func (m *mme) Detach(ctx context.Context, req *s1mme.DetachRequest) (*s1mme.DetachResponse, error) {
	// TODO: implement
	return nil, nil
}

func (m *mme) CreateSession(sess *Session) (*v2.Session, error) {
	br := v2.NewBearer(5, "", &v2.QoSProfile{
		PL: 2, QCI: 255, MBRUL: 0xffffffff, MBRDL: 0xffffffff, GBRUL: 0xffffffff, GBRDL: 0xffffffff,
	})
	var pci, pvi uint8
	if br.PCI {
		pci = 1
	}
	if br.PVI {
		pvi = 1
	}

	raddr, err := net.ResolveUDPAddr("udp", m.sgw.s11IP+":2123")
	if err != nil {
		return nil, err
	}
	session, _, err := m.s11Conn.CreateSession(
		raddr,
		ies.NewIMSI(sess.IMSI),
		ies.NewMSISDN(sess.MSISDN),
		ies.NewMobileEquipmentIdentity(sess.IMEISV),
		ies.NewUserLocationInformation(
			0, 0, 0, 1, 1, 0, 0, 0,
			m.enb.mcc, m.enb.mnc, 0, 0, 0, 0, 1, 1, 0, 0,
		),
		ies.NewRATType(v2.RATTypeEUTRAN),
		ies.NewIndicationFromOctets(0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00),
		m.s11Conn.NewFTEID(v2.IFTypeS11MMEGTPC, m.s11IP, ""),
		m.s11Conn.NewFTEID(v2.IFTypeS5S8PGWGTPC, m.pgw.s5cIP, "").WithInstance(1),
		ies.NewAccessPointName(m.apn),
		ies.NewSelectionMode(v2.SelectionModeMSorNetworkProvidedAPNSubscribedVerified),
		ies.NewPDNType(v2.PDNTypeIPv4),
		ies.NewPDNAddressAllocation(sess.SrcIP),
		ies.NewAPNRestriction(v2.APNRestrictionNoExistingContextsorRestriction),
		ies.NewAggregateMaximumBitRate(0, 0),
		ies.NewBearerContext(
			ies.NewEPSBearerID(br.EBI),
			ies.NewBearerQoS(pci, br.PL, pvi, br.QCI, br.MBRUL, br.MBRDL, br.GBRUL, br.GBRDL),
		),
		ies.NewFullyQualifiedCSID(m.s11IP, 1),
		ies.NewServingNetwork(m.mcc, m.mnc),
		ies.NewUETimeZone(9*time.Hour, 0),
	)
	if err != nil {
		return nil, err
	}

	m.s11Conn.AddSession(session)
	return session, nil
}

func (m *mme) ModifyBearer(sess *v2.Session, sub *Session) (*v2.Bearer, error) {
	teid, err := sess.GetTEID(v2.IFTypeS11S4SGWGTPC)
	if err != nil {
		return nil, err
	}

	fteid := ies.NewFullyQualifiedTEID(v2.IFTypeS1UeNodeBGTPU, sub.itei, m.enb.s1uIP, "")
	if _, err = m.s11Conn.ModifyBearer(
		teid, sess.PeerAddr(), ies.NewIndicationFromOctets(0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00),
		ies.NewBearerContext(ies.NewEPSBearerID(sess.GetDefaultBearer().EBI), fteid, ies.NewPortNumber(2152)),
	); err != nil {
		return nil, err
	}

	return nil, nil
}
