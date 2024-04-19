// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "sync"

// NewGroupedIE creates a new IE with the given IEs.
//
// The IEs with nil value will be ignored.
func NewGroupedIE(itype uint8, ies ...*IE) *IE {
	i := New(itype, 0x00, make([]byte, 0))
	for _, ie := range ies {
		if ie == nil {
			continue
		}

		serialized, err := ie.Marshal()
		if err != nil {
			return nil
		}

		i.Payload = append(i.Payload, serialized...)
		i.ChildIEs = append(i.ChildIEs, ie)
	}
	i.SetLength()

	return i
}

// We're using map to avoid iterating over a list.
// The value `true` is not actually used.
// TODO: consider using a slice with utils in slices package introduced in Go 1.21.
var (
	mu                  sync.RWMutex
	defaultGroupedIEMap = map[uint8]bool{
		BearerContext:              true,
		PDNConnection:              true,
		OverloadControlInformation: true,
		LoadControlInformation:     true,
		RemoteUEContext:            true,
		SCEFPDNConnection:          true,
		V2XContext:                 true,
		PC5QoSParameters:           true,
	}
	isGroupedFun = func(t uint8) bool {
		mu.RLock()
		defer mu.RUnlock()
		_, ok := defaultGroupedIEMap[t]
		return ok
	}
)

// SetIsGroupedFun sets a function to check if an IE is of grouped type or not.
func SetIsGroupedFun(fun func(t uint8) bool) {
	mu.Lock()
	defer mu.Unlock()
	isGroupedFun = fun
}

// AddGroupedIEType adds IE type(s) to the defaultGroupedIEMap.
// This is useful when you want to add new IE types to the defaultGroupedIEMap.
//
// See also SetIsGroupedFun().
func AddGroupedIEType(ts ...uint8) {
	mu.Lock()
	defer mu.Unlock()
	for _, t := range ts {
		defaultGroupedIEMap[t] = true
	}
}

// IsGrouped reports whether an IE is grouped type or not.
//
// By default, this package determines if an IE is grouped type or not by checking
// if the IE type is in the defaultGroupedIEMap.
// You can change this entire behavior by calling SetIsGroupedFun(), or you can add
// new IE types to the defaultGroupedIEMap by calling AddGroupedIEType().
func (i *IE) IsGrouped() bool {
	return isGroupedFun(i.Type)
}

// Add adds variable number of IEs to a grouped IE and update length of it.
// This does nothing if the type of the IE is not grouped (no errors).
func (i *IE) Add(ies ...*IE) {
	if !i.IsGrouped() {
		return
	}

	for _, ie := range ies {
		if ie == nil {
			continue
		}
		i.ChildIEs = append(i.ChildIEs, ie)

		serialized, err := ie.Marshal()
		if err != nil {
			continue
		}
		i.Payload = append(i.Payload, serialized...)
	}
	i.SetLength()
}

// Remove removes an IE looked up by type and instance.
func (i *IE) Remove(typ, instance uint8) {
	if !i.IsGrouped() {
		return
	}

	i.Payload = nil
	newChildren := make([]*IE, len(i.ChildIEs))
	idx := 0
	for _, ie := range i.ChildIEs {
		if ie.Type == typ && ie.Instance() == instance {
			newChildren = newChildren[:len(newChildren)-1]
			continue
		}
		newChildren[idx] = ie
		idx++

		serialized, err := ie.Marshal()
		if err != nil {
			continue
		}
		i.Payload = append(i.Payload, serialized...)
	}
	i.ChildIEs = newChildren
	i.SetLength()
}

// FindByType returns IE looked up by type and instance.
//
// The program may be slower when calling this method multiple times
// because this ranges over a ChildIEs each time it is called.
func (i *IE) FindByType(typ, instance uint8) (*IE, error) {
	if !i.IsGrouped() {
		return nil, ErrInvalidType
	}

	for _, ie := range i.ChildIEs {
		if ie.Type == typ && ie.Instance() == instance {
			return ie, nil
		}
	}
	return nil, ErrIENotFound
}
