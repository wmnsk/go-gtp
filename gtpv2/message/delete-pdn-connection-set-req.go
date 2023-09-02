// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// DeletePDNConnectionSetRequest is a DeletePDNConnectionSetRequest Header and its IEs above.
type DeletePDNConnectionSetRequest struct {
	*Header
	MMEFQCSID        *ie.IE
	SGWFQCSID        *ie.IE
	PGWFQCSID        *ie.IE
	EPDGFQCSID       *ie.IE
	TWANFQCSID       *ie.IE
	PrivateExtension *ie.IE
	AdditionalIEs    []*ie.IE
}

// NewDeletePDNConnectionSetRequest creates a new DeletePDNConnectionSetRequest.
func NewDeletePDNConnectionSetRequest(teid, seq uint32, ies ...*ie.IE) *DeletePDNConnectionSetRequest {
	m := &DeletePDNConnectionSetRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeDeletePDNConnectionSetRequest, teid, seq, nil,
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
			case 2:
				m.PGWFQCSID = i
			case 3:
				m.EPDGFQCSID = i
			case 4:
				m.TWANFQCSID = i
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

// Marshal serializes DeletePDNConnectionSetRequest into bytes.
func (m *DeletePDNConnectionSetRequest) Marshal() ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes DeletePDNConnectionSetRequest into bytes.
func (m *DeletePDNConnectionSetRequest) MarshalTo(b []byte) error {
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
	if ie := m.PGWFQCSID; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.EPDGFQCSID; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.TWANFQCSID; ie != nil {
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

// ParseDeletePDNConnectionSetRequest decodes given bytes as DeletePDNConnectionSetRequest.
func ParseDeletePDNConnectionSetRequest(b []byte) (*DeletePDNConnectionSetRequest, error) {
	m := &DeletePDNConnectionSetRequest{}
	if err := m.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalBinary decodes given bytes as DeletePDNConnectionSetRequest.
func (m *DeletePDNConnectionSetRequest) UnmarshalBinary(b []byte) error {
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
			case 2:
				m.PGWFQCSID = i
			case 3:
				m.EPDGFQCSID = i
			case 4:
				m.TWANFQCSID = i
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
func (m *DeletePDNConnectionSetRequest) MarshalLen() int {
	l := m.Header.MarshalLen() - len(m.Header.Payload)

	if ie := m.MMEFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.SGWFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.PGWFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.EPDGFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.TWANFQCSID; ie != nil {
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
func (m *DeletePDNConnectionSetRequest) SetLength() {
	m.Header.Length = uint16(m.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (m *DeletePDNConnectionSetRequest) MessageTypeName() string {
	return "Delete PDN Connection Set Request"
}

// TEID returns the TEID in uint32.
func (m *DeletePDNConnectionSetRequest) TEID() uint32 {
	return m.Header.teid()
}
