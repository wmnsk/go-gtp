package ie_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv0/ie"
)

func FuzzParse(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		if _, err := ie.Parse(b); err != nil {
			t.Skip()
		}
	})
}
