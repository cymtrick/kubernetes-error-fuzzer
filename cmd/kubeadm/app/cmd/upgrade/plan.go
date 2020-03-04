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

package upgrade

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/version"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	outputapiv1alpha1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/output/v1alpha1"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/upgrade"
	etcdutil "k8s.io/kubernetes/cmd/kubeadm/app/util/etcd"
)

type planFlags struct {
	*applyPlanFlags
}

// NewCmdPlan returns the cobra command for `kubeadm upgrade plan`
func NewCmdPlan(apf *applyPlanFlags) *cobra.Command {
	flags := &planFlags{
		applyPlanFlags: apf,
	}

	cmd := &cobra.Command{
		Use:   "plan [version] [flags]",
		Short: "Check which versions are available to upgrade to and validate whether your current cluster is upgradeable. To skip the internet check, pass in the optional [version] parameter",
		RunE: func(_ *cobra.Command, args []string) error {
			userVersion, err := getK8sVersionFromUserInput(flags.applyPlanFlags, args, false)
			if err != nil {
				return err
			}

			return runPlan(flags, userVersion)
		},
	}

	// Register the common flags for apply and plan
	addApplyPlanFlags(cmd.Flags(), flags.applyPlanFlags)
	return cmd
}

// runPlan takes care of outputting available versions to upgrade to for the user
func runPlan(flags *planFlags, userVersion string) error {
	// Start with the basics, verify that the cluster is healthy, build a client and a versionGetter. Never dry-run when planning.
	klog.V(1).Infoln("[upgrade/plan] verifying health of cluster")
	klog.V(1).Infoln("[upgrade/plan] retrieving configuration from cluster")
	client, versionGetter, cfg, err := enforceRequirements(flags.applyPlanFlags, false, userVersion)
	if err != nil {
		return err
	}

	var etcdClient etcdutil.ClusterInterrogator

	// Currently this is the only method we have for distinguishing
	// external etcd vs static pod etcd
	isExternalEtcd := cfg.Etcd.External != nil
	if isExternalEtcd {
		etcdClient, err = etcdutil.New(
			cfg.Etcd.External.Endpoints,
			cfg.Etcd.External.CAFile,
			cfg.Etcd.External.CertFile,
			cfg.Etcd.External.KeyFile)
	} else {
		// Connects to local/stacked etcd existing in the cluster
		etcdClient, err = etcdutil.NewFromCluster(client, cfg.CertificatesDir)
	}
	if err != nil {
		return err
	}

	// Compute which upgrade possibilities there are
	klog.V(1).Infoln("[upgrade/plan] computing upgrade possibilities")
	availUpgrades, err := upgrade.GetAvailableUpgrades(versionGetter, flags.allowExperimentalUpgrades, flags.allowRCUpgrades, etcdClient, cfg.DNS.Type, client)
	if err != nil {
		return errors.Wrap(err, "[upgrade/versions] FATAL")
	}

	// No upgrades available
	if len(availUpgrades) == 0 {
		klog.V(1).Infoln("[upgrade/plan] Awesome, you're up-to-date! Enjoy!")
		return nil
	}

	// Generate and print upgrade plans
	for _, up := range availUpgrades {
		plan, unstableVersionFlag, err := genUpgradePlan(&up, isExternalEtcd)
		if err != nil {
			return err
		}

		printUpgradePlan(&up, plan, unstableVersionFlag, isExternalEtcd, os.Stdout)
	}
	return nil
}

// newComponentUpgradePlan helper creates outputapiv1alpha1.ComponentUpgradePlan object
func newComponentUpgradePlan(name, currentVersion, newVersion string) outputapiv1alpha1.ComponentUpgradePlan {
	return outputapiv1alpha1.ComponentUpgradePlan{
		Name:           name,
		CurrentVersion: currentVersion,
		NewVersion:     newVersion,
	}
}

// TODO There is currently no way to cleanly output upgrades that involve adding, removing, or changing components
// https://github.com/kubernetes/kubeadm/issues/810 was created to track addressing this.
func appendDNSComponent(components []outputapiv1alpha1.ComponentUpgradePlan, up *upgrade.Upgrade, DNSType kubeadmapi.DNSAddOnType, name string) []outputapiv1alpha1.ComponentUpgradePlan {
	beforeVersion, afterVersion := "", ""
	if up.Before.DNSType == DNSType {
		beforeVersion = up.Before.DNSVersion
	}
	if up.After.DNSType == DNSType {
		afterVersion = up.After.DNSVersion
	}

	if beforeVersion != "" || afterVersion != "" {
		components = append(components, newComponentUpgradePlan(name, beforeVersion, afterVersion))
	}
	return components
}

