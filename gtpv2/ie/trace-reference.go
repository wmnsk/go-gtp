// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"

	"github.com/wmnsk/go-gtp/utils"
)

// NewTraceReference creates a new TraceReference IE.
func NewTraceReference(mcc, mnc string, traceID uint32) *IE {
	v := NewTraceReferenceFields(mcc, mnc, traceID)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(TraceReference, 0x00, b)
}

// TraceReferenceFields is a set of fields in TraceReference IE.
type TraceReferenceFields struct {
	MCC, MNC string
	TraceID  uint32 // 24-bit
}

// NewTraceReferenceFields creates a new TraceReferenceFields.
func NewTraceReferenceFields(mcc, mnc string, traceID uint32) *TraceReferenceFields {
	return &TraceReferenceFields{
		MCC:     mcc,
		MNC:     mnc,
		TraceID: traceID,
	}
}

// Marshal serializes TraceReferenceFields.
func (f *TraceReferenceFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes TraceReferenceFields.
func (f *TraceReferenceFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 3 {
		return io.ErrUnexpectedEOF
	}

	plmn, err := utils.EncodePLMN(f.MCC, f.MNC)
	if err != nil {
		return err
	}
	copy(b[0:3], plmn)

	if l < 6 {
		return io.ErrUnexpectedEOF
	}
	copy(b[3:6], utils.Uint32To24(f.TraceID))

	return nil
}

// ParseTraceReferenceFields decodes TraceReferenceFields.
func ParseTraceReferenceFields(b []byte) (*TraceReferenceFields, error) {
	f := &TraceReferenceFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into TraceReferenceFields.
func (f *TraceReferenceFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 3 {
		return io.ErrUnexpectedEOF
	}

	var err error
	f.MCC, f.MNC, err = utils.DecodePLMN(b[0:3])
	if err != nil {
		return err
	}

	if l < 6 {
		return io.ErrUnexpectedEOF
	}
	f.TraceID = utils.Uint24To32(b[3:6])

	return nil
}

// MarshalLen returns the serial length of TraceReferenceFields in int.
func (f *TraceReferenceFields) MarshalLen() int {
	return 6
}

// TraceID returns TraceID in uint32 if the type of IE matches.
func (i *IE) TraceID() (uint32, error) {
	switch i.Type {
	case TraceReference, TraceInformation:
		if len(i.Payload) < 6 {
			return 0, io.ErrUnexpectedEOF
		}
		return utils.Uint24To32(i.Payload[3:6]), nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustTraceID returns TraceID in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustTraceID() uint32 {
	v, _ := i.TraceID()
	return v
}
