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

// QueryClient is the client for the Query methods of the Costmanagement service.
type QueryClient struct {
	BaseClient
}

// NewQueryClient creates an instance of the QueryClient client.
func NewQueryClient(subscriptionID string) QueryClient {
	return NewQueryClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewQueryClientWithBaseURI creates an instance of the QueryClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewQueryClientWithBaseURI(baseURI string, subscriptionID string) QueryClient {
	return QueryClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// UsageByScope query the usage data for scope defined.
// Parameters:
// scope - the scope associated with query operations. This includes '/subscriptions/{subscriptionId}/' for
// subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup
// scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department
// scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount scope and '/providers/Microsoft.Management/managementGroups/{managementGroupId} for
// Management Group scope..
// parameters - parameters supplied to the CreateOrUpdate Query Config operation.
func (client QueryClient) UsageByScope(ctx context.Context, scope string, parameters QueryDefinition) (result QueryResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryClient.UsageByScope")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TimePeriod", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.TimePeriod.From", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.TimePeriod.To", Name: validation.Null, Rule: true, Chain: nil},
					}},
				{Target: "parameters.Dataset", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.MaxItems, Rule: 2, Chain: nil}}},
						{Target: "parameters.Dataset.Filter", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Or", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Or", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Not", Name: validation.Null, Rule: false, Chain: nil},
								{Target: "parameters.Dataset.Filter.Dimension", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Dimension.Operator", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
								{Target: "parameters.Dataset.Filter.Tag", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Tag.Operator", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
							}},
					}}}}}); err != nil {
		return result, validation.NewError("costmanagement.QueryClient", "UsageByScope", err.Error())
	}

	req, err := client.UsageByScopePreparer(ctx, scope, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.UsageByScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByScope", resp, "Failure sending request")
		return
	}

	result, err = client.UsageByScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByScope", resp, "Failure responding to request")
	}

	return
}

// UsageByScopePreparer prepares the UsageByScope request.
func (client QueryClient) UsageByScopePreparer(ctx context.Context, scope string, parameters QueryDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2019-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CostManagement/query", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UsageByScopeSender sends the UsageByScope request. The method will close the
// http.Response Body if it receives an error.
func (client QueryClient) UsageByScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UsageByScopeResponder handles the response to the UsageByScope request. The method always
// closes the http.Response Body.
func (client QueryClient) UsageByScopeResponder(resp *http.Response) (result QueryResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
