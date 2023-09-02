// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
	"github.com/wmnsk/go-gtp/gtpv2/testutils"
)

func TestUpdatePDNConnectionSetRequest(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewUpdatePDNConnectionSetRequest(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewFullyQualifiedCSID("1.1.1.1", 1),
				ie.NewFullyQualifiedCSID("1.1.1.1", 1).WithInstance(1),
			),
			Serialized: []byte{
				// Header
				0x48, 0xc8, 0x00, 0x1e, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// MME-FQ-CSID
				0x84, 0x00, 0x07, 0x00, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01,
				// SGW-FQ-CSID
				0x84, 0x00, 0x07, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseUpdatePDNConnectionSetRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
