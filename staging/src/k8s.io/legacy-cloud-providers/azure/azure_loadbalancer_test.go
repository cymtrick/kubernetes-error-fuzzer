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

package azure

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-07-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestFindProbe(t *testing.T) {
	tests := []struct {
		msg           string
		existingProbe []network.Probe
		curProbe      network.Probe
		expected      bool
	}{
		{
			msg:      "empty existing probes should return false",
			expected: false,
		},
		{
			msg: "probe names match while ports unmatch should return false",
			existingProbe: []network.Probe{
				{
					Name: to.StringPtr("httpProbe"),
					ProbePropertiesFormat: &network.ProbePropertiesFormat{
						Port: to.Int32Ptr(1),
					},
				},
			},
			curProbe: network.Probe{
				Name: to.StringPtr("httpProbe"),
				ProbePropertiesFormat: &network.ProbePropertiesFormat{
					Port: to.Int32Ptr(2),
				},
			},
			expected: false,
		},
		{
			msg: "probe ports match while names unmatch should return false",
			existingProbe: []network.Probe{
				{
					Name: to.StringPtr("probe1"),
					ProbePropertiesFormat: &network.ProbePropertiesFormat{
						Port: to.Int32Ptr(1),
					},
				},
			},
			curProbe: network.Probe{
				Name: to.StringPtr("probe2"),
				ProbePropertiesFormat: &network.ProbePropertiesFormat{
					Port: to.Int32Ptr(1),
				},
			},
			expected: false,
		},
		{
			msg: "both probe ports and names match should return true",
			existingProbe: []network.Probe{
				{
					Name: to.StringPtr("matchName"),
					ProbePropertiesFormat: &network.ProbePropertiesFormat{
						Port: to.Int32Ptr(1),
					},
				},
			},
			curProbe: network.Probe{
				Name: to.StringPtr("matchName"),
				ProbePropertiesFormat: &network.ProbePropertiesFormat{
					Port: to.Int32Ptr(1),
				},
			},
			expected: true,
		},
	}

	for i, test := range tests {
		findResult := findProbe(test.existingProbe, test.curProbe)
		assert.Equal(t, test.expected, findResult, fmt.Sprintf("TestCase[%d]: %s", i, test.msg))
	}
}

