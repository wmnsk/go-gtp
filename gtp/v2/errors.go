// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v2

import (
	"errors"
	"fmt"
)

var (
	// ErrNoHandlersFound indicates that the handler func is not registered in *Conn
	// for the incoming GTPv2 message. In usual cases this error should not be taken
	// as fatal, as the other endpoint can make your program stop working just by
	// sending unregistered messages.
	ErrNoHandlersFound = errors.New("no handlers found for incoming message, ignoring")

	// ErrUnexpectedType indicates that the type of incoming message is not expected.
	ErrUnexpectedType = errors.New("got unexpected type of message")

	// ErrInvalidVersion indicates that the version of the message specified by the user
	// is not acceptable for the receiver.
	ErrInvalidVersion = errors.New("the version is not acceptable for the receiver")

	// ErrInvalidTEID indicates that the TEID value is different from expected one or
	// not registered in TEIDMap.
	ErrInvalidTEID = errors.New("got invalid TEID")

	// ErrTEIDNotFound indicates that TEID is not registered for the interface specified.
	ErrTEIDNotFound = errors.New("no TEID found")

	// ErrUnknownIMSI indicates that the IMSI is different from expected one.
	ErrUnknownIMSI = errors.New("got unknown IMSI")

	// ErrUnknownAPN indicates that the APN is different from expected one.
	ErrUnknownAPN = errors.New("got unknown APN")

	// ErrTimeout indicates that a handler failed to complete its work due to the
	// absence of messages expected to come from another endpoint.
	ErrTimeout = errors.New("timed out")

	// ErrNoBearerFound indicates that no Bearer found by lookup methods.
	ErrNoBearerFound = errors.New("no Bearer found")

	// ErrNoRemoteAddressFound indicates that no remote address given to send(respond)
	// a message.
	ErrNoRemoteAddressFound = errors.New("no remote address found")

	// ErrDuplicateTEID indicates that the TEID added to a Session already exists.
	// Users should re-generate TEID and add it again.
	ErrDuplicateTEID = errors.New("same TEID cannot exist simultaneously in a Session. Re-generate or request another one")
)

// ErrCauseNotOK indicates that the value in Cause IE is not OK.
type ErrCauseNotOK struct {
	MsgType string
	Cause   uint8
	Msg     string
}

// Error returns error cause with message.
func (e *ErrCauseNotOK) Error() string {
	return fmt.Sprintf("got non-OK Cause: %d in %s; %s", e.Cause, e.MsgType, e.Msg)
}

// ErrRequiredIEMissing indicates that the IE required is missing.
type ErrRequiredIEMissing struct {
	Type uint8
}

// Error returns error with missing IE type.
func (e *ErrRequiredIEMissing) Error() string {
	return fmt.Sprintf("required IE missing: %d", e.Type)
}

// ErrRequiredParameterMissing indicates that no Bearer found by lookup methods.
type ErrRequiredParameterMissing struct {
	Name, Msg string
}

// Error returns missing paramter with message.
func (e *ErrRequiredParameterMissing) Error() string {
	return fmt.Sprintf("required parameter: %s is missing. %s", e.Name, e.Msg)
}
