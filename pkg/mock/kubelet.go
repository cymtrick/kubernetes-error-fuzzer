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

package mock

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"reflect"
	goruntime "runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/golang/mock/gomock"
	cadvisorapi "github.com/google/cadvisor/info/v1"
	cadvisorapiv2 "github.com/google/cadvisor/info/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/mount-utils"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/flowcontrol"
	featuregatetesting "k8s.io/component-base/featuregate/testing"
	internalapi "k8s.io/cri-api/pkg/apis"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"
	"k8s.io/klog/v2/ktesting"
	"k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/pkg/kubelet"
	kubeletconfiginternal "k8s.io/kubernetes/pkg/kubelet/apis/config"
	cadvisortest "k8s.io/kubernetes/pkg/kubelet/cadvisor/testing"
	"k8s.io/kubernetes/pkg/kubelet/clustertrustbundle"
	"k8s.io/kubernetes/pkg/kubelet/cm"
	"k8s.io/kubernetes/pkg/kubelet/config"
	"k8s.io/kubernetes/pkg/kubelet/configmap"
	kubecontainer "k8s.io/kubernetes/pkg/kubelet/container"
	containertest "k8s.io/kubernetes/pkg/kubelet/container/testing"
	"k8s.io/kubernetes/pkg/kubelet/cri/remote"
	fakeremote "k8s.io/kubernetes/pkg/kubelet/cri/remote/fake"
	"k8s.io/kubernetes/pkg/kubelet/eviction"
	"k8s.io/kubernetes/pkg/kubelet/images"
	"k8s.io/kubernetes/pkg/kubelet/kuberuntime"
	"k8s.io/kubernetes/pkg/kubelet/lifecycle"
	"k8s.io/kubernetes/pkg/kubelet/logs"
	"k8s.io/kubernetes/pkg/kubelet/network/dns"
	"k8s.io/kubernetes/pkg/kubelet/nodeshutdown"
	"k8s.io/kubernetes/pkg/kubelet/pleg"
	"k8s.io/kubernetes/pkg/kubelet/pluginmanager"
	kubepod "k8s.io/kubernetes/pkg/kubelet/pod"
	podtest "k8s.io/kubernetes/pkg/kubelet/pod/testing"
	proberesults "k8s.io/kubernetes/pkg/kubelet/prober/results"
	probetest "k8s.io/kubernetes/pkg/kubelet/prober/testing"
	"k8s.io/kubernetes/pkg/kubelet/secret"
	"k8s.io/kubernetes/pkg/kubelet/server"
	serverstats "k8s.io/kubernetes/pkg/kubelet/server/stats"
	"k8s.io/kubernetes/pkg/kubelet/stats"
	"k8s.io/kubernetes/pkg/kubelet/status"
	"k8s.io/kubernetes/pkg/kubelet/status/state"
	statustest "k8s.io/kubernetes/pkg/kubelet/status/testing"
	"k8s.io/kubernetes/pkg/kubelet/token"
	kubetypes "k8s.io/kubernetes/pkg/kubelet/types"
	kubeletutil "k8s.io/kubernetes/pkg/kubelet/util"
	"k8s.io/kubernetes/pkg/kubelet/util/queue"
	kubeletvolume "k8s.io/kubernetes/pkg/kubelet/volumemanager"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/util/oom"
	"k8s.io/kubernetes/pkg/volume"
	_ "k8s.io/kubernetes/pkg/volume/hostpath"
	volumesecret "k8s.io/kubernetes/pkg/volume/secret"
	volumetest "k8s.io/kubernetes/pkg/volume/testing"
	"k8s.io/kubernetes/pkg/volume/util"
	"k8s.io/kubernetes/pkg/volume/util/hostutil"
	"k8s.io/kubernetes/pkg/volume/util/subpath"
	"k8s.io/utils/clock"
	testingclock "k8s.io/utils/clock/testing"
	utilpointer "k8s.io/utils/pointer"
)

func init() {
}

const (
	testKubeletHostname = "127.0.0.1"
	testKubeletHostIP   = "127.0.0.1"
	testKubeletHostIPv6 = "::1"

	// TODO(harry) any global place for these two?
	// Reasonable size range of all container images. 90%ile of images on dockerhub drops into this range.
	minImgSize int64 = 23 * 1024 * 1024
	maxImgSize int64 = 1000 * 1024 * 1024
)

// fakeImageGCManager is a fake image gc manager for testing. It will return image
// list from fake runtime directly instead of caching it.
type fakeImageGCManager struct {
	fakeImageService kubecontainer.ImageService
	images.ImageGCManager
}

func (f *fakeImageGCManager) GetImageList() ([]kubecontainer.Image, error) {
	return f.fakeImageService.ListImages(context.Background())
}

type serviceLister interface {
	List(selector labels.Selector) ([]*v1.Service, error)
}

type testServiceLister struct {
	services []*v1.Service
}

func (ls testServiceLister) List(labels.Selector) ([]*v1.Service, error) {
	return ls.services, nil
}

type TestKubelet struct {
	kubelet              *kubelet.Kubelet
	fakeRuntime          *containertest.FakeRuntime
	fakeContainerManager *cm.FakeContainerManager
	fakeKubeClient       *fake.Clientset
	fakeMirrorClient     *podtest.FakeMirrorClient
	fakeClock            *testingclock.FakeClock
	mounter              mount.Interface
	volumePlugin         *volumetest.FakeVolumePlugin
}

type fakePodWorkers struct {
	lock       sync.Mutex
	syncPodFn  kubelet.SyncPodFnType
	cache      kubecontainer.Cache
	t          TestingInterface
	randomData []byte

	triggeredDeletion []types.UID
	triggeredTerminal []types.UID

	statusLock            sync.Mutex
	running               map[types.UID]bool
	terminating           map[types.UID]bool
	terminated            map[types.UID]bool
	terminationRequested  map[types.UID]bool
	finished              map[types.UID]bool
	removeRuntime         map[types.UID]bool
	removeContent         map[types.UID]bool
	terminatingStaticPods map[string]bool
}

func (tk *TestKubelet) Cleanup() {
	if tk.kubelet != nil {
		os.RemoveAll(*tk.kubelet.GetRootDir())
		tk.kubelet = nil
	}
}

// newTestKubelet returns test kubelet with two images.
func newTestKubelet(t *testing.T, controllerAttachDetachEnabled bool) *TestKubelet {
	imageList := []kubecontainer.Image{
		{
			ID:       "abc",
			RepoTags: []string{"registry.k8s.io:v1", "registry.k8s.io:v2"},
			Size:     123,
		},
		{
			ID:       "efg",
			RepoTags: []string{"registry.k8s.io:v3", "registry.k8s.io:v4"},
			Size:     456,
		},
	}
	return newTestKubeletWithImageList(t, imageList, controllerAttachDetachEnabled, true /*initFakeVolumePlugin*/, true /*localStorageCapacityIsolation*/)
}

func newTestKubeletWithImageList(
	t *testing.T,
	imageList []kubecontainer.Image,
	controllerAttachDetachEnabled bool,
	initFakeVolumePlugin bool,
	localStorageCapacityIsolation bool,
) *TestKubelet {
	logger, _ := ktesting.NewTestContext(t)

	fakeRuntime := &containertest.FakeRuntime{
		ImageList: imageList,
		// Set ready conditions by default.
		RuntimeStatus: &kubecontainer.RuntimeStatus{
			Conditions: []kubecontainer.RuntimeCondition{
				{Type: "RuntimeReady", Status: true},
				{Type: "NetworkReady", Status: true},
			},
		},
		VersionInfo: "1.5.0",
		RuntimeType: "test",
		T:           t,
	}

	fakeRecorder := &record.FakeRecorder{}
	fakeKubeClient := &fake.Clientset{}
	klets := &kubelet.Kubelet{}
	*klets.GetRecorder() = fakeRecorder
	*klets.GetKubeClient() = fakeKubeClient
	*klets.GetHeartbeatClient() = fakeKubeClient
	*klets.GetOS() = &containertest.FakeOS{}
	*klets.GetMounter() = mount.NewFakeMounter(nil)
	*klets.GetHostUtil() = hostutil.NewFakeHostUtil(nil)
	*klets.GetSubPather() = &subpath.FakeSubpath{}

	*klets.GetHostName() = testKubeletHostname
	*klets.GetNodeName() = types.NodeName(testKubeletHostname)
	fmt.Sprintf("%s", kubelet.MaxWaitForContainerRuntime)

	*klets.GetRuntimeState() = *kubelet.NewRuntimeState(kubelet.MaxWaitForContainerRuntime)
	klets.GetRuntimeState().SetNetworkState(nil)
	if tempDir, err := os.MkdirTemp("", "kubelet_test."); err != nil {
		t.Fatalf("can't make a temp rootdir: %v", err)
	} else {
		*klets.GetRootDir() = tempDir
	}
	if err := os.MkdirAll(*klets.GetRootDir(), 0750); err != nil {
		t.Fatalf("can't mkdir(%q): %v", *klets.GetRootDir(), err)
	}
	*klets.GetSourcesReady() = config.NewSourcesReady(func(_ sets.String) bool { return true })
	*klets.GetServiceLister() = testServiceLister{}
	*klets.GetServiceHasSynced() = func() bool { return true }
	*klets.GetNodeHasSynced() = func() bool { return true }
	*klets.GetNodeLister() = testNodeLister{
		nodes: []*v1.Node{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: string(*klets.GetNodeName()),
				},
				Status: v1.NodeStatus{
					Conditions: []v1.NodeCondition{
						{
							Type:    v1.NodeReady,
							Status:  v1.ConditionTrue,
							Reason:  "Ready",
							Message: "Node ready",
						},
					},
					Addresses: []v1.NodeAddress{
						{
							Type:    v1.NodeInternalIP,
							Address: testKubeletHostIP,
						},
						{
							Type:    v1.NodeInternalIP,
							Address: testKubeletHostIPv6,
						},
					},
					VolumesAttached: []v1.AttachedVolume{
						{
							Name:       "fake/fake-device",
							DevicePath: "fake/path",
						},
					},
				},
			},
		},
	}
	*klets.GetRecorder() = fakeRecorder
	if err := klets.SetupDataDirs(); err != nil {
		t.Fatalf("can't initialize kubelet data dirs: %v", err)
	}
	*klets.GetDaemonEndpoints() = v1.NodeDaemonEndpoints{}

	*klets.GetCAdvisor() = &cadvisortest.Fake{}
	machineInfo, _ := (*klets.GetCAdvisor()).MachineInfo()
	klets.SetCachedMachineInfo(machineInfo)
	*klets.GetTracer() = oteltrace.NewNoopTracerProvider().Tracer("")

	fakeMirrorClient := podtest.NewFakeMirrorClient()
	secretManager := secret.NewSimpleSecretManager(*klets.GetKubeClient())
	*klets.GetSecretManager() = secretManager
	configMapManager := configmap.NewSimpleConfigMapManager(*klets.GetKubeClient())
	*klets.GetConfigMapManager() = configMapManager
	*klets.GetMirrorPodManager() = fakeMirrorClient
	*klets.GetPodManager() = kubepod.NewBasicPodManager()
	podStartupLatencyTracker := kubeletutil.NewPodStartupLatencyTracker()
	*klets.GetStatusManager() = status.NewManager(fakeKubeClient, *klets.GetPodManager(), &statustest.FakePodDeletionSafetyProvider{}, podStartupLatencyTracker, *klets.GetRootDir())
	*klets.GetNodeStartupLatencyTracker() = kubeletutil.NewNodeStartupLatencyTracker()

	*klets.GetContainerRuntime() = fakeRuntime
	*klets.GetRuntimeCache() = containertest.NewFakeRuntimeCache(*klets.GetContainerRuntime())
	*klets.GetReasonCache() = *kubelet.NewReasonCache()
	*klets.GetPodCache() = containertest.NewFakeCache(*klets.GetContainerRuntime())
	*klets.GetPodWorkers() = &fakePodWorkers{
		syncPodFn: klets.SyncPod,
		cache:     *klets.GetPodCache(),
		t:         t,
	}

	*klets.GetProbeManager() = probetest.FakeManager{}
	*klets.GetLiveinessManager() = proberesults.NewManager()
	*klets.GetReadinessManager() = proberesults.NewManager()
	*klets.GetStartupManager() = proberesults.NewManager()

	fakeContainerManager := cm.NewFakeContainerManager()
	*klets.GetContainerManager() = fakeContainerManager
	fakeNodeRef := &v1.ObjectReference{
		Kind:      "Node",
		Name:      testKubeletHostname,
		UID:       types.UID(testKubeletHostname),
		Namespace: "",
	}

	volumeStatsAggPeriod := time.Second * 10
	*klets.GetResourceAnalyzer() = serverstats.NewResourceAnalyzer(klets, volumeStatsAggPeriod, *klets.GetRecorder())

	fakeHostStatsProvider := stats.NewFakeHostStatsProvider()

	*klets.GetStatsProvider() = *stats.NewCadvisorStatsProvider(
		*klets.GetCAdvisor(),
		*klets.GetResourceAnalyzer(),
		*klets.GetPodManager(),
		*klets.GetRuntimeCache(),
		fakeRuntime,
		*klets.GetStatusManager(),
		fakeHostStatsProvider,
	)
	fakeImageGCPolicy := images.ImageGCPolicy{
		HighThresholdPercent: 90,
		LowThresholdPercent:  80,
	}
	imageGCManager, err := images.NewImageGCManager(fakeRuntime, *klets.GetStatsProvider(), fakeRecorder, fakeNodeRef, fakeImageGCPolicy, oteltrace.NewNoopTracerProvider())
	assert.NoError(t, err)
	*klets.GetImageManager() = &fakeImageGCManager{
		fakeImageService: fakeRuntime,
		ImageGCManager:   imageGCManager,
	}
	*klets.GetContainerLogManager() = logs.NewStubContainerLogManager()
	containerGCPolicy := kubecontainer.GCPolicy{
		MinAge:             time.Duration(0),
		MaxPerPodContainer: 1,
		MaxContainers:      -1,
	}
	containerGC, err := kubecontainer.NewContainerGC(fakeRuntime, containerGCPolicy, *klets.GetSourcesReady())
	assert.NoError(t, err)
	*klets.GetContainerGC() = containerGC

	fakeClock := testingclock.NewFakeClock(time.Now())
	*klets.GetBackOff() = *flowcontrol.NewBackOff(time.Second, time.Minute)
	klets.GetBackOff().Clock = fakeClock
	*klets.GetResyncInterval() = 10 * time.Second
	*klets.GetWorkQueue() = queue.NewBasicWorkQueue(fakeClock)
	// Relist period does not affect the tests.
	*klets.GetPleg() = pleg.NewGenericPLEG(fakeRuntime, make(chan *pleg.PodLifecycleEvent, 100), &pleg.RelistDuration{RelistPeriod: time.Hour, RelistThreshold: time.Minute * 3}, *klets.GetPodCache(), clock.RealClock{})
	*klets.GetClock() = fakeClock

	nodeRef := &v1.ObjectReference{
		Kind:      "Node",
		Name:      string(*klets.GetNodeName()),
		UID:       types.UID(*klets.GetNodeName()),
		Namespace: "",
	}
	// setup eviction manager
	evictionManager, evictionAdmitHandler := eviction.NewManager(*klets.GetResourceAnalyzer(), eviction.Config{},
		kubelet.KillPodNow(*klets.GetPodWorkers(), fakeRecorder), *klets.GetImageManager(), *klets.GetContainerGC(), fakeRecorder, nodeRef, *klets.GetClock(), klets.SupportLocalStorageCapacityIsolation())

	*klets.GetEvictionManager() = evictionManager
	klets.AdmitHandlers().AddPodAdmitHandler(evictionAdmitHandler)

	// setup shutdown manager
	shutdownManager, shutdownAdmitHandler := nodeshutdown.NewManager(&nodeshutdown.Config{
		Logger:                          logger,
		ProbeManager:                    *klets.GetProbeManager(),
		Recorder:                        fakeRecorder,
		NodeRef:                         nodeRef,
		GetPodsFunc:                     (*klets.GetPodManager()).GetPods,
		KillPodFunc:                     kubelet.KillPodNow(*klets.GetPodWorkers(), fakeRecorder),
		SyncNodeStatusFunc:              func() {},
		ShutdownGracePeriodRequested:    0,
		ShutdownGracePeriodCriticalPods: 0,
	})
	*klets.GetShutdownManager() = shutdownManager
	klets.AdmitHandlers().AddPodAdmitHandler(shutdownAdmitHandler)

	// Add this as cleanup predicate pod admitter
	klets.AdmitHandlers().AddPodAdmitHandler(lifecycle.NewPredicateAdmitHandler(klets.GetNodeAnyWay, lifecycle.NewAdmissionFailureHandlerStub(), (*klets.GetContainerManager()).UpdatePluginResources))

	allPlugins := []volume.VolumePlugin{}
	plug := &volumetest.FakeVolumePlugin{PluginName: "fake", Host: nil}
	if initFakeVolumePlugin {
		allPlugins = append(allPlugins, plug)
	} else {
		allPlugins = append(allPlugins, volumesecret.ProbeVolumePlugins()...)
	}

	var prober volume.DynamicPluginProber // TODO (#51147) inject mock
	volumePluginMgr, err :=
		kubelet.NewInitializedVolumePluginMgr(klets, *klets.GetSecretManager(), *klets.GetConfigMapManager(), token.NewManager(*klets.GetKubeClient()), &clustertrustbundle.NoopManager{}, allPlugins, prober)
	require.NoError(t, err, "Failed to initialize VolumePluginMgr")
	*klets.GetVolumePluginMgr() = *volumePluginMgr

	*klets.GetVolumeManager() = kubeletvolume.NewVolumeManager(
		controllerAttachDetachEnabled,
		*klets.GetNodeName(),
		*klets.GetPodManager(),
		*klets.GetPodWorkers(),
		fakeKubeClient,
		klets.GetVolumePluginMgr(),
		fakeRuntime,
		*klets.GetMounter(),
		*klets.GetHostUtil(),
		klets.GetPodsDirGetters(),
		*klets.GetRecorder(),
		false, /* keepTerminatedPodVolumes */
		volumetest.NewBlockVolumePathHandler())

	*klets.GetPluginManager() = pluginmanager.NewPluginManager(
		klets.GetPluginsRegistrationDir(), /* sockDir */
		*klets.GetRecorder(),
	)
	*klets.SetNodeStatusFuncs() = klets.DefaultNodeStatusFuncs()

	// enable active deadline handler
	activeDeadlineHandler, err := kubelet.NewActiveDeadlineHandler(*klets.GetStatusManager(), *klets.GetRecorder(), *klets.GetClock())
	require.NoError(t, err, "Can't initialize active deadline handler")

	klets.AddPodSyncLoopHandler(activeDeadlineHandler)
	klets.AddPodSyncHandler(activeDeadlineHandler)
	(*klets.GetKubeletConfiguration()).LocalStorageCapacityIsolation = localStorageCapacityIsolation
	return &TestKubelet{klets, fakeRuntime, fakeContainerManager, fakeKubeClient, fakeMirrorClient, fakeClock, nil, plug}
}

