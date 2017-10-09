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

package v1

import (
	"fmt"
	"strconv"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/kubernetes/pkg/api"
	k8s_api_v1 "k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/apis/apps"
	"k8s.io/kubernetes/pkg/apis/extensions"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	// Add non-generated conversion functions to handle the *int32 -> int32
	// conversion. A pointer is useful in the versioned type so we can default
	// it, but a plain int32 is more convenient in the internal type. These
	// functions are the same as the autogenerated ones in every other way.
	err := scheme.AddConversionFuncs(
		Convert_v1_StatefulSetSpec_To_apps_StatefulSetSpec,
		Convert_apps_StatefulSetSpec_To_v1_StatefulSetSpec,
		Convert_v1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy,
		Convert_apps_StatefulSetUpdateStrategy_To_v1_StatefulSetUpdateStrategy,
		Convert_extensions_RollingUpdateDaemonSet_To_v1_RollingUpdateDaemonSet,
		Convert_v1_RollingUpdateDaemonSet_To_extensions_RollingUpdateDaemonSet,
		Convert_v1_StatefulSetStatus_To_apps_StatefulSetStatus,
		Convert_apps_StatefulSetStatus_To_v1_StatefulSetStatus,
		Convert_v1_Deployment_To_extensions_Deployment,
		Convert_extensions_Deployment_To_v1_Deployment,
		Convert_extensions_DaemonSet_To_v1_DaemonSet,
		Convert_v1_DaemonSet_To_extensions_DaemonSet,
		Convert_extensions_DaemonSetSpec_To_v1_DaemonSetSpec,
		Convert_v1_DaemonSetSpec_To_extensions_DaemonSetSpec,
		Convert_extensions_DaemonSetUpdateStrategy_To_v1_DaemonSetUpdateStrategy,
		Convert_v1_DaemonSetUpdateStrategy_To_extensions_DaemonSetUpdateStrategy,
		// extensions
		// TODO: below conversions should be dropped in favor of auto-generated
		// ones, see https://github.com/kubernetes/kubernetes/issues/39865
		Convert_v1_DeploymentSpec_To_extensions_DeploymentSpec,
		Convert_extensions_DeploymentSpec_To_v1_DeploymentSpec,
		Convert_v1_DeploymentStrategy_To_extensions_DeploymentStrategy,
		Convert_extensions_DeploymentStrategy_To_v1_DeploymentStrategy,
		Convert_v1_RollingUpdateDeployment_To_extensions_RollingUpdateDeployment,
		Convert_extensions_RollingUpdateDeployment_To_v1_RollingUpdateDeployment,
		Convert_extensions_ReplicaSetSpec_To_v1_ReplicaSetSpec,
		Convert_v1_ReplicaSetSpec_To_extensions_ReplicaSetSpec,
	)
	if err != nil {
		return err
	}
	return nil
}

func Convert_v1_DeploymentSpec_To_extensions_DeploymentSpec(in *appsv1.DeploymentSpec, out *extensions.DeploymentSpec, s conversion.Scope) error {
	if in.Replicas != nil {
		out.Replicas = *in.Replicas
	}
	out.Selector = in.Selector
	if err := k8s_api_v1.Convert_v1_PodTemplateSpec_To_api_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if err := Convert_v1_DeploymentStrategy_To_extensions_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
		return err
	}
	out.RevisionHistoryLimit = in.RevisionHistoryLimit
	out.MinReadySeconds = in.MinReadySeconds
	out.Paused = in.Paused
	if in.ProgressDeadlineSeconds != nil {
		out.ProgressDeadlineSeconds = new(int32)
		*out.ProgressDeadlineSeconds = *in.ProgressDeadlineSeconds
	}
	return nil
}

