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

package controlplane

import (
	"fmt"
	"net"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/features"
	"k8s.io/kubernetes/cmd/kubeadm/app/images"
	certphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/certs"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	staticpodutil "k8s.io/kubernetes/cmd/kubeadm/app/util/staticpod"
	utilsnet "k8s.io/utils/net"
)

// CreateInitStaticPodManifestFiles will write all static pod manifest files needed to bring up the control plane.
func CreateInitStaticPodManifestFiles(manifestDir, kustomizeDir string, cfg *kubeadmapi.InitConfiguration) error {
	klog.V(1).Infoln("[control-plane] creating static Pod files")
	return CreateStaticPodFiles(manifestDir, kustomizeDir, &cfg.ClusterConfiguration, &cfg.LocalAPIEndpoint, kubeadmconstants.KubeAPIServer, kubeadmconstants.KubeControllerManager, kubeadmconstants.KubeScheduler)
}

// GetStaticPodSpecs returns all staticPodSpecs actualized to the context of the current configuration
// NB. this methods holds the information about how kubeadm creates static pod manifests.
func GetStaticPodSpecs(cfg *kubeadmapi.ClusterConfiguration, endpoint *kubeadmapi.APIEndpoint) map[string]v1.Pod {
	// Get the required hostpath mounts
	mounts := getHostPathVolumesForTheControlPlane(cfg)

	// Prepare static pod specs
	staticPodSpecs := map[string]v1.Pod{
		kubeadmconstants.KubeAPIServer: staticpodutil.ComponentPod(v1.Container{
			Name:            kubeadmconstants.KubeAPIServer,
			Image:           images.GetKubernetesImage(kubeadmconstants.KubeAPIServer, cfg),
			ImagePullPolicy: v1.PullIfNotPresent,
			Command:         getAPIServerCommand(cfg, endpoint),
			VolumeMounts:    staticpodutil.VolumeMountMapToSlice(mounts.GetVolumeMounts(kubeadmconstants.KubeAPIServer)),
			LivenessProbe:   staticpodutil.LivenessProbe(staticpodutil.GetAPIServerProbeAddress(endpoint), "/livez", int(endpoint.BindPort), v1.URISchemeHTTPS),
			ReadinessProbe:  staticpodutil.ReadinessProbe(staticpodutil.GetAPIServerProbeAddress(endpoint), "/readyz", int(endpoint.BindPort), v1.URISchemeHTTPS),
			Resources:       staticpodutil.ComponentResources("250m"),
			Env:             kubeadmutil.GetProxyEnvVars(),
		}, mounts.GetVolumes(kubeadmconstants.KubeAPIServer),
			map[string]string{kubeadmconstants.KubeAPIServerAdvertiseAddressEndpointAnnotationKey: endpoint.String()}),
		kubeadmconstants.KubeControllerManager: staticpodutil.ComponentPod(v1.Container{
			Name:            kubeadmconstants.KubeControllerManager,
			Image:           images.GetKubernetesImage(kubeadmconstants.KubeControllerManager, cfg),
			ImagePullPolicy: v1.PullIfNotPresent,
			Command:         getControllerManagerCommand(cfg),
			VolumeMounts:    staticpodutil.VolumeMountMapToSlice(mounts.GetVolumeMounts(kubeadmconstants.KubeControllerManager)),
			LivenessProbe:   staticpodutil.LivenessProbe(staticpodutil.GetControllerManagerProbeAddress(cfg), "/healthz", kubeadmconstants.KubeControllerManagerPort, v1.URISchemeHTTPS),
			Resources:       staticpodutil.ComponentResources("200m"),
			Env:             kubeadmutil.GetProxyEnvVars(),
		}, mounts.GetVolumes(kubeadmconstants.KubeControllerManager), nil),
		kubeadmconstants.KubeScheduler: staticpodutil.ComponentPod(v1.Container{
			Name:            kubeadmconstants.KubeScheduler,
			Image:           images.GetKubernetesImage(kubeadmconstants.KubeScheduler, cfg),
			ImagePullPolicy: v1.PullIfNotPresent,
			Command:         getSchedulerCommand(cfg),
			VolumeMounts:    staticpodutil.VolumeMountMapToSlice(mounts.GetVolumeMounts(kubeadmconstants.KubeScheduler)),
			LivenessProbe:   staticpodutil.LivenessProbe(staticpodutil.GetSchedulerProbeAddress(cfg), "/healthz", kubeadmconstants.KubeSchedulerPort, v1.URISchemeHTTPS),
			Resources:       staticpodutil.ComponentResources("100m"),
			Env:             kubeadmutil.GetProxyEnvVars(),
		}, mounts.GetVolumes(kubeadmconstants.KubeScheduler), nil),
	}
	return staticPodSpecs
}

