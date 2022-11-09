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

package node

import (
	"context"
	"fmt"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	featuregatetesting "k8s.io/component-base/featuregate/testing"
	"k8s.io/klog/v2"
	podutil "k8s.io/kubernetes/pkg/api/v1/pod"
	"k8s.io/kubernetes/pkg/controller/nodelifecycle"
	"k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/plugin/pkg/admission/defaulttolerationseconds"
	"k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction"
	pluginapi "k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction"
	testutils "k8s.io/kubernetes/test/integration/util"
	imageutils "k8s.io/kubernetes/test/utils/image"
)

// TestEvictionForNoExecuteTaintAddedByUser tests taint-based eviction for a node tainted NoExecute
func TestEvictionForNoExecuteTaintAddedByUser(t *testing.T) {
	// we need at least 2 nodes to prevent lifecycle manager from entering "fully-disrupted" mode
	nodeCount := 3
	nodeIndex := 1 // the exact node doesn't matter, pick one

	tests := map[string]struct {
		enablePodDisruptionConditions bool
	}{
		"Test eviciton for NoExecute taint added by user; pod condition added when PodDisruptionConditions enabled": {
			enablePodDisruptionConditions: true,
		},
		"Test eviciton for NoExecute taint added by user; no pod condition added when PodDisruptionConditions disabled": {
			enablePodDisruptionConditions: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var nodes []*v1.Node
			for i := 0; i < nodeCount; i++ {
				node := &v1.Node{
					ObjectMeta: metav1.ObjectMeta{
						Name:   fmt.Sprintf("testnode-%d", i),
						Labels: map[string]string{"node.kubernetes.io/exclude-disruption": "true"},
					},
					Spec: v1.NodeSpec{},
					Status: v1.NodeStatus{
						Conditions: []v1.NodeCondition{
							{
								Type:   v1.NodeReady,
								Status: v1.ConditionTrue,
							},
						},
					},
				}
				nodes = append(nodes, node)
			}
			testPod := &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "testpod",
				},
				Spec: v1.PodSpec{
					NodeName: nodes[nodeIndex].Name,
					Containers: []v1.Container{
						{Name: "container", Image: imageutils.GetPauseImageName()},
					},
				},
				Status: v1.PodStatus{
					Phase: v1.PodRunning,
					Conditions: []v1.PodCondition{
						{
							Type:   v1.PodReady,
							Status: v1.ConditionTrue,
						},
					},
				},
			}

			defer featuregatetesting.SetFeatureGateDuringTest(t, feature.DefaultFeatureGate, features.PodDisruptionConditions, test.enablePodDisruptionConditions)()
			testCtx := testutils.InitTestAPIServer(t, "taint-no-execute", nil)

			// Build clientset and informers for controllers.
			defer testutils.CleanupTest(t, testCtx)
			cs := testCtx.ClientSet

			// Build clientset and informers for controllers.
			externalClientConfig := restclient.CopyConfig(testCtx.KubeConfig)
			externalClientConfig.QPS = -1
			externalClientset := clientset.NewForConfigOrDie(externalClientConfig)
			externalInformers := informers.NewSharedInformerFactory(externalClientset, time.Second)

			// Start NodeLifecycleController for taint.
			nc, err := nodelifecycle.NewNodeLifecycleController(
				testCtx.Ctx,
				externalInformers.Coordination().V1().Leases(),
				externalInformers.Core().V1().Pods(),
				externalInformers.Core().V1().Nodes(),
				externalInformers.Apps().V1().DaemonSets(),
				cs,
				1*time.Second,    // Node monitor grace period
				time.Minute,      // Node startup grace period
				time.Millisecond, // Node monitor period
				1,                // Pod eviction timeout
				100,              // Eviction limiter QPS
				100,              // Secondary eviction limiter QPS
				50,               // Large cluster threshold
				0.55,             // Unhealthy zone threshold
				true,             // Run taint manager
			)
			if err != nil {
				t.Fatalf("Failed to create node controller: %v", err)
			}

			// Waiting for all controllers to sync
			externalInformers.Start(testCtx.Ctx.Done())
			externalInformers.WaitForCacheSync(testCtx.Ctx.Done())

			// Run all controllers
			go nc.Run(testCtx.Ctx)

			for index := range nodes {
				nodes[index], err = cs.CoreV1().Nodes().Create(testCtx.Ctx, nodes[index], metav1.CreateOptions{})
				if err != nil {
					t.Fatalf("Failed to create node, err: %v", err)
				}
			}

			testPod, err = cs.CoreV1().Pods(testCtx.NS.Name).Create(testCtx.Ctx, testPod, metav1.CreateOptions{})
			if err != nil {
				t.Fatalf("Test Failed: error: %v, while creating pod", err)
			}

			if err := testutils.AddTaintToNode(cs, nodes[nodeIndex].Name, v1.Taint{Key: "CustomTaintByUser", Effect: v1.TaintEffectNoExecute}); err != nil {
				t.Errorf("Failed to taint node in test %s <%s>, err: %v", name, nodes[nodeIndex].Name, err)
			}

			err = wait.PollImmediate(time.Second, time.Second*20, testutils.PodIsGettingEvicted(cs, testPod.Namespace, testPod.Name))
			if err != nil {
				t.Fatalf("Error %q in test %q when waiting for terminating pod: %q", err, name, klog.KObj(testPod))
			}
			testPod, err = cs.CoreV1().Pods(testCtx.NS.Name).Get(testCtx.Ctx, testPod.Name, metav1.GetOptions{})
			if err != nil {
				t.Fatalf("Test Failed: error: %q, while getting updated pod", err)
			}
			_, cond := podutil.GetPodCondition(&testPod.Status, v1.DisruptionTarget)
			if test.enablePodDisruptionConditions == true && cond == nil {
				t.Errorf("Pod %q does not have the expected condition: %q", klog.KObj(testPod), v1.DisruptionTarget)
			} else if test.enablePodDisruptionConditions == false && cond != nil {
				t.Errorf("Pod %q has an unexpected condition: %q", klog.KObj(testPod), v1.DisruptionTarget)
			}
		})
	}
}

