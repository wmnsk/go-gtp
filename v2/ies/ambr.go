// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
)

// NewAggregateMaximumBitRate creates a new AggregateMaximumBitRate IE.
func NewAggregateMaximumBitRate(up, down uint32) *IE {
	return newUint64ValIE(AggregateMaximumBitRate, (uint64(up)<<32 | uint64(down)))
}

// AggregateMaximumBitRateUp returns AggregateMaximumBitRate for Uplink
// if the type of IE matches.
func (i *IE) AggregateMaximumBitRateUp() uint32 {
	if i.Type != AggregateMaximumBitRate {
		return 0
	}
	if len(i.Payload) < 4 {
		return 0
	}

	return binary.BigEndian.Uint32(i.Payload[0:4])
}

// AggregateMaximumBitRateDown returns AggregateMaximumBitRate for Downlink
// if the type of IE matches.
func (i *IE) AggregateMaximumBitRateDown() uint32 {
	if i.Type != AggregateMaximumBitRate {
		return 0
	}
	if len(i.Payload) < 8 {
		return 0
	}

	return binary.BigEndian.Uint32(i.Payload[4:8])
}
