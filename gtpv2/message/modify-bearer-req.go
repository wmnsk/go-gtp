// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// ModifyBearerRequest is a ModifyBearerRequest Header and its IEs above.
type ModifyBearerRequest struct {
	*Header
	MEI                                    *ie.IE
	ULI                                    *ie.IE
	ServingNetwork                         *ie.IE
	RATType                                *ie.IE
	IndicationFlags                        *ie.IE
	SenderFTEIDC                           *ie.IE
	AMBR                                   *ie.IE
	DelayDownlinkPacketNotificationRequest *ie.IE
	BearerContextsToBeModified             []*ie.IE
	BearerContextsToBeRemoved              []*ie.IE
	Recovery                               *ie.IE
	UETimeZone                             *ie.IE
	MMEFQCSID                              *ie.IE
	SGWFQCSID                              *ie.IE
	UCI                                    *ie.IE
	UELocalIPAddress                       *ie.IE
	UEUDPPort                              *ie.IE
	MMESGSNLDN                             *ie.IE
	SGWLDN                                 *ie.IE
	HeNBLocalIPAddress                     *ie.IE
	HeNBUDPPort                            *ie.IE
	MMESGSNIdentifier                      *ie.IE
	CNOperatorSelectionEntity              *ie.IE
	PresenceReportingAreaInformation       []*ie.IE
	MMESGSNOverloadControlInformation      *ie.IE
	SGWOverloadControlInformation          *ie.IE
	EPDGOverloadControlInformation         *ie.IE
	ServingPLMNRateControl                 *ie.IE
	MOExceptionDataCounter                 *ie.IE
	IMSI                                   *ie.IE
	ULIForSGW                              *ie.IE
	WLANLocationInformation                *ie.IE
	WLANLocationTimeStamp                  *ie.IE
	SecondaryRATUsageDataReport            []*ie.IE
	PrivateExtension                       *ie.IE
	AdditionalIEs                          []*ie.IE
}

