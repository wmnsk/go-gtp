// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
	"net"
)

const (
	PCOPPPConfigurationRequest = 0x01
	PCOPPPConfigurationAck     = 0x02
)

// PCOPPP represents a PPP header and its contents used in PCO.
//
// TODO: create another package with full implementation.
type PCOPPP struct {
	Code       uint8
	Identifier uint8
	Length     uint16
	Payload    []byte
}

// NewPCOPPP creates a new PCOPPP.
func NewPCOPPP(code, id uint8, payload []byte) *PCOPPP {
	return &PCOPPP{
		Code:       code,
		Identifier: id,
		Length:     uint16(4 + len(payload)),
		Payload:    payload,
	}
}

// NewPCOPPPWithPAP creates a new PCOPPP with given PAP.
func NewPCOPPPWithPAP(code, id uint8, peer, pass string) *PCOPPP {
	pap, err := NewPAPFields(peer, pass).Marshal()
	if err != nil {
		return nil
	}

	return NewPCOPPP(code, id, pap)
}

// NewPCOPPPWithCHAP creates a new PCOPPP with given CHAP.
func NewPCOPPPWithCHAP(code, id uint8, val []byte, pass string) *PCOPPP {
	pap, err := NewCHAPFields(val, pass).Marshal()
	if err != nil {
		return nil
	}

	return NewPCOPPP(code, id, pap)
}

// NewPCOPPPWithIPCPOptions creates a new PCOPPP with given IPCPOptions.
func NewPCOPPPWithIPCPOptions(code, id uint8, opts ...*IPCPOption) *PCOPPP {
	offset := 0
	b := make([]byte, offset)
	for _, o := range opts {
		l := o.MarshalLen()
		b = append(b, make([]byte, l)...)
		if err := o.MarshalTo(b[offset : offset+l]); err != nil {
			return nil
		}

		offset += l
	}

	return NewPCOPPP(code, id, b)
}

// Marshal serializes PCOPPP.
func (p *PCOPPP) Marshal() ([]byte, error) {
	b := make([]byte, p.MarshalLen())
	if err := p.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes PCOPPP.
func (p *PCOPPP) MarshalTo(b []byte) error {
	if len(b) < 5 {
		return io.ErrUnexpectedEOF
	}
	b[0] = p.Code
	b[1] = p.Identifier
	binary.BigEndian.PutUint16(b[2:4], p.Length)

	copy(b[4:], p.Payload)

	return nil
}

// ParsePCOPPP decodes PCOPPP.
func ParsePCOPPP(b []byte) (*PCOPPP, error) {
	p := &PCOPPP{}
	if err := p.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return p, nil
}

// UnmarshalBinary decodes given bytes into PCOPPP.
func (p *PCOPPP) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 5 {
		return ErrTooShortToParse
	}

	p.Code = b[0]
	p.Identifier = b[1]
	p.Length = binary.BigEndian.Uint16(b[2:4])

	if l < int(p.Length) {
		return io.ErrUnexpectedEOF
	}
	p.Payload = b[4:int(p.Length)]

	return nil
}

// MarshalLen returns the serial length of PCOPPP in int.
func (p *PCOPPP) MarshalLen() int {
	return 4 + len(p.Payload)
}

// PAPFields represents a PAP payload on PPP protocol.
//
// TODO: create another package with full implementation.
type PAPFields struct {
	PeerIDLength   uint8
	PeerID         string
	PasswordLength uint8
	Password       string
}

// NewPAPFields creates a new PAPFields.
func NewPAPFields(id, pass string) *PAPFields {
	return &PAPFields{
		PeerIDLength:   uint8(len([]byte(id))),
		PeerID:         id,
		PasswordLength: uint8(len([]byte(pass))),
		Password:       pass,
	}
}

// Marshal serializes PAPFields.
func (p *PAPFields) Marshal() ([]byte, error) {
	b := make([]byte, p.MarshalLen())
	if err := p.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes PAPFields.
func (p *PAPFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}
	b[0] = p.PeerIDLength
	offset := 1

	if l < offset+int(p.PeerIDLength) {
		return io.ErrUnexpectedEOF
	}
	copy(b[offset:offset+int(p.PeerIDLength)], p.PeerID)
	offset += int(p.PeerIDLength)

	b[offset] = p.PasswordLength
	offset++

	if l < offset+int(p.PasswordLength) {
		return io.ErrUnexpectedEOF
	}
	copy(b[offset:offset+int(p.PasswordLength)], p.Password)

	return nil
}

// ParsePAPFields decodes PAPFields.
func ParsePAPFields(b []byte) (*PAPFields, error) {
	p := &PAPFields{}
	if err := p.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return p, nil
}

// UnmarshalBinary decodes given bytes into PAPFields.
func (p *PAPFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 2 {
		return ErrTooShortToParse
	}

	p.PeerIDLength = b[0]
	offset := 1

	if l < offset+int(p.PeerIDLength) {
		return io.ErrUnexpectedEOF
	}
	p.PeerID = string(b[offset : offset+int(p.PeerIDLength)])
	offset += int(p.PeerIDLength)

	p.PasswordLength = b[offset]
	offset++

	if l < offset+int(p.PasswordLength) {
		return io.ErrUnexpectedEOF
	}
	p.Password = string(b[offset : offset+int(p.PasswordLength)])

	return nil
}

