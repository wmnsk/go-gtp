// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

// TPDU is a TPDU.
type TPDU struct {
	*Header
}

// NewTPDU creates a new G-PDU message.
func NewTPDU(teid uint32, payload []byte) *TPDU {
	t := &TPDU{Header: NewHeader(0x30, MsgTypeTPDU, teid, 0, payload)}

	t.SetLength()
	return t
}

// NewTPDUWithSequence creates a new G-PDU message.
func NewTPDUWithSequence(teid uint32, seq uint16, payload []byte) *TPDU {
	t := &TPDU{Header: NewHeader(0x32, MsgTypeTPDU, teid, seq, payload)}

	t.SetLength()
	return t
}

// Marshal returns the byte sequence generated from a TPDU.
func (t *TPDU) Marshal() ([]byte, error) {
	b := make([]byte, t.MarshalLen())
	if err := t.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (t *TPDU) MarshalTo(b []byte) error {
	if len(b) < t.MarshalLen() {
		return ErrTooShortToMarshal
	}

	t.Header.SetLength()
	return t.Header.MarshalTo(b)
}

// ParseTPDU decodes a given byte sequence as a TPDU.
func ParseTPDU(b []byte) (*TPDU, error) {
	t := &TPDU{}
	if err := t.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return t, nil
}

// UnmarshalBinary decodes a given byte sequence as a TPDU.
func (t *TPDU) UnmarshalBinary(b []byte) error {
	var err error
	t.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}

	return nil
}

// MarshalLen returns the serial length of Data.
func (t *TPDU) MarshalLen() int {
	return t.Header.MarshalLen()
}

// SetLength sets the length in Length field.
func (t *TPDU) SetLength() {
	t.Header.SetLength()
}

// MessageTypeName returns the name of protocol.
func (t *TPDU) MessageTypeName() string {
	return "T-PDU"
}

// TEID returns the TEID in human-readable string.
func (t *TPDU) TEID() uint32 {
	return t.Header.TEID
}

// Decapsulate returns payload as raw []byte.
func (t *TPDU) Decapsulate() []byte {
	return t.Header.Payload
}
