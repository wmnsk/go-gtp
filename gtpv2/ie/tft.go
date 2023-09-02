// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/wmnsk/go-gtp/utils"
)

// TrafficFlowTemplate returns TrafficFlowTemplate struct if the type of IE matches.
func (i *IE) TrafficFlowTemplate() (*TrafficFlowTemplate, error) {
	switch i.Type {
	case BearerTFT, TrafficAggregateDescription:
		return ParseTrafficFlowTemplate(i.Payload)
	case BearerContext:
		ies, err := i.BearerContext()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve BearerContext: %w", err)
		}

		for _, child := range ies {
			if child.Type == BearerTFT {
				return child.TrafficFlowTemplate()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// TrafficFlowTemplate is a set of fields in BearerTFT IE.
type TrafficFlowTemplate struct {
	OperationCode           uint8
	PacketFilters           []*TFTPacketFilter
	PacketFilterIdentifiers []uint8
	Parameters              []*TFTParameter
}

// NewTrafficFlowTemplate creates a new TrafficFlowTemplate.
func NewTrafficFlowTemplate(op uint8, filters []*TFTPacketFilter, ids []uint8, params []*TFTParameter) *TrafficFlowTemplate {
	var fs []*TFTPacketFilter
	for _, f := range filters {
		if f != nil {
			fs = append(fs, f)
		}
	}

	var ps []*TFTParameter
	for _, p := range params {
		if p != nil {
			ps = append(ps, p)
		}
	}

	return &TrafficFlowTemplate{
		OperationCode:           op,
		PacketFilters:           fs,
		PacketFilterIdentifiers: ids,
		Parameters:              ps,
	}
}

// Marshal serializes TrafficFlowTemplate.
func (f *TrafficFlowTemplate) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes TrafficFlowTemplate.
func (f *TrafficFlowTemplate) MarshalTo(b []byte) error {
	if len(b) < 1 {
		return io.ErrUnexpectedEOF
	}

	// 000. .... = TFT operation code
	// ...0 .... = E bit
	// .... 0000 = Number of packet filters
	op := (f.OperationCode & 0b111) << 5
	e := uint8(len(f.Parameters)*1&0b1) << 4
	pf := len(f.PacketFilters)
	if f.OperationCode == TFTOpDeletePacketFiltersFromExistingTFT {
		pf = len(f.PacketFilterIdentifiers)
	}
	b[0] = op | e | uint8(pf&0b1111)

	offset := 1
	switch f.OperationCode {
	case TFTOpCreateNewTFT,
		TFTOpAddPacketFiltersToExistingTFT,
		TFTOpReplacePacketFiltersInExistingTFT:
		for _, filter := range f.PacketFilters {
			if filter == nil {
				continue
			}
			if err := filter.MarshalTo(b[offset:]); err != nil {
				return fmt.Errorf("failed to marshal Packet Filter: %w", err)
			}
			offset += filter.MarshalLen()
		}
	case TFTOpDeletePacketFiltersFromExistingTFT:
		copy(b[offset:offset+pf], f.PacketFilterIdentifiers)
		offset += pf
	}

	for _, param := range f.Parameters {
		if param == nil {
			continue
		}

		if err := param.MarshalTo(b[offset:]); err != nil {
			return fmt.Errorf("failed to marshal Parameter: %w", err)
		}
		offset += param.MarshalLen()
	}

	return nil
}

// ParseTrafficFlowTemplate decodes TrafficFlowTemplate.
func ParseTrafficFlowTemplate(b []byte) (*TrafficFlowTemplate, error) {
	f := &TrafficFlowTemplate{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into TrafficFlowTemplate.
func (f *TrafficFlowTemplate) UnmarshalBinary(b []byte) error {
	if len(b) < 1 {
		return io.ErrUnexpectedEOF
	}

	f.OperationCode = b[0] >> 5
	hasParams := (b[0] >> 4 & 0b1) == 1
	filterLen := int(b[0] & 0b1111)

	offset := 1
	switch f.OperationCode {
	case TFTOpCreateNewTFT,
		TFTOpAddPacketFiltersToExistingTFT,
		TFTOpReplacePacketFiltersInExistingTFT:
		f.PacketFilters = []*TFTPacketFilter{}
		for i := 0; i < filterLen; i++ {
			filter, err := ParseTFTPacketFilter(b[offset:])
			if err != nil {
				return fmt.Errorf("failed to parse Packet Filter: %w", err)
			}
			f.PacketFilters = append(f.PacketFilters, filter)
			offset += filter.MarshalLen()
		}
	case TFTOpDeletePacketFiltersFromExistingTFT:
		if len(b) < offset+filterLen {
			return io.ErrUnexpectedEOF
		}

		f.PacketFilterIdentifiers = b[offset : offset+filterLen]
		offset += filterLen
	}

	if hasParams {
		params, err := ParseMultiTFTParameters(b[offset:])
		if err != nil {
			return fmt.Errorf("failed to parse Parameters: %w", err)
		}
		f.Parameters = params
	}

	return nil
}

// MarshalLen returns the serial length of TrafficFlowTemplate in int.
func (f *TrafficFlowTemplate) MarshalLen() int {
	l := 1

	for _, filter := range f.PacketFilters {
		if filter == nil {
			continue
		}
		l += filter.MarshalLen()
	}

	l += len(f.PacketFilterIdentifiers)

	for _, param := range f.Parameters {
		if param == nil {
			continue
		}
		l += param.MarshalLen()
	}

	return l
}

// TFT Packet Filter Identifier definitions.
const (
	TFTPFPreRel7TFTFilter uint8 = 0
	TFTPFDownlinkOnly     uint8 = 1
	TFTPFUplinkOnly       uint8 = 2
	TFTPFBidirectional    uint8 = 3
)

// TFTPacketFilter represents a PacketFilter in TFT.
type TFTPacketFilter struct {
	Direction            uint8
	Identifier           uint8
	EvaluationPrecedence uint8
	Length               uint8
	Components           []*TFTPFComponent
}

// NewTFTPacketFilter creates a new TFTPacketFilter.
func NewTFTPacketFilter(dir, id, precedence uint8, comps ...*TFTPFComponent) *TFTPacketFilter {
	pf := &TFTPacketFilter{
		Direction:            dir,
		Identifier:           id,
		EvaluationPrecedence: precedence,
		Components:           comps,
	}
	pf.SetLength()

	return pf
}

// Marshal serializes TFTPacketFilter.
func (p *TFTPacketFilter) Marshal() ([]byte, error) {
	b := make([]byte, p.MarshalLen())
	if err := p.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes TFTPacketFilter into b.
func (p *TFTPacketFilter) MarshalTo(b []byte) error {
	l := len(b)
	if l < 3 {
		return io.ErrUnexpectedEOF
	}

	b[0] = (p.Direction&0b11)<<4 | p.Identifier&0b1111
	b[1] = p.EvaluationPrecedence
	b[2] = p.Length
	offset := 3

	for _, comp := range p.Components {
		n := comp.MarshalLen()
		if l < offset+n {
			return io.ErrUnexpectedEOF
		}

		if err := comp.MarshalTo(b[offset : offset+n]); err != nil {
			return err
		}
		offset += n
	}

	return nil
}

// ParseTFTPacketFilter decodes TFTPacketFilter.
func ParseTFTPacketFilter(b []byte) (*TFTPacketFilter, error) {
	p := &TFTPacketFilter{}
	if err := p.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return p, nil
}

// UnmarshalBinary decodes given bytes into TFTPacketFilter.
func (p *TFTPacketFilter) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 3 {
		return io.ErrUnexpectedEOF
	}

	p.Direction = (b[0] >> 4) & 0b11
	p.Identifier = b[0] & 0b1111
	p.EvaluationPrecedence = b[1]
	p.Length = b[2]
	offset := 3

	n := offset + int(p.Length)
	if l < n {
		return io.ErrUnexpectedEOF
	}
	comps, err := ParseMultiTFTPFComponent(b[offset:n])
	if err != nil {
		return err
	}
	p.Components = comps

	return nil
}

// SetLength sets the length in TFTPacketFilter.
func (p *TFTPacketFilter) SetLength() {
	l := 0
	for _, comp := range p.Components {
		l += comp.MarshalLen()
	}

	p.Length = uint8(l)
}

// MarshalLen returns the serial length of TFTPacketFilter in int.
func (p *TFTPacketFilter) MarshalLen() int {
	l := 3

	for _, comp := range p.Components {
		l += comp.MarshalLen()
	}

	return l
}

// Packet Filter Component Type definitions.
const (
	PFCompIPv4RemoteAddress             uint8 = 0b00010000
	PFCompIPv4LocalAddress              uint8 = 0b00010001
	PFCompIPv6RemoteAddress             uint8 = 0b00100000
	PFCompIPv6RemoteAddressPrefixLength uint8 = 0b00100001
	PFCompIPv6LocalAddressPrefixLength  uint8 = 0b00100011
	PFCompProtocolIdentifierNextHeader  uint8 = 0b00110000
	PFCompSingleLocalPort               uint8 = 0b01000000
	PFCompLocalPortRange                uint8 = 0b01000001
	PFCompSingleRemotePort              uint8 = 0b01010000
	PFCompRemotePortRange               uint8 = 0b01010001
	PFCompSecurityParameterIndex        uint8 = 0b01100000
	PFCompTypeOfServiceTrafficClass     uint8 = 0b01110000
	PFCompFlowLabel                     uint8 = 0b10000000
	PFCompDestinationMACAddress         uint8 = 0b10000001
	PFCompSourceMACAddress              uint8 = 0b10000010
	PFCompDot1QCTAGVID                  uint8 = 0b10000011
	PFCompDot1QSTAGVID                  uint8 = 0b10000100
	PFCompDot1QCTAGPCPDEI               uint8 = 0b10000101
	PFCompDot1QSTAGPCPDEI               uint8 = 0b10000110
	PFCompEthertype                     uint8 = 0b10000111
)

// TFTPFComponent represents a component in Packet Fileter in TFT.
type TFTPFComponent struct {
	Type     uint8
	Contents []byte
}

// NewTFTPFComponent creates a new TFTPFComponent.
func NewTFTPFComponent(t uint8, contents []byte) *TFTPFComponent {
	return &TFTPFComponent{
		Type:     t,
		Contents: contents,
	}
}

// NewTFTPFComponentIPv4RemoteAddress creates a new TFTPFComponent of type IPv4RemoteAddress.
func NewTFTPFComponentIPv4RemoteAddress(ip net.IP, mask net.IPMask) *TFTPFComponent {
	return NewTFTPFComponent(PFCompIPv4RemoteAddress, append(ip.To4(), mask...))
}

// NewTFTPFComponentIPv4LocalAddress creates a new TFTPFComponent of type IPv4LocalAddress.
func NewTFTPFComponentIPv4LocalAddress(ip net.IP, mask net.IPMask) *TFTPFComponent {
	return NewTFTPFComponent(PFCompIPv4LocalAddress, append(ip.To4(), mask...))
}

// NewTFTPFComponentIPv6RemoteAddress creates a new TFTPFComponent of type IPv6RemoteAddress.
func NewTFTPFComponentIPv6RemoteAddress(ip net.IP, mask net.IPMask) *TFTPFComponent {
	return NewTFTPFComponent(PFCompIPv6RemoteAddress, append(ip.To16(), mask...))
}

// NewTFTPFComponentIPv6RemoteAddressPrefixLength creates a new TFTPFComponent of type IPv6RemoteAddressPrefixLength.
func NewTFTPFComponentIPv6RemoteAddressPrefixLength(ip net.IP, prefix uint8) *TFTPFComponent {
	return NewTFTPFComponent(PFCompIPv6RemoteAddressPrefixLength, append(ip.To16(), prefix))
}

// NewTFTPFComponentIPv6LocalAddressPrefixLength creates a new TFTPFComponent of type IPv6LocalAddressPrefixLength.
func NewTFTPFComponentIPv6LocalAddressPrefixLength(ip net.IP, prefix uint8) *TFTPFComponent {
	return NewTFTPFComponent(PFCompIPv6LocalAddressPrefixLength, append(ip.To16(), prefix))
}

// NewTFTPFComponentProtocolIdentifierNextHeader creates a new TFTPFComponent of type ProtocolIdentifierNextHeader.
func NewTFTPFComponentProtocolIdentifierNextHeader(id uint8) *TFTPFComponent {
	return NewTFTPFComponent(PFCompProtocolIdentifierNextHeader, []byte{id})
}

// NewTFTPFComponentSingleLocalPort creates a new TFTPFComponent of type SingleLocalPort.
func NewTFTPFComponentSingleLocalPort(port uint16) *TFTPFComponent {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, port)
	return NewTFTPFComponent(PFCompSingleLocalPort, b)
}

// NewTFTPFComponentLocalPortRange creates a new TFTPFComponent of type LocalPortRange.
func NewTFTPFComponentLocalPortRange(low, high uint16) *TFTPFComponent {
	b := make([]byte, 4)
	binary.BigEndian.PutUint16(b[0:2], low)
	binary.BigEndian.PutUint16(b[2:4], high)
	return NewTFTPFComponent(PFCompLocalPortRange, b)
}

// NewTFTPFComponentSingleRemotePort creates a new TFTPFComponent of type SingleRemotePort.
func NewTFTPFComponentSingleRemotePort(port uint16) *TFTPFComponent {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, port)
	return NewTFTPFComponent(PFCompSingleRemotePort, b)
}

// NewTFTPFComponentRemotePortRange creates a new TFTPFComponent of type RemotePortRange.
func NewTFTPFComponentRemotePortRange(low, high uint16) *TFTPFComponent {
	b := make([]byte, 4)
	binary.BigEndian.PutUint16(b[0:2], low)
	binary.BigEndian.PutUint16(b[2:4], high)
	return NewTFTPFComponent(PFCompRemotePortRange, b)
}

// NewTFTPFComponentSecurityParameterIndex creates a new TFTPFComponent of type SecurityParameterIndex.
func NewTFTPFComponentSecurityParameterIndex(idx uint32) *TFTPFComponent {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, idx)
	return NewTFTPFComponent(PFCompSecurityParameterIndex, b)
}

// NewTFTPFComponentTypeOfServiceTrafficClass creates a new TFTPFComponent of type TypeOfServiceTrafficClass.
func NewTFTPFComponentTypeOfServiceTrafficClass(class, mask uint8) *TFTPFComponent {
	return NewTFTPFComponent(PFCompTypeOfServiceTrafficClass, []byte{class, mask})
}

// NewTFTPFComponentFlowLabel creates a new TFTPFComponent of type FlowLabel.
func NewTFTPFComponentFlowLabel(label uint32) *TFTPFComponent {
	return NewTFTPFComponent(PFCompFlowLabel, utils.Uint32To24(label))
}

// NewTFTPFComponentDestinationMACAddress creates a new TFTPFComponent of type DestinationMACAddress.
func NewTFTPFComponentDestinationMACAddress(mac net.HardwareAddr) *TFTPFComponent {
	return NewTFTPFComponent(PFCompDestinationMACAddress, mac)
}

// NewTFTPFComponentSourceMACAddress creates a new TFTPFComponent of type SourceMACAddress.
func NewTFTPFComponentSourceMACAddress(mac net.HardwareAddr) *TFTPFComponent {
	return NewTFTPFComponent(PFCompSourceMACAddress, mac)
}

// NewTFTPFComponentDot1QCTAGVID creates a new TFTPFComponent of type Dot1QCTAGVID.
func NewTFTPFComponentDot1QCTAGVID(vid uint16) *TFTPFComponent {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, vid)
	return NewTFTPFComponent(PFCompDot1QCTAGVID, b)
}

// NewTFTPFComponentDot1QSTAGVID creates a new TFTPFComponent of type Dot1QSTAGVID.
func NewTFTPFComponentDot1QSTAGVID(vid uint16) *TFTPFComponent {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, vid)
	return NewTFTPFComponent(PFCompDot1QSTAGVID, b)
}

