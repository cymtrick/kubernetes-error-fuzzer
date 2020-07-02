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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	clientset "k8s.io/client-go/kubernetes"
	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	fakeadmissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1/fake"
	admissionregistrationv1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	fakeadmissionregistrationv1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1/fake"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	fakeappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1/fake"
	appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	fakeappsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1/fake"
	appsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	fakeappsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2/fake"
	authenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	fakeauthenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1/fake"
	authenticationv1beta1 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	fakeauthenticationv1beta1 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1/fake"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	fakeauthorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1/fake"
	authorizationv1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	fakeauthorizationv1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1/fake"
	autoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	fakeautoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1/fake"
	autoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	fakeautoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1/fake"
	autoscalingv2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	fakeautoscalingv2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2/fake"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	fakebatchv1 "k8s.io/client-go/kubernetes/typed/batch/v1/fake"
	batchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	fakebatchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1/fake"
	batchv2alpha1 "k8s.io/client-go/kubernetes/typed/batch/v2alpha1"
	fakebatchv2alpha1 "k8s.io/client-go/kubernetes/typed/batch/v2alpha1/fake"
	certificatesv1 "k8s.io/client-go/kubernetes/typed/certificates/v1"
	fakecertificatesv1 "k8s.io/client-go/kubernetes/typed/certificates/v1/fake"
	certificatesv1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	fakecertificatesv1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1/fake"
	coordinationv1 "k8s.io/client-go/kubernetes/typed/coordination/v1"
	fakecoordinationv1 "k8s.io/client-go/kubernetes/typed/coordination/v1/fake"
	coordinationv1beta1 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	fakecoordinationv1beta1 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1/fake"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	fakecorev1 "k8s.io/client-go/kubernetes/typed/core/v1/fake"
	discoveryv1alpha1 "k8s.io/client-go/kubernetes/typed/discovery/v1alpha1"
	fakediscoveryv1alpha1 "k8s.io/client-go/kubernetes/typed/discovery/v1alpha1/fake"
	discoveryv1beta1 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
	fakediscoveryv1beta1 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1/fake"
	eventsv1 "k8s.io/client-go/kubernetes/typed/events/v1"
	fakeeventsv1 "k8s.io/client-go/kubernetes/typed/events/v1/fake"
	eventsv1beta1 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	fakeeventsv1beta1 "k8s.io/client-go/kubernetes/typed/events/v1beta1/fake"
	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	fakeextensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1/fake"
	flowcontrolv1alpha1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1"
	fakeflowcontrolv1alpha1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/fake"
	networkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	fakenetworkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1/fake"
	networkingv1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	fakenetworkingv1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1/fake"
	nodev1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	fakenodev1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1/fake"
	nodev1beta1 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	fakenodev1beta1 "k8s.io/client-go/kubernetes/typed/node/v1beta1/fake"
	policyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	fakepolicyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1/fake"
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	fakerbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1/fake"
	rbacv1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	fakerbacv1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1/fake"
	rbacv1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	fakerbacv1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1/fake"
	schedulingv1 "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	fakeschedulingv1 "k8s.io/client-go/kubernetes/typed/scheduling/v1/fake"
	schedulingv1alpha1 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	fakeschedulingv1alpha1 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1/fake"
	schedulingv1beta1 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	fakeschedulingv1beta1 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1/fake"
	settingsv1alpha1 "k8s.io/client-go/kubernetes/typed/settings/v1alpha1"
	fakesettingsv1alpha1 "k8s.io/client-go/kubernetes/typed/settings/v1alpha1/fake"
	storagev1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	fakestoragev1 "k8s.io/client-go/kubernetes/typed/storage/v1/fake"
	storagev1alpha1 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	fakestoragev1alpha1 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1/fake"
	storagev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	fakestoragev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1/fake"
	"k8s.io/client-go/testing"
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

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// AdmissionregistrationV1 retrieves the AdmissionregistrationV1Client
func (c *Clientset) AdmissionregistrationV1() admissionregistrationv1.AdmissionregistrationV1Interface {
	return &fakeadmissionregistrationv1.FakeAdmissionregistrationV1{Fake: &c.Fake}
}

// AdmissionregistrationV1beta1 retrieves the AdmissionregistrationV1beta1Client
func (c *Clientset) AdmissionregistrationV1beta1() admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface {
	return &fakeadmissionregistrationv1beta1.FakeAdmissionregistrationV1beta1{Fake: &c.Fake}
}

// AppsV1 retrieves the AppsV1Client
func (c *Clientset) AppsV1() appsv1.AppsV1Interface {
	return &fakeappsv1.FakeAppsV1{Fake: &c.Fake}
}

