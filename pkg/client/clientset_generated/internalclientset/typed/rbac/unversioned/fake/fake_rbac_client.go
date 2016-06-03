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

package fake

import (
	unversioned "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/rbac/unversioned"
	restclient "k8s.io/kubernetes/pkg/client/restclient"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeRbac struct {
	*core.Fake
}

func (c *FakeRbac) ClusterRoles() unversioned.ClusterRoleInterface {
	return &FakeClusterRoles{c}
}

func (c *FakeRbac) ClusterRoleBindings() unversioned.ClusterRoleBindingInterface {
	return &FakeClusterRoleBindings{c}
}

func (c *FakeRbac) Roles(namespace string) unversioned.RoleInterface {
	return &FakeRoles{c, namespace}
}

func (c *FakeRbac) RoleBindings(namespace string) unversioned.RoleBindingInterface {
	return &FakeRoleBindings{c, namespace}
}

// GetRESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeRbac) GetRESTClient() *restclient.RESTClient {
	return nil
}
