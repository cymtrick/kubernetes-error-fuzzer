/*
Copyright 2018 The Kubernetes Authors.

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

package integration

import (
	"path"
	"reflect"
	"strings"
	"testing"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
	"sigs.k8s.io/yaml"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	serveroptions "k8s.io/apiextensions-apiserver/pkg/cmd/server/options"
	"k8s.io/apiextensions-apiserver/test/integration/fixtures"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/apimachinery/pkg/util/json"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

func TestPostInvalidObjectMeta(t *testing.T) {
	tearDown, apiExtensionClient, dynamicClient, err := fixtures.StartDefaultServerWithClients(t)
	if err != nil {
		t.Fatal(err)
	}
	defer tearDown()

	noxuDefinition := fixtures.NewNoxuCustomResourceDefinition(apiextensionsv1beta1.NamespaceScoped)
	noxuDefinition, err = fixtures.CreateNewCustomResourceDefinition(noxuDefinition, apiExtensionClient, dynamicClient)
	if err != nil {
		t.Fatal(err)
	}

	noxuResourceClient := newNamespacedCustomResourceClient("default", dynamicClient, noxuDefinition)

	obj := fixtures.NewNoxuInstance("default", "foo")
	unstructured.SetNestedField(obj.UnstructuredContent(), int64(42), "metadata", "unknown")
	unstructured.SetNestedField(obj.UnstructuredContent(), map[string]interface{}{"foo": int64(42), "bar": "abc"}, "metadata", "labels")
	_, err = instantiateCustomResource(t, obj, noxuResourceClient, noxuDefinition)
	if err == nil {
		t.Fatalf("unexpected non-error, expected invalid labels to be rejected: %v", err)
	}
	if status, ok := err.(errors.APIStatus); !ok {
		t.Fatalf("expected APIStatus error, but got: %#v", err)
	} else if !errors.IsBadRequest(err) {
		t.Fatalf("expected BadRequst error, but got: %v", errors.ReasonForError(err))
	} else if !strings.Contains(status.Status().Message, "cannot be handled") {
		t.Fatalf("expected 'cannot be handled' error message, got: %v", status.Status().Message)
	}

	unstructured.SetNestedField(obj.UnstructuredContent(), map[string]interface{}{"bar": "abc"}, "metadata", "labels")
	obj, err = instantiateCustomResource(t, obj, noxuResourceClient, noxuDefinition)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if unknown, found, err := unstructured.NestedInt64(obj.UnstructuredContent(), "metadata", "unknown"); err != nil {
		t.Errorf("unexpected error getting metadata.unknown: %v", err)
	} else if found {
		t.Errorf("unexpected metadata.unknown=%#v: expected this to be pruned", unknown)
	}
}

func TestInvalidObjectMetaInStorage(t *testing.T) {
	tearDown, config, options, err := fixtures.StartDefaultServer(t)
	if err != nil {
		t.Fatal(err)
	}
	defer tearDown()

	apiExtensionClient, err := clientset.NewForConfig(config)
	if err != nil {
		t.Fatal(err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		t.Fatal(err)
	}

	noxuDefinition := fixtures.NewNoxuCustomResourceDefinition(apiextensionsv1beta1.NamespaceScoped)
	noxuDefinition.Spec.Validation = &apiextensionsv1beta1.CustomResourceValidation{
		OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{
			Type: "object",
			Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
				"embedded": {
					Type:                   "object",
					XEmbeddedResource:      true,
					XPreserveUnknownFields: pointer.BoolPtr(true),
				},
			},
		},
	}
	noxuDefinition, err = fixtures.CreateNewCustomResourceDefinition(noxuDefinition, apiExtensionClient, dynamicClient)
	if err != nil {
		t.Fatal(err)
	}

	RESTOptionsGetter := serveroptions.NewCRDRESTOptionsGetter(*options.RecommendedOptions.Etcd)
	restOptions, err := RESTOptionsGetter.GetRESTOptions(schema.GroupResource{Group: noxuDefinition.Spec.Group, Resource: noxuDefinition.Spec.Names.Plural})
	if err != nil {
		t.Fatal(err)
	}
	tlsInfo := transport.TLSInfo{
		CertFile: restOptions.StorageConfig.Transport.CertFile,
		KeyFile:  restOptions.StorageConfig.Transport.KeyFile,
		CAFile:   restOptions.StorageConfig.Transport.CAFile,
	}
	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		t.Fatal(err)
	}
	etcdConfig := clientv3.Config{
		Endpoints: restOptions.StorageConfig.Transport.ServerList,
		TLS:       tlsConfig,
	}
	etcdclient, err := clientv3.New(etcdConfig)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Creating object with wrongly typed annotations and non-validating labels manually in etcd")

	original := fixtures.NewNoxuInstance("default", "foo")
	unstructured.SetNestedField(original.UnstructuredContent(), int64(42), "metadata", "unknown")
	unstructured.SetNestedField(original.UnstructuredContent(), map[string]interface{}{"foo": int64(42), "bar": "abc"}, "metadata", "annotations")
	unstructured.SetNestedField(original.UnstructuredContent(), map[string]interface{}{"invalid": "x y"}, "metadata", "labels")
	unstructured.SetNestedField(original.UnstructuredContent(), int64(42), "embedded", "metadata", "unknown")
	unstructured.SetNestedField(original.UnstructuredContent(), map[string]interface{}{"foo": int64(42), "bar": "abc"}, "embedded", "metadata", "annotations")
	unstructured.SetNestedField(original.UnstructuredContent(), map[string]interface{}{"invalid": "x y"}, "embedded", "metadata", "labels")
	unstructured.SetNestedField(original.UnstructuredContent(), "Foo", "embedded", "kind")
	unstructured.SetNestedField(original.UnstructuredContent(), "foo/v1", "embedded", "apiVersion")

	ctx := genericapirequest.WithNamespace(genericapirequest.NewContext(), metav1.NamespaceDefault)
	key := path.Join("/", restOptions.StorageConfig.Prefix, noxuDefinition.Spec.Group, "noxus/default/foo")
	val, _ := json.Marshal(original.UnstructuredContent())
	if _, err := etcdclient.Put(ctx, key, string(val)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Logf("Checking that invalid objects can be deleted")
	noxuResourceClient := newNamespacedCustomResourceClient("default", dynamicClient, noxuDefinition)
	if err := noxuResourceClient.Delete("foo", &metav1.DeleteOptions{}); err != nil {
		t.Fatalf("Unexpected delete error %v", err)
	}
	if _, err := etcdclient.Put(ctx, key, string(val)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Logf("Checking that ObjectMeta is pruned from unknown fields")
	obj, err := noxuResourceClient.Get("foo", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	objJSON, _ := json.Marshal(obj.Object)
	t.Logf("Got object: %v", string(objJSON))

	if unknown, found, err := unstructured.NestedFieldNoCopy(obj.UnstructuredContent(), "metadata", "unknown"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if found {
		t.Errorf("Unexpected to find metadata.unknown=%#v", unknown)
	}
	if unknown, found, err := unstructured.NestedFieldNoCopy(obj.UnstructuredContent(), "embedded", "metadata", "unknown"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if found {
		t.Errorf("Unexpected to find embedded.metadata.unknown=%#v", unknown)
	}

	t.Logf("Checking that ObjectMeta is pruned from wrongly-typed annotations")

	if annotations, found, err := unstructured.NestedStringMap(obj.UnstructuredContent(), "metadata", "annotations"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if found {
		t.Errorf("Unexpected to find metadata.annotations: %#v", annotations)
	}
	if annotations, found, err := unstructured.NestedStringMap(obj.UnstructuredContent(), "embedded", "metadata", "annotations"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if found {
		t.Errorf("Unexpected to find embedded.metadata.annotations: %#v", annotations)
	}

	t.Logf("Checking that ObjectMeta still has the non-validating labels")

	if labels, found, err := unstructured.NestedStringMap(obj.UnstructuredContent(), "metadata", "labels"); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if !found {
		t.Errorf("Expected to find metadata.labels, but didn't")
	} else if expected := map[string]string{"invalid": "x y"}; !reflect.DeepEqual(labels, expected) {
		t.Errorf("Expected metadata.labels to be %#v, got: %#v", expected, labels)
	}
	if labels, found, err := unstructured.NestedStringMap(obj.UnstructuredContent(), "embedded", "metadata", "labels"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if !found {
		t.Errorf("Expected to find embedded.metadata.labels, but didn't")
	} else if expected := map[string]string{"invalid": "x y"}; !reflect.DeepEqual(labels, expected) {
		t.Errorf("Expected embedded.metadata.labels to be %#v, got: %#v", expected, labels)
	}

	t.Logf("Trying to fail on updating with invalid labels")
	unstructured.SetNestedField(obj.Object, "changed", "metadata", "labels", "something")
	if got, err := noxuResourceClient.Update(obj, metav1.UpdateOptions{}); err == nil {
		objJSON, _ := json.Marshal(obj.Object)
		gotJSON, _ := json.Marshal(got.Object)
		t.Fatalf("Expected update error, but didn't get one\nin: %s\nresponse: %v", string(objJSON), string(gotJSON))
	}

	t.Logf("Trying to fail on updating with invalid embedded label")
	unstructured.SetNestedField(obj.Object, "fixed", "metadata", "labels", "invalid")
	if got, err := noxuResourceClient.Update(obj, metav1.UpdateOptions{}); err == nil {
		objJSON, _ := json.Marshal(obj.Object)
		gotJSON, _ := json.Marshal(got.Object)
		t.Fatalf("Expected update error, but didn't get one\nin: %s\nresponse: %v", string(objJSON), string(gotJSON))
	}

	t.Logf("Fixed all labels and update should work")
	unstructured.SetNestedField(obj.Object, "fixed", "embedded", "metadata", "labels", "invalid")
	if _, err := noxuResourceClient.Update(obj, metav1.UpdateOptions{}); err != nil {
		t.Errorf("Unexpected update error with fixed labels: %v", err)
	}

	t.Logf("Trying to fail on updating with wrongly-typed embedded label")
	unstructured.SetNestedField(obj.Object, int64(42), "embedded", "metadata", "labels", "invalid")
	if got, err := noxuResourceClient.Update(obj, metav1.UpdateOptions{}); err == nil {
		objJSON, _ := json.Marshal(obj.Object)
		gotJSON, _ := json.Marshal(got.Object)
		t.Fatalf("Expected update error, but didn't get one\nin: %s\nresponse: %v", string(objJSON), string(gotJSON))
	}
}

var embeddedResourceFixture = &apiextensionsv1beta1.CustomResourceDefinition{
	ObjectMeta: metav1.ObjectMeta{Name: "foos.tests.apiextensions.k8s.io"},
	Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
		Group:   "tests.apiextensions.k8s.io",
		Version: "v1beta1",
		Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
			Plural:   "foos",
			Singular: "foo",
			Kind:     "Foo",
			ListKind: "FooList",
		},
		Scope:                 apiextensionsv1beta1.ClusterScoped,
		PreserveUnknownFields: pointer.BoolPtr(true),
		Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
			Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
		},
	},
}

const (
	embeddedResourceSchema = `
type: object
properties:
  embedded:
    type: object
    x-kubernetes-embedded-resource: true
  noEmbeddedObject:
    type: object
  embeddedNested:
    type: object
    x-kubernetes-embedded-resource: true
    properties:
      embedded:
        type: object
        x-kubernetes-embedded-resource: true
`

	embeddedResourceInstance = `
kind: Foo
apiVersion: tests.apiextensions.k8s.io/v1beta1
metadata:
  name: foo
embedded:
  apiVersion: foo/v1
  kind: Foo
  metadata:
    name: foo
    unspecified: bar
noEmbeddedObject:
  apiVersion: foo/v1
  kind: Foo
  metadata:
    name: foo
    unspecified: bar
embeddedNested:
  apiVersion: foo/v1
  kind: Foo
  metadata:
    name: foo
    unspecified: bar
  embedded:
    apiVersion: foo/v1
    kind: Foo
    metadata:
      name: foo
      unspecified: bar
`

	expectedEmbeddedResourceInstance = `
kind: Foo
apiVersion: tests.apiextensions.k8s.io/v1beta1
embedded:
  apiVersion: foo/v1
  kind: Foo
  metadata:
    name: foo
noEmbeddedObject:
  apiVersion: foo/v1
  kind: Foo
  metadata:
    name: foo
    unspecified: bar
embeddedNested:
  apiVersion: foo/v1
  kind: Foo
  metadata:
    name: foo
  embedded:
    apiVersion: foo/v1
    kind: Foo
    metadata:
      name: foo
`

	wronglyTypedEmbeddedResourceInstance = `
kind: Foo
apiVersion: tests.apiextensions.k8s.io/v1beta1
metadata:
  name: invalid
embedded:
  apiVersion: foo/v1
  kind: Foo
  metadata:
    name: instance
    namespace: 42
`

	invalidEmbeddedResourceInstance = `
kind: Foo
apiVersion: tests.apiextensions.k8s.io/v1beta1
metadata:
  name: invalid
embedded:
  apiVersion: foo/v1
  kind: "%"
  metadata:
    name: ..
embeddedNested:
  apiVersion: foo/v1
  kind: "%"
  metadata:
    name: ..
  embedded:
    apiVersion: foo/v1
    kind: "%"
    metadata:
      name: ..
`
)

func TestEmbeddedResources(t *testing.T) {
	tearDownFn, apiExtensionClient, dynamicClient, err := fixtures.StartDefaultServerWithClients(t)
	if err != nil {
		t.Fatal(err)
	}
	defer tearDownFn()

	crd := embeddedResourceFixture.DeepCopy()
	crd.Spec.Validation = &apiextensionsv1beta1.CustomResourceValidation{}
	if err := yaml.Unmarshal([]byte(embeddedResourceSchema), &crd.Spec.Validation.OpenAPIV3Schema); err != nil {
		t.Fatal(err)
	}

	crd, err = fixtures.CreateNewCustomResourceDefinition(crd, apiExtensionClient, dynamicClient)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Creating CR and expect 'unspecified' fields to be pruned inside ObjectMetas")
	fooClient := dynamicClient.Resource(schema.GroupVersionResource{crd.Spec.Group, crd.Spec.Version, crd.Spec.Names.Plural})
	foo := &unstructured.Unstructured{}
	if err := yaml.Unmarshal([]byte(embeddedResourceInstance), &foo.Object); err != nil {
		t.Fatal(err)
	}
	foo, err = fooClient.Create(foo, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Unable to create CR: %v", err)
	}
	t.Logf("CR created: %#v", foo.UnstructuredContent())

	t.Logf("Checking that everything unknown inside ObjectMeta is gone")
	delete(foo.Object, "metadata")
	var expected map[string]interface{}
	if err := yaml.Unmarshal([]byte(expectedEmbeddedResourceInstance), &expected); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(expected, foo.Object) {
		t.Errorf("unexpected diff: %s", diff.ObjectDiff(expected, foo.Object))
	}

	t.Logf("Trying to create wrongly typed CR")
	invalid := &unstructured.Unstructured{}
	if err := yaml.Unmarshal([]byte(wronglyTypedEmbeddedResourceInstance), &invalid.Object); err != nil {
		t.Fatal(err)
	}
	_, err = fooClient.Create(invalid, metav1.CreateOptions{})
	if err == nil {
		t.Fatal("Expected creation to fail, but didn't")
	}
	t.Logf("Creation of wrongly typed object failed with: %v", err)

	for _, s := range []string{
		`embedded.metadata: Invalid value`,
	} {
		if !strings.Contains(err.Error(), s) {
			t.Errorf("missing error: %s", s)
		}
	}

	t.Logf("Trying to create invalid CR")
	wronglyTyped := &unstructured.Unstructured{}
	if err := yaml.Unmarshal([]byte(invalidEmbeddedResourceInstance), &wronglyTyped.Object); err != nil {
		t.Fatal(err)
	}
	_, err = fooClient.Create(wronglyTyped, metav1.CreateOptions{})
	if err == nil {
		t.Fatal("Expected creation to fail, but didn't")
	}
	t.Logf("Creation of invalid object failed with: %v", err)

	for _, s := range []string{
		`embedded.kind: Invalid value: "%"`,
		`embedded.metadata.name: Invalid value: ".."`,
		`embeddedNested.kind: Invalid value: "%"`,
		`embeddedNested.metadata.name: Invalid value: ".."`,
		`embeddedNested.embedded.kind: Invalid value: "%"`,
		`embeddedNested.embedded.metadata.name: Invalid value: ".."`,
	} {
		if !strings.Contains(err.Error(), s) {
			t.Errorf("missing error: %s", s)
		}
	}
}
