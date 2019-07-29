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

package file

import (
	"github.com/pkg/errors"

	"k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	bootstrapapi "k8s.io/cluster-bootstrap/token/api"
	"k8s.io/klog"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeconfigutil "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig"
)

// RetrieveValidatedConfigInfo connects to the API Server and makes sure it can talk
// securely to the API Server using the provided CA cert and
// optionally refreshes the cluster-info information from the cluster-info ConfigMap
func RetrieveValidatedConfigInfo(filepath, clustername string) (*clientcmdapi.Config, error) {
	config, err := clientcmd.LoadFromFile(filepath)
	if err != nil {
		return nil, err
	}
	return ValidateConfigInfo(config, clustername)
}

// ValidateConfigInfo connects to the API Server and makes sure it can talk
// securely to the API Server using the provided CA cert/client certificates  and
// optionally refreshes the cluster-info information from the cluster-info ConfigMap
func ValidateConfigInfo(config *clientcmdapi.Config, clustername string) (*clientcmdapi.Config, error) {
	err := validateKubeConfig(config)
	if err != nil {
		return nil, err
	}

	var kubeconfig *clientcmdapi.Config

	// If the discovery file config contains authentication credentials
	if kubeconfigutil.HasAuthenticationCredentials(config) {
		klog.V(1).Info("[discovery] Using authentication credentials from the discovery file for validating TLS connection")

		// Use the discovery file config for starting the join process
		kubeconfig = config

		// We should ensure that all the authentication info is embedded in config file, so everything will work also when
		// the kubeconfig file will be stored in /etc/kubernetes/boostrap-kubelet.conf
		if err := kubeconfigutil.EnsureAuthenticationInfoAreEmbedded(kubeconfig); err != nil {
			return nil, errors.Wrap(err, "error while reading client cert file or client key file")
		}
	} else {
		// If the discovery file config does not contains authentication credentials
		klog.V(1).Info("[discovery] Discovery file does not contains authentication credentials, using unauthenticated request for validating TLS connection")

		// Create a new kubeconfig object from the discovery file config, with only the server and the CA cert.
		// NB. We do this in order to not pick up other possible misconfigurations in the clusterinfo file
		var fileCluster = kubeconfigutil.GetClusterFromKubeConfig(config)
		kubeconfig = kubeconfigutil.CreateBasic(
			fileCluster.Server,
			clustername,
			"", // no user provided
			fileCluster.CertificateAuthorityData,
		)
	}

	// Try to read the cluster-info config map; this step was required by the original design in order
	// to validate the TLS connection to the server early in the process
	client, err := kubeconfigutil.ToClientSet(kubeconfig)
	if err != nil {
		return nil, err
	}

	currentCluster := kubeconfigutil.GetClusterFromKubeConfig(kubeconfig)
	klog.V(1).Infof("[discovery] Created cluster-info discovery client, requesting info from %q\n", currentCluster.Server)

	var clusterinfoCM *v1.ConfigMap
	wait.PollInfinite(constants.DiscoveryRetryInterval, func() (bool, error) {
		var err error
		clusterinfoCM, err = client.CoreV1().ConfigMaps(metav1.NamespacePublic).Get(bootstrapapi.ConfigMapClusterInfo, metav1.GetOptions{})
		if err != nil {
			if apierrors.IsForbidden(err) {
				// If the request is unauthorized, the cluster admin has not granted access to the cluster info configmap for unauthenticated users
				// In that case, trust the cluster admin and do not refresh the cluster-info data
				klog.Warningf("[discovery] Could not access the %s ConfigMap for refreshing the cluster-info information, but the TLS cert is valid so proceeding...\n", bootstrapapi.ConfigMapClusterInfo)
				return true, nil
			}
			klog.V(1).Infof("[discovery] Error reading the %s ConfigMap, will try again: %v\n", bootstrapapi.ConfigMapClusterInfo, err)
			return false, nil
		}
		return true, nil
	})

	// If we couldn't fetch the cluster-info ConfigMap, just return the cluster-info object the user provided
	if clusterinfoCM == nil {
		return kubeconfig, nil
	}

	// We somehow got hold of the ConfigMap, try to read some data from it. If we can't, fallback on the user-provided file
	refreshedBaseKubeConfig, err := tryParseClusterInfoFromConfigMap(clusterinfoCM)
	if err != nil {
		klog.V(1).Infof("[discovery] The %s ConfigMap isn't set up properly (%v), but the TLS cert is valid so proceeding...\n", bootstrapapi.ConfigMapClusterInfo, err)
		return kubeconfig, nil
	}

	refreshedCluster := kubeconfigutil.GetClusterFromKubeConfig(refreshedBaseKubeConfig)
	currentCluster.Server = refreshedCluster.Server
	currentCluster.CertificateAuthorityData = refreshedCluster.CertificateAuthorityData

	klog.V(1).Infof("[discovery] Synced Server and CertificateAuthorityData from the %s ConfigMap", bootstrapapi.ConfigMapClusterInfo)
	return kubeconfig, nil
}

// tryParseClusterInfoFromConfigMap tries to parse a kubeconfig file from a ConfigMap key
func tryParseClusterInfoFromConfigMap(cm *v1.ConfigMap) (*clientcmdapi.Config, error) {
	kubeConfigString, ok := cm.Data[bootstrapapi.KubeConfigKey]
	if !ok || len(kubeConfigString) == 0 {
		return nil, errors.Errorf("no %s key in ConfigMap", bootstrapapi.KubeConfigKey)
	}
	parsedKubeConfig, err := clientcmd.Load([]byte(kubeConfigString))
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't parse the kubeconfig file in the %s ConfigMap", bootstrapapi.ConfigMapClusterInfo)
	}
	return parsedKubeConfig, nil
}

// validateKubeConfig makes sure the user-provided kubeconfig file is valid
func validateKubeConfig(config *clientcmdapi.Config) error {
	if len(config.Clusters) < 1 {
		return errors.New("the provided cluster-info kubeconfig file must have at least one Cluster defined")
	}
	defaultCluster := kubeconfigutil.GetClusterFromKubeConfig(config)
	if defaultCluster == nil {
		return errors.New("the provided cluster-info kubeconfig file must have an unnamed Cluster or a CurrentContext that specifies a non-nil Cluster")
	}
	return clientcmd.Validate(*config)
}
