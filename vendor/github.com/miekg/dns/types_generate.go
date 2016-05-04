//+build ignore

// types_generate.go is meant to run with go generate. It will use
// go/{importer,types} to track down all the RR struct types. Then for each type
// it will generate conversion tables (TypeToRR and TypeToString) and banal
// methods (len, Header, copy) based on the struct tags. The generated source is
// written to ztypes.go, and is meant to be checked into git.
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/importer"
	"go/types"
	"log"
	"os"
	"strings"
	"text/template"
)

var skipLen = map[string]struct{}{
	"NSEC":     {},
	"NSEC3":    {},
	"OPT":      {},
	"WKS":      {},
	"IPSECKEY": {},
}

var packageHdr = `
// *** DO NOT MODIFY ***
// AUTOGENERATED BY go generate

package dns

import (
	"encoding/base64"
	"net"
)

`

var TypeToRR = template.Must(template.New("TypeToRR").Parse(`
// TypeToRR is a map of constructors for each RR type.
var TypeToRR = map[uint16]func() RR{
{{range .}}{{if ne . "RFC3597"}}  Type{{.}}:  func() RR { return new({{.}}) },
{{end}}{{end}}                    }

`))

var typeToString = template.Must(template.New("typeToString").Parse(`
// TypeToString is a map of strings for each RR type.
var TypeToString = map[uint16]string{
{{range .}}{{if ne . "NSAPPTR"}}  Type{{.}}: "{{.}}",
{{end}}{{end}}                    TypeNSAPPTR:    "NSAP-PTR",
}

`))

var headerFunc = template.Must(template.New("headerFunc").Parse(`
// Header() functions
{{range .}}  func (rr *{{.}}) Header() *RR_Header { return &rr.Hdr }
{{end}}

`))

// getTypeStruct will take a type and the package scope, and return the
// (innermost) struct if the type is considered a RR type (currently defined as
// those structs beginning with a RR_Header, could be redefined as implementing
// the RR interface). The bool return value indicates if embedded structs were
// resolved.
func getTypeStruct(t types.Type, scope *types.Scope) (*types.Struct, bool) {
	st, ok := t.Underlying().(*types.Struct)
	if !ok {
		return nil, false
	}
	if st.Field(0).Type() == scope.Lookup("RR_Header").Type() {
		return st, false
	}
	if st.Field(0).Anonymous() {
		st, _ := getTypeStruct(st.Field(0).Type(), scope)
		return st, true
	}
	return nil, false
}

