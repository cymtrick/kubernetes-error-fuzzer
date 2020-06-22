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

package scheduler

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	podutil "k8s.io/kubernetes/pkg/api/v1/pod"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/apis/config"
	"k8s.io/kubernetes/pkg/scheduler/apis/config/scheme"
	"k8s.io/kubernetes/pkg/scheduler/core"
	frameworkplugins "k8s.io/kubernetes/pkg/scheduler/framework/plugins"
	frameworkruntime "k8s.io/kubernetes/pkg/scheduler/framework/runtime"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	internalcache "k8s.io/kubernetes/pkg/scheduler/internal/cache"
	internalqueue "k8s.io/kubernetes/pkg/scheduler/internal/queue"
	"k8s.io/kubernetes/pkg/scheduler/metrics"
	"k8s.io/kubernetes/pkg/scheduler/profile"
	"k8s.io/kubernetes/pkg/scheduler/util"
)

const (
	// SchedulerError is the reason recorded for events when an error occurs during scheduling a pod.
	SchedulerError = "SchedulerError"
	// Percentage of plugin metrics to be sampled.
	pluginMetricsSamplePercent = 10
)

// Scheduler watches for new unscheduled pods. It attempts to find
// nodes that they fit on and writes bindings back to the api server.
type Scheduler struct {
	// It is expected that changes made via SchedulerCache will be observed
	// by NodeLister and Algorithm.
	SchedulerCache internalcache.Cache

	Algorithm core.ScheduleAlgorithm

	// NextPod should be a function that blocks until the next pod
	// is available. We don't use a channel for this, because scheduling
	// a pod may take some amount of time and we don't want pods to get
	// stale while they sit in a channel.
	NextPod func() *framework.QueuedPodInfo

	// Error is called if there is an error. It is passed the pod in
	// question, and the error
	Error func(*framework.QueuedPodInfo, error)

	// Close this to shut down the scheduler.
	StopEverything <-chan struct{}

	// Disable pod preemption or not.
	DisablePreemption bool

	// SchedulingQueue holds pods to be scheduled
	SchedulingQueue internalqueue.SchedulingQueue

	// Profiles are the scheduling profiles.
	Profiles profile.Map

	scheduledPodsHasSynced func() bool

	client clientset.Interface
}

// Cache returns the cache in scheduler for test to check the data in scheduler.
func (sched *Scheduler) Cache() internalcache.Cache {
	return sched.SchedulerCache
}

type schedulerOptions struct {
	schedulerAlgorithmSource schedulerapi.SchedulerAlgorithmSource
	disablePreemption        bool
	percentageOfNodesToScore int32
	podInitialBackoffSeconds int64
	podMaxBackoffSeconds     int64
	// Contains out-of-tree plugins to be merged with the in-tree registry.
	frameworkOutOfTreeRegistry frameworkruntime.Registry
	profiles                   []schedulerapi.KubeSchedulerProfile
	extenders                  []schedulerapi.Extender
	frameworkCapturer          FrameworkCapturer
}

// Option configures a Scheduler
type Option func(*schedulerOptions)

// WithProfiles sets profiles for Scheduler. By default, there is one profile
// with the name "default-scheduler".
func WithProfiles(p ...schedulerapi.KubeSchedulerProfile) Option {
	return func(o *schedulerOptions) {
		o.profiles = p
	}
}

// WithAlgorithmSource sets schedulerAlgorithmSource for Scheduler, the default is a source with DefaultProvider.
func WithAlgorithmSource(source schedulerapi.SchedulerAlgorithmSource) Option {
	return func(o *schedulerOptions) {
		o.schedulerAlgorithmSource = source
	}
}

// WithPreemptionDisabled sets disablePreemption for Scheduler, the default value is false
func WithPreemptionDisabled(disablePreemption bool) Option {
	return func(o *schedulerOptions) {
		o.disablePreemption = disablePreemption
	}
}

// WithPercentageOfNodesToScore sets percentageOfNodesToScore for Scheduler, the default value is 50
func WithPercentageOfNodesToScore(percentageOfNodesToScore int32) Option {
	return func(o *schedulerOptions) {
		o.percentageOfNodesToScore = percentageOfNodesToScore
	}
}

