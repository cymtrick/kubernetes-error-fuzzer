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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExecCredential is used by exec-based plugins to communicate credentials to
// HTTP transports.
type ExecCredential struct {
	metav1.TypeMeta `json:",inline"`

	// Spec holds information passed to the plugin by the transport.
	Spec ExecCredentialSpec `json:"spec,omitempty"`

	// Status is filled in by the plugin and holds the credentials that the transport
	// should use to contact the API.
	// +optional
	Status *ExecCredentialStatus `json:"status,omitempty"`
}

// ExecCredentialSpec holds request and runtime specific information provided by
// the transport.
type ExecCredentialSpec struct {
	// Cluster contains information to allow an exec plugin to communicate
	// with the kubernetes cluster being authenticated to.
	Cluster Cluster `json:"cluster"`
}

// ExecCredentialStatus holds credentials for the transport to use.
//
// Token and ClientKeyData are sensitive fields. This data should only be
// transmitted in-memory between client and exec plugin process. Exec plugin
// itself should at least be protected via file permissions.
type ExecCredentialStatus struct {
	// ExpirationTimestamp indicates a time when the provided credentials expire.
	// +optional
	ExpirationTimestamp *metav1.Time `json:"expirationTimestamp,omitempty"`
	// Token is a bearer token used by the client for request authentication.
	Token string `json:"token,omitempty"`
	// PEM-encoded client TLS certificates (including intermediates, if any).
	ClientCertificateData string `json:"clientCertificateData,omitempty"`
	// PEM-encoded private key for the above certificate.
	ClientKeyData string `json:"clientKeyData,omitempty"`
}

// Cluster contains information to allow an exec plugin to communicate
// with the kubernetes cluster being authenticated to.
type Cluster struct {
	// Server is the address of the kubernetes cluster (https://hostname:port).
	Server string `json:"server"`
	// ServerName is passed to the server for SNI and is used in the client to check server
	// certificates against. If ServerName is empty, the hostname used to contact the
	// server is used.
	// +optional
	ServerName string `json:"serverName,omitempty"`
	// CAData contains PEM-encoded certificate authority certificates.
	// If empty, system roots should be used.
	// +listType=atomic
	// +optional
	CAData []byte `json:"caData,omitempty"`
	// Config holds additional config data that is specific to the exec
	// plugin with regards to the cluster being authenticated to.
	//
	// This data is sourced from the clientcmd Cluster object's extensions[exec] field:
	//
	// clusters:
	// - name: my-cluster
	//   cluster:
	//     ...
	//     extensions:
	//     - name: exec  # reserved extension name for per cluster exec config
	//       extension:
	//         audience: 06e3fbd18de8  # arbitrary config
	//
	// In some environments, the user config may be exactly the same across many clusters
	// (i.e. call this exec plugin) minus some details that are specific to each cluster
	// such as the audience.  This field allows the per cluster config to be directly
	// specified with the cluster info.  Using this field to store secret data is not
	// recommended as one of the prime benefits of exec plugins is that no secrets need
	// to be stored directly in the kubeconfig.
	// +optional
	Config runtime.RawExtension `json:"config,omitempty"`
}
