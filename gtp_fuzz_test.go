package gtp_test

import (
	"testing"

	"github.com/wmnsk/go-gtp"
)

func FuzzParse(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		if _, err := gtp.Parse(b); err != nil {
			t.Skip()
		}
	})
}
