// Copyright 2019-2021 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// CreateSessionRequest is a CreateSessionRequest Header and its IEs above.
type CreateSessionRequest struct {
	*Header
	IMSI                               *ie.IE
	MSISDN                             *ie.IE
	MEI                                *ie.IE
	ULI                                *ie.IE
	ServingNetwork                     *ie.IE
	RATType                            *ie.IE
	IndicationFlags                    *ie.IE
	SenderFTEIDC                       *ie.IE
	PGWS5S8FTEIDC                      *ie.IE
	APN                                *ie.IE
	SelectionMode                      *ie.IE
	PDNType                            *ie.IE
	PAA                                *ie.IE
	APNRestriction                     *ie.IE
	AMBR                               *ie.IE
	LinkedEBI                          *ie.IE
	TWMI                               *ie.IE
	PCO                                *ie.IE
	BearerContextsToBeCreated          []*ie.IE
	BearerContextsToBeRemoved          []*ie.IE
	TraceInformation                   *ie.IE
	Recovery                           *ie.IE
	MMEFQCSID                          *ie.IE
	SGWFQCSID                          *ie.IE
	EPDGFQCSID                         *ie.IE
	TWANFQCSID                         *ie.IE
	UETimeZone                         *ie.IE
	UCI                                *ie.IE
	ChargingCharacteristics            *ie.IE
	MMESGSNLDN                         *ie.IE
	SGWLDN                             *ie.IE
	EPDGLDN                            *ie.IE
	TWANLDN                            *ie.IE
	SignallingPriorityIndication       *ie.IE
	UELocalIPAddress                   *ie.IE
	UEUDPPort                          *ie.IE
	APCO                               *ie.IE
	HeNBLocalIPAddress                 *ie.IE
	HeNBUDPPort                        *ie.IE
	MMESGSNIdentifier                  *ie.IE
	TWANIdentifier                     *ie.IE
	EPDGIPAddress                      *ie.IE
	CNOperatorSelectionEntity          *ie.IE
	PresenceReportingAreaInformation   []*ie.IE
	MMESGSNOverloadControlInformation  *ie.IE
	SGWOverloadControlInformation      *ie.IE
	TWANePDGOverloadControlInformation *ie.IE
	OriginationTimeStamp               *ie.IE
	MaximumWaitTime                    *ie.IE
	WLANLocationInformation            *ie.IE
	WLANLocationTimeStamp              *ie.IE
	NBIFOMContainer                    *ie.IE
	RemoteUEContextConnected           []*ie.IE
	TGPPAAAServerIdentifier            *ie.IE
	EPCO                               *ie.IE
	ServingPLMNRateControl             *ie.IE
	MOExceptionDataCounter             *ie.IE
	UETCPPort                          *ie.IE
	MappedUEUsageType                  *ie.IE
	ULIForSGW                          *ie.IE
	SGWUNodeName                       *ie.IE
	SecondaryRATUsageDataReport        []*ie.IE
	UPFunctionSelectionIndicationFlags *ie.IE
	APNRateControlStatus               *ie.IE
	PrivateExtension                   *ie.IE
	AdditionalIEs                      []*ie.IE
}

