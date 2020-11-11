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

package validation

import (
	"fmt"
	"strings"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	featuregatetesting "k8s.io/component-base/featuregate/testing"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/discovery"
	"k8s.io/kubernetes/pkg/features"
	utilpointer "k8s.io/utils/pointer"
)

func TestValidateEndpointSlice(t *testing.T) {
	standardMeta := metav1.ObjectMeta{
		Name:      "hello",
		Namespace: "world",
	}

	testCases := map[string]struct {
		expectedErrors int
		endpointSlice  *discovery.EndpointSlice
	}{
		"good-slice": {
			expectedErrors: 0,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
					Hostname:  utilpointer.StringPtr("valid-123"),
				}},
			},
		},
		"good-fqdns": {
			expectedErrors: 0,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeFQDN,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: []string{"foo.example.com", "example.com", "example.com.", "hyphens-are-good.example.com"},
					Hostname:  utilpointer.StringPtr("valid-123"),
				}},
			},
		},
		"all-protocols": {
			expectedErrors: 0,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("tcp"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}, {
					Name:     utilpointer.StringPtr("udp"),
					Protocol: protocolPtr(api.ProtocolUDP),
				}, {
					Name:     utilpointer.StringPtr("sctp"),
					Protocol: protocolPtr(api.ProtocolSCTP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
					Hostname:  utilpointer.StringPtr("valid-123"),
				}},
			},
		},
		"app-protocols": {
			expectedErrors: 0,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:        utilpointer.StringPtr("one"),
					Protocol:    protocolPtr(api.ProtocolTCP),
					AppProtocol: utilpointer.StringPtr("HTTP"),
				}, {
					Name:        utilpointer.StringPtr("two"),
					Protocol:    protocolPtr(api.ProtocolTCP),
					AppProtocol: utilpointer.StringPtr("https"),
				}, {
					Name:        utilpointer.StringPtr("three"),
					Protocol:    protocolPtr(api.ProtocolTCP),
					AppProtocol: utilpointer.StringPtr("my-protocol"),
				}, {
					Name:        utilpointer.StringPtr("four"),
					Protocol:    protocolPtr(api.ProtocolTCP),
					AppProtocol: utilpointer.StringPtr("example.com/custom-protocol"),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
					Hostname:  utilpointer.StringPtr("valid-123"),
				}},
			},
		},
		"empty-port-name": {
			expectedErrors: 0,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr(""),
					Protocol: protocolPtr(api.ProtocolTCP),
				}, {
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
				}},
			},
		},
		"long-port-name": {
			expectedErrors: 0,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr(strings.Repeat("a", 63)),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
				}},
			},
		},
		"empty-ports-and-endpoints": {
			expectedErrors: 0,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports:       []discovery.EndpointPort{},
				Endpoints:   []discovery.Endpoint{},
			},
		},
		"max-endpoints": {
			expectedErrors: 0,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports:       generatePorts(1),
				Endpoints:   generateEndpoints(maxEndpoints),
			},
		},
		"max-ports": {
			expectedErrors: 0,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports:       generatePorts(maxPorts),
			},
		},
		"max-addresses": {
			expectedErrors: 0,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(maxAddresses),
				}},
			},
		},
		"max-topology-keys": {
			expectedErrors: 0,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
					Topology:  generateTopology(maxTopologyLabels),
				}},
			},
		},

		// expected failures
		"duplicate-port-name": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr(""),
					Protocol: protocolPtr(api.ProtocolTCP),
				}, {
					Name:     utilpointer.StringPtr(""),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{},
			},
		},
		"bad-port-name-caps": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("aCapital"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{},
			},
		},
		"bad-port-name-chars": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("almost_valid"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{},
			},
		},
		"bad-port-name-length": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr(strings.Repeat("a", 64)),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{},
			},
		},
		"invalid-port-protocol": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.Protocol("foo")),
				}},
			},
		},
		"too-many-ports": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports:       generatePorts(maxPorts + 1),
			},
		},
		"too-many-endpoints": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports:       generatePorts(1),
				Endpoints:   generateEndpoints(maxEndpoints + 1),
			},
		},
		"no-endpoint-addresses": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(0),
				}},
			},
		},
		"too-many-addresses": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(maxAddresses + 1),
				}},
			},
		},
		"bad-topology-key": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
					Topology:  map[string]string{"--INVALID": "example"},
				}},
			},
		},
		"too-many-topology-keys": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
					Topology:  generateTopology(maxTopologyLabels + 1),
				}},
			},
		},
		"bad-hostname": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
					Hostname:  utilpointer.StringPtr("--INVALID"),
				}},
			},
		},
		"bad-meta": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "*&^",
					Namespace: "foo",
				},
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
					Hostname:  utilpointer.StringPtr("valid-123"),
				}},
			},
		},
		"bad-ip": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: []string{"123.456.789.012"},
					Hostname:  utilpointer.StringPtr("valid-123"),
				}},
			},
		},
		"bad-ipv4": {
			expectedErrors: 2,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: []string{"123.456.789.012", "2001:4860:4860::8888"},
					Hostname:  utilpointer.StringPtr("valid-123"),
				}},
			},
		},
		"bad-ipv6": {
			expectedErrors: 2,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv6,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: []string{"123.456.789.012", "2001:4860:4860:defg"},
					Hostname:  utilpointer.StringPtr("valid-123"),
				}},
			},
		},
		"bad-fqdns": {
			expectedErrors: 4,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeFQDN,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: []string{"foo.*", "FOO.example.com", "underscores_are_bad.example.com", "*.example.com"},
					Hostname:  utilpointer.StringPtr("valid-123"),
				}},
			},
		},
		"bad-app-protocol": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:        utilpointer.StringPtr("http"),
					Protocol:    protocolPtr(api.ProtocolTCP),
					AppProtocol: utilpointer.StringPtr("--"),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
					Hostname:  utilpointer.StringPtr("valid-123"),
				}},
			},
		},
		"empty-everything": {
			expectedErrors: 3,
			endpointSlice:  &discovery.EndpointSlice{},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			errs := ValidateEndpointSlice(testCase.endpointSlice, true)
			if len(errs) != testCase.expectedErrors {
				t.Errorf("Expected %d errors, got %d errors: %v", testCase.expectedErrors, len(errs), errs)
			}
		})
	}
}

