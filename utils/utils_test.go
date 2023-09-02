// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package utils_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-gtp/utils"
)

func TestBCDEncoding(t *testing.T) {
	cases := []struct {
		description string
		str         string
		bytes       []byte
	}{
		{
			"imsi",
			"123451234567890",
			[]byte{0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0},
		},
	}

	for _, c := range cases {
		t.Run("Str2Bytes/"+c.description, func(t *testing.T) {
			swapped, err := utils.StrToSwappedBytes(c.str, "f")
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(swapped, c.bytes); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("Bytes2Str/"+c.description, func(t *testing.T) {
			str := utils.SwappedBytesToStr(c.bytes, true)

			if diff := cmp.Diff(str, c.str); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestUint32And24(t *testing.T) {
	cases := []struct {
		description string
		u24         []byte
		u32         uint32
	}{
		{
			"Normal",
			[]byte{0xff, 0xff, 0xff},
			0x00ffffff,
		},
	}

	for _, c := range cases {
		t.Run("24To32"+c.description, func(t *testing.T) {
			converted := utils.Uint24To32(c.u24)

			if diff := cmp.Diff(converted, c.u32); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("32To24"+c.description, func(t *testing.T) {
			converted := utils.Uint32To24(c.u32)

			if diff := cmp.Diff(converted, c.u24); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestUint64And40(t *testing.T) {
	cases := []struct {
		description string
		u40         []byte
		u64         uint64
	}{
		{
			"Normal",
			[]byte{0xff, 0xff, 0xff, 0xff, 0xff},
			0x000000ffffffffff,
		},
	}

	for _, c := range cases {
		t.Run("40To64/"+c.description, func(t *testing.T) {
			converted := utils.Uint40To64(c.u40)

			if diff := cmp.Diff(converted, c.u64); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("64To40/"+c.description, func(t *testing.T) {
			converted := utils.Uint64To40(c.u64)

			if diff := cmp.Diff(converted, c.u40); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestPLMN(t *testing.T) {
	cases := []struct {
		description string
		mcc, mnc    string
		encoded     []byte
	}{
		{
			"2-digit",
			"123", "45",
			[]byte{0x21, 0xf3, 0x54},
		}, {
			"3-digit",
			"123", "456",
			[]byte{0x21, 0x63, 0x54},
		},
	}

	for _, c := range cases {
		t.Run("serialize/"+c.description, func(t *testing.T) {
			encoded, err := utils.EncodePLMN(c.mcc, c.mnc)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(encoded, c.encoded); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("Decode/"+c.description, func(t *testing.T) {
			mcc, mnc, err := utils.DecodePLMN(c.encoded)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(mcc, c.mcc); diff != "" {
				t.Error(diff)
			}
			if diff := cmp.Diff(mnc, c.mnc); diff != "" {
				t.Error(diff)
			}
		})
	}
}
