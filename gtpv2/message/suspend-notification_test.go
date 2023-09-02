// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
	"github.com/wmnsk/go-gtp/gtpv2/testutils"
)

func TestSuspendNotification(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewSuspendNotification(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewIMSI("123451234567890"),
				ie.NewUserLocationInformationStruct(
					nil, nil, nil, ie.NewTAI("123", "45", 0x0001),
					ie.NewECGI("123", "45", 0x00000101), nil, nil, nil,
				),
				ie.NewEPSBearerID(5),
				ie.NewPacketTMSI(0xdeadbeef),
				ie.NewNodeType(gtpv2.NodeTypeMME),
				ie.NewIPAddress("1.1.1.1"),
				ie.NewPortNumber(2123),
				ie.NewHopCounter(1),
				ie.NewFullyQualifiedTEID(gtpv2.IFTypeS10MMEGTPC, 0xffffffff, "1.1.1.1", ""),
				ie.NewPrivateExtension(123, []byte{1, 2, 3, 4}),
			),
			Serialized: []byte{
				// Header
				0x48, 162, 0x00, 97, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// IMSI
				0x01, 0x00, 0x08, 0x00, 0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0,
				// UserLocationInformation
				0x56, 0x00, 0x0d, 0x00, 0x18,
				0x21, 0xf3, 0x54, 0x00, 0x01,
				0x21, 0xf3, 0x54, 0x00, 0x00, 0x01, 0x01,
				// EPSBearerID
				0x49, 0x00, 0x01, 0x00, 0x05,
				// PacketTMSI
				0x6f, 0x00, 0x04, 0x00, 0xde, 0xad, 0xbe, 0xef,
				// NodeType
				0x87, 0x00, 0x01, 0x00, 0x01,
				// IPAddress
				0x4a, 0x00, 0x04, 0x00, 0x01, 0x01, 0x01, 0x01,
				// PortNumber
				0x7e, 0x00, 0x02, 0x00, 0x08, 0x4b,
				// HopCounter
				0x71, 0x00, 0x01, 0x00, 0x01,
				// FullyQualifiedTEID
				0x57, 0x00, 0x09, 0x00, 0x8c, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x01,
				// PrivateExtension
				0xff, 0x00, 0x06, 0x00, 0x00, 123, 1, 2, 3, 4,
			},
		},
		{
			Description: "Erlang binary",
			Structured: message.NewSuspendNotification(
				1024, 127,
				ie.NewIPAddress("1.2.3.4"),
				ie.NewPortNumber(2345),
			),
			Serialized: []byte{
				0x48, 0xa2, 0x00, 0x16, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00,
				0x7f, 0x00, 0x4a, 0x00, 0x04, 0x00, 0x01, 0x02, 0x03, 0x04,
				0x7e, 0x00, 0x02, 0x00, 0x09, 0x29,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseSuspendNotification(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
