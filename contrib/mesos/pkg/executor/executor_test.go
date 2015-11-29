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

package executor

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	assertext "k8s.io/kubernetes/contrib/mesos/pkg/assert"
	"k8s.io/kubernetes/contrib/mesos/pkg/executor/messages"
	"k8s.io/kubernetes/contrib/mesos/pkg/podutil"
	kmruntime "k8s.io/kubernetes/contrib/mesos/pkg/runtime"
	"k8s.io/kubernetes/contrib/mesos/pkg/scheduler/podtask"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/testapi"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/client/cache"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/kubelet"
	"k8s.io/kubernetes/pkg/kubelet/dockertools"
	kubetypes "k8s.io/kubernetes/pkg/kubelet/types"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util"
	"k8s.io/kubernetes/pkg/watch"

	"github.com/mesos/mesos-go/mesosproto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestExecutorRegister ensures that the executor thinks it is connected
// after Register is called.
func TestExecutorRegister(t *testing.T) {
	mockDriver := &MockExecutorDriver{}
	executor, updates := NewTestKubernetesExecutor()

	executor.Init(mockDriver)
	executor.Registered(mockDriver, nil, nil, nil)

	initialPodUpdate := kubetypes.PodUpdate{
		Pods: []*api.Pod{},
		Op:   kubetypes.SET,
	}
	receivedInitialPodUpdate := false
	select {
	case update := <-updates:
		if reflect.DeepEqual(initialPodUpdate, update) {
			receivedInitialPodUpdate = true
		}
	case <-time.After(util.ForeverTestTimeout):
	}
	assert.Equal(t, true, receivedInitialPodUpdate,
		"executor should have sent an initial PodUpdate "+
			"to the updates chan upon registration")

	assert.Equal(t, true, executor.isConnected(), "executor should be connected")
	mockDriver.AssertExpectations(t)
}

// TestExecutorDisconnect ensures that the executor thinks that it is not
// connected after a call to Disconnected has occurred.
func TestExecutorDisconnect(t *testing.T) {
	mockDriver := &MockExecutorDriver{}
	executor, _ := NewTestKubernetesExecutor()

	executor.Init(mockDriver)
	executor.Registered(mockDriver, nil, nil, nil)
	executor.Disconnected(mockDriver)

	assert.Equal(t, false, executor.isConnected(),
		"executor should not be connected after Disconnected")
	mockDriver.AssertExpectations(t)
}

// TestExecutorReregister ensures that the executor thinks it is connected
// after a connection problem happens, followed by a call to Reregistered.
func TestExecutorReregister(t *testing.T) {
	mockDriver := &MockExecutorDriver{}
	executor, _ := NewTestKubernetesExecutor()

	executor.Init(mockDriver)
	executor.Registered(mockDriver, nil, nil, nil)
	executor.Disconnected(mockDriver)
	executor.Reregistered(mockDriver, nil)

	assert.Equal(t, true, executor.isConnected(), "executor should be connected")
	mockDriver.AssertExpectations(t)
}

type fakeKubelet struct {
	*kubelet.Kubelet
	hostIP net.IP
}

func (kl *fakeKubelet) GetHostIP() (net.IP, error) {
	return kl.hostIP, nil
}

