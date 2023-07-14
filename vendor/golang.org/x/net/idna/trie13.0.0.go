// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.16
// +build go1.16

package idna

// appendMapping appends the mapping for the respective rune. isMapped must be
// true. A mapping is a categorization of a rune as defined in UTS #46.
func (c info) appendMapping(b []byte, s string) []byte {
	index := int(c >> indexShift)
	if c&xorBit == 0 {
		p := index
		return append(b, mappings[mappingIndex[p]:mappingIndex[p+1]]...)
	}
	b = append(b, s...)
	if c&inlineXOR == inlineXOR {
		// TODO: support and handle two-byte inline masks
		b[len(b)-1] ^= byte(index)
	} else {
		for p := len(b) - int(xorData[index]); p < len(b); p++ {
			index++
			b[p] ^= xorData[index]
		}
	}
	return b
}
