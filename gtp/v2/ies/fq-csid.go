// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"encoding/binary"
	"encoding/hex"
	"net"
)

// Node-ID Type definitions.
const (
	nodeIDIPv4 uint8 = iota
	nodeIDIPv6
	nodeIDOther
)

// NewFullyQualifiedCSID creates a new FullyQualifiedCSID IE.
func NewFullyQualifiedCSID(nodeID string, csIDs ...uint16) *IE {
	var (
		nid   []byte
		ntype uint8
	)
	ip := net.ParseIP(nodeID)
	if ip == nil {
		var err error
		nid, err = hex.DecodeString(nodeID)
		if err != nil {
			return nil
		}
		ntype = nodeIDOther
	} else if v4 := ip.To4(); v4 != nil {
		nid = v4
		ntype = nodeIDIPv4
	} else {
		nid = ip
		ntype = nodeIDIPv6
	}

	i := New(FullyQualifiedCSID, 0x00, make([]byte, 1+len(nid)+len(csIDs)*2))
	i.Payload[0] = ((ntype << 4) & 0xf0) | uint8(len(csIDs)&0x0f)

	offset := len(nid) + 1
	copy(i.Payload[1:offset], nid)
	for n, csid := range csIDs {
		binary.BigEndian.PutUint16(i.Payload[offset+n*2:offset+n*2+2], csid)
	}
	return i
}

// NodeIDType returns NodeIDType in uint8 if the type of IE matches.
func (i *IE) NodeIDType() uint8 {
	switch i.Type {
	case FullyQualifiedCSID:
		return (i.Payload[0] >> 4) & 0x0f
	default:
		return 0
	}
}

// NodeID returns NodeID in []byte if the type of IE matches.
func (i *IE) NodeID() []byte {
	switch i.Type {
	case FullyQualifiedCSID:
		switch (i.Payload[0] >> 4) & 0x0f {
		case nodeIDIPv4, nodeIDOther:
			return i.Payload[1:5]
		case nodeIDIPv6:
			return i.Payload[1:17]
		default:
			return nil
		}
	default:
		return nil
	}
}

// CSIDs returns CSIDs in []uint16 if the type of IE matches.
func (i *IE) CSIDs() []uint16 {
	switch i.Type {
	case FullyQualifiedCSID:
		offset := 0
		switch (i.Payload[0] >> 4) & 0x0f {
		case nodeIDIPv4, nodeIDOther:
			offset += 5
		case nodeIDIPv6:
			offset += 17
		default:
			return nil
		}

		var csids []uint16
		for {
			if offset >= len(i.Payload) {
				break
			}
			csids = append(csids, binary.BigEndian.Uint16(i.Payload[offset:offset+2]))
		}
		return csids
	default:
		return nil
	}
}
