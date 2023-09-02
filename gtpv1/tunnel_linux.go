// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtpv1

import (
	"errors"
	"fmt"
	"net"

	"github.com/vishvananda/netlink"
)

// Role is a role for Kernel GTP-U.
type Role int

// Role definitions.
const (
	RoleGGSN Role = iota
	RoleSGSN
)

// EnableKernelGTP enables Linux Kernel GTP-U.
// Note that this removes all the existing userland tunnels, and cannot be disabled while
// the program is working (at least at this moment).
//
// Using Kernel GTP-U is much more performant than userland, but requires root privilege.
// After enabled, users should add tunnels by AddTunnel func, and also add appropriate
// routing entries. For handling downlink traffic on P-GW, for example;
//
// ip route add <UE's IP> dev <devname> table <table ID>
// ip rule add from <SGi side of I/F> lookup <table ID>
//
// This let the traffic from SGi side of network I/F to be forwarded to GTP device,
// and if the UE's IP is known to Kernel GTP-U(by AddTunnel), it is encapsulated and
// forwarded to the peer(S-GW, in this case).
//
// Please see the examples/gw-tester for how each node handles routing from the program.
func (u *UPlaneConn) EnableKernelGTP(devname string, role Role) error {
	if u.pktConn == nil {
		var err error
		u.pktConn, err = newPktConn(u.laddr)
		if err != nil {
			return err
		}
	}

	f, err := u.pktConn.File()
	if err != nil {
		return fmt.Errorf("failed to retrieve file from conn: %w", err)
	}

	u.KernelGTP.Link = &netlink.GTP{
		LinkAttrs: netlink.LinkAttrs{
			Name: devname,
		},
		FD1:  int(f.Fd()),
		Role: int(role),
	}

	if err := netlink.LinkAdd(u.KernelGTP.Link); err != nil {
		_ = f.Close()
		return fmt.Errorf("failed to add device %s: %w", u.KernelGTP.Link.Name, err)
	}
	if err := netlink.LinkSetUp(u.KernelGTP.Link); err != nil {
		_ = f.Close()
		return fmt.Errorf("failed to setup device %s: %w", u.KernelGTP.Link.Name, err)
	}
	if err := netlink.LinkSetMTU(u.KernelGTP.Link, 1500); err != nil {
		_ = f.Close()
		return fmt.Errorf("failed to set MTU for device %s: %w", u.KernelGTP.Link.Name, err)
	}
	u.KernelGTP.connFile = f
	u.KernelGTP.enabled = true

	// remove relayed userland tunnels if exists
	if len(u.relayMap) != 0 {
		u.mu.Lock()
		u.relayMap = nil
		u.mu.Unlock()
	}

	return nil
}

// AddTunnel adds a GTP-U tunnel with Linux Kernel GTP-U via netlink.
func (u *UPlaneConn) AddTunnel(peerIP, msIP net.IP, otei, itei uint32) error {
	if !u.KernelGTP.enabled {
		return errors.New("cannot call AddTunnel when not using Kernel GTP-U")
	}

	pdp := &netlink.PDP{
		Version:     1,
		PeerAddress: peerIP,
		MSAddress:   msIP,
		OTEI:        otei,
		ITEI:        itei,
	}
	if err := netlink.GTPPDPAdd(u.KernelGTP.Link, pdp); err != nil {
		return fmt.Errorf("failed to add tunnel for %s with %s: %w", msIP, peerIP, err)
	}
	return nil
}

// AddTunnelOverride adds a GTP-U tunnel with Linux Kernel GTP-U via netlink.
// If there is already an existing tunnel that has the same msIP and/or incoming TEID,
// this deletes it before adding the tunnel.
func (u *UPlaneConn) AddTunnelOverride(peerIP, msIP net.IP, otei, itei uint32) error {
	if !u.KernelGTP.enabled {
		return errors.New("cannot call AddTunnelOverride when not using Kernel GTP-U")
	}

	if pdp, _ := netlink.GTPPDPByMSAddress(u.KernelGTP.Link, msIP); pdp != nil {
		// do nothing even this fails
		_ = netlink.GTPPDPDel(u.KernelGTP.Link, pdp)
	}
	if pdp, _ := netlink.GTPPDPByITEI(u.KernelGTP.Link, int(itei)); pdp != nil {
		// do nothing even this fails
		_ = netlink.GTPPDPDel(u.KernelGTP.Link, pdp)
	}

	return u.AddTunnel(peerIP, msIP, otei, itei)
}

// DelTunnelByITEI deletes a Linux Kernel GTP-U tunnel specified with the incoming TEID.
func (u *UPlaneConn) DelTunnelByITEI(itei uint32) error {
	if !u.KernelGTP.enabled {
		return errors.New("cannot call DelTunnel when not using Kernel GTP-U")
	}

	pdp, err := netlink.GTPPDPByITEI(u.KernelGTP.Link, int(itei))
	if err != nil {
		return fmt.Errorf("failed to delete tunnel with %d: %w", itei, err)
	}

	if err := netlink.GTPPDPDel(u.KernelGTP.Link, pdp); err != nil {
		return fmt.Errorf("failed to delete tunnel for %s: %w", pdp, err)
	}

	u.iteiMap.delete(itei)
	return nil
}

// DelTunnelByMSAddress deletes a Linux Kernel GTP-U tunnel specified with the subscriber's IP.
func (u *UPlaneConn) DelTunnelByMSAddress(msIP net.IP) error {
	if !u.KernelGTP.enabled {
		return errors.New("cannot call DelTunnel when not using Kernel GTP-U")
	}

	pdp, err := netlink.GTPPDPByMSAddress(u.KernelGTP.Link, msIP)
	if err != nil {
		return fmt.Errorf("failed to delete tunnel with %s: %w", msIP, err)
	}
	itei := pdp.ITEI

	if err := netlink.GTPPDPDel(u.KernelGTP.Link, pdp); err != nil {
		return fmt.Errorf("failed to delete tunnel for %s: %w", pdp, err)
	}

	u.iteiMap.delete(itei)
	return nil
}