// TestExecutorLaunchAndKillTask ensures that the executor is able to launch
// and kill tasks while properly bookkeping its tasks.
func TestExecutorLaunchAndKillTask(t *testing.T) {
	// create a fake pod watch. We use that below to submit new pods to the scheduler
	podListWatch := NewMockPodsListWatch(api.PodList{})

	// create fake apiserver
	testApiServer := NewTestServer(t, api.NamespaceDefault, &podListWatch.list)
	defer testApiServer.server.Close()

	mockDriver := &MockExecutorDriver{}
	updates := make(chan kubetypes.PodUpdate, 1024)
	config := Config{
		Docker:    dockertools.ConnectToDockerOrDie("fake://"),
		Updates:   updates,
		NodeInfos: make(chan NodeInfo, 1),
		APIClient: client.NewOrDie(&client.Config{
			Host:         testApiServer.server.URL,
			GroupVersion: testapi.Default.GroupVersion(),
		}),
		PodStatusFunc: func(pod *api.Pod) (*api.PodStatus, error) {
			return &api.PodStatus{
				ContainerStatuses: []api.ContainerStatus{
					{
						Name: "foo",
						State: api.ContainerState{
							Running: &api.ContainerStateRunning{},
						},
					},
				},
				Phase:  api.PodRunning,
				HostIP: "127.0.0.1",
			}, nil
		},
		PodLW: &NewMockPodsListWatch(api.PodList{}).ListWatch,
	}
	executor := New(config)

	executor.Init(mockDriver)
	executor.Registered(mockDriver, nil, nil, nil)

	select {
	case <-updates:
	case <-time.After(util.ForeverTestTimeout):
		t.Fatalf("Executor should send an initial update on Registration")
	}

	pod := NewTestPod(1)
	executorinfo := &mesosproto.ExecutorInfo{}
	podTask, err := podtask.New(
		api.NewDefaultContext(),
		"",
		pod,
		executorinfo,
		nil,
	)

	assert.Equal(t, nil, err, "must be able to create a task from a pod")

	podTask.Spec = &podtask.Spec{
		Executor: executorinfo,
	}
	taskInfo, err := podTask.BuildTaskInfo()
	assert.Equal(t, nil, err, "must be able to build task info")

	data, err := testapi.Default.Codec().Encode(pod)
	assert.Equal(t, nil, err, "must be able to encode a pod's spec data")
	taskInfo.Data = data
	var statusUpdateCalls sync.WaitGroup
	statusUpdateDone := func(_ mock.Arguments) { statusUpdateCalls.Done() }

	statusUpdateCalls.Add(1)
	mockDriver.On(
		"SendStatusUpdate",
		mesosproto.TaskState_TASK_STARTING,
	).Return(mesosproto.Status_DRIVER_RUNNING, nil).Run(statusUpdateDone).Once()

	statusUpdateCalls.Add(1)
	mockDriver.On(
		"SendStatusUpdate",
		mesosproto.TaskState_TASK_RUNNING,
	).Return(mesosproto.Status_DRIVER_RUNNING, nil).Run(statusUpdateDone).Once()

	executor.LaunchTask(mockDriver, taskInfo)

	assertext.EventuallyTrue(t, util.ForeverTestTimeout, func() bool {
		executor.lock.Lock()
		defer executor.lock.Unlock()
		return len(executor.tasks) == 1 && len(executor.pods) == 1
	}, "executor must be able to create a task and a pod")

	gotPodUpdate := false
	select {
	case update := <-updates:
		if len(update.Pods) == 1 {
			gotPodUpdate = true
		}
	case <-time.After(util.ForeverTestTimeout):
	}
	assert.Equal(t, true, gotPodUpdate,
		"the executor should send an update about a new pod to "+
			"the updates chan when creating a new one.")

	// Allow some time for asynchronous requests to the driver.
	finished := kmruntime.After(statusUpdateCalls.Wait)
	select {
	case <-finished:
	case <-time.After(util.ForeverTestTimeout):
		t.Fatalf("timed out waiting for status update calls to finish")
	}

	statusUpdateCalls.Add(1)
	mockDriver.On(
		"SendStatusUpdate",
		mesosproto.TaskState_TASK_KILLED,
	).Return(mesosproto.Status_DRIVER_RUNNING, nil).Run(statusUpdateDone).Once()

	executor.KillTask(mockDriver, taskInfo.TaskId)

	assertext.EventuallyTrue(t, util.ForeverTestTimeout, func() bool {
		executor.lock.Lock()
		defer executor.lock.Unlock()
		return len(executor.tasks) == 0 && len(executor.pods) == 0
	}, "executor must be able to kill a created task and pod")

	// Allow some time for asynchronous requests to the driver.
	finished = kmruntime.After(statusUpdateCalls.Wait)
	select {
	case <-finished:
	case <-time.After(util.ForeverTestTimeout):
		t.Fatalf("timed out waiting for status update calls to finish")
	}
	mockDriver.AssertExpectations(t)
}