// MarshalLen returns the serial length of PAPFields in int.
func (p *PAPFields) MarshalLen() int {
	return 2 + len([]byte(p.PeerID)) + len([]byte(p.Password))
}

// CHAPFields represents a PAP payload on PPP protocol.
//
// TODO: create another package with full implementation.
type CHAPFields struct {
	ValueSize uint8
	Value     []byte
	Name      string
}

// NewCHAPFields creates a new CHAPFields.
func NewCHAPFields(val []byte, name string) *CHAPFields {
	return &CHAPFields{
		ValueSize: uint8(len(val)),
		Value:     val,
		Name:      name,
	}
}

// Marshal serializes CHAPFields.
func (c *CHAPFields) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes CHAPFields.
func (c *CHAPFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}
	b[0] = c.ValueSize
	offset := 1

	if l < offset+int(c.ValueSize) {
		return io.ErrUnexpectedEOF
	}
	copy(b[offset:offset+int(c.ValueSize)], c.Value)
	offset += int(c.ValueSize)

	copy(b[offset:], c.Name)

	return nil
}

// ParseCHAPFields decodes CHAPFields.
func ParseCHAPFields(b []byte) (*CHAPFields, error) {
	c := &CHAPFields{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return c, nil
}

// UnmarshalBinary decodes given bytes into CHAPFields.
func (c *CHAPFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 2 {
		return ErrTooShortToParse
	}

	c.ValueSize = b[0]
	offset := 1

	if l < offset+int(c.ValueSize) {
		return io.ErrUnexpectedEOF
	}
	c.Value = b[offset : offset+int(c.ValueSize)]
	offset += int(c.ValueSize)

	c.Name = string(b[offset:])

	return nil
}

// MarshalLen returns the serial length of CHAPFields in int.
func (c *CHAPFields) MarshalLen() int {
	return 1 + len(c.Value) + len([]byte(c.Name))
}

// IPCP Options.
//
// TODO: perhaps there are more options but...
const (
	IPCPOptionIPAddress    uint8 = 3
	IPCPOptionMobileIPv4   uint8 = 4
	IPCPOptionPrimaryDNS   uint8 = 129
	IPCPOptionSecondaryDNS uint8 = 131
)

// IPCPOption is a IPCP Option.
type IPCPOption struct {
	Type    uint8
	Length  uint8
	Payload []byte
}

// NewIPCPOption creates an IPCPOption with given IP address.
func NewIPCPOption(typ uint8, payload []byte) *IPCPOption {
	return &IPCPOption{
		Type:    typ,
		Length:  uint8(2 + len(payload)),
		Payload: payload,
	}
}

// NewIPCPOptionIPAddress creates an IPCPOption with given IP address.
func NewIPCPOptionIPAddress(ip net.IP) *IPCPOption {
	v4 := ip.To4()

	// IPv4
	if v4 != nil {
		return NewIPCPOption(IPCPOptionIPAddress, v4)
	}

	// IPv6
	return NewIPCPOption(IPCPOptionIPAddress, ip)
}

// NewIPCPOptionMobileIPv4 creates an IPCPOption with given IP address.
func NewIPCPOptionMobileIPv4(ip net.IP) *IPCPOption {
	v4 := ip.To4()

	// IPv4
	if v4 != nil {
		return NewIPCPOption(IPCPOptionMobileIPv4, v4)
	}

	// IPv6
	return NewIPCPOption(IPCPOptionMobileIPv4, ip)
}

// NewIPCPOptionPrimaryDNS creates an IPCPOption with given IP address.
func NewIPCPOptionPrimaryDNS(ip net.IP) *IPCPOption {
	v4 := ip.To4()

	// IPv4
	if v4 != nil {
		return NewIPCPOption(IPCPOptionPrimaryDNS, v4)
	}

	// IPv6
	return NewIPCPOption(IPCPOptionPrimaryDNS, ip)
}

// NewIPCPOptionSecondaryDNS creates an IPCPOption with given IP address.
func NewIPCPOptionSecondaryDNS(ip net.IP) *IPCPOption {
	v4 := ip.To4()

	// IPv4
	if v4 != nil {
		return NewIPCPOption(IPCPOptionSecondaryDNS, v4)
	}

	// IPv6
	return NewIPCPOption(IPCPOptionSecondaryDNS, ip)
}

// Marshal serializes IPCPOption.
func (o *IPCPOption) Marshal() ([]byte, error) {
	b := make([]byte, o.MarshalLen())
	if err := o.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes IPCPOption.
func (o *IPCPOption) MarshalTo(b []byte) error {
	if len(b) < 3 {
		return io.ErrUnexpectedEOF
	}
	b[0] = o.Type
	b[1] = o.Length

	copy(b[2:], o.Payload)

	return nil
}

// ParseIPCPOption decodes IPCPOption.
func ParseIPCPOption(b []byte) (*IPCPOption, error) {
	o := &IPCPOption{}
	if err := o.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return o, nil
}

// UnmarshalBinary decodes given bytes into IPCPOption.
func (o *IPCPOption) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 3 {
		return ErrTooShortToParse
	}

	o.Type = b[0]
	o.Length = b[1]

	if l < int(o.Length) {
		return io.ErrUnexpectedEOF
	}
	o.Payload = b[2:int(o.Length)]

	return nil
}

// MarshalLen returns the serial length of IPCPOption in int.
func (o *IPCPOption) MarshalLen() int {
	return 2 + len(o.Payload)
}
