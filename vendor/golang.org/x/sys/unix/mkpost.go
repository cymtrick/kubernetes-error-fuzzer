// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// mkpost processes the output of cgo -godefs to
// modify the generated types. It is used to clean up
// the sys API in an architecture specific manner.
//
// mkpost is run after cgo -godefs; see README.md.
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	// Get the OS and architecture (using GOARCH_TARGET if it exists)
	goos := os.Getenv("GOOS")
	goarch := os.Getenv("GOARCH_TARGET")
	if goarch == "" {
		goarch = os.Getenv("GOARCH")
	}
	// Check that we are using the Docker-based build system if we should be.
	if goos == "linux" {
		if os.Getenv("GOLANG_SYS_BUILD") != "docker" {
			os.Stderr.WriteString("In the Docker-based build system, mkpost should not be called directly.\n")
			os.Stderr.WriteString("See README.md\n")
			os.Exit(1)
		}
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// Intentionally export __val fields in Fsid and Sigset_t
	valRegex := regexp.MustCompile(`type (Fsid|Sigset_t) struct {(\s+)X__val(\s+\S+\s+)}`)
	b = valRegex.ReplaceAll(b, []byte("type $1 struct {${2}Val$3}"))

	// Intentionally export __fds_bits field in FdSet
	fdSetRegex := regexp.MustCompile(`type (FdSet) struct {(\s+)X__fds_bits(\s+\S+\s+)}`)
	b = fdSetRegex.ReplaceAll(b, []byte("type $1 struct {${2}Bits$3}"))

	// If we have empty Ptrace structs, we should delete them. Only s390x emits
	// nonempty Ptrace structs.
	ptraceRexexp := regexp.MustCompile(`type Ptrace((Psw|Fpregs|Per) struct {\s*})`)
	b = ptraceRexexp.ReplaceAll(b, nil)

	// Replace the control_regs union with a blank identifier for now.
	controlRegsRegex := regexp.MustCompile(`(Control_regs)\s+\[0\]uint64`)
	b = controlRegsRegex.ReplaceAll(b, []byte("_ [0]uint64"))

	// Remove fields that are added by glibc
	// Note that this is unstable as the identifers are private.
	removeFieldsRegex := regexp.MustCompile(`X__glibc\S*`)
	b = removeFieldsRegex.ReplaceAll(b, []byte("_"))

	// Convert [65]int8 to [65]byte in Utsname members to simplify
	// conversion to string; see golang.org/issue/20753
	convertUtsnameRegex := regexp.MustCompile(`((Sys|Node|Domain)name|Release|Version|Machine)(\s+)\[(\d+)\]u?int8`)
	b = convertUtsnameRegex.ReplaceAll(b, []byte("$1$3[$4]byte"))

	// Convert [1024]int8 to [1024]byte in Ptmget members
	convertPtmget := regexp.MustCompile(`([SC]n)(\s+)\[(\d+)\]u?int8`)
	b = convertPtmget.ReplaceAll(b, []byte("$1[$3]byte"))

	// Remove spare fields (e.g. in Statx_t)
	spareFieldsRegex := regexp.MustCompile(`X__spare\S*`)
	b = spareFieldsRegex.ReplaceAll(b, []byte("_"))

	// Remove cgo padding fields
	removePaddingFieldsRegex := regexp.MustCompile(`Pad_cgo_\d+`)
	b = removePaddingFieldsRegex.ReplaceAll(b, []byte("_"))

	// Remove padding, hidden, or unused fields
	removeFieldsRegex = regexp.MustCompile(`\b(X_\S+|Padding)`)
	b = removeFieldsRegex.ReplaceAll(b, []byte("_"))

	// Remove the first line of warning from cgo
	b = b[bytes.IndexByte(b, '\n')+1:]
	// Modify the command in the header to include:
	//  mkpost, our own warning, and a build tag.
	replacement := fmt.Sprintf(`$1 | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build %s,%s`, goarch, goos)
	cgoCommandRegex := regexp.MustCompile(`(cgo -godefs .*)`)
	b = cgoCommandRegex.ReplaceAll(b, []byte(replacement))

	// gofmt
	b, err = format.Source(b)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(b)
}
