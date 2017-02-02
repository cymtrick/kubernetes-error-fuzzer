/*
Copyright 2015 The Kubernetes Authors.

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

package deployment

import (
	"fmt"
	"reflect"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	apistorage "k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/apis/extensions/validation"
	"k8s.io/kubernetes/pkg/controller/deployment/util"
)

// deploymentStrategy implements behavior for Deployments.
type deploymentStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

// Strategy is the default logic that applies when creating and updating Deployment
// objects via the REST API.
var Strategy = deploymentStrategy{api.Scheme, names.SimpleNameGenerator}

// DefaultGarbageCollectionPolicy returns Orphan because that's the default
// behavior before the server-side garbage collection is implemented.
func (deploymentStrategy) DefaultGarbageCollectionPolicy() rest.GarbageCollectionPolicy {
	return rest.OrphanDependents
}

// NamespaceScoped is true for deployment.
func (deploymentStrategy) NamespaceScoped() bool {
	return true
}

// PrepareForCreate clears fields that are not allowed to be set by end users on creation.
func (deploymentStrategy) PrepareForCreate(ctx genericapirequest.Context, obj runtime.Object) {
	deployment := obj.(*extensions.Deployment)
	deployment.Status = extensions.DeploymentStatus{}
	deployment.Generation = 1
}

// Validate validates a new deployment.
func (deploymentStrategy) Validate(ctx genericapirequest.Context, obj runtime.Object) field.ErrorList {
	deployment := obj.(*extensions.Deployment)
	return validation.ValidateDeployment(deployment)
}

// Canonicalize normalizes the object after validation.
func (deploymentStrategy) Canonicalize(obj runtime.Object) {
}

// AllowCreateOnUpdate is false for deployments.
func (deploymentStrategy) AllowCreateOnUpdate() bool {
	return false
}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update.
func (deploymentStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
	newDeployment := obj.(*extensions.Deployment)
	oldDeployment := old.(*extensions.Deployment)
	newDeployment.Status = oldDeployment.Status

	// Spec updates bump the generation so that we can distinguish between
	// scaling events and template changes, annotation updates bump the generation
	// because annotations are copied from deployments to their replica sets.
	if !reflect.DeepEqual(newDeployment.Spec, oldDeployment.Spec) ||
		!reflect.DeepEqual(newDeployment.Annotations, oldDeployment.Annotations) {
		newDeployment.Generation = oldDeployment.Generation + 1
	}

	// Records timestamp on selector updates in annotation
	if !reflect.DeepEqual(newDeployment.Spec.Selector, oldDeployment.Spec.Selector) {
		if newDeployment.Annotations == nil {
			newDeployment.Annotations = make(map[string]string)
		}
		now := metav1.Now()
		newDeployment.Annotations[util.SelectorUpdateAnnotation] = now.Format(time.RFC3339)
	}
}

// ValidateUpdate is the default update validation for an end user.
func (deploymentStrategy) ValidateUpdate(ctx genericapirequest.Context, obj, old runtime.Object) field.ErrorList {
	return validation.ValidateDeploymentUpdate(obj.(*extensions.Deployment), old.(*extensions.Deployment))
}

func (deploymentStrategy) AllowUnconditionalUpdate() bool {
	return true
}

type deploymentStatusStrategy struct {
	deploymentStrategy
}

var StatusStrategy = deploymentStatusStrategy{Strategy}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update of status
func (deploymentStatusStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
	newDeployment := obj.(*extensions.Deployment)
	oldDeployment := old.(*extensions.Deployment)
	newDeployment.Spec = oldDeployment.Spec
	newDeployment.Labels = oldDeployment.Labels
}

// ValidateUpdate is the default update validation for an end user updating status
func (deploymentStatusStrategy) ValidateUpdate(ctx genericapirequest.Context, obj, old runtime.Object) field.ErrorList {
	return validation.ValidateDeploymentStatusUpdate(obj.(*extensions.Deployment), old.(*extensions.Deployment))
}

// DeploymentToSelectableFields returns a field set that represents the object.
func DeploymentToSelectableFields(deployment *extensions.Deployment) fields.Set {
	return generic.ObjectMetaFieldsSet(&deployment.ObjectMeta, true)
}

// GetAttrs returns labels and fields of a given object for filtering purposes.
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	deployment, ok := obj.(*extensions.Deployment)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a deployment.")
	}
	return labels.Set(deployment.ObjectMeta.Labels), DeploymentToSelectableFields(deployment), nil
}

// MatchDeployment is the filter used by the generic etcd backend to route
// watch events from etcd to clients of the apiserver only interested in specific
// labels/fields.
func MatchDeployment(label labels.Selector, field fields.Selector) apistorage.SelectionPredicate {
	return apistorage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}
