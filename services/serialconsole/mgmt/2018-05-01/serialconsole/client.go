// Package serialconsole implements the Azure ARM Serialconsole service API version 2018-05-01.
//
// The Azure Serial Console allows you to access the serial console of a Virtual Machine or VM scale set instance
package serialconsole

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

const (
	// DefaultBaseURI is the default URI used for the service Serialconsole
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Serialconsole.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// DisableConsole disables the Serial Console service for all VMs and VM scale sets in the provided subscription
// Parameters:
// defaultParameter - default parameter. Leave the value as "default".
func (client BaseClient) DisableConsole(ctx context.Context, defaultParameter string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.DisableConsole")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DisableConsolePreparer(ctx, defaultParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serialconsole.BaseClient", "DisableConsole", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableConsoleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "serialconsole.BaseClient", "DisableConsole", resp, "Failure sending request")
		return
	}

	result, err = client.DisableConsoleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serialconsole.BaseClient", "DisableConsole", resp, "Failure responding to request")
	}

	return
}

// DisableConsolePreparer prepares the DisableConsole request.
func (client BaseClient) DisableConsolePreparer(ctx context.Context, defaultParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"default":        autorest.Encode("path", defaultParameter),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/consoleServices/{default}/disableConsole", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableConsoleSender sends the DisableConsole request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) DisableConsoleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DisableConsoleResponder handles the response to the DisableConsole request. The method always
// closes the http.Response Body.
func (client BaseClient) DisableConsoleResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// EnableConsole enables the Serial Console service for all VMs and VM scale sets in the provided subscription
// Parameters:
// defaultParameter - default parameter. Leave the value as "default".
func (client BaseClient) EnableConsole(ctx context.Context, defaultParameter string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.EnableConsole")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EnableConsolePreparer(ctx, defaultParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serialconsole.BaseClient", "EnableConsole", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableConsoleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "serialconsole.BaseClient", "EnableConsole", resp, "Failure sending request")
		return
	}

	result, err = client.EnableConsoleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serialconsole.BaseClient", "EnableConsole", resp, "Failure responding to request")
	}

	return
}

// EnableConsolePreparer prepares the EnableConsole request.
func (client BaseClient) EnableConsolePreparer(ctx context.Context, defaultParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"default":        autorest.Encode("path", defaultParameter),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/consoleServices/{default}/enableConsole", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableConsoleSender sends the EnableConsole request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) EnableConsoleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// EnableConsoleResponder handles the response to the EnableConsole request. The method always
// closes the http.Response Body.
func (client BaseClient) EnableConsoleResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetConsoleStatus gets whether or not Serial Console is disabled for a given subscription
// Parameters:
// defaultParameter - default parameter. Leave the value as "default".
func (client BaseClient) GetConsoleStatus(ctx context.Context, defaultParameter string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetConsoleStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetConsoleStatusPreparer(ctx, defaultParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serialconsole.BaseClient", "GetConsoleStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetConsoleStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "serialconsole.BaseClient", "GetConsoleStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetConsoleStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serialconsole.BaseClient", "GetConsoleStatus", resp, "Failure responding to request")
	}

	return
}

// GetConsoleStatusPreparer prepares the GetConsoleStatus request.
func (client BaseClient) GetConsoleStatusPreparer(ctx context.Context, defaultParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"default":        autorest.Encode("path", defaultParameter),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/consoleServices/{default}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetConsoleStatusSender sends the GetConsoleStatus request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetConsoleStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetConsoleStatusResponder handles the response to the GetConsoleStatus request. The method always
// closes the http.Response Body.
func (client BaseClient) GetConsoleStatusResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListOperations gets a list of Serial Console API operations.
func (client BaseClient) ListOperations(ctx context.Context) (result Operations, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ListOperations")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListOperationsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serialconsole.BaseClient", "ListOperations", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListOperationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "serialconsole.BaseClient", "ListOperations", resp, "Failure sending request")
		return
	}

	result, err = client.ListOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serialconsole.BaseClient", "ListOperations", resp, "Failure responding to request")
	}

	return
}

// ListOperationsPreparer prepares the ListOperations request.
func (client BaseClient) ListOperationsPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.SerialConsole/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListOperationsSender sends the ListOperations request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ListOperationsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListOperationsResponder handles the response to the ListOperations request. The method always
// closes the http.Response Body.
func (client BaseClient) ListOperationsResponder(resp *http.Response) (result Operations, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