func TestFindRule(t *testing.T) {
	tests := []struct {
		msg          string
		existingRule []network.LoadBalancingRule
		curRule      network.LoadBalancingRule
		expected     bool
	}{
		{
			msg:      "empty existing rules should return false",
			expected: false,
		},
		{
			msg: "rule names unmatch should return false",
			existingRule: []network.LoadBalancingRule{
				{
					Name: to.StringPtr("httpProbe1"),
					LoadBalancingRulePropertiesFormat: &network.LoadBalancingRulePropertiesFormat{
						FrontendPort: to.Int32Ptr(1),
					},
				},
			},
			curRule: network.LoadBalancingRule{
				Name: to.StringPtr("httpProbe2"),
				LoadBalancingRulePropertiesFormat: &network.LoadBalancingRulePropertiesFormat{
					FrontendPort: to.Int32Ptr(1),
				},
			},
			expected: false,
		},
		{
			msg: "rule names match while frontend ports unmatch should return false",
			existingRule: []network.LoadBalancingRule{
				{
					Name: to.StringPtr("httpProbe"),
					LoadBalancingRulePropertiesFormat: &network.LoadBalancingRulePropertiesFormat{
						FrontendPort: to.Int32Ptr(1),
					},
				},
			},
			curRule: network.LoadBalancingRule{
				Name: to.StringPtr("httpProbe"),
				LoadBalancingRulePropertiesFormat: &network.LoadBalancingRulePropertiesFormat{
					FrontendPort: to.Int32Ptr(2),
				},
			},
			expected: false,
		},
		{
			msg: "rule names match while backend ports unmatch should return false",
			existingRule: []network.LoadBalancingRule{
				{
					Name: to.StringPtr("httpProbe"),
					LoadBalancingRulePropertiesFormat: &network.LoadBalancingRulePropertiesFormat{
						BackendPort: to.Int32Ptr(1),
					},
				},
			},
			curRule: network.LoadBalancingRule{
				Name: to.StringPtr("httpProbe"),
				LoadBalancingRulePropertiesFormat: &network.LoadBalancingRulePropertiesFormat{
					BackendPort: to.Int32Ptr(2),
				},
			},
			expected: false,
		},
		{
			msg: "rule names match while LoadDistribution unmatch should return false",
			existingRule: []network.LoadBalancingRule{
				{
					Name: to.StringPtr("probe1"),
					LoadBalancingRulePropertiesFormat: &network.LoadBalancingRulePropertiesFormat{
						LoadDistribution: network.Default,
					},
				},
			},
			curRule: network.LoadBalancingRule{
				Name: to.StringPtr("probe2"),
				LoadBalancingRulePropertiesFormat: &network.LoadBalancingRulePropertiesFormat{
					LoadDistribution: network.SourceIP,
				},
			},
			expected: false,
		},
		{
			msg: "both rule names and LoadBalancingRulePropertiesFormats match should return true",
			existingRule: []network.LoadBalancingRule{
				{
					Name: to.StringPtr("matchName"),
					LoadBalancingRulePropertiesFormat: &network.LoadBalancingRulePropertiesFormat{
						BackendPort:      to.Int32Ptr(2),
						FrontendPort:     to.Int32Ptr(2),
						LoadDistribution: network.SourceIP,
					},
				},
			},
			curRule: network.LoadBalancingRule{
				Name: to.StringPtr("matchName"),
				LoadBalancingRulePropertiesFormat: &network.LoadBalancingRulePropertiesFormat{
					BackendPort:      to.Int32Ptr(2),
					FrontendPort:     to.Int32Ptr(2),
					LoadDistribution: network.SourceIP,
				},
			},
			expected: true,
		},
	}

	for i, test := range tests {
		findResult := findRule(test.existingRule, test.curRule, true)
		assert.Equal(t, test.expected, findResult, fmt.Sprintf("TestCase[%d]: %s", i, test.msg))
	}
}

func TestGetIdleTimeout(t *testing.T) {
	for _, c := range []struct {
		desc        string
		annotations map[string]string
		i           *int32
		err         bool
	}{
		{desc: "no annotation"},
		{desc: "annotation empty value", annotations: map[string]string{ServiceAnnotationLoadBalancerIdleTimeout: ""}, err: true},
		{desc: "annotation not a number", annotations: map[string]string{ServiceAnnotationLoadBalancerIdleTimeout: "cookies"}, err: true},
		{desc: "annotation negative value", annotations: map[string]string{ServiceAnnotationLoadBalancerIdleTimeout: "-6"}, err: true},
		{desc: "annotation zero value", annotations: map[string]string{ServiceAnnotationLoadBalancerIdleTimeout: "0"}, err: true},
		{desc: "annotation too low value", annotations: map[string]string{ServiceAnnotationLoadBalancerIdleTimeout: "3"}, err: true},
		{desc: "annotation too high value", annotations: map[string]string{ServiceAnnotationLoadBalancerIdleTimeout: "31"}, err: true},
		{desc: "annotation good value", annotations: map[string]string{ServiceAnnotationLoadBalancerIdleTimeout: "24"}, i: to.Int32Ptr(24)},
	} {
		t.Run(c.desc, func(t *testing.T) {
			s := &v1.Service{}
			s.Annotations = c.annotations
			i, err := getIdleTimeout(s)

			if !reflect.DeepEqual(c.i, i) {
				t.Fatalf("got unexpected value: %d", to.Int32(i))
			}
			if (err != nil) != c.err {
				t.Fatalf("expected error=%v, got %v", c.err, err)
			}
		})
	}
}

