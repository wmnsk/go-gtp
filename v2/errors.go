// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v2

import (
	"errors"
	"fmt"

	"github.com/wmnsk/go-gtp/v2/messages"
)

var (
	// ErrTEIDNotFound indicates that TEID is not registered for the interface specified.
	ErrTEIDNotFound = errors.New("no TEID found")

	// ErrTimeout indicates that a handler failed to complete its work due to the
	// absence of messages expected to come from another endpoint.
	ErrTimeout = errors.New("timed out")
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

// Error returns missing parameter with message.
func (e *ErrRequiredParameterMissing) Error() string {
	return fmt.Sprintf("required parameter: %s is missing. %s", e.Name, e.Msg)
}

// ErrUnexpectedType indicates that the type of incoming message is not expected.
type ErrUnexpectedType struct {
	Msg messages.Message
}

// Error returns violating message type.
func (e *ErrUnexpectedType) Error() string {
	return fmt.Sprintf("got unexpected type of message: %T", e.Msg)
}

// ErrUnexpectedIE indicates that the type of incoming message is not expected.
type ErrUnexpectedIE struct {
	IEType uint8
}

// Error returns violating message type.
func (e *ErrUnexpectedIE) Error() string {
	return fmt.Sprintf("got unexpected type of message: %T", e.IEType)
}

// ErrInvalidVersion indicates that the version of the message specified by the user
// is not acceptable for the receiver.
type ErrInvalidVersion struct {
	Version int
}

// Error returns violationg version.
func (e *ErrInvalidVersion) Error() string {
	return fmt.Sprintf("version: %d is not acceptable for the receiver", e.Version)
}

// ErrInvalidSequence indicates that the Sequence Number is invalid.
type ErrInvalidSequence struct {
	Seq uint32
}

// Error returns violating Sequence Number.
func (e *ErrInvalidSequence) Error() string {
	return fmt.Sprintf("got invalid Sequence Number: %d", e.Seq)
}

// ErrInvalidTEID indicates that the TEID value is different from expected one or
// not registered in TEIDMap.
type ErrInvalidTEID struct {
	TEID uint32
}

// Error returns violating TEID.
func (e *ErrInvalidTEID) Error() string {
	return fmt.Sprintf("got invalid TEID: %#08x", e.TEID)
}

// ErrUnknownIMSI indicates that the IMSI is different from expected one.
type ErrUnknownIMSI struct {
	IMSI string
}

// Error returns violating IMSI.
func (e *ErrUnknownIMSI) Error() string {
	return fmt.Sprintf("got unknown IMSI: %s", e.IMSI)
}

// ErrUnknownAPN indicates that the APN is different from expected one.
type ErrUnknownAPN struct {
	APN string
}

// Error returns violating APN.
func (e *ErrUnknownAPN) Error() string {
	return fmt.Sprintf("got unknown APN: %s", e.APN)
}

// ErrNoBearerFound indicates that no Bearer found by lookup methods.
type ErrNoBearerFound struct {
	IMSI string
}

func (e ErrNoBearerFound) Error() string {
	return fmt.Sprintf("no Bearer found: %s", e.IMSI)
}

// ErrNoHandlersFound indicates that the handler func is not registered in *Conn
// for the incoming GTPv2 message. In usual cases this error should not be taken
// as fatal, as the other endpoint can make your program stop working just by
// sending unregistered messages.
type ErrNoHandlersFound struct {
	MsgType string
}

// Error returns violating message type to handle.
func (e *ErrNoHandlersFound) Error() string {
	return fmt.Sprintf("no handlers found for incoming message: %s, ignoring", e.MsgType)
}
