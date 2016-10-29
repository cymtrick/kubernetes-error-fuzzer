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

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apimachinery/registered"
	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/client/restclient/fake"
	"k8s.io/kubernetes/pkg/kubectl"
	cmdtesting "k8s.io/kubernetes/pkg/kubectl/cmd/testing"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/strings"
)

func initTestErrorHandler(t *testing.T) {
	cmdutil.BehaviorOnFatal(func(str string, code int) {
		t.Errorf("Error running command (exit code %d): %s", code, str)
	})
}

func defaultHeader() http.Header {
	header := http.Header{}
	header.Set("Content-Type", runtime.ContentTypeJSON)
	return header
}

func defaultClientConfig() *restclient.Config {
	return &restclient.Config{
		APIPath: "/api",
		ContentConfig: restclient.ContentConfig{
			NegotiatedSerializer: api.Codecs,
			ContentType:          runtime.ContentTypeJSON,
			GroupVersion:         &registered.GroupOrDie(api.GroupName).GroupVersion,
		},
	}
}

func defaultClientConfigForVersion(version *unversioned.GroupVersion) *restclient.Config {
	return &restclient.Config{
		APIPath: "/api",
		ContentConfig: restclient.ContentConfig{
			NegotiatedSerializer: api.Codecs,
			ContentType:          runtime.ContentTypeJSON,
			GroupVersion:         version,
		},
	}
}

type testPrinter struct {
	Objects []runtime.Object
	Err     error
}

func (t *testPrinter) PrintObj(obj runtime.Object, out io.Writer) error {
	t.Objects = append(t.Objects, obj)
	fmt.Fprintf(out, "%#v", obj)
	return t.Err
}

// TODO: implement HandledResources()
func (t *testPrinter) HandledResources() []string {
	return []string{}
}

func (t *testPrinter) AfterPrint(output io.Writer, res string) error {
	return nil
}

type testDescriber struct {
	Name, Namespace string
	Settings        kubectl.DescriberSettings
	Output          string
	Err             error
}

func (t *testDescriber) Describe(namespace, name string, describerSettings kubectl.DescriberSettings) (output string, err error) {
	t.Namespace, t.Name = namespace, name
	t.Settings = describerSettings
	return t.Output, t.Err
}

func objBody(codec runtime.Codec, obj runtime.Object) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(runtime.EncodeOrDie(codec, obj))))
}

func stringBody(body string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(body)))
}

// TODO(jlowdermilk): refactor the Factory so we can test client versions properly,
// with different client/server version skew scenarios.
// Verify that resource.RESTClients constructed from a factory respect mapping.APIVersion
//func TestClientVersions(t *testing.T) {
//	f := cmdutil.NewFactory(nil)
//
//	version := testapi.Default.Version()
//	mapping := &meta.RESTMapping{
//		APIVersion: version,
//	}
//	c, err := f.ClientForMapping(mapping)
//	if err != nil {
//		t.Errorf("unexpected error: %v", err)
//	}
//	client := c.(*client.RESTClient)
//	if client.APIVersion() != version {
//		t.Errorf("unexpected Client APIVersion: %s %v", client.APIVersion, client)
//	}
//}

func Example_printReplicationControllerWithNamespace() {
	f, tf, _, ns := cmdtesting.NewAPIFactory()
	tf.Printer = kubectl.NewHumanReadablePrinter(kubectl.PrintOptions{
		WithNamespace: true,
		ColumnLabels:  []string{},
	})
	tf.Client = &fake.RESTClient{
		NegotiatedSerializer: ns,
		Client:               nil,
	}
	cmd := NewCmdRun(f, os.Stdin, os.Stdout, os.Stderr)
	ctrl := &api.ReplicationController{
		ObjectMeta: api.ObjectMeta{
			Name:              "foo",
			Namespace:         "beep",
			Labels:            map[string]string{"foo": "bar"},
			CreationTimestamp: unversioned.Time{Time: time.Now().AddDate(-10, 0, 0)},
		},
		Spec: api.ReplicationControllerSpec{
			Replicas: 1,
			Selector: map[string]string{"foo": "bar"},
			Template: &api.PodTemplateSpec{
				ObjectMeta: api.ObjectMeta{
					Labels: map[string]string{"foo": "bar"},
				},
				Spec: api.PodSpec{
					Containers: []api.Container{
						{
							Name:  "foo",
							Image: "someimage",
						},
					},
				},
			},
		},
		Status: api.ReplicationControllerStatus{
			Replicas:      1,
			ReadyReplicas: 1,
		},
	}
	mapper, _ := f.Object()
	err := f.PrintObject(cmd, mapper, ctrl, os.Stdout)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
	}
	// Output:
	// NAMESPACE   NAME      DESIRED   CURRENT   READY     AGE
	// beep        foo       1         1         1         10y
}

