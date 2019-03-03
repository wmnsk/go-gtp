// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v2/ies"
)

// ModifyBearerRequest is a ModifyBearerRequest Header and its IEs above.
type ModifyBearerRequest struct {
	*Header
	MEI                                    *ies.IE
	ULI                                    *ies.IE
	ServingNetwork                         *ies.IE
	RATType                                *ies.IE
	IndicationFlags                        *ies.IE
	SenderFTEIDC                           *ies.IE
	AMBR                                   *ies.IE
	DelayDownlinkPacketNotificationRequest *ies.IE
	BearerContextsToBeModified             *ies.IE
	BearerContextsTobeRemoved              *ies.IE
	Recovery                               *ies.IE
	UETimeZone                             *ies.IE
	MMEFQCSID                              *ies.IE
	SGWFQCSID                              *ies.IE
	UCI                                    *ies.IE
	UELocalIPAddress                       *ies.IE
	UEUDPPort                              *ies.IE
	MMESGSNLDN                             *ies.IE
	SGWLDN                                 *ies.IE
	HeNBLocalIPAddress                     *ies.IE
	HeNBUDPPort                            *ies.IE
	MMESGSNIdentifier                      *ies.IE
	CNOperatorSelectionEntity              *ies.IE
	PresenceReportingAreaInformation       *ies.IE
	MMESGSNOverloadControlInformation      *ies.IE
	SGWOverloadControlInformation          *ies.IE
	EPDGOverloadControlInformation         *ies.IE
	ServingPLMNRateControl                 *ies.IE
	MOExceptionDataCounter                 *ies.IE
	IMSI                                   *ies.IE
	ULIForSGW                              *ies.IE
	WLANLocationInformation                *ies.IE
	WLANLocationTimeStamp                  *ies.IE
	SecondaryRATUsageDataReport            *ies.IE
	PrivateExtension                       *ies.IE
	AdditionalIEs                          []*ies.IE
}

