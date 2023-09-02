// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"encoding/hex"
	"io"
	"net"
)

// Node-ID Type definitions.
const (
	nodeIDIPv4 uint8 = iota
	nodeIDIPv6
	nodeIDOther
)

// NewFullyQualifiedCSID creates a new FullyQualifiedCSID IE.
func NewFullyQualifiedCSID(nodeID string, csIDs ...uint16) *IE {
	v := NewFullyQualifiedCSIDFields(nodeID, csIDs...)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(FullyQualifiedCSID, 0x00, b)
}

// FullyQualifiedCSID returns FullyQualifiedCSID in FullyQualifiedCSIDFields type if the type of IE matches.
func (i *IE) FullyQualifiedCSID() (*FullyQualifiedCSIDFields, error) {
	switch i.Type {
	case FullyQualifiedCSID:
		return ParseFullyQualifiedCSIDFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// FullyQualifiedCSIDFields is a set of fields in FullyQualifiedCSID IE.
type FullyQualifiedCSIDFields struct {
	NodeIDType    uint8 // 4-bit
	NumberOfCSIDs uint8 // 4-bit
	NodeID        []byte
	CSIDs         []uint16
}

// NewFullyQualifiedCSIDFields creates a new FullyQualifiedCSIDFields.
func NewFullyQualifiedCSIDFields(nodeID string, csIDs ...uint16) *FullyQualifiedCSIDFields {
	f := &FullyQualifiedCSIDFields{
		NumberOfCSIDs: uint8(len(csIDs)),
		CSIDs:         csIDs,
	}

	ip := net.ParseIP(nodeID)
	if ip == nil {
		var err error
		f.NodeID, err = hex.DecodeString(nodeID)
		if err != nil {
			return nil
		}
		f.NodeIDType = nodeIDOther
	} else if v4 := ip.To4(); v4 != nil {
		f.NodeID = v4
		f.NodeIDType = nodeIDIPv4
	} else {
		f.NodeID = ip
		f.NodeIDType = nodeIDIPv6
	}

	return f
}

// Marshal serializes FullyQualifiedCSIDFields.
func (f *FullyQualifiedCSIDFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes FullyQualifiedCSIDFields.
func (f *FullyQualifiedCSIDFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	b[0] = ((f.NodeIDType << 4) & 0xf0) | (f.NumberOfCSIDs & 0x0f)
	offset := 1

	if l < offset+len(f.NodeID) {
		return io.ErrUnexpectedEOF
	}
	copy(b[offset:offset+len(f.NodeID)], f.NodeID)
	offset += len(f.NodeID)

	for n, csid := range f.CSIDs {
		if l < offset+n*2+2 {
			break
		}
		binary.BigEndian.PutUint16(b[offset+n*2:offset+n*2+2], csid)
	}

	return nil
}

// ParseFullyQualifiedCSIDFields decodes FullyQualifiedCSIDFields.
func ParseFullyQualifiedCSIDFields(b []byte) (*FullyQualifiedCSIDFields, error) {
	f := &FullyQualifiedCSIDFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into FullyQualifiedCSIDFields.
func (f *FullyQualifiedCSIDFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	f.NodeIDType = (b[0] >> 4) & 0x0f
	f.NumberOfCSIDs = b[0] & 0x0f
	offset := 1

	switch f.NodeIDType {
	case nodeIDIPv4, nodeIDOther:
		if l < offset+4 {
			return io.ErrUnexpectedEOF
		}
		f.NodeID = b[offset : offset+4]
		offset += 4
	case nodeIDIPv6:
		if l < offset+16 {
			return io.ErrUnexpectedEOF
		}
		f.NodeID = b[offset : offset+16]
		offset += 16
	default:
		return ErrMalformed
	}

	for {
		if l < offset+2 {
			break
		}
		f.CSIDs = append(f.CSIDs, binary.BigEndian.Uint16(b[offset:offset+2]))
		offset += 2
	}

	return nil
}

// MarshalLen returns the serial length of FullyQualifiedCSIDFields in int.
func (f *FullyQualifiedCSIDFields) MarshalLen() int {
	l := 1

	switch f.NodeIDType {
	case nodeIDIPv4, nodeIDOther:
		l += 4
	case nodeIDIPv6:
		l += 16
	default:

	}

	l += len(f.CSIDs) * 2
	return l
}

// NodeIDType returns NodeIDType in uint8 if the type of IE matches.
func (i *IE) NodeIDType() (uint8, error) {
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case FullyQualifiedCSID:
		return (i.Payload[0] >> 4) & 0x0f, nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustNodeIDType returns NodeIDType in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustNodeIDType() uint8 {
	v, _ := i.NodeIDType()
	return v
}

// NodeID returns NodeID in []byte if the type of IE matches.
func (i *IE) NodeID() ([]byte, error) {
	if len(i.Payload) < 1 {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case FullyQualifiedCSID:
		switch (i.Payload[0] >> 4) & 0x0f {
		case nodeIDIPv4, nodeIDOther:
			if len(i.Payload) < 6 {
				return nil, io.ErrUnexpectedEOF
			}
			return i.Payload[1:5], nil
		case nodeIDIPv6:
			if len(i.Payload) < 18 {
				return nil, io.ErrUnexpectedEOF
			}
			return i.Payload[1:17], nil
		default:
			return nil, ErrMalformed
		}
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MustNodeID returns NodeID in []byte, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustNodeID() []byte {
	v, _ := i.NodeID()
	return v
}

// CSIDs returns CSIDs in []uint16 if the type of IE matches.
func (i *IE) CSIDs() ([]uint16, error) {
	if len(i.Payload) < 1 {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case FullyQualifiedCSID:
		offset := 0
		switch (i.Payload[0] >> 4) & 0x0f {
		case nodeIDIPv4, nodeIDOther:
			offset += 5
		case nodeIDIPv6:
			offset += 17
		default:
			return nil, ErrMalformed
		}

		var csids []uint16
		for {
			if offset+2 > len(i.Payload) {
				break
			}
			csids = append(csids, binary.BigEndian.Uint16(i.Payload[offset:offset+2]))
			offset += 2
		}
		return csids, nil
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MustCSIDs returns CSIDs in []uint16, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustCSIDs() []uint16 {
	v, _ := i.CSIDs()
	return v
}