// WithFrameworkOutOfTreeRegistry sets the registry for out-of-tree plugins. Those plugins
// will be appended to the default registry.
func WithFrameworkOutOfTreeRegistry(registry frameworkruntime.Registry) Option {
	return func(o *schedulerOptions) {
		o.frameworkOutOfTreeRegistry = registry
	}
}

// WithPodInitialBackoffSeconds sets podInitialBackoffSeconds for Scheduler, the default value is 1
func WithPodInitialBackoffSeconds(podInitialBackoffSeconds int64) Option {
	return func(o *schedulerOptions) {
		o.podInitialBackoffSeconds = podInitialBackoffSeconds
	}
}

// WithPodMaxBackoffSeconds sets podMaxBackoffSeconds for Scheduler, the default value is 10
func WithPodMaxBackoffSeconds(podMaxBackoffSeconds int64) Option {
	return func(o *schedulerOptions) {
		o.podMaxBackoffSeconds = podMaxBackoffSeconds
	}
}

// WithExtenders sets extenders for the Scheduler
func WithExtenders(e ...schedulerapi.Extender) Option {
	return func(o *schedulerOptions) {
		o.extenders = e
	}
}

// FrameworkCapturer is used for registering a notify function in building framework.
type FrameworkCapturer func(schedulerapi.KubeSchedulerProfile)

// WithBuildFrameworkCapturer sets a notify function for getting buildFramework details.
func WithBuildFrameworkCapturer(fc FrameworkCapturer) Option {
	return func(o *schedulerOptions) {
		o.frameworkCapturer = fc
	}
}

var defaultSchedulerOptions = schedulerOptions{
	profiles: []schedulerapi.KubeSchedulerProfile{
		// Profiles' default plugins are set from the algorithm provider.
		{SchedulerName: v1.DefaultSchedulerName},
	},
	schedulerAlgorithmSource: schedulerapi.SchedulerAlgorithmSource{
		Provider: defaultAlgorithmSourceProviderName(),
	},
	disablePreemption:        false,
	percentageOfNodesToScore: schedulerapi.DefaultPercentageOfNodesToScore,
	podInitialBackoffSeconds: int64(internalqueue.DefaultPodInitialBackoffDuration.Seconds()),
	podMaxBackoffSeconds:     int64(internalqueue.DefaultPodMaxBackoffDuration.Seconds()),
}

