// Copyright 2019-2023 go-gtp authors. All rights reservev.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "github.com/wmnsk/go-gtp/gtpv2/ie"

// ModifyBearerCommand is a ModifyBearerCommand Header and its IEs above.
type ModifyBearerCommand struct {
	*Header
	APNAMBR                            *ie.IE
	BearerContext                      *ie.IE
	MMESGSNOverloadControlInformation  *ie.IE
	SGWOverloadControlInformation      *ie.IE
	TWANePDGOverloadControlInformation *ie.IE
	SenderFTEIDC                       *ie.IE
	PrivateExtension                   *ie.IE
	AdditionalIEs                      []*ie.IE
}

// NewModifyBearerCommand creates a new ModifyBearerCommand.
func NewModifyBearerCommand(teid, seq uint32, ies ...*ie.IE) *ModifyBearerCommand {
	m := &ModifyBearerCommand{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeModifyBearerCommand, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.AggregateMaximumBitRate:
			m.APNAMBR = i
		case ie.BearerContext:
			m.BearerContext = i
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				m.MMESGSNOverloadControlInformation = i
			case 1:
				m.SGWOverloadControlInformation = i
			case 2:
				m.TWANePDGOverloadControlInformation = i
			}
		case ie.FullyQualifiedTEID:
			m.SenderFTEIDC = i
		case ie.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	m.SetLength()
	return m
}

// Marshal serializes ModifyBearerCommand into bytes.
func (m *ModifyBearerCommand) Marshal() ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes ModifyBearerCommand into bytes.
func (m *ModifyBearerCommand) MarshalTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.MarshalLen()-m.Header.MarshalLen())

	offset := 0
	if ie := m.APNAMBR; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.BearerContext; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.MMESGSNOverloadControlInformation; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.SGWOverloadControlInformation; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.TWANePDGOverloadControlInformation; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.SenderFTEIDC; ie != nil {
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

// ParseModifyBearerCommand decodes given bytes as ModifyBearerCommand.
func ParseModifyBearerCommand(b []byte) (*ModifyBearerCommand, error) {
	m := &ModifyBearerCommand{}
	if err := m.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalBinary decodes given bytes as ModifyBearerCommand.
func (m *ModifyBearerCommand) UnmarshalBinary(b []byte) error {
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
		case ie.AggregateMaximumBitRate:
			m.APNAMBR = i
		case ie.BearerContext:
			m.BearerContext = i
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				m.MMESGSNOverloadControlInformation = i
			case 1:
				m.SGWOverloadControlInformation = i
			case 2:
				m.TWANePDGOverloadControlInformation = i
			}
		case ie.FullyQualifiedTEID:
			m.SenderFTEIDC = i
		case ie.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (m *ModifyBearerCommand) MarshalLen() int {
	l := m.Header.MarshalLen() - len(m.Header.Payload)
	if ie := m.APNAMBR; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.BearerContext; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.MMESGSNOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.SGWOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.TWANePDGOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.SenderFTEIDC; ie != nil {
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
func (m *ModifyBearerCommand) SetLength() {
	m.Header.Length = uint16(m.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (m *ModifyBearerCommand) MessageTypeName() string {
	return "Modify Bearer Command"
}

// TEID returns the TEID in uint32.
func (m *ModifyBearerCommand) TEID() uint32 {
	return m.Header.teid()
}
