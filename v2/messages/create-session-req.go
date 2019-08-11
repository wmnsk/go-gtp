// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v2/ies"
)

// CreateSessionRequest is a CreateSessionRequest Header and its IEs above.
type CreateSessionRequest struct {
	*Header
	IMSI                               *ies.IE
	MSISDN                             *ies.IE
	MEI                                *ies.IE
	ULI                                *ies.IE
	ServingNetwork                     *ies.IE
	RATType                            *ies.IE
	IndicationFlags                    *ies.IE
	SenderFTEIDC                       *ies.IE
	PGWS5S8FTEIDC                      *ies.IE
	APN                                *ies.IE
	SelectionMode                      *ies.IE
	PDNType                            *ies.IE
	PAA                                *ies.IE
	APNRestriction                     *ies.IE
	AMBR                               *ies.IE
	LinkedEBI                          *ies.IE
	TWMI                               *ies.IE
	PCO                                *ies.IE
	BearerContextsToBeCreated          *ies.IE
	BearerContextsToBeRemoved          *ies.IE
	TraceInformation                   *ies.IE
	Recovery                           *ies.IE
	MMEFQCSID                          *ies.IE
	SGWFQCSID                          *ies.IE
	EPDGFQCSID                         *ies.IE
	TWANFQCSID                         *ies.IE
	UETimeZone                         *ies.IE
	UCI                                *ies.IE
	ChargingCharacteristics            *ies.IE
	MMESGSNLDN                         *ies.IE
	SGWLDN                             *ies.IE
	EPDGLDN                            *ies.IE
	TWANLDN                            *ies.IE
	SignallingPriorityIndication       *ies.IE
	UELocalIPAddress                   *ies.IE
	UEUDPPort                          *ies.IE
	APCO                               *ies.IE
	HeNBLocalIPAddress                 *ies.IE
	HeNBUDPPort                        *ies.IE
	MMESGSNIdentifier                  *ies.IE
	TWANIdentifier                     *ies.IE
	EPDGIPAddress                      *ies.IE
	CNOperatorSelectionEntity          *ies.IE
	PresenceReportingAreaInformation   *ies.IE
	MMESGSNOverloadControlInformation  *ies.IE
	SGWOverloadControlInformation      *ies.IE
	TWANePDGOverloadControlInformation *ies.IE
	OriginationTimeStamp               *ies.IE
	MaximumWaitTime                    *ies.IE
	WLANLocationInformation            *ies.IE
	WLANLocationTimeStamp              *ies.IE
	NBIFOMContainer                    *ies.IE
	RemoteUEContextConnected           *ies.IE
	TGPPAAAServerIdentifier            *ies.IE
	EPCO                               *ies.IE
	ServingPLMNRateControl             *ies.IE
	MOExceptionDataCounter             *ies.IE
	UETCPPort                          *ies.IE
	MappedUEUsageType                  *ies.IE
	ULIForSGW                          *ies.IE
	SGWUNodeName                       *ies.IE
	SecondaryRATUsageDataReport        *ies.IE
	UPFunctionSelectionIndicationFlags *ies.IE
	APNRateControlStatus               *ies.IE
	PrivateExtension                   *ies.IE
	AdditionalIEs                      []*ies.IE
}

