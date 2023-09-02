// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewExtensionHeaderTypeList creates a new ExtensionHeaderTypeList IE.
func NewExtensionHeaderTypeList(types ...uint8) *IE {
	return New(ExtensionHeaderTypeList, types)
}

// ExtensionHeaderTypeList returns ExtensionHeaderTypeList in []uint8 if type matches.
func (i *IE) ExtensionHeaderTypeList() ([]uint8, error) {
	if i.Type != ExtensionHeaderTypeList {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// MustExtensionHeaderTypeList returns ExtensionHeaderTypeList in []uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustExtensionHeaderTypeList() []uint8 {
	v, _ := i.ExtensionHeaderTypeList()
	return v
}
