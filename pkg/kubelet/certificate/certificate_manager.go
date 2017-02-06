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

package certificate

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	cryptorand "crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/util/cert"
	certificates "k8s.io/kubernetes/pkg/apis/certificates/v1beta1"
	certificatesclient "k8s.io/kubernetes/pkg/client/clientset_generated/clientset/typed/certificates/v1beta1"
)

const (
	syncPeriod = 1 * time.Hour
)

// Manager maintains and updates the certificates in use by this certificate
// manager. In the background it communicates with the API server to get new
// certificates for certificates about to expire.
type Manager interface {
	// CertificateSigningRequestClient sets the client interface that is used for
	// signing new certificates generated as part of rotation.
	SetCertificateSigningRequestClient(certificatesclient.CertificateSigningRequestInterface) error
	// Start the API server status sync loop.
	Start()
	// Current returns the currently selected certificate from the
	// certificate manager, as well as the associated certificate and key data
	// in PEM format.
	Current() *tls.Certificate
}

// Config is the set of configuration parameters available for a new Manager.
type Config struct {
	// CertificateSigningRequestClient will be used for signing new certificate
	// requests generated when a key rotation occurs. It must be set either at
	// initialization or by using CertificateSigningRequestClient before
	// Manager.Start() is called.
	CertificateSigningRequestClient certificatesclient.CertificateSigningRequestInterface
	// Template is the CertificateRequest that will be used as a template for
	// generating certificate signing requests for all new keys generated as
	// part of rotation. It follows the same rules as the template parameter of
	// crypto.x509.CreateCertificateRequest in the Go standard libraries.
	Template *x509.CertificateRequest
	// Usages is the types of usages that certificates generated by the manager
	// can be used for.
	Usages []certificates.KeyUsage
	// CertificateStore is a persistent store where the current cert/key is
	// kept and future cert/key pairs will be persisted after they are
	// generated.
	CertificateStore Store
	// BootstrapCertificatePEM is the certificate data that will be returned
	// from the Manager if the CertificateStore doesn't have any cert/key pairs
	// currently available and has not yet had a chance to get a new cert/key
	// pair from the API. If the CertificateStore does have a cert/key pair,
	// this will be ignored. If there is no cert/key pair available in the
	// CertificateStore, as soon as Start is called, it will request a new
	// cert/key pair from the CertificateSigningRequestClient. This is intended
	// to allow the first boot of a component to be initialized using a
	// generic, multi-use cert/key pair which will be quickly replaced with a
	// unique cert/key pair.
	BootstrapCertificatePEM []byte
	// BootstrapKeyPEM is the key data that will be returned from the Manager
	// if the CertificateStore doesn't have any cert/key pairs currently
	// available. If the CertificateStore does have a cert/key pair, this will
	// be ignored. If the bootstrap cert/key pair are used, they will be
	// rotated at the first opportunity, possibly well in advance of expiring.
	// This is intended to allow the first boot of a component to be
	// initialized using a generic, multi-use cert/key pair which will be
	// quickly replaced with a unique cert/key pair.
	BootstrapKeyPEM []byte
}

// Store is responsible for getting and updating the current certificate.
// Depending on the concrete implementation, the backing store for this
// behavior may vary.
type Store interface {
	// Current returns the currently selected certificate, as well as the
	// associated certificate and key data in PEM format. If the Store doesn't
	// have a cert/key pair currently, it should return a NoCertKeyError so
	// that the Manager can recover by using bootstrap certificates to request
	// a new cert/key pair.
	Current() (*tls.Certificate, error)
	// Update accepts the PEM data for the cert/key pair and makes the new
	// cert/key pair the 'current' pair, that will be returned by future calls
	// to Current().
	Update(cert, key []byte) (*tls.Certificate, error)
}

// NoCertKeyError indicates there is no cert/key currently available.
type NoCertKeyError string

func (e *NoCertKeyError) Error() string { return string(*e) }

type manager struct {
	certSigningRequestClient certificatesclient.CertificateSigningRequestInterface
	template                 *x509.CertificateRequest
	usages                   []certificates.KeyUsage
	certStore                Store
	certAccessLock           sync.RWMutex
	cert                     *tls.Certificate
	rotationDeadline         time.Time
	forceRotation            bool
}

