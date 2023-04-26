// @generated Code generated by gen-atomicwrapper.

// Copyright (c) 2020-2022 Uber Technologies, Inc.
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
)

// Bool is an atomic type-safe wrapper for bool values.
type Bool struct {
	_ nocmp // disallow non-atomic comparison

	v Uint32
}

var _zeroBool bool

// NewBool creates a new Bool.
func NewBool(val bool) *Bool {
	x := &Bool{}
	if val != _zeroBool {
		x.Store(val)
	}
	return x
}

// Load atomically loads the wrapped bool.
func (x *Bool) Load() bool {
	return truthy(x.v.Load())
}

// Store atomically stores the passed bool.
func (x *Bool) Store(val bool) {
	x.v.Store(boolToInt(val))
}

// CAS is an atomic compare-and-swap for bool values.
//
// Deprecated: Use CompareAndSwap.
func (x *Bool) CAS(old, new bool) (swapped bool) {
	return x.CompareAndSwap(old, new)
}

// CompareAndSwap is an atomic compare-and-swap for bool values.
func (x *Bool) CompareAndSwap(old, new bool) (swapped bool) {
	return x.v.CompareAndSwap(boolToInt(old), boolToInt(new))
}

// Swap atomically stores the given bool and returns the old
// value.
func (x *Bool) Swap(val bool) (old bool) {
	return truthy(x.v.Swap(boolToInt(val)))
}

// MarshalJSON encodes the wrapped bool into JSON.
func (x *Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Load())
}

// UnmarshalJSON decodes a bool from JSON.
func (x *Bool) UnmarshalJSON(b []byte) error {
	var v bool
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	x.Store(v)
	return nil
}
