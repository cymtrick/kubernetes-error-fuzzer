/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

// DO NOT EDIT. THIS FILE IS AUTO-GENERATED BY $KUBEROOT/hack/update-generated-deep-copies.sh.

package v1

import (
	time "time"

	api "k8s.io/kubernetes/pkg/api"
	resource "k8s.io/kubernetes/pkg/api/resource"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	intstr "k8s.io/kubernetes/pkg/util/intstr"
	inf "speter.net/go/exp/math/dec/inf"
)

func deepCopy_resource_Quantity(in resource.Quantity, out *resource.Quantity, c *conversion.Cloner) error {
	if in.Amount != nil {
		if newVal, err := c.DeepCopy(in.Amount); err != nil {
			return err
		} else {
			out.Amount = newVal.(*inf.Dec)
		}
	} else {
		out.Amount = nil
	}
	out.Format = in.Format
	return nil
}

func deepCopy_unversioned_ListMeta(in unversioned.ListMeta, out *unversioned.ListMeta, c *conversion.Cloner) error {
	out.SelfLink = in.SelfLink
	out.ResourceVersion = in.ResourceVersion
	return nil
}

func deepCopy_unversioned_Time(in unversioned.Time, out *unversioned.Time, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.Time); err != nil {
		return err
	} else {
		out.Time = newVal.(time.Time)
	}
	return nil
}

func deepCopy_unversioned_TypeMeta(in unversioned.TypeMeta, out *unversioned.TypeMeta, c *conversion.Cloner) error {
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	return nil
}

