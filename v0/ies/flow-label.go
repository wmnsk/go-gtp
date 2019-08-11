// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "encoding/binary"

// NewFlowLabelDataI creates a new FlowLabelDataI IE.
func NewFlowLabelDataI(label uint16) *IE {
	return newUint16ValIE(FlowLabelDataI, label)
}

// FlowLabelDataI returns FlowLabelDataI if type matches.
func (i *IE) FlowLabelDataI() uint16 {
	if i.Type != FlowLabelDataI {
		return 0
	}
	if len(i.Payload) < 2 {
		return 0
	}

	return binary.BigEndian.Uint16(i.Payload)
}

// NewFlowLabelSignalling creates a new FlowLabelSignalling IE.
func NewFlowLabelSignalling(label uint16) *IE {
	return newUint16ValIE(FlowLabelSignalling, label)
}

// FlowLabelSignalling returns FlowLabelSignalling if type matches.
func (i *IE) FlowLabelSignalling() uint16 {
	if i.Type != FlowLabelSignalling {
		return 0
	}
	if len(i.Payload) < 2 {
		return 0
	}

	return binary.BigEndian.Uint16(i.Payload)
}

// NewFlowLabelDataII creates a new FlowLabelDataII IE.
func NewFlowLabelDataII(nsapi uint8, label uint16) *IE {
	i := New(FlowLabelDataII, make([]byte, 3))
	i.Payload[0] = nsapi | 0xf0
	binary.BigEndian.PutUint16(i.Payload[1:3], label)
	return i
}

// FlowLabelDataII returns FlowLabelDataII if type matches.
func (i *IE) FlowLabelDataII() []byte {
	if i.Type != FlowLabelDataII {
		return nil
	}
	return i.Payload
}

// NSAPI returns NSAPI in uint8 if type matches.
func (i *IE) NSAPI() uint8 {
	if len(i.Payload) == 0 {
		return 0
	}

	switch i.Type {
	case FlowLabelDataII:
		return i.Payload[0] & 0x0f
	default:
		return 0
	}
}

// FlowLabelData returns FlowLabelData in uint16 if type matches.
func (i *IE) FlowLabelData() uint16 {
	if len(i.Payload) < 2 {
		return 0
	}

	switch i.Type {
	case FlowLabelDataI:
		return i.FlowLabelDataI()
	case FlowLabelSignalling:
		return i.FlowLabelSignalling()
	case FlowLabelDataII:
		if len(i.Payload) < 3 {
			return 0
		}
		return binary.BigEndian.Uint16(i.Payload[1:3])
	default:
		return 0
	}
}
