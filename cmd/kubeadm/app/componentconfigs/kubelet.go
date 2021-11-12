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

package componentconfigs

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/version"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	kubeletconfig "k8s.io/kubelet/config/v1beta1"
	utilpointer "k8s.io/utils/pointer"

	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmapiv1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta3"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/initsystem"
)

const (
	// KubeletGroup is a pointer to the used API group name for the kubelet config
	KubeletGroup = kubeletconfig.GroupName

	// kubeletReadOnlyPort specifies the default insecure http server port
	// 0 will disable insecure http server.
	kubeletReadOnlyPort int32 = 0

	// kubeletRotateCertificates specifies the default value to enable certificate rotation
	kubeletRotateCertificates = true

	// kubeletAuthenticationAnonymousEnabled specifies the default value to disable anonymous access
	kubeletAuthenticationAnonymousEnabled = false

	// kubeletAuthenticationWebhookEnabled set the default value to enable authentication webhook
	kubeletAuthenticationWebhookEnabled = true

	// kubeletHealthzBindAddress specifies the default healthz bind address
	kubeletHealthzBindAddress = "127.0.0.1"

	// kubeletSystemdResolverConfig specifies the default resolver config when systemd service is active
	kubeletSystemdResolverConfig = "/run/systemd/resolve/resolv.conf"
)

// kubeletHandler is the handler instance for the kubelet component config
var kubeletHandler = handler{
	GroupVersion: kubeletconfig.SchemeGroupVersion,
	AddToScheme:  kubeletconfig.AddToScheme,
	CreateEmpty: func() kubeadmapi.ComponentConfig {
		return &kubeletConfig{
			configBase: configBase{
				GroupVersion: kubeletconfig.SchemeGroupVersion,
			},
		}
	},
	fromCluster: kubeletConfigFromCluster,
}

func kubeletConfigFromCluster(h *handler, clientset clientset.Interface, clusterCfg *kubeadmapi.ClusterConfiguration) (kubeadmapi.ComponentConfig, error) {
	// Read the ConfigMap from the cluster based on what version the kubelet is
	k8sVersion, err := version.ParseGeneric(clusterCfg.KubernetesVersion)
	if err != nil {
		return nil, err
	}

	// TODO: https://github.com/kubernetes/kubeadm/issues/1582
	// During the first "kubeadm upgrade apply" when the feature gate goes "true" by default and
	// a preferred user value is missing in the ClusterConfiguration, "kubeadm upgrade apply" will try
	// to fetch using the new format and that CM will not exist yet.
	// Tollerate both the old a new format until UnversionedKubeletConfigMap goes GA and is locked.
	// This makes it easier for the users and the code base (avoids changes in /cmd/upgrade/common.go#enforceRequirements).
	configMapNameLegacy := constants.GetKubeletConfigMapName(k8sVersion, true)
	configMapName := constants.GetKubeletConfigMapName(k8sVersion, false)
	klog.V(1).Infof("attempting to download the KubeletConfiguration from the new format location (UnversionedKubeletConfigMap=true)")
	cm, err := h.fromConfigMap(clientset, configMapName, constants.KubeletBaseConfigurationConfigMapKey, true)
	if err != nil {
		klog.V(1).Infof("attempting to download the KubeletConfiguration from the DEPRECATED location (UnversionedKubeletConfigMap=false)")
		cm, err = h.fromConfigMap(clientset, configMapNameLegacy, constants.KubeletBaseConfigurationConfigMapKey, true)
		if err != nil {
			return nil, errors.Wrapf(err, "could not download the kubelet configuration from ConfigMap %q or %q",
				configMapName, configMapNameLegacy)
		}
	}
	return cm, nil
}

// kubeletConfig implements the kubeadmapi.ComponentConfig interface for kubelet
type kubeletConfig struct {
	configBase
	config kubeletconfig.KubeletConfiguration
}

func (kc *kubeletConfig) DeepCopy() kubeadmapi.ComponentConfig {
	result := &kubeletConfig{}
	kc.configBase.DeepCopyInto(&result.configBase)
	kc.config.DeepCopyInto(&result.config)
	return result
}

func (kc *kubeletConfig) Marshal() ([]byte, error) {
	return kc.configBase.Marshal(&kc.config)
}

func (kc *kubeletConfig) Unmarshal(docmap kubeadmapi.DocumentMap) error {
	return kc.configBase.Unmarshal(docmap, &kc.config)
}

func (kc *kubeletConfig) Get() interface{} {
	return &kc.config
}