// CreateStaticPodFiles creates all the requested static pod files.
func CreateStaticPodFiles(manifestDir, kustomizeDir string, cfg *kubeadmapi.ClusterConfiguration, endpoint *kubeadmapi.APIEndpoint, componentNames ...string) error {
	// gets the StaticPodSpecs, actualized for the current ClusterConfiguration
	klog.V(1).Infoln("[control-plane] getting StaticPodSpecs")
	specs := GetStaticPodSpecs(cfg, endpoint)

	// creates required static pod specs
	for _, componentName := range componentNames {
		// retrieves the StaticPodSpec for given component
		spec, exists := specs[componentName]
		if !exists {
			return errors.Errorf("couldn't retrieve StaticPodSpec for %q", componentName)
		}

		// print all volumes that are mounted
		for _, v := range spec.Spec.Volumes {
			klog.V(2).Infof("[control-plane] adding volume %q for component %q", v.Name, componentName)
		}

		// if kustomizeDir is defined, customize the static pod manifest
		if kustomizeDir != "" {
			kustomizedSpec, err := staticpodutil.KustomizeStaticPod(&spec, kustomizeDir)
			if err != nil {
				return errors.Wrapf(err, "failed to kustomize static pod manifest file for %q", componentName)
			}
			spec = *kustomizedSpec
		}

		// writes the StaticPodSpec to disk
		if err := staticpodutil.WriteStaticPodToDisk(componentName, manifestDir, spec); err != nil {
			return errors.Wrapf(err, "failed to create static pod manifest file for %q", componentName)
		}

		klog.V(1).Infof("[control-plane] wrote static Pod manifest for component %q to %q\n", componentName, kubeadmconstants.GetStaticPodFilepath(componentName, manifestDir))
	}

	return nil
}