func Example_printMultiContainersReplicationControllerWithWide() {
	f, tf, _, ns := cmdtesting.NewAPIFactory()
	tf.Printer = kubectl.NewHumanReadablePrinter(kubectl.PrintOptions{
		Wide:         true,
		ColumnLabels: []string{},
	})
	tf.Client = &fake.RESTClient{
		NegotiatedSerializer: ns,
		Client:               nil,
	}
	cmd := NewCmdRun(f, os.Stdin, os.Stdout, os.Stderr)
	ctrl := &api.ReplicationController{
		ObjectMeta: api.ObjectMeta{
			Name:              "foo",
			Labels:            map[string]string{"foo": "bar"},
			CreationTimestamp: unversioned.Time{Time: time.Now().AddDate(-10, 0, 0)},
		},
		Spec: api.ReplicationControllerSpec{
			Replicas: 1,
			Selector: map[string]string{"foo": "bar"},
			Template: &api.PodTemplateSpec{
				ObjectMeta: api.ObjectMeta{
					Labels: map[string]string{"foo": "bar"},
				},
				Spec: api.PodSpec{
					Containers: []api.Container{
						{
							Name:  "foo",
							Image: "someimage",
						},
						{
							Name:  "foo2",
							Image: "someimage2",
						},
					},
				},
			},
		},
		Status: api.ReplicationControllerStatus{
			Replicas: 1,
		},
	}
	mapper, _ := f.Object()
	err := f.PrintObject(cmd, mapper, ctrl, os.Stdout)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
	}
	// Output:
	// NAME      DESIRED   CURRENT   READY     AGE       CONTAINER(S)   IMAGE(S)               SELECTOR
	// foo       1         1         0         10y       foo,foo2       someimage,someimage2   foo=bar
}

func Example_printReplicationController() {
	f, tf, _, ns := cmdtesting.NewAPIFactory()
	tf.Printer = kubectl.NewHumanReadablePrinter(kubectl.PrintOptions{
		ColumnLabels: []string{},
	})
	tf.Client = &fake.RESTClient{
		NegotiatedSerializer: ns,
		Client:               nil,
	}
	cmd := NewCmdRun(f, os.Stdin, os.Stdout, os.Stderr)
	ctrl := &api.ReplicationController{
		ObjectMeta: api.ObjectMeta{
			Name:              "foo",
			Labels:            map[string]string{"foo": "bar"},
			CreationTimestamp: unversioned.Time{Time: time.Now().AddDate(-10, 0, 0)},
		},
		Spec: api.ReplicationControllerSpec{
			Replicas: 1,
			Selector: map[string]string{"foo": "bar"},
			Template: &api.PodTemplateSpec{
				ObjectMeta: api.ObjectMeta{
					Labels: map[string]string{"foo": "bar"},
				},
				Spec: api.PodSpec{
					Containers: []api.Container{
						{
							Name:  "foo",
							Image: "someimage",
						},
						{
							Name:  "foo2",
							Image: "someimage",
						},
					},
				},
			},
		},
		Status: api.ReplicationControllerStatus{
			Replicas: 1,
		},
	}
	mapper, _ := f.Object()
	err := f.PrintObject(cmd, mapper, ctrl, os.Stdout)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
	}
	// Output:
	// NAME      DESIRED   CURRENT   READY     AGE
	// foo       1         1         0         10y
}