// NewCreateSessionRequest creates a new CreateSessionRequest.
func NewCreateSessionRequest(teid, seq uint32, ie ...*ies.IE) *CreateSessionRequest {
	c := &CreateSessionRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeCreateSessionRequest, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.IMSI:
			c.IMSI = i
		case ies.MSISDN:
			c.MSISDN = i
		case ies.MobileEquipmentIdentity:
			c.MEI = i
		case ies.UserLocationInformation:
			switch i.Instance() {
			case 0:
				c.ULI = i
			case 1:
				c.ULIForSGW = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.ServingNetwork:
			c.ServingNetwork = i
		case ies.RATType:
			c.RATType = i
		case ies.Indication:
			c.IndicationFlags = i
		case ies.FullyQualifiedTEID:
			switch i.Instance() {
			case 0:
				c.SenderFTEIDC = i
			case 1:
				c.PGWS5S8FTEIDC = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.AccessPointName:
			c.APN = i
		case ies.SelectionMode:
			c.SelectionMode = i
		case ies.PDNType:
			c.PDNType = i
		case ies.PDNAddressAllocation:
			c.PAA = i
		case ies.APNRestriction:
			c.APNRestriction = i
		case ies.AggregateMaximumBitRate:
			c.AMBR = i
		case ies.EPSBearerID:
			c.LinkedEBI = i
		case ies.TrustedWLANModeIndication:
			c.TWMI = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				c.BearerContextsToBeCreated = i
			case 1:
				c.BearerContextsToBeRemoved = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				c.MMEFQCSID = i
			case 1:
				c.SGWFQCSID = i
			case 2:
				c.EPDGFQCSID = i
			case 3:
				c.TWANFQCSID = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.TraceInformation:
			c.TraceInformation = i
		case ies.Recovery:
			c.Recovery = i
		case ies.UETimeZone:
			c.UETimeZone = i
		case ies.UserCSGInformation:
			c.UCI = i
		case ies.ChargingCharacteristics:
			c.ChargingCharacteristics = i
		case ies.LocalDistinguishedName:
			switch i.Instance() {
			case 0:
				c.MMESGSNLDN = i
			case 1:
				c.SGWLDN = i
			case 2:
				c.EPDGLDN = i
			case 3:
				c.TWANLDN = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.SignallingPriorityIndication:
			c.SignallingPriorityIndication = i
		case ies.IPAddress:
			switch i.Instance() {
			case 0:
				c.UELocalIPAddress = i
			case 1:
				c.HeNBLocalIPAddress = i
			case 2:
				c.MMESGSNIdentifier = i
			case 3:
				c.EPDGIPAddress = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.PortNumber:
			switch i.Instance() {
			case 0:
				c.UEUDPPort = i
			case 1:
				c.HeNBUDPPort = i
			case 2:
				c.UETCPPort = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.AdditionalProtocolConfigurationOptions:
			c.APCO = i
		case ies.TWANIdentifier:
			switch i.Instance() {
			case 0:
				c.TWANIdentifier = i
			case 1:
				c.WLANLocationInformation = i
			}
		case ies.CNOperatorSelectionEntity:
			c.CNOperatorSelectionEntity = i
		case ies.PresenceReportingAreaInformation:
			c.PresenceReportingAreaInformation = i
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				c.MMESGSNOverloadControlInformation = i
			case 1:
				c.SGWOverloadControlInformation = i
			case 2:
				c.TWANePDGOverloadControlInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.MillisecondTimeStamp:
			c.OriginationTimeStamp = i
		case ies.IntegerNumber:
			c.MaximumWaitTime = i
		case ies.TWANIdentifierTimestamp:
			c.WLANLocationTimeStamp = i
		case ies.FContainer:
			c.NBIFOMContainer = i
		case ies.RemoteUEContext:
			c.RemoteUEContextConnected = i
		case ies.NodeIdentifier:
			c.TGPPAAAServerIdentifier = i
		case ies.ExtendedProtocolConfigurationOptions:
			c.EPCO = i
		case ies.ServingPLMNRateControl:
			c.ServingPLMNRateControl = i
		case ies.Counter:
			c.MOExceptionDataCounter = i
		case ies.MappedUEUsageType:
			c.MappedUEUsageType = i
		case ies.FullyQualifiedDomainName:
			c.SGWUNodeName = i
		case ies.SecondaryRATUsageDataReport:
			c.SecondaryRATUsageDataReport = i
		case ies.UPFunctionSelectionIndicationFlags:
			c.UPFunctionSelectionIndicationFlags = i
		case ies.APNRateControlStatus:
			c.APNRateControlStatus = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Marshal serializes CreateSessionRequest into bytes.
func (c *CreateSessionRequest) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes CreateSessionRequest into bytes.
func (c *CreateSessionRequest) MarshalTo(b []byte) error {
	if c.Header.Payload != nil {
		c.Header.Payload = nil
	}
	c.Header.Payload = make([]byte, c.MarshalLen()-c.Header.MarshalLen())

	offset := 0
	if ie := c.IMSI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MSISDN; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MEI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ULI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ServingNetwork; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.RATType; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.IndicationFlags; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SenderFTEIDC; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PGWS5S8FTEIDC; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.APN; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SelectionMode; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PDNType; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PAA; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.APNRestriction; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.AMBR; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.LinkedEBI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TWMI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PCO; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.BearerContextsToBeCreated; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.BearerContextsToBeRemoved; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TraceInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.Recovery; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MMEFQCSID; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGWFQCSID; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.EPDGFQCSID; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TWANFQCSID; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.UETimeZone; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.UCI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ChargingCharacteristics; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MMESGSNLDN; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGWLDN; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.EPDGLDN; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TWANLDN; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SignallingPriorityIndication; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.UELocalIPAddress; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.UEUDPPort; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.APCO; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.HeNBLocalIPAddress; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.HeNBUDPPort; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MMESGSNIdentifier; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TWANIdentifier; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.EPDGIPAddress; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.CNOperatorSelectionEntity; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PresenceReportingAreaInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MMESGSNOverloadControlInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGWOverloadControlInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TWANePDGOverloadControlInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.OriginationTimeStamp; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MaximumWaitTime; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.WLANLocationInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.WLANLocationTimeStamp; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.NBIFOMContainer; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.RemoteUEContextConnected; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TGPPAAAServerIdentifier; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.EPCO; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ServingPLMNRateControl; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MOExceptionDataCounter; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.UETCPPort; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MappedUEUsageType; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ULIForSGW; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGWUNodeName; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SecondaryRATUsageDataReport; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.UPFunctionSelectionIndicationFlags; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.APNRateControlStatus; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
	}

	for _, ie := range c.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(c.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	c.Header.SetLength()
	return c.Header.MarshalTo(b)
}

// ParseCreateSessionRequest decodes given bytes as CreateSessionRequest.
func ParseCreateSessionRequest(b []byte) (*CreateSessionRequest, error) {
	c := &CreateSessionRequest{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes given bytes as CreateSessionRequest.
func (c *CreateSessionRequest) UnmarshalBinary(b []byte) error {
	var err error
	c.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(c.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ies.ParseMultiIEs(c.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.IMSI:
			c.IMSI = i
		case ies.MSISDN:
			c.MSISDN = i
		case ies.MobileEquipmentIdentity:
			c.MEI = i
		case ies.UserLocationInformation:
			switch i.Instance() {
			case 0:
				c.ULI = i
			case 1:
				c.ULIForSGW = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.ServingNetwork:
			c.ServingNetwork = i
		case ies.RATType:
			c.RATType = i
		case ies.Indication:
			c.IndicationFlags = i
		case ies.FullyQualifiedTEID:
			switch i.Instance() {
			case 0:
				c.SenderFTEIDC = i
			case 1:
				c.PGWS5S8FTEIDC = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.AccessPointName:
			c.APN = i
		case ies.SelectionMode:
			c.SelectionMode = i
		case ies.PDNType:
			c.PDNType = i
		case ies.PDNAddressAllocation:
			c.PAA = i
		case ies.APNRestriction:
			c.APNRestriction = i
		case ies.AggregateMaximumBitRate:
			c.AMBR = i
		case ies.EPSBearerID:
			c.LinkedEBI = i
		case ies.TrustedWLANModeIndication:
			c.TWMI = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				c.BearerContextsToBeCreated = i
			case 1:
				c.BearerContextsToBeRemoved = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				c.MMEFQCSID = i
			case 1:
				c.SGWFQCSID = i
			case 2:
				c.EPDGFQCSID = i
			case 3:
				c.TWANFQCSID = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.TraceInformation:
			c.TraceInformation = i
		case ies.Recovery:
			c.Recovery = i
		case ies.UETimeZone:
			c.UETimeZone = i
		case ies.UserCSGInformation:
			c.UCI = i
		case ies.ChargingCharacteristics:
			c.ChargingCharacteristics = i
		case ies.LocalDistinguishedName:
			switch i.Instance() {
			case 0:
				c.MMESGSNLDN = i
			case 1:
				c.SGWLDN = i
			case 2:
				c.EPDGLDN = i
			case 3:
				c.TWANLDN = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.SignallingPriorityIndication:
			c.SignallingPriorityIndication = i
		case ies.IPAddress:
			switch i.Instance() {
			case 0:
				c.UELocalIPAddress = i
			case 1:
				c.HeNBLocalIPAddress = i
			case 2:
				c.MMESGSNIdentifier = i
			case 3:
				c.EPDGIPAddress = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.PortNumber:
			switch i.Instance() {
			case 0:
				c.UEUDPPort = i
			case 1:
				c.HeNBUDPPort = i
			case 2:
				c.UETCPPort = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.AdditionalProtocolConfigurationOptions:
			c.APCO = i
		case ies.TWANIdentifier:
			switch i.Instance() {
			case 0:
				c.TWANIdentifier = i
			case 1:
				c.WLANLocationInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.CNOperatorSelectionEntity:
			c.CNOperatorSelectionEntity = i
		case ies.PresenceReportingAreaInformation:
			c.PresenceReportingAreaInformation = i
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				c.MMESGSNOverloadControlInformation = i
			case 1:
				c.SGWOverloadControlInformation = i
			case 2:
				c.TWANePDGOverloadControlInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.MillisecondTimeStamp:
			c.OriginationTimeStamp = i
		case ies.IntegerNumber:
			c.MaximumWaitTime = i
		case ies.TWANIdentifierTimestamp:
			c.WLANLocationTimeStamp = i
		case ies.FContainer:
			c.NBIFOMContainer = i
		case ies.RemoteUEContext:
			c.RemoteUEContextConnected = i
		case ies.NodeIdentifier:
			c.TGPPAAAServerIdentifier = i
		case ies.ExtendedProtocolConfigurationOptions:
			c.EPCO = i
		case ies.ServingPLMNRateControl:
			c.ServingPLMNRateControl = i
		case ies.Counter:
			c.MOExceptionDataCounter = i
		case ies.MappedUEUsageType:
			c.MappedUEUsageType = i
		case ies.FullyQualifiedDomainName:
			c.SGWUNodeName = i
		case ies.SecondaryRATUsageDataReport:
			c.SecondaryRATUsageDataReport = i
		case ies.UPFunctionSelectionIndicationFlags:
			c.UPFunctionSelectionIndicationFlags = i
		case ies.APNRateControlStatus:
			c.APNRateControlStatus = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (c *CreateSessionRequest) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)

	if ie := c.IMSI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MSISDN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MEI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ULI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ServingNetwork; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.RATType; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SenderFTEIDC; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PGWS5S8FTEIDC; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.APN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SelectionMode; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PDNType; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PAA; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.APNRestriction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.AMBR; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.LinkedEBI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TWMI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.BearerContextsToBeCreated; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.BearerContextsToBeRemoved; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TraceInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MMEFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.EPDGFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TWANFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UETimeZone; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UCI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ChargingCharacteristics; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MMESGSNLDN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWLDN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.EPDGLDN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TWANLDN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SignallingPriorityIndication; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UELocalIPAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UEUDPPort; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.APCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.HeNBLocalIPAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.HeNBUDPPort; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MMESGSNIdentifier; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TWANIdentifier; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.EPDGIPAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.CNOperatorSelectionEntity; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PresenceReportingAreaInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MMESGSNOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TWANePDGOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.OriginationTimeStamp; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MaximumWaitTime; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.WLANLocationInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.WLANLocationTimeStamp; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.NBIFOMContainer; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.RemoteUEContextConnected; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TGPPAAAServerIdentifier; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.EPCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ServingPLMNRateControl; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MOExceptionDataCounter; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UETCPPort; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MappedUEUsageType; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ULIForSGW; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWUNodeName; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SecondaryRATUsageDataReport; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UPFunctionSelectionIndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.APNRateControlStatus; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}

	for _, ie := range c.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length field.
func (c *CreateSessionRequest) SetLength() {
	c.Header.Length = uint16(c.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *CreateSessionRequest) MessageTypeName() string {
	return "Create Session Request"
}

// TEID returns the TEID in uint32.
func (c *CreateSessionRequest) TEID() uint32 {
	return c.Header.teid()
}
