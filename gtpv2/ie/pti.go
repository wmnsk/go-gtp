// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewProcedureTransactionID creates a new ProcedureTransactionID IE.
func NewProcedureTransactionID(pti uint8) *IE {
	return newUint8ValIE(ProcedureTransactionID, pti)
}

// ProcedureTransactionID returns ProcedureTransactionID in uint8 if the type of IE matches.
func (i *IE) ProcedureTransactionID() (uint8, error) {
	if i.Type != ProcedureTransactionID {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustProcedureTransactionID returns ProcedureTransactionID in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustProcedureTransactionID() uint8 {
	v, _ := i.ProcedureTransactionID()
	return v
}
