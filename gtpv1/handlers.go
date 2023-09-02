// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtpv1

import (
	"net"
	"sync"
	"time"

	"github.com/wmnsk/go-gtp/gtpv1/ie"
	"github.com/wmnsk/go-gtp/gtpv1/message"
)

// HandlerFunc is a handler for specific GTPv1 message.
type HandlerFunc func(c Conn, senderAddr net.Addr, msg message.Message) error

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

func newDefaultMsgHandlerMap() *msgHandlerMap {
	return newMsgHandlerMap(
		map[uint8]HandlerFunc{
			message.MsgTypeTPDU:            handleTPDU,
			message.MsgTypeEchoRequest:     handleEchoRequest,
			message.MsgTypeEchoResponse:    handleEchoResponse,
			message.MsgTypeErrorIndication: handleErrorIndication,
		},
	)
}

// handleTPDU responds to sender with ErrorIndication by default.
// By disabling it(DisableErrorIndication), it passes unhandled T-PDU to
// user, which can be caught by calling ReadFromGTP.
func handleTPDU(c Conn, senderAddr net.Addr, msg message.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	pdu, ok := msg.(*message.TPDU)
	if !ok {
		return ErrUnexpectedType
	}

	u, ok := c.(*UPlaneConn)
	if !ok {
		return ErrInvalidConnection
	}

	if u.errIndEnabled {
		if err := u.ErrorIndication(senderAddr, pdu); err != nil {
			logf("failed to send Error Indication to %s: %v", senderAddr, err)
		}
		return nil
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

func handleEchoRequest(c Conn, senderAddr net.Addr, msg message.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	if _, ok := msg.(*message.EchoRequest); !ok {
		return ErrUnexpectedType
	}

	// respond with EchoResponse.
	return c.RespondTo(
		senderAddr, msg, message.NewEchoResponse(0, ie.NewRecovery(c.Restarts())),
	)
}

func handleEchoResponse(c Conn, senderAddr net.Addr, msg message.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	if _, ok := msg.(*message.EchoResponse); !ok {
		return ErrUnexpectedType
	}

	// do nothing.
	return nil
}

func handleErrorIndication(c Conn, senderAddr net.Addr, msg message.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	ind, ok := msg.(*message.ErrorIndication)
	if !ok {
		return ErrUnexpectedType
	}

	// just log and return
	logf("Ignored Error Indication: %v", &ErrorIndicatedError{
		TEID: ind.TEIDDataI.MustTEID(),
		Peer: ind.GTPUPeerAddress.MustIPAddress(),
	})
	return nil
}