// NewTFTPFComponentDot1QCTAGPCPDEI creates a new TFTPFComponent of type Dot1QCTAGPCPDEI.
func NewTFTPFComponentDot1QCTAGPCPDEI(pcpdei uint8) *TFTPFComponent {
	return NewTFTPFComponent(PFCompDot1QCTAGPCPDEI, []byte{pcpdei})
}

// NewTFTPFComponentDot1QSTAGPCPDEI creates a new TFTPFComponent of type Dot1QSTAGPCPDEI.
func NewTFTPFComponentDot1QSTAGPCPDEI(pcpdei uint8) *TFTPFComponent {
	return NewTFTPFComponent(PFCompDot1QSTAGPCPDEI, []byte{pcpdei})
}

// NewTFTPFComponentEthertype creates a new TFTPFComponent of type Ethertype.
func NewTFTPFComponentEthertype(etype uint16) *TFTPFComponent {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, etype)
	return NewTFTPFComponent(PFCompEthertype, b)
}

// IPv4RemoteAddress returns IPv4RemoteAddress in net.IPNet if the type of component matches.
func (c *TFTPFComponent) IPv4RemoteAddress() (*net.IPNet, error) {
	if c.Type != PFCompIPv4RemoteAddress {
		return nil, ErrInvalidType
	}
	if len(c.Contents) < 8 {
		return nil, io.ErrUnexpectedEOF
	}

	return &net.IPNet{IP: c.Contents[:4], Mask: c.Contents[4:8]}, nil
}

