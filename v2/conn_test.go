// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v2_test

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/pkg/errors"
	v2 "github.com/wmnsk/go-gtp/v2"
	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
)

func setup(ctx context.Context, doneCh chan struct{}) (cliConn, srvConn *v2.Conn, err error) {
	cliAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2123")
	if err != nil {
		return nil, nil, err
	}
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.2:2123")
	if err != nil {
		return nil, nil, err
	}

	go func() {
		srvConn = v2.NewConn(srvAddr, 0, []uint8{})
		srvConn.AddHandler(
			messages.MsgTypeCreateSessionRequest,
			func(c *v2.Conn, cliAddr net.Addr, msg messages.Message) error {

				csReq := msg.(*messages.CreateSessionRequest)
				session := v2.NewSession(cliAddr, &v2.Subscriber{Location: &v2.Location{}})

				if ie := csReq.IMSI; ie != nil {
					imsi, err := ie.IMSI()
					if err != nil {
						return err
					}
					if imsi != "123451234567890" {
						return errors.Errorf("unexpected IMSI: %s", imsi)
					}
					session.IMSI = imsi
				}
				c.AddSession(session)

				csRsp := messages.NewCreateSessionResponse(0, 0, ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil))
				if err := c.RespondTo(cliAddr, csReq, csRsp); err != nil {
					return err
				}

				if err := session.Activate(); err != nil {
					return err
				}
				doneCh <- struct{}{}
				return nil
			},
		)

		if err := srvConn.ListenAndServe(ctx); err != nil {
			log.Println(err)
			return
		}
	}()

	// XXX - waiting for server to be well-prepared, should consider better way.
	time.Sleep(1 * time.Second)
	cliConn, err = v2.Dial(ctx, cliAddr, srvAddr, 0, []uint8{})
	if err != nil {
		return nil, nil, err
	}

	return cliConn, srvConn, nil
}

func TestCreateSession(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	doneCh := make(chan struct{})
	rspOK := make(chan struct{})

	cliConn, srvConn, err := setup(ctx, doneCh)
	if err != nil {
		t.Fatal(err)
	}

	cliConn.AddHandler(
		messages.MsgTypeCreateSessionResponse,
		func(c *v2.Conn, srvAddr net.Addr, msg messages.Message) error {
			if srvAddr.String() != "127.0.0.2:2123" {
				t.Errorf("invalid server address: %s", srvAddr)
			}
			if msg.Sequence() != cliConn.SequenceNumber() {
				t.Errorf("invalid sequence number. got: %d, want: %d", msg.Sequence(), cliConn.SequenceNumber())
			}

			// session should be retrieved by msg.TEID() in the real case.
			session, err := c.GetSessionByIMSI("123451234567890")
			if err != nil {
				return err
			}

			csRsp := msg.(*messages.CreateSessionResponse)
			if causeIE := csRsp.Cause; causeIE != nil {
				cause, err := causeIE.Cause()
				if err != nil {
					return err
				}
				if cause != v2.CauseRequestAccepted {
					return &v2.CauseNotOKError{
						MsgType: csRsp.MessageTypeName(),
						Cause:   cause,
						Msg:     "something went wrong",
					}
				}

				if err := session.Activate(); err != nil {
					return err
				}
				rspOK <- struct{}{}
			} else {
				return &v2.RequiredIEMissingError{Type: ies.Cause}
			}

			return nil
		},
	)

	session, _, err := cliConn.CreateSession(srvConn.LocalAddr(), ies.NewIMSI("123451234567890"))
	if err != nil {
		t.Fatal(err)
	}
	cliConn.AddSession(session)

	select {
	case <-rspOK:
		if count := cliConn.SessionCount(); count != 1 {
			t.Errorf("wrong SessionCount in cliConn. want %d, got: %d", 1, count)
		}
		if count := cliConn.BearerCount(); count != 1 {
			t.Errorf("wrong BearerCount in cliConn. want %d, got: %d", 1, count)
		}
		<-doneCh
		if count := srvConn.SessionCount(); count != 1 {
			t.Errorf("wrong SessionCount in srvConn. want %d, got: %d", 1, count)
		}
		if count := srvConn.BearerCount(); count != 1 {
			t.Errorf("wrong BearerCount in srvConn. want %d, got: %d", 1, count)
		}
		return
	case <-time.After(5 * time.Second):
		t.Fatal("timed out while waiting for validating Create Session Response")
	}
}
