/*
Copyright 2021 The Kubernetes Authors.

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

package test

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/pod-security-admission/api"
	"k8s.io/pod-security-admission/policy"
)

// Options hold configuration for running integration tests against an existing server.
type Options struct {
	// ClientConfig is a client configuration with sufficient permission to create, update, and delete
	// namespaces, pods, and pod-template-containing objects.
	// Required.
	ClientConfig *rest.Config

	// CreateNamespace is an optional stub for creating a namespace with the given name and labels.
	// Returning an error fails the test.
	// If nil, DefaultCreateNamespace is used.
	CreateNamespace func(client kubernetes.Interface, name string, labels map[string]string) (*corev1.Namespace, error)

	// These are the check ids/starting versions to exercise.
	// If unset, policy.DefaultChecks() are used.
	Checks []policy.Check

	// ExemptClient is an optional client interface to exercise behavior of an exempt client.
	ExemptClient kubernetes.Interface
	// ExemptNamespaces are optional namespaces not expected to have PodSecurity controls enforced.
	ExemptNamespaces []string
	// ExemptRuntimeClasses are optional runtimeclasses not expected to have PodSecurity controls enforced.
	ExemptRuntimeClasses []string
}

func toJSON(pod *corev1.Pod) string {
	data, _ := json.Marshal(pod)
	return string(data)
}

// checksForLevelAndVersion returns the set of check IDs that apply when evaluating the given level and version.
// checks are assumed to be well-formed and valid to pass to policy.NewEvaluator().
// level must be api.LevelRestricted or api.LevelBaseline
func checksForLevelAndVersion(checks []policy.Check, level api.Level, version api.Version) ([]string, error) {
	retval := []string{}
	for _, check := range checks {
		if !version.Older(check.Versions[0].MinimumVersion) && (level == check.Level || level == api.LevelRestricted) {
			retval = append(retval, check.ID)
		}
	}
	return retval, nil
}

// maxMinorVersionToTest returns the maximum minor version to exercise for a given set of checks.
// checks are assumed to be well-formed and valid to pass to policy.NewEvaluator().
func maxMinorVersionToTest(checks []policy.Check) (int, error) {
	// start with the release under development (1.22 at time of writing).
	// this can be incremented to the current version whenever is convenient.
	maxTestMinor := 22
	for _, check := range checks {
		lastCheckVersion := check.Versions[len(check.Versions)-1].MinimumVersion
		if lastCheckVersion.Major() != 1 {
			return 0, fmt.Errorf("expected major version 1, got %d", lastCheckVersion.Major())
		}
		if lastCheckVersion.Minor() > maxTestMinor {
			maxTestMinor = lastCheckVersion.Minor()
		}
	}
	return maxTestMinor, nil
}

type testWarningHandler struct {
	lock     sync.Mutex
	warnings []string
}

func (t *testWarningHandler) HandleWarningHeader(code int, agent string, warning string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.warnings = append(t.warnings, warning)
}
func (t *testWarningHandler) FlushWarnings() []string {
	t.lock.Lock()
	defer t.lock.Unlock()
	warnings := t.warnings
	t.warnings = nil
	return warnings
}

// and ensures pod fixtures expected to pass and fail against that level/version work as expected.
func Run(t *testing.T, opts Options) {
	warningHandler := &testWarningHandler{}

	configCopy := rest.CopyConfig(opts.ClientConfig)
	configCopy.WarningHandler = warningHandler
	client, err := kubernetes.NewForConfig(configCopy)
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	if opts.CreateNamespace == nil {
		opts.CreateNamespace = DefaultCreateNamespace
	}
	if len(opts.Checks) == 0 {
		opts.Checks = policy.DefaultChecks()
	}
	_, err = policy.NewEvaluator(opts.Checks)
	if err != nil {
		t.Fatalf("invalid checks: %v", err)
	}
	maxMinor, err := maxMinorVersionToTest(opts.Checks)
	if err != nil {
		t.Fatalf("invalid checks: %v", err)
	}

	for _, level := range []api.Level{api.LevelBaseline, api.LevelRestricted} {
		for minor := 0; minor <= maxMinor; minor++ {
			version := api.MajorMinorVersion(1, minor)

			// create test name
			ns := fmt.Sprintf("podsecurity-%s-1-%d", level, minor)

			// create namespace
			_, err := opts.CreateNamespace(client, ns, map[string]string{
				api.EnforceLevelLabel:   string(level),
				api.EnforceVersionLabel: fmt.Sprintf("v1.%d", minor),
				api.WarnLevelLabel:      string(level),
				api.WarnVersionLabel:    fmt.Sprintf("v1.%d", minor),
			})
			if err != nil {
				t.Errorf("failed creating namespace %s: %v", ns, err)
				continue
			}
			t.Cleanup(func() {
				client.CoreV1().Namespaces().Delete(context.Background(), ns, metav1.DeleteOptions{})
			})

			// create service account (to allow pod to pass serviceaccount admission)
			sa, err := client.CoreV1().ServiceAccounts(ns).Create(
				context.Background(),
				&corev1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: "default"}},
				metav1.CreateOptions{},
			)
			if err != nil && !apierrors.IsAlreadyExists(err) {
				t.Errorf("failed creating serviceaccount %s: %v", ns, err)
				continue
			}
			t.Cleanup(func() {
				client.CoreV1().ServiceAccounts(ns).Delete(context.Background(), sa.Name, metav1.DeleteOptions{})
			})

			// create pod
			createPod := func(t *testing.T, i int, pod *corev1.Pod, expectSuccess bool, expectErrorSubstring string) {
				t.Helper()
				// avoid mutating original pod fixture
				pod = pod.DeepCopy()
				// assign pod name and serviceaccount
				pod.Name = "test"
				pod.Spec.ServiceAccountName = "default"
				// dry-run create
				_, err := client.CoreV1().Pods(ns).Create(context.Background(), pod, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
				if !expectSuccess {
					if err == nil {
						t.Errorf("%d: expected error creating %s, got none", i, toJSON(pod))
						return
					}
					if strings.Contains(err.Error(), policy.UnknownForbiddenReason) {
						t.Errorf("%d: unexpected unknown forbidden reason creating %s: %v", i, toJSON(pod), err)
						return
					}
					if !strings.Contains(err.Error(), expectErrorSubstring) {
						t.Errorf("%d: expected error with substring %q, got %v", i, expectErrorSubstring, err)
						return
					}
				}
				if expectSuccess && err != nil {
					t.Errorf("%d: unexpected error creating %s: %v", i, toJSON(pod), err)
				}
			}

			// create controller
			createController := func(t *testing.T, i int, pod *corev1.Pod, expectSuccess bool, expectErrorSubstring string) {
				t.Helper()
				// avoid mutating original pod fixture
				pod = pod.DeepCopy()
				if pod.Labels == nil {
					pod.Labels = map[string]string{}
				}
				pod.Labels["test"] = "true"

				warningHandler.FlushWarnings()
				// dry-run create
				deployment := &appsv1.Deployment{
					ObjectMeta: metav1.ObjectMeta{Name: "test"},
					Spec: appsv1.DeploymentSpec{
						Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"test": "true"}},
						Template: corev1.PodTemplateSpec{
							ObjectMeta: pod.ObjectMeta,
							Spec:       pod.Spec,
						},
					},
				}
				_, err := client.AppsV1().Deployments(ns).Create(context.Background(), deployment, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
				if err != nil {
					t.Errorf("%d: unexpected error creating controller with %s: %v", i, toJSON(pod), err)
					return
				}
				warningText := strings.Join(warningHandler.FlushWarnings(), "; ")
				if !expectSuccess {
					if len(warningText) == 0 {
						t.Errorf("%d: expected warnings creating %s, got none", i, toJSON(pod))
						return
					}
					if strings.Contains(warningText, policy.UnknownForbiddenReason) {
						t.Errorf("%d: unexpected unknown forbidden reason creating %s: %v", i, toJSON(pod), warningText)
						return
					}
					if !strings.Contains(warningText, expectErrorSubstring) {
						t.Errorf("%d: expected warning with substring %q, got %v", i, expectErrorSubstring, warningText)
						return
					}
				}
				if expectSuccess && len(warningText) > 0 {
					t.Errorf("%d: unexpected warning creating %s: %v", i, toJSON(pod), warningText)
				}
			}

			minimalValidPod, err := getMinimalValidPod(level, version)
			if err != nil {
				t.Fatal(err)
			}
			t.Run(ns+"_pass_base", func(t *testing.T) {
				createPod(t, 0, minimalValidPod.DeepCopy(), true, "")
				createController(t, 0, minimalValidPod.DeepCopy(), true, "")
			})

			checkIDs, err := checksForLevelAndVersion(opts.Checks, level, version)
			if err != nil {
				t.Fatal(err)
			}
			if len(checkIDs) == 0 {
				t.Fatal(fmt.Errorf("no checks registered for %s/1.%d", level, minor))
			}
			for _, checkID := range checkIDs {
				checkData, err := getFixtures(fixtureKey{level: level, version: version, check: checkID})
				if err != nil {
					t.Fatal(err)
				}

				t.Run(ns+"_pass_"+checkID, func(t *testing.T) {
					for i, pod := range checkData.pass {
						createPod(t, i, pod, true, "")
						createController(t, i, pod, true, "")
					}
				})
				t.Run(ns+"_fail_"+checkID, func(t *testing.T) {
					for i, pod := range checkData.fail {
						createPod(t, i, pod, false, checkData.expectErrorSubstring)
						createController(t, i, pod, false, checkData.expectErrorSubstring)
					}
				})
			}
		}
	}
}

func DefaultCreateNamespace(client kubernetes.Interface, name string, labels map[string]string) (*corev1.Namespace, error) {
	return client.CoreV1().Namespaces().Create(
		context.Background(),
		&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{Name: name, Labels: labels},
		},
		metav1.CreateOptions{},
	)
}
