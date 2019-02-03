// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/gtp/v1/ies"
)

// CreatePDPContextRequest is a CreatePDPContextRequest Header and its IEs above.
type CreatePDPContextRequest struct {
	*Header
	IMSI             *ies.IE
	RAI              *ies.IE
	Recovery         *ies.IE
	SelectionMode    *ies.IE
	TEIDU            *ies.IE
	TEIDC            *ies.IE
	NSAPI            *ies.IE
	EndUserAddress   *ies.IE
	APN              *ies.IE
	PCO              *ies.IE
	GSNAddrC         *ies.IE
	GSNAddrU         *ies.IE
	MSISDN           *ies.IE
	QoS              *ies.IE
	CommonFlags      *ies.IE
	RATType          *ies.IE
	ULI              *ies.IE
	MSTimeZone       *ies.IE
	PrivateExtension *ies.IE
	AdditionalIEs    []*ies.IE
}

// NewCreatePDPContextRequest creates a new GTPv1 CreatePDPContextRequest.
func NewCreatePDPContextRequest(teid uint32, seq uint16, ie ...*ies.IE) *CreatePDPContextRequest {
	c := &CreatePDPContextRequest{
		Header: NewHeader(0x32, MsgTypeCreatePDPContextRequest, teid, seq, nil),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.IMSI:
			c.IMSI = i
		case ies.RouteingAreaIdentity:
			c.RAI = i
		case ies.Recovery:
			c.Recovery = i
		case ies.SelectionMode:
			c.SelectionMode = i
		case ies.TEIDDataI:
			c.TEIDU = i
		case ies.TEIDCPlane:
			c.TEIDC = i
		case ies.NSAPI:
			c.NSAPI = i
		case ies.EndUserAddress:
			c.EndUserAddress = i
		case ies.AccessPointName:
			c.APN = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.GSNAddress:
			if c.GSNAddrC == nil {
				c.GSNAddrC = i
			} else {
				c.GSNAddrU = i
			}
		case ies.MSISDN:
			c.MSISDN = i
		case ies.QoSProfile:
			c.QoS = i
		case ies.CommonFlags:
			c.CommonFlags = i
		case ies.RATType:
			c.RATType = i
		case ies.UserLocationInformation:
			c.ULI = i
		case ies.MSTimeZone:
			c.MSTimeZone = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Serialize returns the byte sequence generated from a CreatePDPContextRequest.
func (c *CreatePDPContextRequest) Serialize() ([]byte, error) {
	b := make([]byte, c.Len())
	if err := c.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo puts the byte sequence in the byte array given as b.
func (c *CreatePDPContextRequest) SerializeTo(b []byte) error {
	if len(b) < c.Len() {
		return ErrTooShortToSerialize
	}
	c.Header.Payload = make([]byte, c.Len()-c.Header.Len())

	offset := 0
	if ie := c.IMSI; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.RAI; ie != nil {
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
	if ie := c.SelectionMode; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.TEIDU; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.TEIDC; ie != nil {
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
	if ie := c.EndUserAddress; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.APN; ie != nil {
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
	if ie := c.GSNAddrC; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.GSNAddrU; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.MSISDN; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.QoS; ie != nil {
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
	if ie := c.RATType; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.ULI; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.MSTimeZone; ie != nil {
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

// DecodeCreatePDPContextRequest decodes a given byte sequence as a CreatePDPContextRequest.
func DecodeCreatePDPContextRequest(b []byte) (*CreatePDPContextRequest, error) {
	c := &CreatePDPContextRequest{}
	if err := c.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return c, nil
}

// DecodeFromBytes decodes a given byte sequence as a CreatePDPContextRequest.
func (c *CreatePDPContextRequest) DecodeFromBytes(b []byte) error {
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
		case ies.IMSI:
			c.IMSI = i
		case ies.RouteingAreaIdentity:
			c.RAI = i
		case ies.Recovery:
			c.Recovery = i
		case ies.SelectionMode:
			c.SelectionMode = i
		case ies.TEIDDataI:
			c.TEIDU = i
		case ies.TEIDCPlane:
			c.TEIDC = i
		case ies.NSAPI:
			c.NSAPI = i
		case ies.EndUserAddress:
			c.EndUserAddress = i
		case ies.AccessPointName:
			c.APN = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.GSNAddress:
			if c.GSNAddrC == nil {
				c.GSNAddrC = i
			} else {
				c.GSNAddrU = i
			}
		case ies.MSISDN:
			c.MSISDN = i
		case ies.QoSProfile:
			c.QoS = i
		case ies.CommonFlags:
			c.CommonFlags = i
		case ies.RATType:
			c.RATType = i
		case ies.UserLocationInformation:
			c.ULI = i
		case ies.MSTimeZone:
			c.MSTimeZone = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}
	return nil
}

// Len returns the actual length of Data.
func (c *CreatePDPContextRequest) Len() int {
	l := c.Header.Len() - len(c.Header.Payload)

	if ie := c.IMSI; ie != nil {
		l += ie.Len()
	}
	if ie := c.RAI; ie != nil {
		l += ie.Len()
	}
	if ie := c.Recovery; ie != nil {
		l += ie.Len()
	}
	if ie := c.SelectionMode; ie != nil {
		l += ie.Len()
	}
	if ie := c.TEIDU; ie != nil {
		l += ie.Len()
	}
	if ie := c.TEIDC; ie != nil {
		l += ie.Len()
	}
	if ie := c.NSAPI; ie != nil {
		l += ie.Len()
	}
	if ie := c.EndUserAddress; ie != nil {
		l += ie.Len()
	}
	if ie := c.APN; ie != nil {
		l += ie.Len()
	}
	if ie := c.PCO; ie != nil {
		l += ie.Len()
	}
	if ie := c.GSNAddrC; ie != nil {
		l += ie.Len()
	}
	if ie := c.GSNAddrU; ie != nil {
		l += ie.Len()
	}
	if ie := c.MSISDN; ie != nil {
		l += ie.Len()
	}
	if ie := c.QoS; ie != nil {
		l += ie.Len()
	}
	if ie := c.CommonFlags; ie != nil {
		l += ie.Len()
	}
	if ie := c.RATType; ie != nil {
		l += ie.Len()
	}
	if ie := c.ULI; ie != nil {
		l += ie.Len()
	}
	if ie := c.MSTimeZone; ie != nil {
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
func (c *CreatePDPContextRequest) SetLength() {
	c.Length = uint16(c.Len() - 8)
}

// MessageTypeName returns the name of protocol.
func (c *CreatePDPContextRequest) MessageTypeName() string {
	return "Create PDP Context Request"
}

// TEID returns the TEID in human-readable string.
func (c *CreatePDPContextRequest) TEID() uint32 {
	return c.Header.TEID
}
