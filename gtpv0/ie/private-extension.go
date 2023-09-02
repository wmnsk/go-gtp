// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewPrivateExtension creates a new PrivateExtension IE from string.
func NewPrivateExtension(id uint16, val []byte) *IE {
	i := New(PrivateExtension, make([]byte, 2+len(val)))
	binary.BigEndian.PutUint16(i.Payload[:2], id)
	copy(i.Payload[2:], val)
	return i
}

// PrivateExtension returns PrivateExtension value if type matches.
func (i *IE) PrivateExtension() ([]byte, error) {
	if i.Type != PrivateExtension {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return i.Payload, nil
}

// MustPrivateExtension returns PrivateExtension in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustPrivateExtension() []byte {
	v, _ := i.PrivateExtension()
	return v
}

// ExtensionIdentifier returns ExtensionIdentifier value in uint16 if type matches.
func (i *IE) ExtensionIdentifier() (uint16, error) {
	if i.Type != PrivateExtension {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint16(i.Payload[:2]), nil
}

// MustExtensionIdentifier returns ExtensionIdentifier in uint16 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustExtensionIdentifier() uint16 {
	v, _ := i.ExtensionIdentifier()
	return v
}

// ExtensionValue returns ExtensionValue value if type matches.
func (i *IE) ExtensionValue() ([]byte, error) {
	if i.Type != PrivateExtension {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 3 {
		return nil, io.ErrUnexpectedEOF
	}

	return i.Payload[2:], nil
}

// MustExtensionValue returns ExtensionValue in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustExtensionValue() []byte {
	v, _ := i.ExtensionValue()
	return v
}
