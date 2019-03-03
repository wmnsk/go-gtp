// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"fmt"

	"github.com/wmnsk/go-gtp/v2/ies"
)

// Generic is a Generic formed GTP Header and its IEs above.
type Generic struct {
	*Header
	IEs []*ies.IE
}

// NewGeneric creates a new Generic.
func NewGeneric(msgType uint8, teid, seq uint32, ie ...*ies.IE) *Generic {
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
func NewGenericWithoutTEID(msgType uint8, teid, seq uint32, ie ...*ies.IE) *Generic {
	g := &Generic{
		Header: NewHeader(
			NewHeaderFlags(2, 0x00, 0),
			msgType, teid, seq, nil,
		),
		IEs: ie,
	}

	return g
}

// Serialize returns the byte sequence generated from a Generic.
func (g *Generic) Serialize() ([]byte, error) {
	b := make([]byte, g.Len())
	if err := g.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo puts the byte sequence in the byte array given as b.
func (g *Generic) SerializeTo(b []byte) error {
	if g.Header.Payload != nil {
		g.Header.Payload = nil
	}
	g.Header.Payload = make([]byte, g.Len()-g.Header.Len())

	offset := 0
	for _, ie := range g.IEs {
		if ie == nil {
			continue
		}
		if err := ie.SerializeTo(g.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	g.Header.SetLength()
	return g.Header.SerializeTo(b)
}

// DecodeGeneric decodes a given byte sequence as a Generic.
func DecodeGeneric(b []byte) (*Generic, error) {
	g := &Generic{}
	if err := g.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return g, nil
}

// DecodeFromBytes decodes a given byte sequence as a Generic.
func (g *Generic) DecodeFromBytes(b []byte) error {
	var err error
	g.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}
	if len(g.Header.Payload) < 2 {
		return nil
	}

	g.IEs, err = ies.DecodeMultiIEs(g.Header.Payload)
	if err != nil {
		return err
	}
	return nil
}

// Len returns the actual length of Data.
func (g *Generic) Len() int {
	l := g.Header.Len() - len(g.Header.Payload)

	for _, ie := range g.IEs {
		if ie == nil {
			continue
		}
		l += ie.Len()
	}

	return l
}

// SetLength sets the length in Length field.
func (g *Generic) SetLength() {
	l := g.Header.Len() - len(g.Header.Payload) - 4
	for _, ie := range g.IEs {
		l += ie.Len()
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
func (g *Generic) AddIE(ie ...*ies.IE) {
	g.IEs = append(g.IEs, ie...)
	g.SetLength()
}
