// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewNodeType creates a new NodeType IE.
func NewNodeType(nodeType uint8) *IE {
	return newUint8ValIE(NodeType, nodeType)
}

// NodeType returns NodeType in uint8 if the type of IE matches.
func (i *IE) NodeType() uint8 {
	if i.Type != NodeType {
		return 0
	}

	return i.Payload[0]
}