func newTestPods(count int) []*v1.Pod {
	pods := make([]*v1.Pod, count)
	for i := 0; i < count; i++ {
		pods[i] = &v1.Pod{
			Spec: v1.PodSpec{
				HostNetwork: true,
			},
			ObjectMeta: metav1.ObjectMeta{
				UID:  types.UID(strconv.Itoa(10000 + i)),
				Name: fmt.Sprintf("pod%d", i),
			},
		}
	}
	return pods
}

//export TestSyncLoopAbort
func TestSyncLoopAbort(t *testing.T) {
	ctx := context.Background()
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kubelet := testKubelet.kubelet
	(kubelet.GetRuntimeState()).SetRuntimeSync(time.Now())
	// The syncLoop waits on time.After(resyncInterval), set it really big so that we don't race for
	// the channel close
	*kubelet.GetResyncInterval() = time.Second * 30

	ch := make(chan kubetypes.PodUpdate)
	close(ch)

	// sanity check (also prevent this test from hanging in the next step)
	ok := kubelet.SyncLoopIteration(ctx, ch, kubelet, make(chan time.Time), make(chan time.Time), make(chan *pleg.PodLifecycleEvent, 1))
	require.False(t, ok, "Expected syncLoopIteration to return !ok since update chan was closed")

	// this should terminate immediately; if it hangs then the syncLoopIteration isn't aborting properly
	kubelet.SyncLoop(ctx, ch, kubelet)
}

func TestSyncPodsStartPod(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kubelet := testKubelet.kubelet
	fakeRuntime := testKubelet.fakeRuntime
	pods := []*v1.Pod{
		podWithUIDNameNsSpec("12345678", "foo", "new", v1.PodSpec{
			Containers: []v1.Container{
				{Name: "bar"},
			},
		}),
	}
	(*kubelet.GetPodManager()).SetPods(pods)
	kubelet.HandlePodSyncs(pods)
	fakeRuntime.AssertStartedPods([]string{string(pods[0].UID)})
}

func TestHandlePodCleanupsPerQOS(t *testing.T, randomStrings []string) {
	ctx := context.Background()
	fmt.Println("randomStrings: ", randomStrings)
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()

	pod := &kubecontainer.Pod{
		ID:        "12345678",
		Name:      randomStrings[0],
		Namespace: randomStrings[1],
		Containers: []*kubecontainer.Container{
			{Name: randomStrings[2]},
		},
	}

	fakeRuntime := testKubelet.fakeRuntime
	fakeContainerManager := testKubelet.fakeContainerManager
	fakeContainerManager.PodContainerManager.AddPodFromCgroups(pod) // add pod to mock cgroup
	fakeRuntime.PodList = []*containertest.FakePod{
		{Pod: pod},
	}
	kubelet := testKubelet.kubelet
	*kubelet.GetCgroupsPerQOS() = true // enable cgroupsPerQOS to turn on the cgroups cleanup

	// HandlePodCleanups gets called every 2 seconds within the Kubelet's
	// housekeeping routine. This test registers the pod, removes the unwanted pod, then calls into
	// HandlePodCleanups a few more times. We should only see one Destroy() event. podKiller runs
	// within a goroutine so a two second delay should be enough time to
	// mark the pod as killed (within this test case).

	kubelet.HandlePodCleanups(ctx)

	// assert that unwanted pods were killed
	if actual, expected := (*kubelet.GetPodWorkers()).(*fakePodWorkers).triggeredDeletion, []types.UID{"12345678"}; !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %v to be deleted, got %v", expected, actual)
	}
	fakeRuntime.AssertKilledPods([]string(nil))

	// simulate Runtime.KillPod
	fakeRuntime.PodList = nil

	kubelet.HandlePodCleanups(ctx)
	kubelet.HandlePodCleanups(ctx)
	kubelet.HandlePodCleanups(ctx)

	destroyCount := 0
	err := wait.Poll(100*time.Millisecond, 10*time.Second, func() (bool, error) {
		fakeContainerManager.PodContainerManager.Lock()
		defer fakeContainerManager.PodContainerManager.Unlock()
		destroyCount = 0
		for _, functionName := range fakeContainerManager.PodContainerManager.CalledFunctions {
			if functionName == "Destroy" {
				destroyCount = destroyCount + 1
			}
		}
		return destroyCount >= 1, nil
	})

	assert.NoError(t, err, "wait should not return error")
	// housekeeping can get called multiple times. The cgroup Destroy() is
	// done within a goroutine and can get called multiple times, so the
	// Destroy() count in not deterministic on the actual number.
	// https://github.com/kubernetes/kubernetes/blob/29fdbb065b5e0d195299eb2d260b975cbc554673/pkg/kubelet/kubelet_pods.go#L2006
	assert.True(t, destroyCount >= 1, "Expect 1 or more destroys")
}

func TestDispatchWorkOfCompletedPod(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kl := testKubelet.kubelet
	var got bool
	*kl.GetPodWorkers() = &fakePodWorkers{
		syncPodFn: func(ctx context.Context, updateType kubetypes.SyncPodType, pod, mirrorPod *v1.Pod, podStatus *kubecontainer.PodStatus) (bool, error) {
			got = true
			return false, nil
		},
		cache: *kl.GetPodCache(),
		t:     t,
	}
	now := metav1.NewTime(time.Now())
	pods := []*v1.Pod{
		{
			ObjectMeta: metav1.ObjectMeta{
				UID:         "1",
				Name:        "completed-pod1",
				Namespace:   "ns",
				Annotations: make(map[string]string),
			},
			Status: v1.PodStatus{
				Phase: v1.PodFailed,
				ContainerStatuses: []v1.ContainerStatus{
					{
						State: v1.ContainerState{
							Terminated: &v1.ContainerStateTerminated{},
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				UID:         "2",
				Name:        "completed-pod2",
				Namespace:   "ns",
				Annotations: make(map[string]string),
			},
			Status: v1.PodStatus{
				Phase: v1.PodSucceeded,
				ContainerStatuses: []v1.ContainerStatus{
					{
						State: v1.ContainerState{
							Terminated: &v1.ContainerStateTerminated{},
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				UID:               "3",
				Name:              "completed-pod3",
				Namespace:         "ns",
				Annotations:       make(map[string]string),
				DeletionTimestamp: &now,
			},
			Status: v1.PodStatus{
				ContainerStatuses: []v1.ContainerStatus{
					{
						State: v1.ContainerState{
							Terminated: &v1.ContainerStateTerminated{},
						},
					},
				},
			},
		},
	}
	for _, pod := range pods {
		(*kl.GetPodWorkers()).UpdatePod(kubelet.UpdatePodOptions{
			Pod:        pod,
			UpdateType: kubetypes.SyncPodSync,
			StartTime:  time.Now(),
		})
		if !got {
			t.Errorf("Should not skip completed pod %q", pod.Name)
		}
		got = false
	}
}

func TestDispatchWorkOfActivePod(t *testing.T, randomData string) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kl := testKubelet.kubelet
	var got bool
	*kl.GetPodWorkers() = &fakePodWorkers{
		syncPodFn: func(ctx context.Context, updateType kubetypes.SyncPodType, pod, mirrorPod *v1.Pod, podStatus *kubecontainer.PodStatus) (bool, error) {
			got = true
			return false, nil
		},
		cache: *kl.GetPodCache(),
		t:     t,
	}
	pods := []*v1.Pod{
		{
			ObjectMeta: metav1.ObjectMeta{
				UID:         types.UID(randomData),
				Name:        randomData,
				Namespace:   randomData,
				Annotations: make(map[string]string),
			},
			Status: v1.PodStatus{
				Phase: v1.PodRunning,
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				UID:         types.UID(randomData),
				Name:        randomData,
				Namespace:   randomData,
				Annotations: make(map[string]string),
			},
			Status: v1.PodStatus{
				Phase: v1.PodFailed,
				ContainerStatuses: []v1.ContainerStatus{
					{
						State: v1.ContainerState{
							Running: &v1.ContainerStateRunning{},
						},
					},
				},
			},
		},
	}

	for _, pod := range pods {
		(*kl.GetPodWorkers()).UpdatePod(kubelet.UpdatePodOptions{
			Pod:        pod,
			UpdateType: kubetypes.SyncPodSync,
			StartTime:  time.Now(),
		})
		if !got {
			t.Errorf("Should not skip active pod %q", pod.Name)
		}
		got = false
	}
}

func TestHandlePodCleanups(t *testing.T, pod *kubecontainer.Pod) {
	ctx := context.Background()
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()

	fakeRuntime := testKubelet.fakeRuntime
	fakeRuntime.PodList = []*containertest.FakePod{
		{Pod: pod},
	}
	kubelet := testKubelet.kubelet

	kubelet.HandlePodCleanups(ctx)

	// assert that unwanted pods were queued to kill
	if actual, expected := (*kubelet.GetPodWorkers()).(*fakePodWorkers).triggeredDeletion, []types.UID{pod.ID}; !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %v to be deleted, got %v", expected, actual)
	}
	fakeRuntime.AssertKilledPods([]string(nil))
}

func TestHandlePodRemovesWhenSourcesAreReady(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	ready := false

	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()

	fakePod := &kubecontainer.Pod{
		ID:        "1",
		Name:      "foo",
		Namespace: "new",
		Containers: []*kubecontainer.Container{
			{Name: "bar"},
		},
	}

	pods := []*v1.Pod{
		podWithUIDNameNs("1", "foo", "new"),
	}

	fakeRuntime := testKubelet.fakeRuntime
	fakeRuntime.PodList = []*containertest.FakePod{
		{Pod: fakePod},
	}
	kubelet := testKubelet.kubelet
	*kubelet.GetSourcesReady() = config.NewSourcesReady(func(_ sets.String) bool { return ready })

	kubelet.HandlePodRemoves(pods)
	time.Sleep(2 * time.Second)

	// Sources are not ready yet. Don't remove any pods.
	if expect, actual := []types.UID(nil), (*kubelet.GetPodWorkers()).(*fakePodWorkers).triggeredDeletion; !reflect.DeepEqual(expect, actual) {
		t.Fatalf("expected %v kills, got %v", expect, actual)
	}

	ready = true
	kubelet.HandlePodRemoves(pods)
	time.Sleep(2 * time.Second)

	// Sources are ready. Remove unwanted pods.
	if expect, actual := []types.UID{"1"}, (*kubelet.GetPodWorkers()).(*fakePodWorkers).triggeredDeletion; !reflect.DeepEqual(expect, actual) {
		t.Fatalf("expected %v kills, got %v", expect, actual)
	}
}

type testNodeLister struct {
	nodes []*v1.Node
}

func (nl testNodeLister) Get(name string) (*v1.Node, error) {
	for _, node := range nl.nodes {
		if node.Name == name {
			return node, nil
		}
	}
	return nil, fmt.Errorf("Node with name: %s does not exist", name)
}

func (nl testNodeLister) List(_ labels.Selector) (ret []*v1.Node, err error) {
	return nl.nodes, nil
}

func checkPodStatus(t *testing.T, kl *kubelet.Kubelet, pod *v1.Pod, phase v1.PodPhase) {
	t.Helper()
	status, found := (*kl.GetStatusManager()).GetPodStatus(pod.UID)
	require.True(t, found, "Status of pod %q is not found in the status map", pod.UID)
	require.Equal(t, phase, status.Phase)
}

// Tests that we handle port conflicts correctly by setting the failed status in status map.
func TestHandlePortConflicts(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kl := testKubelet.kubelet

	*kl.GetNodeLister() = testNodeLister{nodes: []*v1.Node{
		{
			ObjectMeta: metav1.ObjectMeta{Name: string(*kl.GetNodeName())},
			Status: v1.NodeStatus{
				Allocatable: v1.ResourceList{
					v1.ResourcePods: *resource.NewQuantity(110, resource.DecimalSI),
				},
			},
		},
	}}

	recorder := record.NewFakeRecorder(20)
	nodeRef := &v1.ObjectReference{
		Kind:      "Node",
		Name:      "testNode",
		UID:       types.UID("testNode"),
		Namespace: "",
	}
	testClusterDNSDomain := "TEST"
	*kl.GetDNSConfigurer() = *dns.NewConfigurer(recorder, nodeRef, nil, nil, testClusterDNSDomain, "")

	spec := v1.PodSpec{NodeName: string(*kl.GetNodeName()), Containers: []v1.Container{{Ports: []v1.ContainerPort{{HostPort: 80}}}}}
	pods := []*v1.Pod{
		podWithUIDNameNsSpec("123456789", "newpod", "foo", spec),
		podWithUIDNameNsSpec("987654321", "oldpod", "foo", spec),
	}
	// Make sure the Pods are in the reverse order of creation time.
	pods[1].CreationTimestamp = metav1.NewTime(time.Now())
	pods[0].CreationTimestamp = metav1.NewTime(time.Now().Add(1 * time.Second))
	// The newer pod should be rejected.
	notfittingPod := pods[0]
	fittingPod := pods[1]
	(*kl.GetPodWorkers()).(*fakePodWorkers).running = map[types.UID]bool{
		pods[0].UID: true,
		pods[1].UID: true,
	}

	kl.HandlePodAdditions(pods)

	// Check pod status stored in the status map.
	checkPodStatus(t, kl, notfittingPod, v1.PodFailed)
	checkPodStatus(t, kl, fittingPod, v1.PodPending)
}

// Tests that we handle host name conflicts correctly by setting the failed status in status map.
func TestHandleHostNameConflicts(t *testing.T, randomData string) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kl := testKubelet.kubelet

	*kl.GetNodeLister() = testNodeLister{nodes: []*v1.Node{
		{
			ObjectMeta: metav1.ObjectMeta{Name: "127.0.0.1"},
			Status: v1.NodeStatus{
				Allocatable: v1.ResourceList{
					v1.ResourcePods: *resource.NewQuantity(110, resource.DecimalSI),
				},
			},
		},
	}}

	recorder := record.NewFakeRecorder(20)
	nodeRef := &v1.ObjectReference{
		Kind:      "Node",
		Name:      "testNode",
		UID:       types.UID("testNode"),
		Namespace: randomData,
	}
	testClusterDNSDomain := "TEST"
	*kl.GetDNSConfigurer() = *dns.NewConfigurer(recorder, nodeRef, nil, nil, testClusterDNSDomain, "")

	// default NodeName in test is 127.0.0.1
	pods := []*v1.Pod{
		podWithUIDNameNsSpec(types.UID(randomData), randomData, randomData, v1.PodSpec{NodeName: "127.0.0.2"}),
		podWithUIDNameNsSpec(types.UID(randomData), randomData, randomData, v1.PodSpec{NodeName: "127.0.0.1"}),
	}

	notfittingPod := pods[0]
	fittingPod := pods[1]

	kl.HandlePodAdditions(pods)

	// Check pod status stored in the status map.
	checkPodStatus(t, kl, notfittingPod, v1.PodFailed)
	checkPodStatus(t, kl, fittingPod, v1.PodPending)
}

// Tests that we handle not matching labels selector correctly by setting the failed status in status map.
func TestHandleNodeSelector(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kl := testKubelet.kubelet
	nodes := []*v1.Node{
		{
			ObjectMeta: metav1.ObjectMeta{Name: testKubeletHostname, Labels: map[string]string{"key": "B"}},
			Status: v1.NodeStatus{
				Allocatable: v1.ResourceList{
					v1.ResourcePods: *resource.NewQuantity(110, resource.DecimalSI),
				},
			},
		},
	}
	*kl.GetNodeLister() = testNodeLister{nodes: nodes}

	recorder := record.NewFakeRecorder(20)
	nodeRef := &v1.ObjectReference{
		Kind:      "Node",
		Name:      "testNode",
		UID:       types.UID("testNode"),
		Namespace: "",
	}
	testClusterDNSDomain := "TEST"
	*kl.GetDNSConfigurer() = *dns.NewConfigurer(recorder, nodeRef, nil, nil, testClusterDNSDomain, "")

	pods := []*v1.Pod{
		podWithUIDNameNsSpec("123456789", "podA", "foo", v1.PodSpec{NodeSelector: map[string]string{"key": "A"}}),
		podWithUIDNameNsSpec("987654321", "podB", "foo", v1.PodSpec{NodeSelector: map[string]string{"key": "B"}}),
	}
	// The first pod should be rejected.
	notfittingPod := pods[0]
	fittingPod := pods[1]

	kl.HandlePodAdditions(pods)

	// Check pod status stored in the status map.
	checkPodStatus(t, kl, notfittingPod, v1.PodFailed)
	checkPodStatus(t, kl, fittingPod, v1.PodPending)
}

// Tests that we handle not matching labels selector correctly by setting the failed status in status map.
func TestHandleNodeSelectorBasedOnOS(t *testing.T) {
	tests := []struct {
		name        string
		nodeLabels  map[string]string
		podSelector map[string]string
		podStatus   v1.PodPhase
	}{
		{
			name:        "correct OS label, wrong pod selector, admission denied",
			nodeLabels:  map[string]string{v1.LabelOSStable: goruntime.GOOS, v1.LabelArchStable: goruntime.GOARCH},
			podSelector: map[string]string{v1.LabelOSStable: "dummyOS"},
			podStatus:   v1.PodFailed,
		},
		{
			name:        "correct OS label, correct pod selector, admission denied",
			nodeLabels:  map[string]string{v1.LabelOSStable: goruntime.GOOS, v1.LabelArchStable: goruntime.GOARCH},
			podSelector: map[string]string{v1.LabelOSStable: goruntime.GOOS},
			podStatus:   v1.PodPending,
		},
		{
			// Expect no patching to happen, label B should be preserved and can be used for nodeAffinity.
			name:        "new node label, correct pod selector, admitted",
			nodeLabels:  map[string]string{v1.LabelOSStable: goruntime.GOOS, v1.LabelArchStable: goruntime.GOARCH, "key": "B"},
			podSelector: map[string]string{"key": "B"},
			podStatus:   v1.PodPending,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
			defer testKubelet.Cleanup()
			kl := testKubelet.kubelet
			nodes := []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{Name: testKubeletHostname, Labels: test.nodeLabels},
					Status: v1.NodeStatus{
						Allocatable: v1.ResourceList{
							v1.ResourcePods: *resource.NewQuantity(110, resource.DecimalSI),
						},
					},
				},
			}
			*kl.GetNodeLister() = testNodeLister{nodes: nodes}

			recorder := record.NewFakeRecorder(20)
			nodeRef := &v1.ObjectReference{
				Kind:      "Node",
				Name:      "testNode",
				UID:       types.UID("testNode"),
				Namespace: "",
			}
			testClusterDNSDomain := "TEST"
			*kl.GetDNSConfigurer() = *dns.NewConfigurer(recorder, nodeRef, nil, nil, testClusterDNSDomain, "")

			pod := podWithUIDNameNsSpec("123456789", "podA", "foo", v1.PodSpec{NodeSelector: test.podSelector})

			kl.HandlePodAdditions([]*v1.Pod{pod})

			// Check pod status stored in the status map.
			checkPodStatus(t, kl, pod, test.podStatus)
		})
	}
}

