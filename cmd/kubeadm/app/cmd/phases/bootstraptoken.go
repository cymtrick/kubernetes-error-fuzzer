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
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	clientset "k8s.io/client-go/kubernetes"
	bootstrapapi "k8s.io/client-go/tools/bootstrap/token/api"
	bootstraputil "k8s.io/client-go/tools/bootstrap/token/util"
	kubeadmapiext "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/validation"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/bootstraptoken/clusterinfo"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/bootstraptoken/node"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	kubeconfigutil "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/util/normalizer"
)

var (
	allTokenLongDesc = normalizer.LongDesc(`
		Bootstrap tokens are used for establishing bidirectional trust between a node joining 
		the cluster and a the master node.

		This command makes all the configurations required to make bootstrap tokens works
		and then creates an initial token.
		` + cmdutil.AlphaDisclaimer)

	allTokenExamples = normalizer.Examples(`
		# Makes all the bootstrap token configurations and creates an initial token, functionally 
		# equivalent to what generated by kubeadm init.
		kubeadm alpha phase bootstrap-token all
		`)

	createTokenLongDesc = normalizer.LongDesc(`
		Creates a bootstrap token. If no token value is given, kubeadm will generate a random token instead.
		 
		Alternatively, you can use kubeadm token.
		` + cmdutil.AlphaDisclaimer)

	clusterInfoLongDesc = fmt.Sprintf(normalizer.LongDesc(`
		Uploads the %q ConfigMap in the %q namespace, populating it with cluster information extracted from the 
		given kubeconfig file. The ConfigMap is used for the node bootstrap process in its initial phases, 
		before the client trusts the API server. 

		See online documentation about Authenticating with Bootstrap Tokens for more details.
		`+cmdutil.AlphaDisclaimer), bootstrapapi.ConfigMapClusterInfo, metav1.NamespacePublic)

	nodePostCSRsLongDesc = normalizer.LongDesc(`
		Configures RBAC rules to allow node bootstrap tokens to post a certificate signing request,
		 thus enabling nodes joining the cluster to request long term certificate credentials.

		See online documentation about TLS bootstrapping for more details.
		` + cmdutil.AlphaDisclaimer)

	nodeAutoApproveLongDesc = normalizer.LongDesc(`
		Configures RBAC rules to allow the csrapprover controller to automatically approve 
		certificate signing requests generated by nodes joining the cluster.
		It configures also RBAC rules for certificates rotation (with auto approval of new certificates).
		
		See online documentation about TLS bootstrapping for more details.
		` + cmdutil.AlphaDisclaimer)
)

// NewCmdBootstrapToken returns the Cobra command for running the mark-master phase
func NewCmdBootstrapToken() *cobra.Command {
	var kubeConfigFile string
	cmd := &cobra.Command{
		Use:     "bootstrap-token",
		Short:   "Manage kubeadm-specific bootstrap token functions",
		Long:    cmdutil.MacroCommandLongDescription,
		Aliases: []string{"bootstraptoken"},
	}

	cmd.PersistentFlags().StringVar(&kubeConfigFile, "kubeconfig", "/etc/kubernetes/admin.conf", "The KubeConfig file to use when talking to the cluster")

	// Add subcommands
	cmd.AddCommand(NewSubCmdBootstrapTokenAll(&kubeConfigFile))
	cmd.AddCommand(NewSubCmdBootstrapToken(&kubeConfigFile))
	cmd.AddCommand(NewSubCmdClusterInfo(&kubeConfigFile))
	cmd.AddCommand(NewSubCmdNodeBootstrapToken(&kubeConfigFile))

	return cmd
}

