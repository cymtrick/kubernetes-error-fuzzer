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

package versioned

import (
	glog "github.com/golang/glog"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
	examplev1 "k8s.io/code-generator/_examples/apiserver/clientset/versioned/typed/example/v1"
	secondexamplev1 "k8s.io/code-generator/_examples/apiserver/clientset/versioned/typed/example2/v1"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ExampleV1() examplev1.ExampleV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Example() examplev1.ExampleV1Interface
	SecondExampleV1() secondexamplev1.SecondExampleV1Interface
	// Deprecated: please explicitly pick a version if possible.
	SecondExample() secondexamplev1.SecondExampleV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	exampleV1       *examplev1.ExampleV1Client
	secondExampleV1 *secondexamplev1.SecondExampleV1Client
}

// ExampleV1 retrieves the ExampleV1Client
func (c *Clientset) ExampleV1() examplev1.ExampleV1Interface {
	return c.exampleV1
}

// Deprecated: Example retrieves the default version of ExampleClient.
// Please explicitly pick a version.
func (c *Clientset) Example() examplev1.ExampleV1Interface {
	return c.exampleV1
}

// SecondExampleV1 retrieves the SecondExampleV1Client
func (c *Clientset) SecondExampleV1() secondexamplev1.SecondExampleV1Interface {
	return c.secondExampleV1
}

// Deprecated: SecondExample retrieves the default version of SecondExampleClient.
// Please explicitly pick a version.
func (c *Clientset) SecondExample() secondexamplev1.SecondExampleV1Interface {
	return c.secondExampleV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.exampleV1, err = examplev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.secondExampleV1, err = secondexamplev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		glog.Errorf("failed to create the DiscoveryClient: %v", err)
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.exampleV1 = examplev1.NewForConfigOrDie(c)
	cs.secondExampleV1 = secondexamplev1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.exampleV1 = examplev1.New(c)
	cs.secondExampleV1 = secondexamplev1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
