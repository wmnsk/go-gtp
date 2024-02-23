// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewExtendedCommonFlags creates a new ExtendedCommonFlags IE.
//
// Note: each flag should be set in 1 or 0.
func NewExtendedCommonFlags(uasi, bdwi, pcri, vb, retloc, cpsr, ccrsi, unauthenticatedIMSI int) *IE {
	return New(
		ExtendedCommonFlags,
		[]byte{uint8(
			uasi<<7 | bdwi<<6 | pcri<<5 | vb<<4 | retloc<<3 | cpsr<<2 | ccrsi<<1 | unauthenticatedIMSI,
		)},
	)
}

// ExtendedCommonFlags returns ExtendedCommonFlags value if type matches.
func (i *IE) ExtendedCommonFlags() (uint8, error) {
	if i.Type != ExtendedCommonFlags {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustExtendedCommonFlags returns ExtendedCommonFlags in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustExtendedCommonFlags() uint8 {
	v, _ := i.ExtendedCommonFlags()
	return v
}

// IsUASI checks if UASI flag exists in ExtendedCommonFlags.
func (i *IE) IsUASI() bool {
	return ((i.MustExtendedCommonFlags() >> 7) & 0x01) != 0
}

// IsBDWI checks if BDWI flag exists in ExtendedCommonFlags.
func (i *IE) IsBDWI() bool {
	return ((i.MustExtendedCommonFlags() >> 6) & 0x01) != 0
}

// IsPCRI checks if PCRI flag exists in ExtendedCommonFlags.
func (i *IE) IsPCRI() bool {
	return ((i.MustExtendedCommonFlags() >> 5) & 0x01) != 0
}

// IsVB checks if VB flag exists in ExtendedCommonFlags.
func (i *IE) IsVB() bool {
	return ((i.MustExtendedCommonFlags() >> 4) & 0x01) != 0
}

// IsRetLoc checks if RetLoc flag exists in ExtendedCommonFlags.
func (i *IE) IsRetLoc() bool {
	return ((i.MustExtendedCommonFlags() >> 3) & 0x01) != 0
}

// IsCPSR checks if CPSR flag exists in ExtendedCommonFlags.
func (i *IE) IsCPSR() bool {
	return ((i.MustExtendedCommonFlags() >> 2) & 0x01) != 0
}

// IsCCRSI checks if CCRSI flag exists in ExtendedCommonFlags.
func (i *IE) IsCCRSI() bool {
	return ((i.MustExtendedCommonFlags() >> 1) & 0x01) != 0
}

// IsUnauthenticatedIMSI checks if UnauthenticatedIMSI flag exists in ExtendedCommonFlags.
func (i *IE) IsUnauthenticatedIMSI() bool {
	return (i.MustExtendedCommonFlags() & 0x01) != 0
}
