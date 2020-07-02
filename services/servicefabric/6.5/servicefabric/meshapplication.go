package servicefabric

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

// MeshApplicationClient is the service Fabric REST Client APIs allows management of Service Fabric clusters,
// applications and services.
type MeshApplicationClient struct {
	BaseClient
}

// NewMeshApplicationClient creates an instance of the MeshApplicationClient client.
func NewMeshApplicationClient() MeshApplicationClient {
	return NewMeshApplicationClientWithBaseURI(DefaultBaseURI)
}

// NewMeshApplicationClientWithBaseURI creates an instance of the MeshApplicationClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMeshApplicationClientWithBaseURI(baseURI string) MeshApplicationClient {
	return MeshApplicationClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate creates a Application resource with the specified name, description and properties. If Application
// resource with the same name exists, then it is updated with the specified description and properties.
// Parameters:
// applicationResourceName - the identity of the application.
// applicationResourceDescription - description for creating a Application resource.
func (client MeshApplicationClient) CreateOrUpdate(ctx context.Context, applicationResourceName string, applicationResourceDescription ApplicationResourceDescription) (result ApplicationResourceDescription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MeshApplicationClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: applicationResourceDescription,
			Constraints: []validation.Constraint{{Target: "applicationResourceDescription.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "applicationResourceDescription.ApplicationProperties", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "applicationResourceDescription.Identity", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "applicationResourceDescription.Identity.Type", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("servicefabric.MeshApplicationClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, applicationResourceName, applicationResourceDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshApplicationClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.MeshApplicationClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshApplicationClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client MeshApplicationClient) CreateOrUpdatePreparer(ctx context.Context, applicationResourceName string, applicationResourceDescription ApplicationResourceDescription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationResourceName": applicationResourceName,
	}

	const APIVersion = "6.4-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Resources/Applications/{applicationResourceName}", pathParameters),
		autorest.WithJSON(applicationResourceDescription),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client MeshApplicationClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client MeshApplicationClient) CreateOrUpdateResponder(resp *http.Response) (result ApplicationResourceDescription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the Application resource identified by the name.
// Parameters:
// applicationResourceName - the identity of the application.
func (client MeshApplicationClient) Delete(ctx context.Context, applicationResourceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MeshApplicationClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, applicationResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshApplicationClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicefabric.MeshApplicationClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshApplicationClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client MeshApplicationClient) DeletePreparer(ctx context.Context, applicationResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationResourceName": applicationResourceName,
	}

	const APIVersion = "6.4-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Resources/Applications/{applicationResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MeshApplicationClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client MeshApplicationClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the information about the Application resource with the given name. The information include the description
// and other properties of the Application.
// Parameters:
// applicationResourceName - the identity of the application.
func (client MeshApplicationClient) Get(ctx context.Context, applicationResourceName string) (result ApplicationResourceDescription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MeshApplicationClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, applicationResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshApplicationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.MeshApplicationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshApplicationClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MeshApplicationClient) GetPreparer(ctx context.Context, applicationResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationResourceName": applicationResourceName,
	}

	const APIVersion = "6.4-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Resources/Applications/{applicationResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MeshApplicationClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MeshApplicationClient) GetResponder(resp *http.Response) (result ApplicationResourceDescription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the information about all application resources in a given resource group. The information include the
// description and other properties of the Application.
func (client MeshApplicationClient) List(ctx context.Context) (result PagedApplicationResourceDescriptionList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MeshApplicationClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshApplicationClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.MeshApplicationClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshApplicationClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client MeshApplicationClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "6.4-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/Resources/Applications"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MeshApplicationClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MeshApplicationClient) ListResponder(resp *http.Response) (result PagedApplicationResourceDescriptionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
