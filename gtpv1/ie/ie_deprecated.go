// Copyright 2019-2022 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "log"

// Serialize serializes IE into bytes.
//
// Deprecated: use IE.Marshal instead.
func (i *IE) Serialize() ([]byte, error) {
	log.Println("IE.Serialize is deprecated. use IE.Marshal instead")
	return i.Marshal()
}

// SerializeTo serializes IE into bytes given as b.
//
// Deprecated: use IE.MarshalTo instead.
func (i *IE) SerializeTo(b []byte) error {
	log.Println("IE.SerializeTo is deprecated. use IE.MarshalTo instead")
	return i.MarshalTo(b)
}

// Decode decodes bytes as IE.
//
// Deprecated: use Parse instead.
func Decode(b []byte) (*IE, error) {
	log.Println("Decode is deprecated. use Parse instead")
	return Parse(b)
}

// DecodeFromBytes decodes bytes as IE.
//
// Deprecated: use IE.UnmarshalBinary instead.
func (i *IE) DecodeFromBytes(b []byte) error {
	log.Println("IE.DecodeFromBytes is deprecated. use IE.UnmarshalBinary instead")
	return i.UnmarshalBinary(b)
}

// Len returns the actual length of IE.
//
// Deprecated: use IE.MarshalLen instead.
func (i *IE) Len() int {
	log.Println("IE.Len is deprecated. use IE.MarshalLen instead")
	return i.MarshalLen()
}