func Example_printPodWithWideFormat() {
	f, tf, _, ns := cmdtesting.NewAPIFactory()
	tf.Printer = kubectl.NewHumanReadablePrinter(kubectl.PrintOptions{
		Wide:         true,
		ColumnLabels: []string{},
	})
	tf.Client = &fake.RESTClient{
		NegotiatedSerializer: ns,
		Client:               nil,
	}
	nodeName := "kubernetes-node-abcd"
	cmd := NewCmdRun(f, os.Stdin, os.Stdout, os.Stderr)
	pod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			Name:              "test1",
			CreationTimestamp: unversioned.Time{Time: time.Now().AddDate(-10, 0, 0)},
		},
		Spec: api.PodSpec{
			Containers: make([]api.Container, 2),
			NodeName:   nodeName,
		},
		Status: api.PodStatus{
			Phase: "podPhase",
			ContainerStatuses: []api.ContainerStatus{
				{Ready: true, RestartCount: 3, State: api.ContainerState{Running: &api.ContainerStateRunning{}}},
				{RestartCount: 3},
			},
			PodIP: "10.1.1.3",
		},
	}
	mapper, _ := f.Object()
	err := f.PrintObject(cmd, mapper, pod, os.Stdout)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
	}
	// Output:
	// NAME      READY     STATUS     RESTARTS   AGE       IP         NODE
	// test1     1/2       podPhase   6          10y       10.1.1.3   kubernetes-node-abcd
}

func Example_printPodWithShowLabels() {
	f, tf, _, ns := cmdtesting.NewAPIFactory()
	tf.Printer = kubectl.NewHumanReadablePrinter(kubectl.PrintOptions{
		ShowLabels:   true,
		ColumnLabels: []string{},
	})
	tf.Client = &fake.RESTClient{
		NegotiatedSerializer: ns,
		Client:               nil,
	}
	nodeName := "kubernetes-node-abcd"
	cmd := NewCmdRun(f, os.Stdin, os.Stdout, os.Stderr)
	pod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			Name:              "test1",
			CreationTimestamp: unversioned.Time{Time: time.Now().AddDate(-10, 0, 0)},
			Labels: map[string]string{
				"l1": "key",
				"l2": "value",
			},
		},
		Spec: api.PodSpec{
			Containers: make([]api.Container, 2),
			NodeName:   nodeName,
		},
		Status: api.PodStatus{
			Phase: "podPhase",
			ContainerStatuses: []api.ContainerStatus{
				{Ready: true, RestartCount: 3, State: api.ContainerState{Running: &api.ContainerStateRunning{}}},
				{RestartCount: 3},
			},
		},
	}
	mapper, _ := f.Object()
	err := f.PrintObject(cmd, mapper, pod, os.Stdout)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
	}
	// Output:
	// NAME      READY     STATUS     RESTARTS   AGE       LABELS
	// test1     1/2       podPhase   6          10y       l1=key,l2=value
}

