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
	"path/filepath"
	"strings"

	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
	clientgentypes "k8s.io/kubernetes/cmd/libs/go2idl/client-gen/types"

	"github.com/golang/glog"
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
	return fmt.Sprintf("\n// This file was automatically generated by informer-gen\n\n")
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

func packageForGroup(base string, group clientgentypes.Group) string {
	return filepath.Join(base, group.NonEmpty())
}

func packageForInternalInterfaces(base string) string {
	return filepath.Join(base, "internalinterfaces")
}

// Packages makes the client package definition.
func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	boilerplate, err := arguments.LoadGoBoilerplate()
	if err != nil {
		glog.Fatalf("Failed loading boilerplate: %v", err)
	}

	boilerplate = append(boilerplate, []byte(generatedBy())...)

	customArgs, ok := arguments.CustomArgs.(*CustomArgs)
	if !ok {
		glog.Fatalf("Wrong CustomArgs type: %T", arguments.CustomArgs)
	}

	internalVersionPackagePath := filepath.Join(arguments.OutputPackagePath, "internalversion")
	externalVersionPackagePath := filepath.Join(arguments.OutputPackagePath, "externalversions")

	var packageList generator.Packages
	typesForGroupVersion := make(map[clientgentypes.GroupVersion][]*types.Type)

	externalGroupVersions := make(map[string]clientgentypes.GroupVersions)
	internalGroupVersions := make(map[string]clientgentypes.GroupVersions)
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
		var targetGroupVersions map[string]clientgentypes.GroupVersions

		if internal {
			lastSlash := strings.LastIndex(p.Path, "/")
			if lastSlash == -1 {
				glog.Fatalf("error constructing internal group version for package %q", p.Path)
			}
			gv.Group = clientgentypes.Group(p.Path[lastSlash+1:])
			targetGroupVersions = internalGroupVersions
		} else {
			parts := strings.Split(p.Path, "/")
			gv.Group = clientgentypes.Group(parts[len(parts)-2])
			gv.Version = clientgentypes.Version(parts[len(parts)-1])
			targetGroupVersions = externalGroupVersions
		}

		var typesToGenerate []*types.Type
		for _, t := range p.Types {
			// filter out types which dont have genclient=true.
			if extractBoolTagOrDie("genclient", t.SecondClosestCommentLines) == false {
				continue
			}
			// filter out types which have noMethods
			if extractBoolTagOrDie("noMethods", t.SecondClosestCommentLines) == true {
				continue
			}

			typesToGenerate = append(typesToGenerate, t)

			if _, ok := typesForGroupVersion[gv]; !ok {
				typesForGroupVersion[gv] = []*types.Type{}
			}
			typesForGroupVersion[gv] = append(typesForGroupVersion[gv], t)
		}
		if len(typesToGenerate) == 0 {
			continue
		}

		icGroupName := namer.IC(gv.Group.NonEmpty())
		groupVersionsEntry, ok := targetGroupVersions[icGroupName]
		if !ok {
			groupVersionsEntry = clientgentypes.GroupVersions{
				Group: gv.Group,
			}
		}
		groupVersionsEntry.Versions = append(groupVersionsEntry.Versions, gv.Version)
		targetGroupVersions[icGroupName] = groupVersionsEntry

		orderer := namer.Orderer{Namer: namer.NewPrivateNamer(0)}
		typesToGenerate = orderer.OrderTypes(typesToGenerate)

		if internal {
			packageList = append(packageList, versionPackage(internalVersionPackagePath, gv, boilerplate, typesToGenerate, customArgs.InternalClientSetPackage, customArgs.ListersPackage))
		} else {
			packageList = append(packageList, versionPackage(externalVersionPackagePath, gv, boilerplate, typesToGenerate, customArgs.VersionedClientSetPackage, customArgs.ListersPackage))
		}
	}

	packageList = append(packageList, factoryInterfacePackage(externalVersionPackagePath, boilerplate, customArgs.VersionedClientSetPackage, typesForGroupVersion))
	packageList = append(packageList, factoryPackage(externalVersionPackagePath, boilerplate, externalGroupVersions, customArgs.VersionedClientSetPackage, typesForGroupVersion))
	for _, groupVersionsEntry := range externalGroupVersions {
		packageList = append(packageList, groupPackage(externalVersionPackagePath, groupVersionsEntry, boilerplate))
	}

	packageList = append(packageList, factoryInterfacePackage(internalVersionPackagePath, boilerplate, customArgs.InternalClientSetPackage, typesForGroupVersion))
	packageList = append(packageList, factoryPackage(internalVersionPackagePath, boilerplate, internalGroupVersions, customArgs.InternalClientSetPackage, typesForGroupVersion))
	for _, groupVersionsEntry := range internalGroupVersions {
		packageList = append(packageList, groupPackage(internalVersionPackagePath, groupVersionsEntry, boilerplate))
	}

	return packageList
}

