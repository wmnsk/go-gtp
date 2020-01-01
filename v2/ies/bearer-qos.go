// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"io"

	"github.com/wmnsk/go-gtp/utils"
)

// NewBearerQoS creates a new BearerQoS IE.
func NewBearerQoS(pci, pl, pvi, qci uint8, umbr, dmbr, ugbr, dgbr uint64) *IE {
	i := New(BearerQoS, 0x00, make([]byte, 22))
	i.Payload[0] |= (pci << 6 & 0x40) | (pl << 2 & 0x3c) | (pvi & 0x01)
	i.Payload[1] = qci
	copy(i.Payload[2:7], utils.Uint64To40(umbr))
	copy(i.Payload[7:12], utils.Uint64To40(dmbr))
	copy(i.Payload[12:17], utils.Uint64To40(ugbr))
	copy(i.Payload[17:22], utils.Uint64To40(dgbr))
	return i
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
		return utils.Uint40To64(i.Payload[3:7]), nil
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
