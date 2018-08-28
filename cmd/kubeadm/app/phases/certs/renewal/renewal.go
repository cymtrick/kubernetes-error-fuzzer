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

package renewal

import (
	"crypto/x509"
	"fmt"

	certutil "k8s.io/client-go/util/cert"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/certs/pkiutil"
)

func RenewExistingCert(certsDir, baseName string, impl Interface) error {
	cert, err := pkiutil.TryLoadCertFromDisk(certsDir, baseName)
	if err != nil {
		return fmt.Errorf("failed to load existing certificate %s: %v", baseName, err)
	}

	cfg := certToConfig(cert)
	newCert, newKey, err := impl.Renew(cfg)
	if err != nil {
		return fmt.Errorf("failed to renew certificate %s: %v", baseName, err)
	}

	if err := pkiutil.WriteCertAndKey(certsDir, baseName, newCert, newKey); err != nil {
		return fmt.Errorf("failed to write new certificate %s: %v", baseName, err)
	}
	return nil
}

func certToConfig(cert *x509.Certificate) *certutil.Config {
	return &certutil.Config{
		CommonName:   cert.Subject.CommonName,
		Organization: cert.Subject.Organization,
		AltNames: certutil.AltNames{
			IPs:      cert.IPAddresses,
			DNSNames: cert.DNSNames,
		},
		Usages: cert.ExtKeyUsage,
	}
}
