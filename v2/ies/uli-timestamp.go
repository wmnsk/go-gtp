// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
	"time"
)

// NewULITimestamp creates a new ULITimestamp IE.
func NewULITimestamp(ts time.Time) *IE {
	u64sec := uint64(ts.Sub(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC))) / 1000000000
	return newUint32ValIE(ULITimestamp, uint32(u64sec))
}

// Timestamp returns Timestamp in time.Time if the type of IE matches.
func (i *IE) Timestamp() time.Time {
	if len(i.Payload) < 4 {
		return time.Time{}
	}

	switch i.Type {
	case ULITimestamp, TWANIdentifierTimestamp:
		return time.Unix(int64(binary.BigEndian.Uint32(i.Payload)-2208988800), 0)
	default:
		return time.Time{}
	}
}
