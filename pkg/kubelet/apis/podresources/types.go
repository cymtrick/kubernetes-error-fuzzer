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

package podresources

import (
	"k8s.io/api/core/v1"
	"k8s.io/kubelet/pkg/apis/podresources/v1alpha1"
)

// DevicesProvider knows how to provide the devices used by the given container
type DevicesProvider interface {
	GetDevices(podUID, containerName string) []*v1alpha1.ContainerDevices
	UpdateAllocatedDevices()
}

// PodsProvider knows how to provide the pods admitted by the node
type PodsProvider interface {
	GetPods() []*v1.Pod
}
