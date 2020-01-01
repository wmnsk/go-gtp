// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes CreateSessionResponse into bytes.
//
// DEPRECATED: use CreateSessionResponse.Marshal instead.
func (c *CreateSessionResponse) Serialize() ([]byte, error) {
	log.Println("CreateSessionResponse.Serialize is deprecated. use CreateSessionResponse.Marshal instead")
	return c.Marshal()
}

// SerializeTo serializes CreateSessionResponse into bytes given as b.
//
// DEPRECATED: use CreateSessionResponse.MarshalTo instead.
func (c *CreateSessionResponse) SerializeTo(b []byte) error {
	log.Println("CreateSessionResponse.SerializeTo is deprecated. use CreateSessionResponse.MarshalTo instead")
	return c.MarshalTo(b)
}

// DecodeCreateSessionResponse decodes bytes as CreateSessionResponse.
//
// DEPRECATED: use ParseCreateSessionResponse instead.
func DecodeCreateSessionResponse(b []byte) (*CreateSessionResponse, error) {
	log.Println("DecodeCreateSessionResponse is deprecated. use ParseCreateSessionResponse instead")
	return ParseCreateSessionResponse(b)
}

// DecodeFromBytes decodes bytes as CreateSessionResponse.
//
// DEPRECATED: use CreateSessionResponse.UnmarshalBinary instead.
func (c *CreateSessionResponse) DecodeFromBytes(b []byte) error {
	log.Println("CreateSessionResponse.DecodeFromBytes is deprecated. use CreateSessionResponse.UnmarshalBinary instead")
	return c.UnmarshalBinary(b)
}

// Len returns the actual length of CreateSessionResponse.
//
// DEPRECATED: use CreateSessionResponse.MarshalLen instead.
func (c *CreateSessionResponse) Len() int {
	log.Println("CreateSessionResponse.Len is deprecated. use CreateSessionResponse.MarshalLen instead")
	return c.MarshalLen()
}
