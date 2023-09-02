// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
	"net"
)

// NewS1UDataForwarding creates a new S1UDataForwarding IE.
func NewS1UDataForwarding(ebi uint8, sgwAddr string, sgwTEID uint32) *IE {
	v := NewS1UDataForwardingFields(ebi, net.ParseIP(sgwAddr), sgwTEID)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(S1UDataForwarding, 0x00, b)
}

// NewS1UDataForwardingNetIP creates a new S1UDataForwarding IE.
func NewS1UDataForwardingNetIP(ebi uint8, sgwIP net.IP, sgwTEID uint32) *IE {
	v := NewS1UDataForwardingFields(ebi, sgwIP, sgwTEID)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(S1UDataForwarding, 0x00, b)
}

// S1UDataForwarding returns S1UDataForwarding in S1UDataForwardingFields type if the type of IE matches.
func (i *IE) S1UDataForwarding() (*S1UDataForwardingFields, error) {
	switch i.Type {
	case S1UDataForwarding:
		return ParseS1UDataForwardingFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// S1UDataForwardingFields is a set of fields in S1UDataForwarding IE.
type S1UDataForwardingFields struct {
	EPSBearerID            uint8 // 4-bit
	ServingGWAddressLength uint8
	ServingGWAddress       net.IP
	ServingGWS1UTEID       uint32
}

// NewS1UDataForwardingFields creates a new S1UDataForwardingFields.
func NewS1UDataForwardingFields(ebi uint8, sgwAddr net.IP, sgwTEID uint32) *S1UDataForwardingFields {
	f := &S1UDataForwardingFields{
		EPSBearerID:      ebi,
		ServingGWS1UTEID: sgwTEID,
	}

	if v := sgwAddr.To4(); v != nil {
		f.ServingGWAddressLength = 4
		f.ServingGWAddress = v
		return f
	}

	if v := sgwAddr.To16(); v != nil {
		f.ServingGWAddressLength = 16
		f.ServingGWAddress = v
		return f
	}

	// return IE w/ "something" anyway
	f.ServingGWAddressLength = uint8(len(sgwAddr))
	f.ServingGWAddress = sgwAddr
	return f
}

// Marshal serializes S1UDataForwardingFields.
func (f *S1UDataForwardingFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes S1UDataForwardingFields.
func (f *S1UDataForwardingFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.EPSBearerID & 0x0f
	b[1] = f.ServingGWAddressLength
	offset := 2

	if l < offset+int(f.ServingGWAddressLength) {
		return io.ErrUnexpectedEOF
	}
	copy(b[offset:offset+int(f.ServingGWAddressLength)], f.ServingGWAddress)
	offset += int(f.ServingGWAddressLength)

	if l < offset+4 {
		return io.ErrUnexpectedEOF
	}
	binary.BigEndian.PutUint32(b[offset:offset+4], f.ServingGWS1UTEID)

	return nil
}

// ParseS1UDataForwardingFields decodes S1UDataForwardingFields.
func ParseS1UDataForwardingFields(b []byte) (*S1UDataForwardingFields, error) {
	f := &S1UDataForwardingFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into S1UDataForwardingFields.
func (f *S1UDataForwardingFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}

	f.EPSBearerID = b[0] & 0x0f
	f.ServingGWAddressLength = b[1]
	offset := 2

	if l < offset+int(f.ServingGWAddressLength) {
		return io.ErrUnexpectedEOF
	}
	f.ServingGWAddress = net.IP(b[offset : offset+int(f.ServingGWAddressLength)])
	offset += int(f.ServingGWAddressLength)

	if l < offset+4 {
		return io.ErrUnexpectedEOF
	}
	f.ServingGWS1UTEID = binary.BigEndian.Uint32(b[offset : offset+4])

	return nil
}

// MarshalLen returns the serial length of S1UDataForwardingFields in int.
func (f *S1UDataForwardingFields) MarshalLen() int {
	return 2 + int(f.ServingGWAddressLength) + 4
}

// SGWAddress returns IP address of SGW in string if the type of IE matches.
func (i *IE) SGWAddress() (string, error) {
	if i.Type != S1UDataForwarding {
		return "", &InvalidTypeError{Type: i.Type}
	}

	return i.IPAddress()
}

// MustSGWAddress returns SGWAddress in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustSGWAddress() string {
	v, _ := i.SGWAddress()
	return v
}
