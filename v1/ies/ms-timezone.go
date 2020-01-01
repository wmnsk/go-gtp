// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"io"
	"math"
	"strconv"
	"time"
)

// Timezone adjustment definitions.
const (
	TimeAdjustNoDaylightSaving uint8 = iota
	TimeAdjustDaylightSavingOneHour
	TimeAdjustDaylightSavingTwoHour
)

// NewMSTimeZone creates a new MSTimeZone IE.
func NewMSTimeZone(tz time.Duration, daylightSaving uint8) *IE {
	i := New(MSTimeZone, make([]byte, 2))
	min := tz.Minutes() / 15
	absMin := int(math.Abs(min))
	hex, err := strconv.ParseInt(strconv.Itoa(absMin%10)+strconv.Itoa(absMin/10), 16, 8)
	if err != nil {
		return nil
	}
	if min < 0 {
		hex |= 0x08
	}
	i.Payload[0] = uint8(hex)
	i.Payload[1] = daylightSaving & 0x03

	return i
}

// TimeZone returns TimeZone in time.Duration if the type of IE matches.
func (i *IE) TimeZone() (time.Duration, error) {
	if i.Type != MSTimeZone {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	unsigned := i.Payload[0] & 0xf7
	dec := int((unsigned >> 4) + (unsigned&0x0f)*10)
	if (i.Payload[0]&0x08)>>3 == 1 {
		dec *= -1
	}

	return time.Duration(dec*15) * time.Minute, nil
}

// MustTimeZone returns TimeZone in time.Duration if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustTimeZone() time.Duration {
	v, _ := i.TimeZone()
	return v
}

// DaylightSaving returns DaylightSaving in uint8 if the type of IE matches.
func (i *IE) DaylightSaving() (uint8, error) {
	if i.Type != MSTimeZone {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[1], nil
}

// MustDaylightSaving returns DaylightSaving in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustDaylightSaving() uint8 {
	v, _ := i.DaylightSaving()
	return v
}
