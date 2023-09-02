// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"fmt"

	"github.com/wmnsk/go-gtp/gtpv0/ie"
)

// EchoRequest is a EchoRequest Header and its AdditionalIEs above.
type EchoRequest struct {
	*Header
	PrivateExtension *ie.IE
	AdditionalIEs    []*ie.IE
}

// NewEchoRequest creates a new EchoRequest.
func NewEchoRequest(seq, label uint16, tid uint64, ies ...*ie.IE) *EchoRequest {
	e := &EchoRequest{
		Header: NewHeader(
			0x1e, MsgTypeEchoRequest, seq, label, tid, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.PrivateExtension:
			e.PrivateExtension = i
		default:
			e.AdditionalIEs = append(e.AdditionalIEs, i)
		}
	}

	e.SetLength()
	return e
}

// Marshal returns the byte sequence generated from a EchoRequest.
func (e *EchoRequest) Marshal() ([]byte, error) {
	b := make([]byte, e.MarshalLen())
	if err := e.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (e *EchoRequest) MarshalTo(b []byte) error {
	if e.Header.Payload != nil {
		e.Header.Payload = nil
	}
	e.Header.Payload = make([]byte, e.MarshalLen()-e.Header.MarshalLen())

	offset := 0
	if ie := e.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(e.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	for _, ie := range e.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(e.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	e.Header.SetLength()
	return e.Header.MarshalTo(b)
}

// ParseEchoRequest parses a given byte sequence as a EchoRequest.
func ParseEchoRequest(b []byte) (*EchoRequest, error) {
	e := &EchoRequest{}
	if err := e.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return e, nil
}

// UnmarshalBinary parses a given byte sequence as a EchoRequest.
func (e *EchoRequest) UnmarshalBinary(b []byte) error {
	var err error
	e.Header, err = ParseHeader(b)
	if err != nil {
		return fmt.Errorf("failed to Parse Header: %w", err)
	}
	if len(e.Header.Payload) < 2 {
		return nil
	}

	ies, err := ie.ParseMultiIEs(e.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.PrivateExtension:
			e.PrivateExtension = i
		default:
			e.AdditionalIEs = append(e.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length of Data.
func (e *EchoRequest) MarshalLen() int {
	l := e.Header.MarshalLen() - len(e.Header.Payload)

	if ie := e.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}

	for _, ie := range e.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}

	return l
}

// SetLength sets the length in Length field.
func (e *EchoRequest) SetLength() {
	e.Header.Length = uint16(e.MarshalLen() - 20)
}

// MessageTypeName returns the name of protocol.
func (e *EchoRequest) MessageTypeName() string {
	return "Echo Request"
}

// TID returns the TID in human-readable string.
func (e *EchoRequest) TID() string {
	return e.tid()
}