// New returns a Scheduler
func New(client clientset.Interface,
	informerFactory informers.SharedInformerFactory,
	podInformer coreinformers.PodInformer,
	recorderFactory profile.RecorderFactory,
	stopCh <-chan struct{},
	opts ...Option) (*Scheduler, error) {

	stopEverything := stopCh
	if stopEverything == nil {
		stopEverything = wait.NeverStop
	}

	options := defaultSchedulerOptions
	for _, opt := range opts {
		opt(&options)
	}

	schedulerCache := internalcache.New(30*time.Second, stopEverything)

	registry := frameworkplugins.NewInTreeRegistry()
	if err := registry.Merge(options.frameworkOutOfTreeRegistry); err != nil {
		return nil, err
	}

	snapshot := internalcache.NewEmptySnapshot()

	configurator := &Configurator{
		client:                   client,
		recorderFactory:          recorderFactory,
		informerFactory:          informerFactory,
		podInformer:              podInformer,
		schedulerCache:           schedulerCache,
		StopEverything:           stopEverything,
		disablePreemption:        options.disablePreemption,
		percentageOfNodesToScore: options.percentageOfNodesToScore,
		podInitialBackoffSeconds: options.podInitialBackoffSeconds,
		podMaxBackoffSeconds:     options.podMaxBackoffSeconds,
		profiles:                 append([]schedulerapi.KubeSchedulerProfile(nil), options.profiles...),
		registry:                 registry,
		nodeInfoSnapshot:         snapshot,
		extenders:                options.extenders,
		frameworkCapturer:        options.frameworkCapturer,
	}

	metrics.Register()

	var sched *Scheduler
	source := options.schedulerAlgorithmSource
	switch {
	case source.Provider != nil:
		// Create the config from a named algorithm provider.
		sc, err := configurator.createFromProvider(*source.Provider)
		if err != nil {
			return nil, fmt.Errorf("couldn't create scheduler using provider %q: %v", *source.Provider, err)
		}
		sched = sc
	case source.Policy != nil:
		// Create the config from a user specified policy source.
		policy := &schedulerapi.Policy{}
		switch {
		case source.Policy.File != nil:
			if err := initPolicyFromFile(source.Policy.File.Path, policy); err != nil {
				return nil, err
			}
		case source.Policy.ConfigMap != nil:
			if err := initPolicyFromConfigMap(client, source.Policy.ConfigMap, policy); err != nil {
				return nil, err
			}
		}
		// Set extenders on the configurator now that we've decoded the policy
		// In this case, c.extenders should be nil since we're using a policy (and therefore not componentconfig,
		// which would have set extenders in the above instantiation of Configurator from CC options)
		configurator.extenders = policy.Extenders
		sc, err := configurator.createFromConfig(*policy)
		if err != nil {
			return nil, fmt.Errorf("couldn't create scheduler from policy: %v", err)
		}
		sched = sc
	default:
		return nil, fmt.Errorf("unsupported algorithm source: %v", source)
	}
	// Additional tweaks to the config produced by the configurator.
	sched.DisablePreemption = options.disablePreemption
	sched.StopEverything = stopEverything
	sched.client = client
	sched.scheduledPodsHasSynced = podInformer.Informer().HasSynced

	addAllEventHandlers(sched, informerFactory, podInformer)
	return sched, nil
}

// initPolicyFromFile initialize policy from file
func initPolicyFromFile(policyFile string, policy *schedulerapi.Policy) error {
	// Use a policy serialized in a file.
	_, err := os.Stat(policyFile)
	if err != nil {
		return fmt.Errorf("missing policy config file %s", policyFile)
	}
	data, err := ioutil.ReadFile(policyFile)
	if err != nil {
		return fmt.Errorf("couldn't read policy config: %v", err)
	}
	err = runtime.DecodeInto(scheme.Codecs.UniversalDecoder(), []byte(data), policy)
	if err != nil {
		return fmt.Errorf("invalid policy: %v", err)
	}
	return nil
}

// initPolicyFromConfigMap initialize policy from configMap
func initPolicyFromConfigMap(client clientset.Interface, policyRef *schedulerapi.SchedulerPolicyConfigMapSource, policy *schedulerapi.Policy) error {
	// Use a policy serialized in a config map value.
	policyConfigMap, err := client.CoreV1().ConfigMaps(policyRef.Namespace).Get(context.TODO(), policyRef.Name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("couldn't get policy config map %s/%s: %v", policyRef.Namespace, policyRef.Name, err)
	}
	data, found := policyConfigMap.Data[schedulerapi.SchedulerPolicyConfigMapKey]
	if !found {
		return fmt.Errorf("missing policy config map value at key %q", schedulerapi.SchedulerPolicyConfigMapKey)
	}
	err = runtime.DecodeInto(scheme.Codecs.UniversalDecoder(), []byte(data), policy)
	if err != nil {
		return fmt.Errorf("invalid policy: %v", err)
	}
	return nil
}

// Run begins watching and scheduling. It waits for cache to be synced, then starts scheduling and blocked until the context is done.
func (sched *Scheduler) Run(ctx context.Context) {
	if !cache.WaitForCacheSync(ctx.Done(), sched.scheduledPodsHasSynced) {
		return
	}
	sched.SchedulingQueue.Run()
	wait.UntilWithContext(ctx, sched.scheduleOne, 0)
	sched.SchedulingQueue.Close()
}

