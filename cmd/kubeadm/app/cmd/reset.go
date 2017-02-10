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
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/kubeconfig"
	"k8s.io/kubernetes/cmd/kubeadm/app/preflight"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/pkg/util/initsystem"
)

// NewCmdReset returns the "kubeadm reset" command
func NewCmdReset(out io.Writer) *cobra.Command {
	var skipPreFlight, removeNode bool
	cmd := &cobra.Command{
		Use:   "reset",
		Short: "Run this to revert any changes made to this host by 'kubeadm init' or 'kubeadm join'.",
		Run: func(cmd *cobra.Command, args []string) {
			r, err := NewReset(skipPreFlight, removeNode)
			kubeadmutil.CheckErr(err)
			kubeadmutil.CheckErr(r.Run(out))
		},
	}

	cmd.PersistentFlags().BoolVar(
		&skipPreFlight, "skip-preflight-checks", false,
		"Skip preflight checks normally run before modifying the system",
	)

	cmd.PersistentFlags().BoolVar(
		&removeNode, "remove-node", true,
		"Remove this node from the pool of nodes in this cluster",
	)

	return cmd
}

type Reset struct {
	removeNode bool
}

func NewReset(skipPreFlight, removeNode bool) (*Reset, error) {
	if !skipPreFlight {
		fmt.Println("[preflight] Running pre-flight checks")

		if err := preflight.RunRootCheckOnly(); err != nil {
			return nil, err
		}
	} else {
		fmt.Println("[preflight] Skipping pre-flight checks")
	}

	return &Reset{
		removeNode: removeNode,
	}, nil
}

// Run reverts any changes made to this host by "kubeadm init" or "kubeadm join".
func (r *Reset) Run(out io.Writer) error {

	// Try to drain and remove the node from the cluster
	err := drainAndRemoveNode(r.removeNode)
	if err != nil {
		fmt.Printf("[reset] Failed to cleanup node: [%v]\n", err)
	}

	// Try to stop the kubelet service
	initSystem, err := initsystem.GetInitSystem()
	if err != nil {
		fmt.Println("[reset] WARNING: The kubelet service couldn't be stopped by kubeadm because no supported init system was detected.")
		fmt.Println("[reset] WARNING: Please ensure kubelet is stopped manually.")
	} else {
		fmt.Println("[reset] Stopping the kubelet service")
		if err := initSystem.ServiceStop("kubelet"); err != nil {
			fmt.Printf("[reset] WARNING: The kubelet service couldn't be stopped by kubeadm: [%v]\n", err)
			fmt.Println("[reset] WARNING: Please ensure kubelet is stopped manually.")
		}
	}

	// Try to unmount mounted directories under /var/lib/kubelet in order to be able to remove the /var/lib/kubelet directory later
	fmt.Printf("[reset] Unmounting mounted directories in %q\n", "/var/lib/kubelet")
	umountDirsCmd := "cat /proc/mounts | awk '{print $2}' | grep '/var/lib/kubelet' | xargs -r umount"
	umountOutputBytes, err := exec.Command("sh", "-c", umountDirsCmd).Output()
	if err != nil {
		fmt.Printf("[reset] Failed to unmount mounted directories in /var/lib/kubelet: %s\n", string(umountOutputBytes))
	}

	dockerCheck := preflight.ServiceCheck{Service: "docker", CheckIfActive: true}
	if warnings, errors := dockerCheck.Check(); len(warnings) == 0 && len(errors) == 0 {
		fmt.Println("[reset] Removing kubernetes-managed containers")
		if err := exec.Command("sh", "-c", "docker ps | grep 'k8s_' | awk '{print $1}' | xargs -r docker rm --force --volumes").Run(); err != nil {
			fmt.Println("[reset] Failed to stop the running containers")
		}
	} else {
		fmt.Println("[reset] docker doesn't seem to be running, skipping the removal of running kubernetes containers")
	}

	dirsToClean := []string{"/var/lib/kubelet", "/etc/cni/net.d"}

	// Only clear etcd data when the etcd manifest is found. In case it is not found, we must assume that the user
	// provided external etcd endpoints. In that case, it is his own responsibility to reset etcd
	etcdManifestPath := filepath.Join(kubeadmapi.GlobalEnvParams.KubernetesDir, "manifests/etcd.yaml")
	if _, err := os.Stat(etcdManifestPath); err == nil {
		dirsToClean = append(dirsToClean, "/var/lib/etcd")
	} else {
		fmt.Printf("[reset] No etcd manifest found in %q, assuming external etcd.\n", etcdManifestPath)
	}

	// Then clean contents from the stateful kubelet, etcd and cni directories
	fmt.Printf("[reset] Deleting contents of stateful directories: %v\n", dirsToClean)
	for _, dir := range dirsToClean {
		cleanDir(dir)
	}

	// Remove contents from the config and pki directories
	resetConfigDir(kubeadmapi.GlobalEnvParams.KubernetesDir, kubeadmapi.GlobalEnvParams.HostPKIPath)

	return nil
}