// NewSubCmdBootstrapTokenAll returns the Cobra command for running the token all sub-phase
func NewSubCmdBootstrapTokenAll(kubeConfigFile *string) *cobra.Command {
	cfg := &kubeadmapiext.MasterConfiguration{
		// KubernetesVersion is not used by bootstrap-token, but we set this explicitly to avoid
		// the lookup of the version from the internet when executing ConfigFileAndDefaultsToInternalConfig
		KubernetesVersion: "v1.9.0",
	}

	// Default values for the cobra help text
	legacyscheme.Scheme.Default(cfg)

	var cfgPath, description string
	var usages, extraGroups []string
	var skipTokenPrint bool

	cmd := &cobra.Command{
		Use:     "all",
		Short:   "Makes all the bootstrap token configurations and creates an initial token",
		Long:    allTokenLongDesc,
		Example: allTokenExamples,
		Run: func(cmd *cobra.Command, args []string) {
			err := validation.ValidateMixedArguments(cmd.Flags())
			kubeadmutil.CheckErr(err)

			client, err := kubeconfigutil.ClientSetFromFile(*kubeConfigFile)
			kubeadmutil.CheckErr(err)

			// Creates the bootstap token
			err = createBootstrapToken(*kubeConfigFile, client, cfgPath, cfg, description, usages, extraGroups, skipTokenPrint)
			kubeadmutil.CheckErr(err)

			// Create the cluster-info ConfigMap or update if it already exists
			err = clusterinfo.CreateBootstrapConfigMapIfNotExists(client, *kubeConfigFile)
			kubeadmutil.CheckErr(err)

			// Create the RBAC rules that expose the cluster-info ConfigMap properly
			err = clusterinfo.CreateClusterInfoRBACRules(client)
			kubeadmutil.CheckErr(err)

			// Create RBAC rules that makes the bootstrap tokens able to post CSRs
			err = node.AllowBootstrapTokensToPostCSRs(client)
			kubeadmutil.CheckErr(err)

			// Create RBAC rules that makes the bootstrap tokens able to get their CSRs approved automatically
			err = node.AutoApproveNodeBootstrapTokens(client)
			kubeadmutil.CheckErr(err)

			// Create/update RBAC rules that makes the nodes to rotate certificates and get their CSRs approved automatically
			err = node.AutoApproveNodeCertificateRotation(client)
			kubeadmutil.CheckErr(err)
		},
	}

	// Adds flags to the command
	addBootstrapTokenFlags(cmd.Flags(), cfg, &cfgPath, &description, &usages, &extraGroups, &skipTokenPrint)

	return cmd
}

// NewSubCmdBootstrapToken returns the Cobra command for running the create token phase
func NewSubCmdBootstrapToken(kubeConfigFile *string) *cobra.Command {
	cfg := &kubeadmapiext.MasterConfiguration{
		// KubernetesVersion is not used by bootstrap-token, but we set this explicitly to avoid
		// the lookup of the version from the internet when executing ConfigFileAndDefaultsToInternalConfig
		KubernetesVersion: "v1.9.0",
	}

	// Default values for the cobra help text
	legacyscheme.Scheme.Default(cfg)

	var cfgPath, description string
	var usages, extraGroups []string
	var skipTokenPrint bool

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a bootstrap token to be used for node joining",
		Long:  createTokenLongDesc,
		Run: func(cmd *cobra.Command, args []string) {
			err := validation.ValidateMixedArguments(cmd.Flags())
			kubeadmutil.CheckErr(err)

			client, err := kubeconfigutil.ClientSetFromFile(*kubeConfigFile)
			kubeadmutil.CheckErr(err)

			err = createBootstrapToken(*kubeConfigFile, client, cfgPath, cfg, description, usages, extraGroups, skipTokenPrint)
			kubeadmutil.CheckErr(err)
		},
	}

	// Adds flags to the command
	addBootstrapTokenFlags(cmd.Flags(), cfg, &cfgPath, &description, &usages, &extraGroups, &skipTokenPrint)

	return cmd
}

// NewSubCmdClusterInfo returns the Cobra command for running the cluster-info sub-phase
func NewSubCmdClusterInfo(kubeConfigFile *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cluster-info",
		Short:   "Uploads the cluster-info ConfigMap from the given kubeconfig file",
		Long:    clusterInfoLongDesc,
		Aliases: []string{"clusterinfo"},
		Run: func(cmd *cobra.Command, args []string) {
			client, err := kubeconfigutil.ClientSetFromFile(*kubeConfigFile)
			kubeadmutil.CheckErr(err)

			// Create the cluster-info ConfigMap or update if it already exists
			err = clusterinfo.CreateBootstrapConfigMapIfNotExists(client, *kubeConfigFile)
			kubeadmutil.CheckErr(err)

			// Create the RBAC rules that expose the cluster-info ConfigMap properly
			err = clusterinfo.CreateClusterInfoRBACRules(client)
			kubeadmutil.CheckErr(err)
		},
	}
	return cmd
}

// NewSubCmdNodeBootstrapToken returns the Cobra command for running the node sub-phase
func NewSubCmdNodeBootstrapToken(kubeConfigFile *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "node",
		Short:   "Configures the node bootstrap process",
		Aliases: []string{"clusterinfo"},
		Long:    cmdutil.MacroCommandLongDescription,
	}

	cmd.AddCommand(NewSubCmdNodeBootstrapTokenPostCSRs(kubeConfigFile))
	cmd.AddCommand(NewSubCmdNodeBootstrapTokenAutoApprove(kubeConfigFile))

	return cmd
}

// NewSubCmdNodeBootstrapTokenPostCSRs returns the Cobra command for running the allow-post-csrs sub-phase
func NewSubCmdNodeBootstrapTokenPostCSRs(kubeConfigFile *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allow-post-csrs",
		Short: "Configures RBAC to allow node bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials",
		Long:  nodePostCSRsLongDesc,
		Run: func(cmd *cobra.Command, args []string) {
			client, err := kubeconfigutil.ClientSetFromFile(*kubeConfigFile)
			kubeadmutil.CheckErr(err)

			// Create RBAC rules that makes the bootstrap tokens able to post CSRs
			err = node.AllowBootstrapTokensToPostCSRs(client)
			kubeadmutil.CheckErr(err)
		},
	}
	return cmd
}

