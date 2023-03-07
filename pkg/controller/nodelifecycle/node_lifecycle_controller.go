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

// The Controller sets tainted annotations on nodes.
// Tainted nodes should not be used for new work loads and
// some effort should be given to getting existing work
// loads off of tainted nodes.

package nodelifecycle

import (
	"context"
	"fmt"
	"sync"
	"time"

	"k8s.io/klog/v2"

	coordv1 "k8s.io/api/coordination/v1"
	v1 "k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	appsv1informers "k8s.io/client-go/informers/apps/v1"
	coordinformers "k8s.io/client-go/informers/coordination/v1"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	appsv1listers "k8s.io/client-go/listers/apps/v1"
	coordlisters "k8s.io/client-go/listers/coordination/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/flowcontrol"
	"k8s.io/client-go/util/workqueue"
	nodetopology "k8s.io/component-helpers/node/topology"
	kubeletapis "k8s.io/kubelet/pkg/apis"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/nodelifecycle/scheduler"
	controllerutil "k8s.io/kubernetes/pkg/controller/util/node"
	taintutils "k8s.io/kubernetes/pkg/util/taints"
)

func init() {
	// Register prometheus metrics
	Register()
}

var (
	// UnreachableTaintTemplate is the taint for when a node becomes unreachable.
	UnreachableTaintTemplate = &v1.Taint{
		Key:    v1.TaintNodeUnreachable,
		Effect: v1.TaintEffectNoExecute,
	}

	// NotReadyTaintTemplate is the taint for when a node is not ready for
	// executing pods
	NotReadyTaintTemplate = &v1.Taint{
		Key:    v1.TaintNodeNotReady,
		Effect: v1.TaintEffectNoExecute,
	}

	// map {NodeConditionType: {ConditionStatus: TaintKey}}
	// represents which NodeConditionType under which ConditionStatus should be
	// tainted with which TaintKey
	// for certain NodeConditionType, there are multiple {ConditionStatus,TaintKey} pairs
	nodeConditionToTaintKeyStatusMap = map[v1.NodeConditionType]map[v1.ConditionStatus]string{
		v1.NodeReady: {
			v1.ConditionFalse:   v1.TaintNodeNotReady,
			v1.ConditionUnknown: v1.TaintNodeUnreachable,
		},
		v1.NodeMemoryPressure: {
			v1.ConditionTrue: v1.TaintNodeMemoryPressure,
		},
		v1.NodeDiskPressure: {
			v1.ConditionTrue: v1.TaintNodeDiskPressure,
		},
		v1.NodeNetworkUnavailable: {
			v1.ConditionTrue: v1.TaintNodeNetworkUnavailable,
		},
		v1.NodePIDPressure: {
			v1.ConditionTrue: v1.TaintNodePIDPressure,
		},
	}

	taintKeyToNodeConditionMap = map[string]v1.NodeConditionType{
		v1.TaintNodeNotReady:           v1.NodeReady,
		v1.TaintNodeUnreachable:        v1.NodeReady,
		v1.TaintNodeNetworkUnavailable: v1.NodeNetworkUnavailable,
		v1.TaintNodeMemoryPressure:     v1.NodeMemoryPressure,
		v1.TaintNodeDiskPressure:       v1.NodeDiskPressure,
		v1.TaintNodePIDPressure:        v1.NodePIDPressure,
	}
)

// ZoneState is the state of a given zone.
type ZoneState string

const (
	stateInitial           = ZoneState("Initial")
	stateNormal            = ZoneState("Normal")
	stateFullDisruption    = ZoneState("FullDisruption")
	statePartialDisruption = ZoneState("PartialDisruption")
)

const (
	// The amount of time the nodecontroller should sleep between retrying node health updates
	retrySleepTime   = 20 * time.Millisecond
	nodeNameKeyIndex = "spec.nodeName"
	// podUpdateWorkerSizes assumes that in most cases pod will be handled by monitorNodeHealth pass.
	// Pod update workers will only handle lagging cache pods. 4 workers should be enough.
	podUpdateWorkerSize = 4
)

// labelReconcileInfo lists Node labels to reconcile, and how to reconcile them.
// primaryKey and secondaryKey are keys of labels to reconcile.
//   - If both keys exist, but their values don't match. Use the value from the
//     primaryKey as the source of truth to reconcile.
//   - If ensureSecondaryExists is true, and the secondaryKey does not
//     exist, secondaryKey will be added with the value of the primaryKey.
var labelReconcileInfo = []struct {
	primaryKey            string
	secondaryKey          string
	ensureSecondaryExists bool
}{
	{
		// Reconcile the beta and the stable OS label using the stable label as the source of truth.
		// TODO(#89477): no earlier than 1.22: drop the beta labels if they differ from the GA labels
		primaryKey:            v1.LabelOSStable,
		secondaryKey:          kubeletapis.LabelOS,
		ensureSecondaryExists: true,
	},
	{
		// Reconcile the beta and the stable arch label using the stable label as the source of truth.
		// TODO(#89477): no earlier than 1.22: drop the beta labels if they differ from the GA labels
		primaryKey:            v1.LabelArchStable,
		secondaryKey:          kubeletapis.LabelArch,
		ensureSecondaryExists: true,
	},
}

type nodeHealthData struct {
	probeTimestamp           metav1.Time
	readyTransitionTimestamp metav1.Time
	status                   *v1.NodeStatus
	lease                    *coordv1.Lease
}

func (n *nodeHealthData) deepCopy() *nodeHealthData {
	if n == nil {
		return nil
	}
	return &nodeHealthData{
		probeTimestamp:           n.probeTimestamp,
		readyTransitionTimestamp: n.readyTransitionTimestamp,
		status:                   n.status.DeepCopy(),
		lease:                    n.lease.DeepCopy(),
	}
}

type nodeHealthMap struct {
	lock        sync.RWMutex
	nodeHealths map[string]*nodeHealthData
}

func newNodeHealthMap() *nodeHealthMap {
	return &nodeHealthMap{
		nodeHealths: make(map[string]*nodeHealthData),
	}
}

// getDeepCopy - returns copy of node health data.
// It prevents data being changed after retrieving it from the map.
func (n *nodeHealthMap) getDeepCopy(name string) *nodeHealthData {
	n.lock.RLock()
	defer n.lock.RUnlock()
	return n.nodeHealths[name].deepCopy()
}

func (n *nodeHealthMap) set(name string, data *nodeHealthData) {
	n.lock.Lock()
	defer n.lock.Unlock()
	n.nodeHealths[name] = data
}

type podUpdateItem struct {
	namespace string
	name      string
}

type evictionStatus int

const (
	unmarked = iota
	toBeEvicted
	evicted
)

// nodeEvictionMap stores evictionStatus data for each node.
type nodeEvictionMap struct {
	lock          sync.Mutex
	nodeEvictions map[string]evictionStatus
}

func newNodeEvictionMap() *nodeEvictionMap {
	return &nodeEvictionMap{
		nodeEvictions: make(map[string]evictionStatus),
	}
}

func (n *nodeEvictionMap) registerNode(nodeName string) {
	n.lock.Lock()
	defer n.lock.Unlock()
	n.nodeEvictions[nodeName] = unmarked
}

func (n *nodeEvictionMap) unregisterNode(nodeName string) {
	n.lock.Lock()
	defer n.lock.Unlock()
	delete(n.nodeEvictions, nodeName)
}

func (n *nodeEvictionMap) setStatus(nodeName string, status evictionStatus) bool {
	n.lock.Lock()
	defer n.lock.Unlock()
	if _, exists := n.nodeEvictions[nodeName]; !exists {
		return false
	}
	n.nodeEvictions[nodeName] = status
	return true
}

func (n *nodeEvictionMap) getStatus(nodeName string) (evictionStatus, bool) {
	n.lock.Lock()
	defer n.lock.Unlock()
	if _, exists := n.nodeEvictions[nodeName]; !exists {
		return unmarked, false
	}
	return n.nodeEvictions[nodeName], true
}

