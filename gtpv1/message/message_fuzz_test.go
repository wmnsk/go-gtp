package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv1/message"
)

func FuzzParse(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		if _, err := message.Parse(b); err != nil {
			t.Skip()
		}
	})
}

func FuzzHeaderParse(f *testing.F) {
	f.Add([]byte("10000000"))
	f.Add([]byte("70\x00\x0400000000\x000"))
	f.Add([]byte("70\x00\x0400000000\x0100\x00"))

	f.Fuzz(func(t *testing.T, pkt []byte) {
		header, err := message.ParseHeader(pkt)
		if header == nil && err == nil {
			t.Errorf("nil without error")
		}
	})
}
