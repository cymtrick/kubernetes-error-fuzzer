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

package v1beta1

import (
	unsafe "unsafe"

	v1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/storage/v1beta1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
	storage "k8s.io/kubernetes/pkg/apis/storage"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1beta1.CSIDriver)(nil), (*storage.CSIDriver)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CSIDriver_To_storage_CSIDriver(a.(*v1beta1.CSIDriver), b.(*storage.CSIDriver), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.CSIDriver)(nil), (*v1beta1.CSIDriver)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_CSIDriver_To_v1beta1_CSIDriver(a.(*storage.CSIDriver), b.(*v1beta1.CSIDriver), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.CSIDriverList)(nil), (*storage.CSIDriverList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CSIDriverList_To_storage_CSIDriverList(a.(*v1beta1.CSIDriverList), b.(*storage.CSIDriverList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.CSIDriverList)(nil), (*v1beta1.CSIDriverList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_CSIDriverList_To_v1beta1_CSIDriverList(a.(*storage.CSIDriverList), b.(*v1beta1.CSIDriverList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.CSIDriverSpec)(nil), (*storage.CSIDriverSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CSIDriverSpec_To_storage_CSIDriverSpec(a.(*v1beta1.CSIDriverSpec), b.(*storage.CSIDriverSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.CSIDriverSpec)(nil), (*v1beta1.CSIDriverSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_CSIDriverSpec_To_v1beta1_CSIDriverSpec(a.(*storage.CSIDriverSpec), b.(*v1beta1.CSIDriverSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.CSINode)(nil), (*storage.CSINode)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CSINode_To_storage_CSINode(a.(*v1beta1.CSINode), b.(*storage.CSINode), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.CSINode)(nil), (*v1beta1.CSINode)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_CSINode_To_v1beta1_CSINode(a.(*storage.CSINode), b.(*v1beta1.CSINode), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.CSINodeDriver)(nil), (*storage.CSINodeDriver)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CSINodeDriver_To_storage_CSINodeDriver(a.(*v1beta1.CSINodeDriver), b.(*storage.CSINodeDriver), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.CSINodeDriver)(nil), (*v1beta1.CSINodeDriver)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_CSINodeDriver_To_v1beta1_CSINodeDriver(a.(*storage.CSINodeDriver), b.(*v1beta1.CSINodeDriver), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.CSINodeList)(nil), (*storage.CSINodeList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CSINodeList_To_storage_CSINodeList(a.(*v1beta1.CSINodeList), b.(*storage.CSINodeList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.CSINodeList)(nil), (*v1beta1.CSINodeList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_CSINodeList_To_v1beta1_CSINodeList(a.(*storage.CSINodeList), b.(*v1beta1.CSINodeList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.CSINodeSpec)(nil), (*storage.CSINodeSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CSINodeSpec_To_storage_CSINodeSpec(a.(*v1beta1.CSINodeSpec), b.(*storage.CSINodeSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.CSINodeSpec)(nil), (*v1beta1.CSINodeSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_CSINodeSpec_To_v1beta1_CSINodeSpec(a.(*storage.CSINodeSpec), b.(*v1beta1.CSINodeSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.StorageClass)(nil), (*storage.StorageClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_StorageClass_To_storage_StorageClass(a.(*v1beta1.StorageClass), b.(*storage.StorageClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.StorageClass)(nil), (*v1beta1.StorageClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_StorageClass_To_v1beta1_StorageClass(a.(*storage.StorageClass), b.(*v1beta1.StorageClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.StorageClassList)(nil), (*storage.StorageClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_StorageClassList_To_storage_StorageClassList(a.(*v1beta1.StorageClassList), b.(*storage.StorageClassList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.StorageClassList)(nil), (*v1beta1.StorageClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_StorageClassList_To_v1beta1_StorageClassList(a.(*storage.StorageClassList), b.(*v1beta1.StorageClassList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VolumeAttachment)(nil), (*storage.VolumeAttachment)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VolumeAttachment_To_storage_VolumeAttachment(a.(*v1beta1.VolumeAttachment), b.(*storage.VolumeAttachment), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeAttachment)(nil), (*v1beta1.VolumeAttachment)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeAttachment_To_v1beta1_VolumeAttachment(a.(*storage.VolumeAttachment), b.(*v1beta1.VolumeAttachment), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VolumeAttachmentList)(nil), (*storage.VolumeAttachmentList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VolumeAttachmentList_To_storage_VolumeAttachmentList(a.(*v1beta1.VolumeAttachmentList), b.(*storage.VolumeAttachmentList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeAttachmentList)(nil), (*v1beta1.VolumeAttachmentList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeAttachmentList_To_v1beta1_VolumeAttachmentList(a.(*storage.VolumeAttachmentList), b.(*v1beta1.VolumeAttachmentList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VolumeAttachmentSource)(nil), (*storage.VolumeAttachmentSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(a.(*v1beta1.VolumeAttachmentSource), b.(*storage.VolumeAttachmentSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeAttachmentSource)(nil), (*v1beta1.VolumeAttachmentSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeAttachmentSource_To_v1beta1_VolumeAttachmentSource(a.(*storage.VolumeAttachmentSource), b.(*v1beta1.VolumeAttachmentSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VolumeAttachmentSpec)(nil), (*storage.VolumeAttachmentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(a.(*v1beta1.VolumeAttachmentSpec), b.(*storage.VolumeAttachmentSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeAttachmentSpec)(nil), (*v1beta1.VolumeAttachmentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeAttachmentSpec_To_v1beta1_VolumeAttachmentSpec(a.(*storage.VolumeAttachmentSpec), b.(*v1beta1.VolumeAttachmentSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VolumeAttachmentStatus)(nil), (*storage.VolumeAttachmentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(a.(*v1beta1.VolumeAttachmentStatus), b.(*storage.VolumeAttachmentStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeAttachmentStatus)(nil), (*v1beta1.VolumeAttachmentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeAttachmentStatus_To_v1beta1_VolumeAttachmentStatus(a.(*storage.VolumeAttachmentStatus), b.(*v1beta1.VolumeAttachmentStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VolumeError)(nil), (*storage.VolumeError)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VolumeError_To_storage_VolumeError(a.(*v1beta1.VolumeError), b.(*storage.VolumeError), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeError)(nil), (*v1beta1.VolumeError)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeError_To_v1beta1_VolumeError(a.(*storage.VolumeError), b.(*v1beta1.VolumeError), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VolumeNodeResources)(nil), (*storage.VolumeNodeResources)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VolumeNodeResources_To_storage_VolumeNodeResources(a.(*v1beta1.VolumeNodeResources), b.(*storage.VolumeNodeResources), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeNodeResources)(nil), (*v1beta1.VolumeNodeResources)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeNodeResources_To_v1beta1_VolumeNodeResources(a.(*storage.VolumeNodeResources), b.(*v1beta1.VolumeNodeResources), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_CSIDriver_To_storage_CSIDriver(in *v1beta1.CSIDriver, out *storage.CSIDriver, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_CSIDriverSpec_To_storage_CSIDriverSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CSIDriver_To_storage_CSIDriver is an autogenerated conversion function.
func Convert_v1beta1_CSIDriver_To_storage_CSIDriver(in *v1beta1.CSIDriver, out *storage.CSIDriver, s conversion.Scope) error {
	return autoConvert_v1beta1_CSIDriver_To_storage_CSIDriver(in, out, s)
}

func autoConvert_storage_CSIDriver_To_v1beta1_CSIDriver(in *storage.CSIDriver, out *v1beta1.CSIDriver, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_storage_CSIDriverSpec_To_v1beta1_CSIDriverSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_storage_CSIDriver_To_v1beta1_CSIDriver is an autogenerated conversion function.
func Convert_storage_CSIDriver_To_v1beta1_CSIDriver(in *storage.CSIDriver, out *v1beta1.CSIDriver, s conversion.Scope) error {
	return autoConvert_storage_CSIDriver_To_v1beta1_CSIDriver(in, out, s)
}

func autoConvert_v1beta1_CSIDriverList_To_storage_CSIDriverList(in *v1beta1.CSIDriverList, out *storage.CSIDriverList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]storage.CSIDriver)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_CSIDriverList_To_storage_CSIDriverList is an autogenerated conversion function.
func Convert_v1beta1_CSIDriverList_To_storage_CSIDriverList(in *v1beta1.CSIDriverList, out *storage.CSIDriverList, s conversion.Scope) error {
	return autoConvert_v1beta1_CSIDriverList_To_storage_CSIDriverList(in, out, s)
}

func autoConvert_storage_CSIDriverList_To_v1beta1_CSIDriverList(in *storage.CSIDriverList, out *v1beta1.CSIDriverList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1beta1.CSIDriver)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_storage_CSIDriverList_To_v1beta1_CSIDriverList is an autogenerated conversion function.
func Convert_storage_CSIDriverList_To_v1beta1_CSIDriverList(in *storage.CSIDriverList, out *v1beta1.CSIDriverList, s conversion.Scope) error {
	return autoConvert_storage_CSIDriverList_To_v1beta1_CSIDriverList(in, out, s)
}

func autoConvert_v1beta1_CSIDriverSpec_To_storage_CSIDriverSpec(in *v1beta1.CSIDriverSpec, out *storage.CSIDriverSpec, s conversion.Scope) error {
	out.AttachRequired = (*bool)(unsafe.Pointer(in.AttachRequired))
	out.PodInfoOnMount = (*bool)(unsafe.Pointer(in.PodInfoOnMount))
	out.VolumeLifecycleModes = *(*[]storage.VolumeLifecycleMode)(unsafe.Pointer(&in.VolumeLifecycleModes))
	out.StorageCapacity = (*bool)(unsafe.Pointer(in.StorageCapacity))
	return nil
}

// Convert_v1beta1_CSIDriverSpec_To_storage_CSIDriverSpec is an autogenerated conversion function.
func Convert_v1beta1_CSIDriverSpec_To_storage_CSIDriverSpec(in *v1beta1.CSIDriverSpec, out *storage.CSIDriverSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_CSIDriverSpec_To_storage_CSIDriverSpec(in, out, s)
}

func autoConvert_storage_CSIDriverSpec_To_v1beta1_CSIDriverSpec(in *storage.CSIDriverSpec, out *v1beta1.CSIDriverSpec, s conversion.Scope) error {
	out.AttachRequired = (*bool)(unsafe.Pointer(in.AttachRequired))
	out.PodInfoOnMount = (*bool)(unsafe.Pointer(in.PodInfoOnMount))
	out.VolumeLifecycleModes = *(*[]v1beta1.VolumeLifecycleMode)(unsafe.Pointer(&in.VolumeLifecycleModes))
	out.StorageCapacity = (*bool)(unsafe.Pointer(in.StorageCapacity))
	return nil
}

// Convert_storage_CSIDriverSpec_To_v1beta1_CSIDriverSpec is an autogenerated conversion function.
func Convert_storage_CSIDriverSpec_To_v1beta1_CSIDriverSpec(in *storage.CSIDriverSpec, out *v1beta1.CSIDriverSpec, s conversion.Scope) error {
	return autoConvert_storage_CSIDriverSpec_To_v1beta1_CSIDriverSpec(in, out, s)
}

func autoConvert_v1beta1_CSINode_To_storage_CSINode(in *v1beta1.CSINode, out *storage.CSINode, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_CSINodeSpec_To_storage_CSINodeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CSINode_To_storage_CSINode is an autogenerated conversion function.
func Convert_v1beta1_CSINode_To_storage_CSINode(in *v1beta1.CSINode, out *storage.CSINode, s conversion.Scope) error {
	return autoConvert_v1beta1_CSINode_To_storage_CSINode(in, out, s)
}

func autoConvert_storage_CSINode_To_v1beta1_CSINode(in *storage.CSINode, out *v1beta1.CSINode, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_storage_CSINodeSpec_To_v1beta1_CSINodeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_storage_CSINode_To_v1beta1_CSINode is an autogenerated conversion function.
func Convert_storage_CSINode_To_v1beta1_CSINode(in *storage.CSINode, out *v1beta1.CSINode, s conversion.Scope) error {
	return autoConvert_storage_CSINode_To_v1beta1_CSINode(in, out, s)
}

func autoConvert_v1beta1_CSINodeDriver_To_storage_CSINodeDriver(in *v1beta1.CSINodeDriver, out *storage.CSINodeDriver, s conversion.Scope) error {
	out.Name = in.Name
	out.NodeID = in.NodeID
	out.TopologyKeys = *(*[]string)(unsafe.Pointer(&in.TopologyKeys))
	out.Allocatable = (*storage.VolumeNodeResources)(unsafe.Pointer(in.Allocatable))
	return nil
}

// Convert_v1beta1_CSINodeDriver_To_storage_CSINodeDriver is an autogenerated conversion function.
func Convert_v1beta1_CSINodeDriver_To_storage_CSINodeDriver(in *v1beta1.CSINodeDriver, out *storage.CSINodeDriver, s conversion.Scope) error {
	return autoConvert_v1beta1_CSINodeDriver_To_storage_CSINodeDriver(in, out, s)
}

func autoConvert_storage_CSINodeDriver_To_v1beta1_CSINodeDriver(in *storage.CSINodeDriver, out *v1beta1.CSINodeDriver, s conversion.Scope) error {
	out.Name = in.Name
	out.NodeID = in.NodeID
	out.TopologyKeys = *(*[]string)(unsafe.Pointer(&in.TopologyKeys))
	out.Allocatable = (*v1beta1.VolumeNodeResources)(unsafe.Pointer(in.Allocatable))
	return nil
}

// Convert_storage_CSINodeDriver_To_v1beta1_CSINodeDriver is an autogenerated conversion function.
func Convert_storage_CSINodeDriver_To_v1beta1_CSINodeDriver(in *storage.CSINodeDriver, out *v1beta1.CSINodeDriver, s conversion.Scope) error {
	return autoConvert_storage_CSINodeDriver_To_v1beta1_CSINodeDriver(in, out, s)
}

func autoConvert_v1beta1_CSINodeList_To_storage_CSINodeList(in *v1beta1.CSINodeList, out *storage.CSINodeList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]storage.CSINode)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_CSINodeList_To_storage_CSINodeList is an autogenerated conversion function.
func Convert_v1beta1_CSINodeList_To_storage_CSINodeList(in *v1beta1.CSINodeList, out *storage.CSINodeList, s conversion.Scope) error {
	return autoConvert_v1beta1_CSINodeList_To_storage_CSINodeList(in, out, s)
}

func autoConvert_storage_CSINodeList_To_v1beta1_CSINodeList(in *storage.CSINodeList, out *v1beta1.CSINodeList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1beta1.CSINode)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_storage_CSINodeList_To_v1beta1_CSINodeList is an autogenerated conversion function.
func Convert_storage_CSINodeList_To_v1beta1_CSINodeList(in *storage.CSINodeList, out *v1beta1.CSINodeList, s conversion.Scope) error {
	return autoConvert_storage_CSINodeList_To_v1beta1_CSINodeList(in, out, s)
}

func autoConvert_v1beta1_CSINodeSpec_To_storage_CSINodeSpec(in *v1beta1.CSINodeSpec, out *storage.CSINodeSpec, s conversion.Scope) error {
	out.Drivers = *(*[]storage.CSINodeDriver)(unsafe.Pointer(&in.Drivers))
	return nil
}

// Convert_v1beta1_CSINodeSpec_To_storage_CSINodeSpec is an autogenerated conversion function.
func Convert_v1beta1_CSINodeSpec_To_storage_CSINodeSpec(in *v1beta1.CSINodeSpec, out *storage.CSINodeSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_CSINodeSpec_To_storage_CSINodeSpec(in, out, s)
}

func autoConvert_storage_CSINodeSpec_To_v1beta1_CSINodeSpec(in *storage.CSINodeSpec, out *v1beta1.CSINodeSpec, s conversion.Scope) error {
	out.Drivers = *(*[]v1beta1.CSINodeDriver)(unsafe.Pointer(&in.Drivers))
	return nil
}

// Convert_storage_CSINodeSpec_To_v1beta1_CSINodeSpec is an autogenerated conversion function.
func Convert_storage_CSINodeSpec_To_v1beta1_CSINodeSpec(in *storage.CSINodeSpec, out *v1beta1.CSINodeSpec, s conversion.Scope) error {
	return autoConvert_storage_CSINodeSpec_To_v1beta1_CSINodeSpec(in, out, s)
}

func autoConvert_v1beta1_StorageClass_To_storage_StorageClass(in *v1beta1.StorageClass, out *storage.StorageClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Provisioner = in.Provisioner
	out.Parameters = *(*map[string]string)(unsafe.Pointer(&in.Parameters))
	out.ReclaimPolicy = (*core.PersistentVolumeReclaimPolicy)(unsafe.Pointer(in.ReclaimPolicy))
	out.MountOptions = *(*[]string)(unsafe.Pointer(&in.MountOptions))
	out.AllowVolumeExpansion = (*bool)(unsafe.Pointer(in.AllowVolumeExpansion))
	out.VolumeBindingMode = (*storage.VolumeBindingMode)(unsafe.Pointer(in.VolumeBindingMode))
	out.AllowedTopologies = *(*[]core.TopologySelectorTerm)(unsafe.Pointer(&in.AllowedTopologies))
	return nil
}

// Convert_v1beta1_StorageClass_To_storage_StorageClass is an autogenerated conversion function.
func Convert_v1beta1_StorageClass_To_storage_StorageClass(in *v1beta1.StorageClass, out *storage.StorageClass, s conversion.Scope) error {
	return autoConvert_v1beta1_StorageClass_To_storage_StorageClass(in, out, s)
}

func autoConvert_storage_StorageClass_To_v1beta1_StorageClass(in *storage.StorageClass, out *v1beta1.StorageClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Provisioner = in.Provisioner
	out.Parameters = *(*map[string]string)(unsafe.Pointer(&in.Parameters))
	out.ReclaimPolicy = (*v1.PersistentVolumeReclaimPolicy)(unsafe.Pointer(in.ReclaimPolicy))
	out.MountOptions = *(*[]string)(unsafe.Pointer(&in.MountOptions))
	out.AllowVolumeExpansion = (*bool)(unsafe.Pointer(in.AllowVolumeExpansion))
	out.VolumeBindingMode = (*v1beta1.VolumeBindingMode)(unsafe.Pointer(in.VolumeBindingMode))
	out.AllowedTopologies = *(*[]v1.TopologySelectorTerm)(unsafe.Pointer(&in.AllowedTopologies))
	return nil
}

// Convert_storage_StorageClass_To_v1beta1_StorageClass is an autogenerated conversion function.
func Convert_storage_StorageClass_To_v1beta1_StorageClass(in *storage.StorageClass, out *v1beta1.StorageClass, s conversion.Scope) error {
	return autoConvert_storage_StorageClass_To_v1beta1_StorageClass(in, out, s)
}

func autoConvert_v1beta1_StorageClassList_To_storage_StorageClassList(in *v1beta1.StorageClassList, out *storage.StorageClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]storage.StorageClass)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_StorageClassList_To_storage_StorageClassList is an autogenerated conversion function.
func Convert_v1beta1_StorageClassList_To_storage_StorageClassList(in *v1beta1.StorageClassList, out *storage.StorageClassList, s conversion.Scope) error {
	return autoConvert_v1beta1_StorageClassList_To_storage_StorageClassList(in, out, s)
}

func autoConvert_storage_StorageClassList_To_v1beta1_StorageClassList(in *storage.StorageClassList, out *v1beta1.StorageClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1beta1.StorageClass)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_storage_StorageClassList_To_v1beta1_StorageClassList is an autogenerated conversion function.
func Convert_storage_StorageClassList_To_v1beta1_StorageClassList(in *storage.StorageClassList, out *v1beta1.StorageClassList, s conversion.Scope) error {
	return autoConvert_storage_StorageClassList_To_v1beta1_StorageClassList(in, out, s)
}

func autoConvert_v1beta1_VolumeAttachment_To_storage_VolumeAttachment(in *v1beta1.VolumeAttachment, out *storage.VolumeAttachment, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VolumeAttachment_To_storage_VolumeAttachment is an autogenerated conversion function.
func Convert_v1beta1_VolumeAttachment_To_storage_VolumeAttachment(in *v1beta1.VolumeAttachment, out *storage.VolumeAttachment, s conversion.Scope) error {
	return autoConvert_v1beta1_VolumeAttachment_To_storage_VolumeAttachment(in, out, s)
}

func autoConvert_storage_VolumeAttachment_To_v1beta1_VolumeAttachment(in *storage.VolumeAttachment, out *v1beta1.VolumeAttachment, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_storage_VolumeAttachmentSpec_To_v1beta1_VolumeAttachmentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_storage_VolumeAttachmentStatus_To_v1beta1_VolumeAttachmentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_storage_VolumeAttachment_To_v1beta1_VolumeAttachment is an autogenerated conversion function.
func Convert_storage_VolumeAttachment_To_v1beta1_VolumeAttachment(in *storage.VolumeAttachment, out *v1beta1.VolumeAttachment, s conversion.Scope) error {
	return autoConvert_storage_VolumeAttachment_To_v1beta1_VolumeAttachment(in, out, s)
}

func autoConvert_v1beta1_VolumeAttachmentList_To_storage_VolumeAttachmentList(in *v1beta1.VolumeAttachmentList, out *storage.VolumeAttachmentList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]storage.VolumeAttachment, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_VolumeAttachment_To_storage_VolumeAttachment(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_VolumeAttachmentList_To_storage_VolumeAttachmentList is an autogenerated conversion function.
func Convert_v1beta1_VolumeAttachmentList_To_storage_VolumeAttachmentList(in *v1beta1.VolumeAttachmentList, out *storage.VolumeAttachmentList, s conversion.Scope) error {
	return autoConvert_v1beta1_VolumeAttachmentList_To_storage_VolumeAttachmentList(in, out, s)
}

func autoConvert_storage_VolumeAttachmentList_To_v1beta1_VolumeAttachmentList(in *storage.VolumeAttachmentList, out *v1beta1.VolumeAttachmentList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.VolumeAttachment, len(*in))
		for i := range *in {
			if err := Convert_storage_VolumeAttachment_To_v1beta1_VolumeAttachment(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_storage_VolumeAttachmentList_To_v1beta1_VolumeAttachmentList is an autogenerated conversion function.
func Convert_storage_VolumeAttachmentList_To_v1beta1_VolumeAttachmentList(in *storage.VolumeAttachmentList, out *v1beta1.VolumeAttachmentList, s conversion.Scope) error {
	return autoConvert_storage_VolumeAttachmentList_To_v1beta1_VolumeAttachmentList(in, out, s)
}

func autoConvert_v1beta1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(in *v1beta1.VolumeAttachmentSource, out *storage.VolumeAttachmentSource, s conversion.Scope) error {
	out.PersistentVolumeName = (*string)(unsafe.Pointer(in.PersistentVolumeName))
	if in.InlineVolumeSpec != nil {
		in, out := &in.InlineVolumeSpec, &out.InlineVolumeSpec
		*out = new(core.PersistentVolumeSpec)
		if err := corev1.Convert_v1_PersistentVolumeSpec_To_core_PersistentVolumeSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.InlineVolumeSpec = nil
	}
	return nil
}

// Convert_v1beta1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource is an autogenerated conversion function.
func Convert_v1beta1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(in *v1beta1.VolumeAttachmentSource, out *storage.VolumeAttachmentSource, s conversion.Scope) error {
	return autoConvert_v1beta1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(in, out, s)
}

func autoConvert_storage_VolumeAttachmentSource_To_v1beta1_VolumeAttachmentSource(in *storage.VolumeAttachmentSource, out *v1beta1.VolumeAttachmentSource, s conversion.Scope) error {
	out.PersistentVolumeName = (*string)(unsafe.Pointer(in.PersistentVolumeName))
	if in.InlineVolumeSpec != nil {
		in, out := &in.InlineVolumeSpec, &out.InlineVolumeSpec
		*out = new(v1.PersistentVolumeSpec)
		if err := corev1.Convert_core_PersistentVolumeSpec_To_v1_PersistentVolumeSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.InlineVolumeSpec = nil
	}
	return nil
}

// Convert_storage_VolumeAttachmentSource_To_v1beta1_VolumeAttachmentSource is an autogenerated conversion function.
func Convert_storage_VolumeAttachmentSource_To_v1beta1_VolumeAttachmentSource(in *storage.VolumeAttachmentSource, out *v1beta1.VolumeAttachmentSource, s conversion.Scope) error {
	return autoConvert_storage_VolumeAttachmentSource_To_v1beta1_VolumeAttachmentSource(in, out, s)
}

func autoConvert_v1beta1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(in *v1beta1.VolumeAttachmentSpec, out *storage.VolumeAttachmentSpec, s conversion.Scope) error {
	out.Attacher = in.Attacher
	if err := Convert_v1beta1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(&in.Source, &out.Source, s); err != nil {
		return err
	}
	out.NodeName = in.NodeName
	return nil
}

// Convert_v1beta1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec is an autogenerated conversion function.
func Convert_v1beta1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(in *v1beta1.VolumeAttachmentSpec, out *storage.VolumeAttachmentSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(in, out, s)
}

func autoConvert_storage_VolumeAttachmentSpec_To_v1beta1_VolumeAttachmentSpec(in *storage.VolumeAttachmentSpec, out *v1beta1.VolumeAttachmentSpec, s conversion.Scope) error {
	out.Attacher = in.Attacher
	if err := Convert_storage_VolumeAttachmentSource_To_v1beta1_VolumeAttachmentSource(&in.Source, &out.Source, s); err != nil {
		return err
	}
	out.NodeName = in.NodeName
	return nil
}

// Convert_storage_VolumeAttachmentSpec_To_v1beta1_VolumeAttachmentSpec is an autogenerated conversion function.
func Convert_storage_VolumeAttachmentSpec_To_v1beta1_VolumeAttachmentSpec(in *storage.VolumeAttachmentSpec, out *v1beta1.VolumeAttachmentSpec, s conversion.Scope) error {
	return autoConvert_storage_VolumeAttachmentSpec_To_v1beta1_VolumeAttachmentSpec(in, out, s)
}

func autoConvert_v1beta1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(in *v1beta1.VolumeAttachmentStatus, out *storage.VolumeAttachmentStatus, s conversion.Scope) error {
	out.Attached = in.Attached
	out.AttachmentMetadata = *(*map[string]string)(unsafe.Pointer(&in.AttachmentMetadata))
	out.AttachError = (*storage.VolumeError)(unsafe.Pointer(in.AttachError))
	out.DetachError = (*storage.VolumeError)(unsafe.Pointer(in.DetachError))
	return nil
}

// Convert_v1beta1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus is an autogenerated conversion function.
func Convert_v1beta1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(in *v1beta1.VolumeAttachmentStatus, out *storage.VolumeAttachmentStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(in, out, s)
}

func autoConvert_storage_VolumeAttachmentStatus_To_v1beta1_VolumeAttachmentStatus(in *storage.VolumeAttachmentStatus, out *v1beta1.VolumeAttachmentStatus, s conversion.Scope) error {
	out.Attached = in.Attached
	out.AttachmentMetadata = *(*map[string]string)(unsafe.Pointer(&in.AttachmentMetadata))
	out.AttachError = (*v1beta1.VolumeError)(unsafe.Pointer(in.AttachError))
	out.DetachError = (*v1beta1.VolumeError)(unsafe.Pointer(in.DetachError))
	return nil
}

// Convert_storage_VolumeAttachmentStatus_To_v1beta1_VolumeAttachmentStatus is an autogenerated conversion function.
func Convert_storage_VolumeAttachmentStatus_To_v1beta1_VolumeAttachmentStatus(in *storage.VolumeAttachmentStatus, out *v1beta1.VolumeAttachmentStatus, s conversion.Scope) error {
	return autoConvert_storage_VolumeAttachmentStatus_To_v1beta1_VolumeAttachmentStatus(in, out, s)
}

func autoConvert_v1beta1_VolumeError_To_storage_VolumeError(in *v1beta1.VolumeError, out *storage.VolumeError, s conversion.Scope) error {
	out.Time = in.Time
	out.Message = in.Message
	return nil
}

// Convert_v1beta1_VolumeError_To_storage_VolumeError is an autogenerated conversion function.
func Convert_v1beta1_VolumeError_To_storage_VolumeError(in *v1beta1.VolumeError, out *storage.VolumeError, s conversion.Scope) error {
	return autoConvert_v1beta1_VolumeError_To_storage_VolumeError(in, out, s)
}

func autoConvert_storage_VolumeError_To_v1beta1_VolumeError(in *storage.VolumeError, out *v1beta1.VolumeError, s conversion.Scope) error {
	out.Time = in.Time
	out.Message = in.Message
	return nil
}

// Convert_storage_VolumeError_To_v1beta1_VolumeError is an autogenerated conversion function.
func Convert_storage_VolumeError_To_v1beta1_VolumeError(in *storage.VolumeError, out *v1beta1.VolumeError, s conversion.Scope) error {
	return autoConvert_storage_VolumeError_To_v1beta1_VolumeError(in, out, s)
}

func autoConvert_v1beta1_VolumeNodeResources_To_storage_VolumeNodeResources(in *v1beta1.VolumeNodeResources, out *storage.VolumeNodeResources, s conversion.Scope) error {
	out.Count = (*int32)(unsafe.Pointer(in.Count))
	return nil
}

// Convert_v1beta1_VolumeNodeResources_To_storage_VolumeNodeResources is an autogenerated conversion function.
func Convert_v1beta1_VolumeNodeResources_To_storage_VolumeNodeResources(in *v1beta1.VolumeNodeResources, out *storage.VolumeNodeResources, s conversion.Scope) error {
	return autoConvert_v1beta1_VolumeNodeResources_To_storage_VolumeNodeResources(in, out, s)
}

func autoConvert_storage_VolumeNodeResources_To_v1beta1_VolumeNodeResources(in *storage.VolumeNodeResources, out *v1beta1.VolumeNodeResources, s conversion.Scope) error {
	out.Count = (*int32)(unsafe.Pointer(in.Count))
	return nil
}

// Convert_storage_VolumeNodeResources_To_v1beta1_VolumeNodeResources is an autogenerated conversion function.
func Convert_storage_VolumeNodeResources_To_v1beta1_VolumeNodeResources(in *storage.VolumeNodeResources, out *v1beta1.VolumeNodeResources, s conversion.Scope) error {
	return autoConvert_storage_VolumeNodeResources_To_v1beta1_VolumeNodeResources(in, out, s)
}
