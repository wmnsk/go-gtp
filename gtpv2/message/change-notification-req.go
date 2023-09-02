// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "github.com/wmnsk/go-gtp/gtpv2/ie"

// ChangeNotificationRequest is a ChangeNotificationRequest Header and its IEs above.
type ChangeNotificationRequest struct {
	*Header
	IMSI                             *ie.IE
	MEI                              *ie.IE
	IndicationFlags                  *ie.IE
	RATType                          *ie.IE
	ULI                              *ie.IE
	UCI                              *ie.IE
	PGWS5S8IPAddressForControlPlane  *ie.IE
	LinkedEBI                        *ie.IE
	PresenceReportingAreaInformation []*ie.IE
	MOExceptionDataCounter           *ie.IE
	SecondaryRATUsageDataReport      []*ie.IE
	PrivateExtension                 *ie.IE
	AdditionalIEs                    []*ie.IE
}

// NewChangeNotificationRequest creates a new ChangeNotificationRequest.
func NewChangeNotificationRequest(teid, seq uint32, ies ...*ie.IE) *ChangeNotificationRequest {
	c := &ChangeNotificationRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeChangeNotificationRequest, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.IMSI:
			c.IMSI = i
		case ie.MobileEquipmentIdentity:
			c.MEI = i
		case ie.Indication:
			c.IndicationFlags = i
		case ie.RATType:
			c.RATType = i
		case ie.UserLocationInformation:
			c.ULI = i
		case ie.UserCSGInformation:
			c.UCI = i
		case ie.IPAddress:
			c.PGWS5S8IPAddressForControlPlane = i
		case ie.EPSBearerID:
			c.LinkedEBI = i
		case ie.PresenceReportingAreaInformation:
			c.PresenceReportingAreaInformation = append(c.PresenceReportingAreaInformation, i)
		case ie.Counter:
			c.MOExceptionDataCounter = i
		case ie.SecondaryRATUsageDataReport:
			c.SecondaryRATUsageDataReport = append(c.SecondaryRATUsageDataReport, i)
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Marshal serializes ChangeNotificationRequest into bytes.
func (c *ChangeNotificationRequest) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes ChangeNotificationRequest into bytes.
func (c *ChangeNotificationRequest) MarshalTo(b []byte) error {
	if c.Header.Payload != nil {
		c.Header.Payload = nil
	}
	c.Header.Payload = make([]byte, c.MarshalLen()-c.Header.MarshalLen())

	offset := 0
	if ie := c.IMSI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MEI; ie != nil {
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
	if ie := c.RATType; ie != nil {
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
	if ie := c.UCI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PGWS5S8IPAddressForControlPlane; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.LinkedEBI; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	for _, ie := range c.PresenceReportingAreaInformation {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MOExceptionDataCounter; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	for _, ie := range c.SecondaryRATUsageDataReport {
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

// ParseChangeNotificationRequest decodes given bytes as ChangeNotificationRequest.
func ParseChangeNotificationRequest(b []byte) (*ChangeNotificationRequest, error) {
	c := &ChangeNotificationRequest{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes given bytes as ChangeNotificationRequest.
func (c *ChangeNotificationRequest) UnmarshalBinary(b []byte) error {
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
		case ie.IMSI:
			c.IMSI = i
		case ie.MobileEquipmentIdentity:
			c.MEI = i
		case ie.Indication:
			c.IndicationFlags = i
		case ie.RATType:
			c.RATType = i
		case ie.UserLocationInformation:
			c.ULI = i
		case ie.UserCSGInformation:
			c.UCI = i
		case ie.IPAddress:
			c.PGWS5S8IPAddressForControlPlane = i
		case ie.EPSBearerID:
			c.LinkedEBI = i
		case ie.PresenceReportingAreaInformation:
			c.PresenceReportingAreaInformation = append(c.PresenceReportingAreaInformation, i)
		case ie.Counter:
			c.MOExceptionDataCounter = i
		case ie.SecondaryRATUsageDataReport:
			c.SecondaryRATUsageDataReport = append(c.SecondaryRATUsageDataReport, i)
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (c *ChangeNotificationRequest) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)

	if ie := c.IMSI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MEI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.RATType; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ULI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UCI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PGWS5S8IPAddressForControlPlane; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.LinkedEBI; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range c.PresenceReportingAreaInformation {
		l += ie.MarshalLen()
	}
	if ie := c.MOExceptionDataCounter; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range c.SecondaryRATUsageDataReport {
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
func (c *ChangeNotificationRequest) SetLength() {
	c.Header.Length = uint16(c.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *ChangeNotificationRequest) MessageTypeName() string {
	return "Change Notification Request"
}

// TEID returns the TEID in uint32.
func (c *ChangeNotificationRequest) TEID() uint32 {
	return c.Header.teid()
}
