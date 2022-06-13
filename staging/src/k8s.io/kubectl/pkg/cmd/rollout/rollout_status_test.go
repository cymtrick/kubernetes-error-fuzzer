package rollout

import (
	"bytes"
	"io/ioutil"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest/fake"
	cgtesting "k8s.io/client-go/testing"
	"k8s.io/kubectl/pkg/scheme"
	"net/http"
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdtesting "k8s.io/kubectl/pkg/cmd/testing"
)

var rolloutStatusGroupVersionEncoder = schema.GroupVersion{Group: "apps", Version: "v1"}

func TestRolloutStatus(t *testing.T) {
	deploymentName := "deployment/nginx-deployment"
	ns := scheme.Codecs.WithoutConversion()
	tf := cmdtesting.NewTestFactory().WithNamespace("test")
	tf.ClientConfigVal = cmdtesting.DefaultClientConfig()

	info, _ := runtime.SerializerInfoForMediaType(ns.SupportedMediaTypes(), runtime.ContentTypeJSON)
	encoder := ns.EncoderForVersion(info.Serializer, rolloutStatusGroupVersionEncoder)
	tf.Client = &fake.RESTClient{
		GroupVersion:         rolloutStatusGroupVersionEncoder,
		NegotiatedSerializer: ns,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			dep := &appsv1.Deployment{}
			dep.Name = deploymentName
			body := ioutil.NopCloser(bytes.NewReader([]byte(runtime.EncodeOrDie(encoder, dep))))
			return &http.Response{StatusCode: http.StatusOK, Header: cmdtesting.DefaultHeader(), Body: body}, nil
		}),
	}

	tf.FakeDynamicClient.WatchReactionChain = nil
	tf.FakeDynamicClient.AddWatchReactor("*", func(action cgtesting.Action) (handled bool, ret watch.Interface, err error) {
		fw := watch.NewFake()
		dep := &appsv1.Deployment{}
		dep.Name = deploymentName
		dep.Status = appsv1.DeploymentStatus{
			Replicas:            1,
			UpdatedReplicas:     1,
			ReadyReplicas:       1,
			AvailableReplicas:   1,
			UnavailableReplicas: 0,
			Conditions: []appsv1.DeploymentCondition{{
				Type: appsv1.DeploymentAvailable,
			}},
		}
		c, err := runtime.DefaultUnstructuredConverter.ToUnstructured(dep.DeepCopyObject())
		if err != nil {
			t.Errorf("unexpected err %s", err)
		}
		u := &unstructured.Unstructured{}
		u.SetUnstructuredContent(c)
		go fw.Add(u)
		return true, fw, nil
	})

	streams, _, buf, _ := genericclioptions.NewTestIOStreams()
	cmd := NewCmdRolloutStatus(tf, streams)
	cmd.Run(cmd, []string{deploymentName})

	expectedMsg := "deployment \"deployment/nginx-deployment\" successfully rolled out\n"
	if buf.String() != expectedMsg {
		t.Errorf("expected output: %s, but got: %s", expectedMsg, buf.String())
	}
}

func TestRolloutStatusWithSelector(t *testing.T) {
	deploymentName := "deployment"
	ns := scheme.Codecs.WithoutConversion()
	tf := cmdtesting.NewTestFactory().WithNamespace("test")
	tf.ClientConfigVal = cmdtesting.DefaultClientConfig()

	info, _ := runtime.SerializerInfoForMediaType(ns.SupportedMediaTypes(), runtime.ContentTypeJSON)
	encoder := ns.EncoderForVersion(info.Serializer, rolloutStatusGroupVersionEncoder)
	tf.Client = &fake.RESTClient{
		GroupVersion:         rolloutStatusGroupVersionEncoder,
		NegotiatedSerializer: ns,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			dep := &appsv1.Deployment{}
			dep.Name = deploymentName
			dep.Labels = make(map[string]string)
			dep.Labels["app"] = "api"
			body := ioutil.NopCloser(bytes.NewReader([]byte(runtime.EncodeOrDie(encoder, dep))))
			return &http.Response{StatusCode: http.StatusOK, Header: cmdtesting.DefaultHeader(), Body: body}, nil
		}),
	}

	tf.FakeDynamicClient.WatchReactionChain = nil
	tf.FakeDynamicClient.AddWatchReactor("*", func(action cgtesting.Action) (handled bool, ret watch.Interface, err error) {
		fw := watch.NewFake()
		dep := &appsv1.Deployment{}
		dep.Name = deploymentName
		dep.Status = appsv1.DeploymentStatus{
			Replicas:            1,
			UpdatedReplicas:     1,
			ReadyReplicas:       1,
			AvailableReplicas:   1,
			UnavailableReplicas: 0,
			Conditions: []appsv1.DeploymentCondition{{
				Type: appsv1.DeploymentAvailable,
			}},
		}
		dep.Labels = make(map[string]string)
		dep.Labels["app"] = "api"
		c, err := runtime.DefaultUnstructuredConverter.ToUnstructured(dep.DeepCopyObject())
		if err != nil {
			t.Errorf("unexpected err %s", err)
		}
		u := &unstructured.Unstructured{}
		u.SetUnstructuredContent(c)
		go fw.Add(u)
		return true, fw, nil
	})

	streams, _, buf, _ := genericclioptions.NewTestIOStreams()
	cmd := NewCmdRolloutStatus(tf, streams)
	cmd.Flags().Set("selector", "app=api")
	cmd.Run(cmd, []string{deploymentName})

	expectedMsg := "deployment \"deployment\" successfully rolled out\n"
	if buf.String() != expectedMsg {
		t.Errorf("expected output: %s, but got: %s", expectedMsg, buf.String())
	}
}

