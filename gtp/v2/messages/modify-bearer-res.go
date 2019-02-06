// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/gtp/v2/ies"
)

// ModifyBearerResponse is a ModifyBearerResponse Header and its IEs above.
type ModifyBearerResponse struct {
	*Header
	Cause                          *ies.IE
	MSISDN                         *ies.IE
	LinkedEBI                      *ies.IE
	APNRestriction                 *ies.IE
	PCO                            *ies.IE
	BearerContextsModified         *ies.IE
	BearerContextsMarkedForRemoval *ies.IE
	ChangeReportingAction          *ies.IE
	CSGInformationReportingAction  *ies.IE
	HeNBInformationReporting       *ies.IE
	ChargingGatewayName            *ies.IE
	ChargingGatewayAddress         *ies.IE
	PGWFQCSID                      *ies.IE
	SGWFQCSID                      *ies.IE
	Recovery                       *ies.IE
	SGWLDN                         *ies.IE
	PGWLDN                         *ies.IE
	IndicationFlags                *ies.IE
	PresenceReportingAreaAction    *ies.IE
	PGWNodeLoadControlInformation  *ies.IE
	PGWAPNLoadControlInformation   *ies.IE
	SGWNodeLoadControlInformation  *ies.IE
	PGWOverloadControlInformation  *ies.IE
	SGWOverloadControlInformation  *ies.IE
	PDNConnectionChargingID        *ies.IE
	PrivateExtension               *ies.IE
	AdditionalIEs                  []*ies.IE
}

// NewModifyBearerResponse creates a new ModifyBearerResponse.
func NewModifyBearerResponse(teid, seq uint32, ie ...*ies.IE) *ModifyBearerResponse {
	m := &ModifyBearerResponse{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeModifyBearerResponse, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			m.Cause = i
		case ies.MSISDN:
			m.MSISDN = i
		case ies.EPSBearerID:
			m.LinkedEBI = i
		case ies.APNRestriction:
			m.APNRestriction = i
		case ies.ProtocolConfigurationOptions:
			m.PCO = i
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				m.BearerContextsModified = i
			case 1:
				m.BearerContextsMarkedForRemoval = i
			}
		case ies.ChangeReportingAction:
			m.ChangeReportingAction = i
		case ies.CSGInformationReportingAction:
			m.CSGInformationReportingAction = i
		case ies.HeNBInformationReporting:
			m.HeNBInformationReporting = i
		case ies.FullyQualifiedDomainName:
			m.ChargingGatewayName = i
		case ies.IPAddress:
			m.ChargingGatewayAddress = i
		case ies.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				m.PGWFQCSID = i
			case 1:
				m.SGWFQCSID = i
			}
		case ies.Recovery:
			m.Recovery = i
		case ies.LocalDistinguishedName:
			switch i.Instance() {
			case 0:
				m.SGWLDN = i
			case 1:
				m.PGWLDN = i
			}
		case ies.Indication:
			m.IndicationFlags = i
		case ies.PresenceReportingAreaAction:
			m.PresenceReportingAreaAction = i
		case ies.LoadControlInformation:
			switch i.Instance() {
			case 0:
				m.PGWNodeLoadControlInformation = i
			case 1:
				m.PGWAPNLoadControlInformation = i
			}
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				m.SGWNodeLoadControlInformation = i
			case 1:
				m.PGWOverloadControlInformation = i
			case 2:
				m.SGWOverloadControlInformation = i
			}
		case ies.ChargingID:
			m.PDNConnectionChargingID = i
		case ies.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	m.SetLength()
	return m
}

