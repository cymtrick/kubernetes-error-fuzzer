// @generated Code generated by gen-atomicwrapper.

// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package atomic

import (
	"encoding/json"
	"math"
)

// Float64 is an atomic type-safe wrapper for float64 values.
type Float64 struct {
	_ nocmp // disallow non-atomic comparison

	v Uint64
}

var _zeroFloat64 float64

// NewFloat64 creates a new Float64.
func NewFloat64(v float64) *Float64 {
	x := &Float64{}
	if v != _zeroFloat64 {
		x.Store(v)
	}
	return x
}

// Load atomically loads the wrapped float64.
func (x *Float64) Load() float64 {
	return math.Float64frombits(x.v.Load())
}

// Store atomically stores the passed float64.
func (x *Float64) Store(v float64) {
	x.v.Store(math.Float64bits(v))
}

// CAS is an atomic compare-and-swap for float64 values.
func (x *Float64) CAS(o, n float64) bool {
	return x.v.CAS(math.Float64bits(o), math.Float64bits(n))
}

// MarshalJSON encodes the wrapped float64 into JSON.
func (x *Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Load())
}

// UnmarshalJSON decodes a float64 from JSON.
func (x *Float64) UnmarshalJSON(b []byte) error {
	var v float64
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	x.Store(v)
	return nil
}
