// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2016, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package types

import (
	"time"
)

func NewPopulatedTimestamp(r interface {
	Int63() int64
}, easy bool) *Timestamp {
	this := &Timestamp{}
	ns := int64(r.Int63())
	this.Seconds = ns / 1e9
	this.Nanos = int32(ns % 1e9)
	return this
}

func (ts *Timestamp) String() string {
	return TimestampString(ts)
}

func NewPopulatedStdTime(r interface {
	Int63() int64
}, easy bool) *time.Time {
	timestamp := NewPopulatedTimestamp(r, easy)
	t, err := TimestampFromProto(timestamp)
	if err != nil {
		return nil
	}
	return &t
}

func SizeOfStdTime(t time.Time) int {
	ts, err := TimestampProto(t)
	if err != nil {
		return 0
	}
	return ts.Size()
}

func StdTimeMarshal(t time.Time) ([]byte, error) {
	size := SizeOfStdTime(t)
	buf := make([]byte, size)
	_, err := StdTimeMarshalTo(t, buf)
	return buf, err
}

func StdTimeMarshalTo(t time.Time, data []byte) (int, error) {
	ts, err := TimestampProto(t)
	if err != nil {
		return 0, err
	}
	return ts.MarshalTo(data)
}

func StdTimeUnmarshal(t *time.Time, data []byte) error {
	ts := &Timestamp{}
	if err := ts.Unmarshal(data); err != nil {
		return err
	}
	tt, err := TimestampFromProto(ts)
	if err != nil {
		return err
	}
	*t = tt
	return nil
}
