// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "time"

// NewDelayValue creates a new DelayValue IE.
func NewDelayValue(delay time.Duration) *IE {
	return newUint8ValIE(DelayValue, uint8(delay.Seconds()*1000/50))
}

// DelayValue returns DelayValue in time.Duration if the type of IE matches.
func (i *IE) DelayValue() time.Duration {
	if i.Type != DelayValue {
		return time.Duration(0)
	}

	return time.Duration(i.Payload[0]/50) * time.Millisecond
}
