// Copyright 2019-2022 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "errors"

// Error definitions.
var (
	ErrInvalidLength   = errors.New("length value is invalid")
	ErrTooShortToParse = errors.New("too short to decode as GTP")
)
