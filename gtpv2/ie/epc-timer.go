// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"errors"
	"io"
	"math"
	"time"
)

// NewEPCTimer creates a new Timer IE.
func NewEPCTimer(duration time.Duration) *IE {
	// 8.87 EPC Timer
	// Timer unit
	// Bits 6 to 8 defines the timer value unit as follows: Bits
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

	var unit, value uint8
	switch {
	case duration%(10*time.Hour) == 0:
		unit = 0x80
		value = uint8(duration / (10 * time.Hour))
	case duration%(1*time.Hour) == 0:
		unit = 0x60
		value = uint8(duration / time.Hour)
	case duration%(10*time.Minute) == 0:
		unit = 0x40
		value = uint8(duration / (10 * time.Minute))
	case duration%(1*time.Minute) == 0:
		unit = 0x20
		value = uint8(duration / time.Minute)
	case duration%(2*time.Second) == 0:
		unit = 0x00
		value = uint8(duration / (2 * time.Second))
	default:
		unit = 0xe0
		value = 0
	}

	return newUint8ValIE(EPCTimer, unit+(value&0x1f))
}

// EPCTimer returns EPCTimer in time.Duration if the type of IE matches.
func (i *IE) EPCTimer() (time.Duration, error) {
	return i.Timer()
}

// MustEPCTimer returns EPCTimer in time.Duration, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustEPCTimer() time.Duration {
	v, _ := i.EPCTimer()
	return v
}

// Timer returns Timer in time.Duration if the type of IE matches.
func (i *IE) Timer() (time.Duration, error) {
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case EPCTimer:
		var d time.Duration
		switch (i.Payload[0] & 0xe0) >> 5 {
		case 0x07:
			d = time.Duration(math.MaxInt64)
		case 0x04:
			d = time.Duration(i.Payload[0]&0x1f) * 10 * time.Hour
		case 0x03:
			d = time.Duration(i.Payload[0]&0x1f) * time.Hour
		case 0x02:
			d = time.Duration(i.Payload[0]&0x1f) * 10 * time.Minute
		case 0x01:
			d = time.Duration(i.Payload[0]&0x1f) * time.Minute
		case 0x00:
			d = time.Duration(i.Payload[0]&0x1f) * 2 * time.Second
		default:
			d = 0
		}
		return d, nil
	case OverloadControlInformation:
		return 0, errors.New("not implemented")
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
