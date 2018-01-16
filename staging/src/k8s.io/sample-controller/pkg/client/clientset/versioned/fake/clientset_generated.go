/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
	clientset "k8s.io/sample-controller/pkg/client/clientset/versioned"
	samplecontrollerv1alpha1 "k8s.io/sample-controller/pkg/client/clientset/versioned/typed/samplecontroller/v1alpha1"
	fakesamplecontrollerv1alpha1 "k8s.io/sample-controller/pkg/client/clientset/versioned/typed/samplecontroller/v1alpha1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o))
	fakePtr.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return &Clientset{fakePtr, &fakediscovery.FakeDiscovery{Fake: &fakePtr}}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// SamplecontrollerV1alpha1 retrieves the SamplecontrollerV1alpha1Client
func (c *Clientset) SamplecontrollerV1alpha1() samplecontrollerv1alpha1.SamplecontrollerV1alpha1Interface {
	return &fakesamplecontrollerv1alpha1.FakeSamplecontrollerV1alpha1{Fake: &c.Fake}
}

// Samplecontroller retrieves the SamplecontrollerV1alpha1Client
func (c *Clientset) Samplecontroller() samplecontrollerv1alpha1.SamplecontrollerV1alpha1Interface {
	return &fakesamplecontrollerv1alpha1.FakeSamplecontrollerV1alpha1{Fake: &c.Fake}
}
