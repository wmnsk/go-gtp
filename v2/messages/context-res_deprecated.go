// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes ContextResponse into bytes.
//
// DEPRECATED: use ContextResponse.Marshal instead.
func (c *ContextResponse) Serialize() ([]byte, error) {
	log.Println("ContextResponse.Serialize is deprecated. use ContextResponse.Marshal instead")
	return c.Marshal()
}

// SerializeTo serializes ContextResponse into bytes given as b.
//
// DEPRECATED: use ContextResponse.MarshalTo instead.
func (c *ContextResponse) SerializeTo(b []byte) error {
	log.Println("ContextResponse.SerializeTo is deprecated. use ContextResponse.MarshalTo instead")
	return c.MarshalTo(b)
}

// DecodeContextResponse decodes bytes as ContextResponse.
//
// DEPRECATED: use ParseContextResponse instead.
func DecodeContextResponse(b []byte) (*ContextResponse, error) {
	log.Println("DecodeContextResponse is deprecated. use ParseContextResponse instead")
	return ParseContextResponse(b)
}

// DecodeFromBytes decodes bytes as ContextResponse.
//
// DEPRECATED: use ContextResponse.UnmarshalBinary instead.
func (c *ContextResponse) DecodeFromBytes(b []byte) error {
	log.Println("ContextResponse.DecodeFromBytes is deprecated. use ContextResponse.UnmarshalBinary instead")
	return c.UnmarshalBinary(b)
}

// Len returns the actual length of ContextResponse.
//
// DEPRECATED: use ContextResponse.MarshalLen instead.
func (c *ContextResponse) Len() int {
	log.Println("ContextResponse.Len is deprecated. use ContextResponse.MarshalLen instead")
	return c.MarshalLen()
}
