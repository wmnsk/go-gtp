// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v0/ies"
)

// CreatePDPContextResponse is a CreatePDPContextResponse Header and its IEs above.
type CreatePDPContextResponse struct {
	*Header
	Cause                     *ies.IE
	QoSProfile                *ies.IE
	ReorderingRequired        *ies.IE
	Recovery                  *ies.IE
	FlowLabelDataI            *ies.IE
	FlowLabelSignalling       *ies.IE
	ChargingID                *ies.IE
	EndUserAddress            *ies.IE
	PCO                       *ies.IE
	GGSNAddressForSignalling  *ies.IE
	GGSNAddressForUserTraffic *ies.IE
	ChargingGatewayAddress    *ies.IE
	PrivateExtension          *ies.IE
	AdditionalIEs             []*ies.IE
}

// NewCreatePDPContextResponse creates a new CreatePDPContextResponse.
func NewCreatePDPContextResponse(seq, label uint16, tid uint64, ie ...*ies.IE) *CreatePDPContextResponse {
	c := &CreatePDPContextResponse{
		Header: NewHeader(
			0x1e, MsgTypeCreatePDPContextResponse, seq, label, tid, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			c.Cause = i
		case ies.QualityOfServiceProfile:
			c.QoSProfile = i
		case ies.ReorderingRequired:
			c.ReorderingRequired = i
		case ies.Recovery:
			c.Recovery = i
		case ies.FlowLabelDataI:
			c.FlowLabelDataI = i
		case ies.FlowLabelSignalling:
			c.FlowLabelSignalling = i
		case ies.ChargingID:
			c.ChargingID = i
		case ies.EndUserAddress:
			c.EndUserAddress = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.GSNAddress:
			if c.GGSNAddressForSignalling == nil {
				c.GGSNAddressForSignalling = i
			} else {
				c.GGSNAddressForUserTraffic = i
			}
		case ies.ChargingGatewayAddress:
			c.ChargingGatewayAddress = i
		case ies.PrivateExtension:
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
	if ie := c.QoSProfile; ie != nil {
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
	if ie := c.FlowLabelDataI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.FlowLabelSignalling; ie != nil {
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
	if ie := c.GGSNAddressForSignalling; ie != nil {
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
	if ie := c.ChargingGatewayAddress; ie != nil {
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

// ParseCreatePDPContextResponse Parses a given byte sequence as a CreatePDPContextResponse.
func ParseCreatePDPContextResponse(b []byte) (*CreatePDPContextResponse, error) {
	c := &CreatePDPContextResponse{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary Parses a given byte sequence as a CreatePDPContextResponse.
func (c *CreatePDPContextResponse) UnmarshalBinary(b []byte) error {
	var err error
	c.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(c.Header.Payload) < 2 {
		return nil
	}

	ie, err := ies.ParseMultiIEs(c.Header.Payload)
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
		case ies.QualityOfServiceProfile:
			c.QoSProfile = i
		case ies.ReorderingRequired:
			c.ReorderingRequired = i
		case ies.Recovery:
			c.Recovery = i
		case ies.FlowLabelDataI:
			c.FlowLabelDataI = i
		case ies.FlowLabelSignalling:
			c.FlowLabelSignalling = i
		case ies.ChargingID:
			c.ChargingID = i
		case ies.EndUserAddress:
			c.EndUserAddress = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.GSNAddress:
			if c.GGSNAddressForSignalling == nil {
				c.GGSNAddressForSignalling = i
			} else {
				c.GGSNAddressForUserTraffic = i
			}
		case ies.ChargingGatewayAddress:
			c.ChargingGatewayAddress = i
		case ies.PrivateExtension:
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
	if ie := c.QoSProfile; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ReorderingRequired; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.FlowLabelDataI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.FlowLabelSignalling; ie != nil {
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
	if ie := c.GGSNAddressForSignalling; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.GGSNAddressForUserTraffic; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ChargingGatewayAddress; ie != nil {
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
	c.Header.Length = uint16(c.MarshalLen() - 20)
}

// MessageTypeName returns the name of protocol.
func (c *CreatePDPContextResponse) MessageTypeName() string {
	return "Create PDP Context Response"
}

// TID returns the TID in human-readable string.
func (c *CreatePDPContextResponse) TID() string {
	return c.tid()
}
