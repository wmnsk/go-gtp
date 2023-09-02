// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
	"time"
)

// NewULITimestamp creates a new ULITimestamp IE.
func NewULITimestamp(ts time.Time) *IE {
	u64sec := uint64(ts.Sub(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC))) / 1000000000
	return newUint32ValIE(ULITimestamp, uint32(u64sec))
}

// Timestamp returns Timestamp in time.Time if the type of IE matches.
func (i *IE) Timestamp() (time.Time, error) {
	if len(i.Payload) < 4 {
		return time.Time{}, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case ULITimestamp, TWANIdentifierTimestamp:
		return time.Unix(int64(binary.BigEndian.Uint32(i.Payload)-2208988800), 0), nil
	default:
		return time.Time{}, &InvalidTypeError{Type: i.Type}
	}
}

// MustTimestamp returns Timestamp in time.Time, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustTimestamp() time.Time {
	v, _ := i.Timestamp()
	return v
}
