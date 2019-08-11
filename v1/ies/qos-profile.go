// Copyright 2019 go-gtp authors. All rights reserved.
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
func (i *IE) QoSProfile() []byte {
	if i.Type != QoSProfile {
		return nil
	}
	return i.Payload
}
