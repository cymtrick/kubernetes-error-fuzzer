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

package generators

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"k8s.io/kubernetes/cmd/libs/go2idl/args"
	"k8s.io/kubernetes/cmd/libs/go2idl/generator"
	"k8s.io/kubernetes/cmd/libs/go2idl/namer"
	"k8s.io/kubernetes/cmd/libs/go2idl/types"
	"k8s.io/kubernetes/pkg/util/sets"

	"github.com/golang/glog"
)

// TODO: This is created only to reduce number of changes in a single PR.
// Remove it and use PublicNamer instead.
func deepCopyNamer() *namer.NameStrategy {
	return &namer.NameStrategy{
		Join: func(pre string, in []string, post string) string {
			return strings.Join(in, "_")
		},
		PrependPackageNames: 1,
	}
}

// NameSystems returns the name system used by the generators in this package.
func NameSystems() namer.NameSystems {
	return namer.NameSystems{
		"public": deepCopyNamer(),
		"raw":    namer.NewRawNamer("", nil),
	}
}

// DefaultNameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func DefaultNameSystem() string {
	return "public"
}

func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	boilerplate, err := arguments.LoadGoBoilerplate()
	if err != nil {
		glog.Fatalf("Failed loading boilerplate: %v", err)
	}

	inputs := sets.NewString(arguments.InputDirs...)
	packages := generator.Packages{}
	header := append([]byte(
		`
// +build !ignore_autogenerated

`), boilerplate...)
	header = append(header, []byte(
		`
// This file was autogenerated by deepcopy-gen. Do not edit it manually!

`)...)
	for _, p := range context.Universe {
		copyableType := false
		for _, t := range p.Types {
			if copyableWithinPackage(t) {
				copyableType = true
			}
		}
		if copyableType {
			path := p.Path
			packages = append(packages,
				&generator.DefaultPackage{
					PackageName: filepath.Base(path),
					PackagePath: path,
					HeaderText:  header,
					GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
						generators = []generator.Generator{}
						generators = append(
							generators, NewGenDeepCopy("deep_copy_generated", path, inputs.Has(path)))
						return generators
					},
					FilterFunc: func(c *generator.Context, t *types.Type) bool {
						if t.Name.Package != path {
							return false
						}
						return copyableWithinPackage(t)
					},
				})
		}
	}
	return packages
}

const (
	apiPackagePath        = "k8s.io/kubernetes/pkg/api"
	conversionPackagePath = "k8s.io/kubernetes/pkg/conversion"
)

// genDeepCopy produces a file with a set for a single type.
type genDeepCopy struct {
	generator.DefaultGen
	targetPackage    string
	imports          namer.ImportTracker
	typesForInit     []*types.Type
	generateInitFunc bool
}

func NewGenDeepCopy(sanitizedName, targetPackage string, generateInitFunc bool) generator.Generator {
	return &genDeepCopy{
		DefaultGen: generator.DefaultGen{
			OptionalName: sanitizedName,
		},
		targetPackage:    targetPackage,
		imports:          generator.NewImportTracker(),
		typesForInit:     make([]*types.Type, 0),
		generateInitFunc: generateInitFunc,
	}
}

func (g *genDeepCopy) Namers(c *generator.Context) namer.NameSystems {
	// Have the raw namer for this file track what it imports.
	return namer.NameSystems{"raw": namer.NewRawNamer(g.targetPackage, g.imports)}
}

// Filter ignores all but one type because we're making a single file per type.
func (g *genDeepCopy) Filter(c *generator.Context, t *types.Type) bool {
	// Filter out all types not copyable within the package.
	copyable := copyableWithinPackage(t)
	if copyable {
		g.typesForInit = append(g.typesForInit, t)
	}
	return copyable
}

func copyableWithinPackage(t *types.Type) bool {
	if !strings.HasPrefix(t.Name.Package, "k8s.io/kubernetes/") {
		return false
	}
	if types.ExtractCommentTags("+", t.CommentLines)["gencopy"] == "false" {
		return false
	}
	// TODO: Consider generating functions for other kinds too.
	if t.Kind != types.Struct {
		return false
	}
	// Also, filter out private types.
	if namer.IsPrivateGoName(t.Name.Name) {
		return false
	}
	return true
}

func (g *genDeepCopy) isOtherPackage(pkg string) bool {
	if pkg == g.targetPackage {
		return false
	}
	if strings.HasSuffix(pkg, "\""+g.targetPackage+"\"") {
		return false
	}
	return true
}

func (g *genDeepCopy) Imports(c *generator.Context) (imports []string) {
	importLines := []string{}
	if g.isOtherPackage(apiPackagePath) && g.generateInitFunc {
		importLines = append(importLines, "api \""+apiPackagePath+"\"")
	}
	if g.isOtherPackage(conversionPackagePath) {
		importLines = append(importLines, "conversion \""+conversionPackagePath+"\"")
	}
	for _, singleImport := range g.imports.ImportLines() {
		if g.isOtherPackage(singleImport) {
			importLines = append(importLines, singleImport)
		}
	}
	return importLines
}

