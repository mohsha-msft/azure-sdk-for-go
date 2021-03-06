package runtime

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
	"github.com/satori/go.uuid"
	"net/http"
)

// PredictionClient is the client for the Prediction methods of the Runtime service.
type PredictionClient struct {
	BaseClient
}

// NewPredictionClient creates an instance of the PredictionClient client.
func NewPredictionClient(endpoint string) PredictionClient {
	return PredictionClient{New(endpoint)}
}

// Resolve gets predictions for a given utterance, in the form of intents and entities. The current maximum query size
// is 500 characters.
// Parameters:
// appID - the LUIS application ID (Guid).
// query - the utterance to predict.
// timezoneOffset - the timezone offset for the location of the request.
// verbose - if true, return all intents instead of just the top scoring intent.
// staging - use the staging endpoint slot.
// spellCheck - enable spell checking.
// bingSpellCheckSubscriptionKey - the subscription key to use when enabling Bing spell check
// logParameter - log query (default is true)
func (client PredictionClient) Resolve(ctx context.Context, appID uuid.UUID, query string, timezoneOffset *float64, verbose *bool, staging *bool, spellCheck *bool, bingSpellCheckSubscriptionKey string, logParameter *bool) (result LuisResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PredictionClient.Resolve")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: query,
			Constraints: []validation.Constraint{{Target: "query", Name: validation.MaxLength, Rule: 500, Chain: nil}}}}); err != nil {
		return result, validation.NewError("runtime.PredictionClient", "Resolve", err.Error())
	}

	req, err := client.ResolvePreparer(ctx, appID, query, timezoneOffset, verbose, staging, spellCheck, bingSpellCheckSubscriptionKey, logParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "runtime.PredictionClient", "Resolve", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResolveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "runtime.PredictionClient", "Resolve", resp, "Failure sending request")
		return
	}

	result, err = client.ResolveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "runtime.PredictionClient", "Resolve", resp, "Failure responding to request")
		return
	}

	return
}

// ResolvePreparer prepares the Resolve request.
func (client PredictionClient) ResolvePreparer(ctx context.Context, appID uuid.UUID, query string, timezoneOffset *float64, verbose *bool, staging *bool, spellCheck *bool, bingSpellCheckSubscriptionKey string, logParameter *bool) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId": autorest.Encode("path", appID),
	}

	queryParameters := map[string]interface{}{}
	if timezoneOffset != nil {
		queryParameters["timezoneOffset"] = autorest.Encode("query", *timezoneOffset)
	}
	if verbose != nil {
		queryParameters["verbose"] = autorest.Encode("query", *verbose)
	}
	if staging != nil {
		queryParameters["staging"] = autorest.Encode("query", *staging)
	}
	if spellCheck != nil {
		queryParameters["spellCheck"] = autorest.Encode("query", *spellCheck)
	}
	if len(bingSpellCheckSubscriptionKey) > 0 {
		queryParameters["bing-spell-check-subscription-key"] = autorest.Encode("query", bingSpellCheckSubscriptionKey)
	}
	if logParameter != nil {
		queryParameters["log"] = autorest.Encode("query", *logParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}", pathParameters),
		autorest.WithJSON(query),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResolveSender sends the Resolve request. The method will close the
// http.Response Body if it receives an error.
func (client PredictionClient) ResolveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ResolveResponder handles the response to the Resolve request. The method always
// closes the http.Response Body.
func (client PredictionClient) ResolveResponder(resp *http.Response) (result LuisResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