// NewSubCmdNodeBootstrapTokenAutoApprove returns the Cobra command for running the allow-auto-approve sub-phase
func NewSubCmdNodeBootstrapTokenAutoApprove(kubeConfigFile *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allow-auto-approve",
		Short: "Configures RBAC rules to allow the csrapprover controller automatically approve CSRs from a node bootstrap token",
		Long:  nodeAutoApproveLongDesc,
		Run: func(cmd *cobra.Command, args []string) {
			client, err := kubeconfigutil.ClientSetFromFile(*kubeConfigFile)
			kubeadmutil.CheckErr(err)

			// Create RBAC rules that makes the bootstrap tokens able to get their CSRs approved automatically
			err = node.AutoApproveNodeBootstrapTokens(client)
			kubeadmutil.CheckErr(err)

			// Create/update RBAC rules that makes the nodes to rotate certificates and get their CSRs approved automatically
			err = node.AutoApproveNodeCertificateRotation(client)
			kubeadmutil.CheckErr(err)
		},
	}
	return cmd
}

func addBootstrapTokenFlags(flagSet *pflag.FlagSet, cfg *kubeadmapiext.MasterConfiguration, cfgPath, description *string, usages, extraGroups *[]string, skipTokenPrint *bool) {
	flagSet.StringVar(
		cfgPath, "config", *cfgPath,
		"Path to kubeadm config file (WARNING: Usage of a configuration file is experimental)",
	)
	flagSet.StringVar(
		&cfg.Token, "token", cfg.Token,
		"The token to use for establishing bidirectional trust between nodes and masters",
	)
	flagSet.DurationVar(
		&cfg.TokenTTL.Duration, "ttl", kubeadmconstants.DefaultTokenDuration,
		"The duration before the token is automatically deleted (e.g. 1s, 2m, 3h). If set to '0', the token will never expire",
	)
	flagSet.StringSliceVar(
		usages, "usages", kubeadmconstants.DefaultTokenUsages,
		fmt.Sprintf("Describes the ways in which this token can be used. You can pass --usages multiple times or provide a comma separated list of options. Valid options: [%s]", strings.Join(kubeadmconstants.DefaultTokenUsages, ",")),
	)
	flagSet.StringSliceVar(
		extraGroups, "groups", []string{kubeadmconstants.NodeBootstrapTokenAuthGroup},
		fmt.Sprintf("Extra groups that this token will authenticate as when used for authentication. Must match %q", bootstrapapi.BootstrapGroupPattern),
	)
	flagSet.StringVar(
		description, "description", "The default bootstrap token generated by 'kubeadm init'.",
		"A human friendly description of how this token is used.",
	)
	flagSet.BoolVar(
		skipTokenPrint, "skip-token-print", *skipTokenPrint,
		"Skip printing of the bootstrap token",
	)
}

func createBootstrapToken(kubeConfigFile string, client clientset.Interface, cfgPath string, cfg *kubeadmapiext.MasterConfiguration, description string, usages, extraGroups []string, skipTokenPrint bool) error {
	// adding groups only makes sense for authentication
	usagesSet := sets.NewString(usages...)
	usageAuthentication := strings.TrimPrefix(bootstrapapi.BootstrapTokenUsageAuthentication, bootstrapapi.BootstrapTokenUsagePrefix)
	if len(extraGroups) > 0 && !usagesSet.Has(usageAuthentication) {
		return fmt.Errorf("--groups cannot be specified unless --usages includes %q", usageAuthentication)
	}

	// validate any extra group names
	for _, group := range extraGroups {
		if err := bootstraputil.ValidateBootstrapGroupName(group); err != nil {
			return err
		}
	}

	// This call returns the ready-to-use configuration based on the configuration file that might or might not exist and the default cfg populated by flags
	internalcfg, err := configutil.ConfigFileAndDefaultsToInternalConfig(cfgPath, cfg)
	kubeadmutil.CheckErr(err)

	// Creates or updates the token
	if err := node.UpdateOrCreateToken(client, internalcfg.Token, false, internalcfg.TokenTTL.Duration, usages, extraGroups, description); err != nil {
		return err
	}

	fmt.Println("[bootstraptoken] Bootstrap token Created")
	fmt.Println("[bootstraptoken] You can now join any number of machines by running:")

	joinCommand, err := cmdutil.GetJoinCommand(kubeConfigFile, internalcfg.Token, skipTokenPrint)
	if err != nil {
		return fmt.Errorf("failed to get join command: %v", err)
	}
	fmt.Println(joinCommand)

	return nil
}
