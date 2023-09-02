// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
	"math"
	"time"
)

// NewUETimeZone creates a new UETimeZone IE.
func NewUETimeZone(tz time.Duration, daylightSaving uint8) *IE {
	i := New(UETimeZone, 0x00, make([]byte, 2))
	min := tz.Minutes() / 15
	absMin := int(math.Abs(min))
	hex := byte(((absMin % 10) << 4) | (absMin / 10))
	if min < 0 {
		hex |= 0x08
	}
	i.Payload[0] = hex
	i.Payload[1] = daylightSaving & 0x03

	return i
}

// TimeZone returns TimeZone in time.Duration if the type of IE matches.
func (i *IE) TimeZone() (time.Duration, error) {
	if i.Type != UETimeZone {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	unsigned := i.Payload[0] & 0xf7
	dec := int((unsigned >> 4) + (unsigned&0x0f)*10)
	if (i.Payload[0]&0x08)>>3 == 1 {
		dec *= -1
	}

	return time.Duration(dec*15) * time.Minute, nil
}

// MustTimeZone returns TimeZone in time.Duration, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustTimeZone() time.Duration {
	v, _ := i.TimeZone()
	return v
}

// DaylightSaving returns DaylightSaving in uint8 if the type of IE matches.
func (i *IE) DaylightSaving() (uint8, error) {
	if i.Type != UETimeZone {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[1], nil
}

// MustDaylightSaving returns DaylightSaving in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustDaylightSaving() uint8 {
	v, _ := i.DaylightSaving()
	return v
}
