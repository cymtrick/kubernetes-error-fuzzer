// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentbaseconfigv1alpha1 "k8s.io/component-base/config/v1alpha1"
	v1alpha1 "k8s.io/kube-controller-manager/config/v1alpha1"
	config "k8s.io/kubernetes/pkg/controller/apis/config"
	signerconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/certificates/signer/config/v1alpha1"
	daemonconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/daemon/config/v1alpha1"
	deploymentconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/deployment/config/v1alpha1"
	endpointconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/endpoint/config/v1alpha1"
	endpointsliceconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/endpointslice/config/v1alpha1"
	endpointslicemirroringconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/endpointslicemirroring/config/v1alpha1"
	garbagecollectorconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/garbagecollector/config/v1alpha1"
	jobconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/job/config/v1alpha1"
	namespaceconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/namespace/config/v1alpha1"
	nodeipamconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/nodeipam/config/v1alpha1"
	nodelifecycleconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/nodelifecycle/config/v1alpha1"
	podautoscalerconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/podautoscaler/config/v1alpha1"
	podgcconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/podgc/config/v1alpha1"
	replicasetconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/replicaset/config/v1alpha1"
	replicationconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/replication/config/v1alpha1"
	resourcequotaconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/resourcequota/config/v1alpha1"
	serviceconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/service/config/v1alpha1"
	serviceaccountconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/serviceaccount/config/v1alpha1"
	statefulsetconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/statefulset/config/v1alpha1"
	ttlafterfinishedconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/ttlafterfinished/config/v1alpha1"
	attachdetachconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/volume/attachdetach/config/v1alpha1"
	persistentvolumeconfigv1alpha1 "k8s.io/kubernetes/pkg/controller/volume/persistentvolume/config/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha1.CloudProviderConfiguration)(nil), (*config.CloudProviderConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudProviderConfiguration_To_config_CloudProviderConfiguration(a.(*v1alpha1.CloudProviderConfiguration), b.(*config.CloudProviderConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.CloudProviderConfiguration)(nil), (*v1alpha1.CloudProviderConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_CloudProviderConfiguration_To_v1alpha1_CloudProviderConfiguration(a.(*config.CloudProviderConfiguration), b.(*v1alpha1.CloudProviderConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.DeprecatedControllerConfiguration)(nil), (*config.DeprecatedControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DeprecatedControllerConfiguration_To_config_DeprecatedControllerConfiguration(a.(*v1alpha1.DeprecatedControllerConfiguration), b.(*config.DeprecatedControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.DeprecatedControllerConfiguration)(nil), (*v1alpha1.DeprecatedControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DeprecatedControllerConfiguration_To_v1alpha1_DeprecatedControllerConfiguration(a.(*config.DeprecatedControllerConfiguration), b.(*v1alpha1.DeprecatedControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.GroupResource)(nil), (*v1.GroupResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_GroupResource_To_v1_GroupResource(a.(*v1alpha1.GroupResource), b.(*v1.GroupResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GroupResource)(nil), (*v1alpha1.GroupResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GroupResource_To_v1alpha1_GroupResource(a.(*v1.GroupResource), b.(*v1alpha1.GroupResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.KubeControllerManagerConfiguration)(nil), (*config.KubeControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubeControllerManagerConfiguration_To_config_KubeControllerManagerConfiguration(a.(*v1alpha1.KubeControllerManagerConfiguration), b.(*config.KubeControllerManagerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeControllerManagerConfiguration)(nil), (*v1alpha1.KubeControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration(a.(*config.KubeControllerManagerConfiguration), b.(*v1alpha1.KubeControllerManagerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.GenericControllerManagerConfiguration)(nil), (*v1alpha1.GenericControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_GenericControllerManagerConfiguration_To_v1alpha1_GenericControllerManagerConfiguration(a.(*config.GenericControllerManagerConfiguration), b.(*v1alpha1.GenericControllerManagerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.KubeCloudSharedConfiguration)(nil), (*v1alpha1.KubeCloudSharedConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeCloudSharedConfiguration_To_v1alpha1_KubeCloudSharedConfiguration(a.(*config.KubeCloudSharedConfiguration), b.(*v1alpha1.KubeCloudSharedConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha1.GenericControllerManagerConfiguration)(nil), (*config.GenericControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(a.(*v1alpha1.GenericControllerManagerConfiguration), b.(*config.GenericControllerManagerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha1.KubeCloudSharedConfiguration)(nil), (*config.KubeCloudSharedConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubeCloudSharedConfiguration_To_config_KubeCloudSharedConfiguration(a.(*v1alpha1.KubeCloudSharedConfiguration), b.(*config.KubeCloudSharedConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_CloudProviderConfiguration_To_config_CloudProviderConfiguration(in *v1alpha1.CloudProviderConfiguration, out *config.CloudProviderConfiguration, s conversion.Scope) error {
	out.Name = in.Name
	out.CloudConfigFile = in.CloudConfigFile
	return nil
}

