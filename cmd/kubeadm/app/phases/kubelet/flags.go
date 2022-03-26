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

package kubelet

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"

	versionutil "k8s.io/apimachinery/pkg/util/version"
	componentversion "k8s.io/component-base/version"
	"k8s.io/klog/v2"
	utilsexec "k8s.io/utils/exec"

	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/images"
	preflight "k8s.io/kubernetes/cmd/kubeadm/app/preflight"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
)

type kubeletFlagsOpts struct {
	nodeRegOpts              *kubeadmapi.NodeRegistrationOptions
	pauseImage               string
	registerTaintsUsingFlags bool
	// This is a temporary measure until kubeadm no longer supports a kubelet version with built-in dockershim.
	// TODO: https://github.com/kubernetes/kubeadm/issues/2626
	kubeletVersion *versionutil.Version
}

// GetNodeNameAndHostname obtains the name for this Node using the following precedence
// (from lower to higher):
// - actual hostname
// - NodeRegistrationOptions.Name (same as "--node-name" passed to "kubeadm init/join")
// - "hostname-override" flag in NodeRegistrationOptions.KubeletExtraArgs
// It also returns the hostname or an error if getting the hostname failed.
func GetNodeNameAndHostname(cfg *kubeadmapi.NodeRegistrationOptions) (string, string, error) {
	hostname, err := kubeadmutil.GetHostname("")
	nodeName := hostname
	if cfg.Name != "" {
		nodeName = cfg.Name
	}
	if name, ok := cfg.KubeletExtraArgs["hostname-override"]; ok {
		nodeName = name
	}
	return nodeName, hostname, err
}

// WriteKubeletDynamicEnvFile writes an environment file with dynamic flags to the kubelet.
// Used at "kubeadm init" and "kubeadm join" time.
func WriteKubeletDynamicEnvFile(cfg *kubeadmapi.ClusterConfiguration, nodeReg *kubeadmapi.NodeRegistrationOptions, registerTaintsUsingFlags bool, kubeletDir string) error {
	// This is a temporary measure until kubeadm no longer supports a kubelet version with built-in dockershim.
	// TODO: https://github.com/kubernetes/kubeadm/issues/2626
	kubeletVersion, err := preflight.GetKubeletVersion(utilsexec.New())
	if err != nil {
		// We cannot return an error here, due to the k/k CI, where /cmd/kubeadm/test tests run without
		// a kubelet built on the host. On error, we assume a kubelet version equal to the version
		// of the kubeadm binary. During normal cluster creation this should not happens as kubeadm needs
		// the kubelet binary for init / join.
		kubeletVersion = versionutil.MustParseSemantic(componentversion.Get().GitVersion)
		klog.Warningf("cannot obtain the version of the kubelet while writing dynamic environment file: %v."+
			" Using the version of the kubeadm binary: %s", err, kubeletVersion.String())
	}

	flagOpts := kubeletFlagsOpts{
		nodeRegOpts:              nodeReg,
		pauseImage:               images.GetPauseImage(cfg),
		registerTaintsUsingFlags: registerTaintsUsingFlags,
		kubeletVersion:           kubeletVersion,
	}
	stringMap := buildKubeletArgMap(flagOpts)
	argList := kubeadmutil.BuildArgumentListFromMap(stringMap, nodeReg.KubeletExtraArgs)
	envFileContent := fmt.Sprintf("%s=%q\n", constants.KubeletEnvFileVariableName, strings.Join(argList, " "))

	return writeKubeletFlagBytesToDisk([]byte(envFileContent), kubeletDir)
}

//buildKubeletArgMapCommon takes a kubeletFlagsOpts object and builds based on that a string-string map with flags
//that are common to both Linux and Windows
func buildKubeletArgMapCommon(opts kubeletFlagsOpts) map[string]string {
	kubeletFlags := map[string]string{}

	// This is a temporary measure until kubeadm no longer supports a kubelet version with built-in dockershim.
	// Once that happens only the "remote" branch option should be left.
	// TODO: https://github.com/kubernetes/kubeadm/issues/2626
	hasDockershim := opts.kubeletVersion.Major() == 1 && opts.kubeletVersion.Minor() < 24
	var dockerSocket string
	if runtime.GOOS == "windows" {
		dockerSocket = "npipe:////./pipe/dockershim"
	} else {
		dockerSocket = "unix:///var/run/dockershim.sock"
	}
	if opts.nodeRegOpts.CRISocket == dockerSocket && hasDockershim {
		kubeletFlags["network-plugin"] = "cni"
	} else {
		kubeletFlags["container-runtime"] = "remote"
		kubeletFlags["container-runtime-endpoint"] = opts.nodeRegOpts.CRISocket
	}

	// This flag passes the pod infra container image (e.g. "pause" image) to the kubelet
	// and prevents its garbage collection
	if opts.pauseImage != "" {
		kubeletFlags["pod-infra-container-image"] = opts.pauseImage
	}

	if opts.registerTaintsUsingFlags && opts.nodeRegOpts.Taints != nil && len(opts.nodeRegOpts.Taints) > 0 {
		taintStrs := []string{}
		for _, taint := range opts.nodeRegOpts.Taints {
			taintStrs = append(taintStrs, taint.ToString())
		}

		kubeletFlags["register-with-taints"] = strings.Join(taintStrs, ",")
	}

	// Pass the "--hostname-override" flag to the kubelet only if it's different from the hostname
	nodeName, hostname, err := GetNodeNameAndHostname(opts.nodeRegOpts)
	if err != nil {
		klog.Warning(err)
	}
	if nodeName != hostname {
		klog.V(1).Infof("setting kubelet hostname-override to %q", nodeName)
		kubeletFlags["hostname-override"] = nodeName
	}

	return kubeletFlags
}

// writeKubeletFlagBytesToDisk writes a byte slice down to disk at the specific location of the kubelet flag overrides file
func writeKubeletFlagBytesToDisk(b []byte, kubeletDir string) error {
	kubeletEnvFilePath := filepath.Join(kubeletDir, constants.KubeletEnvFileName)
	fmt.Printf("[kubelet-start] Writing kubelet environment file with flags to file %q\n", kubeletEnvFilePath)

	// creates target folder if not already exists
	if err := os.MkdirAll(kubeletDir, 0700); err != nil {
		return errors.Wrapf(err, "failed to create directory %q", kubeletDir)
	}
	if err := os.WriteFile(kubeletEnvFilePath, b, 0644); err != nil {
		return errors.Wrapf(err, "failed to write kubelet configuration to the file %q", kubeletEnvFilePath)
	}
	return nil
}

// buildKubeletArgMap takes a kubeletFlagsOpts object and builds based on that a string-string map with flags
// that should be given to the local kubelet daemon.
func buildKubeletArgMap(opts kubeletFlagsOpts) map[string]string {
	return buildKubeletArgMapCommon(opts)
}
