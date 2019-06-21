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

package e2e

import (
	"fmt"
	"os/exec"

	"k8s.io/kubernetes/test/e2e/framework"
	e2elog "k8s.io/kubernetes/test/e2e/framework/log"

	"github.com/onsi/ginkgo"
)

var _ = framework.KubeDescribe("GKE node pools [Feature:GKENodePool]", func() {

	f := framework.NewDefaultFramework("node-pools")

	ginkgo.BeforeEach(func() {
		framework.SkipUnlessProviderIs("gke")
	})

	ginkgo.It("should create a cluster with multiple node pools [Feature:GKENodePool]", func() {
		e2elog.Logf("Start create node pool test")
		testCreateDeleteNodePool(f, "test-pool")
	})
})

func testCreateDeleteNodePool(f *framework.Framework, poolName string) {
	e2elog.Logf("Create node pool: %q in cluster: %q", poolName, framework.TestContext.CloudConfig.Cluster)

	clusterStr := fmt.Sprintf("--cluster=%s", framework.TestContext.CloudConfig.Cluster)

	out, err := exec.Command("gcloud", "container", "node-pools", "create",
		poolName,
		clusterStr,
		"--num-nodes=2").CombinedOutput()
	e2elog.Logf("\n%s", string(out))
	if err != nil {
		e2elog.Failf("Failed to create node pool %q. Err: %v\n%v", poolName, err, string(out))
	}
	e2elog.Logf("Successfully created node pool %q.", poolName)

	out, err = exec.Command("gcloud", "container", "node-pools", "list",
		clusterStr).CombinedOutput()
	if err != nil {
		e2elog.Failf("Failed to list node pools from cluster %q. Err: %v\n%v", framework.TestContext.CloudConfig.Cluster, err, string(out))
	}
	e2elog.Logf("Node pools:\n%s", string(out))

	e2elog.Logf("Checking that 2 nodes have the correct node pool label.")
	nodeCount := nodesWithPoolLabel(f, poolName)
	if nodeCount != 2 {
		e2elog.Failf("Wanted 2 nodes with node pool label, got: %v", nodeCount)
	}
	e2elog.Logf("Success, found 2 nodes with correct node pool labels.")

	e2elog.Logf("Deleting node pool: %q in cluster: %q", poolName, framework.TestContext.CloudConfig.Cluster)
	out, err = exec.Command("gcloud", "container", "node-pools", "delete",
		poolName,
		clusterStr,
		"-q").CombinedOutput()
	e2elog.Logf("\n%s", string(out))
	if err != nil {
		e2elog.Failf("Failed to delete node pool %q. Err: %v\n%v", poolName, err, string(out))
	}
	e2elog.Logf("Successfully deleted node pool %q.", poolName)

	out, err = exec.Command("gcloud", "container", "node-pools", "list",
		clusterStr).CombinedOutput()
	if err != nil {
		e2elog.Failf("\nFailed to list node pools from cluster %q. Err: %v\n%v", framework.TestContext.CloudConfig.Cluster, err, string(out))
	}
	e2elog.Logf("\nNode pools:\n%s", string(out))

	e2elog.Logf("Checking that no nodes have the deleted node pool's label.")
	nodeCount = nodesWithPoolLabel(f, poolName)
	if nodeCount != 0 {
		e2elog.Failf("Wanted 0 nodes with node pool label, got: %v", nodeCount)
	}
	e2elog.Logf("Success, found no nodes with the deleted node pool's label.")
}

// nodesWithPoolLabel returns the number of nodes that have the "gke-nodepool"
// label with the given node pool name.
func nodesWithPoolLabel(f *framework.Framework, poolName string) int {
	nodeCount := 0
	nodeList := framework.GetReadySchedulableNodesOrDie(f.ClientSet)
	for _, node := range nodeList.Items {
		if poolLabel := node.Labels["cloud.google.com/gke-nodepool"]; poolLabel == poolName {
			nodeCount++
		}
	}
	return nodeCount
}