// NewManager returns a new certificate manager. A certificate manager is
// responsible for being the authoritative source of certificates in the
// Kubelet and handling updates due to rotation.
func NewManager(config *Config) (Manager, error) {
	cert, forceRotation, err := getCurrentCertificateOrBootstrap(
		config.CertificateStore,
		config.BootstrapCertificatePEM,
		config.BootstrapKeyPEM)
	if err != nil {
		return nil, err
	}

	m := manager{
		certSigningRequestClient: config.CertificateSigningRequestClient,
		template:                 config.Template,
		usages:                   config.Usages,
		certStore:                config.CertificateStore,
		cert:                     cert,
		forceRotation:            forceRotation,
	}

	return &m, nil
}

// Current returns the currently selected certificate from the certificate
// manager. This can be nil if the manager was initialized without a
// certificate and has not yet received one from the
// CertificateSigningRequestClient.
func (m *manager) Current() *tls.Certificate {
	m.certAccessLock.RLock()
	defer m.certAccessLock.RUnlock()
	return m.cert
}

// SetCertificateSigningRequestClient sets the client interface that is used
// for signing new certificates generated as part of rotation. It must be
// called before Start() and can not be used to change the
// CertificateSigningRequestClient that has already been set. This method is to
// support the one specific scenario where the CertificateSigningRequestClient
// uses the CertificateManager.
func (m *manager) SetCertificateSigningRequestClient(certSigningRequestClient certificatesclient.CertificateSigningRequestInterface) error {
	if m.certSigningRequestClient == nil {
		m.certSigningRequestClient = certSigningRequestClient
		return nil
	}
	return fmt.Errorf("CertificateSigningRequestClient is already set.")
}

// Start will start the background work of rotating the certificates.
func (m *manager) Start() {
	// Certificate rotation depends on access to the API server certificate
	// signing API, so don't start the certificate manager if we don't have a
	// client. This will happen on the cluster master, where the kubelet is
	// responsible for bootstrapping the pods of the master components.
	if m.certSigningRequestClient == nil {
		glog.V(2).Infof("Certificate rotation is not enabled, no connection to the apiserver.")
		return
	}

	glog.V(2).Infof("Certificate rotation is enabled.")

	m.setRotationDeadline()

	// Synchronously request a certificate before entering the background
	// loop to allow bootstrap scenarios, where the certificate manager
	// doesn't have a certificate at all yet.
	if m.shouldRotate() {
		_, err := m.rotateCerts()
		if err != nil {
			glog.Errorf("Could not rotate certificates: %v", err)
		}
	}
	backoff := wait.Backoff{
		Duration: 2 * time.Second,
		Factor:   2,
		Jitter:   0.1,
		Steps:    7,
	}
	go wait.Forever(func() {
		time.Sleep(m.rotationDeadline.Sub(time.Now()))
		if err := wait.ExponentialBackoff(backoff, m.rotateCerts); err != nil {
			glog.Errorf("Reached backoff limit, still unable to rotate certs: %v", err)
			wait.PollInfinite(128*time.Second, m.rotateCerts)
		}
	}, 0)
}

func getCurrentCertificateOrBootstrap(
	store Store,
	bootstrapCertificatePEM []byte,
	bootstrapKeyPEM []byte) (cert *tls.Certificate, shouldRotate bool, errResult error) {

	currentCert, err := store.Current()
	if err == nil {
		return currentCert, false, nil
	}

	if _, ok := err.(*NoCertKeyError); !ok {
		return nil, false, err
	}

	if bootstrapCertificatePEM == nil || bootstrapKeyPEM == nil {
		return nil, true, nil
	}

	bootstrapCert, err := tls.X509KeyPair(bootstrapCertificatePEM, bootstrapKeyPEM)
	if err != nil {
		return nil, false, err
	}
	if len(bootstrapCert.Certificate) < 1 {
		return nil, false, fmt.Errorf("no cert/key data found")
	}

	certs, err := x509.ParseCertificates(bootstrapCert.Certificate[0])
	if err != nil {
		return nil, false, fmt.Errorf("unable to parse certificate data: %v", err)
	}
	bootstrapCert.Leaf = certs[0]
	return &bootstrapCert, true, nil
}

// shouldRotate looks at how close the current certificate is to expiring and
// decides if it is time to rotate or not.
func (m *manager) shouldRotate() bool {
	m.certAccessLock.RLock()
	defer m.certAccessLock.RUnlock()
	if m.cert == nil {
		return true
	}
	if m.forceRotation {
		return true
	}
	return time.Now().After(m.rotationDeadline)
}

