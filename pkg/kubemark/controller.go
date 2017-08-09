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

package kubemark

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	informersv1 "k8s.io/client-go/informers/core/v1"
	kubeclient "k8s.io/client-go/kubernetes"
	listersv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/controller"

	"github.com/golang/glog"
)

const (
	namespaceKubemark = "kubemark"
	hollowNodeName    = "hollow-node"
	nodeGroupLabel    = "autoscaling.k8s.io/nodegroup"
	numRetries        = 3
)

// KubemarkController is a simplified version of cloud provider for kubemark. It allows
// to add and delete nodes from a kubemark cluster and introduces nodegroups
// by applying labels to the kubemark's hollow-nodes.
type KubemarkController struct {
	nodeTemplate    *apiv1.ReplicationController
	externalCluster externalCluster
	kubemarkCluster kubemarkCluster
	rand            *rand.Rand
}

// externalCluster is used to communicate with the external cluster that hosts
// kubemark, in order to be able to list, create and delete hollow nodes
// by manipulating the replication controllers.
type externalCluster struct {
	rcLister  listersv1.ReplicationControllerLister
	rcSynced  cache.InformerSynced
	podLister listersv1.PodLister
	podSynced cache.InformerSynced
	client    kubeclient.Interface
}

// kubemarkCluster is used to delete nodes from kubemark cluster once their
// respective replication controllers have been deleted and the nodes have
// become unready. This is to cover for the fact that there is no proper cloud
// provider for kubemark that would care for deleting the nodes.
type kubemarkCluster struct {
	client            kubeclient.Interface
	nodeLister        listersv1.NodeLister
	nodeSynced        cache.InformerSynced
	nodesToDelete     map[string]bool
	nodesToDeleteLock sync.Mutex
}

// NewKubemarkController creates KubemarkController using the provided clients to talk to external
// and kubemark clusters.
func NewKubemarkController(externalClient kubeclient.Interface, externalInformerFactory informers.SharedInformerFactory,
	kubemarkClient kubeclient.Interface, kubemarkNodeInformer informersv1.NodeInformer) (*KubemarkController, error) {
	rcInformer := externalInformerFactory.InformerFor(&apiv1.ReplicationController{}, newReplicationControllerInformer)
	podInformer := externalInformerFactory.InformerFor(&apiv1.Pod{}, newPodInformer)
	controller := &KubemarkController{
		externalCluster: externalCluster{
			rcLister:  listersv1.NewReplicationControllerLister(rcInformer.GetIndexer()),
			rcSynced:  rcInformer.HasSynced,
			podLister: listersv1.NewPodLister(podInformer.GetIndexer()),
			podSynced: podInformer.HasSynced,
			client:    externalClient,
		},
		kubemarkCluster: kubemarkCluster{
			nodeLister:        kubemarkNodeInformer.Lister(),
			nodeSynced:        kubemarkNodeInformer.Informer().HasSynced,
			client:            kubemarkClient,
			nodesToDelete:     make(map[string]bool),
			nodesToDeleteLock: sync.Mutex{},
		},
		rand: rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}

	kubemarkNodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		UpdateFunc: controller.kubemarkCluster.removeUnneededNodes,
	})

	return controller, nil
}

// Init waits for population of caches and populates the node template needed
// for creation of kubemark nodes.
func (kubemarkController *KubemarkController) Init(stopCh chan struct{}) {
	if !controller.WaitForCacheSync("kubemark", stopCh,
		kubemarkController.externalCluster.rcSynced,
		kubemarkController.externalCluster.podSynced,
		kubemarkController.kubemarkCluster.nodeSynced) {
		return
	}

	// Get hollow node template from an existing hollow node to be able to create
	// new nodes based on it.
	nodeTemplate, err := kubemarkController.getNodeTemplate()
	if err != nil {
		glog.Fatalf("Failed to get node template: %s", err)
	}
	kubemarkController.nodeTemplate = nodeTemplate
}

