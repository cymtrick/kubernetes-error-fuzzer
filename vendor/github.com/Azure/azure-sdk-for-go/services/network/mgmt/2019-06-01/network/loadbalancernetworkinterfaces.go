package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// LoadBalancerNetworkInterfacesClient is the network Client
type LoadBalancerNetworkInterfacesClient struct {
	BaseClient
}

// NewLoadBalancerNetworkInterfacesClient creates an instance of the LoadBalancerNetworkInterfacesClient client.
func NewLoadBalancerNetworkInterfacesClient(subscriptionID string) LoadBalancerNetworkInterfacesClient {
	return NewLoadBalancerNetworkInterfacesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLoadBalancerNetworkInterfacesClientWithBaseURI creates an instance of the LoadBalancerNetworkInterfacesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewLoadBalancerNetworkInterfacesClientWithBaseURI(baseURI string, subscriptionID string) LoadBalancerNetworkInterfacesClient {
	return LoadBalancerNetworkInterfacesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets associated load balancer network interfaces.
// Parameters:
// resourceGroupName - the name of the resource group.
// loadBalancerName - the name of the load balancer.
func (client LoadBalancerNetworkInterfacesClient) List(ctx context.Context, resourceGroupName string, loadBalancerName string) (result InterfaceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerNetworkInterfacesClient.List")
		defer func() {
			sc := -1
			if result.ilr.Response.Response != nil {
				sc = result.ilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, loadBalancerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerNetworkInterfacesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerNetworkInterfacesClient", "List", resp, "Failure sending request")
		return
	}

	result.ilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerNetworkInterfacesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.ilr.hasNextLink() && result.ilr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client LoadBalancerNetworkInterfacesClient) ListPreparer(ctx context.Context, resourceGroupName string, loadBalancerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"loadBalancerName":  autorest.Encode("path", loadBalancerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/networkInterfaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerNetworkInterfacesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LoadBalancerNetworkInterfacesClient) ListResponder(resp *http.Response) (result InterfaceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client LoadBalancerNetworkInterfacesClient) listNextResults(ctx context.Context, lastResults InterfaceListResult) (result InterfaceListResult, err error) {
	req, err := lastResults.interfaceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerNetworkInterfacesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerNetworkInterfacesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerNetworkInterfacesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LoadBalancerNetworkInterfacesClient) ListComplete(ctx context.Context, resourceGroupName string, loadBalancerName string) (result InterfaceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerNetworkInterfacesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, loadBalancerName)
	return
}
