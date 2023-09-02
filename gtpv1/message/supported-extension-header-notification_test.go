// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv1/ie"
	"github.com/wmnsk/go-gtp/gtpv1/message"
	"github.com/wmnsk/go-gtp/gtpv1/testutils"
)

func TestSupportedExtensionHeaderNotification(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: message.NewSupportedExtensionHeaderNotification(
				testutils.TestBearerInfo.TEID, 0,
				ie.NewExtensionHeaderTypeList(
					message.ExtHeaderTypePDUSessionContainer,
					message.ExtHeaderTypeUDPPort,
				),
			),
			Serialized: []byte{
				// Header
				0x30, 0x1f, 0x00, 0x04, 0x11, 0x22, 0x33, 0x44,
				// ExtensionHeaderTypeList
				0x8d, 0x02, 0x85, 0x40,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseSupportedExtensionHeaderNotification(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
