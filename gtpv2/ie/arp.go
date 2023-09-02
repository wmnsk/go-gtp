// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewAllocationRetensionPriority creates a new AllocationRetensionPriority IE.
func NewAllocationRetensionPriority(pci, pl, pvi uint8) *IE {
	i := New(AllocationRetensionPriority, 0x00, make([]byte, 1))
	i.Payload[0] |= (pci << 6 & 0x40) | (pl << 2 & 0x3c) | (pvi & 0x01)
	return i
}

func (i *IE) AllocationRetensionPriority() (uint8, error) {
	if i.Type != AllocationRetensionPriority {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// HasPVI reports whether an IE has PVI bit.
func (i *IE) HasPVI() bool {
	v, err := i.AllocationRetensionPriority()
	if err != nil {
		return false
	}

	return has1stBit(v)
}

// HasPCI reports whether an IE has PCI bit.
func (i *IE) HasPCI() bool {
	v, err := i.AllocationRetensionPriority()
	if err != nil {
		return false
	}

	return has7thBit(v)
}

// PriorityLevel returns PriorityLevel in uint8 if the type of IE matches.
func (i *IE) PriorityLevel() (uint8, error) {
	switch i.Type {
	case AllocationRetensionPriority:
		v, err := i.AllocationRetensionPriority()
		if err != nil {
			return 0, err
		}

		return (v & 0x3c) >> 2, nil
	case BearerQoS:
		v, err := i.BearerQoS()
		if err != nil {
			return 0, err
		}

		return (v.ARP & 0x3c) >> 2, nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// PreemptionVulnerability reports whether the preemption vulnerability is set to enabled if the type of IE matches.
//
// Deprecated: use HasPVI instead.
func (i *IE) PreemptionVulnerability() bool {
	return i.HasPVI()
}

// PreemptionCapability reports whether the preemption capability is set to enabled if the type of IE matches.
//
// Deprecated: use HasPCI instead.
func (i *IE) PreemptionCapability() bool {
	return i.HasPCI()
}