// TestExecutorStaticPods test that the ExecutorInfo.data is parsed
// as a zip archive with pod definitions.
func TestExecutorInitializeStaticPodsSource(t *testing.T) {
	// create some zip with static pod definition
	givenPodsDir, err := ioutil.TempDir("/tmp", "executor-givenpods")
	assert.NoError(t, err)
	defer os.RemoveAll(givenPodsDir)

	var wg sync.WaitGroup
	reportErrors := func(errCh <-chan error) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for err := range errCh {
				t.Error(err)
			}
		}()
	}

	createStaticPodFile := func(fileName, name string) {
		spod := `{
			"apiVersion": "v1",
			"kind": "Pod",
			"metadata": {
				"name": "%v",
				"namespace": "staticpods",
				"labels": { "name": "foo", "cluster": "bar" }
			},
			"spec": {
				"containers": [{
					"name": "%v",
					"image": "library/nginx",
					"ports": [{ "containerPort": 80, "name": "http" }]
				}]
			}
		}`
		destfile := filepath.Join(givenPodsDir, fileName)
		err = os.MkdirAll(filepath.Dir(destfile), 0770)
		assert.NoError(t, err)
		err = ioutil.WriteFile(destfile, []byte(fmt.Sprintf(spod, name, name)), 0660)
		assert.NoError(t, err)
	}

	createStaticPodFile("spod.json", "spod-01")
	createStaticPodFile("spod2.json", "spod-02")
	createStaticPodFile("dir/spod.json", "spod-03") // same file name as first one to check for overwriting
	staticpods, errs := podutil.ReadFromDir(givenPodsDir)
	reportErrors(errs)

	gzipped, err := podutil.Gzip(staticpods)
	assert.NoError(t, err)

	expectedStaticPodsNum := 2 // subdirectories are ignored by FileSource, hence only 2

	// temporary directory which is normally located in the executor sandbox
	staticPodsConfigPath, err := ioutil.TempDir("/tmp", "executor-k8sm-archive")
	assert.NoError(t, err)
	defer os.RemoveAll(staticPodsConfigPath)

	executor := &Executor{
		staticPodsConfigPath: staticPodsConfigPath,
	}

	// extract the pods into staticPodsConfigPath
	hostname := "h1"
	err = executor.initializeStaticPodsSource(hostname, gzipped)
	assert.NoError(t, err)

	actualpods, errs := podutil.ReadFromDir(staticPodsConfigPath)
	reportErrors(errs)

	list := podutil.List(actualpods)
	assert.NotNil(t, list)
	assert.Equal(t, expectedStaticPodsNum, len(list.Items))
	wg.Wait()
}

// TestExecutorFrameworkMessage ensures that the executor is able to
// handle messages from the framework, specifically about lost tasks
// and Kamikaze.  When a task is lost, the executor needs to clean up
// its state.  When a Kamikaze message is received, the executor should
// attempt suicide.
func TestExecutorFrameworkMessage(t *testing.T) {
	// create fake apiserver
	podListWatch := NewMockPodsListWatch(api.PodList{})
	testApiServer := NewTestServer(t, api.NamespaceDefault, &podListWatch.list)
	defer testApiServer.server.Close()

	// create and start executor
	mockDriver := &MockExecutorDriver{}
	kubeletFinished := make(chan struct{})
	config := Config{
		Docker:    dockertools.ConnectToDockerOrDie("fake://"),
		Updates:   make(chan kubetypes.PodUpdate, 1024),
		NodeInfos: make(chan NodeInfo, 1),
		APIClient: client.NewOrDie(&client.Config{
			Host:         testApiServer.server.URL,
			GroupVersion: testapi.Default.GroupVersion(),
		}),
		PodStatusFunc: func(pod *api.Pod) (*api.PodStatus, error) {
			return &api.PodStatus{
				ContainerStatuses: []api.ContainerStatus{
					{
						Name: "foo",
						State: api.ContainerState{
							Running: &api.ContainerStateRunning{},
						},
					},
				},
				Phase:  api.PodRunning,
				HostIP: "127.0.0.1",
			}, nil
		},
		ShutdownAlert: func() {
			close(kubeletFinished)
		},
		KubeletFinished: kubeletFinished,
		PodLW:           &NewMockPodsListWatch(api.PodList{}).ListWatch,
	}
	executor := New(config)

	executor.Init(mockDriver)
	executor.Registered(mockDriver, nil, nil, nil)

	executor.FrameworkMessage(mockDriver, "test framework message")

	// set up a pod to then lose
	pod := NewTestPod(1)
	executorinfo := &mesosproto.ExecutorInfo{}
	podTask, _ := podtask.New(
		api.NewDefaultContext(),
		"foo",
		pod,
		executorinfo,
		nil,
	)

	podTask.Spec = &podtask.Spec{
		Executor: executorinfo,
	}
	taskInfo, err := podTask.BuildTaskInfo()
	assert.Equal(t, nil, err, "must be able to build task info")

	data, _ := testapi.Default.Codec().Encode(pod)
	taskInfo.Data = data

	mockDriver.On(
		"SendStatusUpdate",
		mesosproto.TaskState_TASK_STARTING,
	).Return(mesosproto.Status_DRIVER_RUNNING, nil).Once()

	called := make(chan struct{})
	mockDriver.On(
		"SendStatusUpdate",
		mesosproto.TaskState_TASK_RUNNING,
	).Return(mesosproto.Status_DRIVER_RUNNING, nil).Run(func(_ mock.Arguments) { close(called) }).Once()

	executor.LaunchTask(mockDriver, taskInfo)

	// waiting until the pod is really running b/c otherwise a TASK_FAILED could be
	// triggered by the asynchronously running  _launchTask, __launchTask methods
	// when removing the task from k.tasks through the "task-lost:foo" message below.
	select {
	case <-called:
	case <-time.After(util.ForeverTestTimeout):
		t.Fatalf("timed out waiting for SendStatusUpdate for the running task")
	}

	// send task-lost message for it
	called = make(chan struct{})
	mockDriver.On(
		"SendStatusUpdate",
		mesosproto.TaskState_TASK_LOST,
	).Return(mesosproto.Status_DRIVER_RUNNING, nil).Run(func(_ mock.Arguments) { close(called) }).Once()

	executor.FrameworkMessage(mockDriver, "task-lost:foo")
	assertext.EventuallyTrue(t, util.ForeverTestTimeout, func() bool {
		executor.lock.Lock()
		defer executor.lock.Unlock()
		return len(executor.tasks) == 0 && len(executor.pods) == 0
	}, "executor must be able to kill a created task and pod")

	select {
	case <-called:
	case <-time.After(util.ForeverTestTimeout):
		t.Fatalf("timed out waiting for SendStatusUpdate")
	}

	mockDriver.On("Stop").Return(mesosproto.Status_DRIVER_STOPPED, nil).Once()

	executor.FrameworkMessage(mockDriver, messages.Kamikaze)
	assert.Equal(t, true, executor.isDone(),
		"executor should have shut down after receiving a Kamikaze message")

	mockDriver.AssertExpectations(t)
}

