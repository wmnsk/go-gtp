package ie_test

import (
	"io"
	"net"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

func TestPDNAddressAllocationIP(t *testing.T) {

	type IPAddr struct {
		ip  net.IP
		err error
	}
	cases := []struct {
		description string
		paa         *ie.IE
		ipv4        IPAddr
		ipv6        IPAddr
	}{
		{
			"PDNType IPv4",
			ie.NewPDNAddressAllocation("1.2.3.4"),
			IPAddr{
				net.ParseIP("1.2.3.4"),
				nil,
			},
			IPAddr{
				nil,
				ie.ErrIEValueNotFound,
			},
		},
		{
			"PDNType IPv6",
			ie.NewPDNAddressAllocation("::1"),
			IPAddr{
				nil,
				ie.ErrIEValueNotFound,
			},
			IPAddr{
				net.ParseIP("::1"),
				nil,
			},
		},
		{
			"PDNType IPv4v6",
			ie.NewPDNAddressAllocationDual("1.2.3.4", "::1", uint8(64)),
			IPAddr{
				net.ParseIP("1.2.3.4"),
				nil,
			},
			IPAddr{
				net.ParseIP("::1"),
				nil,
			},
		},
		{
			"PDNType NonIP",
			ie.NewPDNAddressAllocation(""),
			IPAddr{
				nil,
				io.ErrUnexpectedEOF,
			},
			IPAddr{
				nil,
				io.ErrUnexpectedEOF,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {

			ipv4, err := c.paa.IPv4()
			if diff := cmp.Diff(ipv4, c.ipv4.ip); diff != "" {
				t.Error(diff)
			}
			if err != c.ipv4.err {
				t.Errorf("IPv4() error = %v, want %v", err, c.ipv4.err)
			}

			ipv6, err := c.paa.IPv6()
			if diff := cmp.Diff(ipv6, c.ipv6.ip); diff != "" {
				t.Error(diff)
			}
			if err != c.ipv6.err {
				t.Errorf("IPv6() error = %v, want %v", err, c.ipv6.err)
			}

		})
	}
}
