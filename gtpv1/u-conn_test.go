// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtpv1_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/wmnsk/go-gtp/gtpv1"
)

type testVal struct {
	teidIn, teidOut uint32
	seq             uint16
	payload         []byte
}

func setup(ctx context.Context) (cliConn, srvConn *gtpv1.UPlaneConn, err error) {
	cliAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2152")
	if err != nil {
		return nil, nil, err
	}
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.2:2152")
	if err != nil {
		return nil, nil, err
	}

	go func() {
		srvConn = gtpv1.NewUPlaneConn(srvAddr)
		srvConn.DisableErrorIndication()
		if err := srvConn.ListenAndServe(ctx); err != nil {
			return
		}
	}()

	// XXX - waiting for server to be well-prepared, should consider better way.
	time.Sleep(1 * time.Second)
	cliConn, err = gtpv1.DialUPlane(ctx, cliAddr, srvAddr)
	if err != nil {
		return nil, nil, err
	}
	cliConn.DisableErrorIndication()

	return cliConn, srvConn, nil
}

func TestClientWrite(t *testing.T) {
	var (
		okCh  = make(chan struct{})
		errCh = make(chan error)
		buf   = make([]byte, 2048)
		tv    = &testVal{
			0x11111111, 0x22222222, 0x3333,
			[]byte{0xde, 0xad, 0xbe, 0xef},
		}
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cliConn, srvConn, err := setup(ctx)
	if err != nil {
		t.Fatal(err)
	}

	go func(tv *testVal) {
		n, addr, teid, err := srvConn.ReadFromGTP(buf)
		if err != nil {
			errCh <- err
			return
		}

		if diff := cmp.Diff(n, len(tv.payload)); diff != "" {
			t.Error(diff)
		}
		if diff := cmp.Diff(addr, cliConn.LocalAddr()); diff != "" {
			t.Error(diff)
		}
		if diff := cmp.Diff(teid, tv.teidOut); diff != "" {
			t.Error(diff)
		}
		if diff := cmp.Diff(buf[:n], tv.payload); diff != "" {
			t.Error(diff)
		}
		okCh <- struct{}{}
	}(tv)

	if _, err := cliConn.WriteToGTP(tv.teidOut, tv.payload, srvConn.LocalAddr()); err != nil {
		t.Fatal(err)
	}

	select {
	case <-okCh:
		return
	case err := <-errCh:
		t.Fatal(err)
	case <-time.After(10 * time.Second):
		t.Fatal("timed out while waiting for response to come")
	}
}
