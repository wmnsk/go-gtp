// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// ResumeAcknowledge is a ResumeAcknowledge Header and its Cause and Private Extension IEs.
type ResumeAcknowledge struct {
	*Header
	Cause            *ie.IE
	PrivateExtension *ie.IE
}

// NewResumeAcknowledge creates a new ResumeAcknowledge.
func NewResumeAcknowledge(teid, seq uint32, ies ...*ie.IE) *ResumeAcknowledge {
	c := &ResumeAcknowledge{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeResumeAcknowledge, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Cause:
			c.Cause = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		}
	}

	c.SetLength()
	return c
}

// Marshal serializes ResumeAcknowledge into bytes.
func (c *ResumeAcknowledge) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes ResumeAcknowledge into bytes.
func (c *ResumeAcknowledge) MarshalTo(b []byte) error {
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
	if ie := c.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
	}

	c.Header.SetLength()
	return c.Header.MarshalTo(b)
}

// ParseResumeAcknowledge decodes given bytes as ResumeAcknowledge.
func ParseResumeAcknowledge(b []byte) (*ResumeAcknowledge, error) {
	c := &ResumeAcknowledge{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes given bytes as ResumeAcknowledge.
func (c *ResumeAcknowledge) UnmarshalBinary(b []byte) error {
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
		case ie.PrivateExtension:
			c.PrivateExtension = i
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (c *ResumeAcknowledge) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)

	if ie := c.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length field.
func (c *ResumeAcknowledge) SetLength() {
	c.Header.Length = uint16(c.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *ResumeAcknowledge) MessageTypeName() string {
	return "Resume Acknowledge"
}

// TEID returns the TEID in uint32.
func (c *ResumeAcknowledge) TEID() uint32 {
	return c.Header.teid()
}
