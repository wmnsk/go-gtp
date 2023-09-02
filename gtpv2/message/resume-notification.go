// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// ResumeNotification is a ResumeNotification Header and its IEs.
type ResumeNotification struct {
	*Header
	IMSI                       *ie.IE
	LinkedEPSBearerID          *ie.IE
	OriginatingNode            *ie.IE
	SenderFTEIDForControlPlane *ie.IE
	PrivateExtension           *ie.IE
}

// NewResumeNotification creates a new ResumeNotification.
func NewResumeNotification(teid, seq uint32, ies ...*ie.IE) *ResumeNotification {
	c := &ResumeNotification{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeResumeNotification, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.IMSI:
			c.IMSI = i
		case ie.EPSBearerID:
			c.LinkedEPSBearerID = i
		case ie.NodeType:
			c.OriginatingNode = i
		case ie.FullyQualifiedTEID:
			c.SenderFTEIDForControlPlane = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		}
	}

	c.SetLength()
	return c
}

// Marshal serializes ResumeNotification into bytes.
func (c *ResumeNotification) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes ResumeNotification into bytes.
func (c *ResumeNotification) MarshalTo(b []byte) error {
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
	if ie := c.LinkedEPSBearerID; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.OriginatingNode; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SenderFTEIDForControlPlane; ie != nil {
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

// ParseResumeNotification decodes given bytes as ResumeNotification.
func ParseResumeNotification(b []byte) (*ResumeNotification, error) {
	c := &ResumeNotification{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes given bytes as ResumeNotification.
func (c *ResumeNotification) UnmarshalBinary(b []byte) error {
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
		case ie.EPSBearerID:
			c.LinkedEPSBearerID = i
		case ie.NodeType:
			c.OriginatingNode = i
		case ie.FullyQualifiedTEID:
			c.SenderFTEIDForControlPlane = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (c *ResumeNotification) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)
	if ie := c.IMSI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.LinkedEPSBearerID; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.OriginatingNode; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SenderFTEIDForControlPlane; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length field.
func (c *ResumeNotification) SetLength() {
	c.Header.Length = uint16(c.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *ResumeNotification) MessageTypeName() string {
	return "Resume Notification"
}

// TEID returns the TEID in uint32.
func (c *ResumeNotification) TEID() uint32 {
	return c.Header.teid()
}