// Create a pod with a given index, requiring one port
func NewTestPod(i int) *api.Pod {
	name := fmt.Sprintf("pod%d", i)
	return &api.Pod{
		TypeMeta: unversioned.TypeMeta{APIVersion: testapi.Default.GroupVersion().String()},
		ObjectMeta: api.ObjectMeta{
			Name:      name,
			Namespace: api.NamespaceDefault,
			SelfLink:  testapi.Default.SelfLink("pods", string(i)),
		},
		Spec: api.PodSpec{
			Containers: []api.Container{
				{
					Name: "foo",
					Ports: []api.ContainerPort{
						{
							ContainerPort: 8000 + i,
							Protocol:      api.ProtocolTCP,
						},
					},
				},
			},
		},
		Status: api.PodStatus{
			Conditions: []api.PodCondition{
				{
					Type:   api.PodReady,
					Status: api.ConditionTrue,
				},
			},
		},
	}
}

// Create mock of pods ListWatch, usually listening on the apiserver pods watch endpoint
type MockPodsListWatch struct {
	ListWatch   cache.ListWatch
	fakeWatcher *watch.FakeWatcher
	list        api.PodList
}

// A apiserver mock which partially mocks the pods API
type TestServer struct {
	server *httptest.Server
	Stats  map[string]uint
	lock   sync.Mutex
}

