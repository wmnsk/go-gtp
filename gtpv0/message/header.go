// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"encoding/binary"
	"fmt"

	"github.com/wmnsk/go-gtp/utils"
)

// Header is a GTPv1 common header.
type Header struct {
	Flags          uint8
	Type           uint8
	Length         uint16
	SequenceNumber uint16
	SndcpNumber    uint8
	FlowLabel      uint16
	TID            uint64
	Payload        []byte
}

// NewHeader creates a new Header.
func NewHeader(flags, mtype uint8, seq, label uint16, tid uint64, payload []byte) *Header {
	h := &Header{
		Flags:          flags,
		Type:           mtype,
		SequenceNumber: seq,
		FlowLabel:      label,
		SndcpNumber:    0xff,
		TID:            tid,
		Payload:        payload,
	}
	h.SetLength()

	return h
}

// HeaderFlags returns a Header Flag built by its components given as arguments.
func HeaderFlags(v, p, s int) uint8 {
	return uint8(
		((v & 0x7) << 5) | ((p & 0x1) << 4) | (s & 0x1) | 0x0e,
	)
}

// Marshal returns the byte sequence generated from an IE instance.
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
	binary.BigEndian.PutUint16(b[4:6], h.SequenceNumber)
	binary.BigEndian.PutUint16(b[6:8], h.FlowLabel)
	binary.BigEndian.PutUint32(b[8:12], uint32(int(h.SndcpNumber)<<24|0xffffff))
	binary.BigEndian.PutUint64(b[12:20], h.TID)
	// two bytes of padding before payload.
	copy(b[20:h.MarshalLen()], h.Payload)
	return nil
}

// ParseHeader Parses given byte sequence as a GTPv1 header.
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
	if l < 20 {
		return ErrTooShortToParse
	}
	h.Flags = b[0]
	h.Type = b[1]
	h.Length = binary.BigEndian.Uint16(b[2:4])
	h.SequenceNumber = binary.BigEndian.Uint16(b[4:6])
	h.FlowLabel = binary.BigEndian.Uint16(b[6:8])
	h.SndcpNumber = b[9]
	h.TID = binary.BigEndian.Uint64(b[12:20])

	if int(h.Length)+20 != l {
		h.Payload = b[20:]
		return nil
	}
	h.Payload = b[20 : 20+h.Length]
	return nil
}

// MarshalLen returns the serial length of Header.
func (h *Header) MarshalLen() int {
	return 20 + len(h.Payload)
}

// SetLength sets the length in Length field.
func (h *Header) SetLength() {
	h.Length = uint16(len(h.Payload))
}

// String returns the GTPv1 header values in human readable format.
func (h *Header) String() string {
	return fmt.Sprintf("{Flags: %#x, Type: %#x, Length: %d, SequenceNumber: %#04x, FlowLabel: %#04x, SndcpNumber: %#02x, TID: %#016x, Payload: %#v}",
		h.Flags,
		h.Type,
		h.Length,
		h.SequenceNumber,
		h.FlowLabel,
		h.SndcpNumber,
		h.TID,
		h.Payload,
	)
}

// tid returns the tid in human-readable string.
func (h *Header) tid() string {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, h.TID)

	return utils.SwappedBytesToStr(b, false)
}

// Version returns the GTP version.
func (h *Header) Version() int {
	return 0
}

// MessageType returns the type of message.
func (h *Header) MessageType() uint8 {
	return h.Type
}
