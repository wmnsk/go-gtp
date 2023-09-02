// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// UpdatePDNConnectionSetRequest is a UpdatePDNConnectionSetRequest Header and its IEs above.
type UpdatePDNConnectionSetRequest struct {
	*Header
	MMEFQCSID        *ie.IE
	SGWFQCSID        *ie.IE
	PrivateExtension *ie.IE
	AdditionalIEs    []*ie.IE
}

// NewUpdatePDNConnectionSetRequest creates a new UpdatePDNConnectionSetRequest.
func NewUpdatePDNConnectionSetRequest(teid, seq uint32, ies ...*ie.IE) *UpdatePDNConnectionSetRequest {
	m := &UpdatePDNConnectionSetRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeUpdatePDNConnectionSetRequest, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				m.MMEFQCSID = i
			case 1:
				m.SGWFQCSID = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	m.SetLength()
	return m
}

// Marshal serializes UpdatePDNConnectionSetRequest into bytes.
func (m *UpdatePDNConnectionSetRequest) Marshal() ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes UpdatePDNConnectionSetRequest into bytes.
func (m *UpdatePDNConnectionSetRequest) MarshalTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.MarshalLen()-m.Header.MarshalLen())

	offset := 0
	if ie := m.MMEFQCSID; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.SGWFQCSID; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
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

// ParseUpdatePDNConnectionSetRequest decodes given bytes as UpdatePDNConnectionSetRequest.
func ParseUpdatePDNConnectionSetRequest(b []byte) (*UpdatePDNConnectionSetRequest, error) {
	m := &UpdatePDNConnectionSetRequest{}
	if err := m.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalBinary decodes given bytes as UpdatePDNConnectionSetRequest.
func (m *UpdatePDNConnectionSetRequest) UnmarshalBinary(b []byte) error {
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
		case ie.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				m.MMEFQCSID = i
			case 1:
				m.SGWFQCSID = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (m *UpdatePDNConnectionSetRequest) MarshalLen() int {
	l := m.Header.MarshalLen() - len(m.Header.Payload)

	if ie := m.MMEFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.SGWFQCSID; ie != nil {
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
func (m *UpdatePDNConnectionSetRequest) SetLength() {
	m.Header.Length = uint16(m.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (m *UpdatePDNConnectionSetRequest) MessageTypeName() string {
	return "Update PDN Connection Set Request"
}

// TEID returns the TEID in uint32.
func (m *UpdatePDNConnectionSetRequest) TEID() uint32 {
	return m.Header.teid()
}
