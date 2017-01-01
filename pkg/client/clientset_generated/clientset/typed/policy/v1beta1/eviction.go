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

package v1beta1

import (
	restclient "k8s.io/kubernetes/pkg/client/restclient"
)

// EvictionsGetter has a method to return a EvictionInterface.
// A group's client should implement this interface.
type EvictionsGetter interface {
	Evictions(namespace string) EvictionInterface
}

// EvictionInterface has methods to work with Eviction resources.
type EvictionInterface interface {
	EvictionExpansion
}

// evictions implements EvictionInterface
type evictions struct {
	client restclient.Interface
	ns     string
}

// newEvictions returns a Evictions
func newEvictions(c *PolicyV1beta1Client, namespace string) *evictions {
	return &evictions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}
