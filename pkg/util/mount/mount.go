/*
Copyright 2014 The Kubernetes Authors.

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

// TODO(thockin): This whole pkg is pretty linux-centric.  As soon as we have
// an alternate platform, we will need to abstract further.
package mount

import (
	"os"
	"path/filepath"
)

type FileType string

const (
	// Default mount command if mounter path is not specified
	defaultMountCommand           = "mount"
	MountsInGlobalPDPath          = "mounts"
	FileTypeDirectory    FileType = "Directory"
	FileTypeFile         FileType = "File"
	FileTypeSocket       FileType = "Socket"
	FileTypeCharDev      FileType = "CharDevice"
	FileTypeBlockDev     FileType = "BlockDevice"
)

type Interface interface {
	// Mount mounts source to target as fstype with given options.
	Mount(source string, target string, fstype string, options []string) error
	// Unmount unmounts given target.
	Unmount(target string) error
	// List returns a list of all mounted filesystems.  This can be large.
	// On some platforms, reading mounts is not guaranteed consistent (i.e.
	// it could change between chunked reads). This is guaranteed to be
	// consistent.
	List() ([]MountPoint, error)
	// IsMountPointMatch determines if the mountpoint matches the dir
	IsMountPointMatch(mp MountPoint, dir string) bool
	// IsNotMountPoint determines if a directory is a mountpoint.
	// It should return ErrNotExist when the directory does not exist.
	// IsNotMountPoint is more expensive than IsLikelyNotMountPoint.
	// IsNotMountPoint detects bind mounts in linux.
	// IsNotMountPoint enumerates all the mountpoints using List() and
	// the list of mountpoints may be large, then it uses
	// IsMountPointMatch to evaluate whether the directory is a mountpoint
	IsNotMountPoint(file string) (bool, error)
	// IsLikelyNotMountPoint uses heuristics to determine if a directory
	// is a mountpoint.
	// It should return ErrNotExist when the directory does not exist.
	// IsLikelyNotMountPoint does NOT properly detect all mountpoint types
	// most notably linux bind mounts.
	IsLikelyNotMountPoint(file string) (bool, error)
	// DeviceOpened determines if the device is in use elsewhere
	// on the system, i.e. still mounted.
	DeviceOpened(pathname string) (bool, error)
	// PathIsDevice determines if a path is a device.
	PathIsDevice(pathname string) (bool, error)
	// GetDeviceNameFromMount finds the device name by checking the mount path
	// to get the global mount path which matches its plugin directory
	GetDeviceNameFromMount(mountPath, pluginDir string) (string, error)
	// MakeRShared checks that given path is on a mount with 'rshared' mount
	// propagation. If not, it bind-mounts the path as rshared.
	MakeRShared(path string) error
	// GetFileType checks for file/directory/socket/block/character devices.
	// Will operate in the host mount namespace if kubelet is running in a container
	GetFileType(pathname string) (FileType, error)
	// MakeFile creates an empty file.
	// Will operate in the host mount namespace if kubelet is running in a container
	MakeFile(pathname string) error
	// MakeDir creates a new directory.
	// Will operate in the host mount namespace if kubelet is running in a container
	MakeDir(pathname string) error
	// ExistsPath checks whether the path exists.
	// Will operate in the host mount namespace if kubelet is running in a container
	ExistsPath(pathname string) bool
}

// Exec executes command where mount utilities are. This can be either the host,
// container where kubelet runs or even a remote pod with mount utilities.
// Usual pkg/util/exec interface is not used because kubelet.RunInContainer does
// not provide stdin/stdout/stderr streams.
type Exec interface {
	// Run executes a command and returns its stdout + stderr combined in one
	// stream.
	Run(cmd string, args ...string) ([]byte, error)
}

// Compile-time check to ensure all Mounter implementations satisfy
// the mount interface
var _ Interface = &Mounter{}

// This represents a single line in /proc/mounts or /etc/fstab.
type MountPoint struct {
	Device string
	Path   string
	Type   string
	Opts   []string
	Freq   int
	Pass   int
}

// SafeFormatAndMount probes a device to see if it is formatted.
// Namely it checks to see if a file system is present. If so it
// mounts it otherwise the device is formatted first then mounted.
type SafeFormatAndMount struct {
	Interface
	Exec
}

// FormatAndMount formats the given disk, if needed, and mounts it.
// That is if the disk is not formatted and it is not being mounted as
// read-only it will format it first then mount it. Otherwise, if the
// disk is already formatted or it is being mounted as read-only, it
// will be mounted without formatting.
func (mounter *SafeFormatAndMount) FormatAndMount(source string, target string, fstype string, options []string) error {
	return mounter.formatAndMount(source, target, fstype, options)
}

// GetMountRefsByDev finds all references to the device provided
// by mountPath; returns a list of paths.
func GetMountRefsByDev(mounter Interface, mountPath string) ([]string, error) {
	mps, err := mounter.List()
	if err != nil {
		return nil, err
	}
	slTarget, err := filepath.EvalSymlinks(mountPath)
	if err != nil {
		slTarget = mountPath
	}

	// Finding the device mounted to mountPath
	diskDev := ""
	for i := range mps {
		if slTarget == mps[i].Path {
			diskDev = mps[i].Device
			break
		}
	}

	// Find all references to the device.
	var refs []string
	for i := range mps {
		if mps[i].Device == diskDev || mps[i].Device == slTarget {
			if mps[i].Path != slTarget {
				refs = append(refs, mps[i].Path)
			}
		}
	}
	return refs, nil
}

// GetDeviceNameFromMount: given a mnt point, find the device from /proc/mounts
// returns the device name, reference count, and error code
func GetDeviceNameFromMount(mounter Interface, mountPath string) (string, int, error) {
	mps, err := mounter.List()
	if err != nil {
		return "", 0, err
	}

	// Find the device name.
	// FIXME if multiple devices mounted on the same mount path, only the first one is returned
	device := ""
	// If mountPath is symlink, need get its target path.
	slTarget, err := filepath.EvalSymlinks(mountPath)
	if err != nil {
		slTarget = mountPath
	}
	for i := range mps {
		if mps[i].Path == slTarget {
			device = mps[i].Device
			break
		}
	}

	// Find all references to the device.
	refCount := 0
	for i := range mps {
		if mps[i].Device == device {
			refCount++
		}
	}
	return device, refCount, nil
}

// IsNotMountPoint determines if a directory is a mountpoint.
// It should return ErrNotExist when the directory does not exist.
// This method uses the List() of all mountpoints
// It is more extensive than IsLikelyNotMountPoint
// and it detects bind mounts in linux
func IsNotMountPoint(mounter Interface, file string) (bool, error) {
	// IsLikelyNotMountPoint provides a quick check
	// to determine whether file IS A mountpoint
	notMnt, notMntErr := mounter.IsLikelyNotMountPoint(file)
	if notMntErr != nil && os.IsPermission(notMntErr) {
		// We were not allowed to do the simple stat() check, e.g. on NFS with
		// root_squash. Fall back to /proc/mounts check below.
		notMnt = true
		notMntErr = nil
	}
	if notMntErr != nil {
		return notMnt, notMntErr
	}
	// identified as mountpoint, so return this fact
	if notMnt == false {
		return notMnt, nil
	}
	// check all mountpoints since IsLikelyNotMountPoint
	// is not reliable for some mountpoint types
	mountPoints, mountPointsErr := mounter.List()
	if mountPointsErr != nil {
		return notMnt, mountPointsErr
	}
	for _, mp := range mountPoints {
		if mounter.IsMountPointMatch(mp, file) {
			notMnt = false
			break
		}
	}
	return notMnt, nil
}

// isBind detects whether a bind mount is being requested and makes the remount options to
// use in case of bind mount, due to the fact that bind mount doesn't respect mount options.
// The list equals:
//   options - 'bind' + 'remount' (no duplicate)
func isBind(options []string) (bool, []string) {
	bindRemountOpts := []string{"remount"}
	bind := false

	if len(options) != 0 {
		for _, option := range options {
			switch option {
			case "bind":
				bind = true
				break
			case "remount":
				break
			default:
				bindRemountOpts = append(bindRemountOpts, option)
			}
		}
	}

	return bind, bindRemountOpts
}
