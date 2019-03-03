// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
)

// NewPrivateExtension creates a new PrivateExtension IE from string.
func NewPrivateExtension(id uint16, val []byte) *IE {
	i := New(PrivateExtension, make([]byte, 2+len(val)))
	binary.BigEndian.PutUint16(i.Payload[:2], id)
	copy(i.Payload[2:], val)
	return i
}

// PrivateExtension returns PrivateExtension value if type matches.
func (i *IE) PrivateExtension() []byte {
	if i.Type != PrivateExtension {
		return nil
	}
	return i.Payload
}

// ExtensionIdentifier returns ExtensionIdentifier value in uint16 if type matches.
func (i *IE) ExtensionIdentifier() uint16 {
	if i.Type != PrivateExtension {
		return 0
	}
	return binary.BigEndian.Uint16(i.Payload[:2])
}

// ExtensionValue returns ExtensionValue value if type matches.
func (i *IE) ExtensionValue() []byte {
	if i.Type != PrivateExtension {
		return nil
	}
	return i.Payload[2:]
}