// IPv4LocalAddress returns IPv4LocalAddress in net.IPNet if the type of component matches.
func (c *TFTPFComponent) IPv4LocalAddress() (*net.IPNet, error) {
	if c.Type != PFCompIPv4LocalAddress {
		return nil, ErrInvalidType
	}
	if len(c.Contents) < 8 {
		return nil, io.ErrUnexpectedEOF
	}

	return &net.IPNet{IP: c.Contents[:4], Mask: c.Contents[4:8]}, nil
}

// IPv6RemoteAddress returns IPv6RemoteAddress in net.IPNet if the type of component matches.
func (c *TFTPFComponent) IPv6RemoteAddress() (*net.IPNet, error) {
	if c.Type != PFCompIPv6RemoteAddress {
		return nil, ErrInvalidType
	}
	if len(c.Contents) < 32 {
		return nil, io.ErrUnexpectedEOF
	}

	return &net.IPNet{IP: c.Contents[:16], Mask: c.Contents[17:32]}, nil
}

// IPv6RemoteAddressPrefixLength returns IPv6RemoteAddressPrefixLength in *net.IPNet
// if the type of component matches.
func (c *TFTPFComponent) IPv6RemoteAddressPrefixLength() (*net.IPNet, error) {
	if c.Type != PFCompIPv6RemoteAddressPrefixLength {
		return nil, ErrInvalidType
	}
	if len(c.Contents) < 17 {
		return nil, io.ErrUnexpectedEOF
	}

	ipnet := &net.IPNet{
		IP:   net.IP(c.Contents[:16]),
		Mask: net.CIDRMask(int(c.Contents[17]), 128),
	}
	return ipnet, nil
}