func deepCopy_v1_AWSElasticBlockStoreVolumeSource(in v1.AWSElasticBlockStoreVolumeSource, out *v1.AWSElasticBlockStoreVolumeSource, c *conversion.Cloner) error {
	out.VolumeID = in.VolumeID
	out.FSType = in.FSType
	out.Partition = in.Partition
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_v1_AzureFileVolumeSource(in v1.AzureFileVolumeSource, out *v1.AzureFileVolumeSource, c *conversion.Cloner) error {
	out.SecretName = in.SecretName
	out.ShareName = in.ShareName
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_v1_Capabilities(in v1.Capabilities, out *v1.Capabilities, c *conversion.Cloner) error {
	if in.Add != nil {
		out.Add = make([]v1.Capability, len(in.Add))
		for i := range in.Add {
			out.Add[i] = in.Add[i]
		}
	} else {
		out.Add = nil
	}
	if in.Drop != nil {
		out.Drop = make([]v1.Capability, len(in.Drop))
		for i := range in.Drop {
			out.Drop[i] = in.Drop[i]
		}
	} else {
		out.Drop = nil
	}
	return nil
}

func deepCopy_v1_CephFSVolumeSource(in v1.CephFSVolumeSource, out *v1.CephFSVolumeSource, c *conversion.Cloner) error {
	if in.Monitors != nil {
		out.Monitors = make([]string, len(in.Monitors))
		for i := range in.Monitors {
			out.Monitors[i] = in.Monitors[i]
		}
	} else {
		out.Monitors = nil
	}
	out.Path = in.Path
	out.User = in.User
	out.SecretFile = in.SecretFile
	if in.SecretRef != nil {
		out.SecretRef = new(v1.LocalObjectReference)
		if err := deepCopy_v1_LocalObjectReference(*in.SecretRef, out.SecretRef, c); err != nil {
			return err
		}
	} else {
		out.SecretRef = nil
	}
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_v1_CinderVolumeSource(in v1.CinderVolumeSource, out *v1.CinderVolumeSource, c *conversion.Cloner) error {
	out.VolumeID = in.VolumeID
	out.FSType = in.FSType
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_v1_ConfigMapKeySelector(in v1.ConfigMapKeySelector, out *v1.ConfigMapKeySelector, c *conversion.Cloner) error {
	if err := deepCopy_v1_LocalObjectReference(in.LocalObjectReference, &out.LocalObjectReference, c); err != nil {
		return err
	}
	out.Key = in.Key
	return nil
}

func deepCopy_v1_ConfigMapVolumeSource(in v1.ConfigMapVolumeSource, out *v1.ConfigMapVolumeSource, c *conversion.Cloner) error {
	if err := deepCopy_v1_LocalObjectReference(in.LocalObjectReference, &out.LocalObjectReference, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]v1.KeyToPath, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_KeyToPath(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_Container(in v1.Container, out *v1.Container, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Image = in.Image
	if in.Command != nil {
		out.Command = make([]string, len(in.Command))
		for i := range in.Command {
			out.Command[i] = in.Command[i]
		}
	} else {
		out.Command = nil
	}
	if in.Args != nil {
		out.Args = make([]string, len(in.Args))
		for i := range in.Args {
			out.Args[i] = in.Args[i]
		}
	} else {
		out.Args = nil
	}
	out.WorkingDir = in.WorkingDir
	if in.Ports != nil {
		out.Ports = make([]v1.ContainerPort, len(in.Ports))
		for i := range in.Ports {
			if err := deepCopy_v1_ContainerPort(in.Ports[i], &out.Ports[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Ports = nil
	}
	if in.Env != nil {
		out.Env = make([]v1.EnvVar, len(in.Env))
		for i := range in.Env {
			if err := deepCopy_v1_EnvVar(in.Env[i], &out.Env[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	if err := deepCopy_v1_ResourceRequirements(in.Resources, &out.Resources, c); err != nil {
		return err
	}
	if in.VolumeMounts != nil {
		out.VolumeMounts = make([]v1.VolumeMount, len(in.VolumeMounts))
		for i := range in.VolumeMounts {
			if err := deepCopy_v1_VolumeMount(in.VolumeMounts[i], &out.VolumeMounts[i], c); err != nil {
				return err
			}
		}
	} else {
		out.VolumeMounts = nil
	}
	if in.LivenessProbe != nil {
		out.LivenessProbe = new(v1.Probe)
		if err := deepCopy_v1_Probe(*in.LivenessProbe, out.LivenessProbe, c); err != nil {
			return err
		}
	} else {
		out.LivenessProbe = nil
	}
	if in.ReadinessProbe != nil {
		out.ReadinessProbe = new(v1.Probe)
		if err := deepCopy_v1_Probe(*in.ReadinessProbe, out.ReadinessProbe, c); err != nil {
			return err
		}
	} else {
		out.ReadinessProbe = nil
	}
	if in.Lifecycle != nil {
		out.Lifecycle = new(v1.Lifecycle)
		if err := deepCopy_v1_Lifecycle(*in.Lifecycle, out.Lifecycle, c); err != nil {
			return err
		}
	} else {
		out.Lifecycle = nil
	}
	out.TerminationMessagePath = in.TerminationMessagePath
	out.ImagePullPolicy = in.ImagePullPolicy
	if in.SecurityContext != nil {
		out.SecurityContext = new(v1.SecurityContext)
		if err := deepCopy_v1_SecurityContext(*in.SecurityContext, out.SecurityContext, c); err != nil {
			return err
		}
	} else {
		out.SecurityContext = nil
	}
	out.Stdin = in.Stdin
	out.StdinOnce = in.StdinOnce
	out.TTY = in.TTY
	return nil
}

func deepCopy_v1_ContainerPort(in v1.ContainerPort, out *v1.ContainerPort, c *conversion.Cloner) error {
	out.Name = in.Name
	out.HostPort = in.HostPort
	out.ContainerPort = in.ContainerPort
	out.Protocol = in.Protocol
	out.HostIP = in.HostIP
	return nil
}

func deepCopy_v1_DownwardAPIVolumeFile(in v1.DownwardAPIVolumeFile, out *v1.DownwardAPIVolumeFile, c *conversion.Cloner) error {
	out.Path = in.Path
	if err := deepCopy_v1_ObjectFieldSelector(in.FieldRef, &out.FieldRef, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_DownwardAPIVolumeSource(in v1.DownwardAPIVolumeSource, out *v1.DownwardAPIVolumeSource, c *conversion.Cloner) error {
	if in.Items != nil {
		out.Items = make([]v1.DownwardAPIVolumeFile, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_DownwardAPIVolumeFile(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_EmptyDirVolumeSource(in v1.EmptyDirVolumeSource, out *v1.EmptyDirVolumeSource, c *conversion.Cloner) error {
	out.Medium = in.Medium
	return nil
}

func deepCopy_v1_EnvVar(in v1.EnvVar, out *v1.EnvVar, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Value = in.Value
	if in.ValueFrom != nil {
		out.ValueFrom = new(v1.EnvVarSource)
		if err := deepCopy_v1_EnvVarSource(*in.ValueFrom, out.ValueFrom, c); err != nil {
			return err
		}
	} else {
		out.ValueFrom = nil
	}
	return nil
}

func deepCopy_v1_EnvVarSource(in v1.EnvVarSource, out *v1.EnvVarSource, c *conversion.Cloner) error {
	if in.FieldRef != nil {
		out.FieldRef = new(v1.ObjectFieldSelector)
		if err := deepCopy_v1_ObjectFieldSelector(*in.FieldRef, out.FieldRef, c); err != nil {
			return err
		}
	} else {
		out.FieldRef = nil
	}
	if in.ConfigMapKeyRef != nil {
		out.ConfigMapKeyRef = new(v1.ConfigMapKeySelector)
		if err := deepCopy_v1_ConfigMapKeySelector(*in.ConfigMapKeyRef, out.ConfigMapKeyRef, c); err != nil {
			return err
		}
	} else {
		out.ConfigMapKeyRef = nil
	}
	if in.SecretKeyRef != nil {
		out.SecretKeyRef = new(v1.SecretKeySelector)
		if err := deepCopy_v1_SecretKeySelector(*in.SecretKeyRef, out.SecretKeyRef, c); err != nil {
			return err
		}
	} else {
		out.SecretKeyRef = nil
	}
	return nil
}

func deepCopy_v1_ExecAction(in v1.ExecAction, out *v1.ExecAction, c *conversion.Cloner) error {
	if in.Command != nil {
		out.Command = make([]string, len(in.Command))
		for i := range in.Command {
			out.Command[i] = in.Command[i]
		}
	} else {
		out.Command = nil
	}
	return nil
}

func deepCopy_v1_FCVolumeSource(in v1.FCVolumeSource, out *v1.FCVolumeSource, c *conversion.Cloner) error {
	if in.TargetWWNs != nil {
		out.TargetWWNs = make([]string, len(in.TargetWWNs))
		for i := range in.TargetWWNs {
			out.TargetWWNs[i] = in.TargetWWNs[i]
		}
	} else {
		out.TargetWWNs = nil
	}
	if in.Lun != nil {
		out.Lun = new(int32)
		*out.Lun = *in.Lun
	} else {
		out.Lun = nil
	}
	out.FSType = in.FSType
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_v1_FlexVolumeSource(in v1.FlexVolumeSource, out *v1.FlexVolumeSource, c *conversion.Cloner) error {
	out.Driver = in.Driver
	out.FSType = in.FSType
	if in.SecretRef != nil {
		out.SecretRef = new(v1.LocalObjectReference)
		if err := deepCopy_v1_LocalObjectReference(*in.SecretRef, out.SecretRef, c); err != nil {
			return err
		}
	} else {
		out.SecretRef = nil
	}
	out.ReadOnly = in.ReadOnly
	if in.Options != nil {
		out.Options = make(map[string]string)
		for key, val := range in.Options {
			out.Options[key] = val
		}
	} else {
		out.Options = nil
	}
	return nil
}

func deepCopy_v1_FlockerVolumeSource(in v1.FlockerVolumeSource, out *v1.FlockerVolumeSource, c *conversion.Cloner) error {
	out.DatasetName = in.DatasetName
	return nil
}

func deepCopy_v1_GCEPersistentDiskVolumeSource(in v1.GCEPersistentDiskVolumeSource, out *v1.GCEPersistentDiskVolumeSource, c *conversion.Cloner) error {
	out.PDName = in.PDName
	out.FSType = in.FSType
	out.Partition = in.Partition
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_v1_GitRepoVolumeSource(in v1.GitRepoVolumeSource, out *v1.GitRepoVolumeSource, c *conversion.Cloner) error {
	out.Repository = in.Repository
	out.Revision = in.Revision
	out.Directory = in.Directory
	return nil
}

func deepCopy_v1_GlusterfsVolumeSource(in v1.GlusterfsVolumeSource, out *v1.GlusterfsVolumeSource, c *conversion.Cloner) error {
	out.EndpointsName = in.EndpointsName
	out.Path = in.Path
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_v1_HTTPGetAction(in v1.HTTPGetAction, out *v1.HTTPGetAction, c *conversion.Cloner) error {
	out.Path = in.Path
	if err := deepCopy_intstr_IntOrString(in.Port, &out.Port, c); err != nil {
		return err
	}
	out.Host = in.Host
	out.Scheme = in.Scheme
	if in.HTTPHeaders != nil {
		out.HTTPHeaders = make([]v1.HTTPHeader, len(in.HTTPHeaders))
		for i := range in.HTTPHeaders {
			if err := deepCopy_v1_HTTPHeader(in.HTTPHeaders[i], &out.HTTPHeaders[i], c); err != nil {
				return err
			}
		}
	} else {
		out.HTTPHeaders = nil
	}
	return nil
}

func deepCopy_v1_HTTPHeader(in v1.HTTPHeader, out *v1.HTTPHeader, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Value = in.Value
	return nil
}

func deepCopy_v1_Handler(in v1.Handler, out *v1.Handler, c *conversion.Cloner) error {
	if in.Exec != nil {
		out.Exec = new(v1.ExecAction)
		if err := deepCopy_v1_ExecAction(*in.Exec, out.Exec, c); err != nil {
			return err
		}
	} else {
		out.Exec = nil
	}
	if in.HTTPGet != nil {
		out.HTTPGet = new(v1.HTTPGetAction)
		if err := deepCopy_v1_HTTPGetAction(*in.HTTPGet, out.HTTPGet, c); err != nil {
			return err
		}
	} else {
		out.HTTPGet = nil
	}
	if in.TCPSocket != nil {
		out.TCPSocket = new(v1.TCPSocketAction)
		if err := deepCopy_v1_TCPSocketAction(*in.TCPSocket, out.TCPSocket, c); err != nil {
			return err
		}
	} else {
		out.TCPSocket = nil
	}
	return nil
}

func deepCopy_v1_HostPathVolumeSource(in v1.HostPathVolumeSource, out *v1.HostPathVolumeSource, c *conversion.Cloner) error {
	out.Path = in.Path
	return nil
}

func deepCopy_v1_ISCSIVolumeSource(in v1.ISCSIVolumeSource, out *v1.ISCSIVolumeSource, c *conversion.Cloner) error {
	out.TargetPortal = in.TargetPortal
	out.IQN = in.IQN
	out.Lun = in.Lun
	out.ISCSIInterface = in.ISCSIInterface
	out.FSType = in.FSType
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_v1_KeyToPath(in v1.KeyToPath, out *v1.KeyToPath, c *conversion.Cloner) error {
	out.Key = in.Key
	out.Path = in.Path
	return nil
}

func deepCopy_v1_Lifecycle(in v1.Lifecycle, out *v1.Lifecycle, c *conversion.Cloner) error {
	if in.PostStart != nil {
		out.PostStart = new(v1.Handler)
		if err := deepCopy_v1_Handler(*in.PostStart, out.PostStart, c); err != nil {
			return err
		}
	} else {
		out.PostStart = nil
	}
	if in.PreStop != nil {
		out.PreStop = new(v1.Handler)
		if err := deepCopy_v1_Handler(*in.PreStop, out.PreStop, c); err != nil {
			return err
		}
	} else {
		out.PreStop = nil
	}
	return nil
}

func deepCopy_v1_LocalObjectReference(in v1.LocalObjectReference, out *v1.LocalObjectReference, c *conversion.Cloner) error {
	out.Name = in.Name
	return nil
}

func deepCopy_v1_NFSVolumeSource(in v1.NFSVolumeSource, out *v1.NFSVolumeSource, c *conversion.Cloner) error {
	out.Server = in.Server
	out.Path = in.Path
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_v1_ObjectFieldSelector(in v1.ObjectFieldSelector, out *v1.ObjectFieldSelector, c *conversion.Cloner) error {
	out.APIVersion = in.APIVersion
	out.FieldPath = in.FieldPath
	return nil
}

func deepCopy_v1_ObjectMeta(in v1.ObjectMeta, out *v1.ObjectMeta, c *conversion.Cloner) error {
	out.Name = in.Name
	out.GenerateName = in.GenerateName
	out.Namespace = in.Namespace
	out.SelfLink = in.SelfLink
	out.UID = in.UID
	out.ResourceVersion = in.ResourceVersion
	out.Generation = in.Generation
	if err := deepCopy_unversioned_Time(in.CreationTimestamp, &out.CreationTimestamp, c); err != nil {
		return err
	}
	if in.DeletionTimestamp != nil {
		out.DeletionTimestamp = new(unversioned.Time)
		if err := deepCopy_unversioned_Time(*in.DeletionTimestamp, out.DeletionTimestamp, c); err != nil {
			return err
		}
	} else {
		out.DeletionTimestamp = nil
	}
	if in.DeletionGracePeriodSeconds != nil {
		out.DeletionGracePeriodSeconds = new(int64)
		*out.DeletionGracePeriodSeconds = *in.DeletionGracePeriodSeconds
	} else {
		out.DeletionGracePeriodSeconds = nil
	}
	if in.Labels != nil {
		out.Labels = make(map[string]string)
		for key, val := range in.Labels {
			out.Labels[key] = val
		}
	} else {
		out.Labels = nil
	}
	if in.Annotations != nil {
		out.Annotations = make(map[string]string)
		for key, val := range in.Annotations {
			out.Annotations[key] = val
		}
	} else {
		out.Annotations = nil
	}
	return nil
}

func deepCopy_v1_PersistentVolumeClaimVolumeSource(in v1.PersistentVolumeClaimVolumeSource, out *v1.PersistentVolumeClaimVolumeSource, c *conversion.Cloner) error {
	out.ClaimName = in.ClaimName
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_v1_PodSecurityContext(in v1.PodSecurityContext, out *v1.PodSecurityContext, c *conversion.Cloner) error {
	if in.SELinuxOptions != nil {
		out.SELinuxOptions = new(v1.SELinuxOptions)
		if err := deepCopy_v1_SELinuxOptions(*in.SELinuxOptions, out.SELinuxOptions, c); err != nil {
			return err
		}
	} else {
		out.SELinuxOptions = nil
	}
	if in.RunAsUser != nil {
		out.RunAsUser = new(int64)
		*out.RunAsUser = *in.RunAsUser
	} else {
		out.RunAsUser = nil
	}
	if in.RunAsNonRoot != nil {
		out.RunAsNonRoot = new(bool)
		*out.RunAsNonRoot = *in.RunAsNonRoot
	} else {
		out.RunAsNonRoot = nil
	}
	if in.SupplementalGroups != nil {
		out.SupplementalGroups = make([]int64, len(in.SupplementalGroups))
		for i := range in.SupplementalGroups {
			out.SupplementalGroups[i] = in.SupplementalGroups[i]
		}
	} else {
		out.SupplementalGroups = nil
	}
	if in.FSGroup != nil {
		out.FSGroup = new(int64)
		*out.FSGroup = *in.FSGroup
	} else {
		out.FSGroup = nil
	}
	return nil
}

func deepCopy_v1_PodSpec(in v1.PodSpec, out *v1.PodSpec, c *conversion.Cloner) error {
	if in.Volumes != nil {
		out.Volumes = make([]v1.Volume, len(in.Volumes))
		for i := range in.Volumes {
			if err := deepCopy_v1_Volume(in.Volumes[i], &out.Volumes[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Volumes = nil
	}
	if in.Containers != nil {
		out.Containers = make([]v1.Container, len(in.Containers))
		for i := range in.Containers {
			if err := deepCopy_v1_Container(in.Containers[i], &out.Containers[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Containers = nil
	}
	out.RestartPolicy = in.RestartPolicy
	if in.TerminationGracePeriodSeconds != nil {
		out.TerminationGracePeriodSeconds = new(int64)
		*out.TerminationGracePeriodSeconds = *in.TerminationGracePeriodSeconds
	} else {
		out.TerminationGracePeriodSeconds = nil
	}
	if in.ActiveDeadlineSeconds != nil {
		out.ActiveDeadlineSeconds = new(int64)
		*out.ActiveDeadlineSeconds = *in.ActiveDeadlineSeconds
	} else {
		out.ActiveDeadlineSeconds = nil
	}
	out.DNSPolicy = in.DNSPolicy
	if in.NodeSelector != nil {
		out.NodeSelector = make(map[string]string)
		for key, val := range in.NodeSelector {
			out.NodeSelector[key] = val
		}
	} else {
		out.NodeSelector = nil
	}
	out.ServiceAccountName = in.ServiceAccountName
	out.DeprecatedServiceAccount = in.DeprecatedServiceAccount
	out.NodeName = in.NodeName
	out.HostNetwork = in.HostNetwork
	out.HostPID = in.HostPID
	out.HostIPC = in.HostIPC
	if in.SecurityContext != nil {
		out.SecurityContext = new(v1.PodSecurityContext)
		if err := deepCopy_v1_PodSecurityContext(*in.SecurityContext, out.SecurityContext, c); err != nil {
			return err
		}
	} else {
		out.SecurityContext = nil
	}
	if in.ImagePullSecrets != nil {
		out.ImagePullSecrets = make([]v1.LocalObjectReference, len(in.ImagePullSecrets))
		for i := range in.ImagePullSecrets {
			if err := deepCopy_v1_LocalObjectReference(in.ImagePullSecrets[i], &out.ImagePullSecrets[i], c); err != nil {
				return err
			}
		}
	} else {
		out.ImagePullSecrets = nil
	}
	return nil
}

func deepCopy_v1_PodTemplateSpec(in v1.PodTemplateSpec, out *v1.PodTemplateSpec, c *conversion.Cloner) error {
	if err := deepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_v1_PodSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_Probe(in v1.Probe, out *v1.Probe, c *conversion.Cloner) error {
	if err := deepCopy_v1_Handler(in.Handler, &out.Handler, c); err != nil {
		return err
	}
	out.InitialDelaySeconds = in.InitialDelaySeconds
	out.TimeoutSeconds = in.TimeoutSeconds
	out.PeriodSeconds = in.PeriodSeconds
	out.SuccessThreshold = in.SuccessThreshold
	out.FailureThreshold = in.FailureThreshold
	return nil
}

func deepCopy_v1_RBDVolumeSource(in v1.RBDVolumeSource, out *v1.RBDVolumeSource, c *conversion.Cloner) error {
	if in.CephMonitors != nil {
		out.CephMonitors = make([]string, len(in.CephMonitors))
		for i := range in.CephMonitors {
			out.CephMonitors[i] = in.CephMonitors[i]
		}
	} else {
		out.CephMonitors = nil
	}
	out.RBDImage = in.RBDImage
	out.FSType = in.FSType
	out.RBDPool = in.RBDPool
	out.RadosUser = in.RadosUser
	out.Keyring = in.Keyring
	if in.SecretRef != nil {
		out.SecretRef = new(v1.LocalObjectReference)
		if err := deepCopy_v1_LocalObjectReference(*in.SecretRef, out.SecretRef, c); err != nil {
			return err
		}
	} else {
		out.SecretRef = nil
	}
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_v1_ResourceRequirements(in v1.ResourceRequirements, out *v1.ResourceRequirements, c *conversion.Cloner) error {
	if in.Limits != nil {
		out.Limits = make(v1.ResourceList)
		for key, val := range in.Limits {
			newVal := new(resource.Quantity)
			if err := deepCopy_resource_Quantity(val, newVal, c); err != nil {
				return err
			}
			out.Limits[key] = *newVal
		}
	} else {
		out.Limits = nil
	}
	if in.Requests != nil {
		out.Requests = make(v1.ResourceList)
		for key, val := range in.Requests {
			newVal := new(resource.Quantity)
			if err := deepCopy_resource_Quantity(val, newVal, c); err != nil {
				return err
			}
			out.Requests[key] = *newVal
		}
	} else {
		out.Requests = nil
	}
	return nil
}

func deepCopy_v1_SELinuxOptions(in v1.SELinuxOptions, out *v1.SELinuxOptions, c *conversion.Cloner) error {
	out.User = in.User
	out.Role = in.Role
	out.Type = in.Type
	out.Level = in.Level
	return nil
}

func deepCopy_v1_SecretKeySelector(in v1.SecretKeySelector, out *v1.SecretKeySelector, c *conversion.Cloner) error {
	if err := deepCopy_v1_LocalObjectReference(in.LocalObjectReference, &out.LocalObjectReference, c); err != nil {
		return err
	}
	out.Key = in.Key
	return nil
}

func deepCopy_v1_SecretVolumeSource(in v1.SecretVolumeSource, out *v1.SecretVolumeSource, c *conversion.Cloner) error {
	out.SecretName = in.SecretName
	return nil
}

func deepCopy_v1_SecurityContext(in v1.SecurityContext, out *v1.SecurityContext, c *conversion.Cloner) error {
	if in.Capabilities != nil {
		out.Capabilities = new(v1.Capabilities)
		if err := deepCopy_v1_Capabilities(*in.Capabilities, out.Capabilities, c); err != nil {
			return err
		}
	} else {
		out.Capabilities = nil
	}
	if in.Privileged != nil {
		out.Privileged = new(bool)
		*out.Privileged = *in.Privileged
	} else {
		out.Privileged = nil
	}
	if in.SELinuxOptions != nil {
		out.SELinuxOptions = new(v1.SELinuxOptions)
		if err := deepCopy_v1_SELinuxOptions(*in.SELinuxOptions, out.SELinuxOptions, c); err != nil {
			return err
		}
	} else {
		out.SELinuxOptions = nil
	}
	if in.RunAsUser != nil {
		out.RunAsUser = new(int64)
		*out.RunAsUser = *in.RunAsUser
	} else {
		out.RunAsUser = nil
	}
	if in.RunAsNonRoot != nil {
		out.RunAsNonRoot = new(bool)
		*out.RunAsNonRoot = *in.RunAsNonRoot
	} else {
		out.RunAsNonRoot = nil
	}
	if in.ReadOnlyRootFilesystem != nil {
		out.ReadOnlyRootFilesystem = new(bool)
		*out.ReadOnlyRootFilesystem = *in.ReadOnlyRootFilesystem
	} else {
		out.ReadOnlyRootFilesystem = nil
	}
	return nil
}

func deepCopy_v1_TCPSocketAction(in v1.TCPSocketAction, out *v1.TCPSocketAction, c *conversion.Cloner) error {
	if err := deepCopy_intstr_IntOrString(in.Port, &out.Port, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_Volume(in v1.Volume, out *v1.Volume, c *conversion.Cloner) error {
	out.Name = in.Name
	if err := deepCopy_v1_VolumeSource(in.VolumeSource, &out.VolumeSource, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_VolumeMount(in v1.VolumeMount, out *v1.VolumeMount, c *conversion.Cloner) error {
	out.Name = in.Name
	out.ReadOnly = in.ReadOnly
	out.MountPath = in.MountPath
	return nil
}

func deepCopy_v1_VolumeSource(in v1.VolumeSource, out *v1.VolumeSource, c *conversion.Cloner) error {
	if in.HostPath != nil {
		out.HostPath = new(v1.HostPathVolumeSource)
		if err := deepCopy_v1_HostPathVolumeSource(*in.HostPath, out.HostPath, c); err != nil {
			return err
		}
	} else {
		out.HostPath = nil
	}
	if in.EmptyDir != nil {
		out.EmptyDir = new(v1.EmptyDirVolumeSource)
		if err := deepCopy_v1_EmptyDirVolumeSource(*in.EmptyDir, out.EmptyDir, c); err != nil {
			return err
		}
	} else {
		out.EmptyDir = nil
	}
	if in.GCEPersistentDisk != nil {
		out.GCEPersistentDisk = new(v1.GCEPersistentDiskVolumeSource)
		if err := deepCopy_v1_GCEPersistentDiskVolumeSource(*in.GCEPersistentDisk, out.GCEPersistentDisk, c); err != nil {
			return err
		}
	} else {
		out.GCEPersistentDisk = nil
	}
	if in.AWSElasticBlockStore != nil {
		out.AWSElasticBlockStore = new(v1.AWSElasticBlockStoreVolumeSource)
		if err := deepCopy_v1_AWSElasticBlockStoreVolumeSource(*in.AWSElasticBlockStore, out.AWSElasticBlockStore, c); err != nil {
			return err
		}
	} else {
		out.AWSElasticBlockStore = nil
	}
	if in.GitRepo != nil {
		out.GitRepo = new(v1.GitRepoVolumeSource)
		if err := deepCopy_v1_GitRepoVolumeSource(*in.GitRepo, out.GitRepo, c); err != nil {
			return err
		}
	} else {
		out.GitRepo = nil
	}
	if in.Secret != nil {
		out.Secret = new(v1.SecretVolumeSource)
		if err := deepCopy_v1_SecretVolumeSource(*in.Secret, out.Secret, c); err != nil {
			return err
		}
	} else {
		out.Secret = nil
	}
	if in.NFS != nil {
		out.NFS = new(v1.NFSVolumeSource)
		if err := deepCopy_v1_NFSVolumeSource(*in.NFS, out.NFS, c); err != nil {
			return err
		}
	} else {
		out.NFS = nil
	}
	if in.ISCSI != nil {
		out.ISCSI = new(v1.ISCSIVolumeSource)
		if err := deepCopy_v1_ISCSIVolumeSource(*in.ISCSI, out.ISCSI, c); err != nil {
			return err
		}
	} else {
		out.ISCSI = nil
	}
	if in.Glusterfs != nil {
		out.Glusterfs = new(v1.GlusterfsVolumeSource)
		if err := deepCopy_v1_GlusterfsVolumeSource(*in.Glusterfs, out.Glusterfs, c); err != nil {
			return err
		}
	} else {
		out.Glusterfs = nil
	}
	if in.PersistentVolumeClaim != nil {
		out.PersistentVolumeClaim = new(v1.PersistentVolumeClaimVolumeSource)
		if err := deepCopy_v1_PersistentVolumeClaimVolumeSource(*in.PersistentVolumeClaim, out.PersistentVolumeClaim, c); err != nil {
			return err
		}
	} else {
		out.PersistentVolumeClaim = nil
	}
	if in.RBD != nil {
		out.RBD = new(v1.RBDVolumeSource)
		if err := deepCopy_v1_RBDVolumeSource(*in.RBD, out.RBD, c); err != nil {
			return err
		}
	} else {
		out.RBD = nil
	}
	if in.FlexVolume != nil {
		out.FlexVolume = new(v1.FlexVolumeSource)
		if err := deepCopy_v1_FlexVolumeSource(*in.FlexVolume, out.FlexVolume, c); err != nil {
			return err
		}
	} else {
		out.FlexVolume = nil
	}
	if in.Cinder != nil {
		out.Cinder = new(v1.CinderVolumeSource)
		if err := deepCopy_v1_CinderVolumeSource(*in.Cinder, out.Cinder, c); err != nil {
			return err
		}
	} else {
		out.Cinder = nil
	}
	if in.CephFS != nil {
		out.CephFS = new(v1.CephFSVolumeSource)
		if err := deepCopy_v1_CephFSVolumeSource(*in.CephFS, out.CephFS, c); err != nil {
			return err
		}
	} else {
		out.CephFS = nil
	}
	if in.Flocker != nil {
		out.Flocker = new(v1.FlockerVolumeSource)
		if err := deepCopy_v1_FlockerVolumeSource(*in.Flocker, out.Flocker, c); err != nil {
			return err
		}
	} else {
		out.Flocker = nil
	}
	if in.DownwardAPI != nil {
		out.DownwardAPI = new(v1.DownwardAPIVolumeSource)
		if err := deepCopy_v1_DownwardAPIVolumeSource(*in.DownwardAPI, out.DownwardAPI, c); err != nil {
			return err
		}
	} else {
		out.DownwardAPI = nil
	}
	if in.FC != nil {
		out.FC = new(v1.FCVolumeSource)
		if err := deepCopy_v1_FCVolumeSource(*in.FC, out.FC, c); err != nil {
			return err
		}
	} else {
		out.FC = nil
	}
	if in.AzureFile != nil {
		out.AzureFile = new(v1.AzureFileVolumeSource)
		if err := deepCopy_v1_AzureFileVolumeSource(*in.AzureFile, out.AzureFile, c); err != nil {
			return err
		}
	} else {
		out.AzureFile = nil
	}
	if in.ConfigMap != nil {
		out.ConfigMap = new(v1.ConfigMapVolumeSource)
		if err := deepCopy_v1_ConfigMapVolumeSource(*in.ConfigMap, out.ConfigMap, c); err != nil {
			return err
		}
	} else {
		out.ConfigMap = nil
	}
	return nil
}

func deepCopy_v1_Job(in Job, out *Job, c *conversion.Cloner) error {
	if err := deepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_v1_JobSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_v1_JobStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_JobCondition(in JobCondition, out *JobCondition, c *conversion.Cloner) error {
	out.Type = in.Type
	out.Status = in.Status
	if err := deepCopy_unversioned_Time(in.LastProbeTime, &out.LastProbeTime, c); err != nil {
		return err
	}
	if err := deepCopy_unversioned_Time(in.LastTransitionTime, &out.LastTransitionTime, c); err != nil {
		return err
	}
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func deepCopy_v1_JobList(in JobList, out *JobList, c *conversion.Cloner) error {
	if err := deepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]Job, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_Job(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_JobSpec(in JobSpec, out *JobSpec, c *conversion.Cloner) error {
	if in.Parallelism != nil {
		out.Parallelism = new(int32)
		*out.Parallelism = *in.Parallelism
	} else {
		out.Parallelism = nil
	}
	if in.Completions != nil {
		out.Completions = new(int32)
		*out.Completions = *in.Completions
	} else {
		out.Completions = nil
	}
	if in.ActiveDeadlineSeconds != nil {
		out.ActiveDeadlineSeconds = new(int64)
		*out.ActiveDeadlineSeconds = *in.ActiveDeadlineSeconds
	} else {
		out.ActiveDeadlineSeconds = nil
	}
	if in.Selector != nil {
		out.Selector = new(LabelSelector)
		if err := deepCopy_v1_LabelSelector(*in.Selector, out.Selector, c); err != nil {
			return err
		}
	} else {
		out.Selector = nil
	}
	if in.ManualSelector != nil {
		out.ManualSelector = new(bool)
		*out.ManualSelector = *in.ManualSelector
	} else {
		out.ManualSelector = nil
	}
	if err := deepCopy_v1_PodTemplateSpec(in.Template, &out.Template, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_JobStatus(in JobStatus, out *JobStatus, c *conversion.Cloner) error {
	if in.Conditions != nil {
		out.Conditions = make([]JobCondition, len(in.Conditions))
		for i := range in.Conditions {
			if err := deepCopy_v1_JobCondition(in.Conditions[i], &out.Conditions[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	if in.StartTime != nil {
		out.StartTime = new(unversioned.Time)
		if err := deepCopy_unversioned_Time(*in.StartTime, out.StartTime, c); err != nil {
			return err
		}
	} else {
		out.StartTime = nil
	}
	if in.CompletionTime != nil {
		out.CompletionTime = new(unversioned.Time)
		if err := deepCopy_unversioned_Time(*in.CompletionTime, out.CompletionTime, c); err != nil {
			return err
		}
	} else {
		out.CompletionTime = nil
	}
	out.Active = in.Active
	out.Succeeded = in.Succeeded
	out.Failed = in.Failed
	return nil
}

func deepCopy_v1_LabelSelector(in LabelSelector, out *LabelSelector, c *conversion.Cloner) error {
	if in.MatchLabels != nil {
		out.MatchLabels = make(map[string]string)
		for key, val := range in.MatchLabels {
			out.MatchLabels[key] = val
		}
	} else {
		out.MatchLabels = nil
	}
	if in.MatchExpressions != nil {
		out.MatchExpressions = make([]LabelSelectorRequirement, len(in.MatchExpressions))
		for i := range in.MatchExpressions {
			if err := deepCopy_v1_LabelSelectorRequirement(in.MatchExpressions[i], &out.MatchExpressions[i], c); err != nil {
				return err
			}
		}
	} else {
		out.MatchExpressions = nil
	}
	return nil
}

func deepCopy_v1_LabelSelectorRequirement(in LabelSelectorRequirement, out *LabelSelectorRequirement, c *conversion.Cloner) error {
	out.Key = in.Key
	out.Operator = in.Operator
	if in.Values != nil {
		out.Values = make([]string, len(in.Values))
		for i := range in.Values {
			out.Values[i] = in.Values[i]
		}
	} else {
		out.Values = nil
	}
	return nil
}

func deepCopy_intstr_IntOrString(in intstr.IntOrString, out *intstr.IntOrString, c *conversion.Cloner) error {
	out.Type = in.Type
	out.IntVal = in.IntVal
	out.StrVal = in.StrVal
	return nil
}

func init() {
	err := api.Scheme.AddGeneratedDeepCopyFuncs(
		deepCopy_resource_Quantity,
		deepCopy_unversioned_ListMeta,
		deepCopy_unversioned_Time,
		deepCopy_unversioned_TypeMeta,
		deepCopy_v1_AWSElasticBlockStoreVolumeSource,
		deepCopy_v1_AzureFileVolumeSource,
		deepCopy_v1_Capabilities,
		deepCopy_v1_CephFSVolumeSource,
		deepCopy_v1_CinderVolumeSource,
		deepCopy_v1_ConfigMapKeySelector,
		deepCopy_v1_ConfigMapVolumeSource,
		deepCopy_v1_Container,
		deepCopy_v1_ContainerPort,
		deepCopy_v1_DownwardAPIVolumeFile,
		deepCopy_v1_DownwardAPIVolumeSource,
		deepCopy_v1_EmptyDirVolumeSource,
		deepCopy_v1_EnvVar,
		deepCopy_v1_EnvVarSource,
		deepCopy_v1_ExecAction,
		deepCopy_v1_FCVolumeSource,
		deepCopy_v1_FlexVolumeSource,
		deepCopy_v1_FlockerVolumeSource,
		deepCopy_v1_GCEPersistentDiskVolumeSource,
		deepCopy_v1_GitRepoVolumeSource,
		deepCopy_v1_GlusterfsVolumeSource,
		deepCopy_v1_HTTPGetAction,
		deepCopy_v1_HTTPHeader,
		deepCopy_v1_Handler,
		deepCopy_v1_HostPathVolumeSource,
		deepCopy_v1_ISCSIVolumeSource,
		deepCopy_v1_KeyToPath,
		deepCopy_v1_Lifecycle,
		deepCopy_v1_LocalObjectReference,
		deepCopy_v1_NFSVolumeSource,
		deepCopy_v1_ObjectFieldSelector,
		deepCopy_v1_ObjectMeta,
		deepCopy_v1_PersistentVolumeClaimVolumeSource,
		deepCopy_v1_PodSecurityContext,
		deepCopy_v1_PodSpec,
		deepCopy_v1_PodTemplateSpec,
		deepCopy_v1_Probe,
		deepCopy_v1_RBDVolumeSource,
		deepCopy_v1_ResourceRequirements,
		deepCopy_v1_SELinuxOptions,
		deepCopy_v1_SecretKeySelector,
		deepCopy_v1_SecretVolumeSource,
		deepCopy_v1_SecurityContext,
		deepCopy_v1_TCPSocketAction,
		deepCopy_v1_Volume,
		deepCopy_v1_VolumeMount,
		deepCopy_v1_VolumeSource,
		deepCopy_v1_Job,
		deepCopy_v1_JobCondition,
		deepCopy_v1_JobList,
		deepCopy_v1_JobSpec,
		deepCopy_v1_JobStatus,
		deepCopy_v1_LabelSelector,
		deepCopy_v1_LabelSelectorRequirement,
		deepCopy_intstr_IntOrString,
	)
	if err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}
