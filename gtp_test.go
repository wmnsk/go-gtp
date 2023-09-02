// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pascaldekloe/goe/verify"

	v0msg "github.com/wmnsk/go-gtp/gtpv0/message"
	v1msg "github.com/wmnsk/go-gtp/gtpv1/message"
	v2ie "github.com/wmnsk/go-gtp/gtpv2/ie"
	v2msg "github.com/wmnsk/go-gtp/gtpv2/message"
)

var v0flow = struct {
	seq   uint16
	label uint16
	tid   uint64
}{1, 0, 0x2143658709214355}

func TestMessage(t *testing.T) {
	cases := []struct {
		description string
		structured  Message
		serialized  []byte
	}{
		{
			"GTPv0 Echo Request",
			v0msg.NewEchoRequest(v0flow.seq, v0flow.label, v0flow.tid),
			[]byte{
				0x1e, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
				0xff, 0xff, 0xff, 0xff, 0x21, 0x43, 0x65, 0x87,
				0x09, 0x21, 0x43, 0x55,
			},
		}, {
			"GTPv1 Echo Request",
			v1msg.NewEchoRequest(0),
			[]byte{
				0x32, 0x01, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
			},
		}, {
			"GTPv2 Echo Request",
			v2msg.NewEchoRequest(0, v2ie.NewRecovery(0x80)),
			[]byte{
				0x40, 0x01, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00,
				0x03, 0x00, 0x01, 0x00, 0x80,
			},
		},
	}

	for _, c := range cases {
		t.Run("Encode/"+c.description, func(t *testing.T) {
			got, err := Marshal(c.structured)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.serialized); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("Parse/"+c.description, func(t *testing.T) {
			v, err := Parse(c.serialized)
			if err != nil {
				t.Fatal(err)
			}

			if got, want := v, c.structured; !verify.Values(t, "", got, want) {
				t.Fail()
			}
		})

		t.Run("MarshalLen/"+c.description, func(t *testing.T) {
			if got, want := c.structured.MarshalLen(), len(c.serialized); got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})

		t.Run("Interface/"+c.description, func(t *testing.T) {
			decoded, err := Parse(c.serialized)
			if err != nil {
				t.Fatal(err)
			}

			if got, want := decoded.Version(), c.structured.(Message).Version(); got != want {
				t.Fatalf("got %v want %v", got, want)
			}
			if got, want := decoded.MessageType(), c.structured.(Message).MessageType(); got != want {
				t.Fatalf("got %v want %v", got, want)
			}
			if got, want := decoded.MessageTypeName(), c.structured.(Message).MessageTypeName(); got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
