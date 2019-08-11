// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewProcedureTransactionID creates a new ProcedureTransactionID IE.
func NewProcedureTransactionID(pti uint8) *IE {
	return newUint8ValIE(ProcedureTransactionID, pti)
}

// ProcedureTransactionID returns ProcedureTransactionID in uint8 if the type of IE matches.
func (i *IE) ProcedureTransactionID() uint8 {
	if i.Type != ProcedureTransactionID {
		return 0
	}

	return i.Payload[0]
}
