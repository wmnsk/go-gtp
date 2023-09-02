// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes ModifyBearerRequest into bytes.
//
// Deprecated: use ModifyBearerRequest.Marshal instead.
func (m *ModifyBearerRequest) Serialize() ([]byte, error) {
	log.Println("ModifyBearerRequest.Serialize is deprecated. use ModifyBearerRequest.Marshal instead")
	return m.Marshal()
}

// SerializeTo serializes ModifyBearerRequest into bytes given as b.
//
// Deprecated: use ModifyBearerRequest.MarshalTo instead.
func (m *ModifyBearerRequest) SerializeTo(b []byte) error {
	log.Println("ModifyBearerRequest.SerializeTo is deprecated. use ModifyBearerRequest.MarshalTo instead")
	return m.MarshalTo(b)
}

// DecodeModifyBearerRequest decodes bytes as ModifyBearerRequest.
//
// Deprecated: use ParseModifyBearerRequest instead.
func DecodeModifyBearerRequest(b []byte) (*ModifyBearerRequest, error) {
	log.Println("DecodeModifyBearerRequest is deprecated. use ParseModifyBearerRequest instead")
	return ParseModifyBearerRequest(b)
}

// DecodeFromBytes decodes bytes as ModifyBearerRequest.
//
// Deprecated: use ModifyBearerRequest.UnmarshalBinary instead.
func (m *ModifyBearerRequest) DecodeFromBytes(b []byte) error {
	log.Println("ModifyBearerRequest.DecodeFromBytes is deprecated. use ModifyBearerRequest.UnmarshalBinary instead")
	return m.UnmarshalBinary(b)
}

// Len returns the actual length of ModifyBearerRequest.
//
// Deprecated: use ModifyBearerRequest.MarshalLen instead.
func (m *ModifyBearerRequest) Len() int {
	log.Println("ModifyBearerRequest.Len is deprecated. use ModifyBearerRequest.MarshalLen instead")
	return m.MarshalLen()
}
