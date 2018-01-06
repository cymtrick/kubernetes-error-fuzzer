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

package cmd

import (
	"bytes"
	"crypto/x509"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"text/template"
	"time"

	"github.com/renstrom/dedent"
	"github.com/spf13/cobra"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/sets"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcertutil "k8s.io/client-go/util/cert"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	tokenphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/bootstraptoken/node"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	kubeconfigutil "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/pubkeypin"
	tokenutil "k8s.io/kubernetes/cmd/kubeadm/app/util/token"
	api "k8s.io/kubernetes/pkg/apis/core"
	bootstrapapi "k8s.io/kubernetes/pkg/bootstrap/api"
	"k8s.io/kubernetes/pkg/printers"
)

var joinCommandTemplate = template.Must(template.New("join").Parse(`` +
	`kubeadm join --token {{.Token}} {{.MasterHostPort}}{{range $h := .CAPubKeyPins}} --discovery-token-ca-cert-hash {{$h}}{{end}}`,
))

// NewCmdToken returns cobra.Command for token management
func NewCmdToken(out io.Writer, errW io.Writer) *cobra.Command {
	var kubeConfigFile string
	var dryRun bool
	tokenCmd := &cobra.Command{
		Use:   "token",
		Short: "Manage bootstrap tokens.",
		Long: dedent.Dedent(`
			This command manages bootstrap tokens. It is optional and needed only for advanced use cases.

			In short, bootstrap tokens are used for establishing bidirectional trust between a client and a server.
			A bootstrap token can be used when a client (for example a node that is about to join the cluster) needs
			to trust the server it is talking to. Then a bootstrap token with the "signing" usage can be used.
			bootstrap tokens can also function as a way to allow short-lived authentication to the API Server
			(the token serves as a way for the API Server to trust the client), for example for doing the TLS Bootstrap.

			What is a bootstrap token more exactly?
			 - It is a Secret in the kube-system namespace of type "bootstrap.kubernetes.io/token".
			 - A bootstrap token must be of the form "[a-z0-9]{6}.[a-z0-9]{16}". The former part is the public token ID,
			   while the latter is the Token Secret and it must be kept private at all circumstances!
			 - The name of the Secret must be named "bootstrap-token-(token-id)".

			You can read more about bootstrap tokens here:
			  https://kubernetes.io/docs/admin/bootstrap-tokens/
		`),

		// Without this callback, if a user runs just the "token"
		// command without a subcommand, or with an invalid subcommand,
		// cobra will print usage information, but still exit cleanly.
		// We want to return an error code in these cases so that the
		// user knows that their command was invalid.
		RunE: cmdutil.SubCmdRunE("token"),
	}

	tokenCmd.PersistentFlags().StringVar(&kubeConfigFile,
		"kubeconfig", "/etc/kubernetes/admin.conf", "The KubeConfig file to use when talking to the cluster")
	tokenCmd.PersistentFlags().BoolVar(&dryRun,
		"dry-run", dryRun, "Whether to enable dry-run mode or not")

	var usages []string
	var extraGroups []string
	var tokenDuration time.Duration
	var description string
	var printJoinCommand bool
	createCmd := &cobra.Command{
		Use:   "create [token]",
		Short: "Create bootstrap tokens on the server.",
		Long: dedent.Dedent(`
			This command will create a bootstrap token for you.
			You can specify the usages for this token, the "time to live" and an optional human friendly description.

			The [token] is the actual token to write.
			This should be a securely generated random token of the form "[a-z0-9]{6}.[a-z0-9]{16}".
			If no [token] is given, kubeadm will generate a random token instead.
		`),
		Run: func(tokenCmd *cobra.Command, args []string) {
			token := ""
			if len(args) != 0 {
				token = args[0]
			}
			client, err := getClientset(kubeConfigFile, dryRun)
			kubeadmutil.CheckErr(err)

			err = RunCreateToken(out, client, token, tokenDuration, usages, extraGroups, description, printJoinCommand, kubeConfigFile)
			kubeadmutil.CheckErr(err)
		},
	}
	createCmd.Flags().DurationVar(&tokenDuration,
		"ttl", kubeadmconstants.DefaultTokenDuration, "The duration before the token is automatically deleted (e.g. 1s, 2m, 3h). If set to '0', the token will never expire.")
	createCmd.Flags().StringSliceVar(&usages,
		"usages", kubeadmconstants.DefaultTokenUsages, fmt.Sprintf("Describes the ways in which this token can be used. You can pass --usages multiple times or provide a comma separated list of options. Valid options: [%s].", strings.Join(kubeadmconstants.DefaultTokenUsages, ",")))
	createCmd.Flags().StringSliceVar(&extraGroups,
		"groups", []string{kubeadmconstants.NodeBootstrapTokenAuthGroup},
		fmt.Sprintf("Extra groups that this token will authenticate as when used for authentication. Must match %q.", bootstrapapi.BootstrapGroupPattern))
	createCmd.Flags().StringVar(&description,
		"description", "", "A human friendly description of how this token is used.")
	createCmd.Flags().BoolVar(&printJoinCommand,
		"print-join-command", false, "Instead of printing only the token, print the full 'kubeadm join' flag needed to join the cluster using the token.")
	tokenCmd.AddCommand(createCmd)

	tokenCmd.AddCommand(NewCmdTokenGenerate(out))

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List bootstrap tokens on the server.",
		Long: dedent.Dedent(`
			This command will list all bootstrap tokens for you.
		`),
		Run: func(tokenCmd *cobra.Command, args []string) {
			client, err := getClientset(kubeConfigFile, dryRun)
			kubeadmutil.CheckErr(err)

			err = RunListTokens(out, errW, client)
			kubeadmutil.CheckErr(err)
		},
	}
	tokenCmd.AddCommand(listCmd)

	deleteCmd := &cobra.Command{
		Use:   "delete [token-value]",
		Short: "Delete bootstrap tokens on the server.",
		Long: dedent.Dedent(`
			This command will delete a given bootstrap token for you.

			The [token-value] is the full Token of the form "[a-z0-9]{6}.[a-z0-9]{16}" or the
			Token ID of the form "[a-z0-9]{6}" to delete.
		`),
		Run: func(tokenCmd *cobra.Command, args []string) {
			if len(args) < 1 {
				kubeadmutil.CheckErr(fmt.Errorf("missing subcommand; 'token delete' is missing token of form [%q]", tokenutil.TokenIDRegexpString))
			}
			client, err := getClientset(kubeConfigFile, dryRun)
			kubeadmutil.CheckErr(err)

			err = RunDeleteToken(out, client, args[0])
			kubeadmutil.CheckErr(err)
		},
	}
	tokenCmd.AddCommand(deleteCmd)

	return tokenCmd
}

