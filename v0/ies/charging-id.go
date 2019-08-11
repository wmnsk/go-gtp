// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
)

// NewChargingID creates a new ChargingID IE.
func NewChargingID(id uint32) *IE {
	return newUint32ValIE(ChargingID, id)
}

// ChargingID returns ChargingID value in uint32 if type matches.
func (i *IE) ChargingID() uint32 {
	if i.Type != ChargingID {
		return 0
	}
	if len(i.Payload) < 4 {
		return 0
	}

	return binary.BigEndian.Uint32(i.Payload)
}
