// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes ModifyBearerRequest into bytes.
//
// DEPRECATED: use ModifyBearerRequest.Marshal instead.
func (m *ModifyBearerRequest) Serialize() ([]byte, error) {
	log.Println("ModifyBearerRequest.Serialize is deprecated. use ModifyBearerRequest.Marshal instead")
	return m.Marshal()
}

// SerializeTo serializes ModifyBearerRequest into bytes given as b.
//
// DEPRECATED: use ModifyBearerRequest.MarshalTo instead.
func (m *ModifyBearerRequest) SerializeTo(b []byte) error {
	log.Println("ModifyBearerRequest.SerializeTo is deprecated. use ModifyBearerRequest.MarshalTo instead")
	return m.MarshalTo(b)
}

// DecodeModifyBearerRequest decodes bytes as ModifyBearerRequest.
//
// DEPRECATED: use ParseModifyBearerRequest instead.
func DecodeModifyBearerRequest(b []byte) (*ModifyBearerRequest, error) {
	log.Println("DecodeModifyBearerRequest is deprecated. use ParseModifyBearerRequest instead")
	return ParseModifyBearerRequest(b)
}

// DecodeFromBytes decodes bytes as ModifyBearerRequest.
//
// DEPRECATED: use ModifyBearerRequest.UnmarshalBinary instead.
func (m *ModifyBearerRequest) DecodeFromBytes(b []byte) error {
	log.Println("ModifyBearerRequest.DecodeFromBytes is deprecated. use ModifyBearerRequest.UnmarshalBinary instead")
	return m.UnmarshalBinary(b)
}

// Len returns the actual length of ModifyBearerRequest.
//
// DEPRECATED: use ModifyBearerRequest.MarshalLen instead.
func (m *ModifyBearerRequest) Len() int {
	log.Println("ModifyBearerRequest.Len is deprecated. use ModifyBearerRequest.MarshalLen instead")
	return m.MarshalLen()
}
