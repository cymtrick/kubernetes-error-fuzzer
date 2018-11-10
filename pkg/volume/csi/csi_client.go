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

package csi

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"time"

	csipb "github.com/container-storage-interface/spec/lib/go/csi/v0"
	"google.golang.org/grpc"
	api "k8s.io/api/core/v1"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/features"
)

type csiClient interface {
	NodeGetInfo(ctx context.Context) (
		nodeID string,
		maxVolumePerNode int64,
		accessibleTopology *csipb.Topology,
		err error)
	NodePublishVolume(
		ctx context.Context,
		volumeid string,
		readOnly bool,
		stagingTargetPath string,
		targetPath string,
		accessMode api.PersistentVolumeAccessMode,
		volumeInfo map[string]string,
		volumeAttribs map[string]string,
		nodePublishSecrets map[string]string,
		fsType string,
		mountOptions []string,
	) error
	NodeUnpublishVolume(
		ctx context.Context,
		volID string,
		targetPath string,
	) error
	NodeStageVolume(ctx context.Context,
		volID string,
		publishVolumeInfo map[string]string,
		stagingTargetPath string,
		fsType string,
		accessMode api.PersistentVolumeAccessMode,
		nodeStageSecrets map[string]string,
		volumeAttribs map[string]string,
	) error
	NodeUnstageVolume(ctx context.Context, volID, stagingTargetPath string) error
	NodeGetCapabilities(ctx context.Context) ([]*csipb.NodeServiceCapability, error)
}

// csiClient encapsulates all csi-plugin methods
type csiDriverClient struct {
	driverName        string
	nodeClientCreator nodeClientCreator
}

var _ csiClient = &csiDriverClient{}

type nodeClientCreator func(driverName string) (
	nodeClient csipb.NodeClient,
	closer io.Closer,
	err error,
)

// newNodeClient creates a new NodeClient with the internally used gRPC
// connection set up. It also returns a closer which must to be called to close
// the gRPC connection when the NodeClient is not used anymore.
// This is the default implementation for the nodeClientCreator, used in
// newCsiDriverClient.
func newNodeClient(driverName string) (nodeClient csipb.NodeClient, closer io.Closer, err error) {
	var conn *grpc.ClientConn
	conn, err = newGrpcConn(driverName)
	if err != nil {
		return nil, nil, err
	}

	nodeClient = csipb.NewNodeClient(conn)
	return nodeClient, conn, nil
}

func newCsiDriverClient(driverName string) *csiDriverClient {
	c := &csiDriverClient{
		driverName:        driverName,
		nodeClientCreator: newNodeClient,
	}
	return c
}

func (c *csiDriverClient) NodeGetInfo(ctx context.Context) (
	nodeID string,
	maxVolumePerNode int64,
	accessibleTopology *csipb.Topology,
	err error) {
	klog.V(4).Info(log("calling NodeGetInfo rpc"))

	nodeClient, closer, err := c.nodeClientCreator(c.driverName)
	if err != nil {
		return "", 0, nil, err
	}
	defer closer.Close()

	res, err := nodeClient.NodeGetInfo(ctx, &csipb.NodeGetInfoRequest{})
	if err != nil {
		return "", 0, nil, err
	}

	return res.GetNodeId(), res.GetMaxVolumesPerNode(), res.GetAccessibleTopology(), nil
}

func (c *csiDriverClient) NodePublishVolume(
	ctx context.Context,
	volID string,
	readOnly bool,
	stagingTargetPath string,
	targetPath string,
	accessMode api.PersistentVolumeAccessMode,
	volumeInfo map[string]string,
	volumeAttribs map[string]string,
	nodePublishSecrets map[string]string,
	fsType string,
	mountOptions []string,
) error {
	klog.V(4).Info(log("calling NodePublishVolume rpc [volid=%s,target_path=%s]", volID, targetPath))
	if volID == "" {
		return errors.New("missing volume id")
	}
	if targetPath == "" {
		return errors.New("missing target path")
	}

	nodeClient, closer, err := c.nodeClientCreator(c.driverName)
	if err != nil {
		return err
	}
	defer closer.Close()

	req := &csipb.NodePublishVolumeRequest{
		VolumeId:           volID,
		TargetPath:         targetPath,
		Readonly:           readOnly,
		PublishInfo:        volumeInfo,
		VolumeAttributes:   volumeAttribs,
		NodePublishSecrets: nodePublishSecrets,
		VolumeCapability: &csipb.VolumeCapability{
			AccessMode: &csipb.VolumeCapability_AccessMode{
				Mode: asCSIAccessMode(accessMode),
			},
		},
	}
	if stagingTargetPath != "" {
		req.StagingTargetPath = stagingTargetPath
	}

	if fsType == fsTypeBlockName {
		req.VolumeCapability.AccessType = &csipb.VolumeCapability_Block{
			Block: &csipb.VolumeCapability_BlockVolume{},
		}
	} else {
		req.VolumeCapability.AccessType = &csipb.VolumeCapability_Mount{
			Mount: &csipb.VolumeCapability_MountVolume{
				FsType:     fsType,
				MountFlags: mountOptions,
			},
		}
	}

	_, err = nodeClient.NodePublishVolume(ctx, req)
	return err
}

