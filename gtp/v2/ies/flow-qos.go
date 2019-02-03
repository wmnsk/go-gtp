// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"github.com/wmnsk/go-gtp/gtp/utils"
)

// NewFlowQoS creates a new FlowQoS IE.
func NewFlowQoS(qci uint8, umbr, dmbr, ugbr, dgbr uint64) *IE {
	i := New(FlowQoS, 0x00, make([]byte, 21))
	i.Payload[0] = qci
	copy(i.Payload[1:6], utils.Uint64To40(umbr))
	copy(i.Payload[6:11], utils.Uint64To40(dmbr))
	copy(i.Payload[11:16], utils.Uint64To40(ugbr))
	copy(i.Payload[16:21], utils.Uint64To40(dgbr))
	return i
}
