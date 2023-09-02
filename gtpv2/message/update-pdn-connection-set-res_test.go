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

func TestUpdatePDNConnectionSetResponse(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewUpdatePDNConnectionSetResponse(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewCause(gtpv2.CauseRequestAccepted, 0, 0, 0, nil),
				ie.NewFullyQualifiedCSID("1.1.1.1", 1),
				ie.NewRecovery(0xff),
			),
			Serialized: []byte{
				// Header
				0x48, 0xc9, 0x00, 0x1e, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				// PGW-FQ-CSID
				0x84, 0x00, 0x07, 0x00, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01,
				// Recovery
				0x03, 0x00, 0x01, 0x00, 0xff,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseUpdatePDNConnectionSetResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
