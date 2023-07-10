/*
Copyright 2022 The Kubernetes Authors.

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

package fuzzer

import (
	fuzz "github.com/google/gofuzz"

	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kubernetes/pkg/apis/flowcontrol"
)

// Funcs returns the fuzzer functions for the flowcontrol api group.
var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		func(obj *flowcontrol.LimitedPriorityLevelConfiguration, c fuzz.Continue) {
			c.FuzzNoCustom(obj) // fuzz self without calling this function again
			if obj.LendablePercent == nil {
				i := int32(0)
				obj.LendablePercent = &i
			}
		},
		func(obj *flowcontrol.ExemptPriorityLevelConfiguration, c fuzz.Continue) {
			c.FuzzNoCustom(obj) // fuzz self without calling this function again
			if obj.NominalConcurrencyShares == nil {
				i := int32(0)
				obj.NominalConcurrencyShares = &i
			}
			if obj.LendablePercent == nil {
				i := int32(0)
				obj.LendablePercent = &i
			}
		},
	}
}
