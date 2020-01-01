// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/v1/messages"
	"github.com/wmnsk/go-gtp/v1/testutils"
)

func TestGeneric(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured:  messages.NewGeneric(messages.MsgTypeEchoRequest, 0, 0),
			Serialized: []byte{
				0x32, 0x01, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := messages.ParseGeneric(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
