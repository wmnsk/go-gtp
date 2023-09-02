// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"
	"time"

	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
	"github.com/wmnsk/go-gtp/gtpv2/testutils"
)

func TestDeleteSessionRequest(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal/FromMMEtoSGW",
			Structured: message.NewDeleteSessionRequest(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewEPSBearerID(5),
				ie.NewUserLocationInformationStruct(
					nil, nil, nil, ie.NewTAI("123", "45", 0x0001),
					ie.NewECGI("123", "45", 0x00000101), nil, nil, nil,
				),
				ie.NewIndicationFromOctets(0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40),
				ie.NewULITimestamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			Serialized: []byte{
				// Header
				0x48, 0x24, 0x00, 0x31, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// EBI
				0x49, 0x00, 0x01, 0x00, 0x05,
				// ULI: TAI ECGI
				0x56, 0x00, 0x0d, 0x00, 0x18,
				0x21, 0xf3, 0x54, 0x00, 0x01,
				0x21, 0xf3, 0x54, 0x00, 0x00, 0x01, 0x01,
				// Indication
				0x4d, 0x00, 0x07, 0x00, 0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40,
				// ULITimestamp
				0xaa, 0x00, 0x04, 0x00, 0xdf, 0xd5, 0x2c, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseDeleteSessionRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