func NewTestServer(t *testing.T, namespace string, pods *api.PodList) *TestServer {
	ts := TestServer{
		Stats: map[string]uint{},
	}
	mux := http.NewServeMux()

	mux.HandleFunc(testapi.Default.ResourcePath("bindings", namespace, ""), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	ts.server = httptest.NewServer(mux)
	return &ts
}

func NewMockPodsListWatch(initialPodList api.PodList) *MockPodsListWatch {
	lw := MockPodsListWatch{
		fakeWatcher: watch.NewFake(),
		list:        initialPodList,
	}
	lw.ListWatch = cache.ListWatch{
		WatchFunc: func(options unversioned.ListOptions) (watch.Interface, error) {
			return lw.fakeWatcher, nil
		},
		ListFunc: func(options unversioned.ListOptions) (runtime.Object, error) {
			return &lw.list, nil
		},
	}
	return &lw
}

// TestExecutorShutdown ensures that the executor properly shuts down
// when Shutdown is called.
func TestExecutorShutdown(t *testing.T) {
	mockDriver := &MockExecutorDriver{}
	kubeletFinished := make(chan struct{})
	var exitCalled int32 = 0
	updates := make(chan kubetypes.PodUpdate, 1024)
	config := Config{
		Docker:    dockertools.ConnectToDockerOrDie("fake://"),
		Updates:   updates,
		NodeInfos: make(chan NodeInfo, 1),
		ShutdownAlert: func() {
			close(kubeletFinished)
		},
		KubeletFinished: kubeletFinished,
		ExitFunc: func(_ int) {
			atomic.AddInt32(&exitCalled, 1)
		},
		PodLW: &NewMockPodsListWatch(api.PodList{}).ListWatch,
	}
	executor := New(config)

	executor.Init(mockDriver)
	executor.Registered(mockDriver, nil, nil, nil)

	mockDriver.On("Stop").Return(mesosproto.Status_DRIVER_STOPPED, nil).Once()

	executor.Shutdown(mockDriver)

	assert.Equal(t, false, executor.isConnected(),
		"executor should not be connected after Shutdown")
	assert.Equal(t, true, executor.isDone(),
		"executor should be in Done state after Shutdown")

	// channel should be closed now, only a constant number of updates left
	num := len(updates)
drainLoop:
	for {
		select {
		case _, ok := <-updates:
			if !ok {
				break drainLoop
			}
			num -= 1
		default:
			t.Fatal("Updates chan should be closed after Shutdown")
		}
	}
	assert.Equal(t, num, 0, "Updates chan should get no new updates after Shutdown")

	assert.Equal(t, true, atomic.LoadInt32(&exitCalled) > 0,
		"the executor should call its ExitFunc when it is ready to close down")

	mockDriver.AssertExpectations(t)
}

func TestExecutorsendFrameworkMessage(t *testing.T) {
	mockDriver := &MockExecutorDriver{}
	executor, _ := NewTestKubernetesExecutor()

	executor.Init(mockDriver)
	executor.Registered(mockDriver, nil, nil, nil)

	called := make(chan struct{})
	mockDriver.On(
		"SendFrameworkMessage",
		"foo bar baz",
	).Return(mesosproto.Status_DRIVER_RUNNING, nil).Run(func(_ mock.Arguments) { close(called) }).Once()
	executor.sendFrameworkMessage(mockDriver, "foo bar baz")

	// guard against data race in mock driver between AssertExpectations and Called
	select {
	case <-called: // expected
	case <-time.After(util.ForeverTestTimeout):
		t.Fatalf("expected call to SendFrameworkMessage")
	}
	mockDriver.AssertExpectations(t)
}

func TestExecutor_updateMetaMap(t *testing.T) {
	for i, tc := range []struct {
		oldmap map[string]string
		newmap map[string]string
		wants  bool
	}{
		{
			oldmap: nil,
			newmap: nil,
			wants:  false,
		},
		{
			oldmap: nil,
			newmap: map[string]string{},
			wants:  false,
		},
		{
			oldmap: map[string]string{},
			newmap: nil,
			wants:  false,
		},
		{
			oldmap: nil,
			newmap: map[string]string{
				"foo": "bar",
			},
			wants: true,
		},
		{
			oldmap: map[string]string{},
			newmap: map[string]string{
				"foo": "bar",
			},
			wants: true,
		},
		{
			oldmap: map[string]string{
				"baz": "qax",
			},
			newmap: map[string]string{
				"foo": "bar",
			},
			wants: true,
		},
		{
			oldmap: map[string]string{
				"baz": "qax",
			},
			newmap: nil,
			wants:  true,
		},
		{
			oldmap: map[string]string{
				"baz": "qax",
				"qwe": "iop",
			},
			newmap: map[string]string{
				"foo": "bar",
				"qwe": "iop",
			},
			wants: true,
		},
	} {
		// do work here
		actual := updateMetaMap(&tc.oldmap, tc.newmap)
		if actual != tc.wants {
			t.Fatalf("test case %d failed, expected %v but got %v instead", i, tc.wants, actual)
		}
		if len(tc.oldmap) != len(tc.newmap) || (len(tc.oldmap) > 0 && !reflect.DeepEqual(tc.oldmap, tc.newmap)) {
			t.Fatalf("test case %d failed, expected %v but got %v instead", i, tc.newmap, tc.oldmap)
		}
	}
}
