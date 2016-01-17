// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build dragonfly freebsd netbsd openbsd

package unix

const ImplementsGetwd = false

func Getwd() (string, error) { return "", ENOTSUP }