func main() {
	// Import and type-check the package
	pkg, err := importer.Default().Import("github.com/miekg/dns")
	fatalIfErr(err)
	scope := pkg.Scope()

	// Collect constants like TypeX
	var numberedTypes []string
	for _, name := range scope.Names() {
		o := scope.Lookup(name)
		if o == nil || !o.Exported() {
			continue
		}
		b, ok := o.Type().(*types.Basic)
		if !ok || b.Kind() != types.Uint16 {
			continue
		}
		if !strings.HasPrefix(o.Name(), "Type") {
			continue
		}
		name := strings.TrimPrefix(o.Name(), "Type")
		if name == "PrivateRR" {
			continue
		}
		numberedTypes = append(numberedTypes, name)
	}

	// Collect actual types (*X)
	var namedTypes []string
	for _, name := range scope.Names() {
		o := scope.Lookup(name)
		if o == nil || !o.Exported() {
			continue
		}
		if st, _ := getTypeStruct(o.Type(), scope); st == nil {
			continue
		}
		if name == "PrivateRR" {
			continue
		}

		// Check if corresponding TypeX exists
		if scope.Lookup("Type"+o.Name()) == nil && o.Name() != "RFC3597" {
			log.Fatalf("Constant Type%s does not exist.", o.Name())
		}

		namedTypes = append(namedTypes, o.Name())
	}

	b := &bytes.Buffer{}
	b.WriteString(packageHdr)

	// Generate TypeToRR
	fatalIfErr(TypeToRR.Execute(b, namedTypes))

	// Generate typeToString
	fatalIfErr(typeToString.Execute(b, numberedTypes))

	// Generate headerFunc
	fatalIfErr(headerFunc.Execute(b, namedTypes))

	// Generate len()
	fmt.Fprint(b, "// len() functions\n")
	for _, name := range namedTypes {
		if _, ok := skipLen[name]; ok {
			continue
		}
		o := scope.Lookup(name)
		st, isEmbedded := getTypeStruct(o.Type(), scope)
		if isEmbedded {
			continue
		}
		fmt.Fprintf(b, "func (rr *%s) len() int {\n", name)
		fmt.Fprintf(b, "l := rr.Hdr.len()\n")
		for i := 1; i < st.NumFields(); i++ {
			o := func(s string) { fmt.Fprintf(b, s, st.Field(i).Name()) }

			if _, ok := st.Field(i).Type().(*types.Slice); ok {
				switch st.Tag(i) {
				case `dns:"-"`:
					// ignored
				case `dns:"cdomain-name"`, `dns:"domain-name"`, `dns:"txt"`:
					o("for _, x := range rr.%s { l += len(x) + 1 }\n")
				default:
					log.Fatalln(name, st.Field(i).Name(), st.Tag(i))
				}
				continue
			}

			switch st.Tag(i) {
			case `dns:"-"`:
				// ignored
			case `dns:"cdomain-name"`, `dns:"domain-name"`:
				o("l += len(rr.%s) + 1\n")
			case `dns:"octet"`:
				o("l += len(rr.%s)\n")
			case `dns:"base64"`:
				o("l += base64.StdEncoding.DecodedLen(len(rr.%s))\n")
			case `dns:"size-hex"`, `dns:"hex"`:
				o("l += len(rr.%s)/2 + 1\n")
			case `dns:"a"`:
				o("l += net.IPv4len // %s\n")
			case `dns:"aaaa"`:
				o("l += net.IPv6len // %s\n")
			case `dns:"txt"`:
				o("for _, t := range rr.%s { l += len(t) + 1 }\n")
			case `dns:"uint48"`:
				o("l += 6 // %s\n")
			case "":
				switch st.Field(i).Type().(*types.Basic).Kind() {
				case types.Uint8:
					o("l += 1 // %s\n")
				case types.Uint16:
					o("l += 2 // %s\n")
				case types.Uint32:
					o("l += 4 // %s\n")
				case types.Uint64:
					o("l += 8 // %s\n")
				case types.String:
					o("l += len(rr.%s) + 1\n")
				default:
					log.Fatalln(name, st.Field(i).Name())
				}
			default:
				log.Fatalln(name, st.Field(i).Name(), st.Tag(i))
			}
		}
		fmt.Fprintf(b, "return l }\n")
	}

	// Generate copy()
	fmt.Fprint(b, "// copy() functions\n")
	for _, name := range namedTypes {
		o := scope.Lookup(name)
		st, isEmbedded := getTypeStruct(o.Type(), scope)
		if isEmbedded {
			continue
		}
		fmt.Fprintf(b, "func (rr *%s) copy() RR {\n", name)
		fields := []string{"*rr.Hdr.copyHeader()"}
		for i := 1; i < st.NumFields(); i++ {
			f := st.Field(i).Name()
			if sl, ok := st.Field(i).Type().(*types.Slice); ok {
				t := sl.Underlying().String()
				t = strings.TrimPrefix(t, "[]")
				t = strings.TrimPrefix(t, "github.com/miekg/dns.")
				fmt.Fprintf(b, "%s := make([]%s, len(rr.%s)); copy(%s, rr.%s)\n",
					f, t, f, f, f)
				fields = append(fields, f)
				continue
			}
			if st.Field(i).Type().String() == "net.IP" {
				fields = append(fields, "copyIP(rr."+f+")")
				continue
			}
			fields = append(fields, "rr."+f)
		}
		fmt.Fprintf(b, "return &%s{%s}\n", name, strings.Join(fields, ","))
		fmt.Fprintf(b, "}\n")
	}

	// gofmt
	res, err := format.Source(b.Bytes())
	if err != nil {
		b.WriteTo(os.Stderr)
		log.Fatal(err)
	}

	// write result
	f, err := os.Create("ztypes.go")
	fatalIfErr(err)
	defer f.Close()
	f.Write(res)
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
