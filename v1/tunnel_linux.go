// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v1

import (
	"net"

	"github.com/pkg/errors"
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
//  ip route add <UE's IP> dev <devname> table <table ID>
//  ip rule add from <SGi side of I/F> lookup <table ID>
//
// This let the traffic from SGi side of network I/F to be forwarded to GTP device,
// and if the UE's IP is known to Kernel GTP-U(by AddTunnel), it is encapsulated and
// forwarded to the peer(S-GW, in this case).
//
// Please see the examples/gw-tester for how each node handles routing from the program.
func (u *UPlaneConn) EnableKernelGTP(devname string, role Role) error {
	if u.pktConn == nil {
		var err error
		u.pktConn, err = net.ListenPacket(u.laddr.Network(), u.laddr.String())
		if err != nil {
			return err
		}
	}

	f, err := u.pktConn.(*net.UDPConn).File()
	if err != nil {
		return errors.Wrapf(err, "failed to retrieve file from conn")
	}

	u.GTPLink = &netlink.GTP{
		LinkAttrs: netlink.LinkAttrs{
			Name: devname,
		},
		FD1:  int(f.Fd()),
		Role: int(role),
	}

	if err := netlink.LinkAdd(u.GTPLink); err != nil {
		return errors.Wrapf(err, "failed to add device: %s", u.GTPLink.Name)
	}
	if err := netlink.LinkSetUp(u.GTPLink); err != nil {
		return errors.Wrapf(err, "failed to setup device: %s", u.GTPLink.Name)
	}
	if err := netlink.LinkSetMTU(u.GTPLink, 1500); err != nil {
		return errors.Wrapf(err, "failed to set MTU for device: %s", u.GTPLink.Name)
	}
	u.kernGTPEnabled = true

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
	if !u.kernGTPEnabled {
		return errors.New("cannot call AddTunnel when not using Kernel GTP-U")
	}

	pdp := &netlink.PDP{
		Version:     1,
		PeerAddress: peerIP,
		MSAddress:   msIP,
		OTEI:        otei,
		ITEI:        itei,
	}
	if err := netlink.GTPPDPAdd(u.GTPLink, pdp); err != nil {
		return errors.Wrapf(err, "failed to add tunnel for %s with %s", msIP, peerIP)
	}
	return nil
}

// AddTunnelOverride adds a GTP-U tunnel with Linux Kernel GTP-U via netlink.
// If there is already an existing tunnel that has the same msIP and/or incoming TEID,
// this deletes it before adding the tunnel.
func (u *UPlaneConn) AddTunnelOverride(peerIP, msIP net.IP, otei, itei uint32) error {
	if !u.kernGTPEnabled {
		return errors.New("cannot call AddTunnelOverride when not using Kernel GTP-U")
	}

	if pdp, _ := netlink.GTPPDPByMSAddress(u.GTPLink, msIP); pdp != nil {
		// do nothing even this fails
		_ = netlink.GTPPDPDel(u.GTPLink, pdp)
	}
	if pdp, _ := netlink.GTPPDPByITEI(u.GTPLink, int(itei)); pdp != nil {
		// do nothing even this fails
		_ = netlink.GTPPDPDel(u.GTPLink, pdp)
	}

	return u.AddTunnel(peerIP, msIP, otei, itei)
}

// DelTunnelByITEI deletes a Linux Kernel GTP-U tunnel specified with the incoming TEID.
func (u *UPlaneConn) DelTunnelByITEI(itei uint32) error {
	if !u.kernGTPEnabled {
		return errors.New("cannot call DelTunnel when not using Kernel GTP-U")
	}

	pdp, err := netlink.GTPPDPByITEI(u.GTPLink, int(itei))
	if err != nil {
		return errors.Wrapf(err, "failed to delete tunnel with %d", itei)
	}

	if err := netlink.GTPPDPDel(u.GTPLink, pdp); err != nil {
		return errors.Wrapf(err, "failed to delete tunnel for %s", pdp)
	}

	u.iteiMap.delete(itei)
	return nil
}

// DelTunnelByMSAddress deletes a Linux Kernel GTP-U tunnel specified with the subscriber's IP.
func (u *UPlaneConn) DelTunnelByMSAddress(msIP net.IP) error {
	if !u.kernGTPEnabled {
		return errors.New("cannot call DelTunnel when not using Kernel GTP-U")
	}

	pdp, err := netlink.GTPPDPByMSAddress(u.GTPLink, msIP)
	if err != nil {
		return errors.Wrapf(err, "failed to delete tunnel with %s", msIP)
	}
	itei := pdp.ITEI

	if err := netlink.GTPPDPDel(u.GTPLink, pdp); err != nil {
		return errors.Wrapf(err, "failed to delete tunnel for %s", pdp)
	}

	u.iteiMap.delete(itei)
	return nil
}
