// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"fmt"

	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// Generic is a Generic formed GTP Header and its IEs above.
type Generic struct {
	*Header
	IEs []*ie.IE
}

// NewGeneric creates a new Generic.
func NewGeneric(msgType uint8, teid, seq uint32, ie ...*ie.IE) *Generic {
	g := &Generic{
		Header: NewHeader(
			NewHeaderFlags(2, 0x00, 1),
			msgType,
			teid, seq,
			nil,
		),
		IEs: ie,
	}
	g.SetLength()

	return g
}

// NewGenericWithoutTEID creates a new Generic.
func NewGenericWithoutTEID(msgType uint8, teid, seq uint32, ie ...*ie.IE) *Generic {
	g := &Generic{
		Header: NewHeader(
			NewHeaderFlags(2, 0x00, 0),
			msgType, teid, seq, nil,
		),
		IEs: ie,
	}

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

// ParseGeneric decodes a given byte sequence as a Generic.
func ParseGeneric(b []byte) (*Generic, error) {
	g := &Generic{}
	if err := g.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return g, nil
}

// UnmarshalBinary decodes a given byte sequence as a Generic.
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
	l := g.Header.MarshalLen() - len(g.Header.Payload) - 4
	for _, ie := range g.IEs {
		l += ie.MarshalLen()
	}
	g.Header.Length = uint16(l)
}

// MessageTypeName returns the name of protocol.
func (g *Generic) MessageTypeName() string {
	return fmt.Sprintf("Unknown (%d)", g.Header.Type)
}

// TEID returns the TEID in uint32.
func (g *Generic) TEID() uint32 {
	return g.Header.teid()
}

// AddIE add IEs to Generic type of GTPv2 message and update Length field.
func (g *Generic) AddIE(ie ...*ie.IE) {
	g.IEs = append(g.IEs, ie...)
	g.SetLength()
}
