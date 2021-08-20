// Copyright 2019-2021 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"encoding/binary"
	"fmt"
)

const (
	fixedHeaderSize = 8
	// sequence number and padding size
	seqSize = 4
)

// Header is a GTPv1 common header.
type Header struct {
	Flags                   uint8
	Type                    uint8
	Length                  uint16
	TEID                    uint32
	SequenceNumber          uint16
	NextExtensionHeaderType uint8
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
	if h.HasSequence() {
		binary.BigEndian.PutUint16(b[offset:offset+2], h.SequenceNumber)
	}
	// one spare byte between SequenceNumber and NextExtensionHeaderType
	if h.HasExtensionHeader() {
		b[offset+3] = h.NextExtensionHeaderType
	}
	if h.HasSequence() || h.HasExtensionHeader() {
		offset += 4
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
	if l < 11 {
		return ErrTooShortToParse
	}
	var offset = 4
	h.Flags = b[0]
	h.Type = b[1]
	h.Length = binary.BigEndian.Uint16(b[2:4])

	h.TEID = binary.BigEndian.Uint32(b[4:8])
	offset += 4

	// The optional 4 bytes will append to mandatory GTP header if any one or more S, E flags are set.
	if h.HasSequence() || h.HasExtensionHeader() {
		if h.Length < seqSize {
			return ErrTooShortToParse
		}
		// Sequence number must be interpreted only if the S bit is on
		if h.HasSequence() {
			h.SequenceNumber = binary.BigEndian.Uint16(b[offset : offset+2])
		}
		// Next Extension Header Type must be interpreted only if the E bit is on.
		if h.HasExtensionHeader() {
			h.NextExtensionHeaderType = b[offset+3]
		}
		offset += 4
	}

	if int(h.Length)+fixedHeaderSize > l {
		h.Payload = b[offset:]
		return nil
	}
	if fixedHeaderSize+h.Length >= uint16(offset) {
		h.Payload = b[offset : fixedHeaderSize+h.Length]
	} else {
		return ErrInvalidLength
	}

	return nil
}

// SetTEID sets the TEIDFlag to 1 and puts the TEID given into TEID field.
func (h *Header) SetTEID(teid uint32) {
	h.Flags |= (1 << 3)
	h.TEID = teid
}

// HasSequence determines whether a GTP Header has TEID inside by checking the flag.
func (h *Header) HasSequence() bool {
	return ((int(h.Flags) >> 1) & 0x1) == 1
}

// Sequence returns SequenceNumber in uint16.
func (h *Header) Sequence() uint16 {
	return h.SequenceNumber
}

// SetSequenceNumber sets the SequenceNumber in Header.
func (h *Header) SetSequenceNumber(seq uint16) {
	h.SequenceNumber = seq
}

// HasExtensionHeader determines whether a GTP Header has extension header inside by checking the flag.
func (h *Header) HasExtensionHeader() bool {
	return ((int(h.Flags) >> 2) & 0x1) == 1
}

// SetSequenceNumber sets the ExtensionHeaderType in Header.
func (h *Header) SetNextExtensionHeaderType(exhType uint8) {
	h.NextExtensionHeaderType = exhType
}

// MarshalLen returns the serial length of Header.
func (h *Header) MarshalLen() int {
	l := len(h.Payload) + 8
	if h.HasSequence() || h.HasExtensionHeader() {
		l += 4
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
		"{Flags: %#x, Type: %#x, Length: %d, TEID: %#08x,"+
			"SequenceNumber: %#04x, ExtensionHeaderType: %#x, Payload: %#v}",
		h.Flags,
		h.Type,
		h.Length,
		h.TEID,
		h.SequenceNumber,
		h.NextExtensionHeaderType,
		h.Payload,
	)
}
