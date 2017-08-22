// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen.go gen_trieval.go

// Package cases provides general and language-specific case mappers.
package cases

import (
	"golang.org/x/text/language"
	"golang.org/x/text/transform"
)

// References:
// - Unicode Reference Manual Chapter 3.13, 4.2, and 5.18.
// - http://www.unicode.org/reports/tr29/
// - http://www.unicode.org/Public/6.3.0/ucd/CaseFolding.txt
// - http://www.unicode.org/Public/6.3.0/ucd/SpecialCasing.txt
// - http://www.unicode.org/Public/6.3.0/ucd/DerivedCoreProperties.txt
// - http://www.unicode.org/Public/6.3.0/ucd/auxiliary/WordBreakProperty.txt
// - http://www.unicode.org/Public/6.3.0/ucd/auxiliary/WordBreakTest.txt
// - http://userguide.icu-project.org/transforms/casemappings

// TODO:
// - Case folding
// - Wide and Narrow?
// - Segmenter option for title casing.
// - ASCII fast paths
// - Encode Soft-Dotted property within trie somehow.

// A Caser transforms given input to a certain case. It implements
// transform.Transformer.
//
// A Caser may be stateful and should therefore not be shared between
// goroutines.
type Caser struct {
	t transform.Transformer
}

// Bytes returns a new byte slice with the result of converting b to the case
// form implemented by c.
func (c Caser) Bytes(b []byte) []byte {
	b, _, _ = transform.Bytes(c.t, b)
	return b
}

// String returns a string with the result of transforming s to the case form
// implemented by c.
func (c Caser) String(s string) string {
	s, _, _ = transform.String(c.t, s)
	return s
}

// Reset resets the Caser to be reused for new input after a previous call to
// Transform.
func (c Caser) Reset() { c.t.Reset() }

// Transform implements the Transformer interface and transforms the given input
// to the case form implemented by c.
func (c Caser) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	return c.t.Transform(dst, src, atEOF)
}

// Upper returns a Caser for language-specific uppercasing.
func Upper(t language.Tag, opts ...Option) Caser {
	return Caser{makeUpper(t, getOpts(opts...))}
}

// Lower returns a Caser for language-specific lowercasing.
func Lower(t language.Tag, opts ...Option) Caser {
	return Caser{makeLower(t, getOpts(opts...))}
}

// Title returns a Caser for language-specific title casing. It uses an
// approximation of the default Unicode Word Break algorithm.
func Title(t language.Tag, opts ...Option) Caser {
	return Caser{makeTitle(t, getOpts(opts...))}
}

// Fold returns a Caser that implements Unicode case folding. The returned Caser
// is stateless and safe to use concurrently by multiple goroutines.
//
// Case folding does not normalize the input and may not preserve a normal form.
// Use the collate or search package for more convenient and linguistically
// sound comparisons.  Use unicode/precis for string comparisons where security
// aspects are a concern.
func Fold(opts ...Option) Caser {
	return Caser{makeFold(getOpts(opts...))}
}

// An Option is used to modify the behavior of a Caser.
type Option func(o *options)

var (
	// NoLower disables the lowercasing of non-leading letters for a title
	// caser.
	NoLower Option = noLower

	// Compact omits mappings in case folding for characters that would grow the
	// input. (Unimplemented.)
	Compact Option = compact
)

// TODO: option to preserve a normal form, if applicable?

type options struct {
	noLower bool
	simple  bool

	// TODO: segmenter, max ignorable, alternative versions, etc.

	noFinalSigma bool // Only used for testing.
}

func getOpts(o ...Option) (res options) {
	for _, f := range o {
		f(&res)
	}
	return
}

func noLower(o *options) {
	o.noLower = true
}

func compact(o *options) {
	o.simple = true
}
