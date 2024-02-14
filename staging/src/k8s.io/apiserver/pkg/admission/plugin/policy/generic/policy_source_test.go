/*
Copyright 2024 The Kubernetes Authors.

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

package generic_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apiserver/pkg/admission/plugin/policy/generic"
	"k8s.io/apiserver/pkg/admission/plugin/policy/matching"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/client-go/tools/cache"
)

func makeTestDispatcher(authorizer.Authorizer, *matching.Matcher) generic.Dispatcher[generic.PolicyHook[*FakePolicy, *FakeBinding, generic.Evaluator]] {
	return nil
}

func TestPolicySourceHasSyncedEmpty(t *testing.T) {
	testContext, testCancel, err := generic.NewPolicyTestContext(
		func(fp *FakePolicy) generic.PolicyAccessor { return fp },
		func(fb *FakeBinding) generic.BindingAccessor { return fb },
		func(fp *FakePolicy) generic.Evaluator { return nil },
		makeTestDispatcher,
		nil,
		nil,
	)
	require.NoError(t, err)
	defer testCancel()
	require.NoError(t, testContext.Start())

	// Should be able to wait for cache sync
	require.True(t, cache.WaitForCacheSync(testContext.Done(), testContext.Source.HasSynced), "cache should sync after informer running")
}

func TestPolicySourceHasSyncedInitialList(t *testing.T) {
	// Create a list of fake policies
	initialObjects := []runtime.Object{
		&FakePolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name: "policy1",
			},
		},
		&FakeBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name: "binding1",
			},
			PolicyName: "policy1",
		},
	}

	testContext, testCancel, err := generic.NewPolicyTestContext(
		func(fp *FakePolicy) generic.PolicyAccessor { return fp },
		func(fb *FakeBinding) generic.BindingAccessor { return fb },
		func(fp *FakePolicy) generic.Evaluator { return nil },
		makeTestDispatcher,
		initialObjects,
		nil,
	)
	require.NoError(t, err)
	defer testCancel()
	require.NoError(t, testContext.Start())
	// Should be able to wait for cache sync
	require.Len(t, testContext.Source.Hooks(), 1, "should have one policy")
	require.NoError(t, testContext.UpdateAndWait(
		&FakePolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name: "policy2",
			},
		},
		&FakeBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name: "binding2",
			},
			PolicyName: "policy2",
		},
	))
	require.Len(t, testContext.Source.Hooks(), 2, "should have two policies")
	require.NoError(t, testContext.UpdateAndWait(
		&FakePolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name: "policy3",
			},
		},
		&FakeBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name: "binding3",
			},
			PolicyName: "policy3",
		},
		&FakePolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name: "policy2",
			},
			ParamKind: &schema.GroupVersionKind{
				Group:   "policy.example.com",
				Version: "v1",
				Kind:    "FakeParam",
			},
		},
	))
	require.Len(t, testContext.Source.Hooks(), 3, "should have 3 policies")

}

func TestPolicySourceBindsToPolicies(t *testing.T) {
	// Create a list of fake policies
	initialObjects := []runtime.Object{
		&FakePolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name: "policy1",
			},
		},
		&FakeBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name: "binding1",
			},
			PolicyName: "policy1",
		},
	}

	testContext, testCancel, err := generic.NewPolicyTestContext(
		func(fp *FakePolicy) generic.PolicyAccessor { return fp },
		func(fb *FakeBinding) generic.BindingAccessor { return fb },
		func(fp *FakePolicy) generic.Evaluator { return nil },
		makeTestDispatcher,
		initialObjects,
		nil,
	)
	require.NoError(t, err)
	defer testCancel()
	require.NoError(t, testContext.Start())

	require.Len(t, testContext.Source.Hooks(), 1, "should have one policy")
	require.Len(t, testContext.Source.Hooks()[0].Bindings, 1, "should have one binding")
	require.Equal(t, "binding1", testContext.Source.Hooks()[0].Bindings[0].GetName(), "should have one binding")

	// Change the binding to another policy (policies without bindings should
	// be ignored, so it should remove the first
	require.NoError(t, testContext.UpdateAndWait(
		&FakePolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name: "policy2",
			},
		},
		&FakeBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name: "binding1",
			},
			PolicyName: "policy2",
		}))
	require.Len(t, testContext.Source.Hooks(), 1, "should have one policy")
	require.Equal(t, "policy2", testContext.Source.Hooks()[0].Policy.GetName(), "policy name should be policy2")
	require.Len(t, testContext.Source.Hooks()[0].Bindings, 1, "should have one binding")
	require.Equal(t, "binding1", testContext.Source.Hooks()[0].Bindings[0].GetName(), "binding name should be binding1")

}

type FakePolicy struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	ParamKind *schema.GroupVersionKind
}

var _ generic.PolicyAccessor = &FakePolicy{}

type FakeBinding struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	PolicyName string
}

var _ generic.BindingAccessor = &FakeBinding{}

func (fp *FakePolicy) GetName() string {
	return fp.Name
}

func (fp *FakePolicy) GetNamespace() string {
	return fp.Namespace
}

func (fp *FakePolicy) GetParamKind() *schema.GroupVersionKind {
	return fp.ParamKind
}

func (fb *FakeBinding) GetName() string {
	return fb.Name
}

func (fb *FakeBinding) GetNamespace() string {
	return fb.Namespace
}

func (fb *FakeBinding) GetPolicyName() types.NamespacedName {
	return types.NamespacedName{
		Name: fb.PolicyName,
	}
}

func (fp *FakePolicy) DeepCopyObject() runtime.Object {
	// totally fudged deepcopy
	newFP := &FakePolicy{}
	*newFP = *fp
	return newFP
}

func (fb *FakeBinding) DeepCopyObject() runtime.Object {
	// totally fudged deepcopy
	newFB := &FakeBinding{}
	*newFB = *fb
	return newFB
}