// Serialize serializes ModifyBearerResponse into bytes.
func (m *ModifyBearerResponse) Serialize() ([]byte, error) {
	b := make([]byte, m.Len())
	if err := m.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes ModifyBearerResponse into bytes.
func (m *ModifyBearerResponse) SerializeTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.Len()-m.Header.Len())

	offset := 0
	if ie := m.Cause; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.MSISDN; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.LinkedEBI; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.APNRestriction; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.PCO; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.BearerContextsModified; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.BearerContextsMarkedForRemoval; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.ChangeReportingAction; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.CSGInformationReportingAction; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.HeNBInformationReporting; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.ChargingGatewayName; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.ChargingGatewayAddress; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.PGWFQCSID; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.SGWFQCSID; ie != nil {
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
	if ie := m.SGWLDN; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.PGWLDN; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.IndicationFlags; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.PresenceReportingAreaAction; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.PGWNodeLoadControlInformation; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.PGWAPNLoadControlInformation; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.SGWNodeLoadControlInformation; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.PGWOverloadControlInformation; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.SGWOverloadControlInformation; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.PDNConnectionChargingID; ie != nil {
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

// DecodeModifyBearerResponse decodes given bytes as ModifyBearerResponse.
func DecodeModifyBearerResponse(b []byte) (*ModifyBearerResponse, error) {
	m := &ModifyBearerResponse{}
	if err := m.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return m, nil
}

// DecodeFromBytes decodes given bytes as ModifyBearerResponse.
func (m *ModifyBearerResponse) DecodeFromBytes(b []byte) error {
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
		case ies.Cause:
			m.Cause = i
		case ies.MSISDN:
			m.MSISDN = i
		case ies.EPSBearerID:
			m.LinkedEBI = i
		case ies.APNRestriction:
			m.APNRestriction = i
		case ies.ProtocolConfigurationOptions:
			m.PCO = i
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				m.BearerContextsModified = i
			case 1:
				m.BearerContextsMarkedForRemoval = i
			}
		case ies.ChangeReportingAction:
			m.ChangeReportingAction = i
		case ies.CSGInformationReportingAction:
			m.CSGInformationReportingAction = i
		case ies.HeNBInformationReporting:
			m.HeNBInformationReporting = i
		case ies.FullyQualifiedDomainName:
			m.ChargingGatewayName = i
		case ies.IPAddress:
			m.ChargingGatewayAddress = i
		case ies.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				m.PGWFQCSID = i
			case 1:
				m.SGWFQCSID = i
			}
		case ies.Recovery:
			m.Recovery = i
		case ies.LocalDistinguishedName:
			switch i.Instance() {
			case 0:
				m.SGWLDN = i
			case 1:
				m.PGWLDN = i
			}
		case ies.Indication:
			m.IndicationFlags = i
		case ies.PresenceReportingAreaAction:
			m.PresenceReportingAreaAction = i
		case ies.LoadControlInformation:
			switch i.Instance() {
			case 0:
				m.PGWNodeLoadControlInformation = i
			case 1:
				m.PGWAPNLoadControlInformation = i
			}
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				m.SGWNodeLoadControlInformation = i
			case 1:
				m.PGWOverloadControlInformation = i
			case 2:
				m.SGWOverloadControlInformation = i
			}
		case ies.ChargingID:
			m.PDNConnectionChargingID = i
		case ies.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	return nil
}

// Len returns the actual length in int.
func (m *ModifyBearerResponse) Len() int {
	l := m.Header.Len() - len(m.Header.Payload)

	if ie := m.Cause; ie != nil {
		l += ie.Len()
	}
	if ie := m.MSISDN; ie != nil {
		l += ie.Len()
	}
	if ie := m.LinkedEBI; ie != nil {
		l += ie.Len()
	}
	if ie := m.APNRestriction; ie != nil {
		l += ie.Len()
	}
	if ie := m.PCO; ie != nil {
		l += ie.Len()
	}
	if ie := m.BearerContextsModified; ie != nil {
		l += ie.Len()
	}
	if ie := m.BearerContextsMarkedForRemoval; ie != nil {
		l += ie.Len()
	}
	if ie := m.ChangeReportingAction; ie != nil {
		l += ie.Len()
	}
	if ie := m.CSGInformationReportingAction; ie != nil {
		l += ie.Len()
	}
	if ie := m.HeNBInformationReporting; ie != nil {
		l += ie.Len()
	}
	if ie := m.ChargingGatewayName; ie != nil {
		l += ie.Len()
	}
	if ie := m.ChargingGatewayAddress; ie != nil {
		l += ie.Len()
	}
	if ie := m.PGWFQCSID; ie != nil {
		l += ie.Len()
	}
	if ie := m.SGWFQCSID; ie != nil {
		l += ie.Len()
	}
	if ie := m.Recovery; ie != nil {
		l += ie.Len()
	}
	if ie := m.SGWLDN; ie != nil {
		l += ie.Len()
	}
	if ie := m.PGWLDN; ie != nil {
		l += ie.Len()
	}
	if ie := m.IndicationFlags; ie != nil {
		l += ie.Len()
	}
	if ie := m.PresenceReportingAreaAction; ie != nil {
		l += ie.Len()
	}
	if ie := m.PGWNodeLoadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := m.PGWAPNLoadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := m.SGWNodeLoadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := m.PGWOverloadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := m.SGWOverloadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := m.PDNConnectionChargingID; ie != nil {
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
func (m *ModifyBearerResponse) SetLength() {
	m.Header.Length = uint16(m.Len() - 4)
}

// MessageTypeName returns the name of protocol.
func (m *ModifyBearerResponse) MessageTypeName() string {
	return "Modify Bearer Response"
}

// TEID returns the TEID in uint32.
func (m *ModifyBearerResponse) TEID() uint32 {
	return m.Header.teid()
}
