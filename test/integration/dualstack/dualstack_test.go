/*
Copyright 2020 The Kubernetes Authors.

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

package dualstack

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"reflect"
	"strings"
	"testing"
	"time"

	jsonpatch "github.com/evanphx/json-patch"

	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/wait"

	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"

	utilfeature "k8s.io/apiserver/pkg/util/feature"
	featuregatetesting "k8s.io/component-base/featuregate/testing"

	"k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/test/integration/framework"
	netutils "k8s.io/utils/net"
)

// TestCreateServiceSingleStackIPv4 test the Service dualstackness in an IPv4 SingleStack cluster
func TestCreateServiceSingleStackIPv4(t *testing.T) {
	// Create an IPv4 single stack control-plane
	serviceCIDR := "10.0.0.0/16"
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.IPv6DualStack, true)()

	cfg := framework.NewIntegrationTestControlPlaneConfig()
	_, cidr, err := net.ParseCIDR(serviceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.ServiceIPRange = *cidr
	_, s, closeFn := framework.RunAnAPIServer(cfg)
	defer closeFn()

	client := clientset.NewForConfigOrDie(&restclient.Config{Host: s.URL})

	// Wait until the default "kubernetes" service is created.
	if err = wait.Poll(250*time.Millisecond, time.Minute, func() (bool, error) {
		_, err := client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), "kubernetes", metav1.GetOptions{})
		if err != nil && !apierrors.IsNotFound(err) {
			return false, err
		}
		return !apierrors.IsNotFound(err), nil
	}); err != nil {
		t.Fatalf("creating kubernetes service timed out")
	}

	var testcases = []struct {
		name               string
		serviceType        v1.ServiceType
		clusterIPs         []string
		ipFamilies         []v1.IPFamily
		ipFamilyPolicy     v1.IPFamilyPolicyType
		expectedIPFamilies []v1.IPFamily
		expectError        bool
	}{
		{
			name:               "Type ClusterIP - Server Allocated IP - Default IP Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         nil,
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - Default IP Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         nil,
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - Default IP Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         nil,
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 IPv6 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 IPv6 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 IPv6 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 IPv4 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 IPv4 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 IPv4 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         nil,
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			expectError:        true,
		},
	}

	for i, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			svc := &v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: fmt.Sprintf("svc-test-%d", i), // use different services for each test
				},
				Spec: v1.ServiceSpec{
					Type:           tc.serviceType,
					ClusterIPs:     tc.clusterIPs,
					IPFamilies:     tc.ipFamilies,
					IPFamilyPolicy: &tc.ipFamilyPolicy,
					Ports: []v1.ServicePort{
						{
							Port:       443,
							TargetPort: intstr.FromInt(443),
						},
					},
				},
			}

			// create the service
			_, err = client.CoreV1().Services(metav1.NamespaceDefault).Create(context.TODO(), svc, metav1.CreateOptions{})
			if (err != nil) != tc.expectError {
				t.Errorf("Test failed expected result: %v received %v ", tc.expectError, err)
			}
			// if no error was expected validate the service otherwise return
			if err != nil {
				return
			}
			// validate the service was created correctly if it was not expected to fail
			svc, err = client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), svc.Name, metav1.GetOptions{})
			if err != nil {
				t.Errorf("Unexpected error to get the service %s %v", svc.Name, err)
			}
			if err := validateServiceAndClusterIPFamily(svc, tc.expectedIPFamilies); err != nil {
				t.Errorf("Unexpected error validating the service %s\n%+v\n%v", svc.Name, svc, err)
			}
		})
	}
}

// TestCreateServiceDualStackIPv6 test the Service dualstackness in an IPv6 only DualStack cluster
func TestCreateServiceDualStackIPv6(t *testing.T) {
	// Create an IPv6 only dual stack control-plane
	serviceCIDR := "2001:db8:1::/48"
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.IPv6DualStack, true)()

	cfg := framework.NewIntegrationTestControlPlaneConfig()
	_, cidr, err := net.ParseCIDR(serviceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.ServiceIPRange = *cidr
	_, s, closeFn := framework.RunAnAPIServer(cfg)
	defer closeFn()

	client := clientset.NewForConfigOrDie(&restclient.Config{Host: s.URL})

	// Wait until the default "kubernetes" service is created.
	if err = wait.Poll(250*time.Millisecond, time.Minute, func() (bool, error) {
		_, err := client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), "kubernetes", metav1.GetOptions{})
		if err != nil && !apierrors.IsNotFound(err) {
			return false, err
		}
		return !apierrors.IsNotFound(err), nil
	}); err != nil {
		t.Fatalf("creating kubernetes service timed out")
	}

	var testcases = []struct {
		name               string
		serviceType        v1.ServiceType
		clusterIPs         []string
		ipFamilies         []v1.IPFamily
		expectedIPFamilies []v1.IPFamily
		ipFamilyPolicy     v1.IPFamilyPolicyType
		expectError        bool
	}{
		{
			name:               "Type ClusterIP - Server Allocated IP - Default IP Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         nil,
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - Default IP Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         nil,
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - Default IP Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         nil,
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 IPv6 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 IPv6 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 IPv6 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 IPv4 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 IPv4 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 IPv4 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        true,
		},
	}

	for i, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			svc := &v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: fmt.Sprintf("svc-test-%d", i), // use different services for each test
				},
				Spec: v1.ServiceSpec{
					Type:           tc.serviceType,
					ClusterIPs:     tc.clusterIPs,
					IPFamilies:     tc.ipFamilies,
					IPFamilyPolicy: &tc.ipFamilyPolicy,
					Ports: []v1.ServicePort{
						{
							Name:       fmt.Sprintf("port-test-%d", i),
							Port:       443,
							TargetPort: intstr.IntOrString{IntVal: 443},
							Protocol:   "TCP",
						},
					},
				},
			}

			// create the service
			_, err = client.CoreV1().Services(metav1.NamespaceDefault).Create(context.TODO(), svc, metav1.CreateOptions{})
			if (err != nil) != tc.expectError {
				t.Errorf("Test failed expected result: %v received %v ", tc.expectError, err)
			}
			// if no error was expected validate the service otherwise return
			if err != nil {
				return
			}
			// validate the service was created correctly if it was not expected to fail
			svc, err = client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), svc.Name, metav1.GetOptions{})
			if err != nil {
				t.Errorf("Unexpected error to get the service %s %v", svc.Name, err)
			}
			if err := validateServiceAndClusterIPFamily(svc, tc.expectedIPFamilies); err != nil {
				t.Errorf("Unexpected error validating the service %s %v", svc.Name, err)
			}
		})
	}
}

// TestCreateServiceDualStackIPv4IPv6 test the Service dualstackness in a IPv4IPv6 DualStack cluster
func TestCreateServiceDualStackIPv4IPv6(t *testing.T) {
	// Create an IPv4IPv6 dual stack control-plane
	serviceCIDR := "10.0.0.0/16"
	secondaryServiceCIDR := "2001:db8:1::/48"
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.IPv6DualStack, true)()

	cfg := framework.NewIntegrationTestControlPlaneConfig()
	_, cidr, err := net.ParseCIDR(serviceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.ServiceIPRange = *cidr

	_, secCidr, err := net.ParseCIDR(secondaryServiceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.SecondaryServiceIPRange = *secCidr

	_, s, closeFn := framework.RunAnAPIServer(cfg)
	defer closeFn()

	client := clientset.NewForConfigOrDie(&restclient.Config{Host: s.URL})

	// Wait until the default "kubernetes" service is created.
	if err = wait.Poll(250*time.Millisecond, time.Minute, func() (bool, error) {
		_, err := client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), "kubernetes", metav1.GetOptions{})
		if err != nil && !apierrors.IsNotFound(err) {
			return false, err
		}
		return !apierrors.IsNotFound(err), nil
	}); err != nil {
		t.Fatalf("creating kubernetes service timed out")
	}

	var testcases = []struct {
		name               string
		serviceType        v1.ServiceType
		clusterIPs         []string
		ipFamilies         []v1.IPFamily
		expectedIPFamilies []v1.IPFamily
		ipFamilyPolicy     v1.IPFamilyPolicyType
		expectError        bool
	}{
		{
			name:               "Type ClusterIP - Server Allocated IP - Default IP Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         nil,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - Default IP Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         nil,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - Default IP Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         nil,
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        false,
		},

		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 IPv6 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 IPv6 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 IPv6 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 IPv4 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 IPv4 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 IPv4 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        false,
		},
	}

	for i, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			svc := &v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: fmt.Sprintf("svc-test-%d", i), // use different services for each test
				},
				Spec: v1.ServiceSpec{
					Type:           tc.serviceType,
					ClusterIPs:     tc.clusterIPs,
					IPFamilies:     tc.ipFamilies,
					IPFamilyPolicy: &tc.ipFamilyPolicy,
					Ports: []v1.ServicePort{
						{
							Port:       443,
							TargetPort: intstr.FromInt(443),
						},
					},
				},
			}

			// create a service
			_, err = client.CoreV1().Services(metav1.NamespaceDefault).Create(context.TODO(), svc, metav1.CreateOptions{})
			if (err != nil) != tc.expectError {
				t.Errorf("Test failed expected result: %v received %v ", tc.expectError, err)
			}
			// if no error was expected validate the service otherwise return
			if err != nil {
				return
			}
			// validate the service was created correctly if it was not expected to fail
			svc, err = client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), svc.Name, metav1.GetOptions{})
			if err != nil {
				t.Errorf("Unexpected error to get the service %s %v", svc.Name, err)
			}

			if err := validateServiceAndClusterIPFamily(svc, tc.expectedIPFamilies); err != nil {
				t.Errorf("Unexpected error validating the service %s %v", svc.Name, err)
			}
		})
	}
}

// TestCreateServiceDualStackIPv6IPv4 test the Service dualstackness in a IPv6IPv4 DualStack cluster
func TestCreateServiceDualStackIPv6IPv4(t *testing.T) {
	// Create an IPv6IPv4 dual stack control-plane
	serviceCIDR := "2001:db8:1::/48"
	secondaryServiceCIDR := "10.0.0.0/16"
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.IPv6DualStack, true)()

	cfg := framework.NewIntegrationTestControlPlaneConfig()
	_, cidr, err := net.ParseCIDR(serviceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.ServiceIPRange = *cidr

	_, secCidr, err := net.ParseCIDR(secondaryServiceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.SecondaryServiceIPRange = *secCidr

	_, s, closeFn := framework.RunAnAPIServer(cfg)
	defer closeFn()

	client := clientset.NewForConfigOrDie(&restclient.Config{Host: s.URL})

	// Wait until the default "kubernetes" service is created.
	if err = wait.Poll(250*time.Millisecond, time.Minute, func() (bool, error) {
		_, err := client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), "kubernetes", metav1.GetOptions{})
		if err != nil && !apierrors.IsNotFound(err) {
			return false, err
		}
		return !apierrors.IsNotFound(err), nil
	}); err != nil {
		t.Fatalf("creating kubernetes service timed out")
	}

	// verify client is working
	if err := wait.PollImmediate(5*time.Second, 2*time.Minute, func() (bool, error) {
		_, err = client.CoreV1().Endpoints("default").Get(context.TODO(), "kubernetes", metav1.GetOptions{})
		if err != nil {
			t.Logf("error fetching endpoints: %v", err)
			return false, nil
		}
		return true, nil
	}); err != nil {
		t.Errorf("server without enabled endpoints failed to register: %v", err)
	}

	var testcases = []struct {
		name               string
		serviceType        v1.ServiceType
		clusterIPs         []string
		ipFamilies         []v1.IPFamily
		expectedIPFamilies []v1.IPFamily
		ipFamilyPolicy     v1.IPFamilyPolicyType
		expectError        bool
	}{
		{
			name:               "Type ClusterIP - Server Allocated IP - Default IP Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         nil,
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - Default IP Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         nil,
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - Default IP Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         nil,
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        false,
		},

		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 IPv6 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv4 IPv6 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 IPv4 Family - Policy Single Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicySingleStack,
			expectError:        true,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 IPv4 Family - Policy Prefer Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyPreferDualStack,
			expectError:        false,
		},
		{
			name:               "Type ClusterIP - Server Allocated IP - IPv6 IPv4 Family - Policy Required Dual Stack",
			serviceType:        v1.ServiceTypeClusterIP,
			clusterIPs:         []string{},
			ipFamilies:         []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			expectedIPFamilies: []v1.IPFamily{v1.IPv6Protocol, v1.IPv4Protocol},
			ipFamilyPolicy:     v1.IPFamilyPolicyRequireDualStack,
			expectError:        false,
		},
	}

	for i, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			svc := &v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: fmt.Sprintf("svc-test-%d", i), // use different services for each test
				},
				Spec: v1.ServiceSpec{
					Type:           tc.serviceType,
					ClusterIPs:     tc.clusterIPs,
					IPFamilies:     tc.ipFamilies,
					IPFamilyPolicy: &tc.ipFamilyPolicy,
					Ports: []v1.ServicePort{
						{
							Port:       443,
							TargetPort: intstr.FromInt(443),
						},
					},
				},
			}

			// create a service
			_, err = client.CoreV1().Services(metav1.NamespaceDefault).Create(context.TODO(), svc, metav1.CreateOptions{})
			if (err != nil) != tc.expectError {
				t.Errorf("Test failed expected result: %v received %v ", tc.expectError, err)
			}
			// if no error was expected validate the service otherwise return
			if err != nil {
				return
			}
			// validate the service was created correctly if it was not expected to fail
			svc, err = client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), svc.Name, metav1.GetOptions{})
			if err != nil {
				t.Errorf("Unexpected error to get the service %s %v", svc.Name, err)
			}

			if err := validateServiceAndClusterIPFamily(svc, tc.expectedIPFamilies); err != nil {
				t.Errorf("Unexpected error validating the service %s %v", svc.Name, err)
			}
		})
	}
}

// TestUpgradeDowngrade tests upgrading and downgrading a service from/to dual-stack
func TestUpgradeDowngrade(t *testing.T) {
	// Create an IPv4IPv6 dual stack control-plane
	serviceCIDR := "10.0.0.0/16"
	secondaryServiceCIDR := "2001:db8:1::/48"
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.IPv6DualStack, true)()

	cfg := framework.NewIntegrationTestControlPlaneConfig()
	_, cidr, err := net.ParseCIDR(serviceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.ServiceIPRange = *cidr

	_, secCidr, err := net.ParseCIDR(secondaryServiceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.SecondaryServiceIPRange = *secCidr

	_, s, closeFn := framework.RunAnAPIServer(cfg)
	defer closeFn()

	client := clientset.NewForConfigOrDie(&restclient.Config{Host: s.URL})

	// Wait until the default "kubernetes" service is created.
	if err = wait.Poll(250*time.Millisecond, time.Minute, func() (bool, error) {
		_, err := client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), "kubernetes", metav1.GetOptions{})
		if err != nil && !apierrors.IsNotFound(err) {
			return false, err
		}
		return !apierrors.IsNotFound(err), nil
	}); err != nil {
		t.Fatalf("creating kubernetes service timed out")
	}

	upgradeServiceName := "svc-upgrade"

	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: upgradeServiceName,
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeClusterIP,
			Ports: []v1.ServicePort{
				{
					Port:       443,
					TargetPort: intstr.FromInt(443),
				},
			},
		},
	}

	// create a service
	_, err = client.CoreV1().Services(metav1.NamespaceDefault).Create(context.TODO(), svc, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("unexpected error while creating service:%v", err)
	}
	// validate the service was created correctly if it was not expected to fail
	svc, err = client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), svc.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Unexpected error to get the service %s %v", svc.Name, err)
	}

	if err := validateServiceAndClusterIPFamily(svc, []v1.IPFamily{v1.IPv4Protocol} /* default cluster config */); err != nil {
		t.Fatalf("Unexpected error validating the service %s %v", svc.Name, err)
	}

	// upgrade it
	requireDualStack := v1.IPFamilyPolicyRequireDualStack
	svc.Spec.IPFamilyPolicy = &requireDualStack
	upgraded, err := client.CoreV1().Services(metav1.NamespaceDefault).Update(context.TODO(), svc, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("unexpected error upgrading service to dual stack. %v", err)
	}
	if err := validateServiceAndClusterIPFamily(upgraded, []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol} /* +1 family */); err != nil {
		t.Fatalf("Unexpected error validating the service(after upgrade) %s %v", svc.Name, err)
	}

	// downgrade it
	singleStack := v1.IPFamilyPolicySingleStack
	upgraded.Spec.IPFamilyPolicy = &singleStack
	upgraded.Spec.ClusterIPs = upgraded.Spec.ClusterIPs[0:1]
	upgraded.Spec.IPFamilies = upgraded.Spec.IPFamilies[0:1]
	downgraded, err := client.CoreV1().Services(metav1.NamespaceDefault).Update(context.TODO(), upgraded, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("unexpected error downgrading service to single stack. %v", err)
	}
	if err := validateServiceAndClusterIPFamily(downgraded, []v1.IPFamily{v1.IPv4Protocol} /* -1 family */); err != nil {
		t.Fatalf("unexpected error validating the service(after downgrade) %s %v", svc.Name, err)
	}

	// run test again this time without removing secondary IPFamily or ClusterIP
	downgraded.Spec.IPFamilyPolicy = &requireDualStack
	upgradedAgain, err := client.CoreV1().Services(metav1.NamespaceDefault).Update(context.TODO(), downgraded, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("unexpected error upgrading service to dual stack. %v", err)
	}
	if err := validateServiceAndClusterIPFamily(upgradedAgain, []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol} /* +1 family */); err != nil {
		t.Fatalf("Unexpected error validating the service(after upgrade) %s %v", svc.Name, err)
	}

	upgradedAgain.Spec.IPFamilyPolicy = &singleStack
	// api-server automatically  removes the secondary ClusterIP and IPFamily
	// when a servie is downgraded.
	downgradedAgain, err := client.CoreV1().Services(metav1.NamespaceDefault).Update(context.TODO(), upgradedAgain, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("unexpected error downgrading service to single stack. %v", err)
	}
	if err := validateServiceAndClusterIPFamily(downgradedAgain, []v1.IPFamily{v1.IPv4Protocol} /* -1 family */); err != nil {
		t.Fatalf("unexpected error validating the service(after downgrade) %s %v", svc.Name, err)
	}
}

