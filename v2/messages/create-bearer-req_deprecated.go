// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes CreateBearerRequest into bytes.
//
// DEPRECATED: use CreateBearerRequest.Marshal instead.
func (c *CreateBearerRequest) Serialize() ([]byte, error) {
	log.Println("CreateBearerRequest.Serialize is deprecated. use CreateBearerRequest.Marshal instead")
	return c.Marshal()
}

// SerializeTo serializes CreateBearerRequest into bytes given as b.
//
// DEPRECATED: use CreateBearerRequest.MarshalTo instead.
func (c *CreateBearerRequest) SerializeTo(b []byte) error {
	log.Println("CreateBearerRequest.SerializeTo is deprecated. use CreateBearerRequest.MarshalTo instead")
	return c.MarshalTo(b)
}

// DecodeCreateBearerRequest decodes bytes as CreateBearerRequest.
//
// DEPRECATED: use ParseCreateBearerRequest instead.
func DecodeCreateBearerRequest(b []byte) (*CreateBearerRequest, error) {
	log.Println("DecodeCreateBearerRequest is deprecated. use ParseCreateBearerRequest instead")
	return ParseCreateBearerRequest(b)
}

// DecodeFromBytes decodes bytes as CreateBearerRequest.
//
// DEPRECATED: use CreateBearerRequest.UnmarshalBinary instead.
func (c *CreateBearerRequest) DecodeFromBytes(b []byte) error {
	log.Println("CreateBearerRequest.DecodeFromBytes is deprecated. use CreateBearerRequest.UnmarshalBinary instead")
	return c.UnmarshalBinary(b)
}

// Len returns the actual length of CreateBearerRequest.
//
// DEPRECATED: use CreateBearerRequest.MarshalLen instead.
func (c *CreateBearerRequest) Len() int {
	log.Println("CreateBearerRequest.Len is deprecated. use CreateBearerRequest.MarshalLen instead")
	return c.MarshalLen()
}
