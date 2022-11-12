/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "k8s.io/client-go/kubernetes/typed/resource/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeResourceV1alpha1 struct {
	*testing.Fake
}

func (c *FakeResourceV1alpha1) PodSchedulings(namespace string) v1alpha1.PodSchedulingInterface {
	return &FakePodSchedulings{c, namespace}
}

func (c *FakeResourceV1alpha1) ResourceClaims(namespace string) v1alpha1.ResourceClaimInterface {
	return &FakeResourceClaims{c, namespace}
}

func (c *FakeResourceV1alpha1) ResourceClaimTemplates(namespace string) v1alpha1.ResourceClaimTemplateInterface {
	return &FakeResourceClaimTemplates{c, namespace}
}

func (c *FakeResourceV1alpha1) ResourceClasses() v1alpha1.ResourceClassInterface {
	return &FakeResourceClasses{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeResourceV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
