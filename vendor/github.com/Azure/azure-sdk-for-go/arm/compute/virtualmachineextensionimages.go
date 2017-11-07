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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// VirtualMachineExtensionImagesClient is the compute Client
type VirtualMachineExtensionImagesClient struct {
	ManagementClient
}

// NewVirtualMachineExtensionImagesClient creates an instance of the VirtualMachineExtensionImagesClient client.
func NewVirtualMachineExtensionImagesClient(subscriptionID string) VirtualMachineExtensionImagesClient {
	return NewVirtualMachineExtensionImagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachineExtensionImagesClientWithBaseURI creates an instance of the VirtualMachineExtensionImagesClient
// client.
func NewVirtualMachineExtensionImagesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineExtensionImagesClient {
	return VirtualMachineExtensionImagesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a virtual machine extension image.
//
// location is the name of a supported Azure region.
func (client VirtualMachineExtensionImagesClient) Get(location string, publisherName string, typeParameter string, version string) (result VirtualMachineExtensionImage, err error) {
	req, err := client.GetPreparer(location, publisherName, typeParameter, version)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionImagesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionImagesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionImagesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachineExtensionImagesClient) GetPreparer(location string, publisherName string, typeParameter string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"publisherName":  autorest.Encode("path", publisherName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"type":           autorest.Encode("path", typeParameter),
		"version":        autorest.Encode("path", version),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions/{version}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineExtensionImagesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachineExtensionImagesClient) GetResponder(resp *http.Response) (result VirtualMachineExtensionImage, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListTypes gets a list of virtual machine extension image types.
//
// location is the name of a supported Azure region.
func (client VirtualMachineExtensionImagesClient) ListTypes(location string, publisherName string) (result ListVirtualMachineExtensionImage, err error) {
	req, err := client.ListTypesPreparer(location, publisherName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionImagesClient", "ListTypes", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListTypesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionImagesClient", "ListTypes", resp, "Failure sending request")
		return
	}

	result, err = client.ListTypesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionImagesClient", "ListTypes", resp, "Failure responding to request")
	}

	return
}

// ListTypesPreparer prepares the ListTypes request.
func (client VirtualMachineExtensionImagesClient) ListTypesPreparer(location string, publisherName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"publisherName":  autorest.Encode("path", publisherName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListTypesSender sends the ListTypes request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineExtensionImagesClient) ListTypesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListTypesResponder handles the response to the ListTypes request. The method always
// closes the http.Response Body.
func (client VirtualMachineExtensionImagesClient) ListTypesResponder(resp *http.Response) (result ListVirtualMachineExtensionImage, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListVersions gets a list of virtual machine extension image versions.
//
// location is the name of a supported Azure region. filter is the filter to apply on the operation.
func (client VirtualMachineExtensionImagesClient) ListVersions(location string, publisherName string, typeParameter string, filter string, top *int32, orderby string) (result ListVirtualMachineExtensionImage, err error) {
	req, err := client.ListVersionsPreparer(location, publisherName, typeParameter, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionImagesClient", "ListVersions", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListVersionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionImagesClient", "ListVersions", resp, "Failure sending request")
		return
	}

	result, err = client.ListVersionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionImagesClient", "ListVersions", resp, "Failure responding to request")
	}

	return
}

// ListVersionsPreparer prepares the ListVersions request.
func (client VirtualMachineExtensionImagesClient) ListVersionsPreparer(location string, publisherName string, typeParameter string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"publisherName":  autorest.Encode("path", publisherName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"type":           autorest.Encode("path", typeParameter),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListVersionsSender sends the ListVersions request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineExtensionImagesClient) ListVersionsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListVersionsResponder handles the response to the ListVersions request. The method always
// closes the http.Response Body.
func (client VirtualMachineExtensionImagesClient) ListVersionsResponder(resp *http.Response) (result ListVirtualMachineExtensionImage, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
