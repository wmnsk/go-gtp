// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewRANAPCause creates a new RANAPCause IE.
func NewRANAPCause(cause uint8) *IE {
	return newUint8ValIE(RANAPCause, cause)
}

// RANAPCause returns RANAPCause in uint8 if type matches.
func (i *IE) RANAPCause() uint8 {
	if i.Type != RANAPCause {
		return 0
	}
	return i.Payload[0]
}
