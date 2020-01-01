// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes ModifyAccessBearersResponse into bytes.
//
// DEPRECATED: use ModifyAccessBearersResponse.Marshal instead.
func (m *ModifyAccessBearersResponse) Serialize() ([]byte, error) {
	log.Println("ModifyAccessBearersResponse.Serialize is deprecated. use ModifyAccessBearersResponse.Marshal instead")
	return m.Marshal()
}

// SerializeTo serializes ModifyAccessBearersResponse into bytes given as b.
//
// DEPRECATED: use ModifyAccessBearersResponse.MarshalTo instead.
func (m *ModifyAccessBearersResponse) SerializeTo(b []byte) error {
	log.Println("ModifyAccessBearersResponse.SerializeTo is deprecated. use ModifyAccessBearersResponse.MarshalTo instead")
	return m.MarshalTo(b)
}

// DecodeModifyAccessBearersResponse decodes bytes as ModifyAccessBearersResponse.
//
// DEPRECATED: use ParseModifyAccessBearersResponse instead.
func DecodeModifyAccessBearersResponse(b []byte) (*ModifyAccessBearersResponse, error) {
	log.Println("DecodeModifyAccessBearersResponse is deprecated. use ParseModifyAccessBearersResponse instead")
	return ParseModifyAccessBearersResponse(b)
}

// DecodeFromBytes decodes bytes as ModifyAccessBearersResponse.
//
// DEPRECATED: use ModifyAccessBearersResponse.UnmarshalBinary instead.
func (m *ModifyAccessBearersResponse) DecodeFromBytes(b []byte) error {
	log.Println("ModifyAccessBearersResponse.DecodeFromBytes is deprecated. use ModifyAccessBearersResponse.UnmarshalBinary instead")
	return m.UnmarshalBinary(b)
}

// Len returns the actual length of ModifyAccessBearersResponse.
//
// DEPRECATED: use ModifyAccessBearersResponse.MarshalLen instead.
func (m *ModifyAccessBearersResponse) Len() int {
	log.Println("ModifyAccessBearersResponse.Len is deprecated. use ModifyAccessBearersResponse.MarshalLen instead")
	return m.MarshalLen()
}