// genUpgradePlan generates output-friendly upgrade plan out of upgrade.Upgrade structure
func genUpgradePlan(up *upgrade.Upgrade, isExternalEtcd bool) (*outputapiv1alpha1.UpgradePlan, string, error) {
	newK8sVersion, err := version.ParseSemantic(up.After.KubeVersion)
	if err != nil {
		return nil, "", errors.Wrapf(err, "Unable to parse normalized version %q as a semantic version", up.After.KubeVersion)
	}

	unstableVersionFlag := ""
	if len(newK8sVersion.PreRelease()) != 0 {
		if strings.HasPrefix(newK8sVersion.PreRelease(), "rc") {
			unstableVersionFlag = " --allow-release-candidate-upgrades"
		} else {
			unstableVersionFlag = " --allow-experimental-upgrades"
		}
	}

	components := []outputapiv1alpha1.ComponentUpgradePlan{}

	if isExternalEtcd && up.CanUpgradeEtcd() {
		components = append(components, newComponentUpgradePlan(constants.Etcd, up.Before.EtcdVersion, up.After.EtcdVersion))
	}

	if up.CanUpgradeKubelets() {
		// The map is of the form <old-version>:<node-count>. Here all the keys are put into a slice and sorted
		// in order to always get the right order. Then the map value is extracted separately
		for _, oldVersion := range sortedSliceFromStringIntMap(up.Before.KubeletVersions) {
			nodeCount := up.Before.KubeletVersions[oldVersion]
			components = append(components, newComponentUpgradePlan(constants.Kubelet, fmt.Sprintf("%d x %s", nodeCount, oldVersion), up.After.KubeVersion))
		}
	}

	components = append(components, newComponentUpgradePlan(constants.KubeAPIServer, up.Before.KubeVersion, up.After.KubeVersion))
	components = append(components, newComponentUpgradePlan(constants.KubeControllerManager, up.Before.KubeVersion, up.After.KubeVersion))
	components = append(components, newComponentUpgradePlan(constants.KubeScheduler, up.Before.KubeVersion, up.After.KubeVersion))
	components = append(components, newComponentUpgradePlan(constants.KubeProxy, up.Before.KubeVersion, up.After.KubeVersion))

	components = appendDNSComponent(components, up, kubeadmapi.CoreDNS, constants.CoreDNS)
	components = appendDNSComponent(components, up, kubeadmapi.KubeDNS, constants.KubeDNS)

	if !isExternalEtcd {
		components = append(components, newComponentUpgradePlan(constants.Etcd, up.Before.EtcdVersion, up.After.EtcdVersion))
	}

	return &outputapiv1alpha1.UpgradePlan{Components: components}, unstableVersionFlag, nil
}

// printUpgradePlan prints a UX-friendly overview of what versions are available to upgrade to
func printUpgradePlan(up *upgrade.Upgrade, plan *outputapiv1alpha1.UpgradePlan, unstableVersionFlag string, isExternalEtcd bool, w io.Writer) {
	// The tab writer writes to the "real" writer w
	tabw := tabwriter.NewWriter(w, 10, 4, 3, ' ', 0)

	// endOfTable helper function flashes table writer
	endOfTable := func() {
		tabw.Flush()
		fmt.Fprintln(w, "")
	}

	printHeader := true
	printManualUpgradeHeader := true
	for _, component := range plan.Components {
		if isExternalEtcd && component.Name == constants.Etcd {
			fmt.Fprintln(w, "External components that should be upgraded manually before you upgrade the control plane with 'kubeadm upgrade apply':")
			fmt.Fprintln(tabw, "COMPONENT\tCURRENT\tAVAILABLE")
			fmt.Fprintf(tabw, "%s\t%s\t%s\n", component.Name, component.CurrentVersion, component.NewVersion)
			// end of external components table
			endOfTable()
		} else if component.Name == constants.Kubelet {
			if printManualUpgradeHeader {
				fmt.Fprintln(w, "Components that must be upgraded manually after you have upgraded the control plane with 'kubeadm upgrade apply':")
				fmt.Fprintln(tabw, "COMPONENT\tCURRENT\tAVAILABLE")
				fmt.Fprintf(tabw, "%s\t%s\t%s\n", component.Name, component.CurrentVersion, component.NewVersion)
				printManualUpgradeHeader = false
			} else {
				fmt.Fprintf(tabw, "%s\t%s\t%s\n", "", component.CurrentVersion, component.NewVersion)
			}
		} else {
			if printHeader {
				// End of manual upgrades table
				endOfTable()

				fmt.Fprintf(w, "Upgrade to the latest %s:\n", up.Description)
				fmt.Fprintln(w, "")
				fmt.Fprintln(tabw, "COMPONENT\tCURRENT\tAVAILABLE")
				printHeader = false
			}
			fmt.Fprintf(tabw, "%s\t%s\t%s\n", component.Name, component.CurrentVersion, component.NewVersion)
		}
	}
	// End of control plane table
	endOfTable()

	//fmt.Fprintln(w, "")
	fmt.Fprintln(w, "You can now apply the upgrade by executing the following command:")
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "\tkubeadm upgrade apply %s%s\n", up.After.KubeVersion, unstableVersionFlag)
	fmt.Fprintln(w, "")

	if up.Before.KubeadmVersion != up.After.KubeadmVersion {
		fmt.Fprintf(w, "Note: Before you can perform this upgrade, you have to update kubeadm to %s.\n", up.After.KubeadmVersion)
		fmt.Fprintln(w, "")
	}

	fmt.Fprintln(w, "_____________________________________________________________________")
	fmt.Fprintln(w, "")
}

// sortedSliceFromStringIntMap returns a slice of the keys in the map sorted alphabetically
func sortedSliceFromStringIntMap(strMap map[string]uint16) []string {
	strSlice := []string{}
	for k := range strMap {
		strSlice = append(strSlice, k)
	}
	sort.Strings(strSlice)
	return strSlice
}
