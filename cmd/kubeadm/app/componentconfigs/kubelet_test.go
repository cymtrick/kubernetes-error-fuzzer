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
	"fmt"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/lithammer/dedent"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	clientsetfake "k8s.io/client-go/kubernetes/fake"
	kubeletconfig "k8s.io/kubelet/config/v1beta1"
	utilpointer "k8s.io/utils/pointer"

	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmapiv1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta3"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
)

// TODO: cleanup after UnversionedKubeletConfigMap goes GA:
// https://github.com/kubernetes/kubeadm/issues/1582
func testKubeletConfigMap(contents string, legacyKubeletConfigMap bool) *v1.ConfigMap {
	return &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      constants.GetKubeletConfigMapName(constants.CurrentKubernetesVersion, legacyKubeletConfigMap),
			Namespace: metav1.NamespaceSystem,
		},
		Data: map[string]string{
			constants.KubeletBaseConfigurationConfigMapKey: dedent.Dedent(contents),
		},
	}
}

func TestKubeletDefault(t *testing.T) {
	var resolverConfig *string
	if isSystemdResolvedActive, _ := isServiceActive("systemd-resolved"); isSystemdResolvedActive {
		// If systemd-resolved is active, we need to set the default resolver config
		resolverConfig = utilpointer.String(kubeletSystemdResolverConfig)
	}

	tests := []struct {
		name       string
		clusterCfg kubeadmapi.ClusterConfiguration
		expected   kubeletConfig
	}{
		{
			name:       "No specific defaulting works",
			clusterCfg: kubeadmapi.ClusterConfiguration{},
			expected: kubeletConfig{
				config: kubeletconfig.KubeletConfiguration{
					FeatureGates:  map[string]bool{},
					StaticPodPath: kubeadmapiv1.DefaultManifestsDir,
					ClusterDNS:    []string{kubeadmapiv1.DefaultClusterDNSIP},
					Authentication: kubeletconfig.KubeletAuthentication{
						X509: kubeletconfig.KubeletX509Authentication{
							ClientCAFile: constants.CACertName,
						},
						Anonymous: kubeletconfig.KubeletAnonymousAuthentication{
							Enabled: utilpointer.BoolPtr(kubeletAuthenticationAnonymousEnabled),
						},
						Webhook: kubeletconfig.KubeletWebhookAuthentication{
							Enabled: utilpointer.BoolPtr(kubeletAuthenticationWebhookEnabled),
						},
					},
					Authorization: kubeletconfig.KubeletAuthorization{
						Mode: kubeletconfig.KubeletAuthorizationModeWebhook,
					},
					HealthzBindAddress: kubeletHealthzBindAddress,
					HealthzPort:        utilpointer.Int32Ptr(constants.KubeletHealthzPort),
					RotateCertificates: kubeletRotateCertificates,
					ResolverConfig:     resolverConfig,
					CgroupDriver:       constants.CgroupDriverSystemd,
				},
			},
		},
		{
			name: "Service subnet, no dual stack defaulting works",
			clusterCfg: kubeadmapi.ClusterConfiguration{
				Networking: kubeadmapi.Networking{
					ServiceSubnet: "192.168.0.0/16",
				},
			},
			expected: kubeletConfig{
				config: kubeletconfig.KubeletConfiguration{
					FeatureGates:  map[string]bool{},
					StaticPodPath: kubeadmapiv1.DefaultManifestsDir,
					ClusterDNS:    []string{"192.168.0.10"},
					Authentication: kubeletconfig.KubeletAuthentication{
						X509: kubeletconfig.KubeletX509Authentication{
							ClientCAFile: constants.CACertName,
						},
						Anonymous: kubeletconfig.KubeletAnonymousAuthentication{
							Enabled: utilpointer.BoolPtr(kubeletAuthenticationAnonymousEnabled),
						},
						Webhook: kubeletconfig.KubeletWebhookAuthentication{
							Enabled: utilpointer.BoolPtr(kubeletAuthenticationWebhookEnabled),
						},
					},
					Authorization: kubeletconfig.KubeletAuthorization{
						Mode: kubeletconfig.KubeletAuthorizationModeWebhook,
					},
					HealthzBindAddress: kubeletHealthzBindAddress,
					HealthzPort:        utilpointer.Int32Ptr(constants.KubeletHealthzPort),
					RotateCertificates: kubeletRotateCertificates,
					ResolverConfig:     resolverConfig,
					CgroupDriver:       constants.CgroupDriverSystemd,
				},
			},
		},
		{
			name: "Service subnet, enabled dual stack defaulting works",
			clusterCfg: kubeadmapi.ClusterConfiguration{
				Networking: kubeadmapi.Networking{
					ServiceSubnet: "192.168.0.0/16",
				},
			},
			expected: kubeletConfig{
				config: kubeletconfig.KubeletConfiguration{
					FeatureGates:  map[string]bool{},
					StaticPodPath: kubeadmapiv1.DefaultManifestsDir,
					ClusterDNS:    []string{"192.168.0.10"},
					Authentication: kubeletconfig.KubeletAuthentication{
						X509: kubeletconfig.KubeletX509Authentication{
							ClientCAFile: constants.CACertName,
						},
						Anonymous: kubeletconfig.KubeletAnonymousAuthentication{
							Enabled: utilpointer.BoolPtr(kubeletAuthenticationAnonymousEnabled),
						},
						Webhook: kubeletconfig.KubeletWebhookAuthentication{
							Enabled: utilpointer.BoolPtr(kubeletAuthenticationWebhookEnabled),
						},
					},
					Authorization: kubeletconfig.KubeletAuthorization{
						Mode: kubeletconfig.KubeletAuthorizationModeWebhook,
					},
					HealthzBindAddress: kubeletHealthzBindAddress,
					HealthzPort:        utilpointer.Int32Ptr(constants.KubeletHealthzPort),
					RotateCertificates: kubeletRotateCertificates,
					ResolverConfig:     resolverConfig,
					CgroupDriver:       constants.CgroupDriverSystemd,
				},
			},
		},
		{
			name: "DNS domain defaulting works",
			clusterCfg: kubeadmapi.ClusterConfiguration{
				Networking: kubeadmapi.Networking{
					DNSDomain: "example.com",
				},
			},
			expected: kubeletConfig{
				config: kubeletconfig.KubeletConfiguration{
					FeatureGates:  map[string]bool{},
					StaticPodPath: kubeadmapiv1.DefaultManifestsDir,
					ClusterDNS:    []string{kubeadmapiv1.DefaultClusterDNSIP},
					ClusterDomain: "example.com",
					Authentication: kubeletconfig.KubeletAuthentication{
						X509: kubeletconfig.KubeletX509Authentication{
							ClientCAFile: constants.CACertName,
						},
						Anonymous: kubeletconfig.KubeletAnonymousAuthentication{
							Enabled: utilpointer.BoolPtr(kubeletAuthenticationAnonymousEnabled),
						},
						Webhook: kubeletconfig.KubeletWebhookAuthentication{
							Enabled: utilpointer.BoolPtr(kubeletAuthenticationWebhookEnabled),
						},
					},
					Authorization: kubeletconfig.KubeletAuthorization{
						Mode: kubeletconfig.KubeletAuthorizationModeWebhook,
					},
					HealthzBindAddress: kubeletHealthzBindAddress,
					HealthzPort:        utilpointer.Int32Ptr(constants.KubeletHealthzPort),
					RotateCertificates: kubeletRotateCertificates,
					ResolverConfig:     resolverConfig,
					CgroupDriver:       constants.CgroupDriverSystemd,
				},
			},
		},
		{
			name: "CertificatesDir defaulting works",
			clusterCfg: kubeadmapi.ClusterConfiguration{
				CertificatesDir: "/path/to/certs",
			},
			expected: kubeletConfig{
				config: kubeletconfig.KubeletConfiguration{
					FeatureGates:  map[string]bool{},
					StaticPodPath: kubeadmapiv1.DefaultManifestsDir,
					ClusterDNS:    []string{kubeadmapiv1.DefaultClusterDNSIP},
					Authentication: kubeletconfig.KubeletAuthentication{
						X509: kubeletconfig.KubeletX509Authentication{
							ClientCAFile: filepath.Join("/path/to/certs", constants.CACertName),
						},
						Anonymous: kubeletconfig.KubeletAnonymousAuthentication{
							Enabled: utilpointer.BoolPtr(kubeletAuthenticationAnonymousEnabled),
						},
						Webhook: kubeletconfig.KubeletWebhookAuthentication{
							Enabled: utilpointer.BoolPtr(kubeletAuthenticationWebhookEnabled),
						},
					},
					Authorization: kubeletconfig.KubeletAuthorization{
						Mode: kubeletconfig.KubeletAuthorizationModeWebhook,
					},
					HealthzBindAddress: kubeletHealthzBindAddress,
					HealthzPort:        utilpointer.Int32Ptr(constants.KubeletHealthzPort),
					RotateCertificates: kubeletRotateCertificates,
					ResolverConfig:     resolverConfig,
					CgroupDriver:       constants.CgroupDriverSystemd,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// This is the same for all test cases so we set it here
			expected := test.expected
			expected.configBase.GroupVersion = kubeletconfig.SchemeGroupVersion

			got := &kubeletConfig{
				configBase: configBase{
					GroupVersion: kubeletconfig.SchemeGroupVersion,
				},
			}
			got.Default(&test.clusterCfg, &kubeadmapi.APIEndpoint{}, &kubeadmapi.NodeRegistrationOptions{})

			if !reflect.DeepEqual(got, &expected) {
				t.Fatalf("Missmatch between expected and got:\nExpected:\n%v\n---\nGot:\n%v", expected, *got)
			}
		})
	}
}

// runKubeletFromTest holds common test case data and evaluation code for kubeletHandler.From* functions
func runKubeletFromTest(t *testing.T, perform func(gvk schema.GroupVersionKind, yaml string) (kubeadmapi.ComponentConfig, error)) {
	const (
		kind          = "KubeletConfiguration"
		clusterDomain = "foo.bar"
	)

	gvk := kubeletHandler.GroupVersion.WithKind(kind)
	yaml := fmt.Sprintf("apiVersion: %s\nkind: %s\nclusterDomain: %s", kubeletHandler.GroupVersion, kind, clusterDomain)

	cfg, err := perform(gvk, yaml)

	if err != nil {
		t.Fatalf("unexpected failure: %v", err)
	}
	if cfg == nil {
		t.Fatal("no config loaded where it should have been")
	}
	if kubeletCfg, ok := cfg.(*kubeletConfig); !ok {
		t.Fatalf("found different object type than expected: %s", reflect.TypeOf(cfg))
	} else if kubeletCfg.config.ClusterDomain != clusterDomain {
		t.Fatalf("unexpected control value (clusterDomain):\n\tgot: %q\n\texpected: %q", kubeletCfg.config.ClusterDomain, clusterDomain)
	}
}

func TestKubeletFromDocumentMap(t *testing.T) {
	runKubeletFromTest(t, func(gvk schema.GroupVersionKind, yaml string) (kubeadmapi.ComponentConfig, error) {
		return kubeletHandler.FromDocumentMap(kubeadmapi.DocumentMap{
			gvk: []byte(yaml),
		})
	})
}

func TestKubeletFromCluster(t *testing.T) {
	runKubeletFromTest(t, func(_ schema.GroupVersionKind, yaml string) (kubeadmapi.ComponentConfig, error) {
		client := clientsetfake.NewSimpleClientset(
			testKubeletConfigMap(yaml, true),
		)
		legacyKubeletConfigMap := true
		return kubeletHandler.FromCluster(client, testClusterCfg(legacyKubeletConfigMap))
	})
	runKubeletFromTest(t, func(_ schema.GroupVersionKind, yaml string) (kubeadmapi.ComponentConfig, error) {
		client := clientsetfake.NewSimpleClientset(
			testKubeletConfigMap(yaml, false),
		)
		legacyKubeletConfigMap := false
		return kubeletHandler.FromCluster(client, testClusterCfg(legacyKubeletConfigMap))
	})
}

func TestMutatePathsOnWindows(t *testing.T) {
	const drive = "C:"
	var fooResolverConfig string = "/foo/resolver"

	tests := []struct {
		name     string
		cfg      *kubeletconfig.KubeletConfiguration
		expected *kubeletconfig.KubeletConfiguration
	}{
		{
			name: "valid: all fields are absolute paths",
			cfg: &kubeletconfig.KubeletConfiguration{
				ResolverConfig: &fooResolverConfig,
				StaticPodPath:  "/foo/staticpods",
				Authentication: kubeletconfig.KubeletAuthentication{
					X509: kubeletconfig.KubeletX509Authentication{
						ClientCAFile: "/foo/ca.crt",
					},
				},
			},
			expected: &kubeletconfig.KubeletConfiguration{
				ResolverConfig: utilpointer.String(""),
				StaticPodPath:  filepath.Join(drive, "/foo/staticpods"),
				Authentication: kubeletconfig.KubeletAuthentication{
					X509: kubeletconfig.KubeletX509Authentication{
						ClientCAFile: filepath.Join(drive, "/foo/ca.crt"),
					},
				},
			},
		},
		{
			name: "valid: some fields are not absolute paths",
			cfg: &kubeletconfig.KubeletConfiguration{
				ResolverConfig: &fooResolverConfig,
				StaticPodPath:  "./foo/staticpods", // not an absolute Unix path
				Authentication: kubeletconfig.KubeletAuthentication{
					X509: kubeletconfig.KubeletX509Authentication{
						ClientCAFile: "/foo/ca.crt",
					},
				},
			},
			expected: &kubeletconfig.KubeletConfiguration{
				ResolverConfig: utilpointer.String(""),
				StaticPodPath:  "./foo/staticpods",
				Authentication: kubeletconfig.KubeletAuthentication{
					X509: kubeletconfig.KubeletX509Authentication{
						ClientCAFile: filepath.Join(drive, "/foo/ca.crt"),
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mutatePathsOnWindows(test.cfg, drive)
			if !reflect.DeepEqual(test.cfg, test.expected) {
				t.Errorf("Missmatch between expected and got:\nExpected:\n%+v\n---\nGot:\n%+v",
					test.expected, test.cfg)
			}
		})
	}
}
