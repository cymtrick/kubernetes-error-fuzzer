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

package webhook

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync/atomic"
	"testing"

	"k8s.io/api/admission/v1alpha1"
	registrationv1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	api "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/client-go/rest"
)

type fakeHookSource struct {
	hooks []registrationv1alpha1.Webhook
	err   error
}

func (f *fakeHookSource) Webhooks() (*registrationv1alpha1.ValidatingWebhookConfiguration, error) {
	if f.err != nil {
		return nil, f.err
	}
	return &registrationv1alpha1.ValidatingWebhookConfiguration{Webhooks: f.hooks}, nil
}

func (f *fakeHookSource) Run(stopCh <-chan struct{}) {}

type fakeServiceResolver struct {
	base url.URL
}

func (f fakeServiceResolver) ResolveEndpoint(namespace, name string) (*url.URL, error) {
	if namespace == "failResolve" {
		return nil, fmt.Errorf("couldn't resolve service location")
	}
	u := f.base
	return &u, nil
}

// TestAdmit tests that GenericAdmissionWebhook#Admit works as expected
func TestAdmit(t *testing.T) {
	scheme := runtime.NewScheme()
	v1alpha1.AddToScheme(scheme)
	api.AddToScheme(scheme)

	testServer := newTestServer(t)
	testServer.StartTLS()
	defer testServer.Close()
	serverURL, err := url.ParseRequestURI(testServer.URL)
	if err != nil {
		t.Fatalf("this should never happen? %v", err)
	}
	wh, err := NewGenericAdmissionWebhook(nil)
	if err != nil {
		t.Fatal(err)
	}
	wh.authInfoResolver = newFakeAuthenticationInfoResolver()
	wh.serviceResolver = fakeServiceResolver{base: *serverURL}
	wh.SetScheme(scheme)

	// Set up a test object for the call
	kind := api.SchemeGroupVersion.WithKind("Pod")
	name := "my-pod"
	namespace := "webhook-test"
	object := api.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"pod.name": name,
			},
			Name:      name,
			Namespace: namespace,
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
	}
	oldObject := api.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: namespace},
	}
	operation := admission.Update
	resource := api.Resource("pods").WithVersion("v1")
	subResource := ""
	userInfo := user.DefaultInfo{
		Name: "webhook-test",
		UID:  "webhook-test",
	}

	type test struct {
		hookSource    fakeHookSource
		path          string
		expectAllow   bool
		errorContains string
	}

	policyFail := registrationv1alpha1.Fail
	policyIgnore := registrationv1alpha1.Ignore

	table := map[string]test{
		"no match": {
			hookSource: fakeHookSource{
				hooks: []registrationv1alpha1.Webhook{{
					Name:         "nomatch",
					ClientConfig: newFakeHookClientConfig("disallow"),
					Rules: []registrationv1alpha1.RuleWithOperations{{
						Operations: []registrationv1alpha1.OperationType{registrationv1alpha1.Create},
					}},
				}},
			},
			expectAllow: true,
		},
		"match & allow": {
			hookSource: fakeHookSource{
				hooks: []registrationv1alpha1.Webhook{{
					Name:         "allow",
					ClientConfig: newFakeHookClientConfig("allow"),
					Rules:        newMatchEverythingRules(),
				}},
			},
			expectAllow: true,
		},
		"match & disallow": {
			hookSource: fakeHookSource{
				hooks: []registrationv1alpha1.Webhook{{
					Name:         "disallow",
					ClientConfig: newFakeHookClientConfig("disallow"),
					Rules:        newMatchEverythingRules(),
				}},
			},
			errorContains: "without explanation",
		},
		"match & disallow ii": {
			hookSource: fakeHookSource{
				hooks: []registrationv1alpha1.Webhook{{
					Name:         "disallowReason",
					ClientConfig: newFakeHookClientConfig("disallowReason"),
					Rules:        newMatchEverythingRules(),
				}},
			},
			errorContains: "you shall not pass",
		},
		"match & fail (but allow because fail open)": {
			hookSource: fakeHookSource{
				hooks: []registrationv1alpha1.Webhook{{
					Name:          "internalErr A",
					ClientConfig:  newFakeHookClientConfig("internalErr"),
					Rules:         newMatchEverythingRules(),
					FailurePolicy: &policyIgnore,
				}, {
					Name:          "internalErr B",
					ClientConfig:  newFakeHookClientConfig("internalErr"),
					Rules:         newMatchEverythingRules(),
					FailurePolicy: &policyIgnore,
				}, {
					Name:          "internalErr C",
					ClientConfig:  newFakeHookClientConfig("internalErr"),
					Rules:         newMatchEverythingRules(),
					FailurePolicy: &policyIgnore,
				}},
			},
			expectAllow: true,
		},
		"match & fail (but disallow because fail closed on nil)": {
			hookSource: fakeHookSource{
				hooks: []registrationv1alpha1.Webhook{{
					Name:         "internalErr A",
					ClientConfig: newFakeHookClientConfig("internalErr"),
					Rules:        newMatchEverythingRules(),
				}, {
					Name:         "internalErr B",
					ClientConfig: newFakeHookClientConfig("internalErr"),
					Rules:        newMatchEverythingRules(),
				}, {
					Name:         "internalErr C",
					ClientConfig: newFakeHookClientConfig("internalErr"),
					Rules:        newMatchEverythingRules(),
				}},
			},
			expectAllow: false,
		},
		"match & fail (but fail because fail closed)": {
			hookSource: fakeHookSource{
				hooks: []registrationv1alpha1.Webhook{{
					Name:          "internalErr A",
					ClientConfig:  newFakeHookClientConfig("internalErr"),
					Rules:         newMatchEverythingRules(),
					FailurePolicy: &policyFail,
				}, {
					Name:          "internalErr B",
					ClientConfig:  newFakeHookClientConfig("internalErr"),
					Rules:         newMatchEverythingRules(),
					FailurePolicy: &policyFail,
				}, {
					Name:          "internalErr C",
					ClientConfig:  newFakeHookClientConfig("internalErr"),
					Rules:         newMatchEverythingRules(),
					FailurePolicy: &policyFail,
				}},
			},
			expectAllow: false,
		},
	}

	for name, tt := range table {
		t.Run(name, func(t *testing.T) {
			wh.hookSource = &tt.hookSource

			err = wh.Admit(admission.NewAttributesRecord(&object, &oldObject, kind, namespace, name, resource, subResource, operation, &userInfo))
			if tt.expectAllow != (err == nil) {
				t.Errorf("expected allowed=%v, but got err=%v", tt.expectAllow, err)
			}
			// ErrWebhookRejected is not an error for our purposes
			if tt.errorContains != "" {
				if err == nil || !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf(" expected an error saying %q, but got %v", tt.errorContains, err)
				}
			}
			if _, isStatusErr := err.(*apierrors.StatusError); err != nil && !isStatusErr {
				t.Errorf("%s: expected a StatusError, got %T", name, err)
			}
		})
	}
}

