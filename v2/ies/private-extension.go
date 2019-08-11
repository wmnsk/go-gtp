// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "encoding/binary"

// NewPrivateExtension creates a new PrivateExtension IE.
func NewPrivateExtension(id uint16, value []byte) *IE {
	i := New(PrivateExtension, 0x00, make([]byte, 2+len(value)))
	binary.BigEndian.PutUint16(i.Payload[0:2], id)
	copy(i.Payload[2:], value)
	return i
}

// EnterpriseID returns EnterpriseID in uint16 if the type of IE matches.
func (i *IE) EnterpriseID() uint16 {
	if i.Type != PrivateExtension {
		return 0
	}
	return binary.BigEndian.Uint16(i.Payload[0:2])

}

// PrivateExtension returns PrivateExtension value in []byte if the type of IE matches.
func (i *IE) PrivateExtension() []byte {
	if i.Type != PrivateExtension {
		return nil
	}
	return i.Payload[2:]
}
