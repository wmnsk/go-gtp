// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v1_test

import (
	"net"
	"testing"

	v1 "github.com/wmnsk/go-gtp/v1"
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
	defer leftConn.Close()
	rightConn, err := v1.ListenAndServeUPlane(rightAddr, 0, errCh)
	if err != nil {
		t.Fatal(err)
	}
	defer rightConn.Close()

	if err := leftConn.RelayTo(rightConn, 0x22222222, 0x11111111, rightAddr); err != nil {
		t.Fatal(err)
	}
	if err := rightConn.RelayTo(leftConn, 0x11111111, 0x22222222, leftAddr); err != nil {
		t.Fatal(err)
	}

	// TODO: add tests to check if the traffic goes through conns.
}
