// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv1/ie"
)

// CreatePDPContextResponse is a CreatePDPContextResponse Header and its IEs above.
type CreatePDPContextResponse struct {
	*Header
	Cause                         *ie.IE
	ReorderingRequired            *ie.IE
	Recovery                      *ie.IE
	TEIDDataI                     *ie.IE
	TEIDCPlane                    *ie.IE
	NSAPI                         *ie.IE
	ChargingID                    *ie.IE
	EndUserAddress                *ie.IE
	PCO                           *ie.IE
	GGSNAddressForCPlane          *ie.IE
	GGSNAddressForUserTraffic     *ie.IE
	AltGGSNAddressForCPlane       *ie.IE
	AltGGSNAddressForUserTraffic  *ie.IE
	QoSProfile                    *ie.IE
	ChargingGatewayAddress        *ie.IE
	AltChargingGatewayAddress     *ie.IE
	CommonFlags                   *ie.IE
	APNRestriction                *ie.IE
	MSInfoChangeReportingAction   *ie.IE
	BearerControlMode             *ie.IE
	EvolvedARPI                   *ie.IE
	ExtendedCommonFlag            *ie.IE
	CSGInformationReportingAction *ie.IE
	APNAMBR                       *ie.IE
	GGSNBackOffTime               *ie.IE
	ExtendedCommonFlagsII         *ie.IE
	PrivateExtension              *ie.IE
	AdditionalIEs                 []*ie.IE
}

