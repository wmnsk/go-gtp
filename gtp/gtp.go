// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtp

import (
	v0msg "github.com/wmnsk/go-gtp/gtp/v0/messages"
	v1msg "github.com/wmnsk/go-gtp/gtp/v1/messages"
	v2msg "github.com/wmnsk/go-gtp/gtp/v2/messages"
)

// Message is an interface that defines all versions of GTP messages.
type Message interface {
	SerializeTo([]byte) error
	DecodeFromBytes(b []byte) error
	Len() int
	Version() int
	MessageType() uint8
	MessageTypeName() string
}

// Serialize returns the byte sequence generated from a Message instance.
// Better to use SerializeXxx instead if you know the name of message to be serialized.
func Serialize(m Message) ([]byte, error) {
	b := make([]byte, m.Len())
	if err := m.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// Decode decodes given bytes as Message.
func Decode(b []byte) (Message, error) {
	if len(b) < 8 {
		return nil, ErrTooShortToDecode
	}

	switch b[0] >> 5 {
	case 0:
		return v0msg.Decode(b)
	case 1:
		return v1msg.Decode(b)
	case 2:
		return v2msg.Decode(b)
	default:
		return nil, ErrInvalidVersion
	}
}
