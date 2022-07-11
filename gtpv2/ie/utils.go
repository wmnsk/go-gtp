// Copyright 2019-2022 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

func has8thBit(f uint8) bool {
	return (f&0x80)>>7 == 1
}

func has7thBit(f uint8) bool {
	return (f&0x40)>>6 == 1
}

func has6thBit(f uint8) bool {
	return (f&0x20)>>5 == 1
}

func has5thBit(f uint8) bool {
	return (f&0x010)>>4 == 1
}

func has4thBit(f uint8) bool {
	return (f&0x08)>>3 == 1
}

func has3rdBit(f uint8) bool {
	return (f&0x04)>>2 == 1
}

func has2ndBit(f uint8) bool {
	return (f&0x02)>>1 == 1
}

func has1stBit(f uint8) bool {
	return (f & 0x01) == 1
}
