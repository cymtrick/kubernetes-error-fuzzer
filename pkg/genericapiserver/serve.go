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

package genericapiserver

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	certutil "k8s.io/kubernetes/pkg/util/cert"
	utilruntime "k8s.io/kubernetes/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/util/validation"

	"github.com/golang/glog"
	"github.com/pkg/errors"
)

const (
	defaultKeepAlivePeriod = 3 * time.Minute
)

// serveSecurely runs the secure http server. It fails only if certificates cannot
// be loaded or the initial listen call fails. The actual server loop (stoppable by closing
// stopCh) runs in a go routine, i.e. serveSecurely does not block.
func (s *GenericAPIServer) serveSecurely(stopCh <-chan struct{}) error {
	secureServer := &http.Server{
		Addr:           s.SecureServingInfo.BindAddress,
		Handler:        s.Handler,
		MaxHeaderBytes: 1 << 20,
		TLSConfig: &tls.Config{
			NameToCertificate: s.SecureServingInfo.SNICerts,
			// Can't use SSLv3 because of POODLE and BEAST
			// Can't use TLSv1.0 because of POODLE and BEAST using CBC cipher
			// Can't use TLSv1.1 because of RC4 cipher usage
			MinVersion: tls.VersionTLS12,
			// enable HTTP2 for go's 1.7 HTTP Server
			NextProtos: []string{"h2", "http/1.1"},
		},
	}

	if s.SecureServingInfo.Cert != nil {
		secureServer.TLSConfig.Certificates = []tls.Certificate{*s.SecureServingInfo.Cert}
	}

	// append all named certs. Otherwise, the go tls stack will think no SNI processing
	// is necessary because there is only one cert anyway.
	// Moreover, if ServerCert.CertFile/ServerCert.KeyFile are not set, the first SNI
	// cert will become the default cert. That's what we expect anyway.
	for _, c := range s.SecureServingInfo.SNICerts {
		secureServer.TLSConfig.Certificates = append(secureServer.TLSConfig.Certificates, *c)
	}

	if len(s.SecureServingInfo.ClientCA) > 0 {
		clientCAs, err := certutil.NewPool(s.SecureServingInfo.ClientCA)
		if err != nil {
			return fmt.Errorf("unable to load client CA file: %v", err)
		}
		// Populate PeerCertificates in requests, but don't reject connections without certificates
		// This allows certificates to be validated by authenticators, while still allowing other auth types
		secureServer.TLSConfig.ClientAuth = tls.RequestClientCert
		// Specify allowed CAs for client certificates
		secureServer.TLSConfig.ClientCAs = clientCAs
	}

	glog.Infof("Serving securely on %s", s.SecureServingInfo.BindAddress)
	var err error
	s.effectiveSecurePort, err = runServer(secureServer, s.SecureServingInfo.BindNetwork, stopCh)
	return err
}

// serveInsecurely run the insecure http server. It fails only if the initial listen
// call fails. The actual server loop (stoppable by closing stopCh) runs in a go
// routine, i.e. serveInsecurely does not block.
func (s *GenericAPIServer) serveInsecurely(stopCh <-chan struct{}) error {
	insecureServer := &http.Server{
		Addr:           s.InsecureServingInfo.BindAddress,
		Handler:        s.InsecureHandler,
		MaxHeaderBytes: 1 << 20,
	}
	glog.Infof("Serving insecurely on %s", s.InsecureServingInfo.BindAddress)
	var err error
	s.effectiveInsecurePort, err = runServer(insecureServer, s.InsecureServingInfo.BindNetwork, stopCh)
	return err
}

// runServer listens on the given port, then spawns a go-routine continuously serving
// until the stopCh is closed. The port is returned. This function does not block.
func runServer(server *http.Server, network string, stopCh <-chan struct{}) (int, error) {
	if len(server.Addr) == 0 {
		return 0, errors.New("address cannot be empty")
	}

	if len(network) == 0 {
		network = "tcp"
	}

	// first listen is synchronous (fail early!)
	ln, err := net.Listen(network, server.Addr)
	if err != nil {
		return 0, fmt.Errorf("failed to listen on %v: %v", server.Addr, err)
	}

	// get port
	tcpAddr, ok := ln.Addr().(*net.TCPAddr)
	if !ok {
		ln.Close()
		return 0, fmt.Errorf("invalid listen address: %q", ln.Addr().String())
	}

	lock := sync.Mutex{} // to avoid we close an old listener during a listen retry
	go func() {
		<-stopCh
		lock.Lock()
		defer lock.Unlock()
		ln.Close()
	}()

	go func() {
		defer utilruntime.HandleCrash()

		for {
			var listener net.Listener
			listener = tcpKeepAliveListener{ln.(*net.TCPListener)}
			if server.TLSConfig != nil {
				listener = tls.NewListener(listener, server.TLSConfig)
			}

			err := server.Serve(listener)
			glog.Errorf("Error serving %v (%v); will try again.", server.Addr, err)

			// listen again
			func() {
				lock.Lock()
				defer lock.Unlock()
				for {
					time.Sleep(15 * time.Second)

					ln, err = net.Listen("tcp", server.Addr)
					if err == nil {
						return
					}
					select {
					case <-stopCh:
						return
					default:
					}
					glog.Errorf("Error listening on %v (%v); will try again.", server.Addr, err)
				}
			}()

			select {
			case <-stopCh:
				return
			default:
			}
		}
	}()

	return tcpAddr.Port, nil
}

type namedTlsCert struct {
	tlsCert tls.Certificate

	// names is a list of domain patterns: fully qualified domain names, possibly prefixed with
	// wildcard segments.
	names []string
}

// getNamedCertificateMap returns a map of *tls.Certificate by name. It's is
// suitable for use in tls.Config#NamedCertificates. Returns an error if any of the certs
// cannot be loaded. Returns nil if len(certs) == 0
func getNamedCertificateMap(certs []namedTlsCert) (map[string]*tls.Certificate, error) {
	// register certs with implicit names first, reverse order such that earlier trump over the later
	byName := map[string]*tls.Certificate{}
	for i := len(certs) - 1; i >= 0; i-- {
		if len(certs[i].names) > 0 {
			continue
		}
		cert := &certs[i].tlsCert

		// read names from certificate common names and DNS names
		if len(cert.Certificate) == 0 {
			return nil, fmt.Errorf("empty SNI certificate, skipping")
		}
		x509Cert, err := x509.ParseCertificate(cert.Certificate[0])
		if err != nil {
			return nil, fmt.Errorf("parse error for SNI certificate: %v", err)
		}
		cn := x509Cert.Subject.CommonName
		if cn == "*" || len(validation.IsDNS1123Subdomain(strings.TrimPrefix(cn, "*."))) == 0 {
			byName[cn] = cert
		}
		for _, san := range x509Cert.DNSNames {
			byName[san] = cert
		}
		// intentionally all IPs in the cert are ignored as SNI forbids passing IPs
		// to select a cert. Before go 1.6 the tls happily passed IPs as SNI values.
	}

	// register certs with explicit names last, overwriting every of the implicit ones,
	// again in reverse order.
	for i := len(certs) - 1; i >= 0; i-- {
		namedCert := &certs[i]
		for _, name := range namedCert.names {
			byName[name] = &certs[i].tlsCert
		}
	}

	return byName, nil
}

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
//
// Copied from Go 1.7.2 net/http/server.go
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(defaultKeepAlivePeriod)
	return tc, nil
}
