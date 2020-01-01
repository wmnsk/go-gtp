// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

/*

// NewRANNASCause creates a new RANNASCause IE.
//
// The cause parameter is set to 0xff automatically if the proto is not ProtoTypeS1AP.
func NewRANNASCause(proto, cause uint8, value uint64) *IE {
	i := New(RANNASCause, 0x00, nil)
	i.Payload[0] = proto

	switch proto {
	case ProtoTypeS1APCause:
		i.Payload[1] = cause
	case ProtoTypeEMMCause, ProtoTypeESMCause:
		i.Payload[1] = 0xff
		i.Payload[2:]
	case ProtoTypeDiameterCause:
		i.Payload[1] = 0xff
		i.Payload[2:]
	case ProtoTypeIKEv2Cause:
		i.Payload[1] = 0xff
		i.Payload[2:]
	default:
		i.Payload[1] = cause
		binary.BigEndian.PutUint64(i.Payload, value)
	}

	return i
}

*/
