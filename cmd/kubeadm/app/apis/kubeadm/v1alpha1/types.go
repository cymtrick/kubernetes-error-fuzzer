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

package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MasterConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	API                API        `json:"api"`
	Etcd               Etcd       `json:"etcd"`
	Networking         Networking `json:"networking"`
	KubernetesVersion  string     `json:"kubernetesVersion"`
	CloudProvider      string     `json:"cloudProvider"`
	NodeName           string     `json:"nodeName"`
	AuthorizationModes []string   `json:"authorizationModes"`

	Token    string        `json:"token"`
	TokenTTL time.Duration `json:"tokenTTL"`

	APIServerExtraArgs         map[string]string `json:"apiServerExtraArgs"`
	ControllerManagerExtraArgs map[string]string `json:"controllerManagerExtraArgs"`
	SchedulerExtraArgs         map[string]string `json:"schedulerExtraArgs"`

	// APIServerCertSANs sets extra Subject Alternative Names for the API Server signing cert
	APIServerCertSANs []string `json:"apiServerCertSANs"`
	// CertificatesDir specifies where to store or look for all required certificates
	CertificatesDir string `json:"certificatesDir"`

	// ImageRepository what container registry to pull control plane images from
	ImageRepository string `json:"imageRepository"`
	// UnifiedControlPlaneImage specifies if a specific container image should be used for all control plane components
	UnifiedControlPlaneImage string `json:"unifiedControlPlaneImage"`

	// FeatureFlags enabled by the user
	FeatureFlags map[string]bool `json:"featureFlags"`
}

type API struct {
	// AdvertiseAddress sets the address for the API server to advertise.
	AdvertiseAddress string `json:"advertiseAddress"`
	// BindPort sets the secure port for the API Server to bind to
	BindPort int32 `json:"bindPort"`
}

type TokenDiscovery struct {
	ID        string   `json:"id"`
	Secret    string   `json:"secret"`
	Addresses []string `json:"addresses"`
}

type Networking struct {
	ServiceSubnet string `json:"serviceSubnet"`
	PodSubnet     string `json:"podSubnet"`
	DNSDomain     string `json:"dnsDomain"`
}

type Etcd struct {
	Endpoints []string          `json:"endpoints"`
	CAFile    string            `json:"caFile"`
	CertFile  string            `json:"certFile"`
	KeyFile   string            `json:"keyFile"`
	DataDir   string            `json:"dataDir"`
	ExtraArgs map[string]string `json:"extraArgs"`
	// Image specifies which container image to use for running etcd. If empty, automatically populated by kubeadm using the image repository and default etcd version
	Image string `json:"image"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NodeConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	CACertPath               string   `json:"caCertPath"`
	DiscoveryFile            string   `json:"discoveryFile"`
	DiscoveryToken           string   `json:"discoveryToken"`
	DiscoveryTokenAPIServers []string `json:"discoveryTokenAPIServers"`
	NodeName                 string   `json:"nodeName"`
	TLSBootstrapToken        string   `json:"tlsBootstrapToken"`
	Token                    string   `json:"token"`
}
