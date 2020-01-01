// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes CreateSessionRequest into bytes.
//
// DEPRECATED: use CreateSessionRequest.Marshal instead.
func (c *CreateSessionRequest) Serialize() ([]byte, error) {
	log.Println("CreateSessionRequest.Serialize is deprecated. use CreateSessionRequest.Marshal instead")
	return c.Marshal()
}

// SerializeTo serializes CreateSessionRequest into bytes given as b.
//
// DEPRECATED: use CreateSessionRequest.MarshalTo instead.
func (c *CreateSessionRequest) SerializeTo(b []byte) error {
	log.Println("CreateSessionRequest.SerializeTo is deprecated. use CreateSessionRequest.MarshalTo instead")
	return c.MarshalTo(b)
}

// DecodeCreateSessionRequest decodes bytes as CreateSessionRequest.
//
// DEPRECATED: use ParseCreateSessionRequest instead.
func DecodeCreateSessionRequest(b []byte) (*CreateSessionRequest, error) {
	log.Println("DecodeCreateSessionRequest is deprecated. use ParseCreateSessionRequest instead")
	return ParseCreateSessionRequest(b)
}

// DecodeFromBytes decodes bytes as CreateSessionRequest.
//
// DEPRECATED: use CreateSessionRequest.UnmarshalBinary instead.
func (c *CreateSessionRequest) DecodeFromBytes(b []byte) error {
	log.Println("CreateSessionRequest.DecodeFromBytes is deprecated. use CreateSessionRequest.UnmarshalBinary instead")
	return c.UnmarshalBinary(b)
}

// Len returns the actual length of CreateSessionRequest.
//
// DEPRECATED: use CreateSessionRequest.MarshalLen instead.
func (c *CreateSessionRequest) Len() int {
	log.Println("CreateSessionRequest.Len is deprecated. use CreateSessionRequest.MarshalLen instead")
	return c.MarshalLen()
}
