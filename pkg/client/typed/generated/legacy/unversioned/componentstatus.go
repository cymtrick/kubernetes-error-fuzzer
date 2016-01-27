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

import (
	api "k8s.io/kubernetes/pkg/api"
	watch "k8s.io/kubernetes/pkg/watch"
)

// ComponentStatusGetter has a method to return a ComponentStatusInterface.
// A group's client should implement this interface.
type ComponentStatusGetter interface {
	ComponentStatus() ComponentStatusInterface
}

// ComponentStatusInterface has methods to work with ComponentStatus resources.
type ComponentStatusInterface interface {
	Create(*api.ComponentStatus) (*api.ComponentStatus, error)
	Update(*api.ComponentStatus) (*api.ComponentStatus, error)
	Delete(name string, options *api.DeleteOptions) error
	DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error
	Get(name string) (*api.ComponentStatus, error)
	List(opts api.ListOptions) (*api.ComponentStatusList, error)
	Watch(opts api.ListOptions) (watch.Interface, error)
	ComponentStatusExpansion
}

// componentStatus implements ComponentStatusInterface
type componentStatus struct {
	client *LegacyClient
}

// newComponentStatus returns a ComponentStatus
func newComponentStatus(c *LegacyClient) *componentStatus {
	return &componentStatus{
		client: c,
	}
}

// Create takes the representation of a componentStatus and creates it.  Returns the server's representation of the componentStatus, and an error, if there is any.
func (c *componentStatus) Create(componentStatus *api.ComponentStatus) (result *api.ComponentStatus, err error) {
	result = &api.ComponentStatus{}
	err = c.client.Post().
		Resource("componentstatus").
		Body(componentStatus).
		Do().
		Into(result)
	return
}

// Update takes the representation of a componentStatus and updates it. Returns the server's representation of the componentStatus, and an error, if there is any.
func (c *componentStatus) Update(componentStatus *api.ComponentStatus) (result *api.ComponentStatus, err error) {
	result = &api.ComponentStatus{}
	err = c.client.Put().
		Resource("componentstatus").
		Name(componentStatus.Name).
		Body(componentStatus).
		Do().
		Into(result)
	return
}

// Delete takes name of the componentStatus and deletes it. Returns an error if one occurs.
func (c *componentStatus) Delete(name string, options *api.DeleteOptions) error {
	return c.client.Delete().
		Resource("componentstatus").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *componentStatus) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	return c.client.Delete().
		Resource("componentstatus").
		VersionedParams(&listOptions, api.Scheme).
		Body(options).
		Do().
		Error()
}

// Get takes name of the componentStatus, and returns the corresponding componentStatus object, and an error if there is any.
func (c *componentStatus) Get(name string) (result *api.ComponentStatus, err error) {
	result = &api.ComponentStatus{}
	err = c.client.Get().
		Resource("componentstatus").
		Name(name).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComponentStatus that match those selectors.
func (c *componentStatus) List(opts api.ListOptions) (result *api.ComponentStatusList, err error) {
	result = &api.ComponentStatusList{}
	err = c.client.Get().
		Resource("componentstatus").
		VersionedParams(&opts, api.Scheme).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested componentStatus.
func (c *componentStatus) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.client.Get().
		Prefix("watch").
		Resource("componentstatus").
		VersionedParams(&opts, api.Scheme).
		Watch()
}
