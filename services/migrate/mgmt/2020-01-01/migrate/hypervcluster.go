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

// HyperVClusterClient is the discover your workloads for Azure.
type HyperVClusterClient struct {
	BaseClient
}

// NewHyperVClusterClient creates an instance of the HyperVClusterClient client.
func NewHyperVClusterClient() HyperVClusterClient {
	return NewHyperVClusterClientWithBaseURI(DefaultBaseURI)
}

// NewHyperVClusterClientWithBaseURI creates an instance of the HyperVClusterClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewHyperVClusterClientWithBaseURI(baseURI string) HyperVClusterClient {
	return HyperVClusterClient{NewWithBaseURI(baseURI)}
}

// GetAllClustersInSite sends the get all clusters in site request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// APIVersion - the API version to use for this operation.
func (client HyperVClusterClient) GetAllClustersInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (result HyperVClusterCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HyperVClusterClient.GetAllClustersInSite")
		defer func() {
			sc := -1
			if result.hvcc.Response.Response != nil {
				sc = result.hvcc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getAllClustersInSiteNextResults
	req, err := client.GetAllClustersInSitePreparer(ctx, subscriptionID, resourceGroupName, siteName, APIVersion, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVClusterClient", "GetAllClustersInSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllClustersInSiteSender(req)
	if err != nil {
		result.hvcc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.HyperVClusterClient", "GetAllClustersInSite", resp, "Failure sending request")
		return
	}

	result.hvcc, err = client.GetAllClustersInSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVClusterClient", "GetAllClustersInSite", resp, "Failure responding to request")
		return
	}
	if result.hvcc.hasNextLink() && result.hvcc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// GetAllClustersInSitePreparer prepares the GetAllClustersInSite request.
func (client HyperVClusterClient) GetAllClustersInSitePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (*http.Request, error) {
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}/clusters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllClustersInSiteSender sends the GetAllClustersInSite request. The method will close the
// http.Response Body if it receives an error.
func (client HyperVClusterClient) GetAllClustersInSiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAllClustersInSiteResponder handles the response to the GetAllClustersInSite request. The method always
// closes the http.Response Body.
func (client HyperVClusterClient) GetAllClustersInSiteResponder(resp *http.Response) (result HyperVClusterCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getAllClustersInSiteNextResults retrieves the next set of results, if any.
func (client HyperVClusterClient) getAllClustersInSiteNextResults(ctx context.Context, lastResults HyperVClusterCollection) (result HyperVClusterCollection, err error) {
	req, err := lastResults.hyperVClusterCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "migrate.HyperVClusterClient", "getAllClustersInSiteNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetAllClustersInSiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "migrate.HyperVClusterClient", "getAllClustersInSiteNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetAllClustersInSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVClusterClient", "getAllClustersInSiteNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// GetAllClustersInSiteComplete enumerates all values, automatically crossing page boundaries as required.
func (client HyperVClusterClient) GetAllClustersInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (result HyperVClusterCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HyperVClusterClient.GetAllClustersInSite")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetAllClustersInSite(ctx, subscriptionID, resourceGroupName, siteName, APIVersion, filter)
	return
}

// GetCluster sends the get cluster request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// clusterName - cluster ARM name.
// APIVersion - the API version to use for this operation.
func (client HyperVClusterClient) GetCluster(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, clusterName string, APIVersion string) (result HyperVCluster, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HyperVClusterClient.GetCluster")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetClusterPreparer(ctx, subscriptionID, resourceGroupName, siteName, clusterName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVClusterClient", "GetCluster", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetClusterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.HyperVClusterClient", "GetCluster", resp, "Failure sending request")
		return
	}

	result, err = client.GetClusterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVClusterClient", "GetCluster", resp, "Failure responding to request")
		return
	}

	return
}

// GetClusterPreparer prepares the GetCluster request.
func (client HyperVClusterClient) GetClusterPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, clusterName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}/clusters/{clusterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetClusterSender sends the GetCluster request. The method will close the
// http.Response Body if it receives an error.
func (client HyperVClusterClient) GetClusterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetClusterResponder handles the response to the GetCluster request. The method always
// closes the http.Response Body.
func (client HyperVClusterClient) GetClusterResponder(resp *http.Response) (result HyperVCluster, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutCluster sends the put cluster request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// clusterName - cluster ARM name.
// body - put cluster body.
// APIVersion - the API version to use for this operation.
func (client HyperVClusterClient) PutCluster(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, clusterName string, body HyperVCluster, APIVersion string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HyperVClusterClient.PutCluster")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutClusterPreparer(ctx, subscriptionID, resourceGroupName, siteName, clusterName, body, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVClusterClient", "PutCluster", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutClusterSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "migrate.HyperVClusterClient", "PutCluster", resp, "Failure sending request")
		return
	}

	result, err = client.PutClusterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.HyperVClusterClient", "PutCluster", resp, "Failure responding to request")
		return
	}

	return
}

// PutClusterPreparer prepares the PutCluster request.
func (client HyperVClusterClient) PutClusterPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, clusterName string, body HyperVCluster, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}/clusters/{clusterName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutClusterSender sends the PutCluster request. The method will close the
// http.Response Body if it receives an error.
func (client HyperVClusterClient) PutClusterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PutClusterResponder handles the response to the PutCluster request. The method always
// closes the http.Response Body.
func (client HyperVClusterClient) PutClusterResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
