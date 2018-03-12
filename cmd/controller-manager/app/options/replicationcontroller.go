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

package options

import (
	"github.com/spf13/pflag"
	"k8s.io/kubernetes/pkg/apis/componentconfig"
)

// ReplicationControllerOptions holds the ReplicationController options.
type ReplicationControllerOptions struct {
	ConcurrentRCSyncs int32
}

// AddFlags adds flags related to ReplicationController for controller manager to the specified FlagSet.
func (o *ReplicationControllerOptions) AddFlags(fs *pflag.FlagSet) {
	if o == nil {
		return
	}

	fs.Int32Var(&o.ConcurrentRCSyncs, "concurrent_rc_syncs", o.ConcurrentRCSyncs, "The number of replication controllers that are allowed to sync concurrently. Larger number = more responsive replica management, but more CPU (and network) load")
}

// ApplyTo fills up ReplicationController config with options.
func (o *ReplicationControllerOptions) ApplyTo(cfg *componentconfig.ReplicationControllerConfiguration) error {
	if o == nil {
		return nil
	}

	cfg.ConcurrentRCSyncs = o.ConcurrentRCSyncs

	return nil
}

// Validate checks validation of ReplicationControllerOptions.
func (o *ReplicationControllerOptions) Validate() []error {
	if o == nil {
		return nil
	}

	errs := []error{}
	return errs
}