// AppsV1beta1 retrieves the AppsV1beta1Client
func (c *Clientset) AppsV1beta1() appsv1beta1.AppsV1beta1Interface {
	return &fakeappsv1beta1.FakeAppsV1beta1{Fake: &c.Fake}
}

// AppsV1beta2 retrieves the AppsV1beta2Client
func (c *Clientset) AppsV1beta2() appsv1beta2.AppsV1beta2Interface {
	return &fakeappsv1beta2.FakeAppsV1beta2{Fake: &c.Fake}
}

// AuthenticationV1 retrieves the AuthenticationV1Client
func (c *Clientset) AuthenticationV1() authenticationv1.AuthenticationV1Interface {
	return &fakeauthenticationv1.FakeAuthenticationV1{Fake: &c.Fake}
}

// AuthenticationV1beta1 retrieves the AuthenticationV1beta1Client
func (c *Clientset) AuthenticationV1beta1() authenticationv1beta1.AuthenticationV1beta1Interface {
	return &fakeauthenticationv1beta1.FakeAuthenticationV1beta1{Fake: &c.Fake}
}

// AuthorizationV1 retrieves the AuthorizationV1Client
func (c *Clientset) AuthorizationV1() authorizationv1.AuthorizationV1Interface {
	return &fakeauthorizationv1.FakeAuthorizationV1{Fake: &c.Fake}
}

// AuthorizationV1beta1 retrieves the AuthorizationV1beta1Client
func (c *Clientset) AuthorizationV1beta1() authorizationv1beta1.AuthorizationV1beta1Interface {
	return &fakeauthorizationv1beta1.FakeAuthorizationV1beta1{Fake: &c.Fake}
}

// AutoscalingV1 retrieves the AutoscalingV1Client
func (c *Clientset) AutoscalingV1() autoscalingv1.AutoscalingV1Interface {
	return &fakeautoscalingv1.FakeAutoscalingV1{Fake: &c.Fake}
}

// AutoscalingV2beta1 retrieves the AutoscalingV2beta1Client
func (c *Clientset) AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface {
	return &fakeautoscalingv2beta1.FakeAutoscalingV2beta1{Fake: &c.Fake}
}

// AutoscalingV2beta2 retrieves the AutoscalingV2beta2Client
func (c *Clientset) AutoscalingV2beta2() autoscalingv2beta2.AutoscalingV2beta2Interface {
	return &fakeautoscalingv2beta2.FakeAutoscalingV2beta2{Fake: &c.Fake}
}

// BatchV1 retrieves the BatchV1Client
func (c *Clientset) BatchV1() batchv1.BatchV1Interface {
	return &fakebatchv1.FakeBatchV1{Fake: &c.Fake}
}

// BatchV1beta1 retrieves the BatchV1beta1Client
func (c *Clientset) BatchV1beta1() batchv1beta1.BatchV1beta1Interface {
	return &fakebatchv1beta1.FakeBatchV1beta1{Fake: &c.Fake}
}

// BatchV2alpha1 retrieves the BatchV2alpha1Client
func (c *Clientset) BatchV2alpha1() batchv2alpha1.BatchV2alpha1Interface {
	return &fakebatchv2alpha1.FakeBatchV2alpha1{Fake: &c.Fake}
}

// CertificatesV1 retrieves the CertificatesV1Client
func (c *Clientset) CertificatesV1() certificatesv1.CertificatesV1Interface {
	return &fakecertificatesv1.FakeCertificatesV1{Fake: &c.Fake}
}

// CertificatesV1beta1 retrieves the CertificatesV1beta1Client
func (c *Clientset) CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1Interface {
	return &fakecertificatesv1beta1.FakeCertificatesV1beta1{Fake: &c.Fake}
}

// CoordinationV1beta1 retrieves the CoordinationV1beta1Client
func (c *Clientset) CoordinationV1beta1() coordinationv1beta1.CoordinationV1beta1Interface {
	return &fakecoordinationv1beta1.FakeCoordinationV1beta1{Fake: &c.Fake}
}

// CoordinationV1 retrieves the CoordinationV1Client
func (c *Clientset) CoordinationV1() coordinationv1.CoordinationV1Interface {
	return &fakecoordinationv1.FakeCoordinationV1{Fake: &c.Fake}
}

// CoreV1 retrieves the CoreV1Client
func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	return &fakecorev1.FakeCoreV1{Fake: &c.Fake}
}

// DiscoveryV1alpha1 retrieves the DiscoveryV1alpha1Client
func (c *Clientset) DiscoveryV1alpha1() discoveryv1alpha1.DiscoveryV1alpha1Interface {
	return &fakediscoveryv1alpha1.FakeDiscoveryV1alpha1{Fake: &c.Fake}
}

// DiscoveryV1beta1 retrieves the DiscoveryV1beta1Client
func (c *Clientset) DiscoveryV1beta1() discoveryv1beta1.DiscoveryV1beta1Interface {
	return &fakediscoveryv1beta1.FakeDiscoveryV1beta1{Fake: &c.Fake}
}

