// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes CreateBearerRequest into bytes.
//
// Deprecated: use CreateBearerRequest.Marshal instead.
func (c *CreateBearerRequest) Serialize() ([]byte, error) {
	log.Println("CreateBearerRequest.Serialize is deprecated. use CreateBearerRequest.Marshal instead")
	return c.Marshal()
}

// SerializeTo serializes CreateBearerRequest into bytes given as b.
//
// Deprecated: use CreateBearerRequest.MarshalTo instead.
func (c *CreateBearerRequest) SerializeTo(b []byte) error {
	log.Println("CreateBearerRequest.SerializeTo is deprecated. use CreateBearerRequest.MarshalTo instead")
	return c.MarshalTo(b)
}

// DecodeCreateBearerRequest decodes bytes as CreateBearerRequest.
//
// Deprecated: use ParseCreateBearerRequest instead.
func DecodeCreateBearerRequest(b []byte) (*CreateBearerRequest, error) {
	log.Println("DecodeCreateBearerRequest is deprecated. use ParseCreateBearerRequest instead")
	return ParseCreateBearerRequest(b)
}

// DecodeFromBytes decodes bytes as CreateBearerRequest.
//
// Deprecated: use CreateBearerRequest.UnmarshalBinary instead.
func (c *CreateBearerRequest) DecodeFromBytes(b []byte) error {
	log.Println("CreateBearerRequest.DecodeFromBytes is deprecated. use CreateBearerRequest.UnmarshalBinary instead")
	return c.UnmarshalBinary(b)
}

// Len returns the actual length of CreateBearerRequest.
//
// Deprecated: use CreateBearerRequest.MarshalLen instead.
func (c *CreateBearerRequest) Len() int {
	log.Println("CreateBearerRequest.Len is deprecated. use CreateBearerRequest.MarshalLen instead")
	return c.MarshalLen()
}
