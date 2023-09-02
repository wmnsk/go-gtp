// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// UpdatePDNConnectionSetResponse is a UpdatePDNConnectionSetResponse Header and its IEs above.
type UpdatePDNConnectionSetResponse struct {
	*Header
	Cause            *ie.IE
	PGWFQCSID        *ie.IE
	Recovery         *ie.IE
	PrivateExtension *ie.IE
	AdditionalIEs    []*ie.IE
}

// NewUpdatePDNConnectionSetResponse creates a new UpdatePDNConnectionSetResponse.
func NewUpdatePDNConnectionSetResponse(teid, seq uint32, ies ...*ie.IE) *UpdatePDNConnectionSetResponse {
	m := &UpdatePDNConnectionSetResponse{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeUpdatePDNConnectionSetResponse, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			m.Cause = i
		case ie.FullyQualifiedCSID:
			m.PGWFQCSID = i
		case ie.Recovery:
			m.Recovery = i
		case ie.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	m.SetLength()
	return m
}

// Marshal returns the byte sequence generated from a UpdatePDNConnectionSetResponse.
func (m *UpdatePDNConnectionSetResponse) Marshal() ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (m *UpdatePDNConnectionSetResponse) MarshalTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.MarshalLen()-m.Header.MarshalLen())

	offset := 0
	if ie := m.Cause; ie != nil {
		if err := ie.MarshalTo(m.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.PGWFQCSID; ie != nil {
		if err := ie.MarshalTo(m.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.Recovery; ie != nil {
		if err := ie.MarshalTo(m.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(m.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	for _, ie := range m.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(m.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	m.Header.SetLength()
	return m.Header.MarshalTo(b)
}

// ParseUpdatePDNConnectionSetResponse decodes a given byte sequence as a UpdatePDNConnectionSetResponse.
func ParseUpdatePDNConnectionSetResponse(b []byte) (*UpdatePDNConnectionSetResponse, error) {
	m := &UpdatePDNConnectionSetResponse{}
	if err := m.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalBinary decodes given byte sequence as UpdatePDNConnectionSetResponse.
func (m *UpdatePDNConnectionSetResponse) UnmarshalBinary(b []byte) error {
	var err error
	m.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(m.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ie.ParseMultiIEs(m.Header.Payload)
	if err != nil {
		return err
	}
	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			m.Cause = i
		case ie.FullyQualifiedCSID:
			m.PGWFQCSID = i
		case ie.Recovery:
			m.Recovery = i
		case ie.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length of Data.
func (m *UpdatePDNConnectionSetResponse) MarshalLen() int {
	l := m.Header.MarshalLen() - len(m.Header.Payload)

	if ie := m.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.PGWFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}

	for _, ie := range m.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length field.
func (m *UpdatePDNConnectionSetResponse) SetLength() {
	m.Header.Length = uint16(m.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (m *UpdatePDNConnectionSetResponse) MessageTypeName() string {
	return "Update PDN Connection Set Response"
}

// TEID returns the TEID in uint32.
func (m *UpdatePDNConnectionSetResponse) TEID() uint32 {
	return m.Header.teid()
}
