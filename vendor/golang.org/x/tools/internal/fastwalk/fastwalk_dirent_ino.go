// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (linux || (darwin && !cgo)) && !appengine
// +build linux darwin,!cgo
// +build !appengine

package fastwalk

import "syscall"

func direntInode(dirent *syscall.Dirent) uint64 {
	return dirent.Ino
}