func TestSubnet(t *testing.T) {
	for i, c := range []struct {
		desc     string
		service  *v1.Service
		expected *string
	}{
		{
			desc:     "No annotation should return nil",
			service:  &v1.Service{},
			expected: nil,
		},
		{
			desc: "annotation with subnet but no ILB should return nil",
			service: &v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						ServiceAnnotationLoadBalancerInternalSubnet: "subnet",
					},
				},
			},
			expected: nil,
		},
		{
			desc: "annotation with subnet but ILB=false should return nil",
			service: &v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						ServiceAnnotationLoadBalancerInternalSubnet: "subnet",
						ServiceAnnotationLoadBalancerInternal:       "false",
					},
				},
			},
			expected: nil,
		},
		{
			desc: "annotation with empty subnet should return nil",
			service: &v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						ServiceAnnotationLoadBalancerInternalSubnet: "",
						ServiceAnnotationLoadBalancerInternal:       "true",
					},
				},
			},
			expected: nil,
		},
		{
			desc: "annotation with subnet and ILB should return subnet",
			service: &v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						ServiceAnnotationLoadBalancerInternalSubnet: "subnet",
						ServiceAnnotationLoadBalancerInternal:       "true",
					},
				},
			},
			expected: to.StringPtr("subnet"),
		},
	} {
		real := subnet(c.service)
		assert.Equal(t, c.expected, real, fmt.Sprintf("TestCase[%d]: %s", i, c.desc))
	}
}

func TestEnsureLoadBalancerDeleted(t *testing.T) {
	const vmCount = 8
	const availabilitySetCount = 4
	const serviceCount = 9

	tests := []struct {
		desc              string
		service           v1.Service
		expectCreateError bool
	}{
		{
			desc:    "external service should be created and deleted successfully",
			service: getTestService("test1", v1.ProtocolTCP, 80),
		},
		{
			desc:    "internal service should be created and deleted successfully",
			service: getInternalTestService("test2", 80),
		},
		{
			desc:    "annotated service with same resourceGroup should be created and deleted successfully",
			service: getResourceGroupTestService("test3", "rg", "", 80),
		},
		{
			desc:              "annotated service with different resourceGroup shouldn't be created but should be deleted successfully",
			service:           getResourceGroupTestService("test4", "random-rg", "1.2.3.4", 80),
			expectCreateError: true,
		},
	}

	az := getTestCloud()
	for i, c := range tests {
		clusterResources := getClusterResources(az, vmCount, availabilitySetCount)
		getTestSecurityGroup(az)
		if c.service.Annotations[ServiceAnnotationLoadBalancerInternal] == "true" {
			addTestSubnet(t, az, &c.service)
		}

		// create the service first.
		lbStatus, err := az.EnsureLoadBalancer(context.TODO(), testClusterName, &c.service, clusterResources.nodes)
		if c.expectCreateError {
			assert.NotNil(t, err, "TestCase[%d]: %s", i, c.desc)
		} else {
			assert.Nil(t, err, "TestCase[%d]: %s", i, c.desc)
			assert.NotNil(t, lbStatus, "TestCase[%d]: %s", i, c.desc)
			result, err := az.LoadBalancerClient.List(context.TODO(), az.Config.ResourceGroup)
			assert.Nil(t, err, "TestCase[%d]: %s", i, c.desc)
			assert.Equal(t, len(result), 1, "TestCase[%d]: %s", i, c.desc)
			assert.Equal(t, len(*result[0].LoadBalancingRules), 1, "TestCase[%d]: %s", i, c.desc)
		}

		// finally, delete it.
		err = az.EnsureLoadBalancerDeleted(context.TODO(), testClusterName, &c.service)
		assert.Nil(t, err, "TestCase[%d]: %s", i, c.desc)
		result, err := az.LoadBalancerClient.List(context.Background(), az.Config.ResourceGroup)
		assert.Nil(t, err, "TestCase[%d]: %s", i, c.desc)
		assert.Equal(t, len(result), 0, "TestCase[%d]: %s", i, c.desc)
	}
}