// Controller is the controller that manages node's life cycle.
type Controller struct {
	taintManager *scheduler.NoExecuteTaintManager

	podLister         corelisters.PodLister
	podInformerSynced cache.InformerSynced
	kubeClient        clientset.Interface

	// This timestamp is to be used instead of LastProbeTime stored in Condition. We do this
	// to avoid the problem with time skew across the cluster.
	now func() metav1.Time

	enterPartialDisruptionFunc func(nodeNum int) float32
	enterFullDisruptionFunc    func(nodeNum int) float32
	computeZoneStateFunc       func(nodeConditions []*v1.NodeCondition) (int, ZoneState)

	knownNodeSet map[string]*v1.Node
	// per Node map storing last observed health together with a local time when it was observed.
	nodeHealthMap *nodeHealthMap

	// evictorLock protects zonePodEvictor and zoneNoExecuteTainter.
	evictorLock     sync.Mutex
	nodeEvictionMap *nodeEvictionMap
	// workers that evicts pods from unresponsive nodes.
	zonePodEvictor map[string]*scheduler.RateLimitedTimedQueue
	// workers that are responsible for tainting nodes.
	zoneNoExecuteTainter map[string]*scheduler.RateLimitedTimedQueue

	nodesToRetry sync.Map

	zoneStates map[string]ZoneState

	daemonSetStore          appsv1listers.DaemonSetLister
	daemonSetInformerSynced cache.InformerSynced

	leaseLister         coordlisters.LeaseLister
	leaseInformerSynced cache.InformerSynced
	nodeLister          corelisters.NodeLister
	nodeInformerSynced  cache.InformerSynced

	getPodsAssignedToNode func(nodeName string) ([]*v1.Pod, error)

	broadcaster record.EventBroadcaster
	recorder    record.EventRecorder

	// Value controlling Controller monitoring period, i.e. how often does Controller
	// check node health signal posted from kubelet. This value should be lower than
	// nodeMonitorGracePeriod.
	// TODO: Change node health monitor to watch based.
	nodeMonitorPeriod time.Duration

	// When node is just created, e.g. cluster bootstrap or node creation, we give
	// a longer grace period.
	nodeStartupGracePeriod time.Duration

	// Controller will not proactively sync node health, but will monitor node
	// health signal updated from kubelet. There are 2 kinds of node healthiness
	// signals: NodeStatus and NodeLease. If it doesn't receive update for this amount
	// of time, it will start posting "NodeReady==ConditionUnknown". The amount of
	// time before which Controller start evicting pods is controlled via flag
	// 'pod-eviction-timeout'.
	// Note: be cautious when changing the constant, it must work with
	// nodeStatusUpdateFrequency in kubelet and renewInterval in NodeLease
	// controller. The node health signal update frequency is the minimal of the
	// two.
	// There are several constraints:
	// 1. nodeMonitorGracePeriod must be N times more than  the node health signal
	//    update frequency, where N means number of retries allowed for kubelet to
	//    post node status/lease. It is pointless to make nodeMonitorGracePeriod
	//    be less than the node health signal update frequency, since there will
	//    only be fresh values from Kubelet at an interval of node health signal
	//    update frequency. The constant must be less than podEvictionTimeout.
	// 2. nodeMonitorGracePeriod can't be too large for user experience - larger
	//    value takes longer for user to see up-to-date node health.
	nodeMonitorGracePeriod time.Duration

	// Number of workers Controller uses to process node monitor health updates.
	// Defaults to scheduler.UpdateWorkerSize.
	nodeUpdateWorkerSize int

	podEvictionTimeout          time.Duration
	evictionLimiterQPS          float32
	secondaryEvictionLimiterQPS float32
	largeClusterThreshold       int32
	unhealthyZoneThreshold      float32

	// if set to true Controller will start TaintManager that will evict Pods from
	// tainted nodes, if they're not tolerated.
	runTaintManager bool

	nodeUpdateQueue workqueue.Interface
	podUpdateQueue  workqueue.RateLimitingInterface
}

// NewNodeLifecycleController returns a new taint controller.
func NewNodeLifecycleController(
	ctx context.Context,
	leaseInformer coordinformers.LeaseInformer,
	podInformer coreinformers.PodInformer,
	nodeInformer coreinformers.NodeInformer,
	daemonSetInformer appsv1informers.DaemonSetInformer,
	kubeClient clientset.Interface,
	nodeMonitorPeriod time.Duration,
	nodeStartupGracePeriod time.Duration,
	nodeMonitorGracePeriod time.Duration,
	podEvictionTimeout time.Duration,
	evictionLimiterQPS float32,
	secondaryEvictionLimiterQPS float32,
	largeClusterThreshold int32,
	unhealthyZoneThreshold float32,
	runTaintManager bool,
) (*Controller, error) {
	logger := klog.LoggerWithName(klog.FromContext(ctx), "NodeLifecycleController")
	if kubeClient == nil {
		logger.Error(nil, "kubeClient is nil when starting nodelifecycle Controller")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	eventBroadcaster := record.NewBroadcaster()
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "node-controller"})

	nc := &Controller{
		kubeClient:                  kubeClient,
		now:                         metav1.Now,
		knownNodeSet:                make(map[string]*v1.Node),
		nodeHealthMap:               newNodeHealthMap(),
		nodeEvictionMap:             newNodeEvictionMap(),
		broadcaster:                 eventBroadcaster,
		recorder:                    recorder,
		nodeMonitorPeriod:           nodeMonitorPeriod,
		nodeStartupGracePeriod:      nodeStartupGracePeriod,
		nodeMonitorGracePeriod:      nodeMonitorGracePeriod,
		nodeUpdateWorkerSize:        scheduler.UpdateWorkerSize,
		zonePodEvictor:              make(map[string]*scheduler.RateLimitedTimedQueue),
		zoneNoExecuteTainter:        make(map[string]*scheduler.RateLimitedTimedQueue),
		nodesToRetry:                sync.Map{},
		zoneStates:                  make(map[string]ZoneState),
		podEvictionTimeout:          podEvictionTimeout,
		evictionLimiterQPS:          evictionLimiterQPS,
		secondaryEvictionLimiterQPS: secondaryEvictionLimiterQPS,
		largeClusterThreshold:       largeClusterThreshold,
		unhealthyZoneThreshold:      unhealthyZoneThreshold,
		runTaintManager:             runTaintManager,
		nodeUpdateQueue:             workqueue.NewNamed("node_lifecycle_controller"),
		podUpdateQueue:              workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "node_lifecycle_controller_pods"),
	}

	nc.enterPartialDisruptionFunc = nc.ReducedQPSFunc
	nc.enterFullDisruptionFunc = nc.HealthyQPSFunc
	nc.computeZoneStateFunc = nc.ComputeZoneState

	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*v1.Pod)
			nc.podUpdated(nil, pod)
			if nc.taintManager != nil {
				nc.taintManager.PodUpdated(nil, pod)
			}
		},
		UpdateFunc: func(prev, obj interface{}) {
			prevPod := prev.(*v1.Pod)
			newPod := obj.(*v1.Pod)
			nc.podUpdated(prevPod, newPod)
			if nc.taintManager != nil {
				nc.taintManager.PodUpdated(prevPod, newPod)
			}
		},
		DeleteFunc: func(obj interface{}) {
			pod, isPod := obj.(*v1.Pod)
			// We can get DeletedFinalStateUnknown instead of *v1.Pod here and we need to handle that correctly.
			if !isPod {
				deletedState, ok := obj.(cache.DeletedFinalStateUnknown)
				if !ok {
					logger.Error(nil, "Received unexpected object", "object", obj)
					return
				}
				pod, ok = deletedState.Obj.(*v1.Pod)
				if !ok {
					logger.Error(nil, "DeletedFinalStateUnknown contained non-Pod object", "object", deletedState.Obj)
					return
				}
			}
			nc.podUpdated(pod, nil)
			if nc.taintManager != nil {
				nc.taintManager.PodUpdated(pod, nil)
			}
		},
	})
	nc.podInformerSynced = podInformer.Informer().HasSynced
	podInformer.Informer().AddIndexers(cache.Indexers{
		nodeNameKeyIndex: func(obj interface{}) ([]string, error) {
			pod, ok := obj.(*v1.Pod)
			if !ok {
				return []string{}, nil
			}
			if len(pod.Spec.NodeName) == 0 {
				return []string{}, nil
			}
			return []string{pod.Spec.NodeName}, nil
		},
	})

	podIndexer := podInformer.Informer().GetIndexer()
	nc.getPodsAssignedToNode = func(nodeName string) ([]*v1.Pod, error) {
		objs, err := podIndexer.ByIndex(nodeNameKeyIndex, nodeName)
		if err != nil {
			return nil, err
		}
		pods := make([]*v1.Pod, 0, len(objs))
		for _, obj := range objs {
			pod, ok := obj.(*v1.Pod)
			if !ok {
				continue
			}
			pods = append(pods, pod)
		}
		return pods, nil
	}
	nc.podLister = podInformer.Lister()
	nc.nodeLister = nodeInformer.Lister()

	if nc.runTaintManager {
		nc.taintManager = scheduler.NewNoExecuteTaintManager(ctx, kubeClient, nc.podLister, nc.nodeLister, nc.getPodsAssignedToNode)
		nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
			AddFunc: controllerutil.CreateAddNodeHandler(func(node *v1.Node) error {
				nc.taintManager.NodeUpdated(nil, node)
				return nil
			}),
			UpdateFunc: controllerutil.CreateUpdateNodeHandler(func(oldNode, newNode *v1.Node) error {
				nc.taintManager.NodeUpdated(oldNode, newNode)
				return nil
			}),
			DeleteFunc: controllerutil.CreateDeleteNodeHandler(func(node *v1.Node) error {
				nc.taintManager.NodeUpdated(node, nil)
				return nil
			}),
		})
	}

	logger.Info("Controller will reconcile labels")
	nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controllerutil.CreateAddNodeHandler(func(node *v1.Node) error {
			nc.nodeUpdateQueue.Add(node.Name)
			nc.nodeEvictionMap.registerNode(node.Name)
			return nil
		}),
		UpdateFunc: controllerutil.CreateUpdateNodeHandler(func(_, newNode *v1.Node) error {
			nc.nodeUpdateQueue.Add(newNode.Name)
			return nil
		}),
		DeleteFunc: controllerutil.CreateDeleteNodeHandler(func(node *v1.Node) error {
			nc.nodesToRetry.Delete(node.Name)
			nc.nodeEvictionMap.unregisterNode(node.Name)
			return nil
		}),
	})

	nc.leaseLister = leaseInformer.Lister()
	nc.leaseInformerSynced = leaseInformer.Informer().HasSynced

	nc.nodeInformerSynced = nodeInformer.Informer().HasSynced

	nc.daemonSetStore = daemonSetInformer.Lister()
	nc.daemonSetInformerSynced = daemonSetInformer.Informer().HasSynced

	return nc, nil
}

