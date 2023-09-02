// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// CreateSessionResponse is a CreateSessionResponse Header and its IEs above.
type CreateSessionResponse struct {
	*Header
	Cause                         *ie.IE
	ChangeReportingAction         *ie.IE
	CSGInformationReportingAction *ie.IE
	HeNBInformationReporting      *ie.IE
	SenderFTEIDC                  *ie.IE
	PGWS5S8FTEIDC                 *ie.IE
	PAA                           *ie.IE
	APNRestriction                *ie.IE
	AMBR                          *ie.IE
	EBI                           *ie.IE
	PCO                           *ie.IE
	BearerContextsCreated         []*ie.IE
	BearerContextMarkedForRemoval *ie.IE
	Recovery                      *ie.IE
	ChargingGatewayName           *ie.IE // = PGWNodeName
	ChargingGatewayAddress        *ie.IE
	PGWFQCSID                     *ie.IE
	SGWFQCSID                     *ie.IE
	SGWLDN                        *ie.IE
	PGWLDN                        *ie.IE
	PGWBackOffTime                *ie.IE
	APCO                          *ie.IE
	TrustedTWANIPv4Parameters     *ie.IE
	IndicationFlags               *ie.IE
	PresenceReportingAreaAction   []*ie.IE
	PGWNodeLoadControlInformation *ie.IE
	PGWAPNLoadControlInformation  *ie.IE
	SGWNodeLoadControlInformation *ie.IE
	PGWOverloadControlInformation *ie.IE
	SGWOverloadControlInformation *ie.IE
	NBIFOMContainer               *ie.IE
	PDNConnectionChargingID       *ie.IE
	EPCO                          *ie.IE
	PrivateExtension              *ie.IE
	AdditionalIEs                 []*ie.IE
}