// Convert_v1alpha1_CloudProviderConfiguration_To_config_CloudProviderConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_CloudProviderConfiguration_To_config_CloudProviderConfiguration(in *v1alpha1.CloudProviderConfiguration, out *config.CloudProviderConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudProviderConfiguration_To_config_CloudProviderConfiguration(in, out, s)
}

func autoConvert_config_CloudProviderConfiguration_To_v1alpha1_CloudProviderConfiguration(in *config.CloudProviderConfiguration, out *v1alpha1.CloudProviderConfiguration, s conversion.Scope) error {
	out.Name = in.Name
	out.CloudConfigFile = in.CloudConfigFile
	return nil
}

// Convert_config_CloudProviderConfiguration_To_v1alpha1_CloudProviderConfiguration is an autogenerated conversion function.
func Convert_config_CloudProviderConfiguration_To_v1alpha1_CloudProviderConfiguration(in *config.CloudProviderConfiguration, out *v1alpha1.CloudProviderConfiguration, s conversion.Scope) error {
	return autoConvert_config_CloudProviderConfiguration_To_v1alpha1_CloudProviderConfiguration(in, out, s)
}

func autoConvert_v1alpha1_DeprecatedControllerConfiguration_To_config_DeprecatedControllerConfiguration(in *v1alpha1.DeprecatedControllerConfiguration, out *config.DeprecatedControllerConfiguration, s conversion.Scope) error {
	out.DeletingPodsQPS = in.DeletingPodsQPS
	out.DeletingPodsBurst = in.DeletingPodsBurst
	out.RegisterRetryCount = in.RegisterRetryCount
	return nil
}

// Convert_v1alpha1_DeprecatedControllerConfiguration_To_config_DeprecatedControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_DeprecatedControllerConfiguration_To_config_DeprecatedControllerConfiguration(in *v1alpha1.DeprecatedControllerConfiguration, out *config.DeprecatedControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_DeprecatedControllerConfiguration_To_config_DeprecatedControllerConfiguration(in, out, s)
}

func autoConvert_config_DeprecatedControllerConfiguration_To_v1alpha1_DeprecatedControllerConfiguration(in *config.DeprecatedControllerConfiguration, out *v1alpha1.DeprecatedControllerConfiguration, s conversion.Scope) error {
	out.DeletingPodsQPS = in.DeletingPodsQPS
	out.DeletingPodsBurst = in.DeletingPodsBurst
	out.RegisterRetryCount = in.RegisterRetryCount
	return nil
}

