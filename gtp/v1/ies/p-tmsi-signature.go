// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "github.com/wmnsk/go-gtp/gtp/utils"

// NewPTMSISignature creates a new PTMSISignature IE.
func NewPTMSISignature(sig uint32) *IE {
	return New(PTMSISignature, utils.Uint32To24(sig))
}

// PTMSISignature returns PTMSISignature value in uint32 if type matches.
func (i *IE) PTMSISignature() uint32 {
	if i.Type != PTMSISignature {
		return 0
	}
	return utils.Uint24To32(i.Payload)
}