// Tests that we handle exceeded resources correctly by setting the failed status in status map.
func TestHandleMemExceeded(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kl := testKubelet.kubelet
	nodes := []*v1.Node{
		{ObjectMeta: metav1.ObjectMeta{Name: testKubeletHostname},
			Status: v1.NodeStatus{Capacity: v1.ResourceList{}, Allocatable: v1.ResourceList{
				v1.ResourceCPU:    *resource.NewMilliQuantity(10, resource.DecimalSI),
				v1.ResourceMemory: *resource.NewQuantity(100, resource.BinarySI),
				v1.ResourcePods:   *resource.NewQuantity(40, resource.DecimalSI),
			}}},
	}
	*kl.GetNodeLister() = testNodeLister{nodes: nodes}

	recorder := record.NewFakeRecorder(20)
	nodeRef := &v1.ObjectReference{
		Kind:      "Node",
		Name:      "testNode",
		UID:       types.UID("testNode"),
		Namespace: "",
	}
	testClusterDNSDomain := "TEST"
	*kl.GetDNSConfigurer() = *dns.NewConfigurer(recorder, nodeRef, nil, nil, testClusterDNSDomain, "")

	spec := v1.PodSpec{NodeName: string(*kl.GetNodeName()),
		Containers: []v1.Container{{Resources: v1.ResourceRequirements{
			Requests: v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("90"),
			},
		}}},
	}
	pods := []*v1.Pod{
		podWithUIDNameNsSpec("123456789", "newpod", "foo", spec),
		podWithUIDNameNsSpec("987654321", "oldpod", "foo", spec),
	}
	// Make sure the Pods are in the reverse order of creation time.
	pods[1].CreationTimestamp = metav1.NewTime(time.Now())
	pods[0].CreationTimestamp = metav1.NewTime(time.Now().Add(1 * time.Second))
	// The newer pod should be rejected.
	notfittingPod := pods[0]
	fittingPod := pods[1]
	(*kl.GetPodWorkers()).(*fakePodWorkers).running = map[types.UID]bool{
		pods[0].UID: true,
		pods[1].UID: true,
	}

	kl.HandlePodAdditions(pods)

	// Check pod status stored in the status map.
	checkPodStatus(t, kl, notfittingPod, v1.PodFailed)
	checkPodStatus(t, kl, fittingPod, v1.PodPending)
}

// Tests that we handle result of interface UpdatePluginResources correctly
// by setting corresponding status in status map.
func TestHandlePluginResources(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kl := testKubelet.kubelet

	adjustedResource := v1.ResourceName("domain1.com/adjustedResource")
	emptyResource := v1.ResourceName("domain2.com/emptyResource")
	missingResource := v1.ResourceName("domain2.com/missingResource")
	failedResource := v1.ResourceName("domain2.com/failedResource")
	resourceQuantity0 := *resource.NewQuantity(int64(0), resource.DecimalSI)
	resourceQuantity1 := *resource.NewQuantity(int64(1), resource.DecimalSI)
	resourceQuantity2 := *resource.NewQuantity(int64(2), resource.DecimalSI)
	resourceQuantityInvalid := *resource.NewQuantity(int64(-1), resource.DecimalSI)
	allowedPodQuantity := *resource.NewQuantity(int64(10), resource.DecimalSI)
	nodes := []*v1.Node{
		{ObjectMeta: metav1.ObjectMeta{Name: testKubeletHostname},
			Status: v1.NodeStatus{Capacity: v1.ResourceList{}, Allocatable: v1.ResourceList{
				adjustedResource: resourceQuantity1,
				emptyResource:    resourceQuantity0,
				v1.ResourcePods:  allowedPodQuantity,
			}}},
	}
	*kl.GetNodeLister() = testNodeLister{nodes: nodes}

	updatePluginResourcesFunc := func(node *schedulerframework.NodeInfo, attrs *lifecycle.PodAdmitAttributes) error {
		// Maps from resourceName to the value we use to set node.allocatableResource[resourceName].
		// A resource with invalid value (< 0) causes the function to return an error
		// to emulate resource Allocation failure.
		// Resources not contained in this map will have their node.allocatableResource
		// quantity unchanged.
		updateResourceMap := map[v1.ResourceName]resource.Quantity{
			adjustedResource: resourceQuantity2,
			emptyResource:    resourceQuantity0,
			failedResource:   resourceQuantityInvalid,
		}
		pod := attrs.Pod
		newAllocatableResource := node.Allocatable.Clone()
		for _, container := range pod.Spec.Containers {
			for resource := range container.Resources.Requests {
				newQuantity, exist := updateResourceMap[resource]
				if !exist {
					continue
				}
				if newQuantity.Value() < 0 {
					return fmt.Errorf("Allocation failed")
				}
				newAllocatableResource.ScalarResources[resource] = newQuantity.Value()
			}
		}
		node.Allocatable = newAllocatableResource
		return nil
	}

	// add updatePluginResourcesFunc to admission handler, to test it's behavior.
	*kl.AdmitHandlers() = lifecycle.PodAdmitHandlers{}
	(*kl.AdmitHandlers()).AddPodAdmitHandler(lifecycle.NewPredicateAdmitHandler(kl.GetNodeAnyWay, lifecycle.NewAdmissionFailureHandlerStub(), updatePluginResourcesFunc))

	recorder := record.NewFakeRecorder(20)
	nodeRef := &v1.ObjectReference{
		Kind:      "Node",
		Name:      "testNode",
		UID:       types.UID("testNode"),
		Namespace: "",
	}
	testClusterDNSDomain := "TEST"
	*kl.GetDNSConfigurer() = *dns.NewConfigurer(recorder, nodeRef, nil, nil, testClusterDNSDomain, "")

	// pod requiring adjustedResource can be successfully allocated because updatePluginResourcesFunc
	// adjusts node.allocatableResource for this resource to a sufficient value.
	fittingPodSpec := v1.PodSpec{NodeName: string(*kl.GetNodeName()),
		Containers: []v1.Container{{Resources: v1.ResourceRequirements{
			Limits: v1.ResourceList{
				adjustedResource: resourceQuantity2,
			},
			Requests: v1.ResourceList{
				adjustedResource: resourceQuantity2,
			},
		}}},
	}
	// pod requiring emptyResource (extended resources with 0 allocatable) will
	// not pass PredicateAdmit.
	emptyPodSpec := v1.PodSpec{NodeName: string(*kl.GetNodeName()),
		Containers: []v1.Container{{Resources: v1.ResourceRequirements{
			Limits: v1.ResourceList{
				emptyResource: resourceQuantity2,
			},
			Requests: v1.ResourceList{
				emptyResource: resourceQuantity2,
			},
		}}},
	}
	// pod requiring missingResource will pass PredicateAdmit.
	//
	// Extended resources missing in node status are ignored in PredicateAdmit.
	// This is required to support extended resources that are not managed by
	// device plugin, such as cluster-level resources.
	missingPodSpec := v1.PodSpec{NodeName: string(*kl.GetNodeName()),
		Containers: []v1.Container{{Resources: v1.ResourceRequirements{
			Limits: v1.ResourceList{
				missingResource: resourceQuantity2,
			},
			Requests: v1.ResourceList{
				missingResource: resourceQuantity2,
			},
		}}},
	}
	// pod requiring failedResource will fail with the resource failed to be allocated.
	failedPodSpec := v1.PodSpec{NodeName: string(*kl.GetNodeName()),
		Containers: []v1.Container{{Resources: v1.ResourceRequirements{
			Limits: v1.ResourceList{
				failedResource: resourceQuantity1,
			},
			Requests: v1.ResourceList{
				failedResource: resourceQuantity1,
			},
		}}},
	}

	fittingPod := podWithUIDNameNsSpec("1", "fittingpod", "foo", fittingPodSpec)
	emptyPod := podWithUIDNameNsSpec("2", "emptypod", "foo", emptyPodSpec)
	missingPod := podWithUIDNameNsSpec("3", "missingpod", "foo", missingPodSpec)
	failedPod := podWithUIDNameNsSpec("4", "failedpod", "foo", failedPodSpec)

	kl.HandlePodAdditions([]*v1.Pod{fittingPod, emptyPod, missingPod, failedPod})

	// Check pod status stored in the status map.
	checkPodStatus(t, kl, fittingPod, v1.PodPending)
	checkPodStatus(t, kl, emptyPod, v1.PodFailed)
	checkPodStatus(t, kl, missingPod, v1.PodPending)
	checkPodStatus(t, kl, failedPod, v1.PodFailed)
}

