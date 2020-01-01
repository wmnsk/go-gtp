// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
	"io"
)

// NewAggregateMaximumBitRate creates a new AggregateMaximumBitRate IE.
func NewAggregateMaximumBitRate(up, down uint32) *IE {
	return newUint64ValIE(AggregateMaximumBitRate, (uint64(up)<<32 | uint64(down)))
}

// AggregateMaximumBitRateUp returns AggregateMaximumBitRate for Uplink
// if the type of IE matches.
func (i *IE) AggregateMaximumBitRateUp() (uint32, error) {
	if i.Type != AggregateMaximumBitRate {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint32(i.Payload[0:4]), nil
}

// MustAggregateMaximumBitRateUp returns AggregateMaximumBitRateUp in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustAggregateMaximumBitRateUp() uint32 {
	v, _ := i.AggregateMaximumBitRateUp()
	return v
}

// AggregateMaximumBitRateDown returns AggregateMaximumBitRate for Downlink
// if the type of IE matches.
func (i *IE) AggregateMaximumBitRateDown() (uint32, error) {
	if i.Type != AggregateMaximumBitRate {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 8 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint32(i.Payload[4:8]), nil
}

// MustAggregateMaximumBitRateDown returns AggregateMaximumBitRateDown in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustAggregateMaximumBitRateDown() uint32 {
	v, _ := i.AggregateMaximumBitRateDown()
	return v
}
