// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// RAC returns RAC value if type matches.
func (i *IE) RAC() (uint8, error) {
	switch i.Type {
	case RouteingAreaIdentity:
		if len(i.Payload) < 6 {
			return 0, io.ErrUnexpectedEOF
		}
		return i.Payload[5], nil
	case UserLocationInformation:
		if len(i.Payload) < 7 {
			return 0, io.ErrUnexpectedEOF
		}
		if i.Payload[0] == locTypeRAI {
			return i.Payload[6], nil
		}
	}
	return 0, &InvalidTypeError{Type: i.Type}
}

// MustRAC returns RAC in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustRAC() uint8 {
	v, _ := i.RAC()
	return v
}
