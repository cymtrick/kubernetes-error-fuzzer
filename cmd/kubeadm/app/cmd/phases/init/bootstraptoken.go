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

package phases

import (
	"fmt"

	"github.com/pkg/errors"

	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
	clusterinfophase "k8s.io/kubernetes/cmd/kubeadm/app/phases/bootstraptoken/clusterinfo"
	nodebootstraptokenphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/bootstraptoken/node"
	"k8s.io/kubernetes/pkg/util/normalizer"
)

var (
	bootstrapTokenLongDesc = normalizer.LongDesc(`
		Bootstrap tokens are used for establishing bidirectional trust between a node joining
		the cluster and a the control-plane node.

		This command makes all the configurations required to make bootstrap tokens works
		and then creates an initial token.
		`)

	bootstrapTokenExamples = normalizer.Examples(`
		# Makes all the bootstrap token configurations and creates an initial token, functionally
		# equivalent to what generated by kubeadm init.
		kubeadm init phase bootstrap-token
		`)
)

// NewBootstrapTokenPhase returns the phase to bootstrapToken
func NewBootstrapTokenPhase() workflow.Phase {
	return workflow.Phase{
		Name:    "bootstrap-token",
		Aliases: []string{"bootstraptoken"},
		Short:   "Generates bootstrap tokens used to join a node to a cluster",
		Example: bootstrapTokenExamples,
		Long:    bootstrapTokenLongDesc,
		InheritFlags: []string{
			options.CfgPath,
			options.KubeconfigPath,
			options.SkipTokenPrint,
		},
		Run: runBootstrapToken,
	}
}

func runBootstrapToken(c workflow.RunData) error {
	data, ok := c.(InitData)
	if !ok {
		return errors.New("bootstrap-token phase invoked with an invalid data struct")
	}

	client, err := data.Client()
	if err != nil {
		return err
	}

	if !data.SkipTokenPrint() {
		tokens := data.Tokens()
		if len(tokens) == 1 {
			fmt.Printf("[bootstrap-token] Using token: %s\n", tokens[0])
		} else if len(tokens) > 1 {
			fmt.Printf("[bootstrap-token] Using tokens: %v\n", tokens)
		}
	}

	fmt.Println("[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles")
	// Create the default node bootstrap token
	if err := nodebootstraptokenphase.UpdateOrCreateTokens(client, false, data.Cfg().BootstrapTokens); err != nil {
		return errors.Wrap(err, "error updating or creating token")
	}
	// Create RBAC rules that makes the bootstrap tokens able to post CSRs
	if err := nodebootstraptokenphase.AllowBootstrapTokensToPostCSRs(client); err != nil {
		return errors.Wrap(err, "error allowing bootstrap tokens to post CSRs")
	}
	// Create RBAC rules that makes the bootstrap tokens able to get their CSRs approved automatically
	if err := nodebootstraptokenphase.AutoApproveNodeBootstrapTokens(client); err != nil {
		return errors.Wrap(err, "error auto-approving node bootstrap tokens")
	}

	// Create/update RBAC rules that makes the nodes to rotate certificates and get their CSRs approved automatically
	if err := nodebootstraptokenphase.AutoApproveNodeCertificateRotation(client); err != nil {
		return err
	}

	// Create the cluster-info ConfigMap with the associated RBAC rules
	if err := clusterinfophase.CreateBootstrapConfigMapIfNotExists(client, data.KubeConfigPath()); err != nil {
		return errors.Wrap(err, "error creating bootstrap ConfigMap")
	}
	if err := clusterinfophase.CreateClusterInfoRBACRules(client); err != nil {
		return errors.Wrap(err, "error creating clusterinfo RBAC rules")
	}
	return nil
}