// GetNodesForNodegroup returns list of the nodes in the node group.
func (kubemarkController *KubemarkController) GetNodeNamesForNodegroup(nodeGroup string) ([]string, error) {
	selector := labels.SelectorFromSet(labels.Set{nodeGroupLabel: nodeGroup})
	pods, err := kubemarkController.externalCluster.podLister.List(selector)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(pods))
	for _, pod := range pods {
		result = append(result, pod.ObjectMeta.Name)
	}
	return result, nil
}

// GetNodeGroupSize returns the current size for the node group.
func (kubemarkController *KubemarkController) GetNodeGroupSize(nodeGroup string) (int, error) {
	selector := labels.SelectorFromSet(labels.Set(map[string]string{nodeGroupLabel: nodeGroup}))
	nodes, err := kubemarkController.externalCluster.rcLister.List(selector)
	if err != nil {
		return 0, err
	}
	return len(nodes), nil
}

// SetNodeGroupSize changes the size of node group by adding or removing nodes.
func (kubemarkController *KubemarkController) SetNodeGroupSize(nodeGroup string, size int) error {
	currSize, err := kubemarkController.GetNodeGroupSize(nodeGroup)
	if err != nil {
		return err
	}
	switch delta := size - currSize; {
	case delta < 0:
		absDelta := -delta
		nodes, err := kubemarkController.GetNodeNamesForNodegroup(nodeGroup)
		if err != nil {
			return err
		}
		if len(nodes) > absDelta {
			return fmt.Errorf("can't remove %d nodes from %s nodegroup, not enough nodes", absDelta, nodeGroup)
		}
		for i, node := range nodes {
			if i == absDelta {
				return nil
			}
			if err := kubemarkController.removeNodeFromNodeGroup(nodeGroup, node); err != nil {
				return err
			}
		}
	case delta > 0:
		for i := 0; i < delta; i++ {
			if err := kubemarkController.addNodeToNodeGroup(nodeGroup); err != nil {
				return err
			}
		}
	}

	return nil
}

func (kubemarkController *KubemarkController) addNodeToNodeGroup(nodeGroup string) error {
	templateCopy, err := api.Scheme.Copy(kubemarkController.nodeTemplate)
	if err != nil {
		return err
	}
	node := templateCopy.(*apiv1.ReplicationController)
	node.Name = fmt.Sprintf("%s-%d", nodeGroup, kubemarkController.rand.Int63())
	node.Labels = map[string]string{nodeGroupLabel: nodeGroup, "name": node.Name}
	node.Spec.Template.Labels = node.Labels

	for i := 0; i < numRetries; i++ {
		_, err = kubemarkController.externalCluster.client.CoreV1().ReplicationControllers(node.Namespace).Create(node)
		if err == nil {
			return nil
		}
	}

	return err
}

func (kubemarkController *KubemarkController) removeNodeFromNodeGroup(nodeGroup string, node string) error {
	pods, err := kubemarkController.externalCluster.podLister.List(labels.Everything())
	if err != nil {
		return err
	}
	for _, pod := range pods {
		if pod.ObjectMeta.Name == node {
			if pod.ObjectMeta.Labels[nodeGroupLabel] != nodeGroup {
				return fmt.Errorf("can't delete node %s from nodegroup %s. Node is not in nodegroup", node, nodeGroup)
			}
			policy := metav1.DeletePropagationForeground
			for i := 0; i < numRetries; i++ {
				err := kubemarkController.externalCluster.client.CoreV1().ReplicationControllers(namespaceKubemark).Delete(
					pod.ObjectMeta.Labels["name"],
					&metav1.DeleteOptions{PropagationPolicy: &policy})
				if err == nil {
					glog.Infof("marking node %s for deletion", node)
					// Mark node for deletion from kubemark cluster.
					// Once it becomes unready after replication controller
					// deletion has been noticed, we will delete it explicitly.
					// This is to cover for the fact that kubemark does not
					// take care of this itself.
					kubemarkController.kubemarkCluster.markNodeForDeletion(node)
					return nil
				}
			}
		}
	}

	return fmt.Errorf("can't delete node %s from nodegroup %s. Node does not exist", node, nodeGroup)
}

