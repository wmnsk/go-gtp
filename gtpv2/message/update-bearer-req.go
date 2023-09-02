// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// UpdateBearerRequest is a UpdateBearerRequest Header and its IEs above.
type UpdateBearerRequest struct {
	*Header
	BearerContexts                []*ie.IE
	PTI                           *ie.IE
	PCO                           *ie.IE
	APNAMBR                       *ie.IE
	ChangeReportingAction         *ie.IE
	CSGInformationReportingAction *ie.IE
	HeNBInformationReporting      *ie.IE
	IndicationFlags               *ie.IE
	PGWFQCSID                     *ie.IE
	SGWFQCSID                     *ie.IE
	PresenceReportingAction       []*ie.IE
	PGWNodeLoadControlInformation *ie.IE
	PGWAPNLoadControlInformation  *ie.IE
	SGWNodeLoadControlInformation *ie.IE
	PGWOverloadControlInformation *ie.IE
	SGWOverloadControlInformation *ie.IE
	NBIFOMContainer               *ie.IE
	PrivateExtension              *ie.IE
	AdditionalIEs                 []*ie.IE
}

// NewUpdateBearerRequest creates a new UpdateBearerRequest.
func NewUpdateBearerRequest(teid, seq uint32, ies ...*ie.IE) *UpdateBearerRequest {
	c := &UpdateBearerRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeUpdateBearerRequest, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.BearerContext:
			c.BearerContexts = append(c.BearerContexts, i)
		case ie.ProcedureTransactionID:
			c.PTI = i
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.AggregateMaximumBitRate:
			c.APNAMBR = i
		case ie.ChangeReportingAction:
			c.ChangeReportingAction = i
		case ie.CSGInformationReportingAction:
			c.CSGInformationReportingAction = i
		case ie.HeNBInformationReporting:
			c.HeNBInformationReporting = i
		case ie.Indication:
			c.IndicationFlags = i
		case ie.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				c.PGWFQCSID = i
			case 1:
				c.SGWFQCSID = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.PresenceReportingAreaAction:
			c.PresenceReportingAction = append(c.PresenceReportingAction, i)
		case ie.LoadControlInformation:
			switch i.Instance() {
			case 0:
				c.PGWNodeLoadControlInformation = i
			case 1:
				c.PGWAPNLoadControlInformation = i
			case 2:
				c.SGWNodeLoadControlInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				c.PGWOverloadControlInformation = i
			case 1:
				c.SGWOverloadControlInformation = i
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

// Marshal serializes UpdateBearerRequest into bytes.
func (c *UpdateBearerRequest) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes UpdateBearerRequest into bytes.
func (c *UpdateBearerRequest) MarshalTo(b []byte) error {
	if c.Header.Payload != nil {
		c.Header.Payload = nil
	}
	c.Header.Payload = make([]byte, c.MarshalLen()-c.Header.MarshalLen())

	offset := 0
	for _, ie := range c.BearerContexts {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PTI; ie != nil {
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
	if ie := c.APNAMBR; ie != nil {
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
	if ie := c.HeNBInformationReporting; ie != nil {
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
	if ie := c.PGWFQCSID; ie != nil {
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
	for _, ie := range c.PresenceReportingAction {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PGWNodeLoadControlInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PGWAPNLoadControlInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGWNodeLoadControlInformation; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.PGWOverloadControlInformation; ie != nil {
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
	if ie := c.NBIFOMContainer; ie != nil {
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

// ParseUpdateBearerRequest decodes given bytes as UpdateBearerRequest.
func ParseUpdateBearerRequest(b []byte) (*UpdateBearerRequest, error) {
	c := &UpdateBearerRequest{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes given bytes as UpdateBearerRequest.
func (c *UpdateBearerRequest) UnmarshalBinary(b []byte) error {
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
		case ie.BearerContext:
			c.BearerContexts = append(c.BearerContexts, i)
		case ie.ProcedureTransactionID:
			c.PTI = i
		case ie.ProtocolConfigurationOptions:
			c.PCO = i
		case ie.AggregateMaximumBitRate:
			c.APNAMBR = i
		case ie.ChangeReportingAction:
			c.ChangeReportingAction = i
		case ie.CSGInformationReportingAction:
			c.CSGInformationReportingAction = i
		case ie.HeNBInformationReporting:
			c.HeNBInformationReporting = i
		case ie.Indication:
			c.IndicationFlags = i
		case ie.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				c.PGWFQCSID = i
			case 1:
				c.SGWFQCSID = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.PresenceReportingAreaAction:
			c.PresenceReportingAction = append(c.PresenceReportingAction, i)
		case ie.LoadControlInformation:
			switch i.Instance() {
			case 0:
				c.PGWNodeLoadControlInformation = i
			case 1:
				c.PGWAPNLoadControlInformation = i
			case 2:
				c.SGWNodeLoadControlInformation = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				c.PGWOverloadControlInformation = i
			case 1:
				c.SGWOverloadControlInformation = i
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
func (c *UpdateBearerRequest) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)

	for _, ie := range c.BearerContexts {
		l += ie.MarshalLen()
	}
	if ie := c.PTI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.APNAMBR; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ChangeReportingAction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.CSGInformationReportingAction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.HeNBInformationReporting; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PGWFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	for _, ie := range c.PresenceReportingAction {
		l += ie.MarshalLen()
	}
	if ie := c.PGWNodeLoadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PGWAPNLoadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWNodeLoadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PGWOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWOverloadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.NBIFOMContainer; ie != nil {
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
func (c *UpdateBearerRequest) SetLength() {
	c.Header.Length = uint16(c.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *UpdateBearerRequest) MessageTypeName() string {
	return "Update Bearer Request"
}

// TEID returns the TEID in uint32.
func (c *UpdateBearerRequest) TEID() uint32 {
	return c.Header.teid()
}