func Convert_extensions_DeploymentSpec_To_v1_DeploymentSpec(in *extensions.DeploymentSpec, out *appsv1.DeploymentSpec, s conversion.Scope) error {
	out.Replicas = &in.Replicas
	out.Selector = in.Selector
	if err := k8s_api_v1.Convert_api_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if err := Convert_extensions_DeploymentStrategy_To_v1_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
		return err
	}
	if in.RevisionHistoryLimit != nil {
		out.RevisionHistoryLimit = new(int32)
		*out.RevisionHistoryLimit = int32(*in.RevisionHistoryLimit)
	}
	out.MinReadySeconds = int32(in.MinReadySeconds)
	out.Paused = in.Paused
	if in.ProgressDeadlineSeconds != nil {
		out.ProgressDeadlineSeconds = new(int32)
		*out.ProgressDeadlineSeconds = *in.ProgressDeadlineSeconds
	}
	return nil
}

func Convert_extensions_DeploymentStrategy_To_v1_DeploymentStrategy(in *extensions.DeploymentStrategy, out *appsv1.DeploymentStrategy, s conversion.Scope) error {
	out.Type = appsv1.DeploymentStrategyType(in.Type)
	if in.RollingUpdate != nil {
		out.RollingUpdate = new(appsv1.RollingUpdateDeployment)
		if err := Convert_extensions_RollingUpdateDeployment_To_v1_RollingUpdateDeployment(in.RollingUpdate, out.RollingUpdate, s); err != nil {
			return err
		}
	} else {
		out.RollingUpdate = nil
	}
	return nil
}

func Convert_v1_DeploymentStrategy_To_extensions_DeploymentStrategy(in *appsv1.DeploymentStrategy, out *extensions.DeploymentStrategy, s conversion.Scope) error {
	out.Type = extensions.DeploymentStrategyType(in.Type)
	if in.RollingUpdate != nil {
		out.RollingUpdate = new(extensions.RollingUpdateDeployment)
		if err := Convert_v1_RollingUpdateDeployment_To_extensions_RollingUpdateDeployment(in.RollingUpdate, out.RollingUpdate, s); err != nil {
			return err
		}
	} else {
		out.RollingUpdate = nil
	}
	return nil
}

func Convert_v1_RollingUpdateDeployment_To_extensions_RollingUpdateDeployment(in *appsv1.RollingUpdateDeployment, out *extensions.RollingUpdateDeployment, s conversion.Scope) error {
	if err := s.Convert(in.MaxUnavailable, &out.MaxUnavailable, 0); err != nil {
		return err
	}
	if err := s.Convert(in.MaxSurge, &out.MaxSurge, 0); err != nil {
		return err
	}
	return nil
}

func Convert_extensions_RollingUpdateDeployment_To_v1_RollingUpdateDeployment(in *extensions.RollingUpdateDeployment, out *appsv1.RollingUpdateDeployment, s conversion.Scope) error {
	if out.MaxUnavailable == nil {
		out.MaxUnavailable = &intstr.IntOrString{}
	}
	if err := s.Convert(&in.MaxUnavailable, out.MaxUnavailable, 0); err != nil {
		return err
	}
	if out.MaxSurge == nil {
		out.MaxSurge = &intstr.IntOrString{}
	}
	if err := s.Convert(&in.MaxSurge, out.MaxSurge, 0); err != nil {
		return err
	}
	return nil
}

