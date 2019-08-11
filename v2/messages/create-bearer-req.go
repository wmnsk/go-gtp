// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v2/ies"
)

// CreateBearerRequest is a CreateBearerRequest Header and its IEs above.
type CreateBearerRequest struct {
	*Header
	PTI                           *ies.IE
	LinkedEBI                     *ies.IE
	PCO                           *ies.IE
	BearerContexts                *ies.IE
	PGWFQCSID                     *ies.IE
	SGWFQCSID                     *ies.IE
	ChangeReportingAction         *ies.IE
	CSGInformationReportingAction *ies.IE
	HeNBInformationReporting      *ies.IE
	PresenceReportingAreaAction   *ies.IE
	IndicationFlags               *ies.IE
	PGWNodeLoadControlInformation *ies.IE
	PGWAPNLoadControlInformation  *ies.IE
	SGWNodeLoadControlInformation *ies.IE
	PGWOverloadControlInformation *ies.IE
	SGWOverloadControlInformation *ies.IE
	NBIFOMContainer               *ies.IE
	PrivateExtension              *ies.IE
	AdditionalIEs                 []*ies.IE
}

// NewCreateBearerRequest creates a new CreateBearerRequest.
func NewCreateBearerRequest(teid, seq uint32, ie ...*ies.IE) *CreateBearerRequest {
	c := &CreateBearerRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeCreateBearerRequest, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.ProcedureTransactionID:
			c.PTI = i
		case ies.EPSBearerID:
			c.LinkedEBI = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.BearerContext:
			c.BearerContexts = i
		case ies.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				c.PGWFQCSID = i
			case 1:
				c.SGWFQCSID = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.ChangeReportingAction:
			c.ChangeReportingAction = i
		case ies.CSGInformationReportingAction:
			c.CSGInformationReportingAction = i
		case ies.HeNBInformationReporting:
			c.HeNBInformationReporting = i
		case ies.PresenceReportingAreaAction:
			c.PresenceReportingAreaAction = i
		case ies.Indication:
			c.IndicationFlags = i
		case ies.LoadControlInformation:
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
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				c.PGWOverloadControlInformation = i
			case 1:
				c.SGWOverloadControlInformation = i
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

// Marshal serializes CreateBearerRequest into bytes.
func (c *CreateBearerRequest) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes CreateBearerRequest into bytes.
func (c *CreateBearerRequest) MarshalTo(b []byte) error {
	if c.Header.Payload != nil {
		c.Header.Payload = nil
	}
	c.Header.Payload = make([]byte, c.MarshalLen()-c.Header.MarshalLen())

	offset := 0
	if ie := c.PTI; ie != nil {
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
	if ie := c.PCO; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.BearerContexts; ie != nil {
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
	if ie := c.PresenceReportingAreaAction; ie != nil {
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

// ParseCreateBearerRequest decodes given bytes as CreateBearerRequest.
func ParseCreateBearerRequest(b []byte) (*CreateBearerRequest, error) {
	c := &CreateBearerRequest{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes given bytes as CreateBearerRequest.
func (c *CreateBearerRequest) UnmarshalBinary(b []byte) error {
	var err error
	c.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(c.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ies.ParseMultiIEs(c.Header.Payload)
	if err != nil {
		return err
	}
	for _, i := range decodedIEs {
		switch i.Type {
		case ies.ProcedureTransactionID:
			c.PTI = i
		case ies.EPSBearerID:
			c.LinkedEBI = i
		case ies.ProtocolConfigurationOptions:
			c.PCO = i
		case ies.BearerContext:
			c.BearerContexts = i
		case ies.FullyQualifiedCSID:
			switch i.Instance() {
			case 0:
				c.PGWFQCSID = i
			case 1:
				c.SGWFQCSID = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.ChangeReportingAction:
			c.ChangeReportingAction = i
		case ies.CSGInformationReportingAction:
			c.CSGInformationReportingAction = i
		case ies.HeNBInformationReporting:
			c.HeNBInformationReporting = i
		case ies.PresenceReportingAreaAction:
			c.PresenceReportingAreaAction = i
		case ies.Indication:
			c.IndicationFlags = i
		case ies.LoadControlInformation:
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
		case ies.OverloadControlInformation:
			switch i.Instance() {
			case 0:
				c.PGWOverloadControlInformation = i
			case 1:
				c.SGWOverloadControlInformation = i
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

	return nil
}

// MarshalLen returns the serial length in int.
func (c *CreateBearerRequest) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)

	if ie := c.PTI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.LinkedEBI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PCO; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.BearerContexts; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PGWFQCSID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGWFQCSID; ie != nil {
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
	if ie := c.PresenceReportingAreaAction; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.IndicationFlags; ie != nil {
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
func (c *CreateBearerRequest) SetLength() {
	c.Header.Length = uint16(c.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *CreateBearerRequest) MessageTypeName() string {
	return "Create Bearer Request"
}

// TEID returns the TEID in uint32.
func (c *CreateBearerRequest) TEID() uint32 {
	return c.Header.teid()
}
