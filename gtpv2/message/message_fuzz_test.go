//go:build go1.18

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv2/message"
)

func FuzzParse(f *testing.F) {
	testcases := [][]byte{
		{0x10, 0x20, 0x30},
		{0x48, 0x20, 0x00, 0x14, 0xff, 0xff, 0xff, 0xff, 0xda, 0xda, 0xda},
	}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, data []byte) {
		if v, err := message.Parse(data); err == nil && v == nil {
			t.Errorf("nil without error")
		}
	})
}

func FuzzHeaderParse(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		if _, err := message.ParseHeader(b); err != nil {
			t.Skip()
		}
	})
}
