// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"fmt"
	"io"

	"github.com/wmnsk/go-gtp/utils"
)

// NewBearerQoS creates a new BearerQoS IE.
func NewBearerQoS(pci, pl, pvi, qci uint8, umbr, dmbr, ugbr, dgbr uint64) *IE {
	v := NewBearerQoSFields(pci, pl, pvi, qci, umbr, dmbr, ugbr, dgbr)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(BearerQoS, 0x00, b)
}

// BearerQoS returns BearerQoS in BearerQoSFields type if the type of IE matches.
func (i *IE) BearerQoS() (*BearerQoSFields, error) {
	switch i.Type {
	case BearerQoS:
		return ParseBearerQoSFields(i.Payload)
	case BearerContext:
		ies, err := i.BearerContext()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve BearerQoS: %w", err)
		}

		for _, child := range ies {
			if child.Type == BearerQoS {
				return child.BearerQoS()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// BearerQoSFields is a set of fields in BearerQoS IE.
type BearerQoSFields struct {
	ARP                          uint8
	QCI                          uint8
	MaximumBitRateForUplink      uint64 // 40 bits
	MaximumBitRateForDownlink    uint64 // 40 bits
	GuaranteedBitRateForUplink   uint64 // 40 bits
	GuaranteedBitRateForDownlink uint64 // 40 bits
}

// NewBearerQoSFields creates a new BearerQoSFields.
func NewBearerQoSFields(pci, pl, pvi, qci uint8, umbr, dmbr, ugbr, dgbr uint64) *BearerQoSFields {
	return &BearerQoSFields{
		ARP:                          (pci << 6 & 0x40) | (pl << 2 & 0x3c) | (pvi & 0x01),
		QCI:                          qci,
		MaximumBitRateForUplink:      umbr,
		MaximumBitRateForDownlink:    dmbr,
		GuaranteedBitRateForUplink:   ugbr,
		GuaranteedBitRateForDownlink: dgbr,
	}
}

// Marshal serializes BearerQoSFields.
func (f *BearerQoSFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes BearerQoSFields.
func (f *BearerQoSFields) MarshalTo(b []byte) error {
	if len(b) < 22 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.ARP
	b[1] = f.QCI

	copy(b[2:7], utils.Uint64To40(f.MaximumBitRateForUplink))
	copy(b[7:12], utils.Uint64To40(f.MaximumBitRateForDownlink))
	copy(b[12:17], utils.Uint64To40(f.GuaranteedBitRateForUplink))
	copy(b[17:22], utils.Uint64To40(f.GuaranteedBitRateForDownlink))

	return nil
}

// ParseBearerQoSFields decodes BearerQoSFields.
func ParseBearerQoSFields(b []byte) (*BearerQoSFields, error) {
	f := &BearerQoSFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into BearerQoSFields.
func (f *BearerQoSFields) UnmarshalBinary(b []byte) error {
	if len(b) < 22 {
		return io.ErrUnexpectedEOF
	}

	f.ARP = b[0]
	f.QCI = b[1]

	f.MaximumBitRateForUplink = utils.Uint40To64(b[2:7])
	f.MaximumBitRateForDownlink = utils.Uint40To64(b[7:12])
	f.GuaranteedBitRateForUplink = utils.Uint40To64(b[12:17])
	f.GuaranteedBitRateForDownlink = utils.Uint40To64(b[17:22])

	return nil
}

// MarshalLen returns the serial length of BearerQoSFields in int.
func (f *BearerQoSFields) MarshalLen() int {
	return 22
}

// QCILabel returns QCILabel in uint8 if the type of IE matches.
func (i *IE) QCILabel() (uint8, error) {
	switch i.Type {
	case BearerQoS:
		if len(i.Payload) < 2 {
			return 0, io.ErrUnexpectedEOF
		}
		return i.Payload[1], nil
	case FlowQoS:
		if len(i.Payload) < 1 {
			return 0, io.ErrUnexpectedEOF
		}
		return i.Payload[0], nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// MBRForUplink returns MBRForUplink in uint64 if the type of IE matches.
func (i *IE) MBRForUplink() (uint64, error) {
	switch i.Type {
	case BearerQoS:
		if len(i.Payload) < 7 {
			return 0, io.ErrUnexpectedEOF
		}
		return utils.Uint40To64(i.Payload[2:7]), nil
	case FlowQoS:
		if len(i.Payload) < 6 {
			return 0, io.ErrUnexpectedEOF
		}
		return utils.Uint40To64(i.Payload[2:6]), nil
	default:
		return 0, io.ErrUnexpectedEOF
	}
}

// MustMBRForUplink returns MBRForUplink in uint64, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustMBRForUplink() uint64 {
	v, _ := i.MBRForUplink()
	return v
}

// MBRForDownlink returns MBRForDownlink in uint64 if the type of IE matches.
func (i *IE) MBRForDownlink() (uint64, error) {
	switch i.Type {
	case BearerQoS:
		if len(i.Payload) < 12 {
			return 0, io.ErrUnexpectedEOF
		}
		return utils.Uint40To64(i.Payload[7:12]), nil
	case FlowQoS:
		if len(i.Payload) < 11 {
			return 0, io.ErrUnexpectedEOF
		}
		return utils.Uint40To64(i.Payload[6:11]), nil
	default:
		return 0, io.ErrUnexpectedEOF
	}
}

// MustMBRForDownlink returns MBRForDownlink in uint64, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustMBRForDownlink() uint64 {
	v, _ := i.MBRForDownlink()
	return v
}

// GBRForUplink returns GBRForUplink in uint64 if the type of IE matches.
func (i *IE) GBRForUplink() (uint64, error) {
	switch i.Type {
	case BearerQoS:
		if len(i.Payload) < 17 {
			return 0, io.ErrUnexpectedEOF
		}
		return utils.Uint40To64(i.Payload[12:17]), nil
	case FlowQoS:
		if len(i.Payload) < 16 {
			return 0, io.ErrUnexpectedEOF
		}
		return utils.Uint40To64(i.Payload[11:16]), nil
	default:
		return 0, io.ErrUnexpectedEOF
	}
}

// MustGBRForUplink returns GBRForUplink in uint64, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustGBRForUplink() uint64 {
	v, _ := i.GBRForUplink()
	return v
}

// GBRForDownlink returns GBRForDownlink in uint64 if the type of IE matches.
func (i *IE) GBRForDownlink() (uint64, error) {
	switch i.Type {
	case BearerQoS:
		if len(i.Payload) < 22 {
			return 0, io.ErrUnexpectedEOF
		}
		return utils.Uint40To64(i.Payload[17:22]), nil
	case FlowQoS:
		if len(i.Payload) < 21 {
			return 0, io.ErrUnexpectedEOF
		}
		return utils.Uint40To64(i.Payload[16:21]), nil
	default:
		return 0, io.ErrUnexpectedEOF
	}
}

// MustGBRForDownlink returns GBRForDownlink in uint64, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustGBRForDownlink() uint64 {
	v, _ := i.GBRForDownlink()
	return v
}
