// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "log"

// Serialize serializes ModifyAccessBearersRequest into bytes.
//
// DEPRECATED: use ModifyAccessBearersRequest.Marshal instead.
func (m *ModifyAccessBearersRequest) Serialize() ([]byte, error) {
	log.Println("ModifyAccessBearersRequest.Serialize is deprecated. use ModifyAccessBearersRequest.Marshal instead")
	return m.Marshal()
}

// SerializeTo serializes ModifyAccessBearersRequest into bytes given as b.
//
// DEPRECATED: use ModifyAccessBearersRequest.MarshalTo instead.
func (m *ModifyAccessBearersRequest) SerializeTo(b []byte) error {
	log.Println("ModifyAccessBearersRequest.SerializeTo is deprecated. use ModifyAccessBearersRequest.MarshalTo instead")
	return m.MarshalTo(b)
}

// DecodeModifyAccessBearersRequest decodes bytes as ModifyAccessBearersRequest.
//
// DEPRECATED: use ParseModifyAccessBearersRequest instead.
func DecodeModifyAccessBearersRequest(b []byte) (*ModifyAccessBearersRequest, error) {
	log.Println("DecodeModifyAccessBearersRequest is deprecated. use ParseModifyAccessBearersRequest instead")
	return ParseModifyAccessBearersRequest(b)
}

// DecodeFromBytes decodes bytes as ModifyAccessBearersRequest.
//
// DEPRECATED: use ModifyAccessBearersRequest.UnmarshalBinary instead.
func (m *ModifyAccessBearersRequest) DecodeFromBytes(b []byte) error {
	log.Println("ModifyAccessBearersRequest.DecodeFromBytes is deprecated. use ModifyAccessBearersRequest.UnmarshalBinary instead")
	return m.UnmarshalBinary(b)
}

// Len returns the actual length of ModifyAccessBearersRequest.
//
// DEPRECATED: use ModifyAccessBearersRequest.MarshalLen instead.
func (m *ModifyAccessBearersRequest) Len() int {
	log.Println("ModifyAccessBearersRequest.Len is deprecated. use ModifyAccessBearersRequest.MarshalLen instead")
	return m.MarshalLen()
}
