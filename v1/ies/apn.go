// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "strings"

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
func (i *IE) AccessPointName() string {
	if i.Type != AccessPointName {
		return ""
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
		apn = append(apn, string(i.Payload[offset+1:offset+l+1]))
		offset += l + 1
	}

	return strings.Join(apn, ".")
}