// Run starts an asynchronous loop that monitors the status of cluster nodes.
func (nc *Controller) Run(ctx context.Context) {
	defer utilruntime.HandleCrash()

	// Start events processing pipeline.
	nc.broadcaster.StartStructuredLogging(0)
	logger := klog.FromContext(ctx)
	logger.Info("Sending events to api server")
	nc.broadcaster.StartRecordingToSink(
		&v1core.EventSinkImpl{
			Interface: v1core.New(nc.kubeClient.CoreV1().RESTClient()).Events(""),
		})
	defer nc.broadcaster.Shutdown()

	// Close node update queue to cleanup go routine.
	defer nc.nodeUpdateQueue.ShutDown()
	defer nc.podUpdateQueue.ShutDown()

	logger.Info("Starting node controller")
	defer logger.Info("Shutting down node controller")

	if !cache.WaitForNamedCacheSync("taint", ctx.Done(), nc.leaseInformerSynced, nc.nodeInformerSynced, nc.podInformerSynced, nc.daemonSetInformerSynced) {
		return
	}

	if nc.runTaintManager {
		go nc.taintManager.Run(ctx)
	}

	// Start workers to reconcile labels and/or update NoSchedule taint for nodes.
	for i := 0; i < scheduler.UpdateWorkerSize; i++ {
		// Thanks to "workqueue", each worker just need to get item from queue, because
		// the item is flagged when got from queue: if new event come, the new item will
		// be re-queued until "Done", so no more than one worker handle the same item and
		// no event missed.
		go wait.UntilWithContext(ctx, nc.doNodeProcessingPassWorker, time.Second)
	}

	for i := 0; i < podUpdateWorkerSize; i++ {
		go wait.UntilWithContext(ctx, nc.doPodProcessingWorker, time.Second)
	}

	if nc.runTaintManager {
		// Handling taint based evictions. Because we don't want a dedicated logic in TaintManager for NC-originated
		// taints and we normally don't rate limit evictions caused by taints, we need to rate limit adding taints.
		go wait.UntilWithContext(ctx, nc.doNoExecuteTaintingPass, scheduler.NodeEvictionPeriod)
	} else {
		// Managing eviction of nodes:
		// When we delete pods off a node, if the node was not empty at the time we then
		// queue an eviction watcher. If we hit an error, retry deletion.
		go wait.UntilWithContext(ctx, nc.doEvictionPass, scheduler.NodeEvictionPeriod)
	}

	// Incorporate the results of node health signal pushed from kubelet to master.
	go wait.UntilWithContext(ctx, func(ctx context.Context) {
		if err := nc.monitorNodeHealth(ctx); err != nil {
			logger.Error(err, "Error monitoring node health")
		}
	}, nc.nodeMonitorPeriod)

	<-ctx.Done()
}

func (nc *Controller) doNodeProcessingPassWorker(ctx context.Context) {
	logger := klog.FromContext(ctx)
	for {
		obj, shutdown := nc.nodeUpdateQueue.Get()
		// "nodeUpdateQueue" will be shutdown when "stopCh" closed;
		// we do not need to re-check "stopCh" again.
		if shutdown {
			return
		}
		nodeName := obj.(string)
		if err := nc.doNoScheduleTaintingPass(ctx, nodeName); err != nil {
			logger.Error(err, "Failed to taint NoSchedule on node, requeue it", "node", klog.KRef("", nodeName))
			// TODO(k82cn): Add nodeName back to the queue
		}
		// TODO: re-evaluate whether there are any labels that need to be
		// reconcile in 1.19. Remove this function if it's no longer necessary.
		if err := nc.reconcileNodeLabels(nodeName); err != nil {
			logger.Error(err, "Failed to reconcile labels for node, requeue it", "node", klog.KRef("", nodeName))
			// TODO(yujuhong): Add nodeName back to the queue
		}
		nc.nodeUpdateQueue.Done(nodeName)
	}
}

func (nc *Controller) doNoScheduleTaintingPass(ctx context.Context, nodeName string) error {
	node, err := nc.nodeLister.Get(nodeName)
	if err != nil {
		// If node not found, just ignore it.
		if apierrors.IsNotFound(err) {
			return nil
		}
		return err
	}

	// Map node's condition to Taints.
	var taints []v1.Taint
	for _, condition := range node.Status.Conditions {
		if taintMap, found := nodeConditionToTaintKeyStatusMap[condition.Type]; found {
			if taintKey, found := taintMap[condition.Status]; found {
				taints = append(taints, v1.Taint{
					Key:    taintKey,
					Effect: v1.TaintEffectNoSchedule,
				})
			}
		}
	}
	if node.Spec.Unschedulable {
		// If unschedulable, append related taint.
		taints = append(taints, v1.Taint{
			Key:    v1.TaintNodeUnschedulable,
			Effect: v1.TaintEffectNoSchedule,
		})
	}

	// Get exist taints of node.
	nodeTaints := taintutils.TaintSetFilter(node.Spec.Taints, func(t *v1.Taint) bool {
		// only NoSchedule taints are candidates to be compared with "taints" later
		if t.Effect != v1.TaintEffectNoSchedule {
			return false
		}
		// Find unschedulable taint of node.
		if t.Key == v1.TaintNodeUnschedulable {
			return true
		}
		// Find node condition taints of node.
		_, found := taintKeyToNodeConditionMap[t.Key]
		return found
	})
	taintsToAdd, taintsToDel := taintutils.TaintSetDiff(taints, nodeTaints)
	// If nothing to add or delete, return true directly.
	if len(taintsToAdd) == 0 && len(taintsToDel) == 0 {
		return nil
	}
	if !controllerutil.SwapNodeControllerTaint(ctx, nc.kubeClient, taintsToAdd, taintsToDel, node) {
		return fmt.Errorf("failed to swap taints of node %+v", node)
	}
	return nil
}

