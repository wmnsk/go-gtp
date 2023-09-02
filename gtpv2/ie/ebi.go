// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"fmt"
	"io"
)

// NewEPSBearerID creates a new EPSBearerID IE.
func NewEPSBearerID(ebi uint8) *IE {
	return newUint8ValIE(EPSBearerID, ebi&0x0f)
}

// EPSBearerID returns EPSBearerID if the type of IE matches.
func (i *IE) EPSBearerID() (uint8, error) {
	switch i.Type {
	case EPSBearerID:
		if len(i.Payload) < 1 {
			return 0, io.ErrUnexpectedEOF
		}

		return i.Payload[0], nil
	case BearerContext:
		ies, err := i.BearerContext()
		if err != nil {
			return 0, fmt.Errorf("failed to retrieve EPSBearerID: %w", err)
		}

		for _, child := range ies {
			if child.Type == EPSBearerID {
				return child.EPSBearerID()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustEPSBearerID returns EPSBearerID in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustEPSBearerID() uint8 {
	v, _ := i.EPSBearerID()
	return v
}
