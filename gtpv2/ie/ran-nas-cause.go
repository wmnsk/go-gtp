// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"fmt"
	"io"
)

// NewRANNASCause creates a new RANNASCause IE.
func NewRANNASCause(pType, cType uint8, cause []byte) *IE {
	v := NewRANNASCauseFields(pType, cType, cause)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(RANNASCause, 0x00, b)
}

// RANNASCause returns RANNASCause in RANNASCauseFields type if the type of IE matches.
func (i *IE) RANNASCause() (*RANNASCauseFields, error) {
	switch i.Type {
	case RANNASCause:
		return ParseRANNASCauseFields(i.Payload)
	case BearerContext:
		ies, err := i.BearerContext()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve RANNASCause: %w", err)
		}

		for _, child := range ies {
			if child.Type == RANNASCause {
				return child.RANNASCause()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// RANNASCauseFields is a set of fields in RANNASCause IE.
type RANNASCauseFields struct {
	ProtocolType uint8  // 4-bit
	CauseType    uint8  // 4-bit
	Cause        []byte // format depends on ProtocolType
}

// NewRANNASCauseFields creates a new RANNASCauseFields.
func NewRANNASCauseFields(pType, cType uint8, cause []byte) *RANNASCauseFields {
	return &RANNASCauseFields{
		ProtocolType: pType,
		CauseType:    cType,
		Cause:        cause,
	}
}

// Marshal serializes RANNASCauseFields.
func (f *RANNASCauseFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes RANNASCauseFields.
func (f *RANNASCauseFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	b[0] = ((f.ProtocolType & 0x0f) << 4) | (f.CauseType & 0x0f)

	if l < 1+len(f.Cause) {
		return io.ErrUnexpectedEOF
	}
	copy(b[1:], f.Cause)

	return nil
}

// ParseRANNASCauseFields decodes RANNASCauseFields.
func ParseRANNASCauseFields(b []byte) (*RANNASCauseFields, error) {
	f := &RANNASCauseFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into RANNASCauseFields.
func (f *RANNASCauseFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}

	f.ProtocolType = (b[0] & 0xf0) >> 4
	f.CauseType = b[0] & 0x0f
	f.Cause = b[1:]

	return nil
}

// MarshalLen returns the serial length of RANNASCauseFields in int.
func (f *RANNASCauseFields) MarshalLen() int {
	return 1 + len(f.Cause)
}
