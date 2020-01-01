// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v1

import (
	"net"
	"sync"
	"time"

	"github.com/wmnsk/go-gtp/v1/ies"
	"github.com/wmnsk/go-gtp/v1/messages"
)

// HandlerFunc is a handler for specific GTPv1 message.
type HandlerFunc func(c Conn, senderAddr net.Addr, msg messages.Message) error

type msgHandlerMap struct {
	syncMap sync.Map
}

func (m *msgHandlerMap) store(msgType uint8, handler HandlerFunc) {
	m.syncMap.Store(msgType, handler)
}

func (m *msgHandlerMap) load(msgType uint8) (HandlerFunc, bool) {
	handler, ok := m.syncMap.Load(msgType)
	if !ok {
		return nil, false
	}

	return handler.(HandlerFunc), true
}

func newMsgHandlerMap(m map[uint8]HandlerFunc) *msgHandlerMap {
	mhm := &msgHandlerMap{syncMap: sync.Map{}}
	for k, v := range m {
		mhm.store(k, v)
	}

	return mhm
}

var defaultHandlerMap = newMsgHandlerMap(
	map[uint8]HandlerFunc{
		messages.MsgTypeTPDU:            handleTPDU,
		messages.MsgTypeEchoRequest:     handleEchoRequest,
		messages.MsgTypeEchoResponse:    handleEchoResponse,
		messages.MsgTypeErrorIndication: handleErrorIndication,
	},
)

func handleTPDU(c Conn, senderAddr net.Addr, msg messages.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	pdu, ok := msg.(*messages.TPDU)
	if !ok {
		return ErrUnexpectedType
	}

	u, ok := c.(*UPlaneConn)
	if !ok {
		return ErrInvalidConnection
	}

	tpdu := &tpduSet{
		raddr:   senderAddr,
		teid:    pdu.TEID(),
		seq:     pdu.Sequence(),
		payload: pdu.Payload,
	}

	// wait for the T-PDU passed to u.tpduCh to be read by ReadFromGTP.
	// if it got stuck for 3 seconds, it discards the T-PDU received.
	go func() {
		select {
		case u.tpduCh <- tpdu:
			return
		case <-time.After(3 * time.Second):
			return
		}
	}()
	return nil
}

func handleEchoRequest(c Conn, senderAddr net.Addr, msg messages.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	if _, ok := msg.(*messages.EchoRequest); !ok {
		return ErrUnexpectedType
	}

	// respond with EchoResponse.
	return c.RespondTo(
		senderAddr, msg, messages.NewEchoResponse(0, ies.NewRecovery(c.Restarts())),
	)
}

func handleEchoResponse(c Conn, senderAddr net.Addr, msg messages.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	if _, ok := msg.(*messages.EchoResponse); !ok {
		return ErrUnexpectedType
	}

	// do nothing.
	return nil
}

func handleErrorIndication(c Conn, senderAddr net.Addr, msg messages.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	ind, ok := msg.(*messages.ErrorIndication)
	if !ok {
		return ErrUnexpectedType
	}

	// let's just return err anyway.
	return &ErrorIndicatedError{
		TEID: ind.TEIDDataI.MustTEID(),
		Peer: ind.GTPUPeerAddress.MustIPAddress(),
	}
}
