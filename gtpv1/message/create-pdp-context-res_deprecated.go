// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes CreatePDPContextResponse into bytes.
//
// Deprecated: use CreatePDPContextResponse.Marshal instead.
func (c *CreatePDPContextResponse) Serialize() ([]byte, error) {
	log.Println("CreatePDPContextResponse.Serialize is deprecated. use CreatePDPContextResponse.Marshal instead")
	return c.Marshal()
}

// SerializeTo serializes CreatePDPContextResponse into bytes given as b.
//
// Deprecated: use CreatePDPContextResponse.MarshalTo instead.
func (c *CreatePDPContextResponse) SerializeTo(b []byte) error {
	log.Println("CreatePDPContextResponse.SerializeTo is deprecated. use CreatePDPContextResponse.MarshalTo instead")
	return c.MarshalTo(b)
}

// DecodeCreatePDPContextResponse decodes bytes as CreatePDPContextResponse.
//
// Deprecated: use ParseCreatePDPContextResponse instead.
func DecodeCreatePDPContextResponse(b []byte) (*CreatePDPContextResponse, error) {
	log.Println("DecodeCreatePDPContextResponse is deprecated. use ParseCreatePDPContextResponse instead")
	return ParseCreatePDPContextResponse(b)
}

// DecodeFromBytes decodes bytes as CreatePDPContextResponse.
//
// Deprecated: use CreatePDPContextResponse.UnmarshalBinary instead.
func (c *CreatePDPContextResponse) DecodeFromBytes(b []byte) error {
	log.Println("CreatePDPContextResponse.DecodeFromBytes is deprecated. use CreatePDPContextResponse.UnmarshalBinary instead")
	return c.UnmarshalBinary(b)
}

// Len returns the actual length of CreatePDPContextResponse.
//
// Deprecated: use CreatePDPContextResponse.MarshalLen instead.
func (c *CreatePDPContextResponse) Len() int {
	log.Println("CreatePDPContextResponse.Len is deprecated. use CreatePDPContextResponse.MarshalLen instead")
	return c.MarshalLen()
}
