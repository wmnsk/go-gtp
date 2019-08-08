// Copyright 2019 go-gtp authors. All rights reservev.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "github.com/wmnsk/go-gtp/v2/ies"

// ModifyAccessBearersRequest is a ModifyAccessBearersRequest Header and its IEs above.
type ModifyAccessBearersRequest struct {
	*Header
	IndicationFlags                        *ies.IE
	SenderFTEIDC                           *ies.IE
	DelayDownlinkPacketNotificationRequest *ies.IE
	BearerContextsToBeModified             *ies.IE
	BearerContextsToBeRemoved              *ies.IE
	Recovery                               *ies.IE
	SecondaryRATUsageDataReport            *ies.IE
	PrivateExtension                       *ies.IE
	AdditionalIEs                          []*ies.IE
}

// NewModifyAccessBearersRequest creates a new ModifyAccessBearersRequest.
func NewModifyAccessBearersRequest(teid, seq uint32, ie ...*ies.IE) *ModifyAccessBearersRequest {
	m := &ModifyAccessBearersRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeModifyAccessBearersRequest, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Indication:
			m.IndicationFlags = i
		case ies.NodeType:
			m.SenderFTEIDC = i
		case ies.DelayValue:
			m.DelayDownlinkPacketNotificationRequest = i
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				m.BearerContextsToBeModified = i
			case 1:
				m.BearerContextsToBeRemoved = i
			}
		case ies.Recovery:
			m.Recovery = i
		case ies.SecondaryRATUsageDataReport:
			m.SecondaryRATUsageDataReport = i
		case ies.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	m.SetLength()
	return m
}

// Serialize serializes ModifyAccessBearersRequest into bytes.
func (m *ModifyAccessBearersRequest) Serialize() ([]byte, error) {
	b := make([]byte, m.Len())
	if err := m.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes ModifyAccessBearersRequest into bytes.
func (m *ModifyAccessBearersRequest) SerializeTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.Len()-m.Header.Len())

	offset := 0
	if ie := m.IndicationFlags; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.SenderFTEIDC; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.DelayDownlinkPacketNotificationRequest; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.BearerContextsToBeModified; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.BearerContextsToBeRemoved; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.Recovery; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.SecondaryRATUsageDataReport; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.PrivateExtension; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	for _, ie := range m.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.SerializeTo(m.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	m.Header.SetLength()
	return m.Header.SerializeTo(b)
}

// DecodeModifyAccessBearersRequest decodes given bytes as ModifyAccessBearersRequest.
func DecodeModifyAccessBearersRequest(b []byte) (*ModifyAccessBearersRequest, error) {
	m := &ModifyAccessBearersRequest{}
	if err := m.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return m, nil
}

// DecodeFromBytes decodes given bytes as ModifyAccessBearersRequest.
func (m *ModifyAccessBearersRequest) DecodeFromBytes(b []byte) error {
	var err error
	m.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}
	if len(m.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ies.DecodeMultiIEs(m.Header.Payload)
	if err != nil {
		return err
	}
	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Indication:
			m.IndicationFlags = i
		case ies.NodeType:
			m.SenderFTEIDC = i
		case ies.DelayValue:
			m.DelayDownlinkPacketNotificationRequest = i
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				m.BearerContextsToBeModified = i
			case 1:
				m.BearerContextsToBeRemoved = i
			}
		case ies.Recovery:
			m.Recovery = i
		case ies.SecondaryRATUsageDataReport:
			m.SecondaryRATUsageDataReport = i
		case ies.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	return nil
}

// Len returns the actual length in int.
func (m *ModifyAccessBearersRequest) Len() int {
	l := m.Header.Len() - len(m.Header.Payload)
	if ie := m.IndicationFlags; ie != nil {
		l += ie.Len()
	}
	if ie := m.SenderFTEIDC; ie != nil {
		l += ie.Len()
	}
	if ie := m.DelayDownlinkPacketNotificationRequest; ie != nil {
		l += ie.Len()
	}
	if ie := m.BearerContextsToBeModified; ie != nil {
		l += ie.Len()
	}
	if ie := m.BearerContextsToBeRemoved; ie != nil {
		l += ie.Len()
	}
	if ie := m.Recovery; ie != nil {
		l += ie.Len()
	}
	if ie := m.SecondaryRATUsageDataReport; ie != nil {
		l += ie.Len()
	}
	if ie := m.PrivateExtension; ie != nil {
		l += ie.Len()
	}

	for _, ie := range m.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.Len()
	}
	return l
}

// SetLength sets the length in Length field.
func (m *ModifyAccessBearersRequest) SetLength() {
	m.Header.Length = uint16(m.Len() - 4)
}

// MessageTypeName returns the name of protocol.
func (m *ModifyAccessBearersRequest) MessageTypeName() string {
	return "Modify Access Bearers Request"
}

// TEID returns the TEID in uint32.
func (m *ModifyAccessBearersRequest) TEID() uint32 {
	return m.Header.teid()
}