func drainAndRemoveNode(removeNode bool) error {
	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("failed to detect node hostname")
	}
	hostname = strings.ToLower(hostname)

	// TODO: Use the "native" k8s client for this once we're confident the versioned is working
	kubeConfigPath := filepath.Join(kubeadmapi.GlobalEnvParams.KubernetesDir, kubeconfig.KubeletKubeConfigFileName)

	getNodesCmd := fmt.Sprintf("kubectl --kubeconfig %s get nodes | grep %s", kubeConfigPath, hostname)
	output, err := exec.Command("sh", "-c", getNodesCmd).Output()
	if err != nil {
		// kubeadm shouldn't drain and/or remove the node when it doesn't exist anymore
		return fmt.Errorf("failed to list nodes: %v", err)
	}
	if len(output) == 0 {
		return fmt.Errorf("list nodes request returned zero entries")
	}

	fmt.Printf("[reset] Draining node: %q\n", hostname)

	_, err = exec.Command("kubectl", "--kubeconfig", kubeConfigPath, "drain", hostname, "--delete-local-data", "--force", "--ignore-daemonsets").Output()
	if err != nil {
		return fmt.Errorf("failed to drain node %q: %v", hostname, err)
	}

	if removeNode {
		fmt.Printf("[reset] Removing node: %q\n", hostname)
		_, err = exec.Command("kubectl", "--kubeconfig", kubeConfigPath, "delete", "node", hostname).Output()
		if err != nil {
			return fmt.Errorf("failed to remove node %q: %v", hostname, err)
		}
	}

	return nil
}

// cleanDir removes everything in a directory, but not the directory itself
func cleanDir(filePath string) error {
	// If the directory doesn't even exist there's nothing to do, and we do
	// not consider this an error
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil
	}

	d, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(filePath, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// resetConfigDir is used to cleanup the files kubeadm writes in /etc/kubernetes/.
func resetConfigDir(configPathDir, pkiPathDir string) {
	dirsToClean := []string{
		filepath.Join(configPathDir, "manifests"),
		pkiPathDir,
	}
	fmt.Printf("[reset] Deleting contents of config directories: %v\n", dirsToClean)
	for _, dir := range dirsToClean {
		err := cleanDir(dir)
		if err != nil {
			fmt.Printf("[reset] Failed to remove directory: %q [%v]\n", dir, err)
		}
	}

	filesToClean := []string{
		filepath.Join(configPathDir, kubeconfig.AdminKubeConfigFileName),
		filepath.Join(configPathDir, kubeconfig.KubeletKubeConfigFileName),
		filepath.Join(configPathDir, kubeadmconstants.CACertName),
	}
	fmt.Printf("[reset] Deleting files: %v\n", filesToClean)
	for _, path := range filesToClean {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Printf("[reset] Failed to remove file: %q [%v]\n", path, err)
		}
	}
}
