// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtpv2

import (
	"net"
	"sync"

	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
)

// HandlerFunc is a handler for specific GTPv2-C message.
type HandlerFunc func(c *Conn, senderAddr net.Addr, msg message.Message) error

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
			message.MsgTypeEchoRequest:                   handleEchoRequest,
			message.MsgTypeEchoResponse:                  handleEchoResponse,
			message.MsgTypeVersionNotSupportedIndication: handleVersionNotSupportedIndication,
		},
	)
}

func handleEchoRequest(c *Conn, senderAddr net.Addr, msg message.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	if _, ok := msg.(*message.EchoRequest); !ok {
		return &UnexpectedTypeError{Msg: msg}
	}

	// respond with EchoResponse.
	return c.RespondTo(
		senderAddr, msg, message.NewEchoResponse(0, ie.NewRecovery(c.RestartCounter)),
	)
}

func handleEchoResponse(c *Conn, senderAddr net.Addr, msg message.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	if _, ok := msg.(*message.EchoResponse); !ok {
		return &UnexpectedTypeError{Msg: msg}
	}

	// do nothing.
	return nil
}

func handleVersionNotSupportedIndication(c *Conn, senderAddr net.Addr, msg message.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	if _, ok := msg.(*message.VersionNotSupportedIndication); !ok {
		return &UnexpectedTypeError{Msg: msg}
	}

	// let's just return err anyway.
	return &InvalidVersionError{Version: msg.Version()}
}
