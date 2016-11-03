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

package photon_pd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/golang/glog"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/cloudprovider"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/photon"
	"k8s.io/kubernetes/pkg/volume"
	volumeutil "k8s.io/kubernetes/pkg/volume/util"
)

const (
	maxRetries         = 10
	checkSleepDuration = time.Second
	diskByIDPath       = "/dev/disk/by-id/"
	diskPhotonPrefix   = "wwn-0x"
)

var ErrProbeVolume = errors.New("Error scanning attached volumes")
var volNameToDeviceName = make(map[string]string)

type PhotonDiskUtil struct{}

func logError(msg string, err error) error {
	s := "Photon Controller utility: " + msg + ". Error [" + err.Error() + "]"
	glog.Errorf(s)
	return fmt.Errorf(s)
}

func removeFromScsiSubsystem(volName string) {
	// TODO: if using pvscsi controller, this won't be needed
	deviceName := volNameToDeviceName[volName]
	fileName := "/sys/block/" + deviceName + "/device/delete"
	data := []byte("1")
	ioutil.WriteFile(fileName, data, 0666)
}

func scsiHostScan() {
	// TODO: if using pvscsi controller, this won't be needed
	scsi_path := "/sys/class/scsi_host/"
	if dirs, err := ioutil.ReadDir(scsi_path); err == nil {
		for _, f := range dirs {
			name := scsi_path + f.Name() + "/scan"
			data := []byte("- - -")
			ioutil.WriteFile(name, data, 0666)
			glog.Errorf("scsiHostScan scan for %s", name)
		}
	}
}

func verifyDevicePath(path string) (string, error) {
	if pathExists, err := volumeutil.PathExists(path); err != nil {
		return "", fmt.Errorf("Error checking if path exists: %v", err)
	} else if pathExists {
		return path, nil
	}

	glog.V(4).Infof("verifyDevicePath: path not exists yet")
	return "", nil
}

// CreateVolume creates a PhotonController persistent disk.
func (util *PhotonDiskUtil) CreateVolume(p *photonPersistentDiskProvisioner) (pdID string, capacityGB int, err error) {
	cloud, err := getCloudProvider(p.plugin.host.GetCloudProvider())
	if err != nil {
		return "", 0, logError("CreateVolume failed to get cloud provider", err)
	}

	capacity := p.options.PVC.Spec.Resources.Requests[api.ResourceName(api.ResourceStorage)]
	volSizeBytes := capacity.Value()
	// PhotonController works with GB, convert to GB with rounding up
	volSizeGB := int(volume.RoundUpSize(volSizeBytes, 1024*1024*1024))
	name := volume.GenerateVolumeName(p.options.ClusterName, p.options.PVName, 255)
	volumeOptions := &photon.VolumeOptions{
		CapacityGB: volSizeGB,
		Tags:       *p.options.CloudTags,
		Name:       name,
	}

	for parameter, value := range p.options.Parameters {
		switch strings.ToLower(parameter) {
		case "flavor":
			volumeOptions.Flavor = value
		default:
			return "", 0, logError("invalid option "+parameter+" for volume plugin "+p.plugin.GetPluginName(), err)
		}
	}

	pdID, err = cloud.CreateDisk(volumeOptions)
	if err != nil {
		return "", 0, logError("failed to CreateDisk", err)
	}

	glog.V(4).Infof("Successfully created Photon Controller persistent disk %s", name)
	return pdID, volSizeGB, nil
}

// DeleteVolume deletes a vSphere volume.
func (util *PhotonDiskUtil) DeleteVolume(pd *photonPersistentDiskDeleter) error {
	cloud, err := getCloudProvider(pd.plugin.host.GetCloudProvider())
	if err != nil {
		return logError("DeleteVolume failed to get cloud provider", err)
	}

	if err = cloud.DeleteDisk(pd.pdID); err != nil {
		return logError("failed to DeleteDisk for pdID "+pd.pdID, err)
	}

	glog.V(4).Infof("Successfully deleted PhotonController persistent disk %s", pd.pdID)
	return nil
}

func getCloudProvider(cloud cloudprovider.Interface) (*photon.PCCloud, error) {
	if cloud == nil {
		return nil, logError("Cloud provider not initialized properly", nil)
	}

	pcc := cloud.(*photon.PCCloud)
	if pcc == nil {
		return nil, logError("Invalid cloud provider: expected Photon Controller", nil)
	}
	return pcc, nil
}