// TODO(filipg): This test should be removed once StatusSyncer can do garbage collection without external signal.
func TestPurgingObsoleteStatusMapEntries(t *testing.T) {
	ctx := context.Background()
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()

	kl := testKubelet.kubelet
	pods := []*v1.Pod{
		{ObjectMeta: metav1.ObjectMeta{Name: "pod1", UID: "1234"}, Spec: v1.PodSpec{Containers: []v1.Container{{Ports: []v1.ContainerPort{{HostPort: 80}}}}}},
		{ObjectMeta: metav1.ObjectMeta{Name: "pod2", UID: "4567"}, Spec: v1.PodSpec{Containers: []v1.Container{{Ports: []v1.ContainerPort{{HostPort: 80}}}}}},
	}
	podToTest := pods[1]
	// Run once to populate the status map.
	kl.HandlePodAdditions(pods)
	if _, found := (*kl.GetStatusManager()).GetPodStatus(podToTest.UID); !found {
		t.Fatalf("expected to have status cached for pod2")
	}
	// Sync with empty pods so that the entry in status map will be removed.
	(*kl.GetPodManager()).SetPods([]*v1.Pod{})
	kl.HandlePodCleanups(ctx)
	if _, found := (*kl.GetStatusManager()).GetPodStatus(podToTest.UID); found {
		t.Fatalf("expected to not have status cached for pod2")
	}
}

func TestValidateContainerLogStatus(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kubelet := testKubelet.kubelet
	containerName := "x"
	testCases := []struct {
		statuses []v1.ContainerStatus
		success  bool // whether getting logs for the container should succeed.
		pSuccess bool // whether getting logs for the previous container should succeed.
	}{
		{
			statuses: []v1.ContainerStatus{
				{
					Name: containerName,
					State: v1.ContainerState{
						Running: &v1.ContainerStateRunning{},
					},
					LastTerminationState: v1.ContainerState{
						Terminated: &v1.ContainerStateTerminated{ContainerID: "docker://fakeid"},
					},
				},
			},
			success:  true,
			pSuccess: true,
		},
		{
			statuses: []v1.ContainerStatus{
				{
					Name: containerName,
					State: v1.ContainerState{
						Running: &v1.ContainerStateRunning{},
					},
				},
			},
			success:  true,
			pSuccess: false,
		},
		{
			statuses: []v1.ContainerStatus{
				{
					Name: containerName,
					State: v1.ContainerState{
						Terminated: &v1.ContainerStateTerminated{},
					},
				},
			},
			success:  false,
			pSuccess: false,
		},
		{
			statuses: []v1.ContainerStatus{
				{
					Name: containerName,
					State: v1.ContainerState{
						Terminated: &v1.ContainerStateTerminated{ContainerID: "docker://fakeid"},
					},
				},
			},
			success:  true,
			pSuccess: false,
		},
		{
			statuses: []v1.ContainerStatus{
				{
					Name: containerName,
					State: v1.ContainerState{
						Terminated: &v1.ContainerStateTerminated{},
					},
					LastTerminationState: v1.ContainerState{
						Terminated: &v1.ContainerStateTerminated{},
					},
				},
			},
			success:  false,
			pSuccess: false,
		},
		{
			statuses: []v1.ContainerStatus{
				{
					Name: containerName,
					State: v1.ContainerState{
						Terminated: &v1.ContainerStateTerminated{},
					},
					LastTerminationState: v1.ContainerState{
						Terminated: &v1.ContainerStateTerminated{ContainerID: "docker://fakeid"},
					},
				},
			},
			success:  true,
			pSuccess: true,
		},
		{
			statuses: []v1.ContainerStatus{
				{
					Name: containerName,
					State: v1.ContainerState{
						Waiting: &v1.ContainerStateWaiting{},
					},
				},
			},
			success:  false,
			pSuccess: false,
		},
		{
			statuses: []v1.ContainerStatus{
				{
					Name:  containerName,
					State: v1.ContainerState{Waiting: &v1.ContainerStateWaiting{Reason: "ErrImagePull"}},
				},
			},
			success:  false,
			pSuccess: false,
		},
		{
			statuses: []v1.ContainerStatus{
				{
					Name:  containerName,
					State: v1.ContainerState{Waiting: &v1.ContainerStateWaiting{Reason: "ErrImagePullBackOff"}},
				},
			},
			success:  false,
			pSuccess: false,
		},
	}

	for i, tc := range testCases {
		// Access the log of the most recent container
		previous := false
		podStatus := &v1.PodStatus{ContainerStatuses: tc.statuses}
		_, err := kubelet.ValidateContainerLogStatus("podName", podStatus, containerName, previous)
		if !tc.success {
			assert.Error(t, err, fmt.Sprintf("[case %d] error", i))
		} else {
			assert.NoError(t, err, "[case %d] error", i)
		}
		// Access the log of the previous, terminated container
		previous = true
		_, err = kubelet.ValidateContainerLogStatus("podName", podStatus, containerName, previous)
		if !tc.pSuccess {
			assert.Error(t, err, fmt.Sprintf("[case %d] error", i))
		} else {
			assert.NoError(t, err, "[case %d] error", i)
		}
		// Access the log of a container that's not in the pod
		_, err = kubelet.ValidateContainerLogStatus("podName", podStatus, "blah", false)
		assert.Error(t, err, fmt.Sprintf("[case %d] invalid container name should cause an error", i))
	}
}

func TestCreateMirrorPod(t *testing.T) {
	tests := []struct {
		name       string
		updateType kubetypes.SyncPodType
	}{
		{
			name:       "SyncPodCreate",
			updateType: kubetypes.SyncPodCreate,
		},
		{
			name:       "SyncPodUpdate",
			updateType: kubetypes.SyncPodUpdate,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
			defer testKubelet.Cleanup()

			kl := testKubelet.kubelet
			manager := testKubelet.fakeMirrorClient
			pod := podWithUIDNameNs("12345678", "bar", "foo")
			pod.Annotations[kubetypes.ConfigSourceAnnotationKey] = "file"
			pods := []*v1.Pod{pod}
			(*kl.GetPodManager()).SetPods(pods)
			isTerminal, err := kl.SyncPod(context.Background(), tt.updateType, pod, nil, &kubecontainer.PodStatus{})
			assert.NoError(t, err)
			if isTerminal {
				t.Fatalf("pod should not be terminal: %#v", pod)
			}
			podFullName := kubecontainer.GetPodFullName(pod)
			assert.True(t, manager.HasPod(podFullName), "Expected mirror pod %q to be created", podFullName)
			assert.Equal(t, 1, manager.NumOfPods(), "Expected only 1 mirror pod %q, got %+v", podFullName, manager.GetPods())
		})
	}
}

func TestDeleteOutdatedMirrorPod(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()

	kl := testKubelet.kubelet
	manager := testKubelet.fakeMirrorClient
	pod := podWithUIDNameNsSpec("12345678", "foo", "ns", v1.PodSpec{
		Containers: []v1.Container{
			{Name: "1234", Image: "foo"},
		},
	})
	pod.Annotations[kubetypes.ConfigSourceAnnotationKey] = "file"

	// Mirror pod has an outdated spec.
	mirrorPod := podWithUIDNameNsSpec("11111111", "foo", "ns", v1.PodSpec{
		Containers: []v1.Container{
			{Name: "1234", Image: "bar"},
		},
	})
	mirrorPod.Annotations[kubetypes.ConfigSourceAnnotationKey] = "api"
	mirrorPod.Annotations[kubetypes.ConfigMirrorAnnotationKey] = "mirror"

	pods := []*v1.Pod{pod, mirrorPod}
	(*kl.GetPodManager()).SetPods(pods)
	isTerminal, err := kl.SyncPod(context.Background(), kubetypes.SyncPodUpdate, pod, mirrorPod, &kubecontainer.PodStatus{})
	assert.NoError(t, err)
	if isTerminal {
		t.Fatalf("pod should not be terminal: %#v", pod)
	}
	name := kubecontainer.GetPodFullName(pod)
	creates, deletes := manager.GetCounts(name)
	if creates != 1 || deletes != 1 {
		t.Errorf("expected 1 creation and 1 deletion of %q, got %d, %d", name, creates, deletes)
	}
}

func TestDeleteOrphanedMirrorPods(t *testing.T) {
	ctx := context.Background()
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()

	kl := testKubelet.kubelet
	manager := testKubelet.fakeMirrorClient
	orphanPods := []*v1.Pod{
		{
			ObjectMeta: metav1.ObjectMeta{
				UID:       "12345678",
				Name:      "pod1",
				Namespace: "ns",
				Annotations: map[string]string{
					kubetypes.ConfigSourceAnnotationKey: "api",
					kubetypes.ConfigMirrorAnnotationKey: "mirror",
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				UID:       "12345679",
				Name:      "pod2",
				Namespace: "ns",
				Annotations: map[string]string{
					kubetypes.ConfigSourceAnnotationKey: "api",
					kubetypes.ConfigMirrorAnnotationKey: "mirror",
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				UID:       "12345670",
				Name:      "pod3",
				Namespace: "ns",
				Annotations: map[string]string{
					kubetypes.ConfigSourceAnnotationKey: "api",
					kubetypes.ConfigMirrorAnnotationKey: "mirror",
				},
			},
		},
	}

	(*kl.GetPodManager()).SetPods(orphanPods)

	// a static pod that is terminating will not be deleted
	(*kl.GetPodWorkers()).(*fakePodWorkers).terminatingStaticPods = map[string]bool{
		kubecontainer.GetPodFullName(orphanPods[2]): true,
	}

	// Sync with an empty pod list to delete all mirror pods.
	kl.HandlePodCleanups(ctx)
	assert.Len(t, manager.GetPods(), 0, "Expected 0 mirror pods")
	for i, pod := range orphanPods {
		name := kubecontainer.GetPodFullName(pod)
		creates, deletes := manager.GetCounts(name)
		switch i {
		case 2:
			if creates != 0 || deletes != 0 {
				t.Errorf("expected 0 creation and 0 deletion of %q, got %d, %d", name, creates, deletes)
			}
		default:
			if creates != 0 || deletes != 1 {
				t.Errorf("expected 0 creation and one deletion of %q, got %d, %d", name, creates, deletes)
			}
		}
	}
}

func TestNetworkErrorsWithoutHostNetwork(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kubelet := testKubelet.kubelet

	(*kubelet.GetRuntimeState()).SetNetworkState(fmt.Errorf("simulated network error"))

	pod := podWithUIDNameNsSpec("12345678", "hostnetwork", "new", v1.PodSpec{
		HostNetwork: false,

		Containers: []v1.Container{
			{Name: "foo"},
		},
	})

	(*kubelet.GetPodManager()).SetPods([]*v1.Pod{pod})
	isTerminal, err := kubelet.SyncPod(context.Background(), kubetypes.SyncPodUpdate, pod, nil, &kubecontainer.PodStatus{})
	assert.Error(t, err, "expected pod with hostNetwork=false to fail when network in error")
	if isTerminal {
		t.Fatalf("pod should not be terminal: %#v", pod)
	}

	pod.Annotations[kubetypes.ConfigSourceAnnotationKey] = kubetypes.FileSource
	pod.Spec.HostNetwork = true
	isTerminal, err = kubelet.SyncPod(context.Background(), kubetypes.SyncPodUpdate, pod, nil, &kubecontainer.PodStatus{})
	assert.NoError(t, err, "expected pod with hostNetwork=true to succeed when network in error")
	if isTerminal {
		t.Fatalf("pod should not be terminal: %#v", pod)
	}
}

func TestFilterOutInactivePods(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kubelet := testKubelet.kubelet
	pods := newTestPods(8)
	now := metav1.NewTime(time.Now())

	// terminal pods are excluded
	pods[0].Status.Phase = v1.PodFailed
	pods[1].Status.Phase = v1.PodSucceeded

	// deleted pod is included unless it's known to be terminated
	pods[2].Status.Phase = v1.PodRunning
	pods[2].DeletionTimestamp = &now
	pods[2].Status.ContainerStatuses = []v1.ContainerStatus{
		{State: v1.ContainerState{
			Running: &v1.ContainerStateRunning{
				StartedAt: now,
			},
		}},
	}

	// pending and running pods are included
	pods[3].Status.Phase = v1.PodPending
	pods[4].Status.Phase = v1.PodRunning

	// pod that is running but has been rejected by admission is excluded
	pods[5].Status.Phase = v1.PodRunning
	(*kubelet.GetStatusManager()).SetPodStatus(pods[5], v1.PodStatus{Phase: v1.PodFailed})

	// pod that is running according to the api but is known terminated is excluded
	pods[6].Status.Phase = v1.PodRunning
	(*kubelet.GetPodWorkers()).(*fakePodWorkers).terminated = map[types.UID]bool{
		pods[6].UID: true,
	}

	// pod that is failed but still terminating is included (it may still be consuming
	// resources)
	pods[7].Status.Phase = v1.PodFailed
	(*kubelet.GetPodWorkers()).(*fakePodWorkers).terminationRequested = map[types.UID]bool{
		pods[7].UID: true,
	}

	expected := []*v1.Pod{pods[2], pods[3], pods[4], pods[7]}
	(*kubelet.GetPodManager()).SetPods(pods)
	actual := kubelet.FilterOutInactivePods(pods)
	assert.Equal(t, expected, actual)
}

func TestCheckpointContainer(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kubelet := testKubelet.kubelet

	fakeRuntime := testKubelet.fakeRuntime
	containerID := kubecontainer.ContainerID{
		Type: "test",
		ID:   "abc1234",
	}

	fakePod := &containertest.FakePod{
		Pod: &kubecontainer.Pod{
			ID:        "12345678",
			Name:      "podFoo",
			Namespace: "nsFoo",
			Containers: []*kubecontainer.Container{
				{
					Name: "containerFoo",
					ID:   containerID,
				},
			},
		},
	}

	fakeRuntime.PodList = []*containertest.FakePod{fakePod}
	wrongContainerName := "wrongContainerName"

	tests := []struct {
		name               string
		containerName      string
		checkpointLocation string
		expectedStatus     error
		expectedLocation   string
	}{
		{
			name:               "Checkpoint with wrong container name",
			containerName:      wrongContainerName,
			checkpointLocation: "",
			expectedStatus:     fmt.Errorf("container %s not found", wrongContainerName),
			expectedLocation:   "",
		},
		{
			name:               "Checkpoint with default checkpoint location",
			containerName:      fakePod.Pod.Containers[0].Name,
			checkpointLocation: "",
			expectedStatus:     nil,
			expectedLocation: filepath.Join(
				kubelet.GetCheckpointsDir(),
				fmt.Sprintf(
					"checkpoint-%s_%s-%s",
					fakePod.Pod.Name,
					fakePod.Pod.Namespace,
					fakePod.Pod.Containers[0].Name,
				),
			),
		},
		{
			name:               "Checkpoint with ignored location",
			containerName:      fakePod.Pod.Containers[0].Name,
			checkpointLocation: "somethingThatWillBeIgnored",
			expectedStatus:     nil,
			expectedLocation: filepath.Join(
				kubelet.GetCheckpointsDir(),
				fmt.Sprintf(
					"checkpoint-%s_%s-%s",
					fakePod.Pod.Name,
					fakePod.Pod.Namespace,
					fakePod.Pod.Containers[0].Name,
				),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			options := &runtimeapi.CheckpointContainerRequest{}
			if test.checkpointLocation != "" {
				options.Location = test.checkpointLocation
			}
			status := kubelet.CheckpointContainer(
				ctx,
				fakePod.Pod.ID,
				fmt.Sprintf(
					"%s_%s",
					fakePod.Pod.Name,
					fakePod.Pod.Namespace,
				),
				test.containerName,
				options,
			)
			require.Equal(t, test.expectedStatus, status)

			if status != nil {
				return
			}

			require.True(
				t,
				strings.HasPrefix(
					options.Location,
					test.expectedLocation,
				),
			)
			require.Equal(
				t,
				options.ContainerId,
				containerID.ID,
			)

		})
	}
}

// TestSyncPodsSetStatusToFailedForPodsThatRunTooLong checking fuzz
func TestSyncPodsSetStatusToFailedForPodsThatRunTooLong(t *testing.T, pod *v1.Pod) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	fakeRuntime := testKubelet.fakeRuntime
	kubelet := testKubelet.kubelet

	pods := []*v1.Pod{pod}

	fakeRuntime.PodList = []*containertest.FakePod{
		{Pod: &kubecontainer.Pod{
			ID:        "12345678",
			Name:      "bar",
			Namespace: "new",
			Containers: []*kubecontainer.Container{
				{Name: "foo"},
			},
		}},
	}

	// Let the pod worker sets the status to fail after this sync.
	kubelet.HandlePodUpdates(pods)
	status, found := (*kubelet.GetStatusManager()).GetPodStatus(pods[0].UID)
	assert.True(t, found, "expected to found status for pod %q", pods[0].UID)
	assert.Equal(t, v1.PodFailed, status.Phase)
	// check pod status contains ContainerStatuses, etc.
	assert.NotNil(t, status.ContainerStatuses)
}

func TestSyncPodsDoesNotSetPodsThatDidNotRunTooLongToFailed(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	fakeRuntime := testKubelet.fakeRuntime

	kubelet := testKubelet.kubelet

	now := metav1.Now()
	startTime := metav1.NewTime(now.Time.Add(-1 * time.Minute))
	exceededActiveDeadlineSeconds := int64(300)

	pods := []*v1.Pod{
		{
			ObjectMeta: metav1.ObjectMeta{
				UID:       "12345678",
				Name:      "bar",
				Namespace: "new",
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{Name: "foo"},
				},
				ActiveDeadlineSeconds: &exceededActiveDeadlineSeconds,
			},
			Status: v1.PodStatus{
				StartTime: &startTime,
			},
		},
	}

	fakeRuntime.PodList = []*containertest.FakePod{
		{Pod: &kubecontainer.Pod{
			ID:        "12345678",
			Name:      "bar",
			Namespace: "new",
			Containers: []*kubecontainer.Container{
				{Name: "foo"},
			},
		}},
	}

	(*kubelet.GetPodManager()).SetPods(pods)
	kubelet.HandlePodUpdates(pods)
	status, found := (*kubelet.GetStatusManager()).GetPodStatus(pods[0].UID)
	assert.True(t, found, "expected to found status for pod %q", pods[0].UID)
	assert.NotEqual(t, v1.PodFailed, status.Phase)
}

