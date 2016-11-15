package compute

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// VirtualMachinesClient is the the Compute Management Client.
type VirtualMachinesClient struct {
	ManagementClient
}

// NewVirtualMachinesClient creates an instance of the VirtualMachinesClient
// client.
func NewVirtualMachinesClient(subscriptionID string) VirtualMachinesClient {
	return NewVirtualMachinesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachinesClientWithBaseURI creates an instance of the
// VirtualMachinesClient client.
func NewVirtualMachinesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachinesClient {
	return VirtualMachinesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Capture captures the VM by copying virtual hard disks of the VM and outputs
// a template that can be used to create similar VMs. This method may poll
// for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the resource group. vmName is the name of
// the virtual machine. parameters is parameters supplied to the Capture
// Virtual Machine operation.
func (client VirtualMachinesClient) Capture(resourceGroupName string, vmName string, parameters VirtualMachineCaptureParameters, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.CapturePreparer(resourceGroupName, vmName, parameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Capture", nil, "Failure preparing request")
	}

	resp, err := client.CaptureSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Capture", resp, "Failure sending request")
	}

	result, err = client.CaptureResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Capture", resp, "Failure responding to request")
	}

	return
}

// CapturePreparer prepares the Capture request.
func (client VirtualMachinesClient) CapturePreparer(resourceGroupName string, vmName string, parameters VirtualMachineCaptureParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", vmName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/capture", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CaptureSender sends the Capture request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) CaptureSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CaptureResponder handles the response to the Capture request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) CaptureResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate the operation to create or update a virtual machine. This
// method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and
// any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. vmName is the name of
// the virtual machine. parameters is parameters supplied to the Create
// Virtual Machine operation.
func (client VirtualMachinesClient) CreateOrUpdate(resourceGroupName string, vmName string, parameters VirtualMachine, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, vmName, parameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualMachinesClient) CreateOrUpdatePreparer(resourceGroupName string, vmName string, parameters VirtualMachine, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", vmName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deallocate shuts down the Virtual Machine and releases the compute
// resources. You are not billed for the compute resources that this Virtual
// Machine uses. This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. vmName is the name of
// the virtual machine.
func (client VirtualMachinesClient) Deallocate(resourceGroupName string, vmName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.DeallocatePreparer(resourceGroupName, vmName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Deallocate", nil, "Failure preparing request")
	}

	resp, err := client.DeallocateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Deallocate", resp, "Failure sending request")
	}

	result, err = client.DeallocateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Deallocate", resp, "Failure responding to request")
	}

	return
}

// DeallocatePreparer prepares the Deallocate request.
func (client VirtualMachinesClient) DeallocatePreparer(resourceGroupName string, vmName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", vmName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/deallocate", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeallocateSender sends the Deallocate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) DeallocateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeallocateResponder handles the response to the Deallocate request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) DeallocateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete the operation to delete a virtual machine. This method may poll for
// completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the resource group. vmName is the name of
// the virtual machine.
func (client VirtualMachinesClient) Delete(resourceGroupName string, vmName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, vmName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualMachinesClient) DeletePreparer(resourceGroupName string, vmName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", vmName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Generalize sets the state of the VM as Generalized.
//
// resourceGroupName is the name of the resource group. vmName is the name of
// the virtual machine.
func (client VirtualMachinesClient) Generalize(resourceGroupName string, vmName string) (result autorest.Response, err error) {
	req, err := client.GeneralizePreparer(resourceGroupName, vmName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Generalize", nil, "Failure preparing request")
	}

	resp, err := client.GeneralizeSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Generalize", resp, "Failure sending request")
	}

	result, err = client.GeneralizeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Generalize", resp, "Failure responding to request")
	}

	return
}

// GeneralizePreparer prepares the Generalize request.
func (client VirtualMachinesClient) GeneralizePreparer(resourceGroupName string, vmName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", vmName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/generalize", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GeneralizeSender sends the Generalize request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) GeneralizeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GeneralizeResponder handles the response to the Generalize request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) GeneralizeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the operation to get a virtual machine.
//
// resourceGroupName is the name of the resource group. vmName is the name of
// the virtual machine. expand is the expand expression to apply on the
// operation. Possible values include: 'instanceView'
func (client VirtualMachinesClient) Get(resourceGroupName string, vmName string, expand InstanceViewTypes) (result VirtualMachine, err error) {
	req, err := client.GetPreparer(resourceGroupName, vmName, expand)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachinesClient) GetPreparer(resourceGroupName string, vmName string, expand InstanceViewTypes) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", vmName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) GetResponder(resp *http.Response) (result VirtualMachine, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the operation to list virtual machines under a resource group.
//
// resourceGroupName is the name of the resource group.
func (client VirtualMachinesClient) List(resourceGroupName string) (result VirtualMachineListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualMachinesClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) ListResponder(resp *http.Response) (result VirtualMachineListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client VirtualMachinesClient) ListNextResults(lastResults VirtualMachineListResult) (result VirtualMachineListResult, err error) {
	req, err := lastResults.VirtualMachineListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "List", resp, "Failure responding to next results request request")
	}

	return
}

