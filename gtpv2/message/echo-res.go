// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// EchoResponse is a EchoResponse Header and its IEs above.
type EchoResponse struct {
	*Header
	Recovery            *ie.IE
	SendingNodeFeatures *ie.IE
	PrivateExtension    *ie.IE
	AdditionalIEs       []*ie.IE
}

// NewEchoResponse creates a new EchoResponse.
func NewEchoResponse(seq uint32, ies ...*ie.IE) *EchoResponse {
	e := &EchoResponse{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 0),
			MsgTypeEchoResponse, 0, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Recovery:
			e.Recovery = i
		case ie.NodeFeatures:
			e.SendingNodeFeatures = i
		case ie.PrivateExtension:
			e.PrivateExtension = i
		default:
			e.AdditionalIEs = append(e.AdditionalIEs, i)
		}
	}

	e.SetLength()
	return e
}

// Marshal returns the byte sequence generated from a EchoResponse.
func (e *EchoResponse) Marshal() ([]byte, error) {
	b := make([]byte, e.MarshalLen())
	if err := e.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (e *EchoResponse) MarshalTo(b []byte) error {
	if e.Header.Payload != nil {
		e.Header.Payload = nil
	}
	e.Header.Payload = make([]byte, e.MarshalLen()-e.Header.MarshalLen())

	offset := 0
	if ie := e.Recovery; ie != nil {
		if err := ie.MarshalTo(e.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := e.SendingNodeFeatures; ie != nil {
		if err := ie.MarshalTo(e.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := e.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(e.Header.Payload[offset:]); err != nil {
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

// ParseEchoResponse decodes a given byte sequence as a EchoResponse.
func ParseEchoResponse(b []byte) (*EchoResponse, error) {
	e := &EchoResponse{}
	if err := e.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return e, nil
}

// UnmarshalBinary decodes a given byte sequence as a EchoResponse.
func (e *EchoResponse) UnmarshalBinary(b []byte) error {
	var err error
	e.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(e.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ie.ParseMultiIEs(e.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Recovery:
			e.Recovery = i
		case ie.NodeFeatures:
			e.SendingNodeFeatures = i
		case ie.PrivateExtension:
			e.PrivateExtension = i
		default:
			e.AdditionalIEs = append(e.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length of Data.
func (e *EchoResponse) MarshalLen() int {
	l := e.Header.MarshalLen() - len(e.Header.Payload)

	if ie := e.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := e.SendingNodeFeatures; ie != nil {
		l += ie.MarshalLen()
	}
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
func (e *EchoResponse) SetLength() {
	e.Header.Length = uint16(e.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (e *EchoResponse) MessageTypeName() string {
	return "Echo Response"
}

// TEID returns the TEID in uint32.
func (e *EchoResponse) TEID() uint32 {
	return e.Header.teid()
}