func (c *csiDriverClient) NodeUnpublishVolume(ctx context.Context, volID string, targetPath string) error {
	klog.V(4).Info(log("calling NodeUnpublishVolume rpc: [volid=%s, target_path=%s", volID, targetPath))
	if volID == "" {
		return errors.New("missing volume id")
	}
	if targetPath == "" {
		return errors.New("missing target path")
	}

	nodeClient, closer, err := c.nodeClientCreator(c.driverName)
	if err != nil {
		return err
	}
	defer closer.Close()

	req := &csipb.NodeUnpublishVolumeRequest{
		VolumeId:   volID,
		TargetPath: targetPath,
	}

	_, err = nodeClient.NodeUnpublishVolume(ctx, req)
	return err
}

func (c *csiDriverClient) NodeStageVolume(ctx context.Context,
	volID string,
	publishInfo map[string]string,
	stagingTargetPath string,
	fsType string,
	accessMode api.PersistentVolumeAccessMode,
	nodeStageSecrets map[string]string,
	volumeAttribs map[string]string,
) error {
	klog.V(4).Info(log("calling NodeStageVolume rpc [volid=%s,staging_target_path=%s]", volID, stagingTargetPath))
	if volID == "" {
		return errors.New("missing volume id")
	}
	if stagingTargetPath == "" {
		return errors.New("missing staging target path")
	}

	nodeClient, closer, err := c.nodeClientCreator(c.driverName)
	if err != nil {
		return err
	}
	defer closer.Close()

	req := &csipb.NodeStageVolumeRequest{
		VolumeId:          volID,
		PublishInfo:       publishInfo,
		StagingTargetPath: stagingTargetPath,
		VolumeCapability: &csipb.VolumeCapability{
			AccessMode: &csipb.VolumeCapability_AccessMode{
				Mode: asCSIAccessMode(accessMode),
			},
		},
		NodeStageSecrets: nodeStageSecrets,
		VolumeAttributes: volumeAttribs,
	}

	if fsType == fsTypeBlockName {
		req.VolumeCapability.AccessType = &csipb.VolumeCapability_Block{
			Block: &csipb.VolumeCapability_BlockVolume{},
		}
	} else {
		req.VolumeCapability.AccessType = &csipb.VolumeCapability_Mount{
			Mount: &csipb.VolumeCapability_MountVolume{
				FsType: fsType,
			},
		}
	}

	_, err = nodeClient.NodeStageVolume(ctx, req)
	return err
}

func (c *csiDriverClient) NodeUnstageVolume(ctx context.Context, volID, stagingTargetPath string) error {
	klog.V(4).Info(log("calling NodeUnstageVolume rpc [volid=%s,staging_target_path=%s]", volID, stagingTargetPath))
	if volID == "" {
		return errors.New("missing volume id")
	}
	if stagingTargetPath == "" {
		return errors.New("missing staging target path")
	}

	nodeClient, closer, err := c.nodeClientCreator(c.driverName)
	if err != nil {
		return err
	}
	defer closer.Close()

	req := &csipb.NodeUnstageVolumeRequest{
		VolumeId:          volID,
		StagingTargetPath: stagingTargetPath,
	}
	_, err = nodeClient.NodeUnstageVolume(ctx, req)
	return err
}

func (c *csiDriverClient) NodeGetCapabilities(ctx context.Context) ([]*csipb.NodeServiceCapability, error) {
	klog.V(4).Info(log("calling NodeGetCapabilities rpc"))

	nodeClient, closer, err := c.nodeClientCreator(c.driverName)
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	req := &csipb.NodeGetCapabilitiesRequest{}
	resp, err := nodeClient.NodeGetCapabilities(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.GetCapabilities(), nil
}

func asCSIAccessMode(am api.PersistentVolumeAccessMode) csipb.VolumeCapability_AccessMode_Mode {
	switch am {
	case api.ReadWriteOnce:
		return csipb.VolumeCapability_AccessMode_SINGLE_NODE_WRITER
	case api.ReadOnlyMany:
		return csipb.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY
	case api.ReadWriteMany:
		return csipb.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER
	}
	return csipb.VolumeCapability_AccessMode_UNKNOWN
}

func newGrpcConn(driverName string) (*grpc.ClientConn, error) {
	if driverName == "" {
		return nil, fmt.Errorf("driver name is empty")
	}
	addr := fmt.Sprintf(csiAddrTemplate, driverName)
	// TODO once KubeletPluginsWatcher graduates to beta, remove FeatureGate check
	if utilfeature.DefaultFeatureGate.Enabled(features.KubeletPluginsWatcher) {
		csiDrivers.RLock()
		driver, ok := csiDrivers.driversMap[driverName]
		csiDrivers.RUnlock()

		if !ok {
			return nil, fmt.Errorf("driver name %s not found in the list of registered CSI drivers", driverName)
		}
		addr = driver.driverEndpoint
	}
	network := "unix"
	klog.V(4).Infof(log("creating new gRPC connection for [%s://%s]", network, addr))

	return grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithDialer(func(target string, timeout time.Duration) (net.Conn, error) {
			return net.Dial(network, target)
		}),
	)
}
