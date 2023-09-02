// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
	"net"
)

// NewFullyQualifiedTEID creates a new FullyQualifiedTEID IE.
func NewFullyQualifiedTEID(ifType uint8, teid uint32, v4, v6 string) *IE {
	v := NewFullyQualifiedTEIDFields(ifType, teid, net.ParseIP(v4), net.ParseIP(v6))
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(FullyQualifiedTEID, 0x00, b)
}

// NewFullyQualifiedTEIDNetIP creates a new FullyQualifiedTEID IE from net.IP instead of string.
func NewFullyQualifiedTEIDNetIP(ifType uint8, teid uint32, v4, v6 net.IP) *IE {
	v := NewFullyQualifiedTEIDFields(ifType, teid, v4, v6)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(FullyQualifiedTEID, 0x00, b)
}

// FullyQualifiedTEID returns FullyQualifiedTEID in FullyQualifiedTEIDFields type if the type of IE matches.
func (i *IE) FullyQualifiedTEID() (*FullyQualifiedTEIDFields, error) {
	switch i.Type {
	case FullyQualifiedTEID:
		return ParseFullyQualifiedTEIDFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// FullyQualifiedTEIDFields is a set of fields in FullyQualifiedTEID IE.
type FullyQualifiedTEIDFields struct {
	Flags         uint8 // 7-8th bit, in the same octet as InterfaceType
	InterfaceType uint8 // 1-6-bit, in the same octet as Flags
	TEIDGREKey    uint32
	IPv4Address   net.IP
	IPv6Address   net.IP
}

// NewFullyQualifiedTEIDFields creates a new FullyQualifiedTEIDFields.
func NewFullyQualifiedTEIDFields(ifType uint8, teid uint32, v4, v6 net.IP) *FullyQualifiedTEIDFields {
	f := &FullyQualifiedTEIDFields{
		InterfaceType: ifType,
		TEIDGREKey:    teid,
	}

	if v := v4.To4(); v != nil {
		f.Flags |= 0x80
		f.IPv4Address = v
	}
	if v := v6.To16(); v != nil {
		f.Flags |= 0x40
		f.IPv6Address = v
	}

	return f
}

// Marshal serializes FullyQualifiedTEIDFields.
func (f *FullyQualifiedTEIDFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes FullyQualifiedTEIDFields.
func (f *FullyQualifiedTEIDFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 5 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.InterfaceType
	binary.BigEndian.PutUint32(b[1:5], f.TEIDGREKey)
	offset := 5

	if f.IPv4Address != nil {
		if l < offset+4 {
			return io.ErrUnexpectedEOF
		}
		if v := f.IPv4Address.To4(); v != nil {
			b[0] |= 0x80
			copy(b[offset:offset+4], v)
			offset += 4
		}
	}
	if f.IPv6Address != nil {
		if l < offset+16 {
			return io.ErrUnexpectedEOF
		}
		if v := f.IPv6Address.To16(); v != nil {
			b[0] |= 0x40
			copy(b[offset:offset+16], v)
		}
	}

	return nil
}

// ParseFullyQualifiedTEIDFields decodes FullyQualifiedTEIDFields.
func ParseFullyQualifiedTEIDFields(b []byte) (*FullyQualifiedTEIDFields, error) {
	f := &FullyQualifiedTEIDFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into FullyQualifiedTEIDFields.
func (f *FullyQualifiedTEIDFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 5 {
		return io.ErrUnexpectedEOF
	}

	f.Flags = b[0] & 0xc0
	f.InterfaceType = b[0] & 0x3f
	f.TEIDGREKey = binary.BigEndian.Uint32(b[1:5])
	offset := 5

	if has8thBit(f.Flags) { // has IPv4Address
		if l < offset+4 {
			return io.ErrUnexpectedEOF
		}
		if v := net.IP(b[offset : offset+4]).To4(); v != nil {
			f.IPv4Address = v
			offset += 4
		}
	}
	if has7thBit(f.Flags) { // has IPv6Address
		if l < offset+16 {
			return io.ErrUnexpectedEOF
		}
		if v := net.IP(b[offset : offset+16]).To16(); v != nil {
			f.IPv6Address = v
		}
	}

	return nil
}

// MarshalLen returns the serial length of FullyQualifiedTEIDFields in int.
func (f *FullyQualifiedTEIDFields) MarshalLen() int {
	l := 5
	if has8thBit(f.Flags) { // has IPv4Address
		l += 4
	}
	if has7thBit(f.Flags) { // has IPv6Address
		l += 16
	}

	return l
}

// HasIPv4 reports whether an IE has IPv4 bit.
func (i *IE) HasIPv4() bool {
	switch i.Type {
	case FullyQualifiedTEID:
		if len(i.Payload) < 1 {
			return false
		}

		return has8thBit(i.Payload[0])
	default:
		return false
	}
}

// HasIPv6 reports whether an IE has IPv6 bit.
func (i *IE) HasIPv6() bool {
	switch i.Type {
	case FullyQualifiedTEID:
		if len(i.Payload) < 1 {
			return false
		}

		return has7thBit(i.Payload[0])
	default:
		return false
	}
}

// InterfaceType returns InterfaceType in uint8 if the type of IE matches.
func (i *IE) InterfaceType() (uint8, error) {
	if i.Type != FullyQualifiedTEID {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0] & 0x3f, nil
}

// MustInterfaceType returns InterfaceType in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustInterfaceType() uint8 {
	v, _ := i.InterfaceType()
	return v
}

// GREKey returns GREKey in uint32 if the type of IE matches.
func (i *IE) GREKey() (uint32, error) {
	if len(i.Payload) < 6 {
		return 0, io.ErrUnexpectedEOF
	}
	switch i.Type {
	case FullyQualifiedTEID:
		return binary.BigEndian.Uint32(i.Payload[1:5]), nil
	case S103PDNDataForwardingInfo:
		switch i.Payload[0] {
		case 4:
			if len(i.Payload) < 9 {
				return 0, io.ErrUnexpectedEOF
			}
			return binary.BigEndian.Uint32(i.Payload[5:9]), nil
		case 16:
			if len(i.Payload) < 21 {
				return 0, io.ErrUnexpectedEOF
			}
			return binary.BigEndian.Uint32(i.Payload[17:21]), nil
		default:
			return 0, ErrMalformed
		}
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustGREKey returns GREKey in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustGREKey() uint32 {
	v, _ := i.GREKey()
	return v
}

// TEID returns TEID in uint32 if the type of IE matches.
func (i *IE) TEID() (uint32, error) {
	if len(i.Payload) < 5 {
		return 0, io.ErrUnexpectedEOF
	}
	switch i.Type {
	case FullyQualifiedTEID:
		return binary.BigEndian.Uint32(i.Payload[1:5]), nil
	case S1UDataForwarding:
		switch i.Payload[0] {
		case 4:
			if len(i.Payload) < 9 {
				return 0, io.ErrUnexpectedEOF
			}
			return binary.BigEndian.Uint32(i.Payload[5:9]), nil
		case 16:
			if len(i.Payload) < 21 {
				return 0, io.ErrUnexpectedEOF
			}
			return binary.BigEndian.Uint32(i.Payload[17:21]), nil
		default:
			return 0, ErrMalformed
		}
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}

}

// MustTEID returns TEID in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustTEID() uint32 {
	v, _ := i.TEID()
	return v
}
