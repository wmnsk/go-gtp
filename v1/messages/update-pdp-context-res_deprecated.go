// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes UpdatePDPContextResponse into bytes.
//
// DEPRECATED: use UpdatePDPContextResponse.Marshal instead.
func (u *UpdatePDPContextResponse) Serialize() ([]byte, error) {
	log.Println("UpdatePDPContextResponse.Serialize is deprecated. use UpdatePDPContextResponse.Marshal instead")
	return u.Marshal()
}

// SerializeTo serializes UpdatePDPContextResponse into bytes given as b.
//
// DEPRECATED: use UpdatePDPContextResponse.MarshalTo instead.
func (u *UpdatePDPContextResponse) SerializeTo(b []byte) error {
	log.Println("UpdatePDPContextResponse.SerializeTo is deprecated. use UpdatePDPContextResponse.MarshalTo instead")
	return u.MarshalTo(b)
}

// DecodeUpdatePDPContextResponse decodes bytes as UpdatePDPContextResponse.
//
// DEPRECATED: use ParseUpdatePDPContextResponse instead.
func DecodeUpdatePDPContextResponse(b []byte) (*UpdatePDPContextResponse, error) {
	log.Println("DecodeUpdatePDPContextResponse is deprecated. use ParseUpdatePDPContextResponse instead")
	return ParseUpdatePDPContextResponse(b)
}

// DecodeFromBytes decodes bytes as UpdatePDPContextResponse.
//
// DEPRECATED: use UpdatePDPContextResponse.UnmarshalBinary instead.
func (u *UpdatePDPContextResponse) DecodeFromBytes(b []byte) error {
	log.Println("UpdatePDPContextResponse.DecodeFromBytes is deprecated. use UpdatePDPContextResponse.UnmarshalBinary instead")
	return u.UnmarshalBinary(b)
}

// Len returns the actual length of UpdatePDPContextResponse.
//
// DEPRECATED: use UpdatePDPContextResponse.MarshalLen instead.
func (u *UpdatePDPContextResponse) Len() int {
	log.Println("UpdatePDPContextResponse.Len is deprecated. use UpdatePDPContextResponse.MarshalLen instead")
	return u.MarshalLen()
}
