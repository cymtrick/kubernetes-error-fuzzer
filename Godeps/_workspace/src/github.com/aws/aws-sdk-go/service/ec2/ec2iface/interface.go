// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package ec2iface provides an interface for the Amazon Elastic Compute Cloud.
package ec2iface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// EC2API is the interface type for ec2.EC2.
type EC2API interface {
	AcceptVpcPeeringConnectionRequest(*ec2.AcceptVpcPeeringConnectionInput) (*request.Request, *ec2.AcceptVpcPeeringConnectionOutput)

	AcceptVpcPeeringConnection(*ec2.AcceptVpcPeeringConnectionInput) (*ec2.AcceptVpcPeeringConnectionOutput, error)

	AllocateAddressRequest(*ec2.AllocateAddressInput) (*request.Request, *ec2.AllocateAddressOutput)

	AllocateAddress(*ec2.AllocateAddressInput) (*ec2.AllocateAddressOutput, error)

	AssignPrivateIpAddressesRequest(*ec2.AssignPrivateIpAddressesInput) (*request.Request, *ec2.AssignPrivateIpAddressesOutput)

	AssignPrivateIpAddresses(*ec2.AssignPrivateIpAddressesInput) (*ec2.AssignPrivateIpAddressesOutput, error)

	AssociateAddressRequest(*ec2.AssociateAddressInput) (*request.Request, *ec2.AssociateAddressOutput)

	AssociateAddress(*ec2.AssociateAddressInput) (*ec2.AssociateAddressOutput, error)

	AssociateDhcpOptionsRequest(*ec2.AssociateDhcpOptionsInput) (*request.Request, *ec2.AssociateDhcpOptionsOutput)

	AssociateDhcpOptions(*ec2.AssociateDhcpOptionsInput) (*ec2.AssociateDhcpOptionsOutput, error)

	AssociateRouteTableRequest(*ec2.AssociateRouteTableInput) (*request.Request, *ec2.AssociateRouteTableOutput)

	AssociateRouteTable(*ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error)

	AttachClassicLinkVpcRequest(*ec2.AttachClassicLinkVpcInput) (*request.Request, *ec2.AttachClassicLinkVpcOutput)

	AttachClassicLinkVpc(*ec2.AttachClassicLinkVpcInput) (*ec2.AttachClassicLinkVpcOutput, error)

	AttachInternetGatewayRequest(*ec2.AttachInternetGatewayInput) (*request.Request, *ec2.AttachInternetGatewayOutput)

	AttachInternetGateway(*ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error)

	AttachNetworkInterfaceRequest(*ec2.AttachNetworkInterfaceInput) (*request.Request, *ec2.AttachNetworkInterfaceOutput)

	AttachNetworkInterface(*ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error)

	AttachVolumeRequest(*ec2.AttachVolumeInput) (*request.Request, *ec2.VolumeAttachment)

	AttachVolume(*ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error)

	AttachVpnGatewayRequest(*ec2.AttachVpnGatewayInput) (*request.Request, *ec2.AttachVpnGatewayOutput)

	AttachVpnGateway(*ec2.AttachVpnGatewayInput) (*ec2.AttachVpnGatewayOutput, error)

	AuthorizeSecurityGroupEgressRequest(*ec2.AuthorizeSecurityGroupEgressInput) (*request.Request, *ec2.AuthorizeSecurityGroupEgressOutput)

	AuthorizeSecurityGroupEgress(*ec2.AuthorizeSecurityGroupEgressInput) (*ec2.AuthorizeSecurityGroupEgressOutput, error)

	AuthorizeSecurityGroupIngressRequest(*ec2.AuthorizeSecurityGroupIngressInput) (*request.Request, *ec2.AuthorizeSecurityGroupIngressOutput)

	AuthorizeSecurityGroupIngress(*ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error)

	BundleInstanceRequest(*ec2.BundleInstanceInput) (*request.Request, *ec2.BundleInstanceOutput)

	BundleInstance(*ec2.BundleInstanceInput) (*ec2.BundleInstanceOutput, error)

	CancelBundleTaskRequest(*ec2.CancelBundleTaskInput) (*request.Request, *ec2.CancelBundleTaskOutput)

	CancelBundleTask(*ec2.CancelBundleTaskInput) (*ec2.CancelBundleTaskOutput, error)

	CancelConversionTaskRequest(*ec2.CancelConversionTaskInput) (*request.Request, *ec2.CancelConversionTaskOutput)

	CancelConversionTask(*ec2.CancelConversionTaskInput) (*ec2.CancelConversionTaskOutput, error)

	CancelExportTaskRequest(*ec2.CancelExportTaskInput) (*request.Request, *ec2.CancelExportTaskOutput)

	CancelExportTask(*ec2.CancelExportTaskInput) (*ec2.CancelExportTaskOutput, error)

	CancelImportTaskRequest(*ec2.CancelImportTaskInput) (*request.Request, *ec2.CancelImportTaskOutput)

	CancelImportTask(*ec2.CancelImportTaskInput) (*ec2.CancelImportTaskOutput, error)

	CancelReservedInstancesListingRequest(*ec2.CancelReservedInstancesListingInput) (*request.Request, *ec2.CancelReservedInstancesListingOutput)

	CancelReservedInstancesListing(*ec2.CancelReservedInstancesListingInput) (*ec2.CancelReservedInstancesListingOutput, error)

	CancelSpotFleetRequestsRequest(*ec2.CancelSpotFleetRequestsInput) (*request.Request, *ec2.CancelSpotFleetRequestsOutput)

	CancelSpotFleetRequests(*ec2.CancelSpotFleetRequestsInput) (*ec2.CancelSpotFleetRequestsOutput, error)

	CancelSpotInstanceRequestsRequest(*ec2.CancelSpotInstanceRequestsInput) (*request.Request, *ec2.CancelSpotInstanceRequestsOutput)

	CancelSpotInstanceRequests(*ec2.CancelSpotInstanceRequestsInput) (*ec2.CancelSpotInstanceRequestsOutput, error)

	ConfirmProductInstanceRequest(*ec2.ConfirmProductInstanceInput) (*request.Request, *ec2.ConfirmProductInstanceOutput)

	ConfirmProductInstance(*ec2.ConfirmProductInstanceInput) (*ec2.ConfirmProductInstanceOutput, error)

	CopyImageRequest(*ec2.CopyImageInput) (*request.Request, *ec2.CopyImageOutput)

	CopyImage(*ec2.CopyImageInput) (*ec2.CopyImageOutput, error)

	CopySnapshotRequest(*ec2.CopySnapshotInput) (*request.Request, *ec2.CopySnapshotOutput)

	CopySnapshot(*ec2.CopySnapshotInput) (*ec2.CopySnapshotOutput, error)

	CreateCustomerGatewayRequest(*ec2.CreateCustomerGatewayInput) (*request.Request, *ec2.CreateCustomerGatewayOutput)

	CreateCustomerGateway(*ec2.CreateCustomerGatewayInput) (*ec2.CreateCustomerGatewayOutput, error)

	CreateDhcpOptionsRequest(*ec2.CreateDhcpOptionsInput) (*request.Request, *ec2.CreateDhcpOptionsOutput)

	CreateDhcpOptions(*ec2.CreateDhcpOptionsInput) (*ec2.CreateDhcpOptionsOutput, error)

	CreateFlowLogsRequest(*ec2.CreateFlowLogsInput) (*request.Request, *ec2.CreateFlowLogsOutput)

	CreateFlowLogs(*ec2.CreateFlowLogsInput) (*ec2.CreateFlowLogsOutput, error)

	CreateImageRequest(*ec2.CreateImageInput) (*request.Request, *ec2.CreateImageOutput)

	CreateImage(*ec2.CreateImageInput) (*ec2.CreateImageOutput, error)

	CreateInstanceExportTaskRequest(*ec2.CreateInstanceExportTaskInput) (*request.Request, *ec2.CreateInstanceExportTaskOutput)

	CreateInstanceExportTask(*ec2.CreateInstanceExportTaskInput) (*ec2.CreateInstanceExportTaskOutput, error)

	CreateInternetGatewayRequest(*ec2.CreateInternetGatewayInput) (*request.Request, *ec2.CreateInternetGatewayOutput)

	CreateInternetGateway(*ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error)

	CreateKeyPairRequest(*ec2.CreateKeyPairInput) (*request.Request, *ec2.CreateKeyPairOutput)

	CreateKeyPair(*ec2.CreateKeyPairInput) (*ec2.CreateKeyPairOutput, error)

	CreateNetworkAclRequest(*ec2.CreateNetworkAclInput) (*request.Request, *ec2.CreateNetworkAclOutput)

	CreateNetworkAcl(*ec2.CreateNetworkAclInput) (*ec2.CreateNetworkAclOutput, error)

	CreateNetworkAclEntryRequest(*ec2.CreateNetworkAclEntryInput) (*request.Request, *ec2.CreateNetworkAclEntryOutput)

	CreateNetworkAclEntry(*ec2.CreateNetworkAclEntryInput) (*ec2.CreateNetworkAclEntryOutput, error)

	CreateNetworkInterfaceRequest(*ec2.CreateNetworkInterfaceInput) (*request.Request, *ec2.CreateNetworkInterfaceOutput)

	CreateNetworkInterface(*ec2.CreateNetworkInterfaceInput) (*ec2.CreateNetworkInterfaceOutput, error)

	CreatePlacementGroupRequest(*ec2.CreatePlacementGroupInput) (*request.Request, *ec2.CreatePlacementGroupOutput)

	CreatePlacementGroup(*ec2.CreatePlacementGroupInput) (*ec2.CreatePlacementGroupOutput, error)

	CreateReservedInstancesListingRequest(*ec2.CreateReservedInstancesListingInput) (*request.Request, *ec2.CreateReservedInstancesListingOutput)

	CreateReservedInstancesListing(*ec2.CreateReservedInstancesListingInput) (*ec2.CreateReservedInstancesListingOutput, error)

	CreateRouteRequest(*ec2.CreateRouteInput) (*request.Request, *ec2.CreateRouteOutput)

	CreateRoute(*ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error)

	CreateRouteTableRequest(*ec2.CreateRouteTableInput) (*request.Request, *ec2.CreateRouteTableOutput)

	CreateRouteTable(*ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error)

	CreateSecurityGroupRequest(*ec2.CreateSecurityGroupInput) (*request.Request, *ec2.CreateSecurityGroupOutput)

	CreateSecurityGroup(*ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error)

	CreateSnapshotRequest(*ec2.CreateSnapshotInput) (*request.Request, *ec2.Snapshot)

	CreateSnapshot(*ec2.CreateSnapshotInput) (*ec2.Snapshot, error)

	CreateSpotDatafeedSubscriptionRequest(*ec2.CreateSpotDatafeedSubscriptionInput) (*request.Request, *ec2.CreateSpotDatafeedSubscriptionOutput)

	CreateSpotDatafeedSubscription(*ec2.CreateSpotDatafeedSubscriptionInput) (*ec2.CreateSpotDatafeedSubscriptionOutput, error)

	CreateSubnetRequest(*ec2.CreateSubnetInput) (*request.Request, *ec2.CreateSubnetOutput)

	CreateSubnet(*ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error)

	CreateTagsRequest(*ec2.CreateTagsInput) (*request.Request, *ec2.CreateTagsOutput)

	CreateTags(*ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error)

	CreateVolumeRequest(*ec2.CreateVolumeInput) (*request.Request, *ec2.Volume)

	CreateVolume(*ec2.CreateVolumeInput) (*ec2.Volume, error)

	CreateVpcRequest(*ec2.CreateVpcInput) (*request.Request, *ec2.CreateVpcOutput)

	CreateVpc(*ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error)

	CreateVpcEndpointRequest(*ec2.CreateVpcEndpointInput) (*request.Request, *ec2.CreateVpcEndpointOutput)

	CreateVpcEndpoint(*ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error)

	CreateVpcPeeringConnectionRequest(*ec2.CreateVpcPeeringConnectionInput) (*request.Request, *ec2.CreateVpcPeeringConnectionOutput)

	CreateVpcPeeringConnection(*ec2.CreateVpcPeeringConnectionInput) (*ec2.CreateVpcPeeringConnectionOutput, error)

	CreateVpnConnectionRequest(*ec2.CreateVpnConnectionInput) (*request.Request, *ec2.CreateVpnConnectionOutput)

	CreateVpnConnection(*ec2.CreateVpnConnectionInput) (*ec2.CreateVpnConnectionOutput, error)

	CreateVpnConnectionRouteRequest(*ec2.CreateVpnConnectionRouteInput) (*request.Request, *ec2.CreateVpnConnectionRouteOutput)

	CreateVpnConnectionRoute(*ec2.CreateVpnConnectionRouteInput) (*ec2.CreateVpnConnectionRouteOutput, error)

	CreateVpnGatewayRequest(*ec2.CreateVpnGatewayInput) (*request.Request, *ec2.CreateVpnGatewayOutput)

	CreateVpnGateway(*ec2.CreateVpnGatewayInput) (*ec2.CreateVpnGatewayOutput, error)

	DeleteCustomerGatewayRequest(*ec2.DeleteCustomerGatewayInput) (*request.Request, *ec2.DeleteCustomerGatewayOutput)

	DeleteCustomerGateway(*ec2.DeleteCustomerGatewayInput) (*ec2.DeleteCustomerGatewayOutput, error)

	DeleteDhcpOptionsRequest(*ec2.DeleteDhcpOptionsInput) (*request.Request, *ec2.DeleteDhcpOptionsOutput)

	DeleteDhcpOptions(*ec2.DeleteDhcpOptionsInput) (*ec2.DeleteDhcpOptionsOutput, error)

	DeleteFlowLogsRequest(*ec2.DeleteFlowLogsInput) (*request.Request, *ec2.DeleteFlowLogsOutput)

	DeleteFlowLogs(*ec2.DeleteFlowLogsInput) (*ec2.DeleteFlowLogsOutput, error)

	DeleteInternetGatewayRequest(*ec2.DeleteInternetGatewayInput) (*request.Request, *ec2.DeleteInternetGatewayOutput)

	DeleteInternetGateway(*ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error)

	DeleteKeyPairRequest(*ec2.DeleteKeyPairInput) (*request.Request, *ec2.DeleteKeyPairOutput)

	DeleteKeyPair(*ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error)

	DeleteNetworkAclRequest(*ec2.DeleteNetworkAclInput) (*request.Request, *ec2.DeleteNetworkAclOutput)

	DeleteNetworkAcl(*ec2.DeleteNetworkAclInput) (*ec2.DeleteNetworkAclOutput, error)

	DeleteNetworkAclEntryRequest(*ec2.DeleteNetworkAclEntryInput) (*request.Request, *ec2.DeleteNetworkAclEntryOutput)

	DeleteNetworkAclEntry(*ec2.DeleteNetworkAclEntryInput) (*ec2.DeleteNetworkAclEntryOutput, error)

	DeleteNetworkInterfaceRequest(*ec2.DeleteNetworkInterfaceInput) (*request.Request, *ec2.DeleteNetworkInterfaceOutput)

	DeleteNetworkInterface(*ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error)

	DeletePlacementGroupRequest(*ec2.DeletePlacementGroupInput) (*request.Request, *ec2.DeletePlacementGroupOutput)

	DeletePlacementGroup(*ec2.DeletePlacementGroupInput) (*ec2.DeletePlacementGroupOutput, error)

	DeleteRouteRequest(*ec2.DeleteRouteInput) (*request.Request, *ec2.DeleteRouteOutput)

	DeleteRoute(*ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error)

	DeleteRouteTableRequest(*ec2.DeleteRouteTableInput) (*request.Request, *ec2.DeleteRouteTableOutput)

	DeleteRouteTable(*ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error)

	DeleteSecurityGroupRequest(*ec2.DeleteSecurityGroupInput) (*request.Request, *ec2.DeleteSecurityGroupOutput)

	DeleteSecurityGroup(*ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error)

	DeleteSnapshotRequest(*ec2.DeleteSnapshotInput) (*request.Request, *ec2.DeleteSnapshotOutput)

	DeleteSnapshot(*ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error)

	DeleteSpotDatafeedSubscriptionRequest(*ec2.DeleteSpotDatafeedSubscriptionInput) (*request.Request, *ec2.DeleteSpotDatafeedSubscriptionOutput)

	DeleteSpotDatafeedSubscription(*ec2.DeleteSpotDatafeedSubscriptionInput) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error)

	DeleteSubnetRequest(*ec2.DeleteSubnetInput) (*request.Request, *ec2.DeleteSubnetOutput)

	DeleteSubnet(*ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error)

	DeleteTagsRequest(*ec2.DeleteTagsInput) (*request.Request, *ec2.DeleteTagsOutput)

	DeleteTags(*ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error)

	DeleteVolumeRequest(*ec2.DeleteVolumeInput) (*request.Request, *ec2.DeleteVolumeOutput)

	DeleteVolume(*ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error)

	DeleteVpcRequest(*ec2.DeleteVpcInput) (*request.Request, *ec2.DeleteVpcOutput)

	DeleteVpc(*ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error)

	DeleteVpcEndpointsRequest(*ec2.DeleteVpcEndpointsInput) (*request.Request, *ec2.DeleteVpcEndpointsOutput)

	DeleteVpcEndpoints(*ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error)

	DeleteVpcPeeringConnectionRequest(*ec2.DeleteVpcPeeringConnectionInput) (*request.Request, *ec2.DeleteVpcPeeringConnectionOutput)

	DeleteVpcPeeringConnection(*ec2.DeleteVpcPeeringConnectionInput) (*ec2.DeleteVpcPeeringConnectionOutput, error)

	DeleteVpnConnectionRequest(*ec2.DeleteVpnConnectionInput) (*request.Request, *ec2.DeleteVpnConnectionOutput)

	DeleteVpnConnection(*ec2.DeleteVpnConnectionInput) (*ec2.DeleteVpnConnectionOutput, error)

	DeleteVpnConnectionRouteRequest(*ec2.DeleteVpnConnectionRouteInput) (*request.Request, *ec2.DeleteVpnConnectionRouteOutput)

	DeleteVpnConnectionRoute(*ec2.DeleteVpnConnectionRouteInput) (*ec2.DeleteVpnConnectionRouteOutput, error)

	DeleteVpnGatewayRequest(*ec2.DeleteVpnGatewayInput) (*request.Request, *ec2.DeleteVpnGatewayOutput)

	DeleteVpnGateway(*ec2.DeleteVpnGatewayInput) (*ec2.DeleteVpnGatewayOutput, error)

	DeregisterImageRequest(*ec2.DeregisterImageInput) (*request.Request, *ec2.DeregisterImageOutput)

	DeregisterImage(*ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error)

	DescribeAccountAttributesRequest(*ec2.DescribeAccountAttributesInput) (*request.Request, *ec2.DescribeAccountAttributesOutput)

	DescribeAccountAttributes(*ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error)

	DescribeAddressesRequest(*ec2.DescribeAddressesInput) (*request.Request, *ec2.DescribeAddressesOutput)

	DescribeAddresses(*ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error)

	DescribeAvailabilityZonesRequest(*ec2.DescribeAvailabilityZonesInput) (*request.Request, *ec2.DescribeAvailabilityZonesOutput)

	DescribeAvailabilityZones(*ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error)

	DescribeBundleTasksRequest(*ec2.DescribeBundleTasksInput) (*request.Request, *ec2.DescribeBundleTasksOutput)

	DescribeBundleTasks(*ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error)

	DescribeClassicLinkInstancesRequest(*ec2.DescribeClassicLinkInstancesInput) (*request.Request, *ec2.DescribeClassicLinkInstancesOutput)

	DescribeClassicLinkInstances(*ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error)

	DescribeConversionTasksRequest(*ec2.DescribeConversionTasksInput) (*request.Request, *ec2.DescribeConversionTasksOutput)

	DescribeConversionTasks(*ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error)

	DescribeCustomerGatewaysRequest(*ec2.DescribeCustomerGatewaysInput) (*request.Request, *ec2.DescribeCustomerGatewaysOutput)

	DescribeCustomerGateways(*ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error)

	DescribeDhcpOptionsRequest(*ec2.DescribeDhcpOptionsInput) (*request.Request, *ec2.DescribeDhcpOptionsOutput)

	DescribeDhcpOptions(*ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error)

	DescribeExportTasksRequest(*ec2.DescribeExportTasksInput) (*request.Request, *ec2.DescribeExportTasksOutput)

	DescribeExportTasks(*ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error)

	DescribeFlowLogsRequest(*ec2.DescribeFlowLogsInput) (*request.Request, *ec2.DescribeFlowLogsOutput)

	DescribeFlowLogs(*ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error)

	DescribeImageAttributeRequest(*ec2.DescribeImageAttributeInput) (*request.Request, *ec2.DescribeImageAttributeOutput)

	DescribeImageAttribute(*ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error)

	DescribeImagesRequest(*ec2.DescribeImagesInput) (*request.Request, *ec2.DescribeImagesOutput)

	DescribeImages(*ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error)

	DescribeImportImageTasksRequest(*ec2.DescribeImportImageTasksInput) (*request.Request, *ec2.DescribeImportImageTasksOutput)

	DescribeImportImageTasks(*ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error)

	DescribeImportSnapshotTasksRequest(*ec2.DescribeImportSnapshotTasksInput) (*request.Request, *ec2.DescribeImportSnapshotTasksOutput)

	DescribeImportSnapshotTasks(*ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error)

	DescribeInstanceAttributeRequest(*ec2.DescribeInstanceAttributeInput) (*request.Request, *ec2.DescribeInstanceAttributeOutput)

	DescribeInstanceAttribute(*ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error)

	DescribeInstanceStatusRequest(*ec2.DescribeInstanceStatusInput) (*request.Request, *ec2.DescribeInstanceStatusOutput)

	DescribeInstanceStatus(*ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error)

	DescribeInstanceStatusPages(*ec2.DescribeInstanceStatusInput, func(*ec2.DescribeInstanceStatusOutput, bool) bool) error

	DescribeInstancesRequest(*ec2.DescribeInstancesInput) (*request.Request, *ec2.DescribeInstancesOutput)

	DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error)

	DescribeInstancesPages(*ec2.DescribeInstancesInput, func(*ec2.DescribeInstancesOutput, bool) bool) error

	DescribeInternetGatewaysRequest(*ec2.DescribeInternetGatewaysInput) (*request.Request, *ec2.DescribeInternetGatewaysOutput)

	DescribeInternetGateways(*ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error)

	DescribeKeyPairsRequest(*ec2.DescribeKeyPairsInput) (*request.Request, *ec2.DescribeKeyPairsOutput)

	DescribeKeyPairs(*ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error)

	DescribeMovingAddressesRequest(*ec2.DescribeMovingAddressesInput) (*request.Request, *ec2.DescribeMovingAddressesOutput)

	DescribeMovingAddresses(*ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error)

	DescribeNetworkAclsRequest(*ec2.DescribeNetworkAclsInput) (*request.Request, *ec2.DescribeNetworkAclsOutput)

	DescribeNetworkAcls(*ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error)

	DescribeNetworkInterfaceAttributeRequest(*ec2.DescribeNetworkInterfaceAttributeInput) (*request.Request, *ec2.DescribeNetworkInterfaceAttributeOutput)

	DescribeNetworkInterfaceAttribute(*ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error)

	DescribeNetworkInterfacesRequest(*ec2.DescribeNetworkInterfacesInput) (*request.Request, *ec2.DescribeNetworkInterfacesOutput)

	DescribeNetworkInterfaces(*ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error)

	DescribePlacementGroupsRequest(*ec2.DescribePlacementGroupsInput) (*request.Request, *ec2.DescribePlacementGroupsOutput)

	DescribePlacementGroups(*ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error)

	DescribePrefixListsRequest(*ec2.DescribePrefixListsInput) (*request.Request, *ec2.DescribePrefixListsOutput)

	DescribePrefixLists(*ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error)

	DescribeRegionsRequest(*ec2.DescribeRegionsInput) (*request.Request, *ec2.DescribeRegionsOutput)

	DescribeRegions(*ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error)

	DescribeReservedInstancesRequest(*ec2.DescribeReservedInstancesInput) (*request.Request, *ec2.DescribeReservedInstancesOutput)

	DescribeReservedInstances(*ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error)

	DescribeReservedInstancesListingsRequest(*ec2.DescribeReservedInstancesListingsInput) (*request.Request, *ec2.DescribeReservedInstancesListingsOutput)

	DescribeReservedInstancesListings(*ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error)

	DescribeReservedInstancesModificationsRequest(*ec2.DescribeReservedInstancesModificationsInput) (*request.Request, *ec2.DescribeReservedInstancesModificationsOutput)

	DescribeReservedInstancesModifications(*ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error)

	DescribeReservedInstancesModificationsPages(*ec2.DescribeReservedInstancesModificationsInput, func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool) error

	DescribeReservedInstancesOfferingsRequest(*ec2.DescribeReservedInstancesOfferingsInput) (*request.Request, *ec2.DescribeReservedInstancesOfferingsOutput)

	DescribeReservedInstancesOfferings(*ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error)

	DescribeReservedInstancesOfferingsPages(*ec2.DescribeReservedInstancesOfferingsInput, func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool) error

	DescribeRouteTablesRequest(*ec2.DescribeRouteTablesInput) (*request.Request, *ec2.DescribeRouteTablesOutput)

	DescribeRouteTables(*ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error)

	DescribeSecurityGroupsRequest(*ec2.DescribeSecurityGroupsInput) (*request.Request, *ec2.DescribeSecurityGroupsOutput)

	DescribeSecurityGroups(*ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error)

	DescribeSnapshotAttributeRequest(*ec2.DescribeSnapshotAttributeInput) (*request.Request, *ec2.DescribeSnapshotAttributeOutput)

	DescribeSnapshotAttribute(*ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error)

	DescribeSnapshotsRequest(*ec2.DescribeSnapshotsInput) (*request.Request, *ec2.DescribeSnapshotsOutput)

	DescribeSnapshots(*ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error)

	DescribeSnapshotsPages(*ec2.DescribeSnapshotsInput, func(*ec2.DescribeSnapshotsOutput, bool) bool) error

	DescribeSpotDatafeedSubscriptionRequest(*ec2.DescribeSpotDatafeedSubscriptionInput) (*request.Request, *ec2.DescribeSpotDatafeedSubscriptionOutput)

	DescribeSpotDatafeedSubscription(*ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error)

	DescribeSpotFleetInstancesRequest(*ec2.DescribeSpotFleetInstancesInput) (*request.Request, *ec2.DescribeSpotFleetInstancesOutput)

	DescribeSpotFleetInstances(*ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error)

	DescribeSpotFleetRequestHistoryRequest(*ec2.DescribeSpotFleetRequestHistoryInput) (*request.Request, *ec2.DescribeSpotFleetRequestHistoryOutput)

	DescribeSpotFleetRequestHistory(*ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error)

	DescribeSpotFleetRequestsRequest(*ec2.DescribeSpotFleetRequestsInput) (*request.Request, *ec2.DescribeSpotFleetRequestsOutput)

	DescribeSpotFleetRequests(*ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error)

	DescribeSpotInstanceRequestsRequest(*ec2.DescribeSpotInstanceRequestsInput) (*request.Request, *ec2.DescribeSpotInstanceRequestsOutput)

	DescribeSpotInstanceRequests(*ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error)

	DescribeSpotPriceHistoryRequest(*ec2.DescribeSpotPriceHistoryInput) (*request.Request, *ec2.DescribeSpotPriceHistoryOutput)

	DescribeSpotPriceHistory(*ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error)

	DescribeSpotPriceHistoryPages(*ec2.DescribeSpotPriceHistoryInput, func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool) error

	DescribeSubnetsRequest(*ec2.DescribeSubnetsInput) (*request.Request, *ec2.DescribeSubnetsOutput)

	DescribeSubnets(*ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error)

	DescribeTagsRequest(*ec2.DescribeTagsInput) (*request.Request, *ec2.DescribeTagsOutput)

	DescribeTags(*ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error)

	DescribeVolumeAttributeRequest(*ec2.DescribeVolumeAttributeInput) (*request.Request, *ec2.DescribeVolumeAttributeOutput)

	DescribeVolumeAttribute(*ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error)

	DescribeVolumeStatusRequest(*ec2.DescribeVolumeStatusInput) (*request.Request, *ec2.DescribeVolumeStatusOutput)

	DescribeVolumeStatus(*ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error)

	DescribeVolumeStatusPages(*ec2.DescribeVolumeStatusInput, func(*ec2.DescribeVolumeStatusOutput, bool) bool) error

	DescribeVolumesRequest(*ec2.DescribeVolumesInput) (*request.Request, *ec2.DescribeVolumesOutput)

	DescribeVolumes(*ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error)

	DescribeVolumesPages(*ec2.DescribeVolumesInput, func(*ec2.DescribeVolumesOutput, bool) bool) error

	DescribeVpcAttributeRequest(*ec2.DescribeVpcAttributeInput) (*request.Request, *ec2.DescribeVpcAttributeOutput)

	DescribeVpcAttribute(*ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error)

	DescribeVpcClassicLinkRequest(*ec2.DescribeVpcClassicLinkInput) (*request.Request, *ec2.DescribeVpcClassicLinkOutput)

	DescribeVpcClassicLink(*ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error)

	DescribeVpcEndpointServicesRequest(*ec2.DescribeVpcEndpointServicesInput) (*request.Request, *ec2.DescribeVpcEndpointServicesOutput)

	DescribeVpcEndpointServices(*ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error)

	DescribeVpcEndpointsRequest(*ec2.DescribeVpcEndpointsInput) (*request.Request, *ec2.DescribeVpcEndpointsOutput)

	DescribeVpcEndpoints(*ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error)

	DescribeVpcPeeringConnectionsRequest(*ec2.DescribeVpcPeeringConnectionsInput) (*request.Request, *ec2.DescribeVpcPeeringConnectionsOutput)

	DescribeVpcPeeringConnections(*ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error)

	DescribeVpcsRequest(*ec2.DescribeVpcsInput) (*request.Request, *ec2.DescribeVpcsOutput)

	DescribeVpcs(*ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error)

	DescribeVpnConnectionsRequest(*ec2.DescribeVpnConnectionsInput) (*request.Request, *ec2.DescribeVpnConnectionsOutput)

	DescribeVpnConnections(*ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error)

	DescribeVpnGatewaysRequest(*ec2.DescribeVpnGatewaysInput) (*request.Request, *ec2.DescribeVpnGatewaysOutput)

	DescribeVpnGateways(*ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error)

	DetachClassicLinkVpcRequest(*ec2.DetachClassicLinkVpcInput) (*request.Request, *ec2.DetachClassicLinkVpcOutput)

	DetachClassicLinkVpc(*ec2.DetachClassicLinkVpcInput) (*ec2.DetachClassicLinkVpcOutput, error)

	DetachInternetGatewayRequest(*ec2.DetachInternetGatewayInput) (*request.Request, *ec2.DetachInternetGatewayOutput)

	DetachInternetGateway(*ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error)

	DetachNetworkInterfaceRequest(*ec2.DetachNetworkInterfaceInput) (*request.Request, *ec2.DetachNetworkInterfaceOutput)

	DetachNetworkInterface(*ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error)

	DetachVolumeRequest(*ec2.DetachVolumeInput) (*request.Request, *ec2.VolumeAttachment)

	DetachVolume(*ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error)

	DetachVpnGatewayRequest(*ec2.DetachVpnGatewayInput) (*request.Request, *ec2.DetachVpnGatewayOutput)

	DetachVpnGateway(*ec2.DetachVpnGatewayInput) (*ec2.DetachVpnGatewayOutput, error)

	DisableVgwRoutePropagationRequest(*ec2.DisableVgwRoutePropagationInput) (*request.Request, *ec2.DisableVgwRoutePropagationOutput)

	DisableVgwRoutePropagation(*ec2.DisableVgwRoutePropagationInput) (*ec2.DisableVgwRoutePropagationOutput, error)

	DisableVpcClassicLinkRequest(*ec2.DisableVpcClassicLinkInput) (*request.Request, *ec2.DisableVpcClassicLinkOutput)

	DisableVpcClassicLink(*ec2.DisableVpcClassicLinkInput) (*ec2.DisableVpcClassicLinkOutput, error)

	DisassociateAddressRequest(*ec2.DisassociateAddressInput) (*request.Request, *ec2.DisassociateAddressOutput)

	DisassociateAddress(*ec2.DisassociateAddressInput) (*ec2.DisassociateAddressOutput, error)

	DisassociateRouteTableRequest(*ec2.DisassociateRouteTableInput) (*request.Request, *ec2.DisassociateRouteTableOutput)

	DisassociateRouteTable(*ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error)

	EnableVgwRoutePropagationRequest(*ec2.EnableVgwRoutePropagationInput) (*request.Request, *ec2.EnableVgwRoutePropagationOutput)

	EnableVgwRoutePropagation(*ec2.EnableVgwRoutePropagationInput) (*ec2.EnableVgwRoutePropagationOutput, error)

	EnableVolumeIORequest(*ec2.EnableVolumeIOInput) (*request.Request, *ec2.EnableVolumeIOOutput)

	EnableVolumeIO(*ec2.EnableVolumeIOInput) (*ec2.EnableVolumeIOOutput, error)

	EnableVpcClassicLinkRequest(*ec2.EnableVpcClassicLinkInput) (*request.Request, *ec2.EnableVpcClassicLinkOutput)

	EnableVpcClassicLink(*ec2.EnableVpcClassicLinkInput) (*ec2.EnableVpcClassicLinkOutput, error)

	GetConsoleOutputRequest(*ec2.GetConsoleOutputInput) (*request.Request, *ec2.GetConsoleOutputOutput)

	GetConsoleOutput(*ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error)

	GetPasswordDataRequest(*ec2.GetPasswordDataInput) (*request.Request, *ec2.GetPasswordDataOutput)

	GetPasswordData(*ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error)

	ImportImageRequest(*ec2.ImportImageInput) (*request.Request, *ec2.ImportImageOutput)

	ImportImage(*ec2.ImportImageInput) (*ec2.ImportImageOutput, error)

	ImportInstanceRequest(*ec2.ImportInstanceInput) (*request.Request, *ec2.ImportInstanceOutput)

	ImportInstance(*ec2.ImportInstanceInput) (*ec2.ImportInstanceOutput, error)

	ImportKeyPairRequest(*ec2.ImportKeyPairInput) (*request.Request, *ec2.ImportKeyPairOutput)

	ImportKeyPair(*ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error)

	ImportSnapshotRequest(*ec2.ImportSnapshotInput) (*request.Request, *ec2.ImportSnapshotOutput)

	ImportSnapshot(*ec2.ImportSnapshotInput) (*ec2.ImportSnapshotOutput, error)

	ImportVolumeRequest(*ec2.ImportVolumeInput) (*request.Request, *ec2.ImportVolumeOutput)

	ImportVolume(*ec2.ImportVolumeInput) (*ec2.ImportVolumeOutput, error)

	ModifyImageAttributeRequest(*ec2.ModifyImageAttributeInput) (*request.Request, *ec2.ModifyImageAttributeOutput)

	ModifyImageAttribute(*ec2.ModifyImageAttributeInput) (*ec2.ModifyImageAttributeOutput, error)

	ModifyInstanceAttributeRequest(*ec2.ModifyInstanceAttributeInput) (*request.Request, *ec2.ModifyInstanceAttributeOutput)

	ModifyInstanceAttribute(*ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error)

	ModifyNetworkInterfaceAttributeRequest(*ec2.ModifyNetworkInterfaceAttributeInput) (*request.Request, *ec2.ModifyNetworkInterfaceAttributeOutput)

	ModifyNetworkInterfaceAttribute(*ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error)

	ModifyReservedInstancesRequest(*ec2.ModifyReservedInstancesInput) (*request.Request, *ec2.ModifyReservedInstancesOutput)

	ModifyReservedInstances(*ec2.ModifyReservedInstancesInput) (*ec2.ModifyReservedInstancesOutput, error)

	ModifySnapshotAttributeRequest(*ec2.ModifySnapshotAttributeInput) (*request.Request, *ec2.ModifySnapshotAttributeOutput)

	ModifySnapshotAttribute(*ec2.ModifySnapshotAttributeInput) (*ec2.ModifySnapshotAttributeOutput, error)

	ModifySpotFleetRequestRequest(*ec2.ModifySpotFleetRequestInput) (*request.Request, *ec2.ModifySpotFleetRequestOutput)

	ModifySpotFleetRequest(*ec2.ModifySpotFleetRequestInput) (*ec2.ModifySpotFleetRequestOutput, error)

	ModifySubnetAttributeRequest(*ec2.ModifySubnetAttributeInput) (*request.Request, *ec2.ModifySubnetAttributeOutput)

	ModifySubnetAttribute(*ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error)

	ModifyVolumeAttributeRequest(*ec2.ModifyVolumeAttributeInput) (*request.Request, *ec2.ModifyVolumeAttributeOutput)

	ModifyVolumeAttribute(*ec2.ModifyVolumeAttributeInput) (*ec2.ModifyVolumeAttributeOutput, error)

	ModifyVpcAttributeRequest(*ec2.ModifyVpcAttributeInput) (*request.Request, *ec2.ModifyVpcAttributeOutput)

	ModifyVpcAttribute(*ec2.ModifyVpcAttributeInput) (*ec2.ModifyVpcAttributeOutput, error)

	ModifyVpcEndpointRequest(*ec2.ModifyVpcEndpointInput) (*request.Request, *ec2.ModifyVpcEndpointOutput)

	ModifyVpcEndpoint(*ec2.ModifyVpcEndpointInput) (*ec2.ModifyVpcEndpointOutput, error)

	MonitorInstancesRequest(*ec2.MonitorInstancesInput) (*request.Request, *ec2.MonitorInstancesOutput)

	MonitorInstances(*ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error)

	MoveAddressToVpcRequest(*ec2.MoveAddressToVpcInput) (*request.Request, *ec2.MoveAddressToVpcOutput)

	MoveAddressToVpc(*ec2.MoveAddressToVpcInput) (*ec2.MoveAddressToVpcOutput, error)

	PurchaseReservedInstancesOfferingRequest(*ec2.PurchaseReservedInstancesOfferingInput) (*request.Request, *ec2.PurchaseReservedInstancesOfferingOutput)

	PurchaseReservedInstancesOffering(*ec2.PurchaseReservedInstancesOfferingInput) (*ec2.PurchaseReservedInstancesOfferingOutput, error)

	RebootInstancesRequest(*ec2.RebootInstancesInput) (*request.Request, *ec2.RebootInstancesOutput)

	RebootInstances(*ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error)

	RegisterImageRequest(*ec2.RegisterImageInput) (*request.Request, *ec2.RegisterImageOutput)

	RegisterImage(*ec2.RegisterImageInput) (*ec2.RegisterImageOutput, error)

	RejectVpcPeeringConnectionRequest(*ec2.RejectVpcPeeringConnectionInput) (*request.Request, *ec2.RejectVpcPeeringConnectionOutput)

	RejectVpcPeeringConnection(*ec2.RejectVpcPeeringConnectionInput) (*ec2.RejectVpcPeeringConnectionOutput, error)

	ReleaseAddressRequest(*ec2.ReleaseAddressInput) (*request.Request, *ec2.ReleaseAddressOutput)

	ReleaseAddress(*ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error)

	ReplaceNetworkAclAssociationRequest(*ec2.ReplaceNetworkAclAssociationInput) (*request.Request, *ec2.ReplaceNetworkAclAssociationOutput)

	ReplaceNetworkAclAssociation(*ec2.ReplaceNetworkAclAssociationInput) (*ec2.ReplaceNetworkAclAssociationOutput, error)

	ReplaceNetworkAclEntryRequest(*ec2.ReplaceNetworkAclEntryInput) (*request.Request, *ec2.ReplaceNetworkAclEntryOutput)

	ReplaceNetworkAclEntry(*ec2.ReplaceNetworkAclEntryInput) (*ec2.ReplaceNetworkAclEntryOutput, error)

	ReplaceRouteRequest(*ec2.ReplaceRouteInput) (*request.Request, *ec2.ReplaceRouteOutput)

	ReplaceRoute(*ec2.ReplaceRouteInput) (*ec2.ReplaceRouteOutput, error)

	ReplaceRouteTableAssociationRequest(*ec2.ReplaceRouteTableAssociationInput) (*request.Request, *ec2.ReplaceRouteTableAssociationOutput)

	ReplaceRouteTableAssociation(*ec2.ReplaceRouteTableAssociationInput) (*ec2.ReplaceRouteTableAssociationOutput, error)

	ReportInstanceStatusRequest(*ec2.ReportInstanceStatusInput) (*request.Request, *ec2.ReportInstanceStatusOutput)

	ReportInstanceStatus(*ec2.ReportInstanceStatusInput) (*ec2.ReportInstanceStatusOutput, error)

	RequestSpotFleetRequest(*ec2.RequestSpotFleetInput) (*request.Request, *ec2.RequestSpotFleetOutput)

	RequestSpotFleet(*ec2.RequestSpotFleetInput) (*ec2.RequestSpotFleetOutput, error)

	RequestSpotInstancesRequest(*ec2.RequestSpotInstancesInput) (*request.Request, *ec2.RequestSpotInstancesOutput)

	RequestSpotInstances(*ec2.RequestSpotInstancesInput) (*ec2.RequestSpotInstancesOutput, error)

	ResetImageAttributeRequest(*ec2.ResetImageAttributeInput) (*request.Request, *ec2.ResetImageAttributeOutput)

	ResetImageAttribute(*ec2.ResetImageAttributeInput) (*ec2.ResetImageAttributeOutput, error)

	ResetInstanceAttributeRequest(*ec2.ResetInstanceAttributeInput) (*request.Request, *ec2.ResetInstanceAttributeOutput)

	ResetInstanceAttribute(*ec2.ResetInstanceAttributeInput) (*ec2.ResetInstanceAttributeOutput, error)

	ResetNetworkInterfaceAttributeRequest(*ec2.ResetNetworkInterfaceAttributeInput) (*request.Request, *ec2.ResetNetworkInterfaceAttributeOutput)

	ResetNetworkInterfaceAttribute(*ec2.ResetNetworkInterfaceAttributeInput) (*ec2.ResetNetworkInterfaceAttributeOutput, error)

	ResetSnapshotAttributeRequest(*ec2.ResetSnapshotAttributeInput) (*request.Request, *ec2.ResetSnapshotAttributeOutput)

	ResetSnapshotAttribute(*ec2.ResetSnapshotAttributeInput) (*ec2.ResetSnapshotAttributeOutput, error)

	RestoreAddressToClassicRequest(*ec2.RestoreAddressToClassicInput) (*request.Request, *ec2.RestoreAddressToClassicOutput)

	RestoreAddressToClassic(*ec2.RestoreAddressToClassicInput) (*ec2.RestoreAddressToClassicOutput, error)

	RevokeSecurityGroupEgressRequest(*ec2.RevokeSecurityGroupEgressInput) (*request.Request, *ec2.RevokeSecurityGroupEgressOutput)

	RevokeSecurityGroupEgress(*ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error)

	RevokeSecurityGroupIngressRequest(*ec2.RevokeSecurityGroupIngressInput) (*request.Request, *ec2.RevokeSecurityGroupIngressOutput)

	RevokeSecurityGroupIngress(*ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error)

	RunInstancesRequest(*ec2.RunInstancesInput) (*request.Request, *ec2.Reservation)

	RunInstances(*ec2.RunInstancesInput) (*ec2.Reservation, error)

	StartInstancesRequest(*ec2.StartInstancesInput) (*request.Request, *ec2.StartInstancesOutput)

	StartInstances(*ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error)

	StopInstancesRequest(*ec2.StopInstancesInput) (*request.Request, *ec2.StopInstancesOutput)

	StopInstances(*ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error)

	TerminateInstancesRequest(*ec2.TerminateInstancesInput) (*request.Request, *ec2.TerminateInstancesOutput)

	TerminateInstances(*ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error)

	UnassignPrivateIpAddressesRequest(*ec2.UnassignPrivateIpAddressesInput) (*request.Request, *ec2.UnassignPrivateIpAddressesOutput)

	UnassignPrivateIpAddresses(*ec2.UnassignPrivateIpAddressesInput) (*ec2.UnassignPrivateIpAddressesOutput, error)

	UnmonitorInstancesRequest(*ec2.UnmonitorInstancesInput) (*request.Request, *ec2.UnmonitorInstancesOutput)

	UnmonitorInstances(*ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error)
}
