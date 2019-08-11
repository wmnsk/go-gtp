// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v2/ies"
)

// ContextAcknowledge is a ContextAcknowledge Header and its IEs above.
type ContextAcknowledge struct {
	*Header
	Cause                  *ies.IE
	IndicationFlags        *ies.IE
	ForwardingFTEID        *ies.IE
	BearerContexts         *ies.IE
	SGSNNumber             *ies.IE
	MMENumberForMTSMS      *ies.IE
	SGSNIdentifierForMTSMS *ies.IE
	MMEIdentifierForMTSMS  *ies.IE
	PrivateExtension       *ies.IE
	AdditionalIEs          []*ies.IE
}

// NewContextAcknowledge creates a new ContextAcknowledge.
func NewContextAcknowledge(teid, seq uint32, ie ...*ies.IE) *ContextAcknowledge {
	c := &ContextAcknowledge{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeContextAcknowledge, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			c.Cause = i
		case ies.Indication:
			c.IndicationFlags = i
		case ies.FullyQualifiedTEID:
			c.ForwardingFTEID = i
		case ies.BearerContext:
			c.BearerContexts = i
		case ies.NodeNumber:
			switch i.Instance() {
			case 0:
				c.SGSNNumber = i
			case 1:
				c.MMENumberForMTSMS = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.NodeIdentifier:
			switch i.Instance() {
			case 0:
				c.SGSNIdentifierForMTSMS = i
			case 1:
				c.MMEIdentifierForMTSMS = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Marshal serializes ContextAcknowledge into bytes.
func (c *ContextAcknowledge) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes ContextAcknowledge into bytes.
func (c *ContextAcknowledge) MarshalTo(b []byte) error {
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
	if ie := c.IndicationFlags; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.ForwardingFTEID; ie != nil {
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
	if ie := c.SGSNNumber; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MMENumberForMTSMS; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SGSNIdentifierForMTSMS; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.MMEIdentifierForMTSMS; ie != nil {
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

// ParseContextAcknowledge decodes given bytes as ContextAcknowledge.
func ParseContextAcknowledge(b []byte) (*ContextAcknowledge, error) {
	c := &ContextAcknowledge{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes given bytes as ContextAcknowledge.
func (c *ContextAcknowledge) UnmarshalBinary(b []byte) error {
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
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			c.Cause = i
		case ies.Indication:
			c.IndicationFlags = i
		case ies.FullyQualifiedTEID:
			c.ForwardingFTEID = i
		case ies.BearerContext:
			c.BearerContexts = i
		case ies.NodeNumber:
			switch i.Instance() {
			case 0:
				c.SGSNNumber = i
			case 1:
				c.MMENumberForMTSMS = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.NodeIdentifier:
			switch i.Instance() {
			case 0:
				c.SGSNIdentifierForMTSMS = i
			case 1:
				c.MMEIdentifierForMTSMS = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (c *ContextAcknowledge) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)

	if ie := c.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.ForwardingFTEID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.BearerContexts; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGSNNumber; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MMENumberForMTSMS; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SGSNIdentifierForMTSMS; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.MMEIdentifierForMTSMS; ie != nil {
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
func (c *ContextAcknowledge) SetLength() {
	c.Header.Length = uint16(c.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *ContextAcknowledge) MessageTypeName() string {
	return "Context Acknowledge"
}

// TEID returns the TEID in uint32.
func (c *ContextAcknowledge) TEID() uint32 {
	return c.Header.teid()
}
