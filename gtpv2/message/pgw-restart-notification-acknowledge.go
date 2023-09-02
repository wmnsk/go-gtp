// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// PGWRestartNotificationAcknowledge is a PGWRestartNotificationAcknowledge Header and its IEs above.
type PGWRestartNotificationAcknowledge struct {
	*Header
	Cause            *ie.IE
	PrivateExtension *ie.IE
	AdditionalIEs    []*ie.IE
}

// NewPGWRestartNotificationAcknowledge creates a new PGWRestartNotificationAcknowledge.
func NewPGWRestartNotificationAcknowledge(teid, seq uint32, ies ...*ie.IE) *PGWRestartNotificationAcknowledge {
	m := &PGWRestartNotificationAcknowledge{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypePGWRestartNotificationAcknowledge, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			m.Cause = i
		case ie.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	m.SetLength()
	return m
}

// Marshal serializes PGWRestartNotificationAcknowledge into bytes.
func (m *PGWRestartNotificationAcknowledge) Marshal() ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes PGWRestartNotificationAcknowledge into bytes.
func (m *PGWRestartNotificationAcknowledge) MarshalTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.MarshalLen()-m.Header.MarshalLen())

	offset := 0
	if ie := m.Cause; ie != nil {
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

// ParsePGWRestartNotificationAcknowledge decodes given bytes as PGWRestartNotificationAcknowledge.
func ParsePGWRestartNotificationAcknowledge(b []byte) (*PGWRestartNotificationAcknowledge, error) {
	m := &PGWRestartNotificationAcknowledge{}
	if err := m.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalBinary decodes given bytes as PGWRestartNotificationAcknowledge.
func (m *PGWRestartNotificationAcknowledge) UnmarshalBinary(b []byte) error {
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
		case ie.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (m *PGWRestartNotificationAcknowledge) MarshalLen() int {
	l := m.Header.MarshalLen() - len(m.Header.Payload)

	if ie := m.Cause; ie != nil {
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
func (m *PGWRestartNotificationAcknowledge) SetLength() {
	m.Header.Length = uint16(m.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (m *PGWRestartNotificationAcknowledge) MessageTypeName() string {
	return "PGW Restart Notification Acknowledge"
}

// TEID returns the TEID in uint32.
func (m *PGWRestartNotificationAcknowledge) TEID() uint32 {
	return m.Header.teid()
}
