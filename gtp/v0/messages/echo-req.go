// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/pkg/errors"
	"github.com/wmnsk/go-gtp/gtp/v0/ies"
)

// EchoRequest is a EchoRequest Header and its AdditionalIEs above.
type EchoRequest struct {
	*Header
	PrivateExtension *ies.IE
	AdditionalIEs    []*ies.IE
}

// NewEchoRequest creates a new EchoRequest.
func NewEchoRequest(seq, label uint16, tid uint64, ie ...*ies.IE) *EchoRequest {
	e := &EchoRequest{
		Header: NewHeader(
			0x1e, MsgTypeEchoRequest, seq, label, tid, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.PrivateExtension:
			e.PrivateExtension = i
		default:
			e.AdditionalIEs = append(e.AdditionalIEs, i)
		}
	}

	e.SetLength()
	return e
}

// Serialize returns the byte sequence generated from a EchoRequest.
func (e *EchoRequest) Serialize() ([]byte, error) {
	b := make([]byte, e.Len())
	if err := e.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo puts the byte sequence in the byte array given as b.
func (e *EchoRequest) SerializeTo(b []byte) error {
	if e.Header.Payload != nil {
		e.Header.Payload = nil
	}
	e.Header.Payload = make([]byte, e.Len()-e.Header.Len())

	offset := 0
	if ie := e.PrivateExtension; ie != nil {
		if err := ie.SerializeTo(e.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	for _, ie := range e.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.SerializeTo(e.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	e.Header.SetLength()
	return e.Header.SerializeTo(b)
}

// DecodeEchoRequest decodes a given byte sequence as a EchoRequest.
func DecodeEchoRequest(b []byte) (*EchoRequest, error) {
	e := &EchoRequest{}
	if err := e.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return e, nil
}

// DecodeFromBytes decodes a given byte sequence as a EchoRequest.
func (e *EchoRequest) DecodeFromBytes(b []byte) error {
	var err error
	e.Header, err = DecodeHeader(b)
	if err != nil {
		return errors.Wrap(err, "failed to decode Header:")
	}
	if len(e.Header.Payload) < 2 {
		return nil
	}

	ie, err := ies.DecodeMultiIEs(e.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.PrivateExtension:
			e.PrivateExtension = i
		default:
			e.AdditionalIEs = append(e.AdditionalIEs, i)
		}
	}

	return nil
}

// Len returns the actual length of Data.
func (e *EchoRequest) Len() int {
	l := e.Header.Len() - len(e.Header.Payload)

	if ie := e.PrivateExtension; ie != nil {
		l += ie.Len()
	}

	for _, ie := range e.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.Len()
	}

	return l
}

// SetLength sets the length in Length field.
func (e *EchoRequest) SetLength() {
	e.Header.Length = uint16(e.Len() - 20)
}

// MessageTypeName returns the name of protocol.
func (e *EchoRequest) MessageTypeName() string {
	return "Echo Request"
}

// TID returns the TID in human-readable string.
func (e *EchoRequest) TID() string {
	return e.tid()
}
