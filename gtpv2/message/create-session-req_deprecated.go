// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes CreateSessionRequest into bytes.
//
// Deprecated: use CreateSessionRequest.Marshal instead.
func (c *CreateSessionRequest) Serialize() ([]byte, error) {
	log.Println("CreateSessionRequest.Serialize is deprecated. use CreateSessionRequest.Marshal instead")
	return c.Marshal()
}

// SerializeTo serializes CreateSessionRequest into bytes given as b.
//
// Deprecated: use CreateSessionRequest.MarshalTo instead.
func (c *CreateSessionRequest) SerializeTo(b []byte) error {
	log.Println("CreateSessionRequest.SerializeTo is deprecated. use CreateSessionRequest.MarshalTo instead")
	return c.MarshalTo(b)
}

// DecodeCreateSessionRequest decodes bytes as CreateSessionRequest.
//
// Deprecated: use ParseCreateSessionRequest instead.
func DecodeCreateSessionRequest(b []byte) (*CreateSessionRequest, error) {
	log.Println("DecodeCreateSessionRequest is deprecated. use ParseCreateSessionRequest instead")
	return ParseCreateSessionRequest(b)
}

// DecodeFromBytes decodes bytes as CreateSessionRequest.
//
// Deprecated: use CreateSessionRequest.UnmarshalBinary instead.
func (c *CreateSessionRequest) DecodeFromBytes(b []byte) error {
	log.Println("CreateSessionRequest.DecodeFromBytes is deprecated. use CreateSessionRequest.UnmarshalBinary instead")
	return c.UnmarshalBinary(b)
}

// Len returns the actual length of CreateSessionRequest.
//
// Deprecated: use CreateSessionRequest.MarshalLen instead.
func (c *CreateSessionRequest) Len() int {
	log.Println("CreateSessionRequest.Len is deprecated. use CreateSessionRequest.MarshalLen instead")
	return c.MarshalLen()
}