// IPv6LocalAddressPrefixLength returns IPv6LocalAddressPrefixLength in *net.IPNet
// if the type of component matches.
func (c *TFTPFComponent) IPv6LocalAddressPrefixLength() (*net.IPNet, error) {
	if c.Type != PFCompIPv6LocalAddressPrefixLength {
		return nil, ErrInvalidType
	}
	if len(c.Contents) < 17 {
		return nil, io.ErrUnexpectedEOF
	}

	ipnet := &net.IPNet{
		IP:   net.IP(c.Contents[:16]),
		Mask: net.CIDRMask(int(c.Contents[17]), 128),
	}
	return ipnet, nil
}

// ProtocolIdentifierNextHeader returns ProtocolIdentifierNextHeader in uint8
// if the type of component matches.
func (c *TFTPFComponent) ProtocolIdentifierNextHeader() (uint8, error) {
	if c.Type != PFCompProtocolIdentifierNextHeader {
		return 0, ErrInvalidType
	}
	if len(c.Contents) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return c.Contents[0], nil
}

// SingleLocalPort returns SingleLocalPort in uint16 if the type of component matches.
func (c *TFTPFComponent) SingleLocalPort() (uint16, error) {
	if c.Type != PFCompSingleLocalPort {
		return 0, ErrInvalidType
	}
	if len(c.Contents) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint16(c.Contents[:2]), nil
}

