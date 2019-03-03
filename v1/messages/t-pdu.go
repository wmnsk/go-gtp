// Copyright 2019 go-gtp authors. All rights reserved.
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

// Serialize returns the byte sequence generated from a TPDU.
func (t *TPDU) Serialize() ([]byte, error) {
	b := make([]byte, t.Len())
	if err := t.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo puts the byte sequence in the byte array given as b.
func (t *TPDU) SerializeTo(b []byte) error {
	if len(b) < t.Len() {
		return ErrTooShortToSerialize
	}

	t.Header.SetLength()
	return t.Header.SerializeTo(b)
}

// DecodeTPDU decodes a given byte sequence as a TPDU.
func DecodeTPDU(b []byte) (*TPDU, error) {
	t := &TPDU{}
	if err := t.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return t, nil
}

// DecodeFromBytes decodes a given byte sequence as a TPDU.
func (t *TPDU) DecodeFromBytes(b []byte) error {
	var err error
	t.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	return nil
}

// Len returns the actual length of Data.
func (t *TPDU) Len() int {
	return t.Header.Len()
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