// TestConvertToFromExternalName tests the compatibility with old clients that
// may not clear ClusterIPs
func TestConvertToFromExternalName(t *testing.T) {
	// Create an IPv4IPv6 dual stack control-plane
	serviceCIDR := "10.0.0.0/16"
	secondaryServiceCIDR := "2001:db8:1::/48"
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.IPv6DualStack, true)()

	cfg := framework.NewIntegrationTestControlPlaneConfig()
	_, cidr, err := net.ParseCIDR(serviceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.ServiceIPRange = *cidr

	_, secCidr, err := net.ParseCIDR(secondaryServiceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.SecondaryServiceIPRange = *secCidr

	_, s, closeFn := framework.RunAnAPIServer(cfg)
	defer closeFn()

	client := clientset.NewForConfigOrDie(&restclient.Config{Host: s.URL})

	// Wait until the default "kubernetes" service is created.
	if err = wait.Poll(250*time.Millisecond, time.Minute, func() (bool, error) {
		_, err := client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), "kubernetes", metav1.GetOptions{})
		if err != nil && !apierrors.IsNotFound(err) {
			return false, err
		}
		return !apierrors.IsNotFound(err), nil
	}); err != nil {
		t.Fatalf("creating kubernetes service timed out")
	}

	serviceName := "svc-ext-name"
	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeClusterIP,
			Ports: []v1.ServicePort{
				{
					Port:       443,
					TargetPort: intstr.FromInt(443),
				},
			},
		},
	}

	// create a service
	_, err = client.CoreV1().Services(metav1.NamespaceDefault).Create(context.TODO(), svc, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("unexpected error while creating service:%v", err)
	}
	// validate the service was created correctly if it was not expected to fail
	svc, err = client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), svc.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Unexpected error to get the service %s %v", svc.Name, err)
	}

	if err := validateServiceAndClusterIPFamily(svc, []v1.IPFamily{v1.IPv4Protocol}); err != nil {
		t.Fatalf("Unexpected error validating the service %s %v", svc.Name, err)
	}

	// convert to ExternalName
	svc.Spec.Type = v1.ServiceTypeExternalName
	svc.Spec.ClusterIP = "" // not clearing ClusterIPs
	svc.Spec.ExternalName = "something.somewhere"

	externalNameSvc, err := client.CoreV1().Services(metav1.NamespaceDefault).Update(context.TODO(), svc, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("unexpected error converting service to external name. %v", err)
	}

	if len(externalNameSvc.Spec.ClusterIPs) > 0 || len(externalNameSvc.Spec.ClusterIP) > 0 || len(externalNameSvc.Spec.IPFamilies) > 0 {
		t.Fatalf("unpexpected externalname service with ClusterIPs %v or ClusterIP %v or IPFamilies %v", externalNameSvc.Spec.ClusterIPs, externalNameSvc.Spec.ClusterIP, externalNameSvc.Spec.IPFamilies)
	}

	// convert to a ClusterIP service
	externalNameSvc.Spec.Type = v1.ServiceTypeClusterIP
	externalNameSvc.Spec.ExternalName = ""
	clusterIPSvc, err := client.CoreV1().Services(metav1.NamespaceDefault).Update(context.TODO(), externalNameSvc, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("unexpected error converting service to ClusterIP. %v", err)
	}
	if err := validateServiceAndClusterIPFamily(clusterIPSvc, []v1.IPFamily{v1.IPv4Protocol}); err != nil {
		t.Fatalf("Unexpected error validating the service %s %v", svc.Name, err)
	}
}