// LocalPortRange returns LocalPortRange in two uint16(low, high) if the type of
// component matches.
func (c *TFTPFComponent) LocalPortRange() (uint16, uint16, error) {
	if c.Type != PFCompLocalPortRange {
		return 0, 0, ErrInvalidType
	}
	if len(c.Contents) < 4 {
		return 0, 0, io.ErrUnexpectedEOF
	}

	low := binary.BigEndian.Uint16(c.Contents[0:2])
	high := binary.BigEndian.Uint16(c.Contents[2:4])
	return low, high, nil
}

// SingleRemotePort returns SingleRemotePort in uint16 if the type of component matches.
func (c *TFTPFComponent) SingleRemotePort() (uint16, error) {
	if c.Type != PFCompSingleRemotePort {
		return 0, ErrInvalidType
	}
	if len(c.Contents) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint16(c.Contents[:2]), nil
}

// RemotePortRange returns RemotePortRange in two uint16(low, high) if the type of
// component matches.
func (c *TFTPFComponent) RemotePortRange() (uint16, uint16, error) {
	if c.Type != PFCompRemotePortRange {
		return 0, 0, ErrInvalidType
	}
	if len(c.Contents) < 4 {
		return 0, 0, io.ErrUnexpectedEOF
	}

	low := binary.BigEndian.Uint16(c.Contents[0:2])
	high := binary.BigEndian.Uint16(c.Contents[2:4])
	return low, high, nil
}

