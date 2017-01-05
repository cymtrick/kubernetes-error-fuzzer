/*
Copyright 2015 The Kubernetes Authors.

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

	"k8s.io/gengo/args"
	"k8s.io/gengo/examples/set-gen/sets"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"

	"github.com/golang/glog"
)

// CustomArgs is used tby the go2idl framework to pass args specific to this
// generator.
type CustomArgs struct {
	BoundingDirs []string // Only deal with types rooted under these dirs.
}

// This is the comment tag that carries parameters for deep-copy generation.
const tagName = "k8s:deepcopy-gen"

// Known values for the comment tag.
const tagValuePackage = "package"

// tagValue holds parameters from a tagName tag.
type tagValue struct {
	value    string
	register bool
}

func extractTag(comments []string) *tagValue {
	tagVals := types.ExtractCommentTags("+", comments)[tagName]
	if tagVals == nil {
		// No match for the tag.
		return nil
	}
	// If there are multiple values, abort.
	if len(tagVals) > 1 {
		glog.Fatalf("Found %d %s tags: %q", len(tagVals), tagName, tagVals)
	}

	// If we got here we are returning something.
	tag := &tagValue{}

	// Get the primary value.
	parts := strings.Split(tagVals[0], ",")
	if len(parts) >= 1 {
		tag.value = parts[0]
	}

	// Parse extra arguments.
	parts = parts[1:]
	for i := range parts {
		kv := strings.SplitN(parts[i], "=", 2)
		k := kv[0]
		v := ""
		if len(kv) == 2 {
			v = kv[1]
		}
		switch k {
		case "register":
			if v != "false" {
				tag.register = true
			}
		default:
			glog.Fatalf("Unsupported %s param: %q", tagName, parts[i])
		}
	}
	return tag
}

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

	inputs := sets.NewString(context.Inputs...)
	packages := generator.Packages{}
	header := append([]byte(fmt.Sprintf("// +build !%s\n\n", arguments.GeneratedBuildTag)), boilerplate...)
	header = append(header, []byte(`
	    // This file was autogenerated by deepcopy-gen. Do not edit it manually!

		`)...)

	boundingDirs := []string{}
	if customArgs, ok := arguments.CustomArgs.(*CustomArgs); ok {
		if customArgs.BoundingDirs == nil {
			customArgs.BoundingDirs = context.Inputs
		}
		for i := range customArgs.BoundingDirs {
			// Strip any trailing slashes - they are not exactly "correct" but
			// this is friendlier.
			boundingDirs = append(boundingDirs, strings.TrimRight(customArgs.BoundingDirs[i], "/"))
		}
	}

	for i := range inputs {
		glog.V(5).Infof("Considering pkg %q", i)
		pkg := context.Universe[i]
		if pkg == nil {
			// If the input had no Go files, for example.
			continue
		}

		ptag := extractTag(pkg.Comments)
		ptagValue := ""
		ptagRegister := false
		if ptag != nil {
			ptagValue = ptag.value
			if ptagValue != tagValuePackage {
				glog.Fatalf("Package %v: unsupported %s value: %q", i, tagName, ptagValue)
			}
			ptagRegister = ptag.register
			glog.V(5).Infof("  tag.value: %q, tag.register: %t", ptagValue, ptagRegister)
		} else {
			glog.V(5).Infof("  no tag")
		}

		// If the pkg-scoped tag says to generate, we can skip scanning types.
		pkgNeedsGeneration := (ptagValue == tagValuePackage)
		if !pkgNeedsGeneration {
			// If the pkg-scoped tag did not exist, scan all types for one that
			// explicitly wants generation.
			for _, t := range pkg.Types {
				glog.V(5).Infof("  considering type %q", t.Name.String())
				ttag := extractTag(t.CommentLines)
				if ttag != nil && ttag.value == "true" {
					glog.V(5).Infof("    tag=true")
					if !copyableType(t) {
						glog.Fatalf("Type %v requests deepcopy generation but is not copyable", t)
					}
					pkgNeedsGeneration = true
					break
				}
			}
		}

		if pkgNeedsGeneration {
			glog.V(3).Infof("Package %q needs generation", i)
			packages = append(packages,
				&generator.DefaultPackage{
					PackageName: strings.Split(filepath.Base(pkg.Path), ".")[0],
					PackagePath: pkg.Path,
					HeaderText:  header,
					GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
						return []generator.Generator{
							NewGenDeepCopy(arguments.OutputFileBaseName, pkg.Path, boundingDirs, (ptagValue == tagValuePackage), ptagRegister),
						}
					},
					FilterFunc: func(c *generator.Context, t *types.Type) bool {
						return t.Name.Package == pkg.Path
					},
				})
		}
	}
	return packages
}

const (
	conversionPackagePath = "k8s.io/kubernetes/pkg/conversion"
	runtimePackagePath    = "k8s.io/kubernetes/pkg/runtime"
)

// genDeepCopy produces a file with autogenerated deep-copy functions.
type genDeepCopy struct {
	generator.DefaultGen
	targetPackage string
	boundingDirs  []string
	allTypes      bool
	registerTypes bool
	imports       namer.ImportTracker
	typesForInit  []*types.Type
}

func NewGenDeepCopy(sanitizedName, targetPackage string, boundingDirs []string, allTypes, registerTypes bool) generator.Generator {
	return &genDeepCopy{
		DefaultGen: generator.DefaultGen{
			OptionalName: sanitizedName,
		},
		targetPackage: targetPackage,
		boundingDirs:  boundingDirs,
		allTypes:      allTypes,
		registerTypes: registerTypes,
		imports:       generator.NewImportTracker(),
		typesForInit:  make([]*types.Type, 0),
	}
}

func (g *genDeepCopy) Namers(c *generator.Context) namer.NameSystems {
	// Have the raw namer for this file track what it imports.
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.targetPackage, g.imports),
		"dcFnName": &dcFnNamer{
			public:    deepCopyNamer(),
			tracker:   g.imports,
			myPackage: g.targetPackage,
		},
	}
}

func (g *genDeepCopy) Filter(c *generator.Context, t *types.Type) bool {
	// Filter out types not being processed or not copyable within the package.
	enabled := g.allTypes
	if !enabled {
		ttag := extractTag(t.CommentLines)
		if ttag != nil && ttag.value == "true" {
			enabled = true
		}
	}
	if !enabled {
		return false
	}
	if !copyableType(t) {
		glog.V(2).Infof("Type %v is not copyable", t)
		return false
	}
	glog.V(4).Infof("Type %v is copyable", t)
	g.typesForInit = append(g.typesForInit, t)
	return true
}

func (g *genDeepCopy) copyableAndInBounds(t *types.Type) bool {
	if !copyableType(t) {
		return false
	}
	// Only packages within the restricted range can be processed.
	if !isRootedUnder(t.Name.Package, g.boundingDirs) {
		return false
	}
	return true
}

// hasDeepCopyMethod returns true if an appropriate DeepCopy() method is
// defined for the given type.  This allows more efficient deep copy
// implementations to be defined by the type's author.  The correct signature
// for a type T is:
//    func (t T) DeepCopy() T
// or:
//    func (t *T) DeepCopyt() T
func hasDeepCopyMethod(t *types.Type) bool {
	for mn, mt := range t.Methods {
		if mn != "DeepCopy" {
			continue
		}
		if len(mt.Signature.Parameters) != 0 {
			return false
		}
		if len(mt.Signature.Results) != 1 || mt.Signature.Results[0].Name != t.Name {
			return false
		}
		return true
	}
	return false
}

func isRootedUnder(pkg string, roots []string) bool {
	// Add trailing / to avoid false matches, e.g. foo/bar vs foo/barn.  This
	// assumes that bounding dirs do not have trailing slashes.
	pkg = pkg + "/"
	for _, root := range roots {
		if strings.HasPrefix(pkg, root+"/") {
			return true
		}
	}
	return false
}

func copyableType(t *types.Type) bool {
	// If the type opts out of copy-generation, stop.
	ttag := extractTag(t.CommentLines)
	if ttag != nil && ttag.value == "false" {
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
	for _, singleImport := range g.imports.ImportLines() {
		if g.isOtherPackage(singleImport) {
			importLines = append(importLines, singleImport)
		}
	}
	return importLines
}

func argsFromType(t *types.Type) generator.Args {
	return generator.Args{
		"type": t,
	}
}

type dcFnNamer struct {
	public    namer.Namer
	tracker   namer.ImportTracker
	myPackage string
}

func (n *dcFnNamer) Name(t *types.Type) string {
	pubName := n.public.Name(t)
	n.tracker.AddType(t)
	if t.Name.Package == n.myPackage {
		return "DeepCopy_" + pubName
	}
	return fmt.Sprintf("%s.DeepCopy_%s", n.tracker.LocalNameOf(t.Name.Package), pubName)
}

func (g *genDeepCopy) Init(c *generator.Context, w io.Writer) error {
	cloner := c.Universe.Type(types.Name{Package: conversionPackagePath, Name: "Cloner"})
	g.imports.AddType(cloner)
	if !g.registerTypes {
		sw := generator.NewSnippetWriter(w, c, "$", "$")
		sw.Do("// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.\n", nil)
		sw.Do("func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc{\n", nil)
		sw.Do("return []conversion.GeneratedDeepCopyFunc{\n", nil)
		for _, t := range g.typesForInit {
			args := argsFromType(t).
				With("typeof", c.Universe.Package("reflect").Function("TypeOf"))
			sw.Do("{Fn: $.type|dcFnName$, InType: $.typeof|raw$(&$.type|raw${})},\n", args)
		}
		sw.Do("}\n", nil)
		sw.Do("}\n\n", nil)
		return sw.Error()
	}
	glog.V(5).Infof("Registering types in pkg %q", g.targetPackage)

	sw := generator.NewSnippetWriter(w, c, "$", "$")
	sw.Do("func init() {\n", nil)
	sw.Do("SchemeBuilder.Register(RegisterDeepCopies)\n", nil)
	sw.Do("}\n\n", nil)

	scheme := c.Universe.Type(types.Name{Package: runtimePackagePath, Name: "Scheme"})
	schemePtr := &types.Type{
		Kind: types.Pointer,
		Elem: scheme,
	}
	sw.Do("// RegisterDeepCopies adds deep-copy functions to the given scheme. Public\n", nil)
	sw.Do("// to allow building arbitrary schemes.\n", nil)
	sw.Do("func RegisterDeepCopies(scheme $.|raw$) error {\n", schemePtr)
	sw.Do("return scheme.AddGeneratedDeepCopyFuncs(\n", nil)
	for _, t := range g.typesForInit {
		args := argsFromType(t).
			With("typeof", c.Universe.Package("reflect").Function("TypeOf"))
		sw.Do("conversion.GeneratedDeepCopyFunc{Fn: $.type|dcFnName$, InType: $.typeof|raw$(&$.type|raw${})},\n", args)
	}
	sw.Do(")\n", nil)
	sw.Do("}\n\n", nil)
	return sw.Error()
}

func (g *genDeepCopy) needsGeneration(t *types.Type) bool {
	tag := extractTag(t.CommentLines)
	tv := ""
	if tag != nil {
		tv = tag.value
		if tv != "true" && tv != "false" {
			glog.Fatalf("Type %v: unsupported %s value: %q", t, tagName, tag.value)
		}
	}
	if g.allTypes && tv == "false" {
		// The whole package is being generated, but this type has opted out.
		glog.V(5).Infof("Not generating for type %v because type opted out", t)
		return false
	}
	if !g.allTypes && tv != "true" {
		// The whole package is NOT being generated, and this type has NOT opted in.
		glog.V(5).Infof("Not generating for type %v because type did not opt in", t)
		return false
	}
	return true
}

func (g *genDeepCopy) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	if !g.needsGeneration(t) {
		return nil
	}
	glog.V(5).Infof("Generating deepcopy function for type %v", t)

	sw := generator.NewSnippetWriter(w, c, "$", "$")
	args := argsFromType(t).
		With("clonerType", types.Ref(conversionPackagePath, "Cloner"))
	sw.Do("func $.type|dcFnName$(in interface{}, out interface{}, c *$.clonerType|raw$) error {{\n", args)
	sw.Do("in := in.(*$.type|raw$)\nout := out.(*$.type|raw$)\n", argsFromType(t))
	g.generateFor(t, sw)
	sw.Do("return nil\n", nil)
	sw.Do("}}\n\n", nil)
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
	sw.Do("*out = *in\n", nil)
}

func (g *genDeepCopy) doMap(t *types.Type, sw *generator.SnippetWriter) {
	sw.Do("*out = make($.|raw$)\n", t)
	if t.Key.IsAssignable() {
		switch {
		case hasDeepCopyMethod(t.Elem):
			sw.Do("for key, val := range *in {\n", nil)
			sw.Do("(*out)[key] = val.DeepCopy()\n", nil)
			sw.Do("}\n", nil)
		case t.Elem.IsAnonymousStruct():
			sw.Do("for key := range *in {\n", nil)
			sw.Do("(*out)[key] = struct{}{}\n", nil)
			sw.Do("}\n", nil)
		case t.Elem.IsAssignable():
			sw.Do("for key, val := range *in {\n", nil)
			sw.Do("(*out)[key] = val\n", nil)
			sw.Do("}\n", nil)
		default:
			sw.Do("for key, val := range *in {\n", nil)
			if g.copyableAndInBounds(t.Elem) {
				sw.Do("newVal := new($.|raw$)\n", t.Elem)
				sw.Do("if err := $.type|dcFnName$(&val, newVal, c); err != nil {\n", argsFromType(t.Elem))
				sw.Do("return err\n", nil)
				sw.Do("}\n", nil)
				sw.Do("(*out)[key] = *newVal\n", nil)
			} else {
				sw.Do("if newVal, err := c.DeepCopy(&val); err != nil {\n", nil)
				sw.Do("return err\n", nil)
				sw.Do("} else {\n", nil)
				sw.Do("(*out)[key] = *newVal.(*$.|raw$)\n", t.Elem)
				sw.Do("}\n", nil)
			}
			sw.Do("}\n", nil)
		}
	} else {
		// TODO: Implement it when necessary.
		sw.Do("for range *in {\n", nil)
		sw.Do("// FIXME: Copying unassignable keys unsupported $.|raw$\n", t.Key)
		sw.Do("}\n", nil)
	}
}

func (g *genDeepCopy) doSlice(t *types.Type, sw *generator.SnippetWriter) {
	sw.Do("*out = make($.|raw$, len(*in))\n", t)
	if t.Elem.Kind == types.Builtin {
		sw.Do("copy(*out, *in)\n", nil)
	} else {
		sw.Do("for i := range *in {\n", nil)
		if hasDeepCopyMethod(t.Elem) {
			sw.Do("(*out)[i] = (*in)[i].DeepCopy()\n", nil)
		} else if t.Elem.IsAssignable() {
			sw.Do("(*out)[i] = (*in)[i]\n", nil)
		} else if g.copyableAndInBounds(t.Elem) {
			sw.Do("if err := $.type|dcFnName$(&(*in)[i], &(*out)[i], c); err != nil {\n", argsFromType(t.Elem))
			sw.Do("return err\n", nil)
			sw.Do("}\n", nil)
		} else {
			sw.Do("if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {\n", nil)
			sw.Do("return err\n", nil)
			sw.Do("} else {\n", nil)
			sw.Do("(*out)[i] = *newVal.(*$.|raw$)\n", t.Elem)
			sw.Do("}\n", nil)
		}
		sw.Do("}\n", nil)
	}
}

func (g *genDeepCopy) doStruct(t *types.Type, sw *generator.SnippetWriter) {
	if hasDeepCopyMethod(t) {
		sw.Do("*out = in.DeepCopy()\n", nil)
		return
	}

	// Simple copy covers a lot of cases.
	sw.Do("*out = *in\n", nil)

	// Now fix-up fields as needed.
	for _, m := range t.Members {
		t := m.Type
		hasMethod := hasDeepCopyMethod(t)
		if t.Kind == types.Alias {
			copied := *t.Underlying
			copied.Name = t.Name
			t = &copied
		}
		args := generator.Args{
			"type": t,
			"kind": t.Kind,
			"name": m.Name,
		}
		switch t.Kind {
		case types.Builtin:
			if hasMethod {
				sw.Do("out.$.name$ = in.$.name$.DeepCopy()\n", args)
			}
		case types.Map, types.Slice, types.Pointer:
			if hasMethod {
				sw.Do("if in.$.name$ != nil {\n", args)
				sw.Do("out.$.name$ = in.$.name$.DeepCopy()\n", args)
				sw.Do("}\n", nil)
			} else {
				// Fixup non-nil reference-sematic types.
				sw.Do("if in.$.name$ != nil {\n", args)
				sw.Do("in, out := &in.$.name$, &out.$.name$\n", args)
				g.generateFor(t, sw)
				sw.Do("}\n", nil)
			}
		case types.Struct:
			if hasMethod {
				sw.Do("out.$.name$ = in.$.name$.DeepCopy()\n", args)
			} else if t.IsAssignable() {
				// Nothing else needed.
			} else if g.copyableAndInBounds(t) {
				// Not assignable but should have a deepcopy function.
				// TODO: do a topological sort of packages and ensure that this works, else inline it.
				sw.Do("if err := $.type|dcFnName$(&in.$.name$, &out.$.name$, c); err != nil {\n", args)
				sw.Do("return err\n", nil)
				sw.Do("}\n", nil)
			} else {
				// Fall back on the slow-path and hope it works.
				// TODO: don't depend on kubernetes code for this
				sw.Do("if newVal, err := c.DeepCopy(&in.$.name$); err != nil {\n", args)
				sw.Do("return err\n", nil)
				sw.Do("} else {\n", nil)
				sw.Do("out.$.name$ = *newVal.(*$.type|raw$)\n", args)
				sw.Do("}\n", nil)
			}
		default:
			// Interfaces, Arrays, and other Kinds we don't understand.
			sw.Do("// in.$.name$ is kind '$.kind$'\n", args)
			if hasMethod {
				sw.Do("if in.$.name$ != nil {\n", args)
				sw.Do("out.$.name$ = in.$.name$.DeepCopy()\n", args)
				sw.Do("}\n", args)
			} else {
				// TODO: don't depend on kubernetes code for this
				sw.Do("if in.$.name$ != nil {\n", args)
				sw.Do("if newVal, err := c.DeepCopy(&in.$.name$); err != nil {\n", args)
				sw.Do("return err\n", nil)
				sw.Do("} else {\n", nil)
				sw.Do("out.$.name$ = *newVal.(*$.type|raw$)\n", args)
				sw.Do("}\n", nil)
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
	if hasDeepCopyMethod(t.Elem) {
		sw.Do("*out = new($.Elem|raw$)\n", t)
		sw.Do("**out = (*in).DeepCopy()\n", nil)
	} else if t.Elem.IsAssignable() {
		sw.Do("*out = new($.Elem|raw$)\n", t)
		sw.Do("**out = **in", nil)
	} else if g.copyableAndInBounds(t.Elem) {
		sw.Do("*out = new($.Elem|raw$)\n", t)
		sw.Do("if err := $.type|dcFnName$(*in, *out, c); err != nil {\n", argsFromType(t.Elem))
		sw.Do("return err\n", nil)
		sw.Do("}\n", nil)
	} else {
		sw.Do("if newVal, err := c.DeepCopy(*in); err != nil {\n", nil)
		sw.Do("return err\n", nil)
		sw.Do("} else {\n", nil)
		sw.Do("*out = newVal.(*$.|raw$)\n", t.Elem)
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