// Convert_config_DeprecatedControllerConfiguration_To_v1alpha1_DeprecatedControllerConfiguration is an autogenerated conversion function.
func Convert_config_DeprecatedControllerConfiguration_To_v1alpha1_DeprecatedControllerConfiguration(in *config.DeprecatedControllerConfiguration, out *v1alpha1.DeprecatedControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_DeprecatedControllerConfiguration_To_v1alpha1_DeprecatedControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(in *v1alpha1.GenericControllerManagerConfiguration, out *config.GenericControllerManagerConfiguration, s conversion.Scope) error {
	out.Port = in.Port
	out.Address = in.Address
	out.MinResyncPeriod = in.MinResyncPeriod
	if err := componentbaseconfigv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.ControllerStartInterval = in.ControllerStartInterval
	if err := componentbaseconfigv1alpha1.Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.Controllers = *(*[]string)(unsafe.Pointer(&in.Controllers))
	if err := componentbaseconfigv1alpha1.Convert_v1alpha1_DebuggingConfiguration_To_config_DebuggingConfiguration(&in.Debugging, &out.Debugging, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_config_GenericControllerManagerConfiguration_To_v1alpha1_GenericControllerManagerConfiguration(in *config.GenericControllerManagerConfiguration, out *v1alpha1.GenericControllerManagerConfiguration, s conversion.Scope) error {
	out.Port = in.Port
	out.Address = in.Address
	out.MinResyncPeriod = in.MinResyncPeriod
	if err := componentbaseconfigv1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.ControllerStartInterval = in.ControllerStartInterval
	if err := componentbaseconfigv1alpha1.Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.Controllers = *(*[]string)(unsafe.Pointer(&in.Controllers))
	if err := componentbaseconfigv1alpha1.Convert_config_DebuggingConfiguration_To_v1alpha1_DebuggingConfiguration(&in.Debugging, &out.Debugging, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_GroupResource_To_v1_GroupResource(in *v1alpha1.GroupResource, out *v1.GroupResource, s conversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}

// Convert_v1alpha1_GroupResource_To_v1_GroupResource is an autogenerated conversion function.
func Convert_v1alpha1_GroupResource_To_v1_GroupResource(in *v1alpha1.GroupResource, out *v1.GroupResource, s conversion.Scope) error {
	return autoConvert_v1alpha1_GroupResource_To_v1_GroupResource(in, out, s)
}

func autoConvert_v1_GroupResource_To_v1alpha1_GroupResource(in *v1.GroupResource, out *v1alpha1.GroupResource, s conversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}

// Convert_v1_GroupResource_To_v1alpha1_GroupResource is an autogenerated conversion function.
func Convert_v1_GroupResource_To_v1alpha1_GroupResource(in *v1.GroupResource, out *v1alpha1.GroupResource, s conversion.Scope) error {
	return autoConvert_v1_GroupResource_To_v1alpha1_GroupResource(in, out, s)
}

func autoConvert_v1alpha1_KubeCloudSharedConfiguration_To_config_KubeCloudSharedConfiguration(in *v1alpha1.KubeCloudSharedConfiguration, out *config.KubeCloudSharedConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_CloudProviderConfiguration_To_config_CloudProviderConfiguration(&in.CloudProvider, &out.CloudProvider, s); err != nil {
		return err
	}
	out.ExternalCloudVolumePlugin = in.ExternalCloudVolumePlugin
	out.UseServiceAccountCredentials = in.UseServiceAccountCredentials
	out.AllowUntaggedCloud = in.AllowUntaggedCloud
	out.RouteReconciliationPeriod = in.RouteReconciliationPeriod
	out.NodeMonitorPeriod = in.NodeMonitorPeriod
	out.ClusterName = in.ClusterName
	out.ClusterCIDR = in.ClusterCIDR
	out.AllocateNodeCIDRs = in.AllocateNodeCIDRs
	out.CIDRAllocatorType = in.CIDRAllocatorType
	if err := v1.Convert_Pointer_bool_To_bool(&in.ConfigureCloudRoutes, &out.ConfigureCloudRoutes, s); err != nil {
		return err
	}
	out.NodeSyncPeriod = in.NodeSyncPeriod
	return nil
}

func autoConvert_config_KubeCloudSharedConfiguration_To_v1alpha1_KubeCloudSharedConfiguration(in *config.KubeCloudSharedConfiguration, out *v1alpha1.KubeCloudSharedConfiguration, s conversion.Scope) error {
	if err := Convert_config_CloudProviderConfiguration_To_v1alpha1_CloudProviderConfiguration(&in.CloudProvider, &out.CloudProvider, s); err != nil {
		return err
	}
	out.ExternalCloudVolumePlugin = in.ExternalCloudVolumePlugin
	out.UseServiceAccountCredentials = in.UseServiceAccountCredentials
	out.AllowUntaggedCloud = in.AllowUntaggedCloud
	out.RouteReconciliationPeriod = in.RouteReconciliationPeriod
	out.NodeMonitorPeriod = in.NodeMonitorPeriod
	out.ClusterName = in.ClusterName
	out.ClusterCIDR = in.ClusterCIDR
	out.AllocateNodeCIDRs = in.AllocateNodeCIDRs
	out.CIDRAllocatorType = in.CIDRAllocatorType
	if err := v1.Convert_bool_To_Pointer_bool(&in.ConfigureCloudRoutes, &out.ConfigureCloudRoutes, s); err != nil {
		return err
	}
	out.NodeSyncPeriod = in.NodeSyncPeriod
	return nil
}

func autoConvert_v1alpha1_KubeControllerManagerConfiguration_To_config_KubeControllerManagerConfiguration(in *v1alpha1.KubeControllerManagerConfiguration, out *config.KubeControllerManagerConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(&in.Generic, &out.Generic, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_KubeCloudSharedConfiguration_To_config_KubeCloudSharedConfiguration(&in.KubeCloudShared, &out.KubeCloudShared, s); err != nil {
		return err
	}
	if err := attachdetachconfigv1alpha1.Convert_v1alpha1_AttachDetachControllerConfiguration_To_config_AttachDetachControllerConfiguration(&in.AttachDetachController, &out.AttachDetachController, s); err != nil {
		return err
	}
	if err := signerconfigv1alpha1.Convert_v1alpha1_CSRSigningControllerConfiguration_To_config_CSRSigningControllerConfiguration(&in.CSRSigningController, &out.CSRSigningController, s); err != nil {
		return err
	}
	if err := daemonconfigv1alpha1.Convert_v1alpha1_DaemonSetControllerConfiguration_To_config_DaemonSetControllerConfiguration(&in.DaemonSetController, &out.DaemonSetController, s); err != nil {
		return err
	}
	if err := deploymentconfigv1alpha1.Convert_v1alpha1_DeploymentControllerConfiguration_To_config_DeploymentControllerConfiguration(&in.DeploymentController, &out.DeploymentController, s); err != nil {
		return err
	}
	if err := statefulsetconfigv1alpha1.Convert_v1alpha1_StatefulSetControllerConfiguration_To_config_StatefulSetControllerConfiguration(&in.StatefulSetController, &out.StatefulSetController, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_DeprecatedControllerConfiguration_To_config_DeprecatedControllerConfiguration(&in.DeprecatedController, &out.DeprecatedController, s); err != nil {
		return err
	}
	if err := endpointconfigv1alpha1.Convert_v1alpha1_EndpointControllerConfiguration_To_config_EndpointControllerConfiguration(&in.EndpointController, &out.EndpointController, s); err != nil {
		return err
	}
	if err := endpointsliceconfigv1alpha1.Convert_v1alpha1_EndpointSliceControllerConfiguration_To_config_EndpointSliceControllerConfiguration(&in.EndpointSliceController, &out.EndpointSliceController, s); err != nil {
		return err
	}
	if err := endpointslicemirroringconfigv1alpha1.Convert_v1alpha1_EndpointSliceMirroringControllerConfiguration_To_config_EndpointSliceMirroringControllerConfiguration(&in.EndpointSliceMirroringController, &out.EndpointSliceMirroringController, s); err != nil {
		return err
	}
	if err := garbagecollectorconfigv1alpha1.Convert_v1alpha1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(&in.GarbageCollectorController, &out.GarbageCollectorController, s); err != nil {
		return err
	}
	if err := podautoscalerconfigv1alpha1.Convert_v1alpha1_HPAControllerConfiguration_To_config_HPAControllerConfiguration(&in.HPAController, &out.HPAController, s); err != nil {
		return err
	}
	if err := jobconfigv1alpha1.Convert_v1alpha1_JobControllerConfiguration_To_config_JobControllerConfiguration(&in.JobController, &out.JobController, s); err != nil {
		return err
	}
	if err := namespaceconfigv1alpha1.Convert_v1alpha1_NamespaceControllerConfiguration_To_config_NamespaceControllerConfiguration(&in.NamespaceController, &out.NamespaceController, s); err != nil {
		return err
	}
	if err := nodeipamconfigv1alpha1.Convert_v1alpha1_NodeIPAMControllerConfiguration_To_config_NodeIPAMControllerConfiguration(&in.NodeIPAMController, &out.NodeIPAMController, s); err != nil {
		return err
	}
	if err := nodelifecycleconfigv1alpha1.Convert_v1alpha1_NodeLifecycleControllerConfiguration_To_config_NodeLifecycleControllerConfiguration(&in.NodeLifecycleController, &out.NodeLifecycleController, s); err != nil {
		return err
	}
	if err := persistentvolumeconfigv1alpha1.Convert_v1alpha1_PersistentVolumeBinderControllerConfiguration_To_config_PersistentVolumeBinderControllerConfiguration(&in.PersistentVolumeBinderController, &out.PersistentVolumeBinderController, s); err != nil {
		return err
	}
	if err := podgcconfigv1alpha1.Convert_v1alpha1_PodGCControllerConfiguration_To_config_PodGCControllerConfiguration(&in.PodGCController, &out.PodGCController, s); err != nil {
		return err
	}
	if err := replicasetconfigv1alpha1.Convert_v1alpha1_ReplicaSetControllerConfiguration_To_config_ReplicaSetControllerConfiguration(&in.ReplicaSetController, &out.ReplicaSetController, s); err != nil {
		return err
	}
	if err := replicationconfigv1alpha1.Convert_v1alpha1_ReplicationControllerConfiguration_To_config_ReplicationControllerConfiguration(&in.ReplicationController, &out.ReplicationController, s); err != nil {
		return err
	}
	if err := resourcequotaconfigv1alpha1.Convert_v1alpha1_ResourceQuotaControllerConfiguration_To_config_ResourceQuotaControllerConfiguration(&in.ResourceQuotaController, &out.ResourceQuotaController, s); err != nil {
		return err
	}
	if err := serviceaccountconfigv1alpha1.Convert_v1alpha1_SAControllerConfiguration_To_config_SAControllerConfiguration(&in.SAController, &out.SAController, s); err != nil {
		return err
	}
	if err := serviceconfigv1alpha1.Convert_v1alpha1_ServiceControllerConfiguration_To_config_ServiceControllerConfiguration(&in.ServiceController, &out.ServiceController, s); err != nil {
		return err
	}
	if err := ttlafterfinishedconfigv1alpha1.Convert_v1alpha1_TTLAfterFinishedControllerConfiguration_To_config_TTLAfterFinishedControllerConfiguration(&in.TTLAfterFinishedController, &out.TTLAfterFinishedController, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_KubeControllerManagerConfiguration_To_config_KubeControllerManagerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeControllerManagerConfiguration_To_config_KubeControllerManagerConfiguration(in *v1alpha1.KubeControllerManagerConfiguration, out *config.KubeControllerManagerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeControllerManagerConfiguration_To_config_KubeControllerManagerConfiguration(in, out, s)
}

func autoConvert_config_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration(in *config.KubeControllerManagerConfiguration, out *v1alpha1.KubeControllerManagerConfiguration, s conversion.Scope) error {
	if err := Convert_config_GenericControllerManagerConfiguration_To_v1alpha1_GenericControllerManagerConfiguration(&in.Generic, &out.Generic, s); err != nil {
		return err
	}
	if err := Convert_config_KubeCloudSharedConfiguration_To_v1alpha1_KubeCloudSharedConfiguration(&in.KubeCloudShared, &out.KubeCloudShared, s); err != nil {
		return err
	}
	if err := attachdetachconfigv1alpha1.Convert_config_AttachDetachControllerConfiguration_To_v1alpha1_AttachDetachControllerConfiguration(&in.AttachDetachController, &out.AttachDetachController, s); err != nil {
		return err
	}
	if err := signerconfigv1alpha1.Convert_config_CSRSigningControllerConfiguration_To_v1alpha1_CSRSigningControllerConfiguration(&in.CSRSigningController, &out.CSRSigningController, s); err != nil {
		return err
	}
	if err := daemonconfigv1alpha1.Convert_config_DaemonSetControllerConfiguration_To_v1alpha1_DaemonSetControllerConfiguration(&in.DaemonSetController, &out.DaemonSetController, s); err != nil {
		return err
	}
	if err := deploymentconfigv1alpha1.Convert_config_DeploymentControllerConfiguration_To_v1alpha1_DeploymentControllerConfiguration(&in.DeploymentController, &out.DeploymentController, s); err != nil {
		return err
	}
	if err := statefulsetconfigv1alpha1.Convert_config_StatefulSetControllerConfiguration_To_v1alpha1_StatefulSetControllerConfiguration(&in.StatefulSetController, &out.StatefulSetController, s); err != nil {
		return err
	}
	if err := Convert_config_DeprecatedControllerConfiguration_To_v1alpha1_DeprecatedControllerConfiguration(&in.DeprecatedController, &out.DeprecatedController, s); err != nil {
		return err
	}
	if err := endpointconfigv1alpha1.Convert_config_EndpointControllerConfiguration_To_v1alpha1_EndpointControllerConfiguration(&in.EndpointController, &out.EndpointController, s); err != nil {
		return err
	}
	if err := endpointsliceconfigv1alpha1.Convert_config_EndpointSliceControllerConfiguration_To_v1alpha1_EndpointSliceControllerConfiguration(&in.EndpointSliceController, &out.EndpointSliceController, s); err != nil {
		return err
	}
	if err := endpointslicemirroringconfigv1alpha1.Convert_config_EndpointSliceMirroringControllerConfiguration_To_v1alpha1_EndpointSliceMirroringControllerConfiguration(&in.EndpointSliceMirroringController, &out.EndpointSliceMirroringController, s); err != nil {
		return err
	}
	if err := garbagecollectorconfigv1alpha1.Convert_config_GarbageCollectorControllerConfiguration_To_v1alpha1_GarbageCollectorControllerConfiguration(&in.GarbageCollectorController, &out.GarbageCollectorController, s); err != nil {
		return err
	}
	if err := podautoscalerconfigv1alpha1.Convert_config_HPAControllerConfiguration_To_v1alpha1_HPAControllerConfiguration(&in.HPAController, &out.HPAController, s); err != nil {
		return err
	}
	if err := jobconfigv1alpha1.Convert_config_JobControllerConfiguration_To_v1alpha1_JobControllerConfiguration(&in.JobController, &out.JobController, s); err != nil {
		return err
	}
	if err := namespaceconfigv1alpha1.Convert_config_NamespaceControllerConfiguration_To_v1alpha1_NamespaceControllerConfiguration(&in.NamespaceController, &out.NamespaceController, s); err != nil {
		return err
	}
	if err := nodeipamconfigv1alpha1.Convert_config_NodeIPAMControllerConfiguration_To_v1alpha1_NodeIPAMControllerConfiguration(&in.NodeIPAMController, &out.NodeIPAMController, s); err != nil {
		return err
	}
	if err := nodelifecycleconfigv1alpha1.Convert_config_NodeLifecycleControllerConfiguration_To_v1alpha1_NodeLifecycleControllerConfiguration(&in.NodeLifecycleController, &out.NodeLifecycleController, s); err != nil {
		return err
	}
	if err := persistentvolumeconfigv1alpha1.Convert_config_PersistentVolumeBinderControllerConfiguration_To_v1alpha1_PersistentVolumeBinderControllerConfiguration(&in.PersistentVolumeBinderController, &out.PersistentVolumeBinderController, s); err != nil {
		return err
	}
	if err := podgcconfigv1alpha1.Convert_config_PodGCControllerConfiguration_To_v1alpha1_PodGCControllerConfiguration(&in.PodGCController, &out.PodGCController, s); err != nil {
		return err
	}
	if err := replicasetconfigv1alpha1.Convert_config_ReplicaSetControllerConfiguration_To_v1alpha1_ReplicaSetControllerConfiguration(&in.ReplicaSetController, &out.ReplicaSetController, s); err != nil {
		return err
	}
	if err := replicationconfigv1alpha1.Convert_config_ReplicationControllerConfiguration_To_v1alpha1_ReplicationControllerConfiguration(&in.ReplicationController, &out.ReplicationController, s); err != nil {
		return err
	}
	if err := resourcequotaconfigv1alpha1.Convert_config_ResourceQuotaControllerConfiguration_To_v1alpha1_ResourceQuotaControllerConfiguration(&in.ResourceQuotaController, &out.ResourceQuotaController, s); err != nil {
		return err
	}
	if err := serviceaccountconfigv1alpha1.Convert_config_SAControllerConfiguration_To_v1alpha1_SAControllerConfiguration(&in.SAController, &out.SAController, s); err != nil {
		return err
	}
	if err := serviceconfigv1alpha1.Convert_config_ServiceControllerConfiguration_To_v1alpha1_ServiceControllerConfiguration(&in.ServiceController, &out.ServiceController, s); err != nil {
		return err
	}
	if err := ttlafterfinishedconfigv1alpha1.Convert_config_TTLAfterFinishedControllerConfiguration_To_v1alpha1_TTLAfterFinishedControllerConfiguration(&in.TTLAfterFinishedController, &out.TTLAfterFinishedController, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration is an autogenerated conversion function.
func Convert_config_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration(in *config.KubeControllerManagerConfiguration, out *v1alpha1.KubeControllerManagerConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration(in, out, s)
}
