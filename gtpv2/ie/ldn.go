// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewLocalDistinguishedName creates a new LocalDistinguishedName IE.
func NewLocalDistinguishedName(name string) *IE {
	return NewStringIE(LocalDistinguishedName, name)
}

// LocalDistinguishedName returns LocalDistinguishedName in string if the type of IE matches.
func (i *IE) LocalDistinguishedName() (string, error) {
	if i.Type != LocalDistinguishedName {
		return "", &InvalidTypeError{Type: i.Type}
	}
	return i.ValueAsString()
}

// MustLocalDistinguishedName returns LocalDistinguishedName in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustLocalDistinguishedName() string {
	v, _ := i.LocalDistinguishedName()
	return v
}