func Convert_v1_Deployment_To_extensions_Deployment(in *appsv1.Deployment, out *extensions.Deployment, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_DeploymentSpec_To_extensions_DeploymentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}

	// Copy annotation to deprecated rollbackTo field for roundtrip
	// TODO: remove this conversion after we delete extensions/v1beta1 and apps/v1beta1 Deployment
	if revision, _ := in.Annotations[appsv1.DeprecatedRollbackTo]; revision != "" {
		if revision64, err := strconv.ParseInt(revision, 10, 64); err != nil {
			return fmt.Errorf("failed to parse annotation[%s]=%s as int64: %v", appsv1.DeprecatedRollbackTo, revision, err)
		} else {
			out.Spec.RollbackTo = new(extensions.RollbackConfig)
			out.Spec.RollbackTo.Revision = revision64
		}
		delete(out.Annotations, appsv1.DeprecatedRollbackTo)
	} else {
		out.Spec.RollbackTo = nil
	}

	if err := Convert_v1_DeploymentStatus_To_extensions_DeploymentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_extensions_Deployment_To_v1_Deployment(in *extensions.Deployment, out *appsv1.Deployment, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_extensions_DeploymentSpec_To_v1_DeploymentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}

	// Copy deprecated rollbackTo field to annotation for roundtrip
	// TODO: remove this conversion after we delete extensions/v1beta1 and apps/v1beta1 Deployment
	if in.Spec.RollbackTo != nil {
		if out.Annotations == nil {
			out.Annotations = make(map[string]string)
		}
		out.Annotations[appsv1.DeprecatedRollbackTo] = strconv.FormatInt(in.Spec.RollbackTo.Revision, 10)
	} else {
		delete(out.Annotations, appsv1.DeprecatedRollbackTo)
	}

	if err := Convert_extensions_DeploymentStatus_To_v1_DeploymentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_extensions_RollingUpdateDaemonSet_To_v1_RollingUpdateDaemonSet(in *extensions.RollingUpdateDaemonSet, out *appsv1.RollingUpdateDaemonSet, s conversion.Scope) error {
	if out.MaxUnavailable == nil {
		out.MaxUnavailable = &intstr.IntOrString{}
	}
	if err := s.Convert(&in.MaxUnavailable, out.MaxUnavailable, 0); err != nil {
		return err
	}
	return nil
}

func Convert_v1_RollingUpdateDaemonSet_To_extensions_RollingUpdateDaemonSet(in *appsv1.RollingUpdateDaemonSet, out *extensions.RollingUpdateDaemonSet, s conversion.Scope) error {
	if err := s.Convert(in.MaxUnavailable, &out.MaxUnavailable, 0); err != nil {
		return err
	}
	return nil
}

func Convert_extensions_DaemonSet_To_v1_DaemonSet(in *extensions.DaemonSet, out *appsv1.DaemonSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if out.Annotations == nil {
		out.Annotations = make(map[string]string)
	}
	out.Annotations[appsv1.DeprecatedTemplateGeneration] = strconv.FormatInt(in.Spec.TemplateGeneration, 10)
	if err := Convert_extensions_DaemonSetSpec_To_v1_DaemonSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := s.Convert(&in.Status, &out.Status, 0); err != nil {
		return err
	}
	return nil
}

func Convert_extensions_DaemonSetSpec_To_v1_DaemonSetSpec(in *extensions.DaemonSetSpec, out *appsv1.DaemonSetSpec, s conversion.Scope) error {
	out.Selector = in.Selector
	if err := k8s_api_v1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if err := Convert_extensions_DaemonSetUpdateStrategy_To_v1_DaemonSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
		return err
	}
	out.MinReadySeconds = int32(in.MinReadySeconds)
	if in.RevisionHistoryLimit != nil {
		out.RevisionHistoryLimit = new(int32)
		*out.RevisionHistoryLimit = *in.RevisionHistoryLimit
	} else {
		out.RevisionHistoryLimit = nil
	}
	return nil
}

func Convert_extensions_DaemonSetUpdateStrategy_To_v1_DaemonSetUpdateStrategy(in *extensions.DaemonSetUpdateStrategy, out *appsv1.DaemonSetUpdateStrategy, s conversion.Scope) error {
	out.Type = appsv1.DaemonSetUpdateStrategyType(in.Type)
	if in.RollingUpdate != nil {
		out.RollingUpdate = &appsv1.RollingUpdateDaemonSet{}
		if err := Convert_extensions_RollingUpdateDaemonSet_To_v1_RollingUpdateDaemonSet(in.RollingUpdate, out.RollingUpdate, s); err != nil {
			return err
		}
	}
	return nil
}

