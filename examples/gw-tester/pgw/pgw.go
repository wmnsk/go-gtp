// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"log"
	"net"

	"github.com/vishvananda/netlink"
	v1 "github.com/wmnsk/go-gtp/v1"

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

	errCh chan error
}

func newPGW(cfg *Config) (*pgw, error) {
	p := &pgw{
		s5c:   cfg.LocalAddrs.S5C,
		s5u:   cfg.LocalAddrs.S5U,
		sgiIF: cfg.SGiIFName,

		errCh: make(chan error, 1),
	}

	laddr, err := net.ResolveUDPAddr("udp", p.s5c)
	if err != nil {
		return nil, err
	}

	p.cConn, err = v2.ListenAndServe(laddr, 0, p.errCh)
	if err != nil {
		return nil, err
	}

	_, p.routeSubnet, err = net.ParseCIDR(cfg.RouteSubnet)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *pgw) run(ctx context.Context) error {
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
	if p.uConn == nil {
		laddr, err := net.ResolveUDPAddr("udp", p.s5u)
		if err != nil {
			return err
		}
		p.uConn, err = v1.ListenAndServeUPlaneKernel("gtp-pgw", v1.RoleGGSN, laddr, 0, p.errCh)
		if err != nil {
			return err
		}
		log.Printf("Started listening on %s", p.uConn.LocalAddr())
	}

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
