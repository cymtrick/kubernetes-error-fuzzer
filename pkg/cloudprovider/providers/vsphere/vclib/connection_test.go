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

package vclib_test

import (
	"context"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"k8s.io/kubernetes/pkg/cloudprovider/providers/vsphere/vclib"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/vsphere/vclib/fixtures"
)

func createTestServer(
	t *testing.T,
	caCertPath string,
	serverCertPath string,
	serverKeyPath string,
	handler http.HandlerFunc,
) (*httptest.Server, string) {
	caCertPEM, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		t.Fatalf("Could not read ca cert from file")
	}

	serverCert, err := tls.LoadX509KeyPair(serverCertPath, serverKeyPath)
	if err != nil {
		t.Fatalf("Could not load server cert and server key from files: %#v", err)
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
		t.Fatalf("Cannot add CA to CAPool")
	}

	server := httptest.NewUnstartedServer(http.HandlerFunc(handler))
	server.TLS = &tls.Config{
		Certificates: []tls.Certificate{
			serverCert,
		},
		RootCAs: certPool,
	}

	// calculate the leaf certificate's fingerprint
	x509LeafCert := server.TLS.Certificates[0].Certificate[0]
	tpBytes := sha1.Sum(x509LeafCert)
	tpString := fmt.Sprintf("%x", tpBytes)

	return server, tpString
}

func TestWithValidCaCert(t *testing.T) {
	handler, verifyConnectionWasMade := getRequestVerifier(t)

	server, _ := createTestServer(t, fixtures.CaCertPath, fixtures.ServerCertPath, fixtures.ServerKeyPath, handler)
	server.StartTLS()
	u := mustParseUrl(t, server.URL)

	connection := &vclib.VSphereConnection{
		Hostname: u.Hostname(),
		Port:     u.Port(),
		CACert:   fixtures.CaCertPath,
	}

	// Ignoring error here, because we only care about the TLS connection
	connection.NewClient(context.Background())

	verifyConnectionWasMade()
}

func TestWithVerificationWithWrongThumbprint(t *testing.T) {
	handler, _ := getRequestVerifier(t)

	server, _ := createTestServer(t, fixtures.CaCertPath, fixtures.ServerCertPath, fixtures.ServerKeyPath, handler)
	server.StartTLS()
	u := mustParseUrl(t, server.URL)

	connection := &vclib.VSphereConnection{
		Hostname:   u.Hostname(),
		Port:       u.Port(),
		Thumbprint: "obviously wrong",
	}

	_, err := connection.NewClient(context.Background())

	if msg := err.Error(); !strings.Contains(msg, "thumbprint does not match") {
		t.Fatalf("Expected wrong thumbprint error, got '%s'", msg)
	}
}

func TestWithVerificationWithoutCaCertOrThumbprint(t *testing.T) {
	handler, _ := getRequestVerifier(t)

	server, _ := createTestServer(t, fixtures.CaCertPath, fixtures.ServerCertPath, fixtures.ServerKeyPath, handler)
	server.StartTLS()
	u := mustParseUrl(t, server.URL)

	connection := &vclib.VSphereConnection{
		Hostname: u.Hostname(),
		Port:     u.Port(),
	}

	_, err := connection.NewClient(context.Background())

	verifyWrappedX509UnkownAuthorityErr(t, err)
}

func TestWithValidThumbprint(t *testing.T) {
	handler, verifyConnectionWasMade := getRequestVerifier(t)

	server, thumbprint :=
		createTestServer(t, fixtures.CaCertPath, fixtures.ServerCertPath, fixtures.ServerKeyPath, handler)
	server.StartTLS()
	u := mustParseUrl(t, server.URL)

	connection := &vclib.VSphereConnection{
		Hostname:   u.Hostname(),
		Port:       u.Port(),
		Thumbprint: thumbprint,
	}

	// Ignoring error here, because we only care about the TLS connection
	connection.NewClient(context.Background())

	verifyConnectionWasMade()
}

func TestWithValidThumbprintAlternativeFormat(t *testing.T) {
	handler, verifyConnectionWasMade := getRequestVerifier(t)

	server, thumbprint :=
		createTestServer(t, fixtures.CaCertPath, fixtures.ServerCertPath, fixtures.ServerKeyPath, handler)
	server.StartTLS()
	u := mustParseUrl(t, server.URL)

	// lowercase, remove the ':'
	tpDifferentFormat := strings.Replace(strings.ToLower(thumbprint), ":", "", -1)

	connection := &vclib.VSphereConnection{
		Hostname:   u.Hostname(),
		Port:       u.Port(),
		Thumbprint: tpDifferentFormat,
	}

	// Ignoring error here, because we only care about the TLS connection
	connection.NewClient(context.Background())

	verifyConnectionWasMade()
}

func TestWithInvalidCaCertPath(t *testing.T) {
	connection := &vclib.VSphereConnection{
		Hostname: "should-not-matter",
		Port:     "should-not-matter",
		CACert:   "invalid-path",
	}

	_, err := connection.NewClient(context.Background())
	if _, ok := err.(*os.PathError); !ok {
		t.Fatalf("Expected an os.PathError, got: '%s' (%#v)", err.Error(), err)
	}
}

func TestInvalidCaCert(t *testing.T) {
	t.Skip("Waiting for https://github.com/vmware/govmomi/pull/1154")

	connection := &vclib.VSphereConnection{
		Hostname: "should-not-matter",
		Port:     "should-not-matter",
		CACert:   fixtures.InvalidCertPath,
	}

	_, err := connection.NewClient(context.Background())

	if err != vclib.ErrCaCertInvalid {
		t.Fatalf("ErrCaCertInvalid should have occurred, instead got: %v", err)
	}
}

func verifyWrappedX509UnkownAuthorityErr(t *testing.T, err error) {
	urlErr, ok := err.(*url.Error)
	if !ok {
		t.Fatalf("Expected to receive an url.Error, got '%s' (%#v)", err.Error(), err)
	}
	x509Err, ok := urlErr.Err.(x509.UnknownAuthorityError)
	if !ok {
		t.Fatalf("Expected to receive a wrapped x509.UnknownAuthorityError, got: '%s' (%#v)", urlErr.Error(), urlErr)
	}
	if msg := x509Err.Error(); msg != "x509: certificate signed by unknown authority" {
		t.Fatalf("Expected 'signed by unknown authority' error, got: '%s'", msg)
	}
}

func getRequestVerifier(t *testing.T) (http.HandlerFunc, func()) {
	gotRequest := false

	handler := func(w http.ResponseWriter, r *http.Request) {
		gotRequest = true
	}

	checker := func() {
		if !gotRequest {
			t.Fatalf("Never saw a request, maybe TLS connection could not be established?")
		}
	}

	return handler, checker
}

func mustParseUrl(t *testing.T, i string) *url.URL {
	u, err := url.Parse(i)
	if err != nil {
		t.Fatalf("Cannot parse URL: %v", err)
	}
	return u
}
