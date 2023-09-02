// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes UpdatePDPContextRequest into bytes.
//
// Deprecated: use UpdatePDPContextRequest.Marshal instead.
func (u *UpdatePDPContextRequest) Serialize() ([]byte, error) {
	log.Println("UpdatePDPContextRequest.Serialize is deprecated. use UpdatePDPContextRequest.Marshal instead")
	return u.Marshal()
}

// SerializeTo serializes UpdatePDPContextRequest into bytes given as b.
//
// Deprecated: use UpdatePDPContextRequest.MarshalTo instead.
func (u *UpdatePDPContextRequest) SerializeTo(b []byte) error {
	log.Println("UpdatePDPContextRequest.SerializeTo is deprecated. use UpdatePDPContextRequest.MarshalTo instead")
	return u.MarshalTo(b)
}

// DecodeUpdatePDPContextRequest decodes bytes as UpdatePDPContextRequest.
//
// Deprecated: use ParseUpdatePDPContextRequest instead.
func DecodeUpdatePDPContextRequest(b []byte) (*UpdatePDPContextRequest, error) {
	log.Println("DecodeUpdatePDPContextRequest is deprecated. use ParseUpdatePDPContextRequest instead")
	return ParseUpdatePDPContextRequest(b)
}

// DecodeFromBytes decodes bytes as UpdatePDPContextRequest.
//
// Deprecated: use UpdatePDPContextRequest.UnmarshalBinary instead.
func (u *UpdatePDPContextRequest) DecodeFromBytes(b []byte) error {
	log.Println("UpdatePDPContextRequest.DecodeFromBytes is deprecated. use UpdatePDPContextRequest.UnmarshalBinary instead")
	return u.UnmarshalBinary(b)
}

// Len returns the actual length of UpdatePDPContextRequest.
//
// Deprecated: use UpdatePDPContextRequest.MarshalLen instead.
func (u *UpdatePDPContextRequest) Len() int {
	log.Println("UpdatePDPContextRequest.Len is deprecated. use UpdatePDPContextRequest.MarshalLen instead")
	return u.MarshalLen()
}