// TestExistingServiceDefaulting tests that existing services are defaulted correctly after an upgrade
func TestExistingServiceDefaulting(t *testing.T) {
	// Create an IPv4IPv6 dual stack control-plane
	serviceCIDR := "10.0.0.0/16"
	cfg := framework.NewIntegrationTestControlPlaneConfig()
	_, cidr, err := net.ParseCIDR(serviceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.ServiceIPRange = *cidr

	_, s, closeFn := framework.RunAnAPIServer(cfg)
	defer closeFn()

	client := clientset.NewForConfigOrDie(&restclient.Config{Host: s.URL})

	// Wait until the default "kubernetes" service is created.
	if err = wait.Poll(250*time.Millisecond, time.Minute, func() (bool, error) {
		_, err := client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), "kubernetes", metav1.GetOptions{})
		if err != nil && !apierrors.IsNotFound(err) {
			return false, err
		}
		return !apierrors.IsNotFound(err), nil
	}); err != nil {
		t.Fatalf("creating kubernetes service timed out")
	}

	serviceName := "svc-ext-name"
	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeClusterIP,
			Ports: []v1.ServicePort{
				{
					Port:       443,
					TargetPort: intstr.FromInt(443),
				},
			},
		},
	}
	// make sure gate is off (cluster is not running as dual stack)
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.IPv6DualStack, false)()

	// create a service
	_, err = client.CoreV1().Services(metav1.NamespaceDefault).Create(context.TODO(), svc, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("unexpected error while creating service:%v", err)
	}

	// validate the service was created correctly if it was not expected to fail
	svc, err = client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), svc.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Unexpected error to get the service %s %v", svc.Name, err)
	}

	// turn gate on
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.IPv6DualStack, true)()

	// validate the service was created correctly if it was not expected to fail
	defaultedSvc, err := client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), svc.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Unexpected error to get the service %s %v", svc.Name, err)
	}

	if defaultedSvc.Spec.IPFamilyPolicy == nil {
		t.Fatalf("service should have an IPFamily Policy")
	} else {
		// must be equal to single stack
		if *(defaultedSvc.Spec.IPFamilyPolicy) != v1.IPFamilyPolicySingleStack {
			t.Fatalf("service should have a SingleStack as IPFamilyPolicy instead we got:%v", defaultedSvc.Spec.IPFamilyPolicy)
		}
	}

	if err := validateServiceAndClusterIPFamily(defaultedSvc, []v1.IPFamily{v1.IPv4Protocol}); err != nil {
		t.Fatalf("Unexpected error validating the service %s %v", svc.Name, err)
	}
}

