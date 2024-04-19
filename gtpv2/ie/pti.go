// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewProcedureTransactionID creates a new ProcedureTransactionID IE.
func NewProcedureTransactionID(pti uint8) *IE {
	return NewUint8IE(ProcedureTransactionID, pti)
}

// ProcedureTransactionID returns ProcedureTransactionID in uint8 if the type of IE matches.
func (i *IE) ProcedureTransactionID() (uint8, error) {
	if i.Type != ProcedureTransactionID {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	return i.ValueAsUint8()
}

// MustProcedureTransactionID returns ProcedureTransactionID in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustProcedureTransactionID() uint8 {
	v, _ := i.ProcedureTransactionID()
	return v
}
