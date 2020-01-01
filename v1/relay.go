// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v1

import (
	"log"
	"net"
	"sync"
	"time"
)

// Relay is to relay packets between two UPlaneConn.
//
// DEPRECATED. Use UPlaneConn.RelayTo() instead.
type Relay struct {
	mu                  sync.Mutex
	closeCh             chan struct{}
	leftConn, rightConn *UPlaneConn
	teidPair            map[uint32]*peer
}

// NewRelay creates a new Relay.
//
// DEPRECATED. Use UPlaneConn.RelayTo() instead.
func NewRelay(leftConn, rightConn *UPlaneConn) *Relay {
	log.Println("Relay is deprecated. Use UPlaneConn.RelayTo() instead.")
	return &Relay{
		mu:        sync.Mutex{},
		closeCh:   make(chan struct{}),
		leftConn:  leftConn,
		rightConn: rightConn,
		teidPair:  map[uint32]*peer{},
	}
}

// Run starts listening on both UPlaneConn.
// Until peer information is registered by AddPeer(), it just drops packets.
//
// DEPRECATED. Use UPlaneConn.RelayTo() instead.
func (r *Relay) Run() {
	// from left to right
	go func() {
		buf := make([]byte, 1600)
		for {
			select {
			case <-r.closed():
				return
			default:
				// do nothing and go forward.
			}

			n, _, teid, err := r.leftConn.ReadFromGTP(buf)
			if err != nil {
				continue
			}

			peer, ok := r.getPeer(teid)
			if !ok {
				continue
			}
			if _, err := r.rightConn.WriteToGTP(peer.teid, buf[:n], peer.addr); err != nil {
				continue
			}
		}
	}()

	// from right to left
	go func() {
		buf := make([]byte, 1600)
		for {
			select {
			case <-r.closed():
				return
			default:
				// do nothing and go forward.
			}

			n, _, teid, err := r.rightConn.ReadFromGTP(buf)
			if err != nil {
				continue
			}

			peer, ok := r.getPeer(teid)
			if !ok {
				continue
			}
			if _, err := r.leftConn.WriteToGTP(peer.teid, buf[:n], peer.addr); err != nil {
				continue
			}
		}
	}()
}

// Close closes Relay. It does not close the UPlaneConn given at first.
//
// DEPRECATED. Use UPlaneConn.RelayTo() instead.
func (r *Relay) Close() error {
	if err := r.leftConn.SetReadDeadline(time.Now().Add(time.Duration(1 * time.Millisecond))); err != nil {
		return err
	}
	if err := r.rightConn.SetReadDeadline(time.Now().Add(time.Duration(1 * time.Millisecond))); err != nil {
		return err
	}
	close(r.closeCh)
	return nil
}

func (r *Relay) closed() <-chan struct{} {
	return r.closeCh
}

// AddPeer adds a peer information with the TEID contained in the incoming meesage.
//
// DEPRECATED. Use UPlaneConn.RelayTo() instead.
func (r *Relay) AddPeer(teidIn, teidOut uint32, raddr net.Addr) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.teidPair[teidIn] = &peer{teid: teidOut, addr: raddr}
}

func (r *Relay) getPeer(teid uint32) (*peer, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, ok := r.teidPair[teid]
	return p, ok
}
