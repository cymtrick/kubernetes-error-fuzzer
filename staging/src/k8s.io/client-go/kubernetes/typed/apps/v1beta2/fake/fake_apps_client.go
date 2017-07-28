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

package fake

import (
	v1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAppsV1beta2 struct {
	*testing.Fake
}

func (c *FakeAppsV1beta2) DaemonSets(namespace string) v1beta2.DaemonSetInterface {
	return &FakeDaemonSets{c, namespace}
}

func (c *FakeAppsV1beta2) Deployments(namespace string) v1beta2.DeploymentInterface {
	return &FakeDeployments{c, namespace}
}

func (c *FakeAppsV1beta2) ReplicaSets(namespace string) v1beta2.ReplicaSetInterface {
	return &FakeReplicaSets{c, namespace}
}

func (c *FakeAppsV1beta2) Scales(namespace string) v1beta2.ScaleInterface {
	return &FakeScales{c, namespace}
}

func (c *FakeAppsV1beta2) StatefulSets(namespace string) v1beta2.StatefulSetInterface {
	return &FakeStatefulSets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAppsV1beta2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
