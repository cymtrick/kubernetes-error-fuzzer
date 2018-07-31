// Code generated by gocc; DO NOT EDIT.

// This file is dual licensed under CC0 and The gonum license.
//
// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Copyright ©2017 Robin Eklind.
// This file is made available under a Creative Commons CC0 1.0
// Universal Public Domain Dedication.

package parser

import (
	"fmt"
)

type action interface {
	act()
	String() string
}

type (
	accept bool
	shift  int // value is next state index
	reduce int // value is production index
)

func (this accept) act() {}
func (this shift) act()  {}
func (this reduce) act() {}

func (this accept) Equal(that action) bool {
	if _, ok := that.(accept); ok {
		return true
	}
	return false
}

func (this reduce) Equal(that action) bool {
	that1, ok := that.(reduce)
	if !ok {
		return false
	}
	return this == that1
}

func (this shift) Equal(that action) bool {
	that1, ok := that.(shift)
	if !ok {
		return false
	}
	return this == that1
}

func (this accept) String() string { return "accept(0)" }
func (this shift) String() string  { return fmt.Sprintf("shift:%d", this) }
func (this reduce) String() string {
	return fmt.Sprintf("reduce:%d(%s)", this, productionsTable[this].String)
}
