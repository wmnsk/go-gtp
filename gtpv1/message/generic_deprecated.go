// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes Generic into bytes.
//
// Deprecated: use Generic.Marshal instead.
func (g *Generic) Serialize() ([]byte, error) {
	log.Println("Generic.Serialize is deprecated. use Generic.Marshal instead")
	return g.Marshal()
}

// SerializeTo serializes Generic into bytes given as b.
//
// Deprecated: use Generic.MarshalTo instead.
func (g *Generic) SerializeTo(b []byte) error {
	log.Println("Generic.SerializeTo is deprecated. use Generic.MarshalTo instead")
	return g.MarshalTo(b)
}

// DecodeGeneric decodes bytes as Generic.
//
// Deprecated: use ParseGeneric instead.
func DecodeGeneric(b []byte) (*Generic, error) {
	log.Println("DecodeGeneric is deprecated. use ParseGeneric instead")
	return ParseGeneric(b)
}

// DecodeFromBytes decodes bytes as Generic.
//
// Deprecated: use Generic.UnmarshalBinary instead.
func (g *Generic) DecodeFromBytes(b []byte) error {
	log.Println("Generic.DecodeFromBytes is deprecated. use Generic.UnmarshalBinary instead")
	return g.UnmarshalBinary(b)
}

// Len returns the actual length of Generic.
//
// Deprecated: use Generic.MarshalLen instead.
func (g *Generic) Len() int {
	log.Println("Generic.Len is deprecated. use Generic.MarshalLen instead")
	return g.MarshalLen()
}
