// Copyright 2019-2022 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtpv1

import "github.com/wmnsk/go-gtp/gtpv1/message"

// Encapsulate encapsulates given bytes with GTPv1-U Header and returns in message.TPDU.
func Encapsulate(teid uint32, payload []byte) *message.TPDU {
	pdu := message.NewTPDU(teid, payload)
	return pdu
}

// Decapsulate decapsulates given bytes and returns TEID, and Payload.
func Decapsulate(b []byte) (uint32, []byte, error) {
	header, err := message.ParseHeader(b)
	if err != nil {
		return 0, nil, err
	}

	if header.Type != message.MsgTypeTPDU {
		return 0, nil, message.ErrInvalidMessageType
	}
	return header.TEID, header.Payload, nil
}
