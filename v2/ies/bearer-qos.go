// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
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
func (i *IE) QCILabel() uint8 {
	switch i.Type {
	case BearerQoS:
		return i.Payload[1]
	case FlowQoS:
		return i.Payload[0]
	default:
		return 0
	}
}

// MBRForUplink returns MBRForUplink in uint64 if the type of IE matches.
func (i *IE) MBRForUplink() uint64 {
	switch i.Type {
	case BearerQoS:
		return utils.Uint40To64(i.Payload[3:7])
	case FlowQoS:
		return utils.Uint40To64(i.Payload[2:6])
	default:
		return 0
	}
}

// MBRForDownlink returns MBRForDownlink in uint64 if the type of IE matches.
func (i *IE) MBRForDownlink() uint64 {
	switch i.Type {
	case BearerQoS:
		return utils.Uint40To64(i.Payload[7:12])
	case FlowQoS:
		return utils.Uint40To64(i.Payload[6:11])
	default:
		return 0
	}
}

// GBRForUplink returns GBRForUplink in uint64 if the type of IE matches.
func (i *IE) GBRForUplink() uint64 {
	switch i.Type {
	case BearerQoS:
		return utils.Uint40To64(i.Payload[12:17])
	case FlowQoS:
		return utils.Uint40To64(i.Payload[11:16])
	default:
		return 0
	}
}

// GBRForDownlink returns GBRForDownlink in uint64 if the type of IE matches.
func (i *IE) GBRForDownlink() uint64 {
	switch i.Type {
	case BearerQoS:
		return utils.Uint40To64(i.Payload[17:22])
	case FlowQoS:
		return utils.Uint40To64(i.Payload[16:21])
	default:
		return 0
	}
}
