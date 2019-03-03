// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewLocalDistinguishedName creates a new LocalDistinguishedName IE.
func NewLocalDistinguishedName(name string) *IE {
	return newStringIE(LocalDistinguishedName, name)
}

// LocalDistinguishedName returns LocalDistinguishedName in string if the type of IE matches.
func (i *IE) LocalDistinguishedName() string {
	if i.Type != LocalDistinguishedName {
		return ""
	}

	return string(i.Payload)
}
