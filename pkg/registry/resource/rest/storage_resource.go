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

package rest

import (
	resourcev1alpha1 "k8s.io/api/resource/v1alpha1"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/resource"
	podschedulingstore "k8s.io/kubernetes/pkg/registry/resource/podscheduling/storage"
	resourceclaimstore "k8s.io/kubernetes/pkg/registry/resource/resourceclaim/storage"
	resourceclaimtemplatestore "k8s.io/kubernetes/pkg/registry/resource/resourceclaimtemplate/storage"
	resourceclassstore "k8s.io/kubernetes/pkg/registry/resource/resourceclass/storage"
)

type RESTStorageProvider struct{}

func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, error) {
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(resource.GroupName, legacyscheme.Scheme, legacyscheme.ParameterCodec, legacyscheme.Codecs)
	// If you add a version here, be sure to add an entry in `k8s.io/kubernetes/cmd/kube-apiserver/app/aggregator.go with specific priorities.
	// TODO refactor the plumbing to provide the information in the APIGroupInfo

	if storageMap, err := p.v1alpha1Storage(apiResourceConfigSource, restOptionsGetter); err != nil {
		return genericapiserver.APIGroupInfo{}, err
	} else if len(storageMap) > 0 {
		apiGroupInfo.VersionedResourcesStorageMap[resourcev1alpha1.SchemeGroupVersion.Version] = storageMap
	}

	return apiGroupInfo, nil
}

func (p RESTStorageProvider) v1alpha1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (map[string]rest.Storage, error) {
	storage := map[string]rest.Storage{}

	if resource := "resourceclasses"; apiResourceConfigSource.ResourceEnabled(resourcev1alpha1.SchemeGroupVersion.WithResource(resource)) {
		resourceClassStorage, err := resourceclassstore.NewREST(restOptionsGetter)
		if err != nil {
			return nil, err
		}
		storage[resource] = resourceClassStorage
	}

	if resource := "resourceclaims"; apiResourceConfigSource.ResourceEnabled(resourcev1alpha1.SchemeGroupVersion.WithResource(resource)) {
		resourceClaimStorage, resourceClaimStatusStorage, err := resourceclaimstore.NewREST(restOptionsGetter)
		if err != nil {
			return nil, err
		}
		storage[resource] = resourceClaimStorage
		storage[resource+"/status"] = resourceClaimStatusStorage
	}

	if resource := "resourceclaimtemplates"; apiResourceConfigSource.ResourceEnabled(resourcev1alpha1.SchemeGroupVersion.WithResource(resource)) {
		resourceClaimTemplateStorage, err := resourceclaimtemplatestore.NewREST(restOptionsGetter)
		if err != nil {
			return nil, err
		}
		storage[resource] = resourceClaimTemplateStorage
	}

	if resource := "podschedulings"; apiResourceConfigSource.ResourceEnabled(resourcev1alpha1.SchemeGroupVersion.WithResource(resource)) {
		podSchedulingStorage, podSchedulingStatusStorage, err := podschedulingstore.NewREST(restOptionsGetter)
		if err != nil {
			return nil, err
		}
		storage[resource] = podSchedulingStorage
		storage[resource+"/status"] = podSchedulingStatusStorage
	}

	return storage, nil
}

func (p RESTStorageProvider) GroupName() string {
	return resource.GroupName
}