// SecurityParameterIndex returns SecurityParameterIndex in uint32 if the type of
// component matches.
func (c *TFTPFComponent) SecurityParameterIndex() (uint32, error) {
	if c.Type != PFCompSecurityParameterIndex {
		return 0, ErrInvalidType
	}
	if len(c.Contents) < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint32(c.Contents[:4]), nil
}

// TypeOfServiceTrafficClass returns TypeOfServiceTrafficClass in two uint8
// (class, mask) if the type of component matches.
func (c *TFTPFComponent) TypeOfServiceTrafficClass() (uint8, uint8, error) {
	if c.Type != PFCompTypeOfServiceTrafficClass {
		return 0, 0, ErrInvalidType
	}
	if len(c.Contents) < 2 {
		return 0, 0, io.ErrUnexpectedEOF
	}

	return c.Contents[0], c.Contents[1], nil
}

// FlowLabel returns FlowLabel in uint32 if the type of component matches.
func (c *TFTPFComponent) FlowLabel() (uint32, error) {
	if c.Type != PFCompFlowLabel {
		return 0, ErrInvalidType
	}
	if len(c.Contents) < 3 {
		return 0, io.ErrUnexpectedEOF
	}

	return utils.Uint24To32(c.Contents[:3]), nil
}

// DestinationMACAddress returns DestinationMACAddress in net.HardwareAddr if
// the type of component matches.
func (c *TFTPFComponent) DestinationMACAddress() (net.HardwareAddr, error) {
	if c.Type != PFCompDestinationMACAddress {
		return nil, ErrInvalidType
	}
	if len(c.Contents) < 6 {
		return nil, io.ErrUnexpectedEOF
	}

	return net.HardwareAddr(c.Contents[:6]), nil
}

