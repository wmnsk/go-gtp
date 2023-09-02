// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewFlowLabelDataI creates a new FlowLabelDataI IE.
func NewFlowLabelDataI(label uint16) *IE {
	return newUint16ValIE(FlowLabelDataI, label)
}

// FlowLabelDataI returns FlowLabelDataI if type matches.
func (i *IE) FlowLabelDataI() (uint16, error) {
	if i.Type != FlowLabelDataI {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint16(i.Payload), nil
}

// MustFlowLabelDataI returns FlowLabelDataI in uint16 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustFlowLabelDataI() uint16 {
	v, _ := i.FlowLabelDataI()
	return v
}

// NewFlowLabelSignalling creates a new FlowLabelSignalling IE.
func NewFlowLabelSignalling(label uint16) *IE {
	return newUint16ValIE(FlowLabelSignalling, label)
}

// FlowLabelSignalling returns FlowLabelSignalling if type matches.
func (i *IE) FlowLabelSignalling() (uint16, error) {
	if i.Type != FlowLabelSignalling {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint16(i.Payload), nil
}

// MustFlowLabelSignalling returns FlowLabelSignalling in uint16 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustFlowLabelSignalling() uint16 {
	v, _ := i.FlowLabelSignalling()
	return v
}

// NewFlowLabelDataII creates a new FlowLabelDataII IE.
func NewFlowLabelDataII(nsapi uint8, label uint16) *IE {
	i := New(FlowLabelDataII, make([]byte, 3))
	i.Payload[0] = nsapi | 0xf0
	binary.BigEndian.PutUint16(i.Payload[1:3], label)
	return i
}

// FlowLabelDataII returns FlowLabelDataII if type matches.
func (i *IE) FlowLabelDataII() ([]byte, error) {
	if i.Type != FlowLabelDataII {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return i.Payload, nil
}

// MustFlowLabelDataII returns FlowLabelDataII in []byte if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustFlowLabelDataII() []byte {
	v, _ := i.FlowLabelDataII()
	return v
}

// NSAPI returns NSAPI in uint8 if type matches.
func (i *IE) NSAPI() (uint8, error) {
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case FlowLabelDataII:
		return i.Payload[0] & 0x0f, nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustNSAPI returns NSAPI in uint8 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustNSAPI() uint8 {
	v, _ := i.NSAPI()
	return v
}

// FlowLabelData returns FlowLabelData in uint16 if type matches.
func (i *IE) FlowLabelData() (uint16, error) {
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case FlowLabelDataI:
		return i.FlowLabelDataI()
	case FlowLabelSignalling:
		return i.FlowLabelSignalling()
	case FlowLabelDataII:
		if len(i.Payload) < 3 {
			return 0, io.ErrUnexpectedEOF
		}
		return binary.BigEndian.Uint16(i.Payload[1:3]), nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MustFlowLabelData returns FlowLabelData in uint16 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustFlowLabelData() uint16 {
	v, _ := i.FlowLabelData()
	return v
}
