// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"crypto/rand"
	"encoding/binary"
	"net"
)

// NewTEIDDataI creates a new TEIDDataI IE.
func NewTEIDDataI(teid uint32) *IE {
	return newUint32ValIE(TEIDDataI, teid)
}

// NewTEIDCPlane creates a new TEID C-Plane IE.
func NewTEIDCPlane(teid uint32) *IE {
	return newUint32ValIE(TEIDCPlane, teid)
}

// NewTEIDDataII creates a new TEIDDataII IE.
func NewTEIDDataII(teid uint32) *IE {
	return newUint32ValIE(TEIDDataII, teid)
}

// NewTEIDDataIRandom creates a new TEIDDataI IE with random value.
func NewTEIDDataIRandom() *IE {
	teid := make([]byte, 4)
	rand.Read(teid)
	return NewTEIDDataI(binary.BigEndian.Uint32(teid))
}

// NewTEIDCPlaneRandom creates a new TEID C-Plane IE with random value.
func NewTEIDCPlaneRandom() *IE {
	teid := make([]byte, 4)
	rand.Read(teid)
	return NewTEIDCPlane(binary.BigEndian.Uint32(teid))
}

// NewTEIDDataIIRandom creates a new TEIDDataII IE with random value.
func NewTEIDDataIIRandom() *IE {
	teid := make([]byte, 4)
	rand.Read(teid)
	return NewTEIDDataII(binary.BigEndian.Uint32(teid))
}

// TEID returns TEID value if type matches.
func (i *IE) TEID() uint32 {
	switch i.Type {
	case TEIDCPlane, TEIDDataI, TEIDDataII:
		return binary.BigEndian.Uint32(i.Payload)
	default:
		return 0
	}
}

// NewSelectionMode creates a new SelectionMode IE.
// Note that exactly one of the parameters should be set to true.
// Otherwise, you'll get the unexpected result.
func NewSelectionMode(mode uint8) *IE {
	return newUint8ValIE(SelectionMode, mode)
}

// SelectionMode returns SelectionMode value if type matches.
func (i *IE) SelectionMode() uint8 {
	if i.Type != SelectionMode {
		return 0
	}
	return i.Payload[0]
}

// NewNSAPI creates a new NSAPI IE.
func NewNSAPI(nsapi uint8) *IE {
	return newUint8ValIE(NSAPI, nsapi)
}

// NSAPI returns NSAPI value if type matches.
func (i *IE) NSAPI() uint8 {
	if i.Type != NSAPI {
		return 0
	}
	return i.Payload[0]
}

const (
	pdpTypeETSI uint8 = iota | 0xf0
	pdpTypeIETF
)

// NewEndUserAddress creates a new EndUserAddress IE from the given IP Address in string.
//
// The addr can be either IPv4 or IPv6. If the address type is PPP,
// just put "ppp" in addr parameter or use NewEndUserAddressPPP func instead.
func NewEndUserAddress(addr string) *IE {
	if addr == "ppp" {
		return NewEndUserAddressPPP()
	}
	ip := net.ParseIP(addr)
	v4 := ip.To4()

	// IPv4
	if v4 != nil {
		return newEUAddrV4(v4)
	}

	return newEUAddrV6(ip)
}

// NewEndUserAddressIPv4 creates a new EndUserAddress IE with IPv4.
func NewEndUserAddressIPv4(addr string) *IE {
	v4 := net.ParseIP(addr).To4()
	if v4 == nil {
		return New(EndUserAddress, []byte{0xf1, 0x21})
	}

	return newEUAddrV4(v4)
}

// NewEndUserAddressIPv6 creates a new EndUserAddress IE with IPv6.
func NewEndUserAddressIPv6(addr string) *IE {
	v6 := net.ParseIP(addr).To16()
	if v6 == nil {
		return New(EndUserAddress, []byte{0xf1, 0x57})
	}

	return newEUAddrV6(v6)
}

func newEUAddrV4(v4 []byte) *IE {
	e := New(
		EndUserAddress,
		make([]byte, 6),
	)
	e.Payload[0] = pdpTypeIETF
	e.Payload[1] = 0x21
	copy(e.Payload[2:], v4)

	return e
}

func newEUAddrV6(v6 []byte) *IE {
	e := New(
		EndUserAddress,
		make([]byte, 18),
	)
	e.Payload = make([]byte, 18)
	e.Payload[1] = 0x57
	copy(e.Payload[2:], v6)

	return e
}