// getAPIServerCommand builds the right API server command from the given config object and version
func getAPIServerCommand(cfg *kubeadmapi.ClusterConfiguration, localAPIEndpoint *kubeadmapi.APIEndpoint) []string {
	defaultArguments := map[string]string{
		"advertise-address":               localAPIEndpoint.AdvertiseAddress,
		"insecure-port":                   "0",
		"enable-admission-plugins":        "NodeRestriction",
		"service-cluster-ip-range":        cfg.Networking.ServiceSubnet,
		"service-account-key-file":        filepath.Join(cfg.CertificatesDir, kubeadmconstants.ServiceAccountPublicKeyName),
		"client-ca-file":                  filepath.Join(cfg.CertificatesDir, kubeadmconstants.CACertName),
		"tls-cert-file":                   filepath.Join(cfg.CertificatesDir, kubeadmconstants.APIServerCertName),
		"tls-private-key-file":            filepath.Join(cfg.CertificatesDir, kubeadmconstants.APIServerKeyName),
		"kubelet-client-certificate":      filepath.Join(cfg.CertificatesDir, kubeadmconstants.APIServerKubeletClientCertName),
		"kubelet-client-key":              filepath.Join(cfg.CertificatesDir, kubeadmconstants.APIServerKubeletClientKeyName),
		"enable-bootstrap-token-auth":     "true",
		"secure-port":                     fmt.Sprintf("%d", localAPIEndpoint.BindPort),
		"allow-privileged":                "true",
		"kubelet-preferred-address-types": "InternalIP,ExternalIP,Hostname",
		// add options to configure the front proxy.  Without the generated client cert, this will never be useable
		// so add it unconditionally with recommended values
		"requestheader-username-headers":     "X-Remote-User",
		"requestheader-group-headers":        "X-Remote-Group",
		"requestheader-extra-headers-prefix": "X-Remote-Extra-",
		"requestheader-client-ca-file":       filepath.Join(cfg.CertificatesDir, kubeadmconstants.FrontProxyCACertName),
		"requestheader-allowed-names":        "front-proxy-client",
		"proxy-client-cert-file":             filepath.Join(cfg.CertificatesDir, kubeadmconstants.FrontProxyClientCertName),
		"proxy-client-key-file":              filepath.Join(cfg.CertificatesDir, kubeadmconstants.FrontProxyClientKeyName),
	}

	command := []string{"kube-apiserver"}

	// If the user set endpoints for an external etcd cluster
	if cfg.Etcd.External != nil {
		defaultArguments["etcd-servers"] = strings.Join(cfg.Etcd.External.Endpoints, ",")

		// Use any user supplied etcd certificates
		if cfg.Etcd.External.CAFile != "" {
			defaultArguments["etcd-cafile"] = cfg.Etcd.External.CAFile
		}
		if cfg.Etcd.External.CertFile != "" && cfg.Etcd.External.KeyFile != "" {
			defaultArguments["etcd-certfile"] = cfg.Etcd.External.CertFile
			defaultArguments["etcd-keyfile"] = cfg.Etcd.External.KeyFile
		}
	} else {
		// Default to etcd static pod on localhost
		// localhost IP family should be the same that the AdvertiseAddress
		etcdLocalhostAddress := "127.0.0.1"
		if utilsnet.IsIPv6String(localAPIEndpoint.AdvertiseAddress) {
			etcdLocalhostAddress = "::1"
		}
		defaultArguments["etcd-servers"] = fmt.Sprintf("https://%s", net.JoinHostPort(etcdLocalhostAddress, strconv.Itoa(kubeadmconstants.EtcdListenClientPort)))
		defaultArguments["etcd-cafile"] = filepath.Join(cfg.CertificatesDir, kubeadmconstants.EtcdCACertName)
		defaultArguments["etcd-certfile"] = filepath.Join(cfg.CertificatesDir, kubeadmconstants.APIServerEtcdClientCertName)
		defaultArguments["etcd-keyfile"] = filepath.Join(cfg.CertificatesDir, kubeadmconstants.APIServerEtcdClientKeyName)

		// Apply user configurations for local etcd
		if cfg.Etcd.Local != nil {
			if value, ok := cfg.Etcd.Local.ExtraArgs["advertise-client-urls"]; ok {
				defaultArguments["etcd-servers"] = value
			}
		}
	}

	// TODO: The following code should be removed after dual-stack is GA.
	// Note: The user still retains the ability to explicitly set feature-gates and that value will overwrite this base value.
	if enabled, present := cfg.FeatureGates[features.IPv6DualStack]; present {
		defaultArguments["feature-gates"] = fmt.Sprintf("%s=%t", features.IPv6DualStack, enabled)
	}

	if cfg.APIServer.ExtraArgs == nil {
		cfg.APIServer.ExtraArgs = map[string]string{}
	}
	cfg.APIServer.ExtraArgs["authorization-mode"] = getAuthzModes(cfg.APIServer.ExtraArgs["authorization-mode"])
	command = append(command, kubeadmutil.BuildArgumentListFromMap(defaultArguments, cfg.APIServer.ExtraArgs)...)

	return command
}

