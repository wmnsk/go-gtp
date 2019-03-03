// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v2/ies"
)

// CreateSessionResponse is a CreateSessionResponse Header and its IEs above.
type CreateSessionResponse struct {
	*Header
	Cause                         *ies.IE
	ChangeReportingAction         *ies.IE
	HeNBInformationReporting      *ies.IE
	SenderFTEIDC                  *ies.IE
	PGWS5S8FTEIDC                 *ies.IE
	PAA                           *ies.IE
	APNRestriction                *ies.IE
	AMBR                          *ies.IE
	EBI                           *ies.IE
	PCO                           *ies.IE
	BearerContextsCreated         *ies.IE
	BearerContextMarkedForRemoval *ies.IE
	Recovery                      *ies.IE
	ChargingGatewayName           *ies.IE
	ChargingGatewayAddress        *ies.IE
	PGWFQCSID                     *ies.IE
	SGWFQCSID                     *ies.IE
	PGWLDN                        *ies.IE
	SGWLDN                        *ies.IE
	PGWBackOffTime                *ies.IE
	APCO                          *ies.IE
	TrustedTWANIPv4Parameters     *ies.IE
	IndicationFlags               *ies.IE
	PresenceReportingAreaAction   *ies.IE
	PGWNodeLoadControlInformation *ies.IE
	PGWAPNLoadControlInformation  *ies.IE
	SGWNodeLoadControlInformation *ies.IE
	PGWOverloadControlInformation *ies.IE
	SGWOverloadControlInformation *ies.IE
	NBIFOMContainer               *ies.IE
	PDNConnectionChargingID       *ies.IE
	EPCO                          *ies.IE
	PrivateExtension              *ies.IE
	AdditionalIEs                 []*ies.IE
}

