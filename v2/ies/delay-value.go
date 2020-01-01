// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"io"
	"time"
)

// NewDelayValue creates a new DelayValue IE.
func NewDelayValue(delay time.Duration) *IE {
	return newUint8ValIE(DelayValue, uint8(delay.Seconds()*1000/50))
}

// DelayValue returns DelayValue in time.Duration if the type of IE matches.
func (i *IE) DelayValue() (time.Duration, error) {
	if i.Type != DelayValue {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
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
