// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

import (
	"strconv"
)

// TODO: add getter/setters
// TODO: consider using struct

// NewIndication creates a new Indication IE.
// Note that each parameters should be 0 if false and 1 if true. Otherwise,
// the value won't be set as expected in the bitwise operations.
func NewIndication(
	daf, dtf, hi, dfi, oi, isrsi, israi, sgwci,
	sqci, uimsi, cfsi, crsi, ps, pt, si, msv,
	retLoc, pbic, srni, s6af, s4af, mbmdt, israu, ccrsi,
	cprai, arrl, ppoff, ppon, ppsi, csfbi, clii, cpsr,
	nsi, uasi, dtci, bdwi, psci, pcri, aosi, aopi,
	roaai, epcosi, cpopci, pmtmsi, s11tf, pnsi, unaccsi, wpmsi,
	eevrsi, ltemui, ltempi, enbcrsi, tspcmi uint8,
) *IE {
	i := New(Indication, 0x00, make([]byte, 7))
	i.Payload[0] |= sgwci
	i.Payload[0] |= (israi << 1)
	i.Payload[0] |= (isrsi << 2)
	i.Payload[0] |= (oi << 3)
	i.Payload[0] |= (dfi << 4)
	i.Payload[0] |= (hi << 5)
	i.Payload[0] |= (dtf << 6)
	i.Payload[0] |= (daf << 7)

	i.Payload[1] |= msv
	i.Payload[1] |= (si << 1)
	i.Payload[1] |= (pt << 2)
	i.Payload[1] |= (ps << 3)
	i.Payload[1] |= (crsi << 4)
	i.Payload[1] |= (cfsi << 5)
	i.Payload[1] |= (uimsi << 6)
	i.Payload[1] |= (sqci << 7)

	i.Payload[2] |= ccrsi
	i.Payload[2] |= (israu << 1)
	i.Payload[2] |= (mbmdt << 2)
	i.Payload[2] |= (s4af << 3)
	i.Payload[2] |= (s6af << 4)
	i.Payload[2] |= (srni << 5)
	i.Payload[2] |= (pbic << 6)
	i.Payload[2] |= (retLoc << 7)

	i.Payload[3] |= cpsr
	i.Payload[3] |= (clii << 1)
	i.Payload[3] |= (csfbi << 2)
	i.Payload[3] |= (ppsi << 3)
	i.Payload[3] |= (ppon << 4)
	i.Payload[3] |= (ppoff << 5)
	i.Payload[3] |= (arrl << 6)
	i.Payload[3] |= (cprai << 7)

	i.Payload[4] |= aopi
	i.Payload[4] |= (aosi << 1)
	i.Payload[4] |= (pcri << 2)
	i.Payload[4] |= (psci << 3)
	i.Payload[4] |= (bdwi << 4)
	i.Payload[4] |= (dtci << 5)
	i.Payload[4] |= (uasi << 6)
	i.Payload[4] |= (nsi << 7)

	i.Payload[5] |= wpmsi
	i.Payload[5] |= (unaccsi << 1)
	i.Payload[5] |= (pnsi << 2)
	i.Payload[5] |= (s11tf << 3)
	i.Payload[5] |= (pmtmsi << 4)
	i.Payload[5] |= (cpopci << 5)
	i.Payload[5] |= (epcosi << 6)
	i.Payload[5] |= (roaai << 7)

	i.Payload[6] |= (tspcmi << 3)
	i.Payload[6] |= (enbcrsi << 4)
	i.Payload[6] |= (ltempi << 5)
	i.Payload[6] |= (ltemui << 6)
	i.Payload[6] |= (eevrsi << 7)

	return i
}

// NewIndicationFromBitSequence creates a new Indication IE from string-formatted
// sequence of bits. The sequence should look the same as Wireshark appearance.
// The input is to be like "10100001000010000001010100010000100010001000000101000".
func NewIndicationFromBitSequence(bitSequence string) *IE {
	if len(bitSequence) != 53 {
		return nil
	}

	ie := New(Indication, 0x00, make([]byte, 7))
	for i, r := range bitSequence {
		bit, err := strconv.Atoi(string(r))
		if err != nil {
			return nil
		}
		if bit > 1 {
			bit = 1
		}

		// index:
		//   0 =< i < 8 : set bits in first octet
		//   8 =< i < 16: set bits in second octet ...
		// bit shift:
		//   0 =< i < 8 : shift i to left
		//   8 =< i < 16: shift i - 8 to left ...
		ie.Payload[i/8] |= uint8(bit << uint8(8*(i/8+1)-i-1))
	}

	return ie
}

// NewIndicationFromOctets creates a new IndicationFromOctets IE from the set of octets.
func NewIndicationFromOctets(octs ...uint8) *IE {
	ie := New(Indication, 0x00, make([]byte, 0))
	for i, o := range octs {
		if i >= 7 {
			break
		}
		ie.Payload = append(ie.Payload, o)
	}
	ie.SetLength()
	return ie
}