// TestPreferDualStack preferDualstack on create and update
func TestPreferDualStack(t *testing.T) {
	// Create an IPv4IPv6 dual stack control-plane
	serviceCIDR := "10.0.0.0/16"
	secondaryServiceCIDR := "2001:db8:1::/48"
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.IPv6DualStack, true)()

	cfg := framework.NewIntegrationTestControlPlaneConfig()
	_, cidr, err := net.ParseCIDR(serviceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.ServiceIPRange = *cidr

	_, secCidr, err := net.ParseCIDR(secondaryServiceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.SecondaryServiceIPRange = *secCidr

	_, s, closeFn := framework.RunAnAPIServer(cfg)
	defer closeFn()

	client := clientset.NewForConfigOrDie(&restclient.Config{Host: s.URL})

	// Wait until the default "kubernetes" service is created.
	if err = wait.Poll(250*time.Millisecond, time.Minute, func() (bool, error) {
		_, err := client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), "kubernetes", metav1.GetOptions{})
		if err != nil && !apierrors.IsNotFound(err) {
			return false, err
		}
		return !apierrors.IsNotFound(err), nil
	}); err != nil {
		t.Fatalf("creating kubernetes service timed out")
	}

	preferDualStack := v1.IPFamilyPolicyPreferDualStack

	serviceName := "svc-upgrade"

	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: v1.ServiceSpec{
			Type:           v1.ServiceTypeClusterIP,
			IPFamilyPolicy: &preferDualStack,
			Ports: []v1.ServicePort{
				{
					Port:       443,
					TargetPort: intstr.FromInt(443),
				},
			},
		},
	}

	// create a service
	_, err = client.CoreV1().Services(metav1.NamespaceDefault).Create(context.TODO(), svc, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("unexpected error while creating service:%v", err)
	}
	// validate the service was created correctly if it was not expected to fail
	svc, err = client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), svc.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Unexpected error to get the service %s %v", svc.Name, err)
	}

	if err := validateServiceAndClusterIPFamily(svc, []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol}); err != nil {
		t.Fatalf("Unexpected error validating the service %s %v", svc.Name, err)
	}

	// update it
	svc.Spec.Selector = map[string]string{"foo": "bar"}
	upgraded, err := client.CoreV1().Services(metav1.NamespaceDefault).Update(context.TODO(), svc, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("unexpected error upgrading service to dual stack. %v", err)
	}
	if err := validateServiceAndClusterIPFamily(upgraded, []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol}); err != nil {
		t.Fatalf("Unexpected error validating the service(after upgrade) %s %v", svc.Name, err)
	}
}

