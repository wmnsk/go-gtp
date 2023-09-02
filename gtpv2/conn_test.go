// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtpv2_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
)

func setup(ctx context.Context, doneCh chan struct{}) (cliConn, srvConn *gtpv2.Conn, err error) {
	cliAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1"+gtpv2.GTPCPort)
	if err != nil {
		return nil, nil, err
	}
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.2"+gtpv2.GTPCPort)
	if err != nil {
		return nil, nil, err
	}

	srvCreated := make(chan struct{})
	go func() {
		srvConn = gtpv2.NewConn(srvAddr, gtpv2.IFTypeS11S4SGWGTPC, 0)
		srvConn.AddHandler(
			message.MsgTypeCreateSessionRequest,
			func(c *gtpv2.Conn, cliAddr net.Addr, msg message.Message) error {

				csReq := msg.(*message.CreateSessionRequest)
				session := gtpv2.NewSession(cliAddr, &gtpv2.Subscriber{Location: &gtpv2.Location{}})

				var otei uint32
				if imsiIE := csReq.IMSI; imsiIE != nil {
					imsi, err := imsiIE.IMSI()
					if err != nil {
						return err
					}
					if imsi != "123451234567890" {
						return fmt.Errorf("unexpected IMSI: %s", imsi)
					}
					session.IMSI = imsi
				} else {
					return &gtpv2.RequiredIEMissingError{Type: ie.IMSI}
				}

				if fteidcIE := csReq.SenderFTEIDC; fteidcIE != nil {
					ip, err := fteidcIE.IPAddress()
					if err != nil {
						return err
					}
					if ip != "127.0.0.1" {
						return fmt.Errorf("unexpected IP in F-TEID: %s", ip)
					}

					ifType, err := fteidcIE.InterfaceType()
					if err != nil {
						return err
					}
					otei, err = fteidcIE.TEID()
					if err != nil {
						return err
					}
					session.AddTEID(ifType, otei)
				} else {
					return &gtpv2.RequiredIEMissingError{Type: ie.IMSI}
				}

				fTEID := srvConn.NewSenderFTEID("127.0.0.2", "")
				srvConn.RegisterSession(fTEID.MustTEID(), session)
				csRsp := message.NewCreateSessionResponse(
					otei, msg.Sequence(), ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil), fTEID,
				)
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

		if err := srvConn.Listen(ctx); err != nil {
			log.Println(err)
			return
		}
		srvCreated <- struct{}{}
		if err := srvConn.Serve(ctx); err != nil {
			log.Println(err)
			return
		}
	}()

	select {
	case <-srvCreated:
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout waiting for server creation")
	}
	cliConn, err = gtpv2.Dial(ctx, cliAddr, srvAddr, gtpv2.IFTypeS11MMEGTPC, 0)
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
		message.MsgTypeCreateSessionResponse,
		func(c *gtpv2.Conn, srvAddr net.Addr, msg message.Message) error {
			if srvAddr.String() != "127.0.0.2"+gtpv2.GTPCPort {
				t.Errorf("invalid server address: %s", srvAddr)
			}
			if msg.Sequence() != cliConn.SequenceNumber() {
				t.Errorf("invalid sequence number. got: %d, want: %d", msg.Sequence(), cliConn.SequenceNumber())
			}

			session, err := c.GetSessionByTEID(msg.TEID(), srvAddr)
			if err != nil {
				return err
			}

			csRsp := msg.(*message.CreateSessionResponse)
			if causeIE := csRsp.Cause; causeIE != nil {
				cause, err := causeIE.Cause()
				if err != nil {
					return err
				}
				if cause != gtpv2.CauseRequestAccepted {
					return &gtpv2.CauseNotOKError{
						MsgType: csRsp.MessageTypeName(),
						Cause:   cause,
						Msg:     "something went wrong",
					}
				}
			} else {
				return &gtpv2.RequiredIEMissingError{Type: ie.Cause}
			}

			if fteidIE := csRsp.SenderFTEIDC; fteidIE != nil {
				it, err := fteidIE.InterfaceType()
				if err != nil {
					return err
				}
				if it != gtpv2.IFTypeS11S4SGWGTPC {
					return fmt.Errorf("invalid InterfaceType: %v", it)
				}
				otei, err := fteidIE.TEID()
				if err != nil {
					return err
				}
				session.AddTEID(it, otei)

				ip, err := fteidIE.IPAddress()
				if err != nil {
					return err
				}
				if ip != "127.0.0.2" {
					return fmt.Errorf("unexpected IP in F-TEID: %s", ip)
				}
			} else {
				return &gtpv2.RequiredIEMissingError{Type: ie.Cause}
			}

			if err := session.Activate(); err != nil {
				return err
			}
			rspOK <- struct{}{}
			return nil
		},
	)

	fTEID := cliConn.NewSenderFTEID("127.0.0.1", "")
	_, _, err = cliConn.CreateSession(srvConn.LocalAddr(), ie.NewIMSI("123451234567890"), fTEID)
	if err != nil {
		t.Fatal(err)
	}

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
	case <-time.After(5 * time.Second):
		t.Fatal("timed out while waiting for validating Create Session Response")
	}
}
