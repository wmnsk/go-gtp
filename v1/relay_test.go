// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v1_test

import (
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	leftConn := v1.NewUPlaneConn(leftAddr)
	go func() {
		if err := leftConn.ListenAndServe(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	rightConn := v1.NewUPlaneConn(rightAddr)
	go func() {
		if err := rightConn.ListenAndServe(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	if err := leftConn.RelayTo(rightConn, 0x22222222, 0x11111111, rightAddr); err != nil {
		t.Fatal(err)
	}
	if err := rightConn.RelayTo(leftConn, 0x11111111, 0x22222222, leftAddr); err != nil {
		t.Fatal(err)
	}

	// TODO: add tests to check if the traffic goes through conns.
}
