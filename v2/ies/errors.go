// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "errors"

// Error definitions.
var (
	ErrTooShortToParse = errors.New("too short to decode as GTP")
	ErrInvalidLength    = errors.New("length value is invalid")

	ErrInvalidType = errors.New("invalid type")
	ErrIENotFound  = errors.New("could not find the specified IE in a grouped IE")
)
