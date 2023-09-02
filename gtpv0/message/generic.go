// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"fmt"

	"github.com/wmnsk/go-gtp/gtpv0/ie"
)

// Generic is a Generic Header and its IEs above.
type Generic struct {
	*Header
	IEs []*ie.IE
}

// NewGeneric creates a new GTPv0 Generic.
func NewGeneric(msgType uint8, seq, label uint16, tid uint64, ie ...*ie.IE) *Generic {
	g := &Generic{
		Header: NewHeader(0x1e, msgType, seq, label, tid, nil),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		g.IEs = append(g.IEs, i)
	}

	g.SetLength()
	return g
}

// Marshal returns the byte sequence generated from a Generic.
func (g *Generic) Marshal() ([]byte, error) {
	b := make([]byte, g.MarshalLen())
	if err := g.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (g *Generic) MarshalTo(b []byte) error {
	if g.Header.Payload != nil {
		g.Header.Payload = nil
	}
	g.Header.Payload = make([]byte, g.MarshalLen()-g.Header.MarshalLen())

	offset := 0
	for _, ie := range g.IEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(g.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	g.Header.SetLength()
	return g.Header.MarshalTo(b)
}

// ParseGeneric parses a given byte sequence as a Generic.
func ParseGeneric(b []byte) (*Generic, error) {
	g := &Generic{}
	if err := g.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return g, nil
}

// UnmarshalBinary parses a given byte sequence as a Generic.
func (g *Generic) UnmarshalBinary(b []byte) error {
	var err error
	g.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(g.Header.Payload) < 2 {
		return nil
	}

	g.IEs, err = ie.ParseMultiIEs(g.Header.Payload)
	if err != nil {
		return err
	}
	return nil
}

// MarshalLen returns the serial length of Data.
func (g *Generic) MarshalLen() int {
	l := g.Header.MarshalLen() - len(g.Header.Payload)
	for _, ie := range g.IEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}

	return l
}

// SetLength sets the length in Length field.
func (g *Generic) SetLength() {
	g.Header.Length = uint16(g.MarshalLen() - 20)
}

// MessageTypeName returns the name of protocol.
func (g *Generic) MessageTypeName() string {
	return fmt.Sprintf("Unknown (%d)", g.Type)
}

// TID returns the TID in human-readable string.
func (g *Generic) TID() string {
	return g.tid()
}

// AddIE add IEs to Generic type of GTPv2 message and update Length field.
func (g *Generic) AddIE(ie ...*ie.IE) {
	g.IEs = append(g.IEs, ie...)
	g.SetLength()
}