func (nc *Controller) doNoExecuteTaintingPass(ctx context.Context) {
	// Extract out the keys of the map in order to not hold
	// the evictorLock for the entire function and hold it
	// only when nescessary.
	var zoneNoExecuteTainterKeys []string
	func() {
		nc.evictorLock.Lock()
		defer nc.evictorLock.Unlock()

		zoneNoExecuteTainterKeys = make([]string, 0, len(nc.zoneNoExecuteTainter))
		for k := range nc.zoneNoExecuteTainter {
			zoneNoExecuteTainterKeys = append(zoneNoExecuteTainterKeys, k)
		}
	}()
	logger := klog.FromContext(ctx)
	for _, k := range zoneNoExecuteTainterKeys {
		var zoneNoExecuteTainterWorker *scheduler.RateLimitedTimedQueue
		func() {
			nc.evictorLock.Lock()
			defer nc.evictorLock.Unlock()
			// Extracting the value without checking if the key
			// exists or not is safe to do here since zones do
			// not get removed, and consequently pod evictors for
			// these zones also do not get removed, only added.
			zoneNoExecuteTainterWorker = nc.zoneNoExecuteTainter[k]
		}()
		// Function should return 'false' and a time after which it should be retried, or 'true' if it shouldn't (it succeeded).
		zoneNoExecuteTainterWorker.Try(logger, func(value scheduler.TimedValue) (bool, time.Duration) {
			node, err := nc.nodeLister.Get(value.Value)
			if apierrors.IsNotFound(err) {
				logger.Info("Node no longer present in nodeLister", "node", klog.KRef("", value.Value))
				return true, 0
			} else if err != nil {
				logger.Info("Failed to get Node from the nodeLister", "node", klog.KRef("", value.Value), "err", err)
				// retry in 50 millisecond
				return false, 50 * time.Millisecond
			}
			_, condition := controllerutil.GetNodeCondition(&node.Status, v1.NodeReady)
			// Because we want to mimic NodeStatus.Condition["Ready"] we make "unreachable" and "not ready" taints mutually exclusive.
			taintToAdd := v1.Taint{}
			oppositeTaint := v1.Taint{}
			switch condition.Status {
			case v1.ConditionFalse:
				taintToAdd = *NotReadyTaintTemplate
				oppositeTaint = *UnreachableTaintTemplate
			case v1.ConditionUnknown:
				taintToAdd = *UnreachableTaintTemplate
				oppositeTaint = *NotReadyTaintTemplate
			default:
				// It seems that the Node is ready again, so there's no need to taint it.
				logger.V(4).Info("Node was in a taint queue, but it's ready now. Ignoring taint request", "node", klog.KRef("", value.Value))
				return true, 0
			}
			result := controllerutil.SwapNodeControllerTaint(ctx, nc.kubeClient, []*v1.Taint{&taintToAdd}, []*v1.Taint{&oppositeTaint}, node)
			if result {
				//count the evictionsNumber
				zone := nodetopology.GetZoneKey(node)
				evictionsNumber.WithLabelValues(zone).Inc()
				evictionsTotal.WithLabelValues(zone).Inc()
			}

			return result, 0
		})
	}
}

func (nc *Controller) doEvictionPass(ctx context.Context) {
	// Extract out the keys of the map in order to not hold
	// the evictorLock for the entire function and hold it
	// only when nescessary.
	var zonePodEvictorKeys []string
	func() {
		nc.evictorLock.Lock()
		defer nc.evictorLock.Unlock()

		zonePodEvictorKeys = make([]string, 0, len(nc.zonePodEvictor))
		for k := range nc.zonePodEvictor {
			zonePodEvictorKeys = append(zonePodEvictorKeys, k)
		}
	}()
	logger := klog.FromContext(ctx)
	for _, k := range zonePodEvictorKeys {
		var zonePodEvictionWorker *scheduler.RateLimitedTimedQueue
		func() {
			nc.evictorLock.Lock()
			defer nc.evictorLock.Unlock()
			// Extracting the value without checking if the key
			// exists or not is safe to do here since zones do
			// not get removed, and consequently pod evictors for
			// these zones also do not get removed, only added.
			zonePodEvictionWorker = nc.zonePodEvictor[k]
		}()

		// Function should return 'false' and a time after which it should be retried, or 'true' if it shouldn't (it succeeded).
		zonePodEvictionWorker.Try(logger, func(value scheduler.TimedValue) (bool, time.Duration) {
			node, err := nc.nodeLister.Get(value.Value)
			if apierrors.IsNotFound(err) {
				logger.Info("Node no longer present in nodeLister", "node", klog.KRef("", value.Value))
			} else if err != nil {
				logger.Info("Failed to get Node from the nodeLister", "node", klog.KRef("", value.Value), "err", err)
			}
			nodeUID, _ := value.UID.(string)
			pods, err := nc.getPodsAssignedToNode(value.Value)
			if err != nil {
				utilruntime.HandleError(fmt.Errorf("unable to list pods from node %q: %v", value.Value, err))
				return false, 0
			}
			remaining, err := controllerutil.DeletePods(ctx, nc.kubeClient, pods, nc.recorder, value.Value, nodeUID, nc.daemonSetStore)
			if err != nil {
				// We are not setting eviction status here.
				// New pods will be handled by zonePodEvictor retry
				// instead of immediate pod eviction.
				utilruntime.HandleError(fmt.Errorf("unable to evict node %q: %v", value.Value, err))
				return false, 0
			}
			if !nc.nodeEvictionMap.setStatus(value.Value, evicted) {
				logger.V(2).Info("Node was unregistered in the meantime - skipping setting status", "node", klog.KRef("", value.Value))
			}
			if remaining {
				logger.Info("Pods awaiting deletion due to Controller eviction")
			}

			if node != nil {
				zone := nodetopology.GetZoneKey(node)
				evictionsNumber.WithLabelValues(zone).Inc()
				evictionsTotal.WithLabelValues(zone).Inc()
			}

			return true, 0
		})
	}
}

// monitorNodeHealth verifies node health are constantly updated by kubelet, and
// if not, post "NodeReady==ConditionUnknown".
// This function will taint nodes who are not ready or not reachable for a long period of time.
func (nc *Controller) monitorNodeHealth(ctx context.Context) error {
	start := nc.now()
	defer func() {
		updateAllNodesHealthDuration.Observe(time.Since(start.Time).Seconds())
	}()

	// We are listing nodes from local cache as we can tolerate some small delays
	// comparing to state from etcd and there is eventual consistency anyway.
	nodes, err := nc.nodeLister.List(labels.Everything())
	if err != nil {
		return err
	}
	added, deleted, newZoneRepresentatives := nc.classifyNodes(nodes)
	logger := klog.FromContext(ctx)
	for i := range newZoneRepresentatives {
		nc.addPodEvictorForNewZone(logger, newZoneRepresentatives[i])
	}
	for i := range added {
		logger.V(1).Info("Controller observed a new Node", "node", klog.KRef("", added[i].Name))
		controllerutil.RecordNodeEvent(nc.recorder, added[i].Name, string(added[i].UID), v1.EventTypeNormal, "RegisteredNode", fmt.Sprintf("Registered Node %v in Controller", added[i].Name))
		nc.knownNodeSet[added[i].Name] = added[i]
		nc.addPodEvictorForNewZone(logger, added[i])
		if nc.runTaintManager {
			nc.markNodeAsReachable(ctx, added[i])
		} else {
			nc.cancelPodEviction(logger, added[i])
		}
	}

	for i := range deleted {
		logger.V(1).Info("Controller observed a Node deletion", "node", klog.KRef("", deleted[i].Name))
		controllerutil.RecordNodeEvent(nc.recorder, deleted[i].Name, string(deleted[i].UID), v1.EventTypeNormal, "RemovingNode", fmt.Sprintf("Removing Node %v from Controller", deleted[i].Name))
		delete(nc.knownNodeSet, deleted[i].Name)
	}

	var zoneToNodeConditionsLock sync.Mutex
	zoneToNodeConditions := map[string][]*v1.NodeCondition{}
	updateNodeFunc := func(piece int) {
		start := nc.now()
		defer func() {
			updateNodeHealthDuration.Observe(time.Since(start.Time).Seconds())
		}()

		var gracePeriod time.Duration
		var observedReadyCondition v1.NodeCondition
		var currentReadyCondition *v1.NodeCondition
		node := nodes[piece].DeepCopy()

		if err := wait.PollImmediate(retrySleepTime, retrySleepTime*scheduler.NodeHealthUpdateRetry, func() (bool, error) {
			var err error
			gracePeriod, observedReadyCondition, currentReadyCondition, err = nc.tryUpdateNodeHealth(ctx, node)
			if err == nil {
				return true, nil
			}
			name := node.Name
			node, err = nc.kubeClient.CoreV1().Nodes().Get(ctx, name, metav1.GetOptions{})
			if err != nil {
				logger.Error(nil, "Failed while getting a Node to retry updating node health. Probably Node was deleted", "node", klog.KRef("", name))
				return false, err
			}
			return false, nil
		}); err != nil {
			logger.Error(err, "Update health of Node from Controller error, Skipping - no pods will be evicted", "node", klog.KObj(node))
			return
		}

		// Some nodes may be excluded from disruption checking
		if !isNodeExcludedFromDisruptionChecks(node) {
			zoneToNodeConditionsLock.Lock()
			zoneToNodeConditions[nodetopology.GetZoneKey(node)] = append(zoneToNodeConditions[nodetopology.GetZoneKey(node)], currentReadyCondition)
			zoneToNodeConditionsLock.Unlock()
		}

		if currentReadyCondition != nil {
			pods, err := nc.getPodsAssignedToNode(node.Name)
			if err != nil {
				utilruntime.HandleError(fmt.Errorf("unable to list pods of node %v: %v", node.Name, err))
				if currentReadyCondition.Status != v1.ConditionTrue && observedReadyCondition.Status == v1.ConditionTrue {
					// If error happened during node status transition (Ready -> NotReady)
					// we need to mark node for retry to force MarkPodsNotReady execution
					// in the next iteration.
					nc.nodesToRetry.Store(node.Name, struct{}{})
				}
				return
			}
			if nc.runTaintManager {
				nc.processTaintBaseEviction(ctx, node, &observedReadyCondition)
			} else {
				if err := nc.processNoTaintBaseEviction(ctx, node, &observedReadyCondition, gracePeriod, pods); err != nil {
					utilruntime.HandleError(fmt.Errorf("unable to evict all pods from node %v: %v; queuing for retry", node.Name, err))
				}
			}

			_, needsRetry := nc.nodesToRetry.Load(node.Name)
			switch {
			case currentReadyCondition.Status != v1.ConditionTrue && observedReadyCondition.Status == v1.ConditionTrue:
				// Report node event only once when status changed.
				controllerutil.RecordNodeStatusChange(nc.recorder, node, "NodeNotReady")
				fallthrough
			case needsRetry && observedReadyCondition.Status != v1.ConditionTrue:
				if err = controllerutil.MarkPodsNotReady(ctx, nc.kubeClient, nc.recorder, pods, node.Name); err != nil {
					utilruntime.HandleError(fmt.Errorf("unable to mark all pods NotReady on node %v: %v; queuing for retry", node.Name, err))
					nc.nodesToRetry.Store(node.Name, struct{}{})
					return
				}
			}
		}
		nc.nodesToRetry.Delete(node.Name)
	}

	// Marking the pods not ready on a node requires looping over them and
	// updating each pod's status one at a time. This is performed serially, and
	// can take a while if we're processing each node serially as well. So we
	// process them with bounded concurrency instead, since most of the time is
	// spent waiting on io.
	workqueue.ParallelizeUntil(ctx, nc.nodeUpdateWorkerSize, len(nodes), updateNodeFunc)

	nc.handleDisruption(ctx, zoneToNodeConditions, nodes)

	return nil
}

