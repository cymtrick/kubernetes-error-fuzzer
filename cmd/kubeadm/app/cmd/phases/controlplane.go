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
	"errors"
	"fmt"

	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/controlplane"
	"k8s.io/kubernetes/pkg/util/normalizer"
)

var (
	controlPlaneExample = normalizer.Examples(`
		# Generates all static Pod manifest files for control plane components, 
		# functionally equivalent to what is generated by kubeadm init.
		kubeadm init phase control-plane

		# Generates all static Pod manifest files using options read from a configuration file.
		kubeadm init phase control-plane --config config.yaml
		`)

	controlPlanePhaseProperties = map[string]struct {
		name  string
		short string
	}{
		kubeadmconstants.KubeAPIServer: {
			name:  "apiserver",
			short: getPhaseDescription(kubeadmconstants.KubeAPIServer),
		},
		kubeadmconstants.KubeControllerManager: {
			name:  "controller-manager",
			short: getPhaseDescription(kubeadmconstants.KubeControllerManager),
		},
		kubeadmconstants.KubeScheduler: {
			name:  "scheduler",
			short: getPhaseDescription(kubeadmconstants.KubeScheduler),
		},
	}
)

type controlPlaneData interface {
	Cfg() *kubeadmapi.InitConfiguration
	KubeConfigDir() string
	ManifestDir() string
}

func getPhaseDescription(component string) string {
	return fmt.Sprintf("Generates the %s static Pod manifest", component)
}

// NewControlPlanePhase creates a kubeadm workflow phase that implements bootstrapping the control plane.
func NewControlPlanePhase() workflow.Phase {
	phase := workflow.Phase{
		Name:  "control-plane",
		Short: "Generates all static Pod manifest files necessary to establish the control plane",
		Long:  cmdutil.MacroCommandLongDescription,
		Phases: []workflow.Phase{
			{
				Name:           "all",
				Short:          "Generates all static Pod manifest files",
				InheritFlags:   getControlPlanePhaseFlags("all"),
				RunAllSiblings: true,
			},
			newControlPlaneSubPhase(kubeadmconstants.KubeAPIServer),
			newControlPlaneSubPhase(kubeadmconstants.KubeControllerManager),
			newControlPlaneSubPhase(kubeadmconstants.KubeScheduler),
		},
		Run: runControlPlanePhase,
	}
	return phase
}

func newControlPlaneSubPhase(component string) workflow.Phase {
	phase := workflow.Phase{
		Name:         controlPlanePhaseProperties[component].name,
		Short:        controlPlanePhaseProperties[component].short,
		Run:          runControlPlaneSubPhase(component),
		InheritFlags: getControlPlanePhaseFlags(component),
	}
	return phase
}

func getControlPlanePhaseFlags(name string) []string {
	flags := []string{
		options.CfgPath,
		options.CertificatesDir,
		options.KubernetesVersion,
	}
	if name == "all" || name == kubeadmconstants.KubeAPIServer {
		flags = append(flags,
			options.APIServerAdvertiseAddress,
			options.APIServerBindPort,
			options.APIServerExtraArgs,
			options.FeatureGatesString,
			options.NetworkingServiceSubnet,
		)
	}
	if name == "all" || name == kubeadmconstants.KubeControllerManager {
		flags = append(flags,
			options.ControllerManagerExtraArgs,
			options.NetworkingPodSubnet,
		)
	}
	if name == "all" || name == kubeadmconstants.KubeScheduler {
		flags = append(flags,
			options.SchedulerExtraArgs,
		)
	}
	return flags
}

func runControlPlanePhase(c workflow.RunData) error {
	data, ok := c.(controlPlaneData)
	if !ok {
		return errors.New("control-plane phase invoked with an invalid data struct")
	}

	fmt.Printf("[control-plane] Using manifest folder %q\n", data.ManifestDir())
	return nil
}

func runControlPlaneSubPhase(component string) func(c workflow.RunData) error {
	return func(c workflow.RunData) error {
		data, ok := c.(controlPlaneData)
		if !ok {
			return errors.New("control-plane phase invoked with an invalid data struct")
		}
		cfg := data.Cfg()

		fmt.Printf("[control-plane] Creating static Pod manifest for %q\n", component)
		if err := controlplane.CreateStaticPodFiles(data.ManifestDir(), cfg, component); err != nil {
			return err
		}
		return nil
	}
}
