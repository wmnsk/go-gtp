// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vishvananda/netlink"
	v1 "github.com/wmnsk/go-gtp/v1"
	"github.com/wmnsk/go-gtp/v2/messages"

	"github.com/pkg/errors"
	v2 "github.com/wmnsk/go-gtp/v2"
)

type pgw struct {
	cConn *v2.Conn
	uConn *v1.UPlaneConn

	s5c, s5u string
	sgiIF    string

	routeSubnet *net.IPNet
	addedRoutes []*netlink.Route
	addedRules  []*netlink.Rule

	promAddr string
	mc       *metricsCollector

	errCh chan error
}

func newPGW(cfg *Config) (*pgw, error) {
	p := &pgw{
		s5c:   cfg.LocalAddrs.S5CAddr,
		s5u:   cfg.LocalAddrs.S5UAddr,
		sgiIF: cfg.SGiIFName,

		errCh: make(chan error, 1),
	}

	var err error
	_, p.routeSubnet, err = net.ParseCIDR(cfg.RouteSubnet)
	if err != nil {
		return nil, err
	}

	if cfg.PromAddr != "" {
		// validate if the address is valid or not.
		if _, err = net.ResolveTCPAddr("tcp", cfg.PromAddr); err != nil {
			return nil, err
		}
		p.promAddr = cfg.PromAddr
	}

	return p, nil
}

func (p *pgw) run(ctx context.Context) error {
	cAddr, err := net.ResolveUDPAddr("udp", p.s5c)
	if err != nil {
		return err
	}
	p.cConn = v2.NewConn(cAddr, 0)
	go func() {
		if err := p.cConn.ListenAndServe(ctx); err != nil {
			log.Println(err)
			return
		}
	}()
	log.Printf("Started serving S5-C on %s", cAddr)

	// register handlers for ALL the messages you expect remote endpoint to send.
	p.cConn.AddHandlers(map[uint8]v2.HandlerFunc{
		messages.MsgTypeCreateSessionRequest: p.handleCreateSessionRequest,
		messages.MsgTypeDeleteSessionRequest: p.handleDeleteSessionRequest,
	})

	uAddr, err := net.ResolveUDPAddr("udp", p.s5u)
	if err != nil {
		return err
	}
	p.uConn = v1.NewUPlaneConn(uAddr)
	if err := p.uConn.EnableKernelGTP("gtp-pgw", v1.RoleGGSN); err != nil {
		return err
	}
	go func() {
		if err = p.uConn.ListenAndServe(ctx); err != nil {
			log.Println(err)
			return
		}
		log.Println("uConn.ListenAndServe exitted")
	}()
	log.Printf("Started serving S5-U on %s", uAddr)

	// start serving Prometheus, if address is given
	if p.promAddr != "" {
		if err := p.runMetricsCollector(); err != nil {
			return err
		}

		http.Handle("/metrics", promhttp.Handler())
		go func() {
			if err := http.ListenAndServe(p.promAddr, nil); err != nil {
				log.Println(err)
			}
		}()
		log.Printf("Started serving Prometheus on %s", p.promAddr)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-p.errCh:
			log.Printf("Warning: %s", err)
		}
	}
}

func (p *pgw) close() error {
	var errs []error
	for _, r := range p.addedRoutes {
		if err := netlink.RouteDel(r); err != nil {
			errs = append(errs, err)
		}
	}
	for _, r := range p.addedRules {
		if err := netlink.RuleDel(r); err != nil {
			errs = append(errs, err)
		}
	}

	if p.uConn != nil {
		if err := netlink.LinkDel(p.uConn.GTPLink); err != nil {
			errs = append(errs, err)
		}
		if err := p.uConn.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if err := p.cConn.Close(); err != nil {
		errs = append(errs, err)
	}

	close(p.errCh)

	if len(errs) > 0 {
		return errors.Errorf("errors while closing S-GW: %v", errs)
	}
	return nil
}

func (p *pgw) setupUPlane(peerIP, msIP net.IP, otei, itei uint32) error {
	if err := p.uConn.AddTunnelOverride(peerIP, msIP, otei, itei); err != nil {
		return err
	}

	ms32 := &net.IPNet{IP: msIP, Mask: net.CIDRMask(32, 32)}
	dlroute := &netlink.Route{ // ip route replace
		Dst:       ms32,                          // UE's IP
		LinkIndex: p.uConn.GTPLink.Attrs().Index, // dev gtp-pgw
		Scope:     netlink.SCOPE_LINK,            // scope link
		Protocol:  4,                             // proto static
		Priority:  1,                             // metric 1
		Table:     3001,                          // table 3001
	}
	if err := netlink.RouteReplace(dlroute); err != nil {
		return err
	}
	p.addedRoutes = append(p.addedRoutes, dlroute)

	link, err := netlink.LinkByName(p.sgiIF)
	if err != nil {
		return err
	}

	ulroute := &netlink.Route{ // ip route replace
		Dst:       p.routeSubnet,      // dst network via SGi
		LinkIndex: link.Attrs().Index, // SGi I/F name
		Scope:     netlink.SCOPE_LINK, // scope link
		Protocol:  4,                  // proto static
		Priority:  1,                  // metric 1
	}
	if err := netlink.RouteReplace(ulroute); err != nil {
		return err
	}
	p.addedRoutes = append(p.addedRoutes, ulroute)

	rules, err := netlink.RuleList(0)
	if err != nil {
		return err
	}
	for _, r := range rules {
		if r.IifName == link.Attrs().Name && r.Dst == ms32 {
			return nil
		}
	}

	rule := netlink.NewRule()
	rule.IifName = link.Attrs().Name
	rule.Dst = ms32
	rule.Table = 3001
	if err := netlink.RuleAdd(rule); err != nil {
		return err
	}
	p.addedRules = append(p.addedRules, rule)

	return nil
}
