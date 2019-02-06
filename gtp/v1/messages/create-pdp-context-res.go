// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/gtp/v1/ies"
)

// CreatePDPContextResponse is a CreatePDPContextResponse Header and its IEs above.
type CreatePDPContextResponse struct {
	*Header
	Cause                                *ies.IE
	ReorderingRequired                   *ies.IE
	Recovery                             *ies.IE
	TEIDDataI                            *ies.IE
	TEIDCPlane                           *ies.IE
	NSAPI                                *ies.IE
	ChargingID                           *ies.IE
	EndUserAddress                       *ies.IE
	PCO                                  *ies.IE
	GGSNAddressForCPlane                 *ies.IE
	GGSNAddressForUsertraffic            *ies.IE
	AlternativeGGSNAddressForCPlane      *ies.IE
	AlternativeGGSNAddressForUsertraffic *ies.IE
	QoSProfile                           *ies.IE
	ChargingGatewayAddress               *ies.IE
	AlternativeChargingGatewayAddress    *ies.IE
	CommonFlags                          *ies.IE
	APNRestriction                       *ies.IE
	MSInfoChangeReportingAction          *ies.IE
	BearerControlMode                    *ies.IE
	EvolvedAllocationRetentionPriorityI  *ies.IE
	ExtendedCommonFlag                   *ies.IE
	CSGInformationReportingAction        *ies.IE
	APNAMBR                              *ies.IE
	GGSNBackOffTime                      *ies.IE
	ExtendedCommonFlagsII                *ies.IE
	PrivateExtension                     *ies.IE
	AdditionalIEs                        []*ies.IE
}