func podWithUIDNameNs(uid types.UID, name, namespace string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			UID:         uid,
			Name:        name,
			Namespace:   namespace,
			Annotations: map[string]string{},
		},
	}
}

func podWithUIDNameNsSpec(uid types.UID, name, namespace string, spec v1.PodSpec) *v1.Pod {
	pod := podWithUIDNameNs(uid, name, namespace)
	pod.Spec = spec
	return pod
}

func TestDeletePodDirsForDeletedPods(t *testing.T) {
	ctx := context.Background()
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kl := testKubelet.kubelet
	pods := []*v1.Pod{
		podWithUIDNameNs("12345678", "pod1", "ns"),
		podWithUIDNameNs("12345679", "pod2", "ns"),
	}

	(*kl.GetPodManager()).SetPods(pods)
	// Sync to create pod directories.
	kl.HandlePodSyncs((*kl.GetPodManager()).GetPods())
	for i := range pods {
		assert.True(t, dirExists(kl.GetPodDir(pods[i].UID)), "Expected directory to exist for pod %d", i)
	}

	// Pod 1 has been deleted and no longer exists.
	(*kl.GetPodManager()).SetPods([]*v1.Pod{pods[0]})
	kl.HandlePodCleanups(ctx)
	assert.True(t, dirExists(kl.GetPodDir(pods[0].UID)), "Expected directory to exist for pod 0")
	assert.False(t, dirExists(kl.GetPodDir(pods[1].UID)), "Expected directory to be deleted for pod 1")
}

func syncAndVerifyPodDir(t *testing.T, testKubelet *TestKubelet, pods []*v1.Pod, podsToCheck []*v1.Pod, shouldExist bool) {
	ctx := context.Background()
	t.Helper()
	kl := testKubelet.kubelet

	(*kl.GetPodManager()).SetPods(pods)
	kl.HandlePodSyncs(pods)
	kl.HandlePodCleanups(ctx)
	for i, pod := range podsToCheck {
		exist := dirExists(kl.GetPodDir(pod.UID))
		assert.Equal(t, shouldExist, exist, "directory of pod %d", i)
	}
}

func TestDoesNotDeletePodDirsForTerminatedPods(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kl := testKubelet.kubelet
	pods := []*v1.Pod{
		podWithUIDNameNs("12345678", "pod1", "ns"),
		podWithUIDNameNs("12345679", "pod2", "ns"),
		podWithUIDNameNs("12345680", "pod3", "ns"),
	}
	syncAndVerifyPodDir(t, testKubelet, pods, pods, true)
	// Pod 1 failed, and pod 2 succeeded. None of the pod directories should be
	// deleted.
	(*kl.GetStatusManager()).SetPodStatus(pods[1], v1.PodStatus{Phase: v1.PodFailed})
	(*kl.GetStatusManager()).SetPodStatus(pods[2], v1.PodStatus{Phase: v1.PodSucceeded})
	syncAndVerifyPodDir(t, testKubelet, pods, pods, true)
}

func TestDoesNotDeletePodDirsIfContainerIsRunning(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	runningPod := &kubecontainer.Pod{
		ID:        "12345678",
		Name:      "pod1",
		Namespace: "ns",
	}
	apiPod := podWithUIDNameNs(runningPod.ID, runningPod.Name, runningPod.Namespace)

	// Sync once to create pod directory; confirm that the pod directory has
	// already been created.
	pods := []*v1.Pod{apiPod}
	(*testKubelet.kubelet.GetPodWorkers()).(*fakePodWorkers).running = map[types.UID]bool{apiPod.UID: true}
	syncAndVerifyPodDir(t, testKubelet, pods, []*v1.Pod{apiPod}, true)

	// Pretend the pod is deleted from apiserver, but is still active on the node.
	// The pod directory should not be removed.
	pods = []*v1.Pod{}
	testKubelet.fakeRuntime.PodList = []*containertest.FakePod{{Pod: runningPod, NetnsPath: ""}}
	syncAndVerifyPodDir(t, testKubelet, pods, []*v1.Pod{apiPod}, true)

	// The pod is deleted and also not active on the node. The pod directory
	// should be removed.
	pods = []*v1.Pod{}
	testKubelet.fakeRuntime.PodList = []*containertest.FakePod{}
	(*testKubelet.kubelet.GetPodWorkers()).(*fakePodWorkers).running = nil
	syncAndVerifyPodDir(t, testKubelet, pods, []*v1.Pod{apiPod}, false)
}

func TestGetPodsToSync(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kubelet := testKubelet.kubelet
	clock := testKubelet.fakeClock
	pods := newTestPods(5)

	exceededActiveDeadlineSeconds := int64(30)
	notYetActiveDeadlineSeconds := int64(120)
	startTime := metav1.NewTime(clock.Now())
	pods[0].Status.StartTime = &startTime
	pods[0].Spec.ActiveDeadlineSeconds = &exceededActiveDeadlineSeconds
	pods[1].Status.StartTime = &startTime
	pods[1].Spec.ActiveDeadlineSeconds = &notYetActiveDeadlineSeconds
	pods[2].Status.StartTime = &startTime
	pods[2].Spec.ActiveDeadlineSeconds = &exceededActiveDeadlineSeconds

	(*kubelet.GetPodManager()).SetPods(pods)
	(*kubelet.GetWorkQueue()).Enqueue(pods[2].UID, 0)
	(*kubelet.GetWorkQueue()).Enqueue(pods[3].UID, 30*time.Second)
	(*kubelet.GetWorkQueue()).Enqueue(pods[4].UID, 2*time.Minute)

	clock.Step(1 * time.Minute)

	expected := []*v1.Pod{pods[2], pods[3], pods[0]}
	podsToSync := kubelet.GetPodsToSync()
	sort.Sort(podsByUID(expected))
	sort.Sort(podsByUID(podsToSync))
	assert.Equal(t, expected, podsToSync)
}

func TestGenerateAPIPodStatusWithSortedContainers(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kubelet := testKubelet.kubelet
	numContainers := 10
	expectedOrder := []string{}
	cStatuses := []*kubecontainer.Status{}
	specContainerList := []v1.Container{}
	for i := 0; i < numContainers; i++ {
		id := fmt.Sprintf("%v", i)
		containerName := fmt.Sprintf("%vcontainer", id)
		expectedOrder = append(expectedOrder, containerName)
		cStatus := &kubecontainer.Status{
			ID:   kubecontainer.BuildContainerID("test", id),
			Name: containerName,
		}
		// Rearrange container statuses
		if i%2 == 0 {
			cStatuses = append(cStatuses, cStatus)
		} else {
			cStatuses = append([]*kubecontainer.Status{cStatus}, cStatuses...)
		}
		specContainerList = append(specContainerList, v1.Container{Name: containerName})
	}
	pod := podWithUIDNameNs("uid1", "foo", "test")
	pod.Spec = v1.PodSpec{
		Containers: specContainerList,
	}

	status := &kubecontainer.PodStatus{
		ID:                pod.UID,
		Name:              pod.Name,
		Namespace:         pod.Namespace,
		ContainerStatuses: cStatuses,
	}
	for i := 0; i < 5; i++ {
		apiStatus := kubelet.GenerateAPIPodStatus(pod, status, false)
		for i, c := range apiStatus.ContainerStatuses {
			if expectedOrder[i] != c.Name {
				t.Fatalf("Container status not sorted, expected %v at index %d, but found %v", expectedOrder[i], i, c.Name)
			}
		}
	}
}

func verifyContainerStatuses(t *testing.T, statuses []v1.ContainerStatus, expectedState, expectedLastTerminationState map[string]v1.ContainerState, message string) {
	for _, s := range statuses {
		assert.Equal(t, expectedState[s.Name], s.State, "%s: state", message)
		assert.Equal(t, expectedLastTerminationState[s.Name], s.LastTerminationState, "%s: last terminated state", message)
	}
}

// Test generateAPIPodStatus with different reason cache and old api pod status.
func TestGenerateAPIPodStatusWithReasonCache(t *testing.T) {
	// The following waiting reason and message  are generated in convertStatusToAPIStatus()
	testTimestamp := time.Unix(123456789, 987654321)
	testErrorReason := fmt.Errorf("test-error")
	emptyContainerID := (&kubecontainer.ContainerID{}).String()
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kl := testKubelet.kubelet
	pod := podWithUIDNameNs("12345678", "foo", "new")
	pod.Spec = v1.PodSpec{RestartPolicy: v1.RestartPolicyOnFailure}

	podStatus := &kubecontainer.PodStatus{
		ID:        pod.UID,
		Name:      pod.Name,
		Namespace: pod.Namespace,
	}
	tests := []struct {
		containers    []v1.Container
		statuses      []*kubecontainer.Status
		reasons       map[string]error
		oldStatuses   []v1.ContainerStatus
		expectedState map[string]v1.ContainerState
		// Only set expectedInitState when it is different from expectedState
		expectedInitState            map[string]v1.ContainerState
		expectedLastTerminationState map[string]v1.ContainerState
	}{
		// For container with no historical record, State should be Waiting, LastTerminationState should be retrieved from
		// old status from apiserver.
		{
			containers: []v1.Container{{Name: "without-old-record"}, {Name: "with-old-record"}},
			statuses:   []*kubecontainer.Status{},
			reasons:    map[string]error{},
			oldStatuses: []v1.ContainerStatus{{
				Name:                 "with-old-record",
				LastTerminationState: v1.ContainerState{Terminated: &v1.ContainerStateTerminated{}},
			}},
			expectedState: map[string]v1.ContainerState{
				"without-old-record": {Waiting: &v1.ContainerStateWaiting{
					Reason: "ContainerCreating",
				}},
				"with-old-record": {Waiting: &v1.ContainerStateWaiting{
					Reason: "ContainerCreating",
				}},
			},
			expectedInitState: map[string]v1.ContainerState{
				"without-old-record": {Waiting: &v1.ContainerStateWaiting{
					Reason: "PodInitializing",
				}},
				"with-old-record": {Waiting: &v1.ContainerStateWaiting{
					Reason: "PodInitializing",
				}},
			},
			expectedLastTerminationState: map[string]v1.ContainerState{
				"with-old-record": {Terminated: &v1.ContainerStateTerminated{}},
			},
		},
		// For running container, State should be Running, LastTerminationState should be retrieved from latest terminated status.
		{
			containers: []v1.Container{{Name: "running"}},
			statuses: []*kubecontainer.Status{
				{
					Name:      "running",
					State:     kubecontainer.ContainerStateRunning,
					StartedAt: testTimestamp,
				},
				{
					Name:     "running",
					State:    kubecontainer.ContainerStateExited,
					ExitCode: 1,
				},
			},
			reasons:     map[string]error{},
			oldStatuses: []v1.ContainerStatus{},
			expectedState: map[string]v1.ContainerState{
				"running": {Running: &v1.ContainerStateRunning{
					StartedAt: metav1.NewTime(testTimestamp),
				}},
			},
			expectedLastTerminationState: map[string]v1.ContainerState{
				"running": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    1,
					ContainerID: emptyContainerID,
				}},
			},
		},
		// For terminated container:
		// * If there is no recent start error record, State should be Terminated, LastTerminationState should be retrieved from
		// second latest terminated status;
		// * If there is recent start error record, State should be Waiting, LastTerminationState should be retrieved from latest
		// terminated status;
		// * If ExitCode = 0, restart policy is RestartPolicyOnFailure, the container shouldn't be restarted. No matter there is
		// recent start error or not, State should be Terminated, LastTerminationState should be retrieved from second latest
		// terminated status.
		{
			containers: []v1.Container{{Name: "without-reason"}, {Name: "with-reason"}},
			statuses: []*kubecontainer.Status{
				{
					Name:     "without-reason",
					State:    kubecontainer.ContainerStateExited,
					ExitCode: 1,
				},
				{
					Name:     "with-reason",
					State:    kubecontainer.ContainerStateExited,
					ExitCode: 2,
				},
				{
					Name:     "without-reason",
					State:    kubecontainer.ContainerStateExited,
					ExitCode: 3,
				},
				{
					Name:     "with-reason",
					State:    kubecontainer.ContainerStateExited,
					ExitCode: 4,
				},
				{
					Name:     "succeed",
					State:    kubecontainer.ContainerStateExited,
					ExitCode: 0,
				},
				{
					Name:     "succeed",
					State:    kubecontainer.ContainerStateExited,
					ExitCode: 5,
				},
			},
			reasons:     map[string]error{"with-reason": testErrorReason, "succeed": testErrorReason},
			oldStatuses: []v1.ContainerStatus{},
			expectedState: map[string]v1.ContainerState{
				"without-reason": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    1,
					ContainerID: emptyContainerID,
				}},
				"with-reason": {Waiting: &v1.ContainerStateWaiting{Reason: testErrorReason.Error()}},
				"succeed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    0,
					ContainerID: emptyContainerID,
				}},
			},
			expectedLastTerminationState: map[string]v1.ContainerState{
				"without-reason": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    3,
					ContainerID: emptyContainerID,
				}},
				"with-reason": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    2,
					ContainerID: emptyContainerID,
				}},
				"succeed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    5,
					ContainerID: emptyContainerID,
				}},
			},
		},
		// For Unknown Container Status:
		// * In certain situations a container can be running and fail to retrieve the status which results in
		// * a transition to the Unknown state. Prior to this fix, a container would make an invalid transition
		// * from Running->Waiting. This test validates the correct behavior of transitioning from Running->Terminated.
		{
			containers: []v1.Container{{Name: "unknown"}},
			statuses: []*kubecontainer.Status{
				{
					Name:  "unknown",
					State: kubecontainer.ContainerStateUnknown,
				},
				{
					Name:  "unknown",
					State: kubecontainer.ContainerStateRunning,
				},
			},
			reasons: map[string]error{},
			oldStatuses: []v1.ContainerStatus{{
				Name:  "unknown",
				State: v1.ContainerState{Running: &v1.ContainerStateRunning{}},
			}},
			expectedState: map[string]v1.ContainerState{
				"unknown": {Terminated: &v1.ContainerStateTerminated{
					ExitCode: 137,
					Message:  "The container could not be located when the pod was terminated",
					Reason:   "ContainerStatusUnknown",
				}},
			},
			expectedLastTerminationState: map[string]v1.ContainerState{
				"unknown": {Running: &v1.ContainerStateRunning{}},
			},
		},
	}

	for i, test := range tests {
		*kl.GetReasonCache() = *kubelet.NewReasonCache()
		for n, e := range test.reasons {
			(*kl.GetReasonCache()).Add(pod.UID, n, e, "")
		}
		pod.Spec.Containers = test.containers
		pod.Status.ContainerStatuses = test.oldStatuses
		podStatus.ContainerStatuses = test.statuses
		apiStatus := kl.GenerateAPIPodStatus(pod, podStatus, false)
		verifyContainerStatuses(t, apiStatus.ContainerStatuses, test.expectedState, test.expectedLastTerminationState, fmt.Sprintf("case %d", i))
	}

	// Everything should be the same for init containers
	for i, test := range tests {
		(*kl.GetReasonCache()) = *kubelet.NewReasonCache()
		for n, e := range test.reasons {
			(*kl.GetReasonCache()).Add(pod.UID, n, e, "")
		}
		pod.Spec.InitContainers = test.containers
		pod.Status.InitContainerStatuses = test.oldStatuses
		podStatus.ContainerStatuses = test.statuses
		apiStatus := kl.GenerateAPIPodStatus(pod, podStatus, false)
		expectedState := test.expectedState
		if test.expectedInitState != nil {
			expectedState = test.expectedInitState
		}
		verifyContainerStatuses(t, apiStatus.InitContainerStatuses, expectedState, test.expectedLastTerminationState, fmt.Sprintf("case %d", i))
	}
}