func Convert_v1_DaemonSet_To_extensions_DaemonSet(in *appsv1.DaemonSet, out *extensions.DaemonSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_DaemonSetSpec_To_extensions_DaemonSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if value, ok := in.Annotations[appsv1.DeprecatedTemplateGeneration]; ok {
		if value64, err := strconv.ParseInt(value, 10, 64); err != nil {
			return err
		} else {
			out.Spec.TemplateGeneration = value64
			delete(out.Annotations, appsv1.DeprecatedTemplateGeneration)
		}
	}
	if err := s.Convert(&in.Status, &out.Status, 0); err != nil {
		return err
	}
	return nil
}

func Convert_v1_DaemonSetSpec_To_extensions_DaemonSetSpec(in *appsv1.DaemonSetSpec, out *extensions.DaemonSetSpec, s conversion.Scope) error {
	out.Selector = in.Selector
	if err := k8s_api_v1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if err := Convert_v1_DaemonSetUpdateStrategy_To_extensions_DaemonSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
		return err
	}
	if in.RevisionHistoryLimit != nil {
		out.RevisionHistoryLimit = new(int32)
		*out.RevisionHistoryLimit = *in.RevisionHistoryLimit
	} else {
		out.RevisionHistoryLimit = nil
	}
	out.MinReadySeconds = in.MinReadySeconds
	return nil
}

func Convert_v1_DaemonSetUpdateStrategy_To_extensions_DaemonSetUpdateStrategy(in *appsv1.DaemonSetUpdateStrategy, out *extensions.DaemonSetUpdateStrategy, s conversion.Scope) error {
	out.Type = extensions.DaemonSetUpdateStrategyType(in.Type)
	if in.RollingUpdate != nil {
		out.RollingUpdate = &extensions.RollingUpdateDaemonSet{}
		if err := Convert_v1_RollingUpdateDaemonSet_To_extensions_RollingUpdateDaemonSet(in.RollingUpdate, out.RollingUpdate, s); err != nil {
			return err
		}
	}
	return nil
}

