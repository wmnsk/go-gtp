// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
)

// NewNodeFeatures creates a new NodeFeatures IE.
func NewNodeFeatures(flags uint8) *IE {
	return newUint8ValIE(NodeFeatures, flags)
}

// NodeFeatures returns NodeFeatures in uint8 if the type of IE matches.
func (i *IE) NodeFeatures() (uint8, error) {
	if i.Type != NodeFeatures {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// MustNodeFeatures returns NodeFeatures in uint8 if the type of IE matches.
func (i *IE) MustNodeFeatures() uint8 {
	v, _ := i.NodeFeatures()
	return v
}

// HasPRN reports whether an IE has PRN bit.
func (i *IE) HasPRN() bool {
	v, err := i.NodeFeatures()
	if err != nil {
		return false
	}

	return has1stBit(v)
}

// HasMABR reports whether an IE has MABR bit.
func (i *IE) HasMABR() bool {
	v, err := i.NodeFeatures()
	if err != nil {
		return false
	}

	return has2ndBit(v)
}

// HasNTSR reports whether an IE has NTSR bit.
func (i *IE) HasNTSR() bool {
	v, err := i.NodeFeatures()
	if err != nil {
		return false
	}

	return has3rdBit(v)
}

// HasCIOT reports whether an IE has CIOT bit.
func (i *IE) HasCIOT() bool {
	v, err := i.NodeFeatures()
	if err != nil {
		return false
	}

	return has4thBit(v)
}

// HasS1UN reports whether an IE has S1UN bit.
func (i *IE) HasS1UN() bool {
	v, err := i.NodeFeatures()
	if err != nil {
		return false
	}

	return has5thBit(v)
}

// HasETH reports whether an IE has ETH bit.
func (i *IE) HasETH() bool {
	v, err := i.NodeFeatures()
	if err != nil {
		return false
	}

	return has6thBit(v)
}

// HasMTEDT reports whether an IE has MTEDT bit.
func (i *IE) HasMTEDT() bool {
	v, err := i.NodeFeatures()
	if err != nil {
		return false
	}

	return has7thBit(v)
}
