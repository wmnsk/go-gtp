// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewMSNotReachableReason creates a new MSNotReachableReason IE.
func NewMSNotReachableReason(reason uint8) *IE {
	return newUint8ValIE(MSNotReachableReason, reason)
}

// MSNotReachableReason returns MSNotReachableReason value if type matches.
func (i *IE) MSNotReachableReason() uint8 {
	if i.Type != MSNotReachableReason {
		return 0
	}
	if len(i.Payload) == 0 {
		return 0
	}

	return i.Payload[0]
}
