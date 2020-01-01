// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes Header into bytes.
//
// DEPRECATED: use Header.Marshal instead.
func (h *Header) Serialize() ([]byte, error) {
	log.Println("Header.Serialize is deprecated. use Header.Marshal instead")
	return h.Marshal()
}

// SerializeTo serializes Header into bytes given as b.
//
// DEPRECATED: use Header.MarshalTo instead.
func (h *Header) SerializeTo(b []byte) error {
	log.Println("Header.SerializeTo is deprecated. use Header.MarshalTo instead")
	return h.MarshalTo(b)
}

// DecodeHeader decodes bytes as Header.
//
// DEPRECATED: use ParseHeader instead.
func DecodeHeader(b []byte) (*Header, error) {
	log.Println("DecodeHeader is deprecated. use ParseHeader instead")
	return ParseHeader(b)
}

// DecodeFromBytes decodes bytes as Header.
//
// DEPRECATED: use Header.UnmarshalBinary instead.
func (h *Header) DecodeFromBytes(b []byte) error {
	log.Println("Header.DecodeFromBytes is deprecated. use Header.UnmarshalBinary instead")
	return h.UnmarshalBinary(b)
}

// Len returns the actual length of Header.
//
// DEPRECATED: use Header.MarshalLen instead.
func (h *Header) Len() int {
	log.Println("Header.Len is deprecated. use Header.MarshalLen instead")
	return h.MarshalLen()
}
