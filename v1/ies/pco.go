// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import "encoding/binary"

// ConfigurationProtocolOption represents a Configuration protocol option in PCO.
type ConfigurationProtocolOption struct {
	ProtocolID uint16
	Length     uint8
	Contents   []byte
}

// NewConfigurationProtocolOption creates a new ConfigurationProtocolOption.
func NewConfigurationProtocolOption(pid uint16, contents []byte) *ConfigurationProtocolOption {
	c := &ConfigurationProtocolOption{
		ProtocolID: pid,
		Length:     uint8(len(contents)),
		Contents:   contents,
	}
	return c
}

// Marshal serializes ConfigurationProtocolOption.
func (c *ConfigurationProtocolOption) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes ConfigurationProtocolOption.
func (c *ConfigurationProtocolOption) MarshalTo(b []byte) error {
	binary.BigEndian.PutUint16(b[0:2], c.ProtocolID)
	b[2] = c.Length
	if c.Length != 0 {
		copy(b[3:], c.Contents)
	}

	return nil
}

// ParseConfigurationProtocolOption decodes ConfigurationProtocolOption.
func ParseConfigurationProtocolOption(b []byte) (*ConfigurationProtocolOption, error) {
	c := &ConfigurationProtocolOption{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return c, nil
}

// UnmarshalBinary decodes given bytes into ConfigurationProtocolOption.
func (c *ConfigurationProtocolOption) UnmarshalBinary(b []byte) error {
	if len(b) < 4 {
		return ErrTooShortToParse
	}
	c.ProtocolID = binary.BigEndian.Uint16(b[0:2])
	c.Length = b[2]
	if c.Length != 0 {
		copy(c.Contents, b[3:])
	}

	return nil
}

// MarshalLen returns the serial length of ConfigurationProtocolOption in int.
func (c *ConfigurationProtocolOption) MarshalLen() int {
	return 3 + len(c.Contents)
}

// PCOPayload is a Payload of ProtocolConfigurationPayload IE.
type PCOPayload struct {
	ConfigurationProtocol        uint8
	ConfigurationProtocolOptions []*ConfigurationProtocolOption
}

// NewPCOPayload creates a new PCOPayload.
func NewPCOPayload(configProto uint8, opts ...*ConfigurationProtocolOption) *PCOPayload {
	p := &PCOPayload{ConfigurationProtocol: configProto}
	p.ConfigurationProtocolOptions = append(p.ConfigurationProtocolOptions, opts...)

	return p
}

// Marshal serializes PCOPayload.
func (p *PCOPayload) Marshal() ([]byte, error) {
	b := make([]byte, p.MarshalLen())
	if err := p.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes PCOPayload.
func (p *PCOPayload) MarshalTo(b []byte) error {
	b[0] = (p.ConfigurationProtocol & 0x07) | 0x80
	offset := 1
	for _, opt := range p.ConfigurationProtocolOptions {
		if err := opt.MarshalTo(b[offset:]); err != nil {
			return err
		}
		offset += opt.MarshalLen()
	}

	return nil
}

// ParsePCOPayload decodes PCOPayload.
func ParsePCOPayload(b []byte) (*PCOPayload, error) {
	p := &PCOPayload{}
	if err := p.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return p, nil
}

// UnmarshalBinary decodes given bytes into PCOPayload.
func (p *PCOPayload) UnmarshalBinary(b []byte) error {
	if len(b) == 0 {
		return ErrTooShortToParse
	}

	p.ConfigurationProtocol = b[0] & 0x07

	offset := 1
	for {
		if offset >= len(b) {
			return nil
		}
		opt, err := ParseConfigurationProtocolOption(b[offset:])
		if err != nil {
			return err
		}
		p.ConfigurationProtocolOptions = append(p.ConfigurationProtocolOptions, opt)
		offset += opt.MarshalLen()
	}
}

// MarshalLen returns the serial length of PCOPayload in int.
func (p *PCOPayload) MarshalLen() int {
	l := 1
	for _, opt := range p.ConfigurationProtocolOptions {
		l += opt.MarshalLen()
	}

	return l
}

// NewProtocolConfigurationOptions creates a new ProtocolConfigurationOptions IE.
func NewProtocolConfigurationOptions(configProto uint8, options ...*ConfigurationProtocolOption) *IE {
	pco := NewPCOPayload(configProto, options...)

	i := New(ProtocolConfigurationOptions, make([]byte, pco.MarshalLen()))
	if err := pco.MarshalTo(i.Payload); err != nil {
		return nil
	}

	return i
}

// ProtocolConfigurationOptions returns ProtocolConfigurationOptions in
// PCOPayload type if the type of IE matches.
func (i *IE) ProtocolConfigurationOptions() (*PCOPayload, error) {
	if i.Type != ProtocolConfigurationOptions {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	pco, err := ParsePCOPayload(i.Payload)
	if err != nil {
		return nil, err
	}
	return pco, nil
}

// MustProtocolConfigurationOptions returns ProtocolConfigurationOptions in uint32 if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustProtocolConfigurationOptions() *PCOPayload {
	v, _ := i.ProtocolConfigurationOptions()
	return v
}