// NewCmdTokenGenerate returns cobra.Command to generate new token
func NewCmdTokenGenerate(out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "Generate and print a bootstrap token, but do not create it on the server.",
		Long: dedent.Dedent(`
			This command will print out a randomly-generated bootstrap token that can be used with
			the "init" and "join" commands.

			You don't have to use this command in order to generate a token. You can do so
			yourself as long as it is in the format "[a-z0-9]{6}.[a-z0-9]{16}". This
			command is provided for convenience to generate tokens in the given format.

			You can also use "kubeadm init" without specifying a token and it will
			generate and print one for you.
		`),
		Run: func(cmd *cobra.Command, args []string) {
			err := RunGenerateToken(out)
			kubeadmutil.CheckErr(err)
		},
	}
}

// RunCreateToken generates a new bootstrap token and stores it as a secret on the server.
func RunCreateToken(out io.Writer, client clientset.Interface, token string, tokenDuration time.Duration, usages []string, extraGroups []string, description string, printJoinCommand bool, kubeConfigFile string) error {
	if len(token) == 0 {
		var err error
		token, err = tokenutil.GenerateToken()
		if err != nil {
			return err
		}
	} else {
		_, _, err := tokenutil.ParseToken(token)
		if err != nil {
			return err
		}
	}

	// adding groups only makes sense for authentication
	usagesSet := sets.NewString(usages...)
	usageAuthentication := strings.TrimPrefix(bootstrapapi.BootstrapTokenUsageAuthentication, bootstrapapi.BootstrapTokenUsagePrefix)
	if len(extraGroups) > 0 && !usagesSet.Has(usageAuthentication) {
		return fmt.Errorf("--groups cannot be specified unless --usages includes %q", usageAuthentication)
	}

	// validate any extra group names
	for _, group := range extraGroups {
		if err := bootstrapapi.ValidateBootstrapGroupName(group); err != nil {
			return err
		}
	}

	// validate usages
	if err := bootstrapapi.ValidateUsages(usages); err != nil {
		return err
	}

	err := tokenphase.CreateNewToken(client, token, tokenDuration, usages, extraGroups, description)
	if err != nil {
		return err
	}

	// if --print-join-command was specified, print the full `kubeadm join` command
	// otherwise, just print the token
	if printJoinCommand {
		joinCommand, err := getJoinCommand(token, kubeConfigFile)
		if err != nil {
			return fmt.Errorf("failed to get join command: %v", err)
		}
		fmt.Fprintln(out, joinCommand)
	} else {
		fmt.Fprintln(out, token)
	}

	return nil
}

