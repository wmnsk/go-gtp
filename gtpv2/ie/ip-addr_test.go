package ie_test

import (
	"io"
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
		pdnType     uint8
		ipv4        net.IP
		ipv6        net.IP
	}{
		{
			"PDNType IPv4",
			ie.NewPDNAddressAllocation("1.2.3.4"),
			gtpv2.PDNTypeIPv4,
			net.ParseIP("1.2.3.4"),
			nil,
		},
		{
			"PDNType IPv6",
			ie.NewPDNAddressAllocation("::1"),
			gtpv2.PDNTypeIPv6,
			nil,
			net.ParseIP("::1"),
		},
		{
			"PDNType IPv4v6",
			ie.NewPDNAddressAllocationDual("1.2.3.4", "::1", 64),
			gtpv2.PDNTypeIPv4v6,
			net.ParseIP("1.2.3.4"),
			net.ParseIP("::1"),
		},
		{
			"PDNType NonIP",
			ie.NewPDNAddressAllocation(""),
			gtpv2.PDNTypeNonIP,
			nil,
			nil,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			pdnType := c.paa.MustPDNType()
			if diff := cmp.Diff(pdnType, c.pdnType); diff != "" {
				t.Error(diff)
			}

			ipv4, _ := c.paa.IPv4()
			if diff := cmp.Diff(ipv4, c.ipv4); diff != "" {
				t.Error(diff)
			}

			ipv6, _ := c.paa.IPv6()
			if diff := cmp.Diff(ipv6, c.ipv6); diff != "" {
				t.Error(diff)
			}

			ip, err := c.paa.IP()
			if err == nil {
				v := ipv4
				if pdnType == gtpv2.PDNTypeIPv6 {
					v = ipv6
				}
				if diff := cmp.Diff(ip, v); diff != "" {
					t.Error(diff)
				}
			} else if err != io.ErrUnexpectedEOF {
				t.Error(err)
			}
		})
	}
}
