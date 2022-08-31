package ie_test

import (
	"net"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

func TestPDNAddressAllocationIP(t *testing.T) {
	cases := []struct {
		description string
		paa         *ie.IE
		ipv4        net.IP
		ipv6        net.IP
	}{
		{
			"PDNType IPv4",
			ie.NewPDNAddressAllocation("1.2.3.4"),
			net.ParseIP("1.2.3.4"),
			nil,
		},
		{
			"PDNType IPv6",
			ie.NewPDNAddressAllocation("::1"),
			nil,
			net.ParseIP("::1"),
		},
		{
			"PDNType IPv4v6",
			ie.NewPDNAddressAllocationDual("1.2.3.4", "::1", uint8(64)),
			net.ParseIP("1.2.3.4"),
			net.ParseIP("::1"),
		},
		{
			"PDNType NonIP",
			ie.NewPDNAddressAllocation(""),
			nil,
			nil,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			pdnType := c.paa.MustPDNType()
			if pdnType == gtpv2.PDNTypeIPv4 || pdnType == gtpv2.PDNTypeIPv4v6 {
				ipv4, err := c.paa.IPv4()
				if err != nil {
					t.Fatal(err)
				}
				if diff := cmp.Diff(ipv4, c.ipv4); diff != "" {
					t.Error(diff)
				}
			}
			if pdnType == gtpv2.PDNTypeIPv6 || pdnType == gtpv2.PDNTypeIPv4v6 {
				ipv6, err := c.paa.IPv6()
				if err != nil {
					t.Fatal(err)
				}
				if diff := cmp.Diff(ipv6, c.ipv6); diff != "" {
					t.Error(diff)
				}
			}
			if pdnType == gtpv2.PDNTypeNonIP {
				_, err := c.paa.IP()
				if err == ie.ErrIEValueNotFound {
					t.Error(err)
				}
			}

		})
	}
}
