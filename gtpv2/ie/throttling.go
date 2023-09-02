// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
	"time"
)

// NewThrottling creates a new Timer IE.
func NewThrottling(delay time.Duration, factor uint8) *IE {
	v := NewThrottlingFields(delay, factor)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(Throttling, 0x00, b)
}

// Throttling returns Throttling in time.Duration if the type of IE matches.
func (i *IE) Throttling() (*ThrottlingFields, error) {
	switch i.Type {
	case Throttling:
		return ParseThrottlingFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// ThrottlingFields is a set of fields in Throttling IE.
type ThrottlingFields struct {
	DelayUnit  uint8
	DelayValue time.Duration
	Factor     uint8
}

// NewThrottlingFields creates a new ThrottlingFields.
func NewThrottlingFields(delay time.Duration, factor uint8) *ThrottlingFields {
	// 8.85 Throttling
	// Bits 5 to 1 represent the binary coded timer value.
	// Bits 6 to 8 defines the timer value unit as follows:
	// Bits
	// 8 7 6
	// 0 0 0 value is incremented in multiples of 2 seconds
	// 0 0 1 value is incremented in multiples of 1 minute
	// 0 1 0 value is incremented in multiples of 10 minutes
	// 0 1 1 value is incremented in multiples of 1 hour
	// 1 0 0 value is incremented in multiples of 10 hours
	// 1 1 1 value indicates that the timer is infinite
	//
	// Other values shall be interpreted as multiples of 1 minute in this version of the protocol.
	// Timer unit and Timer value both set to all "zeros" shall be interpreted as an indication that the timer is stopped.

	var unit uint8
	switch {
	case delay%(10*time.Hour) == 0:
		unit = 0x04
	case delay%(1*time.Hour) == 0:
		unit = 0x03
	case delay%(10*time.Minute) == 0:
		unit = 0x02
	case delay%(1*time.Minute) == 0:
		unit = 0x01
	case delay%(2*time.Second) == 0:
		unit = 0x00
	default:
		unit = 0xe0
	}

	return &ThrottlingFields{
		DelayUnit:  unit,
		DelayValue: delay,
		Factor:     factor,
	}
}

// Marshal serializes ThrottlingFields.
func (f *ThrottlingFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes ThrottlingFields.
func (f *ThrottlingFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}

	var value uint8
	switch f.DelayUnit {
	case 0x04:
		value = uint8(f.DelayValue / (10 * time.Hour))
	case 0x03:
		value = uint8(f.DelayValue / time.Hour)
	case 0x02:
		value = uint8(f.DelayValue / (10 * time.Minute))
	case 0x01:
		value = uint8(f.DelayValue / time.Minute)
	case 0x00:
		value = uint8(f.DelayValue / (2 * time.Second))
	default:
		value = 0
	}

	b[0] = ((f.DelayUnit << 5) & 0xe0) | (value & 0x1f)
	b[1] = f.Factor

	return nil
}

// ParseThrottlingFields decodes ThrottlingFields.
func ParseThrottlingFields(b []byte) (*ThrottlingFields, error) {
	f := &ThrottlingFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into ThrottlingFields.
func (f *ThrottlingFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}

	f.DelayUnit = (b[0] & 0xe0) >> 5

	switch f.DelayUnit {
	case 0x04:
		f.DelayValue = time.Duration(b[0]&0x1f) * (10 * time.Hour)
	case 0x03:
		f.DelayValue = time.Duration(b[0]&0x1f) * time.Hour
	case 0x02:
		f.DelayValue = time.Duration(b[0]&0x1f) * (10 * time.Minute)
	case 0x01:
		f.DelayValue = time.Duration(b[0]&0x1f) * time.Minute
	case 0x00:
		f.DelayValue = time.Duration(b[0]&0x1f) * (2 * time.Second)
	default:
		f.DelayValue = 0
	}

	f.Factor = b[1]

	return nil
}

// MarshalLen returns the serial length of ThrottlingFields in int.
func (f *ThrottlingFields) MarshalLen() int {
	return 2
}