// SourceMACAddress returns SourceMACAddress in net.HardwareAddr if the type of
// component matches.
func (c *TFTPFComponent) SourceMACAddress() (net.HardwareAddr, error) {
	if c.Type != PFCompSourceMACAddress {
		return nil, ErrInvalidType
	}
	if len(c.Contents) < 6 {
		return nil, io.ErrUnexpectedEOF
	}

	return net.HardwareAddr(c.Contents[:6]), nil
}

// Dot1QCTAGVID returns Dot1QCTAGVID in uint16 if the type of component matches.
func (c *TFTPFComponent) Dot1QCTAGVID() (uint16, error) {
	if c.Type != PFCompDot1QCTAGVID {
		return 0, ErrInvalidType
	}
	if len(c.Contents) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint16(c.Contents[:2]), nil
}

// Dot1QSTAGVID returns Dot1QSTAGVID in uint16 if the type of component matches.
func (c *TFTPFComponent) Dot1QSTAGVID() (uint16, error) {
	if c.Type != PFCompDot1QSTAGVID {
		return 0, ErrInvalidType
	}
	if len(c.Contents) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint16(c.Contents[:2]), nil
}

// Dot1QCTAGPCPDEI returns Dot1QCTAGPCPDEI in uint8 if the type of component matches.
func (c *TFTPFComponent) Dot1QCTAGPCPDEI() (uint8, error) {
	if c.Type != PFCompDot1QCTAGPCPDEI {
		return 0, ErrInvalidType
	}
	if len(c.Contents) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return c.Contents[0], nil
}

// Dot1QSTAGPCPDEI returns Dot1QSTAGPCPDEI in uint8 if the type of component matches.
func (c *TFTPFComponent) Dot1QSTAGPCPDEI() (uint8, error) {
	if c.Type != PFCompDot1QSTAGPCPDEI {
		return 0, ErrInvalidType
	}
	if len(c.Contents) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return c.Contents[0], nil
}

// Ethertype returns Ethertype in uint16 if the type of component matches.
func (c *TFTPFComponent) Ethertype() (uint16, error) {
	if c.Type != PFCompEthertype {
		return 0, ErrInvalidType
	}
	if len(c.Contents) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint16(c.Contents[:2]), nil
}

