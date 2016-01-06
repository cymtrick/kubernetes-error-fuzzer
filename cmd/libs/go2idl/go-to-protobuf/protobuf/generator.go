/*
Copyright 2015 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package protobuf

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"k8s.io/kubernetes/cmd/libs/go2idl/generator"
	"k8s.io/kubernetes/cmd/libs/go2idl/namer"
	"k8s.io/kubernetes/cmd/libs/go2idl/types"
)

// genProtoIDL produces a .proto IDL.
type genProtoIDL struct {
	generator.DefaultGen
	localPackage   types.Name
	localGoPackage types.Name
	imports        *ImportTracker

	generateAll    bool
	omitGogo       bool
	omitFieldTypes map[types.Name]struct{}
}

func (g *genProtoIDL) PackageVars(c *generator.Context) []string {
	if g.omitGogo {
		return []string{
			fmt.Sprintf("option go_package = %q;", g.localGoPackage.Name),
		}
	}
	return []string{
		"option (gogoproto.marshaler_all) = true;",
		"option (gogoproto.sizer_all) = true;",
		"option (gogoproto.unmarshaler_all) = true;",
		"option (gogoproto.goproto_unrecognized_all) = false;",
		"option (gogoproto.goproto_enum_prefix_all) = false;",
		"option (gogoproto.goproto_getters_all) = false;",
		fmt.Sprintf("option go_package = %q;", g.localGoPackage.Name),
	}
}
func (g *genProtoIDL) Filename() string { return g.OptionalName + ".proto" }
func (g *genProtoIDL) FileType() string { return "protoidl" }
func (g *genProtoIDL) Namers(c *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		// The local namer returns the correct protobuf name for a proto type
		// in the context of a package
		"local": localNamer{g.localPackage},
	}
}

// Filter ignores types that are identified as not exportable.
func (g *genProtoIDL) Filter(c *generator.Context, t *types.Type) bool {
	flags := types.ExtractCommentTags("+", t.CommentLines)
	switch {
	case flags["protobuf"] == "false":
		return false
	case flags["protobuf"] == "true":
		return true
	case !g.generateAll:
		return false
	}
	seen := map[*types.Type]bool{}
	ok := isProtoable(seen, t)
	return ok
}

func isProtoable(seen map[*types.Type]bool, t *types.Type) bool {
	if seen[t] {
		// be optimistic in the case of type cycles.
		return true
	}
	seen[t] = true
	switch t.Kind {
	case types.Builtin:
		return true
	case types.Alias:
		return isProtoable(seen, t.Underlying)
	case types.Slice, types.Pointer:
		return isProtoable(seen, t.Elem)
	case types.Map:
		return isProtoable(seen, t.Key) && isProtoable(seen, t.Elem)
	case types.Struct:
		for _, m := range t.Members {
			if isProtoable(seen, m.Type) {
				return true
			}
		}
		return false
	case types.Func, types.Chan:
		return false
	case types.DeclarationOf, types.Unknown, types.Unsupported:
		return false
	case types.Interface:
		return false
	default:
		log.Printf("WARNING: type %q is not protable: %s", t.Kind, t.Name)
		return false
	}
}

func (g *genProtoIDL) Imports(c *generator.Context) (imports []string) {
	lines := []string{}
	// TODO: this could be expressed more cleanly
	for _, line := range g.imports.ImportLines() {
		if g.omitGogo && line == "github.com/gogo/protobuf/gogoproto/gogo.proto" {
			continue
		}
		lines = append(lines, line)
	}
	return lines
}

// GenerateType makes the body of a file implementing a set for type t.
func (g *genProtoIDL) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "$", "$")
	b := bodyGen{
		locator: &protobufLocator{
			namer:   c.Namers["proto"].(ProtobufFromGoNamer),
			tracker: g.imports,

			localGoPackage: g.localGoPackage.Package,
		},
		localPackage: g.localPackage,

		omitGogo:       g.omitGogo,
		omitFieldTypes: g.omitFieldTypes,

		t: t,
	}
	switch t.Kind {
	case types.Struct:
		return b.doStruct(sw)
	default:
		return b.unknown(sw)
	}
	return sw.Error()
}

// ProtobufFromGoNamer finds the protobuf name of a type (and its package, and
// the package path) from its Go name.
type ProtobufFromGoNamer interface {
	GoNameToProtoName(name types.Name) types.Name
}

type ProtobufLocator interface {
	ProtoTypeFor(t *types.Type) (*types.Type, error)
	CastTypeName(name types.Name) string
}

type protobufLocator struct {
	namer   ProtobufFromGoNamer
	tracker namer.ImportTracker

	localGoPackage string
}

// CastTypeName returns the cast type name of a Go type
// TODO: delegate to a new localgo namer?
func (p protobufLocator) CastTypeName(name types.Name) string {
	if name.Package == p.localGoPackage {
		return name.Name
	}
	return name.String()
}

// ProtoTypeFor locates a Protobuf type for the provided Go type (if possible).
func (p protobufLocator) ProtoTypeFor(t *types.Type) (*types.Type, error) {
	switch {
	// we've already converted the type, or it's a map
	case t.Kind == typesKindProtobuf || t.Kind == types.Map:
		p.tracker.AddType(t)
		return t, nil
	}
	// it's a fundamental type
	if t, ok := isFundamentalProtoType(t); ok {
		p.tracker.AddType(t)
		return t, nil
	}
	// it's a message
	if t.Kind == types.Struct {
		t := &types.Type{
			Name: p.namer.GoNameToProtoName(t.Name),
			Kind: typesKindProtobuf,

			CommentLines: t.CommentLines,
		}
		p.tracker.AddType(t)
		return t, nil
	}
	return nil, errUnrecognizedType
}

type bodyGen struct {
	locator        ProtobufLocator
	localPackage   types.Name
	omitGogo       bool
	omitFieldTypes map[types.Name]struct{}

	t *types.Type
}

func (b bodyGen) unknown(sw *generator.SnippetWriter) error {
	return fmt.Errorf("not sure how to generate: %#v", b.t)
}

func (b bodyGen) doStruct(sw *generator.SnippetWriter) error {
	if len(b.t.Name.Name) == 0 {
		return nil
	}
	if isPrivateGoName(b.t.Name.Name) {
		return nil
	}

	var fields []protoField
	options := []string{}
	allOptions := types.ExtractCommentTags("+", b.t.CommentLines)
	for k, v := range allOptions {
		switch {
		case strings.HasPrefix(k, "protobuf.options."):
			key := strings.TrimPrefix(k, "protobuf.options.")
			switch key {
			case "marshal":
				if v == "false" {
					if !b.omitGogo {
						options = append(options,
							"(gogoproto.marshaler) = false",
							"(gogoproto.unmarshaler) = false",
							"(gogoproto.sizer) = false",
						)
					}
				}
			default:
				if !b.omitGogo || !strings.HasPrefix(key, "(gogoproto.") {
					options = append(options, fmt.Sprintf("%s = %s", key, v))
				}
			}
		case k == "protobuf.embed":
			fields = []protoField{
				{
					Tag:  1,
					Name: v,
					Type: &types.Type{
						Name: types.Name{
							Name:    v,
							Package: b.localPackage.Package,
							Path:    b.localPackage.Path,
						},
					},
				},
			}
		}
	}

	if fields == nil {
		memberFields, err := membersToFields(b.locator, b.t, b.localPackage, b.omitFieldTypes)
		if err != nil {
			return fmt.Errorf("type %v cannot be converted to protobuf: %v", b.t, err)
		}
		fields = memberFields
	}

	out := sw.Out()
	genComment(out, b.t.CommentLines, "")
	sw.Do(`message $.Name.Name$ {
`, b.t)

	if len(options) > 0 {
		sort.Sort(sort.StringSlice(options))
		for _, s := range options {
			fmt.Fprintf(out, "  option %s;\n", s)
		}
		fmt.Fprintln(out)
	}

	for i, field := range fields {
		genComment(out, field.CommentLines, "  ")
		fmt.Fprintf(out, "  ")
		switch {
		case field.Map:
		case field.Repeated:
			fmt.Fprintf(out, "repeated ")
		case field.Required:
			fmt.Fprintf(out, "required ")
		default:
			fmt.Fprintf(out, "optional ")
		}
		sw.Do(`$.Type|local$ $.Name$ = $.Tag$`, field)
		if len(field.Extras) > 0 {
			extras := []string{}
			for k, v := range field.Extras {
				if b.omitGogo && strings.HasPrefix(k, "(gogoproto.") {
					continue
				}
				extras = append(extras, fmt.Sprintf("%s = %s", k, v))
			}
			sort.Sort(sort.StringSlice(extras))
			if len(extras) > 0 {
				fmt.Fprintf(out, " [")
				fmt.Fprint(out, strings.Join(extras, ", "))
				fmt.Fprintf(out, "]")
			}
		}
		fmt.Fprintf(out, ";\n")
		if i != len(fields)-1 {
			fmt.Fprintf(out, "\n")
		}
	}
	fmt.Fprintf(out, "}\n\n")
	return nil
}

type protoField struct {
	LocalPackage types.Name

	Tag      int
	Name     string
	Type     *types.Type
	Map      bool
	Repeated bool
	Optional bool
	Required bool
	Nullable bool
	Extras   map[string]string

	CommentLines string

	OptionalSet bool
}

var (
	errUnrecognizedType = fmt.Errorf("did not recognize the provided type")
)

func isFundamentalProtoType(t *types.Type) (*types.Type, bool) {
	// TODO: when we enable proto3, also include other fundamental types in the google.protobuf package
	// switch {
	// case t.Kind == types.Struct && t.Name == types.Name{Package: "time", Name: "Time"}:
	// 	return &types.Type{
	// 		Kind: typesKindProtobuf,
	// 		Name: types.Name{Path: "google/protobuf/timestamp.proto", Package: "google.protobuf", Name: "Timestamp"},
	// 	}, true
	// }
	switch t.Kind {
	case types.Slice:
		if t.Elem.Name.Name == "byte" && len(t.Elem.Name.Package) == 0 {
			return &types.Type{Name: types.Name{Name: "bytes"}, Kind: typesKindProtobuf}, true
		}
	case types.Builtin:
		switch t.Name.Name {
		case "string", "uint32", "int32", "uint64", "int64", "bool":
			return &types.Type{Name: types.Name{Name: t.Name.Name}, Kind: typesKindProtobuf}, true
		case "int":
			return &types.Type{Name: types.Name{Name: "int64"}, Kind: typesKindProtobuf}, true
		case "uint":
			return &types.Type{Name: types.Name{Name: "uint64"}, Kind: typesKindProtobuf}, true
		case "float64", "float":
			return &types.Type{Name: types.Name{Name: "double"}, Kind: typesKindProtobuf}, true
		case "float32":
			return &types.Type{Name: types.Name{Name: "float"}, Kind: typesKindProtobuf}, true
		case "uintptr":
			return &types.Type{Name: types.Name{Name: "uint64"}, Kind: typesKindProtobuf}, true
		}
		// TODO: complex?
	}
	return t, false
}

func memberTypeToProtobufField(locator ProtobufLocator, field *protoField, t *types.Type) error {
	var err error
	switch t.Kind {
	case typesKindProtobuf:
		field.Type, err = locator.ProtoTypeFor(t)
	case types.Builtin:
		field.Type, err = locator.ProtoTypeFor(t)
	case types.Map:
		valueField := &protoField{}
		if err := memberTypeToProtobufField(locator, valueField, t.Elem); err != nil {
			return err
		}
		keyField := &protoField{}
		if err := memberTypeToProtobufField(locator, keyField, t.Key); err != nil {
			return err
		}
		field.Type = &types.Type{
			Kind: types.Map,
			Key:  keyField.Type,
			Elem: valueField.Type,
		}
		if !strings.HasPrefix(t.Name.Name, "map[") {
			field.Extras["(gogoproto.casttype)"] = strconv.Quote(locator.CastTypeName(t.Name))
		}
		if k, ok := keyField.Extras["(gogoproto.casttype)"]; ok {
			field.Extras["(gogoproto.castkey)"] = k
		}
		if v, ok := valueField.Extras["(gogoproto.casttype)"]; ok {
			field.Extras["(gogoproto.castvalue)"] = v
		}
		field.Map = true
	case types.Pointer:
		if err := memberTypeToProtobufField(locator, field, t.Elem); err != nil {
			return err
		}
		field.Nullable = true
	case types.Alias:
		if err := memberTypeToProtobufField(locator, field, t.Underlying); err != nil {
			log.Printf("failed to alias: %s %s: err", t.Name, t.Underlying.Name, err)
			return err
		}
		if field.Extras == nil {
			field.Extras = make(map[string]string)
		}
		field.Extras["(gogoproto.casttype)"] = strconv.Quote(locator.CastTypeName(t.Name))
	case types.Slice:
		if t.Elem.Name.Name == "byte" && len(t.Elem.Name.Package) == 0 {
			field.Type = &types.Type{Name: types.Name{Name: "bytes"}, Kind: typesKindProtobuf}
			return nil
		}
		if err := memberTypeToProtobufField(locator, field, t.Elem); err != nil {
			return err
		}
		field.Repeated = true
	case types.Struct:
		if len(t.Name.Name) == 0 {
			return errUnrecognizedType
		}
		field.Type, err = locator.ProtoTypeFor(t)
		field.Nullable = false
	default:
		return errUnrecognizedType
	}
	return err
}

// protobufTagToField extracts information from an existing protobuf tag
// TODO: take a current package
func protobufTagToField(tag string, field *protoField, m types.Member, t *types.Type, localPackage types.Name) error {
	if len(tag) == 0 {
		return nil
	}

	// protobuf:"bytes,3,opt,name=Id,customtype=github.com/gogo/protobuf/test.Uuid"
	parts := strings.Split(tag, ",")
	if len(parts) < 3 {
		return fmt.Errorf("member %q of %q malformed 'protobuf' tag, not enough segments\n", m.Name, t.Name)
	}
	protoTag, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("member %q of %q malformed 'protobuf' tag, field ID is %q which is not an integer: %v\n", m.Name, t.Name, parts[1], err)
	}
	field.Tag = protoTag
	// TODO: we are converting a Protobuf type back into an internal type, which is questionable
	if last := strings.LastIndex(parts[0], "."); last != -1 {
		prefix := parts[0][:last]
		field.Type = &types.Type{
			Name: types.Name{
				Name:    parts[0][last+1:],
				Package: prefix,
				// TODO: this probably needs to be a lookup into a namer
				Path: strings.Replace(prefix, ".", "/", -1),
			},
			Kind: typesKindProtobuf,
		}
	} else {
		switch parts[0] {
		case "varint", "bytes", "fixed64":
		default:
			field.Type = &types.Type{
				Name: types.Name{
					Name:    parts[0],
					Package: localPackage.Package,
					Path:    localPackage.Path,
				},
				Kind: typesKindProtobuf,
			}
		}
	}
	field.OptionalSet = true

	protoExtra := make(map[string]string)
	for i, extra := range parts[3:] {
		parts := strings.SplitN(extra, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("member %q of %q malformed 'protobuf' tag, tag %d should be key=value, got %q\n", m.Name, t.Name, i+4, extra)
		}
		switch parts[0] {
		case "casttype":
			parts[0] = fmt.Sprintf("(gogoproto.%s)", parts[0])
			protoExtra[parts[0]] = parts[1]
		}
	}

	field.Extras = protoExtra
	if name, ok := protoExtra["name"]; ok {
		field.Name = name
		delete(protoExtra, "name")
	}

	return nil
}

func membersToFields(locator ProtobufLocator, t *types.Type, localPackage types.Name, omitFieldTypes map[types.Name]struct{}) ([]protoField, error) {
	fields := []protoField{}

	for _, m := range t.Members {
		if isPrivateGoName(m.Name) {
			// skip private fields
			continue
		}
		if _, ok := omitFieldTypes[types.Name{Name: m.Type.Name.Name, Package: m.Type.Name.Package}]; ok {
			continue
		}
		tags := reflect.StructTag(m.Tags)
		field := protoField{
			LocalPackage: localPackage,

			Tag:    -1,
			Extras: make(map[string]string),
		}

		if err := protobufTagToField(tags.Get("protobuf"), &field, m, t, localPackage); err != nil {
			return nil, err
		}

		// extract information from JSON field tag
		if tag := tags.Get("json"); len(tag) > 0 {
			parts := strings.Split(tag, ",")
			if len(field.Name) == 0 && len(parts[0]) != 0 {
				field.Name = parts[0]
			}
			if field.Tag == -1 && field.Name == "-" {
				continue
			}
		}

		if field.Type == nil {
			if err := memberTypeToProtobufField(locator, &field, m.Type); err != nil {
				return nil, fmt.Errorf("unable to embed type %q as field %q in %q: %v", m.Type, field.Name, t.Name, err)
			}
		}
		if len(field.Name) == 0 {
			field.Name = strings.ToLower(m.Name[:1]) + m.Name[1:]
		}

		if field.Map && field.Repeated {
			// maps cannot be repeated
			field.Repeated = false
			field.Nullable = true
		}

		if !field.Nullable {
			field.Extras["(gogoproto.nullable)"] = "false"
		}
		if (field.Type.Name.Name == "bytes" && field.Type.Name.Package == "") || (field.Repeated && field.Type.Name.Package == "" && isPrivateGoName(field.Type.Name.Name)) {
			delete(field.Extras, "(gogoproto.nullable)")
		}
		if field.Name != m.Name {
			field.Extras["(gogoproto.customname)"] = strconv.Quote(m.Name)
		}
		field.CommentLines = m.CommentLines
		fields = append(fields, field)
	}

	// assign tags
	highest := 0
	byTag := make(map[int]*protoField)
	// fields are in Go struct order, which we preserve
	for i := range fields {
		field := &fields[i]
		tag := field.Tag
		if tag != -1 {
			if existing, ok := byTag[tag]; ok {
				return nil, fmt.Errorf("field %q and %q in %q both have tag %d", field.Name, existing.Name, tag)
			}
			byTag[tag] = field
		}
		if tag > highest {
			highest = tag
		}
	}
	// starting from the highest observed tag, assign new field tags
	for i := range fields {
		field := &fields[i]
		if field.Tag != -1 {
			continue
		}
		highest++
		field.Tag = highest
		byTag[field.Tag] = field
	}
	return fields, nil
}

func genComment(out io.Writer, comment, indent string) {
	lines := strings.Split(comment, "\n")
	for {
		l := len(lines)
		if l == 0 || len(lines[l-1]) != 0 {
			break
		}
		lines = lines[:l-1]
	}
	for _, c := range lines {
		fmt.Fprintf(out, "%s// %s\n", indent, c)
	}
}

type protoIDLFileType struct{}

func (ft protoIDLFileType) AssembleFile(f *generator.File, pathname string) error {
	log.Printf("Assembling IDL file %q", pathname)
	destFile, err := os.Create(pathname)
	if err != nil {
		return err
	}
	defer destFile.Close()

	b := &bytes.Buffer{}
	et := generator.NewErrorTracker(b)
	ft.assemble(et, f)
	if et.Error() != nil {
		return et.Error()
	}

	// TODO: is there an IDL formatter?
	_, err = destFile.Write(b.Bytes())
	return err
}

func (ft protoIDLFileType) VerifyFile(f *generator.File, pathname string) error {
	log.Printf("Verifying IDL file %q", pathname)
	friendlyName := filepath.Join(f.PackageName, f.Name)
	b := &bytes.Buffer{}
	et := generator.NewErrorTracker(b)
	ft.assemble(et, f)
	if et.Error() != nil {
		return et.Error()
	}
	formatted := b.Bytes()
	existing, err := ioutil.ReadFile(pathname)
	if err != nil {
		return fmt.Errorf("unable to read file %q for comparison: %v", friendlyName, err)
	}
	if bytes.Compare(formatted, existing) == 0 {
		return nil
	}
	// Be nice and find the first place where they differ
	i := 0
	for i < len(formatted) && i < len(existing) && formatted[i] == existing[i] {
		i++
	}
	eDiff, fDiff := existing[i:], formatted[i:]
	if len(eDiff) > 100 {
		eDiff = eDiff[:100]
	}
	if len(fDiff) > 100 {
		fDiff = fDiff[:100]
	}
	return fmt.Errorf("output for %q differs; first existing/expected diff: \n  %q\n  %q", friendlyName, string(eDiff), string(fDiff))
}

func (ft protoIDLFileType) assemble(w io.Writer, f *generator.File) {
	w.Write(f.Header)

	fmt.Fprint(w, "syntax = 'proto2';\n\n")

	if len(f.PackageName) > 0 {
		fmt.Fprintf(w, "package %v;\n\n", f.PackageName)
	}

	if len(f.Imports) > 0 {
		imports := []string{}
		for i := range f.Imports {
			imports = append(imports, i)
		}
		sort.Strings(imports)
		for _, s := range imports {
			fmt.Fprintf(w, "import %q;\n", s)
		}
		fmt.Fprint(w, "\n")
	}

	if f.Vars.Len() > 0 {
		fmt.Fprintf(w, "%s\n", f.Vars.String())
	}

	w.Write(f.Body.Bytes())
}

func isPackable(t *types.Type) bool {
	if t.Kind != typesKindProtobuf {
		return false
	}
	switch t.Name.Name {
	case "int32", "int64", "varint":
		return true
	default:
		return false
	}
}

func isPrivateGoName(name string) bool {
	if len(name) == 0 {
		return true
	}
	return strings.ToLower(name[:1]) == name[:1]
}