// NewCreatePDPContextResponse creates a new GTPv1 CreatePDPContextResponse.
func NewCreatePDPContextResponse(teid uint32, seq uint16, ies ...*ie.IE) *CreatePDPContextResponse {
	c := &CreatePDPContextResponse{
		Header: NewHeader(0x32, MsgTypeCreatePDPContextResponse, teid, seq, nil),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			c.Cause = i
		case ie.ReorderingRequired:
			c.ReorderingRequired = i
		case ie.Recovery:
			c.Recovery = i
		case ie.TEIDDataI:
			c.TEIDDataI = i
		case ie.TEIDCPlane:
			c.TEIDCPlane = i
		case ie.NSAPI:
			c.NSAPI = i
		case ie.ChargingID:
			c.ChargingID = i
		case ie.EndUserAddress:
			c.EndUserAddress = i
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.GSNAddress:
			if c.GGSNAddressForCPlane == nil {
				c.GGSNAddressForCPlane = i
			} else if c.GGSNAddressForUserTraffic == nil {
				c.GGSNAddressForUserTraffic = i
			} else if c.AltGGSNAddressForCPlane == nil {
				c.AltGGSNAddressForCPlane = i
			} else if c.AltGGSNAddressForUserTraffic == nil {
				c.AltGGSNAddressForUserTraffic = i
			}
		case ie.QoSProfile:
			c.QoSProfile = i
		case ie.ChargingGatewayAddress:
			if c.ChargingGatewayAddress == nil {
				c.ChargingGatewayAddress = i
			} else if c.AltChargingGatewayAddress == nil {
				c.AltChargingGatewayAddress = i
			}
		case ie.CommonFlags:
			c.CommonFlags = i
		case ie.APNRestriction:
			c.APNRestriction = i
		case ie.MSInfoChangeReportingAction:
			c.MSInfoChangeReportingAction = i
		case ie.BearerControlMode:
			c.BearerControlMode = i
		case ie.EvolvedAllocationRetentionPriorityI:
			c.EvolvedARPI = i
		case ie.ExtendedCommonFlags:
			c.ExtendedCommonFlag = i
		case ie.CSGInformationReportingAction:
			c.CSGInformationReportingAction = i
		case ie.AggregateMaximumBitRate:
			c.APNAMBR = i
		case ie.GGSNBackOffTime:
			c.GGSNBackOffTime = i
		case ie.ExtendedCommonFlagsII:
			c.ExtendedCommonFlagsII = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Marshal returns the byte sequence generated from a CreatePDPContextResponse.
func (c *CreatePDPContextResponse) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (c *CreatePDPContextResponse) MarshalTo(b []byte) error {
	if len(b) < c.MarshalLen() {
		return ErrTooShortToMarshal
	}
	c.Header.Payload = make([]byte, c.MarshalLen()-c.Header.MarshalLen())

	offset := 0
	if ie := c.Cause; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ReorderingRequired; ie != nil {
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
	if ie := c.TEIDDataI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.TEIDCPlane; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.NSAPI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ChargingID; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.EndUserAddress; ie != nil {
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
	if ie := c.GGSNAddressForCPlane; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.GGSNAddressForUserTraffic; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.AltGGSNAddressForCPlane; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.AltGGSNAddressForUserTraffic; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.QoSProfile; ie != nil {
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
	if ie := c.AltChargingGatewayAddress; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.CommonFlags; ie != nil {
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
	if ie := c.MSInfoChangeReportingAction; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.BearerControlMode; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.EvolvedARPI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ExtendedCommonFlag; ie != nil {
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
	if ie := c.APNAMBR; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.GGSNBackOffTime; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ExtendedCommonFlagsII; ie != nil {
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

// ParseCreatePDPContextResponse decodes a given byte sequence as a CreatePDPContextResponse.
func ParseCreatePDPContextResponse(b []byte) (*CreatePDPContextResponse, error) {
	c := &CreatePDPContextResponse{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes a given byte sequence as a CreatePDPContextResponse.
func (c *CreatePDPContextResponse) UnmarshalBinary(b []byte) error {
	var err error
	c.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(c.Header.Payload) < 2 {
		return nil
	}

	ies, err := ie.ParseMultiIEs(c.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			c.Cause = i
		case ie.ReorderingRequired:
			c.ReorderingRequired = i
		case ie.Recovery:
			c.Recovery = i
		case ie.TEIDDataI:
			c.TEIDDataI = i
		case ie.TEIDCPlane:
			c.TEIDCPlane = i
		case ie.NSAPI:
			c.NSAPI = i
		case ie.ChargingID:
			c.ChargingID = i
		case ie.EndUserAddress:
			c.EndUserAddress = i
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.GSNAddress:
			if c.GGSNAddressForCPlane == nil {
				c.GGSNAddressForCPlane = i
			} else if c.GGSNAddressForUserTraffic == nil {
				c.GGSNAddressForUserTraffic = i
			} else if c.AltGGSNAddressForCPlane == nil {
				c.AltGGSNAddressForCPlane = i
			} else if c.AltGGSNAddressForUserTraffic == nil {
				c.AltGGSNAddressForUserTraffic = i
			}
		case ie.QoSProfile:
			c.QoSProfile = i
		case ie.ChargingGatewayAddress:
			if c.ChargingGatewayAddress == nil {
				c.ChargingGatewayAddress = i
			} else if c.AltChargingGatewayAddress == nil {
				c.AltChargingGatewayAddress = i
			}
		case ie.CommonFlags:
			c.CommonFlags = i
		case ie.APNRestriction:
			c.APNRestriction = i
		case ie.MSInfoChangeReportingAction:
			c.MSInfoChangeReportingAction = i
		case ie.BearerControlMode:
			c.BearerControlMode = i
		case ie.EvolvedAllocationRetentionPriorityI:
			c.EvolvedARPI = i
		case ie.ExtendedCommonFlags:
			c.ExtendedCommonFlag = i
		case ie.CSGInformationReportingAction:
			c.CSGInformationReportingAction = i
		case ie.AggregateMaximumBitRate:
			c.APNAMBR = i
		case ie.GGSNBackOffTime:
			c.GGSNBackOffTime = i
		case ie.ExtendedCommonFlagsII:
			c.ExtendedCommonFlagsII = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length of Data.
func (c *CreatePDPContextResponse) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)

	if ie := c.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ReorderingRequired; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TEIDDataI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TEIDCPlane; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.NSAPI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ChargingID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.EndUserAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.GGSNAddressForCPlane; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.GGSNAddressForUserTraffic; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.AltGGSNAddressForCPlane; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.AltGGSNAddressForUserTraffic; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.QoSProfile; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ChargingGatewayAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.AltChargingGatewayAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.CommonFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.APNRestriction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MSInfoChangeReportingAction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.BearerControlMode; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.EvolvedARPI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ExtendedCommonFlag; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.CSGInformationReportingAction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.APNAMBR; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.GGSNBackOffTime; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ExtendedCommonFlagsII; ie != nil {
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
func (c *CreatePDPContextResponse) SetLength() {
	c.Length = uint16(c.MarshalLen() - 8)
}

// MessageTypeName returns the name of protocol.
func (c *CreatePDPContextResponse) MessageTypeName() string {
	return "Create PDP Context Response"
}

// TEID returns the TEID in human-readable string.
func (c *CreatePDPContextResponse) TEID() uint32 {
	return c.Header.TEID
}