// TestTaintBasedEvictions tests related cases for the TaintBasedEvictions feature
func TestTaintBasedEvictions(t *testing.T) {
	// we need at least 2 nodes to prevent lifecycle manager from entering "fully-disrupted" mode
	nodeCount := 3
	nodeIndex := 1 // the exact node doesn't matter, pick one
	zero := int64(0)
	gracePeriod := int64(1)
	testPod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "testpod1", DeletionGracePeriodSeconds: &zero},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{Name: "container", Image: imageutils.GetPauseImageName()},
			},
			Tolerations: []v1.Toleration{
				{
					Key:      v1.TaintNodeNotReady,
					Operator: v1.TolerationOpExists,
					Effect:   v1.TaintEffectNoExecute,
				},
			},
			TerminationGracePeriodSeconds: &gracePeriod,
		},
	}
	tests := []struct {
		name                        string
		nodeTaints                  []v1.Taint
		nodeConditions              []v1.NodeCondition
		pod                         *v1.Pod
		tolerationSeconds           int64
		expectedWaitForPodCondition string
	}{
		{
			name:                        "Taint based evictions for NodeNotReady and 200 tolerationseconds",
			nodeTaints:                  []v1.Taint{{Key: v1.TaintNodeNotReady, Effect: v1.TaintEffectNoExecute}},
			nodeConditions:              []v1.NodeCondition{{Type: v1.NodeReady, Status: v1.ConditionFalse}},
			pod:                         testPod.DeepCopy(),
			tolerationSeconds:           200,
			expectedWaitForPodCondition: "updated with tolerationSeconds of 200",
		},
		{
			name:           "Taint based evictions for NodeNotReady with no pod tolerations",
			nodeTaints:     []v1.Taint{{Key: v1.TaintNodeNotReady, Effect: v1.TaintEffectNoExecute}},
			nodeConditions: []v1.NodeCondition{{Type: v1.NodeReady, Status: v1.ConditionFalse}},
			pod: &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{Name: "testpod1"},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{Name: "container", Image: imageutils.GetPauseImageName()},
					},
				},
			},
			tolerationSeconds:           300,
			expectedWaitForPodCondition: "updated with tolerationSeconds=300",
		},
		{
			name:                        "Taint based evictions for NodeNotReady and 0 tolerationseconds",
			nodeTaints:                  []v1.Taint{{Key: v1.TaintNodeNotReady, Effect: v1.TaintEffectNoExecute}},
			nodeConditions:              []v1.NodeCondition{{Type: v1.NodeReady, Status: v1.ConditionFalse}},
			pod:                         testPod.DeepCopy(),
			tolerationSeconds:           0,
			expectedWaitForPodCondition: "terminating",
		},
		{
			name:           "Taint based evictions for NodeUnreachable",
			nodeTaints:     []v1.Taint{{Key: v1.TaintNodeUnreachable, Effect: v1.TaintEffectNoExecute}},
			nodeConditions: []v1.NodeCondition{{Type: v1.NodeReady, Status: v1.ConditionUnknown}},
		},
	}

	// Build admission chain handler.
	podTolerations := podtolerationrestriction.NewPodTolerationsPlugin(&pluginapi.Configuration{})
	admission := admission.NewChainHandler(
		podTolerations,
		defaulttolerationseconds.NewDefaultTolerationSeconds(),
	)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testCtx := testutils.InitTestAPIServer(t, "taint-based-evictions", admission)

			// Build clientset and informers for controllers.
			externalClientConfig := restclient.CopyConfig(testCtx.KubeConfig)
			externalClientConfig.QPS = -1
			externalClientset := clientset.NewForConfigOrDie(externalClientConfig)
			externalInformers := informers.NewSharedInformerFactory(externalClientset, time.Second)
			podTolerations.SetExternalKubeClientSet(externalClientset)
			podTolerations.SetExternalKubeInformerFactory(externalInformers)

			defer testutils.CleanupTest(t, testCtx)
			cs := testCtx.ClientSet

			// Start NodeLifecycleController for taint.
			nc, err := nodelifecycle.NewNodeLifecycleController(
				testCtx.Ctx,
				externalInformers.Coordination().V1().Leases(),
				externalInformers.Core().V1().Pods(),
				externalInformers.Core().V1().Nodes(),
				externalInformers.Apps().V1().DaemonSets(),
				cs,
				1*time.Second,    // Node monitor grace period
				time.Minute,      // Node startup grace period
				time.Millisecond, // Node monitor period
				time.Second,      // Pod eviction timeout
				100,              // Eviction limiter QPS
				100,              // Secondary eviction limiter QPS
				50,               // Large cluster threshold
				0.55,             // Unhealthy zone threshold
				true,             // Run taint manager
			)
			if err != nil {
				t.Fatalf("Failed to create node controller: %v", err)
			}

			// Waiting for all controllers to sync
			externalInformers.Start(testCtx.Ctx.Done())
			externalInformers.WaitForCacheSync(testCtx.Ctx.Done())

			// Run the controller
			go nc.Run(testCtx.Ctx)

			nodeRes := v1.ResourceList{
				v1.ResourceCPU:    resource.MustParse("4000m"),
				v1.ResourceMemory: resource.MustParse("16Gi"),
				v1.ResourcePods:   resource.MustParse("110"),
			}

			var nodes []*v1.Node
			for i := 0; i < nodeCount; i++ {
				node := &v1.Node{
					ObjectMeta: metav1.ObjectMeta{
						Name: fmt.Sprintf("node-%d", i),
						Labels: map[string]string{
							v1.LabelTopologyRegion:                  "region1",
							v1.LabelTopologyZone:                    "zone1",
							"node.kubernetes.io/exclude-disruption": "true",
						},
					},
					Spec: v1.NodeSpec{},
					Status: v1.NodeStatus{
						Capacity:    nodeRes,
						Allocatable: nodeRes,
					},
				}
				if i == nodeIndex {
					node.Status.Conditions = append(node.Status.Conditions, test.nodeConditions...)
				} else {
					node.Status.Conditions = append(node.Status.Conditions, v1.NodeCondition{
						Type:   v1.NodeReady,
						Status: v1.ConditionTrue,
					})
				}
				nodes = append(nodes, node)
				if _, err := cs.CoreV1().Nodes().Create(context.TODO(), node, metav1.CreateOptions{}); err != nil {
					t.Fatalf("Failed to create node: %q, err: %v", klog.KObj(node), err)
				}
			}

			if test.pod != nil {
				test.pod.Spec.NodeName = nodes[nodeIndex].Name
				test.pod.Name = "testpod"
				if len(test.pod.Spec.Tolerations) > 0 {
					test.pod.Spec.Tolerations[0].TolerationSeconds = &test.tolerationSeconds
				}

				test.pod, err = cs.CoreV1().Pods(testCtx.NS.Name).Create(context.TODO(), test.pod, metav1.CreateOptions{})
				if err != nil {
					t.Fatalf("Test Failed: error: %q, while creating pod %q", err, klog.KObj(test.pod))
				}
			}

			if err := testutils.WaitForNodeTaints(cs, nodes[nodeIndex], test.nodeTaints); err != nil {
				t.Errorf("Failed to taint node %q, err: %v", klog.KObj(nodes[nodeIndex]), err)
			}

			if test.pod != nil {
				err = wait.PollImmediate(time.Second, time.Second*15, func() (bool, error) {
					pod, err := cs.CoreV1().Pods(test.pod.Namespace).Get(context.TODO(), test.pod.Name, metav1.GetOptions{})
					if err != nil {
						return false, err
					}
					// as node is unreachable, pod0 is expected to be in Terminating status
					// rather than getting deleted
					if test.tolerationSeconds == 0 {
						return pod.DeletionTimestamp != nil, nil
					}
					if seconds, err := testutils.GetTolerationSeconds(pod.Spec.Tolerations); err == nil {
						return seconds == test.tolerationSeconds, nil
					}
					return false, nil
				})
				if err != nil {
					pod, _ := cs.CoreV1().Pods(testCtx.NS.Name).Get(context.TODO(), test.pod.Name, metav1.GetOptions{})
					t.Fatalf("Error: %v, Expected test pod to be %s but it's %v", err, test.expectedWaitForPodCondition, pod)
				}
				testutils.CleanupPods(cs, t, []*v1.Pod{test.pod})
			}
			testutils.CleanupNodes(cs, t)
		})
	}
}
