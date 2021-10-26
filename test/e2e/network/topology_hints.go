/*
Copyright 2015 The Kubernetes Authors.

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

package network

import (
	"context"
	"fmt"
	"time"

	"github.com/onsi/ginkgo"
	v1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/test/e2e/framework"
	e2edaemonset "k8s.io/kubernetes/test/e2e/framework/daemonset"
	e2epod "k8s.io/kubernetes/test/e2e/framework/pod"
	e2eservice "k8s.io/kubernetes/test/e2e/framework/service"
	e2eskipper "k8s.io/kubernetes/test/e2e/framework/skipper"
	"k8s.io/kubernetes/test/e2e/network/common"
	imageutils "k8s.io/kubernetes/test/utils/image"
)

var _ = common.SIGDescribe("Feature:Topology Hints", func() {
	f := framework.NewDefaultFramework("topology-hints")

	// filled in BeforeEach
	var c clientset.Interface

	ginkgo.BeforeEach(func() {
		c = f.ClientSet
		e2eskipper.SkipUnlessMultizone(c)
	})

	ginkgo.It("should distribute endpoints evenly", func() {
		portNum := 9376
		thLabels := map[string]string{labelKey: clientLabelValue}
		img := imageutils.GetE2EImage(imageutils.Agnhost)
		ports := []v1.ContainerPort{{ContainerPort: int32(portNum)}}
		dsConf := e2edaemonset.NewDaemonSet("topology-serve-hostname", img, thLabels, nil, nil, ports, "serve-hostname")
		ds, err := c.AppsV1().DaemonSets(f.Namespace.Name).Create(context.TODO(), dsConf, metav1.CreateOptions{})
		framework.ExpectNoError(err, "error creating DaemonSet")

		svc := createServiceReportErr(c, f.Namespace.Name, &v1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name: "topology-hints",
				Annotations: map[string]string{
					v1.AnnotationTopologyAwareHints: "Auto",
				},
			},
			Spec: v1.ServiceSpec{
				Selector:                 thLabels,
				PublishNotReadyAddresses: true,
				Ports: []v1.ServicePort{{
					Name:       "example",
					Port:       80,
					TargetPort: intstr.FromInt(portNum),
					Protocol:   v1.ProtocolTCP,
				}},
			},
		})

		err = wait.Poll(5*time.Second, framework.PodStartTimeout, func() (bool, error) {
			return e2edaemonset.CheckRunningOnAllNodes(f, ds)
		})
		framework.ExpectNoError(err, "timed out waiting for DaemonSets to be ready")

		nodeNames := e2edaemonset.SchedulableNodes(c, ds)
		framework.Logf("Waiting for %d endpoints to be tracked in EndpointSlices", len(nodeNames))

		var finalSlices []discoveryv1.EndpointSlice
		err = wait.Poll(5*time.Second, 3*time.Minute, func() (bool, error) {
			slices, listErr := c.DiscoveryV1().EndpointSlices(f.Namespace.Name).List(context.TODO(), metav1.ListOptions{LabelSelector: fmt.Sprintf("%s=%s", discoveryv1.LabelServiceName, svc.Name)})
			if listErr != nil {
				return false, listErr
			}

			numEndpoints := 0
			for _, slice := range slices.Items {
				numEndpoints += len(slice.Endpoints)
			}
			if len(nodeNames) > numEndpoints {
				framework.Logf("Expected %d endpoints, got %d", len(nodeNames), numEndpoints)
				return false, nil
			}

			finalSlices = slices.Items
			return true, nil
		})
		framework.ExpectNoError(err, "timed out waiting for EndpointSlices to be ready")

		ginkgo.By("having hints set for each endpoint")
		for _, slice := range finalSlices {
			for _, ep := range slice.Endpoints {
				if ep.Zone == nil {
					framework.Failf("Expected endpoint in %s to have zone: %v", slice.Name, ep)
				}
				if ep.Hints == nil || len(ep.Hints.ForZones) == 0 {
					framework.Failf("Expected endpoint in %s to have hints: %v", slice.Name, ep)
				}
				if len(ep.Hints.ForZones) > 1 {
					framework.Failf("Expected endpoint in %s to have exactly 1 zone hint, got %d: %v", slice.Name, len(ep.Hints.ForZones), ep)
				}
				if *ep.Zone != ep.Hints.ForZones[0].Name {
					framework.Failf("Expected endpoint in %s to have same zone hint, got %s: %v", slice.Name, *ep.Zone, ep)
				}
			}
		}

		nodeList, err := c.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
		framework.ExpectNoError(err)
		nodesByZone := map[string]string{}
		for _, node := range nodeList.Items {
			if zone, ok := node.Labels[v1.LabelTopologyZone]; ok {
				nodesByZone[node.Name] = zone
			}
		}

		podList, err := c.CoreV1().Pods(f.Namespace.Name).List(context.TODO(), metav1.ListOptions{})
		framework.ExpectNoError(err)
		podsByZone := map[string]string{}
		for _, pod := range podList.Items {
			if zone, ok := nodesByZone[pod.Spec.NodeName]; ok {
				podsByZone[pod.Name] = zone
			}
		}

		ginkgo.By("keeping requests in the same zone")
		for i, nodeName := range nodeNames {
			// Iterate through max of 3 nodes
			if i > 2 {
				break
			}
			fromZone, ok := nodesByZone[nodeName]
			if !ok {
				framework.Failf("Expected zone to be specified for %s node", nodeName)
			}
			execPod := e2epod.CreateExecPodOrFail(f.ClientSet, f.Namespace.Name, "execpod-", func(pod *v1.Pod) {
				pod.Spec.NodeName = nodeName
			})
			framework.Logf("Ensuring that requests from %s pod on %s node stay in %s zone", execPod.Name, nodeName, fromZone)

			cmd := fmt.Sprintf("curl -q -s --connect-timeout 2 http://%s:80/", svc.Name)
			consecutiveRequests := 0
			var stdout string
			if pollErr := wait.PollImmediate(framework.Poll, e2eservice.KubeProxyLagTimeout, func() (bool, error) {
				var err error
				stdout, err = framework.RunHostCmd(f.Namespace.Name, execPod.Name, cmd)
				if err != nil {
					framework.Logf("unexpected error running %s from %s", cmd, execPod.Name)
					return false, nil
				}
				destZone, ok := podsByZone[stdout]
				if !ok {
					framework.Logf("could not determine dest zone from output: %s", stdout)
					return false, nil
				}
				if fromZone != destZone {
					framework.Logf("expected request from %s to stay in %s zone, delivered to %s zone", execPod.Name, fromZone, destZone)
					return false, nil
				}
				consecutiveRequests++
				if consecutiveRequests < 5 {
					return false, nil
				}
				return true, nil
			}); pollErr != nil {
				framework.Failf("expected 5 consecutive requests from %s to stay in %s zone %v within %v, stdout: %v", execPod.Name, fromZone, e2eservice.KubeProxyLagTimeout, stdout)
			}
		}
	})
})
