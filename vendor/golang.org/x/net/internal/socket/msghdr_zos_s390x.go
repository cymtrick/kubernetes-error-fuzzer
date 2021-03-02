// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build s390x && zos
// +build s390x,zos

package socket

import "unsafe"

func (h *msghdr) pack(vs []iovec, bs [][]byte, oob []byte, sa []byte) {
	for i := range vs {
		vs[i].set(bs[i])
	}
	if len(vs) > 0 {
		h.Iov = &vs[0]
		h.Iovlen = int32(len(vs))
	}
	if len(oob) > 0 {
		h.Control = (*byte)(unsafe.Pointer(&oob[0]))
		h.Controllen = uint32(len(oob))
	}
	if sa != nil {
		h.Name = (*byte)(unsafe.Pointer(&sa[0]))
		h.Namelen = uint32(len(sa))
	}
}

func (h *msghdr) controllen() int {
	return int(h.Controllen)
}

func (h *msghdr) flags() int {
	return int(h.Flags)
}
