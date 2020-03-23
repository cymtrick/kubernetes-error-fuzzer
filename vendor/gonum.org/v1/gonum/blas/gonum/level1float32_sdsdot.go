// Code generated by "go generate gonum.org/v1/gonum/blas/gonum”; DO NOT EDIT.

// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonum

import (
	"gonum.org/v1/gonum/internal/asm/f32"
)

// Sdsdot computes the dot product of the two vectors plus a constant
//  alpha + \sum_i x[i]*y[i]
//
// Float32 implementations are autogenerated and not directly tested.
func (Implementation) Sdsdot(n int, alpha float32, x []float32, incX int, y []float32, incY int) float32 {
	if incX == 0 {
		panic(zeroIncX)
	}
	if incY == 0 {
		panic(zeroIncY)
	}
	if n <= 0 {
		if n == 0 {
			return 0
		}
		panic(nLT0)
	}
	if incX == 1 && incY == 1 {
		if len(x) < n {
			panic(shortX)
		}
		if len(y) < n {
			panic(shortY)
		}
		return alpha + float32(f32.DdotUnitary(x[:n], y[:n]))
	}
	var ix, iy int
	if incX < 0 {
		ix = (-n + 1) * incX
	}
	if incY < 0 {
		iy = (-n + 1) * incY
	}
	if ix >= len(x) || ix+(n-1)*incX >= len(x) {
		panic(shortX)
	}
	if iy >= len(y) || iy+(n-1)*incY >= len(y) {
		panic(shortY)
	}
	return alpha + float32(f32.DdotInc(x, y, uintptr(n), uintptr(incX), uintptr(incY), uintptr(ix), uintptr(iy)))
}