func (nc *Controller) processTaintBaseEviction(ctx context.Context, node *v1.Node, observedReadyCondition *v1.NodeCondition) {
	decisionTimestamp := nc.now()
	// Check eviction timeout against decisionTimestamp
	logger := klog.FromContext(ctx)
	switch observedReadyCondition.Status {
	case v1.ConditionFalse:
		// We want to update the taint straight away if Node is already tainted with the UnreachableTaint
		if taintutils.TaintExists(node.Spec.Taints, UnreachableTaintTemplate) {
			taintToAdd := *NotReadyTaintTemplate
			if !controllerutil.SwapNodeControllerTaint(ctx, nc.kubeClient, []*v1.Taint{&taintToAdd}, []*v1.Taint{UnreachableTaintTemplate}, node) {
				logger.Error(nil, "Failed to instantly swap UnreachableTaint to NotReadyTaint. Will try again in the next cycle")
			}
		} else if nc.markNodeForTainting(node, v1.ConditionFalse) {
			logger.V(2).Info("Node is NotReady. Adding it to the Taint queue", "node", klog.KObj(node), "timeStamp", decisionTimestamp)
		}
	case v1.ConditionUnknown:
		// We want to update the taint straight away if Node is already tainted with the UnreachableTaint
		if taintutils.TaintExists(node.Spec.Taints, NotReadyTaintTemplate) {
			taintToAdd := *UnreachableTaintTemplate
			if !controllerutil.SwapNodeControllerTaint(ctx, nc.kubeClient, []*v1.Taint{&taintToAdd}, []*v1.Taint{NotReadyTaintTemplate}, node) {
				logger.Error(nil, "Failed to instantly swap NotReadyTaint to UnreachableTaint. Will try again in the next cycle")
			}
		} else if nc.markNodeForTainting(node, v1.ConditionUnknown) {
			logger.V(2).Info("Node is unresponsive. Adding it to the Taint queue", "node", klog.KObj(node), "timeStamp", decisionTimestamp)
		}
	case v1.ConditionTrue:
		removed, err := nc.markNodeAsReachable(ctx, node)
		if err != nil {
			logger.Error(nil, "Failed to remove taints from node. Will retry in next iteration", "node", klog.KObj(node))
		}
		if removed {
			logger.V(2).Info("Node is healthy again, removing all taints", "node", klog.KObj(node))
		}
	}
}

func (nc *Controller) processNoTaintBaseEviction(ctx context.Context, node *v1.Node, observedReadyCondition *v1.NodeCondition, gracePeriod time.Duration, pods []*v1.Pod) error {
	decisionTimestamp := nc.now()
	nodeHealthData := nc.nodeHealthMap.getDeepCopy(node.Name)
	if nodeHealthData == nil {
		return fmt.Errorf("health data doesn't exist for node %q", node.Name)
	}
	// Check eviction timeout against decisionTimestamp
	logger := klog.FromContext(ctx)
	switch observedReadyCondition.Status {
	case v1.ConditionFalse:
		if decisionTimestamp.After(nodeHealthData.readyTransitionTimestamp.Add(nc.podEvictionTimeout)) {
			enqueued, err := nc.evictPods(ctx, node, pods)
			if err != nil {
				return err
			}
			if enqueued {
				logger.V(2).Info("Node is NotReady. Adding Pods on Node to eviction queue: decisionTimestamp is later than readyTransitionTimestamp + podEvictionTimeout",
					"node", klog.KObj(node),
					"decisionTimestamp", decisionTimestamp,
					"readyTransitionTimestamp", nodeHealthData.readyTransitionTimestamp,
					"podEvictionTimeout", nc.podEvictionTimeout,
				)
			}
		}
	case v1.ConditionUnknown:
		if decisionTimestamp.After(nodeHealthData.probeTimestamp.Add(nc.podEvictionTimeout)) {
			enqueued, err := nc.evictPods(ctx, node, pods)
			if err != nil {
				return err
			}
			if enqueued {
				logger.V(2).Info("Node is unresponsive. Adding Pods on Node to eviction queues: decisionTimestamp is later than readyTransitionTimestamp + podEvictionTimeout-gracePeriod",
					"node", klog.KObj(node),
					"decisionTimestamp", decisionTimestamp,
					"readyTransitionTimestamp", nodeHealthData.readyTransitionTimestamp,
					"podEvictionTimeoutGracePeriod", nc.podEvictionTimeout-gracePeriod,
				)
			}
		}
	case v1.ConditionTrue:
		if nc.cancelPodEviction(logger, node) {
			logger.V(2).Info("Node is ready again, cancelled pod eviction", "node", klog.KObj(node))
		}
	}
	return nil
}

// labelNodeDisruptionExclusion is a label on nodes that controls whether they are
// excluded from being considered for disruption checks by the node controller.
const labelNodeDisruptionExclusion = "node.kubernetes.io/exclude-disruption"

func isNodeExcludedFromDisruptionChecks(node *v1.Node) bool {
	if _, ok := node.Labels[labelNodeDisruptionExclusion]; ok {
		return true
	}
	return false
}