func TestServiceOwnsPublicIP(t *testing.T) {
	tests := []struct {
		desc        string
		pip         *network.PublicIPAddress
		clusterName string
		serviceName string
		expected    bool
	}{
		{
			desc:        "false should be returned when pip is nil",
			clusterName: "kubernetes",
			serviceName: "nginx",
			expected:    false,
		},
		{
			desc: "false should be returned when service name tag doesn't match",
			pip: &network.PublicIPAddress{
				Tags: map[string]*string{
					serviceTagKey: to.StringPtr("nginx"),
				},
			},
			serviceName: "web",
			expected:    false,
		},
		{
			desc: "true should be returned when service name tag matches and cluster name tag is not set",
			pip: &network.PublicIPAddress{
				Tags: map[string]*string{
					serviceTagKey: to.StringPtr("nginx"),
				},
			},
			clusterName: "kubernetes",
			serviceName: "nginx",
			expected:    true,
		},
		{
			desc: "false should be returned when cluster name doesn't match",
			pip: &network.PublicIPAddress{
				Tags: map[string]*string{
					serviceTagKey:  to.StringPtr("nginx"),
					clusterNameKey: to.StringPtr("kubernetes"),
				},
			},
			clusterName: "k8s",
			serviceName: "nginx",
			expected:    false,
		},
		{
			desc: "false should be returned when cluster name matches while service name doesn't match",
			pip: &network.PublicIPAddress{
				Tags: map[string]*string{
					serviceTagKey:  to.StringPtr("web"),
					clusterNameKey: to.StringPtr("kubernetes"),
				},
			},
			clusterName: "kubernetes",
			serviceName: "nginx",
			expected:    false,
		},
		{
			desc: "true should be returned when both service name tag and cluster name match",
			pip: &network.PublicIPAddress{
				Tags: map[string]*string{
					serviceTagKey:  to.StringPtr("nginx"),
					clusterNameKey: to.StringPtr("kubernetes"),
				},
			},
			clusterName: "kubernetes",
			serviceName: "nginx",
			expected:    true,
		},
	}

	for i, c := range tests {
		owns := serviceOwnsPublicIP(c.pip, c.clusterName, c.serviceName)
		assert.Equal(t, owns, c.expected, "TestCase[%d]: %s", i, c.desc)
	}
}

func TestGetPublicIPAddressResourceGroup(t *testing.T) {
	az := getTestCloud()

	for i, c := range []struct {
		desc        string
		annotations map[string]string
		expected    string
	}{
		{
			desc:     "no annotation",
			expected: "rg",
		},
		{
			desc:        "annoation with empty string resource group",
			annotations: map[string]string{ServiceAnnotationLoadBalancerResourceGroup: ""},
			expected:    "rg",
		},
		{
			desc:        "annoation with non-empty resource group ",
			annotations: map[string]string{ServiceAnnotationLoadBalancerResourceGroup: "rg2"},
			expected:    "rg2",
		},
	} {
		t.Run(c.desc, func(t *testing.T) {
			s := &v1.Service{}
			s.Annotations = c.annotations
			real := az.getPublicIPAddressResourceGroup(s)
			assert.Equal(t, c.expected, real, "TestCase[%d]: %s", i, c.desc)
		})
	}
}

func TestGetServiceTags(t *testing.T) {
	tests := []struct {
		desc     string
		service  *v1.Service
		expected []string
	}{
		{
			desc:     "nil should be returned when service is nil",
			service:  nil,
			expected: nil,
		},
		{
			desc:     "nil should be returned when service has no annotations",
			service:  &v1.Service{},
			expected: nil,
		},
		{
			desc: "single tag should be returned when service has set one annotations",
			service: &v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						ServiceAnnotationAllowedServiceTag: "tag1",
					},
				},
			},
			expected: []string{"tag1"},
		},
		{
			desc: "multiple tags should be returned when service has set multi-annotations",
			service: &v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						ServiceAnnotationAllowedServiceTag: "tag1, tag2",
					},
				},
			},
			expected: []string{"tag1", "tag2"},
		},
		{
			desc: "correct tags should be returned when comma or spaces are included in the annotations",
			service: &v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						ServiceAnnotationAllowedServiceTag: ", tag1, ",
					},
				},
			},
			expected: []string{"tag1"},
		},
	}

	for i, c := range tests {
		tags := getServiceTags(c.service)
		assert.Equal(t, tags, c.expected, "TestCase[%d]: %s", i, c.desc)
	}
}
