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

// This file was automatically generated by informer-gen

package v1beta1

import (
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// StorageClasses returns a StorageClassInformer.
	StorageClasses() StorageClassInformer
	// VolumeAttachments returns a VolumeAttachmentInformer.
	VolumeAttachments() VolumeAttachmentInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// StorageClasses returns a StorageClassInformer.
func (v *version) StorageClasses() StorageClassInformer {
	return &storageClassInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// VolumeAttachments returns a VolumeAttachmentInformer.
func (v *version) VolumeAttachments() VolumeAttachmentInformer {
	return &volumeAttachmentInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
