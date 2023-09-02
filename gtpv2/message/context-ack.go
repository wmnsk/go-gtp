// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// ContextAcknowledge is a ContextAcknowledge Header and its IEs above.
type ContextAcknowledge struct {
	*Header
	Cause                  *ie.IE
	IndicationFlags        *ie.IE
	ForwardingFTEID        *ie.IE
	BearerContexts         []*ie.IE
	SGSNNumber             *ie.IE
	MMENumberForMTSMS      *ie.IE
	SGSNIdentifierForMTSMS *ie.IE
	MMEIdentifierForMTSMS  *ie.IE
	PrivateExtension       *ie.IE
	AdditionalIEs          []*ie.IE
}

// NewContextAcknowledge creates a new ContextAcknowledge.
func NewContextAcknowledge(teid, seq uint32, ies ...*ie.IE) *ContextAcknowledge {
	c := &ContextAcknowledge{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeContextAcknowledge, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			c.Cause = i
		case ie.Indication:
			c.IndicationFlags = i
		case ie.FullyQualifiedTEID:
			c.ForwardingFTEID = i
		case ie.BearerContext:
			c.BearerContexts = append(c.BearerContexts, i)
		case ie.NodeNumber:
			switch i.Instance() {
			case 0:
				c.SGSNNumber = i
			case 1:
				c.MMENumberForMTSMS = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.NodeIdentifier:
			switch i.Instance() {
			case 0:
				c.SGSNIdentifierForMTSMS = i
			case 1:
				c.MMEIdentifierForMTSMS = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.PrivateExtension:
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
	for _, ie := range c.BearerContexts {
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

	decodedIEs, err := ie.ParseMultiIEs(c.Header.Payload)
	if err != nil {
		return err
	}
	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			c.Cause = i
		case ie.Indication:
			c.IndicationFlags = i
		case ie.FullyQualifiedTEID:
			c.ForwardingFTEID = i
		case ie.BearerContext:
			c.BearerContexts = append(c.BearerContexts, i)
		case ie.NodeNumber:
			switch i.Instance() {
			case 0:
				c.SGSNNumber = i
			case 1:
				c.MMENumberForMTSMS = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.NodeIdentifier:
			switch i.Instance() {
			case 0:
				c.SGSNIdentifierForMTSMS = i
			case 1:
				c.MMEIdentifierForMTSMS = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ie.PrivateExtension:
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
	for _, ie := range c.BearerContexts {
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
