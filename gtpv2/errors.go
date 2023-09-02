// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtpv2

import (
	"errors"
	"fmt"

	"github.com/wmnsk/go-gtp/gtpv2/message"
)

var (
	// ErrTEIDNotFound indicates that TEID is not registered for the interface specified.
	ErrTEIDNotFound = errors.New("no TEID found")

	// ErrTimeout indicates that a handler failed to complete its work due to the
	// absence of message expected to come from another endpoint.
	ErrTimeout = errors.New("timed out")
)

// CauseNotOKError indicates that the value in Cause IE is not OK.
type CauseNotOKError struct {
	MsgType string
	Cause   uint8
	Msg     string
}

// Error returns error cause with message.
func (e *CauseNotOKError) Error() string {
	return fmt.Sprintf("got non-OK Cause: %d in %s; %s", e.Cause, e.MsgType, e.Msg)
}

// RequiredIEMissingError indicates that the IE required is missing.
type RequiredIEMissingError struct {
	Type uint8
}

// Error returns error with missing IE type.
func (e *RequiredIEMissingError) Error() string {
	return fmt.Sprintf("required IE missing: %d", e.Type)
}

// RequiredParameterMissingError indicates that no Bearer found by lookup methods.
type RequiredParameterMissingError struct {
	Name, Msg string
}

// Error returns missing parameter with message.
func (e *RequiredParameterMissingError) Error() string {
	return fmt.Sprintf("required parameter: %s is missing. %s", e.Name, e.Msg)
}

// UnexpectedTypeError indicates that the type of incoming message is not expected.
type UnexpectedTypeError struct {
	Msg message.Message
}

// Error returns violating message type.
func (e *UnexpectedTypeError) Error() string {
	return fmt.Sprintf("got unexpected type of message: %T", e.Msg)
}

// UnexpectedIEError indicates that the type of incoming message is not expected.
type UnexpectedIEError struct {
	IEType uint8
}

// Error returns violating message type.
func (e *UnexpectedIEError) Error() string {
	return fmt.Sprintf("got unexpected type of message: %T", e.IEType)
}

// InvalidVersionError indicates that the version of the message specified by the user
// is not acceptable for the receiver.
type InvalidVersionError struct {
	Version int
}

// Error returns violationg version.
func (e *InvalidVersionError) Error() string {
	return fmt.Sprintf("version: %d is not acceptable for the receiver", e.Version)
}

// InvalidSequenceError indicates that the Sequence Number is invalid.
type InvalidSequenceError struct {
	Seq uint32
}

// Error returns violating Sequence Number.
func (e *InvalidSequenceError) Error() string {
	return fmt.Sprintf("got invalid Sequence Number: %d", e.Seq)
}

// InvalidTEIDError indicates that the TEID value is different from expected one or
// not registered in TEIDMap.
type InvalidTEIDError struct {
	TEID uint32
}

// Error returns violating TEID.
func (e *InvalidTEIDError) Error() string {
	return fmt.Sprintf("got invalid TEID: %#08x", e.TEID)
}

// UnknownIMSIError indicates that the IMSI is different from expected one.
type UnknownIMSIError struct {
	IMSI string
}

// Error returns violating IMSI.
func (e *UnknownIMSIError) Error() string {
	return fmt.Sprintf("got unknown IMSI: %s", e.IMSI)
}

// UnknownAPNError indicates that the APN is different from expected one.
type UnknownAPNError struct {
	APN string
}

// Error returns violating APN.
func (e *UnknownAPNError) Error() string {
	return fmt.Sprintf("got unknown APN: %s", e.APN)
}

// InvalidSessionError indicates that something went wrong with Session.
type InvalidSessionError struct {
	IMSI string
}

// Error returns message with IMSI associated with Session if available.
func (e *InvalidSessionError) Error() string {
	return fmt.Sprintf("invalid session, IMSI: %s", e.IMSI)
}

// BearerNotFoundError indicates that no Bearer found by lookup methods.
type BearerNotFoundError struct {
	IMSI string
}

// Error returns message with IMSI associated with Bearer if available.
func (e *BearerNotFoundError) Error() string {
	return fmt.Sprintf("no Bearer found: %s", e.IMSI)
}

// HandlerNotFoundError indicates that the handler func is not registered in *Conn
// for the incoming GTPv2 message. In usual cases this error should not be taken
// as fatal, as the other endpoint can make your program stop working just by
// sending unregistered message.
type HandlerNotFoundError struct {
	MsgType string
}

// Error returns violating message type to handle.
func (e *HandlerNotFoundError) Error() string {
	return fmt.Sprintf("no handlers found for incoming message: %s, ignoring", e.MsgType)
}
