// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package idna

// This file implements the Punycode algorithm from RFC 3492.

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

// These parameter values are specified in section 5.
//
// All computation is done with int32s, so that overflow behavior is identical
// regardless of whether int is 32-bit or 64-bit.
const (
	base        int32 = 36
	damp        int32 = 700
	initialBias int32 = 72
	initialN    int32 = 128
	skew        int32 = 38
	tmax        int32 = 26
	tmin        int32 = 1
)

// decode decodes a string as specified in section 6.2.
func decode(encoded string) (string, error) {
	if encoded == "" {
		return "", nil
	}
	pos := 1 + strings.LastIndex(encoded, "-")
	if pos == 1 {
		return "", fmt.Errorf("idna: invalid label %q", encoded)
	}
	if pos == len(encoded) {
		return encoded[:len(encoded)-1], nil
	}
	output := make([]rune, 0, len(encoded))
	if pos != 0 {
		for _, r := range encoded[:pos-1] {
			output = append(output, r)
		}
	}
	i, n, bias := int32(0), initialN, initialBias
	for pos < len(encoded) {
		oldI, w := i, int32(1)
		for k := base; ; k += base {
			if pos == len(encoded) {
				return "", fmt.Errorf("idna: invalid label %q", encoded)
			}
			digit, ok := decodeDigit(encoded[pos])
			if !ok {
				return "", fmt.Errorf("idna: invalid label %q", encoded)
			}
			pos++
			i += digit * w
			if i < 0 {
				return "", fmt.Errorf("idna: invalid label %q", encoded)
			}
			t := k - bias
			if t < tmin {
				t = tmin
			} else if t > tmax {
				t = tmax
			}
			if digit < t {
				break
			}
			w *= base - t
			if w >= math.MaxInt32/base {
				return "", fmt.Errorf("idna: invalid label %q", encoded)
			}
		}
		x := int32(len(output) + 1)
		bias = adapt(i-oldI, x, oldI == 0)
		n += i / x
		i %= x
		if n > utf8.MaxRune || len(output) >= 1024 {
			return "", fmt.Errorf("idna: invalid label %q", encoded)
		}
		output = append(output, 0)
		copy(output[i+1:], output[i:])
		output[i] = n
		i++
	}
	return string(output), nil
}

// encode encodes a string as specified in section 6.3 and prepends prefix to
// the result.
//
// The "while h < length(input)" line in the specification becomes "for
// remaining != 0" in the Go code, because len(s) in Go is in bytes, not runes.
func encode(prefix, s string) (string, error) {
	output := make([]byte, len(prefix), len(prefix)+1+2*len(s))
	copy(output, prefix)
	delta, n, bias := int32(0), initialN, initialBias
	b, remaining := int32(0), int32(0)
	for _, r := range s {
		if r < 0x80 {
			b++
			output = append(output, byte(r))
		} else {
			remaining++
		}
	}
	h := b
	if b > 0 {
		output = append(output, '-')
	}
	for remaining != 0 {
		m := int32(0x7fffffff)
		for _, r := range s {
			if m > r && r >= n {
				m = r
			}
		}
		delta += (m - n) * (h + 1)
		if delta < 0 {
			return "", fmt.Errorf("idna: invalid label %q", s)
		}
		n = m
		for _, r := range s {
			if r < n {
				delta++
				if delta < 0 {
					return "", fmt.Errorf("idna: invalid label %q", s)
				}
				continue
			}
			if r > n {
				continue
			}
			q := delta
			for k := base; ; k += base {
				t := k - bias
				if t < tmin {
					t = tmin
				} else if t > tmax {
					t = tmax
				}
				if q < t {
					break
				}
				output = append(output, encodeDigit(t+(q-t)%(base-t)))
				q = (q - t) / (base - t)
			}
			output = append(output, encodeDigit(q))
			bias = adapt(delta, h+1, h == b)
			delta = 0
			h++
			remaining--
		}
		delta++
		n++
	}
	return string(output), nil
}

func decodeDigit(x byte) (digit int32, ok bool) {
	switch {
	case '0' <= x && x <= '9':
		return int32(x - ('0' - 26)), true
	case 'A' <= x && x <= 'Z':
		return int32(x - 'A'), true
	case 'a' <= x && x <= 'z':
		return int32(x - 'a'), true
	}
	return 0, false
}

func encodeDigit(digit int32) byte {
	switch {
	case 0 <= digit && digit < 26:
		return byte(digit + 'a')
	case 26 <= digit && digit < 36:
		return byte(digit + ('0' - 26))
	}
	panic("idna: internal error in punycode encoding")
}

// adapt is the bias adaptation function specified in section 6.1.
func adapt(delta, numPoints int32, firstTime bool) int32 {
	if firstTime {
		delta /= damp
	} else {
		delta /= 2
	}
	delta += delta / numPoints
	k := int32(0)
	for delta > ((base-tmin)*tmax)/2 {
		delta /= base - tmin
		k += base
	}
	return k + (base-tmin+1)*delta/(delta+skew)
}
