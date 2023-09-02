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

func TestUpdateBearerResponse(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewUpdateBearerResponse(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
				ie.NewBearerContext(
					ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
					ie.NewEPSBearerID(0x05),
					ie.NewFullyQualifiedTEID(gtpv2.IFTypeS1USGWGTPU, 0xffffffff, "1.1.1.3", ""),
				),
				ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
				ie.NewBearerContext(
					ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
					ie.NewEPSBearerID(0x06),
					ie.NewFullyQualifiedTEID(gtpv2.IFTypeS1USGWGTPU, 0xffffffff, "1.1.1.3", ""),
				),
			),
			Serialized: []byte{
				// Header
				0x48, 0x62, 0x00, 0x46, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				// BearerContext 1
				0x5d, 0x00, 0x18, 0x00,
				//   Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				//   EBI
				0x49, 0x00, 0x01, 0x00, 0x05,
				//   FTEID
				0x57, 0x00, 0x09, 0x00, 0x81, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x03,
				// BearerContext 2
				0x5d, 0x00, 0x18, 0x00,
				//   Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				//   EBI
				0x49, 0x00, 0x01, 0x00, 0x06,
				//   FTEID
				0x57, 0x00, 0x09, 0x00, 0x81, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x03,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseUpdateBearerResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
