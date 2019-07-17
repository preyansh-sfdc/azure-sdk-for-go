package billing

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

// SubscriptionsClient is the billing client provides access to billing resources for Azure subscriptions.
type SubscriptionsClient struct {
	BaseClient
}

// NewSubscriptionsClient creates an instance of the SubscriptionsClient client.
func NewSubscriptionsClient(subscriptionID string) SubscriptionsClient {
	return NewSubscriptionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSubscriptionsClientWithBaseURI creates an instance of the SubscriptionsClient client.
func NewSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsClient {
	return SubscriptionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a single billing subscription by name.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
// billingSubscriptionName - billing Subscription Id.
func (client SubscriptionsClient) Get(ctx context.Context, billingAccountName string, invoiceSectionName string, billingSubscriptionName string) (result SubscriptionSummary, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountName, invoiceSectionName, billingSubscriptionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SubscriptionsClient) GetPreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, billingSubscriptionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName":      autorest.Encode("path", billingAccountName),
		"billingSubscriptionName": autorest.Encode("path", billingSubscriptionName),
		"invoiceSectionName":      autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/billingSubscriptions/{billingSubscriptionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) GetResponder(resp *http.Response) (result SubscriptionSummary, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingAccountName lists billing subscriptions by billing account name.
// Parameters:
// billingAccountName - billing Account Id.
func (client SubscriptionsClient) ListByBillingAccountName(ctx context.Context, billingAccountName string) (result SubscriptionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ListByBillingAccountName")
		defer func() {
			sc := -1
			if result.slr.Response.Response != nil {
				sc = result.slr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingAccountNameNextResults
	req, err := client.ListByBillingAccountNamePreparer(ctx, billingAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByBillingAccountName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountNameSender(req)
	if err != nil {
		result.slr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByBillingAccountName", resp, "Failure sending request")
		return
	}

	result.slr, err = client.ListByBillingAccountNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByBillingAccountName", resp, "Failure responding to request")
	}

	return
}

// ListByBillingAccountNamePreparer prepares the ListByBillingAccountName request.
func (client SubscriptionsClient) ListByBillingAccountNamePreparer(ctx context.Context, billingAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountNameSender sends the ListByBillingAccountName request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) ListByBillingAccountNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByBillingAccountNameResponder handles the response to the ListByBillingAccountName request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) ListByBillingAccountNameResponder(resp *http.Response) (result SubscriptionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingAccountNameNextResults retrieves the next set of results, if any.
func (client SubscriptionsClient) listByBillingAccountNameNextResults(ctx context.Context, lastResults SubscriptionsListResult) (result SubscriptionsListResult, err error) {
	req, err := lastResults.subscriptionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByBillingAccountNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingAccountNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByBillingAccountNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingAccountNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByBillingAccountNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingAccountNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client SubscriptionsClient) ListByBillingAccountNameComplete(ctx context.Context, billingAccountName string) (result SubscriptionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ListByBillingAccountName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingAccountName(ctx, billingAccountName)
	return
}

// ListByBillingProfileName lists billing subscriptions by billing profile name.
// Parameters:
// billingAccountName - billing Account Id.
// billingProfileName - billing Profile Id.
func (client SubscriptionsClient) ListByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string) (result SubscriptionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ListByBillingProfileName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByBillingProfileNamePreparer(ctx, billingAccountName, billingProfileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByBillingProfileName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingProfileNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByBillingProfileName", resp, "Failure sending request")
		return
	}

	result, err = client.ListByBillingProfileNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByBillingProfileName", resp, "Failure responding to request")
	}

	return
}

// ListByBillingProfileNamePreparer prepares the ListByBillingProfileName request.
func (client SubscriptionsClient) ListByBillingProfileNamePreparer(ctx context.Context, billingAccountName string, billingProfileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingProfileNameSender sends the ListByBillingProfileName request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) ListByBillingProfileNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByBillingProfileNameResponder handles the response to the ListByBillingProfileName request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) ListByBillingProfileNameResponder(resp *http.Response) (result SubscriptionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByInvoiceSectionName lists billing subscription by invoice section name.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
func (client SubscriptionsClient) ListByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string) (result SubscriptionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ListByInvoiceSectionName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByInvoiceSectionNamePreparer(ctx, billingAccountName, invoiceSectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByInvoiceSectionName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInvoiceSectionNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByInvoiceSectionName", resp, "Failure sending request")
		return
	}

	result, err = client.ListByInvoiceSectionNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByInvoiceSectionName", resp, "Failure responding to request")
	}

	return
}

// ListByInvoiceSectionNamePreparer prepares the ListByInvoiceSectionName request.
func (client SubscriptionsClient) ListByInvoiceSectionNamePreparer(ctx context.Context, billingAccountName string, invoiceSectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/billingSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInvoiceSectionNameSender sends the ListByInvoiceSectionName request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) ListByInvoiceSectionNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByInvoiceSectionNameResponder handles the response to the ListByInvoiceSectionName request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) ListByInvoiceSectionNameResponder(resp *http.Response) (result SubscriptionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Transfer transfers the subscription from one invoice section to another within a billing account.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
// billingSubscriptionName - billing Subscription Id.
// parameters - parameters supplied to the Transfer Billing Subscription operation.
func (client SubscriptionsClient) Transfer(ctx context.Context, billingAccountName string, invoiceSectionName string, billingSubscriptionName string, parameters TransferBillingSubscriptionRequestProperties) (result SubscriptionsTransferFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.Transfer")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TransferPreparer(ctx, billingAccountName, invoiceSectionName, billingSubscriptionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "Transfer", nil, "Failure preparing request")
		return
	}

	result, err = client.TransferSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "Transfer", result.Response(), "Failure sending request")
		return
	}

	return
}

