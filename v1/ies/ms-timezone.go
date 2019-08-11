// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
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
func (i *IE) TimeZone() time.Duration {
	if i.Type != MSTimeZone {
		return 0
	}
	unsigned := i.Payload[0] & 0xf7
	dec := int((unsigned >> 4) + (unsigned&0x0f)*10)
	if (i.Payload[0]&0x08)>>3 == 1 {
		dec *= -1
	}

	return time.Duration(dec*15) * time.Minute
}

// DaylightSaving returns DaylightSaving in uint8 if the type of IE matches.
func (i *IE) DaylightSaving() uint8 {
	if i.Type != MSTimeZone {
		return 0
	}

	return i.Payload[1]
}
