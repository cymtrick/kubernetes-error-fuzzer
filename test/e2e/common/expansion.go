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
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/kubernetes/test/e2e/framework"
	imageutils "k8s.io/kubernetes/test/utils/image"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// These tests exercise the Kubernetes expansion syntax $(VAR).
// For more information, see:
// https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/expansion.md
var _ = framework.KubeDescribe("Variable Expansion", func() {
	f := framework.NewDefaultFramework("var-expansion")

	/*
		Release : v1.9
		Testname: Environment variables, expansion
		Description: Create a Pod with environment variables. Environment variables defined using previously defined environment variables MUST expand to proper values.
	*/
	framework.ConformanceIt("should allow composing env vars into new env vars [NodeConformance]", func() {
		podName := "var-expansion-" + string(uuid.NewUUID())
		pod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:   podName,
				Labels: map[string]string{"name": podName},
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:    "dapi-container",
						Image:   imageutils.GetE2EImage(imageutils.BusyBox),
						Command: []string{"sh", "-c", "env"},
						Env: []v1.EnvVar{
							{
								Name:  "FOO",
								Value: "foo-value",
							},
							{
								Name:  "BAR",
								Value: "bar-value",
							},
							{
								Name:  "FOOBAR",
								Value: "$(FOO);;$(BAR)",
							},
						},
					},
				},
				RestartPolicy: v1.RestartPolicyNever,
			},
		}

		f.TestContainerOutput("env composition", pod, 0, []string{
			"FOO=foo-value",
			"BAR=bar-value",
			"FOOBAR=foo-value;;bar-value",
		})
	})

	/*
		Release : v1.9
		Testname: Environment variables, command expansion
		Description: Create a Pod with environment variables and container command using them. Container command using the  defined environment variables MUST expand to proper values.
	*/
	framework.ConformanceIt("should allow substituting values in a container's command [NodeConformance]", func() {
		podName := "var-expansion-" + string(uuid.NewUUID())
		pod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:   podName,
				Labels: map[string]string{"name": podName},
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:    "dapi-container",
						Image:   imageutils.GetE2EImage(imageutils.BusyBox),
						Command: []string{"sh", "-c", "TEST_VAR=wrong echo \"$(TEST_VAR)\""},
						Env: []v1.EnvVar{
							{
								Name:  "TEST_VAR",
								Value: "test-value",
							},
						},
					},
				},
				RestartPolicy: v1.RestartPolicyNever,
			},
		}

		f.TestContainerOutput("substitution in container's command", pod, 0, []string{
			"test-value",
		})
	})

	/*
		Release : v1.9
		Testname: Environment variables, command argument expansion
		Description: Create a Pod with environment variables and container command arguments using them. Container command arguments using the  defined environment variables MUST expand to proper values.
	*/
	framework.ConformanceIt("should allow substituting values in a container's args [NodeConformance]", func() {
		podName := "var-expansion-" + string(uuid.NewUUID())
		pod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:   podName,
				Labels: map[string]string{"name": podName},
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:    "dapi-container",
						Image:   imageutils.GetE2EImage(imageutils.BusyBox),
						Command: []string{"sh", "-c"},
						Args:    []string{"TEST_VAR=wrong echo \"$(TEST_VAR)\""},
						Env: []v1.EnvVar{
							{
								Name:  "TEST_VAR",
								Value: "test-value",
							},
						},
					},
				},
				RestartPolicy: v1.RestartPolicyNever,
			},
		}

		f.TestContainerOutput("substitution in container's args", pod, 0, []string{
			"test-value",
		})
	})

	/*
		    Testname: var-expansion-subpath
		    Description: Make sure a container's subpath can be set using an
			expansion of environment variables.
	*/
	It("should allow substituting values in a volume subpath [Feature:VolumeSubpathEnvExpansion][NodeAlphaFeature:VolumeSubpathEnvExpansion]", func() {
		podName := "var-expansion-" + string(uuid.NewUUID())
		pod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:   podName,
				Labels: map[string]string{"name": podName},
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:    "dapi-container",
						Image:   imageutils.GetE2EImage(imageutils.BusyBox),
						Command: []string{"sh", "-c", "test -d /testcontainer/" + podName + ";echo $?"},
						Env: []v1.EnvVar{
							{
								Name:  "POD_NAME",
								Value: podName,
							},
						},
						VolumeMounts: []v1.VolumeMount{
							{
								Name:        "workdir1",
								MountPath:   "/logscontainer",
								SubPathExpr: "$(POD_NAME)",
							},
							{
								Name:      "workdir2",
								MountPath: "/testcontainer",
							},
						},
					},
				},
				RestartPolicy: v1.RestartPolicyNever,
				Volumes: []v1.Volume{
					{
						Name: "workdir1",
						VolumeSource: v1.VolumeSource{
							HostPath: &v1.HostPathVolumeSource{Path: "/tmp"},
						},
					},
					{
						Name: "workdir2",
						VolumeSource: v1.VolumeSource{
							HostPath: &v1.HostPathVolumeSource{Path: "/tmp"},
						},
					},
				},
			},
		}

		f.TestContainerOutput("substitution in volume subpath", pod, 0, []string{
			"0",
		})
	})

	/*
		    Testname: var-expansion-subpath-with-backticks
		    Description: Make sure a container's subpath can not be set using an
			expansion of environment variables when backticks are supplied.
	*/
	It("should fail substituting values in a volume subpath with backticks [Feature:VolumeSubpathEnvExpansion][NodeAlphaFeature:VolumeSubpathEnvExpansion][Slow]", func() {

		podName := "var-expansion-" + string(uuid.NewUUID())
		pod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:   podName,
				Labels: map[string]string{"name": podName},
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:  "dapi-container",
						Image: imageutils.GetE2EImage(imageutils.BusyBox),
						Env: []v1.EnvVar{
							{
								Name:  "POD_NAME",
								Value: "..",
							},
						},
						VolumeMounts: []v1.VolumeMount{
							{
								Name:        "workdir1",
								MountPath:   "/logscontainer",
								SubPathExpr: "$(POD_NAME)",
							},
						},
					},
				},
				RestartPolicy: v1.RestartPolicyNever,
				Volumes: []v1.Volume{
					{
						Name: "workdir1",
						VolumeSource: v1.VolumeSource{
							EmptyDir: &v1.EmptyDirVolumeSource{},
						},
					},
				},
			},
		}

		// Pod should fail
		testPodFailSubpath(f, pod)
	})

	/*
		    Testname: var-expansion-subpath-with-absolute-path
		    Description: Make sure a container's subpath can not be set using an
			expansion of environment variables when absolute path is supplied.
	*/
	It("should fail substituting values in a volume subpath with absolute path [Feature:VolumeSubpathEnvExpansion][NodeAlphaFeature:VolumeSubpathEnvExpansion][Slow]", func() {

		podName := "var-expansion-" + string(uuid.NewUUID())
		pod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:   podName,
				Labels: map[string]string{"name": podName},
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:  "dapi-container",
						Image: imageutils.GetE2EImage(imageutils.BusyBox),
						Env: []v1.EnvVar{
							{
								Name:  "POD_NAME",
								Value: "/tmp",
							},
						},
						VolumeMounts: []v1.VolumeMount{
							{
								Name:        "workdir1",
								MountPath:   "/logscontainer",
								SubPathExpr: "$(POD_NAME)",
							},
						},
					},
				},
				RestartPolicy: v1.RestartPolicyNever,
				Volumes: []v1.Volume{
					{
						Name: "workdir1",
						VolumeSource: v1.VolumeSource{
							EmptyDir: &v1.EmptyDirVolumeSource{},
						},
					},
				},
			},
		}

		// Pod should fail
		testPodFailSubpath(f, pod)
	})

	/*
	   Testname: var-expansion-subpath-ready-from-failed-state
	   Description: Verify that a failing subpath expansion can be modified during the lifecycle of a container.
	*/
	It("should verify that a failing subpath expansion can be modified during the lifecycle of a container [Feature:VolumeSubpathEnvExpansion][NodeAlphaFeature:VolumeSubpathEnvExpansion][Slow]", func() {

		podName := "var-expansion-" + string(uuid.NewUUID())
		containerName := "dapi-container"
		pod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:        podName,
				Labels:      map[string]string{"name": podName},
				Annotations: map[string]string{"notmysubpath": "mypath"},
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:    containerName,
						Image:   imageutils.GetE2EImage(imageutils.BusyBox),
						Command: []string{"sh", "-c", "tail -f /dev/null"},
						Env: []v1.EnvVar{
							{
								Name:  "POD_NAME",
								Value: "foo",
							},
							{
								Name: "ANNOTATION",
								ValueFrom: &v1.EnvVarSource{
									FieldRef: &v1.ObjectFieldSelector{
										APIVersion: "v1",
										FieldPath:  "metadata.annotations['mysubpath']",
									},
								},
							},
						},
						VolumeMounts: []v1.VolumeMount{
							{
								Name:        "workdir1",
								MountPath:   "/subpath_mount",
								SubPathExpr: "$(ANNOTATION)/$(POD_NAME)",
							},
							{
								Name:      "workdir2",
								MountPath: "/volume_mount",
							},
						},
					},
				},
				Volumes: []v1.Volume{
					{
						Name: "workdir1",
						VolumeSource: v1.VolumeSource{
							HostPath: &v1.HostPathVolumeSource{Path: "/tmp"},
						},
					},
					{
						Name: "workdir2",
						VolumeSource: v1.VolumeSource{
							HostPath: &v1.HostPathVolumeSource{Path: "/tmp"},
						},
					},
				},
			},
		}

		By("creating the pod with failed condition")
		pod, err := f.ClientSet.CoreV1().Pods(f.Namespace.Name).Create(pod)
		Expect(err).ToNot(HaveOccurred(), "while creating pod")

		err = framework.WaitTimeoutForPodRunningInNamespace(f.ClientSet, pod.Name, pod.Namespace, framework.PodStartShortTimeout)
		Expect(err).To(HaveOccurred(), "while waiting for pod to be running")

		var podClient *framework.PodClient
		podClient = f.PodClient()

		By("updating the pod")
		podClient.Update(podName, func(pod *v1.Pod) {
			pod.ObjectMeta.Annotations = map[string]string{"mysubpath": "mypath"}
		})

		By("waiting for pod running")
		err = framework.WaitTimeoutForPodRunningInNamespace(f.ClientSet, pod.Name, pod.Namespace, framework.PodStartShortTimeout)
		Expect(err).NotTo(HaveOccurred(), "while waiting for pod to be running")

		By("deleting the pod gracefully")
		err = framework.DeletePodWithWait(f, f.ClientSet, pod)
		Expect(err).NotTo(HaveOccurred(), "failed to delete pod")
	})

	/*
		    Testname: var-expansion-subpath-test-writes
		    Description: Verify that a subpath expansion can be used to write files into subpaths.
			1.	valid subpathexpr starts a container running
			2.	test for valid subpath writes
			3.	successful expansion of the subpathexpr isn't required for volume cleanup

	*/
	It("should succeed in writing subpaths in container [Feature:VolumeSubpathEnvExpansion][NodeAlphaFeature:VolumeSubpathEnvExpansion][Slow]", func() {

		podName := "var-expansion-" + string(uuid.NewUUID())
		containerName := "dapi-container"
		pod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:        podName,
				Labels:      map[string]string{"name": podName},
				Annotations: map[string]string{"mysubpath": "mypath"},
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:    containerName,
						Image:   imageutils.GetE2EImage(imageutils.BusyBox),
						Command: []string{"sh", "-c", "tail -f /dev/null"},
						Env: []v1.EnvVar{
							{
								Name:  "POD_NAME",
								Value: "foo",
							},
							{
								Name: "ANNOTATION",
								ValueFrom: &v1.EnvVarSource{
									FieldRef: &v1.ObjectFieldSelector{
										APIVersion: "v1",
										FieldPath:  "metadata.annotations['mysubpath']",
									},
								},
							},
						},
						VolumeMounts: []v1.VolumeMount{
							{
								Name:        "workdir1",
								MountPath:   "/subpath_mount",
								SubPathExpr: "$(ANNOTATION)/$(POD_NAME)",
							},
							{
								Name:      "workdir2",
								MountPath: "/volume_mount",
							},
						},
					},
				},
				RestartPolicy: v1.RestartPolicyNever,
				Volumes: []v1.Volume{
					{
						Name: "workdir1",
						VolumeSource: v1.VolumeSource{
							HostPath: &v1.HostPathVolumeSource{Path: "/tmp"},
						},
					},
					{
						Name: "workdir2",
						VolumeSource: v1.VolumeSource{
							HostPath: &v1.HostPathVolumeSource{Path: "/tmp"},
						},
					},
				},
			},
		}

		By("creating the pod")
		pod, err := f.ClientSet.CoreV1().Pods(f.Namespace.Name).Create(pod)

		By("waiting for pod running")
		err = framework.WaitTimeoutForPodRunningInNamespace(f.ClientSet, pod.Name, pod.Namespace, framework.PodStartShortTimeout)
		Expect(err).NotTo(HaveOccurred(), "while waiting for pod to be running")

		By("creating a file in subpath")
		cmd := "touch /volume_mount/mypath/foo/test.log"
		_, err = framework.RunHostCmd(pod.Namespace, pod.Name, cmd)
		if err != nil {
			framework.Failf("expected to be able to write to subpath")
		}

		By("test for file in mounted path")
		cmd = "test -f /subpath_mount/test.log"
		_, err = framework.RunHostCmd(pod.Namespace, pod.Name, cmd)
		if err != nil {
			framework.Failf("expected to be able to verify file")
		}

		By("updating the annotation value")
		var podClient *framework.PodClient
		podClient = f.PodClient()

		podClient.Update(podName, func(pod *v1.Pod) {
			pod.ObjectMeta.Annotations["mysubpath"] = "mynewpath"
		})

		By("waiting for annotated pod running")
		err = framework.WaitTimeoutForPodRunningInNamespace(f.ClientSet, pod.Name, pod.Namespace, framework.PodStartShortTimeout)
		Expect(err).NotTo(HaveOccurred(), "while waiting for annotated pod to be running")

		By("deleting the pod gracefully")
		err = framework.DeletePodWithWait(f, f.ClientSet, pod)
		Expect(err).NotTo(HaveOccurred(), "failed to delete pod")
	})

	/*
		    Testname: var-expansion-subpath-lifecycle
		    Description: Verify should not change the subpath mount on a container restart if the environment variable changes
			1.	valid subpathexpr starts a container running
			2.	test for valid subpath writes
			3.	container restarts
			4.	delete cleanly

	*/

	It("should not change the subpath mount on a container restart if the environment variable changes [Feature:VolumeSubpathEnvExpansion][NodeAlphaFeature:VolumeSubpathEnvExpansion][Slow]", func() {

		suffix := string(uuid.NewUUID())
		podName := fmt.Sprintf("var-expansion-%s", suffix)
		containerName := "dapi-container"
		pod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:        podName,
				Labels:      map[string]string{"name": podName},
				Annotations: map[string]string{"mysubpath": "foo"},
			},
			Spec: v1.PodSpec{
				InitContainers: []v1.Container{
					{
						Name:    fmt.Sprintf("init-volume-%s", suffix),
						Image:   imageutils.GetE2EImage(imageutils.BusyBox),
						Command: []string{"sh", "-c", "mkdir -p /volume_mount/foo; touch /volume_mount/foo/test.log"},
						VolumeMounts: []v1.VolumeMount{
							{
								Name:      "workdir1",
								MountPath: "/subpath_mount",
							},
							{
								Name:      "workdir2",
								MountPath: "/volume_mount",
							},
						},
					},
				},

				Containers: []v1.Container{
					{
						Name:    containerName,
						Image:   imageutils.GetE2EImage(imageutils.BusyBox),
						Command: []string{"/bin/sh", "-ec", "sleep 100000"},
						Env: []v1.EnvVar{
							{
								Name: "POD_NAME",
								ValueFrom: &v1.EnvVarSource{
									FieldRef: &v1.ObjectFieldSelector{
										APIVersion: "v1",
										FieldPath:  "metadata.annotations['mysubpath']",
									},
								},
							},
						},
						VolumeMounts: []v1.VolumeMount{
							{
								Name:        "workdir1",
								MountPath:   "/subpath_mount",
								SubPathExpr: "$(POD_NAME)",
							},
							{
								Name:      "workdir2",
								MountPath: "/volume_mount",
							},
						},
					},
				},
				RestartPolicy: v1.RestartPolicyOnFailure,
				Volumes: []v1.Volume{
					{
						Name: "workdir1",
						VolumeSource: v1.VolumeSource{
							HostPath: &v1.HostPathVolumeSource{Path: "/tmp"},
						},
					},
					{
						Name: "workdir2",
						VolumeSource: v1.VolumeSource{
							HostPath: &v1.HostPathVolumeSource{Path: "/tmp"},
						},
					},
				},
			},
		}

		// Add liveness probe to subpath container
		pod.Spec.Containers[0].LivenessProbe = &v1.Probe{
			Handler: v1.Handler{
				Exec: &v1.ExecAction{

					Command: []string{"cat", "/subpath_mount/test.log"},
				},
			},
			InitialDelaySeconds: 1,
			FailureThreshold:    1,
			PeriodSeconds:       2,
		}

		// Start pod
		By(fmt.Sprintf("Creating pod %s", pod.Name))
		pod, err := f.ClientSet.CoreV1().Pods(f.Namespace.Name).Create(pod)
		Expect(err).ToNot(HaveOccurred(), "while creating pod")
		defer func() {
			framework.DeletePodWithWait(f, f.ClientSet, pod)
		}()
		err = framework.WaitForPodRunningInNamespace(f.ClientSet, pod)
		Expect(err).ToNot(HaveOccurred(), "while waiting for pod to be running")

		var podClient *framework.PodClient
		podClient = f.PodClient()

		By("updating the pod")
		podClient.Update(podName, func(pod *v1.Pod) {
			pod.ObjectMeta.Annotations = map[string]string{"mysubpath": "newsubpath"}
		})

		By("waiting for pod and container restart")
		waitForPodContainerRestart(f, pod, "/volume_mount/foo/test.log")

		By("test for subpath mounted with old value")
		cmd := "test -f /volume_mount/foo/test.log"
		_, err = framework.RunHostCmd(pod.Namespace, pod.Name, cmd)
		if err != nil {
			framework.Failf("expected to be able to verify old file exists")
		}

		cmd = "test ! -f /volume_mount/newsubpath/test.log"
		_, err = framework.RunHostCmd(pod.Namespace, pod.Name, cmd)
		if err != nil {
			framework.Failf("expected to be able to verify new file does not exist")
		}
	})
})