func isInternalVersion(gv clientgentypes.GroupVersion) bool {
	return len(gv.Version) == 0
}

func factoryPackage(basePackage string, boilerplate []byte, groupVersions map[string]clientgentypes.GroupVersions, clientSetPackage string, typesForGroupVersion map[clientgentypes.GroupVersion][]*types.Type) generator.Package {
	return &generator.DefaultPackage{
		PackageName: filepath.Base(basePackage),
		PackagePath: basePackage,
		HeaderText:  boilerplate,
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			generators = append(generators, &factoryGenerator{
				DefaultGen: generator.DefaultGen{
					OptionalName: "factory",
				},
				outputPackage:             basePackage,
				imports:                   generator.NewImportTracker(),
				groupVersions:             groupVersions,
				clientSetPackage:          clientSetPackage,
				internalInterfacesPackage: packageForInternalInterfaces(basePackage),
			})

			generators = append(generators, &genericGenerator{
				DefaultGen: generator.DefaultGen{
					OptionalName: "generic",
				},
				outputPackage:        basePackage,
				imports:              generator.NewImportTracker(),
				groupVersions:        groupVersions,
				typesForGroupVersion: typesForGroupVersion,
			})

			return generators
		},
	}
}

func factoryInterfacePackage(basePackage string, boilerplate []byte, clientSetPackage string, typesForGroupVersion map[clientgentypes.GroupVersion][]*types.Type) generator.Package {
	packagePath := packageForInternalInterfaces(basePackage)

	return &generator.DefaultPackage{
		PackageName: filepath.Base(packagePath),
		PackagePath: packagePath,
		HeaderText:  boilerplate,
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			generators = append(generators, &factoryInterfaceGenerator{
				DefaultGen: generator.DefaultGen{
					OptionalName: "factory_interfaces",
				},
				outputPackage:    packagePath,
				imports:          generator.NewImportTracker(),
				clientSetPackage: clientSetPackage,
			})

			return generators
		},
	}
}

func groupPackage(basePackage string, groupVersions clientgentypes.GroupVersions, boilerplate []byte) generator.Package {
	packagePath := filepath.Join(basePackage, strings.ToLower(groupVersions.Group.NonEmpty()))

	return &generator.DefaultPackage{
		PackageName: strings.ToLower(groupVersions.Group.NonEmpty()),
		PackagePath: packagePath,
		HeaderText:  boilerplate,
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			generators = append(generators, &groupInterfaceGenerator{
				DefaultGen: generator.DefaultGen{
					OptionalName: "interface",
				},
				outputPackage:             packagePath,
				groupVersions:             groupVersions,
				imports:                   generator.NewImportTracker(),
				internalInterfacesPackage: packageForInternalInterfaces(basePackage),
			})
			return generators
		},
		FilterFunc: func(c *generator.Context, t *types.Type) bool {
			// piggy-back on types that are tagged for client-gen
			return extractBoolTagOrDie("genclient", t.SecondClosestCommentLines) == true
		},
	}
}

func versionPackage(basePackage string, gv clientgentypes.GroupVersion, boilerplate []byte, typesToGenerate []*types.Type, clientSetPackage, listersPackage string) generator.Package {
	packagePath := filepath.Join(basePackage, strings.ToLower(gv.Group.NonEmpty()), strings.ToLower(gv.Version.NonEmpty()))

	return &generator.DefaultPackage{
		PackageName: strings.ToLower(gv.Version.NonEmpty()),
		PackagePath: packagePath,
		HeaderText:  boilerplate,
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			generators = append(generators, &versionInterfaceGenerator{
				DefaultGen: generator.DefaultGen{
					OptionalName: "interface",
				},
				outputPackage: packagePath,
				imports:       generator.NewImportTracker(),
				types:         typesToGenerate,
				internalInterfacesPackage: packageForInternalInterfaces(basePackage),
			})

			for _, t := range typesToGenerate {
				generators = append(generators, &informerGenerator{
					DefaultGen: generator.DefaultGen{
						OptionalName: strings.ToLower(t.Name.Name),
					},
					outputPackage:             packagePath,
					groupVersion:              gv,
					typeToGenerate:            t,
					imports:                   generator.NewImportTracker(),
					clientSetPackage:          clientSetPackage,
					listersPackage:            listersPackage,
					internalInterfacesPackage: packageForInternalInterfaces(basePackage),
				})
			}
			return generators
		},
		FilterFunc: func(c *generator.Context, t *types.Type) bool {
			// piggy-back on types that are tagged for client-gen
			return extractBoolTagOrDie("genclient", t.SecondClosestCommentLines) == true
		},
	}
}
