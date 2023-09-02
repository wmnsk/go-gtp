// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv1/message"
	"github.com/wmnsk/go-gtp/gtpv1/testutils"
)

func TestEndMarker(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured:  message.NewEndMarker(),
			Serialized: []byte{
				0x30, 0xfe, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseEndMarker(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
