// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
	"net"
)

// NewS103PDNDataForwardingInfo creates a new S103PDNDataForwardingInfo IE.
func NewS103PDNDataForwardingInfo(hsgwAddr string, greKey uint32, ebis ...uint8) *IE {
	v := NewS103PDNDataForwardingInfoFields(net.ParseIP(hsgwAddr), greKey, ebis...)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(S103PDNDataForwardingInfo, 0x00, b)
}

// NewS103PDNDataForwardingInfoNetIP creates a new S103PDNDataForwardingInfo IE.
func NewS103PDNDataForwardingInfoNetIP(hsgwIP net.IP, greKey uint32, ebis ...uint8) *IE {
	v := NewS103PDNDataForwardingInfoFields(hsgwIP, greKey, ebis...)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(S103PDNDataForwardingInfo, 0x00, b)
}

// S103PDNDataForwardingInfo returns S103PDNDataForwardingInfo in S103PDNDataForwardingInfoFields type if the type of IE matches.
func (i *IE) S103PDNDataForwardingInfo() (*S103PDNDataForwardingInfoFields, error) {
	switch i.Type {
	case S103PDNDataForwardingInfo:
		return ParseS103PDNDataForwardingInfoFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// S103PDNDataForwardingInfoFields is a set of fields in S103PDNDataForwardingInfo IE.
type S103PDNDataForwardingInfoFields struct {
	HSGWAddressForForwardingLength uint8
	HSGWAddressForForwarding       net.IP
	GREKey                         uint32
	EPSBearerIDNumber              uint8
	EPSBearerIDs                   []uint8
}

// NewS103PDNDataForwardingInfoFields creates a new S103PDNDataForwardingInfoFields.
func NewS103PDNDataForwardingInfoFields(hsgwIP net.IP, greKey uint32, ebis ...uint8) *S103PDNDataForwardingInfoFields {
	f := &S103PDNDataForwardingInfoFields{
		GREKey:            greKey,
		EPSBearerIDNumber: uint8(len(ebis)),
		EPSBearerIDs:      ebis,
	}

	if v := hsgwIP.To4(); v != nil {
		f.HSGWAddressForForwardingLength = 4
		f.HSGWAddressForForwarding = v
		return f
	}

	if v := hsgwIP.To16(); v != nil {
		f.HSGWAddressForForwardingLength = 16
		f.HSGWAddressForForwarding = v
		return f
	}

	// return IE w/ "something" anyway
	f.HSGWAddressForForwardingLength = uint8(len(hsgwIP))
	f.HSGWAddressForForwarding = hsgwIP
	return f
}

// Marshal serializes S103PDNDataForwardingInfoFields.
func (f *S103PDNDataForwardingInfoFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes S103PDNDataForwardingInfoFields.
func (f *S103PDNDataForwardingInfoFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.HSGWAddressForForwardingLength
	offset := 1

	if l < offset+int(f.HSGWAddressForForwardingLength) {
		return io.ErrUnexpectedEOF
	}
	copy(b[offset:offset+int(f.HSGWAddressForForwardingLength)], f.HSGWAddressForForwarding)
	offset += int(f.HSGWAddressForForwardingLength)

	if l < offset+4 {
		return io.ErrUnexpectedEOF
	}
	binary.BigEndian.PutUint32(b[offset:offset+4], f.GREKey)
	offset += 4

	b[offset] = f.EPSBearerIDNumber
	offset++

	if l < offset+int(f.EPSBearerIDNumber) {
		return io.ErrUnexpectedEOF
	}
	for _, ebi := range f.EPSBearerIDs {
		b[offset] = ebi & 0x0f
		offset++
	}

	return nil
}

// ParseS103PDNDataForwardingInfoFields decodes S103PDNDataForwardingInfoFields.
func ParseS103PDNDataForwardingInfoFields(b []byte) (*S103PDNDataForwardingInfoFields, error) {
	f := &S103PDNDataForwardingInfoFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into S103PDNDataForwardingInfoFields.
func (f *S103PDNDataForwardingInfoFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}

	f.HSGWAddressForForwardingLength = b[0]
	offset := 1

	if l < offset+int(f.HSGWAddressForForwardingLength) {
		return io.ErrUnexpectedEOF
	}
	f.HSGWAddressForForwarding = net.IP(b[offset : offset+int(f.HSGWAddressForForwardingLength)])
	offset += int(f.HSGWAddressForForwardingLength)

	if l < offset+4 {
		return io.ErrUnexpectedEOF
	}
	f.GREKey = binary.BigEndian.Uint32(b[offset : offset+4])
	offset += 4

	if l <= offset {
		return io.ErrUnexpectedEOF
	}

	f.EPSBearerIDNumber = b[offset]
	offset++

	if l < offset+int(f.EPSBearerIDNumber) {
		return io.ErrUnexpectedEOF
	}

	f.EPSBearerIDs = make([]uint8, f.EPSBearerIDNumber)
	for n := 0; n < int(f.EPSBearerIDNumber); n++ {
		f.EPSBearerIDs[n] = b[offset]
		offset++
	}

	return nil
}

// MarshalLen returns the serial length of S103PDNDataForwardingInfoFields in int.
func (f *S103PDNDataForwardingInfoFields) MarshalLen() int {
	return 1 + int(f.HSGWAddressForForwardingLength) + 4 + 1 + int(f.EPSBearerIDNumber)
}

// HSGWAddress returns IP address of HSGW in string if the type of IE matches.
func (i *IE) HSGWAddress() (string, error) {
	if i.Type != S103PDNDataForwardingInfo {
		return "", &InvalidTypeError{Type: i.Type}
	}

	return i.IPAddress()
}

// MustHSGWAddress returns HSGWAddress in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustHSGWAddress() string {
	v, _ := i.HSGWAddress()
	return v
}

// EBIs returns the EBIs in []uint8 if the type of IE matches.
func (i *IE) EBIs() ([]uint8, error) {
	if i.Type != S103PDNDataForwardingInfo {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return nil, io.ErrUnexpectedEOF
	}

	var n, offset int
	switch i.Payload[0] {
	case 4:
		if len(i.Payload) <= 9 {
			return nil, io.ErrUnexpectedEOF
		}
		n = int(i.Payload[9])
		offset = 10
	case 16:
		if len(i.Payload) <= 21 {
			return nil, io.ErrUnexpectedEOF
		}
		n = int(i.Payload[21])
		offset = 22
	default:
		return nil, ErrMalformed
	}

	var ebis []uint8
	for x := 0; x < n; x++ {
		if len(i.Payload) <= offset+x {
			break
		}
		ebis = append(ebis, i.Payload[offset+x])
	}
	return ebis, nil
}

// MustEBIs returns EBIs in []uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustEBIs() []uint8 {
	v, _ := i.EBIs()
	return v
}
