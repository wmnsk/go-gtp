// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v2_test

import (
	"net"
	"testing"
	"time"

	"github.com/pkg/errors"
	v2 "github.com/wmnsk/go-gtp/v2"
	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
)

func setup(doneCh chan struct{}, errCh chan error) (cliConn, srvConn *v2.Conn, err error) {
	cliAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2123")
	if err != nil {
		return nil, nil, err
	}
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.2:2123")
	if err != nil {
		return nil, nil, err
	}

	connCh := make(chan struct{})
	fatalCh := make(chan error)
	go func() {
		srvConn, err = v2.ListenAndServe(srvAddr, 0, errCh)
		if err != nil {
			fatalCh <- err
			return
		}
		srvConn.AddHandler(
			messages.MsgTypeCreateSessionRequest,
			func(c *v2.Conn, cliAddr net.Addr, msg messages.Message) error {
				csReq := msg.(*messages.CreateSessionRequest)
				if imsi := csReq.IMSI.IMSI(); imsi != "123451234567890" {
					return errors.Errorf("unexpected IMSI: %s", imsi)
				}

				csRsp := messages.NewCreateSessionResponse(0, 0, ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil))
				if err := c.RespondTo(cliAddr, csReq, csRsp); err != nil {
					return err
				}
				doneCh <- struct{}{}
				return nil
			},
		)
		connCh <- struct{}{}
	}()

	// XXX - waiting for server to be well-prepared, should consider better way.
	time.Sleep(1 * time.Second)
	cliConn, err = v2.Dial(cliAddr, srvAddr, 0, errCh)
	if err != nil {
		return nil, nil, err
	}

	select {
	case <-connCh:
		return cliConn, srvConn, nil
	case err := <-fatalCh:
		return nil, nil, err
	case <-time.After(1 * time.Second):
		return nil, nil, err
	}
}

func TestCreateSession(t *testing.T) {
	var (
		rspSent = make(chan struct{})
		rspOK   = make(chan struct{})
		errCh   = make(chan error)
	)
	cliConn, srvConn, err := setup(rspSent, errCh)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { cliConn.Close(); srvConn.Close() }()

	cliConn.AddHandler(
		messages.MsgTypeCreateSessionResponse,
		func(c *v2.Conn, srvAddr net.Addr, msg messages.Message) error {
			if srvAddr.String() != "127.0.0.2:2123" {
				t.Fatal("invalid server address")
			}

			csRsp := msg.(*messages.CreateSessionResponse)
			if causeIE := csRsp.Cause; causeIE != nil {
				if cause := causeIE.Cause(); cause != v2.CauseRequestAccepted {
					return &v2.ErrCauseNotOK{
						MsgType: csRsp.MessageTypeName(),
						Cause:   cause,
						Msg:     "something went wrong",
					}
				}
				rspOK <- struct{}{}
			} else {
				return &v2.ErrRequiredIEMissing{Type: ies.Cause}
			}

			return nil
		},
	)

	sess, err := cliConn.CreateSession(srvConn.LocalAddr(), ies.NewIMSI("123451234567890"))
	if err != nil {
		t.Fatal(err)
	}
	cliConn.AddSession(sess)

	select {
	case <-rspSent:
		select {
		case <-rspOK:
			return
		case <-time.After(1 * time.Second):
			t.Fatal("timed out while waiting for validating Create Session Response")
		}
	case err := <-errCh:
		t.Fatal(err)
	case <-time.After(1 * time.Second):
		t.Fatal("timed out while waiting for Create Session Response")
	}
}