func argsFromType(t *types.Type) interface{} {
	return map[string]interface{}{
		"type": t,
	}
}

func (g *genDeepCopy) funcNameTmpl(t *types.Type) string {
	tmpl := "DeepCopy_$.type|public$"
	g.imports.AddType(t)
	if t.Name.Package != g.targetPackage {
		tmpl = g.imports.LocalNameOf(t.Name.Package) + "." + tmpl
	}
	return tmpl
}

func (g *genDeepCopy) Init(c *generator.Context, w io.Writer) error {
	if !g.generateInitFunc {
		// TODO: We should come up with a solution to register all generated
		// deep-copy functions. However, for now, to avoid import cycles
		// we register only those explicitly requested.
		return nil
	}
	sw := generator.NewSnippetWriter(w, c, "$", "$")
	sw.Do("func init() {\n", nil)
	if g.targetPackage == apiPackagePath {
		sw.Do("if err := Scheme.AddGeneratedDeepCopyFuncs(\n", nil)
	} else {
		sw.Do("if err := api.Scheme.AddGeneratedDeepCopyFuncs(\n", nil)
	}
	for _, t := range g.typesForInit {
		sw.Do(fmt.Sprintf("%s,\n", g.funcNameTmpl(t)), argsFromType(t))
	}
	sw.Do("); err != nil {\n", nil)
	sw.Do("// if one of the deep copy functions is malformed, detect it immediately.\n", nil)
	sw.Do("panic(err)\n", nil)
	sw.Do("}\n", nil)
	sw.Do("}\n\n", nil)
	return sw.Error()
}

// GenerateType makes the body of a file implementing a set for type t.
func (g *genDeepCopy) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "$", "$")
	funcName := g.funcNameTmpl(t)
	if g.targetPackage == conversionPackagePath {
		sw.Do(fmt.Sprintf("func %s(in $.type|raw$, out *$.type|raw$, c *Cloner) error {\n", funcName), argsFromType(t))
	} else {
		sw.Do(fmt.Sprintf("func %s(in $.type|raw$, out *$.type|raw$, c *conversion.Cloner) error {\n", funcName), argsFromType(t))
	}
	g.generateFor(t, sw)
	sw.Do("return nil\n", nil)
	sw.Do("}\n\n", nil)
	return sw.Error()
}

// we use the system of shadowing 'in' and 'out' so that the same code is valid
// at any nesting level. This makes the autogenerator easy to understand, and
// the compiler shouldn't care.
func (g *genDeepCopy) generateFor(t *types.Type, sw *generator.SnippetWriter) {
	var f func(*types.Type, *generator.SnippetWriter)
	switch t.Kind {
	case types.Builtin:
		f = g.doBuiltin
	case types.Map:
		f = g.doMap
	case types.Slice:
		f = g.doSlice
	case types.Struct:
		f = g.doStruct
	case types.Interface:
		f = g.doInterface
	case types.Pointer:
		f = g.doPointer
	case types.Alias:
		f = g.doAlias
	default:
		f = g.doUnknown
	}
	f(t, sw)
}

func (g *genDeepCopy) doBuiltin(t *types.Type, sw *generator.SnippetWriter) {
	sw.Do("*out = in\n", nil)
}

func (g *genDeepCopy) doMap(t *types.Type, sw *generator.SnippetWriter) {
	sw.Do("*out = make($.|raw$)\n", t)
	if t.Key.IsAssignable() {
		sw.Do("for key, val := range in {\n", nil)
		if t.Elem.IsAssignable() {
			sw.Do("(*out)[key] = val\n", nil)
		} else {
			if copyableWithinPackage(t.Elem) {
				sw.Do("newVal := new($.|raw$)\n", t.Elem)
				funcName := g.funcNameTmpl(t.Elem)
				sw.Do(fmt.Sprintf("if err := %s(val, newVal, c); err != nil {\n", funcName), argsFromType(t.Elem))
				sw.Do("return err\n", nil)
				sw.Do("}\n", nil)
				sw.Do("(*out)[key] = *newVal\n", nil)
			} else {
				sw.Do("if newVal, err := c.DeepCopy(val); err != nil {\n", nil)
				sw.Do("return err\n", nil)
				sw.Do("} else {\n", nil)
				sw.Do("(*out)[key] = newVal.($.|raw$)\n", t.Elem)
				sw.Do("}\n", nil)
			}
		}
	} else {
		// TODO: Implement it when necessary.
		sw.Do("for range in {\n", nil)
		sw.Do("// FIXME: Copying unassignable keys unsupported $.|raw$\n", t.Key)
	}
	sw.Do("}\n", nil)
}

