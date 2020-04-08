/*
Copyright 2016 The Kubernetes Authors.

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
	appsapiv1 "k8s.io/api/apps/v1"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/apps"
	controllerrevisionsstore "k8s.io/kubernetes/pkg/registry/apps/controllerrevision/storage"
	daemonsetstore "k8s.io/kubernetes/pkg/registry/apps/daemonset/storage"
	deploymentstore "k8s.io/kubernetes/pkg/registry/apps/deployment/storage"
	replicasetstore "k8s.io/kubernetes/pkg/registry/apps/replicaset/storage"
	statefulsetstore "k8s.io/kubernetes/pkg/registry/apps/statefulset/storage"
)

// StorageProvider is a struct for apps REST storage.
type StorageProvider struct{}

// NewRESTStorage returns APIGroupInfo object.
func (p StorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool, error) {
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(apps.GroupName, legacyscheme.Scheme, legacyscheme.ParameterCodec, legacyscheme.Codecs)
	// If you add a version here, be sure to add an entry in `k8s.io/kubernetes/cmd/kube-apiserver/app/aggregator.go with specific priorities.
	// TODO refactor the plumbing to provide the information in the APIGroupInfo

	if apiResourceConfigSource.VersionEnabled(appsapiv1.SchemeGroupVersion) {
		storageMap, err := p.v1Storage(apiResourceConfigSource, restOptionsGetter)
		if err != nil {
			return genericapiserver.APIGroupInfo{}, false, err
		}
		apiGroupInfo.VersionedResourcesStorageMap[appsapiv1.SchemeGroupVersion.Version] = storageMap
	}

	return apiGroupInfo, true, nil
}

func (p StorageProvider) v1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (map[string]rest.Storage, error) {
	storage := map[string]rest.Storage{}

	// deployments
	deploymentStorage, err := deploymentstore.NewStorage(restOptionsGetter)
	if err != nil {
		return storage, err
	}
	storage["deployments"] = deploymentStorage.Deployment
	storage["deployments/status"] = deploymentStorage.Status
	storage["deployments/scale"] = deploymentStorage.Scale

	// statefulsets
	statefulSetStorage, err := statefulsetstore.NewStorage(restOptionsGetter)
	if err != nil {
		return storage, err
	}
	storage["statefulsets"] = statefulSetStorage.StatefulSet
	storage["statefulsets/status"] = statefulSetStorage.Status
	storage["statefulsets/scale"] = statefulSetStorage.Scale

	// daemonsets
	daemonSetStorage, daemonSetStatusStorage, err := daemonsetstore.NewREST(restOptionsGetter)
	if err != nil {
		return storage, err
	}
	storage["daemonsets"] = daemonSetStorage
	storage["daemonsets/status"] = daemonSetStatusStorage

	// replicasets
	replicaSetStorage, err := replicasetstore.NewStorage(restOptionsGetter)
	if err != nil {
		return storage, err
	}
	storage["replicasets"] = replicaSetStorage.ReplicaSet
	storage["replicasets/status"] = replicaSetStorage.Status
	storage["replicasets/scale"] = replicaSetStorage.Scale

	// controllerrevisions
	historyStorage, err := controllerrevisionsstore.NewREST(restOptionsGetter)
	if err != nil {
		return storage, err
	}
	storage["controllerrevisions"] = historyStorage

	return storage, nil
}

// GroupName returns name of the group
func (p StorageProvider) GroupName() string {
	return apps.GroupName
}
