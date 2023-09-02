// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"errors"
	"fmt"
)

// Error definitions.
var (
	ErrInvalidLength      = errors.New("got invalid length ")
	ErrTooShortToMarshal  = errors.New("too short to serialize")
	ErrTooShortToParse    = errors.New("too short to decode as GTPv1")
	ErrInvalidMessageType = errors.New("got invalid message type")
)

// InvalidTypeError indicates the type of an ExtensionHeader is invalid.
type InvalidTypeError struct {
	Type uint8
}

// Error returns message with the invalid type given.
func (e *InvalidTypeError) Error() string {
	return fmt.Sprintf("got invalid type: %v", e.Type)
}
