// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-gtp/gtpv0"
	"github.com/wmnsk/go-gtp/gtpv0/ie"
)

func TestIE(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		Serialized  []byte
	}{
		{
			"Cause",
			ie.NewCause(gtpv0.CauseRequestAccepted),
			[]byte{0x01, 0x80},
		}, {
			"IMSI",
			ie.NewIMSI("123450123456789"),
			[]byte{0x02, 0x21, 0x43, 0x05, 0x21, 0x43, 0x65, 0x87, 0xf9},
		}, {
			"RAI",
			ie.NewRouteingAreaIdentity("123", "45", 0x1111, 0x22),
			[]byte{0x03, 0x21, 0xf3, 0x54, 0x11, 0x11, 0x22},
		}, {
			"TLLI",
			ie.NewTemporaryLogicalLinkIdentity(0xff00ff00),
			[]byte{0x04, 0xff, 0x00, 0xff, 0x00},
		}, {
			"PacketTMSI",
			ie.NewPacketTMSI(0xdeadbeef),
			[]byte{0x05, 0xde, 0xad, 0xbe, 0xef},
		}, { // XXX - not implemented fully
			"QoS Profile",
			ie.NewQualityOfServiceProfile(1, 1, 1, 1, 1),
			[]byte{0x06, 0x09, 0x11, 0x01},
		}, {
			"ReorderingRequired",
			ie.NewReorderingRequired(false),
			[]byte{0x08, 0xfe},
		},
		/* XXX - not implemented
		{
			"AuthenticationTriplet",
			ie.NewAuthenticationTriplet(),
			[]byte{},
		}, {
			"MAPCause",
			ie.NewMAPCause(),
			[]byte{},
		}, {
			"PacketTMSISignature",
			ie.NewPacketTMSISignature(),
			[]byte{},
		}, {
			"MSValidated",
			ie.NewMSValidated(),
			[]byte{},
		},*/
		{
			"PTMSISignature",
			ie.NewPTMSISignature(0xbeebee),
			[]byte{0x0c, 0xbe, 0xeb, 0xee},
		}, {
			"Recovery",
			ie.NewRecovery(0x80),
			[]byte{0x0e, 0x80},
		}, {
			"SelectionMode",
			ie.NewSelectionMode(0xff),
			[]byte{0x0f, 0xff},
		}, {
			"FlowLabelDataI",
			ie.NewFlowLabelDataI(0x0001),
			[]byte{0x10, 0x00, 0x01},
		}, {
			"FlowLabelSignalling",
			ie.NewFlowLabelSignalling(0x0001),
			[]byte{0x11, 0x00, 0x01},
		}, {
			"FlowLabelDataII",
			ie.NewFlowLabelDataII(5, 0x0001),
			[]byte{0x12, 0xf5, 0x00, 0x01},
		}, {
			"MSNotReachableReason",
			ie.NewMSNotReachableReason(0xff),
			[]byte{0x13, 0xff},
		}, {
			"ChargingID",
			ie.NewChargingID(0xff00ff00),
			[]byte{0x7f, 0xff, 0x00, 0xff, 0x00},
		}, {
			"EndUserAddress/v4",
			ie.NewEndUserAddressIPv4("1.1.1.1"),
			[]byte{
				// Type, Length
				0x80, 0x00, 0x06,
				// Value
				0xf1, 0x21, 0x01, 0x01, 0x01, 0x01,
			},
		}, {
			"EndUserAddress/v6",
			ie.NewEndUserAddressIPv6("2001::1"),
			[]byte{
				// Type, Length
				0x80, 0x00, 0x12,
				// Value
				0xf1, 0x57, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		}, {
			"EndUserAddress/ppp",
			ie.NewEndUserAddressPPP(),
			[]byte{
				// Type, Length
				0x80, 0x00, 0x02,
				// Value
				0xf0, 0xf1,
			},
		},
		/* XXX - not implemented
		{
			"MMContext",
			ie.NewMMContext(),
			[]byte{
				// Type, Length
				0x81, 0x00, 0x00,
				// Value
			},
		}, {
			"PDPContext",
			ie.NewPDPContext(),
			[]byte{
				// Type, Length
				0x81, 0x00, 0x00,
				// Value
			},
		},
		*/
		{
			"AccessPointName",
			ie.NewAccessPointName("some.apn.example"),
			[]byte{
				// Type, Length
				0x83, 0x00, 0x11,
				// Value
				0x04, 0x73, 0x6f, 0x6d, 0x65, 0x03, 0x61, 0x70, 0x6e, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
			},
		},
		/* XXX - not implemented
		{
			"PCO",
			ie.NewProtocolConfigurationOption(),
			[]byte{},
		}, */
		{
			"GSNAddress/v4",
			ie.NewGSNAddress("1.1.1.1"),
			[]byte{
				// Type, Length
				0x85, 0x00, 0x04,
				// Value
				0x01, 0x01, 0x01, 0x01,
			},
		}, {
			"GSNAddress/v6",
			ie.NewGSNAddress("2001::1"),
			[]byte{
				// Type, Length
				0x85, 0x00, 0x10,
				// Value
				0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		}, {
			"MSISDN",
			ie.NewMSISDN("819012345678"),
			[]byte{
				// Type, Length
				0x86, 0x00, 0x07,
				// Value
				0x91, 0x18, 0x09, 0x21, 0x43, 0x65, 0x87,
			},
		}, {
			"ChargingGatewayAddress/v4",
			ie.NewChargingGatewayAddress("1.1.1.1"),
			[]byte{
				// Type, Length
				0xfb, 0x00, 0x04,
				// Value
				0x01, 0x01, 0x01, 0x01,
			},
		}, {
			"ChargingGatewayAddress/v6",
			ie.NewChargingGatewayAddress("2001::1"),
			[]byte{
				// Type, Length
				0xfb, 0x00, 0x10,
				// Value
				0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		}, {
			"PrivateExtension",
			ie.NewPrivateExtension(0x0080, []byte{0xde, 0xad, 0xbe, 0xef}),
			[]byte{
				// Type, Length
				0xff, 0x00, 0x06,
				// Value
				0x00, 0x80, 0xde, 0xad, 0xbe, 0xef,
			},
		},
	}

	for _, c := range cases {
		t.Run("Marshal/"+c.description, func(t *testing.T) {
			got, err := c.structured.Marshal()
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.Serialized); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("Parse/"+c.description, func(t *testing.T) {
			got, err := ie.Parse(c.Serialized)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.structured); diff != "" {
				t.Error(err)
			}
		})
	}
}