func TestRolloutStatusWatchDisabled(t *testing.T) {
	deploymentName := "deployment/nginx-deployment"
	ns := scheme.Codecs.WithoutConversion()
	tf := cmdtesting.NewTestFactory().WithNamespace("test")
	tf.ClientConfigVal = cmdtesting.DefaultClientConfig()

	info, _ := runtime.SerializerInfoForMediaType(ns.SupportedMediaTypes(), runtime.ContentTypeJSON)
	encoder := ns.EncoderForVersion(info.Serializer, rolloutStatusGroupVersionEncoder)
	tf.Client = &fake.RESTClient{
		GroupVersion:         rolloutStatusGroupVersionEncoder,
		NegotiatedSerializer: ns,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			dep := &appsv1.Deployment{}
			dep.Name = deploymentName
			body := ioutil.NopCloser(bytes.NewReader([]byte(runtime.EncodeOrDie(encoder, dep))))
			return &http.Response{StatusCode: http.StatusOK, Header: cmdtesting.DefaultHeader(), Body: body}, nil
		}),
	}

	tf.FakeDynamicClient.WatchReactionChain = nil
	tf.FakeDynamicClient.AddWatchReactor("*", func(action cgtesting.Action) (handled bool, ret watch.Interface, err error) {
		fw := watch.NewFake()
		dep := &appsv1.Deployment{}
		dep.Name = deploymentName
		dep.Status = appsv1.DeploymentStatus{
			Replicas:            1,
			UpdatedReplicas:     1,
			ReadyReplicas:       1,
			AvailableReplicas:   1,
			UnavailableReplicas: 0,
			Conditions: []appsv1.DeploymentCondition{{
				Type: appsv1.DeploymentAvailable,
			}},
		}
		c, err := runtime.DefaultUnstructuredConverter.ToUnstructured(dep.DeepCopyObject())
		if err != nil {
			t.Errorf("unexpected err %s", err)
		}
		u := &unstructured.Unstructured{}
		u.SetUnstructuredContent(c)
		go fw.Add(u)
		return true, fw, nil
	})

	streams, _, buf, _ := genericclioptions.NewTestIOStreams()
	cmd := NewCmdRolloutStatus(tf, streams)
	cmd.Flags().Set("watch", "false")
	cmd.Run(cmd, []string{deploymentName})

	expectedMsg := "deployment \"deployment/nginx-deployment\" successfully rolled out\n"
	if buf.String() != expectedMsg {
		t.Errorf("expected output: %s, but got: %s", expectedMsg, buf.String())
	}
}

func TestRolloutStatusWatchDisabledUnavailable(t *testing.T) {
	deploymentName := "deployment/nginx-deployment"
	ns := scheme.Codecs.WithoutConversion()
	tf := cmdtesting.NewTestFactory().WithNamespace("test")
	tf.ClientConfigVal = cmdtesting.DefaultClientConfig()

	info, _ := runtime.SerializerInfoForMediaType(ns.SupportedMediaTypes(), runtime.ContentTypeJSON)
	encoder := ns.EncoderForVersion(info.Serializer, rolloutStatusGroupVersionEncoder)
	tf.Client = &fake.RESTClient{
		GroupVersion:         rolloutStatusGroupVersionEncoder,
		NegotiatedSerializer: ns,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			dep := &appsv1.Deployment{}
			dep.Name = deploymentName
			body := ioutil.NopCloser(bytes.NewReader([]byte(runtime.EncodeOrDie(encoder, dep))))
			return &http.Response{StatusCode: http.StatusOK, Header: cmdtesting.DefaultHeader(), Body: body}, nil
		}),
	}

	tf.FakeDynamicClient.WatchReactionChain = nil
	tf.FakeDynamicClient.AddWatchReactor("*", func(action cgtesting.Action) (handled bool, ret watch.Interface, err error) {
		fw := watch.NewFake()
		dep := &appsv1.Deployment{}
		dep.Name = deploymentName
		dep.Status = appsv1.DeploymentStatus{
			Replicas:            1,
			UpdatedReplicas:     1,
			ReadyReplicas:       1,
			AvailableReplicas:   0,
			UnavailableReplicas: 0,
			Conditions: []appsv1.DeploymentCondition{{
				Type: appsv1.DeploymentAvailable,
			}},
		}
		c, err := runtime.DefaultUnstructuredConverter.ToUnstructured(dep.DeepCopyObject())
		if err != nil {
			t.Errorf("unexpected err %s", err)
		}
		u := &unstructured.Unstructured{}
		u.SetUnstructuredContent(c)
		go fw.Add(u)
		return true, fw, nil
	})

	streams, _, buf, _ := genericclioptions.NewTestIOStreams()
	cmd := NewCmdRolloutStatus(tf, streams)
	cmd.Flags().Set("watch", "false")
	cmd.Run(cmd, []string{deploymentName})

	expectedMsg := "Waiting for deployment \"deployment/nginx-deployment\" rollout to finish: 0 of 1 updated replicas are available...\n"
	if buf.String() != expectedMsg {
		t.Errorf("expected output: %s, but got: %s", expectedMsg, buf.String())
	}
}
