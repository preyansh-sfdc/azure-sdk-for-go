package storsimple

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// BackupsClient is the client for the Backups methods of the Storsimple service.
type BackupsClient struct {
	BaseClient
}

// NewBackupsClient creates an instance of the BackupsClient client.
func NewBackupsClient(subscriptionID string) BackupsClient {
	return NewBackupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupsClientWithBaseURI creates an instance of the BackupsClient client.
func NewBackupsClientWithBaseURI(baseURI string, subscriptionID string) BackupsClient {
	return BackupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Clone clones the backup element as a new volume.
// Parameters:
// deviceName - the device name
// backupName - the backup name.
// backupElementName - the backup element name.
// parameters - the clone request object.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupsClient) Clone(ctx context.Context, deviceName string, backupName string, backupElementName string, parameters CloneRequest, resourceGroupName string, managerName string) (result BackupsCloneFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.Clone")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.TargetDeviceID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TargetVolumeName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TargetAccessControlRecordIds", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.BackupElement", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.BackupElement.ElementID", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.BackupElement.ElementName", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.BackupElement.ElementType", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.BackupElement.SizeInBytes", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.BackupElement.VolumeName", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.BackupElement.VolumeContainerID", Name: validation.Null, Rule: true, Chain: nil},
					}}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupsClient", "Clone", err.Error())
	}

	req, err := client.ClonePreparer(ctx, deviceName, backupName, backupElementName, parameters, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "Clone", nil, "Failure preparing request")
		return
	}

	result, err = client.CloneSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "Clone", result.Response(), "Failure sending request")
		return
	}

	return
}

// ClonePreparer prepares the Clone request.
func (client BackupsClient) ClonePreparer(ctx context.Context, deviceName string, backupName string, backupElementName string, parameters CloneRequest, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupElementName": backupElementName,
		"backupName":        backupName,
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups/{backupName}/elements/{backupElementName}/clone", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CloneSender sends the Clone request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) CloneSender(req *http.Request) (future BackupsCloneFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CloneResponder handles the response to the Clone request. The method always
// closes the http.Response Body.
func (client BackupsClient) CloneResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes the backup.
// Parameters:
// deviceName - the device name
// backupName - the backup name.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupsClient) Delete(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string) (result BackupsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, deviceName, backupName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BackupsClient) DeletePreparer(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupName":        backupName,
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups/{backupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) DeleteSender(req *http.Request) (future BackupsDeleteFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BackupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByDevice retrieves all the backups in a device.
// Parameters:
// deviceName - the device name
// resourceGroupName - the resource group name
// managerName - the manager name
// filter - oData Filter options
func (client BackupsClient) ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string, filter string) (result BackupListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.ListByDevice")
		defer func() {
			sc := -1
			if result.bl.Response.Response != nil {
				sc = result.bl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupsClient", "ListByDevice", err.Error())
	}

	result.fn = client.listByDeviceNextResults
	req, err := client.ListByDevicePreparer(ctx, deviceName, resourceGroupName, managerName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "ListByDevice", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDeviceSender(req)
	if err != nil {
		result.bl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "ListByDevice", resp, "Failure sending request")
		return
	}

	result.bl, err = client.ListByDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "ListByDevice", resp, "Failure responding to request")
	}

	return
}

// ListByDevicePreparer prepares the ListByDevice request.
func (client BackupsClient) ListByDevicePreparer(ctx context.Context, deviceName string, resourceGroupName string, managerName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDeviceSender sends the ListByDevice request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) ListByDeviceSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByDeviceResponder handles the response to the ListByDevice request. The method always
// closes the http.Response Body.
func (client BackupsClient) ListByDeviceResponder(resp *http.Response) (result BackupList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDeviceNextResults retrieves the next set of results, if any.
func (client BackupsClient) listByDeviceNextResults(ctx context.Context, lastResults BackupList) (result BackupList, err error) {
	req, err := lastResults.backupListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storsimple.BackupsClient", "listByDeviceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storsimple.BackupsClient", "listByDeviceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "listByDeviceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDeviceComplete enumerates all values, automatically crossing page boundaries as required.
func (client BackupsClient) ListByDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string, managerName string, filter string) (result BackupListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.ListByDevice")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDevice(ctx, deviceName, resourceGroupName, managerName, filter)
	return
}

// Restore restores the backup on the device.
// Parameters:
// deviceName - the device name
// backupName - the backupSet name
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupsClient) Restore(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string) (result BackupsRestoreFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.Restore")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupsClient", "Restore", err.Error())
	}

	req, err := client.RestorePreparer(ctx, deviceName, backupName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "Restore", nil, "Failure preparing request")
		return
	}

	result, err = client.RestoreSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "Restore", result.Response(), "Failure sending request")
		return
	}

	return
}

// RestorePreparer prepares the Restore request.
func (client BackupsClient) RestorePreparer(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupName":        backupName,
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups/{backupName}/restore", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestoreSender sends the Restore request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) RestoreSender(req *http.Request) (future BackupsRestoreFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// RestoreResponder handles the response to the Restore request. The method always
// closes the http.Response Body.
func (client BackupsClient) RestoreResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
