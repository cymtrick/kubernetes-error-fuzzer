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

package azure

import (
	"time"

	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/Azure/azure-sdk-for-go/arm/compute"
	"github.com/Azure/azure-sdk-for-go/arm/network"
	"github.com/Azure/go-autorest/autorest"
)

const (
	operationPollInterval        = 3 * time.Second
	operationPollTimeoutDuration = time.Hour
)

// CreateOrUpdateSGWithRetry invokes az.SecurityGroupsClient.CreateOrUpdate with exponential backoff retry
func (az *Cloud) CreateOrUpdateSGWithRetry(sg network.SecurityGroup) error {
	return wait.Poll(operationPollInterval, operationPollTimeoutDuration, func() (bool, error) {
		resp, err := az.SecurityGroupsClient.CreateOrUpdate(az.ResourceGroup, *sg.Name, sg, nil)
		return processRetryResponse(resp, err)
	})
}

// CreateOrUpdateLBWithRetry invokes az.LoadBalancerClient.CreateOrUpdate with exponential backoff retry
func (az *Cloud) CreateOrUpdateLBWithRetry(lb network.LoadBalancer) error {
	return wait.Poll(operationPollInterval, operationPollTimeoutDuration, func() (bool, error) {
		resp, err := az.LoadBalancerClient.CreateOrUpdate(az.ResourceGroup, *lb.Name, lb, nil)
		return processRetryResponse(resp, err)
	})
}

// CreateOrUpdatePIPWithRetry invokes az.PublicIPAddressesClient.CreateOrUpdate with exponential backoff retry
func (az *Cloud) CreateOrUpdatePIPWithRetry(pip network.PublicIPAddress) error {
	return wait.Poll(operationPollInterval, operationPollTimeoutDuration, func() (bool, error) {
		resp, err := az.PublicIPAddressesClient.CreateOrUpdate(az.ResourceGroup, *pip.Name, pip, nil)
		return processRetryResponse(resp, err)
	})
}

// CreateOrUpdateInterfaceWithRetry invokes az.PublicIPAddressesClient.CreateOrUpdate with exponential backoff retry
func (az *Cloud) CreateOrUpdateInterfaceWithRetry(nic network.Interface) error {
	return wait.Poll(operationPollInterval, operationPollTimeoutDuration, func() (bool, error) {
		resp, err := az.InterfacesClient.CreateOrUpdate(az.ResourceGroup, *nic.Name, nic, nil)
		return processRetryResponse(resp, err)
	})
}

// DeletePublicIPWithRetry invokes az.PublicIPAddressesClient.Delete with exponential backoff retry
func (az *Cloud) DeletePublicIPWithRetry(pipName string) error {
	return wait.Poll(operationPollInterval, operationPollTimeoutDuration, func() (bool, error) {
		resp, err := az.PublicIPAddressesClient.Delete(az.ResourceGroup, pipName, nil)
		return processRetryResponse(resp, err)
	})
}

// DeleteLBWithRetry invokes az.LoadBalancerClient.Delete with exponential backoff retry
func (az *Cloud) DeleteLBWithRetry(lbName string) error {
	return wait.Poll(operationPollInterval, operationPollTimeoutDuration, func() (bool, error) {
		resp, err := az.LoadBalancerClient.Delete(az.ResourceGroup, lbName, nil)
		return processRetryResponse(resp, err)
	})
}

// CreateOrUpdateRouteTableWithRetry invokes az.RouteTablesClient.CreateOrUpdate with exponential backoff retry
func (az *Cloud) CreateOrUpdateRouteTableWithRetry(routeTable network.RouteTable) error {
	return wait.Poll(operationPollInterval, operationPollTimeoutDuration, func() (bool, error) {
		resp, err := az.RouteTablesClient.CreateOrUpdate(az.ResourceGroup, az.RouteTableName, routeTable, nil)
		return processRetryResponse(resp, err)
	})
}

// CreateOrUpdateRouteWithRetry invokes az.RoutesClient.CreateOrUpdate with exponential backoff retry
func (az *Cloud) CreateOrUpdateRouteWithRetry(route network.Route) error {
	return wait.Poll(operationPollInterval, operationPollTimeoutDuration, func() (bool, error) {
		resp, err := az.RoutesClient.CreateOrUpdate(az.ResourceGroup, az.RouteTableName, *route.Name, route, nil)
		return processRetryResponse(resp, err)
	})
}

// DeleteRouteWithRetry invokes az.RoutesClient.Delete with exponential backoff retry
func (az *Cloud) DeleteRouteWithRetry(routeName string) error {
	return wait.Poll(operationPollInterval, operationPollTimeoutDuration, func() (bool, error) {
		resp, err := az.RoutesClient.Delete(az.ResourceGroup, az.RouteTableName, routeName, nil)
		return processRetryResponse(resp, err)
	})
}

// CreateOrUpdateVMWithRetry invokes az.VirtualMachinesClient.CreateOrUpdate with exponential backoff retry
func (az *Cloud) CreateOrUpdateVMWithRetry(vmName string, newVM compute.VirtualMachine) error {
	return wait.Poll(operationPollInterval, operationPollTimeoutDuration, func() (bool, error) {
		resp, err := az.VirtualMachinesClient.CreateOrUpdate(az.ResourceGroup, vmName, newVM, nil)
		return processRetryResponse(resp, err)
	})
}

// An in-progress convenience function to deal with common HTTP backoff response conditions
func processRetryResponse(resp autorest.Response, err error) (bool, error) {
	if isSuccessHTTPResponse(resp) {
		return true, nil
	}
	if shouldRetryAPIRequest(resp, err) {
		return false, err
	}
	// TODO determine the complete set of short-circuit conditions
	if err != nil {
		return false, err
	}
	// Fall-through: stop periodic backoff, return error object from most recent request
	return true, err
}

// shouldRetryAPIRequest determines if the response from an HTTP request suggests periodic retry behavior
func shouldRetryAPIRequest(resp autorest.Response, err error) bool {
	// TODO determine the complete set of retry conditions
	if err != nil {
		return true
	}
	if resp.StatusCode == 429 || resp.StatusCode == 500 {
		return true
	}
	return false
}

// isSuccessHTTPResponse determines if the response from an HTTP request suggests success
func isSuccessHTTPResponse(resp autorest.Response) bool {
	// TODO determine the complete set of success conditions
	if resp.StatusCode == 200 || resp.StatusCode == 201 || resp.StatusCode == 202 {
		return true
	}
	return false
}