func testPodFailSubpath(f *framework.Framework, pod *v1.Pod) {

	pod, err := f.ClientSet.CoreV1().Pods(f.Namespace.Name).Create(pod)
	Expect(err).ToNot(HaveOccurred(), "while creating pod")

	defer func() {
		framework.DeletePodWithWait(f, f.ClientSet, pod)
	}()

	err = framework.WaitTimeoutForPodRunningInNamespace(f.ClientSet, pod.Name, pod.Namespace, framework.PodStartShortTimeout)
	Expect(err).To(HaveOccurred(), "while waiting for pod to be running")
}

// Tests that the existing subpath mount is detected when a container restarts
func waitForPodContainerRestart(f *framework.Framework, pod *v1.Pod, volumeMount string) {

	By("Failing liveness probe")
	out, err := framework.RunKubectl("exec", fmt.Sprintf("--namespace=%s", pod.Namespace), pod.Name, "--container", pod.Spec.Containers[0].Name, "--", "/bin/sh", "-c", fmt.Sprintf("rm %v", volumeMount))

	framework.Logf("Pod exec output: %v", out)
	Expect(err).ToNot(HaveOccurred(), "while failing liveness probe")

	// Check that container has restarted
	By("Waiting for container to restart")
	restarts := int32(0)
	err = wait.PollImmediate(10*time.Second, 2*time.Minute, func() (bool, error) {
		pod, err := f.ClientSet.CoreV1().Pods(f.Namespace.Name).Get(pod.Name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		for _, status := range pod.Status.ContainerStatuses {
			if status.Name == pod.Spec.Containers[0].Name {
				framework.Logf("Container %v, restarts: %v", status.Name, status.RestartCount)
				restarts = status.RestartCount
				if restarts > 0 {
					framework.Logf("Container has restart count: %v", restarts)
					return true, nil
				}
			}
		}
		return false, nil
	})
	Expect(err).ToNot(HaveOccurred(), "while waiting for container to restart")

	// Fix liveness probe
	By("Rewriting the file")
	out, err = framework.RunKubectl("exec", fmt.Sprintf("--namespace=%s", pod.Namespace), pod.Name, "--container", pod.Spec.Containers[0].Name, "--", "/bin/sh", "-c", fmt.Sprintf("echo test-after > %v", volumeMount))
	framework.Logf("Pod exec output: %v", out)
	Expect(err).ToNot(HaveOccurred(), "while rewriting the probe file")

	// Wait for container restarts to stabilize
	By("Waiting for container to stop restarting")
	stableCount := int(0)
	stableThreshold := int(time.Minute / framework.Poll)
	err = wait.PollImmediate(framework.Poll, 2*time.Minute, func() (bool, error) {
		pod, err := f.ClientSet.CoreV1().Pods(f.Namespace.Name).Get(pod.Name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		for _, status := range pod.Status.ContainerStatuses {
			if status.Name == pod.Spec.Containers[0].Name {
				if status.RestartCount == restarts {
					stableCount++
					if stableCount > stableThreshold {
						framework.Logf("Container restart has stabilized")
						return true, nil
					}
				} else {
					restarts = status.RestartCount
					stableCount = 0
					framework.Logf("Container has restart count: %v", restarts)
				}
				break
			}
		}
		return false, nil
	})
	Expect(err).ToNot(HaveOccurred(), "while waiting for container to stabilize")
}
