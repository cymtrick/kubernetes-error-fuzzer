// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wal

import (
	"bufio"
	"encoding/binary"
	"hash"
	"io"
	"sync"

	"github.com/coreos/etcd/pkg/crc"
	"github.com/coreos/etcd/wal/walpb"
)

type encoder struct {
	mu sync.Mutex
	bw *bufio.Writer

	crc       hash.Hash32
	buf       []byte
	uint64buf []byte
}

func newEncoder(w io.Writer, prevCrc uint32) *encoder {
	return &encoder{
		bw:  bufio.NewWriter(w),
		crc: crc.New(prevCrc, crcTable),
		// 1MB buffer
		buf:       make([]byte, 1024*1024),
		uint64buf: make([]byte, 8),
	}
}

func (e *encoder) encode(rec *walpb.Record) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.crc.Write(rec.Data)
	rec.Crc = e.crc.Sum32()
	var (
		data []byte
		err  error
		n    int
	)

	if rec.Size() > len(e.buf) {
		data, err = rec.Marshal()
		if err != nil {
			return err
		}
	} else {
		n, err = rec.MarshalTo(e.buf)
		if err != nil {
			return err
		}
		data = e.buf[:n]
	}

	lenField, padBytes := encodeFrameSize(len(data))
	if err = writeUint64(e.bw, lenField, e.uint64buf); err != nil {
		return err
	}

	if padBytes != 0 {
		data = append(data, make([]byte, padBytes)...)
	}
	_, err = e.bw.Write(data)
	return err
}

func encodeFrameSize(dataBytes int) (lenField uint64, padBytes int) {
	lenField = uint64(dataBytes)
	// force 8 byte alignment so length never gets a torn write
	padBytes = (8 - (dataBytes % 8)) % 8
	if padBytes != 0 {
		lenField |= uint64(0x80|padBytes) << 56
	}
	return
}

func (e *encoder) flush() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.bw.Flush()
}

func writeUint64(w io.Writer, n uint64, buf []byte) error {
	// http://golang.org/src/encoding/binary/binary.go
	binary.LittleEndian.PutUint64(buf, n)
	_, err := w.Write(buf)
	return err
}
