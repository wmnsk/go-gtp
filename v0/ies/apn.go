// Copyright 2019-2020 go-gtp. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"io"
	"strings"
)

// NewAccessPointName creates a new AccessPointName IE.
func NewAccessPointName(apn string) *IE {
	i := New(AccessPointName, make([]byte, len(apn)+1))
	var offset = 0
	for _, label := range strings.Split(apn, ".") {
		l := len(label)
		i.Payload[offset] = uint8(l)
		copy(i.Payload[offset+1:], []byte(label))
		offset += l + 1
	}

	return i
}

// AccessPointName returns AccessPointName in string if type of IE matches.
func (i *IE) AccessPointName() (string, error) {
	if i.Type != AccessPointName {
		return "", &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return "", io.ErrUnexpectedEOF
	}

	var (
		apn    []string
		offset int
	)

	max := len(i.Payload)
	for {
		if offset >= max {
			break
		}
		l := int(i.Payload[offset])
		if offset+l+1 >= max {
			return "", io.ErrUnexpectedEOF
		}
		apn = append(apn, string(i.Payload[offset+1:offset+l+1]))
		offset += l + 1
	}

	return strings.Join(apn, "."), nil
}

// MustAccessPointName returns AccessPointName in string if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustAccessPointName() string {
	v, _ := i.AccessPointName()
	return v
}