// TestAdmitCachedClient tests that GenericAdmissionWebhook#Admit should cache restClient
func TestAdmitCachedClient(t *testing.T) {
	scheme := runtime.NewScheme()
	v1alpha1.AddToScheme(scheme)
	api.AddToScheme(scheme)

	testServer := newTestServer(t)
	testServer.StartTLS()
	defer testServer.Close()
	serverURL, err := url.ParseRequestURI(testServer.URL)
	if err != nil {
		t.Fatalf("this should never happen? %v", err)
	}
	wh, err := NewGenericAdmissionWebhook(nil)
	if err != nil {
		t.Fatal(err)
	}
	wh.authInfoResolver = newFakeAuthenticationInfoResolver()
	wh.serviceResolver = fakeServiceResolver{base: *serverURL}
	wh.SetScheme(scheme)

	// Set up a test object for the call
	kind := api.SchemeGroupVersion.WithKind("Pod")
	name := "my-pod"
	namespace := "webhook-test"
	object := api.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"pod.name": name,
			},
			Name:      name,
			Namespace: namespace,
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
	}
	oldObject := api.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: namespace},
	}
	operation := admission.Update
	resource := api.Resource("pods").WithVersion("v1")
	subResource := ""
	userInfo := user.DefaultInfo{
		Name: "webhook-test",
		UID:  "webhook-test",
	}

	type test struct {
		name        string
		hookSource  fakeHookSource
		expectAllow bool
		expectCache bool
	}

	policyIgnore := registrationv1alpha1.Ignore
	cases := []test{
		{
			name: "cache 1",
			hookSource: fakeHookSource{
				hooks: []registrationv1alpha1.Webhook{{
					Name:          "cache1",
					ClientConfig:  newFakeHookClientConfig("allow"),
					Rules:         newMatchEverythingRules(),
					FailurePolicy: &policyIgnore,
				}},
			},
			expectAllow: true,
			expectCache: true,
		},
		{
			name: "cache 2",
			hookSource: fakeHookSource{
				hooks: []registrationv1alpha1.Webhook{{
					Name:          "cache2",
					ClientConfig:  newFakeHookClientConfig("internalErr"),
					Rules:         newMatchEverythingRules(),
					FailurePolicy: &policyIgnore,
				}},
			},
			expectAllow: true,
			expectCache: true,
		},
		{
			name: "cache 3",
			hookSource: fakeHookSource{
				hooks: []registrationv1alpha1.Webhook{{
					Name:          "cache3",
					ClientConfig:  newFakeHookClientConfig("allow"),
					Rules:         newMatchEverythingRules(),
					FailurePolicy: &policyIgnore,
				}},
			},
			expectAllow: true,
			expectCache: false,
		},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			wh.hookSource = &testcase.hookSource
			wh.authInfoResolver.(*fakeAuthenticationInfoResolver).cachedCount = 0

			err = wh.Admit(admission.NewAttributesRecord(&object, &oldObject, kind, namespace, testcase.name, resource, subResource, operation, &userInfo))
			if testcase.expectAllow != (err == nil) {
				t.Errorf("expected allowed=%v, but got err=%v", testcase.expectAllow, err)
			}

			if testcase.expectCache && wh.authInfoResolver.(*fakeAuthenticationInfoResolver).cachedCount != 1 {
				t.Errorf("expected cacheclient, but got none")
			}

			if !testcase.expectCache && wh.authInfoResolver.(*fakeAuthenticationInfoResolver).cachedCount != 0 {
				t.Errorf("expected not cacheclient, but got cache")
			}
		})
	}

}