func Convert_extensions_ReplicaSetSpec_To_v1_ReplicaSetSpec(in *extensions.ReplicaSetSpec, out *appsv1.ReplicaSetSpec, s conversion.Scope) error {
	out.Replicas = new(int32)
	*out.Replicas = int32(in.Replicas)
	out.MinReadySeconds = in.MinReadySeconds
	out.Selector = in.Selector
	if err := k8s_api_v1.Convert_api_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_ReplicaSetSpec_To_extensions_ReplicaSetSpec(in *appsv1.ReplicaSetSpec, out *extensions.ReplicaSetSpec, s conversion.Scope) error {
	if in.Replicas != nil {
		out.Replicas = *in.Replicas
	}
	out.MinReadySeconds = in.MinReadySeconds
	out.Selector = in.Selector
	if err := k8s_api_v1.Convert_v1_PodTemplateSpec_To_api_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_StatefulSetSpec_To_apps_StatefulSetSpec(in *appsv1.StatefulSetSpec, out *apps.StatefulSetSpec, s conversion.Scope) error {
	if in.Replicas != nil {
		out.Replicas = *in.Replicas
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		if err := s.Convert(*in, *out, 0); err != nil {
			return err
		}
	} else {
		out.Selector = nil
	}
	if err := k8s_api_v1.Convert_v1_PodTemplateSpec_To_api_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]api.PersistentVolumeClaim, len(*in))
		for i := range *in {
			if err := s.Convert(&(*in)[i], &(*out)[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.VolumeClaimTemplates = nil
	}
	if err := Convert_v1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
		return err
	}
	if in.RevisionHistoryLimit != nil {
		out.RevisionHistoryLimit = new(int32)
		*out.RevisionHistoryLimit = *in.RevisionHistoryLimit
	} else {
		out.RevisionHistoryLimit = nil
	}
	out.ServiceName = in.ServiceName
	out.PodManagementPolicy = apps.PodManagementPolicyType(in.PodManagementPolicy)
	return nil
}

func Convert_apps_StatefulSetSpec_To_v1_StatefulSetSpec(in *apps.StatefulSetSpec, out *appsv1.StatefulSetSpec, s conversion.Scope) error {
	out.Replicas = new(int32)
	*out.Replicas = in.Replicas
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		if err := s.Convert(*in, *out, 0); err != nil {
			return err
		}
	} else {
		out.Selector = nil
	}
	if err := k8s_api_v1.Convert_api_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]v1.PersistentVolumeClaim, len(*in))
		for i := range *in {
			if err := s.Convert(&(*in)[i], &(*out)[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.VolumeClaimTemplates = nil
	}
	if in.RevisionHistoryLimit != nil {
		out.RevisionHistoryLimit = new(int32)
		*out.RevisionHistoryLimit = *in.RevisionHistoryLimit
	} else {
		out.RevisionHistoryLimit = nil
	}
	out.ServiceName = in.ServiceName
	out.PodManagementPolicy = appsv1.PodManagementPolicyType(in.PodManagementPolicy)
	if err := Convert_apps_StatefulSetUpdateStrategy_To_v1_StatefulSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(in *appsv1.StatefulSetUpdateStrategy, out *apps.StatefulSetUpdateStrategy, s conversion.Scope) error {
	out.Type = apps.StatefulSetUpdateStrategyType(in.Type)
	if in.RollingUpdate != nil {
		out.RollingUpdate = new(apps.RollingUpdateStatefulSetStrategy)
		out.RollingUpdate.Partition = *in.RollingUpdate.Partition
	} else {
		out.RollingUpdate = nil
	}
	return nil
}

func Convert_apps_StatefulSetUpdateStrategy_To_v1_StatefulSetUpdateStrategy(in *apps.StatefulSetUpdateStrategy, out *appsv1.StatefulSetUpdateStrategy, s conversion.Scope) error {
	out.Type = appsv1.StatefulSetUpdateStrategyType(in.Type)
	if in.RollingUpdate != nil {
		out.RollingUpdate = new(appsv1.RollingUpdateStatefulSetStrategy)
		out.RollingUpdate.Partition = new(int32)
		*out.RollingUpdate.Partition = in.RollingUpdate.Partition
	} else {
		out.RollingUpdate = nil
	}
	return nil
}

func Convert_v1_StatefulSetStatus_To_apps_StatefulSetStatus(in *appsv1.StatefulSetStatus, out *apps.StatefulSetStatus, s conversion.Scope) error {
	out.ObservedGeneration = new(int64)
	*out.ObservedGeneration = in.ObservedGeneration
	out.Replicas = in.Replicas
	out.ReadyReplicas = in.ReadyReplicas
	out.CurrentReplicas = in.CurrentReplicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.CurrentRevision = in.CurrentRevision
	out.UpdateRevision = in.UpdateRevision
	if in.CollisionCount != nil {
		out.CollisionCount = new(int32)
		*out.CollisionCount = *in.CollisionCount
	}
	return nil
}

func Convert_apps_StatefulSetStatus_To_v1_StatefulSetStatus(in *apps.StatefulSetStatus, out *appsv1.StatefulSetStatus, s conversion.Scope) error {
	if in.ObservedGeneration != nil {
		out.ObservedGeneration = *in.ObservedGeneration
	}
	out.Replicas = in.Replicas
	out.ReadyReplicas = in.ReadyReplicas
	out.CurrentReplicas = in.CurrentReplicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.CurrentRevision = in.CurrentRevision
	out.UpdateRevision = in.UpdateRevision
	if in.CollisionCount != nil {
		out.CollisionCount = new(int32)
		*out.CollisionCount = *in.CollisionCount
	}
	return nil
}
