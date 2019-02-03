// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v1_test

import (
	"net"
	"testing"

	v1 "github.com/wmnsk/go-gtp/gtp/v1"
)

func TestRelay(t *testing.T) {
	leftAddr, err := net.ResolveUDPAddr("udp", "127.0.0.11:2152")
	if err != nil {
		t.Fatal(err)
	}
	rightAddr, err := net.ResolveUDPAddr("udp", "127.0.0.12:2152")
	if err != nil {
		t.Fatal(err)
	}

	errCh := make(chan error)
	leftConn, err := v1.ListenAndServeUPlane(leftAddr, 0, errCh)
	if err != nil {
		t.Fatal(err)
	}
	rightConn, err := v1.ListenAndServeUPlane(rightAddr, 0, errCh)
	if err != nil {
		t.Fatal(err)
	}

	// XXX - should be updated
	relay := v1.NewRelay(leftConn, rightConn)
	leftTEID, rightTEID := uint32(0x22222222), uint32(0x11111111)
	relay.AddPeer(leftTEID, rightTEID, rightAddr)
	relay.AddPeer(rightTEID, leftTEID, leftAddr)
	go relay.Run()
	defer relay.Close()

	if _, err := leftConn.WriteToGTP(rightTEID, []byte{0xde, 0xad, 0xbe, 0xef}, rightAddr); err != nil {
		t.Fatal(err)
	}
}