// NewEndUserAddressPPP creates a new EndUserAddress IE with PPP.
func NewEndUserAddressPPP() *IE {
	e := New(EndUserAddress, make([]byte, 2))
	e.Payload[0] = pdpTypeETSI
	e.Payload[1] = pdpTypeIETF

	e.SetLength()
	return e
}

// EndUserAddress returns EndUserAddress value if type matches.
func (i *IE) EndUserAddress() []byte {
	if i.Type != EndUserAddress {
		return nil
	}
	return i.Payload
}

// PDPTypeOrganization returns PDPTypeOrganization if type matches.
func (i *IE) PDPTypeOrganization() uint8 {
	if i.Type != EndUserAddress {
		return 0
	}
	return i.Payload[0]
}

// PDPTypeNumber returns PDPTypeNumber if type matches.
func (i *IE) PDPTypeNumber() uint8 {
	if i.Type != EndUserAddress {
		return 0
	}
	return i.Payload[1]
}

// IPAddress returns IPAddress if type matches.
func (i *IE) IPAddress() string {
	switch i.Type {
	case EndUserAddress:
		if i.PDPTypeOrganization() != pdpTypeIETF {
			return ""
		}
		if len(i.Payload) < 3 {
			return ""
		}
		return net.IP(i.Payload[2:]).String()
	case GSNAddress:
		return net.IP(i.Payload).String()
	default:
		return ""
	}
}

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
	for _, opt := range opts {
		p.ConfigurationProtocolOptions = append(p.ConfigurationProtocolOptions, opt)
	}

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

// NewQoSProfile creates a new QoSProfile IE.
// XXX - NOT IMPLEMENTED YET. RETURNS EMPTY IE.
func NewQoSProfile() *IE {
	return New(QoSProfile, []byte{})
}

// QoSProfile returns QoSProfile if type matches.
// XXX - NOT IMPLEMENTED YET. RETURNS NIL.
func (i *IE) QoSProfile() []byte {
	if i.Type != QoSProfile {
		return nil
	}
	return nil
}

// NewCommonFlags creates a new CommonFlags IE.
//
// Note: each flag should be set in 1 or 0.
func NewCommonFlags(dualAddr, upgradeQoS, nrsn, noQoS, mbmsCount, ranReady, mbmsService, prohibitComp int) *IE {
	return New(
		CommonFlags,
		[]byte{uint8(
			dualAddr<<7 | upgradeQoS<<6 | nrsn<<5 | noQoS<<4 | mbmsCount<<3 | ranReady<<2 | mbmsService<<1 | prohibitComp,
		)},
	)
}

// CommonFlags returns CommonFlags value if type matches.
func (i *IE) CommonFlags() uint8 {
	if i.Type != CommonFlags {
		return 0
	}
	return i.Payload[0]
}

// IsDualAddressBearer checks if DualAddressBearer flag exists in CommonFlags.
func (i *IE) IsDualAddressBearer() bool {
	return ((i.CommonFlags() >> 7) & 0x01) != 0
}

// IsUpgradeQoSSupported checks if UpgradeQoSSupported flag exists in CommonFlags.
func (i *IE) IsUpgradeQoSSupported() bool {
	return ((i.CommonFlags() >> 6) & 0x01) != 0
}

// IsNRSN checks if NRSN flag exists in CommonFlags.
func (i *IE) IsNRSN() bool {
	return ((i.CommonFlags() >> 5) & 0x01) != 0
}

// IsNoQoSNegotiation checks if NoQoSNegotiation flag exists in CommonFlags.
func (i *IE) IsNoQoSNegotiation() bool {
	return ((i.CommonFlags() >> 4) & 0x01) != 0
}

// IsMBMSCountingInformation checks if MBMSCountingInformation flag exists in CommonFlags.
func (i *IE) IsMBMSCountingInformation() bool {
	return ((i.CommonFlags() >> 3) & 0x01) != 0
}

// IsRANProceduresReady checks if RANProceduresReady flag exists in CommonFlags.
func (i *IE) IsRANProceduresReady() bool {
	return ((i.CommonFlags() >> 2) & 0x01) != 0
}

// IsMBMSServiceType checks if MBMSServiceType flag exists in CommonFlags.
func (i *IE) IsMBMSServiceType() bool {
	return ((i.CommonFlags() >> 1) & 0x01) != 0
}

// IsProhibitPayloadCompression checks if ProhibitPayloadCompression flag exists in CommonFlags.
func (i *IE) IsProhibitPayloadCompression() bool {
	return (i.CommonFlags() & 0x01) != 0
}
