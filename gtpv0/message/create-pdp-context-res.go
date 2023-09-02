// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv0/ie"
)

// CreatePDPContextResponse is a CreatePDPContextResponse Header and its IEs above.
type CreatePDPContextResponse struct {
	*Header
	Cause                     *ie.IE
	QoSProfile                *ie.IE
	ReorderingRequired        *ie.IE
	Recovery                  *ie.IE
	FlowLabelDataI            *ie.IE
	FlowLabelSignalling       *ie.IE
	ChargingID                *ie.IE
	EndUserAddress            *ie.IE
	PCO                       *ie.IE
	GGSNAddressForSignalling  *ie.IE
	GGSNAddressForUserTraffic *ie.IE
	ChargingGatewayAddress    *ie.IE
	PrivateExtension          *ie.IE
	AdditionalIEs             []*ie.IE
}

// NewCreatePDPContextResponse creates a new CreatePDPContextResponse.
func NewCreatePDPContextResponse(seq, label uint16, tid uint64, ies ...*ie.IE) *CreatePDPContextResponse {
	c := &CreatePDPContextResponse{
		Header: NewHeader(
			0x1e, MsgTypeCreatePDPContextResponse, seq, label, tid, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			c.Cause = i
		case ie.QualityOfServiceProfile:
			c.QoSProfile = i
		case ie.ReorderingRequired:
			c.ReorderingRequired = i
		case ie.Recovery:
			c.Recovery = i
		case ie.FlowLabelDataI:
			c.FlowLabelDataI = i
		case ie.FlowLabelSignalling:
			c.FlowLabelSignalling = i
		case ie.ChargingID:
			c.ChargingID = i
		case ie.EndUserAddress:
			c.EndUserAddress = i
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.GSNAddress:
			if c.GGSNAddressForSignalling == nil {
				c.GGSNAddressForSignalling = i
			} else {
				c.GGSNAddressForUserTraffic = i
			}
		case ie.ChargingGatewayAddress:
			c.ChargingGatewayAddress = i
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

// ParseCreatePDPContextResponse parses a given byte sequence as a CreatePDPContextResponse.
func ParseCreatePDPContextResponse(b []byte) (*CreatePDPContextResponse, error) {
	c := &CreatePDPContextResponse{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary parses a given byte sequence as a CreatePDPContextResponse.
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
		case ie.QualityOfServiceProfile:
			c.QoSProfile = i
		case ie.ReorderingRequired:
			c.ReorderingRequired = i
		case ie.Recovery:
			c.Recovery = i
		case ie.FlowLabelDataI:
			c.FlowLabelDataI = i
		case ie.FlowLabelSignalling:
			c.FlowLabelSignalling = i
		case ie.ChargingID:
			c.ChargingID = i
		case ie.EndUserAddress:
			c.EndUserAddress = i
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.GSNAddress:
			if c.GGSNAddressForSignalling == nil {
				c.GGSNAddressForSignalling = i
			} else {
				c.GGSNAddressForUserTraffic = i
			}
		case ie.ChargingGatewayAddress:
			c.ChargingGatewayAddress = i
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