func TestValidateEndpointSliceCreate(t *testing.T) {
	standardMeta := metav1.ObjectMeta{
		Name:      "hello",
		Namespace: "world",
	}

	testCases := map[string]struct {
		expectedErrors      int
		endpointSlice       *discovery.EndpointSlice
		nodeNameGateEnabled bool
	}{
		"good-slice": {
			expectedErrors: 0,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
					Hostname:  utilpointer.StringPtr("valid-123"),
				}},
			},
		},
		"good-slice-node-name": {
			expectedErrors:      0,
			nodeNameGateEnabled: true,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
					Hostname:  utilpointer.StringPtr("valid-123"),
					NodeName:  utilpointer.StringPtr("valid-node-name"),
				}},
			},
		},

		// expected failures
		"bad-node-name": {
			expectedErrors:      1,
			nodeNameGateEnabled: true,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
					Hostname:  utilpointer.StringPtr("valid-123"),
					NodeName:  utilpointer.StringPtr("INvalid-node-name"),
				}},
			},
		},
		"node-name-disabled": {
			expectedErrors:      1,
			nodeNameGateEnabled: false,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
					Hostname:  utilpointer.StringPtr("valid-123"),
					NodeName:  utilpointer.StringPtr("valid-node-name"),
				}},
			},
		},
		"deprecated-address-type": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressType("IP"),
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
				}},
			},
		},
		"bad-address-type": {
			expectedErrors: 1,
			endpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressType("other"),
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr("http"),
					Protocol: protocolPtr(api.ProtocolTCP),
				}},
				Endpoints: []discovery.Endpoint{{
					Addresses: generateIPAddresses(1),
				}},
			},
		},
	}

	for name, testCase := range testCases {
		defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.EndpointSliceNodeName, testCase.nodeNameGateEnabled)()

		t.Run(name, func(t *testing.T) {
			errs := ValidateEndpointSliceCreate(testCase.endpointSlice)
			if len(errs) != testCase.expectedErrors {
				t.Errorf("Expected %d errors, got %d errors: %v", testCase.expectedErrors, len(errs), errs)
			}
		})
	}
}

