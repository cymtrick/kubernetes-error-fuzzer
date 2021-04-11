// Copyright (c) 2019 Uber Technologies, Inc.
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

// +build go1.13

package multierr

import "errors"

// As attempts to find the first error in the error list that matches the type
// of the value that target points to.
//
// This function allows errors.As to traverse the values stored on the
// multierr error.
func (merr *multiError) As(target interface{}) bool {
	for _, err := range merr.Errors() {
		if errors.As(err, target) {
			return true
		}
	}
	return false
}

// Is attempts to match the provided error against errors in the error list.
//
// This function allows errors.Is to traverse the values stored on the
// multierr error.
func (merr *multiError) Is(target error) bool {
	for _, err := range merr.Errors() {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}
