// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewAggregateMaximumBitRate creates a new AggregateMaximumBitRate IE.
func NewAggregateMaximumBitRate(up, down uint32) *IE {
	// this is more efficient but removed for consistency with other structured IEs.
	// return newUint64ValIE(AggregateMaximumBitRate, (uint64(up)<<32 | uint64(down)))

	v := NewAggregateMaximumBitRateFields(up, down)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(AggregateMaximumBitRate, 0x00, b)
}

// AggregateMaximumBitRate returns AggregateMaximumBitRate in AggregateMaximumBitRateFields type if the type of IE matches.
func (i *IE) AggregateMaximumBitRate() (*AggregateMaximumBitRateFields, error) {
	switch i.Type {
	case AggregateMaximumBitRate:
		return ParseAggregateMaximumBitRateFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// AggregateMaximumBitRateFields is a set of fields in AggregateMaximumBitRate IE.
type AggregateMaximumBitRateFields struct {
	APNAMBRForUplink   uint32
	APNAMBRForDownlink uint32
}

// NewAggregateMaximumBitRateFields creates a new AggregateMaximumBitRateFields.
func NewAggregateMaximumBitRateFields(up, down uint32) *AggregateMaximumBitRateFields {
	return &AggregateMaximumBitRateFields{
		APNAMBRForUplink:   up,
		APNAMBRForDownlink: down,
	}
}

// Marshal serializes AggregateMaximumBitRateFields.
func (f *AggregateMaximumBitRateFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes AggregateMaximumBitRateFields.
func (f *AggregateMaximumBitRateFields) MarshalTo(b []byte) error {
	if len(b) < 8 {
		return io.ErrUnexpectedEOF
	}

	binary.BigEndian.PutUint32(b[0:4], f.APNAMBRForUplink)
	binary.BigEndian.PutUint32(b[4:8], f.APNAMBRForDownlink)

	return nil
}

// ParseAggregateMaximumBitRateFields decodes AggregateMaximumBitRateFields.
func ParseAggregateMaximumBitRateFields(b []byte) (*AggregateMaximumBitRateFields, error) {
	f := &AggregateMaximumBitRateFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into AggregateMaximumBitRateFields.
func (f *AggregateMaximumBitRateFields) UnmarshalBinary(b []byte) error {
	if len(b) < 8 {
		return io.ErrUnexpectedEOF
	}

	f.APNAMBRForUplink = binary.BigEndian.Uint32(b[0:4])
	f.APNAMBRForDownlink = binary.BigEndian.Uint32(b[4:8])

	return nil
}

// MarshalLen returns the serial length of AggregateMaximumBitRateFields in int.
func (f *AggregateMaximumBitRateFields) MarshalLen() int {
	return 8
}

// AggregateMaximumBitRateUp returns AggregateMaximumBitRate for Uplink
// if the type of IE matches.
func (i *IE) AggregateMaximumBitRateUp() (uint32, error) {
	if i.Type != AggregateMaximumBitRate {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint32(i.Payload[0:4]), nil
}

// MustAggregateMaximumBitRateUp returns AggregateMaximumBitRateUp in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustAggregateMaximumBitRateUp() uint32 {
	v, _ := i.AggregateMaximumBitRateUp()
	return v
}

// AggregateMaximumBitRateDown returns AggregateMaximumBitRate for Downlink
// if the type of IE matches.
func (i *IE) AggregateMaximumBitRateDown() (uint32, error) {
	if i.Type != AggregateMaximumBitRate {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 8 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint32(i.Payload[4:8]), nil
}

// MustAggregateMaximumBitRateDown returns AggregateMaximumBitRateDown in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustAggregateMaximumBitRateDown() uint32 {
	v, _ := i.AggregateMaximumBitRateDown()
	return v
}