func newAllPhasePodList() *api.PodList {
	nodeName := "kubernetes-node-abcd"
	return &api.PodList{
		Items: []api.Pod{
			{
				ObjectMeta: api.ObjectMeta{
					Name:              "test1",
					CreationTimestamp: unversioned.Time{Time: time.Now().AddDate(-10, 0, 0)},
				},
				Spec: api.PodSpec{
					Containers: make([]api.Container, 2),
					NodeName:   nodeName,
				},
				Status: api.PodStatus{
					Phase: api.PodPending,
					ContainerStatuses: []api.ContainerStatus{
						{Ready: true, RestartCount: 3, State: api.ContainerState{Running: &api.ContainerStateRunning{}}},
						{RestartCount: 3},
					},
				},
			},
			{
				ObjectMeta: api.ObjectMeta{
					Name:              "test2",
					CreationTimestamp: unversioned.Time{Time: time.Now().AddDate(-10, 0, 0)},
				},
				Spec: api.PodSpec{
					Containers: make([]api.Container, 2),
					NodeName:   nodeName,
				},
				Status: api.PodStatus{
					Phase: api.PodRunning,
					ContainerStatuses: []api.ContainerStatus{
						{Ready: true, RestartCount: 3, State: api.ContainerState{Running: &api.ContainerStateRunning{}}},
						{RestartCount: 3},
					},
				},
			},
			{
				ObjectMeta: api.ObjectMeta{
					Name:              "test3",
					CreationTimestamp: unversioned.Time{Time: time.Now().AddDate(-10, 0, 0)},
				},
				Spec: api.PodSpec{
					Containers: make([]api.Container, 2),
					NodeName:   nodeName,
				},
				Status: api.PodStatus{
					Phase: api.PodSucceeded,
					ContainerStatuses: []api.ContainerStatus{
						{Ready: true, RestartCount: 3, State: api.ContainerState{Running: &api.ContainerStateRunning{}}},
						{RestartCount: 3},
					},
				},
			},
			{
				ObjectMeta: api.ObjectMeta{
					Name:              "test4",
					CreationTimestamp: unversioned.Time{Time: time.Now().AddDate(-10, 0, 0)},
				},
				Spec: api.PodSpec{
					Containers: make([]api.Container, 2),
					NodeName:   nodeName,
				},
				Status: api.PodStatus{
					Phase: api.PodFailed,
					ContainerStatuses: []api.ContainerStatus{
						{Ready: true, RestartCount: 3, State: api.ContainerState{Running: &api.ContainerStateRunning{}}},
						{RestartCount: 3},
					},
				},
			},
			{
				ObjectMeta: api.ObjectMeta{
					Name:              "test5",
					CreationTimestamp: unversioned.Time{Time: time.Now().AddDate(-10, 0, 0)},
				},
				Spec: api.PodSpec{
					Containers: make([]api.Container, 2),
					NodeName:   nodeName,
				},
				Status: api.PodStatus{
					Phase: api.PodUnknown,
					ContainerStatuses: []api.ContainerStatus{
						{Ready: true, RestartCount: 3, State: api.ContainerState{Running: &api.ContainerStateRunning{}}},
						{RestartCount: 3},
					},
				},
			}},
	}
}

func Example_printPodHideTerminated() {
	f, tf, _, ns := cmdtesting.NewAPIFactory()
	tf.Printer = kubectl.NewHumanReadablePrinter(kubectl.PrintOptions{
		ColumnLabels: []string{},
	})
	tf.Client = &fake.RESTClient{
		NegotiatedSerializer: ns,
		Client:               nil,
	}
	cmd := NewCmdRun(f, os.Stdin, os.Stdout, os.Stderr)
	podList := newAllPhasePodList()
	// filter pods
	filterFuncs := f.DefaultResourceFilterFunc()
	filterOpts := f.DefaultResourceFilterOptions(cmd, false)
	_, filteredPodList, errs := cmdutil.FilterResourceList(podList, filterFuncs, filterOpts)
	if errs != nil {
		fmt.Printf("Unexpected filter error: %v\n", errs)
	}
	for _, pod := range filteredPodList {
		mapper, _ := f.Object()
		err := f.PrintObject(cmd, mapper, pod, os.Stdout)
		if err != nil {
			fmt.Printf("Unexpected error: %v", err)
		}
	}
	// Output:
	// NAME      READY     STATUS    RESTARTS   AGE
	// test1     1/2       Pending   6          10y
	// test2     1/2       Running   6         10y
	// test5     1/2       Unknown   6         10y
}

func Example_printPodShowAll() {
	f, tf, _, ns := cmdtesting.NewAPIFactory()
	tf.Printer = kubectl.NewHumanReadablePrinter(kubectl.PrintOptions{
		ShowAll:      true,
		ColumnLabels: []string{},
	})
	tf.Client = &fake.RESTClient{
		NegotiatedSerializer: ns,
		Client:               nil,
	}
	cmd := NewCmdRun(f, os.Stdin, os.Stdout, os.Stderr)
	podList := newAllPhasePodList()
	mapper, _ := f.Object()
	err := f.PrintObject(cmd, mapper, podList, os.Stdout)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
	}
	// Output:
	// NAME      READY     STATUS      RESTARTS   AGE
	// test1     1/2       Pending     6          10y
	// test2     1/2       Running     6          10y
	// test3     1/2       Succeeded   6          10y
	// test4     1/2       Failed      6          10y
	// test5     1/2       Unknown     6          10y
}