func (g *genDeepCopy) doSlice(t *types.Type, sw *generator.SnippetWriter) {
	sw.Do("*out = make($.|raw$, len(in))\n", t)
	if t.Elem.Kind == types.Builtin {
		sw.Do("copy(*out, in)\n", nil)
	} else {
		sw.Do("for i := range in {\n", nil)
		if t.Elem.IsAssignable() {
			sw.Do("(*out)[i] = in[i]\n", nil)
		} else if copyableWithinPackage(t.Elem) {
			funcName := g.funcNameTmpl(t.Elem)
			sw.Do(fmt.Sprintf("if err := %s(in[i], &(*out)[i], c); err != nil {\n", funcName), argsFromType(t.Elem))
			sw.Do("return err\n", nil)
			sw.Do("}\n", nil)
		} else {
			sw.Do("if newVal, err := c.DeepCopy(in[i]); err != nil {\n", nil)
			sw.Do("return err\n", nil)
			sw.Do("} else {\n", nil)
			sw.Do("(*out)[i] = newVal.($.|raw$)\n", t.Elem)
			sw.Do("}\n", nil)
		}
		sw.Do("}\n", nil)
	}
}

func (g *genDeepCopy) doStruct(t *types.Type, sw *generator.SnippetWriter) {
	for _, m := range t.Members {
		args := map[string]interface{}{
			"type": m.Type,
			"name": m.Name,
		}
		switch m.Type.Kind {
		case types.Builtin:
			sw.Do("out.$.name$ = in.$.name$\n", args)
		case types.Map, types.Slice, types.Pointer:
			sw.Do("if in.$.name$ != nil {\n", args)
			sw.Do("in, out := in.$.name$, &out.$.name$\n", args)
			g.generateFor(m.Type, sw)
			sw.Do("} else {\n", nil)
			sw.Do("out.$.name$ = nil\n", args)
			sw.Do("}\n", nil)
		case types.Struct:
			if copyableWithinPackage(m.Type) {
				funcName := g.funcNameTmpl(m.Type)
				sw.Do(fmt.Sprintf("if err := %s(in.$.name$, &out.$.name$, c); err != nil {\n", funcName), args)
				sw.Do("return err\n", nil)
				sw.Do("}\n", nil)
			} else {
				sw.Do("if newVal, err := c.DeepCopy(in.$.name$); err != nil {\n", args)
				sw.Do("return err\n", nil)
				sw.Do("} else {\n", nil)
				sw.Do("out.$.name$ = newVal.($.type|raw$)\n", args)
				sw.Do("}\n", nil)
			}
		default:
			if m.Type.Kind == types.Alias && m.Type.Underlying.Kind == types.Builtin {
				sw.Do("out.$.name$ = in.$.name$\n", args)
			} else {
				sw.Do("if in.$.name$ == nil {\n", args)
				sw.Do("out.$.name$ = nil\n", args)
				sw.Do("} else if newVal, err := c.DeepCopy(in.$.name$); err != nil {\n", args)
				sw.Do("return err\n", nil)
				sw.Do("} else {\n", nil)
				sw.Do("out.$.name$ = newVal.($.type|raw$)\n", args)
				sw.Do("}\n", nil)
			}
		}
	}
}

func (g *genDeepCopy) doInterface(t *types.Type, sw *generator.SnippetWriter) {
	// TODO: Add support for interfaces.
	g.doUnknown(t, sw)
}

func (g *genDeepCopy) doPointer(t *types.Type, sw *generator.SnippetWriter) {
	sw.Do("*out = new($.Elem|raw$)\n", t)
	if t.Elem.Kind == types.Builtin {
		sw.Do("**out = *in", nil)
	} else if copyableWithinPackage(t.Elem) {
		funcName := g.funcNameTmpl(t.Elem)
		sw.Do(fmt.Sprintf("if err := %s(*in, *out, c); err != nil {\n", funcName), argsFromType(t.Elem))
		sw.Do("return err\n", nil)
		sw.Do("}\n", nil)
	} else {
		sw.Do("if newVal, err := c.DeepCopy(*in); err != nil {\n", nil)
		sw.Do("return err\n", nil)
		sw.Do("} else {\n", nil)
		sw.Do("**out = newVal.($.|raw$)\n", t.Elem)
		sw.Do("}\n", nil)
	}
}

func (g *genDeepCopy) doAlias(t *types.Type, sw *generator.SnippetWriter) {
	// TODO: Add support for aliases.
	g.doUnknown(t, sw)
}

func (g *genDeepCopy) doUnknown(t *types.Type, sw *generator.SnippetWriter) {
	sw.Do("// FIXME: Type $.|raw$ is unsupported.\n", t)
}
