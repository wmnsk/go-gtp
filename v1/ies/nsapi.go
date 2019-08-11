// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewNSAPI creates a new NSAPI IE.
func NewNSAPI(nsapi uint8) *IE {
	return newUint8ValIE(NSAPI, nsapi)
}

// NSAPI returns NSAPI value if type matches.
func (i *IE) NSAPI() uint8 {
	if i.Type != NSAPI {
		return 0
	}
	return i.Payload[0]
}
