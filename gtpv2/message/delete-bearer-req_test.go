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

func TestDeleteBearerRequest(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewDeleteBearerRequest(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewEPSBearerID(5),
				ie.NewCause(gtpv2.CauseISRDeactivation, 0, 0, 0, nil),
			),
			Serialized: []byte{
				// Header
				0x48, 0x63, 0x00, 0x13, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// EBI
				0x49, 0x00, 0x01, 0x00, 0x05,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x05, 0x00,
			},
		}, {
			Description: "EBIs",
			Structured: message.NewDeleteBearerRequest(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewEPSBearerID(5).WithInstance(1),
				ie.NewEPSBearerID(6).WithInstance(1),
				ie.NewCause(gtpv2.CauseISRDeactivation, 0, 0, 0, nil),
			),
			Serialized: []byte{
				// Header
				0x48, 0x63, 0x00, 0x18, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// EBIs 1
				0x49, 0x00, 0x01, 0x01, 0x05,
				// EBIs 2
				0x49, 0x00, 0x01, 0x01, 0x06,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x05, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseDeleteBearerRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
