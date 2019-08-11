// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v0/ies"
)

// CreatePDPContextRequest is a CreatePDPContextRequest Header and its IEs above.
type CreatePDPContextRequest struct {
	*Header
	RAI                       *ies.IE
	QoSProfile                *ies.IE
	Recovery                  *ies.IE
	SelectionMode             *ies.IE
	FlowLabelDataI            *ies.IE
	FlowLabelSignalling       *ies.IE
	EndUserAddress            *ies.IE
	APN                       *ies.IE
	PCO                       *ies.IE
	SGSNAddressForSignalling  *ies.IE
	SGSNAddressForUserTraffic *ies.IE
	MSISDN                    *ies.IE
	PrivateExtension          *ies.IE
	AdditionalIEs             []*ies.IE
}

// NewCreatePDPContextRequest creates a new CreatePDPContextRequest.
func NewCreatePDPContextRequest(seq, label uint16, tid uint64, ie ...*ies.IE) *CreatePDPContextRequest {
	c := &CreatePDPContextRequest{
		Header: NewHeader(
			0x1e, MsgTypeCreatePDPContextRequest, seq, label, tid, nil,
		),
	}

	// Optional IEs and Private Extensions, or any arbitrary IE.
	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.RouteingAreaIdentity:
			c.RAI = i
		case ies.QualityOfServiceProfile:
			c.QoSProfile = i
		case ies.Recovery:
			c.Recovery = i
		case ies.SelectionMode:
			c.SelectionMode = i
		case ies.FlowLabelDataI:
			c.FlowLabelDataI = i
		case ies.FlowLabelSignalling:
			c.FlowLabelSignalling = i
		case ies.EndUserAddress:
			c.EndUserAddress = i
		case ies.AccessPointName:
			c.APN = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.GSNAddress:
			if c.SGSNAddressForSignalling == nil {
				c.SGSNAddressForSignalling = i
			} else {
				c.SGSNAddressForUserTraffic = i
			}
		case ies.MSISDN:
			c.MSISDN = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Marshal returns the byte sequence generated from a CreatePDPContextRequest.
func (c *CreatePDPContextRequest) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (c *CreatePDPContextRequest) MarshalTo(b []byte) error {
	if c.Header.Payload != nil {
		c.Header.Payload = nil
	}
	c.Header.Payload = make([]byte, c.MarshalLen()-c.Header.MarshalLen())

	offset := 0
	if ie := c.RAI; ie != nil {
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
	if ie := c.Recovery; ie != nil {
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
	if ie := c.EndUserAddress; ie != nil {
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
	if ie := c.PCO; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGSNAddressForSignalling; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGSNAddressForUserTraffic; ie != nil {
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

// ParseCreatePDPContextRequest parses a given byte sequence as a CreatePDPContextRequest.
func ParseCreatePDPContextRequest(b []byte) (*CreatePDPContextRequest, error) {
	c := &CreatePDPContextRequest{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary parses a given byte sequence as a CreatePDPContextRequest.
func (c *CreatePDPContextRequest) UnmarshalBinary(b []byte) error {
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
		case ies.RouteingAreaIdentity:
			c.RAI = i
		case ies.QualityOfServiceProfile:
			c.QoSProfile = i
		case ies.Recovery:
			c.Recovery = i
		case ies.SelectionMode:
			c.SelectionMode = i
		case ies.FlowLabelDataI:
			c.FlowLabelDataI = i
		case ies.FlowLabelSignalling:
			c.FlowLabelSignalling = i
		case ies.EndUserAddress:
			c.EndUserAddress = i
		case ies.AccessPointName:
			c.APN = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.GSNAddress:
			if c.SGSNAddressForSignalling == nil {
				c.SGSNAddressForSignalling = i
			} else {
				c.SGSNAddressForUserTraffic = i
			}
		case ies.MSISDN:
			c.MSISDN = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length of Data.
func (c *CreatePDPContextRequest) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)

	if ie := c.RAI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.QoSProfile; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SelectionMode; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.FlowLabelDataI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.FlowLabelSignalling; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.EndUserAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.APN; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGSNAddressForSignalling; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGSNAddressForUserTraffic; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MSISDN; ie != nil {
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
func (c *CreatePDPContextRequest) SetLength() {
	c.Header.Length = uint16(c.MarshalLen() - 20)
}

// MessageTypeName returns the name of protocol.
func (c *CreatePDPContextRequest) MessageTypeName() string {
	return "Create PDP Context Request"
}

// TID returns the TID in human-readable string.
func (c *CreatePDPContextRequest) TID() string {
	return c.tid()
}