// RunGenerateToken just generates a random token for the user
func RunGenerateToken(out io.Writer) error {
	token, err := tokenutil.GenerateToken()
	if err != nil {
		return err
	}

	fmt.Fprintln(out, token)
	return nil
}

// RunListTokens lists details on all existing bootstrap tokens on the server.
func RunListTokens(out io.Writer, errW io.Writer, client clientset.Interface) error {
	// First, build our selector for bootstrap tokens only
	tokenSelector := fields.SelectorFromSet(
		map[string]string{
			api.SecretTypeField: string(bootstrapapi.SecretTypeBootstrapToken),
		},
	)
	listOptions := metav1.ListOptions{
		FieldSelector: tokenSelector.String(),
	}

	secrets, err := client.CoreV1().Secrets(metav1.NamespaceSystem).List(listOptions)
	if err != nil {
		return fmt.Errorf("failed to list bootstrap tokens [%v]", err)
	}

	w := tabwriter.NewWriter(out, 10, 4, 3, ' ', 0)
	fmt.Fprintln(w, "TOKEN\tTTL\tEXPIRES\tUSAGES\tDESCRIPTION\tEXTRA GROUPS")
	for _, secret := range secrets.Items {
		tokenID := getSecretString(&secret, bootstrapapi.BootstrapTokenIDKey)
		if len(tokenID) == 0 {
			fmt.Fprintf(errW, "bootstrap token has no token-id data: %s\n", secret.Name)
			continue
		}

		// enforce the right naming convention
		if secret.Name != fmt.Sprintf("%s%s", bootstrapapi.BootstrapTokenSecretPrefix, tokenID) {
			fmt.Fprintf(errW, "bootstrap token name is not of the form '%s(token-id)': %s\n", bootstrapapi.BootstrapTokenSecretPrefix, secret.Name)
			continue
		}

		tokenSecret := getSecretString(&secret, bootstrapapi.BootstrapTokenSecretKey)
		if len(tokenSecret) == 0 {
			fmt.Fprintf(errW, "bootstrap token has no token-secret data: %s\n", secret.Name)
			continue
		}
		td := &kubeadmapi.TokenDiscovery{ID: tokenID, Secret: tokenSecret}

		// Expiration time is optional, if not specified this implies the token
		// never expires.
		ttl := "<forever>"
		expires := "<never>"
		secretExpiration := getSecretString(&secret, bootstrapapi.BootstrapTokenExpirationKey)
		if len(secretExpiration) > 0 {
			expireTime, err := time.Parse(time.RFC3339, secretExpiration)
			if err != nil {
				fmt.Fprintf(errW, "can't parse expiration time of bootstrap token %s\n", secret.Name)
				continue
			}
			ttl = printers.ShortHumanDuration(expireTime.Sub(time.Now()))
			expires = expireTime.Format(time.RFC3339)
		}

		usages := []string{}
		for k, v := range secret.Data {
			// Skip all fields that don't include this prefix
			if !strings.HasPrefix(k, bootstrapapi.BootstrapTokenUsagePrefix) {
				continue
			}
			// Skip those that don't have this usage set to true
			if string(v) != "true" {
				continue
			}
			usages = append(usages, strings.TrimPrefix(k, bootstrapapi.BootstrapTokenUsagePrefix))
		}
		sort.Strings(usages)
		usageString := strings.Join(usages, ",")
		if len(usageString) == 0 {
			usageString = "<none>"
		}

		description := getSecretString(&secret, bootstrapapi.BootstrapTokenDescriptionKey)
		if len(description) == 0 {
			description = "<none>"
		}

		groups := getSecretString(&secret, bootstrapapi.BootstrapTokenExtraGroupsKey)
		if len(groups) == 0 {
			groups = "<none>"
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", tokenutil.BearerToken(td), ttl, expires, usageString, description, groups)
	}
	w.Flush()
	return nil
}

// RunDeleteToken removes a bootstrap token from the server.
func RunDeleteToken(out io.Writer, client clientset.Interface, tokenIDOrToken string) error {
	// Assume the given first argument is a token id and try to parse it
	tokenID := tokenIDOrToken
	if err := tokenutil.ParseTokenID(tokenIDOrToken); err != nil {
		if tokenID, _, err = tokenutil.ParseToken(tokenIDOrToken); err != nil {
			return fmt.Errorf("given token or token id %q didn't match pattern [%q] or [%q]", tokenIDOrToken, tokenutil.TokenIDRegexpString, tokenutil.TokenRegexpString)
		}
	}

	tokenSecretName := fmt.Sprintf("%s%s", bootstrapapi.BootstrapTokenSecretPrefix, tokenID)
	if err := client.CoreV1().Secrets(metav1.NamespaceSystem).Delete(tokenSecretName, nil); err != nil {
		return fmt.Errorf("failed to delete bootstrap token [%v]", err)
	}
	fmt.Fprintf(out, "bootstrap token with id %q deleted\n", tokenID)
	return nil
}

func getSecretString(secret *v1.Secret, key string) string {
	if secret.Data == nil {
		return ""
	}
	if val, ok := secret.Data[key]; ok {
		return string(val)
	}
	return ""
}

func getClientset(file string, dryRun bool) (clientset.Interface, error) {
	if dryRun {
		dryRunGetter, err := apiclient.NewClientBackedDryRunGetterFromKubeconfig(file)
		if err != nil {
			return nil, err
		}
		return apiclient.NewDryRunClient(dryRunGetter, os.Stdout), nil
	}
	return kubeconfigutil.ClientSetFromFile(file)
}

func getJoinCommand(token string, kubeConfigFile string) (string, error) {
	// load the kubeconfig file to get the CA certificate and endpoint
	config, err := clientcmd.LoadFromFile(kubeConfigFile)
	if err != nil {
		return "", fmt.Errorf("failed to load kubeconfig: %v", err)
	}

	// load the default cluster config
	clusterConfig := kubeconfigutil.GetClusterFromKubeConfig(config)
	if clusterConfig == nil {
		return "", fmt.Errorf("failed to get default cluster config")
	}

	// load CA certificates from the kubeconfig (either from PEM data or by file path)
	var caCerts []*x509.Certificate
	if clusterConfig.CertificateAuthorityData != nil {
		caCerts, err = clientcertutil.ParseCertsPEM(clusterConfig.CertificateAuthorityData)
		if err != nil {
			return "", fmt.Errorf("failed to parse CA certificate from kubeconfig: %v", err)
		}
	} else if clusterConfig.CertificateAuthority != "" {
		caCerts, err = clientcertutil.CertsFromFile(clusterConfig.CertificateAuthority)
		if err != nil {
			return "", fmt.Errorf("failed to load CA certificate referenced by kubeconfig: %v", err)
		}
	} else {
		return "", fmt.Errorf("no CA certificates found in kubeconfig")
	}

	// hash all the CA certs and include their public key pins as trusted values
	publicKeyPins := make([]string, 0, len(caCerts))
	for _, caCert := range caCerts {
		publicKeyPins = append(publicKeyPins, pubkeypin.Hash(caCert))
	}

	ctx := map[string]interface{}{
		"Token":          token,
		"CAPubKeyPins":   publicKeyPins,
		"MasterHostPort": strings.Replace(clusterConfig.Server, "https://", "", -1),
	}

	var out bytes.Buffer
	err = joinCommandTemplate.Execute(&out, ctx)
	if err != nil {
		return "", fmt.Errorf("failed to render join command template: %v", err)
	}
	return out.String(), nil
}
