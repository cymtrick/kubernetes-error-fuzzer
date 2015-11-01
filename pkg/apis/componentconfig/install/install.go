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

// Package install installs the experimental API group, making it available as
// an option to all of the API encoding/decoding machinery.
package install

import (
	"fmt"
	"strings"

	"github.com/golang/glog"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/latest"
	"k8s.io/kubernetes/pkg/api/meta"
	"k8s.io/kubernetes/pkg/api/registered"
	apiutil "k8s.io/kubernetes/pkg/api/util"
	_ "k8s.io/kubernetes/pkg/apis/componentconfig"
	"k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/sets"
)

const importPrefix = "k8s.io/kubernetes/pkg/apis/componentconfig"

var accessor = meta.NewAccessor()

func init() {
	groupMeta, err := latest.RegisterGroup("componentconfig")
	if err != nil {
		glog.V(4).Infof("%v", err)
		return
	}
	registeredGroupVersions := registered.GroupVersionsForGroup("componentconfig")
	groupVersion := registeredGroupVersions[0]
	*groupMeta = latest.GroupMeta{
		GroupVersion: groupVersion,
		Group:        apiutil.GetGroup(groupVersion),
		Version:      apiutil.GetVersion(groupVersion),
		Codec:        runtime.CodecFor(api.Scheme, groupVersion),
	}
	var versions []string
	var groupVersions []string
	for i := len(registeredGroupVersions) - 1; i >= 0; i-- {
		versions = append(versions, apiutil.GetVersion(registeredGroupVersions[i]))
		groupVersions = append(groupVersions, registeredGroupVersions[i])
	}
	groupMeta.Versions = versions
	groupMeta.GroupVersions = groupVersions

	groupMeta.SelfLinker = runtime.SelfLinker(accessor)

	// the list of kinds that are scoped at the root of the api hierarchy
	// if a kind is not enumerated here, it is assumed to have a namespace scope
	rootScoped := sets.NewString()

	ignoredKinds := sets.NewString()

	groupMeta.RESTMapper = api.NewDefaultRESTMapper("componentconfig", groupVersions, interfacesFor, importPrefix, ignoredKinds, rootScoped)
	api.RegisterRESTMapper(groupMeta.RESTMapper)
	groupMeta.InterfacesFor = interfacesFor
}

// interfacesFor returns the default Codec and ResourceVersioner for a given version
// string, or an error if the version is not known.
func interfacesFor(version string) (*meta.VersionInterfaces, error) {
	switch version {
	case "componentconfig/v1alpha1":
		return &meta.VersionInterfaces{
			Codec:            v1alpha1.Codec,
			ObjectConvertor:  api.Scheme,
			MetadataAccessor: accessor,
		}, nil
	default:
		g, _ := latest.Group("componentconfig")
		return nil, fmt.Errorf("unsupported storage version: %s (valid: %s)", version, strings.Join(g.Versions, ", "))
	}
}
