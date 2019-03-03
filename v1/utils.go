// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v1

import "github.com/wmnsk/go-gtp/v1/messages"

// Encapsulate encapsulates given bytes with GTPv1-U Header and returns in message.TPDU.
func Encapsulate(teid uint32, payload []byte) *messages.TPDU {
	pdu := messages.NewTPDU(teid, payload)
	return pdu
}

// Decapsulate decapsulates given bytes and returns TEID, and Payload.
func Decapsulate(b []byte) (uint32, []byte, error) {
	header, err := messages.DecodeHeader(b)
	if err != nil {
		return 0, nil, err
	}

	if header.Type != messages.MsgTypeTPDU {
		return 0, nil, messages.ErrInvalidMessageType
	}
	return header.TEID, header.Payload, nil
}
