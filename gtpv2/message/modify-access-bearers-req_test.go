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

func TestModifyAccessBearersRequest(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal/NoIE",
			Structured: message.NewModifyAccessBearersRequest(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
			),
			Serialized: []byte{
				// Header
				0x48, 0xd3, 0x00, 0x08, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
			},
		}, {
			Description: "Normal/WithIndication",
			Structured: message.NewModifyAccessBearersRequest(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewIndicationFromOctets(0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40),
				ie.NewBearerContext(
					ie.NewEPSBearerID(0x05),
					ie.NewFullyQualifiedTEID(gtpv2.IFTypeS1UeNodeBGTPU, 0xffffffff, "1.1.1.4", ""),
				),
				ie.NewBearerContext(
					ie.NewEPSBearerID(0x06),
					ie.NewFullyQualifiedTEID(gtpv2.IFTypeS1UeNodeBGTPU, 0xffffffff, "1.1.1.4", ""),
				),
			),
			Serialized: []byte{
				// Header
				0x48, 0xd3, 0x00, 0x3f, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Indication
				0x4d, 0x00, 0x07, 0x00, 0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40,
				// BearerContext 1
				0x5d, 0x00, 0x12, 0x00,
				//   EBI
				0x49, 0x00, 0x01, 0x00, 0x05,
				//   FTEID
				0x57, 0x00, 0x09, 0x00, 0x80, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x04,
				// BearerContext 2
				0x5d, 0x00, 0x12, 0x00,
				//   EBI
				0x49, 0x00, 0x01, 0x00, 0x06,
				//   FTEID
				0x57, 0x00, 0x09, 0x00, 0x80, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x04,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseModifyAccessBearersRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