func Example_printServiceWithNamespacesAndLabels() {
	f, tf, _, ns := cmdtesting.NewAPIFactory()
	tf.Printer = kubectl.NewHumanReadablePrinter(kubectl.PrintOptions{
		WithNamespace: true,
		ColumnLabels:  []string{"l1"},
	})
	tf.Client = &fake.RESTClient{
		NegotiatedSerializer: ns,
		Client:               nil,
	}
	cmd := NewCmdRun(f, os.Stdin, os.Stdout, os.Stderr)
	svc := &api.ServiceList{
		Items: []api.Service{
			{
				ObjectMeta: api.ObjectMeta{
					Name:              "svc1",
					Namespace:         "ns1",
					CreationTimestamp: unversioned.Time{Time: time.Now().AddDate(-10, 0, 0)},
					Labels: map[string]string{
						"l1": "value",
					},
				},
				Spec: api.ServiceSpec{
					Ports: []api.ServicePort{
						{Protocol: "UDP", Port: 53},
						{Protocol: "TCP", Port: 53},
					},
					Selector: map[string]string{
						"s": "magic",
					},
					ClusterIP: "10.1.1.1",
				},
				Status: api.ServiceStatus{},
			},
			{
				ObjectMeta: api.ObjectMeta{
					Name:              "svc2",
					Namespace:         "ns2",
					CreationTimestamp: unversioned.Time{Time: time.Now().AddDate(-10, 0, 0)},
					Labels: map[string]string{
						"l1": "dolla-bill-yall",
					},
				},
				Spec: api.ServiceSpec{
					Ports: []api.ServicePort{
						{Protocol: "TCP", Port: 80},
						{Protocol: "TCP", Port: 8080},
					},
					Selector: map[string]string{
						"s": "kazam",
					},
					ClusterIP: "10.1.1.2",
				},
				Status: api.ServiceStatus{},
			}},
	}
	ld := strings.NewLineDelimiter(os.Stdout, "|")
	defer ld.Flush()

	mapper, _ := f.Object()
	err := f.PrintObject(cmd, mapper, svc, ld)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
	}
	// Output:
	// |NAMESPACE   NAME      CLUSTER-IP   EXTERNAL-IP   PORT(S)           AGE       L1|
	// |ns1         svc1      10.1.1.1     <unknown>     53/UDP,53/TCP     10y       value|
	// |ns2         svc2      10.1.1.2     <unknown>     80/TCP,8080/TCP   10y       dolla-bill-yall|
	// ||
}

func TestNormalizationFuncGlobalExistence(t *testing.T) {
	// This test can be safely deleted when we will not support multiple flag formats
	root := NewKubectlCommand(cmdutil.NewFactory(nil), os.Stdin, os.Stdout, os.Stderr)

	if root.Parent() != nil {
		t.Fatal("We expect the root command to be returned")
	}
	if root.GlobalNormalizationFunc() == nil {
		t.Fatal("We expect that root command has a global normalization function")
	}

	if reflect.ValueOf(root.GlobalNormalizationFunc()).Pointer() != reflect.ValueOf(root.Flags().GetNormalizeFunc()).Pointer() {
		t.Fatal("root command seems to have a wrong normalization function")
	}

	sub := root
	for sub.HasSubCommands() {
		sub = sub.Commands()[0]
	}

	// In case of failure of this test check this PR: spf13/cobra#110
	if reflect.ValueOf(sub.Flags().GetNormalizeFunc()).Pointer() != reflect.ValueOf(root.Flags().GetNormalizeFunc()).Pointer() {
		t.Fatal("child and root commands should have the same normalization functions")
	}
}
