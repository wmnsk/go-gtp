// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v1

import (
	"errors"
	"net"
)

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

	u.iteiMap.delete(teidIn)
	return nil
}
