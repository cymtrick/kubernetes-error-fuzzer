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

package options

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"

	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/apiserver/pkg/authorization/authorizerfactory"
	"k8s.io/apiserver/pkg/authorization/path"
	"k8s.io/apiserver/pkg/authorization/union"
	"k8s.io/apiserver/pkg/server"
	authorizationclient "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// DelegatingAuthorizationOptions provides an easy way for composing API servers to delegate their authorization to
// the root kube API server.
// WARNING: never assume that every authenticated incoming request already does authorization.
//          The aggregator in the kube API server does this today, but this behaviour is not
//          guaranteed in the future.
type DelegatingAuthorizationOptions struct {
	// RemoteKubeConfigFile is the file to use to connect to a "normal" kube API server which hosts the
	// SubjectAccessReview.authorization.k8s.io endpoint for checking tokens.
	RemoteKubeConfigFile string

	// AllowCacheTTL is the length of time that a successful authorization response will be cached
	AllowCacheTTL time.Duration

	// DenyCacheTTL is the length of time that an unsuccessful authorization response will be cached.
	// You generally want more responsive, "deny, try again" flows.
	DenyCacheTTL time.Duration

	// AlwaysAllowPaths are HTTP paths which are excluded from authorization. They can be plain
	// paths or end in * in which case prefix-match is applied. A leading / is optional.
	AlwaysAllowPaths []string
}

func NewDelegatingAuthorizationOptions() *DelegatingAuthorizationOptions {
	return &DelegatingAuthorizationOptions{
		// very low for responsiveness, but high enough to handle storms
		AllowCacheTTL: 10 * time.Second,
		DenyCacheTTL:  10 * time.Second,
	}
}

func (s *DelegatingAuthorizationOptions) Validate() []error {
	allErrors := []error{}
	return allErrors
}

func (s *DelegatingAuthorizationOptions) AddFlags(fs *pflag.FlagSet) {
	if s == nil {
		return
	}

	fs.StringVar(&s.RemoteKubeConfigFile, "authorization-kubeconfig", s.RemoteKubeConfigFile,
		"kubeconfig file pointing at the 'core' kubernetes server with enough rights to create "+
			" subjectaccessreviews.authorization.k8s.io.")

	fs.DurationVar(&s.AllowCacheTTL, "authorization-webhook-cache-authorized-ttl",
		s.AllowCacheTTL,
		"The duration to cache 'authorized' responses from the webhook authorizer.")

	fs.DurationVar(&s.DenyCacheTTL,
		"authorization-webhook-cache-unauthorized-ttl", s.DenyCacheTTL,
		"The duration to cache 'unauthorized' responses from the webhook authorizer.")

	fs.StringSliceVar(&s.AlwaysAllowPaths, "authorization-always-allow-paths", s.AlwaysAllowPaths,
		"A list of HTTP paths to skip during authorization, i.e. these are authorized without "+
			"contacting the 'core' kubernetes server.")
}

func (s *DelegatingAuthorizationOptions) ApplyTo(c *server.AuthorizationInfo) error {
	if s == nil {
		c.Authorizer = authorizerfactory.NewAlwaysAllowAuthorizer()
		return nil
	}

	a, err := s.ToAuthorization()
	if err != nil {
		return err
	}
	c.Authorizer = a
	return nil
}

func (s *DelegatingAuthorizationOptions) ToAuthorization() (authorizer.Authorizer, error) {
	var authorizers []authorizer.Authorizer

	if len(s.AlwaysAllowPaths) > 0 {
		a, err := path.NewAuthorizer(s.AlwaysAllowPaths)
		if err != nil {
			return nil, err
		}
		authorizers = append(authorizers, a)
	}

	sarClient, err := s.newSubjectAccessReview()
	if err != nil {
		return nil, err
	}
	cfg := authorizerfactory.DelegatingAuthorizerConfig{
		SubjectAccessReviewClient: sarClient,
		AllowCacheTTL:             s.AllowCacheTTL,
		DenyCacheTTL:              s.DenyCacheTTL,
	}
	a, err := cfg.New()
	if err != nil {
		return nil, err
	}
	authorizers = append(authorizers, a)

	return union.New(authorizers...), nil
}

func (s *DelegatingAuthorizationOptions) newSubjectAccessReview() (authorizationclient.SubjectAccessReviewInterface, error) {
	var clientConfig *rest.Config
	var err error
	if len(s.RemoteKubeConfigFile) > 0 {
		loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: s.RemoteKubeConfigFile}
		loader := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})

		clientConfig, err = loader.ClientConfig()

	} else {
		// without the remote kubeconfig file, try to use the in-cluster config.  Most addon API servers will
		// use this path
		clientConfig, err = rest.InClusterConfig()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get delegated authorization kubeconfig: %v", err)
	}

	// set high qps/burst limits since this will effectively limit API server responsiveness
	clientConfig.QPS = 200
	clientConfig.Burst = 400

	client, err := authorizationclient.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}

	return client.SubjectAccessReviews(), nil
}
