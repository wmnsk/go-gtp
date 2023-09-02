// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"encoding/binary"
	"fmt"

	"github.com/wmnsk/go-gtp/utils"
)

const (
	fixedHeaderSize  = 4
	seqSpareSize     = 4
	teidSize         = 4
	noTEIDHeaderSize = fixedHeaderSize + seqSpareSize
	teidHeaderSize   = noTEIDHeaderSize + teidSize
)

// Header is a GTPv2 common header
type Header struct {
	Flags          uint8
	Type           uint8
	Length         uint16
	TEID           uint32
	SequenceNumber uint32
	Spare          uint8
	Payload        []byte
}

// NewHeader creates a new Header
func NewHeader(flags, mtype uint8, teid, seqnum uint32, data []byte) *Header {
	h := &Header{
		Flags:          flags,
		Type:           mtype,
		TEID:           teid,
		SequenceNumber: seqnum,
		Spare:          0,
		Payload:        data,
	}
	h.SetLength()

	return h
}

// NewHeaderFlags returns a Header Flag built by its components given as arguments.
func NewHeaderFlags(v, p, t int) uint8 {
	return uint8(
		((v & 0x7) << 5) | ((p & 0x1) << 4) | ((t & 0x1) << 3),
	)
}

// Marshal returns the byte sequence generated from a Header instance.
func (h *Header) Marshal() ([]byte, error) {
	b := make([]byte, h.MarshalLen())
	if err := h.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (h *Header) MarshalTo(b []byte) error {
	b[0] = h.Flags
	b[1] = h.Type
	binary.BigEndian.PutUint16(b[2:4], h.Length)
	offset := 4
	if h.HasTEID() {
		binary.BigEndian.PutUint32(b[offset:offset+4], h.TEID)
		offset += 4
	}
	copy(b[offset:offset+3], utils.Uint32To24(h.SequenceNumber))
	b[offset+3] = h.Spare
	copy(b[offset+4:h.MarshalLen()], h.Payload)

	return nil
}

// ParseHeader decodes given byte sequence as a GTPv2 header.
func ParseHeader(b []byte) (*Header, error) {
	h := &Header{}
	if err := h.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return h, nil
}

// UnmarshalBinary sets the values retrieved from byte sequence in GTPv2 header.
func (h *Header) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 12 {
		return ErrTooShortToParse
	}
	h.Flags = b[0]
	h.Type = b[1]
	h.Length = binary.BigEndian.Uint16(b[2:4])
	if h.Length < seqSpareSize {
		return ErrTooShortToParse
	}
	if h.HasTEID() {
		if h.Length < seqSpareSize+teidSize {
			return ErrTooShortToParse
		}
		h.TEID = binary.BigEndian.Uint32(b[4:8])
		h.SequenceNumber = utils.Uint24To32(b[8:11])
		h.Spare = b[11]

		if int(h.Length)+fixedHeaderSize > l {
			h.Payload = b[teidHeaderSize:]
			return nil
		}
		if fixedHeaderSize+h.Length >= teidHeaderSize {
			h.Payload = b[teidHeaderSize : fixedHeaderSize+h.Length]
		} else {
			return ErrInvalidLength
		}
		return nil
	}
	h.SequenceNumber = utils.Uint24To32(b[4:7])
	h.Spare = b[7]

	if int(h.Length)+fixedHeaderSize > l {
		h.Payload = b[noTEIDHeaderSize:]
		return nil
	}
	if fixedHeaderSize+h.Length >= noTEIDHeaderSize {
		h.Payload = b[noTEIDHeaderSize : fixedHeaderSize+h.Length]
	} else {
		return ErrInvalidLength
	}

	return nil
}

// MarshalLen returns field length in integer.
func (h *Header) MarshalLen() int {
	l := 8 + len(h.Payload)
	if h.HasTEID() {
		l += 4
	}

	return l
}

// SetLength sets the length in Length field.
func (h *Header) SetLength() {
	h.Length = uint16(4 + len(h.Payload))
	if h.HasTEID() {
		h.Length += 4
	}
}

// String returns the GTPv2 header values in human readable format.
func (h *Header) String() string {
	return fmt.Sprintf("{Flags: %#x, Type: %d, Length: %d, TEID: %#x, SequenceNumber: %#x, Spare: %d, Payload: %#v}",
		h.Flags,
		h.Type,
		h.Length,
		h.TEID,
		h.SequenceNumber,
		h.Spare,
		h.Payload,
	)
}

// IsPiggybacking reports whether the message has the trailing(piggybacked) message.
func (h *Header) IsPiggybacking() bool {
	return (int(h.Flags)>>4)&0x01 == 1
}

// SetPiggybacking sets the Piggybacking flag.
//
// The given value should only be 0 or 1. Otherwise it may cause the unexpected result.
func (h *Header) SetPiggybacking(val uint8) {
	h.Flags = (h.Flags & 0xef) | (val & 0x01 << 4)
}

// HasTEID determines whether a GTPv2 has TEID inside by checking the flag.
func (h *Header) HasTEID() bool {
	return (int(h.Flags)>>3)&0x01 == 1
}

func (h *Header) teid() uint32 {
	if !h.HasTEID() {
		return 0
	}
	return h.TEID
}

// SetTEID sets the TEIDFlag to 1 and puts the TEID given into TEID field.
func (h *Header) SetTEID(teid uint32) {
	h.Flags |= (1 << 3)
	h.TEID = teid
}

// Sequence returns SequenceNumber in uint32.
func (h *Header) Sequence() uint32 {
	return h.SequenceNumber
}

// SetSequenceNumber sets the SequenceNumber in Header.
func (h *Header) SetSequenceNumber(seq uint32) {
	h.SequenceNumber = seq
}

// HasMessagePriority reports whether the message has MessagePriority field
func (h *Header) HasMessagePriority() bool {
	return (int(h.Flags)>>2)&0x01 == 1
}

// SetMessagePriority sets the MessagePriorityFlag to 1 and puts the MessagePriority
// given into MessagePriority field.
func (h *Header) SetMessagePriority(mp uint8) {
	h.Flags |= (1 << 2)
	h.Spare = (mp & 0xf0)
}

// MessagePriority returns the value of MessagePriority.
//
// Note that this returns the value set in the field even if the MessagePriorityFlag
// is not set to 1.
func (h *Header) MessagePriority() uint8 {
	return h.Spare & 0xf0
}

// Version returns the GTP version.
func (h *Header) Version() int {
	return 2
}

// MessageType returns the type of messagg.
func (h *Header) MessageType() uint8 {
	return h.Type
}