// NewCreatePDPContextResponse creates a new GTPv1 CreatePDPContextResponse.
func NewCreatePDPContextResponse(teid uint32, seq uint16, ie ...*ies.IE) *CreatePDPContextResponse {
	c := &CreatePDPContextResponse{
		Header: NewHeader(0x32, MsgTypeCreatePDPContextResponse, teid, seq, nil),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			c.Cause = i
		case ies.ReorderingRequired:
			c.ReorderingRequired = i
		case ies.Recovery:
			c.Recovery = i
		case ies.TEIDDataI:
			c.TEIDDataI = i
		case ies.TEIDCPlane:
			c.TEIDCPlane = i
		case ies.NSAPI:
			c.NSAPI = i
		case ies.ChargingID:
			c.ChargingID = i
		case ies.EndUserAddress:
			c.EndUserAddress = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.GSNAddress:
			if c.GGSNAddressForCPlane == nil {
				c.GGSNAddressForCPlane = i
			} else if c.GGSNAddressForUsertraffic == nil {
				c.GGSNAddressForUsertraffic = i
			} else if c.AlternativeGGSNAddressForCPlane == nil {
				c.AlternativeGGSNAddressForCPlane = i
			} else if c.AlternativeGGSNAddressForUsertraffic == nil {
				c.AlternativeGGSNAddressForUsertraffic = i
			}
		case ies.QoSProfile:
			c.QoSProfile = i
		case ies.ChargingGatewayAddress:
			if c.ChargingGatewayAddress == nil {
				c.ChargingGatewayAddress = i
			} else if c.AlternativeChargingGatewayAddress == nil {
				c.AlternativeChargingGatewayAddress = i
			}
		case ies.CommonFlags:
			c.CommonFlags = i
		case ies.APNRestriction:
			c.APNRestriction = i
		case ies.MSInfoChangeReportingAction:
			c.MSInfoChangeReportingAction = i
		case ies.BearerControlMode:
			c.BearerControlMode = i
		case ies.EvolvedAllocationRetentionPriorityI:
			c.EvolvedAllocationRetentionPriorityI = i
		case ies.ExtendedCommonFlags:
			c.ExtendedCommonFlag = i
		case ies.CSGInformationReportingAction:
			c.CSGInformationReportingAction = i
		case ies.AggregateMaximumBitRate:
			c.APNAMBR = i
		case ies.GGSNBackOffTime:
			c.GGSNBackOffTime = i
		case ies.ExtendedCommonFlagsII:
			c.ExtendedCommonFlagsII = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Serialize returns the byte sequence generated from a CreatePDPContextResponse.
func (c *CreatePDPContextResponse) Serialize() ([]byte, error) {
	b := make([]byte, c.Len())
	if err := c.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo puts the byte sequence in the byte array given as b.
func (c *CreatePDPContextResponse) SerializeTo(b []byte) error {
	if len(b) < c.Len() {
		return ErrTooShortToSerialize
	}
	c.Header.Payload = make([]byte, c.Len()-c.Header.Len())

	offset := 0
	if ie := c.Cause; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.ReorderingRequired; ie != nil {
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
	if ie := c.TEIDDataI; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.TEIDCPlane; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.NSAPI; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.ChargingID; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.EndUserAddress; ie != nil {
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
	if ie := c.GGSNAddressForCPlane; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.GGSNAddressForUsertraffic; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.AlternativeGGSNAddressForCPlane; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.AlternativeGGSNAddressForUsertraffic; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.QoSProfile; ie != nil {
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
	if ie := c.AlternativeChargingGatewayAddress; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.CommonFlags; ie != nil {
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
	if ie := c.MSInfoChangeReportingAction; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.BearerControlMode; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.EvolvedAllocationRetentionPriorityI; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.ExtendedCommonFlag; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.CSGInformationReportingAction; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.APNAMBR; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.GGSNBackOffTime; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.ExtendedCommonFlagsII; ie != nil {
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

// DecodeCreatePDPContextResponse decodes a given byte sequence as a CreatePDPContextResponse.
func DecodeCreatePDPContextResponse(b []byte) (*CreatePDPContextResponse, error) {
	c := &CreatePDPContextResponse{}
	if err := c.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return c, nil
}

// DecodeFromBytes decodes a given byte sequence as a CreatePDPContextResponse.
func (c *CreatePDPContextResponse) DecodeFromBytes(b []byte) error {
	var err error
	c.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}
	if len(c.Header.Payload) < 2 {
		return nil
	}

	ie, err := ies.DecodeMultiIEs(c.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			c.Cause = i
		case ies.ReorderingRequired:
			c.ReorderingRequired = i
		case ies.Recovery:
			c.Recovery = i
		case ies.TEIDDataI:
			c.TEIDDataI = i
		case ies.TEIDCPlane:
			c.TEIDCPlane = i
		case ies.NSAPI:
			c.NSAPI = i
		case ies.ChargingID:
			c.ChargingID = i
		case ies.EndUserAddress:
			c.EndUserAddress = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.GSNAddress:
			if c.GGSNAddressForCPlane == nil {
				c.GGSNAddressForCPlane = i
			} else if c.GGSNAddressForUsertraffic == nil {
				c.GGSNAddressForUsertraffic = i
			} else if c.AlternativeGGSNAddressForCPlane == nil {
				c.AlternativeGGSNAddressForCPlane = i
			} else if c.AlternativeGGSNAddressForUsertraffic == nil {
				c.AlternativeGGSNAddressForUsertraffic = i
			}
		case ies.QoSProfile:
			c.QoSProfile = i
		case ies.ChargingGatewayAddress:
			if c.ChargingGatewayAddress == nil {
				c.ChargingGatewayAddress = i
			} else if c.AlternativeChargingGatewayAddress == nil {
				c.AlternativeChargingGatewayAddress = i
			}
		case ies.CommonFlags:
			c.CommonFlags = i
		case ies.APNRestriction:
			c.APNRestriction = i
		case ies.MSInfoChangeReportingAction:
			c.MSInfoChangeReportingAction = i
		case ies.BearerControlMode:
			c.BearerControlMode = i
		case ies.EvolvedAllocationRetentionPriorityI:
			c.EvolvedAllocationRetentionPriorityI = i
		case ies.ExtendedCommonFlags:
			c.ExtendedCommonFlag = i
		case ies.CSGInformationReportingAction:
			c.CSGInformationReportingAction = i
		case ies.AggregateMaximumBitRate:
			c.APNAMBR = i
		case ies.GGSNBackOffTime:
			c.GGSNBackOffTime = i
		case ies.ExtendedCommonFlagsII:
			c.ExtendedCommonFlagsII = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// Len returns the actual length of Data.
func (c *CreatePDPContextResponse) Len() int {
	l := c.Header.Len() - len(c.Header.Payload)

	if ie := c.Cause; ie != nil {
		l += ie.Len()
	}
	if ie := c.ReorderingRequired; ie != nil {
		l += ie.Len()
	}
	if ie := c.Recovery; ie != nil {
		l += ie.Len()
	}
	if ie := c.TEIDDataI; ie != nil {
		l += ie.Len()
	}
	if ie := c.TEIDCPlane; ie != nil {
		l += ie.Len()
	}
	if ie := c.NSAPI; ie != nil {
		l += ie.Len()
	}
	if ie := c.ChargingID; ie != nil {
		l += ie.Len()
	}
	if ie := c.EndUserAddress; ie != nil {
		l += ie.Len()
	}
	if ie := c.PCO; ie != nil {
		l += ie.Len()
	}
	if ie := c.GGSNAddressForCPlane; ie != nil {
		l += ie.Len()
	}
	if ie := c.GGSNAddressForUsertraffic; ie != nil {
		l += ie.Len()
	}
	if ie := c.AlternativeGGSNAddressForCPlane; ie != nil {
		l += ie.Len()
	}
	if ie := c.AlternativeGGSNAddressForUsertraffic; ie != nil {
		l += ie.Len()
	}
	if ie := c.QoSProfile; ie != nil {
		l += ie.Len()
	}
	if ie := c.ChargingGatewayAddress; ie != nil {
		l += ie.Len()
	}
	if ie := c.AlternativeChargingGatewayAddress; ie != nil {
		l += ie.Len()
	}
	if ie := c.CommonFlags; ie != nil {
		l += ie.Len()
	}
	if ie := c.APNRestriction; ie != nil {
		l += ie.Len()
	}
	if ie := c.MSInfoChangeReportingAction; ie != nil {
		l += ie.Len()
	}
	if ie := c.BearerControlMode; ie != nil {
		l += ie.Len()
	}
	if ie := c.EvolvedAllocationRetentionPriorityI; ie != nil {
		l += ie.Len()
	}
	if ie := c.ExtendedCommonFlag; ie != nil {
		l += ie.Len()
	}
	if ie := c.CSGInformationReportingAction; ie != nil {
		l += ie.Len()
	}
	if ie := c.APNAMBR; ie != nil {
		l += ie.Len()
	}
	if ie := c.GGSNBackOffTime; ie != nil {
		l += ie.Len()
	}
	if ie := c.ExtendedCommonFlagsII; ie != nil {
		l += ie.Len()
	}
	if ie := c.PrivateExtension; ie != nil {
		l += ie.Len()
	}

	for _, ie := range c.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.Len()
	}
	return l
}

// SetLength sets the length in Length field.
func (c *CreatePDPContextResponse) SetLength() {
	c.Length = uint16(c.Len() - 8)
}

// MessageTypeName returns the name of protocol.
func (c *CreatePDPContextResponse) MessageTypeName() string {
	return "Create PDP Context Response"
}

// TEID returns the TEID in human-readable string.
func (c *CreatePDPContextResponse) TEID() uint32 {
	return c.Header.TEID
}
