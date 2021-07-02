/*
Copyright 2017 The Kubernetes Authors.

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
	"time"

	fuzz "github.com/google/gofuzz"

	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/util/certificate/csr"
	"k8s.io/kubernetes/pkg/apis/certificates"
	api "k8s.io/kubernetes/pkg/apis/core"
)

// Funcs returns the fuzzer functions for the certificates api group.
var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		func(obj *certificates.CertificateSigningRequestSpec, c fuzz.Continue) {
			c.FuzzNoCustom(obj) // fuzz self without calling this function again
			obj.Usages = []certificates.KeyUsage{certificates.UsageKeyEncipherment}
			obj.SignerName = "example.com/custom-sample-signer"
			obj.ExpirationSeconds = csr.DurationToExpirationSeconds(time.Hour + time.Minute + time.Second)
		},
		func(obj *certificates.CertificateSigningRequestCondition, c fuzz.Continue) {
			c.FuzzNoCustom(obj) // fuzz self without calling this function again
			if len(obj.Status) == 0 {
				obj.Status = api.ConditionTrue
			}
		},
	}
}
