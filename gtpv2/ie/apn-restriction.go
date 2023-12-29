// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewAPNRestriction creates a new APNRestriction IE.
func NewAPNRestriction(restriction uint8) *IE {
	return NewUint8IE(APNRestriction, restriction)
}

// APNRestriction returns APNRestriction in uint8 if the type of IE matches.
func (i *IE) APNRestriction() (uint8, error) {
	if i.Type != APNRestriction {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.ValueAsUint8()
}

// MustAPNRestriction returns APNRestriction in uint8, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustAPNRestriction() uint8 {
	v, _ := i.APNRestriction()
	return v
}

// RestrictionType returns RestrictionType in uint8 if the type of IE matches.
func (i *IE) RestrictionType() (uint8, error) {
	return i.APNRestriction()
}
