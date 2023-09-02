// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"encoding/binary"
	"fmt"
)

// Header is a GTPv1 common header.
type Header struct {
	Flags                   uint8
	Type                    uint8
	Length                  uint16
	TEID                    uint32
	SequenceNumber          uint16
	NPDUNumber              uint8
	NextExtensionHeaderType uint8
	ExtensionHeaders        []*ExtensionHeader
	Payload                 []byte
}

// NewHeader creates a new Header.
func NewHeader(flags, mtype uint8, teid uint32, seqnum uint16, payload []byte) *Header {
	h := &Header{
		Flags:                   flags,
		Type:                    mtype,
		TEID:                    teid,
		SequenceNumber:          seqnum,
		NextExtensionHeaderType: ExtHeaderTypeNoMoreExtensionHeaders,
		Payload:                 payload,
	}

	h.SetLength()
	return h
}

// NewHeaderWithNPDUNumber creates a new Header with NPDUNumber.
func NewHeaderWithNPDUNumber(flags, mtype uint8, teid uint32, seqnum uint16, npdu uint8, payload []byte) *Header {
	h := NewHeader(flags, mtype, teid, seqnum, payload)
	h.SetNPDUNumber(npdu)
	return h
}

// NewHeaderWithExtensionHeaders creates a new Header with ExtensionHeaders.
func NewHeaderWithExtensionHeaders(flags, mtype uint8, teid uint32, seqnum uint16, payload []byte, extHdrs ...*ExtensionHeader) *Header {
	h := NewHeader(flags, mtype, teid, seqnum, payload)
	_ = h.AddExtensionHeaders(extHdrs...)
	return h
}

// NewHeaderFlags returns a Header Flag built by its components given as arguments.
func NewHeaderFlags(v, p, e, s, n int) uint8 {
	return uint8(
		((v & 0x7) << 5) | ((p & 0x1) << 4) | ((e & 0x1) << 2) | ((s & 0x1) << 1) | (n & 0x1),
	)
}

// Marshal returns the byte sequence generated from a Header.
func (h *Header) Marshal() ([]byte, error) {
	b := make([]byte, h.MarshalLen())
	if err := h.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (h *Header) MarshalTo(b []byte) error {
	if len(b) < h.MarshalLen() {
		return ErrTooShortToMarshal
	}

	b[0] = h.Flags
	b[1] = h.Type
	binary.BigEndian.PutUint16(b[2:4], h.Length)
	binary.BigEndian.PutUint32(b[4:8], h.TEID)

	offset := 8
	s, pn, e := h.HasSequence(), h.HasNPDUNumber(), h.HasExtensionHeader()
	if s {
		binary.BigEndian.PutUint16(b[offset:offset+2], h.SequenceNumber)
	}

	if pn {
		b[offset+2] = h.NPDUNumber
	}

	if e {
		b[offset+3] = h.NextExtensionHeaderType
	}

	if s || pn || e {
		offset += 4
	}

	if e {
		for _, eh := range h.ExtensionHeaders {
			if err := eh.MarshalTo(b[offset:]); err != nil {
				return err
			}
			offset += eh.MarshalLen()
		}
	}

	copy(b[offset:], h.Payload)
	return nil
}

// ParseHeader decodes given byte sequence as a GTPv1 header.
func ParseHeader(b []byte) (*Header, error) {
	h := &Header{}
	if err := h.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return h, nil
}

// UnmarshalBinary sets the values retrieved from byte sequence in GTPv1 header.
func (h *Header) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 8 {
		return ErrTooShortToParse
	}
	var offset = 4
	h.Flags = b[0]
	h.Type = b[1]
	h.Length = binary.BigEndian.Uint16(b[2:4])

	h.TEID = binary.BigEndian.Uint32(b[4:8])
	offset += 4

	if l < int(h.Length)+8 {
		return ErrTooShortToParse
	}

	s, pn, e := h.HasSequence(), h.HasNPDUNumber(), h.HasExtensionHeader()
	// The optional 4 bytes will append to mandatory GTP header if any one or more S, E flags are set.
	if s || pn || e {
		if h.Length < 4 {
			return ErrTooShortToParse
		}
		// Sequence number must be interpreted only if the S bit is on
		if s {
			h.SequenceNumber = binary.BigEndian.Uint16(b[offset : offset+2])
		}
		// NPDUNumber must be interpreted only if the PN bit is on
		if pn {
			h.NPDUNumber = b[offset+2]
		}
		// Next Extension Header Type must be interpreted only if the E bit is on.
		if e {
			h.NextExtensionHeaderType = b[offset+3]
		}
		offset += 4
	}

	if e {
		var err error
		h.ExtensionHeaders, err = ParseMultiExtensionHeaders(b[offset:])
		if err != nil {
			return err
		}
		h.ExtensionHeaders[0].Type = h.NextExtensionHeaderType

		for _, eh := range h.ExtensionHeaders {
			offset += eh.MarshalLen()
		}
	}

	if offset > int(h.Length)+8 {
		return ErrInvalidLength
	}
	h.Payload = b[offset : 8+h.Length]
	return nil
}