// Marshal serializes TFTPFComponent.
func (c *TFTPFComponent) Marshal() ([]byte, error) {
	b := make([]byte, c.MarshalLen())
	if err := c.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes TFTPFComponent into b.
func (c *TFTPFComponent) MarshalTo(b []byte) error {
	if len(b) < 1+len(c.Contents) {
		return io.ErrUnexpectedEOF
	}

	b[0] = c.Type
	copy(b[1:1+len(c.Contents)], c.Contents)

	return nil
}

// ParseTFTPFComponent decodes TFTPFComponent.
func ParseTFTPFComponent(b []byte) (*TFTPFComponent, error) {
	c := &TFTPFComponent{}
	if err := c.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return c, nil
}

// ParseMultiTFTPFComponent decodes TFTPFComponent.
func ParseMultiTFTPFComponent(b []byte) ([]*TFTPFComponent, error) {
	var comps []*TFTPFComponent
	for {
		if len(b) == 0 {
			break
		}

		i, err := ParseTFTPFComponent(b)
		if err != nil {
			return nil, err
		}
		comps = append(comps, i)
		b = b[i.MarshalLen():]
	}
	return comps, nil
}

// UnmarshalBinary decodes given bytes into TFTPFComponent.
func (c *TFTPFComponent) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	c.Type = b[0]

	n := 0
	switch c.Type {
	case PFCompIPv4RemoteAddress, PFCompIPv4LocalAddress:
		n = 8
	case PFCompIPv6RemoteAddress:
		n = 32
	case PFCompIPv6RemoteAddressPrefixLength, PFCompIPv6LocalAddressPrefixLength:
		n = 17
	case PFCompProtocolIdentifierNextHeader:
		n = 1
	case PFCompSingleLocalPort, PFCompSingleRemotePort:
		n = 2
	case PFCompLocalPortRange, PFCompRemotePortRange:
		n = 4
	case PFCompSecurityParameterIndex:
		n = 4
	case PFCompTypeOfServiceTrafficClass:
		n = 2
	case PFCompFlowLabel:
		n = 3
	case PFCompDestinationMACAddress, PFCompSourceMACAddress:
		n = 6
	case PFCompDot1QCTAGVID, PFCompDot1QSTAGVID:
		n = 2
	case PFCompDot1QCTAGPCPDEI, PFCompDot1QSTAGPCPDEI:
		n = 1
	case PFCompEthertype:
		n = 2
	}

	if l < 1+n {
		return io.ErrUnexpectedEOF
	}
	c.Contents = b[1 : 1+n]
	return nil
}

// MarshalLen returns the serial length of TFTPFComponent in int.
func (c *TFTPFComponent) MarshalLen() int {
	return 1 + len(c.Contents)
}

// TFT Parameter Identifier definitions.
const (
	TFTParamIDAuthorizationToken      uint8 = 1
	TFTParamIDFlowIdentifier          uint8 = 2
	TFTParamIDPacketFileterIdentifier uint8 = 3
)

// TFTParameter represents a Parameter in TFT.
type TFTParameter struct {
	Identifier uint8
	Length     uint8
	Contents   []byte
}

// NewTFTParameter creates a new TFTParameter.
func NewTFTParameter(id uint8, contents []byte) *TFTParameter {
	return &TFTParameter{
		Identifier: id,
		Length:     uint8(len(contents)),
		Contents:   contents,
	}
}

// Marshal serializes TFTParameter.
func (p *TFTParameter) Marshal() ([]byte, error) {
	b := make([]byte, p.MarshalLen())
	if err := p.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes TFTParameter into b.
func (p *TFTParameter) MarshalTo(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}

	b[0] = p.Identifier
	b[1] = p.Length
	offset := 2

	if l < offset+len(p.Contents) {
		return io.ErrUnexpectedEOF
	}
	copy(b[offset:offset+len(p.Contents)], p.Contents)

	return nil
}

// ParseTFTParameter decodes TFTParameter.
func ParseTFTParameter(b []byte) (*TFTParameter, error) {
	p := &TFTParameter{}
	if err := p.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return p, nil
}

// ParseMultiTFTParameters decodes TFTParameter.
func ParseMultiTFTParameters(b []byte) ([]*TFTParameter, error) {
	var params []*TFTParameter
	for {
		if len(b) == 0 {
			break
		}

		i, err := ParseTFTParameter(b)
		if err != nil {
			return nil, err
		}
		params = append(params, i)
		b = b[i.MarshalLen():]
	}
	return params, nil
}

// UnmarshalBinary decodes given bytes into TFTParameter.
func (p *TFTParameter) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}

	p.Identifier = b[0]
	p.Length = b[1]
	if l < 2+int(p.Length) {
		return io.ErrUnexpectedEOF
	}
	p.Contents = b[2 : 2+int(p.Length)]

	return nil
}

// SetLength sets the length in TFTParameter.
func (p *TFTParameter) SetLength() {
	p.Length = uint8(len(p.Contents))
}

// MarshalLen returns the serial length of TFTParameter in int.
func (p *TFTParameter) MarshalLen() int {
	return 2 + len(p.Contents)
}
