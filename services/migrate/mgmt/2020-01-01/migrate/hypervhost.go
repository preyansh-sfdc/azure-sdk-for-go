package migrate

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

// HyperVHostClient is the discover your workloads for Azure.
type HyperVHostClient struct {
	BaseClient
}

// NewHyperVHostClient creates an instance of the HyperVHostClient client.
func NewHyperVHostClient() HyperVHostClient {
	return NewHyperVHostClientWithBaseURI(DefaultBaseURI)
}

// NewHyperVHostClientWithBaseURI creates an instance of the HyperVHostClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewHyperVHostClientWithBaseURI(baseURI string) HyperVHostClient {
	return HyperVHostClient{NewWithBaseURI(baseURI)}
}

// GetAllHostsInSite sends the get all hosts in site request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// APIVersion - the API version to use for this operation.
func (client HyperVHostClient) GetAllHostsInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (result HyperVHostCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HyperVHostClient.GetAllHostsInSite")
		defer func() {
			sc := -1
			if result.hvhc.Response.Response != nil {
				sc = result.hvhc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getAllHostsInSiteNextResults
	req, err := client.GetAllHostsInSitePreparer(ctx, subscriptionID, resourceGroupName, siteName, APIVersion, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVHostClient", "GetAllHostsInSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllHostsInSiteSender(req)
	if err != nil {
		result.hvhc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.HyperVHostClient", "GetAllHostsInSite", resp, "Failure sending request")
		return
	}

	result.hvhc, err = client.GetAllHostsInSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVHostClient", "GetAllHostsInSite", resp, "Failure responding to request")
		return
	}
	if result.hvhc.hasNextLink() && result.hvhc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetAllHostsInSitePreparer prepares the GetAllHostsInSite request.
func (client HyperVHostClient) GetAllHostsInSitePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}/hosts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllHostsInSiteSender sends the GetAllHostsInSite request. The method will close the
// http.Response Body if it receives an error.
func (client HyperVHostClient) GetAllHostsInSiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAllHostsInSiteResponder handles the response to the GetAllHostsInSite request. The method always
// closes the http.Response Body.
func (client HyperVHostClient) GetAllHostsInSiteResponder(resp *http.Response) (result HyperVHostCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getAllHostsInSiteNextResults retrieves the next set of results, if any.
func (client HyperVHostClient) getAllHostsInSiteNextResults(ctx context.Context, lastResults HyperVHostCollection) (result HyperVHostCollection, err error) {
	req, err := lastResults.hyperVHostCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "migrate.HyperVHostClient", "getAllHostsInSiteNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetAllHostsInSiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "migrate.HyperVHostClient", "getAllHostsInSiteNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetAllHostsInSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVHostClient", "getAllHostsInSiteNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetAllHostsInSiteComplete enumerates all values, automatically crossing page boundaries as required.
func (client HyperVHostClient) GetAllHostsInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (result HyperVHostCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HyperVHostClient.GetAllHostsInSite")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetAllHostsInSite(ctx, subscriptionID, resourceGroupName, siteName, APIVersion, filter)
	return
}

// GetHost sends the get host request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// hostName - host ARM name.
// APIVersion - the API version to use for this operation.
func (client HyperVHostClient) GetHost(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, hostName string, APIVersion string) (result HyperVHost, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HyperVHostClient.GetHost")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetHostPreparer(ctx, subscriptionID, resourceGroupName, siteName, hostName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVHostClient", "GetHost", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetHostSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.HyperVHostClient", "GetHost", resp, "Failure sending request")
		return
	}

	result, err = client.GetHostResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVHostClient", "GetHost", resp, "Failure responding to request")
		return
	}

	return
}

// GetHostPreparer prepares the GetHost request.
func (client HyperVHostClient) GetHostPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, hostName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":          autorest.Encode("path", hostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}/hosts/{hostName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetHostSender sends the GetHost request. The method will close the
// http.Response Body if it receives an error.
func (client HyperVHostClient) GetHostSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetHostResponder handles the response to the GetHost request. The method always
// closes the http.Response Body.
func (client HyperVHostClient) GetHostResponder(resp *http.Response) (result HyperVHost, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutHost sends the put host request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// hostName - host ARM name.
// body - put host body.
// APIVersion - the API version to use for this operation.
func (client HyperVHostClient) PutHost(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, hostName string, body HyperVHost, APIVersion string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HyperVHostClient.PutHost")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutHostPreparer(ctx, subscriptionID, resourceGroupName, siteName, hostName, body, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVHostClient", "PutHost", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutHostSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "migrate.HyperVHostClient", "PutHost", resp, "Failure sending request")
		return
	}

	result, err = client.PutHostResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVHostClient", "PutHost", resp, "Failure responding to request")
		return
	}

	return
}

// PutHostPreparer prepares the PutHost request.
func (client HyperVHostClient) PutHostPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, hostName string, body HyperVHost, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":          autorest.Encode("path", hostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}/hosts/{hostName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutHostSender sends the PutHost request. The method will close the
// http.Response Body if it receives an error.
func (client HyperVHostClient) PutHostSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PutHostResponder handles the response to the PutHost request. The method always
// closes the http.Response Body.
func (client HyperVHostClient) PutHostResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
