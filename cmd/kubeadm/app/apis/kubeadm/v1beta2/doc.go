/*
Copyright 2019 The Kubernetes Authors.

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

// +k8s:defaulter-gen=TypeMeta
// +groupName=kubeadm.k8s.io
// +k8s:deepcopy-gen=package
// +k8s:conversion-gen=k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm

// Package v1beta2 has been DEPRECATED by v1beta3
//
// Package v1beta2 defines the v1beta2 version of the kubeadm configuration file format.
// This version improves on the v1beta1 format by fixing some minor issues and adding a few new fields.
//
// A list of changes since v1beta1:
//	- "certificateKey" field is added to InitConfiguration and JoinConfiguration.
//	- "ignorePreflightErrors" field is added to the NodeRegistrationOptions.
//	- The JSON "omitempty" tag is used in a more places where appropriate.
//	- The JSON "omitempty" tag of the "taints" field (inside NodeRegistrationOptions) is removed.
//	See the Kubernetes 1.15 changelog for further details.
//
// Migration from old kubeadm config versions
//
// Please convert your v1beta1 configuration files to v1beta2 using the "kubeadm config migrate" command of kubeadm v1.15.x
// (conversion from older releases of kubeadm config files requires older release of kubeadm as well e.g.
//	kubeadm v1.11 should be used to migrate v1alpha1 to v1alpha2; kubeadm v1.12 should be used to translate v1alpha2 to v1alpha3;
//	kubeadm v1.13 or v1.14 should be used to translate v1alpha3 to v1beta1)
//
// Nevertheless, kubeadm v1.15.x will support reading from v1beta1 version of the kubeadm config file format.
//
// Basics
//
// The preferred way to configure kubeadm is to pass an YAML configuration file with the --config option. Some of the
// configuration options defined in the kubeadm config file are also available as command line flags, but only
// the most common/simple use case are supported with this approach.
//
// A kubeadm config file could contain multiple configuration types separated using three dashes (“---”).
//
// kubeadm supports the following configuration types:
//
//     apiVersion: kubeadm.k8s.io/v1beta2
//     kind: InitConfiguration
//
//     apiVersion: kubeadm.k8s.io/v1beta2
//     kind: ClusterConfiguration
//
//     apiVersion: kubelet.config.k8s.io/v1beta1
//     kind: KubeletConfiguration
//
//     apiVersion: kubeproxy.config.k8s.io/v1alpha1
//     kind: KubeProxyConfiguration
//
//     apiVersion: kubeadm.k8s.io/v1beta2
//     kind: JoinConfiguration
//
// To print the defaults for "init" and "join" actions use the following commands:
//  kubeadm config print init-defaults
//  kubeadm config print join-defaults
//
// The list of configuration types that must be included in a configuration file depends by the action you are
// performing (init or join) and by the configuration options you are going to use (defaults or advanced customization).
//
// If some configuration types are not provided, or provided only partially, kubeadm will use default values; defaults
// provided by kubeadm includes also enforcing consistency of values across components when required (e.g.
// cluster-cidr flag on controller manager and clusterCIDR on kube-proxy).
//
// Users are always allowed to override default values, with the only exception of a small subset of setting with
// relevance for security (e.g. enforce authorization-mode Node and RBAC on api server)
//
// If the user provides a configuration types that is not expected for the action you are performing, kubeadm will
// ignore those types and print a warning.
//
// Kubeadm init configuration types
//
// When executing kubeadm init with the --config option, the following configuration types could be used:
// InitConfiguration, ClusterConfiguration, KubeProxyConfiguration, KubeletConfiguration, but only one
// between InitConfiguration and ClusterConfiguration is mandatory.
//
//     apiVersion: kubeadm.k8s.io/v1beta2
//     kind: InitConfiguration
//     bootstrapTokens:
//         ...
//     nodeRegistration:
//         ...
//
// The InitConfiguration type should be used to configure runtime settings, that in case of kubeadm init
// are the configuration of the bootstrap token and all the setting which are specific to the node where kubeadm
// is executed, including:
//
// - NodeRegistration, that holds fields that relate to registering the new node to the cluster;
// use it to customize the node name, the CRI socket to use or any other settings that should apply to this
// node only (e.g. the node ip).
//
// - LocalAPIEndpoint, that represents the endpoint of the instance of the API server to be deployed on this node;
// use it e.g. to customize the API server advertise address.
//
//     apiVersion: kubeadm.k8s.io/v1beta2
//     kind: ClusterConfiguration
//     networking:
//         ...
//     etcd:
//         ...
//     apiServer:
//       extraArgs:
//         ...
//       extraVolumes:
//         ...
//     ...
//
// The ClusterConfiguration type should be used to configure cluster-wide settings,
// including settings for:
//
// - Networking, that holds configuration for the networking topology of the cluster; use it e.g. to customize
// pod subnet or services subnet.
//
// - Etcd configurations; use it e.g. to customize the local etcd or to configure the API server
// for using an external etcd cluster.
//
// - kube-apiserver, kube-scheduler, kube-controller-manager configurations; use it to customize control-plane
// components by adding customized setting or overriding kubeadm default settings.
//
//    apiVersion: kubeproxy.config.k8s.io/v1alpha1
//    kind: KubeProxyConfiguration
//       ...
//
// The KubeProxyConfiguration type should be used to change the configuration passed to kube-proxy instances deployed
// in the cluster. If this object is not provided or provided only partially, kubeadm applies defaults.
//
// See https://kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/ or https://godoc.org/k8s.io/kube-proxy/config/v1alpha1#KubeProxyConfiguration
// for kube proxy official documentation.
//
//    apiVersion: kubelet.config.k8s.io/v1beta1
//    kind: KubeletConfiguration
//       ...
//
// The KubeletConfiguration type should be used to change the configurations that will be passed to all kubelet instances
// deployed in the cluster. If this object is not provided or provided only partially, kubeadm applies defaults.
//
// See https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/ or https://godoc.org/k8s.io/kubelet/config/v1beta1#KubeletConfiguration
// for kubelet official documentation.
//
// Here is a fully populated example of a single YAML file containing multiple
// configuration types to be used during a `kubeadm init` run.
//
// 	apiVersion: kubeadm.k8s.io/v1beta2
// 	kind: InitConfiguration
// 	bootstrapTokens:
// 	- token: "9a08jv.c0izixklcxtmnze7"
// 	  description: "kubeadm bootstrap token"
// 	  ttl: "24h"
// 	- token: "783bde.3f89s0fje9f38fhf"
// 	  description: "another bootstrap token"
// 	  usages:
// 	  - authentication
// 	  - signing
// 	  groups:
// 	  - system:bootstrappers:kubeadm:default-node-token
// 	nodeRegistration:
// 	  name: "ec2-10-100-0-1"
// 	  criSocket: "unix:///var/run/containerd/containerd.sock"
// 	  taints:
// 	  - key: "kubeadmNode"
// 	    value: "someValue"
// 	    effect: "NoSchedule"
// 	  kubeletExtraArgs:
// 	    v: 4
// 	  ignorePreflightErrors:
// 	  - IsPrivilegedUser
// 	localAPIEndpoint:
// 	  advertiseAddress: "10.100.0.1"
// 	  bindPort: 6443
// 	certificateKey: "e6a2eb8581237ab72a4f494f30285ec12a9694d750b9785706a83bfcbbbd2204"
// 	---
// 	apiVersion: kubeadm.k8s.io/v1beta2
// 	kind: ClusterConfiguration
// 	etcd:
// 	  # one of local or external
// 	  local:
// 	    imageRepository: "registry.k8s.io"
// 	    imageTag: "3.2.24"
// 	    dataDir: "/var/lib/etcd"
// 	    extraArgs:
// 	      listen-client-urls: "http://10.100.0.1:2379"
// 	    serverCertSANs:
// 	    -  "ec2-10-100-0-1.compute-1.amazonaws.com"
// 	    peerCertSANs:
// 	    - "10.100.0.1"
// 	  # external:
// 	    # endpoints:
// 	    # - "10.100.0.1:2379"
// 	    # - "10.100.0.2:2379"
// 	    # caFile: "/etcd/kubernetes/pki/etcd/etcd-ca.crt"
// 	    # certFile: "/etcd/kubernetes/pki/etcd/etcd.crt"
// 	    # keyFile: "/etcd/kubernetes/pki/etcd/etcd.key"
// 	networking:
// 	  serviceSubnet: "10.96.0.0/16"
// 	  podSubnet: "10.244.0.0/24"
// 	  dnsDomain: "cluster.local"
// 	kubernetesVersion: "v1.12.0"
// 	controlPlaneEndpoint: "10.100.0.1:6443"
// 	apiServer:
// 	  extraArgs:
// 	    authorization-mode: "Node,RBAC"
// 	  extraVolumes:
// 	  - name: "some-volume"
// 	    hostPath: "/etc/some-path"
// 	    mountPath: "/etc/some-pod-path"
// 	    readOnly: false
// 	    pathType: File
// 	  certSANs:
// 	  - "10.100.1.1"
// 	  - "ec2-10-100-0-1.compute-1.amazonaws.com"
// 	  timeoutForControlPlane: 4m0s
// 	controllerManager:
// 	  extraArgs:
// 	    "node-cidr-mask-size": "20"
// 	  extraVolumes:
// 	  - name: "some-volume"
// 	    hostPath: "/etc/some-path"
// 	    mountPath: "/etc/some-pod-path"
// 	    readOnly: false
// 	    pathType: File
// 	scheduler:
// 	  extraArgs:
// 	    address: "10.100.0.1"
// 	  extraVolumes:
// 	  - name: "some-volume"
// 	    hostPath: "/etc/some-path"
// 	    mountPath: "/etc/some-pod-path"
// 	    readOnly: false
// 	    pathType: File
// 	certificatesDir: "/etc/kubernetes/pki"
// 	imageRepository: "registry.k8s.io"
// 	useHyperKubeImage: false
// 	clusterName: "example-cluster"
// 	---
// 	apiVersion: kubelet.config.k8s.io/v1beta1
// 	kind: KubeletConfiguration
// 	# kubelet specific options here
// 	---
// 	apiVersion: kubeproxy.config.k8s.io/v1alpha1
// 	kind: KubeProxyConfiguration
// 	# kube-proxy specific options here
//
// Kubeadm join configuration types
//
// When executing kubeadm join with the --config option, the JoinConfiguration type should be provided.
//
//    apiVersion: kubeadm.k8s.io/v1beta2
//    kind: JoinConfiguration
//       ...
//
// The JoinConfiguration type should be used to configure runtime settings, that in case of kubeadm join
// are the discovery method used for accessing the cluster info and all the setting which are specific
// to the node where kubeadm is executed, including:
//
// - NodeRegistration, that holds fields that relate to registering the new node to the cluster;
// use it to customize the node name, the CRI socket to use or any other settings that should apply to this
// node only (e.g. the node ip).
//
// - APIEndpoint, that represents the endpoint of the instance of the API server to be eventually deployed on this node.
//
package v1beta2 // import "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta2"

//TODO: The BootstrapTokenString object should move out to either k8s.io/client-go or k8s.io/api in the future
//(probably as part of Bootstrap Tokens going GA). It should not be staged under the kubeadm API as it is now.
