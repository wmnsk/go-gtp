// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtp

import "errors"

// Common error definitions.
var (
	ErrInvalidVersion    = errors.New("got invalid version")
	ErrInvalidLength     = errors.New("length value is invalid")
	ErrTooShortToParse   = errors.New("too short to decode as GTP")
	ErrTooShortToMarshal = errors.New("too short to serialize")
)
