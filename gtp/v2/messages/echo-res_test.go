// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtp/v2/ies"
	"github.com/wmnsk/go-gtp/gtp/v2/messages"
	"github.com/wmnsk/go-gtp/gtp/v2/testutils"
)

func TestEchoResponse(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured:  messages.NewEchoResponse(0, ies.NewRecovery(0x80)),
			Serialized: []byte{
				0x40, 0x02, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00,
				0x03, 0x00, 0x01, 0x00, 0x80,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializeable, error) {
		v, err := messages.DecodeEchoResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
