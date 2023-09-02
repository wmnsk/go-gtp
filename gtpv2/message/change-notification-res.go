// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "github.com/wmnsk/go-gtp/gtpv2/ie"

// ChangeNotificationResponse is a ChangeNotificationResponse Header and its IEs above.
type ChangeNotificationResponse struct {
	*Header
	IMSI                          *ie.IE
	MEI                           *ie.IE
	Cause                         *ie.IE
	ChangeReportingAction         *ie.IE
	CSGInformationReportingAction *ie.IE
	PresenceReportingAreaAction   []*ie.IE
	PrivateExtension              *ie.IE
	AdditionalIEs                 []*ie.IE
}

// NewChangeNotificationResponse creates a new ChangeNotificationResponse.
func NewChangeNotificationResponse(teid, seq uint32, ies ...*ie.IE) *ChangeNotificationResponse {
	c := &ChangeNotificationResponse{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeChangeNotificationResponse, teid, seq, nil,
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
		case ie.Cause:
			c.Cause = i
		case ie.ChangeReportingAction:
			c.ChangeReportingAction = i
		case ie.CSGInformationReportingAction:
			c.CSGInformationReportingAction = i
		case ie.PresenceReportingAreaAction:
			c.PresenceReportingAreaAction = append(c.PresenceReportingAreaAction, i)
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Marshal serializes ChangeNotificationResponse into bytes.
func (c *ChangeNotificationResponse) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes ChangeNotificationResponse into bytes.
func (c *ChangeNotificationResponse) MarshalTo(b []byte) error {
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
	if ie := c.Cause; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ChangeReportingAction; ie != nil {
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
	for _, ie := range c.PresenceReportingAreaAction {
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

// ParseChangeNotificationResponse decodes given bytes as ChangeNotificationResponse.
func ParseChangeNotificationResponse(b []byte) (*ChangeNotificationResponse, error) {
	c := &ChangeNotificationResponse{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes given bytes as ChangeNotificationResponse.
func (c *ChangeNotificationResponse) UnmarshalBinary(b []byte) error {
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
		case ie.IMSI:
			c.IMSI = i
		case ie.MobileEquipmentIdentity:
			c.MEI = i
		case ie.Cause:
			c.Cause = i
		case ie.ChangeReportingAction:
			c.ChangeReportingAction = i
		case ie.CSGInformationReportingAction:
			c.CSGInformationReportingAction = i
		case ie.PresenceReportingAreaAction:
			c.PresenceReportingAreaAction = append(c.PresenceReportingAreaAction, i)
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (c *ChangeNotificationResponse) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)

	if ie := c.IMSI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MEI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ChangeReportingAction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.CSGInformationReportingAction; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range c.PresenceReportingAreaAction {
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
func (c *ChangeNotificationResponse) SetLength() {
	c.Header.Length = uint16(c.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *ChangeNotificationResponse) MessageTypeName() string {
	return "Change Notification Response"
}

// TEID returns the TEID in uint32.
func (c *ChangeNotificationResponse) TEID() uint32 {
	return c.Header.teid()
}
