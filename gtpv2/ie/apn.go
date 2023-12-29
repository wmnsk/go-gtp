// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewAccessPointName creates a new AccessPointName IE.
func NewAccessPointName(apn string) *IE {
	return NewFQDNIE(AccessPointName, apn)
}

// AccessPointName returns AccessPointName in string if the type of IE matches.
func (i *IE) AccessPointName() (string, error) {
	if i.Type != AccessPointName {
		return "", &InvalidTypeError{Type: i.Type}
	}
	return i.ValueAsFQDN()
}

// MustAccessPointName returns AccessPointName in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustAccessPointName() string {
	v, _ := i.AccessPointName()
	return v
}
