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

package kubeadm

import (
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
)

// SetDefaultTimeouts sets an internal Timeouts struct to its default values.
func SetDefaultTimeouts(t **Timeouts) {
	*t = &Timeouts{
		ControlPlaneComponentHealthCheck: &metav1.Duration{Duration: constants.ControlPlaneComponentHealthCheckTimeout},
		KubeletHealthCheck:               &metav1.Duration{Duration: constants.KubeletHealthCheckTimeout},
		KubernetesAPICall:                &metav1.Duration{Duration: constants.KubernetesAPICallTimeout},
		EtcdAPICall:                      &metav1.Duration{Duration: constants.EtcdAPICallTimeout},
		TLSBootstrap:                     &metav1.Duration{Duration: constants.TLSBootstrapTimeout},
		Discovery:                        &metav1.Duration{Duration: constants.DiscoveryTimeout},
	}
}

var (
	activeTimeouts *Timeouts = nil
	timeoutMutex             = &sync.RWMutex{}
)

func init() {
	SetDefaultTimeouts(&activeTimeouts)
}

// GetActiveTimeouts gets the active timeouts structure.
func GetActiveTimeouts() *Timeouts {
	timeoutMutex.RLock()
	defer timeoutMutex.RUnlock()
	return activeTimeouts
}

// SetActiveTimeouts sets the active timeouts structure.
func SetActiveTimeouts(timeouts *Timeouts) {
	timeoutMutex.Lock()
	activeTimeouts = timeouts.DeepCopy()
	timeoutMutex.Unlock()
}