func (m *manager) rotateCerts() (bool, error) {
	csrPEM, keyPEM, err := m.generateCSR()
	if err != nil {
		glog.Errorf("Unable to generate a certificate signing request: %v", err)
		return false, nil
	}

	// Call the Certificate Signing Request API to get a certificate for the
	// new private key.
	crtPEM, err := requestCertificate(m.certSigningRequestClient, csrPEM, m.usages)
	if err != nil {
		glog.Errorf("Failed while requesting a signed certificate from the master: %v", err)
		return false, nil
	}

	cert, err := m.certStore.Update(crtPEM, keyPEM)
	if err != nil {
		glog.Errorf("Unable to store the new cert/key pair: %v", err)
		return false, nil
	}

	m.updateCached(cert)
	m.setRotationDeadline()
	m.forceRotation = false
	return true, nil
}

// setRotationDeadline sets a cached value for the threshold at which the
// current certificate should be rotated, 80%+/-10% of the expiration of the
// certificate.
func (m *manager) setRotationDeadline() {
	m.certAccessLock.RLock()
	defer m.certAccessLock.RUnlock()
	if m.cert == nil {
		m.rotationDeadline = time.Now()
		return
	}

	notAfter := m.cert.Leaf.NotAfter
	totalDuration := float64(notAfter.Sub(m.cert.Leaf.NotBefore))

	// Use some jitter to set the rotation threshold so each node will rotate
	// at approximately 70-90% of the total lifetime of the certificate.  With
	// jitter, if a number of nodes are added to a cluster at approximately the
	// same time (such as cluster creation time), they won't all try to rotate
	// certificates at the same time for the rest of the life of the cluster.
	jitteryDuration := wait.Jitter(time.Duration(totalDuration), 0.2) - time.Duration(totalDuration*0.3)

	m.rotationDeadline = m.cert.Leaf.NotBefore.Add(jitteryDuration)
}

func (m *manager) updateCached(cert *tls.Certificate) {
	m.certAccessLock.Lock()
	defer m.certAccessLock.Unlock()
	m.cert = cert
}

func (m *manager) generateCSR() (csrPEM []byte, keyPEM []byte, err error) {
	// Generate a new private key.
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), cryptorand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to generate a new private key: %v", err)
	}
	der, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to marshal the new key to DER: %v", err)
	}

	keyPEM = pem.EncodeToMemory(&pem.Block{Type: cert.ECPrivateKeyBlockType, Bytes: der})

	csrPEM, err = cert.MakeCSRFromTemplate(privateKey, m.template)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to create a csr from the private key: %v", err)
	}
	return csrPEM, keyPEM, nil
}

// requestCertificate will create a certificate signing request using the PEM
// encoded CSR and send it to API server, then it will watch the object's
// status, once approved by API server, it will return the API server's issued
// certificate (pem-encoded). If there is any errors, or the watch timeouts, it
// will return an error.
//
// NOTE This is a copy of a function with the same name in
// k8s.io/kubernetes/pkg/kubelet/util/csr/csr.go, changing only the package that
// CertificateSigningRequestInterface and KeyUsage are imported from.
func requestCertificate(client certificatesclient.CertificateSigningRequestInterface, csrData []byte, usages []certificates.KeyUsage) (certData []byte, err error) {
	glog.Infof("Requesting new certificate.")
	req, err := client.Create(&certificates.CertificateSigningRequest{
		// Username, UID, Groups will be injected by API server.
		TypeMeta:   metav1.TypeMeta{Kind: "CertificateSigningRequest"},
		ObjectMeta: metav1.ObjectMeta{GenerateName: "csr-"},

		Spec: certificates.CertificateSigningRequestSpec{
			Request: csrData,
			Usages:  usages,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("cannot create certificate signing request: %v", err)
	}

	// Make a default timeout = 3600s.
	var defaultTimeoutSeconds int64 = 3600
	certWatch, err := client.Watch(metav1.ListOptions{
		Watch:          true,
		TimeoutSeconds: &defaultTimeoutSeconds,
		FieldSelector:  fields.OneTermEqualSelector("metadata.name", req.Name).String(),
	})
	if err != nil {
		return nil, fmt.Errorf("cannot watch on the certificate signing request: %v", err)
	}
	defer certWatch.Stop()
	ch := certWatch.ResultChan()

	for {
		event, ok := <-ch
		if !ok {
			break
		}

		if event.Type == watch.Modified || event.Type == watch.Added {
			if event.Object.(*certificates.CertificateSigningRequest).UID != req.UID {
				continue
			}
			status := event.Object.(*certificates.CertificateSigningRequest).Status
			for _, c := range status.Conditions {
				if c.Type == certificates.CertificateDenied {
					return nil, fmt.Errorf("certificate signing request is not approved, reason: %v, message: %v", c.Reason, c.Message)
				}
				if c.Type == certificates.CertificateApproved && status.Certificate != nil {
					return status.Certificate, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("watch channel closed")
}
