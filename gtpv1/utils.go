// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtpv1

import "github.com/wmnsk/go-gtp/gtpv1/message"

// Encapsulate encapsulates given payload with GTPv1-U Header and returns message.TPDU.
func Encapsulate(teid uint32, payload []byte) *message.TPDU {
	pdu := message.NewTPDU(teid, payload)
	return pdu
}

// EncapsulateWithExtensionHeader encapsulates given payload and Extension Headers
// with GTPv1-U Header and returns message.TPDU.
func EncapsulateWithExtensionHeader(teid uint32, payload []byte, extHdrs ...*message.ExtensionHeader) *message.TPDU {
	pdu := message.NewTPDUWithExtentionHeader(teid, payload, extHdrs...)
	return pdu
}

// Decapsulate decapsulates given bytes and returns TEID and Payload.
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

// DecapsulateWithExtensionHeader decapsulates given bytes and returns TEID,
// Payload, and Extension Headers.
// It is always safe to use this even if there may be no Extension Headers.
func DecapsulateWithExtensionHeader(b []byte) (uint32, []byte, []*message.ExtensionHeader, error) {
	h, err := message.ParseHeader(b)
	if err != nil {
		return 0, nil, nil, err
	}

	if h.Type != message.MsgTypeTPDU {
		return 0, nil, nil, message.ErrInvalidMessageType
	}
	return h.TEID, h.Payload, h.ExtensionHeaders, nil
}
