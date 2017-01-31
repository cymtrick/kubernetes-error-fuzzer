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

package example

import (
	"github.com/google/gofuzz"

	apitesting "k8s.io/apimachinery/pkg/api/testing"
	"k8s.io/apimachinery/pkg/runtime"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	examplev1 "k8s.io/apiserver/pkg/apis/example/v1"
)

// overrideGenericFuncs override some generic fuzzer funcs from k8s.io/apiserver in order to have more realistic
// values in a Kubernetes context.
func overrideGenericFuncs(t apitesting.TestingCommon, codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		func(j *runtime.Object, c fuzz.Continue) {
			// TODO: uncomment when round trip starts from a versioned object
			if true { //c.RandBool() {
				*j = &runtime.Unknown{
					// We do not set TypeMeta here because it is not carried through a round trip
					Raw:         []byte(`{"apiVersion":"unknown.group/unknown","kind":"Something","someKey":"someValue"}`),
					ContentType: runtime.ContentTypeJSON,
				}
			} else {
				types := []runtime.Object{&Pod{}}
				t := types[c.Rand.Intn(len(types))]
				c.Fuzz(t)
				*j = t
			}
		},
		func(r *runtime.RawExtension, c fuzz.Continue) {
			// Pick an arbitrary type and fuzz it
			types := []runtime.Object{&Pod{}}
			obj := types[c.Rand.Intn(len(types))]
			c.Fuzz(obj)

			// Convert the object to raw bytes
			bytes, err := runtime.Encode(apitesting.TestCodec(codecs, examplev1.SchemeGroupVersion), obj)
			if err != nil {
				t.Errorf("Failed to encode object: %v", err)
				return
			}

			// Set the bytes field on the RawExtension
			r.Raw = bytes
		},
	}
}

func exampleFuncs(t apitesting.TestingCommon) []interface{} {
	return []interface{}{
		func(s *PodSpec, c fuzz.Continue) {
			c.FuzzNoCustom(s)
			// has a default value
			ttl := int64(30)
			if c.RandBool() {
				ttl = int64(c.Uint32())
			}
			s.TerminationGracePeriodSeconds = &ttl

			if s.SchedulerName == "" {
				s.SchedulerName = "default-scheduler"
			}
		},
		func(j *PodPhase, c fuzz.Continue) {
			statuses := []PodPhase{"Pending", "Running", "Succeeded", "Failed", "Unknown"}
			*j = statuses[c.Rand.Intn(len(statuses))]
		},
		func(rp *RestartPolicy, c fuzz.Continue) {
			policies := []RestartPolicy{"Always", "Never", "OnFailure"}
			*rp = policies[c.Rand.Intn(len(policies))]
		},
	}
}

func FuzzerFuncs(t apitesting.TestingCommon, codecs runtimeserializer.CodecFactory) []interface{} {
	return apitesting.MergeFuzzerFuncs(t,
		apitesting.GenericFuzzerFuncs(t, codecs),
		overrideGenericFuncs(t, codecs),
		exampleFuncs(t),
	)
}
