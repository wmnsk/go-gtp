// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes CreateBearerResponse into bytes.
//
// Deprecated: use CreateBearerResponse.Marshal instead.
func (c *CreateBearerResponse) Serialize() ([]byte, error) {
	log.Println("CreateBearerResponse.Serialize is deprecated. use CreateBearerResponse.Marshal instead")
	return c.Marshal()
}

// SerializeTo serializes CreateBearerResponse into bytes given as b.
//
// Deprecated: use CreateBearerResponse.MarshalTo instead.
func (c *CreateBearerResponse) SerializeTo(b []byte) error {
	log.Println("CreateBearerResponse.SerializeTo is deprecated. use CreateBearerResponse.MarshalTo instead")
	return c.MarshalTo(b)
}

// DecodeCreateBearerResponse decodes bytes as CreateBearerResponse.
//
// Deprecated: use ParseCreateBearerResponse instead.
func DecodeCreateBearerResponse(b []byte) (*CreateBearerResponse, error) {
	log.Println("DecodeCreateBearerResponse is deprecated. use ParseCreateBearerResponse instead")
	return ParseCreateBearerResponse(b)
}

// DecodeFromBytes decodes bytes as CreateBearerResponse.
//
// Deprecated: use CreateBearerResponse.UnmarshalBinary instead.
func (c *CreateBearerResponse) DecodeFromBytes(b []byte) error {
	log.Println("CreateBearerResponse.DecodeFromBytes is deprecated. use CreateBearerResponse.UnmarshalBinary instead")
	return c.UnmarshalBinary(b)
}

// Len returns the actual length of CreateBearerResponse.
//
// Deprecated: use CreateBearerResponse.MarshalLen instead.
func (c *CreateBearerResponse) Len() int {
	log.Println("CreateBearerResponse.Len is deprecated. use CreateBearerResponse.MarshalLen instead")
	return c.MarshalLen()
}