// tryUpdateNodeHealth checks a given node's conditions and tries to update it. Returns grace period to
// which given node is entitled, state of current and last observed Ready Condition, and an error if it occurred.
func (nc *Controller) tryUpdateNodeHealth(ctx context.Context, node *v1.Node) (time.Duration, v1.NodeCondition, *v1.NodeCondition, error) {
	nodeHealth := nc.nodeHealthMap.getDeepCopy(node.Name)
	defer func() {
		nc.nodeHealthMap.set(node.Name, nodeHealth)
	}()

	var gracePeriod time.Duration
	var observedReadyCondition v1.NodeCondition
	_, currentReadyCondition := controllerutil.GetNodeCondition(&node.Status, v1.NodeReady)
	if currentReadyCondition == nil {
		// If ready condition is nil, then kubelet (or nodecontroller) never posted node status.
		// A fake ready condition is created, where LastHeartbeatTime and LastTransitionTime is set
		// to node.CreationTimestamp to avoid handle the corner case.
		observedReadyCondition = v1.NodeCondition{
			Type:               v1.NodeReady,
			Status:             v1.ConditionUnknown,
			LastHeartbeatTime:  node.CreationTimestamp,
			LastTransitionTime: node.CreationTimestamp,
		}
		gracePeriod = nc.nodeStartupGracePeriod
		if nodeHealth != nil {
			nodeHealth.status = &node.Status
		} else {
			nodeHealth = &nodeHealthData{
				status:                   &node.Status,
				probeTimestamp:           node.CreationTimestamp,
				readyTransitionTimestamp: node.CreationTimestamp,
			}
		}
	} else {
		// If ready condition is not nil, make a copy of it, since we may modify it in place later.
		observedReadyCondition = *currentReadyCondition
		gracePeriod = nc.nodeMonitorGracePeriod
	}
	// There are following cases to check:
	// - both saved and new status have no Ready Condition set - we leave everything as it is,
	// - saved status have no Ready Condition, but current one does - Controller was restarted with Node data already present in etcd,
	// - saved status have some Ready Condition, but current one does not - it's an error, but we fill it up because that's probably a good thing to do,
	// - both saved and current statuses have Ready Conditions and they have the same LastProbeTime - nothing happened on that Node, it may be
	//   unresponsive, so we leave it as it is,
	// - both saved and current statuses have Ready Conditions, they have different LastProbeTimes, but the same Ready Condition State -
	//   everything's in order, no transition occurred, we update only probeTimestamp,
	// - both saved and current statuses have Ready Conditions, different LastProbeTimes and different Ready Condition State -
	//   Ready Condition changed it state since we last seen it, so we update both probeTimestamp and readyTransitionTimestamp.
	// TODO: things to consider:
	//   - if 'LastProbeTime' have gone back in time its probably an error, currently we ignore it,
	//   - currently only correct Ready State transition outside of Node Controller is marking it ready by Kubelet, we don't check
	//     if that's the case, but it does not seem necessary.
	var savedCondition *v1.NodeCondition
	var savedLease *coordv1.Lease
	if nodeHealth != nil {
		_, savedCondition = controllerutil.GetNodeCondition(nodeHealth.status, v1.NodeReady)
		savedLease = nodeHealth.lease
	}
	logger := klog.FromContext(ctx)
	if nodeHealth == nil {
		logger.Info("Missing timestamp for Node. Assuming now as a timestamp", "node", klog.KObj(node))
		nodeHealth = &nodeHealthData{
			status:                   &node.Status,
			probeTimestamp:           nc.now(),
			readyTransitionTimestamp: nc.now(),
		}
	} else if savedCondition == nil && currentReadyCondition != nil {
		logger.V(1).Info("Creating timestamp entry for newly observed Node", "node", klog.KObj(node))
		nodeHealth = &nodeHealthData{
			status:                   &node.Status,
			probeTimestamp:           nc.now(),
			readyTransitionTimestamp: nc.now(),
		}
	} else if savedCondition != nil && currentReadyCondition == nil {
		logger.Error(nil, "ReadyCondition was removed from Status of Node", "node", klog.KObj(node))
		// TODO: figure out what to do in this case. For now we do the same thing as above.
		nodeHealth = &nodeHealthData{
			status:                   &node.Status,
			probeTimestamp:           nc.now(),
			readyTransitionTimestamp: nc.now(),
		}
	} else if savedCondition != nil && currentReadyCondition != nil && savedCondition.LastHeartbeatTime != currentReadyCondition.LastHeartbeatTime {
		var transitionTime metav1.Time
		// If ReadyCondition changed since the last time we checked, we update the transition timestamp to "now",
		// otherwise we leave it as it is.
		if savedCondition.LastTransitionTime != currentReadyCondition.LastTransitionTime {
			logger.V(3).Info("ReadyCondition for Node transitioned from savedCondition to currentReadyCondition", "node", klog.KObj(node), "savedCondition", savedCondition, "currentReadyCondition", currentReadyCondition)
			transitionTime = nc.now()
		} else {
			transitionTime = nodeHealth.readyTransitionTimestamp
		}
		if loggerV := logger.V(5); loggerV.Enabled() {
			loggerV.Info("Node ReadyCondition updated. Updating timestamp", "node", klog.KObj(node), "nodeHealthStatus", nodeHealth.status, "nodeStatus", node.Status)
		} else {
			logger.V(3).Info("Node ReadyCondition updated. Updating timestamp", "node", klog.KObj(node))
		}
		nodeHealth = &nodeHealthData{
			status:                   &node.Status,
			probeTimestamp:           nc.now(),
			readyTransitionTimestamp: transitionTime,
		}
	}
	// Always update the probe time if node lease is renewed.
	// Note: If kubelet never posted the node status, but continues renewing the
	// heartbeat leases, the node controller will assume the node is healthy and
	// take no action.
	observedLease, _ := nc.leaseLister.Leases(v1.NamespaceNodeLease).Get(node.Name)
	if observedLease != nil && (savedLease == nil || savedLease.Spec.RenewTime.Before(observedLease.Spec.RenewTime)) {
		nodeHealth.lease = observedLease
		nodeHealth.probeTimestamp = nc.now()
	}

	if nc.now().After(nodeHealth.probeTimestamp.Add(gracePeriod)) {
		// NodeReady condition or lease was last set longer ago than gracePeriod, so
		// update it to Unknown (regardless of its current value) in the master.

		nodeConditionTypes := []v1.NodeConditionType{
			v1.NodeReady,
			v1.NodeMemoryPressure,
			v1.NodeDiskPressure,
			v1.NodePIDPressure,
			// We don't change 'NodeNetworkUnavailable' condition, as it's managed on a control plane level.
			// v1.NodeNetworkUnavailable,
		}

		nowTimestamp := nc.now()
		for _, nodeConditionType := range nodeConditionTypes {
			_, currentCondition := controllerutil.GetNodeCondition(&node.Status, nodeConditionType)
			if currentCondition == nil {
				logger.V(2).Info("Condition of node was never updated by kubelet", "nodeConditionType", nodeConditionType, "node", klog.KObj(node))
				node.Status.Conditions = append(node.Status.Conditions, v1.NodeCondition{
					Type:               nodeConditionType,
					Status:             v1.ConditionUnknown,
					Reason:             "NodeStatusNeverUpdated",
					Message:            "Kubelet never posted node status.",
					LastHeartbeatTime:  node.CreationTimestamp,
					LastTransitionTime: nowTimestamp,
				})
			} else {
				logger.V(2).Info("Node hasn't been updated",
					"node", klog.KObj(node), "duration", nc.now().Time.Sub(nodeHealth.probeTimestamp.Time), "nodeConditionType", nodeConditionType, "currentCondition", currentCondition)
				if currentCondition.Status != v1.ConditionUnknown {
					currentCondition.Status = v1.ConditionUnknown
					currentCondition.Reason = "NodeStatusUnknown"
					currentCondition.Message = "Kubelet stopped posting node status."
					currentCondition.LastTransitionTime = nowTimestamp
				}
			}
		}
		// We need to update currentReadyCondition due to its value potentially changed.
		_, currentReadyCondition = controllerutil.GetNodeCondition(&node.Status, v1.NodeReady)

		if !apiequality.Semantic.DeepEqual(currentReadyCondition, &observedReadyCondition) {
			if _, err := nc.kubeClient.CoreV1().Nodes().UpdateStatus(ctx, node, metav1.UpdateOptions{}); err != nil {
				logger.Error(err, "Error updating node", "node", klog.KObj(node))
				return gracePeriod, observedReadyCondition, currentReadyCondition, err
			}
			nodeHealth = &nodeHealthData{
				status:                   &node.Status,
				probeTimestamp:           nodeHealth.probeTimestamp,
				readyTransitionTimestamp: nc.now(),
				lease:                    observedLease,
			}
			return gracePeriod, observedReadyCondition, currentReadyCondition, nil
		}
	}

	return gracePeriod, observedReadyCondition, currentReadyCondition, nil
}

