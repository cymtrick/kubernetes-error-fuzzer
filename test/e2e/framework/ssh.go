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

package framework

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	sshutil "k8s.io/kubernetes/pkg/ssh"
)

// GetSigner returns an ssh.Signer for the provider ("gce", etc.) that can be
// used to SSH to their nodes.
func GetSigner(provider string) (ssh.Signer, error) {
	// honor a consistent SSH key across all providers
	if path := os.Getenv("KUBE_SSH_KEY_PATH"); len(path) > 0 {
		return sshutil.MakePrivateKeySignerFromFile(path)
	}

	// Select the key itself to use. When implementing more providers here,
	// please also add them to any SSH tests that are disabled because of signer
	// support.
	keyfile := ""
	switch provider {
	case "gce", "gke", "kubemark":
		keyfile = os.Getenv("GCE_SSH_KEY")
		if keyfile == "" {
			keyfile = "google_compute_engine"
		}
	case "aws", "eks":
		keyfile = os.Getenv("AWS_SSH_KEY")
		if keyfile == "" {
			keyfile = "kube_aws_rsa"
		}
	case "local", "vsphere":
		keyfile = os.Getenv("LOCAL_SSH_KEY")
		if keyfile == "" {
			keyfile = "id_rsa"
		}
	case "skeleton":
		keyfile = os.Getenv("KUBE_SSH_KEY")
		if keyfile == "" {
			keyfile = "id_rsa"
		}
	default:
		return nil, fmt.Errorf("GetSigner(...) not implemented for %s", provider)
	}

	// Respect absolute paths for keys given by user, fallback to assuming
	// relative paths are in ~/.ssh
	if !filepath.IsAbs(keyfile) {
		keydir := filepath.Join(os.Getenv("HOME"), ".ssh")
		keyfile = filepath.Join(keydir, keyfile)
	}

	return sshutil.MakePrivateKeySignerFromFile(keyfile)
}

// NodeSSHHosts returns SSH-able host names for all schedulable nodes - this
// excludes master node. If it can't find any external IPs, it falls back to
// looking for internal IPs. If it can't find an internal IP for every node it
// returns an error, though it still returns all hosts that it found in that
// case.
func NodeSSHHosts(c clientset.Interface) ([]string, error) {
	nodelist := waitListSchedulableNodesOrDie(c)

	hosts := NodeAddresses(nodelist, v1.NodeExternalIP)
	// If ExternalIPs aren't set, assume the test programs can reach the
	// InternalIP. Simplified exception logic here assumes that the hosts will
	// either all have ExternalIP or none will. Simplifies handling here and
	// should be adequate since the setting of the external IPs is provider
	// specific: they should either all have them or none of them will.
	if len(hosts) == 0 {
		Logf("No external IP address on nodes, falling back to internal IPs")
		hosts = NodeAddresses(nodelist, v1.NodeInternalIP)
	}

	// Error if any node didn't have an external/internal IP.
	if len(hosts) != len(nodelist.Items) {
		return hosts, fmt.Errorf(
			"only found %d IPs on nodes, but found %d nodes. Nodelist: %v",
			len(hosts), len(nodelist.Items), nodelist)
	}

	sshHosts := make([]string, 0, len(hosts))
	for _, h := range hosts {
		sshHosts = append(sshHosts, net.JoinHostPort(h, sshPort))
	}
	return sshHosts, nil
}

type SSHResult struct {
	User   string
	Host   string
	Cmd    string
	Stdout string
	Stderr string
	Code   int
}

// NodeExec execs the given cmd on node via SSH. Note that the nodeName is an sshable name,
// eg: the name returned by framework.GetMasterHost(). This is also not guaranteed to work across
// cloud providers since it involves ssh.
func NodeExec(nodeName, cmd string) (SSHResult, error) {
	return SSH(cmd, net.JoinHostPort(nodeName, sshPort), TestContext.Provider)
}

// SSH synchronously SSHs to a node running on provider and runs cmd. If there
// is no error performing the SSH, the stdout, stderr, and exit code are
// returned.
func SSH(cmd, host, provider string) (SSHResult, error) {
	result := SSHResult{Host: host, Cmd: cmd}

	// Get a signer for the provider.
	signer, err := GetSigner(provider)
	if err != nil {
		return result, fmt.Errorf("error getting signer for provider %s: '%v'", provider, err)
	}

	// RunSSHCommand will default to Getenv("USER") if user == "", but we're
	// defaulting here as well for logging clarity.
	result.User = os.Getenv("KUBE_SSH_USER")
	if result.User == "" {
		result.User = os.Getenv("USER")
	}

	stdout, stderr, code, err := sshutil.RunSSHCommand(cmd, result.User, host, signer)
	result.Stdout = stdout
	result.Stderr = stderr
	result.Code = code

	return result, err
}

func LogSSHResult(result SSHResult) {
	remote := fmt.Sprintf("%s@%s", result.User, result.Host)
	Logf("ssh %s: command:   %s", remote, result.Cmd)
	Logf("ssh %s: stdout:    %q", remote, result.Stdout)
	Logf("ssh %s: stderr:    %q", remote, result.Stderr)
	Logf("ssh %s: exit code: %d", remote, result.Code)
}

func IssueSSHCommandWithResult(cmd, provider string, node *v1.Node) (*SSHResult, error) {
	Logf("Getting external IP address for %s", node.Name)
	host := ""
	for _, a := range node.Status.Addresses {
		if a.Type == v1.NodeExternalIP {
			host = net.JoinHostPort(a.Address, sshPort)
			break
		}
	}

	if host == "" {
		// No external IPs were found, let's try to use internal as plan B
		for _, a := range node.Status.Addresses {
			if a.Type == v1.NodeInternalIP {
				host = net.JoinHostPort(a.Address, sshPort)
				break
			}
		}
	}

	if host == "" {
		return nil, fmt.Errorf("couldn't find any IP address for node %s", node.Name)
	}

	Logf("SSH %q on %s(%s)", cmd, node.Name, host)
	result, err := SSH(cmd, host, provider)
	LogSSHResult(result)

	if result.Code != 0 || err != nil {
		return nil, fmt.Errorf("failed running %q: %v (exit code %d, stderr %v)",
			cmd, err, result.Code, result.Stderr)
	}

	return &result, nil
}

func IssueSSHCommand(cmd, provider string, node *v1.Node) error {
	_, err := IssueSSHCommandWithResult(cmd, provider, node)
	if err != nil {
		return err
	}
	return nil
}
