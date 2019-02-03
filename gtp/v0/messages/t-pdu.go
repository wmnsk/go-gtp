// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

// TPDU represents a T-PDU type of GTPv0 message.
type TPDU struct {
	*Header
}

// NewTPDU creates a new TPDU.
func NewTPDU(seq, label uint16, tid uint64, payload []byte) *TPDU {
	t := &TPDU{
		Header: NewHeader(
			0x1e, MsgTypeTPDU, seq, label, tid, payload,
		),
	}
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

	t.Header.Payload = t.Payload
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
	t.Payload = t.Header.Payload

	return nil
}

// Len returns the actual length of Data.
func (t *TPDU) Len() int {
	return t.Header.Len() - len(t.Header.Payload) + len(t.Payload)
}

// SetLength sets the length in Length field.
func (t *TPDU) SetLength() {
	t.Header.Length = uint16(len(t.Payload))
}

// MessageTypeName returns the name of protocol.
func (t *TPDU) MessageTypeName() string {
	return "T-PDU"
}

// TID returns the TID in human-readable string.
func (t *TPDU) TID() string {
	return t.tid()
}
