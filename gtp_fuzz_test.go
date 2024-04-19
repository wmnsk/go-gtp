// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtp_test

import (
	"testing"

	"github.com/wmnsk/go-gtp"
)

func FuzzParse(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		if _, err := gtp.Parse(b); err != nil {
			t.Skip()
		}
	})
}
