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

package v1alpha1

import (
	"reflect"
	"strings"

	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	// Add non-generated conversion functions
	err := scheme.AddConversionFuncs(
		Convert_v1alpha1_MasterConfiguration_To_kubeadm_MasterConfiguration,
		Convert_v1alpha1_Etcd_To_kubeadm_Etcd,
		Convert_kubeadm_Etcd_To_v1alpha1_Etcd,
	)
	if err != nil {
		return err
	}

	return nil
}

func Convert_v1alpha1_MasterConfiguration_To_kubeadm_MasterConfiguration(in *MasterConfiguration, out *kubeadm.MasterConfiguration, s conversion.Scope) error {
	if err := autoConvert_v1alpha1_MasterConfiguration_To_kubeadm_MasterConfiguration(in, out, s); err != nil {
		return err
	}

	UpgradeCloudProvider(in, out)
	UpgradeAuthorizationModes(in, out)
	// We don't support migrating information from the .PrivilegedPods field which was removed in v1alpha2
	// We don't support migrating information from the .ImagePullPolicy field which was removed in v1alpha2

	return nil
}

func Convert_v1alpha1_Etcd_To_kubeadm_Etcd(in *Etcd, out *kubeadm.Etcd, s conversion.Scope) error {
	if err := autoConvert_v1alpha1_Etcd_To_kubeadm_Etcd(in, out, s); err != nil {
		return err
	}

	// The .Etcd schema changed between v1alpha1 and v1alpha2 API types. The change was to basically only split up the fields into two sub-structs, which can be seen here
	if len(in.Endpoints) != 0 {
		out.External = &kubeadm.ExternalEtcd{
			Endpoints: in.Endpoints,
			CAFile:    in.CAFile,
			CertFile:  in.CertFile,
			KeyFile:   in.KeyFile,
		}
	} else {
		out.Local = &kubeadm.LocalEtcd{
			Image:          in.Image,
			DataDir:        in.DataDir,
			ExtraArgs:      in.ExtraArgs,
			ServerCertSANs: in.ServerCertSANs,
			PeerCertSANs:   in.PeerCertSANs,
		}
	}

	// No need to transfer information about .Etcd.Selfhosted to v1alpha2
	return nil
}

// no-op, as we don't support converting from newer API to old alpha API
func Convert_kubeadm_Etcd_To_v1alpha1_Etcd(in *kubeadm.Etcd, out *Etcd, s conversion.Scope) error {

	if in.External != nil {
		out.Endpoints = in.External.Endpoints
		out.CAFile = in.External.CAFile
		out.CertFile = in.External.CertFile
		out.KeyFile = in.External.KeyFile
	} else {
		out.Image = in.Local.Image
		out.DataDir = in.Local.DataDir
		out.ExtraArgs = in.Local.ExtraArgs
		out.ServerCertSANs = in.Local.ServerCertSANs
		out.PeerCertSANs = in.Local.PeerCertSANs
	}

	return nil
}

// UpgradeCloudProvider handles the removal of .CloudProvider as smoothly as possible
func UpgradeCloudProvider(in *MasterConfiguration, out *kubeadm.MasterConfiguration) {
	if len(in.CloudProvider) != 0 {
		if out.APIServerExtraArgs == nil {
			out.APIServerExtraArgs = map[string]string{}
		}
		if out.ControllerManagerExtraArgs == nil {
			out.ControllerManagerExtraArgs = map[string]string{}
		}

		out.APIServerExtraArgs["cloud-provider"] = in.CloudProvider
		out.ControllerManagerExtraArgs["cloud-provider"] = in.CloudProvider
	}
}

func UpgradeAuthorizationModes(in *MasterConfiguration, out *kubeadm.MasterConfiguration) {
	// If .AuthorizationModes was set to something else than the default, preserve the information via extraargs
	if !reflect.DeepEqual(in.AuthorizationModes, strings.Split(DefaultAuthorizationModes, ",")) {

		if out.APIServerExtraArgs == nil {
			out.APIServerExtraArgs = map[string]string{}
		}
		out.APIServerExtraArgs["authorization-mode"] = strings.Join(in.AuthorizationModes, ",")
	}
}