func (kubemarkController *KubemarkController) getReplicationControllerByName(name string) *apiv1.ReplicationController {
	rcs, err := kubemarkController.externalCluster.rcLister.List(labels.Everything())
	if err != nil {
		return nil
	}
	for _, rc := range rcs {
		if rc.ObjectMeta.Name == name {
			return rc
		}
	}
	return nil
}

func (kubemarkController *KubemarkController) getNodeNameForPod(podName string) (string, error) {
	pods, err := kubemarkController.externalCluster.podLister.List(labels.Everything())
	if err != nil {
		return "", err
	}
	for _, pod := range pods {
		if pod.ObjectMeta.Name == podName {
			return pod.Labels["name"], nil
		}
	}
	return "", fmt.Errorf("pod %s not found", podName)
}

// getNodeTemplate returns the template for hollow node replication controllers
// by looking for an existing hollow node specification. This requires at least
// one kubemark node to be present on startup.
func (kubemarkController *KubemarkController) getNodeTemplate() (*apiv1.ReplicationController, error) {
	podName, err := kubemarkController.kubemarkCluster.getHollowNodeName()
	if err != nil {
		return nil, err
	}
	hollowNodeName, err := kubemarkController.getNodeNameForPod(podName)
	if err != nil {
		return nil, err
	}
	if hollowNode := kubemarkController.getReplicationControllerByName(hollowNodeName); hollowNode != nil {
		nodeTemplate := &apiv1.ReplicationController{
			Spec: apiv1.ReplicationControllerSpec{
				Template: hollowNode.Spec.Template,
			},
		}

		nodeTemplate.Spec.Selector = nil
		nodeTemplate.Namespace = namespaceKubemark
		one := int32(1)
		nodeTemplate.Spec.Replicas = &one

		return nodeTemplate, nil
	}
	return nil, fmt.Errorf("can't get hollow node template")
}

func (kubemarkCluster *kubemarkCluster) getHollowNodeName() (string, error) {
	nodes, err := kubemarkCluster.nodeLister.List(labels.Everything())
	if err != nil {
		return "", err
	}
	for _, node := range nodes {
		return node.Name, nil
	}
	return "", fmt.Errorf("did not find any hollow nodes in the cluster")
}

func (kubemarkCluster *kubemarkCluster) removeUnneededNodes(oldObj interface{}, newObj interface{}) {
	node, ok := newObj.(*apiv1.Node)
	if !ok {
		return
	}
	for _, condition := range node.Status.Conditions {
		// Delete node if it is in unready state, and it has been
		// explicitly marked for deletion.
		if condition.Type == apiv1.NodeReady && condition.Status != apiv1.ConditionTrue {
			kubemarkCluster.nodesToDeleteLock.Lock()
			defer kubemarkCluster.nodesToDeleteLock.Unlock()
			if kubemarkCluster.nodesToDelete[node.Name] {
				kubemarkCluster.nodesToDelete[node.Name] = false
				if err := kubemarkCluster.client.CoreV1().Nodes().Delete(node.Name, &metav1.DeleteOptions{}); err != nil {
					glog.Errorf("failed to delete node %s from kubemark cluster", node.Name)
				}
			}
			return
		}
	}
}

func (kubemarkCluster *kubemarkCluster) markNodeForDeletion(name string) {
	kubemarkCluster.nodesToDeleteLock.Lock()
	defer kubemarkCluster.nodesToDeleteLock.Unlock()
	kubemarkCluster.nodesToDelete[name] = true
}

func newReplicationControllerInformer(kubeClient kubeclient.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	rcListWatch := cache.NewListWatchFromClient(kubeClient.CoreV1().RESTClient(), "replicationcontrollers", namespaceKubemark, fields.Everything())
	return cache.NewSharedIndexInformer(rcListWatch, &apiv1.ReplicationController{}, resyncPeriod, nil)
}

func newPodInformer(kubeClient kubeclient.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	podListWatch := cache.NewListWatchFromClient(kubeClient.CoreV1().RESTClient(), "pods", namespaceKubemark, fields.Everything())
	return cache.NewSharedIndexInformer(podListWatch, &apiv1.Pod{}, resyncPeriod, nil)
}
