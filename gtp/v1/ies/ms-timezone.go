// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// Timezone adjustment definitions.
const (
	TimeAdjustNoDaylightSaving uint8 = iota
	TimeAdjustDaylightSavingOneHour
	TimeAdjustDaylightSavingTwoHour
)

// NewMSTimeZone creates a new MSTimeZone IE.
//
// XXX - Should be implemented properly.
func NewMSTimeZone(timezone, adj uint8) *IE {
	t := append([]byte{timezone}, []byte{adj}...)
	return New(MSTimeZone, t)
}

// MSTimeZone returns MSTimeZone value if type matches.
//
// XXX - Should be implemented properly.
func (i *IE) MSTimeZone() []byte {
	if i.Type != MSTimeZone {
		return nil
	}
	return i.Payload
}

// TimeZone returns TimeZone value if type matches.
//
// XXX - Should be implemented properly.
func (i *IE) TimeZone() uint8 {
	if i.Type != MSTimeZone {
		return 0
	}
	return i.Payload[0]
}

// DaylightSaving returns DaylightSaving value if type matches.
//
// XXX - Should be implemented properly.
func (i *IE) DaylightSaving() uint8 {
	if i.Type != MSTimeZone {
		return 0
	}
	return i.Payload[1]
}