func TestValidateEndpointSliceUpdate(t *testing.T) {
	standardMeta := metav1.ObjectMeta{Name: "es1", Namespace: "test"}

	testCases := map[string]struct {
		expectedErrors      int
		nodeNameGateEnabled bool
		newEndpointSlice    *discovery.EndpointSlice
		oldEndpointSlice    *discovery.EndpointSlice
	}{
		"valid and identical slices": {
			newEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv6,
			},
			oldEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv6,
			},
			expectedErrors: 0,
		},
		"node name set before + after, gate disabled": {
			newEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Endpoints: []discovery.Endpoint{{
					Addresses: []string{"10.1.2.3"},
					NodeName:  utilpointer.StringPtr("foo"),
				}},
			},
			oldEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Endpoints: []discovery.Endpoint{{
					Addresses: []string{"10.1.2.3"},
					NodeName:  utilpointer.StringPtr("foo"),
				}},
			},
			expectedErrors: 0,
		},
		"node name set after, gate enabled": {
			nodeNameGateEnabled: true,
			newEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Endpoints: []discovery.Endpoint{{
					Addresses: []string{"10.1.2.3"},
				}},
			},
			oldEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Endpoints: []discovery.Endpoint{{
					Addresses: []string{"10.1.2.3"},
					NodeName:  utilpointer.StringPtr("foo"),
				}},
			},
			expectedErrors: 0,
		},

		// expected errors
		"invalide node name set after, gate enabled": {
			nodeNameGateEnabled: true,
			newEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Endpoints: []discovery.Endpoint{{
					Addresses: []string{"10.1.2.3"},
					NodeName:  utilpointer.StringPtr("INVALID foo"),
				}},
			},
			oldEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Endpoints: []discovery.Endpoint{{
					Addresses: []string{"10.1.2.3"},
				}},
			},
			expectedErrors: 1,
		},
		"node name set after, gate disabled": {
			nodeNameGateEnabled: false,
			newEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Endpoints: []discovery.Endpoint{{
					Addresses: []string{"10.1.2.3"},
				}},
			},
			oldEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Endpoints: []discovery.Endpoint{{
					Addresses: []string{"10.1.2.3"},
					NodeName:  utilpointer.StringPtr("foo"),
				}},
			},
			expectedErrors: 0,
		},
		"deprecated address type": {
			expectedErrors: 1,
			newEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressType("IP"),
			},
			oldEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressType("IP"),
			},
		},
		"valid and identical slices with different address types": {
			newEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
			},
			oldEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressType("other"),
			},
			expectedErrors: 1,
		},
		"invalid slices with valid address types": {
			newEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
				Ports: []discovery.EndpointPort{{
					Name:     utilpointer.StringPtr(""),
					Protocol: protocolPtr(api.Protocol("invalid")),
				}},
			},
			oldEndpointSlice: &discovery.EndpointSlice{
				ObjectMeta:  standardMeta,
				AddressType: discovery.AddressTypeIPv4,
			},
			expectedErrors: 1,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.EndpointSliceNodeName, testCase.nodeNameGateEnabled)()

			errs := ValidateEndpointSliceUpdate(testCase.newEndpointSlice, testCase.oldEndpointSlice)
			if len(errs) != testCase.expectedErrors {
				t.Errorf("Expected %d errors, got %d errors: %v", testCase.expectedErrors, len(errs), errs)
			}
		})
	}
}

// Test helpers

func protocolPtr(protocol api.Protocol) *api.Protocol {
	return &protocol
}

func generatePorts(n int) []discovery.EndpointPort {
	ports := []discovery.EndpointPort{}
	for i := 0; i < n; i++ {
		ports = append(ports, discovery.EndpointPort{
			Name:     utilpointer.StringPtr(fmt.Sprintf("http-%d", i)),
			Protocol: protocolPtr(api.ProtocolTCP),
		})
	}
	return ports
}

func generateEndpoints(n int) []discovery.Endpoint {
	endpoints := []discovery.Endpoint{}
	for i := 0; i < n; i++ {
		endpoints = append(endpoints, discovery.Endpoint{
			Addresses: []string{fmt.Sprintf("10.1.2.%d", i%255)},
		})
	}
	return endpoints
}

func generateIPAddresses(n int) []string {
	addresses := []string{}
	for i := 0; i < n; i++ {
		addresses = append(addresses, fmt.Sprintf("10.1.2.%d", i%255))
	}
	return addresses
}

func generateTopology(n int) map[string]string {
	topology := map[string]string{}
	for i := 0; i < n; i++ {
		topology[fmt.Sprintf("topology-%d", i)] = "example"
	}
	return topology
}