// getAuthzModes gets the authorization-related parameters to the api server
// Node,RBAC is the default mode if nothing is passed to kubeadm. User provided modes override
// the default.
func getAuthzModes(authzModeExtraArgs string) string {
	defaultMode := []string{
		kubeadmconstants.ModeNode,
		kubeadmconstants.ModeRBAC,
	}

	if len(authzModeExtraArgs) > 0 {
		mode := []string{}
		for _, requested := range strings.Split(authzModeExtraArgs, ",") {
			if isValidAuthzMode(requested) {
				mode = append(mode, requested)
			} else {
				klog.Warningf("ignoring unknown kube-apiserver authorization-mode %q", requested)
			}
		}

		// only return the user provided mode if at least one was valid
		if len(mode) > 0 {
			if !compareAuthzModes(defaultMode, mode) {
				klog.Warningf("the default kube-apiserver authorization-mode is %q; using %q",
					strings.Join(defaultMode, ","),
					strings.Join(mode, ","),
				)
			}
			return strings.Join(mode, ",")
		}
	}
	return strings.Join(defaultMode, ",")
}

// compareAuthzModes compares two given authz modes and returns false if they do not match
func compareAuthzModes(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, m := range a {
		if m != b[i] {
			return false
		}
	}
	return true
}

func isValidAuthzMode(authzMode string) bool {
	allModes := []string{
		kubeadmconstants.ModeNode,
		kubeadmconstants.ModeRBAC,
		kubeadmconstants.ModeWebhook,
		kubeadmconstants.ModeABAC,
		kubeadmconstants.ModeAlwaysAllow,
		kubeadmconstants.ModeAlwaysDeny,
	}

	for _, mode := range allModes {
		if authzMode == mode {
			return true
		}
	}
	return false
}

// calcNodeCidrSize determines the size of the subnets used on each node, based
// on the pod subnet provided.  For IPv4, we assume that the pod subnet will
// be /16 and use /24. If the pod subnet cannot be parsed, the IPv4 value will
// be used (/24).
//
// For IPv6, the algorithm will do two three. First, the node CIDR will be set
// to a multiple of 8, using the available bits for easier readability by user.
// Second, the number of nodes will be 512 to 64K to attempt to maximize the
// number of nodes (see NOTE below). Third, pod networks of /113 and larger will
// be rejected, as the amount of bits available is too small.
//
// A special case is when the pod network size is /112, where /120 will be used,
// only allowing 256 nodes and 256 pods.
//
// If the pod network size is /113 or larger, the node CIDR will be set to the same
// size and this will be rejected later in validation.
//
// NOTE: Currently, the design allows a maximum of 64K nodes. This algorithm splits
// the available bits to maximize the number used for nodes, but still have the node
// CIDR be a multiple of eight.
//
func calcNodeCidrSize(podSubnet string) (string, bool) {
	maskSize := "24"
	isIPv6 := false
	if ip, podCidr, err := net.ParseCIDR(podSubnet); err == nil {
		if utilsnet.IsIPv6(ip) {
			var nodeCidrSize int
			isIPv6 = true
			podNetSize, totalBits := podCidr.Mask.Size()
			switch {
			case podNetSize == 112:
				// Special case, allows 256 nodes, 256 pods/node
				nodeCidrSize = 120
			case podNetSize < 112:
				// Use multiple of 8 for node CIDR, with 512 to 64K nodes
				nodeCidrSize = totalBits - ((totalBits-podNetSize-1)/8-1)*8
			default:
				// Not enough bits, will fail later, when validate
				nodeCidrSize = podNetSize
			}
			maskSize = strconv.Itoa(nodeCidrSize)
		}
	}
	return maskSize, isIPv6
}

