/*
Copyright 2022 The Kubernetes Authors.

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

package options

import (
	"github.com/spf13/pflag"
	nodeconfig "k8s.io/cloud-provider/controllers/node/config"
)

// NodeControllerOptions holds the ServiceController options.
type NodeControllerOptions struct {
	*nodeconfig.NodeControllerConfiguration
}

// AddFlags adds flags related to ServiceController for controller manager to the specified FlagSet.
func (o *NodeControllerOptions) AddFlags(fs *pflag.FlagSet) {
	if o == nil {
		return
	}

	fs.Int32Var(&o.WorkerCount, "workerCount", o.WorkerCount, "Number of workers synchronizing node status.")
}

// ApplyTo fills up ServiceController config with options.
func (o *NodeControllerOptions) ApplyTo(cfg *nodeconfig.NodeControllerConfiguration) error {
	if o == nil {
		return nil
	}

	cfg.WorkerCount = o.WorkerCount

	return nil
}

// Validate checks validation of NodeControllerOptions.
func (o *NodeControllerOptions) Validate() []error {
	if o == nil {
		return nil
	}

	return []error{}
}