// EventsV1 retrieves the EventsV1Client
func (c *Clientset) EventsV1() eventsv1.EventsV1Interface {
	return &fakeeventsv1.FakeEventsV1{Fake: &c.Fake}
}

// EventsV1beta1 retrieves the EventsV1beta1Client
func (c *Clientset) EventsV1beta1() eventsv1beta1.EventsV1beta1Interface {
	return &fakeeventsv1beta1.FakeEventsV1beta1{Fake: &c.Fake}
}

// ExtensionsV1beta1 retrieves the ExtensionsV1beta1Client
func (c *Clientset) ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface {
	return &fakeextensionsv1beta1.FakeExtensionsV1beta1{Fake: &c.Fake}
}

// FlowcontrolV1alpha1 retrieves the FlowcontrolV1alpha1Client
func (c *Clientset) FlowcontrolV1alpha1() flowcontrolv1alpha1.FlowcontrolV1alpha1Interface {
	return &fakeflowcontrolv1alpha1.FakeFlowcontrolV1alpha1{Fake: &c.Fake}
}

// NetworkingV1 retrieves the NetworkingV1Client
func (c *Clientset) NetworkingV1() networkingv1.NetworkingV1Interface {
	return &fakenetworkingv1.FakeNetworkingV1{Fake: &c.Fake}
}

// NetworkingV1beta1 retrieves the NetworkingV1beta1Client
func (c *Clientset) NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1Interface {
	return &fakenetworkingv1beta1.FakeNetworkingV1beta1{Fake: &c.Fake}
}

// NodeV1alpha1 retrieves the NodeV1alpha1Client
func (c *Clientset) NodeV1alpha1() nodev1alpha1.NodeV1alpha1Interface {
	return &fakenodev1alpha1.FakeNodeV1alpha1{Fake: &c.Fake}
}

// NodeV1beta1 retrieves the NodeV1beta1Client
func (c *Clientset) NodeV1beta1() nodev1beta1.NodeV1beta1Interface {
	return &fakenodev1beta1.FakeNodeV1beta1{Fake: &c.Fake}
}

// PolicyV1beta1 retrieves the PolicyV1beta1Client
func (c *Clientset) PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface {
	return &fakepolicyv1beta1.FakePolicyV1beta1{Fake: &c.Fake}
}

// RbacV1 retrieves the RbacV1Client
func (c *Clientset) RbacV1() rbacv1.RbacV1Interface {
	return &fakerbacv1.FakeRbacV1{Fake: &c.Fake}
}

// RbacV1beta1 retrieves the RbacV1beta1Client
func (c *Clientset) RbacV1beta1() rbacv1beta1.RbacV1beta1Interface {
	return &fakerbacv1beta1.FakeRbacV1beta1{Fake: &c.Fake}
}

// RbacV1alpha1 retrieves the RbacV1alpha1Client
func (c *Clientset) RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface {
	return &fakerbacv1alpha1.FakeRbacV1alpha1{Fake: &c.Fake}
}

// SchedulingV1alpha1 retrieves the SchedulingV1alpha1Client
func (c *Clientset) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return &fakeschedulingv1alpha1.FakeSchedulingV1alpha1{Fake: &c.Fake}
}

// SchedulingV1beta1 retrieves the SchedulingV1beta1Client
func (c *Clientset) SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1Interface {
	return &fakeschedulingv1beta1.FakeSchedulingV1beta1{Fake: &c.Fake}
}

// SchedulingV1 retrieves the SchedulingV1Client
func (c *Clientset) SchedulingV1() schedulingv1.SchedulingV1Interface {
	return &fakeschedulingv1.FakeSchedulingV1{Fake: &c.Fake}
}

// SettingsV1alpha1 retrieves the SettingsV1alpha1Client
func (c *Clientset) SettingsV1alpha1() settingsv1alpha1.SettingsV1alpha1Interface {
	return &fakesettingsv1alpha1.FakeSettingsV1alpha1{Fake: &c.Fake}
}

// StorageV1beta1 retrieves the StorageV1beta1Client
func (c *Clientset) StorageV1beta1() storagev1beta1.StorageV1beta1Interface {
	return &fakestoragev1beta1.FakeStorageV1beta1{Fake: &c.Fake}
}

// StorageV1 retrieves the StorageV1Client
func (c *Clientset) StorageV1() storagev1.StorageV1Interface {
	return &fakestoragev1.FakeStorageV1{Fake: &c.Fake}
}

// StorageV1alpha1 retrieves the StorageV1alpha1Client
func (c *Clientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return &fakestoragev1alpha1.FakeStorageV1alpha1{Fake: &c.Fake}
}
