/*
Copyright 2017 The Kubernetes Authors.

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

package phases

import (
	"fmt"

	"github.com/spf13/cobra"

	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmscheme "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	kubeadmapiv1alpha2 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha2"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/validation"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	certsphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/certs"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	"k8s.io/kubernetes/pkg/util/normalizer"
)

var (
	allCertsLongDesc = normalizer.LongDesc(`
		Generates a self-signed CA to provision identities for each component in the cluster (including nodes)
		and client certificates to be used by various components.

		If a given certificate and private key pair both exist, kubeadm skips the generation step and  
		existing files will be used.
		` + cmdutil.AlphaDisclaimer)

	allCertsExample = normalizer.Examples(`
		# Creates all PKI assets necessary to establish the control plane, 
		# functionally equivalent to what generated by kubeadm init.
		kubeadm alpha phase certs all

		# Creates all PKI assets using options read from a configuration file.
		kubeadm alpha phase certs all --config masterconfiguration.yaml
		`)

	caCertLongDesc = fmt.Sprintf(normalizer.LongDesc(`
		Generates the self-signed kubernetes certificate authority and related key, and saves them into %s and %s files. 

		If both files already exist, kubeadm skips the generation step and existing files will be used.
		`+cmdutil.AlphaDisclaimer), kubeadmconstants.CACertName, kubeadmconstants.CAKeyName)

	apiServerCertLongDesc = fmt.Sprintf(normalizer.LongDesc(`
		Generates the API server serving certificate and key and saves them into %s and %s files.
		
		The certificate includes default subject alternative names and additional SANs provided by the user;
		default SANs are: <node-name>, <apiserver-advertise-address>, kubernetes, kubernetes.default, kubernetes.default.svc,
		kubernetes.default.svc.<service-dns-domain>, <internalAPIServerVirtualIP> (that is the .10 address in <service-cidr> address space).

		If both files already exist, kubeadm skips the generation step and existing files will be used.
		`+cmdutil.AlphaDisclaimer), kubeadmconstants.APIServerCertName, kubeadmconstants.APIServerKeyName)

	apiServerKubeletCertLongDesc = fmt.Sprintf(normalizer.LongDesc(`
		Generates the client certificate for the API server to connect to the kubelet securely and the respective key,
		and saves them into %s and %s files.

		If both files already exist, kubeadm skips the generation step and existing files will be used.
		`+cmdutil.AlphaDisclaimer), kubeadmconstants.APIServerKubeletClientCertName, kubeadmconstants.APIServerKubeletClientKeyName)

	etcdCaCertLongDesc = fmt.Sprintf(normalizer.LongDesc(`
		Generates the self-signed etcd certificate authority and related key and saves them into %s and %s files. 

		If both files already exist, kubeadm skips the generation step and existing files will be used.
		`+cmdutil.AlphaDisclaimer), kubeadmconstants.EtcdCACertName, kubeadmconstants.EtcdCAKeyName)

	etcdServerCertLongDesc = fmt.Sprintf(normalizer.LongDesc(`
		Generates the etcd serving certificate and key and saves them into %s and %s files.
		
		The certificate includes default subject alternative names and additional SANs provided by the user;
		default SANs are: localhost, 127.0.0.1.

		If both files already exist, kubeadm skips the generation step and existing files will be used.
		`+cmdutil.AlphaDisclaimer), kubeadmconstants.EtcdServerCertName, kubeadmconstants.EtcdServerKeyName)

	etcdPeerCertLongDesc = fmt.Sprintf(normalizer.LongDesc(`
		Generates the etcd peer certificate and key and saves them into %s and %s files.
		
		The certificate includes default subject alternative names and additional SANs provided by the user;
		default SANs are: <node-name>, <apiserver-advertise-address>.

		If both files already exist, kubeadm skips the generation step and existing files will be used.
		`+cmdutil.AlphaDisclaimer), kubeadmconstants.EtcdPeerCertName, kubeadmconstants.EtcdPeerKeyName)

	etcdHealthcheckClientCertLongDesc = fmt.Sprintf(normalizer.LongDesc(`
		Generates the client certificate for liveness probes to healthcheck etcd and the respective key,
		and saves them into %s and %s files.

		If both files already exist, kubeadm skips the generation step and existing files will be used.
		`+cmdutil.AlphaDisclaimer), kubeadmconstants.EtcdHealthcheckClientCertName, kubeadmconstants.EtcdHealthcheckClientKeyName)

	apiServerEtcdServerCertLongDesc = fmt.Sprintf(normalizer.LongDesc(`
		Generates the client certificate for the API server to connect to etcd securely and the respective key,
		and saves them into %s and %s files.

		If both files already exist, kubeadm skips the generation step and existing files will be used.
		`+cmdutil.AlphaDisclaimer), kubeadmconstants.APIServerEtcdClientCertName, kubeadmconstants.APIServerEtcdClientKeyName)

	saKeyLongDesc = fmt.Sprintf(normalizer.LongDesc(`
		Generates the private key for signing service account tokens along with its public key, and saves them into
		%s and %s files.

		If both files already exist, kubeadm skips the generation step and existing files will be used.
		`+cmdutil.AlphaDisclaimer), kubeadmconstants.ServiceAccountPrivateKeyName, kubeadmconstants.ServiceAccountPublicKeyName)

	frontProxyCaCertLongDesc = fmt.Sprintf(normalizer.LongDesc(`
		Generates the front proxy CA certificate and key and saves them into %s and %s files.

		If both files already exist, kubeadm skips the generation step and existing files will be used.
		`+cmdutil.AlphaDisclaimer), kubeadmconstants.FrontProxyCACertName, kubeadmconstants.FrontProxyCAKeyName)

	frontProxyClientCertLongDesc = fmt.Sprintf(normalizer.LongDesc(`
		Generates the front proxy client certificate and key and saves them into %s and %s files.

		If both files already exist, kubeadm skips the generation step and existing files will be used.
		`+cmdutil.AlphaDisclaimer), kubeadmconstants.FrontProxyClientCertName, kubeadmconstants.FrontProxyClientKeyName)
)

// NewCmdCerts returns main command for certs phase
func NewCmdCerts() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "certs",
		Aliases: []string{"certificates"},
		Short:   "Generates certificates for a Kubernetes cluster",
		Long:    cmdutil.MacroCommandLongDescription,
	}

	cmd.AddCommand(getCertsSubCommands("")...)
	return cmd
}

// getCertsSubCommands returns sub commands for certs phase
func getCertsSubCommands(defaultKubernetesVersion string) []*cobra.Command {

	cfg := &kubeadmapiv1alpha2.MasterConfiguration{}

	// This is used for unit testing only...
	// If we wouldn't set this to something, the code would dynamically look up the version from the internet
	// By setting this explicitly for tests workarounds that
	if defaultKubernetesVersion != "" {
		cfg.KubernetesVersion = defaultKubernetesVersion
	}

	// Default values for the cobra help text
	kubeadmscheme.Scheme.Default(cfg)

	var cfgPath string
	var subCmds []*cobra.Command

	subCmdProperties := []struct {
		use      string
		short    string
		long     string
		examples string
		cmdFunc  func(cfg *kubeadmapi.MasterConfiguration) error
	}{
		{
			use:      "all",
			short:    "Generates all PKI assets necessary to establish the control plane",
			long:     allCertsLongDesc,
			examples: allCertsExample,
			cmdFunc:  certsphase.CreatePKIAssets,
		},
		{
			use:     "ca",
			short:   "Generates a self-signed kubernetes CA to provision identities for components of the cluster",
			long:    caCertLongDesc,
			cmdFunc: certsphase.CreateCACertAndKeyFiles,
		},
		{
			use:     "apiserver",
			short:   "Generates an API server serving certificate and key",
			long:    apiServerCertLongDesc,
			cmdFunc: certsphase.CreateAPIServerCertAndKeyFiles,
		},
		{
			use:     "apiserver-kubelet-client",
			short:   "Generates a client certificate for the API server to connect to the kubelets securely",
			long:    apiServerKubeletCertLongDesc,
			cmdFunc: certsphase.CreateAPIServerKubeletClientCertAndKeyFiles,
		},
		{
			use:     "etcd-ca",
			short:   "Generates a self-signed CA to provision identities for etcd",
			long:    etcdCaCertLongDesc,
			cmdFunc: certsphase.CreateEtcdCACertAndKeyFiles,
		},
		{
			use:     "etcd-server",
			short:   "Generates an etcd serving certificate and key",
			long:    etcdServerCertLongDesc,
			cmdFunc: certsphase.CreateEtcdServerCertAndKeyFiles,
		},
		{
			use:     "etcd-peer",
			short:   "Generates an etcd peer certificate and key",
			long:    etcdPeerCertLongDesc,
			cmdFunc: certsphase.CreateEtcdPeerCertAndKeyFiles,
		},
		{
			use:     "etcd-healthcheck-client",
			short:   "Generates a client certificate for liveness probes to healthcheck etcd",
			long:    etcdHealthcheckClientCertLongDesc,
			cmdFunc: certsphase.CreateEtcdHealthcheckClientCertAndKeyFiles,
		},
		{
			use:     "apiserver-etcd-client",
			short:   "Generates a client certificate for the API server to connect to etcd securely",
			long:    apiServerEtcdServerCertLongDesc,
			cmdFunc: certsphase.CreateAPIServerEtcdClientCertAndKeyFiles,
		},
		{
			use:     "sa",
			short:   "Generates a private key for signing service account tokens along with its public key",
			long:    saKeyLongDesc,
			cmdFunc: certsphase.CreateServiceAccountKeyAndPublicKeyFiles,
		},
		{
			use:     "front-proxy-ca",
			short:   "Generates a front proxy CA certificate and key for a Kubernetes cluster",
			long:    frontProxyCaCertLongDesc,
			cmdFunc: certsphase.CreateFrontProxyCACertAndKeyFiles,
		},
		{
			use:     "front-proxy-client",
			short:   "Generates a front proxy CA client certificate and key for a Kubernetes cluster",
			long:    frontProxyClientCertLongDesc,
			cmdFunc: certsphase.CreateFrontProxyClientCertAndKeyFiles,
		},
	}

	for _, properties := range subCmdProperties {
		// Creates the UX Command
		cmd := &cobra.Command{
			Use:     properties.use,
			Short:   properties.short,
			Long:    properties.long,
			Example: properties.examples,
			Run:     runCmdFunc(properties.cmdFunc, &cfgPath, cfg),
		}

		// Add flags to the command
		cmd.Flags().StringVar(&cfgPath, "config", cfgPath, "Path to kubeadm config file. WARNING: Usage of a configuration file is experimental")
		cmd.Flags().StringVar(&cfg.CertificatesDir, "cert-dir", cfg.CertificatesDir, "The path where to save the certificates")
		if properties.use == "all" || properties.use == "apiserver" {
			cmd.Flags().StringVar(&cfg.Networking.DNSDomain, "service-dns-domain", cfg.Networking.DNSDomain, "Alternative domain for services, to use for the API server serving cert")
			cmd.Flags().StringVar(&cfg.Networking.ServiceSubnet, "service-cidr", cfg.Networking.ServiceSubnet, "Alternative range of IP address for service VIPs, from which derives the internal API server VIP that will be added to the API Server serving cert")
			cmd.Flags().StringSliceVar(&cfg.APIServerCertSANs, "apiserver-cert-extra-sans", []string{}, "Optional extra altnames to use for the API server serving cert. Can be both IP addresses and DNS names")
			cmd.Flags().StringVar(&cfg.API.AdvertiseAddress, "apiserver-advertise-address", cfg.API.AdvertiseAddress, "The IP address the API server is accessible on, to use for the API server serving cert")
		}

		subCmds = append(subCmds, cmd)
	}

	return subCmds
}

// runCmdFunc creates a cobra.Command Run function, by composing the call to the given cmdFunc with necessary additional steps (e.g preparation of input parameters)
func runCmdFunc(cmdFunc func(cfg *kubeadmapi.MasterConfiguration) error, cfgPath *string, cfg *kubeadmapiv1alpha2.MasterConfiguration) func(cmd *cobra.Command, args []string) {

	// the following statement build a closure that wraps a call to a cmdFunc, binding
	// the function itself with the specific parameters of each sub command.
	// Please note that specific parameter should be passed as value, while other parameters - passed as reference -
	// are shared between sub commands and gets access to current value e.g. flags value.

	return func(cmd *cobra.Command, args []string) {
		if err := validation.ValidateMixedArguments(cmd.Flags()); err != nil {
			kubeadmutil.CheckErr(err)
		}

		// This call returns the ready-to-use configuration based on the configuration file that might or might not exist and the default cfg populated by flags
		internalcfg, err := configutil.ConfigFileAndDefaultsToInternalConfig(*cfgPath, cfg)
		kubeadmutil.CheckErr(err)

		// Execute the cmdFunc
		err = cmdFunc(internalcfg)
		kubeadmutil.CheckErr(err)
	}
}
