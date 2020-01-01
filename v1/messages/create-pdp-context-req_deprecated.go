// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes CreatePDPContextRequest into bytes.
//
// DEPRECATED: use CreatePDPContextRequest.Marshal instead.
func (c *CreatePDPContextRequest) Serialize() ([]byte, error) {
	log.Println("CreatePDPContextRequest.Serialize is deprecated. use CreatePDPContextRequest.Marshal instead")
	return c.Marshal()
}

// SerializeTo serializes CreatePDPContextRequest into bytes given as b.
//
// DEPRECATED: use CreatePDPContextRequest.MarshalTo instead.
func (c *CreatePDPContextRequest) SerializeTo(b []byte) error {
	log.Println("CreatePDPContextRequest.SerializeTo is deprecated. use CreatePDPContextRequest.MarshalTo instead")
	return c.MarshalTo(b)
}

// DecodeCreatePDPContextRequest decodes bytes as CreatePDPContextRequest.
//
// DEPRECATED: use ParseCreatePDPContextRequest instead.
func DecodeCreatePDPContextRequest(b []byte) (*CreatePDPContextRequest, error) {
	log.Println("DecodeCreatePDPContextRequest is deprecated. use ParseCreatePDPContextRequest instead")
	return ParseCreatePDPContextRequest(b)
}

// DecodeFromBytes decodes bytes as CreatePDPContextRequest.
//
// DEPRECATED: use CreatePDPContextRequest.UnmarshalBinary instead.
func (c *CreatePDPContextRequest) DecodeFromBytes(b []byte) error {
	log.Println("CreatePDPContextRequest.DecodeFromBytes is deprecated. use CreatePDPContextRequest.UnmarshalBinary instead")
	return c.UnmarshalBinary(b)
}

// Len returns the actual length of CreatePDPContextRequest.
//
// DEPRECATED: use CreatePDPContextRequest.MarshalLen instead.
func (c *CreatePDPContextRequest) Len() int {
	log.Println("CreatePDPContextRequest.Len is deprecated. use CreatePDPContextRequest.MarshalLen instead")
	return c.MarshalLen()
}
