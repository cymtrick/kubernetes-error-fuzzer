// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This code was translated into a form compatible with 6a from the public
// domain sources in SUPERCOP: https://bench.cr.yp.to/supercop.html

// +build amd64,!gccgo,!appengine

#include "const_amd64.h"

// func square(out, in *[5]uint64)
TEXT ·square(SB),7,$0-16
	MOVQ out+0(FP), DI
	MOVQ in+8(FP), SI

	MOVQ 0(SI),AX
	MULQ 0(SI)
	MOVQ AX,CX
	MOVQ DX,R8
	MOVQ 0(SI),AX
	SHLQ $1,AX
	MULQ 8(SI)
	MOVQ AX,R9
	MOVQ DX,R10
	MOVQ 0(SI),AX
	SHLQ $1,AX
	MULQ 16(SI)
	MOVQ AX,R11
	MOVQ DX,R12
	MOVQ 0(SI),AX
	SHLQ $1,AX
	MULQ 24(SI)
	MOVQ AX,R13
	MOVQ DX,R14
	MOVQ 0(SI),AX
	SHLQ $1,AX
	MULQ 32(SI)
	MOVQ AX,R15
	MOVQ DX,BX
	MOVQ 8(SI),AX
	MULQ 8(SI)
	ADDQ AX,R11
	ADCQ DX,R12
	MOVQ 8(SI),AX
	SHLQ $1,AX
	MULQ 16(SI)
	ADDQ AX,R13
	ADCQ DX,R14
	MOVQ 8(SI),AX
	SHLQ $1,AX
	MULQ 24(SI)
	ADDQ AX,R15
	ADCQ DX,BX
	MOVQ 8(SI),DX
	IMUL3Q $38,DX,AX
	MULQ 32(SI)
	ADDQ AX,CX
	ADCQ DX,R8
	MOVQ 16(SI),AX
	MULQ 16(SI)
	ADDQ AX,R15
	ADCQ DX,BX
	MOVQ 16(SI),DX
	IMUL3Q $38,DX,AX
	MULQ 24(SI)
	ADDQ AX,CX
	ADCQ DX,R8
	MOVQ 16(SI),DX
	IMUL3Q $38,DX,AX
	MULQ 32(SI)
	ADDQ AX,R9
	ADCQ DX,R10
	MOVQ 24(SI),DX
	IMUL3Q $19,DX,AX
	MULQ 24(SI)
	ADDQ AX,R9
	ADCQ DX,R10
	MOVQ 24(SI),DX
	IMUL3Q $38,DX,AX
	MULQ 32(SI)
	ADDQ AX,R11
	ADCQ DX,R12
	MOVQ 32(SI),DX
	IMUL3Q $19,DX,AX
	MULQ 32(SI)
	ADDQ AX,R13
	ADCQ DX,R14
	MOVQ $REDMASK51,SI
	SHLQ $13,R8:CX
	ANDQ SI,CX
	SHLQ $13,R10:R9
	ANDQ SI,R9
	ADDQ R8,R9
	SHLQ $13,R12:R11
	ANDQ SI,R11
	ADDQ R10,R11
	SHLQ $13,R14:R13
	ANDQ SI,R13
	ADDQ R12,R13
	SHLQ $13,BX:R15
	ANDQ SI,R15
	ADDQ R14,R15
	IMUL3Q $19,BX,DX
	ADDQ DX,CX
	MOVQ CX,DX
	SHRQ $51,DX
	ADDQ R9,DX
	ANDQ SI,CX
	MOVQ DX,R8
	SHRQ $51,DX
	ADDQ R11,DX
	ANDQ SI,R8
	MOVQ DX,R9
	SHRQ $51,DX
	ADDQ R13,DX
	ANDQ SI,R9
	MOVQ DX,AX
	SHRQ $51,DX
	ADDQ R15,DX
	ANDQ SI,AX
	MOVQ DX,R10
	SHRQ $51,DX
	IMUL3Q $19,DX,DX
	ADDQ DX,CX
	ANDQ SI,R10
	MOVQ CX,0(DI)
	MOVQ R8,8(DI)
	MOVQ R9,16(DI)
	MOVQ AX,24(DI)
	MOVQ R10,32(DI)
	RET
