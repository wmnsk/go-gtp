// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes ModifyBearerResponse into bytes.
//
// Deprecated: use ModifyBearerResponse.Marshal instead.
func (m *ModifyBearerResponse) Serialize() ([]byte, error) {
	log.Println("ModifyBearerResponse.Serialize is deprecated. use ModifyBearerResponse.Marshal instead")
	return m.Marshal()
}

// SerializeTo serializes ModifyBearerResponse into bytes given as b.
//
// Deprecated: use ModifyBearerResponse.MarshalTo instead.
func (m *ModifyBearerResponse) SerializeTo(b []byte) error {
	log.Println("ModifyBearerResponse.SerializeTo is deprecated. use ModifyBearerResponse.MarshalTo instead")
	return m.MarshalTo(b)
}

// DecodeModifyBearerResponse decodes bytes as ModifyBearerResponse.
//
// Deprecated: use ParseModifyBearerResponse instead.
func DecodeModifyBearerResponse(b []byte) (*ModifyBearerResponse, error) {
	log.Println("DecodeModifyBearerResponse is deprecated. use ParseModifyBearerResponse instead")
	return ParseModifyBearerResponse(b)
}

// DecodeFromBytes decodes bytes as ModifyBearerResponse.
//
// Deprecated: use ModifyBearerResponse.UnmarshalBinary instead.
func (m *ModifyBearerResponse) DecodeFromBytes(b []byte) error {
	log.Println("ModifyBearerResponse.DecodeFromBytes is deprecated. use ModifyBearerResponse.UnmarshalBinary instead")
	return m.UnmarshalBinary(b)
}

// Len returns the actual length of ModifyBearerResponse.
//
// Deprecated: use ModifyBearerResponse.MarshalLen instead.
func (m *ModifyBearerResponse) Len() int {
	log.Println("ModifyBearerResponse.Len is deprecated. use ModifyBearerResponse.MarshalLen instead")
	return m.MarshalLen()
}
