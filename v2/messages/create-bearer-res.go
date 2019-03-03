// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v2/ies"
)

// CreateBearerResponse is a CreateBearerResponse Header and its IEs above.
type CreateBearerResponse struct {
	*Header
	Cause                              *ies.IE
	BearerContexts                     *ies.IE
	Recovery                           *ies.IE
	MMEFQCSID                          *ies.IE
	SGWFQCSID                          *ies.IE
	EPDGFQCSID                         *ies.IE
	TWANFQCSID                         *ies.IE
	PCO                                *ies.IE
	UETimeZone                         *ies.IE
	ULI                                *ies.IE
	TWANIdentifier                     *ies.IE
	TWANIdentifierTimestamp            *ies.IE
	MMEOverloadControlInformation      *ies.IE
	SGWOverloadControlInformation      *ies.IE
	PresenceReportingAction            *ies.IE
	MMESGSNIdentifier                  *ies.IE
	TWANePDGOverloadControlInformation *ies.IE
	WLANLocationInformation            *ies.IE
	WLANLocationTimestamp              *ies.IE
	UELocalIPAddress                   *ies.IE
	UEUDPPort                          *ies.IE
	NBIFOMContainer                    *ies.IE
	UETCPPort                          *ies.IE
	PrivateExtension                   *ies.IE
	AdditionalIEs                      []*ies.IE
}