// NewModifyBearerRequest creates a new ModifyBearerRequest.
func NewModifyBearerRequest(teid, seq uint32, ies ...*ie.IE) *ModifyBearerRequest {
	m := &ModifyBearerRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeModifyBearerRequest, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.MobileEquipmentIdentity:
			m.MEI = i
		case ie.UserLocationInformation:
			switch i.Instance() {
			case 0:
				m.ULI = i
			case 1:
				m.ULIForSGW = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.ServingNetwork:
			m.ServingNetwork = i
		case ie.RATType:
			m.RATType = i
		case ie.Indication:
			m.IndicationFlags = i
		case ie.FullyQualifiedTEID:
			m.SenderFTEIDC = i
		case ie.AggregateMaximumBitRate:
			m.AMBR = i
		case ie.DelayValue:
			m.DelayDownlinkPacketNotificationRequest = i
		case ie.BearerContext:
			switch i.Instance() {
			case 0:
				m.BearerContextsToBeModified = append(m.BearerContextsToBeModified, i)
			case 1:
				m.BearerContextsToBeRemoved = append(m.BearerContextsToBeRemoved, i)
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.Recovery:
			m.Recovery = i
		case ie.UETimeZone:
			m.UETimeZone = i
		case ie.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				m.MMEFQCSID = i
			case 1:
				m.SGWFQCSID = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.UserCSGInformation:
			m.UCI = i
		case ie.IPAddress:
			switch i.Instance() {
			case 0:
				m.UELocalIPAddress = i
			case 1:
				m.HeNBLocalIPAddress = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.PortNumber:
			switch i.Instance() {
			case 0:
				m.UEUDPPort = i
			case 1:
				m.HeNBUDPPort = i
			case 2:
				m.MMESGSNIdentifier = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.LocalDistinguishedName:
			switch i.Instance() {
			case 0:
				m.MMESGSNLDN = i
			case 1:
				m.SGWLDN = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.CNOperatorSelectionEntity:
			m.CNOperatorSelectionEntity = i
		case ie.PresenceReportingAreaInformation:
			m.PresenceReportingAreaInformation = append(m.PresenceReportingAreaInformation, i)
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				m.MMESGSNOverloadControlInformation = i
			case 1:
				m.SGWOverloadControlInformation = i
			case 2:
				m.EPDGOverloadControlInformation = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.ServingPLMNRateControl:
			m.ServingPLMNRateControl = i
		case ie.Counter:
			m.MOExceptionDataCounter = i
		case ie.IMSI:
			m.IMSI = i
		case ie.TWANIdentifier:
			m.WLANLocationInformation = i
		case ie.TWANIdentifierTimestamp:
			m.WLANLocationTimeStamp = i
		case ie.SecondaryRATUsageDataReport:
			m.SecondaryRATUsageDataReport = append(m.SecondaryRATUsageDataReport, i)
		case ie.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	m.SetLength()
	return m
}

// Marshal serializes ModifyBearerRequest into bytes.
func (m *ModifyBearerRequest) Marshal() ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes ModifyBearerRequest into bytes.
func (m *ModifyBearerRequest) MarshalTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.MarshalLen()-m.Header.MarshalLen())

	offset := 0
	if ie := m.MEI; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.ULI; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.ServingNetwork; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.RATType; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.IndicationFlags; ie != nil {
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
	if ie := m.AMBR; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.DelayDownlinkPacketNotificationRequest; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	for _, ie := range m.BearerContextsToBeModified {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	for _, ie := range m.BearerContextsToBeRemoved {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.Recovery; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.UETimeZone; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
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
	if ie := m.UCI; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.UELocalIPAddress; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.UEUDPPort; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.MMESGSNLDN; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.SGWLDN; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.HeNBLocalIPAddress; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.HeNBUDPPort; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.MMESGSNIdentifier; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.CNOperatorSelectionEntity; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	for _, ie := range m.PresenceReportingAreaInformation {
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
	if ie := m.EPDGOverloadControlInformation; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.ServingPLMNRateControl; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.MOExceptionDataCounter; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.IMSI; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.ULIForSGW; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.WLANLocationInformation; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.WLANLocationTimeStamp; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	for _, ie := range m.SecondaryRATUsageDataReport {
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

// ParseModifyBearerRequest decodes given bytes as ModifyBearerRequest.
func ParseModifyBearerRequest(b []byte) (*ModifyBearerRequest, error) {
	c := &ModifyBearerRequest{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes given bytes as ModifyBearerRequest.
func (m *ModifyBearerRequest) UnmarshalBinary(b []byte) error {
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
		case ie.MobileEquipmentIdentity:
			m.MEI = i
		case ie.UserLocationInformation:
			switch i.Instance() {
			case 0:
				m.ULI = i
			case 1:
				m.ULIForSGW = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.ServingNetwork:
			m.ServingNetwork = i
		case ie.RATType:
			m.RATType = i
		case ie.Indication:
			m.IndicationFlags = i
		case ie.FullyQualifiedTEID:
			m.SenderFTEIDC = i
		case ie.AggregateMaximumBitRate:
			m.AMBR = i
		case ie.DelayValue:
			m.DelayDownlinkPacketNotificationRequest = i
		case ie.BearerContext:
			switch i.Instance() {
			case 0:
				m.BearerContextsToBeModified = append(m.BearerContextsToBeModified, i)
			case 1:
				m.BearerContextsToBeRemoved = append(m.BearerContextsToBeRemoved, i)
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.Recovery:
			m.Recovery = i
		case ie.UETimeZone:
			m.UETimeZone = i
		case ie.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				m.MMEFQCSID = i
			case 1:
				m.SGWFQCSID = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.UserCSGInformation:
			m.UCI = i
		case ie.IPAddress:
			switch i.Instance() {
			case 0:
				m.UELocalIPAddress = i
			case 1:
				m.HeNBLocalIPAddress = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.PortNumber:
			switch i.Instance() {
			case 0:
				m.UEUDPPort = i
			case 1:
				m.HeNBUDPPort = i
			case 2:
				m.MMESGSNIdentifier = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.LocalDistinguishedName:
			switch i.Instance() {
			case 0:
				m.MMESGSNLDN = i
			case 1:
				m.SGWLDN = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.CNOperatorSelectionEntity:
			m.CNOperatorSelectionEntity = i
		case ie.PresenceReportingAreaInformation:
			m.PresenceReportingAreaInformation = append(m.PresenceReportingAreaInformation, i)
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				m.MMESGSNOverloadControlInformation = i
			case 1:
				m.SGWOverloadControlInformation = i
			case 2:
				m.EPDGOverloadControlInformation = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ie.ServingPLMNRateControl:
			m.ServingPLMNRateControl = i
		case ie.Counter:
			m.MOExceptionDataCounter = i
		case ie.IMSI:
			m.IMSI = i
		case ie.TWANIdentifier:
			m.WLANLocationInformation = i
		case ie.TWANIdentifierTimestamp:
			m.WLANLocationTimeStamp = i
		case ie.SecondaryRATUsageDataReport:
			m.SecondaryRATUsageDataReport = append(m.SecondaryRATUsageDataReport, i)
		case ie.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (m *ModifyBearerRequest) MarshalLen() int {
	l := m.Header.MarshalLen() - len(m.Header.Payload)

	if ie := m.MEI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.ULI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.ServingNetwork; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.RATType; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.SenderFTEIDC; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.AMBR; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.DelayDownlinkPacketNotificationRequest; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range m.BearerContextsToBeModified {
		l += ie.MarshalLen()
	}
	for _, ie := range m.BearerContextsToBeRemoved {
		l += ie.MarshalLen()
	}
	if ie := m.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.UETimeZone; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.MMEFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.SGWFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.UCI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.UELocalIPAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.UEUDPPort; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.MMESGSNLDN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.SGWLDN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.HeNBLocalIPAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.HeNBUDPPort; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.MMESGSNIdentifier; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.CNOperatorSelectionEntity; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range m.PresenceReportingAreaInformation {
		l += ie.MarshalLen()
	}
	if ie := m.MMESGSNOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.SGWOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.EPDGOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.ServingPLMNRateControl; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.MOExceptionDataCounter; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.IMSI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.ULIForSGW; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.WLANLocationInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.WLANLocationTimeStamp; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range m.SecondaryRATUsageDataReport {
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
func (m *ModifyBearerRequest) SetLength() {
	m.Header.Length = uint16(m.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (m *ModifyBearerRequest) MessageTypeName() string {
	return "Modify Bearer Request"
}

// TEID returns the TEID in uint32.
func (m *ModifyBearerRequest) TEID() uint32 {
	return m.Header.teid()
}