// recordSchedulingFailure records an event for the pod that indicates the
// pod has failed to schedule. Also, update the pod condition and nominated node name if set.
func (sched *Scheduler) recordSchedulingFailure(prof *profile.Profile, podInfo *framework.QueuedPodInfo, err error, reason string, nominatedNode string) {
	sched.Error(podInfo, err)

	// Update the scheduling queue with the nominated pod information. Without
	// this, there would be a race condition between the next scheduling cycle
	// and the time the scheduler receives a Pod Update for the nominated pod.
	// Here we check for nil only for tests.
	if sched.SchedulingQueue != nil {
		sched.SchedulingQueue.AddNominatedPod(podInfo.Pod, nominatedNode)
	}

	pod := podInfo.Pod
	prof.Recorder.Eventf(pod, nil, v1.EventTypeWarning, "FailedScheduling", "Scheduling", err.Error())
	if err := updatePod(sched.client, pod, &v1.PodCondition{
		Type:    v1.PodScheduled,
		Status:  v1.ConditionFalse,
		Reason:  reason,
		Message: err.Error(),
	}, nominatedNode); err != nil {
		klog.Errorf("Error updating pod %s/%s: %v", pod.Namespace, pod.Name, err)
	}
}

func updatePod(client clientset.Interface, pod *v1.Pod, condition *v1.PodCondition, nominatedNode string) error {
	klog.V(3).Infof("Updating pod condition for %s/%s to (%s==%s, Reason=%s)", pod.Namespace, pod.Name, condition.Type, condition.Status, condition.Reason)
	podCopy := pod.DeepCopy()
	// NominatedNodeName is updated only if we are trying to set it, and the value is
	// different from the existing one.
	if !podutil.UpdatePodCondition(&podCopy.Status, condition) &&
		(len(nominatedNode) == 0 || pod.Status.NominatedNodeName == nominatedNode) {
		return nil
	}
	if nominatedNode != "" {
		podCopy.Status.NominatedNodeName = nominatedNode
	}
	return util.PatchPod(client, pod, podCopy)
}

// assume signals to the cache that a pod is already in the cache, so that binding can be asynchronous.
// assume modifies `assumed`.
func (sched *Scheduler) assume(assumed *v1.Pod, host string) error {
	// Optimistically assume that the binding will succeed and send it to apiserver
	// in the background.
	// If the binding fails, scheduler will release resources allocated to assumed pod
	// immediately.
	assumed.Spec.NodeName = host

	if err := sched.SchedulerCache.AssumePod(assumed); err != nil {
		klog.Errorf("scheduler cache AssumePod failed: %v", err)
		return err
	}
	// if "assumed" is a nominated pod, we should remove it from internal cache
	if sched.SchedulingQueue != nil {
		sched.SchedulingQueue.DeleteNominatedPodIfExists(assumed)
	}

	return nil
}

// bind binds a pod to a given node defined in a binding object.
// The precedence for binding is: (1) extenders and (2) framework plugins.
// We expect this to run asynchronously, so we handle binding metrics internally.
func (sched *Scheduler) bind(ctx context.Context, prof *profile.Profile, assumed *v1.Pod, targetNode string, state *framework.CycleState) (err error) {
	start := time.Now()
	defer func() {
		sched.finishBinding(prof, assumed, targetNode, start, err)
	}()

	bound, err := sched.extendersBinding(assumed, targetNode)
	if bound {
		return err
	}
	bindStatus := prof.RunBindPlugins(ctx, state, assumed, targetNode)
	if bindStatus.IsSuccess() {
		return nil
	}
	if bindStatus.Code() == framework.Error {
		return bindStatus.AsError()
	}
	return fmt.Errorf("bind status: %s, %v", bindStatus.Code().String(), bindStatus.Message())
}

// TODO(#87159): Move this to a Plugin.
func (sched *Scheduler) extendersBinding(pod *v1.Pod, node string) (bool, error) {
	for _, extender := range sched.Algorithm.Extenders() {
		if !extender.IsBinder() || !extender.IsInterested(pod) {
			continue
		}
		return true, extender.Bind(&v1.Binding{
			ObjectMeta: metav1.ObjectMeta{Namespace: pod.Namespace, Name: pod.Name, UID: pod.UID},
			Target:     v1.ObjectReference{Kind: "Node", Name: node},
		})
	}
	return false, nil
}

