package message_test

import (
	"net"
	"testing"

	"github.com/wmnsk/go-gtp/gtpv2/message"
)

var (
	mac1, _ = net.ParseMAC("12:34:56:78:90:01")
	mac2, _ = net.ParseMAC("12:34:56:78:90:02")
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
