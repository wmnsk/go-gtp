// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// UpdateBearerResponse is a UpdateBearerResponse Header and its IEs above.
type UpdateBearerResponse struct {
	*Header
	Cause                              *ie.IE
	BearerContexts                     []*ie.IE
	PCO                                *ie.IE
	Recovery                           *ie.IE
	MMEFQCSID                          *ie.IE
	SGWFQCSID                          *ie.IE
	EPDGFQCSID                         *ie.IE
	TWANFQCSID                         *ie.IE
	IndicationFlags                    *ie.IE
	UETimeZone                         *ie.IE
	ULI                                *ie.IE
	TWANIdentifier                     *ie.IE
	MMESGSNOverloadControlInformation  *ie.IE
	SGWOverloadControlInformation      *ie.IE
	PresenceReportingAreaAction        []*ie.IE
	MMESGSNIdentifier                  *ie.IE
	TWANePDGOverloadControlInformation *ie.IE
	WLANLocationInformation            *ie.IE
	WLANLocationTimeStamp              *ie.IE
	UELocalIPAddress                   *ie.IE
	UEUDPPort                          *ie.IE
	NBIFOMContainer                    *ie.IE
	UETCPPort                          *ie.IE
	PrivateExtension                   *ie.IE
	AdditionalIEs                      []*ie.IE
}

// NewUpdateBearerRequest creates a new UpdateBearerRequest.
func NewUpdateBearerResponse(teid, seq uint32, ies ...*ie.IE) *UpdateBearerResponse {
	c := &UpdateBearerResponse{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeUpdateBearerResponse, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			c.Cause = i
		case ie.BearerContext:
			c.BearerContexts = append(c.BearerContexts, i)
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.Recovery:
			c.Recovery = i
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
		case ie.Indication:
			c.IndicationFlags = i
		case ie.UETimeZone:
			c.UETimeZone = i
		case ie.UserLocationInformation:
			c.ULI = i
		case ie.TWANIdentifier:
			switch i.Instance() {
			case 0:
				c.TWANIdentifier = i
			case 1:
				c.WLANLocationInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
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
		case ie.PresenceReportingAreaAction:
			c.PresenceReportingAreaAction = append(c.PresenceReportingAreaAction, i)
		case ie.IPAddress:
			switch i.Instance() {
			case 0:
				c.MMESGSNIdentifier = i
			case 1:
				c.UELocalIPAddress = i
			}
		case ie.TWANIdentifierTimestamp:
			switch i.Instance() {
			case 1:
				c.WLANLocationTimeStamp = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.PortNumber:
			switch i.Instance() {
			case 0:
				c.UEUDPPort = i
			case 1:
				c.UETCPPort = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.FContainer:
			c.NBIFOMContainer = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Marshal serializes UpdateBearerResponse into bytes.
func (c *UpdateBearerResponse) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes UpdateBearerResponse into bytes.
func (c *UpdateBearerResponse) MarshalTo(b []byte) error {
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
	for _, ie := range c.BearerContexts {
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
	if ie := c.IndicationFlags; ie != nil {
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
	if ie := c.ULI; ie != nil {
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
	for _, ie := range c.PresenceReportingAreaAction {
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
	if ie := c.TWANePDGOverloadControlInformation; ie != nil {
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
	if ie := c.NBIFOMContainer; ie != nil {
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

// ParseUpdateBearerResponse decodes given bytes as UpdateBearerResponse.
func ParseUpdateBearerResponse(b []byte) (*UpdateBearerResponse, error) {
	c := &UpdateBearerResponse{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes given bytes as UpdateBearerResponse.
func (c *UpdateBearerResponse) UnmarshalBinary(b []byte) error {
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
		switch i.Type {
		case ie.Cause:
			c.Cause = i
		case ie.BearerContext:
			c.BearerContexts = append(c.BearerContexts, i)
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.Recovery:
			c.Recovery = i
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
		case ie.Indication:
			c.IndicationFlags = i
		case ie.UETimeZone:
			c.UETimeZone = i
		case ie.UserLocationInformation:
			c.ULI = i
		case ie.TWANIdentifier:
			switch i.Instance() {
			case 0:
				c.TWANIdentifier = i
			case 1:
				c.WLANLocationInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
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
		case ie.PresenceReportingAreaAction:
			c.PresenceReportingAreaAction = append(c.PresenceReportingAreaAction, i)
		case ie.IPAddress:
			switch i.Instance() {
			case 0:
				c.MMESGSNIdentifier = i
			case 1:
				c.UELocalIPAddress = i
			}
		case ie.TWANIdentifierTimestamp:
			switch i.Instance() {
			case 1:
				c.WLANLocationTimeStamp = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.PortNumber:
			switch i.Instance() {
			case 0:
				c.UEUDPPort = i
			case 1:
				c.UETCPPort = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.FContainer:
			c.NBIFOMContainer = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (c *UpdateBearerResponse) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)

	if ie := c.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range c.BearerContexts {
		l += ie.MarshalLen()
	}
	if ie := c.PCO; ie != nil {
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
	if ie := c.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UETimeZone; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ULI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TWANIdentifier; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MMESGSNOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range c.PresenceReportingAreaAction {
		l += ie.MarshalLen()
	}
	if ie := c.MMESGSNIdentifier; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.TWANePDGOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.WLANLocationInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.WLANLocationTimeStamp; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UELocalIPAddress; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UEUDPPort; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.NBIFOMContainer; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UETCPPort; ie != nil {
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
func (c *UpdateBearerResponse) SetLength() {
	c.Header.Length = uint16(c.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *UpdateBearerResponse) MessageTypeName() string {
	return "Update Bearer Response"
}

// TEID returns the TEID in uint32.
func (c *UpdateBearerResponse) TEID() uint32 {
	return c.Header.teid()
}