// SetTEID sets the TEIDFlag to 1 and puts the TEID given into TEID field.
func (h *Header) SetTEID(teid uint32) {
	h.Flags |= (1 << 3)
	h.TEID = teid
}

// Sequence returns SequenceNumber in uint16.
func (h *Header) Sequence() uint16 {
	return h.SequenceNumber
}

// HasSequence reports whether a Header has SequenceNumber by checking the flag.
func (h *Header) HasSequence() bool {
	return ((int(h.Flags) >> 1) & 0x1) == 1
}

// SetSequenceNumber sets the SequenceNumber in Header.
func (h *Header) SetSequenceNumber(seq uint16) {
	h.Flags |= 0x02
	h.SequenceNumber = seq
	h.SetLength()
}

// WithSequenceNumber returns the Header with SequenceNumber added.
func (h *Header) WithSequenceNumber(seq uint16) *Header {
	h.SetSequenceNumber(seq)
	return h
}

// HasNPDUNumber reports whether a Header has N-PDU Number by checking the flag.
func (h *Header) HasNPDUNumber() bool {
	return (int(h.Flags) & 0x1) == 1
}

// SetNPDUNumber sets the NPDUNumber in Header.
func (h *Header) SetNPDUNumber(npdu uint8) {
	h.Flags |= 0x01
	h.NPDUNumber = npdu
	h.SetLength()
}

// WithNPDUNumber returns the Header with NPDUNumber added.
func (h *Header) WithNPDUNumber(npdu uint8) *Header {
	h.SetNPDUNumber(npdu)
	return h
}

// HasExtensionHeader reports whether a Header has extension header by checking the flag.
func (h *Header) HasExtensionHeader() bool {
	return ((int(h.Flags) >> 2) & 0x1) == 1
}

// SetNextExtensionHeaderType sets the ExtensionHeaderType in Header.
func (h *Header) SetNextExtensionHeaderType(exhType uint8) {
	h.Flags |= 0x04
	h.NextExtensionHeaderType = exhType
}

// WithExtensionHeaders returns the Header with ExtensionHeaders added.
func (h *Header) WithExtensionHeaders(extHdrs ...*ExtensionHeader) *Header {
	_ = h.AddExtensionHeaders(extHdrs...)
	return h
}

// MarshalLen returns the serial length of Header.
func (h *Header) MarshalLen() int {
	l := len(h.Payload) + 8
	if h.HasSequence() || h.HasNPDUNumber() || h.HasExtensionHeader() {
		l += 4
	}

	if h.HasExtensionHeader() {
		for _, eh := range h.ExtensionHeaders {
			l += eh.MarshalLen()
		}
	}

	return l
}

// SetLength sets the length in Length field.
func (h *Header) SetLength() {
	h.Length = uint16(h.MarshalLen() - 8)
}

// Version returns GTP version in int.
func (h *Header) Version() int {
	return 1
}

// MessageType returns the type of message.
func (h *Header) MessageType() uint8 {
	return h.Type
}

// String returns the GTPv1 header values in human readable format.
func (h *Header) String() string {
	return fmt.Sprintf(
		"{Flags: %#x, Type: %#x, Length: %d, TEID: %#08x, SequenceNumber: %#04x, NPDUNumber: %#x, "+
			"ExtensionHeaderType: %#x, ExtensionHeaders: %v, Payload: %#v}",
		h.Flags,
		h.Type,
		h.Length,
		h.TEID,
		h.SequenceNumber,
		h.NPDUNumber,
		h.NextExtensionHeaderType,
		h.ExtensionHeaders,
		h.Payload,
	)
}

// AddExtensionHeaders adds ExtensionHeader(s) to Header.
//
// This function validates if the next extension header type matches the actual one for safety.
// To create arbitrary(possibly malformed) Header, access ExtensionHeaders field on your own.
func (h *Header) AddExtensionHeaders(extHdrs ...*ExtensionHeader) error {
	if len(extHdrs) < 1 {
		return nil
	}

	h.Flags |= 0x04

	h.NextExtensionHeaderType = extHdrs[0].Type
	next := h.NextExtensionHeaderType
	for _, eh := range extHdrs {
		if next != eh.Type {
			return fmt.Errorf("next type: %x does not match the current type: %x", next, eh.Type)
		}
		h.ExtensionHeaders = append(h.ExtensionHeaders, eh)
		next = eh.NextType
	}

	if next != ExtHeaderTypeNoMoreExtensionHeaders {
		return fmt.Errorf("non-empty next type: %x is specified but does not exist", next)
	}

	h.SetLength()
	return nil
}
