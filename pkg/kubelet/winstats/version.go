// +build windows

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

package winstats

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
)

type OSInfo struct {
	BuildNumber, ProductName string
	MajorVersion, MinorVersion, UBR uint64
}

// GetOSInfo reads Windows version information from the registry
func GetOSInfo() (*OSInfo, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	buildNumber, _, err := k.GetStringValue("CurrentBuildNumber")
	if err != nil {
		return nil, err
	}

	majorVersionNumber, _, err := k.GetIntegerValue("CurrentMajorVersionNumber")
	if err != nil {
		return nil, err
	}

	minorVersionNumber, _, err := k.GetIntegerValue("CurrentMinorVersionNumber")
	if err != nil {
		return nil, err
	}

	revision, _, err := k.GetIntegerValue("UBR")
	if err != nil {
		return nil, err
	}

	productName, _, err := k.GetStringValue("ProductName")
	if err != nil {
		return nil, nil
	}

	return &OSInfo{
		BuildNumber:  buildNumber,
		ProductName:  productName,
		MajorVersion: majorVersionNumber,
		MinorVersion: minorVersionNumber,
		UBR:          revision,
	}, nil
}

func (o *OSInfo) GetPatchVersion() string {
	return fmt.Sprintf("%d.%d.%s.%d", o.MajorVersion, o.MinorVersion, o.BuildNumber, o.UBR)
}

func (o *OSInfo) GetBuild() string {
	return fmt.Sprintf("%d.%d.%s", o.MajorVersion, o.MinorVersion, o.BuildNumber)
}
