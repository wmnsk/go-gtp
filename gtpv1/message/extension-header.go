// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "fmt"

// ExtensionHeaderType definitions.
const (
	ExtHeaderTypeNoMoreExtensionHeaders                 uint8 = 0b00000000
	ExtHeaderTypeMBMSSupportIndication                  uint8 = 0b00000001
	ExtHeaderTypeMSInfoChangeReportingSupportIndication uint8 = 0b00000010
	ExtHeaderTypeLongPDCPPDUNumber                      uint8 = 0b00000011
	ExtHeaderTypeServiceClassIndicator                  uint8 = 0b00100000
	ExtHeaderTypeUDPPort                                uint8 = 0b01000000
	ExtHeaderTypeRANContainer                           uint8 = 0b10000001
	ExtHeaderTypeLongPDCPPDUNumberRequired              uint8 = 0b10000010
	ExtHeaderTypeXwRANContainer                         uint8 = 0b10000011
	ExtHeaderTypeNRRANContainer                         uint8 = 0b10000100
	ExtHeaderTypePDUSessionContainer                    uint8 = 0b10000101
	ExtHeaderTypePDCPPDUNumber                          uint8 = 0b11000000
	ExtHeaderTypeSuspendRequest                         uint8 = 0b11000001
	ExtHeaderTypeSuspendResponse                        uint8 = 0b11000010
)

// ExtensionHeader represents an Extension Header defined in §5.2, TS 29.281 and §6.1 TS 29.060.
type ExtensionHeader struct {
	Type     uint8 // this doesn't exist in the spec, but it's apparently helpful to have
	Length   uint8
	Content  []byte
	NextType uint8
}

// NewExtensionHeader creates a new ExtensionHeader.
//
// ExtensionHeader struct has its own type while it does not actually exist in the packet.
// Be sure to set an appropriate one - putting a wrong type may cause unexpected errors while
// using method depends on ExtensionHeader struct.
func NewExtensionHeader(typ uint8, content []byte, nextType uint8) *ExtensionHeader {
	eh := &ExtensionHeader{
		Type:     typ,
		Content:  content,
		NextType: nextType,
	}

	eh.SetLength()
	return eh
}

// Marshal returns the byte sequence generated from an ExtensionHeader.
func (e *ExtensionHeader) Marshal() ([]byte, error) {
	b := make([]byte, e.MarshalLen())
	if err := e.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (e *ExtensionHeader) MarshalTo(b []byte) error {
	l := len(b)
	if l < e.MarshalLen() {
		return ErrTooShortToMarshal
	}

	b[0] = e.Length
	offset := int(e.Length)*4 - 1
	if l < offset+1 {
		return ErrTooShortToMarshal
	}
	copy(b[1:offset], e.Content)
	b[offset] = e.NextType

	return nil
}

// ParseExtensionHeader decodes given byte sequence as a ExtensionHeader.
func ParseExtensionHeader(b []byte) (*ExtensionHeader, error) {
	e := &ExtensionHeader{}
	if err := e.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return e, nil
}

// UnmarshalBinary sets the values retrieved from byte sequence in ExtensionHeader.
func (e *ExtensionHeader) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 2 {
		return ErrTooShortToParse
	}

	e.Length = b[0]
	if e.Length == 0 {
		return ErrInvalidLength
	}
	offset := int(e.Length)*4 - 1
	if l < offset+1 {
		return ErrTooShortToParse
	}

	e.Content = b[1:offset]
	e.NextType = b[offset]

	return nil
}

// ParseMultiExtensionHeaders parses given bytes sequence as multiple ExtensionHeaders.
func ParseMultiExtensionHeaders(b []byte) ([]*ExtensionHeader, error) {
	var ehs []*ExtensionHeader
	next := ExtHeaderTypeNoMoreExtensionHeaders
	for {
		eh, err := ParseExtensionHeader(b)
		if err != nil {
			return nil, err
		}
		eh.Type = next

		ehs = append(ehs, eh)
		b = b[eh.MarshalLen():]
		next = eh.NextType
		if next == ExtHeaderTypeNoMoreExtensionHeaders {
			break
		}
		continue
	}
	return ehs, nil
}

// MarshalLen returns the serial length of ExtensionHeader.
func (e *ExtensionHeader) MarshalLen() int {
	return pad4Len(len(e.Content) + 2)
}

// SetLength sets the length calculated from the length of contents to Length field.
func (e *ExtensionHeader) SetLength() {
	e.Length = uint8(pad4Len(len(e.Content)+2) / 4)
}

// String returns an ExtensionHeader fields in human readable format.
func (e *ExtensionHeader) String() string {
	return fmt.Sprintf("{Type: %#x, Length: %d, Content: %#x, NextType: %x}",
		e.Type,
		e.Length,
		e.Content,
		e.NextType,
	)
}

// IsComprehensionRequired reports whether the comprehension of the ExtensionHeader is
// required or not.
func (e *ExtensionHeader) IsComprehensionRequired() bool {
	return e.Type>>7 == 1
}

func pad4Len(n int) int {
	return n + ((4 - n) & 0b11)
}
