// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

// Download https://encoding.spec.whatwg.org/encodings.json and use it to
// generate table.go.

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type enc struct {
	Name   string
	Labels []string
}

type group struct {
	Encodings []enc
	Heading   string
}

const specURL = "https://encoding.spec.whatwg.org/encodings.json"

func main() {
	resp, err := http.Get(specURL)
	if err != nil {
		log.Fatalf("error fetching %s: %s", specURL, err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("error fetching %s: HTTP status %s", specURL, resp.Status)
	}
	defer resp.Body.Close()

	var groups []group
	d := json.NewDecoder(resp.Body)
	err = d.Decode(&groups)
	if err != nil {
		log.Fatalf("error reading encodings.json: %s", err)
	}

	fmt.Println("// generated by go run gen.go; DO NOT EDIT")
	fmt.Println()
	fmt.Println("package charset")
	fmt.Println()

	fmt.Println("import (")
	fmt.Println(`"golang.org/x/text/encoding"`)
	for _, pkg := range []string{"charmap", "japanese", "korean", "simplifiedchinese", "traditionalchinese", "unicode"} {
		fmt.Printf("\"golang.org/x/text/encoding/%s\"\n", pkg)
	}
	fmt.Println(")")
	fmt.Println()

	fmt.Println("var encodings = map[string]struct{e encoding.Encoding; name string} {")
	for _, g := range groups {
		for _, e := range g.Encodings {
			goName, ok := miscNames[e.Name]
			if !ok {
				for k, v := range prefixes {
					if strings.HasPrefix(e.Name, k) {
						goName = v + e.Name[len(k):]
						break
					}
				}
				if goName == "" {
					log.Fatalf("unrecognized encoding name: %s", e.Name)
				}
			}

			for _, label := range e.Labels {
				fmt.Printf("%q: {%s, %q},\n", label, goName, e.Name)
			}
		}
	}
	fmt.Println("}")
}

var prefixes = map[string]string{
	"iso-8859-": "charmap.ISO8859_",
	"windows-":  "charmap.Windows",
}

var miscNames = map[string]string{
	"utf-8":          "encoding.Nop",
	"ibm866":         "charmap.CodePage866",
	"iso-8859-8-i":   "charmap.ISO8859_8",
	"koi8-r":         "charmap.KOI8R",
	"koi8-u":         "charmap.KOI8U",
	"macintosh":      "charmap.Macintosh",
	"x-mac-cyrillic": "charmap.MacintoshCyrillic",
	"gbk":            "simplifiedchinese.GBK",
	"gb18030":        "simplifiedchinese.GB18030",
	"hz-gb-2312":     "simplifiedchinese.HZGB2312",
	"big5":           "traditionalchinese.Big5",
	"euc-jp":         "japanese.EUCJP",
	"iso-2022-jp":    "japanese.ISO2022JP",
	"shift_jis":      "japanese.ShiftJIS",
	"euc-kr":         "korean.EUCKR",
	"replacement":    "encoding.Replacement",
	"utf-16be":       "unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)",
	"utf-16le":       "unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)",
	"x-user-defined": "charmap.XUserDefined",
}
