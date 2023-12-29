// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
	"time"
)

// NewDelayValue creates a new DelayValue IE.
func NewDelayValue(delay time.Duration) *IE {
	return NewUint8IE(DelayValue, uint8(delay.Seconds()*1000/50))
}

// NewDelayValueRaw creates a new DelayValue IE from a uint8 value.
//
// The value should be in multiples of 50ms or zero.
func NewDelayValueRaw(delay uint8) *IE {
	return NewUint8IE(DelayValue, delay)
}

// DelayValue returns DelayValue in time.Duration if the type of IE matches.
//
// The returned value is in time.Duration. To get the value in multiples of 50ms,
// use ValueAsUint8 or access Payload field directly instead.
func (i *IE) DelayValue() (time.Duration, error) {
	if i.Type != DelayValue {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return time.Duration(i.Payload[0]/50) * time.Millisecond, nil
}

// MustDelayValue returns DelayValue in time.Duration, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustDelayValue() time.Duration {
	v, _ := i.DelayValue()
	return v
}
