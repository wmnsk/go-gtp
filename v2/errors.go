// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v2

import (
	"errors"
	"fmt"

	"github.com/ErvinsK/go-gtp/v2/messages"
)

var (
	// ErrTEIDNotFound indicates that TEID is not registered for the interface specified.
	ErrTEIDNotFound = errors.New("no TEID found")

	// ErrTimeout indicates that a handler failed to complete its work due to the
	// absence of messages expected to come from another endpoint.
	ErrTimeout = errors.New("timed out")
)

// CauseNotOKError indicates that the value in Cause IE is not OK.
type CauseNotOKError struct {
	MsgType string
	Cause   uint8
	Msg     string
}

//x Error returns error cause with message.
func (e *CauseNotOKError) Error() string {
	return fmt.Sprintf("got non-OK Cause: %d in %s; %s", e.Cause, e.MsgType, e.Msg)
}

// RequiredIEMissingError indicates that the IE required is missing.
type RequiredIEMissingError struct {
	Type uint8
}

//x Error returns error with missing IE type.
func (e *RequiredIEMissingError) Error() string {
	return fmt.Sprintf("required IE missing: %d", e.Type)
}

// RequiredParameterMissingError indicates that no Bearer found by lookup methods.
type RequiredParameterMissingError struct {
	Name, Msg string
}

//x Error returns missing parameter with message.
func (e *RequiredParameterMissingError) Error() string {
	return fmt.Sprintf("required parameter: %s is missing. %s", e.Name, e.Msg)
}

// UnexpectedTypeError indicates that the type of incoming message is not expected.
type UnexpectedTypeError struct {
	Msg messages.Message
}

//x Error returns violating message type.
func (e *UnexpectedTypeError) Error() string {
	return fmt.Sprintf("got unexpected type of message: %T", e.Msg)
}

// UnexpectedIEError indicates that the type of incoming message is not expected.
type UnexpectedIEError struct {
	IEType uint8
}

//x Error returns violating message type.
func (e *UnexpectedIEError) Error() string {
	return fmt.Sprintf("got unexpected type of message: %T", e.IEType)
}

// InvalidVersionError indicates that the version of the message specified by the user
// is not acceptable for the receiver.
type InvalidVersionError struct {
	Version int
}

//x Error returns violationg version.
func (e *InvalidVersionError) Error() string {
	return fmt.Sprintf("version: %d is not acceptable for the receiver", e.Version)
}

// InvalidSequenceError indicates that the Sequence Number is invalid.
type InvalidSequenceError struct {
	Seq uint32
}

//x Error returns violating Sequence Number.
func (e *InvalidSequenceError) Error() string {
	return fmt.Sprintf("got invalid Sequence Number: %d", e.Seq)
}

// InvalidTEIDError indicates that the TEID value is different from expected one or
// not registered in TEIDMap.
type InvalidTEIDError struct {
	TEID uint32
}

//x Error returns violating TEID.
func (e *InvalidTEIDError) Error() string {
	return fmt.Sprintf("got invalid TEID: %#08x", e.TEID)
}

// UnknownIMSIError indicates that the IMSI is different from expected one.
type UnknownIMSIError struct {
	IMSI string
}

//x Error returns violating IMSI.
func (e *UnknownIMSIError) Error() string {
	return fmt.Sprintf("got unknown IMSI: %s", e.IMSI)
}

// UnknownAPNError indicates that the APN is different from expected one.
type UnknownAPNError struct {
	APN string
}

//x Error returns violating APN.
func (e *UnknownAPNError) Error() string {
	return fmt.Sprintf("got unknown APN: %s", e.APN)
}

// InvalidSessionError indicates that something went wrong with Session.
type InvalidSessionError struct {
	IMSI string
}

//x Error returns message with IMSI associated with Session if available.
func (e *InvalidSessionError) Error() string {
	return fmt.Sprintf("invalid session, IMSI: %s", e.IMSI)
}

// BearerNotFoundError indicates that no Bearer found by lookup methods.
type BearerNotFoundError struct {
	IMSI string
}

//x Error returns message with IMSI associated with Bearer if available.
func (e *BearerNotFoundError) Error() string {
	return fmt.Sprintf("no Bearer found: %s", e.IMSI)
}

// HandlerNotFoundError indicates that the handler func is not registered in *Conn
// for the incoming GTPv2 message. In usual cases this error should not be taken
// as fatal, as the other endpoint can make your program stop working just by
// sending unregistered messages.
type HandlerNotFoundError struct {
	MsgType string
}

//x Error returns violating message type to handle.
func (e *HandlerNotFoundError) Error() string {
	return fmt.Sprintf("no handlers found for incoming message: %s, ignoring", e.MsgType)
}
