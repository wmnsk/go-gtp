// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "io"

// NewFullyQualifiedDomainName creates a new FullyQualifiedDomainName IE.
func NewFullyQualifiedDomainName(fqdn string) *IE {
	return newStringIE(FullyQualifiedDomainName, fqdn)
}

// FullyQualifiedDomainName returns FullyQualifiedDomainName in string if the type of IE matches.
func (i *IE) FullyQualifiedDomainName() (string, error) {
	if i.Type != FullyQualifiedDomainName {
		return "", &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return "", io.ErrUnexpectedEOF
	}

	return string(i.Payload), nil
}

// MustFullyQualifiedDomainName returns FullyQualifiedDomainName in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustFullyQualifiedDomainName() string {
	v, _ := i.FullyQualifiedDomainName()
	return v
}
