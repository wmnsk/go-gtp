// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"fmt"
	"io"
)

// ConfigurationProtocol definitions.
const (
	ConfigurationProtocolPPPForUseWithIPPDPTypeOrIPPDNType uint8 = 0b000
)

// NewProtocolConfigurationOptions creates a new ProtocolConfigurationOptions IE.
func NewProtocolConfigurationOptions(proto uint8, options ...*PCOContainer) *IE {
	v := NewProtocolConfigurationOptionsFields(proto, options...)
	b, err := v.Marshal()
	if err != nil {
		return nil
	}

	return New(ProtocolConfigurationOptions, 0x00, b)
}

// ProtocolConfigurationOptions returns ProtocolConfigurationOptions in
// ProtocolConfigurationOptionsFields type if the type of IE matches.
func (i *IE) ProtocolConfigurationOptions() (*ProtocolConfigurationOptionsFields, error) {
	switch i.Type {
	case ProtocolConfigurationOptions:
		if len(i.Payload) < 1 {
			return nil, io.ErrUnexpectedEOF
		}

		return ParseProtocolConfigurationOptionsFields(i.Payload)
	case BearerContext:
		ies, err := i.BearerContext()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve ProtocolConfigurationOptions: %w", err)
		}

		for _, child := range ies {
			if child.Type == ProtocolConfigurationOptions {
				return child.ProtocolConfigurationOptions()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MustProtocolConfigurationOptions returns ProtocolConfigurationOptions in *ProtocolConfigurationOptionsFields, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustProtocolConfigurationOptions() *ProtocolConfigurationOptionsFields {
	v, _ := i.ProtocolConfigurationOptions()
	return v
}

// ProtocolConfigurationOptionsFields is a set of fields in ProtocolConfigurationOptions IE.
type ProtocolConfigurationOptionsFields struct {
	Extension             uint8 // bit 8 of octet 1
	ConfigurationProtocol uint8 // bit 1-3 of octet 1
	ProtocolOrContainers  []*PCOContainer
}

// NewProtocolConfigurationOptionsFields creates a new ProtocolConfigurationOptionsFields.
func NewProtocolConfigurationOptionsFields(proto uint8, opts ...*PCOContainer) *ProtocolConfigurationOptionsFields {
	f := &ProtocolConfigurationOptionsFields{ConfigurationProtocol: proto}
	f.ProtocolOrContainers = append(f.ProtocolOrContainers, opts...)

	return f
}

// Marshal serializes ProtocolConfigurationOptionsFields.
func (f *ProtocolConfigurationOptionsFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes ProtocolConfigurationOptionsFields.
func (f *ProtocolConfigurationOptionsFields) MarshalTo(b []byte) error {
	b[0] = (f.ConfigurationProtocol & 0x07) | 0x80
	offset := 1
	for _, opt := range f.ProtocolOrContainers {
		if err := opt.MarshalTo(b[offset:]); err != nil {
			return err
		}
		offset += opt.MarshalLen()
	}

	return nil
}

// ParseProtocolConfigurationOptionsFields decodes ProtocolConfigurationOptionsFields.
func ParseProtocolConfigurationOptionsFields(b []byte) (*ProtocolConfigurationOptionsFields, error) {
	f := &ProtocolConfigurationOptionsFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into ProtocolConfigurationOptionsFields.
func (f *ProtocolConfigurationOptionsFields) UnmarshalBinary(b []byte) error {
	if len(b) < 1 {
		return ErrTooShortToParse
	}

	f.Extension = (b[0] >> 7) & 0x01
	f.ConfigurationProtocol = b[0] & 0x07

	offset := 1
	for {
		if offset >= len(b) {
			return nil
		}
		opt, err := ParsePCOContainer(b[offset:])
		if err != nil {
			return err
		}
		f.ProtocolOrContainers = append(f.ProtocolOrContainers, opt)
		offset += opt.MarshalLen()
	}
}

// MarshalLen returns the serial length of ProtocolConfigurationOptionsFields in int.
func (f *ProtocolConfigurationOptionsFields) MarshalLen() int {
	l := 1
	for _, opt := range f.ProtocolOrContainers {
		l += opt.MarshalLen()
	}

	return l
}

// ProtocolIdentifier definitions.
//
// [Table 10.5.154/3GPP TS 24.008]
//
// At least the following protocol identifiers (as defined in RFC 3232 [103]) shall be
// supported in this version of the protocol:
// - C021H (LCP);
// - C023H (PAP);
// - C223H (CHAP); and
// - 8021H (IPCP).
const (
	PCOProtocolIdentifierLCP  uint16 = 0xc021
	PCOProtocolIdentifierPAP  uint16 = 0xc023
	PCOProtocolIdentifierCHAP uint16 = 0xc223
	PCOProtocolIdentifierIPCP uint16 = 0x8021
)

// PCOContainer is either of a Configuration protocol option or Additional parameters in PCO,
// which are not distinguishable without meta information(link direction) but fortunately
// the format is the same.
type PCOContainer struct {
	ID       uint16
	Length   uint8
	Contents []byte
}

// NewPCOContainer creates a new PCOContainer.
func NewPCOContainer(pid uint16, contents []byte) *PCOContainer {
	c := &PCOContainer{
		ID:       pid,
		Length:   uint8(len(contents)),
		Contents: contents,
	}
	return c
}

// Marshal serializes PCOContainer.
func (c *PCOContainer) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes PCOContainer.
func (c *PCOContainer) MarshalTo(b []byte) error {
	binary.BigEndian.PutUint16(b[0:2], c.ID)
	b[2] = c.Length
	if c.Length != 0 {
		copy(b[3:], c.Contents)
	}

	return nil
}

// ParsePCOContainer decodes PCOContainer.
func ParsePCOContainer(b []byte) (*PCOContainer, error) {
	c := &PCOContainer{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return c, nil
}

// UnmarshalBinary decodes given bytes into PCOContainer.
func (c *PCOContainer) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 3 {
		return ErrTooShortToParse
	}

	c.ID = binary.BigEndian.Uint16(b[0:2])
	c.Length = b[2]

	if c.Length != 0 && l >= 3+int(c.Length) {
		c.Contents = make([]byte, c.Length)
		copy(c.Contents, b[3:3+int(c.Length)])
	}

	return nil
}

// MarshalLen returns the serial length of PCOContainer in int.
func (c *PCOContainer) MarshalLen() int {
	return 3 + len(c.Contents)
}
