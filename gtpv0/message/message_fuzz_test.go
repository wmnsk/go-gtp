package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv0/message"
)

func FuzzParse(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		if _, err := message.Parse(b); err != nil {
			t.Skip()
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
