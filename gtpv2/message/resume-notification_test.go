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

func TestResumeNotification(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewResumeNotification(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewIMSI("123451234567890"),
				ie.NewEPSBearerID(5),
				ie.NewNodeType(gtpv2.NodeTypeMME),
				ie.NewFullyQualifiedTEID(gtpv2.IFTypeS10MMEGTPC, 0xffffffff, "1.1.1.1", ""),
				ie.NewPrivateExtension(256, []byte{4, 5, 6, 7}),
			),
			Serialized: []byte{
				// Header
				0x48, 164, 0x00, 53, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// IMSI
				0x01, 0x00, 0x08, 0x00, 0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0,
				// EPSBearerID
				0x49, 0x00, 0x01, 0x00, 0x05,
				// NodeType
				0x87, 0x00, 0x01, 0x00, 0x01,
				// FullyQualifiedTEID
				0x57, 0x00, 0x09, 0x00, 0x8c, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x01,
				// PrivateExtension
				0xff, 0x00, 0x06, 0x00, 1, 0, 4, 5, 6, 7,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseResumeNotification(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