func (nc *Controller) handleDisruption(ctx context.Context, zoneToNodeConditions map[string][]*v1.NodeCondition, nodes []*v1.Node) {
	newZoneStates := map[string]ZoneState{}
	allAreFullyDisrupted := true
	logger := klog.FromContext(ctx)
	for k, v := range zoneToNodeConditions {
		zoneSize.WithLabelValues(k).Set(float64(len(v)))
		unhealthy, newState := nc.computeZoneStateFunc(v)
		zoneHealth.WithLabelValues(k).Set(float64(100*(len(v)-unhealthy)) / float64(len(v)))
		unhealthyNodes.WithLabelValues(k).Set(float64(unhealthy))
		if newState != stateFullDisruption {
			allAreFullyDisrupted = false
		}
		newZoneStates[k] = newState
		if _, had := nc.zoneStates[k]; !had {
			logger.Error(nil, "Setting initial state for unseen zone", "zone", k)
			nc.zoneStates[k] = stateInitial
		}
	}

	allWasFullyDisrupted := true
	for k, v := range nc.zoneStates {
		if _, have := zoneToNodeConditions[k]; !have {
			zoneSize.WithLabelValues(k).Set(0)
			zoneHealth.WithLabelValues(k).Set(100)
			unhealthyNodes.WithLabelValues(k).Set(0)
			delete(nc.zoneStates, k)
			continue
		}
		if v != stateFullDisruption {
			allWasFullyDisrupted = false
			break
		}
	}

	// At least one node was responding in previous pass or in the current pass. Semantics is as follows:
	// - if the new state is "partialDisruption" we call a user defined function that returns a new limiter to use,
	// - if the new state is "normal" we resume normal operation (go back to default limiter settings),
	// - if new state is "fullDisruption" we restore normal eviction rate,
	//   - unless all zones in the cluster are in "fullDisruption" - in that case we stop all evictions.
	if !allAreFullyDisrupted || !allWasFullyDisrupted {
		// We're switching to full disruption mode
		if allAreFullyDisrupted {
			logger.Info("Controller detected that all Nodes are not-Ready. Entering master disruption mode")
			for i := range nodes {
				if nc.runTaintManager {
					_, err := nc.markNodeAsReachable(ctx, nodes[i])
					if err != nil {
						logger.Error(nil, "Failed to remove taints from Node", "node", klog.KObj(nodes[i]))
					}
				} else {
					nc.cancelPodEviction(logger, nodes[i])
				}
			}
			// We stop all evictions.
			for k := range nc.zoneStates {
				if nc.runTaintManager {
					nc.zoneNoExecuteTainter[k].SwapLimiter(0)
				} else {
					nc.zonePodEvictor[k].SwapLimiter(0)
				}
			}
			for k := range nc.zoneStates {
				nc.zoneStates[k] = stateFullDisruption
			}
			// All rate limiters are updated, so we can return early here.
			return
		}
		// We're exiting full disruption mode
		if allWasFullyDisrupted {
			logger.Info("Controller detected that some Nodes are Ready. Exiting master disruption mode")
			// When exiting disruption mode update probe timestamps on all Nodes.
			now := nc.now()
			for i := range nodes {
				v := nc.nodeHealthMap.getDeepCopy(nodes[i].Name)
				v.probeTimestamp = now
				v.readyTransitionTimestamp = now
				nc.nodeHealthMap.set(nodes[i].Name, v)
			}
			// We reset all rate limiters to settings appropriate for the given state.
			for k := range nc.zoneStates {
				nc.setLimiterInZone(k, len(zoneToNodeConditions[k]), newZoneStates[k])
				nc.zoneStates[k] = newZoneStates[k]
			}
			return
		}
		// We know that there's at least one not-fully disrupted so,
		// we can use default behavior for rate limiters
		for k, v := range nc.zoneStates {
			newState := newZoneStates[k]
			if v == newState {
				continue
			}
			logger.Info("Controller detected that zone is now in new state", "zone", k, "newState", newState)
			nc.setLimiterInZone(k, len(zoneToNodeConditions[k]), newState)
			nc.zoneStates[k] = newState
		}
	}
}

func (nc *Controller) podUpdated(oldPod, newPod *v1.Pod) {
	if newPod == nil {
		return
	}
	if len(newPod.Spec.NodeName) != 0 && (oldPod == nil || newPod.Spec.NodeName != oldPod.Spec.NodeName) {
		podItem := podUpdateItem{newPod.Namespace, newPod.Name}
		nc.podUpdateQueue.Add(podItem)
	}
}

func (nc *Controller) doPodProcessingWorker(ctx context.Context) {
	for {
		obj, shutdown := nc.podUpdateQueue.Get()
		// "podUpdateQueue" will be shutdown when "stopCh" closed;
		// we do not need to re-check "stopCh" again.
		if shutdown {
			return
		}

		podItem := obj.(podUpdateItem)
		nc.processPod(ctx, podItem)
	}
}

// processPod is processing events of assigning pods to nodes. In particular:
// 1. for NodeReady=true node, taint eviction for this pod will be cancelled
// 2. for NodeReady=false or unknown node, taint eviction of pod will happen and pod will be marked as not ready
// 3. if node doesn't exist in cache, it will be skipped and handled later by doEvictionPass
func (nc *Controller) processPod(ctx context.Context, podItem podUpdateItem) {
	defer nc.podUpdateQueue.Done(podItem)
	pod, err := nc.podLister.Pods(podItem.namespace).Get(podItem.name)
	logger := klog.FromContext(ctx)
	if err != nil {
		if apierrors.IsNotFound(err) {
			// If the pod was deleted, there is no need to requeue.
			return
		}
		logger.Info("Failed to read pod", "pod", klog.KRef(podItem.namespace, podItem.name), "err", err)
		nc.podUpdateQueue.AddRateLimited(podItem)
		return
	}

	nodeName := pod.Spec.NodeName

	nodeHealth := nc.nodeHealthMap.getDeepCopy(nodeName)
	if nodeHealth == nil {
		// Node data is not gathered yet or node has beed removed in the meantime.
		// Pod will be handled by doEvictionPass method.
		return
	}

	node, err := nc.nodeLister.Get(nodeName)
	if err != nil {
		logger.Info("Failed to read node", "node", klog.KRef("", nodeName), "err", err)
		nc.podUpdateQueue.AddRateLimited(podItem)
		return
	}

	_, currentReadyCondition := controllerutil.GetNodeCondition(nodeHealth.status, v1.NodeReady)
	if currentReadyCondition == nil {
		// Lack of NodeReady condition may only happen after node addition (or if it will be maliciously deleted).
		// In both cases, the pod will be handled correctly (evicted if needed) during processing
		// of the next node update event.
		return
	}

	pods := []*v1.Pod{pod}
	// In taint-based eviction mode, only node updates are processed by NodeLifecycleController.
	// Pods are processed by TaintManager.
	if !nc.runTaintManager {
		if err := nc.processNoTaintBaseEviction(ctx, node, currentReadyCondition, nc.nodeMonitorGracePeriod, pods); err != nil {
			logger.Info("Unable to process pod eviction from node", "pod", klog.KRef(podItem.namespace, podItem.name), "node", klog.KRef("", nodeName), "err", err)
			nc.podUpdateQueue.AddRateLimited(podItem)
			return
		}
	}

	if currentReadyCondition.Status != v1.ConditionTrue {
		if err := controllerutil.MarkPodsNotReady(ctx, nc.kubeClient, nc.recorder, pods, nodeName); err != nil {
			logger.Info("Unable to mark pod NotReady on node", "pod", klog.KRef(podItem.namespace, podItem.name), "node", klog.KRef("", nodeName), "err", err)
			nc.podUpdateQueue.AddRateLimited(podItem)
		}
	}
}

func (nc *Controller) setLimiterInZone(zone string, zoneSize int, state ZoneState) {
	switch state {
	case stateNormal:
		if nc.runTaintManager {
			nc.zoneNoExecuteTainter[zone].SwapLimiter(nc.evictionLimiterQPS)
		} else {
			nc.zonePodEvictor[zone].SwapLimiter(nc.evictionLimiterQPS)
		}
	case statePartialDisruption:
		if nc.runTaintManager {
			nc.zoneNoExecuteTainter[zone].SwapLimiter(
				nc.enterPartialDisruptionFunc(zoneSize))
		} else {
			nc.zonePodEvictor[zone].SwapLimiter(
				nc.enterPartialDisruptionFunc(zoneSize))
		}
	case stateFullDisruption:
		if nc.runTaintManager {
			nc.zoneNoExecuteTainter[zone].SwapLimiter(
				nc.enterFullDisruptionFunc(zoneSize))
		} else {
			nc.zonePodEvictor[zone].SwapLimiter(
				nc.enterFullDisruptionFunc(zoneSize))
		}
	}
}

// classifyNodes classifies the allNodes to three categories:
//  1. added: the nodes that in 'allNodes', but not in 'knownNodeSet'
//  2. deleted: the nodes that in 'knownNodeSet', but not in 'allNodes'
//  3. newZoneRepresentatives: the nodes that in both 'knownNodeSet' and 'allNodes', but no zone states
func (nc *Controller) classifyNodes(allNodes []*v1.Node) (added, deleted, newZoneRepresentatives []*v1.Node) {
	for i := range allNodes {
		if _, has := nc.knownNodeSet[allNodes[i].Name]; !has {
			added = append(added, allNodes[i])
		} else {
			// Currently, we only consider new zone as updated.
			zone := nodetopology.GetZoneKey(allNodes[i])
			if _, found := nc.zoneStates[zone]; !found {
				newZoneRepresentatives = append(newZoneRepresentatives, allNodes[i])
			}
		}
	}

	// If there's a difference between lengths of known Nodes and observed nodes
	// we must have removed some Node.
	if len(nc.knownNodeSet)+len(added) != len(allNodes) {
		knowSetCopy := map[string]*v1.Node{}
		for k, v := range nc.knownNodeSet {
			knowSetCopy[k] = v
		}
		for i := range allNodes {
			delete(knowSetCopy, allNodes[i].Name)
		}
		for i := range knowSetCopy {
			deleted = append(deleted, knowSetCopy[i])
		}
	}
	return
}

