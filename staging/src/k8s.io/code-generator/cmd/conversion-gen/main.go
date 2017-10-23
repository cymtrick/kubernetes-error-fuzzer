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

// conversion-gen is a tool for auto-generating Conversion functions.
//
// Given a list of input directories, it will scan for "peer" packages and
// generate functions that efficiently convert between same-name types in each
// package.  For any pair of types that has a
//     `Convert_<pkg1>_<type>_To_<pkg2>_<Type()`
// function (and its reciprocal), it will simply call that.  use standard value
// assignment whenever possible.  The resulting file will be stored in the same
// directory as the processed source package.
//
// Generation is governed by comment tags in the source.  Any package may
// request Conversion generation by including a comment in the file-comments of
// one file, of the form:
//   // +k8s:conversion-gen=<import-path-of-peer-package>
//
// When generating for a package, individual types or fields of structs may opt
// out of Conversion generation by specifying a comment on the of the form:
//   // +k8s:conversion-gen=false
package main

import (
	"path/filepath"

	"k8s.io/code-generator/cmd/conversion-gen/generators"
	"k8s.io/gengo/args"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

func main() {
	arguments := args.Default()

	// Custom args.
	customArgs := &generators.CustomArgs{
		BasePeerDirs: generators.DefaultBasePeerDirs,
		SkipUnsafe:   false,
	}
	pflag.CommandLine.StringSliceVar(&customArgs.BasePeerDirs, "base-peer-dirs", customArgs.BasePeerDirs,
		"Comma-separated list of apimachinery import paths which are considered, after tag-specified peers, for conversions. Only change these if you have very good reasons.")
	pflag.CommandLine.StringSliceVar(&customArgs.ExtraPeerDirs, "extra-peer-dirs", customArgs.ExtraPeerDirs,
		"Application specific comma-separated list of import paths which are considered, after tag-specified peers and base-peer-dirs, for conversions.")
	pflag.CommandLine.BoolVar(&customArgs.SkipUnsafe, "skip-unsafe", customArgs.SkipUnsafe,
		"If true, will not generate code using unsafe pointer conversions; resulting code may be slower.")

	// Override defaults.
	arguments.GoHeaderFilePath = filepath.Join(args.DefaultSourceTree(), "k8s.io/kubernetes/hack/boilerplate/boilerplate.go.txt")
	arguments.OutputFileBaseName = "conversion_generated"
	arguments.CustomArgs = customArgs

	// Run it.
	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		glog.Fatalf("Error: %v", err)
	}
	glog.V(2).Info("Completed successfully.")
}