// NewCreateSessionResponse creates a new CreateSessionResponse.
func NewCreateSessionResponse(teid, seq uint32, ie ...*ies.IE) *CreateSessionResponse {
	c := &CreateSessionResponse{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeCreateSessionResponse, teid, seq, nil,
		),
	}
	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			c.Cause = i
		case ies.ChangeReportingAction:
			c.ChangeReportingAction = i
		case ies.HeNBInformationReporting:
			c.HeNBInformationReporting = i
		case ies.FullyQualifiedTEID:
			switch i.Instance() {
			case 0:
				c.SenderFTEIDC = i
			case 1:
				c.PGWS5S8FTEIDC = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.PDNAddressAllocation:
			c.PAA = i
		case ies.APNRestriction:
			c.APNRestriction = i
		case ies.AggregateMaximumBitRate:
			c.AMBR = i
		case ies.EPSBearerID:
			c.EBI = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				c.BearerContextsCreated = i
			case 1:
				c.BearerContextMarkedForRemoval = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.Recovery:
			c.Recovery = i
		case ies.FullyQualifiedDomainName:
			c.ChargingGatewayName = i
		case ies.IPAddress:
			c.ChargingGatewayAddress = i
		case ies.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				c.PGWFQCSID = i
			case 1:
				c.SGWFQCSID = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.LocalDistinguishedName:
			switch i.Instance() {
			case 0:
				c.PGWLDN = i
			case 1:
				c.SGWLDN = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.EPCTimer:
			c.PGWBackOffTime = i
		case ies.AdditionalProtocolConfigurationOptions:
			c.APCO = i
		case ies.IPv4ConfigurationParameters:
			c.TrustedTWANIPv4Parameters = i
		case ies.Indication:
			c.IndicationFlags = i
		case ies.PresenceReportingAreaAction:
			c.PresenceReportingAreaAction = i
		case ies.LoadControlInformation:
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
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				c.PGWOverloadControlInformation = i
			case 1:
				c.SGWOverloadControlInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.FContainer:
			c.NBIFOMContainer = i
		case ies.ChargingID:
			c.PDNConnectionChargingID = i
		case ies.ExtendedProtocolConfigurationOptions:
			c.EPCO = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Serialize serializes CreateSessionResponse into bytes.
func (c *CreateSessionResponse) Serialize() ([]byte, error) {
	b := make([]byte, c.Len())
	if err := c.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes CreateSessionResponse into bytes.
func (c *CreateSessionResponse) SerializeTo(b []byte) error {
	if c.Header.Payload != nil {
		c.Header.Payload = nil
	}
	c.Header.Payload = make([]byte, c.Len()-c.Header.Len())

	offset := 0
	if ie := c.Cause; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.ChangeReportingAction; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.HeNBInformationReporting; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.SenderFTEIDC; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PGWS5S8FTEIDC; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PAA; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.APNRestriction; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.AMBR; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.EBI; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PCO; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.BearerContextsCreated; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.BearerContextMarkedForRemoval; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.Recovery; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.ChargingGatewayName; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.ChargingGatewayAddress; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PGWFQCSID; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.SGWFQCSID; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PGWLDN; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.SGWLDN; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PGWBackOffTime; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.APCO; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.TrustedTWANIPv4Parameters; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.IndicationFlags; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PresenceReportingAreaAction; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PGWNodeLoadControlInformation; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PGWAPNLoadControlInformation; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.SGWNodeLoadControlInformation; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PGWOverloadControlInformation; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.SGWOverloadControlInformation; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.NBIFOMContainer; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PDNConnectionChargingID; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.EPCO; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PrivateExtension; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	for _, ie := range c.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.SerializeTo(c.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	c.Header.SetLength()
	return c.Header.SerializeTo(b)
}

// DecodeCreateSessionResponse decodes given bytes as CreateSessionResponse.
func DecodeCreateSessionResponse(b []byte) (*CreateSessionResponse, error) {
	c := &CreateSessionResponse{}
	if err := c.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return c, nil
}

// DecodeFromBytes decodes given bytes as CreateSessionResponse.
func (c *CreateSessionResponse) DecodeFromBytes(b []byte) error {
	var err error
	c.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}
	if len(c.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ies.DecodeMultiIEs(c.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			c.Cause = i
		case ies.ChangeReportingAction:
			c.ChangeReportingAction = i
		case ies.HeNBInformationReporting:
			c.HeNBInformationReporting = i
		case ies.FullyQualifiedTEID:
			switch i.Instance() {
			case 0:
				c.SenderFTEIDC = i
			case 1:
				c.PGWS5S8FTEIDC = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.PDNAddressAllocation:
			c.PAA = i
		case ies.APNRestriction:
			c.APNRestriction = i
		case ies.AggregateMaximumBitRate:
			c.AMBR = i
		case ies.EPSBearerID:
			c.EBI = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				c.BearerContextsCreated = i
			case 1:
				c.BearerContextMarkedForRemoval = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.Recovery:
			c.Recovery = i
		case ies.FullyQualifiedDomainName:
			c.ChargingGatewayName = i
		case ies.IPAddress:
			c.ChargingGatewayAddress = i
		case ies.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				c.PGWFQCSID = i
			case 1:
				c.SGWFQCSID = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.LocalDistinguishedName:
			switch i.Instance() {
			case 0:
				c.PGWLDN = i
			case 1:
				c.SGWLDN = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.EPCTimer:
			c.PGWBackOffTime = i
		case ies.AdditionalProtocolConfigurationOptions:
			c.APCO = i
		case ies.IPv4ConfigurationParameters:
			c.TrustedTWANIPv4Parameters = i
		case ies.Indication:
			c.IndicationFlags = i
		case ies.PresenceReportingAreaAction:
			c.PresenceReportingAreaAction = i
		case ies.LoadControlInformation:
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
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				c.PGWOverloadControlInformation = i
			case 1:
				c.SGWOverloadControlInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.FContainer:
			c.NBIFOMContainer = i
		case ies.ChargingID:
			c.PDNConnectionChargingID = i
		case ies.ExtendedProtocolConfigurationOptions:
			c.EPCO = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// Len returns the actual length in int.
func (c *CreateSessionResponse) Len() int {
	l := c.Header.Len() - len(c.Header.Payload)

	if ie := c.Cause; ie != nil {
		l += ie.Len()
	}
	if ie := c.ChangeReportingAction; ie != nil {
		l += ie.Len()
	}
	if ie := c.HeNBInformationReporting; ie != nil {
		l += ie.Len()
	}
	if ie := c.SenderFTEIDC; ie != nil {
		l += ie.Len()
	}
	if ie := c.PGWS5S8FTEIDC; ie != nil {
		l += ie.Len()
	}
	if ie := c.PAA; ie != nil {
		l += ie.Len()
	}
	if ie := c.APNRestriction; ie != nil {
		l += ie.Len()
	}
	if ie := c.AMBR; ie != nil {
		l += ie.Len()
	}
	if ie := c.EBI; ie != nil {
		l += ie.Len()
	}
	if ie := c.PCO; ie != nil {
		l += ie.Len()
	}
	if ie := c.BearerContextsCreated; ie != nil {
		l += ie.Len()
	}
	if ie := c.BearerContextMarkedForRemoval; ie != nil {
		l += ie.Len()
	}
	if ie := c.Recovery; ie != nil {
		l += ie.Len()
	}
	if ie := c.ChargingGatewayName; ie != nil {
		l += ie.Len()
	}
	if ie := c.ChargingGatewayAddress; ie != nil {
		l += ie.Len()
	}
	if ie := c.PGWFQCSID; ie != nil {
		l += ie.Len()
	}
	if ie := c.SGWFQCSID; ie != nil {
		l += ie.Len()
	}
	if ie := c.PGWLDN; ie != nil {
		l += ie.Len()
	}
	if ie := c.SGWLDN; ie != nil {
		l += ie.Len()
	}
	if ie := c.PGWBackOffTime; ie != nil {
		l += ie.Len()
	}
	if ie := c.APCO; ie != nil {
		l += ie.Len()
	}
	if ie := c.TrustedTWANIPv4Parameters; ie != nil {
		l += ie.Len()
	}
	if ie := c.IndicationFlags; ie != nil {
		l += ie.Len()
	}
	if ie := c.PresenceReportingAreaAction; ie != nil {
		l += ie.Len()
	}
	if ie := c.PGWNodeLoadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := c.PGWAPNLoadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := c.SGWNodeLoadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := c.PGWOverloadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := c.SGWOverloadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := c.NBIFOMContainer; ie != nil {
		l += ie.Len()
	}
	if ie := c.PDNConnectionChargingID; ie != nil {
		l += ie.Len()
	}
	if ie := c.EPCO; ie != nil {
		l += ie.Len()
	}
	if ie := c.PrivateExtension; ie != nil {
		l += ie.Len()
	}

	for _, ie := range c.AdditionalIEs {
		l += ie.Len()
	}
	return l
}

// SetLength sets the length in Length field.
func (c *CreateSessionResponse) SetLength() {
	c.Header.Length = uint16(c.Len() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *CreateSessionResponse) MessageTypeName() string {
	return "Create Session Response"
}

// TEID returns the TEID in uint32.
func (c *CreateSessionResponse) TEID() uint32 {
	return c.Header.teid()
}
