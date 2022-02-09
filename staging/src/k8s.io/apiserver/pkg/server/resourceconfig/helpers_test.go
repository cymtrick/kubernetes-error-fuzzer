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

package resourceconfig

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	extensionsapiv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	serverstore "k8s.io/apiserver/pkg/server/storage"
)

func TestParseRuntimeConfig(t *testing.T) {
	scheme := newFakeScheme(t)
	apiv1GroupVersion := apiv1.SchemeGroupVersion
	testCases := []struct {
		name                  string
		runtimeConfig         map[string]string
		defaultResourceConfig func() *serverstore.ResourceConfig
		expectedAPIConfig     func() *serverstore.ResourceConfig
		expectedEnabledAPIs   map[schema.GroupVersionResource]bool
		err                   bool
	}{
		{
			name: "using-kind",
			runtimeConfig: map[string]string{
				"apps/v1/Deployment": "false",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedEnabledAPIs: defaultFakeEnabledResources(),
			err:                 true,
		},
		{
			name:          "everything-default-value",
			runtimeConfig: map[string]string{},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedEnabledAPIs: defaultFakeEnabledResources(),
			err:                 false,
		},
		{
			name:          "no-runtimeConfig-override",
			runtimeConfig: map[string]string{},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableVersions(extensionsapiv1beta1.SchemeGroupVersion)
				return config
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableVersions(extensionsapiv1beta1.SchemeGroupVersion)
				return config
			},
			expectedEnabledAPIs: defaultFakeEnabledResources(),
			err:                 false,
		},
		{
			name: "version-enabled-by-runtimeConfig-override",
			runtimeConfig: map[string]string{
				"apps/v1": "",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				return config
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				return config
			},
			expectedEnabledAPIs: defaultFakeEnabledResources(),
			err:                 false,
		},
		{
			name: "disable-v1",
			runtimeConfig: map[string]string{
				"/v1": "false",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableVersions(apiv1GroupVersion)
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               true,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       false,
			},
			err: false,
		},
		{
			name: "invalid-runtime-config",
			runtimeConfig: map[string]string{
				"invalidgroup/version": "false",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedEnabledAPIs: defaultFakeEnabledResources(),
			err:                 false,
		},
		{
			name: "enable-all",
			runtimeConfig: map[string]string{
				"api/all": "true",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.EnableVersions(scheme.PrioritizedVersionsAllGroups()...)
				// disabling groups of APIs removes the individual resource preferences from the default
				config.RemoveMatchingResourcePreferences(matchAllExplicitResourcesForFake)
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  true,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               true,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       true,
			},
			err: false,
		},
		{
			name: "only-enable-v1",
			runtimeConfig: map[string]string{
				"api/all": "false",
				"/v1":     "true",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableVersions(appsv1.SchemeGroupVersion)
				config.DisableVersions(extensionsapiv1beta1.SchemeGroupVersion)
				// disabling groups of APIs removes the individual resource preferences from the default
				config.RemoveMatchingResourcePreferences(matchAllExplicitResourcesForFake)
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               false,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       true,
			},
			err: false,
		},
		{
			name: "enable-specific-extensions-resources",
			runtimeConfig: map[string]string{
				"extensions/v1beta1/deployments": "true",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.EnableResources(extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"))
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               true,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       true,
			}, err: false,
		},
		{
			name: "disable-specific-extensions-resources",
			runtimeConfig: map[string]string{
				"extensions/v1beta1/ingresses": "false",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableResources(extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"))
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               true,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       true,
			}, err: false,
		},
		{
			name: "disable-all-extensions-resources",
			runtimeConfig: map[string]string{
				"extensions/v1beta1": "false",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableVersions(extensionsapiv1beta1.SchemeGroupVersion)
				// disabling groups of APIs removes the individual resource preferences from the default
				config.RemoveMatchingResourcePreferences(matchAllExplicitResourcesForFake)
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               true,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       true,
			}, err: false,
		},
		{
			name: "disable-a-no-extensions-resources",
			runtimeConfig: map[string]string{
				"apps/v1/deployments": "false",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableResources(appsv1.SchemeGroupVersion.WithResource("deployments"))
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               false,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       true,
			},
			err: false, // no error for backwards compatibility
		},
		{
			name: "disable-all-beta-resources",
			runtimeConfig: map[string]string{
				"api/beta": "false",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableVersions(extensionsapiv1beta1.SchemeGroupVersion)
				// disabling groups of APIs removes the individual resource preferences from the default
				config.RemoveMatchingResourcePreferences(matchAllExplicitResourcesForFake)
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               true,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       true,
			},
			err: false, // no error for backwards compatibility
		},
		{
			name: "user-explicit-disable-resource-over-user-version-enable",
			runtimeConfig: map[string]string{
				"apps/v1":             "true",
				"apps/v1/deployments": "false",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableResources(appsv1.SchemeGroupVersion.WithResource("deployments"))
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               false,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       true,
			},
			err: false, // no error for backwards compatibility
		},
		{
			name: "user-explicit-enable-resource-over-user-version-disable",
			runtimeConfig: map[string]string{
				"apps/v1":             "false",
				"apps/v1/deployments": "true",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableVersions(appsv1.SchemeGroupVersion)
				config.EnableResources(appsv1.SchemeGroupVersion.WithResource("deployments"))
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               true,
				appsv1.SchemeGroupVersion.WithResource("other"):                     false,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       true,
			},
			err: false, // no error for backwards compatibility
		},
		{
			name: "user-explicit-disable-resource-over-user-stability-enable",
			runtimeConfig: map[string]string{
				"api/ga":              "true",
				"apps/v1/deployments": "false",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableResources(appsv1.SchemeGroupVersion.WithResource("deployments"))
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               false,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       true,
			},
			err: false, // no error for backwards compatibility
		},
		{
			name: "user-explicit-enable-resource-over-user-stability-disable",
			runtimeConfig: map[string]string{
				"api/ga":              "false",
				"apps/v1/deployments": "true",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableVersions(apiv1.SchemeGroupVersion)
				config.DisableVersions(appsv1.SchemeGroupVersion)
				config.EnableResources(appsv1.SchemeGroupVersion.WithResource("deployments"))
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               true,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       false,
			},
			err: false, // no error for backwards compatibility
		},
		{
			name: "user-explicit-disable-resource-over-user-version-enable-over-user-stability-disable",
			runtimeConfig: map[string]string{
				"api/ga":              "false",
				"apps/v1":             "true",
				"apps/v1/deployments": "false",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableVersions(apiv1.SchemeGroupVersion)
				config.EnableVersions(appsv1.SchemeGroupVersion)
				config.DisableResources(appsv1.SchemeGroupVersion.WithResource("deployments"))
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               false,
				appsv1.SchemeGroupVersion.WithResource("other"):                     true,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       false,
			},
			err: false, // no error for backwards compatibility
		},
		{
			name: "user-explicit-enable-resource-over-user-version-disable-over-user-stability-disable",
			runtimeConfig: map[string]string{
				"api/ga":              "false",
				"apps/v1":             "false",
				"apps/v1/deployments": "true",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableVersions(apiv1.SchemeGroupVersion)
				config.DisableVersions(appsv1.SchemeGroupVersion)
				config.EnableResources(appsv1.SchemeGroupVersion.WithResource("deployments"))
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               true,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       false,
			},
			err: false, // no error for backwards compatibility
		},
		{
			name: "user-explicit-disable-resource-over-user-version-enable-over-user-stability-enable",
			runtimeConfig: map[string]string{
				"api/ga":              "true",
				"apps/v1":             "true",
				"apps/v1/deployments": "false",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.EnableVersions(appsv1.SchemeGroupVersion)
				config.DisableResources(appsv1.SchemeGroupVersion.WithResource("deployments"))
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               false,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       true,
			},
			err: false, // no error for backwards compatibility
		},
		{
			name: "user-explicit-enable-resource-over-user-version-disable-over-user-stability-enable",
			runtimeConfig: map[string]string{
				"api/ga":              "true",
				"apps/v1":             "false",
				"apps/v1/deployments": "true",
			},
			defaultResourceConfig: func() *serverstore.ResourceConfig {
				return newFakeAPIResourceConfigSource()
			},
			expectedAPIConfig: func() *serverstore.ResourceConfig {
				config := newFakeAPIResourceConfigSource()
				config.DisableVersions(appsv1.SchemeGroupVersion)
				config.EnableResources(appsv1.SchemeGroupVersion.WithResource("deployments"))
				return config
			},
			expectedEnabledAPIs: map[schema.GroupVersionResource]bool{
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   true,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
				extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
				appsv1.SchemeGroupVersion.WithResource("deployments"):               true,
				appsv1.SchemeGroupVersion.WithResource("other"):                     false,
				apiv1.SchemeGroupVersion.WithResource("pods"):                       true,
			},
			err: false, // no error for backwards compatibility
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Log(scheme.PrioritizedVersionsAllGroups())
			actualDisablers, err := MergeAPIResourceConfigs(test.defaultResourceConfig(), test.runtimeConfig, scheme)
			if err == nil && test.err {
				t.Fatalf("expected error")
			} else if err != nil && !test.err {
				t.Fatalf("unexpected error: %s, for test: %v", err, test)
			}
			if err != nil {
				return
			}

			expectedConfig := test.expectedAPIConfig()
			if !reflect.DeepEqual(actualDisablers, expectedConfig) {
				t.Fatalf("%v: unexpected apiResourceDisablers. Actual: %v\n expected: %v", test.runtimeConfig, actualDisablers, expectedConfig)
			}

			for _, resourceToCheck := range apiResourcesToCheck() {
				actual := actualDisablers.ResourceEnabled(resourceToCheck)
				expected := test.expectedEnabledAPIs[resourceToCheck]
				if actual != expected {
					t.Errorf("for %v, actual=%v, expected=%v", resourceToCheck, actual, expected)
				}
			}
			for resourceToCheck, expected := range test.expectedEnabledAPIs {
				actual := actualDisablers.ResourceEnabled(resourceToCheck)
				if actual != expected {
					t.Errorf("for %v, actual=%v, expected=%v", resourceToCheck, actual, expected)
				}
			}
		})
	}
}

func newFakeAPIResourceConfigSource() *serverstore.ResourceConfig {
	ret := serverstore.NewResourceConfig()
	// NOTE: GroupVersions listed here will be enabled by default. Don't put alpha versions in the list.
	ret.EnableVersions(
		apiv1.SchemeGroupVersion,
		appsv1.SchemeGroupVersion,
		extensionsapiv1beta1.SchemeGroupVersion,
	)
	ret.EnableResources(
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"),
	)
	ret.DisableResources(
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"),
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"),
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"),
	)

	return ret
}

func matchAllExplicitResourcesForFake(gvr schema.GroupVersionResource) bool {
	switch gvr {
	case extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"),
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"),
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"),
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):
		return true
	}
	return false
}

// apiResourcesToCheck are the apis we use in this set of unit tests.  They will be check for enable/disable status
func apiResourcesToCheck() []schema.GroupVersionResource {
	return []schema.GroupVersionResource{
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"),
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"),
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"),
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"),
		appsv1.SchemeGroupVersion.WithResource("deployments"),
		apiv1.SchemeGroupVersion.WithResource("pods"),
	}
}

func defaultFakeEnabledResources() map[schema.GroupVersionResource]bool {
	return map[schema.GroupVersionResource]bool{
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("ingresses"):   true,
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("deployments"): false,
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("replicasets"): false,
		extensionsapiv1beta1.SchemeGroupVersion.WithResource("daemonsets"):  false,
		appsv1.SchemeGroupVersion.WithResource("deployments"):               true,
		apiv1.SchemeGroupVersion.WithResource("pods"):                       true,
	}
}

func newFakeScheme(t *testing.T) *runtime.Scheme {
	ret := runtime.NewScheme()
	require.NoError(t, apiv1.AddToScheme(ret))
	require.NoError(t, appsv1.AddToScheme(ret))
	require.NoError(t, extensionsapiv1beta1.AddToScheme(ret))

	require.NoError(t, ret.SetVersionPriority(apiv1.SchemeGroupVersion))
	require.NoError(t, ret.SetVersionPriority(extensionsapiv1beta1.SchemeGroupVersion))

	return ret
}
