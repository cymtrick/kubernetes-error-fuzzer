/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package unversioned

import "k8s.io/kubernetes/pkg/api"

type NamespaceExpansion interface {
	Finalize(item *api.Namespace) (*api.Namespace, error)
}

// Finalize takes the representation of a namespace to update.  Returns the server's representation of the namespace, and an error, if it occurs.
func (c *namespaces) Finalize(namespace *api.Namespace) (result *api.Namespace, err error) {
	result = &api.Namespace{}
	err = c.client.Put().Resource("namespaces").Name(namespace.Name).SubResource("finalize").Body(namespace).Do().Into(result)
	return
}
