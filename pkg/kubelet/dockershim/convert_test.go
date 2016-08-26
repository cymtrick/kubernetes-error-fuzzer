/*
Copyright 2016 The Kubernetes Authors.

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

package dockershim

import (
	"testing"

	"github.com/stretchr/testify/assert"

	runtimeApi "k8s.io/kubernetes/pkg/kubelet/api/v1alpha1/runtime"
)

func TestConvertDockerStatusToRuntimeAPIState(t *testing.T) {
	testCases := []struct {
		input    string
		expected runtimeApi.ContainerState
	}{
		{input: "Up 5 hours", expected: runtimeApi.ContainerState_RUNNING},
		{input: "Exited (0) 2 hours ago", expected: runtimeApi.ContainerState_EXITED},
		{input: "Created", expected: runtimeApi.ContainerState_CREATED},
		{input: "Random string", expected: runtimeApi.ContainerState_UNKNOWN},
	}

	for _, test := range testCases {
		actual := toRuntimeAPIContainerState(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
