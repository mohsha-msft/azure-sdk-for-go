package costmanagement

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

// SubscriptionDimensionsClient is the client for the SubscriptionDimensions methods of the Costmanagement service.
type SubscriptionDimensionsClient struct {
	BaseClient
}

// NewSubscriptionDimensionsClient creates an instance of the SubscriptionDimensionsClient client.
func NewSubscriptionDimensionsClient(subscriptionID string) SubscriptionDimensionsClient {
	return NewSubscriptionDimensionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSubscriptionDimensionsClientWithBaseURI creates an instance of the SubscriptionDimensionsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewSubscriptionDimensionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionDimensionsClient {
	return SubscriptionDimensionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the dimensions by subscription Id.
// Parameters:
// filter - may be used to filter dimensions by properties/category, properties/usageStart,
// properties/usageEnd. Supported operators are 'eq','lt', 'gt', 'le', 'ge'.
// expand - may be used to expand the properties/data within a dimension category. By default, data is not
// included when listing dimensions.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N dimension data.
func (client SubscriptionDimensionsClient) List(ctx context.Context, filter string, expand string, skiptoken string, top *int32) (result DimensionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionDimensionsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.SubscriptionDimensionsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, filter, expand, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.SubscriptionDimensionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.SubscriptionDimensionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.SubscriptionDimensionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SubscriptionDimensionsClient) ListPreparer(ctx context.Context, filter string, expand string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/dimensions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionDimensionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SubscriptionDimensionsClient) ListResponder(resp *http.Response) (result DimensionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
