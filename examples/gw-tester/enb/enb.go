// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vishvananda/netlink"
	"google.golang.org/grpc"

	"github.com/wmnsk/go-gtp/examples/gw-tester/s1mme"
	"github.com/wmnsk/go-gtp/gtpv1"
)

type enb struct {
	mu sync.Mutex

	// S1-MME
	mmeAddr     net.Addr
	cConn       *grpc.ClientConn
	s1mmeClient s1mme.AttacherClient

	// S1-U
	uAddr net.Addr
	uConn *gtpv1.UPlaneConn

	location *s1mme.Location

	candidateSubs []*Subscriber
	sessions      []*Subscriber
	attachCh      chan *Subscriber

	useKernelGTP bool

	addedAddrs  map[netlink.Link]*netlink.Addr
	addedRoutes []*netlink.Route
	addedRules  []*netlink.Rule

	promAddr string
	mc       *metricsCollector

	errCh chan error
}

func newENB(cfg *Config) (*enb, error) {
	e := &enb{
		mu: sync.Mutex{},
		location: &s1mme.Location{
			Mcc:     cfg.MCC,
			Mnc:     cfg.MNC,
			RatType: s1mme.Location_RATType(cfg.RATType),
			Tai:     uint32(cfg.TAI),
			Eci:     cfg.ECI,
		},
		candidateSubs: cfg.Subscribers,
		useKernelGTP:  cfg.UseKernelGTP,
		addedAddrs:    make(map[netlink.Link]*netlink.Addr),

		errCh: make(chan error, 1),
	}

	var err error
	e.uAddr, err = net.ResolveUDPAddr("udp", cfg.LocalAddrs.S1UIP+gtpv1.GTPUPort)
	if err != nil {
		return nil, err
	}

	e.mmeAddr, err = net.ResolveTCPAddr("tcp", cfg.MMEAddr)
	if err != nil {
		return nil, err
	}

	if cfg.PromAddr != "" {
		// validate if the address is valid or not.
		if _, err = net.ResolveTCPAddr("tcp", cfg.PromAddr); err != nil {
			return nil, err
		}
		e.promAddr = cfg.PromAddr
	}

	if !e.useKernelGTP {
		log.Println("WARN: U-Plane does not work without GTP kernel module")
	}

	return e, nil
}