// NewCreateSessionRequest creates a new CreateSessionRequest.
func NewCreateSessionRequest(teid, seq uint32, ies ...*ie.IE) *CreateSessionRequest {
	c := &CreateSessionRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeCreateSessionRequest, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.IMSI:
			c.IMSI = i
		case ie.MSISDN:
			c.MSISDN = i
		case ie.MobileEquipmentIdentity:
			c.MEI = i
		case ie.UserLocationInformation:
			switch i.Instance() {
			case 0:
				c.ULI = i
			case 1:
				c.ULIForSGW = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.ServingNetwork:
			c.ServingNetwork = i
		case ie.RATType:
			c.RATType = i
		case ie.Indication:
			c.IndicationFlags = i
		case ie.FullyQualifiedTEID:
			switch i.Instance() {
			case 0:
				c.SenderFTEIDC = i
			case 1:
				c.PGWS5S8FTEIDC = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.AccessPointName:
			c.APN = i
		case ie.SelectionMode:
			c.SelectionMode = i
		case ie.PDNType:
			c.PDNType = i
		case ie.PDNAddressAllocation:
			c.PAA = i
		case ie.APNRestriction:
			c.APNRestriction = i
		case ie.AggregateMaximumBitRate:
			c.AMBR = i
		case ie.EPSBearerID:
			c.LinkedEBI = i
		case ie.TrustedWLANModeIndication:
			c.TWMI = i
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.BearerContext:
			switch i.Instance() {
			case 0:
				c.BearerContextsToBeCreated = append(c.BearerContextsToBeCreated, i)
			case 1:
				c.BearerContextsToBeRemoved = append(c.BearerContextsToBeCreated, i)
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.FullyQualifiedCSID:
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
		case ie.TraceInformation:
			c.TraceInformation = i
		case ie.Recovery:
			c.Recovery = i
		case ie.UETimeZone:
			c.UETimeZone = i
		case ie.UserCSGInformation:
			c.UCI = i
		case ie.ChargingCharacteristics:
			c.ChargingCharacteristics = i
		case ie.LocalDistinguishedName:
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
		case ie.SignallingPriorityIndication:
			c.SignallingPriorityIndication = i
		case ie.IPAddress:
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
		case ie.PortNumber:
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
		case ie.AdditionalProtocolConfigurationOptions:
			c.APCO = i
		case ie.TWANIdentifier:
			switch i.Instance() {
			case 0:
				c.TWANIdentifier = i
			case 1:
				c.WLANLocationInformation = i
			}
		case ie.CNOperatorSelectionEntity:
			c.CNOperatorSelectionEntity = i
		case ie.PresenceReportingAreaInformation:
			c.PresenceReportingAreaInformation = append(c.PresenceReportingAreaInformation, i)
		case ie.OverloadControlInformation:
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
		case ie.MillisecondTimeStamp:
			c.OriginationTimeStamp = i
		case ie.IntegerNumber:
			c.MaximumWaitTime = i
		case ie.TWANIdentifierTimestamp:
			c.WLANLocationTimeStamp = i
		case ie.FContainer:
			c.NBIFOMContainer = i
		case ie.RemoteUEContext:
			c.RemoteUEContextConnected = append(c.RemoteUEContextConnected, i)
		case ie.NodeIdentifier:
			c.TGPPAAAServerIdentifier = i
		case ie.ExtendedProtocolConfigurationOptions:
			c.EPCO = i
		case ie.ServingPLMNRateControl:
			c.ServingPLMNRateControl = i
		case ie.Counter:
			c.MOExceptionDataCounter = i
		case ie.MappedUEUsageType:
			c.MappedUEUsageType = i
		case ie.FullyQualifiedDomainName:
			c.SGWUNodeName = i
		case ie.SecondaryRATUsageDataReport:
			c.SecondaryRATUsageDataReport = append(c.SecondaryRATUsageDataReport, i)
		case ie.UPFunctionSelectionIndicationFlags:
			c.UPFunctionSelectionIndicationFlags = i
		case ie.APNRateControlStatus:
			c.APNRateControlStatus = i
		case ie.PrivateExtension:
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
	for _, ie := range c.BearerContextsToBeCreated {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	for _, ie := range c.BearerContextsToBeRemoved {
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
	for _, ie := range c.PresenceReportingAreaInformation {
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
	for _, ie := range c.RemoteUEContextConnected {
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
	for _, ie := range c.SecondaryRATUsageDataReport {
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

	decodedIEs, err := ie.ParseMultiIEs(c.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.IMSI:
			c.IMSI = i
		case ie.MSISDN:
			c.MSISDN = i
		case ie.MobileEquipmentIdentity:
			c.MEI = i
		case ie.UserLocationInformation:
			switch i.Instance() {
			case 0:
				c.ULI = i
			case 1:
				c.ULIForSGW = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.ServingNetwork:
			c.ServingNetwork = i
		case ie.RATType:
			c.RATType = i
		case ie.Indication:
			c.IndicationFlags = i
		case ie.FullyQualifiedTEID:
			switch i.Instance() {
			case 0:
				c.SenderFTEIDC = i
			case 1:
				c.PGWS5S8FTEIDC = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.AccessPointName:
			c.APN = i
		case ie.SelectionMode:
			c.SelectionMode = i
		case ie.PDNType:
			c.PDNType = i
		case ie.PDNAddressAllocation:
			c.PAA = i
		case ie.APNRestriction:
			c.APNRestriction = i
		case ie.AggregateMaximumBitRate:
			c.AMBR = i
		case ie.EPSBearerID:
			c.LinkedEBI = i
		case ie.TrustedWLANModeIndication:
			c.TWMI = i
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.BearerContext:
			switch i.Instance() {
			case 0:
				c.BearerContextsToBeCreated = append(c.BearerContextsToBeCreated, i)
			case 1:
				c.BearerContextsToBeRemoved = append(c.BearerContextsToBeCreated, i)
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.FullyQualifiedCSID:
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
		case ie.TraceInformation:
			c.TraceInformation = i
		case ie.Recovery:
			c.Recovery = i
		case ie.UETimeZone:
			c.UETimeZone = i
		case ie.UserCSGInformation:
			c.UCI = i
		case ie.ChargingCharacteristics:
			c.ChargingCharacteristics = i
		case ie.LocalDistinguishedName:
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
		case ie.SignallingPriorityIndication:
			c.SignallingPriorityIndication = i
		case ie.IPAddress:
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
		case ie.PortNumber:
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
		case ie.AdditionalProtocolConfigurationOptions:
			c.APCO = i
		case ie.TWANIdentifier:
			switch i.Instance() {
			case 0:
				c.TWANIdentifier = i
			case 1:
				c.WLANLocationInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.CNOperatorSelectionEntity:
			c.CNOperatorSelectionEntity = i
		case ie.PresenceReportingAreaInformation:
			c.PresenceReportingAreaInformation = append(c.PresenceReportingAreaInformation, i)
		case ie.OverloadControlInformation:
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
		case ie.MillisecondTimeStamp:
			c.OriginationTimeStamp = i
		case ie.IntegerNumber:
			c.MaximumWaitTime = i
		case ie.TWANIdentifierTimestamp:
			c.WLANLocationTimeStamp = i
		case ie.FContainer:
			c.NBIFOMContainer = i
		case ie.RemoteUEContext:
			c.RemoteUEContextConnected = append(c.PresenceReportingAreaInformation, i)
		case ie.NodeIdentifier:
			c.TGPPAAAServerIdentifier = i
		case ie.ExtendedProtocolConfigurationOptions:
			c.EPCO = i
		case ie.ServingPLMNRateControl:
			c.ServingPLMNRateControl = i
		case ie.Counter:
			c.MOExceptionDataCounter = i
		case ie.MappedUEUsageType:
			c.MappedUEUsageType = i
		case ie.FullyQualifiedDomainName:
			c.SGWUNodeName = i
		case ie.SecondaryRATUsageDataReport:
			c.SecondaryRATUsageDataReport = append(c.PresenceReportingAreaInformation, i)
		case ie.UPFunctionSelectionIndicationFlags:
			c.UPFunctionSelectionIndicationFlags = i
		case ie.APNRateControlStatus:
			c.APNRateControlStatus = i
		case ie.PrivateExtension:
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
	for _, ie := range c.BearerContextsToBeCreated {
		l += ie.MarshalLen()
	}
	for _, ie := range c.BearerContextsToBeRemoved {
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
	for _, ie := range c.PresenceReportingAreaInformation {
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
	for _, ie := range c.RemoteUEContextConnected {
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
	for _, ie := range c.SecondaryRATUsageDataReport {
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

//////////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *CreateSessionRequest) MarshalTo1(Payload []byte, marshalLen int, debug bool) error {

	c.Header.Payload = make([]byte, marshalLen-c.Header.MarshalLen())

	itemVal := reflect.ValueOf(*c)
	var IElen, offset int64
	for i := 1; i < itemVal.NumField(); i++ { //loop over fields in msg
		fieldVal := itemVal.Field(i)        //get a field
		if fieldVal.Kind() == reflect.Ptr { //check that it's a pointer
			fieldVal1 := fieldVal.Elem() // This is only helpful if the field is not nil
			if debug {
				fmt.Println(i, " ==> fieldVal ", fieldVal.Kind(), "   ", fieldVal1.Kind())
			}
			if fieldVal1.Kind() != reflect.Invalid { //check that Kind is valid
				//fmt.Println(i, " ==> fieldVal1.Type().Name() ;", fieldVal1.Type().Name(), ";")

				MarshalLenMethod := fieldVal.MethodByName("MarshalLen")      //get method (i *IE) MarshalLen()
				MarshalLenResult := MarshalLenMethod.Call([]reflect.Value{}) //call method (i *IE) MarshalLen()
				IElen = MarshalLenResult[0].Int()

				MarshalToMethod := fieldVal.MethodByName("MarshalTo")                                                //get method (i *IE) MarshalTo()
				MarshalToResult := MarshalToMethod.Call([]reflect.Value{reflect.ValueOf(c.Header.Payload[offset:])}) //call method (i *IE) MarshalTo()
				if debug {
					fmt.Println(i, " MarshalToResult[0] ==> ", MarshalToResult[0], "IElen = ", IElen)
					fmt.Println(i, " ==> fieldVal1 ", fieldVal1.Kind(),
						"  reflect.TypeOf(fieldVal1) =  ", fieldVal1.Type().Name())
				}
				offset += IElen //update offset

			}
			//fmt.Println(i, " ==> fieldVal ", fieldVal.Kind(), "   ", fieldVal1.Kind())
		} else {
			if debug {
				fmt.Println(i, " ==> fieldVal ", fieldVal.Kind())
			}
		}
	}
	if debug {
		fmt.Println("MarshalTo(msg *CreateSessionRequest) ==> msg.Header.Payload =  ", hex.Dump(c.Header.Payload))
	}
	c.Header.SetLength()
	return c.Header.MarshalTo(Payload)

}

func (c *CreateSessionRequest) MarshalLen1(debug bool) int64 {
	if debug {
		fmt.Printf(" CreateSessionRequest ==> len(msg.Header.Payload) =  %d \n", len(c.Header.Payload))
	}
	itemVal := reflect.ValueOf(*c)
	var Totlen, IElen int64
	for i := 0; i < itemVal.NumField(); i++ { //loop over fields in msg
		fieldVal := itemVal.Field(i)        //get a field
		if fieldVal.Kind() == reflect.Ptr { //check that it's a pointer
			fieldVal1 := fieldVal.Elem() // This is only helpful if the field is not nil
			if debug {
				fmt.Println(i, " ==> fieldVal ", fieldVal.Kind(), "   ", fieldVal1.Kind())
			}
			if fieldVal1.Kind() != reflect.Invalid { //check that Kind is valid
				//fmt.Println(i, " ==> fieldVal1.Type().Name() ;", fieldVal1.Type().Name(), ";")

				MarshalLenMethod := fieldVal.MethodByName("MarshalLen")      //get method (i *IE) MarshalLen()
				MarshalLenResult := MarshalLenMethod.Call([]reflect.Value{}) //call method (i *IE) MarshalLen()
				IElen = MarshalLenResult[0].Int()                            // get result from (i *IE) MarshalLen()
				Totlen += IElen                                              //add MarshalLen to total length
				if debug {
					fmt.Println(i, "IElen = ", IElen)
				}

			}
			//fmt.Println(i, " ==> fieldVal ", fieldVal.Kind(), "   ", fieldVal1.Kind())
		} else {
			if debug {
				fmt.Println(i, " ==> fieldVal ", fieldVal.Kind())
			}
		}
	}
	return Totlen
}
