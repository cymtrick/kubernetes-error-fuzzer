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
	"context"
	json "encoding/json"
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsbatchv1 "k8s.io/client-go/applyconfigurations/batch/v1"
	testing "k8s.io/client-go/testing"
)

// FakeCronJobs implements CronJobInterface
type FakeCronJobs struct {
	Fake *FakeBatchV1
	ns   string
}

var cronjobsResource = schema.GroupVersionResource{Group: "batch", Version: "v1", Resource: "cronjobs"}

var cronjobsKind = schema.GroupVersionKind{Group: "batch", Version: "v1", Kind: "CronJob"}

// Get takes name of the cronJob, and returns the corresponding cronJob object, and an error if there is any.
func (c *FakeCronJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *batchv1.CronJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cronjobsResource, c.ns, name), &batchv1.CronJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1.CronJob), err
}

// List takes label and field selectors, and returns the list of CronJobs that match those selectors.
func (c *FakeCronJobs) List(ctx context.Context, opts v1.ListOptions) (result *batchv1.CronJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cronjobsResource, cronjobsKind, c.ns, opts), &batchv1.CronJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &batchv1.CronJobList{ListMeta: obj.(*batchv1.CronJobList).ListMeta}
	for _, item := range obj.(*batchv1.CronJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cronJobs.
func (c *FakeCronJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cronjobsResource, c.ns, opts))

}

// Create takes the representation of a cronJob and creates it.  Returns the server's representation of the cronJob, and an error, if there is any.
func (c *FakeCronJobs) Create(ctx context.Context, cronJob *batchv1.CronJob, opts v1.CreateOptions) (result *batchv1.CronJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cronjobsResource, c.ns, cronJob), &batchv1.CronJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1.CronJob), err
}

// Update takes the representation of a cronJob and updates it. Returns the server's representation of the cronJob, and an error, if there is any.
func (c *FakeCronJobs) Update(ctx context.Context, cronJob *batchv1.CronJob, opts v1.UpdateOptions) (result *batchv1.CronJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cronjobsResource, c.ns, cronJob), &batchv1.CronJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1.CronJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCronJobs) UpdateStatus(ctx context.Context, cronJob *batchv1.CronJob, opts v1.UpdateOptions) (*batchv1.CronJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cronjobsResource, "status", c.ns, cronJob), &batchv1.CronJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1.CronJob), err
}

// Delete takes name of the cronJob and deletes it. Returns an error if one occurs.
func (c *FakeCronJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cronjobsResource, c.ns, name, opts), &batchv1.CronJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCronJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cronjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &batchv1.CronJobList{})
	return err
}

// Patch applies the patch and returns the patched cronJob.
func (c *FakeCronJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *batchv1.CronJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cronjobsResource, c.ns, name, pt, data, subresources...), &batchv1.CronJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1.CronJob), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied cronJob.
func (c *FakeCronJobs) Apply(ctx context.Context, cronJob *applyconfigurationsbatchv1.CronJobApplyConfiguration, opts v1.ApplyOptions) (result *batchv1.CronJob, err error) {
	if cronJob == nil {
		return nil, fmt.Errorf("cronJob provided to Apply must not be nil")
	}
	data, err := json.Marshal(cronJob)
	if err != nil {
		return nil, err
	}
	name := cronJob.Name
	if name == nil {
		return nil, fmt.Errorf("cronJob.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cronjobsResource, c.ns, *name, types.ApplyPatchType, data), &batchv1.CronJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1.CronJob), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeCronJobs) ApplyStatus(ctx context.Context, cronJob *applyconfigurationsbatchv1.CronJobApplyConfiguration, opts v1.ApplyOptions) (result *batchv1.CronJob, err error) {
	if cronJob == nil {
		return nil, fmt.Errorf("cronJob provided to Apply must not be nil")
	}
	data, err := json.Marshal(cronJob)
	if err != nil {
		return nil, err
	}
	name := cronJob.Name
	if name == nil {
		return nil, fmt.Errorf("cronJob.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cronjobsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &batchv1.CronJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1.CronJob), err
}
