/*
Copyright 2015 The Kubernetes Authors.

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

package registry

import (
	"github.com/golang/glog"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	etcdstorage "k8s.io/apiserver/pkg/storage/etcd"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	"k8s.io/apiserver/pkg/storage/storagebackend/factory"
)

// Creates a cacher based given storageConfig.
func StorageWithCacher(capacity int) generic.StorageDecorator {
	return func(
		storageConfig *storagebackend.Config,
		objectType runtime.Object,
		resourcePrefix string,
		keyFunc func(obj runtime.Object) (string, error),
		newListFunc func() runtime.Object,
		getAttrsFunc storage.AttrFunc,
		triggerFunc storage.TriggerPublisherFunc) (storage.Interface, factory.DestroyFunc) {

		s, d := generic.NewRawStorage(storageConfig)
		if capacity == 0 {
			glog.V(5).Infof("Storage caching is disabled for %T", objectType)
			return s, d
		}
		glog.V(5).Infof("Storage caching is enabled for %T with capacity %v", objectType, capacity)

		// TODO: we would change this later to make storage always have cacher and hide low level KV layer inside.
		// Currently it has two layers of same storage interface -- cacher and low level kv.
		cacherConfig := storage.CacherConfig{
			CacheCapacity:        capacity,
			Storage:              s,
			Versioner:            etcdstorage.APIObjectVersioner{},
			Type:                 objectType,
			ResourcePrefix:       resourcePrefix,
			KeyFunc:              keyFunc,
			NewListFunc:          newListFunc,
			GetAttrsFunc:         getAttrsFunc,
			TriggerPublisherFunc: triggerFunc,
			Codec:                storageConfig.Codec,
		}
		cacher := storage.NewCacherFromConfig(cacherConfig)
		destroyFunc := func() {
			cacher.Stop()
			d()
		}

		return cacher, destroyFunc
	}
}
