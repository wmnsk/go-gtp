// Copyright 2019-2021 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv2/message"
	"github.com/wmnsk/go-gtp/gtpv2/testutils"

	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

func TestCreateBearerRequest(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewCreateBearerRequest(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewEPSBearerID(5),
				ie.NewBearerContext(
					ie.NewEPSBearerID(0),
					// ie.NewBearerTFT(), XXX - not implemented yet.
					ie.NewBearerQoS(1, 2, 1, 0xff, 0x1111111111, 0x2222222222, 0x1111111111, 0x2222222222),
				),
			),
			Serialized: []byte{
				// Header
				0x48, 0x5f, 0x00, 0x30, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// EBI
				0x49, 0x00, 0x01, 0x00, 0x05,
				// BearerContext
				0x5d, 0x00, 0x1f, 0x00,
				//   EBI
				0x49, 0x00, 0x01, 0x00, 0x00,
				//   BearerQoS
				0x50, 0x00, 0x16, 0x00, 0x49, 0xff,
				0x11, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x22,
				0x11, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x22,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseCreateBearerRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