// ListAll gets the list of Virtual Machines in the subscription. Use nextLink
// property in the response to get the next page of Virtual Machines. Do this
// till nextLink is not null to fetch all the Virtual Machines.
func (client VirtualMachinesClient) ListAll() (result VirtualMachineListResult, err error) {
	req, err := client.ListAllPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "ListAll", nil, "Failure preparing request")
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "ListAll", resp, "Failure sending request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client VirtualMachinesClient) ListAllPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/virtualMachines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) ListAllResponder(resp *http.Response) (result VirtualMachineListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllNextResults retrieves the next set of results, if any.
func (client VirtualMachinesClient) ListAllNextResults(lastResults VirtualMachineListResult) (result VirtualMachineListResult, err error) {
	req, err := lastResults.VirtualMachineListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "ListAll", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "ListAll", resp, "Failure sending next results request request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "ListAll", resp, "Failure responding to next results request request")
	}

	return
}

// ListAvailableSizes lists all available virtual machine sizes it can be
// resized to for a virtual machine.
//
// resourceGroupName is the name of the resource group. vmName is the name of
// the virtual machine.
func (client VirtualMachinesClient) ListAvailableSizes(resourceGroupName string, vmName string) (result VirtualMachineSizeListResult, err error) {
	req, err := client.ListAvailableSizesPreparer(resourceGroupName, vmName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "ListAvailableSizes", nil, "Failure preparing request")
	}

	resp, err := client.ListAvailableSizesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "ListAvailableSizes", resp, "Failure sending request")
	}

	result, err = client.ListAvailableSizesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "ListAvailableSizes", resp, "Failure responding to request")
	}

	return
}

// ListAvailableSizesPreparer prepares the ListAvailableSizes request.
func (client VirtualMachinesClient) ListAvailableSizesPreparer(resourceGroupName string, vmName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", vmName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/vmSizes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListAvailableSizesSender sends the ListAvailableSizes request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) ListAvailableSizesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAvailableSizesResponder handles the response to the ListAvailableSizes request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) ListAvailableSizesResponder(resp *http.Response) (result VirtualMachineSizeListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PowerOff the operation to power off (stop) a virtual machine. This method
// may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. vmName is the name of
// the virtual machine.
func (client VirtualMachinesClient) PowerOff(resourceGroupName string, vmName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.PowerOffPreparer(resourceGroupName, vmName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "PowerOff", nil, "Failure preparing request")
	}

	resp, err := client.PowerOffSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "PowerOff", resp, "Failure sending request")
	}

	result, err = client.PowerOffResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "PowerOff", resp, "Failure responding to request")
	}

	return
}

// PowerOffPreparer prepares the PowerOff request.
func (client VirtualMachinesClient) PowerOffPreparer(resourceGroupName string, vmName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", vmName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/powerOff", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// PowerOffSender sends the PowerOff request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) PowerOffSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// PowerOffResponder handles the response to the PowerOff request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) PowerOffResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Redeploy the operation to redeploy a virtual machine. This method may poll
// for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the resource group. vmName is the name of
// the virtual machine.
func (client VirtualMachinesClient) Redeploy(resourceGroupName string, vmName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.RedeployPreparer(resourceGroupName, vmName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Redeploy", nil, "Failure preparing request")
	}

	resp, err := client.RedeploySender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Redeploy", resp, "Failure sending request")
	}

	result, err = client.RedeployResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Redeploy", resp, "Failure responding to request")
	}

	return
}

// RedeployPreparer prepares the Redeploy request.
func (client VirtualMachinesClient) RedeployPreparer(resourceGroupName string, vmName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", vmName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/redeploy", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// RedeploySender sends the Redeploy request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) RedeploySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// RedeployResponder handles the response to the Redeploy request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) RedeployResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Restart the operation to restart a virtual machine. This method may poll
// for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the resource group. vmName is the name of
// the virtual machine.
func (client VirtualMachinesClient) Restart(resourceGroupName string, vmName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.RestartPreparer(resourceGroupName, vmName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Restart", nil, "Failure preparing request")
	}

	resp, err := client.RestartSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Restart", resp, "Failure sending request")
	}

	result, err = client.RestartResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Restart", resp, "Failure responding to request")
	}

	return
}

// RestartPreparer prepares the Restart request.
func (client VirtualMachinesClient) RestartPreparer(resourceGroupName string, vmName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", vmName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/restart", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// RestartSender sends the Restart request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) RestartSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// RestartResponder handles the response to the Restart request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) RestartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Start the operation to start a virtual machine. This method may poll for
// completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the resource group. vmName is the name of
// the virtual machine.
func (client VirtualMachinesClient) Start(resourceGroupName string, vmName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.StartPreparer(resourceGroupName, vmName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Start", nil, "Failure preparing request")
	}

	resp, err := client.StartSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Start", resp, "Failure sending request")
	}

	result, err = client.StartResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachinesClient", "Start", resp, "Failure responding to request")
	}

	return
}

// StartPreparer prepares the Start request.
func (client VirtualMachinesClient) StartPreparer(resourceGroupName string, vmName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", vmName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) StartSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