func (sched *Scheduler) finishBinding(prof *profile.Profile, assumed *v1.Pod, targetNode string, start time.Time, err error) {
	if finErr := sched.SchedulerCache.FinishBinding(assumed); finErr != nil {
		klog.Errorf("scheduler cache FinishBinding failed: %v", finErr)
	}
	if err != nil {
		klog.V(1).Infof("Failed to bind pod: %v/%v", assumed.Namespace, assumed.Name)
		if err := sched.SchedulerCache.ForgetPod(assumed); err != nil {
			klog.Errorf("scheduler cache ForgetPod failed: %v", err)
		}
		return
	}

	metrics.BindingLatency.Observe(metrics.SinceInSeconds(start))
	metrics.DeprecatedSchedulingDuration.WithLabelValues(metrics.Binding).Observe(metrics.SinceInSeconds(start))
	prof.Recorder.Eventf(assumed, nil, v1.EventTypeNormal, "Scheduled", "Binding", "Successfully assigned %v/%v to %v", assumed.Namespace, assumed.Name, targetNode)
}

// scheduleOne does the entire scheduling workflow for a single pod.  It is serialized on the scheduling algorithm's host fitting.
func (sched *Scheduler) scheduleOne(ctx context.Context) {
	podInfo := sched.NextPod()
	// pod could be nil when schedulerQueue is closed
	if podInfo == nil || podInfo.Pod == nil {
		return
	}
	pod := podInfo.Pod
	prof, err := sched.profileForPod(pod)
	if err != nil {
		// This shouldn't happen, because we only accept for scheduling the pods
		// which specify a scheduler name that matches one of the profiles.
		klog.Error(err)
		return
	}
	if sched.skipPodSchedule(prof, pod) {
		return
	}

	klog.V(3).Infof("Attempting to schedule pod: %v/%v", pod.Namespace, pod.Name)

	// Synchronously attempt to find a fit for the pod.
	start := time.Now()
	state := framework.NewCycleState()
	state.SetRecordPluginMetrics(rand.Intn(100) < pluginMetricsSamplePercent)
	schedulingCycleCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	scheduleResult, err := sched.Algorithm.Schedule(schedulingCycleCtx, prof, state, pod)
	if err != nil {
		// Schedule() may have failed because the pod would not fit on any host, so we try to
		// preempt, with the expectation that the next time the pod is tried for scheduling it
		// will fit due to the preemption. It is also possible that a different pod will schedule
		// into the resources that were preempted, but this is harmless.
		nominatedNode := ""
		if fitError, ok := err.(*core.FitError); ok {
			if sched.DisablePreemption {
				klog.V(3).Infof("Pod priority feature is not enabled or preemption is disabled by scheduler configuration." +
					" No preemption is performed.")
			} else {
				preemptionStartTime := time.Now()
				// TODO(Huang-Wei): implement the preemption logic as a PostFilter plugin.
				nominatedNode, _ = core.Preempt(schedulingCycleCtx, prof, state, pod, fitError.FilteredNodesStatuses)
				metrics.PreemptionAttempts.Inc()
				metrics.SchedulingAlgorithmPreemptionEvaluationDuration.Observe(metrics.SinceInSeconds(preemptionStartTime))
				metrics.DeprecatedSchedulingDuration.WithLabelValues(metrics.PreemptionEvaluation).Observe(metrics.SinceInSeconds(preemptionStartTime))

				// Run PostFilter plugins to try to make the pod schedulable in a future scheduling cycle.
				result, status := prof.RunPostFilterPlugins(ctx, state, pod, fitError.FilteredNodesStatuses)
				if status.Code() == framework.Error {
					klog.Errorf("Status after running PostFilter plugins for pod %v/%v: %v", pod.Namespace, pod.Name, status)
				} else {
					klog.V(5).Infof("Status after running PostFilter plugins for pod %v/%v: %v", pod.Namespace, pod.Name, status)
				}
				if status.IsSuccess() && result != nil {
					nominatedNode = result.NominatedNodeName
				}
			}
			// Pod did not fit anywhere, so it is counted as a failure. If preemption
			// succeeds, the pod should get counted as a success the next time we try to
			// schedule it. (hopefully)
			metrics.PodScheduleFailures.Inc()
		} else if err == core.ErrNoNodesAvailable {
			// No nodes available is counted as unschedulable rather than an error.
			metrics.PodScheduleFailures.Inc()
		} else {
			klog.ErrorS(err, "Error selecting node for pod", "pod", klog.KObj(pod))
			metrics.PodScheduleErrors.Inc()
		}
		sched.recordSchedulingFailure(prof, podInfo, err, v1.PodReasonUnschedulable, nominatedNode)
		return
	}
	metrics.SchedulingAlgorithmLatency.Observe(metrics.SinceInSeconds(start))
	// Tell the cache to assume that a pod now is running on a given node, even though it hasn't been bound yet.
	// This allows us to keep scheduling without waiting on binding to occur.
	assumedPodInfo := podInfo.DeepCopy()
	assumedPod := assumedPodInfo.Pod

	// Run "reserve" plugins.
	if sts := prof.RunReservePlugins(schedulingCycleCtx, state, assumedPod, scheduleResult.SuggestedHost); !sts.IsSuccess() {
		sched.recordSchedulingFailure(prof, assumedPodInfo, sts.AsError(), SchedulerError, "")
		metrics.PodScheduleErrors.Inc()
		return
	}

	// assume modifies `assumedPod` by setting NodeName=scheduleResult.SuggestedHost
	err = sched.assume(assumedPod, scheduleResult.SuggestedHost)
	if err != nil {
		// This is most probably result of a BUG in retrying logic.
		// We report an error here so that pod scheduling can be retried.
		// This relies on the fact that Error will check if the pod has been bound
		// to a node and if so will not add it back to the unscheduled pods queue
		// (otherwise this would cause an infinite loop).
		sched.recordSchedulingFailure(prof, assumedPodInfo, err, SchedulerError, "")
		metrics.PodScheduleErrors.Inc()
		// trigger un-reserve plugins to clean up state associated with the reserved Pod
		prof.RunUnreservePlugins(schedulingCycleCtx, state, assumedPod, scheduleResult.SuggestedHost)
		return
	}

	// Run "permit" plugins.
	runPermitStatus := prof.RunPermitPlugins(schedulingCycleCtx, state, assumedPod, scheduleResult.SuggestedHost)
	if runPermitStatus.Code() != framework.Wait && !runPermitStatus.IsSuccess() {
		var reason string
		if runPermitStatus.IsUnschedulable() {
			metrics.PodScheduleFailures.Inc()
			reason = v1.PodReasonUnschedulable
		} else {
			metrics.PodScheduleErrors.Inc()
			reason = SchedulerError
		}
		if forgetErr := sched.Cache().ForgetPod(assumedPod); forgetErr != nil {
			klog.Errorf("scheduler cache ForgetPod failed: %v", forgetErr)
		}
		// One of the plugins returned status different than success or wait.
		prof.RunUnreservePlugins(schedulingCycleCtx, state, assumedPod, scheduleResult.SuggestedHost)
		sched.recordSchedulingFailure(prof, assumedPodInfo, runPermitStatus.AsError(), reason, "")
		return
	}

	// bind the pod to its host asynchronously (we can do this b/c of the assumption step above).
	go func() {
		bindingCycleCtx, cancel := context.WithCancel(ctx)
		defer cancel()
		metrics.SchedulerGoroutines.WithLabelValues("binding").Inc()
		defer metrics.SchedulerGoroutines.WithLabelValues("binding").Dec()

		waitOnPermitStatus := prof.WaitOnPermit(bindingCycleCtx, assumedPod)
		if !waitOnPermitStatus.IsSuccess() {
			var reason string
			if waitOnPermitStatus.IsUnschedulable() {
				metrics.PodScheduleFailures.Inc()
				reason = v1.PodReasonUnschedulable
			} else {
				metrics.PodScheduleErrors.Inc()
				reason = SchedulerError
			}
			if forgetErr := sched.Cache().ForgetPod(assumedPod); forgetErr != nil {
				klog.Errorf("scheduler cache ForgetPod failed: %v", forgetErr)
			}
			// trigger un-reserve plugins to clean up state associated with the reserved Pod
			prof.RunUnreservePlugins(bindingCycleCtx, state, assumedPod, scheduleResult.SuggestedHost)
			sched.recordSchedulingFailure(prof, assumedPodInfo, waitOnPermitStatus.AsError(), reason, "")
			return
		}

		// Run "prebind" plugins.
		preBindStatus := prof.RunPreBindPlugins(bindingCycleCtx, state, assumedPod, scheduleResult.SuggestedHost)
		if !preBindStatus.IsSuccess() {
			var reason string
			metrics.PodScheduleErrors.Inc()
			reason = SchedulerError
			if forgetErr := sched.Cache().ForgetPod(assumedPod); forgetErr != nil {
				klog.Errorf("scheduler cache ForgetPod failed: %v", forgetErr)
			}
			// trigger un-reserve plugins to clean up state associated with the reserved Pod
			prof.RunUnreservePlugins(bindingCycleCtx, state, assumedPod, scheduleResult.SuggestedHost)
			sched.recordSchedulingFailure(prof, assumedPodInfo, preBindStatus.AsError(), reason, "")
			return
		}

		err := sched.bind(bindingCycleCtx, prof, assumedPod, scheduleResult.SuggestedHost, state)
		metrics.E2eSchedulingLatency.Observe(metrics.SinceInSeconds(start))
		if err != nil {
			metrics.PodScheduleErrors.Inc()
			// trigger un-reserve plugins to clean up state associated with the reserved Pod
			prof.RunUnreservePlugins(bindingCycleCtx, state, assumedPod, scheduleResult.SuggestedHost)
			sched.recordSchedulingFailure(prof, assumedPodInfo, fmt.Errorf("Binding rejected: %v", err), SchedulerError, "")
		} else {
			// Calculating nodeResourceString can be heavy. Avoid it if klog verbosity is below 2.
			if klog.V(2).Enabled() {
				klog.InfoS("Successfully bound pod to node", "pod", klog.KObj(pod), "node", scheduleResult.SuggestedHost, "evaluatedNodes", scheduleResult.EvaluatedNodes, "feasibleNodes", scheduleResult.FeasibleNodes)
			}

			metrics.PodScheduleSuccesses.Inc()
			metrics.PodSchedulingAttempts.Observe(float64(podInfo.Attempts))
			metrics.PodSchedulingDuration.Observe(metrics.SinceInSeconds(podInfo.InitialAttemptTimestamp))

			// Run "postbind" plugins.
			prof.RunPostBindPlugins(bindingCycleCtx, state, assumedPod, scheduleResult.SuggestedHost)
		}
	}()
}

