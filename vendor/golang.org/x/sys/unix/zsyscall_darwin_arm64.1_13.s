// go run mkasm_darwin.go arm64
// Code generated by the command above; DO NOT EDIT.

// +build go1.13

#include "textflag.h"
TEXT ·libc_fdopendir_trampoline(SB),NOSPLIT,$0-0
	JMP	libc_fdopendir(SB)
TEXT ·libc_closedir_trampoline(SB),NOSPLIT,$0-0
	JMP	libc_closedir(SB)
TEXT ·libc_readdir_r_trampoline(SB),NOSPLIT,$0-0
	JMP	libc_readdir_r(SB)