func (kc *kubeletConfig) Set(cfg interface{}) {
	kc.config = *cfg.(*kubeletconfig.KubeletConfiguration)
}

func (kc *kubeletConfig) Default(cfg *kubeadmapi.ClusterConfiguration, _ *kubeadmapi.APIEndpoint, nodeRegOpts *kubeadmapi.NodeRegistrationOptions) {
	const kind = "KubeletConfiguration"

	if kc.config.FeatureGates == nil {
		kc.config.FeatureGates = map[string]bool{}
	}

	if kc.config.StaticPodPath == "" {
		kc.config.StaticPodPath = kubeadmapiv1.DefaultManifestsDir
	} else if kc.config.StaticPodPath != kubeadmapiv1.DefaultManifestsDir {
		warnDefaultComponentConfigValue(kind, "staticPodPath", kubeadmapiv1.DefaultManifestsDir, kc.config.StaticPodPath)
	}

	clusterDNS := ""
	dnsIP, err := constants.GetDNSIP(cfg.Networking.ServiceSubnet)
	if err != nil {
		clusterDNS = kubeadmapiv1.DefaultClusterDNSIP
	} else {
		clusterDNS = dnsIP.String()
	}

	if kc.config.ClusterDNS == nil {
		kc.config.ClusterDNS = []string{clusterDNS}
	} else if len(kc.config.ClusterDNS) != 1 || kc.config.ClusterDNS[0] != clusterDNS {
		warnDefaultComponentConfigValue(kind, "clusterDNS", []string{clusterDNS}, kc.config.ClusterDNS)
	}

	if kc.config.ClusterDomain == "" {
		kc.config.ClusterDomain = cfg.Networking.DNSDomain
	} else if cfg.Networking.DNSDomain != "" && kc.config.ClusterDomain != cfg.Networking.DNSDomain {
		warnDefaultComponentConfigValue(kind, "clusterDomain", cfg.Networking.DNSDomain, kc.config.ClusterDomain)
	}

	// Require all clients to the kubelet API to have client certs signed by the cluster CA
	clientCAFile := filepath.Join(cfg.CertificatesDir, constants.CACertName)
	if kc.config.Authentication.X509.ClientCAFile == "" {
		kc.config.Authentication.X509.ClientCAFile = clientCAFile
	} else if kc.config.Authentication.X509.ClientCAFile != clientCAFile {
		warnDefaultComponentConfigValue(kind, "authentication.x509.clientCAFile", clientCAFile, kc.config.Authentication.X509.ClientCAFile)
	}

	if kc.config.Authentication.Anonymous.Enabled == nil {
		kc.config.Authentication.Anonymous.Enabled = utilpointer.BoolPtr(kubeletAuthenticationAnonymousEnabled)
	} else if *kc.config.Authentication.Anonymous.Enabled {
		warnDefaultComponentConfigValue(kind, "authentication.anonymous.enabled", kubeletAuthenticationAnonymousEnabled, *kc.config.Authentication.Anonymous.Enabled)
	}

	// On every client request to the kubelet API, execute a webhook (SubjectAccessReview request) to the API server
	// and ask it whether the client is authorized to access the kubelet API
	if kc.config.Authorization.Mode == "" {
		kc.config.Authorization.Mode = kubeletconfig.KubeletAuthorizationModeWebhook
	} else if kc.config.Authorization.Mode != kubeletconfig.KubeletAuthorizationModeWebhook {
		warnDefaultComponentConfigValue(kind, "authorization.mode", kubeletconfig.KubeletAuthorizationModeWebhook, kc.config.Authorization.Mode)
	}

	// Let clients using other authentication methods like ServiceAccount tokens also access the kubelet API
	if kc.config.Authentication.Webhook.Enabled == nil {
		kc.config.Authentication.Webhook.Enabled = utilpointer.BoolPtr(kubeletAuthenticationWebhookEnabled)
	} else if !*kc.config.Authentication.Webhook.Enabled {
		warnDefaultComponentConfigValue(kind, "authentication.webhook.enabled", kubeletAuthenticationWebhookEnabled, *kc.config.Authentication.Webhook.Enabled)
	}

	// Serve a /healthz webserver on localhost:10248 that kubeadm can talk to
	if kc.config.HealthzBindAddress == "" {
		kc.config.HealthzBindAddress = kubeletHealthzBindAddress
	} else if kc.config.HealthzBindAddress != kubeletHealthzBindAddress {
		warnDefaultComponentConfigValue(kind, "healthzBindAddress", kubeletHealthzBindAddress, kc.config.HealthzBindAddress)
	}

	if kc.config.HealthzPort == nil {
		kc.config.HealthzPort = utilpointer.Int32Ptr(constants.KubeletHealthzPort)
	} else if *kc.config.HealthzPort != constants.KubeletHealthzPort {
		warnDefaultComponentConfigValue(kind, "healthzPort", constants.KubeletHealthzPort, *kc.config.HealthzPort)
	}

	if kc.config.ReadOnlyPort != kubeletReadOnlyPort {
		warnDefaultComponentConfigValue(kind, "readOnlyPort", kubeletReadOnlyPort, kc.config.ReadOnlyPort)
	}

	// We cannot show a warning for RotateCertificates==false and we must hardcode it to true.
	// There is no way to determine if the user has set this or not, given the field is a non-pointer.
	kc.config.RotateCertificates = kubeletRotateCertificates

	if len(kc.config.CgroupDriver) == 0 {
		klog.V(1).Infof("the value of KubeletConfiguration.cgroupDriver is empty; setting it to %q", constants.CgroupDriverSystemd)
		kc.config.CgroupDriver = constants.CgroupDriverSystemd
	}

	ok, err := isServiceActive("systemd-resolved")
	if err != nil {
		klog.Warningf("cannot determine if systemd-resolved is active: %v", err)
	}
	if ok {
		if kc.config.ResolverConfig == nil {
			kc.config.ResolverConfig = utilpointer.String(kubeletSystemdResolverConfig)
		} else {
			if kc.config.ResolverConfig != utilpointer.String(kubeletSystemdResolverConfig) {
				warnDefaultComponentConfigValue(kind, "resolvConf", kubeletSystemdResolverConfig, *kc.config.ResolverConfig)
			}
		}
	}
}

