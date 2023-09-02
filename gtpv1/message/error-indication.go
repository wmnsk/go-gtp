// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "github.com/wmnsk/go-gtp/gtpv1/ie"

// ErrorIndication is a ErrorIndication Header and its IEs above.
type ErrorIndication struct {
	*Header
	TEIDDataI        *ie.IE
	GTPUPeerAddress  *ie.IE
	PrivateExtension *ie.IE
	AdditionalIEs    []*ie.IE
}

// NewErrorIndication creates a new GTPv1 ErrorIndication.
func NewErrorIndication(teid uint32, seq uint16, ies ...*ie.IE) *ErrorIndication {
	e := &ErrorIndication{
		Header: NewHeader(0x32, MsgTypeErrorIndication, teid, seq, nil),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.TEIDDataI:
			e.TEIDDataI = i
		case ie.GSNAddress:
			e.GTPUPeerAddress = i
		case ie.PrivateExtension:
			e.PrivateExtension = i
		default:
			e.AdditionalIEs = append(e.AdditionalIEs, i)
		}
	}

	e.SetLength()
	return e
}

// Marshal returns the byte sequence generated from a ErrorIndication.
func (e *ErrorIndication) Marshal() ([]byte, error) {
	b := make([]byte, e.MarshalLen())
	if err := e.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (e *ErrorIndication) MarshalTo(b []byte) error {
	if len(b) < e.MarshalLen() {
		return ErrTooShortToMarshal
	}
	e.Header.Payload = make([]byte, e.MarshalLen()-e.Header.MarshalLen())

	offset := 0
	if ie := e.TEIDDataI; ie != nil {
		if err := ie.MarshalTo(e.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := e.GTPUPeerAddress; ie != nil {
		if err := ie.MarshalTo(e.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
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

// ParseErrorIndication decodes a given byte sequence as a ErrorIndication.
func ParseErrorIndication(b []byte) (*ErrorIndication, error) {
	e := &ErrorIndication{}
	if err := e.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return e, nil
}

// UnmarshalBinary decodes a given byte sequence as a ErrorIndication.
func (e *ErrorIndication) UnmarshalBinary(b []byte) error {
	var err error
	e.Header, err = ParseHeader(b)
	if err != nil {
		return err
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
		case ie.TEIDDataI:
			e.TEIDDataI = i
		case ie.GSNAddress:
			e.GTPUPeerAddress = i
		case ie.PrivateExtension:
			e.PrivateExtension = i
		default:
			e.AdditionalIEs = append(e.AdditionalIEs, i)
		}
	}
	return nil
}

// MarshalLen returns the serial length of Data.
func (e *ErrorIndication) MarshalLen() int {
	l := e.Header.MarshalLen() - len(e.Header.Payload)

	if ie := e.TEIDDataI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := e.GTPUPeerAddress; ie != nil {
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
func (e *ErrorIndication) SetLength() {
	e.Length = uint16(e.MarshalLen() - 8)
}

// MessageTypeName returns the name of protocol.
func (e *ErrorIndication) MessageTypeName() string {
	return "Errror Indication"
}

// TEID returns the TEID in human-readable string.
func (e *ErrorIndication) TEID() uint32 {
	return e.Header.TEID
}
