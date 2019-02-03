// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v2

import (
	"net"
	"sync"

	"github.com/wmnsk/go-gtp/gtp/v2/ies"
	"github.com/wmnsk/go-gtp/gtp/v2/messages"
)

// HandlerFunc is a handler for specific GTPv2-C message.
type HandlerFunc func(c *Conn, senderAddr net.Addr, msg messages.Message) error

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
		messages.MsgTypeEchoRequest:                   handleEchoRequest,
		messages.MsgTypeEchoResponse:                  handleEchoResponse,
		messages.MsgTypeVersionNotSupportedIndication: handleVersionNotSupportedIndication,
	},
)

func handleEchoRequest(c *Conn, senderAddr net.Addr, msg messages.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	if _, ok := msg.(*messages.EchoRequest); !ok {
		return ErrUnexpectedType
	}

	// respond with EchoResponse.
	if err := c.RespondTo(
		senderAddr,
		msg, messages.NewEchoResponse(0, ies.NewRecovery(c.RestartCounter)),
	); err != nil {
		return err
	}

	return nil
}

func handleEchoResponse(c *Conn, senderAddr net.Addr, msg messages.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	if _, ok := msg.(*messages.EchoResponse); !ok {
		return ErrUnexpectedType
	}

	// do nothing.
	return nil
}

func handleVersionNotSupportedIndication(c *Conn, senderAddr net.Addr, msg messages.Message) error {
	// this should never happen, as the type should have been assured by
	// msgHandlerMap before this function is called.
	if _, ok := msg.(*messages.VersionNotSupportedIndication); !ok {
		return ErrUnexpectedType
	}

	// let's just return err anyway.
	return ErrInvalidVersion
}
