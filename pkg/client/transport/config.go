/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package transport

import "net/http"

// Config holds various options for establishing a transport.
type Config struct {
	// UserAgent is an optional field that specifies the caller of this
	// request.
	UserAgent string

	// The base TLS configuration for this transport.
	TLS TLSConfig

	// Username and password for basic authentication
	Username string
	Password string

	// Bearer token for authentication
	BearerToken string

	// Transport may be used for custom HTTP behavior. This attribute may
	// not be specified with the TLS client certificate options. Use
	// WrapTransport for most client level operations.
	Transport http.RoundTripper

	// WrapTransport will be invoked for custom HTTP behavior after the
	// underlying transport is initialized (either the transport created
	// from TLSClientConfig, Transport, or http.DefaultTransport). The
	// config may layer other RoundTrippers on top of the returned
	// RoundTripper.
	WrapTransport func(rt http.RoundTripper) http.RoundTripper
}

// HasCA returns whether the configuration has a certificate authority or not.
func (c *Config) HasCA() bool {
	return len(c.TLS.CAData) > 0 || len(c.TLS.CAFile) > 0
}

// HasBasicAuth returns whether the configuration has basic authentication or not.
func (c *Config) HasBasicAuth() bool {
	return len(c.Username) != 0
}

// HasTokenAuth returns whether the configuration has token authentication or not.
func (c *Config) HasTokenAuth() bool {
	return len(c.BearerToken) != 0
}

// HasCertAuth returns whether the configuration has certificate authentication or not.
func (c *Config) HasCertAuth() bool {
	return len(c.TLS.CertData) != 0 || len(c.TLS.CertFile) != 0
}

// TLSConfig holds the information needed to set up a TLS transport.
type TLSConfig struct {
	CAFile   string // Path of the PEM-encoded server trusted root certificates.
	CertFile string // Path of the PEM-encoded client certificate.
	KeyFile  string // Path of the PEM-encoded client key.

	Insecure bool // Server should be accessed without verifying the certificate. For testing only.

	CAData   []byte // Bytes of the PEM-encoded server trusted root certificates. Supercedes CAFile.
	CertData []byte // Bytes of the PEM-encoded client certificate. Supercedes CertFile.
	KeyData  []byte // Bytes of the PEM-encoded client key. Supercedes KeyFile.
}
