// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtp

import (
	v0msg "github.com/wmnsk/go-gtp/gtpv0/message"
	v1msg "github.com/wmnsk/go-gtp/gtpv1/message"
	v2msg "github.com/wmnsk/go-gtp/gtpv2/message"
)

// Message is an interface that defines all versions of GTP message.
type Message interface {
	MarshalTo([]byte) error
	UnmarshalBinary(b []byte) error
	MarshalLen() int
	Version() int
	MessageType() uint8
	MessageTypeName() string

	// deprecated
	SerializeTo([]byte) error
	DecodeFromBytes(b []byte) error
}

// Marshal returns the byte sequence generated from a Message instance.
// Better to use (*MessageName).Marshal instead if you know the name of message to be serialized.
func Marshal(m Message) ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// Parse decodes given bytes as Message.
func Parse(b []byte) (Message, error) {
	if len(b) < 8 {
		return nil, ErrTooShortToParse
	}

	switch b[0] >> 5 {
	case 0:
		return v0msg.Parse(b)
	case 1:
		return v1msg.Parse(b)
	case 2:
		return v2msg.Parse(b)
	default:
		return nil, ErrInvalidVersion
	}
}
