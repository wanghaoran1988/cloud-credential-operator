package vmwarecloudsimple

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VirtualNetworksByPCClient is the description of the new service
type VirtualNetworksByPCClient struct {
	BaseClient
}

// NewVirtualNetworksByPCClient creates an instance of the VirtualNetworksByPCClient client.
func NewVirtualNetworksByPCClient(referer string, regionID string, subscriptionID string) VirtualNetworksByPCClient {
	return NewVirtualNetworksByPCClientWithBaseURI(DefaultBaseURI, referer, regionID, subscriptionID)
}

// NewVirtualNetworksByPCClientWithBaseURI creates an instance of the VirtualNetworksByPCClient client.
func NewVirtualNetworksByPCClientWithBaseURI(baseURI string, referer string, regionID string, subscriptionID string) VirtualNetworksByPCClient {
	return VirtualNetworksByPCClient{NewWithBaseURI(baseURI, referer, regionID, subscriptionID)}
}

// List return list of virtual networks in location for private cloud
// Parameters:
// pcName - the private cloud name
// resourcePoolName - resource pool used to derive vSphere cluster which contains virtual networks
func (client VirtualNetworksByPCClient) List(ctx context.Context, pcName string, resourcePoolName string) (result VirtualNetworkListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksByPCClient.List")
		defer func() {
			sc := -1
			if result.vnlr.Response.Response != nil {
				sc = result.vnlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, pcName, resourcePoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualNetworksByPCClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.vnlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualNetworksByPCClient", "List", resp, "Failure sending request")
		return
	}

	result.vnlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualNetworksByPCClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualNetworksByPCClient) ListPreparer(ctx context.Context, pcName string, resourcePoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pcName":         autorest.Encode("path", pcName),
		"regionId":       autorest.Encode("path", client.RegionID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version":      APIVersion,
		"resourcePoolName": autorest.Encode("query", resourcePoolName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualNetworks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksByPCClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualNetworksByPCClient) ListResponder(resp *http.Response) (result VirtualNetworkListResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VirtualNetworksByPCClient) listNextResults(ctx context.Context, lastResults VirtualNetworkListResponse) (result VirtualNetworkListResponse, err error) {
	req, err := lastResults.virtualNetworkListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualNetworksByPCClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualNetworksByPCClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualNetworksByPCClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualNetworksByPCClient) ListComplete(ctx context.Context, pcName string, resourcePoolName string) (result VirtualNetworkListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksByPCClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, pcName, resourcePoolName)
	return
}