// Test generateAPIPodStatus with different restart policies.
func TestGenerateAPIPodStatusWithDifferentRestartPolicies(t *testing.T) {
	testErrorReason := fmt.Errorf("test-error")
	emptyContainerID := (&kubecontainer.ContainerID{}).String()
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kubelet := testKubelet.kubelet
	pod := podWithUIDNameNs("12345678", "foo", "new")
	containers := []v1.Container{{Name: "succeed"}, {Name: "failed"}}
	podStatus := &kubecontainer.PodStatus{
		ID:        pod.UID,
		Name:      pod.Name,
		Namespace: pod.Namespace,
		ContainerStatuses: []*kubecontainer.Status{
			{
				Name:     "succeed",
				State:    kubecontainer.ContainerStateExited,
				ExitCode: 0,
			},
			{
				Name:     "failed",
				State:    kubecontainer.ContainerStateExited,
				ExitCode: 1,
			},
			{
				Name:     "succeed",
				State:    kubecontainer.ContainerStateExited,
				ExitCode: 2,
			},
			{
				Name:     "failed",
				State:    kubecontainer.ContainerStateExited,
				ExitCode: 3,
			},
		},
	}
	(*kubelet.GetReasonCache()).Add(pod.UID, "succeed", testErrorReason, "")
	(*kubelet.GetReasonCache()).Add(pod.UID, "failed", testErrorReason, "")
	for c, test := range []struct {
		restartPolicy                v1.RestartPolicy
		expectedState                map[string]v1.ContainerState
		expectedLastTerminationState map[string]v1.ContainerState
		// Only set expectedInitState when it is different from expectedState
		expectedInitState map[string]v1.ContainerState
		// Only set expectedInitLastTerminationState when it is different from expectedLastTerminationState
		expectedInitLastTerminationState map[string]v1.ContainerState
	}{
		{
			restartPolicy: v1.RestartPolicyNever,
			expectedState: map[string]v1.ContainerState{
				"succeed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    0,
					ContainerID: emptyContainerID,
				}},
				"failed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    1,
					ContainerID: emptyContainerID,
				}},
			},
			expectedLastTerminationState: map[string]v1.ContainerState{
				"succeed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    2,
					ContainerID: emptyContainerID,
				}},
				"failed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    3,
					ContainerID: emptyContainerID,
				}},
			},
		},
		{
			restartPolicy: v1.RestartPolicyOnFailure,
			expectedState: map[string]v1.ContainerState{
				"succeed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    0,
					ContainerID: emptyContainerID,
				}},
				"failed": {Waiting: &v1.ContainerStateWaiting{Reason: testErrorReason.Error()}},
			},
			expectedLastTerminationState: map[string]v1.ContainerState{
				"succeed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    2,
					ContainerID: emptyContainerID,
				}},
				"failed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    1,
					ContainerID: emptyContainerID,
				}},
			},
		},
		{
			restartPolicy: v1.RestartPolicyAlways,
			expectedState: map[string]v1.ContainerState{
				"succeed": {Waiting: &v1.ContainerStateWaiting{Reason: testErrorReason.Error()}},
				"failed":  {Waiting: &v1.ContainerStateWaiting{Reason: testErrorReason.Error()}},
			},
			expectedLastTerminationState: map[string]v1.ContainerState{
				"succeed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    0,
					ContainerID: emptyContainerID,
				}},
				"failed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    1,
					ContainerID: emptyContainerID,
				}},
			},
			// If the init container is terminated with exit code 0, it won't be restarted even when the
			// restart policy is RestartAlways.
			expectedInitState: map[string]v1.ContainerState{
				"succeed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    0,
					ContainerID: emptyContainerID,
				}},
				"failed": {Waiting: &v1.ContainerStateWaiting{Reason: testErrorReason.Error()}},
			},
			expectedInitLastTerminationState: map[string]v1.ContainerState{
				"succeed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    2,
					ContainerID: emptyContainerID,
				}},
				"failed": {Terminated: &v1.ContainerStateTerminated{
					ExitCode:    1,
					ContainerID: emptyContainerID,
				}},
			},
		},
	} {
		pod.Spec.RestartPolicy = test.restartPolicy
		// Test normal containers
		pod.Spec.Containers = containers
		apiStatus := kubelet.GenerateAPIPodStatus(pod, podStatus, false)
		expectedState, expectedLastTerminationState := test.expectedState, test.expectedLastTerminationState
		verifyContainerStatuses(t, apiStatus.ContainerStatuses, expectedState, expectedLastTerminationState, fmt.Sprintf("case %d", c))
		pod.Spec.Containers = nil

		// Test init containers
		pod.Spec.InitContainers = containers
		apiStatus = kubelet.GenerateAPIPodStatus(pod, podStatus, false)
		if test.expectedInitState != nil {
			expectedState = test.expectedInitState
		}
		if test.expectedInitLastTerminationState != nil {
			expectedLastTerminationState = test.expectedInitLastTerminationState
		}
		verifyContainerStatuses(t, apiStatus.InitContainerStatuses, expectedState, expectedLastTerminationState, fmt.Sprintf("case %d", c))
		pod.Spec.InitContainers = nil
	}
}

// testPodAdmitHandler is a lifecycle.PodAdmitHandler for testing.
type testPodAdmitHandler struct {
	// list of pods to reject.
	podsToReject []*v1.Pod
}

// Admit rejects all pods in the podsToReject list with a matching UID.
func (a *testPodAdmitHandler) Admit(attrs *lifecycle.PodAdmitAttributes) lifecycle.PodAdmitResult {
	for _, podToReject := range a.podsToReject {
		if podToReject.UID == attrs.Pod.UID {
			return lifecycle.PodAdmitResult{Admit: false, Reason: "Rejected", Message: "Pod is rejected"}
		}
	}
	return lifecycle.PodAdmitResult{Admit: true}
}

// Test verifies that the kubelet invokes an admission handler during HandlePodAdditions.
func TestHandlePodAdditionsInvokesPodAdmitHandlers(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kl := testKubelet.kubelet
	*kl.GetNodeLister() = testNodeLister{nodes: []*v1.Node{
		{
			ObjectMeta: metav1.ObjectMeta{Name: string(*kl.GetNodeName())},
			Status: v1.NodeStatus{
				Allocatable: v1.ResourceList{
					v1.ResourcePods: *resource.NewQuantity(110, resource.DecimalSI),
				},
			},
		},
	}}

	pods := []*v1.Pod{
		{
			ObjectMeta: metav1.ObjectMeta{
				UID:       "123456789",
				Name:      "podA",
				Namespace: "foo",
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				UID:       "987654321",
				Name:      "podB",
				Namespace: "foo",
			},
		},
	}
	podToReject := pods[0]
	podToAdmit := pods[1]
	podsToReject := []*v1.Pod{podToReject}

	(*kl.AdmitHandlers()).AddPodAdmitHandler(&testPodAdmitHandler{podsToReject: podsToReject})

	kl.HandlePodAdditions(pods)

	// Check pod status stored in the status map.
	checkPodStatus(t, kl, podToReject, v1.PodFailed)
	checkPodStatus(t, kl, podToAdmit, v1.PodPending)
}

func TestPodResourceAllocationReset(t *testing.T) {
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.InPlacePodVerticalScaling, true)()
	testKubelet := newTestKubelet(t, false)
	defer testKubelet.Cleanup()
	kubelet := testKubelet.kubelet
	(*kubelet.GetStatusManager()) = status.NewFakeManager()

	nodes := []*v1.Node{
		{
			ObjectMeta: metav1.ObjectMeta{Name: testKubeletHostname},
			Status: v1.NodeStatus{
				Capacity: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("8"),
					v1.ResourceMemory: resource.MustParse("8Gi"),
				},
				Allocatable: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("4"),
					v1.ResourceMemory: resource.MustParse("4Gi"),
					v1.ResourcePods:   *resource.NewQuantity(40, resource.DecimalSI),
				},
			},
		},
	}
	(*kubelet.GetNodeLister()) = testNodeLister{nodes: nodes}

	cpu500m := resource.MustParse("500m")
	cpu800m := resource.MustParse("800m")
	mem500M := resource.MustParse("500Mi")
	mem800M := resource.MustParse("800Mi")
	cpu500mMem500MPodSpec := &v1.PodSpec{
		Containers: []v1.Container{
			{
				Name: "c1",
				Resources: v1.ResourceRequirements{
					Requests: v1.ResourceList{v1.ResourceCPU: cpu500m, v1.ResourceMemory: mem500M},
				},
			},
		},
	}
	cpu800mMem800MPodSpec := cpu500mMem500MPodSpec.DeepCopy()
	cpu800mMem800MPodSpec.Containers[0].Resources.Requests = v1.ResourceList{v1.ResourceCPU: cpu800m, v1.ResourceMemory: mem800M}
	cpu800mPodSpec := cpu500mMem500MPodSpec.DeepCopy()
	cpu800mPodSpec.Containers[0].Resources.Requests = v1.ResourceList{v1.ResourceCPU: cpu800m}
	mem800MPodSpec := cpu500mMem500MPodSpec.DeepCopy()
	mem800MPodSpec.Containers[0].Resources.Requests = v1.ResourceList{v1.ResourceMemory: mem800M}

	cpu500mPodSpec := cpu500mMem500MPodSpec.DeepCopy()
	cpu500mPodSpec.Containers[0].Resources.Requests = v1.ResourceList{v1.ResourceCPU: cpu500m}
	mem500MPodSpec := cpu500mMem500MPodSpec.DeepCopy()
	mem500MPodSpec.Containers[0].Resources.Requests = v1.ResourceList{v1.ResourceMemory: mem500M}
	emptyPodSpec := cpu500mMem500MPodSpec.DeepCopy()
	emptyPodSpec.Containers[0].Resources.Requests = v1.ResourceList{}

	tests := []struct {
		name                          string
		pod                           *v1.Pod
		existingPodAllocation         *v1.Pod
		expectedPodResourceAllocation state.PodResourceAllocation
	}{
		{
			name: "Having both memory and cpu, resource allocation not exists",
			pod:  podWithUIDNameNsSpec("1", "pod1", "foo", *cpu500mMem500MPodSpec),
			expectedPodResourceAllocation: state.PodResourceAllocation{
				"1": map[string]v1.ResourceList{
					cpu500mMem500MPodSpec.Containers[0].Name: cpu500mMem500MPodSpec.Containers[0].Resources.Requests,
				},
			},
		},
		{
			name:                  "Having both memory and cpu, resource allocation exists",
			pod:                   podWithUIDNameNsSpec("2", "pod2", "foo", *cpu500mMem500MPodSpec),
			existingPodAllocation: podWithUIDNameNsSpec("2", "pod2", "foo", *cpu500mMem500MPodSpec),
			expectedPodResourceAllocation: state.PodResourceAllocation{
				"2": map[string]v1.ResourceList{
					cpu500mMem500MPodSpec.Containers[0].Name: cpu500mMem500MPodSpec.Containers[0].Resources.Requests,
				},
			},
		},
		{
			name:                  "Having both memory and cpu, resource allocation exists (with different value)",
			pod:                   podWithUIDNameNsSpec("3", "pod3", "foo", *cpu500mMem500MPodSpec),
			existingPodAllocation: podWithUIDNameNsSpec("3", "pod3", "foo", *cpu800mMem800MPodSpec),
			expectedPodResourceAllocation: state.PodResourceAllocation{
				"3": map[string]v1.ResourceList{
					cpu800mMem800MPodSpec.Containers[0].Name: cpu800mMem800MPodSpec.Containers[0].Resources.Requests,
				},
			},
		},
		{
			name: "Only has cpu, resource allocation not exists",
			pod:  podWithUIDNameNsSpec("4", "pod5", "foo", *cpu500mPodSpec),
			expectedPodResourceAllocation: state.PodResourceAllocation{
				"4": map[string]v1.ResourceList{
					cpu500mPodSpec.Containers[0].Name: cpu500mPodSpec.Containers[0].Resources.Requests,
				},
			},
		},
		{
			name:                  "Only has cpu, resource allocation exists",
			pod:                   podWithUIDNameNsSpec("5", "pod5", "foo", *cpu500mPodSpec),
			existingPodAllocation: podWithUIDNameNsSpec("5", "pod5", "foo", *cpu500mPodSpec),
			expectedPodResourceAllocation: state.PodResourceAllocation{
				"5": map[string]v1.ResourceList{
					cpu500mPodSpec.Containers[0].Name: cpu500mPodSpec.Containers[0].Resources.Requests,
				},
			},
		},
		{
			name:                  "Only has cpu, resource allocation exists (with different value)",
			pod:                   podWithUIDNameNsSpec("6", "pod6", "foo", *cpu500mPodSpec),
			existingPodAllocation: podWithUIDNameNsSpec("6", "pod6", "foo", *cpu800mPodSpec),
			expectedPodResourceAllocation: state.PodResourceAllocation{
				"6": map[string]v1.ResourceList{
					cpu800mPodSpec.Containers[0].Name: cpu800mPodSpec.Containers[0].Resources.Requests,
				},
			},
		},
		{
			name: "Only has memory, resource allocation not exists",
			pod:  podWithUIDNameNsSpec("7", "pod7", "foo", *mem500MPodSpec),
			expectedPodResourceAllocation: state.PodResourceAllocation{
				"7": map[string]v1.ResourceList{
					mem500MPodSpec.Containers[0].Name: mem500MPodSpec.Containers[0].Resources.Requests,
				},
			},
		},
		{
			name:                  "Only has memory, resource allocation exists",
			pod:                   podWithUIDNameNsSpec("8", "pod8", "foo", *mem500MPodSpec),
			existingPodAllocation: podWithUIDNameNsSpec("8", "pod8", "foo", *mem500MPodSpec),
			expectedPodResourceAllocation: state.PodResourceAllocation{
				"8": map[string]v1.ResourceList{
					mem500MPodSpec.Containers[0].Name: mem500MPodSpec.Containers[0].Resources.Requests,
				},
			},
		},
		{
			name:                  "Only has memory, resource allocation exists (with different value)",
			pod:                   podWithUIDNameNsSpec("9", "pod9", "foo", *mem500MPodSpec),
			existingPodAllocation: podWithUIDNameNsSpec("9", "pod9", "foo", *mem800MPodSpec),
			expectedPodResourceAllocation: state.PodResourceAllocation{
				"9": map[string]v1.ResourceList{
					mem800MPodSpec.Containers[0].Name: mem800MPodSpec.Containers[0].Resources.Requests,
				},
			},
		},
		{
			name: "No CPU and memory, resource allocation not exists",
			pod:  podWithUIDNameNsSpec("10", "pod10", "foo", *emptyPodSpec),
			expectedPodResourceAllocation: state.PodResourceAllocation{
				"10": map[string]v1.ResourceList{
					emptyPodSpec.Containers[0].Name: emptyPodSpec.Containers[0].Resources.Requests,
				},
			},
		},
		{
			name:                  "No CPU and memory, resource allocation exists",
			pod:                   podWithUIDNameNsSpec("11", "pod11", "foo", *emptyPodSpec),
			existingPodAllocation: podWithUIDNameNsSpec("11", "pod11", "foo", *emptyPodSpec),
			expectedPodResourceAllocation: state.PodResourceAllocation{
				"11": map[string]v1.ResourceList{
					emptyPodSpec.Containers[0].Name: emptyPodSpec.Containers[0].Resources.Requests,
				},
			},
		},
	}
	for _, tc := range tests {
		if tc.existingPodAllocation != nil {
			// when kubelet restarts, AllocatedResources has already existed before adding pod
			err := (*kubelet.GetStatusManager()).SetPodAllocation(tc.existingPodAllocation)
			if err != nil {
				t.Fatalf("failed to set pod allocation: %v", err)
			}
		}
		kubelet.HandlePodAdditions([]*v1.Pod{tc.pod})

		allocatedResources, found := (*kubelet.GetStatusManager()).GetContainerResourceAllocation(string(tc.pod.UID), tc.pod.Spec.Containers[0].Name)
		if !found {
			t.Fatalf("resource allocation should exist: (pod: %#v, container: %s)", tc.pod, tc.pod.Spec.Containers[0].Name)
		}
		assert.Equal(t, tc.expectedPodResourceAllocation[string(tc.pod.UID)][tc.pod.Spec.Containers[0].Name], allocatedResources, tc.name)
	}
}

