// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewFullyQualifiedDomainName creates a new FullyQualifiedDomainName IE.
func NewFullyQualifiedDomainName(fqdn string) *IE {
	return newStringIE(FullyQualifiedDomainName, fqdn)
}

// FullyQualifiedDomainName returns FullyQualifiedDomainName in string if the type of IE matches.
func (i *IE) FullyQualifiedDomainName() string {
	if i.Type != FullyQualifiedDomainName {
		return ""
	}
	if len(i.Payload) == 0 {
		return ""
	}

	return string(i.Payload)
}
