// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes TPDU into bytes.
//
// Deprecated: use TPDU.Marshal instead.
func (t *TPDU) Serialize() ([]byte, error) {
	log.Println("TPDU.Serialize is deprecated. use TPDU.Marshal instead")
	return t.Marshal()
}

// SerializeTo serializes TPDU into bytes given as b.
//
// Deprecated: use TPDU.MarshalTo instead.
func (t *TPDU) SerializeTo(b []byte) error {
	log.Println("TPDU.SerializeTo is deprecated. use TPDU.MarshalTo instead")
	return t.MarshalTo(b)
}

// DecodeTPDU decodes bytes as TPDU.
//
// Deprecated: use ParseTPDU instead.
func DecodeTPDU(b []byte) (*TPDU, error) {
	log.Println("DecodeTPDU is deprecated. use ParseTPDU instead")
	return ParseTPDU(b)
}

// DecodeFromBytes decodes bytes as TPDU.
//
// Deprecated: use TPDU.UnmarshalBinary instead.
func (t *TPDU) DecodeFromBytes(b []byte) error {
	log.Println("TPDU.DecodeFromBytes is deprecated. use TPDU.UnmarshalBinary instead")
	return t.UnmarshalBinary(b)
}

// Len returns the actual length of TPDU.
//
// Deprecated: use TPDU.MarshalLen instead.
func (t *TPDU) Len() int {
	log.Println("TPDU.Len is deprecated. use TPDU.MarshalLen instead")
	return t.MarshalLen()
}
