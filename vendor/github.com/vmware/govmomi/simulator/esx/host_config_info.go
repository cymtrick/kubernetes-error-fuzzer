/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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

package esx

import "github.com/vmware/govmomi/vim25/types"

// HostConfigInfo is the default template for the HostSystem config property.
// Capture method:
//   govc object.collect -s -dump HostSystem:ha-host config
var HostConfigInfo = types.HostConfigInfo{
	Host: types.ManagedObjectReference{Type: "HostSystem", Value: "ha-host"},
	Product: types.AboutInfo{
		Name:                  "VMware ESXi",
		FullName:              "VMware ESXi 6.5.0 build-5969303",
		Vendor:                "VMware, Inc.",
		Version:               "6.5.0",
		Build:                 "5969303",
		LocaleVersion:         "INTL",
		LocaleBuild:           "000",
		OsType:                "vmnix-x86",
		ProductLineId:         "embeddedEsx",
		ApiType:               "HostAgent",
		ApiVersion:            "6.5",
		InstanceUuid:          "",
		LicenseProductName:    "VMware ESX Server",
		LicenseProductVersion: "6.0",
	},
	DeploymentInfo: &types.HostDeploymentInfo{
		BootedFromStatelessCache: types.NewBool(false),
	},
	HyperThread: &types.HostHyperThreadScheduleInfo{
		Available: false,
		Active:    false,
		Config:    true,
	},
	ConsoleReservation:        (*types.ServiceConsoleReservationInfo)(nil),
	VirtualMachineReservation: (*types.VirtualMachineMemoryReservationInfo)(nil),
	StorageDevice:             &HostStorageDeviceInfo,
	SystemFile:                nil,
	Network: &types.HostNetworkInfo{
		Vswitch: []types.HostVirtualSwitch{
			{
				Name:              "vSwitch0",
				Key:               "key-vim.host.VirtualSwitch-vSwitch0",
				NumPorts:          1536,
				NumPortsAvailable: 1530,
				Mtu:               1500,
				Portgroup:         []string{"key-vim.host.PortGroup-VM Network", "key-vim.host.PortGroup-Management Network"},
				Pnic:              []string{"key-vim.host.PhysicalNic-vmnic0"},
				Spec: types.HostVirtualSwitchSpec{
					NumPorts: 128,
					Bridge: &types.HostVirtualSwitchBondBridge{
						HostVirtualSwitchBridge: types.HostVirtualSwitchBridge{},
						NicDevice:               []string{"vmnic0"},
						Beacon: &types.HostVirtualSwitchBeaconConfig{
							Interval: 1,
						},
						LinkDiscoveryProtocolConfig: &types.LinkDiscoveryProtocolConfig{
							Protocol:  "cdp",
							Operation: "listen",
						},
					},
					Policy: &types.HostNetworkPolicy{
						Security: &types.HostNetworkSecurityPolicy{
							AllowPromiscuous: types.NewBool(false),
							MacChanges:       types.NewBool(true),
							ForgedTransmits:  types.NewBool(true),
						},
						NicTeaming: &types.HostNicTeamingPolicy{
							Policy:         "loadbalance_srcid",
							ReversePolicy:  types.NewBool(true),
							NotifySwitches: types.NewBool(true),
							RollingOrder:   types.NewBool(false),
							FailureCriteria: &types.HostNicFailureCriteria{
								CheckSpeed:        "minimum",
								Speed:             10,
								CheckDuplex:       types.NewBool(false),
								FullDuplex:        types.NewBool(false),
								CheckErrorPercent: types.NewBool(false),
								Percentage:        0,
								CheckBeacon:       types.NewBool(false),
							},
							NicOrder: &types.HostNicOrderPolicy{
								ActiveNic:  []string{"vmnic0"},
								StandbyNic: nil,
							},
						},
						OffloadPolicy: &types.HostNetOffloadCapabilities{
							CsumOffload:     types.NewBool(true),
							TcpSegmentation: types.NewBool(true),
							ZeroCopyXmit:    types.NewBool(true),
						},
						ShapingPolicy: &types.HostNetworkTrafficShapingPolicy{
							Enabled:          types.NewBool(false),
							AverageBandwidth: 0,
							PeakBandwidth:    0,
							BurstSize:        0,
						},
					},
					Mtu: 0,
				},
			},
		},
		ProxySwitch: nil,
		Portgroup: []types.HostPortGroup{
			{
				Key:     "key-vim.host.PortGroup-VM Network",
				Port:    nil,
				Vswitch: "key-vim.host.VirtualSwitch-vSwitch0",
				ComputedPolicy: types.HostNetworkPolicy{
					Security: &types.HostNetworkSecurityPolicy{
						AllowPromiscuous: types.NewBool(false),
						MacChanges:       types.NewBool(true),
						ForgedTransmits:  types.NewBool(true),
					},
					NicTeaming: &types.HostNicTeamingPolicy{
						Policy:         "loadbalance_srcid",
						ReversePolicy:  types.NewBool(true),
						NotifySwitches: types.NewBool(true),
						RollingOrder:   types.NewBool(false),
						FailureCriteria: &types.HostNicFailureCriteria{
							CheckSpeed:        "minimum",
							Speed:             10,
							CheckDuplex:       types.NewBool(false),
							FullDuplex:        types.NewBool(false),
							CheckErrorPercent: types.NewBool(false),
							Percentage:        0,
							CheckBeacon:       types.NewBool(false),
						},
						NicOrder: &types.HostNicOrderPolicy{
							ActiveNic:  []string{"vmnic0"},
							StandbyNic: nil,
						},
					},
					OffloadPolicy: &types.HostNetOffloadCapabilities{
						CsumOffload:     types.NewBool(true),
						TcpSegmentation: types.NewBool(true),
						ZeroCopyXmit:    types.NewBool(true),
					},
					ShapingPolicy: &types.HostNetworkTrafficShapingPolicy{
						Enabled:          types.NewBool(false),
						AverageBandwidth: 0,
						PeakBandwidth:    0,
						BurstSize:        0,
					},
				},
				Spec: types.HostPortGroupSpec{
					Name:        "VM Network",
					VlanId:      0,
					VswitchName: "vSwitch0",
					Policy: types.HostNetworkPolicy{
						Security: &types.HostNetworkSecurityPolicy{},
						NicTeaming: &types.HostNicTeamingPolicy{
							Policy:          "",
							ReversePolicy:   (*bool)(nil),
							NotifySwitches:  (*bool)(nil),
							RollingOrder:    (*bool)(nil),
							FailureCriteria: &types.HostNicFailureCriteria{},
							NicOrder:        (*types.HostNicOrderPolicy)(nil),
						},
						OffloadPolicy: &types.HostNetOffloadCapabilities{},
						ShapingPolicy: &types.HostNetworkTrafficShapingPolicy{},
					},
				},
			},
			{
				Key: "key-vim.host.PortGroup-Management Network",
				Port: []types.HostPortGroupPort{
					{
						Key:  "key-vim.host.PortGroup.Port-33554436",
						Mac:  []string{"00:0c:29:81:d8:a0"},
						Type: "host",
					},
				},
				Vswitch: "key-vim.host.VirtualSwitch-vSwitch0",
				ComputedPolicy: types.HostNetworkPolicy{
					Security: &types.HostNetworkSecurityPolicy{
						AllowPromiscuous: types.NewBool(false),
						MacChanges:       types.NewBool(true),
						ForgedTransmits:  types.NewBool(true),
					},
					NicTeaming: &types.HostNicTeamingPolicy{
						Policy:         "loadbalance_srcid",
						ReversePolicy:  types.NewBool(true),
						NotifySwitches: types.NewBool(true),
						RollingOrder:   types.NewBool(false),
						FailureCriteria: &types.HostNicFailureCriteria{
							CheckSpeed:        "minimum",
							Speed:             10,
							CheckDuplex:       types.NewBool(false),
							FullDuplex:        types.NewBool(false),
							CheckErrorPercent: types.NewBool(false),
							Percentage:        0,
							CheckBeacon:       types.NewBool(false),
						},
						NicOrder: &types.HostNicOrderPolicy{
							ActiveNic:  []string{"vmnic0"},
							StandbyNic: nil,
						},
					},
					OffloadPolicy: &types.HostNetOffloadCapabilities{
						CsumOffload:     types.NewBool(true),
						TcpSegmentation: types.NewBool(true),
						ZeroCopyXmit:    types.NewBool(true),
					},
					ShapingPolicy: &types.HostNetworkTrafficShapingPolicy{
						Enabled:          types.NewBool(false),
						AverageBandwidth: 0,
						PeakBandwidth:    0,
						BurstSize:        0,
					},
				},
				Spec: types.HostPortGroupSpec{
					Name:        "Management Network",
					VlanId:      0,
					VswitchName: "vSwitch0",
					Policy: types.HostNetworkPolicy{
						Security: &types.HostNetworkSecurityPolicy{},
						NicTeaming: &types.HostNicTeamingPolicy{
							Policy:         "loadbalance_srcid",
							ReversePolicy:  (*bool)(nil),
							NotifySwitches: types.NewBool(true),
							RollingOrder:   types.NewBool(false),
							FailureCriteria: &types.HostNicFailureCriteria{
								CheckSpeed:        "",
								Speed:             0,
								CheckDuplex:       (*bool)(nil),
								FullDuplex:        (*bool)(nil),
								CheckErrorPercent: (*bool)(nil),
								Percentage:        0,
								CheckBeacon:       types.NewBool(false),
							},
							NicOrder: &types.HostNicOrderPolicy{
								ActiveNic:  []string{"vmnic0"},
								StandbyNic: nil,
							},
						},
						OffloadPolicy: &types.HostNetOffloadCapabilities{},
						ShapingPolicy: &types.HostNetworkTrafficShapingPolicy{},
					},
				},
			},
		},
		Pnic: []types.PhysicalNic{
			{
				Key:    "key-vim.host.PhysicalNic-vmnic0",
				Device: "vmnic0",
				Pci:    "0000:0b:00.0",
				Driver: "nvmxnet3",
				LinkSpeed: &types.PhysicalNicLinkInfo{
					SpeedMb: 10000,
					Duplex:  true,
				},
				ValidLinkSpecification: []types.PhysicalNicLinkInfo{
					{
						SpeedMb: 10000,
						Duplex:  true,
					},
				},
				Spec: types.PhysicalNicSpec{
					Ip: &types.HostIpConfig{},
					LinkSpeed: &types.PhysicalNicLinkInfo{
						SpeedMb: 10000,
						Duplex:  true,
					},
				},
				WakeOnLanSupported: false,
				Mac:                "00:0c:29:81:d8:a0",
				FcoeConfiguration: &types.FcoeConfig{
					PriorityClass: 3,
					SourceMac:     "00:0c:29:81:d8:a0",
					VlanRange: []types.FcoeConfigVlanRange{
						{},
					},
					Capabilities: types.FcoeConfigFcoeCapabilities{
						PriorityClass:    false,
						SourceMacAddress: false,
						VlanRange:        true,
					},
					FcoeActive: false,
				},
				VmDirectPathGen2Supported:             types.NewBool(false),
				VmDirectPathGen2SupportedMode:         "",
				ResourcePoolSchedulerAllowed:          types.NewBool(true),
				ResourcePoolSchedulerDisallowedReason: nil,
				AutoNegotiateSupported:                types.NewBool(false),
			},
			{
				Key:    "key-vim.host.PhysicalNic-vmnic1",
				Device: "vmnic1",
				Pci:    "0000:13:00.0",
				Driver: "nvmxnet3",
				LinkSpeed: &types.PhysicalNicLinkInfo{
					SpeedMb: 10000,
					Duplex:  true,
				},
				ValidLinkSpecification: []types.PhysicalNicLinkInfo{
					{
						SpeedMb: 10000,
						Duplex:  true,
					},
				},
				Spec: types.PhysicalNicSpec{
					Ip: &types.HostIpConfig{},
					LinkSpeed: &types.PhysicalNicLinkInfo{
						SpeedMb: 10000,
						Duplex:  true,
					},
				},
				WakeOnLanSupported: false,
				Mac:                "00:0c:29:81:d8:aa",
				FcoeConfiguration: &types.FcoeConfig{
					PriorityClass: 3,
					SourceMac:     "00:0c:29:81:d8:aa",
					VlanRange: []types.FcoeConfigVlanRange{
						{},
					},
					Capabilities: types.FcoeConfigFcoeCapabilities{
						PriorityClass:    false,
						SourceMacAddress: false,
						VlanRange:        true,
					},
					FcoeActive: false,
				},
				VmDirectPathGen2Supported:             types.NewBool(false),
				VmDirectPathGen2SupportedMode:         "",
				ResourcePoolSchedulerAllowed:          types.NewBool(true),
				ResourcePoolSchedulerDisallowedReason: nil,
				AutoNegotiateSupported:                types.NewBool(false),
			},
		},
		Vnic: []types.HostVirtualNic{
			{
				Device:    "vmk0",
				Key:       "key-vim.host.VirtualNic-vmk0",
				Portgroup: "Management Network",
				Spec: types.HostVirtualNicSpec{
					Ip: &types.HostIpConfig{
						Dhcp:       true,
						IpAddress:  "127.0.0.1",
						SubnetMask: "255.0.0.0",
						IpV6Config: (*types.HostIpConfigIpV6AddressConfiguration)(nil),
					},
					Mac:                    "00:0c:29:81:d8:a0",
					DistributedVirtualPort: (*types.DistributedVirtualSwitchPortConnection)(nil),
					Portgroup:              "Management Network",
					Mtu:                    1500,
					TsoEnabled:             types.NewBool(true),
					NetStackInstanceKey:    "defaultTcpipStack",
					OpaqueNetwork:          (*types.HostVirtualNicOpaqueNetworkSpec)(nil),
					ExternalId:             "",
					PinnedPnic:             "",
					IpRouteSpec:            (*types.HostVirtualNicIpRouteSpec)(nil),
				},
				Port: "key-vim.host.PortGroup.Port-33554436",
			},
		},
		ConsoleVnic: nil,
		DnsConfig: &types.HostDnsConfig{
			Dhcp:             true,
			VirtualNicDevice: "vmk0",
			HostName:         "localhost",
			DomainName:       "localdomain",
			Address:          []string{"8.8.8.8"},
			SearchDomain:     []string{"localdomain"},
		},
		IpRouteConfig: &types.HostIpRouteConfig{
			DefaultGateway:     "127.0.0.1",
			GatewayDevice:      "",
			IpV6DefaultGateway: "",
			IpV6GatewayDevice:  "",
		},
		ConsoleIpRouteConfig: nil,
		RouteTableInfo: &types.HostIpRouteTableInfo{
			IpRoute: []types.HostIpRouteEntry{
				{
					Network:      "0.0.0.0",
					PrefixLength: 0,
					Gateway:      "127.0.0.1",
					DeviceName:   "vmk0",
				},
				{
					Network:      "127.0.0.0",
					PrefixLength: 8,
					Gateway:      "0.0.0.0",
					DeviceName:   "vmk0",
				},
			},
			Ipv6Route: nil,
		},
		Dhcp:              nil,
		Nat:               nil,
		IpV6Enabled:       types.NewBool(false),
		AtBootIpV6Enabled: types.NewBool(false),
		NetStackInstance: []types.HostNetStackInstance{
			{
				Key:                             "vSphereProvisioning",
				Name:                            "",
				DnsConfig:                       &types.HostDnsConfig{},
				IpRouteConfig:                   &types.HostIpRouteConfig{},
				RequestedMaxNumberOfConnections: 11000,
				CongestionControlAlgorithm:      "newreno",
				IpV6Enabled:                     types.NewBool(true),
				RouteTableConfig:                (*types.HostIpRouteTableConfig)(nil),
			},
			{
				Key:                             "vmotion",
				Name:                            "",
				DnsConfig:                       &types.HostDnsConfig{},
				IpRouteConfig:                   &types.HostIpRouteConfig{},
				RequestedMaxNumberOfConnections: 11000,
				CongestionControlAlgorithm:      "newreno",
				IpV6Enabled:                     types.NewBool(true),
				RouteTableConfig:                (*types.HostIpRouteTableConfig)(nil),
			},
			{
				Key:  "defaultTcpipStack",
				Name: "defaultTcpipStack",
				DnsConfig: &types.HostDnsConfig{
					Dhcp:             true,
					VirtualNicDevice: "vmk0",
					HostName:         "localhost",
					DomainName:       "localdomain",
					Address:          []string{"8.8.8.8"},
					SearchDomain:     []string{"localdomain"},
				},
				IpRouteConfig: &types.HostIpRouteConfig{
					DefaultGateway:     "127.0.0.1",
					GatewayDevice:      "",
					IpV6DefaultGateway: "",
					IpV6GatewayDevice:  "",
				},
				RequestedMaxNumberOfConnections: 11000,
				CongestionControlAlgorithm:      "newreno",
				IpV6Enabled:                     types.NewBool(true),
				RouteTableConfig: &types.HostIpRouteTableConfig{
					IpRoute: []types.HostIpRouteOp{
						{
							ChangeOperation: "ignore",
							Route: types.HostIpRouteEntry{
								Network:      "0.0.0.0",
								PrefixLength: 0,
								Gateway:      "127.0.0.1",
								DeviceName:   "vmk0",
							},
						},
						{
							ChangeOperation: "ignore",
							Route: types.HostIpRouteEntry{
								Network:      "127.0.0.0",
								PrefixLength: 8,
								Gateway:      "0.0.0.0",
								DeviceName:   "vmk0",
							},
						},
					},
					Ipv6Route: nil,
				},
			},
		},
		OpaqueSwitch:  nil,
		OpaqueNetwork: nil,
	},
	Vmotion: &types.HostVMotionInfo{
		NetConfig: &types.HostVMotionNetConfig{
			CandidateVnic: []types.HostVirtualNic{
				{
					Device:    "vmk0",
					Key:       "VMotionConfig.vmotion.key-vim.host.VirtualNic-vmk0",
					Portgroup: "Management Network",
					Spec: types.HostVirtualNicSpec{
						Ip: &types.HostIpConfig{
							Dhcp:       true,
							IpAddress:  "127.0.0.1",
							SubnetMask: "255.0.0.0",
							IpV6Config: (*types.HostIpConfigIpV6AddressConfiguration)(nil),
						},
						Mac:                    "00:0c:29:81:d8:a0",
						DistributedVirtualPort: (*types.DistributedVirtualSwitchPortConnection)(nil),
						Portgroup:              "Management Network",
						Mtu:                    1500,
						TsoEnabled:             types.NewBool(true),
						NetStackInstanceKey:    "defaultTcpipStack",
						OpaqueNetwork:          (*types.HostVirtualNicOpaqueNetworkSpec)(nil),
						ExternalId:             "",
						PinnedPnic:             "",
						IpRouteSpec:            (*types.HostVirtualNicIpRouteSpec)(nil),
					},
					Port: "",
				},
			},
			SelectedVnic: "",
		},
		IpConfig: (*types.HostIpConfig)(nil),
	},
	VirtualNicManagerInfo: &types.HostVirtualNicManagerInfo{
		NetConfig: []types.VirtualNicManagerNetConfig{
			{
				NicType:            "faultToleranceLogging",
				MultiSelectAllowed: true,
				CandidateVnic: []types.HostVirtualNic{
					{
						Device:    "vmk0",
						Key:       "faultToleranceLogging.key-vim.host.VirtualNic-vmk0",
						Portgroup: "Management Network",
						Spec: types.HostVirtualNicSpec{
							Ip: &types.HostIpConfig{
								Dhcp:       true,
								IpAddress:  "127.0.0.1",
								SubnetMask: "255.0.0.0",
								IpV6Config: (*types.HostIpConfigIpV6AddressConfiguration)(nil),
							},
							Mac:                    "00:0c:29:81:d8:a0",
							DistributedVirtualPort: (*types.DistributedVirtualSwitchPortConnection)(nil),
							Portgroup:              "Management Network",
							Mtu:                    1500,
							TsoEnabled:             types.NewBool(true),
							NetStackInstanceKey:    "defaultTcpipStack",
							OpaqueNetwork:          (*types.HostVirtualNicOpaqueNetworkSpec)(nil),
							ExternalId:             "",
							PinnedPnic:             "",
							IpRouteSpec:            (*types.HostVirtualNicIpRouteSpec)(nil),
						},
						Port: "",
					},
				},
				SelectedVnic: nil,
			},
			{
				NicType:            "management",
				MultiSelectAllowed: true,
				CandidateVnic: []types.HostVirtualNic{
					{
						Device:    "vmk1",
						Key:       "management.key-vim.host.VirtualNic-vmk1",
						Portgroup: "",
						Spec: types.HostVirtualNicSpec{
							Ip: &types.HostIpConfig{
								Dhcp:       true,
								IpAddress:  "192.168.0.1",
								SubnetMask: "255.0.0.0",
								IpV6Config: (*types.HostIpConfigIpV6AddressConfiguration)(nil),
							},
							Mac:                    "00:0c:29:81:d8:00",
							DistributedVirtualPort: (*types.DistributedVirtualSwitchPortConnection)(nil),
							Portgroup:              "Management Network",
							Mtu:                    1500,
							TsoEnabled:             types.NewBool(true),
							NetStackInstanceKey:    "defaultTcpipStack",
							OpaqueNetwork:          (*types.HostVirtualNicOpaqueNetworkSpec)(nil),
							ExternalId:             "",
							PinnedPnic:             "",
							IpRouteSpec:            (*types.HostVirtualNicIpRouteSpec)(nil),
						},
						Port: "",
					},
					{
						Device:    "vmk0",
						Key:       "management.key-vim.host.VirtualNic-vmk0",
						Portgroup: "Management Network",
						Spec: types.HostVirtualNicSpec{
							Ip: &types.HostIpConfig{
								Dhcp:       true,
								IpAddress:  "127.0.0.1",
								SubnetMask: "255.0.0.0",
								IpV6Config: (*types.HostIpConfigIpV6AddressConfiguration)(nil),
							},
							Mac:                    "00:0c:29:81:d8:a0",
							DistributedVirtualPort: (*types.DistributedVirtualSwitchPortConnection)(nil),
							Portgroup:              "Management Network",
							Mtu:                    1500,
							TsoEnabled:             types.NewBool(true),
							NetStackInstanceKey:    "defaultTcpipStack",
							OpaqueNetwork:          (*types.HostVirtualNicOpaqueNetworkSpec)(nil),
							ExternalId:             "",
							PinnedPnic:             "",
							IpRouteSpec:            (*types.HostVirtualNicIpRouteSpec)(nil),
						},
						Port: "",
					},
				},
				SelectedVnic: []string{"management.key-vim.host.VirtualNic-vmk0"},
			},
			{
				NicType:            "vSphereProvisioning",
				MultiSelectAllowed: true,
				CandidateVnic: []types.HostVirtualNic{
					{
						Device:    "vmk0",
						Key:       "vSphereProvisioning.key-vim.host.VirtualNic-vmk0",
						Portgroup: "Management Network",
						Spec: types.HostVirtualNicSpec{
							Ip: &types.HostIpConfig{
								Dhcp:       true,
								IpAddress:  "127.0.0.1",
								SubnetMask: "255.0.0.0",
								IpV6Config: (*types.HostIpConfigIpV6AddressConfiguration)(nil),
							},
							Mac:                    "00:0c:29:81:d8:a0",
							DistributedVirtualPort: (*types.DistributedVirtualSwitchPortConnection)(nil),
							Portgroup:              "Management Network",
							Mtu:                    1500,
							TsoEnabled:             types.NewBool(true),
							NetStackInstanceKey:    "defaultTcpipStack",
							OpaqueNetwork:          (*types.HostVirtualNicOpaqueNetworkSpec)(nil),
							ExternalId:             "",
							PinnedPnic:             "",
							IpRouteSpec:            (*types.HostVirtualNicIpRouteSpec)(nil),
						},
						Port: "",
					},
				},
				SelectedVnic: nil,
			},
			{
				NicType:            "vSphereReplication",
				MultiSelectAllowed: true,
				CandidateVnic: []types.HostVirtualNic{
					{
						Device:    "vmk0",
						Key:       "vSphereReplication.key-vim.host.VirtualNic-vmk0",
						Portgroup: "Management Network",
						Spec: types.HostVirtualNicSpec{
							Ip: &types.HostIpConfig{
								Dhcp:       true,
								IpAddress:  "127.0.0.1",
								SubnetMask: "255.0.0.0",
								IpV6Config: (*types.HostIpConfigIpV6AddressConfiguration)(nil),
							},
							Mac:                    "00:0c:29:81:d8:a0",
							DistributedVirtualPort: (*types.DistributedVirtualSwitchPortConnection)(nil),
							Portgroup:              "Management Network",
							Mtu:                    1500,
							TsoEnabled:             types.NewBool(true),
							NetStackInstanceKey:    "defaultTcpipStack",
							OpaqueNetwork:          (*types.HostVirtualNicOpaqueNetworkSpec)(nil),
							ExternalId:             "",
							PinnedPnic:             "",
							IpRouteSpec:            (*types.HostVirtualNicIpRouteSpec)(nil),
						},
						Port: "",
					},
				},
				SelectedVnic: nil,
			},
			{
				NicType:            "vSphereReplicationNFC",
				MultiSelectAllowed: true,
				CandidateVnic: []types.HostVirtualNic{
					{
						Device:    "vmk0",
						Key:       "vSphereReplicationNFC.key-vim.host.VirtualNic-vmk0",
						Portgroup: "Management Network",
						Spec: types.HostVirtualNicSpec{
							Ip: &types.HostIpConfig{
								Dhcp:       true,
								IpAddress:  "127.0.0.1",
								SubnetMask: "255.0.0.0",
								IpV6Config: (*types.HostIpConfigIpV6AddressConfiguration)(nil),
							},
							Mac:                    "00:0c:29:81:d8:a0",
							DistributedVirtualPort: (*types.DistributedVirtualSwitchPortConnection)(nil),
							Portgroup:              "Management Network",
							Mtu:                    1500,
							TsoEnabled:             types.NewBool(true),
							NetStackInstanceKey:    "defaultTcpipStack",
							OpaqueNetwork:          (*types.HostVirtualNicOpaqueNetworkSpec)(nil),
							ExternalId:             "",
							PinnedPnic:             "",
							IpRouteSpec:            (*types.HostVirtualNicIpRouteSpec)(nil),
						},
						Port: "",
					},
				},
				SelectedVnic: nil,
			},
			{
				NicType:            "vmotion",
				MultiSelectAllowed: true,
				CandidateVnic: []types.HostVirtualNic{
					{
						Device:    "vmk0",
						Key:       "vmotion.key-vim.host.VirtualNic-vmk0",
						Portgroup: "Management Network",
						Spec: types.HostVirtualNicSpec{
							Ip: &types.HostIpConfig{
								Dhcp:       true,
								IpAddress:  "127.0.0.1",
								SubnetMask: "255.0.0.0",
								IpV6Config: (*types.HostIpConfigIpV6AddressConfiguration)(nil),
							},
							Mac:                    "00:0c:29:81:d8:a0",
							DistributedVirtualPort: (*types.DistributedVirtualSwitchPortConnection)(nil),
							Portgroup:              "Management Network",
							Mtu:                    1500,
							TsoEnabled:             types.NewBool(true),
							NetStackInstanceKey:    "defaultTcpipStack",
							OpaqueNetwork:          (*types.HostVirtualNicOpaqueNetworkSpec)(nil),
							ExternalId:             "",
							PinnedPnic:             "",
							IpRouteSpec:            (*types.HostVirtualNicIpRouteSpec)(nil),
						},
						Port: "",
					},
				},
				SelectedVnic: nil,
			},
			{
				NicType:            "vsan",
				MultiSelectAllowed: true,
				CandidateVnic: []types.HostVirtualNic{
					{
						Device:    "vmk0",
						Key:       "vsan.key-vim.host.VirtualNic-vmk0",
						Portgroup: "Management Network",
						Spec: types.HostVirtualNicSpec{
							Ip: &types.HostIpConfig{
								Dhcp:       true,
								IpAddress:  "127.0.0.1",
								SubnetMask: "255.0.0.0",
								IpV6Config: (*types.HostIpConfigIpV6AddressConfiguration)(nil),
							},
							Mac:                    "00:0c:29:81:d8:a0",
							DistributedVirtualPort: (*types.DistributedVirtualSwitchPortConnection)(nil),
							Portgroup:              "Management Network",
							Mtu:                    1500,
							TsoEnabled:             types.NewBool(true),
							NetStackInstanceKey:    "defaultTcpipStack",
							OpaqueNetwork:          (*types.HostVirtualNicOpaqueNetworkSpec)(nil),
							ExternalId:             "",
							PinnedPnic:             "",
							IpRouteSpec:            (*types.HostVirtualNicIpRouteSpec)(nil),
						},
						Port: "",
					},
				},
				SelectedVnic: nil,
			},
			{
				NicType:            "vsanWitness",
				MultiSelectAllowed: true,
				CandidateVnic: []types.HostVirtualNic{
					{
						Device:    "vmk0",
						Key:       "vsanWitness.key-vim.host.VirtualNic-vmk0",
						Portgroup: "Management Network",
						Spec: types.HostVirtualNicSpec{
							Ip: &types.HostIpConfig{
								Dhcp:       true,
								IpAddress:  "127.0.0.1",
								SubnetMask: "255.0.0.0",
								IpV6Config: (*types.HostIpConfigIpV6AddressConfiguration)(nil),
							},
							Mac:                    "00:0c:29:81:d8:a0",
							DistributedVirtualPort: (*types.DistributedVirtualSwitchPortConnection)(nil),
							Portgroup:              "Management Network",
							Mtu:                    1500,
							TsoEnabled:             types.NewBool(true),
							NetStackInstanceKey:    "defaultTcpipStack",
							OpaqueNetwork:          (*types.HostVirtualNicOpaqueNetworkSpec)(nil),
							ExternalId:             "",
							PinnedPnic:             "",
							IpRouteSpec:            (*types.HostVirtualNicIpRouteSpec)(nil),
						},
						Port: "",
					},
				},
				SelectedVnic: nil,
			},
		},
	},
	Capabilities: &types.HostNetCapabilities{
		CanSetPhysicalNicLinkSpeed: true,
		SupportsNicTeaming:         true,
		NicTeamingPolicy:           []string{"loadbalance_ip", "loadbalance_srcmac", "loadbalance_srcid", "failover_explicit"},
		SupportsVlan:               true,
		UsesServiceConsoleNic:      false,
		SupportsNetworkHints:       true,
		MaxPortGroupsPerVswitch:    0,
		VswitchConfigSupported:     true,
		VnicConfigSupported:        true,
		IpRouteConfigSupported:     true,
		DnsConfigSupported:         true,
		DhcpOnVnicSupported:        true,
		IpV6Supported:              types.NewBool(true),
	},
	DatastoreCapabilities: &types.HostDatastoreSystemCapabilities{
		NfsMountCreationRequired:     true,
		NfsMountCreationSupported:    true,
		LocalDatastoreSupported:      false,
		VmfsExtentExpansionSupported: types.NewBool(true),
	},
	OffloadCapabilities: &types.HostNetOffloadCapabilities{
		CsumOffload:     types.NewBool(true),
		TcpSegmentation: types.NewBool(true),
		ZeroCopyXmit:    types.NewBool(true),
	},
	Service: &types.HostServiceInfo{
		Service: []types.HostService{
			{
				Key:           "DCUI",
				Label:         "Direct Console UI",
				Required:      false,
				Uninstallable: false,
				Running:       true,
				Ruleset:       nil,
				Policy:        "on",
				SourcePackage: &types.HostServiceSourcePackage{
					SourcePackageName: "esx-base",
					Description:       "This VIB contains all of the base functionality of vSphere ESXi.",
				},
			},
			{
				Key:           "TSM",
				Label:         "ESXi Shell",
				Required:      false,
				Uninstallable: false,
				Running:       false,
				Ruleset:       nil,
				Policy:        "off",
				SourcePackage: &types.HostServiceSourcePackage{
					SourcePackageName: "esx-base",
					Description:       "This VIB contains all of the base functionality of vSphere ESXi.",
				},
			},
			{
				Key:           "TSM-SSH",
				Label:         "SSH",
				Required:      false,
				Uninstallable: false,
				Running:       false,
				Ruleset:       nil,
				Policy:        "off",
				SourcePackage: &types.HostServiceSourcePackage{
					SourcePackageName: "esx-base",
					Description:       "This VIB contains all of the base functionality of vSphere ESXi.",
				},
			},
			{
				Key:           "lbtd",
				Label:         "Load-Based Teaming Daemon",
				Required:      false,
				Uninstallable: false,
				Running:       true,
				Ruleset:       nil,
				Policy:        "on",
				SourcePackage: &types.HostServiceSourcePackage{
					SourcePackageName: "esx-base",
					Description:       "This VIB contains all of the base functionality of vSphere ESXi.",
				},
			},
			{
				Key:           "lwsmd",
				Label:         "Active Directory Service",
				Required:      false,
				Uninstallable: false,
				Running:       false,
				Ruleset:       nil,
				Policy:        "off",
				SourcePackage: &types.HostServiceSourcePackage{
					SourcePackageName: "esx-base",
					Description:       "This VIB contains all of the base functionality of vSphere ESXi.",
				},
			},
			{
				Key:           "ntpd",
				Label:         "NTP Daemon",
				Required:      false,
				Uninstallable: false,
				Running:       false,
				Ruleset:       []string{"ntpClient"},
				Policy:        "off",
				SourcePackage: &types.HostServiceSourcePackage{
					SourcePackageName: "esx-base",
					Description:       "This VIB contains all of the base functionality of vSphere ESXi.",
				},
			},
			{
				Key:           "pcscd",
				Label:         "PC/SC Smart Card Daemon",
				Required:      false,
				Uninstallable: false,
				Running:       false,
				Ruleset:       nil,
				Policy:        "off",
				SourcePackage: &types.HostServiceSourcePackage{
					SourcePackageName: "esx-base",
					Description:       "This VIB contains all of the base functionality of vSphere ESXi.",
				},
			},
			{
				Key:           "sfcbd-watchdog",
				Label:         "CIM Server",
				Required:      false,
				Uninstallable: false,
				Running:       false,
				Ruleset:       []string{"CIMHttpServer", "CIMHttpsServer"},
				Policy:        "on",
				SourcePackage: &types.HostServiceSourcePackage{
					SourcePackageName: "esx-base",
					Description:       "This VIB contains all of the base functionality of vSphere ESXi.",
				},
			},
			{
				Key:           "snmpd",
				Label:         "SNMP Server",
				Required:      false,
				Uninstallable: false,
				Running:       false,
				Ruleset:       []string{"snmp"},
				Policy:        "on",
				SourcePackage: &types.HostServiceSourcePackage{
					SourcePackageName: "esx-base",
					Description:       "This VIB contains all of the base functionality of vSphere ESXi.",
				},
			},
			{
				Key:           "vmsyslogd",
				Label:         "Syslog Server",
				Required:      true,
				Uninstallable: false,
				Running:       true,
				Ruleset:       nil,
				Policy:        "on",
				SourcePackage: &types.HostServiceSourcePackage{
					SourcePackageName: "esx-base",
					Description:       "This VIB contains all of the base functionality of vSphere ESXi.",
				},
			},
			{
				Key:           "vpxa",
				Label:         "VMware vCenter Agent",
				Required:      false,
				Uninstallable: false,
				Running:       false,
				Ruleset:       []string{"vpxHeartbeats"},
				Policy:        "on",
				SourcePackage: &types.HostServiceSourcePackage{
					SourcePackageName: "esx-base",
					Description:       "This VIB contains all of the base functionality of vSphere ESXi.",
				},
			},
			{
				Key:           "xorg",
				Label:         "X.Org Server",
				Required:      false,
				Uninstallable: false,
				Running:       false,
				Ruleset:       nil,
				Policy:        "on",
				SourcePackage: &types.HostServiceSourcePackage{
					SourcePackageName: "esx-xserver",
					Description:       "This VIB contains X Server used for virtual machine 3D hardware acceleration.",
				},
			},
		},
	},
	Firewall: &HostFirewallInfo,
	AutoStart: &types.HostAutoStartManagerConfig{
		Defaults: &types.AutoStartDefaults{
			Enabled:          (*bool)(nil),
			StartDelay:       120,
			StopDelay:        120,
			WaitForHeartbeat: types.NewBool(false),
			StopAction:       "PowerOff",
		},
		PowerInfo: nil,
	},
	ActiveDiagnosticPartition: &types.HostDiagnosticPartition{
		StorageType:    "directAttached",
		DiagnosticType: "singleHost",
		Slots:          -15,
		Id: types.HostScsiDiskPartition{
			DiskName:  "mpx.vmhba0:C0:T0:L0",
			Partition: 9,
		},
	},
	Option:            nil,
	OptionDef:         nil,
	Flags:             &types.HostFlagInfo{},
	AdminDisabled:     (*bool)(nil),
	LockdownMode:      "lockdownDisabled",
	Ipmi:              (*types.HostIpmiInfo)(nil),
	SslThumbprintInfo: (*types.HostSslThumbprintInfo)(nil),
	SslThumbprintData: nil,
	Certificate:       []uint8{0x31, 0x30},
	PciPassthruInfo:   nil,
	AuthenticationManagerInfo: &types.HostAuthenticationManagerInfo{
		AuthConfig: []types.BaseHostAuthenticationStoreInfo{
			&types.HostLocalAuthenticationInfo{
				HostAuthenticationStoreInfo: types.HostAuthenticationStoreInfo{
					Enabled: true,
				},
			},
			&types.HostActiveDirectoryInfo{
				HostDirectoryStoreInfo:         types.HostDirectoryStoreInfo{},
				JoinedDomain:                   "",
				TrustedDomain:                  nil,
				DomainMembershipStatus:         "",
				SmartCardAuthenticationEnabled: types.NewBool(false),
			},
		},
	},
	FeatureVersion: nil,
	PowerSystemCapability: &types.PowerSystemCapability{
		AvailablePolicy: []types.HostPowerPolicy{
			{
				Key:         1,
				Name:        "PowerPolicy.static.name",
				ShortName:   "static",
				Description: "PowerPolicy.static.description",
			},
			{
				Key:         2,
				Name:        "PowerPolicy.dynamic.name",
				ShortName:   "dynamic",
				Description: "PowerPolicy.dynamic.description",
			},
			{
				Key:         3,
				Name:        "PowerPolicy.low.name",
				ShortName:   "low",
				Description: "PowerPolicy.low.description",
			},
			{
				Key:         4,
				Name:        "PowerPolicy.custom.name",
				ShortName:   "custom",
				Description: "PowerPolicy.custom.description",
			},
		},
	},
	PowerSystemInfo: &types.PowerSystemInfo{
		CurrentPolicy: types.HostPowerPolicy{
			Key:         2,
			Name:        "PowerPolicy.dynamic.name",
			ShortName:   "dynamic",
			Description: "PowerPolicy.dynamic.description",
		},
	},
	CacheConfigurationInfo: []types.HostCacheConfigurationInfo{
		{
			Key:      types.ManagedObjectReference{Type: "Datastore", Value: "5980f676-21a5db76-9eef-000c2981d8a0"},
			SwapSize: 0,
		},
	},
	WakeOnLanCapable:        types.NewBool(false),
	FeatureCapability:       nil,
	MaskedFeatureCapability: nil,
	VFlashConfigInfo:        nil,
	VsanHostConfig: &types.VsanHostConfigInfo{
		Enabled:     types.NewBool(false),
		HostSystem:  &types.ManagedObjectReference{Type: "HostSystem", Value: "ha-host"},
		ClusterInfo: &types.VsanHostConfigInfoClusterInfo{},
		StorageInfo: &types.VsanHostConfigInfoStorageInfo{
			AutoClaimStorage: types.NewBool(false),
			DiskMapping:      nil,
			DiskMapInfo:      nil,
			ChecksumEnabled:  (*bool)(nil),
		},
		NetworkInfo:     &types.VsanHostConfigInfoNetworkInfo{},
		FaultDomainInfo: &types.VsanHostFaultDomainInfo{},
	},
	DomainList:             nil,
	ScriptCheckSum:         nil,
	HostConfigCheckSum:     nil,
	GraphicsInfo:           nil,
	SharedPassthruGpuTypes: nil,
	GraphicsConfig: &types.HostGraphicsConfig{
		HostDefaultGraphicsType:        "shared",
		SharedPassthruAssignmentPolicy: "performance",
		DeviceType:                     nil,
	},
	IoFilterInfo: []types.HostIoFilterInfo{
		{
			IoFilterInfo: types.IoFilterInfo{
				Id:          "VMW_spm_1.0.0",
				Name:        "spm",
				Vendor:      "VMW",
				Version:     "1.0.230",
				Type:        "datastoreIoControl",
				Summary:     "VMware Storage I/O Control",
				ReleaseDate: "2016-07-21",
			},
			Available: true,
		},
		{
			IoFilterInfo: types.IoFilterInfo{
				Id:          "VMW_vmwarevmcrypt_1.0.0",
				Name:        "vmwarevmcrypt",
				Vendor:      "VMW",
				Version:     "1.0.0",
				Type:        "encryption",
				Summary:     "VMcrypt IO Filter",
				ReleaseDate: "2016-07-21",
			},
			Available: true,
		},
	},
	SriovDevicePool: nil,
}