func (e *enb) run(ctx context.Context) error {
	// TODO: bind local address(cfg.LocalAddrs.S1CIP) with WithDialer option?
	conn, err := grpc.Dial(e.mmeAddr.String(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	e.s1mmeClient = s1mme.NewAttacherClient(conn)
	log.Printf("Established S1-MME connection with %s", e.mmeAddr)

	e.uConn = gtpv1.NewUPlaneConn(e.uAddr)
	if e.useKernelGTP {
		if err := e.uConn.EnableKernelGTP("gtp-enb", gtpv1.RoleSGSN); err != nil {
			return err
		}
	}
	go func() {
		if err := e.uConn.ListenAndServe(ctx); err != nil {
			log.Println(err)
			return
		}
		log.Println("uConn.ListenAndServe exitted")
	}()
	log.Printf("Started serving S1-U on %s", e.uAddr)

	// start serving Prometheus, if address is given
	if e.promAddr != "" {
		if err := e.runMetricsCollector(); err != nil {
			return err
		}

		http.Handle("/metrics", promhttp.Handler())
		go func() {
			if err := http.ListenAndServe(e.promAddr, nil); err != nil {
				log.Println(err)
			}
		}()
		log.Printf("Started serving Prometheus on %s", e.promAddr)
	}

	for _, sub := range e.candidateSubs {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(10 * time.Millisecond):
			// not to load too much in case of many subscribers
		}

		if err := e.attach(ctx, sub); err != nil {
			return err
		}
	}

	// wait for new subscribers to be attached
	e.attachCh = make(chan *Subscriber)
	defer close(e.attachCh)
	for {
		select {
		case <-ctx.Done():
			return nil
		case sub := <-e.attachCh:
			go func() {
				if err := e.attach(ctx, sub); err != nil {
					log.Printf("Failed to attach: %s: %s", sub, err)
				}
			}()
			// wait 10ms after dispatching
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (e *enb) reload(cfg *Config) error {
	e.location = &s1mme.Location{
		Mcc:     cfg.MCC,
		Mnc:     cfg.MNC,
		RatType: s1mme.Location_RATType(cfg.RATType),
		Tai:     uint32(cfg.TAI),
		Eci:     cfg.ECI,
	}

	// TODO: consider more efficient way
	var attachSubs []*Subscriber
	for _, sub := range cfg.Subscribers {
		var existing bool
		for _, attached := range e.sessions {
			if sub.IMSI == attached.IMSI {
				existing = true
			}
		}

		if existing {
			if !sub.Reattach {
				continue
			}

			if err := e.uConn.DelTunnelByITEI(sub.ITEI); err != nil {
				continue
			}
		}
		attachSubs = append(attachSubs, sub)
	}

	for _, sub := range attachSubs {
		e.attachCh <- sub
	}
	return nil
}

func (e *enb) close() error {
	var errs []error
	e.mu.Lock()
	defer e.mu.Unlock()

	for l, a := range e.addedAddrs {
		if err := netlink.AddrDel(l, a); err != nil {
			errs = append(errs, err)
		}
	}

	for _, r := range e.addedRoutes {
		if err := netlink.RouteDel(r); err != nil {
			errs = append(errs, err)
		}
	}

	for _, r := range e.addedRules {
		if err := netlink.RuleDel(r); err != nil {
			errs = append(errs, err)
		}
	}

	if c := e.uConn; c != nil {
		if err := c.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if c := e.cConn; c != nil {
		if err := c.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	return fmt.Errorf("errors while closing eNB: %v", errs)
}

func (e *enb) attach(ctx context.Context, sub *Subscriber) error {
	// allocate random TEID if 0 is specified in config.
	if sub.ITEI == 0 {
		sub.ITEI = e.newTEID()
	}

	req := &s1mme.AttachRequest{
		Imsi:     sub.IMSI,
		Msisdn:   sub.MSISDN,
		Imeisv:   sub.IMEISV,
		S1UAddr:  e.uConn.LocalAddr().String(),
		SrcIp:    sub.SrcIP,
		ITei:     sub.ITEI,
		Location: e.location,
	}

	rsp, err := e.s1mmeClient.Attach(ctx, req)
	if err != nil {
		return err
	}
	if e.mc != nil {
		e.mc.messagesSent.WithLabelValues(e.mmeAddr.String(), "Attach Request").Inc()
		e.mc.messagesReceived.WithLabelValues(e.mmeAddr.String(), "Attach Response").Inc()
	}

	switch rsp.Cause {
	case s1mme.Cause_SUCCESS:
		sgwIP, _, err := net.SplitHostPort(rsp.SgwAddr)
		if err != nil {
			return err
		}

		if e.useKernelGTP {
			if err := e.uConn.AddTunnelOverride(
				net.ParseIP(sgwIP), net.ParseIP(req.SrcIp), rsp.OTei, req.ITei,
			); err != nil {
				log.Println(net.ParseIP(sgwIP), net.ParseIP(req.SrcIp), rsp.OTei, req.ITei)
				e.errCh <- fmt.Errorf("failed to create tunnel for %s: %w", sub.IMSI, err)
				return nil
			}
			if err := e.addRoute(); err != nil {
				e.errCh <- fmt.Errorf("failed to add route for %s: %w", sub.IMSI, err)
				return nil
			}
		}

		sub.sgwAddr = rsp.SgwAddr
		sub.otei = rsp.OTei

		e.sessions = append(e.sessions, sub)
		if e.useKernelGTP {
			if err := e.setupUPlane(ctx, sub); err != nil {
				e.errCh <- fmt.Errorf("failed to setup U-Plane for %s: %w", sub.IMSI, err)
				return nil
			}
			log.Printf("Successfully established tunnel for %s", sub.IMSI)
		}
	default:
		e.errCh <- fmt.Errorf("got unexpected Cause for %s: %s", rsp.Cause, sub.IMSI)
		return nil
	}

	return nil
}

func (e *enb) newTEID() uint32 {
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return 0
	}

	generated := binary.BigEndian.Uint32(b)
	for _, s := range e.sessions {
		if generated == s.ITEI {
			return e.newTEID()
		}
	}

	return generated
}

func (e *enb) setupUPlane(ctx context.Context, sub *Subscriber) error {
	switch sub.TrafficType {
	case "http_get":
		if err := e.addIP(sub); err != nil {
			return err
		}
		if err := e.addRuleLocal(sub); err != nil {
			return err
		}
		go func() {
			if err := e.runHTTPProbe(ctx, sub); err != nil {
				e.errCh <- err
			}
		}()
	case "external":
		if err := e.addRuleExternal(sub); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown/unimplemented type: %s specified for 'type' in subscriber", sub.TrafficType)
	}

	return nil
}

func (e *enb) addRoute() error {
	route := &netlink.Route{
		Dst:       &net.IPNet{IP: net.IPv4zero, Mask: net.CIDRMask(0, 32)}, // default
		LinkIndex: e.uConn.KernelGTP.Link.Attrs().Index,                    // dev gtp-<ECI>
		Scope:     netlink.SCOPE_LINK,                                      // scope link
		Protocol:  4,                                                       // proto static
		Priority:  1,                                                       // metric 1
		Table:     1001,                                                    // table <ECI>
	}

	e.addedRoutes = append(e.addedRoutes, route)
	return netlink.RouteReplace(route)
}

func (e *enb) addRuleExternal(sub *Subscriber) error {
	rules, err := netlink.RuleList(0)
	if err != nil {
		return err
	}

	mask32 := &net.IPNet{IP: net.ParseIP(sub.SrcIP), Mask: net.CIDRMask(32, 32)}
	for _, r := range rules {
		if r.IifName == sub.EUuIFName && r.Src == mask32 && r.Table == 1001 {
			return nil
		}
	}

	rule := netlink.NewRule()
	rule.IifName = sub.EUuIFName
	rule.Src = mask32
	rule.Table = 1001

	e.addedRules = append(e.addedRules, rule)
	return netlink.RuleAdd(rule)
}

func (e *enb) addRuleLocal(sub *Subscriber) error {
	rules, err := netlink.RuleList(0)
	if err != nil {
		return err
	}

	mask32 := &net.IPNet{IP: net.ParseIP(sub.SrcIP), Mask: net.CIDRMask(32, 32)}
	for _, r := range rules {
		if r.Src == mask32 && r.Table == 1001 {
			return nil
		}
	}

	rule := netlink.NewRule()
	rule.Src = mask32
	rule.Table = 1001

	e.addedRules = append(e.addedRules, rule)
	return netlink.RuleAdd(rule)
}

func (e *enb) runHTTPProbe(ctx context.Context, sub *Subscriber) error {
	laddr, err := net.ResolveTCPAddr("tcp", sub.SrcIP+":0")
	if err != nil {
		return err
	}
	dialer := net.Dialer{LocalAddr: laddr}
	client := http.Client{
		Transport: &http.Transport{Dial: dialer.Dial},
		Timeout:   3 * time.Second,
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(5 * time.Second):
			// do nothing here and go forward
		}

		rsp, err := client.Get(sub.HTTPURL)
		if err != nil {
			e.errCh <- fmt.Errorf("failed to GET %s: %w", sub.HTTPURL, err)
			continue
		}

		if rsp.StatusCode == http.StatusOK {
			log.Printf("[HTTP Probe;%s] Successfully GET %s: Status: %s", sub.IMSI, sub.HTTPURL, rsp.Status)
			rsp.Body.Close()
			continue
		}
		rsp.Body.Close()
		e.errCh <- fmt.Errorf("got invalid response on HTTP probe: %v", rsp.StatusCode)
	}
}

func (e *enb) addIP(sub *Subscriber) error {
	link, err := netlink.LinkByName(sub.EUuIFName)
	if err != nil {
		return err
	}
	addrs, err := netlink.AddrList(link, netlink.FAMILY_ALL)
	if err != nil {
		return err
	}

	netToAdd := &net.IPNet{IP: net.ParseIP(sub.SrcIP), Mask: net.CIDRMask(24, 32)}
	var addr netlink.Addr
	var found bool
	for _, a := range addrs {
		if a.Label == sub.EUuIFName {
			if a.IPNet.String() == netToAdd.String() {
				return nil
			}
			addr = a
			found = true
		}
	}
	if !found {
		return fmt.Errorf("cannot find the interface to add address: %s", sub.EUuIFName)
	}

	addr.IPNet = netToAdd
	if err := netlink.AddrAdd(link, &addr); err != nil {
		return err
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	e.addedAddrs[link] = &addr

	return nil
}