func TestHandlePodResourcesResize(t *testing.T) {
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.InPlacePodVerticalScaling, true)()
	testKubelet := newTestKubelet(t, false)
	defer testKubelet.Cleanup()
	kubelet := testKubelet.kubelet
	(*kubelet.GetStatusManager()) = status.NewFakeManager()

	cpu500m := resource.MustParse("500m")
	cpu1000m := resource.MustParse("1")
	cpu1500m := resource.MustParse("1500m")
	cpu2500m := resource.MustParse("2500m")
	cpu5000m := resource.MustParse("5000m")
	mem500M := resource.MustParse("500Mi")
	mem1000M := resource.MustParse("1Gi")
	mem1500M := resource.MustParse("1500Mi")
	mem2500M := resource.MustParse("2500Mi")
	mem4500M := resource.MustParse("4500Mi")

	nodes := []*v1.Node{
		{
			ObjectMeta: metav1.ObjectMeta{Name: testKubeletHostname},
			Status: v1.NodeStatus{
				Capacity: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("8"),
					v1.ResourceMemory: resource.MustParse("8Gi"),
				},
				Allocatable: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("4"),
					v1.ResourceMemory: resource.MustParse("4Gi"),
					v1.ResourcePods:   *resource.NewQuantity(40, resource.DecimalSI),
				},
			},
		},
	}
	(*kubelet.GetNodeLister()) = testNodeLister{nodes: nodes}

	testPod1 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			UID:       "1111",
			Name:      "pod1",
			Namespace: "ns1",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "c1",
					Image: "i1",
					Resources: v1.ResourceRequirements{
						Requests: v1.ResourceList{v1.ResourceCPU: cpu1000m, v1.ResourceMemory: mem1000M},
					},
				},
			},
		},
		Status: v1.PodStatus{
			Phase: v1.PodRunning,
			ContainerStatuses: []v1.ContainerStatus{
				{
					Name:               "c1",
					AllocatedResources: v1.ResourceList{v1.ResourceCPU: cpu1000m, v1.ResourceMemory: mem1000M},
					Resources:          &v1.ResourceRequirements{},
				},
			},
		},
	}
	testPod2 := testPod1.DeepCopy()
	testPod2.UID = "2222"
	testPod2.Name = "pod2"
	testPod2.Namespace = "ns2"
	testPod3 := testPod1.DeepCopy()
	testPod3.UID = "3333"
	testPod3.Name = "pod3"
	testPod3.Namespace = "ns2"

	testKubelet.fakeKubeClient = fake.NewSimpleClientset(testPod1, testPod2, testPod3)
	(*kubelet.GetKubeClient()) = testKubelet.fakeKubeClient
	defer testKubelet.fakeKubeClient.ClearActions()
	(*kubelet.GetPodManager()).AddPod(testPod1)
	(*kubelet.GetPodManager()).AddPod(testPod2)
	(*kubelet.GetPodManager()).AddPod(testPod3)
	(*kubelet.GetPodWorkers()).(*fakePodWorkers).running = map[types.UID]bool{
		testPod1.UID: true,
		testPod2.UID: true,
		testPod3.UID: true,
	}
	defer (*kubelet.GetPodManager()).RemovePod(testPod3)
	defer (*kubelet.GetPodManager()).RemovePod(testPod2)
	defer (*kubelet.GetPodManager()).RemovePod(testPod1)

	tests := []struct {
		name                string
		pod                 *v1.Pod
		newRequests         v1.ResourceList
		expectedAllocations v1.ResourceList
		expectedResize      v1.PodResizeStatus
	}{
		{
			name:                "Request CPU and memory decrease - expect InProgress",
			pod:                 testPod2,
			newRequests:         v1.ResourceList{v1.ResourceCPU: cpu500m, v1.ResourceMemory: mem500M},
			expectedAllocations: v1.ResourceList{v1.ResourceCPU: cpu500m, v1.ResourceMemory: mem500M},
			expectedResize:      v1.PodResizeStatusInProgress,
		},
		{
			name:                "Request CPU increase, memory decrease - expect InProgress",
			pod:                 testPod2,
			newRequests:         v1.ResourceList{v1.ResourceCPU: cpu1500m, v1.ResourceMemory: mem500M},
			expectedAllocations: v1.ResourceList{v1.ResourceCPU: cpu1500m, v1.ResourceMemory: mem500M},
			expectedResize:      v1.PodResizeStatusInProgress,
		},
		{
			name:                "Request CPU decrease, memory increase - expect InProgress",
			pod:                 testPod2,
			newRequests:         v1.ResourceList{v1.ResourceCPU: cpu500m, v1.ResourceMemory: mem1500M},
			expectedAllocations: v1.ResourceList{v1.ResourceCPU: cpu500m, v1.ResourceMemory: mem1500M},
			expectedResize:      v1.PodResizeStatusInProgress,
		},
		{
			name:                "Request CPU and memory increase beyond current capacity - expect Deferred",
			pod:                 testPod2,
			newRequests:         v1.ResourceList{v1.ResourceCPU: cpu2500m, v1.ResourceMemory: mem2500M},
			expectedAllocations: v1.ResourceList{v1.ResourceCPU: cpu1000m, v1.ResourceMemory: mem1000M},
			expectedResize:      v1.PodResizeStatusDeferred,
		},
		{
			name:                "Request CPU decrease and memory increase beyond current capacity - expect Deferred",
			pod:                 testPod2,
			newRequests:         v1.ResourceList{v1.ResourceCPU: cpu500m, v1.ResourceMemory: mem2500M},
			expectedAllocations: v1.ResourceList{v1.ResourceCPU: cpu1000m, v1.ResourceMemory: mem1000M},
			expectedResize:      v1.PodResizeStatusDeferred,
		},
		{
			name:                "Request memory increase beyond node capacity - expect Infeasible",
			pod:                 testPod2,
			newRequests:         v1.ResourceList{v1.ResourceCPU: cpu1000m, v1.ResourceMemory: mem4500M},
			expectedAllocations: v1.ResourceList{v1.ResourceCPU: cpu1000m, v1.ResourceMemory: mem1000M},
			expectedResize:      v1.PodResizeStatusInfeasible,
		},
		{
			name:                "Request CPU increase beyond node capacity - expect Infeasible",
			pod:                 testPod2,
			newRequests:         v1.ResourceList{v1.ResourceCPU: cpu5000m, v1.ResourceMemory: mem1000M},
			expectedAllocations: v1.ResourceList{v1.ResourceCPU: cpu1000m, v1.ResourceMemory: mem1000M},
			expectedResize:      v1.PodResizeStatusInfeasible,
		},
	}

	for _, tt := range tests {
		tt.pod.Spec.Containers[0].Resources.Requests = tt.newRequests
		tt.pod.Status.ContainerStatuses[0].AllocatedResources = v1.ResourceList{v1.ResourceCPU: cpu1000m, v1.ResourceMemory: mem1000M}
		kubelet.HandlePodResourcesResize(tt.pod)
		updatedPod, found := (*kubelet.GetPodManager()).GetPodByName(tt.pod.Namespace, tt.pod.Name)
		assert.True(t, found, "expected to find pod %s", tt.pod.Name)
		assert.Equal(t, tt.expectedAllocations, updatedPod.Status.ContainerStatuses[0].AllocatedResources, tt.name)
		assert.Equal(t, tt.expectedResize, updatedPod.Status.Resize, tt.name)
		testKubelet.fakeKubeClient.ClearActions()
	}
}

// testPodSyncLoopHandler is a lifecycle.PodSyncLoopHandler that is used for testing.
type testPodSyncLoopHandler struct {
	// list of pods to sync
	podsToSync []*v1.Pod
}

// ShouldSync evaluates if the pod should be synced from the kubelet.
func (a *testPodSyncLoopHandler) ShouldSync(pod *v1.Pod) bool {
	for _, podToSync := range a.podsToSync {
		if podToSync.UID == pod.UID {
			return true
		}
	}
	return false
}

// TestGetPodsToSyncInvokesPodSyncLoopHandlers ensures that the get pods to sync routine invokes the handler.
func TestGetPodsToSyncInvokesPodSyncLoopHandlers(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kubelet := testKubelet.kubelet
	pods := newTestPods(5)
	expected := []*v1.Pod{pods[0]}
	kubelet.AddPodSyncLoopHandler(&testPodSyncLoopHandler{expected})
	(*kubelet.GetPodManager()).SetPods(pods)

	podsToSync := kubelet.GetPodsToSync()
	sort.Sort(podsByUID(expected))
	sort.Sort(podsByUID(podsToSync))
	assert.Equal(t, expected, podsToSync)
}

// testPodSyncHandler is a lifecycle.PodSyncHandler that is used for testing.
type testPodSyncHandler struct {
	// list of pods to evict.
	podsToEvict []*v1.Pod
	// the reason for the eviction
	reason string
	// the message for the eviction
	message string
}

// ShouldEvict evaluates if the pod should be evicted from the kubelet.
func (a *testPodSyncHandler) ShouldEvict(pod *v1.Pod) lifecycle.ShouldEvictResponse {
	for _, podToEvict := range a.podsToEvict {
		if podToEvict.UID == pod.UID {
			return lifecycle.ShouldEvictResponse{Evict: true, Reason: a.reason, Message: a.message}
		}
	}
	return lifecycle.ShouldEvictResponse{Evict: false}
}

// TestGenerateAPIPodStatusInvokesPodSyncHandlers invokes the handlers and reports the proper status
func TestGenerateAPIPodStatusInvokesPodSyncHandlers(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kubelet := testKubelet.kubelet
	pod := newTestPods(1)[0]
	podsToEvict := []*v1.Pod{pod}
	kubelet.AddPodSyncHandler(&testPodSyncHandler{podsToEvict, "Evicted", "because"})
	status := &kubecontainer.PodStatus{
		ID:        pod.UID,
		Name:      pod.Name,
		Namespace: pod.Namespace,
	}
	apiStatus := kubelet.GenerateAPIPodStatus(pod, status, false)
	require.Equal(t, v1.PodFailed, apiStatus.Phase)
	require.Equal(t, "Evicted", apiStatus.Reason)
	require.Equal(t, "because", apiStatus.Message)
}

func TestSyncTerminatingPodKillPod(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	defer testKubelet.Cleanup()
	kl := testKubelet.kubelet
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			UID:       "12345678",
			Name:      "bar",
			Namespace: "foo",
		},
	}
	pods := []*v1.Pod{pod}
	(*kl.GetPodManager()).SetPods(pods)
	podStatus := &kubecontainer.PodStatus{ID: pod.UID}
	gracePeriodOverride := int64(0)
	err := kl.SyncTerminatingPod(context.Background(), pod, podStatus, &gracePeriodOverride, func(podStatus *v1.PodStatus) {
		podStatus.Phase = v1.PodFailed
		podStatus.Reason = "reason"
		podStatus.Message = "message"
	})
	require.NoError(t, err)

	// Check pod status stored in the status map.
	checkPodStatus(t, kl, pod, v1.PodFailed)
}

// func TestSyncLabels(t *testing.T) {
// 	tests := []struct {
// 		name             string
// 		existingNode     *v1.Node
// 		isPatchingNeeded bool
// 	}{
// 		{
// 			name:             "no labels",
// 			existingNode:     &v1.Node{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{}}},
// 			isPatchingNeeded: true,
// 		},
// 		{
// 			name:             "wrong labels",
// 			existingNode:     &v1.Node{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{v1.LabelOSStable: "dummyOS", v1.LabelArchStable: "dummyArch"}}},
// 			isPatchingNeeded: true,
// 		},
// 		{
// 			name:             "correct labels",
// 			existingNode:     &v1.Node{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{v1.LabelOSStable: goruntime.GOOS, v1.LabelArchStable: goruntime.GOARCH}}},
// 			isPatchingNeeded: false,
// 		},
// 		{
// 			name:             "partially correct labels",
// 			existingNode:     &v1.Node{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{v1.LabelOSStable: goruntime.GOOS, v1.LabelArchStable: "dummyArch"}}},
// 			isPatchingNeeded: true,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			testKubelet := newTestKubelet(t, false)
// 			defer testKubelet.Cleanup()
// 			kl := testKubelet.kubelet
// 			kubeClient := testKubelet.fakeKubeClient

