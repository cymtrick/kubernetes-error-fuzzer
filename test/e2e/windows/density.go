/*
Copyright 2018 The Kubernetes Authors.

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

package windows

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	stats "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"
	"k8s.io/kubernetes/test/e2e/framework"
	imageutils "k8s.io/kubernetes/test/utils/image"
	"k8s.io/apimachinery/pkg/util/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = SIGDescribe("Density [Serial] [Slow]", func() {

	f := framework.NewDefaultFramework("density-test-windows")

	Context("create a batch of pods", func() {
		// TODO(coufon): the values are generous, set more precise limits with benchmark data
		// and add more tests
		dTests := []densityTest{
			{
				podsNr:   10,
				interval: 0 * time.Millisecond,
				cpuLimits: framework.ContainersCPUSummary{
					stats.SystemContainerKubelet: {0.50: 0.30, 0.95: 0.50},
					stats.SystemContainerRuntime: {0.50: 0.40, 0.95: 0.60},
				},
				memLimits: framework.ResourceUsagePerContainer{
					stats.SystemContainerKubelet: &framework.ContainerResourceUsage{MemoryRSSInBytes: 100 * 1024 * 1024},
					stats.SystemContainerRuntime: &framework.ContainerResourceUsage{MemoryRSSInBytes: 500 * 1024 * 1024},
				},
				// percentile limit of single pod startup latency
				podStartupLimits: framework.LatencyMetric{
					Perc50: 30 * time.Second,
					Perc90: 54 * time.Second,
					Perc99: 59 * time.Second,
				},
				// upbound of startup latency of a batch of pods
				podBatchStartupLimit: 10 * time.Minute,
			},
		}

		for _, testArg := range dTests {
			itArg := testArg
			desc := fmt.Sprintf("latency/resource should be within limit when create %d pods with %v interval", itArg.podsNr, itArg.interval)
			It(desc, func() {
				itArg.createMethod = "batch"

				runDensityBatchTest(f, itArg)

			})
		}
	})


})

type densityTest struct {
	// number of pods
	podsNr int
	// number of background pods
	bgPodsNr int
	// interval between creating pod (rate control)
	interval time.Duration
	// create pods in 'batch' or 'sequence'
	createMethod string
	// API QPS limit
	APIQPSLimit int
	// performance limits
	cpuLimits            framework.ContainersCPUSummary
	memLimits            framework.ResourceUsagePerContainer
	podStartupLimits     framework.LatencyMetric
	podBatchStartupLimit time.Duration
}

// runDensityBatchTest runs the density batch pod creation test
func runDensityBatchTest(f *framework.Framework, testArg densityTest) (time.Duration, []framework.PodLatencyData) {
	const (
		podType               = "density_test_pod"
		sleepBeforeCreatePods = 30 * time.Second
	)
	var (
		mutex      = &sync.Mutex{}
		watchTimes = make(map[string]metav1.Time, 0)
		stopCh     = make(chan struct{})
	)

	// create test pod data structure
	pods := newTestPods(testArg.podsNr, false, imageutils.GetPauseImageName(), podType)

	// the controller watches the change of pod status
	controller := newInformerWatchPod(f, mutex, watchTimes, podType)
	go controller.Run(stopCh)
	defer close(stopCh)

	// TODO(coufon): in the test we found kubelet starts while it is busy on something, as a result 'syncLoop'
	// does not response to pod creation immediately. Creating the first pod has a delay around 5s.
	// The node status has already been 'ready' so `wait and check node being ready does not help here.
	// Now wait here for a grace period to let 'syncLoop' be ready
	time.Sleep(sleepBeforeCreatePods)

	//rc.Start()

	By("Creating a batch of pods")
	// It returns a map['pod name']'creation time' containing the creation timestamps
	createTimes := createBatchPodWithRateControl(f, pods, testArg.interval)

	By("Waiting for all Pods to be observed by the watch...")

	Eventually(func() bool {
		return len(watchTimes) == testArg.podsNr
	}, 10*time.Minute, 10*time.Second).Should(BeTrue())

	if len(watchTimes) < testArg.podsNr {
		framework.Failf("Timeout reached waiting for all Pods to be observed by the watch.")
	}

	// Analyze results
	var (
		firstCreate metav1.Time
		lastRunning metav1.Time
		init        = true
		e2eLags     = make([]framework.PodLatencyData, 0)
	)

	for name, create := range createTimes {
		watch, ok := watchTimes[name]
		Expect(ok).To(Equal(true))

		e2eLags = append(e2eLags,
			framework.PodLatencyData{Name: name, Latency: watch.Time.Sub(create.Time)})

		if !init {
			if firstCreate.Time.After(create.Time) {
				firstCreate = create
			}
			if lastRunning.Time.Before(watch.Time) {
				lastRunning = watch
			}
		} else {
			init = false
			firstCreate, lastRunning = create, watch
		}
	}

	sort.Sort(framework.LatencySlice(e2eLags))
	batchLag := lastRunning.Time.Sub(firstCreate.Time)

	deletePodsSync(f, pods)

	return batchLag, e2eLags
}

// createBatchPodWithRateControl creates a batch of pods concurrently, uses one goroutine for each creation.
// between creations there is an interval for throughput control
func createBatchPodWithRateControl(f *framework.Framework, pods []*v1.Pod, interval time.Duration) map[string]metav1.Time {
	createTimes := make(map[string]metav1.Time)
	for _, pod := range pods {
		createTimes[pod.ObjectMeta.Name] = metav1.Now()
		go f.PodClient().Create(pod)
		time.Sleep(interval)
	}
	return createTimes
}

// newInformerWatchPod creates an informer to check whether all pods are running.
func newInformerWatchPod(f *framework.Framework, mutex *sync.Mutex, watchTimes map[string]metav1.Time, podType string) cache.Controller {
	ns := f.Namespace.Name
	checkPodRunning := func(p *v1.Pod) {
		mutex.Lock()
		defer mutex.Unlock()
		defer GinkgoRecover()

		if p.Status.Phase == v1.PodRunning {
			if _, found := watchTimes[p.Name]; !found {
				watchTimes[p.Name] = metav1.Now()
			}
		}
	}

	_, controller := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				options.LabelSelector = labels.SelectorFromSet(labels.Set{"type": podType}).String()
				obj, err := f.ClientSet.CoreV1().Pods(ns).List(options)
				return runtime.Object(obj), err
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				options.LabelSelector = labels.SelectorFromSet(labels.Set{"type": podType}).String()
				return f.ClientSet.CoreV1().Pods(ns).Watch(options)
			},
		},
		&v1.Pod{},
		0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				p, ok := obj.(*v1.Pod)
				Expect(ok).To(Equal(true))
				go checkPodRunning(p)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				p, ok := newObj.(*v1.Pod)
				Expect(ok).To(Equal(true))
				go checkPodRunning(p)
			},
		},
	)
	return controller
}


// newTestPods creates a list of pods (specification) for test.
func newTestPods(numPods int, volume bool, imageName, podType string) []*v1.Pod {
	var pods []*v1.Pod
	for i := 0; i < numPods; i++ {
		podName := "test-" + string(uuid.NewUUID())
		labels := map[string]string{
			"type": podType,
			"name": podName,
		}
		if volume {
			pods = append(pods,
				&v1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:   podName,
						Labels: labels,
					},
					Spec: v1.PodSpec{
						// Restart policy is always (default).
						Containers: []v1.Container{
							{
								Image: imageName,
								Name:  podName,
								VolumeMounts: []v1.VolumeMount{
									{MountPath: "/test-volume-mnt", Name: podName + "-volume"},
								},
							},
						},
						NodeSelector: map[string]string{
							"beta.kubernetes.io/os": "windows",
						},
						Volumes: []v1.Volume{
							{Name: podName + "-volume", VolumeSource: v1.VolumeSource{EmptyDir: &v1.EmptyDirVolumeSource{}}},
						},
					},
				})
		} else {
			pods = append(pods,
				&v1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:   podName,
						Labels: labels,
					},
					Spec: v1.PodSpec{
						// Restart policy is always (default).
						Containers: []v1.Container{
							{
								Image: imageName,
								Name:  podName,
							},
						},
						NodeSelector: map[string]string{
							"beta.kubernetes.io/os": "windows",
						},
					},
				})
		}

	}
	return pods
}

// deletePodsSync deletes a list of pods and block until pods disappear.
func deletePodsSync(f *framework.Framework, pods []*v1.Pod) {
	var wg sync.WaitGroup
	for _, pod := range pods {
		wg.Add(1)
		go func(pod *v1.Pod) {
			defer GinkgoRecover()
			defer wg.Done()

			err := f.PodClient().Delete(pod.ObjectMeta.Name, metav1.NewDeleteOptions(30))
			Expect(err).NotTo(HaveOccurred())

			Expect(framework.WaitForPodToDisappear(f.ClientSet, f.Namespace.Name, pod.ObjectMeta.Name, labels.Everything(),
				30*time.Second, 10*time.Minute)).NotTo(HaveOccurred())
		}(pod)
	}
	wg.Wait()
	return
}

