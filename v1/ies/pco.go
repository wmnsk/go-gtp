// Copyright 2019 go-gtp authors. All rights reserved.
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

// Serialize serializes ConfigurationProtocolOption.
func (c *ConfigurationProtocolOption) Serialize() ([]byte, error) {
	b := make([]byte, c.Len())
	if err := c.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo serializes ConfigurationProtocolOption.
func (c *ConfigurationProtocolOption) SerializeTo(b []byte) error {
	binary.BigEndian.PutUint16(b[0:2], c.ProtocolID)
	b[2] = c.Length
	if c.Length != 0 {
		copy(b[3:], c.Contents)
	}

	return nil
}

// DecodeConfigurationProtocolOption decodes ConfigurationProtocolOption.
func DecodeConfigurationProtocolOption(b []byte) (*ConfigurationProtocolOption, error) {
	c := &ConfigurationProtocolOption{}
	if err := c.DecodeFromBytes(b); err != nil {
		return nil, err
	}

	return c, nil
}

// DecodeFromBytes decodes given bytes into ConfigurationProtocolOption.
func (c *ConfigurationProtocolOption) DecodeFromBytes(b []byte) error {
	if len(b) < 4 {
		return ErrTooShortToDecode
	}
	c.ProtocolID = binary.BigEndian.Uint16(b[0:2])
	c.Length = b[2]
	if c.Length != 0 {
		copy(c.Contents, b[3:])
	}

	return nil
}

// Len returns the actual length of ConfigurationProtocolOption in int.
func (c *ConfigurationProtocolOption) Len() int {
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

// Serialize serializes PCOPayload.
func (p *PCOPayload) Serialize() ([]byte, error) {
	b := make([]byte, p.Len())
	if err := p.SerializeTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// SerializeTo serializes PCOPayload.
func (p *PCOPayload) SerializeTo(b []byte) error {
	b[0] = (p.ConfigurationProtocol & 0x07) | 0x80
	offset := 1
	for _, opt := range p.ConfigurationProtocolOptions {
		if err := opt.SerializeTo(b[offset:]); err != nil {
			return err
		}
		offset += opt.Len()
	}

	return nil
}

// DecodePCOPayload decodes PCOPayload.
func DecodePCOPayload(b []byte) (*PCOPayload, error) {
	p := &PCOPayload{}
	if err := p.DecodeFromBytes(b); err != nil {
		return nil, err
	}

	return p, nil
}

// DecodeFromBytes decodes given bytes into PCOPayload.
func (p *PCOPayload) DecodeFromBytes(b []byte) error {
	p.ConfigurationProtocol = b[0] & 0x07

	offset := 1
	for {
		if offset >= len(b) {
			return nil
		}
		opt, err := DecodeConfigurationProtocolOption(b[offset:])
		if err != nil {
			return err
		}
		p.ConfigurationProtocolOptions = append(p.ConfigurationProtocolOptions, opt)
	}
}

// Len returns the actual length of PCOPayload in int.
func (p *PCOPayload) Len() int {
	l := 1
	for _, opt := range p.ConfigurationProtocolOptions {
		l += opt.Len()
	}

	return l
}

// NewProtocolConfigurationOptions creates a new ProtocolConfigurationOptions IE.
func NewProtocolConfigurationOptions(configProto uint8, options ...*ConfigurationProtocolOption) *IE {
	pco := NewPCOPayload(configProto, options...)

	i := New(ProtocolConfigurationOptions, make([]byte, pco.Len()))
	if err := pco.SerializeTo(i.Payload); err != nil {
		return nil
	}

	return i
}

// ProtocolConfigurationOptions returns ProtocolConfigurationOptions in
// PCOPayload type if the type of IE matches.
func (i *IE) ProtocolConfigurationOptions() *PCOPayload {
	if i.Type != ProtocolConfigurationOptions {
		return nil
	}

	pco, err := DecodePCOPayload(i.Payload)
	if err != nil {
		return nil
	}
	return pco
}