// HealthyQPSFunc returns the default value for cluster eviction rate - we take
// nodeNum for consistency with ReducedQPSFunc.
func (nc *Controller) HealthyQPSFunc(nodeNum int) float32 {
	return nc.evictionLimiterQPS
}

// ReducedQPSFunc returns the QPS for when a the cluster is large make
// evictions slower, if they're small stop evictions altogether.
func (nc *Controller) ReducedQPSFunc(nodeNum int) float32 {
	if int32(nodeNum) > nc.largeClusterThreshold {
		return nc.secondaryEvictionLimiterQPS
	}
	return 0
}

// addPodEvictorForNewZone checks if new zone appeared, and if so add new evictor.
func (nc *Controller) addPodEvictorForNewZone(logger klog.Logger, node *v1.Node) {
	nc.evictorLock.Lock()
	defer nc.evictorLock.Unlock()
	zone := nodetopology.GetZoneKey(node)
	if _, found := nc.zoneStates[zone]; !found {
		nc.zoneStates[zone] = stateInitial
		if !nc.runTaintManager {
			nc.zonePodEvictor[zone] =
				scheduler.NewRateLimitedTimedQueue(
					flowcontrol.NewTokenBucketRateLimiter(nc.evictionLimiterQPS, scheduler.EvictionRateLimiterBurst))
		} else {
			nc.zoneNoExecuteTainter[zone] =
				scheduler.NewRateLimitedTimedQueue(
					flowcontrol.NewTokenBucketRateLimiter(nc.evictionLimiterQPS, scheduler.EvictionRateLimiterBurst))
		}
		// Init the metric for the new zone.
		logger.Info("Initializing eviction metric for zone", "zone", zone)
		evictionsNumber.WithLabelValues(zone).Add(0)
		evictionsTotal.WithLabelValues(zone).Add(0)
	}
}

// cancelPodEviction removes any queued evictions, typically because the node is available again. It
// returns true if an eviction was queued.
func (nc *Controller) cancelPodEviction(logger klog.Logger, node *v1.Node) bool {
	zone := nodetopology.GetZoneKey(node)
	if !nc.nodeEvictionMap.setStatus(node.Name, unmarked) {
		logger.V(2).Info("Node was unregistered in the meantime - skipping setting status", "node", klog.KObj(node))
	}
	nc.evictorLock.Lock()
	defer nc.evictorLock.Unlock()
	wasDeleting := nc.zonePodEvictor[zone].Remove(node.Name)
	if wasDeleting {
		logger.V(2).Info("Cancelling pod Eviction on Node", "node", klog.KObj(node))
		return true
	}
	return false
}

// evictPods:
//   - adds node to evictor queue if the node is not marked as evicted.
//     Returns false if the node name was already enqueued.
//   - deletes pods immediately if node is already marked as evicted.
//     Returns false, because the node wasn't added to the queue.
func (nc *Controller) evictPods(ctx context.Context, node *v1.Node, pods []*v1.Pod) (bool, error) {
	status, ok := nc.nodeEvictionMap.getStatus(node.Name)
	if ok && status == evicted {
		// Node eviction already happened for this node.
		// Handling immediate pod deletion.
		_, err := controllerutil.DeletePods(ctx, nc.kubeClient, pods, nc.recorder, node.Name, string(node.UID), nc.daemonSetStore)
		if err != nil {
			return false, fmt.Errorf("unable to delete pods from node %q: %v", node.Name, err)
		}
		return false, nil
	}
	logger := klog.FromContext(ctx)
	if !nc.nodeEvictionMap.setStatus(node.Name, toBeEvicted) {
		logger.V(2).Info("Node was unregistered in the meantime - skipping setting status", "node", klog.KObj(node))
	}

	nc.evictorLock.Lock()
	defer nc.evictorLock.Unlock()

	return nc.zonePodEvictor[nodetopology.GetZoneKey(node)].Add(node.Name, string(node.UID)), nil
}

func (nc *Controller) markNodeForTainting(node *v1.Node, status v1.ConditionStatus) bool {
	nc.evictorLock.Lock()
	defer nc.evictorLock.Unlock()
	if status == v1.ConditionFalse {
		if !taintutils.TaintExists(node.Spec.Taints, NotReadyTaintTemplate) {
			nc.zoneNoExecuteTainter[nodetopology.GetZoneKey(node)].Remove(node.Name)
		}
	}

	if status == v1.ConditionUnknown {
		if !taintutils.TaintExists(node.Spec.Taints, UnreachableTaintTemplate) {
			nc.zoneNoExecuteTainter[nodetopology.GetZoneKey(node)].Remove(node.Name)
		}
	}

	return nc.zoneNoExecuteTainter[nodetopology.GetZoneKey(node)].Add(node.Name, string(node.UID))
}

func (nc *Controller) markNodeAsReachable(ctx context.Context, node *v1.Node) (bool, error) {
	err := controller.RemoveTaintOffNode(ctx, nc.kubeClient, node.Name, node, UnreachableTaintTemplate)
	logger := klog.FromContext(ctx)
	if err != nil {
		logger.Error(err, "Failed to remove taint from node", "node", klog.KObj(node))
		return false, err
	}
	err = controller.RemoveTaintOffNode(ctx, nc.kubeClient, node.Name, node, NotReadyTaintTemplate)
	if err != nil {
		logger.Error(err, "Failed to remove taint from node", "node", klog.KObj(node))
		return false, err
	}
	nc.evictorLock.Lock()
	defer nc.evictorLock.Unlock()

	return nc.zoneNoExecuteTainter[nodetopology.GetZoneKey(node)].Remove(node.Name), nil
}

// ComputeZoneState returns a slice of NodeReadyConditions for all Nodes in a given zone.
// The zone is considered:
// - fullyDisrupted if there're no Ready Nodes,
// - partiallyDisrupted if at least than nc.unhealthyZoneThreshold percent of Nodes are not Ready,
// - normal otherwise
func (nc *Controller) ComputeZoneState(nodeReadyConditions []*v1.NodeCondition) (int, ZoneState) {
	readyNodes := 0
	notReadyNodes := 0
	for i := range nodeReadyConditions {
		if nodeReadyConditions[i] != nil && nodeReadyConditions[i].Status == v1.ConditionTrue {
			readyNodes++
		} else {
			notReadyNodes++
		}
	}
	switch {
	case readyNodes == 0 && notReadyNodes > 0:
		return notReadyNodes, stateFullDisruption
	case notReadyNodes > 2 && float32(notReadyNodes)/float32(notReadyNodes+readyNodes) >= nc.unhealthyZoneThreshold:
		return notReadyNodes, statePartialDisruption
	default:
		return notReadyNodes, stateNormal
	}
}

// reconcileNodeLabels reconciles node labels.
func (nc *Controller) reconcileNodeLabels(nodeName string) error {
	node, err := nc.nodeLister.Get(nodeName)
	if err != nil {
		// If node not found, just ignore it.
		if apierrors.IsNotFound(err) {
			return nil
		}
		return err
	}

	if node.Labels == nil {
		// Nothing to reconcile.
		return nil
	}

	labelsToUpdate := map[string]string{}
	for _, r := range labelReconcileInfo {
		primaryValue, primaryExists := node.Labels[r.primaryKey]
		secondaryValue, secondaryExists := node.Labels[r.secondaryKey]

		if !primaryExists {
			// The primary label key does not exist. This should not happen
			// within our supported version skew range, when no external
			// components/factors modifying the node object. Ignore this case.
			continue
		}
		if secondaryExists && primaryValue != secondaryValue {
			// Secondary label exists, but not consistent with the primary
			// label. Need to reconcile.
			labelsToUpdate[r.secondaryKey] = primaryValue

		} else if !secondaryExists && r.ensureSecondaryExists {
			// Apply secondary label based on primary label.
			labelsToUpdate[r.secondaryKey] = primaryValue
		}
	}

	if len(labelsToUpdate) == 0 {
		return nil
	}
	if !controllerutil.AddOrUpdateLabelsOnNode(nc.kubeClient, labelsToUpdate, node) {
		return fmt.Errorf("failed update labels for node %+v", node)
	}
	return nil
}
