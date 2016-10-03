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

package common

import (
	"strings"
	"time"

	"github.com/golang/glog"
	. "github.com/onsi/ginkgo"
	api "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/test/e2e/framework"
)

const (
	etcHostsImageName          = "gcr.io/google_containers/netexec:1.7"
	etcHostsPodName            = "test-pod"
	etcHostsHostNetworkPodName = "test-host-network-pod"
	etcHostsPartialContent     = "# Kubernetes-managed hosts file."
)

type KubeletManagedHostConfig struct {
	hostNetworkPod *api.Pod
	pod            *api.Pod
	f              *framework.Framework
}

var _ = framework.KubeDescribe("KubeletManagedEtcHosts", func() {
	f := framework.NewDefaultFramework("e2e-kubelet-etc-hosts")
	config := &KubeletManagedHostConfig{
		f: f,
	}

	It("should test kubelet managed /etc/hosts file", func() {
		By("Setting up the test")
		config.setup()

		By("Running the test")
		config.verifyEtcHosts()
	})
})

func (config *KubeletManagedHostConfig) verifyEtcHosts() {
	By("Verifying /etc/hosts of container is kubelet-managed for pod with hostNetwork=false")
	assertManagedStatus(config, etcHostsPodName, true, "busybox-1")
	assertManagedStatus(config, etcHostsPodName, true, "busybox-2")

	By("Verifying /etc/hosts of container is not kubelet-managed since container specifies /etc/hosts mount")
	assertManagedStatus(config, etcHostsPodName, false, "busybox-3")

	By("Verifying /etc/hosts content of container is not kubelet-managed for pod with hostNetwork=true")
	assertManagedStatus(config, etcHostsHostNetworkPodName, false, "busybox-1")
	assertManagedStatus(config, etcHostsHostNetworkPodName, false, "busybox-2")
}

func (config *KubeletManagedHostConfig) setup() {
	By("Creating hostNetwork=false pod")
	config.createPodWithoutHostNetwork()

	By("Creating hostNetwork=true pod")
	config.createPodWithHostNetwork()
}

func (config *KubeletManagedHostConfig) createPodWithoutHostNetwork() {
	podSpec := config.createPodSpec(etcHostsPodName)
	config.pod = config.f.PodClient().CreateSync(podSpec)
}

func (config *KubeletManagedHostConfig) createPodWithHostNetwork() {
	podSpec := config.createPodSpecWithHostNetwork(etcHostsHostNetworkPodName)
	config.hostNetworkPod = config.f.PodClient().CreateSync(podSpec)
}

func assertManagedStatus(
	config *KubeletManagedHostConfig, podName string, expectedIsManaged bool, name string) {
	// See https://github.com/kubernetes/kubernetes/issues/27023
	//
	// Retry until timeout for the right contents of /etc/hosts to show
	// up. There may be a low probability race here. We still fail the
	// test if retry was necessary, but at least we will know whether or
	// not it resolves or seems to be a permanent condition.
	//
	// If /etc/hosts is properly mounted, then this will succeed
	// immediately.
	const retryTimeout = 30 * time.Second

	retryCount := 0
	etcHostsContent := ""
	matched := false

	for startTime := time.Now(); time.Since(startTime) < retryTimeout; {
		etcHostsContent = config.getEtcHostsContent(podName, name)
		isManaged := strings.Contains(etcHostsContent, etcHostsPartialContent)

		if expectedIsManaged == isManaged {
			matched = true
			break
		}

		glog.Errorf(
			"For pod: %s, name: %s, expected %t, actual %t (/etc/hosts was %q), retryCount: %d",
			podName, name, expectedIsManaged, isManaged, etcHostsContent, retryCount)

		retryCount++
		time.Sleep(100 * time.Millisecond)
	}

	if retryCount > 0 {
		if matched {
			conditionText := "should"
			if !expectedIsManaged {
				conditionText = "should not"
			}

			framework.Failf(
				"/etc/hosts file %s be kubelet managed (name: %s, retries: %d). /etc/hosts contains %q",
				conditionText, name, retryCount, etcHostsContent)
		} else {
			framework.Failf(
				"had to retry %d times to get matching content in /etc/hosts (name: %s)",
				retryCount, name)
		}
	}
}

func (config *KubeletManagedHostConfig) getEtcHostsContent(podName, containerName string) string {
	return config.f.ExecCommandInContainer(podName, containerName, "cat", "/etc/hosts")
}

func (config *KubeletManagedHostConfig) createPodSpec(podName string) *api.Pod {
	pod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			Name: podName,
		},
		Spec: api.PodSpec{
			Containers: []api.Container{
				{
					Name:            "busybox-1",
					Image:           etcHostsImageName,
					ImagePullPolicy: api.PullIfNotPresent,
					Command: []string{
						"sleep",
						"900",
					},
				},
				{
					Name:            "busybox-2",
					Image:           etcHostsImageName,
					ImagePullPolicy: api.PullIfNotPresent,
					Command: []string{
						"sleep",
						"900",
					},
				},
				{
					Name:            "busybox-3",
					Image:           etcHostsImageName,
					ImagePullPolicy: api.PullIfNotPresent,
					Command: []string{
						"sleep",
						"900",
					},
					VolumeMounts: []api.VolumeMount{
						{
							Name:      "host-etc-hosts",
							MountPath: "/etc/hosts",
						},
					},
				},
			},
			Volumes: []api.Volume{
				{
					Name: "host-etc-hosts",
					VolumeSource: api.VolumeSource{
						HostPath: &api.HostPathVolumeSource{
							Path: "/etc/hosts",
						},
					},
				},
			},
		},
	}
	return pod
}

func (config *KubeletManagedHostConfig) createPodSpecWithHostNetwork(podName string) *api.Pod {
	pod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			Name: podName,
		},
		Spec: api.PodSpec{
			SecurityContext: &api.PodSecurityContext{
				HostNetwork: true,
			},
			Containers: []api.Container{
				{
					Name:            "busybox-1",
					Image:           etcHostsImageName,
					ImagePullPolicy: api.PullIfNotPresent,
					Command: []string{
						"sleep",
						"900",
					},
				},
				{
					Name:            "busybox-2",
					Image:           etcHostsImageName,
					ImagePullPolicy: api.PullIfNotPresent,
					Command: []string{
						"sleep",
						"900",
					},
				},
			},
		},
	}
	return pod
}
