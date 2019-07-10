package v1

import (
	"net"

	"github.com/pkg/errors"
	"github.com/vishvananda/netlink"
)

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
		netlink.GTPPDPDel(u.GTPLink, pdp)
	}
	if pdp, _ := netlink.GTPPDPByITEI(u.GTPLink, int(itei)); pdp != nil {
		netlink.GTPPDPDel(u.GTPLink, pdp)
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

	if err := netlink.GTPPDPDel(u.GTPLink, pdp); err != nil {
		return errors.Wrapf(err, "failed to delete tunnel for %s", pdp)
	}
	return nil
}

type peer struct {
	teid    uint32
	addr    net.Addr
	srcConn *UPlaneConn
}

// RelayTo relays T-PDU type of packet to peer node(specified by raddr) from the UPlaneConn given.
//
// By using this, owner of UPlaneConn won't be able to Read and Write the packets that has teidIn.
func (u *UPlaneConn) RelayTo(c *UPlaneConn, teidIn, teidOut uint32, raddr net.Addr) error {
	if u.kernGTPEnabled {
		return errors.New("cannot call RelayTo when using Kernel GTP-U")
	}

	u.mu.Lock()
	defer u.mu.Unlock()
	if u.relayMap == nil {
		u.relayMap = map[uint32]*peer{}
	}
	u.relayMap[teidIn] = &peer{teid: teidOut, addr: raddr, srcConn: c}
	return nil
}

// CloseRelay stops relaying T-PDU from a conn to conn.
func (u *UPlaneConn) CloseRelay(teidIn uint32) error {
	if u.kernGTPEnabled {
		return errors.New("cannot call CloseRelay when using Kernel GTP-U")
	}

	u.mu.Lock()
	delete(u.relayMap, teidIn)
	u.mu.Unlock()
	return nil
}
