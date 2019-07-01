/*
Copyright 2019 The Kubernetes Authors.

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

// Package runtimeclass contains an admission controller for modifying and validating new Pods to
// take RuntimeClass into account. For RuntimeClass definitions which describe an overhead associated
// with running a pod, this admission controller will set the pod.Spec.Overhead field accordingly. This
// field should only be set through this controller, so vaidation will be carried out to ensure the pod's
// value matches what is defined in the coresponding RuntimeClass.
package runtimeclass

import (
	"fmt"
	"io"

	v1beta1 "k8s.io/api/node/v1beta1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apiserver/pkg/admission"
	genericadmissioninitailizer "k8s.io/apiserver/pkg/admission/initializer"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/informers"
	nodev1beta1listers "k8s.io/client-go/listers/node/v1beta1"
	api "k8s.io/kubernetes/pkg/apis/core"
	node "k8s.io/kubernetes/pkg/apis/node"
	nodev1beta1 "k8s.io/kubernetes/pkg/apis/node/v1beta1"
	"k8s.io/kubernetes/pkg/features"
)

// PluginName indicates name of admission plugin.
const PluginName = "RuntimeClass"

// Register registers a plugin
func Register(plugins *admission.Plugins) {
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return NewRuntimeClass(), nil
	})
}

// RuntimeClass is an implementation of admission.Interface.
// It looks at all new pods and sets pod.Spec.Overhead if a RuntimeClass is specified which
// defines an Overhead. If pod.Spec.Overhead is set but a RuntimeClass with matching overhead is
// not specified, the pod is rejected.
type RuntimeClass struct {
	*admission.Handler
	runtimeClassLister nodev1beta1listers.RuntimeClassLister
}

var _ admission.MutationInterface = &RuntimeClass{}
var _ admission.ValidationInterface = &RuntimeClass{}

var _ genericadmissioninitailizer.WantsExternalKubeInformerFactory = &RuntimeClass{}

// SetExternalKubeInformerFactory implements the WantsExternalKubeInformerFactory interface.
func (r *RuntimeClass) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	runtimeClassInformer := f.Node().V1beta1().RuntimeClasses()
	r.SetReadyFunc(runtimeClassInformer.Informer().HasSynced)
	r.runtimeClassLister = runtimeClassInformer.Lister()
}

// ValidateInitialization implements the WantsExternalKubeInformerFactory interface.
func (r *RuntimeClass) ValidateInitialization() error {
	if r.runtimeClassLister == nil {
		return fmt.Errorf("missing RuntimeClass lister")
	}
	return nil
}

// Admit makes an admission decision based on the request attributes
func (r *RuntimeClass) Admit(attributes admission.Attributes, o admission.ObjectInterfaces) error {

	// Ignore all calls to subresources or resources other than pods.
	if shouldIgnore(attributes) {
		return nil
	}

	pod, runtimeClass, err := r.prepareObjects(attributes)
	if err != nil {
		return err
	}
	if utilfeature.DefaultFeatureGate.Enabled(features.PodOverhead) {
		err = setOverhead(attributes, pod, runtimeClass)
		if err != nil {
			return err
		}
	}

	return nil
}

// Validate makes sure that pod adhere's to RuntimeClass's definition
func (r *RuntimeClass) Validate(attributes admission.Attributes, o admission.ObjectInterfaces) error {

	// Ignore all calls to subresources or resources other than pods.
	if shouldIgnore(attributes) {
		return nil
	}

	pod, runtimeClass, err := r.prepareObjects(attributes)
	if err != nil {
		return err
	}

	if utilfeature.DefaultFeatureGate.Enabled(features.PodOverhead) {
		err = validateOverhead(attributes, pod, runtimeClass)
	}

	return nil
}

// NewRuntimeClass creates a new RuntimeClass admission control handler
func NewRuntimeClass() *RuntimeClass {
	return &RuntimeClass{
		Handler: admission.NewHandler(admission.Create),
	}
}

// prepareObjects returns pod and runtimeClass types from the given admission attributes
func (r *RuntimeClass) prepareObjects(attributes admission.Attributes) (pod *api.Pod, runtimeClass *v1beta1.RuntimeClass, err error) {

	pod, ok := attributes.GetObject().(*api.Pod)
	if !ok {
		return nil, nil, apierrors.NewBadRequest("Resource was marked with kind Pod but was unable to be converted")
	}

	// get RuntimeClass object
	runtimeClass, err = r.getRuntimeClass(pod, pod.Spec.RuntimeClassName)
	if err != nil {
		return pod, nil, err
	}

	// return the pod and runtimeClass. If no RuntimeClass is specified in PodSpec, runtimeClass will be nil
	return pod, runtimeClass, nil
}

// getRuntimeClass will return a reference to the RuntimeClass object if it is found. If it cannot be found, or a RuntimeClassName
// is not provided in the pod spec, *node.RuntimeClass returned will be nil
func (r *RuntimeClass) getRuntimeClass(pod *api.Pod, runtimeClassName *string) (runtimeClass *v1beta1.RuntimeClass, err error) {

	runtimeClass = nil

	if runtimeClassName != nil {
		runtimeClass, err = r.runtimeClassLister.Get(*runtimeClassName)
	}

	return runtimeClass, err
}

func setOverhead(a admission.Attributes, pod *api.Pod, runtimeClass *v1beta1.RuntimeClass) (err error) {

	if runtimeClass == nil || runtimeClass.Overhead == nil {
		return nil
	}

	// convert to internal type and assign to pod's Overhead
	nodeOverhead := &node.Overhead{}
	err = nodev1beta1.Convert_v1beta1_Overhead_To_node_Overhead(runtimeClass.Overhead, nodeOverhead, nil)
	if err != nil {
		return err
	}

	// reject pod if Overhead is already set that differs from what is defined in RuntimeClass
	if pod.Spec.Overhead != nil && !apiequality.Semantic.DeepEqual(nodeOverhead.PodFixed, pod.Spec.Overhead) {
		return admission.NewForbidden(a, fmt.Errorf("pod rejected: Pod's Overhead doesn't match RuntimeClass's defined Overhead"))
	}

	pod.Spec.Overhead = nodeOverhead.PodFixed

	return nil
}

func validateOverhead(a admission.Attributes, pod *api.Pod, runtimeClass *v1beta1.RuntimeClass) (err error) {

	if runtimeClass != nil && runtimeClass.Overhead != nil {
		// If the Overhead set doesn't match what is provided in the RuntimeClass definition, reject the pod
		nodeOverhead := &node.Overhead{}
		err := nodev1beta1.Convert_v1beta1_Overhead_To_node_Overhead(runtimeClass.Overhead, nodeOverhead, nil)
		if err != nil {
			return err
		}
		if !apiequality.Semantic.DeepEqual(nodeOverhead.PodFixed, pod.Spec.Overhead) {
			return admission.NewForbidden(a, fmt.Errorf("pod rejected: Pod's Overhead doesn't match RuntimeClass's defined Overhead"))
		}
	} else {
		// If RuntimeClass with Overhead is not defined but an Overhead is set for pod, reject the pod
		if pod.Spec.Overhead != nil {
			return admission.NewForbidden(a, fmt.Errorf("pod rejected: Pod Overhead set without corresponding RuntimeClass defined Overhead"))
		}
	}

	return nil
}

func shouldIgnore(attributes admission.Attributes) bool {
	// Ignore all calls to subresources or resources other than pods.
	if len(attributes.GetSubresource()) != 0 || attributes.GetResource().GroupResource() != api.Resource("pods") {
		return true
	}

	return false
}
