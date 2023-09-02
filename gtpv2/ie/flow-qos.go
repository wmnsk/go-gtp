// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"

	"github.com/wmnsk/go-gtp/utils"
)

// NewFlowQoS creates a new FlowQoS IE.
func NewFlowQoS(qci uint8, umbr, dmbr, ugbr, dgbr uint64) *IE {
	v := NewFlowQoSFields(qci, umbr, dmbr, ugbr, dgbr)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(FlowQoS, 0x00, b)
}

// FlowQoS returns FlowQoS in FlowQoSFields type if the type of IE matches.
func (i *IE) FlowQoS() (*FlowQoSFields, error) {
	switch i.Type {
	case FlowQoS:
		return ParseFlowQoSFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// FlowQoSFields is a set of fields in FlowQoS IE.
type FlowQoSFields struct {
	QCI                          uint8
	MaximumBitRateForUplink      uint64 // 40 bits
	MaximumBitRateForDownlink    uint64 // 40 bits
	GuaranteedBitRateForUplink   uint64 // 40 bits
	GuaranteedBitRateForDownlink uint64 // 40 bits
}

// NewFlowQoSFields creates a new FlowQoSFields.
func NewFlowQoSFields(qci uint8, umbr, dmbr, ugbr, dgbr uint64) *FlowQoSFields {
	return &FlowQoSFields{
		QCI:                          qci,
		MaximumBitRateForUplink:      umbr,
		MaximumBitRateForDownlink:    dmbr,
		GuaranteedBitRateForUplink:   ugbr,
		GuaranteedBitRateForDownlink: dgbr,
	}
}

// Marshal serializes FlowQoSFields.
func (f *FlowQoSFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes FlowQoSFields.
func (f *FlowQoSFields) MarshalTo(b []byte) error {
	if len(b) < 21 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.QCI

	copy(b[1:6], utils.Uint64To40(f.MaximumBitRateForUplink))
	copy(b[6:11], utils.Uint64To40(f.MaximumBitRateForDownlink))
	copy(b[11:16], utils.Uint64To40(f.GuaranteedBitRateForUplink))
	copy(b[16:21], utils.Uint64To40(f.GuaranteedBitRateForDownlink))

	return nil
}

// ParseFlowQoSFields decodes FlowQoSFields.
func ParseFlowQoSFields(b []byte) (*FlowQoSFields, error) {
	f := &FlowQoSFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into FlowQoSFields.
func (f *FlowQoSFields) UnmarshalBinary(b []byte) error {
	if len(b) < 21 {
		return io.ErrUnexpectedEOF
	}

	f.QCI = b[0]

	f.MaximumBitRateForUplink = utils.Uint40To64(b[1:6])
	f.MaximumBitRateForDownlink = utils.Uint40To64(b[6:11])
	f.GuaranteedBitRateForUplink = utils.Uint40To64(b[11:16])
	f.GuaranteedBitRateForDownlink = utils.Uint40To64(b[16:21])

	return nil
}

// MarshalLen returns the serial length of FlowQoSFields in int.
func (f *FlowQoSFields) MarshalLen() int {
	return 21
}
