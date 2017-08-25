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

package upgrade

import (
	clientset "k8s.io/client-go/kubernetes"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/pkg/util/version"
)

// PerformStaticPodControlPlaneUpgrade upgrades a static pod-hosted control plane
func PerformStaticPodControlPlaneUpgrade(_ clientset.Interface, _ *kubeadmapi.MasterConfiguration, _ *version.Version) error {
	// No-op for now; doesn't do anything yet
	return nil
}
