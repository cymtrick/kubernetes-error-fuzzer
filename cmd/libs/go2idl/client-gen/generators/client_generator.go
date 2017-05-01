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

// Package generators has the generators for the client-gen utility.
package generators

import (
	"fmt"
	"path/filepath"
	"strings"

	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
	clientgenargs "k8s.io/kubernetes/cmd/libs/go2idl/client-gen/args"
	"k8s.io/kubernetes/cmd/libs/go2idl/client-gen/generators/fake"
	"k8s.io/kubernetes/cmd/libs/go2idl/client-gen/generators/scheme"
	clientgentypes "k8s.io/kubernetes/cmd/libs/go2idl/client-gen/types"

	"github.com/golang/glog"
)

// NameSystems returns the name system used by the generators in this package.
func NameSystems() namer.NameSystems {
	pluralExceptions := map[string]string{
		"Endpoints": "Endpoints",
	}
	lowercaseNamer := namer.NewAllLowercasePluralNamer(pluralExceptions)
	return namer.NameSystems{
		"public":             namer.NewPublicNamer(0),
		"private":            namer.NewPrivateNamer(0),
		"raw":                namer.NewRawNamer("", nil),
		"publicPlural":       namer.NewPublicPluralNamer(pluralExceptions),
		"privatePlural":      namer.NewPrivatePluralNamer(pluralExceptions),
		"allLowercasePlural": lowercaseNamer,
		"resource":           NewTagOverrideNamer("resourceName", lowercaseNamer),
	}
}

// DefaultNameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func DefaultNameSystem() string {
	return "public"
}

func generatedBy(customArgs clientgenargs.Args) string {
	if len(customArgs.CmdArgs) != 0 {
		return fmt.Sprintf("\n// This package is generated by client-gen with custom arguments.\n\n")
	}
	return fmt.Sprintf("\n// This package is generated by client-gen with the default arguments.\n\n")
}

func packageForGroup(gv clientgentypes.GroupVersion, typeList []*types.Type, clientsetPackage string, apiPath string, srcTreePath string, inputPackage string, boilerplate []byte, generatedBy string) generator.Package {
	groupVersionClientPackage := strings.ToLower(filepath.Join(clientsetPackage, "typed", gv.Group.NonEmpty(), gv.Version.NonEmpty()))
	return &generator.DefaultPackage{
		PackageName: strings.ToLower(gv.Version.NonEmpty()),
		PackagePath: groupVersionClientPackage,
		HeaderText:  boilerplate,
		PackageDocumentation: []byte(
			generatedBy +
				`// This package has the automatically generated typed clients.
`),
		// GeneratorFunc returns a list of generators. Each generator makes a
		// single file.
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			generators = []generator.Generator{
				// Always generate a "doc.go" file.
				generator.DefaultGen{OptionalName: "doc"},
			}
			// Since we want a file per type that we generate a client for, we
			// have to provide a function for this.
			for _, t := range typeList {
				generators = append(generators, &genClientForType{
					DefaultGen: generator.DefaultGen{
						OptionalName: strings.ToLower(c.Namers["private"].Name(t)),
					},
					outputPackage:    groupVersionClientPackage,
					clientsetPackage: clientsetPackage,
					group:            gv.Group.NonEmpty(),
					version:          gv.Version.String(),
					typeToMatch:      t,
					imports:          generator.NewImportTracker(),
				})
			}

			generators = append(generators, &genGroup{
				DefaultGen: generator.DefaultGen{
					OptionalName: gv.Group.NonEmpty() + "_client",
				},
				outputPackage:    groupVersionClientPackage,
				inputPackage:     inputPackage,
				clientsetPackage: clientsetPackage,
				group:            gv.Group.NonEmpty(),
				version:          gv.Version.String(),
				apiPath:          apiPath,
				types:            typeList,
				imports:          generator.NewImportTracker(),
			})

			expansionFileName := "generated_expansion"
			generators = append(generators, &genExpansion{
				groupPackagePath: filepath.Join(srcTreePath, groupVersionClientPackage),
				DefaultGen: generator.DefaultGen{
					OptionalName: expansionFileName,
				},
				types: typeList,
			})

			return generators
		},
		FilterFunc: func(c *generator.Context, t *types.Type) bool {
			return extractBoolTagOrDie("genclient", t.SecondClosestCommentLines) == true
		},
	}
}

func packageForClientset(customArgs clientgenargs.Args, clientsetPackage string, boilerplate []byte, generatedBy string) generator.Package {
	return &generator.DefaultPackage{
		PackageName: customArgs.ClientsetName,
		PackagePath: clientsetPackage,
		HeaderText:  boilerplate,
		PackageDocumentation: []byte(
			generatedBy +
				`// This package has the automatically generated clientset.
`),
		// GeneratorFunc returns a list of generators. Each generator generates a
		// single file.
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			generators = []generator.Generator{
				// Always generate a "doc.go" file.
				generator.DefaultGen{OptionalName: "doc"},

				&genClientset{
					DefaultGen: generator.DefaultGen{
						OptionalName: "clientset",
					},
					groups:           customArgs.Groups,
					clientsetPackage: clientsetPackage,
					outputPackage:    customArgs.ClientsetName,
					imports:          generator.NewImportTracker(),
				},
			}
			return generators
		},
	}
}

