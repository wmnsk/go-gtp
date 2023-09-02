// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// LAC returns LAC value if type matches.
func (i *IE) LAC() (uint16, error) {
	switch i.Type {
	case RouteingAreaIdentity:
		if len(i.Payload) < 5 {
			return 0, io.ErrUnexpectedEOF
		}
		return binary.BigEndian.Uint16(i.Payload[3:5]), nil
	case UserLocationInformation:
		if len(i.Payload) < 6 {
			return 0, io.ErrUnexpectedEOF
		}
		return binary.BigEndian.Uint16(i.Payload[4:6]), nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustLAC returns LAC in uint16 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustLAC() uint16 {
	v, _ := i.LAC()
	return v
}
