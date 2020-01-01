// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	v2 "github.com/wmnsk/go-gtp/v2"
	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
	"github.com/wmnsk/go-gtp/v2/testutils"
)

func TestModifyBearerResponse(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal/FromSGWtoMME",
			Structured: messages.NewModifyBearerResponse(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
				ies.NewBearerContext(
					ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
					ies.NewEPSBearerID(0x05),
					ies.NewFullyQualifiedTEID(v2.IFTypeS1USGWGTPU, 0xffffffff, "1.1.1.3", ""),
				),
			),
			Serialized: []byte{
				// Header
				0x48, 0x23, 0x00, 0x2a, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				// BearerContext
				0x5d, 0x00, 0x18, 0x00,
				//   Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				//   EBI
				0x49, 0x00, 0x01, 0x00, 0x05,
				//   FTEID
				0x57, 0x00, 0x09, 0x00, 0x81, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x03,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := messages.ParseModifyBearerResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
