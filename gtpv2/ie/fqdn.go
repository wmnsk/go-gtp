// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"strings"
)

// NewFullyQualifiedDomainName creates a new FullyQualifiedDomainName IE.
func NewFullyQualifiedDomainName(fqdn string) *IE {
	i := New(FullyQualifiedDomainName, 0x00, make([]byte, len(fqdn)+1))
	var offset = 0
	for _, label := range strings.Split(fqdn, ".") {
		l := len(label)
		i.Payload[offset] = uint8(l)
		copy(i.Payload[offset+1:], label)
		offset += l + 1
	}

	return i
}

// FullyQualifiedDomainName returns FullyQualifiedDomainName in string if the type of IE matches.
func (i *IE) FullyQualifiedDomainName() (string, error) {
	if i.Type != FullyQualifiedDomainName {
		return "", &InvalidTypeError{Type: i.Type}
	}

	var (
		fqdn   []string
		offset int
	)
	max := len(i.Payload)
	for {
		if offset >= max {
			break
		}
		l := int(i.Payload[offset])
		if offset+l+1 > max {
			break
		}
		fqdn = append(fqdn, string(i.Payload[offset+1:offset+l+1]))
		offset += l + 1
	}

	return strings.Join(fqdn, "."), nil
}

// MustFullyQualifiedDomainName returns FullyQualifiedDomainName in string, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustFullyQualifiedDomainName() string {
	v, _ := i.FullyQualifiedDomainName()
	return v
}
