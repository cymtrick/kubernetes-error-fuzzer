/*
Copyright 2015 The Kubernetes Authors.

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

package node

import (
	"fmt"
	"net/http"
	"net/url"
	"path"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/rest"
	"k8s.io/kubernetes/pkg/capabilities"
	genericapirequest "k8s.io/kubernetes/pkg/genericapiserver/api/request"
	"k8s.io/kubernetes/pkg/kubelet/client"
	"k8s.io/kubernetes/pkg/registry/core/node"
	genericregistry "k8s.io/kubernetes/pkg/registry/generic/registry"
	genericrest "k8s.io/kubernetes/pkg/registry/generic/rest"
)

// ProxyREST implements the proxy subresource for a Node
type ProxyREST struct {
	Store          *genericregistry.Store
	Connection     client.ConnectionInfoGetter
	ProxyTransport http.RoundTripper
}

// Implement Connecter
var _ = rest.Connecter(&ProxyREST{})

var proxyMethods = []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}

// New returns an empty service resource
func (r *ProxyREST) New() runtime.Object {
	return &api.Node{}
}

// ConnectMethods returns the list of HTTP methods that can be proxied
func (r *ProxyREST) ConnectMethods() []string {
	return proxyMethods
}

// NewConnectOptions returns versioned resource that represents proxy parameters
func (r *ProxyREST) NewConnectOptions() (runtime.Object, bool, string) {
	return &api.NodeProxyOptions{}, true, "path"
}

// Connect returns a handler for the node proxy
func (r *ProxyREST) Connect(ctx genericapirequest.Context, id string, opts runtime.Object, responder rest.Responder) (http.Handler, error) {
	proxyOpts, ok := opts.(*api.NodeProxyOptions)
	if !ok {
		return nil, fmt.Errorf("Invalid options object: %#v", opts)
	}
	location, transport, err := node.ResourceLocation(r.Store, r.Connection, r.ProxyTransport, ctx, id)
	if err != nil {
		return nil, err
	}
	location.Path = path.Join(location.Path, proxyOpts.Path)
	// Return a proxy handler that uses the desired transport, wrapped with additional proxy handling (to get URL rewriting, X-Forwarded-* headers, etc)
	return newThrottledUpgradeAwareProxyHandler(location, transport, true, false, responder), nil
}

func newThrottledUpgradeAwareProxyHandler(location *url.URL, transport http.RoundTripper, wrapTransport, upgradeRequired bool, responder rest.Responder) *genericrest.UpgradeAwareProxyHandler {
	handler := genericrest.NewUpgradeAwareProxyHandler(location, transport, wrapTransport, upgradeRequired, responder)
	handler.MaxBytesPerSec = capabilities.Get().PerConnectionBandwidthLimitBytesPerSec
	return handler
}
