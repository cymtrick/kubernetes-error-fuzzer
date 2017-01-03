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
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"k8s.io/kubernetes/pkg/kubectl"
	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

var (
	roleBindingLong = templates.LongDesc(`
		Create a RoleBinding for a particular Role or ClusterRole.`)

	roleBindingExample = templates.Examples(`
		  # Create a RoleBinding for user1, user2, and group1 using the admin ClusterRole
		  kubectl create rolebinding admin --clusterrole=admin --user=user1 --user=user2 --group=group1`)
)

// RoleBinding is a command to ease creating RoleBindings.
func NewCmdCreateRoleBinding(f cmdutil.Factory, cmdOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rolebinding NAME --clusterrole=NAME|--role=NAME [--user=username] [--group=groupname] [--serviceaccount=namespace:serviceaccountname] [--dry-run]",
		Short:   "Create a RoleBinding for a particular Role or ClusterRole",
		Long:    roleBindingLong,
		Example: roleBindingExample,
		Run: func(cmd *cobra.Command, args []string) {
			err := CreateRoleBinding(f, cmdOut, cmd, args)
			cmdutil.CheckErr(err)
		},
	}
	cmdutil.AddApplyAnnotationFlags(cmd)
	cmdutil.AddValidateFlags(cmd)
	cmdutil.AddPrinterFlags(cmd)
	cmdutil.AddGeneratorFlags(cmd, cmdutil.RoleBindingV1GeneratorName)
	cmd.Flags().String("clusterrole", "", "ClusterRole this RoleBinding should reference")
	cmd.Flags().String("role", "", "Role this RoleBinding should reference")
	cmd.Flags().StringSlice("user", []string{}, "usernames to bind to the role")
	cmd.Flags().StringSlice("group", []string{}, "groups to bind to the role")
	cmd.Flags().StringSlice("serviceaccount", []string{}, "service accounts to bind to the role")
	return cmd
}

func CreateRoleBinding(f cmdutil.Factory, cmdOut io.Writer, cmd *cobra.Command, args []string) error {
	name, err := NameFromCommandArgs(cmd, args)
	if err != nil {
		return err
	}
	var generator kubectl.StructuredGenerator
	switch generatorName := cmdutil.GetFlagString(cmd, "generator"); generatorName {
	case cmdutil.RoleBindingV1GeneratorName:
		generator = &kubectl.RoleBindingGeneratorV1{
			Name:            name,
			ClusterRole:     cmdutil.GetFlagString(cmd, "clusterrole"),
			Role:            cmdutil.GetFlagString(cmd, "role"),
			Users:           cmdutil.GetFlagStringSlice(cmd, "user"),
			Groups:          cmdutil.GetFlagStringSlice(cmd, "group"),
			ServiceAccounts: cmdutil.GetFlagStringSlice(cmd, "serviceaccount"),
		}
	default:
		return cmdutil.UsageError(cmd, fmt.Sprintf("Generator: %s not supported.", generatorName))
	}
	return RunCreateSubcommand(f, cmd, cmdOut, &CreateSubcommandOptions{
		Name:                name,
		StructuredGenerator: generator,
		DryRun:              cmdutil.GetDryRunFlag(cmd),
		OutputFormat:        cmdutil.GetFlagString(cmd, "output"),
	})
}