// Mutate modifies absolute path fields in the KubeletConfiguration to be Windows compatible absolute paths.
func (kc *kubeletConfig) Mutate() error {
	// TODO: use build tags and move the Windows related logic to _windows.go files
	// once the kubeadm code base is unit tested for Windows as part of CI - "GOOS=windows go test ...".
	if runtime.GOOS != "windows" {
		return nil
	}

	// When "kubeadm join" downloads the KubeletConfiguration from the cluster on Windows
	// nodes, it would contain absolute paths that may lack drive letters, since the config
	// could have been generated on a Linux control-plane node. On Windows the
	// Golang path.IsAbs() function returns false unless the path contains a drive letter.
	// This trips client-go and the kubelet, creating problems on Windows nodes.
	// Fixing it in client-go or the kubelet is a breaking change to existing Windows
	// users that rely on relative paths:
	//   https://github.com/kubernetes/kubernetes/pull/77710#issuecomment-491989621
	//
	// Thus, a workaround here is to adapt the KubeletConfiguration paths for Windows.
	// Note this is currently bound to KubeletConfiguration v1beta1.
	klog.V(2).Infoln("[componentconfig] Adapting the paths in the KubeletConfiguration for Windows...")

	// Get the drive from where the kubeadm binary was called.
	exe, err := os.Executable()
	if err != nil {
		return errors.Wrap(err, "could not obtain information about the kubeadm executable")
	}
	drive := filepath.VolumeName(filepath.Dir(exe))
	klog.V(2).Infof("[componentconfig] Assuming Windows drive %q", drive)

	// Mutate the paths in the config.
	mutatePathsOnWindows(&kc.config, drive)
	return nil
}

func mutatePathsOnWindows(cfg *kubeletconfig.KubeletConfiguration, drive string) {
	mutateStringField := func(name string, field *string) {
		// path.IsAbs() is not reliable here in the Windows runtime, so check if the
		// path starts with "/" instead. This means the path originated from a Unix node and
		// is an absolute path.
		if !strings.HasPrefix(*field, "/") {
			return
		}
		// Prepend the drive letter to the path and update the field.
		*field = filepath.Join(drive, *field)
		klog.V(2).Infof("[componentconfig] kubelet/Windows: adapted path for field %q to %q", name, *field)
	}

	// Mutate the fields we care about.
	klog.V(2).Infof("[componentconfig] kubelet/Windows: changing field \"resolverConfig\" to empty")
	cfg.ResolverConfig = utilpointer.String("")
	mutateStringField("staticPodPath", &cfg.StaticPodPath)
	mutateStringField("authentication.x509.clientCAFile", &cfg.Authentication.X509.ClientCAFile)
}

// isServiceActive checks whether the given service exists and is running
func isServiceActive(name string) (bool, error) {
	initSystem, err := initsystem.GetInitSystem()
	if err != nil {
		return false, err
	}
	return initSystem.ServiceIsActive(name), nil
}
