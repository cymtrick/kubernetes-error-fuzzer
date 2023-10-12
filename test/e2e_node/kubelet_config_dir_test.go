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

package e2enode

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeletconfig "k8s.io/kubernetes/pkg/kubelet/apis/config"
	"k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/nodefeature"
)

var _ = SIGDescribe("Kubelet Config", framework.WithSlow(), framework.WithSerial(), framework.WithDisruptive(), nodefeature.KubeletConfigDropInDir, func() {
	f := framework.NewDefaultFramework("kubelet-config-drop-in-dir-test")
	ginkgo.Context("when merging drop-in configs", func() {
		var oldcfg *kubeletconfig.KubeletConfiguration
		ginkgo.BeforeEach(func(ctx context.Context) {
			var err error
			oldcfg, err = getCurrentKubeletConfig(ctx)
			framework.ExpectNoError(err)
		})
		ginkgo.AfterEach(func(ctx context.Context) {
			files, err := filepath.Glob(filepath.Join(framework.TestContext.KubeletConfigDropinDir, "*"+".conf"))
			framework.ExpectNoError(err)
			for _, file := range files {
				err := os.Remove(file)
				framework.ExpectNoError(err)
			}
			updateKubeletConfig(ctx, f, oldcfg, true)
		})
		ginkgo.It("should merge kubelet configs correctly", func(ctx context.Context) {
			// Get the initial kubelet configuration
			initialConfig, err := getCurrentKubeletConfig(ctx)
			framework.ExpectNoError(err)

			ginkgo.By("Stopping the kubelet")
			restartKubelet := stopKubelet()

			// wait until the kubelet health check will fail
			gomega.Eventually(ctx, func() bool {
				return kubeletHealthCheck(kubeletHealthCheckURL)
			}, f.Timeouts.PodStart, f.Timeouts.Poll).Should(gomega.BeFalse())

			configDir := framework.TestContext.KubeletConfigDropinDir

			contents := []byte(`apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
port: 10255
readOnlyPort: 10257
clusterDNS:
- 192.168.1.10
systemReserved:
  memory: 1Gi`)
			framework.ExpectNoError(os.WriteFile(filepath.Join(configDir, "10-kubelet.conf"), contents, 0755))
			contents = []byte(`apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
clusterDNS:
- 192.168.1.1
- 192.168.1.5
- 192.168.1.8
port: 8080
cpuManagerReconcilePeriod: 0s
systemReserved:
  memory: 2Gi`)
			framework.ExpectNoError(os.WriteFile(filepath.Join(configDir, "20-kubelet.conf"), contents, 0755))
			ginkgo.By("Restarting the kubelet")
			restartKubelet()
			// wait until the kubelet health check will succeed
			gomega.Eventually(ctx, func() bool {
				return kubeletHealthCheck(kubeletHealthCheckURL)
			}, f.Timeouts.PodStart, f.Timeouts.Poll).Should(gomega.BeTrue())

			mergedConfig, err := getCurrentKubeletConfig(ctx)
			framework.ExpectNoError(err)

			// Replace specific fields in the initial configuration with expectedConfig values
			initialConfig.Port = int32(8080)                  // not overridden by second file, should be retained.
			initialConfig.ReadOnlyPort = int32(10257)         // overridden by second file.
			initialConfig.SystemReserved = map[string]string{ // overridden by map in second file.
				"memory": "2Gi",
			}
			initialConfig.ClusterDNS = []string{"192.168.1.1", "192.168.1.5", "192.168.1.8"} // overridden by slice in second file.
			// This value was explicitly set in the drop-in, make sure it is retained
			initialConfig.CPUManagerReconcilePeriod = metav1.Duration{Duration: time.Duration(0)}
			// Meanwhile, this value was not explicitly set, but could have been overridden by a "default" of 0 for the type.
			// Ensure the true default persists.
			initialConfig.CPUCFSQuotaPeriod = metav1.Duration{Duration: time.Duration(100000000)}
			// Compare the expected config with the merged config
			gomega.Expect(initialConfig).To(gomega.BeComparableTo(mergedConfig), "Merged kubelet config does not match the expected configuration.")
		})
	})

})
