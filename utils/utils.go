// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Package utils provides some utilities which might be useful specifically for GTP(or other telco protocols).
package utils

import (
	"encoding/binary"
	"encoding/hex"
)

// StrToSwappedBytes returns swapped bits from a byte.
// It is used for some values where some values are represented in swapped format.
//
// The second parameter is the hex character(0-f) to fill the last digit when
// handling a odd number. "f" is used In most cases.
func StrToSwappedBytes(s, filler string) ([]byte, error) {
	var raw []byte
	var err error
	if len(s)%2 == 0 {
		raw, err = hex.DecodeString(s)
	} else {
		raw, err = hex.DecodeString(s + filler)
	}
	if err != nil {
		return nil, err
	}

	return swap(raw), nil
}

// SwappedBytesToStr decodes raw swapped bytes into string.
// It is used for some values where some values are represented in swapped format.
//
// The second parameter is to decide whether to cut the last digit or not.
func SwappedBytesToStr(raw []byte, cutLastDigit bool) string {
	if len(raw) == 0 {
		return ""
	}

	s := hex.EncodeToString(swap(raw))
	if cutLastDigit {
		s = s[:len(s)-1]
	}

	return s
}

func swap(raw []byte) []byte {
	var swapped []byte
	for n := range raw {
		t := ((raw[n] >> 4) & 0xf) + ((raw[n] << 4) & 0xf0)
		swapped = append(swapped, t)
	}
	return swapped
}

// Uint24To32 converts 24bits-length []byte value into the uint32 with 8bits of zeros as prefix.
// This function is used for the fields with 3 octets.
func Uint24To32(b []byte) uint32 {
	if len(b) != 3 {
		return 0
	}
	return uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2])
}

// Uint32To24 converts the uint32 value into 24bits-length []byte. The values in 25-32 bit are cut off.
// This function is used for the fields with 3 octets.
func Uint32To24(n uint32) []byte {
	return []byte{uint8(n >> 16), uint8(n >> 8), uint8(n)}
}

// Uint40To64 converts 40bits-length []byte value into the uint64 with 8bits of zeros as prefix.
// This function is used for the fields with 3 octets.
func Uint40To64(b []byte) uint64 {
	if len(b) != 5 {
		return 0
	}
	return uint64(b[0])<<32 | uint64(b[1])<<24 | uint64(b[2])<<16 | uint64(b[3])<<8 | uint64(b[4])
}

// Uint64To40 converts the uint64 value into 40bits-length []byte. The values in 25-64 bit are cut off.
// This function is used for the fields with 3 octets.
func Uint64To40(n uint64) []byte {
	return []byte{uint8(n >> 32), uint8(n >> 24), uint8(n >> 16), uint8(n >> 8), uint8(n)}
}

// EncodePLMN encodes MCC and MNC as BCD-encoded bytes.
func EncodePLMN(mcc, mnc string) ([]byte, error) {
	c, err := StrToSwappedBytes(mcc, "f")
	if err != nil {
		return nil, err
	}
	n, err := StrToSwappedBytes(mnc, "f")
	if err != nil {
		return nil, err
	}

	// 2-digit
	b := make([]byte, 3)
	if len(mnc) == 2 {
		b = append(c, n...)
		return b, nil
	}

	// 3-digit
	b[0] = c[0]
	b[1] = (c[1] & 0x0f) | (n[1] << 4 & 0xf0)
	b[2] = n[0]

	return b, nil
}

// DecodePLMN decodes BCD-encoded bytes into MCC and MNC.
func DecodePLMN(b []byte) (mcc, mnc string, err error) {
	raw := hex.EncodeToString(b)
	mcc = string(raw[1]) + string(raw[0]) + string(raw[3])
	mnc = string(raw[5]) + string(raw[4])
	if string(raw[2]) != "f" {
		mnc += string(raw[2])
	}

	return
}

// ParseECI decodes ECI uint32 into e-NodeB ID and Cell ID.
func ParseECI(eci uint32) (enbID uint32, cellID uint8, err error) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, eci)
	cellID = uint8(buf[3])
	enbID = binary.BigEndian.Uint32([]byte{0, buf[0], buf[1], buf[2]})
	return
}
