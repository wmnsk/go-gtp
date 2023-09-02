// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"net"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

var (
	ip1 = net.ParseIP("1.1.1.1")
	ip2 = net.ParseIP("2.2.2.2")
)

func TestPCOPPP(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.PCOPPP
		serialized  []byte
	}{
		{
			"PAP",
			ie.NewPCOPPPWithPAP(ie.PCOPPPConfigurationRequest, 0, "foo", "bar"),
			[]byte{0x01, 0x00, 0x00, 0x0c, 0x03, 0x66, 0x6f, 0x6f, 0x03, 0x62, 0x61, 0x72},
		}, {
			"CHAP",
			ie.NewPCOPPPWithCHAP(ie.PCOPPPConfigurationRequest, 0, []byte{0xde, 0xad, 0xbe, 0xef}, "foo"),
			[]byte{0x01, 0x00, 0x00, 0x0c, 0x04, 0xde, 0xad, 0xbe, 0xef, 0x66, 0x6f, 0x6f},
		}, {
			"IPCPOptions/IPAddress+PrimaryDNS",
			ie.NewPCOPPPWithIPCPOptions(
				ie.PCOPPPConfigurationRequest,
				0,
				ie.NewIPCPOptionIPAddress(ip1),
				ie.NewIPCPOptionPrimaryDNS(ip2),
			),
			[]byte{0x01, 0x00, 0x00, 0x10, 0x03, 0x06, 0x01, 0x01, 0x01, 0x01, 0x81, 0x06, 0x02, 0x02, 0x02, 0x02},
		},
	}

	for _, c := range cases {
		t.Run("serialize/"+c.description, func(t *testing.T) {
			got, err := c.structured.Marshal()
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.serialized); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("decode/"+c.description, func(t *testing.T) {
			got, err := ie.ParsePCOPPP(c.serialized)
			if err != nil {
				t.Fatal(err)
			}

			opt := cmp.AllowUnexported(*got, *c.structured)
			if diff := cmp.Diff(got, c.structured, opt); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestPCOPPPPAPFields(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.PAPFields
		serialized  []byte
	}{
		{
			"Normal",
			ie.NewPAPFields("foo", "bar"),
			[]byte{0x03, 0x66, 0x6f, 0x6f, 0x03, 0x62, 0x61, 0x72},
		},
	}

	for _, c := range cases {
		t.Run("serialize/"+c.description, func(t *testing.T) {
			got, err := c.structured.Marshal()
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.serialized); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("decode/"+c.description, func(t *testing.T) {
			got, err := ie.ParsePAPFields(c.serialized)
			if err != nil {
				t.Fatal(err)
			}

			opt := cmp.AllowUnexported(*got, *c.structured)
			if diff := cmp.Diff(got, c.structured, opt); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestPCOPPPCHAPFields(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.CHAPFields
		serialized  []byte
	}{
		{
			"Normal",
			ie.NewCHAPFields([]byte{0xde, 0xad, 0xbe, 0xef}, "foo"),
			[]byte{0x04, 0xde, 0xad, 0xbe, 0xef, 0x66, 0x6f, 0x6f},
		},
	}

	for _, c := range cases {
		t.Run("serialize/"+c.description, func(t *testing.T) {
			got, err := c.structured.Marshal()
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.serialized); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("decode/"+c.description, func(t *testing.T) {
			got, err := ie.ParseCHAPFields(c.serialized)
			if err != nil {
				t.Fatal(err)
			}

			opt := cmp.AllowUnexported(*got, *c.structured)
			if diff := cmp.Diff(got, c.structured, opt); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestPCOPPPIPCPOption(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IPCPOption
		serialized  []byte
	}{
		{
			"IPAddress",
			ie.NewIPCPOptionIPAddress(ip1),
			[]byte{0x03, 0x06, 0x01, 0x01, 0x01, 0x01},
		}, {
			"PrimaryDNS",
			ie.NewIPCPOptionPrimaryDNS(ip2),
			[]byte{0x81, 0x06, 0x02, 0x02, 0x02, 0x02},
		},
	}

	for _, c := range cases {
		t.Run("serialize/"+c.description, func(t *testing.T) {
			got, err := c.structured.Marshal()
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.serialized); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("decode/"+c.description, func(t *testing.T) {
			got, err := ie.ParseIPCPOption(c.serialized)
			if err != nil {
				t.Fatal(err)
			}

			opt := cmp.AllowUnexported(*got, *c.structured)
			if diff := cmp.Diff(got, c.structured, opt); diff != "" {
				t.Error(diff)
			}
		})
	}
}