type labelsForMergePatch struct {
	Labels map[string]string `json:"lables,omitempty"`
}

// tests an update service while dualstack flag is off
func TestServiceUpdate(t *testing.T) {
	// Create an IPv4 single stack control-plane
	serviceCIDR := "10.0.0.0/16"
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.IPv6DualStack, false)()

	cfg := framework.NewIntegrationTestControlPlaneConfig()
	_, cidr, err := net.ParseCIDR(serviceCIDR)
	if err != nil {
		t.Fatalf("bad cidr: %v", err)
	}
	cfg.ExtraConfig.ServiceIPRange = *cidr
	_, s, closeFn := framework.RunAnAPIServer(cfg)
	defer closeFn()

	client := clientset.NewForConfigOrDie(&restclient.Config{Host: s.URL})

	// Wait until the default "kubernetes" service is created.
	if err = wait.Poll(250*time.Millisecond, time.Minute, func() (bool, error) {
		_, err := client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), "kubernetes", metav1.GetOptions{})
		if err != nil && !apierrors.IsNotFound(err) {
			return false, err
		}
		return !apierrors.IsNotFound(err), nil
	}); err != nil {
		t.Fatalf("creating kubernetes service timed out")
	}

	serviceName := "test-service"
	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeClusterIP,
			Ports: []v1.ServicePort{
				{
					Port:       443,
					TargetPort: intstr.FromInt(443),
				},
			},
		},
	}

	// create the service
	_, err = client.CoreV1().Services(metav1.NamespaceDefault).Create(context.TODO(), svc, metav1.CreateOptions{})
	// if no error was expected validate the service otherwise return
	if err != nil {
		t.Errorf("unexpected error creating service:%v", err)
		return
	}

	svc, err = client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), svc.Name, metav1.GetOptions{})
	if err != nil {
		t.Errorf("Unexpected error to get the service %s %v", svc.Name, err)
	}

	// update using put
	svc.Labels = map[string]string{"x": "y"}
	_, err = client.CoreV1().Services(metav1.NamespaceDefault).Update(context.TODO(), svc, metav1.UpdateOptions{})
	if err != nil {
		t.Errorf("Unexpected error updating the service %s %v", svc.Name, err)
	}

	_, err = client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), svc.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Unexpected error to get the service %s %v", svc.Name, err)
	}

	// update using StrategicMergePatchType
	labels := labelsForMergePatch{
		Labels: map[string]string{"foo": "bar"},
	}

	patchBytes, err := json.Marshal(&labels)
	if err != nil {
		t.Fatalf("failed to json.Marshal labels: %v", err)
	}

	_, err = client.CoreV1().Services(metav1.NamespaceDefault).Patch(context.TODO(), svc.Name, types.StrategicMergePatchType, patchBytes, metav1.PatchOptions{})
	if err != nil {
		t.Fatalf("unexpected error patching service using strategic merge patch. %v", err)
	}

	current, err := client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), svc.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Unexpected error to get the service %s %v", svc.Name, err)
	}

	// update using json patch
	toUpdate := current.DeepCopy()
	currentJSON, err := json.Marshal(current)
	if err != nil {
		t.Fatalf("unexpected error marshal current service. %v", err)
	}
	toUpdate.Labels = map[string]string{"alpha": "bravo"}
	toUpdateJSON, err := json.Marshal(toUpdate)
	if err != nil {
		t.Fatalf("unexpected error marshal toupdate service. %v", err)
	}

	patchBytes, err = jsonpatch.CreateMergePatch(currentJSON, toUpdateJSON)
	if err != nil {
		t.Fatalf("unexpected error creating json patch. %v", err)
	}

	_, err = client.CoreV1().Services(metav1.NamespaceDefault).Patch(context.TODO(), svc.Name, types.MergePatchType, patchBytes, metav1.PatchOptions{})
	if err != nil {
		t.Fatalf("unexpected error patching service using merge patch. %v", err)
	}

	// validate the service was created correctly if it was not expected to fail
	_, err = client.CoreV1().Services(metav1.NamespaceDefault).Get(context.TODO(), svc.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Unexpected error to get the service %s %v", svc.Name, err)
	}

}

