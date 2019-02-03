// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "github.com/wmnsk/go-gtp/gtp/v1/ies"

// ErrorIndication is a ErrorIndication Header and its IEs above.
type ErrorIndication struct {
	*Header
	TEIDDataI        *ies.IE
	GTPUPeerAddress  *ies.IE
	PrivateExtension *ies.IE
	AdditionalIEs    []*ies.IE
}

// NewErrorIndication creates a new GTPv1 NewErrorIndication.
func NewErrorIndication(teid uint32, seq uint16, ie ...*ies.IE) *ErrorIndication {
	e := &ErrorIndication{
		Header: NewHeader(0x32, MsgTypeErrorIndication, teid, seq, nil),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.TEIDDataI:
			e.TEIDDataI = i
		case ies.GSNAddress:
			e.GTPUPeerAddress = i
		case ies.PrivateExtension:
			e.PrivateExtension = i
		default:
			e.AdditionalIEs = append(e.AdditionalIEs, i)
		}
	}

	e.SetLength()
	return e
}

// Serialize returns the byte sequence generated from a ErrorIndication.
func (e *ErrorIndication) Serialize() ([]byte, error) {
	b := make([]byte, e.Len())
	if err := e.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo puts the byte sequence in the byte array given as b.
func (e *ErrorIndication) SerializeTo(b []byte) error {
	if len(b) < e.Len() {
		return ErrTooShortToSerialize
	}
	e.Header.Payload = make([]byte, e.Len()-e.Header.Len())

	offset := 0
	if ie := e.TEIDDataI; ie != nil {
		if err := ie.SerializeTo(e.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := e.GTPUPeerAddress; ie != nil {
		if err := ie.SerializeTo(e.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
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

// DecodeErrorIndication decodes a given byte sequence as a ErrorIndication.
func DecodeErrorIndication(b []byte) (*ErrorIndication, error) {
	e := &ErrorIndication{}
	if err := e.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return e, nil
}

// DecodeFromBytes decodes a given byte sequence as a ErrorIndication.
func (e *ErrorIndication) DecodeFromBytes(b []byte) error {
	var err error
	e.Header, err = DecodeHeader(b)
	if err != nil {
		return err
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
		case ies.TEIDDataI:
			e.TEIDDataI = i
		case ies.GSNAddress:
			e.GTPUPeerAddress = i
		case ies.PrivateExtension:
			e.PrivateExtension = i
		default:
			e.AdditionalIEs = append(e.AdditionalIEs, i)
		}
	}
	return nil
}

// Len returns the actual length of Data.
func (e *ErrorIndication) Len() int {
	l := e.Header.Len() - len(e.Header.Payload)

	if ie := e.TEIDDataI; ie != nil {
		l += ie.Len()
	}
	if ie := e.GTPUPeerAddress; ie != nil {
		l += ie.Len()
	}
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
func (e *ErrorIndication) SetLength() {
	e.Length = uint16(e.Len() - 8)
}

// MessageTypeName returns the name of protocol.
func (e *ErrorIndication) MessageTypeName() string {
	return "Errror Indication"
}

// TEID returns the TEID in human-readable string.
func (e *ErrorIndication) TEID() uint32 {
	return e.Header.TEID
}