func newTestServer(t *testing.T) *httptest.Server {
	// Create the test webhook server
	sCert, err := tls.X509KeyPair(serverCert, serverKey)
	if err != nil {
		t.Fatal(err)
	}
	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(caCert)
	testServer := httptest.NewUnstartedServer(http.HandlerFunc(webhookHandler))
	testServer.TLS = &tls.Config{
		Certificates: []tls.Certificate{sCert},
		ClientCAs:    rootCAs,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	return testServer
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got req: %v\n", r.URL.Path)
	switch r.URL.Path {
	case "/internalErr":
		http.Error(w, "webhook internal server error", http.StatusInternalServerError)
		return
	case "/invalidReq":
		w.WriteHeader(http.StatusSwitchingProtocols)
		w.Write([]byte("webhook invalid request"))
		return
	case "/invalidResp":
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("webhook invalid response"))
	case "/disallow":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&v1alpha1.AdmissionReview{
			Status: v1alpha1.AdmissionReviewStatus{
				Allowed: false,
			},
		})
	case "/disallowReason":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&v1alpha1.AdmissionReview{
			Status: v1alpha1.AdmissionReviewStatus{
				Allowed: false,
				Result: &metav1.Status{
					Message: "you shall not pass",
				},
			},
		})
	case "/allow":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&v1alpha1.AdmissionReview{
			Status: v1alpha1.AdmissionReviewStatus{
				Allowed: true,
			},
		})
	default:
		http.NotFound(w, r)
	}
}

func newFakeAuthenticationInfoResolver() *fakeAuthenticationInfoResolver {
	return &fakeAuthenticationInfoResolver{
		restConfig: &rest.Config{
			TLSClientConfig: rest.TLSClientConfig{
				CAData:   caCert,
				CertData: clientCert,
				KeyData:  clientKey,
			},
		},
	}
}

type fakeAuthenticationInfoResolver struct {
	restConfig  *rest.Config
	cachedCount int32
}

func (c *fakeAuthenticationInfoResolver) ClientConfigFor(server string) (*rest.Config, error) {
	atomic.AddInt32(&c.cachedCount, 1)
	return c.restConfig, nil
}

func TestToStatusErr(t *testing.T) {
	hookName := "foo"
	deniedBy := fmt.Sprintf("admission webhook %q denied the request", hookName)
	tests := []struct {
		name          string
		result        *metav1.Status
		expectedError string
	}{
		{
			"nil result",
			nil,
			deniedBy + " without explanation",
		},
		{
			"only message",
			&metav1.Status{
				Message: "you shall not pass",
			},
			deniedBy + ": you shall not pass",
		},
		{
			"only reason",
			&metav1.Status{
				Reason: metav1.StatusReasonForbidden,
			},
			deniedBy + ": Forbidden",
		},
		{
			"message and reason",
			&metav1.Status{
				Message: "you shall not pass",
				Reason:  metav1.StatusReasonForbidden,
			},
			deniedBy + ": you shall not pass",
		},
		{
			"no message, no reason",
			&metav1.Status{},
			deniedBy + " without explanation",
		},
	}
	for _, test := range tests {
		err := toStatusErr(hookName, test.result)
		if err == nil || err.Error() != test.expectedError {
			t.Errorf("%s: expected an error saying %q, but got %v", test.name, test.expectedError, err)
		}
	}
}

func newFakeHookClientConfig(urlPath string) registrationv1alpha1.WebhookClientConfig {
	return registrationv1alpha1.WebhookClientConfig{
		Service: registrationv1alpha1.ServiceReference{
			Name:      "webhook-test",
			Namespace: "default",
		},
		URLPath:  urlPath,
		CABundle: caCert,
	}
}

func newMatchEverythingRules() []registrationv1alpha1.RuleWithOperations {
	return []registrationv1alpha1.RuleWithOperations{{
		Operations: []registrationv1alpha1.OperationType{registrationv1alpha1.OperationAll},
		Rule: registrationv1alpha1.Rule{
			APIGroups:   []string{"*"},
			APIVersions: []string{"*"},
			Resources:   []string{"*/*"},
		},
	}}
}