func packageForScheme(customArgs clientgenargs.Args, clientsetPackage string, srcTreePath string, boilerplate []byte, generatedBy string) generator.Package {
	schemePackage := filepath.Join(clientsetPackage, "scheme")

	// create runtime.Registry for internal client because it has to know about group versions
	internalClient := false
NextGroup:
	for _, group := range customArgs.Groups {
		for _, v := range group.Versions {
			if v == "" {
				internalClient = true
				break NextGroup
			}
		}
	}

	return &generator.DefaultPackage{
		PackageName: "scheme",
		PackagePath: schemePackage,
		HeaderText:  boilerplate,
		PackageDocumentation: []byte(
			generatedBy +
				`// This package contains the scheme of the automatically generated clientset.
`),
		// GeneratorFunc returns a list of generators. Each generator generates a
		// single file.
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			generators = []generator.Generator{
				// Always generate a "doc.go" file.
				generator.DefaultGen{OptionalName: "doc"},

				&scheme.GenScheme{
					DefaultGen: generator.DefaultGen{
						OptionalName: "register",
					},
					InputPackages:  customArgs.GroupVersionToInputPath,
					OutputPackage:  schemePackage,
					OutputPath:     filepath.Join(srcTreePath, schemePackage),
					Groups:         customArgs.Groups,
					ImportTracker:  generator.NewImportTracker(),
					CreateRegistry: internalClient,
				},
			}
			return generators
		},
	}
}

// Packages makes the client package definition.
func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	boilerplate, err := arguments.LoadGoBoilerplate()
	if err != nil {
		glog.Fatalf("Failed loading boilerplate: %v", err)
	}

	customArgs, ok := arguments.CustomArgs.(clientgenargs.Args)
	if !ok {
		glog.Fatalf("cannot convert arguments.CustomArgs to clientgenargs.Args")
	}
	includedTypesOverrides := customArgs.IncludedTypesOverrides

	generatedBy := generatedBy(customArgs)

	gvToTypes := map[clientgentypes.GroupVersion][]*types.Type{}
	for gv, inputDir := range customArgs.GroupVersionToInputPath {
		p := context.Universe.Package(inputDir)
		for n, t := range p.Types {
			// filter out types which are not included in user specified overrides.
			typesOverride, ok := includedTypesOverrides[gv]
			if ok {
				found := false
				for _, typeStr := range typesOverride {
					if typeStr == n {
						found = true
						break
					}
				}
				if !found {
					continue
				}
			} else {
				// User has not specified any override for this group version.
				// filter out types which dont have genclient=true.
				if extractBoolTagOrDie("genclient", t.SecondClosestCommentLines) == false {
					continue
				}
			}
			if _, found := gvToTypes[gv]; !found {
				gvToTypes[gv] = []*types.Type{}
			}
			gvToTypes[gv] = append(gvToTypes[gv], t)
		}
	}

	var packageList []generator.Package
	clientsetPackage := filepath.Join(customArgs.ClientsetOutputPath, customArgs.ClientsetName)

	packageList = append(packageList, packageForClientset(customArgs, clientsetPackage, boilerplate, generatedBy))
	packageList = append(packageList, packageForScheme(customArgs, clientsetPackage, arguments.OutputBase, boilerplate, generatedBy))
	if customArgs.FakeClient {
		packageList = append(packageList, fake.PackageForClientset(customArgs, clientsetPackage, boilerplate, generatedBy))
	}

	// If --clientset-only=true, we don't regenerate the individual typed clients.
	if customArgs.ClientsetOnly {
		return generator.Packages(packageList)
	}

	orderer := namer.Orderer{Namer: namer.NewPrivateNamer(0)}
	for _, group := range customArgs.Groups {
		for _, version := range group.Versions {
			gv := clientgentypes.GroupVersion{Group: group.Group, Version: version}
			types := gvToTypes[gv]
			inputPath := customArgs.GroupVersionToInputPath[gv]
			packageList = append(packageList, packageForGroup(gv, orderer.OrderTypes(types), clientsetPackage, customArgs.ClientsetAPIPath, arguments.OutputBase, inputPath, boilerplate, generatedBy))
			if customArgs.FakeClient {
				packageList = append(packageList, fake.PackageForGroup(gv, orderer.OrderTypes(types), clientsetPackage, inputPath, boilerplate, generatedBy))
			}
		}
	}

	return generator.Packages(packageList)
}

// tagOverrideNamer is a namer which pulls names from a given tag, if specified,
// and otherwise falls back to a different namer.
type tagOverrideNamer struct {
	tagName  string
	fallback namer.Namer
}

func (n *tagOverrideNamer) Name(t *types.Type) string {
	if nameOverride := extractTag(n.tagName, t.SecondClosestCommentLines); nameOverride != "" {
		return nameOverride
	}

	return n.fallback.Name(t)
}

// NewTagOverrideNamer creates a namer.Namer which uses the contents of the given tag as
// the name, or falls back to another Namer if the tag is not present.
func NewTagOverrideNamer(tagName string, fallback namer.Namer) namer.Namer {
	return &tagOverrideNamer{
		tagName:  tagName,
		fallback: fallback,
	}
}
