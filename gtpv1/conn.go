// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtpv1

import (
	"net"

	"github.com/wmnsk/go-gtp/gtpv1/message"
)

// Conn is an abstraction of both GTPv1-C and GTPv1-U Conn.
type Conn interface {
	net.PacketConn
	AddHandler(uint8, HandlerFunc)
	RespondTo(net.Addr, message.Message, message.Message) error
	Restarts() uint8
}
