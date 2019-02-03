// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v0 "github.com/wmnsk/go-gtp/gtp/v0"
	"github.com/wmnsk/go-gtp/gtp/v0/ies"
)

func TestIE(t *testing.T) {
	cases := []struct {
		description string
		structured  *ies.IE
		serialized  []byte
	}{
		{
			"Cause",
			ies.NewCause(v0.CauseRequestAccepted),
			[]byte{0x01, 0x80},
		}, {
			"IMSI",
			ies.NewIMSI("123450123456789"),
			[]byte{0x02, 0x21, 0x43, 0x05, 0x21, 0x43, 0x65, 0x87, 0xf9},
		}, {
			"RAI",
			ies.NewRouteingAreaIdentity("123", "45", 0x1111, 0x22),
			[]byte{0x03, 0x21, 0xf3, 0x54, 0x11, 0x11, 0x22},
		}, {
			"TLLI",
			ies.NewTemporaryLogicalLinkIdentity(0xff00ff00),
			[]byte{0x04, 0xff, 0x00, 0xff, 0x00},
		}, {
			"P-TMSI",
			ies.NewPacketTMSI(0xff00ff00),
			[]byte{0x05, 0xff, 0x00, 0xff, 0x00},
		}, { // XXX - not implemented fully
			"QoS Profile",
			ies.NewQualityOfServiceProfile(1, 1, 1, 1, 1),
			[]byte{0x06, 0x09, 0x11, 0x01},
		}, {
			"ReorderingRequired",
			ies.NewReorderingRequired(false),
			[]byte{0x08, 0xfe},
		},
		/* XXX - not implemented
		{
			"AuthenticationTriplet",
			ies.NewAuthenticationTriplet(),
			[]byte{},
		}, {
			"MAPCause",
			ies.NewMAPCause(),
			[]byte{},
		}, {
			"PacketTMSISignature",
			ies.NewPacketTMSISignature(),
			[]byte{},
		}, {
			"MSValidated",
			ies.NewMSValidated(),
			[]byte{},
		},*/
		{
			"Recovery",
			ies.NewRecovery(0x80),
			[]byte{0x0e, 0x80},
		}, {
			"SelectionMode",
			ies.NewSelectionMode(0xff),
			[]byte{0x0f, 0xff},
		}, {
			"FlowLabelDataI",
			ies.NewFlowLabelDataI(0x0001),
			[]byte{0x10, 0x00, 0x01},
		}, {
			"FlowLabelSignalling",
			ies.NewFlowLabelSignalling(0x0001),
			[]byte{0x11, 0x00, 0x01},
		}, {
			"FlowLabelDataII",
			ies.NewFlowLabelDataII(5, 0x0001),
			[]byte{0x12, 0xf5, 0x00, 0x01},
		}, {
			"MSNotReachableReason",
			ies.NewMSNotReachableReason(0xff),
			[]byte{0x13, 0xff},
		}, {
			"ChargingID",
			ies.NewChargingID(0xff00ff00),
			[]byte{0x7f, 0xff, 0x00, 0xff, 0x00},
		}, {
			"EndUserAddress/v4",
			ies.NewEndUserAddressIPv4("1.1.1.1"),
			[]byte{
				// Type, Length
				0x80, 0x00, 0x06,
				// Value
				0xf1, 0x21, 0x01, 0x01, 0x01, 0x01,
			},
		}, {
			"EndUserAddress/v6",
			ies.NewEndUserAddressIPv6("2001::1"),
			[]byte{
				// Type, Length
				0x80, 0x00, 0x12,
				// Value
				0xf1, 0x57, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		}, {
			"EndUserAddress/ppp",
			ies.NewEndUserAddressPPP(),
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
			ies.NewMMContext(),
			[]byte{
				// Type, Length
				0x81, 0x00, 0x00,
				// Value
			},
		}, {
			"PDPContext",
			ies.NewPDPContext(),
			[]byte{
				// Type, Length
				0x81, 0x00, 0x00,
				// Value
			},
		},
		*/
		{
			"AccessPointName",
			ies.NewAccessPointName("some.apn.example"),
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
			ies.NewProtocolConfigurationOption(),
			[]byte{},
		}, */
		{
			"GSNAddress/v4",
			ies.NewGSNAddress("1.1.1.1"),
			[]byte{
				// Type, Length
				0x85, 0x00, 0x04,
				// Value
				0x01, 0x01, 0x01, 0x01,
			},
		}, {
			"GSNAddress/v6",
			ies.NewGSNAddress("2001::1"),
			[]byte{
				// Type, Length
				0x85, 0x00, 0x10,
				// Value
				0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		}, {
			"MSISDN",
			ies.NewMSISDN("819012345678"),
			[]byte{
				// Type, Length
				0x86, 0x00, 0x07,
				// Value
				0x91, 0x18, 0x09, 0x21, 0x43, 0x65, 0x87,
			},
		}, {
			"ChargingGatewayAddress/v4",
			ies.NewChargingGatewayAddress("1.1.1.1"),
			[]byte{
				// Type, Length
				0xfb, 0x00, 0x04,
				// Value
				0x01, 0x01, 0x01, 0x01,
			},
		}, {
			"ChargingGatewayAddress/v6",
			ies.NewChargingGatewayAddress("2001::1"),
			[]byte{
				// Type, Length
				0xfb, 0x00, 0x10,
				// Value
				0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		}, {
			"PrivateExtension",
			ies.NewPrivateExtension(0x0080, []byte{0xde, 0xad, 0xbe, 0xef}),
			[]byte{
				// Type, Length
				0xff, 0x00, 0x06,
				// Value
				0x00, 0x80, 0xde, 0xad, 0xbe, 0xef,
			},
		},
	}

	for _, c := range cases {
		t.Run("serialize/"+c.description, func(t *testing.T) {
			got, err := c.structured.Serialize()
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.serialized); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("decode/"+c.description, func(t *testing.T) {
			got, err := ies.Decode(c.serialized)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.structured); diff != "" {
				t.Error(err)
			}
		})
	}
}