// TransferPreparer prepares the Transfer request.
func (client SubscriptionsClient) TransferPreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, billingSubscriptionName string, parameters TransferBillingSubscriptionRequestProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName":      autorest.Encode("path", billingAccountName),
		"billingSubscriptionName": autorest.Encode("path", billingSubscriptionName),
		"invoiceSectionName":      autorest.Encode("path", invoiceSectionName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/billingSubscriptions/{billingSubscriptionName}/transfer", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TransferSender sends the Transfer request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) TransferSender(req *http.Request) (future SubscriptionsTransferFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// TransferResponder handles the response to the Transfer request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) TransferResponder(resp *http.Response) (result TransferBillingSubscriptionResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidateTransfer validates the transfer of billing subscriptions across invoice sections.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
// billingSubscriptionName - billing Subscription Id.
// parameters - parameters supplied to the Transfer Billing Subscription operation.
func (client SubscriptionsClient) ValidateTransfer(ctx context.Context, billingAccountName string, invoiceSectionName string, billingSubscriptionName string, parameters TransferBillingSubscriptionRequestProperties) (result ValidateSubscriptionTransferEligibilityResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ValidateTransfer")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ValidateTransferPreparer(ctx, billingAccountName, invoiceSectionName, billingSubscriptionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ValidateTransfer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateTransferSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ValidateTransfer", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateTransferResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ValidateTransfer", resp, "Failure responding to request")
	}

	return
}

// ValidateTransferPreparer prepares the ValidateTransfer request.
func (client SubscriptionsClient) ValidateTransferPreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, billingSubscriptionName string, parameters TransferBillingSubscriptionRequestProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName":      autorest.Encode("path", billingAccountName),
		"billingSubscriptionName": autorest.Encode("path", billingSubscriptionName),
		"invoiceSectionName":      autorest.Encode("path", invoiceSectionName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/billingSubscriptions/{billingSubscriptionName}/validateTransferEligibility", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateTransferSender sends the ValidateTransfer request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) ValidateTransferSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ValidateTransferResponder handles the response to the ValidateTransfer request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) ValidateTransferResponder(resp *http.Response) (result ValidateSubscriptionTransferEligibilityResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
