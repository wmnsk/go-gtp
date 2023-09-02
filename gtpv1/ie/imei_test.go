// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "testing"

func TestIE_IMEISV(t *testing.T) {
	t.Run("Encode/Decode IMEI", func(t *testing.T) {
		ie := NewIMEISV("123456789012345")

		got := ie.MustIMEISV()
		if got != "123456789012345" {
			t.Errorf("wrong IMEI, got: %v", got)
		}
	})

	t.Run("Encode/Decode IMEISV", func(t *testing.T) {
		ie := NewIMEISV("1234567890123456")

		got := ie.MustIMEISV()
		if got != "1234567890123456" {
			t.Errorf("wrong IMEISV, got: %v", got)
		}
	})
}
