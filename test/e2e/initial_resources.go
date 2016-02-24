/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package e2e

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/kubernetes/pkg/api"
)

// [Feature:InitialResources]: Initial resources is an experimental feature, so
// these tests are not run by default.
//
// Flaky issue #20272
var _ = Describe("Initial Resources [Feature:InitialResources] [Flaky]", func() {
	f := NewDefaultFramework("initial-resources")

	It("should set initial resources based on historical data", func() {
		// TODO(piosz): Add cleanup data in InfluxDB that left from previous tests.
		cpu := 100
		mem := 200
		for i := 0; i < 10; i++ {
			rc := NewStaticResourceConsumer(fmt.Sprintf("ir-%d", i), 1, cpu, mem, int64(2*cpu), int64(2*mem), f)
			defer rc.CleanUp()
		}
		// Wait some time to make sure usage data is gathered.
		time.Sleep(10 * time.Minute)

		pod := runPod(f, "ir-test", resourceConsumerImage)
		r := pod.Spec.Containers[0].Resources.Requests
		Expect(r.Cpu().MilliValue()).Should(BeNumerically("~", cpu, 10))
		Expect(r.Memory().Value()).Should(BeNumerically("~", mem*1024*1024, 20*1024*1024))
	})
})

func runPod(f *Framework, name, image string) *api.Pod {
	pod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			Name: name,
		},
		Spec: api.PodSpec{
			Containers: []api.Container{
				{
					Name:  name,
					Image: image,
				},
			},
		},
	}
	createdPod, err := f.Client.Pods(f.Namespace.Name).Create(pod)
	expectNoError(err)
	expectNoError(waitForPodRunningInNamespace(f.Client, name, f.Namespace.Name))
	return createdPod
}