// validateServiceAndClusterIPFamily checks that the service has the expected IPFamilies
func validateServiceAndClusterIPFamily(svc *v1.Service, expectedIPFamilies []v1.IPFamily) error {
	// create a slice for the errors
	var errstrings []string

	if svc.Spec.IPFamilies == nil {
		return fmt.Errorf("service ip family nil for service %s/%s", svc.Namespace, svc.Name)
	}
	if !reflect.DeepEqual(svc.Spec.IPFamilies, expectedIPFamilies) {
		return fmt.Errorf("ip families mismatch for service: %s/%s, expected: %s, actual: %s", svc.Namespace, svc.Name, expectedIPFamilies, svc.Spec.IPFamilies)
	}

	if len(svc.Spec.ClusterIPs) == 0 {
		return fmt.Errorf("svc %s is invalid it does not have ClusterIP", svc.Name)
	}

	// not headless
	if len(svc.Spec.ClusterIPs) > 0 && svc.Spec.ClusterIPs[0] != v1.ClusterIPNone {
		if len(svc.Spec.ClusterIPs) != len(svc.Spec.IPFamilies) {
			return fmt.Errorf("svc %v is invalid len(ClusterIPs:%v) != len(IPFamilies:%v)", svc.Name, svc.Spec.ClusterIPs, svc.Spec.IPFamilies)
		}
	}

	for j, ip := range svc.Spec.ClusterIPs {
		// we should never be here
		if ip == v1.ClusterIPNone && len(svc.Spec.ClusterIPs) > 1 {
			errstrings = append(errstrings, fmt.Sprintf("Error validating Service: %s, None is used with +1 clusterIPs (%v)", svc.Name, svc.Spec.ClusterIPs))
		}

		if ip == v1.ClusterIPNone {
			break // the service is headless. the rest of family check is pointless
		}

		// the clusterIP assigned should have the same IPFamily requested
		if netutils.IsIPv6String(ip) != (expectedIPFamilies[j] == v1.IPv6Protocol) {
			errstrings = append(errstrings, fmt.Sprintf("got unexpected service ip %s, should belong to %s ip family", ip, expectedIPFamilies[j]))
		}
	}

	if len(errstrings) > 0 {
		errstrings = append(errstrings, fmt.Sprintf("Error validating Service: %s, ClusterIPs: %v Expected IPFamilies %v", svc.Name, svc.Spec.ClusterIPs, expectedIPFamilies))
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}

	return nil
}
