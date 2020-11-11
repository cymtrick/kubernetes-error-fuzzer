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

package quobyte

import (
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	v1 "k8s.io/api/core/v1"
	volumehelpers "k8s.io/cloud-provider/volume/helpers"
	proxyutil "k8s.io/kubernetes/pkg/proxy/util"

	quobyteapi "github.com/quobyte/api"
	"k8s.io/klog/v2"
)

type quobyteVolumeManager struct {
	config      *quobyteAPIConfig
	dialOptions *proxyutil.FilteredDialOptions
}

func (manager *quobyteVolumeManager) createVolume(provisioner *quobyteVolumeProvisioner, createQuota bool) (quobyte *v1.QuobyteVolumeSource, size int, err error) {
	capacity := provisioner.options.PVC.Spec.Resources.Requests[v1.ResourceName(v1.ResourceStorage)]
	volumeSize, err := volumehelpers.RoundUpToGiBInt(capacity)
	if err != nil {
		return nil, 0, err
	}
	// Quobyte has the concept of Volumes which doesn't have a specific size (they can grow unlimited)
	// to simulate a size constraint we set here a Quota for logical space
	volumeRequest := &quobyteapi.CreateVolumeRequest{
		Name:              provisioner.volume,
		RootUserID:        provisioner.user,
		RootGroupID:       provisioner.group,
		TenantID:          provisioner.tenant,
		ConfigurationName: provisioner.config,
	}

	quobyteClient := manager.createQuobyteClient()
	volumeUUID, err := quobyteClient.CreateVolume(volumeRequest)
	if err != nil {
		return &v1.QuobyteVolumeSource{}, volumeSize, err
	}

	// Set Quota for Volume with specified byte size
	if createQuota {
		err = quobyteClient.SetVolumeQuota(volumeUUID, uint64(capacity.Value()))
		if err != nil {
			return &v1.QuobyteVolumeSource{}, volumeSize, err
		}
	}

	klog.V(4).Infof("Created Quobyte volume %s", provisioner.volume)
	return &v1.QuobyteVolumeSource{
		Registry: provisioner.registry,
		Volume:   provisioner.volume,
		User:     provisioner.user,
		Group:    provisioner.group,
	}, volumeSize, nil
}

func (manager *quobyteVolumeManager) deleteVolume(deleter *quobyteVolumeDeleter) error {
	return manager.createQuobyteClient().DeleteVolumeByName(deleter.volume, deleter.tenant)
}

func (manager *quobyteVolumeManager) createQuobyteClient() *quobyteapi.QuobyteClient {
	client := quobyteapi.NewQuobyteClient(
		manager.config.quobyteAPIServer,
		manager.config.quobyteUser,
		manager.config.quobytePassword,
	)
	// quobyte client library @v0.1.7 uses a zero-value http.Client with a nil
	// transport which is equivalent to using http.DefaultTransport.
	rt := http.DefaultTransport.(*http.Transport).Clone()
	rt.DialContext = proxyutil.NewFilteredDialContext(rt.DialContext, nil, manager.dialOptions)
	client.SetTransport(rt)
	return client
}

func (mounter *quobyteMounter) pluginDirIsMounted(pluginDir string) (bool, error) {
	mounts, err := mounter.mounter.List()
	if err != nil {
		return false, err
	}

	for _, mountPoint := range mounts {
		if strings.HasPrefix(mountPoint.Type, "quobyte") {
			continue
		}

		if mountPoint.Path == pluginDir {
			klog.V(4).Infof("quobyte: found mountpoint %s in /proc/mounts", mountPoint.Path)
			return true, nil
		}
	}

	return false, nil
}

func (mounter *quobyteMounter) correctTraillingSlash(regStr string) string {
	return filepath.Clean(regStr) + string(os.PathSeparator)
}

func validateRegistry(registry string) bool {
	if len(registry) == 0 {
		return false
	}

	for _, hostPortPair := range strings.Split(registry, ",") {
		if _, _, err := net.SplitHostPort(hostPortPair); err != nil {
			return false
		}
	}

	return true
}
