/*
Copyright 2023 The Kubernetes Authors.

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

package util

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "k8s.io/api/core/v1"

	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
)

func TestMergeKubeadmEnvVars(t *testing.T) {
	baseEnv := []kubeadmapi.EnvVar{}
	extraEnv := []kubeadmapi.EnvVar{}
	MergeKubeadmEnvVars(append(baseEnv, extraEnv...))
	var tests = []struct {
		name      string
		proxyEnv  []kubeadmapi.EnvVar
		extraEnv  []kubeadmapi.EnvVar
		mergedEnv []v1.EnvVar
	}{
		{
			name: "normal case without duplicated env",
			proxyEnv: []kubeadmapi.EnvVar{
				{
					EnvVar: v1.EnvVar{Name: "Foo1", Value: "Bar1"},
				},
				{
					EnvVar: v1.EnvVar{Name: "Foo2", Value: "Bar2"},
				},
			},
			extraEnv: []kubeadmapi.EnvVar{
				{
					EnvVar: v1.EnvVar{Name: "Foo3", Value: "Bar3"},
				},
			},
			mergedEnv: []v1.EnvVar{
				{Name: "Foo1", Value: "Bar1"},
				{Name: "Foo2", Value: "Bar2"},
				{Name: "Foo3", Value: "Bar3"},
			},
		},
		{
			name: "extraEnv env take precedence over the proxyEnv",
			proxyEnv: []kubeadmapi.EnvVar{
				{
					EnvVar: v1.EnvVar{Name: "Foo1", Value: "Bar1"},
				},
				{
					EnvVar: v1.EnvVar{Name: "Foo2", Value: "Bar2"},
				},
			},
			extraEnv: []kubeadmapi.EnvVar{
				{
					EnvVar: v1.EnvVar{Name: "Foo2", Value: "Bar3"},
				},
			},
			mergedEnv: []v1.EnvVar{
				{Name: "Foo1", Value: "Bar1"},
				{Name: "Foo2", Value: "Bar3"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			envs := MergeKubeadmEnvVars(test.proxyEnv, test.extraEnv)
			if !assert.ElementsMatch(t, envs, test.mergedEnv) {
				t.Errorf("expected env: %v, got: %v", test.mergedEnv, envs)
			}
		})
	}
}
