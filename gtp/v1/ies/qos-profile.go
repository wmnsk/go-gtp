// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewQoSProfile creates a new QoSProfile IE.
// XXX - NOT IMPLEMENTED YET. RETURNS EMPTY IE.
func NewQoSProfile() *IE {
	return New(QoSProfile, []byte{})
}

// QoSProfile returns QoSProfile if type matches.
// XXX - NOT IMPLEMENTED YET. RETURNS NIL.
func (i *IE) QoSProfile() []byte {
	if i.Type != QoSProfile {
		return nil
	}
	return nil
}
