/*
Copyright 2016 The Kubernetes Authors.

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
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
	clientgentypes "k8s.io/kubernetes/cmd/libs/go2idl/client-gen/types"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

// NameSystems returns the name system used by the generators in this package.
func NameSystems() namer.NameSystems {
	pluralExceptions := map[string]string{
		"Endpoints": "Endpoints",
	}
	return namer.NameSystems{
		"public":             namer.NewPublicNamer(0),
		"private":            namer.NewPrivateNamer(0),
		"raw":                namer.NewRawNamer("", nil),
		"publicPlural":       namer.NewPublicPluralNamer(pluralExceptions),
		"allLowercasePlural": namer.NewAllLowercasePluralNamer(pluralExceptions),
		"lowercaseSingular":  &lowercaseSingularNamer{},
	}
}

// lowercaseSingularNamer implements Namer
type lowercaseSingularNamer struct{}

// Name returns t's name in all lowercase.
func (n *lowercaseSingularNamer) Name(t *types.Type) string {
	return strings.ToLower(t.Name.Name)
}

// DefaultNameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func DefaultNameSystem() string {
	return "public"
}

// generatedBy returns information about the arguments used to invoke
// lister-gen.
func generatedBy() string {
	var cmdArgs string
	pflag.VisitAll(func(f *pflag.Flag) {
		if !f.Changed || f.Name == "verify-only" {
			return
		}
		cmdArgs += fmt.Sprintf("--%s=%s ", f.Name, f.Value)
	})
	return fmt.Sprintf("\n// This file was automatically generated by lister-gen with arguments: %s\n\n", cmdArgs)
}

// Packages makes the client package definition.
func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	boilerplate, err := arguments.LoadGoBoilerplate()
	if err != nil {
		glog.Fatalf("Failed loading boilerplate: %v", err)
	}

	boilerplate = append(boilerplate, []byte(generatedBy())...)

	var packageList generator.Packages
	for _, inputDir := range arguments.InputDirs {
		p := context.Universe.Package(inputDir)

		objectMeta, internal, err := objectMetaForPackage(p)
		if err != nil {
			glog.Fatal(err)
		}
		if objectMeta == nil {
			// no types in this package had genclient
			continue
		}

		var gv clientgentypes.GroupVersion
		var internalGVPkg string

		if internal {
			lastSlash := strings.LastIndex(p.Path, "/")
			if lastSlash == -1 {
				glog.Fatalf("error constructing internal group version for package %q", p.Path)
			}
			gv.Group = clientgentypes.Group(p.Path[lastSlash+1:])
			internalGVPkg = p.Path
		} else {
			parts := strings.Split(p.Path, "/")
			gv.Group = clientgentypes.Group(parts[len(parts)-2])
			gv.Version = clientgentypes.Version(parts[len(parts)-1])

			internalGVPkg = strings.Join(parts[0:len(parts)-1], "/")
		}

		var typesToGenerate []*types.Type
		for _, t := range p.Types {
			// filter out types which dont have genclient=true.
			if extractBoolTagOrDie("genclient", t.SecondClosestCommentLines) == false {
				continue
			}
			typesToGenerate = append(typesToGenerate, t)
		}
		orderer := namer.Orderer{Namer: namer.NewPrivateNamer(0)}
		typesToGenerate = orderer.OrderTypes(typesToGenerate)

		packagePath := filepath.Join(arguments.OutputPackagePath, strings.ToLower(gv.Group.NonEmpty()), strings.ToLower(gv.Version.NonEmpty()))
		packageList = append(packageList, &generator.DefaultPackage{
			PackageName: strings.ToLower(gv.Version.NonEmpty()),
			PackagePath: packagePath,
			HeaderText:  boilerplate,
			GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
				generators = append(generators, &expansionGenerator{
					DefaultGen: generator.DefaultGen{
						OptionalName: "expansion_generated",
					},
					packagePath: filepath.Join(arguments.OutputBase, packagePath),
					types:       typesToGenerate,
				})

				for _, t := range typesToGenerate {
					generators = append(generators, &listerGenerator{
						DefaultGen: generator.DefaultGen{
							OptionalName: strings.ToLower(t.Name.Name),
						},
						outputPackage:  arguments.OutputPackagePath,
						groupVersion:   gv,
						internalGVPkg:  internalGVPkg,
						typeToGenerate: t,
						imports:        generator.NewImportTracker(),
						objectMeta:     objectMeta,
					})
				}
				return generators
			},
			FilterFunc: func(c *generator.Context, t *types.Type) bool {
				// piggy-back on types that are tagged for client-gen
				return extractBoolTagOrDie("genclient", t.SecondClosestCommentLines) == true
			},
		})
	}

	return packageList
}

// objectMetaForPackage returns the type of ObjectMeta used by package p.
func objectMetaForPackage(p *types.Package) (*types.Type, bool, error) {
	generatingForPackage := false
	for _, t := range p.Types {
		// filter out types which dont have genclient=true.
		if extractBoolTagOrDie("genclient", t.SecondClosestCommentLines) == false {
			continue
		}
		generatingForPackage = true
		for _, member := range t.Members {
			if member.Name == "ObjectMeta" {
				return member.Type, isInternal(member), nil
			}
		}
	}
	if generatingForPackage {
		return nil, false, fmt.Errorf("unable to find ObjectMeta for any types in package %s", p.Path)
	}
	return nil, false, nil
}

// isInternal returns true if the tags for a member do not contain a json tag
func isInternal(m types.Member) bool {
	return !strings.Contains(m.Tags, "json")
}

// listerGenerator produces a file of listers for a given GroupVersion and
// type.
type listerGenerator struct {
	generator.DefaultGen
	outputPackage  string
	groupVersion   clientgentypes.GroupVersion
	internalGVPkg  string
	typeToGenerate *types.Type
	imports        namer.ImportTracker
	objectMeta     *types.Type
}

var _ generator.Generator = &listerGenerator{}

func (g *listerGenerator) Filter(c *generator.Context, t *types.Type) bool {
	return t == g.typeToGenerate
}

func (g *listerGenerator) Namers(c *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
	}
}

func (g *listerGenerator) Imports(c *generator.Context) (imports []string) {
	imports = append(imports, g.imports.ImportLines()...)
	imports = append(imports, "k8s.io/apimachinery/pkg/api/errors")
	imports = append(imports, "k8s.io/apimachinery/pkg/labels")
	// for Indexer
	imports = append(imports, "k8s.io/kubernetes/pkg/client/cache")
	return
}

func (g *listerGenerator) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "$", "$")

	glog.V(5).Infof("processing type %v", t)
	m := map[string]interface{}{
		"Resource":   c.Universe.Function(types.Name{Package: g.internalGVPkg, Name: "Resource"}),
		"type":       t,
		"objectMeta": g.objectMeta,
	}

	namespaced := !extractBoolTagOrDie("nonNamespaced", t.SecondClosestCommentLines)
	if namespaced {
		sw.Do(typeListerInterface, m)
	} else {
		sw.Do(typeListerInterface_NonNamespaced, m)
	}

	sw.Do(typeListerStruct, m)
	sw.Do(typeListerConstructor, m)
	sw.Do(typeLister_List, m)

	if namespaced {
		sw.Do(typeLister_NamespaceLister, m)
		sw.Do(namespaceListerInterface, m)
		sw.Do(namespaceListerStruct, m)
		sw.Do(namespaceLister_List, m)
		sw.Do(namespaceLister_Get, m)
	} else {
		sw.Do(typeLister_NonNamespacedGet, m)
	}

	return sw.Error()
}

var typeListerInterface = `
// $.type|public$Lister helps list $.type|publicPlural$.
type $.type|public$Lister interface {
	// List lists all $.type|publicPlural$ in the indexer.
	List(selector labels.Selector) (ret []*$.type|raw$, err error)
	// $.type|publicPlural$ returns an object that can list and get $.type|publicPlural$.
	$.type|publicPlural$(namespace string) $.type|public$NamespaceLister
	$.type|public$ListerExpansion
}
`

var typeListerInterface_NonNamespaced = `
// $.type|public$Lister helps list $.type|publicPlural$.
type $.type|public$Lister interface {
	// List lists all $.type|publicPlural$ in the indexer.
	List(selector labels.Selector) (ret []*$.type|raw$, err error)
	// Get retrieves the $.type|public$ from the index for a given name.
	Get(name string) (*$.type|raw$, error)
	$.type|public$ListerExpansion
}
`

var typeListerStruct = `
// $.type|private$Lister implements the $.type|public$Lister interface.
type $.type|private$Lister struct {
	indexer cache.Indexer
}
`

var typeListerConstructor = `
// New$.type|public$Lister returns a new $.type|public$Lister.
func New$.type|public$Lister(indexer cache.Indexer) $.type|public$Lister {
	return &$.type|private$Lister{indexer: indexer}
}
`

var typeLister_List = `
// List lists all $.type|publicPlural$ in the indexer.
func (s *$.type|private$Lister) List(selector labels.Selector) (ret []*$.type|raw$, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*$.type|raw$))
	})
	return ret, err
}
`

var typeLister_NamespaceLister = `
// $.type|publicPlural$ returns an object that can list and get $.type|publicPlural$.
func (s *$.type|private$Lister) $.type|publicPlural$(namespace string) $.type|public$NamespaceLister {
	return $.type|private$NamespaceLister{indexer: s.indexer, namespace: namespace}
}
`

var typeLister_NonNamespacedGet = `
// Get retrieves the $.type|public$ from the index for a given name.
func (s *$.type|private$Lister) Get(name string) (*$.type|raw$, error) {
  key := &$.type|raw${ObjectMeta: $.objectMeta|raw${Name: name}}
  obj, exists, err := s.indexer.Get(key)
  if err != nil {
    return nil, err
  }
  if !exists {
    return nil, errors.NewNotFound($.Resource|raw$("$.type|lowercaseSingular$"), name)
  }
  return obj.(*$.type|raw$), nil
}
`

var namespaceListerInterface = `
// $.type|public$NamespaceLister helps list and get $.type|publicPlural$.
type $.type|public$NamespaceLister interface {
	// List lists all $.type|publicPlural$ in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*$.type|raw$, err error)
	// Get retrieves the $.type|public$ from the indexer for a given namespace and name.
	Get(name string) (*$.type|raw$, error)
	$.type|public$NamespaceListerExpansion
}
`

var namespaceListerStruct = `
// $.type|private$NamespaceLister implements the $.type|public$NamespaceLister
// interface.
type $.type|private$NamespaceLister struct {
	indexer cache.Indexer
	namespace string
}
`

var namespaceLister_List = `
// List lists all $.type|publicPlural$ in the indexer for a given namespace.
func (s $.type|private$NamespaceLister) List(selector labels.Selector) (ret []*$.type|raw$, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*$.type|raw$))
	})
	return ret, err
}
`

var namespaceLister_Get = `
// Get retrieves the $.type|public$ from the indexer for a given namespace and name.
func (s $.type|private$NamespaceLister) Get(name string) (*$.type|raw$, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound($.Resource|raw$("$.type|lowercaseSingular$"), name)
	}
	return obj.(*$.type|raw$), nil
}
`