// 			test.existingNode.Name = string(*kl.GetNodeName())

// 			*kl.GetNodeLister() = testNodeLister{nodes: []*v1.Node{test.existingNode}}
// 			go func() { kl.SyncNodeStatus() }()

// 			err := retryWithExponentialBackOff(
// 				100*time.Millisecond,
// 				func() (bool, error) {
// 					var savedNode *v1.Node
// 					if test.isPatchingNeeded {
// 						actions := kubeClient.Actions()
// 						if len(actions) == 0 {
// 							t.Logf("No action yet")
// 							return false, nil
// 						}
// 						for _, action := range actions {
// 							if action.GetVerb() == "patch" {
// 								var (
// 									err          error
// 									patchAction  = action.(core.PatchActionImpl)
// 									patchContent = patchAction.GetPatch()
// 								)
// 								savedNode, err = kubelet.ApplyNodeStatusPatch(test.existingNode, patchContent)
// 								if err != nil {
// 									t.Logf("node patching failed, %v", err)
// 									return false, nil
// 								}
// 							}
// 						}
// 					} else {
// 						savedNode = test.existingNode
// 					}
// 					if savedNode == nil || savedNode.Labels == nil {
// 						t.Logf("savedNode.Labels should not be nil")
// 						return false, nil
// 					}
// 					val, ok := savedNode.Labels[v1.LabelOSStable]
// 					if !ok {
// 						t.Logf("expected kubernetes.io/os label to be present")
// 						return false, nil
// 					}
// 					if val != goruntime.GOOS {
// 						t.Logf("expected kubernetes.io/os to match runtime.GOOS but got %v", val)
// 						return false, nil
// 					}
// 					val, ok = savedNode.Labels[v1.LabelArchStable]
// 					if !ok {
// 						t.Logf("expected kubernetes.io/arch label to be present")
// 						return false, nil
// 					}
// 					if val != goruntime.GOARCH {
// 						t.Logf("expected kubernetes.io/arch to match runtime.GOARCH but got %v", val)
// 						return false, nil
// 					}
// 					return true, nil
// 				},
// 			)
// 			if err != nil {
// 				t.Fatalf("expected labels to be reconciled but it failed with %v", err)
// 			}
// 		})
// 	}
// }

func waitForVolumeUnmount(
	volumeManager kubeletvolume.VolumeManager,
	pod *v1.Pod) error {
	var podVolumes kubecontainer.VolumeMap
	err := retryWithExponentialBackOff(
		time.Duration(50*time.Millisecond),
		func() (bool, error) {
			// Verify volumes detached
			podVolumes = volumeManager.GetMountedVolumesForPod(
				util.GetUniquePodName(pod))

			if len(podVolumes) != 0 {
				return false, nil
			}

			return true, nil
		},
	)

	if err != nil {
		return fmt.Errorf(
			"Expected volumes to be unmounted. But some volumes are still mounted: %#v", podVolumes)
	}

	return nil
}

func waitForVolumeDetach(
	volumeName v1.UniqueVolumeName,
	volumeManager kubeletvolume.VolumeManager) error {
	attachedVolumes := []v1.UniqueVolumeName{}
	err := retryWithExponentialBackOff(
		time.Duration(50*time.Millisecond),
		func() (bool, error) {
			// Verify volumes detached
			volumeAttached := volumeManager.VolumeIsAttached(volumeName)
			return !volumeAttached, nil
		},
	)

	if err != nil {
		return fmt.Errorf(
			"Expected volumes to be detached. But some volumes are still attached: %#v", attachedVolumes)
	}

	return nil
}

func retryWithExponentialBackOff(initialDuration time.Duration, fn wait.ConditionFunc) error {
	backoff := wait.Backoff{
		Duration: initialDuration,
		Factor:   3,
		Jitter:   0,
		Steps:    6,
	}
	return wait.ExponentialBackoff(backoff, fn)
}

func simulateVolumeInUseUpdate(
	volumeName v1.UniqueVolumeName,
	stopCh <-chan struct{},
	volumeManager kubeletvolume.VolumeManager) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			volumeManager.MarkVolumesAsReportedInUse(
				[]v1.UniqueVolumeName{volumeName})
		case <-stopCh:
			return
		}
	}
}

func runVolumeManager(kubelet *kubelet.Kubelet) chan struct{} {
	stopCh := make(chan struct{})
	go (*kubelet.GetVolumeManager()).Run(*kubelet.GetSourcesReady(), stopCh)
	return stopCh
}

// dirExists returns true if the path exists and represents a directory.
func dirExists(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Sort pods by UID.
type podsByUID []*v1.Pod

func (p podsByUID) Len() int           { return len(p) }
func (p podsByUID) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p podsByUID) Less(i, j int) bool { return p[i].UID < p[j].UID }

// createAndStartFakeRemoteRuntime creates and starts fakeremote.RemoteRuntime.
// It returns the RemoteRuntime, endpoint on success.
// Users should call fakeRuntime.Stop() to cleanup the server.
func createAndStartFakeRemoteRuntime(t *testing.T) (*fakeremote.RemoteRuntime, string) {
	endpoint, err := fakeremote.GenerateEndpoint()
	require.NoError(t, err)

	fakeRuntime := fakeremote.NewFakeRemoteRuntime()
	fakeRuntime.Start(endpoint)

	return fakeRuntime, endpoint
}

func createRemoteRuntimeService(endpoint string, t *testing.T, tp oteltrace.TracerProvider) internalapi.RuntimeService {
	runtimeService, err := remote.NewRemoteRuntimeService(endpoint, 15*time.Second, tp)
	require.NoError(t, err)
	return runtimeService
}

func TestNewMainKubeletStandAlone(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "logs")
	kubelet.ContainerLogsDir = tempDir
	assert.NoError(t, err)
	defer os.RemoveAll(kubelet.ContainerLogsDir)
	kubeCfg := &kubeletconfiginternal.KubeletConfiguration{
		SyncFrequency: metav1.Duration{Duration: time.Minute},
		ConfigMapAndSecretChangeDetectionStrategy: kubeletconfiginternal.WatchChangeDetectionStrategy,
		ContainerLogMaxSize:                       "10Mi",
		ContainerLogMaxFiles:                      5,
		MemoryThrottlingFactor:                    utilpointer.Float64(0),
	}
	var prober volume.DynamicPluginProber
	tp := oteltrace.NewNoopTracerProvider()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	cadvisor := cadvisortest.NewMockInterface(mockCtrl)
	cadvisor.EXPECT().MachineInfo().Return(&cadvisorapi.MachineInfo{}, nil).AnyTimes()
	cadvisor.EXPECT().ImagesFsInfo().Return(cadvisorapiv2.FsInfo{
		Usage:     400,
		Capacity:  1000,
		Available: 600,
	}, nil).AnyTimes()
	tlsOptions := &server.TLSOptions{
		Config: &tls.Config{
			MinVersion: 0,
		},
	}
	fakeRuntime, endpoint := createAndStartFakeRemoteRuntime(t)
	defer func() {
		fakeRuntime.Stop()
	}()
	fakeRecorder := &record.FakeRecorder{}
	rtSvc := createRemoteRuntimeService(endpoint, t, oteltrace.NewNoopTracerProvider())
	kubeDep := &kubelet.Dependencies{
		Auth:                 nil,
		CAdvisorInterface:    cadvisor,
		Cloud:                nil,
		ContainerManager:     cm.NewStubContainerManager(),
		KubeClient:           nil, // standalone mode
		HeartbeatClient:      nil,
		EventClient:          nil,
		TracerProvider:       tp,
		HostUtil:             hostutil.NewFakeHostUtil(nil),
		Mounter:              mount.NewFakeMounter(nil),
		Recorder:             fakeRecorder,
		RemoteRuntimeService: rtSvc,
		RemoteImageService:   fakeRuntime.ImageService,
		Subpather:            &subpath.FakeSubpath{},
		OOMAdjuster:          oom.NewOOMAdjuster(),
		OSInterface:          kubecontainer.RealOS{},
		DynamicPluginProber:  prober,
		TLSOptions:           tlsOptions,
	}
	crOptions := &config.ContainerRuntimeOptions{}

	testMainKubelet, err := kubelet.NewMainKubelet(
		kubeCfg,
		kubeDep,
		crOptions,
		"hostname",
		false,
		"hostname",
		[]net.IP{},
		"",
		"external",
		"/tmp/cert",
		"/tmp/rootdir",
		"",
		"",
		false,
		[]v1.Taint{},
		[]string{},
		"",
		false,
		false,
		metav1.Duration{Duration: time.Minute},
		1024,
		110,
		true,
		true,
		map[string]string{},
		1024,
		false,
	)
	assert.NoError(t, err, "NewMainKubelet should succeed")
	assert.NotNil(t, testMainKubelet, "testMainKubelet should not be nil")

	testMainKubelet.BirthCry()
	testMainKubelet.StartGarbageCollection()
	// Nil pointer panic can be reproduced if configmap manager is not nil.
	// See https://github.com/kubernetes/kubernetes/issues/113492
	// pod := &v1.Pod{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		UID:       "12345678",
	// 		Name:      "bar",
	// 		Namespace: "foo",
	// 	},
	// 	Spec: v1.PodSpec{
	// 		Containers: []v1.Container{{
	// 			EnvFrom: []v1.EnvFromSource{{
	// 				ConfigMapRef: &v1.ConfigMapEnvSource{
	// 					LocalObjectReference: v1.LocalObjectReference{Name: "config-map"}}},
	// 			}}},
	// 		Volumes: []v1.Volume{{
	// 			VolumeSource: v1.VolumeSource{
	// 				ConfigMap: &v1.ConfigMapVolumeSource{
	// 					LocalObjectReference: v1.LocalObjectReference{
	// 						Name: "config-map"}}}}},
	// 	},
	// }
	// testMainKubelet.configMapManager.RegisterPod(pod)
	// testMainKubelet.secretManager.RegisterPod(pod)
	assert.Nil(t, testMainKubelet.GetConfigMapManager(), "configmap manager should be nil if kubelet is in standalone mode")
	assert.Nil(t, testMainKubelet.GetSecretManager(), "secret manager should be nil if kubelet is in standalone mode")
}

func TestSyncPodSpans(t *testing.T) {
	testKubelet := newTestKubelet(t, false)
	kl := testKubelet.kubelet

	recorder := record.NewFakeRecorder(20)
	nodeRef := &v1.ObjectReference{
		Kind:      "Node",
		Name:      "testNode",
		UID:       types.UID("testNode"),
		Namespace: "",
	}
	*kl.GetDNSConfigurer() = *dns.NewConfigurer(recorder, nodeRef, nil, nil, "TEST", "")

	kubeCfg := &kubeletconfiginternal.KubeletConfiguration{
		SyncFrequency: metav1.Duration{Duration: time.Minute},
		ConfigMapAndSecretChangeDetectionStrategy: kubeletconfiginternal.WatchChangeDetectionStrategy,
		ContainerLogMaxSize:                       "10Mi",
		ContainerLogMaxFiles:                      5,
		MemoryThrottlingFactor:                    utilpointer.Float64(0),
	}

	exp := tracetest.NewInMemoryExporter()
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSyncer(exp),
	)
	*kl.GetTracer() = tp.Tracer("k8s.io/kubernetes/pkg/kubelet")

	fakeRuntime, endpoint := createAndStartFakeRemoteRuntime(t)
	defer func() {
		fakeRuntime.Stop()
	}()
	runtimeSvc := createRemoteRuntimeService(endpoint, t, tp)
	*kl.GetRuntimeService() = runtimeSvc

	fakeRuntime.ImageService.SetFakeImageSize(100)
	fakeRuntime.ImageService.SetFakeImages([]string{"test:latest"})
	imageSvc, err := remote.NewRemoteImageService(endpoint, 15*time.Second, tp)
	assert.NoError(t, err)

	*kl.GetContainerRuntime(), err = kuberuntime.NewKubeGenericRuntimeManager(
		*kl.GetRecorder(),
		*kl.GetLiveinessManager(),
		*kl.GetReadinessManager(),
		*kl.GetStartupManager(),
		*kl.GetRootDir(),
		kl.GetMachineInfo(),
		*kl.GetPodWorkers(),
		*kl.GetOS(),
		kl,
		nil,
		kl.GetBackOff(),
		kubeCfg.SerializeImagePulls,
		kubeCfg.MaxParallelImagePulls,
		float32(kubeCfg.RegistryPullQPS),
		int(kubeCfg.RegistryBurst),
		"",
		"",
		kubeCfg.CPUCFSQuota,
		kubeCfg.CPUCFSQuotaPeriod,
		runtimeSvc,
		imageSvc,
		*kl.GetContainerManager(),
		*kl.GetContainerLogManager(),
		kl.GetRuntimeClassManager(),
		false,
		kubeCfg.MemorySwap.SwapBehavior,
		(*kl.GetContainerManager()).GetNodeAllocatableAbsolute,
		*kubeCfg.MemoryThrottlingFactor,
		kubeletutil.NewPodStartupLatencyTracker(),
		tp,
	)
	assert.NoError(t, err)

	pod := podWithUIDNameNsSpec("12345678", "foo", "new", v1.PodSpec{
		Containers: []v1.Container{
			{
				Name:            "bar",
				Image:           "test:latest",
				ImagePullPolicy: v1.PullAlways,
			},
		},
		EnableServiceLinks: utilpointer.Bool(false),
	})

	_, err = kl.SyncPod(context.Background(), kubetypes.SyncPodCreate, pod, nil, &kubecontainer.PodStatus{})
	require.NoError(t, err)

	require.NoError(t, err)
	assert.NotEmpty(t, exp.GetSpans())

	// find root span for syncPod
	var rootSpan *tracetest.SpanStub
	spans := exp.GetSpans()
	for i, span := range spans {
		if span.Name == "syncPod" {
			rootSpan = &spans[i]
			break
		}
	}
	assert.NotNil(t, rootSpan)

	imageServiceSpans := make([]tracetest.SpanStub, 0)
	runtimeServiceSpans := make([]tracetest.SpanStub, 0)
	for _, span := range exp.GetSpans() {
		if span.SpanContext.TraceID() == rootSpan.SpanContext.TraceID() {
			switch {
			case strings.HasPrefix(span.Name, "runtime.v1.ImageService"):
				imageServiceSpans = append(imageServiceSpans, span)
			case strings.HasPrefix(span.Name, "runtime.v1.RuntimeService"):
				runtimeServiceSpans = append(runtimeServiceSpans, span)
			}
		}
	}
	assert.NotEmpty(t, imageServiceSpans, "syncPod trace should have image service spans")
	assert.NotEmpty(t, runtimeServiceSpans, "syncPod trace should have runtime service spans")

	for _, span := range imageServiceSpans {
		assert.Equal(t, span.Parent.SpanID(), rootSpan.SpanContext.SpanID(), fmt.Sprintf("image service span %s %s should be child of root span", span.Name, span.Parent.SpanID()))
	}

	for _, span := range runtimeServiceSpans {
		assert.Equal(t, span.Parent.SpanID(), rootSpan.SpanContext.SpanID(), fmt.Sprintf("runtime service span %s %s should be child of root span", span.Name, span.Parent.SpanID()))
	}
}
