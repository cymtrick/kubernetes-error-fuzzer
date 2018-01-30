/*
Copyright 2016 The Kubernetes Authors.

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

package preflight

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/renstrom/dedent"

	"net/http"
	"os"

	"k8s.io/apimachinery/pkg/util/sets"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/utils/exec"
)

var (
	externalEtcdRootCAFileContent = dedent.Dedent(`
		-----BEGIN CERTIFICATE-----
		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
		Z9JkQayesM6D7uwbQJXd5rgy
		-----END CERTIFICATE-----
	`)

	externalEtcdCertFileContent = dedent.Dedent(`
		-----BEGIN CERTIFICATE-----
		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
		XRMVQjawUTknbAgHuE7/rEX8c27WUA==
		-----END CERTIFICATE-----
	`)
	externalEtcdKeyFileContent = dedent.Dedent(`
		-----BEGIN RSA PRIVATE KEY-----
		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
		-----END RSA PRIVATE KEY-----
	`)
)

type preflightCheckTest struct {
	msg string
}

func (pfct preflightCheckTest) Name() string {
	return "preflightCheckTest"
}

func (pfct preflightCheckTest) Check() (warning, errors []error) {
	if pfct.msg == "warning" {
		return []error{fmt.Errorf("warning")}, nil
	}
	if pfct.msg != "" {
		return nil, []error{fmt.Errorf("fake error")}
	}
	return
}

func TestRunInitMasterChecks(t *testing.T) {
	var tests = []struct {
		name     string
		cfg      *kubeadmapi.MasterConfiguration
		expected bool
	}{
		{name: "Test valid advertised address",
			cfg: &kubeadmapi.MasterConfiguration{
				API: kubeadmapi.API{AdvertiseAddress: "foo"},
			},
			expected: false,
		},
		{
			name: "Test CA file exists if specfied",
			cfg: &kubeadmapi.MasterConfiguration{
				Etcd: kubeadmapi.Etcd{CAFile: "/foo"},
			},
			expected: false,
		},
		{
			name: "Test Cert file exists if specfied",
			cfg: &kubeadmapi.MasterConfiguration{
				Etcd: kubeadmapi.Etcd{CertFile: "/foo"},
			},
			expected: false,
		},
		{
			name: "Test Key file exists if specfied",
			cfg: &kubeadmapi.MasterConfiguration{
				Etcd: kubeadmapi.Etcd{CertFile: "/foo"},
			},
			expected: false,
		},
		{
			cfg: &kubeadmapi.MasterConfiguration{
				API: kubeadmapi.API{AdvertiseAddress: "2001:1234::1:15"},
			},
			expected: false,
		},
	}

	for _, rt := range tests {
		actual := RunInitMasterChecks(exec.New(), rt.cfg, sets.NewString())
		if (actual == nil) != rt.expected {
			t.Errorf(
				"failed RunInitMasterChecks:\n\texpected: %t\n\t  actual: %t\n\t error: %v",
				rt.expected,
				(actual == nil),
				actual,
			)
		}
	}
}

func TestRunJoinNodeChecks(t *testing.T) {
	var tests = []struct {
		cfg      *kubeadmapi.NodeConfiguration
		expected bool
	}{
		{
			cfg:      &kubeadmapi.NodeConfiguration{},
			expected: false,
		},
		{
			cfg: &kubeadmapi.NodeConfiguration{
				DiscoveryTokenAPIServers: []string{"192.168.1.15"},
			},
			expected: false,
		},
		{
			cfg: &kubeadmapi.NodeConfiguration{
				DiscoveryTokenAPIServers: []string{"2001:1234::1:15"},
			},
			expected: false,
		},
	}

	for _, rt := range tests {
		actual := RunJoinNodeChecks(exec.New(), rt.cfg, "", sets.NewString())
		if (actual == nil) != rt.expected {
			t.Errorf(
				"failed RunJoinNodeChecks:\n\texpected: %t\n\t  actual: %t",
				rt.expected,
				(actual != nil),
			)
		}
	}
}

func TestRunChecks(t *testing.T) {
	var tokenTest = []struct {
		p        []Checker
		expected bool
		output   string
	}{
		{[]Checker{}, true, ""},
		{[]Checker{preflightCheckTest{"warning"}}, true, "\t[WARNING preflightCheckTest]: warning\n"}, // should just print warning
		{[]Checker{preflightCheckTest{"error"}}, false, ""},
		{[]Checker{preflightCheckTest{"test"}}, false, ""},
		{[]Checker{DirAvailableCheck{Path: "/does/not/exist"}}, true, ""},
		{[]Checker{DirAvailableCheck{Path: "/"}}, false, ""},
		{[]Checker{FileAvailableCheck{Path: "/does/not/exist"}}, true, ""},
		{[]Checker{FileContentCheck{Path: "/does/not/exist"}}, false, ""},
		{[]Checker{FileContentCheck{Path: "/"}}, true, ""},
		{[]Checker{FileContentCheck{Path: "/", Content: []byte("does not exist")}}, false, ""},
		{[]Checker{InPathCheck{executable: "foobarbaz", exec: exec.New()}}, true, "\t[WARNING FileExisting-foobarbaz]: foobarbaz not found in system path\n"},
		{[]Checker{InPathCheck{executable: "foobarbaz", mandatory: true, exec: exec.New()}}, false, ""},
		{[]Checker{ExtraArgsCheck{
			APIServerExtraArgs:         map[string]string{"secure-port": "1234"},
			ControllerManagerExtraArgs: map[string]string{"use-service-account-credentials": "true"},
			SchedulerExtraArgs:         map[string]string{"leader-elect": "true"},
		}}, true, ""},
		{[]Checker{ExtraArgsCheck{
			APIServerExtraArgs: map[string]string{"secure-port": "foo"},
		}}, true, "\t[WARNING ExtraArgs]: kube-apiserver: failed to parse extra argument --secure-port=foo\n"},
		{[]Checker{ExtraArgsCheck{
			APIServerExtraArgs: map[string]string{"invalid-argument": "foo"},
		}}, true, "\t[WARNING ExtraArgs]: kube-apiserver: failed to parse extra argument --invalid-argument=foo\n"},
		{[]Checker{InPathCheck{executable: "foobar", mandatory: false, exec: exec.New(), suggestion: "install foobar"}}, true, "\t[WARNING FileExisting-foobar]: foobar not found in system path\nSuggestion: install foobar\n"},
	}
	for _, rt := range tokenTest {
		buf := new(bytes.Buffer)
		actual := RunChecks(rt.p, buf, sets.NewString())
		if (actual == nil) != rt.expected {
			t.Errorf(
				"failed RunChecks:\n\texpected: %t\n\t  actual: %t",
				rt.expected,
				(actual == nil),
			)
		}
		if buf.String() != rt.output {
			t.Errorf(
				"failed RunChecks:\n\texpected: %s\n\t  actual: %s",
				rt.output,
				buf.String(),
			)
		}
	}
}
func TestConfigRootCAs(t *testing.T) {
	f, err := ioutil.TempFile(os.TempDir(), "kubeadm-external-etcd-test-cafile")
	if err != nil {
		t.Errorf("failed configRootCAs:\n\texpected: succeed creating temp CA file\n\tactual:%v", err)
	}
	defer os.Remove(f.Name())
	if err := ioutil.WriteFile(f.Name(), []byte(externalEtcdRootCAFileContent), 0644); err != nil {
		t.Errorf("failed configRootCAs:\n\texpected: succeed writing contents to temp CA file %s\n\tactual:%v", f.Name(), err)
	}

	c := ExternalEtcdVersionCheck{Etcd: kubeadmapi.Etcd{CAFile: f.Name()}}

	config, err := c.configRootCAs(nil)
	if err != nil {
		t.Errorf(
			"failed configRootCAs:\n\texpected: has no error\n\tactual:%v",
			err,
		)
	}
	if config.RootCAs == nil {
		t.Errorf(
			"failed configRootCAs:\n\texpected: RootCAs not equal to nil\n\tactual:%v",
			config.RootCAs,
		)
	}
}
func TestConfigCertAndKey(t *testing.T) {
	certFile, err := ioutil.TempFile(os.TempDir(), "kubeadm-external-etcd-test-certfile")
	if err != nil {
		t.Errorf(
			"failed configCertAndKey:\n\texpected: succeed creating temp CertFile file\n\tactual:%v",
			err,
		)
	}
	defer os.Remove(certFile.Name())
	if err := ioutil.WriteFile(certFile.Name(), []byte(externalEtcdCertFileContent), 0644); err != nil {
		t.Errorf(
			"failed configCertAndKey:\n\texpected: succeed writing contents to temp CertFile file %s\n\tactual:%v",
			certFile.Name(),
			err,
		)
	}

	keyFile, err := ioutil.TempFile(os.TempDir(), "kubeadm-external-etcd-test-keyfile")
	if err != nil {
		t.Errorf(
			"failed configCertAndKey:\n\texpected: succeed creating temp KeyFile file\n\tactual:%v",
			err,
		)
	}
	defer os.Remove(keyFile.Name())
	if err := ioutil.WriteFile(keyFile.Name(), []byte(externalEtcdKeyFileContent), 0644); err != nil {
		t.Errorf(
			"failed configCertAndKey:\n\texpected: succeed writing contents to temp KeyFile file %s\n\tactual:%v",
			keyFile.Name(),
			err,
		)
	}
	c := ExternalEtcdVersionCheck{Etcd: kubeadmapi.Etcd{
		CertFile: certFile.Name(),
		KeyFile:  keyFile.Name(),
	}}

	config, err := c.configCertAndKey(nil)
	if err != nil {
		t.Errorf(
			"failed configCertAndKey:\n\texpected: has no error\n\tactual:%v",
			err,
		)
	}
	if config.Certificates == nil {
		t.Errorf(
			"failed configCertAndKey:\n\texpected: Certificates not equal to nil\n\tactual:%v",
			config.Certificates,
		)
	}
}

func TestKubernetesVersionCheck(t *testing.T) {
	var tests = []struct {
		check          KubernetesVersionCheck
		expectWarnings bool
	}{
		{
			check: KubernetesVersionCheck{
				KubeadmVersion:    "v1.6.6", //Same version
				KubernetesVersion: "v1.6.6",
			},
			expectWarnings: false,
		},
		{
			check: KubernetesVersionCheck{
				KubeadmVersion:    "v1.6.6", //KubernetesVersion version older than KubeadmVersion
				KubernetesVersion: "v1.5.5",
			},
			expectWarnings: false,
		},
		{
			check: KubernetesVersionCheck{
				KubeadmVersion:    "v1.6.6", //KubernetesVersion newer than KubeadmVersion, within the same minor release (new patch)
				KubernetesVersion: "v1.6.7",
			},
			expectWarnings: false,
		},
		{
			check: KubernetesVersionCheck{
				KubeadmVersion:    "v1.6.6", //KubernetesVersion newer than KubeadmVersion, in a different minor/in pre-release
				KubernetesVersion: "v1.7.0-alpha.0",
			},
			expectWarnings: true,
		},
		{
			check: KubernetesVersionCheck{
				KubeadmVersion:    "v1.6.6", //KubernetesVersion newer than KubeadmVersion, in a different minor/stable
				KubernetesVersion: "v1.7.0",
			},
			expectWarnings: true,
		},
		{
			check: KubernetesVersionCheck{
				KubeadmVersion:    "v0.0.0", //"super-custom" builds - Skip this check
				KubernetesVersion: "v1.7.0",
			},
			expectWarnings: false,
		},
	}

	for _, rt := range tests {
		warning, _ := rt.check.Check()
		if (warning != nil) != rt.expectWarnings {
			t.Errorf(
				"failed KubernetesVersionCheck:\n\texpected: %t\n\t  actual: %t (KubeadmVersion:%s, KubernetesVersion: %s)",
				rt.expectWarnings,
				(warning != nil),
				rt.check.KubeadmVersion,
				rt.check.KubernetesVersion,
			)
		}
	}
}

func TestHTTPProxyCIDRCheck(t *testing.T) {
	var tests = []struct {
		check          HTTPProxyCIDRCheck
		expectWarnings bool
	}{
		{
			check: HTTPProxyCIDRCheck{
				Proto: "https",
				CIDR:  "127.0.0.0/8",
			}, // Loopback addresses never should produce proxy warnings
			expectWarnings: false,
		},
		{
			check: HTTPProxyCIDRCheck{
				Proto: "https",
				CIDR:  "10.96.0.0/12",
			}, // Expected to be accessed directly, we set NO_PROXY to 10.0.0.0/8
			expectWarnings: false,
		},
		{
			check: HTTPProxyCIDRCheck{
				Proto: "https",
				CIDR:  "192.168.0.0/16",
			}, // Expected to go via proxy as this range is not listed in NO_PROXY
			expectWarnings: true,
		},
		{
			check: HTTPProxyCIDRCheck{
				Proto: "https",
				CIDR:  "2001:db8::/56",
			}, // Expected to be accessed directly, part of 2001:db8::/48 in NO_PROXY
			expectWarnings: false,
		},
		{
			check: HTTPProxyCIDRCheck{
				Proto: "https",
				CIDR:  "2001:db8:1::/56",
			}, // Expected to go via proxy, range is not in 2001:db8::/48
			expectWarnings: true,
		},
	}

	// Save current content of *_proxy and *_PROXY variables.
	savedEnv := resetProxyEnv()
	defer restoreEnv(savedEnv)
	t.Log("Saved environment: ", savedEnv)

	os.Setenv("HTTP_PROXY", "http://proxy.example.com:3128")
	os.Setenv("NO_PROXY", "example.com,10.0.0.0/8,2001:db8::/48")

	// Check if we can reliably execute tests:
	// ProxyFromEnvironment caches the *_proxy environment variables and
	// if ProxyFromEnvironment already executed before our test with empty
	// HTTP_PROXY it will make these tests return false positive failures
	req, err := http.NewRequest("GET", "http://host.fake.tld/", nil)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	proxy, err := http.ProxyFromEnvironment(req)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if proxy == nil {
		t.Skip("test skipped as ProxyFromEnvironment already initialized in environment without defined HTTP proxy")
	}
	t.Log("http.ProxyFromEnvironment is usable, continue executing test")

	for _, rt := range tests {
		warning, _ := rt.check.Check()
		if (warning != nil) != rt.expectWarnings {
			t.Errorf(
				"failed HTTPProxyCIDRCheck:\n\texpected: %t\n\t  actual: %t (CIDR:%s). Warnings: %v",
				rt.expectWarnings,
				(warning != nil),
				rt.check.CIDR,
				warning,
			)
		}
	}
}

// resetProxyEnv is helper function that unsets all *_proxy variables
// and return previously set values as map. This can be used to restore
// original state of the environment.
func resetProxyEnv() map[string]string {
	savedEnv := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if strings.HasSuffix(strings.ToLower(pair[0]), "_proxy") {
			savedEnv[pair[0]] = pair[1]
			os.Unsetenv(pair[0])
		}
	}
	return savedEnv
}

// restoreEnv is helper function to restores values
// of environment variables from saved state in the map
func restoreEnv(e map[string]string) {
	for k, v := range e {
		os.Setenv(k, v)
	}
}

func TestKubeletVersionCheck(t *testing.T) {
	type T struct {
		kubeletVersion string
		k8sVersion     string
		expectErrors   bool
		expectWarnings bool
	}

	cases := []T{
		{"v1.10.2", "", false, false},              // check minimally supported version when there is no information about control plane
		{"v1.7.3", "v1.7.8", true, false},          // too old kubelet (older than kubeadmconstants.MinimumKubeletVersion), should fail.
		{"v1.9.0", "v1.9.5", false, false},         // kubelet within same major.minor as control plane
		{"v1.9.5", "v1.9.1", false, false},         // kubelet is newer, but still within same major.minor as control plane
		{"v1.9.0", "v1.10.1", false, false},        // kubelet is lower than control plane, but newer than minimally supported
		{"v1.10.0-alpha.1", "v1.9.1", true, false}, // kubelet is newer (development build) than control plane, should fail.
		{"v1.10.0", "v1.9.5", true, false},         // kubelet is newer (release) than control plane, should fail.
	}

	dir, err := ioutil.TempDir("", "test-kubelet-version-check")
	if err != nil {
		t.Errorf("Failed to create directory for testing GetKubeletVersion: %v", err)
	}
	defer os.RemoveAll(dir)

	// We don't want to call real kubelet or something else in $PATH
	oldPATH := os.Getenv("PATH")
	defer os.Setenv("PATH", oldPATH)

	os.Setenv("PATH", dir)

	kubeletFn := filepath.Join(dir, "kubelet")
	for _, tc := range cases {

		content := []byte(fmt.Sprintf("#!/bin/sh\necho 'Kubernetes %s'", tc.kubeletVersion))
		if err := ioutil.WriteFile(kubeletFn, content, 0755); err != nil {
			t.Errorf("Error creating test stub file %s: %v", kubeletFn, err)
		}

		check := KubeletVersionCheck{KubernetesVersion: tc.k8sVersion}
		warnings, errors := check.Check()

		switch {
		case warnings != nil && !tc.expectWarnings:
			t.Errorf("KubeletVersionCheck: unexpected warnings for kubelet version %q and kubernetes version %q. Warnings: %v", tc.kubeletVersion, tc.k8sVersion, warnings)
		case warnings == nil && tc.expectWarnings:
			t.Errorf("KubeletVersionCheck: expected warnings for kubelet version %q and kubernetes version %q but got nothing", tc.kubeletVersion, tc.k8sVersion)
		case errors != nil && !tc.expectErrors:
			t.Errorf("KubeletVersionCheck: unexpected errors for kubelet version %q and kubernetes version %q. errors: %v", tc.kubeletVersion, tc.k8sVersion, errors)
		case errors == nil && tc.expectErrors:
			t.Errorf("KubeletVersionCheck: expected errors for kubelet version %q and kubernetes version %q but got nothing", tc.kubeletVersion, tc.k8sVersion)
		}

	}

}

func TestSetHasItemOrAll(t *testing.T) {
	var tests = []struct {
		ignoreSet      sets.String
		testString     string
		expectedResult bool
	}{
		{sets.NewString(), "foo", false},
		{sets.NewString("all"), "foo", true},
		{sets.NewString("all", "bar"), "foo", true},
		{sets.NewString("bar"), "foo", false},
		{sets.NewString("baz", "foo", "bar"), "foo", true},
		{sets.NewString("baz", "bar", "foo"), "Foo", true},
	}

	for _, rt := range tests {
		result := setHasItemOrAll(rt.ignoreSet, rt.testString)
		if result != rt.expectedResult {
			t.Errorf(
				"setHasItemOrAll: expected: %v actual: %v (arguments: %q, %q)",
				rt.expectedResult, result,
				rt.ignoreSet,
				rt.testString,
			)
		}
	}
}