// getControllerManagerCommand builds the right controller manager command from the given config object and version
func getControllerManagerCommand(cfg *kubeadmapi.ClusterConfiguration) []string {

	kubeconfigFile := filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.ControllerManagerKubeConfigFileName)
	caFile := filepath.Join(cfg.CertificatesDir, kubeadmconstants.CACertName)

	defaultArguments := map[string]string{
		"bind-address":                     "127.0.0.1",
		"leader-elect":                     "true",
		"kubeconfig":                       kubeconfigFile,
		"authentication-kubeconfig":        kubeconfigFile,
		"authorization-kubeconfig":         kubeconfigFile,
		"client-ca-file":                   caFile,
		"requestheader-client-ca-file":     filepath.Join(cfg.CertificatesDir, kubeadmconstants.FrontProxyCACertName),
		"root-ca-file":                     caFile,
		"service-account-private-key-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.ServiceAccountPrivateKeyName),
		"cluster-signing-cert-file":        caFile,
		"cluster-signing-key-file":         filepath.Join(cfg.CertificatesDir, kubeadmconstants.CAKeyName),
		"use-service-account-credentials":  "true",
		"controllers":                      "*,bootstrapsigner,tokencleaner",
	}

	// If using external CA, pass empty string to controller manager instead of ca.key/ca.crt path,
	// so that the csrsigning controller fails to start
	if res, _ := certphase.UsingExternalCA(cfg); res {
		defaultArguments["cluster-signing-key-file"] = ""
		defaultArguments["cluster-signing-cert-file"] = ""
	}

	// Let the controller-manager allocate Node CIDRs for the Pod network.
	// Each node will get a subspace of the address CIDR provided with --pod-network-cidr.
	if cfg.Networking.PodSubnet != "" {
		defaultArguments["allocate-node-cidrs"] = "true"
		defaultArguments["cluster-cidr"] = cfg.Networking.PodSubnet
		if cfg.Networking.ServiceSubnet != "" {
			defaultArguments["service-cluster-ip-range"] = cfg.Networking.ServiceSubnet
		}
	}

	// Set cluster name
	if cfg.ClusterName != "" {
		defaultArguments["cluster-name"] = cfg.ClusterName
	}

	// TODO: The following code should be remvoved after dual-stack is GA.
	// Note: The user still retains the ability to explicitly set feature-gates and that value will overwrite this base value.
	enabled, present := cfg.FeatureGates[features.IPv6DualStack]
	if present {
		defaultArguments["feature-gates"] = fmt.Sprintf("%s=%t", features.IPv6DualStack, enabled)
	}
	if cfg.Networking.PodSubnet != "" {
		if enabled {
			// any errors will be caught during validation
			subnets := strings.Split(cfg.Networking.PodSubnet, ",")
			for _, podSubnet := range subnets {
				if maskSize, isIPv6 := calcNodeCidrSize(podSubnet); isIPv6 {
					defaultArguments["node-cidr-mask-size-ipv6"] = maskSize
				} else {
					defaultArguments["node-cidr-mask-size-ipv4"] = maskSize
				}
			}
		} else {
			maskSize, _ := calcNodeCidrSize(cfg.Networking.PodSubnet)
			defaultArguments["node-cidr-mask-size"] = maskSize
		}
	}

	command := []string{"kube-controller-manager"}
	command = append(command, kubeadmutil.BuildArgumentListFromMap(defaultArguments, cfg.ControllerManager.ExtraArgs)...)

	return command
}

// getSchedulerCommand builds the right scheduler command from the given config object and version
func getSchedulerCommand(cfg *kubeadmapi.ClusterConfiguration) []string {
	kubeconfigFile := filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.SchedulerKubeConfigFileName)
	defaultArguments := map[string]string{
		"bind-address":              "127.0.0.1",
		"leader-elect":              "true",
		"kubeconfig":                kubeconfigFile,
		"authentication-kubeconfig": kubeconfigFile,
		"authorization-kubeconfig":  kubeconfigFile,
	}

	// TODO: The following code should be remvoved after dual-stack is GA.
	// Note: The user still retains the ability to explicitly set feature-gates and that value will overwrite this base value.
	if enabled, present := cfg.FeatureGates[features.IPv6DualStack]; present {
		defaultArguments["feature-gates"] = fmt.Sprintf("%s=%t", features.IPv6DualStack, enabled)
	}

	command := []string{"kube-scheduler"}
	command = append(command, kubeadmutil.BuildArgumentListFromMap(defaultArguments, cfg.Scheduler.ExtraArgs)...)
	return command
}