// NewCreateSessionResponse creates a new CreateSessionResponse.
func NewCreateSessionResponse(teid, seq uint32, ies ...*ie.IE) *CreateSessionResponse {
	c := &CreateSessionResponse{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeCreateSessionResponse, teid, seq, nil,
		),
	}
	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			c.Cause = i
		case ie.ChangeReportingAction:
			c.ChangeReportingAction = i
		case ie.CSGInformationReportingAction:
			c.CSGInformationReportingAction = i
		case ie.HeNBInformationReporting:
			c.HeNBInformationReporting = i
		case ie.FullyQualifiedTEID:
			switch i.Instance() {
			case 0:
				c.SenderFTEIDC = i
			case 1:
				c.PGWS5S8FTEIDC = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.PDNAddressAllocation:
			c.PAA = i
		case ie.APNRestriction:
			c.APNRestriction = i
		case ie.AggregateMaximumBitRate:
			c.AMBR = i
		case ie.EPSBearerID:
			c.EBI = i
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.BearerContext:
			switch i.Instance() {
			case 0:
				c.BearerContextsCreated = append(c.BearerContextsCreated, i)
			case 1:
				c.BearerContextMarkedForRemoval = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.Recovery:
			c.Recovery = i
		case ie.FullyQualifiedDomainName:
			c.ChargingGatewayName = i
		case ie.IPAddress:
			c.ChargingGatewayAddress = i
		case ie.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				c.PGWFQCSID = i
			case 1:
				c.SGWFQCSID = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.LocalDistinguishedName:
			switch i.Instance() {
			case 0:
				c.PGWLDN = i
			case 1:
				c.SGWLDN = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.EPCTimer:
			c.PGWBackOffTime = i
		case ie.AdditionalProtocolConfigurationOptions:
			c.APCO = i
		case ie.IPv4ConfigurationParameters:
			c.TrustedTWANIPv4Parameters = i
		case ie.Indication:
			c.IndicationFlags = i
		case ie.PresenceReportingAreaAction:
			c.PresenceReportingAreaAction = append(c.PresenceReportingAreaAction, i)
		case ie.LoadControlInformation:
			switch i.Instance() {
			case 0:
				c.PGWNodeLoadControlInformation = i
			case 1:
				c.PGWAPNLoadControlInformation = i
			case 2:
				c.SGWNodeLoadControlInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				c.PGWOverloadControlInformation = i
			case 1:
				c.SGWOverloadControlInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.FContainer:
			c.NBIFOMContainer = i
		case ie.ChargingID:
			c.PDNConnectionChargingID = i
		case ie.ExtendedProtocolConfigurationOptions:
			c.EPCO = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Marshal serializes CreateSessionResponse into bytes.
func (c *CreateSessionResponse) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes CreateSessionResponse into bytes.
func (c *CreateSessionResponse) MarshalTo(b []byte) error {
	if c.Header.Payload != nil {
		c.Header.Payload = nil
	}
	c.Header.Payload = make([]byte, c.MarshalLen()-c.Header.MarshalLen())

	offset := 0
	if ie := c.Cause; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ChangeReportingAction; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.CSGInformationReportingAction; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.HeNBInformationReporting; ie != nil {
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
	if ie := c.EBI; ie != nil {
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
	for _, ie := range c.BearerContextsCreated {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.BearerContextMarkedForRemoval; ie != nil {
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
	if ie := c.ChargingGatewayName; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ChargingGatewayAddress; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PGWFQCSID; ie != nil {
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
	if ie := c.PGWLDN; ie != nil {
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
	if ie := c.PGWBackOffTime; ie != nil {
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
	if ie := c.TrustedTWANIPv4Parameters; ie != nil {
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
	for _, ie := range c.PresenceReportingAreaAction {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PGWNodeLoadControlInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PGWAPNLoadControlInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGWNodeLoadControlInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PGWOverloadControlInformation; ie != nil {
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
	if ie := c.NBIFOMContainer; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PDNConnectionChargingID; ie != nil {
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
	if ie := c.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
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

// ParseCreateSessionResponse decodes given bytes as CreateSessionResponse.
func ParseCreateSessionResponse(b []byte) (*CreateSessionResponse, error) {
	c := &CreateSessionResponse{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes given bytes as CreateSessionResponse.
func (c *CreateSessionResponse) UnmarshalBinary(b []byte) error {
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
		case ie.Cause:
			c.Cause = i
		case ie.ChangeReportingAction:
			c.ChangeReportingAction = i
		case ie.CSGInformationReportingAction:
			c.CSGInformationReportingAction = i
		case ie.HeNBInformationReporting:
			c.HeNBInformationReporting = i
		case ie.FullyQualifiedTEID:
			switch i.Instance() {
			case 0:
				c.SenderFTEIDC = i
			case 1:
				c.PGWS5S8FTEIDC = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.PDNAddressAllocation:
			c.PAA = i
		case ie.APNRestriction:
			c.APNRestriction = i
		case ie.AggregateMaximumBitRate:
			c.AMBR = i
		case ie.EPSBearerID:
			c.EBI = i
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.BearerContext:
			switch i.Instance() {
			case 0:
				c.BearerContextsCreated = append(c.BearerContextsCreated, i)
			case 1:
				c.BearerContextMarkedForRemoval = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.Recovery:
			c.Recovery = i
		case ie.FullyQualifiedDomainName:
			c.ChargingGatewayName = i
		case ie.IPAddress:
			c.ChargingGatewayAddress = i
		case ie.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				c.PGWFQCSID = i
			case 1:
				c.SGWFQCSID = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.LocalDistinguishedName:
			switch i.Instance() {
			case 0:
				c.PGWLDN = i
			case 1:
				c.SGWLDN = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.EPCTimer:
			c.PGWBackOffTime = i
		case ie.AdditionalProtocolConfigurationOptions:
			c.APCO = i
		case ie.IPv4ConfigurationParameters:
			c.TrustedTWANIPv4Parameters = i
		case ie.Indication:
			c.IndicationFlags = i
		case ie.PresenceReportingAreaAction:
			c.PresenceReportingAreaAction = append(c.PresenceReportingAreaAction, i)
		case ie.LoadControlInformation:
			switch i.Instance() {
			case 0:
				c.PGWNodeLoadControlInformation = i
			case 1:
				c.PGWAPNLoadControlInformation = i
			case 2:
				c.SGWNodeLoadControlInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				c.PGWOverloadControlInformation = i
			case 1:
				c.SGWOverloadControlInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.FContainer:
			c.NBIFOMContainer = i
		case ie.ChargingID:
			c.PDNConnectionChargingID = i
		case ie.ExtendedProtocolConfigurationOptions:
			c.EPCO = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (c *CreateSessionResponse) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)

	if ie := c.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ChangeReportingAction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.CSGInformationReportingAction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.HeNBInformationReporting; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SenderFTEIDC; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PGWS5S8FTEIDC; ie != nil {
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
	if ie := c.EBI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PCO; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range c.BearerContextsCreated {
		l += ie.MarshalLen()
	}
	if ie := c.BearerContextMarkedForRemoval; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ChargingGatewayName; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ChargingGatewayAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PGWFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PGWLDN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWLDN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PGWBackOffTime; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.APCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TrustedTWANIPv4Parameters; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range c.PresenceReportingAreaAction {
		l += ie.MarshalLen()
	}
	if ie := c.PGWNodeLoadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PGWAPNLoadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWNodeLoadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PGWOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.NBIFOMContainer; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PDNConnectionChargingID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.EPCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}

	for _, ie := range c.AdditionalIEs {
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length field.
func (c *CreateSessionResponse) SetLength() {
	c.Header.Length = uint16(c.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *CreateSessionResponse) MessageTypeName() string {
	return "Create Session Response"
}

// TEID returns the TEID in uint32.
func (c *CreateSessionResponse) TEID() uint32 {
	return c.Header.teid()
}
