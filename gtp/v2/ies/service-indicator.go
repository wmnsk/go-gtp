// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewServiceIndicator creates a new ServiceIndicator IE.
func NewServiceIndicator(ind uint8) *IE {
	return newUint8ValIE(ServiceIndicator, ind)
}

// ServiceIndicator returns ServiceIndicator in uint8 if the type of IE matches.
func (i *IE) ServiceIndicator() uint8 {
	if i.Type != ServiceIndicator {
		return 0
	}

	return i.Payload[0]
}