func (sched *Scheduler) profileForPod(pod *v1.Pod) (*profile.Profile, error) {
	prof, ok := sched.Profiles[pod.Spec.SchedulerName]
	if !ok {
		return nil, fmt.Errorf("profile not found for scheduler name %q", pod.Spec.SchedulerName)
	}
	return prof, nil
}

// skipPodSchedule returns true if we could skip scheduling the pod for specified cases.
func (sched *Scheduler) skipPodSchedule(prof *profile.Profile, pod *v1.Pod) bool {
	// Case 1: pod is being deleted.
	if pod.DeletionTimestamp != nil {
		prof.Recorder.Eventf(pod, nil, v1.EventTypeWarning, "FailedScheduling", "Scheduling", "skip schedule deleting pod: %v/%v", pod.Namespace, pod.Name)
		klog.V(3).Infof("Skip schedule deleting pod: %v/%v", pod.Namespace, pod.Name)
		return true
	}

	// Case 2: pod has been assumed and pod updates could be skipped.
	// An assumed pod can be added again to the scheduling queue if it got an update event
	// during its previous scheduling cycle but before getting assumed.
	if sched.skipPodUpdate(pod) {
		return true
	}

	return false
}

func defaultAlgorithmSourceProviderName() *string {
	provider := schedulerapi.SchedulerDefaultProviderName
	return &provider
}
