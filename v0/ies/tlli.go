// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
)

// NewTemporaryLogicalLinkIdentity creates a new TemporaryLogicalLinkIdentity IE.
func NewTemporaryLogicalLinkIdentity(tlli uint32) *IE {
	return newUint32ValIE(TemporaryLogicalLinkIdentity, tlli)
}

// TemporaryLogicalLinkIdentity returns TemporaryLogicalLinkIdentity value in uint32 if type matches.
func (i *IE) TemporaryLogicalLinkIdentity() uint32 {
	if i.Type != TemporaryLogicalLinkIdentity {
		return 0
	}
	return binary.BigEndian.Uint32(i.Payload)
}