// NewModifyBearerRequest creates a new ModifyBearerRequest.
func NewModifyBearerRequest(teid, seq uint32, ie ...*ies.IE) *ModifyBearerRequest {
	m := &ModifyBearerRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeModifyBearerRequest, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.MobileEquipmentIdentity:
			m.MEI = i
		case ies.UserLocationInformation:
			switch i.Instance() {
			case 0:
				m.ULI = i
			case 1:
				m.ULIForSGW = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ies.ServingNetwork:
			m.ServingNetwork = i
		case ies.RATType:
			m.RATType = i
		case ies.Indication:
			m.IndicationFlags = i
		case ies.FullyQualifiedTEID:
			m.SenderFTEIDC = i
		case ies.AggregateMaximumBitRate:
			m.AMBR = i
		case ies.DelayValue:
			m.DelayDownlinkPacketNotificationRequest = i
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				m.BearerContextsToBeModified = i
			case 1:
				m.BearerContextsTobeRemoved = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ies.Recovery:
			m.Recovery = i
		case ies.UETimeZone:
			m.UETimeZone = i
		case ies.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				m.MMEFQCSID = i
			case 1:
				m.SGWFQCSID = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ies.UserCSGInformation:
			m.UCI = i
		case ies.IPAddress:
			switch i.Instance() {
			case 0:
				m.UELocalIPAddress = i
			case 1:
				m.HeNBLocalIPAddress = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ies.PortNumber:
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
		case ies.LocalDistinguishedName:
			switch i.Instance() {
			case 0:
				m.MMESGSNLDN = i
			case 1:
				m.SGWLDN = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ies.CNOperatorSelectionEntity:
			m.CNOperatorSelectionEntity = i
		case ies.PresenceReportingAreaInformation:
			m.PresenceReportingAreaInformation = i
		case ies.OverloadControlInformation:
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
		case ies.ServingPLMNRateControl:
			m.ServingPLMNRateControl = i
		case ies.Counter:
			m.MOExceptionDataCounter = i
		case ies.IMSI:
			m.IMSI = i
		case ies.TWANIdentifier:
			m.WLANLocationInformation = i
		case ies.TWANIdentifierTimestamp:
			m.WLANLocationTimeStamp = i
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

// Serialize serializes ModifyBearerRequest into bytes.
func (m *ModifyBearerRequest) Serialize() ([]byte, error) {
	b := make([]byte, m.Len())
	if err := m.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes ModifyBearerRequest into bytes.
func (m *ModifyBearerRequest) SerializeTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.Len()-m.Header.Len())

	offset := 0
	if ie := m.MEI; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.ULI; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.ServingNetwork; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.RATType; ie != nil {
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
	if ie := m.SenderFTEIDC; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.AMBR; ie != nil {
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
	if ie := m.BearerContextsTobeRemoved; ie != nil {
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
	if ie := m.UETimeZone; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.MMEFQCSID; ie != nil {
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
	if ie := m.UCI; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.UELocalIPAddress; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.UEUDPPort; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.MMESGSNLDN; ie != nil {
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
	if ie := m.HeNBLocalIPAddress; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.HeNBUDPPort; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.MMESGSNIdentifier; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.CNOperatorSelectionEntity; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.PresenceReportingAreaInformation; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.MMESGSNOverloadControlInformation; ie != nil {
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
	if ie := m.EPDGOverloadControlInformation; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.ServingPLMNRateControl; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.MOExceptionDataCounter; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.IMSI; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.ULIForSGW; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.WLANLocationInformation; ie != nil {
		if err := ie.SerializeTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := m.WLANLocationTimeStamp; ie != nil {
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

// DecodeModifyBearerRequest decodes given bytes as ModifyBearerRequest.
func DecodeModifyBearerRequest(b []byte) (*ModifyBearerRequest, error) {
	c := &ModifyBearerRequest{}
	if err := c.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return c, nil
}

// DecodeFromBytes decodes given bytes as ModifyBearerRequest.
func (m *ModifyBearerRequest) DecodeFromBytes(b []byte) error {
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
		case ies.MobileEquipmentIdentity:
			m.MEI = i
		case ies.UserLocationInformation:
			switch i.Instance() {
			case 0:
				m.ULI = i
			case 1:
				m.ULIForSGW = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ies.ServingNetwork:
			m.ServingNetwork = i
		case ies.RATType:
			m.RATType = i
		case ies.Indication:
			m.IndicationFlags = i
		case ies.FullyQualifiedTEID:
			m.SenderFTEIDC = i
		case ies.AggregateMaximumBitRate:
			m.AMBR = i
		case ies.DelayValue:
			m.DelayDownlinkPacketNotificationRequest = i
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				m.BearerContextsToBeModified = i
			case 1:
				m.BearerContextsTobeRemoved = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ies.Recovery:
			m.Recovery = i
		case ies.UETimeZone:
			m.UETimeZone = i
		case ies.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				m.MMEFQCSID = i
			case 1:
				m.SGWFQCSID = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ies.UserCSGInformation:
			m.UCI = i
		case ies.IPAddress:
			switch i.Instance() {
			case 0:
				m.UELocalIPAddress = i
			case 1:
				m.HeNBLocalIPAddress = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ies.PortNumber:
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
		case ies.LocalDistinguishedName:
			switch i.Instance() {
			case 0:
				m.MMESGSNLDN = i
			case 1:
				m.SGWLDN = i
			default:
				m.AdditionalIEs = append(m.AdditionalIEs, i)
			}
		case ies.CNOperatorSelectionEntity:
			m.CNOperatorSelectionEntity = i
		case ies.PresenceReportingAreaInformation:
			m.PresenceReportingAreaInformation = i
		case ies.OverloadControlInformation:
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
		case ies.ServingPLMNRateControl:
			m.ServingPLMNRateControl = i
		case ies.Counter:
			m.MOExceptionDataCounter = i
		case ies.IMSI:
			m.IMSI = i
		case ies.TWANIdentifier:
			m.WLANLocationInformation = i
		case ies.TWANIdentifierTimestamp:
			m.WLANLocationTimeStamp = i
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
func (m *ModifyBearerRequest) Len() int {
	l := m.Header.Len() - len(m.Header.Payload)

	if ie := m.MEI; ie != nil {
		l += ie.Len()
	}
	if ie := m.ULI; ie != nil {
		l += ie.Len()
	}
	if ie := m.ServingNetwork; ie != nil {
		l += ie.Len()
	}
	if ie := m.RATType; ie != nil {
		l += ie.Len()
	}
	if ie := m.IndicationFlags; ie != nil {
		l += ie.Len()
	}
	if ie := m.SenderFTEIDC; ie != nil {
		l += ie.Len()
	}
	if ie := m.AMBR; ie != nil {
		l += ie.Len()
	}
	if ie := m.DelayDownlinkPacketNotificationRequest; ie != nil {
		l += ie.Len()
	}
	if ie := m.BearerContextsToBeModified; ie != nil {
		l += ie.Len()
	}
	if ie := m.BearerContextsTobeRemoved; ie != nil {
		l += ie.Len()
	}
	if ie := m.Recovery; ie != nil {
		l += ie.Len()
	}
	if ie := m.UETimeZone; ie != nil {
		l += ie.Len()
	}
	if ie := m.MMEFQCSID; ie != nil {
		l += ie.Len()
	}
	if ie := m.SGWFQCSID; ie != nil {
		l += ie.Len()
	}
	if ie := m.UCI; ie != nil {
		l += ie.Len()
	}
	if ie := m.UELocalIPAddress; ie != nil {
		l += ie.Len()
	}
	if ie := m.UEUDPPort; ie != nil {
		l += ie.Len()
	}
	if ie := m.MMESGSNLDN; ie != nil {
		l += ie.Len()
	}
	if ie := m.SGWLDN; ie != nil {
		l += ie.Len()
	}
	if ie := m.HeNBLocalIPAddress; ie != nil {
		l += ie.Len()
	}
	if ie := m.HeNBUDPPort; ie != nil {
		l += ie.Len()
	}
	if ie := m.MMESGSNIdentifier; ie != nil {
		l += ie.Len()
	}
	if ie := m.CNOperatorSelectionEntity; ie != nil {
		l += ie.Len()
	}
	if ie := m.PresenceReportingAreaInformation; ie != nil {
		l += ie.Len()
	}
	if ie := m.MMESGSNOverloadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := m.SGWOverloadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := m.EPDGOverloadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := m.ServingPLMNRateControl; ie != nil {
		l += ie.Len()
	}
	if ie := m.MOExceptionDataCounter; ie != nil {
		l += ie.Len()
	}
	if ie := m.IMSI; ie != nil {
		l += ie.Len()
	}
	if ie := m.ULIForSGW; ie != nil {
		l += ie.Len()
	}
	if ie := m.WLANLocationInformation; ie != nil {
		l += ie.Len()
	}
	if ie := m.WLANLocationTimeStamp; ie != nil {
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
func (m *ModifyBearerRequest) SetLength() {
	m.Header.Length = uint16(m.Len() - 4)
}

// MessageTypeName returns the name of protocol.
func (m *ModifyBearerRequest) MessageTypeName() string {
	return "Modify Bearer Request"
}

// TEID returns the TEID in uint32.
func (m *ModifyBearerRequest) TEID() uint32 {
	return m.Header.teid()
}
