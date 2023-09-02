// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// SuspendNotification is a SuspendNotification Header and its IEs above.
type SuspendNotification struct {
	*Header
	IMSI                   *ie.IE
	RAI                    *ie.IE
	LinkedEBI              *ie.IE
	PTMSI                  *ie.IE
	OriginatingNode        *ie.IE
	AddressForControlPlane *ie.IE
	UDPSourcePortNumber    *ie.IE
	HopCounter             *ie.IE
	SenderFTEIDC           *ie.IE
	PrivateExtension       *ie.IE
	AdditionalIEs          []*ie.IE
}

// NewSuspendNotification creates a new SuspendNotification.
func NewSuspendNotification(teid, seq uint32, ies ...*ie.IE) *SuspendNotification {
	c := &SuspendNotification{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeSuspendNotification, teid, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.IMSI:
			c.IMSI = i
		case ie.UserLocationInformation:
			c.RAI = i
		case ie.EPSBearerID:
			c.LinkedEBI = i
		case ie.PacketTMSI:
			c.PTMSI = i
		case ie.NodeType:
			c.OriginatingNode = i
		case ie.IPAddress:
			c.AddressForControlPlane = i
		case ie.PortNumber:
			c.UDPSourcePortNumber = i
		case ie.HopCounter:
			c.HopCounter = i
		case ie.FullyQualifiedTEID:
			c.SenderFTEIDC = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Marshal serializes SuspendNotification into bytes.
func (c *SuspendNotification) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes SuspendNotification into bytes.
func (c *SuspendNotification) MarshalTo(b []byte) error {
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
	if ie := c.RAI; ie != nil {
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
	if ie := c.PTMSI; ie != nil {
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
	if ie := c.AddressForControlPlane; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.UDPSourcePortNumber; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.HopCounter; ie != nil {
		if err := ie.MarshalTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := c.SenderFTEIDC; ie != nil {
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

// ParseSuspendNotification decodes given bytes as SuspendNotification.
func ParseSuspendNotification(b []byte) (*SuspendNotification, error) {
	c := &SuspendNotification{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalBinary decodes given bytes as SuspendNotification.
func (c *SuspendNotification) UnmarshalBinary(b []byte) error {
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
		case ie.UserLocationInformation:
			c.RAI = i
		case ie.PacketTMSI:
			c.PTMSI = i
		case ie.EPSBearerID:
			c.LinkedEBI = i
		case ie.NodeType:
			c.OriginatingNode = i
		case ie.IPAddress:
			c.AddressForControlPlane = i
		case ie.PortNumber:
			c.UDPSourcePortNumber = i
		case ie.HopCounter:
			c.HopCounter = i
		case ie.FullyQualifiedTEID:
			c.SenderFTEIDC = i
		case ie.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (c *SuspendNotification) MarshalLen() int {
	l := c.Header.MarshalLen() - len(c.Header.Payload)
	if ie := c.IMSI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.RAI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.PTMSI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.LinkedEBI; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.OriginatingNode; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.AddressForControlPlane; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.UDPSourcePortNumber; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.HopCounter; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := c.SenderFTEIDC; ie != nil {
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
func (c *SuspendNotification) SetLength() {
	c.Header.Length = uint16(c.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *SuspendNotification) MessageTypeName() string {
	return "Suspend Notification"
}

// TEID returns the TEID in uint32.
func (c *SuspendNotification) TEID() uint32 {
	return c.Header.teid()
}
