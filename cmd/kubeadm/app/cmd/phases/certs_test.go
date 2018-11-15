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

package phases

import (
	"os"
	"testing"

	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/certs"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/pkiutil"
	testutil "k8s.io/kubernetes/cmd/kubeadm/test"
)

type testCertsData struct {
	cfg *kubeadmapi.InitConfiguration
}

func (t *testCertsData) Cfg() *kubeadmapi.InitConfiguration { return t.cfg }
func (t *testCertsData) ExternalCA() bool                   { return false }
func (t *testCertsData) CertificateDir() string             { return t.cfg.CertificatesDir }
func (t *testCertsData) CertificateWriteDir() string        { return t.cfg.CertificatesDir }

func TestCertsWithCSRs(t *testing.T) {
	csrDir := testutil.SetupTempDir(t)
	defer os.RemoveAll(csrDir)
	certDir := testutil.SetupTempDir(t)
	defer os.RemoveAll(certDir)
	cert := &certs.KubeadmCertAPIServer

	certsData := &testCertsData{
		cfg: testutil.GetDefaultInternalConfig(t),
	}
	certsData.cfg.CertificatesDir = certDir

	// global vars
	csrOnly = true
	csrDir = certDir

	phase := NewCertsPhase()
	// find the api cert phase
	var apiServerPhase *workflow.Phase
	for _, phase := range phase.Phases {
		if phase.Name == cert.Name {
			apiServerPhase = &phase
			break
		}
	}

	if apiServerPhase == nil {
		t.Fatalf("couldn't find apiserver phase")
	}

	err := apiServerPhase.Run(certsData)
	if err != nil {
		t.Fatalf("couldn't run API server phase: %v", err)
	}

	if _, _, err := pkiutil.TryLoadCSRAndKeyFromDisk(csrDir, cert.BaseName); err != nil {
		t.Fatalf("couldn't load certificate %q: %v", cert.BaseName, err)
	}
}
