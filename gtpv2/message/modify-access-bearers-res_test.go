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

func TestModifyAccessBearersResponse(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal/CauseOnly",
			Structured: message.NewModifyAccessBearersResponse(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
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
				0x48, 0xd4, 0x00, 0x3a, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
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
		v, err := message.ParseModifyAccessBearersResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
