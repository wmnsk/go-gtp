// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewQoSProfile creates a new QoSProfile IE.
//
// XXX - NOT Fully implemented. Users need to put the whole payload in []byte.
func NewQoSProfile(payload []byte) *IE {
	return New(QoSProfile, payload)
}

// QoSProfile returns QoSProfile if type matches.
//
// XXX - NOT Fully implemented. This method just returns the whole payload in []byte.
func (i *IE) QoSProfile() ([]byte, error) {
	if i.Type != QoSProfile {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return i.Payload, nil
}

// MustQoSProfile returns QoSProfile in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustQoSProfile() []byte {
	v, _ := i.QoSProfile()
	return v
}
