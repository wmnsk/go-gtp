// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"log"
	"net"

	"github.com/pkg/errors"
	"github.com/vishvananda/netlink"

	v1 "github.com/wmnsk/go-gtp/v1"
	v2 "github.com/wmnsk/go-gtp/v2"
)

type sgw struct {
	s11Conn, s5cConn *v2.Conn
	s1uConn, s5uConn *v1.UPlaneConn

	s11IP, s5cIP, s1uIP, s5uIP string

	enableKernGTP bool

	addedRoutes []*netlink.Route
	addedRules  []*netlink.Rule

	errCh chan error
}

func newSGW(cfg *Config) (*sgw, error) {
	s := &sgw{
		enableKernGTP: cfg.EnableKernGTP,
		errCh:         make(chan error, 1),
	}

	s11, err := net.ResolveUDPAddr("udp", cfg.LocalAddrs.S11)
	if err != nil {
		return nil, err
	}
	s.s11Conn, err = v2.ListenAndServe(s11, 0, s.errCh)
	if err != nil {
		return nil, err
	}
	s.s11IP, _, err = net.SplitHostPort(s11.String())
	if err != nil {
		return nil, err
	}
	log.Printf("Started serving on %s", s.s11Conn.LocalAddr())

	s5c, err := net.ResolveUDPAddr("udp", cfg.LocalAddrs.S5C)
	if err != nil {
		return nil, err
	}
	s.s5cConn, err = v2.ListenAndServe(s5c, 0, s.errCh)
	if err != nil {
		return nil, err
	}
	s.s5cIP, _, err = net.SplitHostPort(s5c.String())
	if err != nil {
		return nil, err
	}
	log.Printf("Started serving on %s", s.s5cConn.LocalAddr())

	s1u, err := net.ResolveUDPAddr("udp", cfg.LocalAddrs.S1U)
	if err != nil {
		return nil, err
	}
	s5u, err := net.ResolveUDPAddr("udp", cfg.LocalAddrs.S5U)
	if err != nil {
		return nil, err
	}

	if s.enableKernGTP {
		s.s1uConn, err = v1.ListenAndServeUPlaneKernel("gtp-sgw-s1", v1.RoleGGSN, s1u, 0, s.errCh)
		if err != nil {
			return nil, err
		}

		s.s5uConn, err = v1.ListenAndServeUPlaneKernel("gtp-sgw-s5", v1.RoleSGSN, s5u, 0, s.errCh)
		if err != nil {
			return nil, err
		}

		if err := s.addRoutes(); err != nil {
			return nil, err
		}
	} else {
		s.s1uConn, err = v1.ListenAndServeUPlane(s1u, 0, s.errCh)
		if err != nil {
			return nil, err
		}

		s.s5uConn, err = v1.ListenAndServeUPlane(s5u, 0, s.errCh)
		if err != nil {
			return nil, err
		}
	}

	s.s1uIP, _, err = net.SplitHostPort(s1u.String())
	if err != nil {
		return nil, err
	}
	s.s5uIP, _, err = net.SplitHostPort(s5u.String())
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *sgw) run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-s.errCh:
			log.Printf("Warning: %s", errors.WithStack(err))
		}
	}
}

func (s *sgw) close() error {
	var errs []error

	if s.enableKernGTP {
		for _, r := range s.addedRoutes {
			if err := netlink.RouteDel(r); err != nil {
				errs = append(errs, err)
			}
		}
		for _, r := range s.addedRules {
			if err := netlink.RuleDel(r); err != nil {
				errs = append(errs, err)
			}
		}
	}

	if s.s1uConn != nil {
		if err := s.s1uConn.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if s.s5uConn != nil {
		if err := s.s5uConn.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if s.s11Conn != nil {
		if err := s.s11Conn.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if s.s5cConn != nil {
		if err := s.s5cConn.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	close(s.errCh)

	if len(errs) > 0 {
		return errors.Errorf("errors while closing S-GW: %v", errs)
	}
	return nil
}

func (s *sgw) addRoutes() error {
	defnet := &net.IPNet{IP: net.IPv4zero, Mask: net.CIDRMask(0, 32)}
	s1route := &netlink.Route{ // ip route replace
		Dst:       defnet,                  // default
		LinkIndex: s.s5uConn.GTPLink.Index, // dev gtp-s5
		Scope:     netlink.SCOPE_LINK,      // scope link
		Protocol:  4,                       // proto static
		Priority:  1,                       // metric 1
		Table:     2001,                    // table 2001
	}

	if err := netlink.RouteReplace(s1route); err != nil {
		return err
	}
	s.addedRoutes = append(s.addedRoutes, s1route)

	s5route := &netlink.Route{ // ip route replace
		Dst:       defnet,                          // default
		LinkIndex: s.s1uConn.GTPLink.Attrs().Index, // dev gtp-s1
		Scope:     netlink.SCOPE_LINK,              // scope link
		Protocol:  4,                               // proto static
		Priority:  1,                               // metric 1
		Table:     2005,                            // table 2005
	}

	if err := netlink.RouteReplace(s5route); err != nil {
		return err
	}
	s.addedRoutes = append(s.addedRoutes, s1route)

	rules, err := netlink.RuleList(0)
	if err != nil {
		return err
	}

	var s1found, s5found bool
	for _, r := range rules {
		if s1found && s5found {
			break
		}

		if r.IifName == s.s1uConn.GTPLink.Name && r.Table == 2001 {
			s1found = true
		}
		if r.IifName == s.s5uConn.GTPLink.Name && r.Table == 2005 {
			s5found = true
		}
	}

	if !s1found {
		rule := netlink.NewRule()
		rule.IifName = s.s1uConn.GTPLink.Name
		rule.Table = 2001

		if err := netlink.RuleAdd(rule); err != nil {
			return err
		}
		s.addedRules = append(s.addedRules, rule)
	}

	if !s5found {
		rule := netlink.NewRule()
		rule.IifName = s.s5uConn.GTPLink.Name
		rule.Table = 2005

		if err := netlink.RuleAdd(rule); err != nil {
			return err
		}
		s.addedRules = append(s.addedRules, rule)
	}

	return nil
}