// NewCreateBearerResponse creates a new CreateBearerResponse.
func NewCreateBearerResponse(teid, seq uint32, ie ...*ies.IE) *CreateBearerResponse {
	c := &CreateBearerResponse{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeCreateBearerResponse, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			c.Cause = i
		case ies.BearerContext:
			c.BearerContexts = i
		case ies.Recovery:
			c.Recovery = i
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
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.UETimeZone:
			c.UETimeZone = i
		case ies.UserLocationInformation:
			c.ULI = i
		case ies.TWANIdentifier:
			switch i.Instance() {
			case 0:
				c.TWANIdentifier = i
			case 1:
				c.WLANLocationInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.TWANIdentifierTimestamp:
			switch i.Instance() {
			case 0:
				c.TWANIdentifierTimestamp = i
			case 1:
				c.WLANLocationTimestamp = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				c.MMEOverloadControlInformation = i
			case 1:
				c.SGWOverloadControlInformation = i
			case 2:
				c.TWANePDGOverloadControlInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.PresenceReportingAreaAction:
			c.PresenceReportingAction = i
		case ies.IPAddress:
			switch i.Instance() {
			case 0:
				c.MMESGSNIdentifier = i
			case 1:
				c.UELocalIPAddress = i
			}
		case ies.PortNumber:
			switch i.Instance() {
			case 0:
				c.UEUDPPort = i
			case 1:
				c.UETCPPort = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.FContainer:
			c.NBIFOMContainer = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Serialize serializes CreateBearerResponse into bytes.
func (c *CreateBearerResponse) Serialize() ([]byte, error) {
	b := make([]byte, c.Len())
	if err := c.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes CreateBearerResponse into bytes.
func (c *CreateBearerResponse) SerializeTo(b []byte) error {
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
	if ie := c.BearerContexts; ie != nil {
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
	if ie := c.MMEFQCSID; ie != nil {
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
	if ie := c.EPDGFQCSID; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.TWANFQCSID; ie != nil {
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
	if ie := c.UETimeZone; ie != nil {
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
	if ie := c.TWANIdentifier; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.TWANIdentifierTimestamp; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.MMEOverloadControlInformation; ie != nil {
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
	if ie := c.PresenceReportingAction; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.MMESGSNIdentifier; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.TWANePDGOverloadControlInformation; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.WLANLocationInformation; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.WLANLocationTimestamp; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.UELocalIPAddress; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.UEUDPPort; ie != nil {
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
	if ie := c.UETCPPort; ie != nil {
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

// DecodeCreateBearerResponse decodes given bytes as CreateBearerResponse.
func DecodeCreateBearerResponse(b []byte) (*CreateBearerResponse, error) {
	c := &CreateBearerResponse{}
	if err := c.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return c, nil
}

// DecodeFromBytes decodes given bytes as CreateBearerResponse.
func (c *CreateBearerResponse) DecodeFromBytes(b []byte) error {
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
		case ies.BearerContext:
			c.BearerContexts = i
		case ies.Recovery:
			c.Recovery = i
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
			}
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.UETimeZone:
			c.UETimeZone = i
		case ies.UserLocationInformation:
			c.ULI = i
		case ies.TWANIdentifier:
			switch i.Instance() {
			case 0:
				c.TWANIdentifier = i
			case 1:
				c.WLANLocationInformation = i
			}
		case ies.TWANIdentifierTimestamp:
			switch i.Instance() {
			case 0:
				c.TWANIdentifierTimestamp = i
			case 1:
				c.WLANLocationTimestamp = i
			}
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				c.MMEOverloadControlInformation = i
			case 1:
				c.SGWOverloadControlInformation = i
			case 2:
				c.TWANePDGOverloadControlInformation = i
			}
		case ies.PresenceReportingAreaAction:
			c.PresenceReportingAction = i
		case ies.IPAddress:
			switch i.Instance() {
			case 0:
				c.MMESGSNIdentifier = i
			case 1:
				c.UELocalIPAddress = i
			}
		case ies.PortNumber:
			switch i.Instance() {
			case 0:
				c.UEUDPPort = i
			case 1:
				c.UETCPPort = i
			}
		case ies.FContainer:
			c.NBIFOMContainer = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// Len returns the actual length in int.
func (c *CreateBearerResponse) Len() int {
	l := c.Header.Len() - len(c.Header.Payload)

	if ie := c.Cause; ie != nil {
		l += ie.Len()
	}
	if ie := c.BearerContexts; ie != nil {
		l += ie.Len()
	}
	if ie := c.Recovery; ie != nil {
		l += ie.Len()
	}
	if ie := c.MMEFQCSID; ie != nil {
		l += ie.Len()
	}
	if ie := c.SGWFQCSID; ie != nil {
		l += ie.Len()
	}
	if ie := c.EPDGFQCSID; ie != nil {
		l += ie.Len()
	}
	if ie := c.TWANFQCSID; ie != nil {
		l += ie.Len()
	}
	if ie := c.PCO; ie != nil {
		l += ie.Len()
	}
	if ie := c.UETimeZone; ie != nil {
		l += ie.Len()
	}
	if ie := c.ULI; ie != nil {
		l += ie.Len()
	}
	if ie := c.TWANIdentifier; ie != nil {
		l += ie.Len()
	}
	if ie := c.TWANIdentifierTimestamp; ie != nil {
		l += ie.Len()
	}
	if ie := c.MMEOverloadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := c.SGWOverloadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := c.PresenceReportingAction; ie != nil {
		l += ie.Len()
	}
	if ie := c.MMESGSNIdentifier; ie != nil {
		l += ie.Len()
	}
	if ie := c.TWANePDGOverloadControlInformation; ie != nil {
		l += ie.Len()
	}
	if ie := c.WLANLocationInformation; ie != nil {
		l += ie.Len()
	}
	if ie := c.WLANLocationTimestamp; ie != nil {
		l += ie.Len()
	}
	if ie := c.UELocalIPAddress; ie != nil {
		l += ie.Len()
	}
	if ie := c.UEUDPPort; ie != nil {
		l += ie.Len()
	}
	if ie := c.NBIFOMContainer; ie != nil {
		l += ie.Len()
	}
	if ie := c.UETCPPort; ie != nil {
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
func (c *CreateBearerResponse) SetLength() {
	c.Header.Length = uint16(c.Len() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *CreateBearerResponse) MessageTypeName() string {
	return "Create Bearer Response"
}

// TEID returns the TEID in uint32.
func (c *CreateBearerResponse) TEID() uint32 {
	return c.Header.teid()
}
