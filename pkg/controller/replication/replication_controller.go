/*
Copyright 2014 The Kubernetes Authors.

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

// If you make changes to this file, you should also make the corresponding change in ReplicaSet.

package replication

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
	"time"

	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	utiltrace "k8s.io/apiserver/pkg/util/trace"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	clientv1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
	coreinformers "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/core/v1"
	corelisters "k8s.io/kubernetes/pkg/client/listers/core/v1"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/util/metrics"
)

const (
	// Realistic value of the burstReplica field for the replication manager based off
	// performance requirements for kubernetes 1.0.
	BurstReplicas = 500

	// The number of times we retry updating a replication controller's status.
	statusUpdateRetries = 1
)

func getRCKind() schema.GroupVersionKind {
	return v1.SchemeGroupVersion.WithKind("ReplicationController")
}

// ReplicationManager is responsible for synchronizing ReplicationController objects stored
// in the system with actual running pods.
// TODO: this really should be called ReplicationController. The only reason why it's a Manager
// is to distinguish this type from API object "ReplicationController". We should fix this.
type ReplicationManager struct {
	kubeClient clientset.Interface
	podControl controller.PodControlInterface

	// An rc is temporarily suspended after creating/deleting these many replicas.
	// It resumes normal action after observing the watch events for them.
	burstReplicas int
	// To allow injection of syncReplicationController for testing.
	syncHandler func(rcKey string) error

	// A TTLCache of pod creates/deletes each rc expects to see.
	expectations *controller.UIDTrackingControllerExpectations

	rcLister       corelisters.ReplicationControllerLister
	rcListerSynced cache.InformerSynced

	podLister corelisters.PodLister
	// podListerSynced returns true if the pod store has been synced at least once.
	// Added as a member to the struct to allow injection for testing.
	podListerSynced cache.InformerSynced

	lookupCache *controller.MatchingCache

	// Controllers that need to be synced
	queue workqueue.RateLimitingInterface
}

// NewReplicationManager configures a replication manager with the specified event recorder
func NewReplicationManager(podInformer coreinformers.PodInformer, rcInformer coreinformers.ReplicationControllerInformer, kubeClient clientset.Interface, burstReplicas int, lookupCacheSize int) *ReplicationManager {
	if kubeClient != nil && kubeClient.Core().RESTClient().GetRateLimiter() != nil {
		metrics.RegisterMetricAndTrackRateLimiterUsage("replication_controller", kubeClient.Core().RESTClient().GetRateLimiter())
	}

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: v1core.New(kubeClient.Core().RESTClient()).Events("")})

	rm := &ReplicationManager{
		kubeClient: kubeClient,
		podControl: controller.RealPodControl{
			KubeClient: kubeClient,
			Recorder:   eventBroadcaster.NewRecorder(api.Scheme, clientv1.EventSource{Component: "replication-controller"}),
		},
		burstReplicas: burstReplicas,
		expectations:  controller.NewUIDTrackingControllerExpectations(controller.NewControllerExpectations()),
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "replicationmanager"),
	}

	rcInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    rm.enqueueController,
		UpdateFunc: rm.updateRC,
		// This will enter the sync loop and no-op, because the controller has been deleted from the store.
		// Note that deleting a controller immediately after scaling it to 0 will not work. The recommended
		// way of achieving this is by performing a `stop` operation on the controller.
		DeleteFunc: rm.enqueueController,
	})
	rm.rcLister = rcInformer.Lister()
	rm.rcListerSynced = rcInformer.Informer().HasSynced

	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: rm.addPod,
		// This invokes the rc for every pod change, eg: host assignment. Though this might seem like overkill
		// the most frequent pod update is status, and the associated rc will only list from local storage, so
		// it should be ok.
		UpdateFunc: rm.updatePod,
		DeleteFunc: rm.deletePod,
	})
	rm.podLister = podInformer.Lister()
	rm.podListerSynced = podInformer.Informer().HasSynced

	rm.syncHandler = rm.syncReplicationController
	rm.lookupCache = controller.NewMatchingCache(lookupCacheSize)
	return rm
}

// SetEventRecorder replaces the event recorder used by the replication manager
// with the given recorder. Only used for testing.
func (rm *ReplicationManager) SetEventRecorder(recorder record.EventRecorder) {
	// TODO: Hack. We can't cleanly shutdown the event recorder, so benchmarks
	// need to pass in a fake.
	rm.podControl = controller.RealPodControl{KubeClient: rm.kubeClient, Recorder: recorder}
}

// Run begins watching and syncing.
func (rm *ReplicationManager) Run(workers int, stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	defer rm.queue.ShutDown()

	glog.Infof("Starting RC Manager")

	if !cache.WaitForCacheSync(stopCh, rm.podListerSynced, rm.rcListerSynced) {
		utilruntime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}

	for i := 0; i < workers; i++ {
		go wait.Until(rm.worker, time.Second, stopCh)
	}

	<-stopCh
	glog.Infof("Shutting down RC Manager")
}

// getPodController returns the controller managing the given pod.
// TODO: Surface that we are ignoring multiple controllers for a single pod.
// TODO: use ownerReference.Controller to determine if the rc controls the pod.
func (rm *ReplicationManager) getPodController(pod *v1.Pod) *v1.ReplicationController {
	// look up in the cache, if cached and the cache is valid, just return cached value
	if obj, cached := rm.lookupCache.GetMatchingObject(pod); cached {
		controller, ok := obj.(*v1.ReplicationController)
		if !ok {
			// This should not happen
			utilruntime.HandleError(fmt.Errorf("lookup cache does not return a ReplicationController object"))
			return nil
		}
		if cached && rm.isCacheValid(pod, controller) {
			return controller
		}
	}

	// if not cached or cached value is invalid, search all the rc to find the matching one, and update cache
	controllers, err := rm.rcLister.GetPodControllers(pod)
	if err != nil {
		glog.V(4).Infof("No controllers found for pod %v, replication manager will avoid syncing", pod.Name)
		return nil
	}
	// In theory, overlapping controllers is user error. This sorting will not prevent
	// oscillation of replicas in all cases, eg:
	// rc1 (older rc): [(k1=v1)], replicas=1 rc2: [(k2=v2)], replicas=2
	// pod: [(k1:v1), (k2:v2)] will wake both rc1 and rc2, and we will sync rc1.
	// pod: [(k2:v2)] will wake rc2 which creates a new replica.
	if len(controllers) > 1 {
		// More than two items in this list indicates user error. If two replication-controller
		// overlap, sort by creation timestamp, subsort by name, then pick
		// the first.
		utilruntime.HandleError(fmt.Errorf("user error! more than one replication controller is selecting pods with labels: %+v", pod.Labels))
		sort.Sort(OverlappingControllers(controllers))
	}

	// update lookup cache
	rm.lookupCache.Update(pod, controllers[0])

	return controllers[0]
}

// isCacheValid check if the cache is valid
func (rm *ReplicationManager) isCacheValid(pod *v1.Pod, cachedRC *v1.ReplicationController) bool {
	_, err := rm.rcLister.ReplicationControllers(cachedRC.Namespace).Get(cachedRC.Name)
	// rc has been deleted or updated, cache is invalid
	if err != nil || !isControllerMatch(pod, cachedRC) {
		return false
	}
	return true
}

// isControllerMatch take a Pod and ReplicationController, return whether the Pod and ReplicationController are matching
// TODO(mqliang): This logic is a copy from GetPodControllers(), remove the duplication
func isControllerMatch(pod *v1.Pod, rc *v1.ReplicationController) bool {
	if rc.Namespace != pod.Namespace {
		return false
	}
	selector := labels.Set(rc.Spec.Selector).AsSelectorPreValidated()

	// If an rc with a nil or empty selector creeps in, it should match nothing, not everything.
	if selector.Empty() || !selector.Matches(labels.Set(pod.Labels)) {
		return false
	}
	return true
}

// callback when RC is updated
func (rm *ReplicationManager) updateRC(old, cur interface{}) {
	oldRC := old.(*v1.ReplicationController)
	curRC := cur.(*v1.ReplicationController)

	// We should invalidate the whole lookup cache if a RC's selector has been updated.
	//
	// Imagine that you have two RCs:
	// * old RC1
	// * new RC2
	// You also have a pod that is attached to RC2 (because it doesn't match RC1 selector).
	// Now imagine that you are changing RC1 selector so that it is now matching that pod,
	// in such case, we must invalidate the whole cache so that pod could be adopted by RC1
	//
	// This makes the lookup cache less helpful, but selector update does not happen often,
	// so it's not a big problem
	if !reflect.DeepEqual(oldRC.Spec.Selector, curRC.Spec.Selector) {
		rm.lookupCache.InvalidateAll()
	}
	// TODO: Remove when #31981 is resolved!
	glog.Infof("Observed updated replication controller %v. Desired pod count change: %d->%d", curRC.Name, *(oldRC.Spec.Replicas), *(curRC.Spec.Replicas))

	// You might imagine that we only really need to enqueue the
	// controller when Spec changes, but it is safer to sync any
	// time this function is triggered. That way a full informer
	// resync can requeue any controllers that don't yet have pods
	// but whose last attempts at creating a pod have failed (since
	// we don't block on creation of pods) instead of those
	// controllers stalling indefinitely. Enqueueing every time
	// does result in some spurious syncs (like when Status.Replica
	// is updated and the watch notification from it retriggers
	// this function), but in general extra resyncs shouldn't be
	// that bad as rcs that haven't met expectations yet won't
	// sync, and all the listing is done using local stores.
	if oldRC.Status.Replicas != curRC.Status.Replicas {
		// TODO: Should we log status or spec?
		glog.V(4).Infof("Observed updated replica count for rc: %v, %d->%d", curRC.Name, oldRC.Status.Replicas, curRC.Status.Replicas)
	}
	rm.enqueueController(cur)
}

// When a pod is created, enqueue the controller that manages it and update it's expectations.
func (rm *ReplicationManager) addPod(obj interface{}) {
	pod := obj.(*v1.Pod)

	rc := rm.getPodController(pod)
	if rc == nil {
		return
	}
	rcKey, err := controller.KeyFunc(rc)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for replication controller %#v: %v", rc, err))
		return
	}

	if pod.DeletionTimestamp != nil {
		// on a restart of the controller manager, it's possible a new pod shows up in a state that
		// is already pending deletion. Prevent the pod from being a creation observation.
		rm.deletePod(pod)
		return
	}
	rm.expectations.CreationObserved(rcKey)
	rm.enqueueController(rc)
}

// When a pod is updated, figure out what controller/s manage it and wake them
// up. If the labels of the pod have changed we need to awaken both the old
// and new controller. old and cur must be *v1.Pod types.
func (rm *ReplicationManager) updatePod(old, cur interface{}) {
	curPod := cur.(*v1.Pod)
	oldPod := old.(*v1.Pod)
	if curPod.ResourceVersion == oldPod.ResourceVersion {
		// Periodic resync will send update events for all known pods.
		// Two different versions of the same pod will always have different RVs.
		return
	}
	glog.V(4).Infof("Pod %s updated, objectMeta %+v -> %+v.", curPod.Name, oldPod.ObjectMeta, curPod.ObjectMeta)
	labelChanged := !reflect.DeepEqual(curPod.Labels, oldPod.Labels)
	if curPod.DeletionTimestamp != nil {
		// when a pod is deleted gracefully it's deletion timestamp is first modified to reflect a grace period,
		// and after such time has passed, the kubelet actually deletes it from the store. We receive an update
		// for modification of the deletion timestamp and expect an rc to create more replicas asap, not wait
		// until the kubelet actually deletes the pod. This is different from the Phase of a pod changing, because
		// an rc never initiates a phase change, and so is never asleep waiting for the same.
		rm.deletePod(curPod)
		if labelChanged {
			// we don't need to check the oldPod.DeletionTimestamp because DeletionTimestamp cannot be unset.
			rm.deletePod(oldPod)
		}
		return
	}

	// Only need to get the old controller if the labels changed.
	// Enqueue the oldRC before the curRC to give curRC a chance to adopt the oldPod.
	if labelChanged {
		// If the old and new rc are the same, the first one that syncs
		// will set expectations preventing any damage from the second.
		if oldRC := rm.getPodController(oldPod); oldRC != nil {
			rm.enqueueController(oldRC)
		}
	}

	changedToReady := !v1.IsPodReady(oldPod) && v1.IsPodReady(curPod)
	if curRC := rm.getPodController(curPod); curRC != nil {
		rm.enqueueController(curRC)
		// TODO: MinReadySeconds in the Pod will generate an Available condition to be added in
		// the Pod status which in turn will trigger a requeue of the owning replication controller
		// thus having its status updated with the newly available replica. For now, we can fake the
		// update by resyncing the controller MinReadySeconds after the it is requeued because a Pod
		// transitioned to Ready.
		// Note that this still suffers from #29229, we are just moving the problem one level
		// "closer" to kubelet (from the deployment to the replication controller manager).
		if changedToReady && curRC.Spec.MinReadySeconds > 0 {
			glog.V(2).Infof("ReplicationController %q will be enqueued after %ds for availability check", curRC.Name, curRC.Spec.MinReadySeconds)
			rm.enqueueControllerAfter(curRC, time.Duration(curRC.Spec.MinReadySeconds)*time.Second)
		}
	}
}

// When a pod is deleted, enqueue the controller that manages the pod and update its expectations.
// obj could be an *v1.Pod, or a DeletionFinalStateUnknown marker item.
func (rm *ReplicationManager) deletePod(obj interface{}) {
	pod, ok := obj.(*v1.Pod)

	// When a delete is dropped, the relist will notice a pod in the store not
	// in the list, leading to the insertion of a tombstone object which contains
	// the deleted key/value. Note that this value might be stale. If the pod
	// changed labels the new rc will not be woken up till the periodic resync.
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
			return
		}
		pod, ok = tombstone.Obj.(*v1.Pod)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a pod %#v", obj))
			return
		}
	}
	glog.V(4).Infof("Pod %s/%s deleted through %v, timestamp %+v, labels %+v.", pod.Namespace, pod.Name, utilruntime.GetCaller(), pod.DeletionTimestamp, pod.Labels)
	if rc := rm.getPodController(pod); rc != nil {
		rcKey, err := controller.KeyFunc(rc)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Couldn't get key for replication controller %#v: %v", rc, err))
			return
		}
		rm.expectations.DeletionObserved(rcKey, controller.PodKey(pod))
		rm.enqueueController(rc)
	}
}

// obj could be an *v1.ReplicationController, or a DeletionFinalStateUnknown marker item.
func (rm *ReplicationManager) enqueueController(obj interface{}) {
	key, err := controller.KeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %+v: %v", obj, err))
		return
	}

	// TODO: Handle overlapping controllers better. Either disallow them at admission time or
	// deterministically avoid syncing controllers that fight over pods. Currently, we only
	// ensure that the same controller is synced for a given pod. When we periodically relist
	// all controllers there will still be some replica instability. One way to handle this is
	// by querying the store for all controllers that this rc overlaps, as well as all
	// controllers that overlap this rc, and sorting them.
	rm.queue.Add(key)
}

// obj could be an *v1.ReplicationController, or a DeletionFinalStateUnknown marker item.
func (rm *ReplicationManager) enqueueControllerAfter(obj interface{}, after time.Duration) {
	key, err := controller.KeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %+v: %v", obj, err))
		return
	}

	// TODO: Handle overlapping controllers better. Either disallow them at admission time or
	// deterministically avoid syncing controllers that fight over pods. Currently, we only
	// ensure that the same controller is synced for a given pod. When we periodically relist
	// all controllers there will still be some replica instability. One way to handle this is
	// by querying the store for all controllers that this rc overlaps, as well as all
	// controllers that overlap this rc, and sorting them.
	rm.queue.AddAfter(key, after)
}

// worker runs a worker thread that just dequeues items, processes them, and marks them done.
// It enforces that the syncHandler is never invoked concurrently with the same key.
func (rm *ReplicationManager) worker() {
	workFunc := func() bool {
		key, quit := rm.queue.Get()
		if quit {
			return true
		}
		defer rm.queue.Done(key)

		err := rm.syncHandler(key.(string))
		if err == nil {
			rm.queue.Forget(key)
			return false
		}

		rm.queue.AddRateLimited(key)
		utilruntime.HandleError(err)
		return false
	}
	for {
		if quit := workFunc(); quit {
			glog.Infof("replication controller worker shutting down")
			return
		}
	}
}

// manageReplicas checks and updates replicas for the given replication controller.
// Does NOT modify <filteredPods>.
func (rm *ReplicationManager) manageReplicas(filteredPods []*v1.Pod, rc *v1.ReplicationController) error {
	diff := len(filteredPods) - int(*(rc.Spec.Replicas))
	rcKey, err := controller.KeyFunc(rc)
	if err != nil {
		return err
	}
	if diff == 0 {
		return nil
	}

	if diff < 0 {
		diff *= -1
		if diff > rm.burstReplicas {
			diff = rm.burstReplicas
		}
		// TODO: Track UIDs of creates just like deletes. The problem currently
		// is we'd need to wait on the result of a create to record the pod's
		// UID, which would require locking *across* the create, which will turn
		// into a performance bottleneck. We should generate a UID for the pod
		// beforehand and store it via ExpectCreations.
		errCh := make(chan error, diff)
		rm.expectations.ExpectCreations(rcKey, diff)
		var wg sync.WaitGroup
		wg.Add(diff)
		glog.V(2).Infof("Too few %q/%q replicas, need %d, creating %d", rc.Namespace, rc.Name, *(rc.Spec.Replicas), diff)
		for i := 0; i < diff; i++ {
			go func() {
				defer wg.Done()
				var err error
				var trueVar = true
				controllerRef := &metav1.OwnerReference{
					APIVersion: getRCKind().GroupVersion().String(),
					Kind:       getRCKind().Kind,
					Name:       rc.Name,
					UID:        rc.UID,
					Controller: &trueVar,
				}
				err = rm.podControl.CreatePodsWithControllerRef(rc.Namespace, rc.Spec.Template, rc, controllerRef)
				if err != nil {
					// Decrement the expected number of creates because the informer won't observe this pod
					glog.V(2).Infof("Failed creation, decrementing expectations for controller %q/%q", rc.Namespace, rc.Name)
					rm.expectations.CreationObserved(rcKey)
					errCh <- err
					utilruntime.HandleError(err)
				}
			}()
		}
		wg.Wait()

		select {
		case err := <-errCh:
			// all errors have been reported before and they're likely to be the same, so we'll only return the first one we hit.
			if err != nil {
				return err
			}
		default:
		}

		return nil
	}

	if diff > rm.burstReplicas {
		diff = rm.burstReplicas
	}
	glog.V(2).Infof("Too many %q/%q replicas, need %d, deleting %d", rc.Namespace, rc.Name, *(rc.Spec.Replicas), diff)
	// No need to sort pods if we are about to delete all of them
	if *(rc.Spec.Replicas) != 0 {
		// Sort the pods in the order such that not-ready < ready, unscheduled
		// < scheduled, and pending < running. This ensures that we delete pods
		// in the earlier stages whenever possible.
		sort.Sort(controller.ActivePods(filteredPods))
	}
	// Snapshot the UIDs (ns/name) of the pods we're expecting to see
	// deleted, so we know to record their expectations exactly once either
	// when we see it as an update of the deletion timestamp, or as a delete.
	// Note that if the labels on a pod/rc change in a way that the pod gets
	// orphaned, the rs will only wake up after the expectations have
	// expired even if other pods are deleted.
	deletedPodKeys := []string{}
	for i := 0; i < diff; i++ {
		deletedPodKeys = append(deletedPodKeys, controller.PodKey(filteredPods[i]))
	}
	// We use pod namespace/name as a UID to wait for deletions, so if the
	// labels on a pod/rc change in a way that the pod gets orphaned, the
	// rc will only wake up after the expectation has expired.
	errCh := make(chan error, diff)
	rm.expectations.ExpectDeletions(rcKey, deletedPodKeys)
	var wg sync.WaitGroup
	wg.Add(diff)
	for i := 0; i < diff; i++ {
		go func(ix int) {
			defer wg.Done()
			if err := rm.podControl.DeletePod(rc.Namespace, filteredPods[ix].Name, rc); err != nil {
				// Decrement the expected number of deletes because the informer won't observe this deletion
				podKey := controller.PodKey(filteredPods[ix])
				glog.V(2).Infof("Failed to delete %v due to %v, decrementing expectations for controller %q/%q", podKey, err, rc.Namespace, rc.Name)
				rm.expectations.DeletionObserved(rcKey, podKey)
				errCh <- err
				utilruntime.HandleError(err)
			}
		}(i)
	}
	wg.Wait()

	select {
	case err := <-errCh:
		// all errors have been reported before and they're likely to be the same, so we'll only return the first one we hit.
		if err != nil {
			return err
		}
	default:
	}

	return nil

}

// syncReplicationController will sync the rc with the given key if it has had its expectations fulfilled, meaning
// it did not expect to see any more of its pods created or deleted. This function is not meant to be invoked
// concurrently with the same key.
func (rm *ReplicationManager) syncReplicationController(key string) error {
	trace := utiltrace.New("syncReplicationController: " + key)
	defer trace.LogIfLong(250 * time.Millisecond)

	startTime := time.Now()
	defer func() {
		glog.V(4).Infof("Finished syncing controller %q (%v)", key, time.Now().Sub(startTime))
	}()

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	rc, err := rm.rcLister.ReplicationControllers(namespace).Get(name)
	if errors.IsNotFound(err) {
		glog.Infof("Replication Controller has been deleted %v", key)
		rm.expectations.DeleteExpectations(key)
		return nil
	}
	if err != nil {
		return err
	}

	trace.Step("ReplicationController restored")
	rcNeedsSync := rm.expectations.SatisfiedExpectations(key)
	trace.Step("Expectations restored")

	// NOTE: filteredPods are pointing to objects from cache - if you need to
	// modify them, you need to copy it first.
	// TODO: Do the List and Filter in a single pass, or use an index.
	var filteredPods []*v1.Pod
	// list all pods to include the pods that don't match the rc's selector
	// anymore but has the stale controller ref.
	pods, err := rm.podLister.Pods(rc.Namespace).List(labels.Everything())
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Error getting pods for rc %q: %v", key, err))
		rm.queue.Add(key)
		return err
	}
	cm := controller.NewPodControllerRefManager(rm.podControl, rc, labels.Set(rc.Spec.Selector).AsSelectorPreValidated(), getRCKind())
	filteredPods, err = cm.ClaimPods(pods)
	if err != nil {
		// Something went wrong with adoption or release.
		// Requeue and try again so we don't leave orphans sitting around.
		rm.queue.Add(key)
		return err
	}

	var manageReplicasErr error
	if rcNeedsSync && rc.DeletionTimestamp == nil {
		manageReplicasErr = rm.manageReplicas(filteredPods, rc)
	}
	trace.Step("manageReplicas done")

	copy, err := api.Scheme.DeepCopy(rc)
	if err != nil {
		return err
	}
	rc = copy.(*v1.ReplicationController)

	newStatus := calculateStatus(rc, filteredPods, manageReplicasErr)

	// Always updates status as pods come up or die.
	updatedRC, err := updateReplicationControllerStatus(rm.kubeClient.Core().ReplicationControllers(rc.Namespace), *rc, newStatus)
	if err != nil {
		// Multiple things could lead to this update failing.  Returning an error causes a requeue without forcing a hotloop
		return err
	}
	// Resync the ReplicationController after MinReadySeconds as a last line of defense to guard against clock-skew.
	if manageReplicasErr == nil && updatedRC.Spec.MinReadySeconds > 0 &&
		updatedRC.Status.ReadyReplicas == *(updatedRC.Spec.Replicas) &&
		updatedRC.Status.AvailableReplicas != *(updatedRC.Spec.Replicas) {
		rm.enqueueControllerAfter(updatedRC, time.Duration(updatedRC.Spec.MinReadySeconds)*time.Second)
	}
	return manageReplicasErr
}
