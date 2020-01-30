// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "io"

// NewAllocationRetensionPriority creates a new AllocationRetensionPriority IE.
func NewAllocationRetensionPriority(pci, pl, pvi uint8) *IE {
	i := New(AllocationRetensionPriority, 0x00, make([]byte, 1))
	i.Payload[0] |= (pci << 6 & 0x40) | (pl << 2 & 0x3c) | (pvi & 0x01)
	return i
}

// PreemptionCapability reports whether the preemption capability is set to enabled if the type of IE matches.
func (i *IE) PreemptionCapability() bool {
	if len(i.Payload) == 0 {
		return false
	}

	switch i.Type {
	case AllocationRetensionPriority, BearerQoS:
		return (i.Payload[0] & 0x40) != 1
	default:
		return false
	}
}

// PriorityLevel returns PriorityLevel in uint8 if the type of IE matches.
func (i *IE) PriorityLevel() (uint8, error) {
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case AllocationRetensionPriority, BearerQoS:
		return (i.Payload[0] & 0x3c) >> 2, nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// PreemptionVulnerability reports whether the preemption vulnerability is set to enabled if the type of IE matches.
func (i *IE) PreemptionVulnerability() bool {
	if len(i.Payload) == 0 {
		return false
	}

	switch i.Type {
	case AllocationRetensionPriority, BearerQoS:
		return (i.Payload[0] & 0x01) != 1
	default:
		return false
	}
}
