// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

type serviceClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client serviceClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// FilterBlobs - The Filter Blobs operation enables callers to list blobs across all containers whose tags match a given search expression. Filter blobs
// searches across all containers within a storage account but can
// be scoped within the expression to a single container.
func (client serviceClient) FilterBlobs(ctx context.Context, options *ServiceFilterBlobsOptions) (FilterBlobSegmentResponse, error) {
	req, err := client.filterBlobsCreateRequest(ctx, options)
	if err != nil {
		return FilterBlobSegmentResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return FilterBlobSegmentResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FilterBlobSegmentResponse{}, client.filterBlobsHandleError(resp)
	}
	result, err := client.filterBlobsHandleResponse(resp)
	if err != nil {
		return FilterBlobSegmentResponse{}, err
	}
	return result, nil
}

// filterBlobsCreateRequest creates the FilterBlobs request.
func (client serviceClient) filterBlobsCreateRequest(ctx context.Context, options *ServiceFilterBlobsOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("comp", "blobs")
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Where != nil {
		query.Set("where", *options.Where)
	}
	if options != nil && options.Marker != nil {
		query.Set("marker", *options.Marker)
	}
	if options != nil && options.Maxresults != nil {
		query.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-12-12")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// filterBlobsHandleResponse handles the FilterBlobs response.
func (client serviceClient) filterBlobsHandleResponse(resp *azcore.Response) (FilterBlobSegmentResponse, error) {
	result := FilterBlobSegmentResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return FilterBlobSegmentResponse{}, err
		}
		result.Date = &date
	}
	err := resp.UnmarshalAsXML(&result.EnumerationResults)
	return result, err
}

// filterBlobsHandleError handles the FilterBlobs error response.
func (client serviceClient) filterBlobsHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetAccountInfo - Returns the sku name and account kind
func (client serviceClient) GetAccountInfo(ctx context.Context, options *ServiceGetAccountInfoOptions) (ServiceGetAccountInfoResponse, error) {
	req, err := client.getAccountInfoCreateRequest(ctx, options)
	if err != nil {
		return ServiceGetAccountInfoResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ServiceGetAccountInfoResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServiceGetAccountInfoResponse{}, client.getAccountInfoHandleError(resp)
	}
	result, err := client.getAccountInfoHandleResponse(resp)
	if err != nil {
		return ServiceGetAccountInfoResponse{}, err
	}
	return result, nil
}

// getAccountInfoCreateRequest creates the GetAccountInfo request.
func (client serviceClient) getAccountInfoCreateRequest(ctx context.Context, options *ServiceGetAccountInfoOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("restype", "account")
	query.Set("comp", "properties")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-12-12")
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// getAccountInfoHandleResponse handles the GetAccountInfo response.
func (client serviceClient) getAccountInfoHandleResponse(resp *azcore.Response) (ServiceGetAccountInfoResponse, error) {
	result := ServiceGetAccountInfoResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ServiceGetAccountInfoResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-sku-name"); val != "" {
		result.SKUName = (*SKUName)(&val)
	}
	if val := resp.Header.Get("x-ms-account-kind"); val != "" {
		result.AccountKind = (*AccountKind)(&val)
	}
	return result, nil
}

// getAccountInfoHandleError handles the GetAccountInfo error response.
func (client serviceClient) getAccountInfoHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetProperties - gets the properties of a storage account's Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing)
// rules.
func (client serviceClient) GetProperties(ctx context.Context, options *ServiceGetPropertiesOptions) (StorageServicePropertiesResponse, error) {
	req, err := client.getPropertiesCreateRequest(ctx, options)
	if err != nil {
		return StorageServicePropertiesResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StorageServicePropertiesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StorageServicePropertiesResponse{}, client.getPropertiesHandleError(resp)
	}
	result, err := client.getPropertiesHandleResponse(resp)
	if err != nil {
		return StorageServicePropertiesResponse{}, err
	}
	return result, nil
}

// getPropertiesCreateRequest creates the GetProperties request.
func (client serviceClient) getPropertiesCreateRequest(ctx context.Context, options *ServiceGetPropertiesOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("restype", "service")
	query.Set("comp", "properties")
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-12-12")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// getPropertiesHandleResponse handles the GetProperties response.
func (client serviceClient) getPropertiesHandleResponse(resp *azcore.Response) (StorageServicePropertiesResponse, error) {
	result := StorageServicePropertiesResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	err := resp.UnmarshalAsXML(&result.StorageServiceProperties)
	return result, err
}

// getPropertiesHandleError handles the GetProperties error response.
func (client serviceClient) getPropertiesHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetStatistics - Retrieves statistics related to replication for the Blob service. It is only available on the secondary location endpoint when read-access
// geo-redundant replication is enabled for the storage account.
func (client serviceClient) GetStatistics(ctx context.Context, options *ServiceGetStatisticsOptions) (StorageServiceStatsResponse, error) {
	req, err := client.getStatisticsCreateRequest(ctx, options)
	if err != nil {
		return StorageServiceStatsResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StorageServiceStatsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StorageServiceStatsResponse{}, client.getStatisticsHandleError(resp)
	}
	result, err := client.getStatisticsHandleResponse(resp)
	if err != nil {
		return StorageServiceStatsResponse{}, err
	}
	return result, nil
}

// getStatisticsCreateRequest creates the GetStatistics request.
func (client serviceClient) getStatisticsCreateRequest(ctx context.Context, options *ServiceGetStatisticsOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("restype", "service")
	query.Set("comp", "stats")
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-12-12")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// getStatisticsHandleResponse handles the GetStatistics response.
func (client serviceClient) getStatisticsHandleResponse(resp *azcore.Response) (StorageServiceStatsResponse, error) {
	result := StorageServiceStatsResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return StorageServiceStatsResponse{}, err
		}
		result.Date = &date
	}
	err := resp.UnmarshalAsXML(&result.StorageServiceStats)
	return result, err
}

// getStatisticsHandleError handles the GetStatistics error response.
func (client serviceClient) getStatisticsHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetUserDelegationKey - Retrieves a user delegation key for the Blob service. This is only a valid operation when using bearer token authentication.
func (client serviceClient) GetUserDelegationKey(ctx context.Context, keyInfo KeyInfo, options *ServiceGetUserDelegationKeyOptions) (UserDelegationKeyResponse, error) {
	req, err := client.getUserDelegationKeyCreateRequest(ctx, keyInfo, options)
	if err != nil {
		return UserDelegationKeyResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return UserDelegationKeyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return UserDelegationKeyResponse{}, client.getUserDelegationKeyHandleError(resp)
	}
	result, err := client.getUserDelegationKeyHandleResponse(resp)
	if err != nil {
		return UserDelegationKeyResponse{}, err
	}
	return result, nil
}

// getUserDelegationKeyCreateRequest creates the GetUserDelegationKey request.
func (client serviceClient) getUserDelegationKeyCreateRequest(ctx context.Context, keyInfo KeyInfo, options *ServiceGetUserDelegationKeyOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPost, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("restype", "service")
	query.Set("comp", "userdelegationkey")
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-12-12")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(keyInfo)
}

// getUserDelegationKeyHandleResponse handles the GetUserDelegationKey response.
func (client serviceClient) getUserDelegationKeyHandleResponse(resp *azcore.Response) (UserDelegationKeyResponse, error) {
	result := UserDelegationKeyResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return UserDelegationKeyResponse{}, err
		}
		result.Date = &date
	}
	err := resp.UnmarshalAsXML(&result.UserDelegationKey)
	return result, err
}

// getUserDelegationKeyHandleError handles the GetUserDelegationKey error response.
func (client serviceClient) getUserDelegationKeyHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListContainersSegment - The List Containers Segment operation returns a list of the containers under the specified account
func (client serviceClient) ListContainersSegment(options *ServiceListContainersSegmentOptions) ListContainersSegmentResponsePager {
	return &listContainersSegmentResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listContainersSegmentCreateRequest(ctx, options)
		},
		responder: client.listContainersSegmentHandleResponse,
		errorer:   client.listContainersSegmentHandleError,
		advancer: func(ctx context.Context, resp ListContainersSegmentResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.EnumerationResults.NextMarker)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listContainersSegmentCreateRequest creates the ListContainersSegment request.
func (client serviceClient) listContainersSegmentCreateRequest(ctx context.Context, options *ServiceListContainersSegmentOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("comp", "list")
	if options != nil && options.Prefix != nil {
		query.Set("prefix", *options.Prefix)
	}
	if options != nil && options.Marker != nil {
		query.Set("marker", *options.Marker)
	}
	if options != nil && options.Maxresults != nil {
		query.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	if options != nil && options.Include != nil {
		query.Set("include", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(*options.Include), "[]")), ","))
	}
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-12-12")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// listContainersSegmentHandleResponse handles the ListContainersSegment response.
func (client serviceClient) listContainersSegmentHandleResponse(resp *azcore.Response) (ListContainersSegmentResponseResponse, error) {
	result := ListContainersSegmentResponseResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	err := resp.UnmarshalAsXML(&result.EnumerationResults)
	return result, err
}

// listContainersSegmentHandleError handles the ListContainersSegment error response.
func (client serviceClient) listContainersSegmentHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// SetProperties - Sets properties for a storage account's Blob service endpoint, including properties for Storage Analytics and CORS (Cross-Origin Resource
// Sharing) rules
func (client serviceClient) SetProperties(ctx context.Context, storageServiceProperties StorageServiceProperties, options *ServiceSetPropertiesOptions) (ServiceSetPropertiesResponse, error) {
	req, err := client.setPropertiesCreateRequest(ctx, storageServiceProperties, options)
	if err != nil {
		return ServiceSetPropertiesResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ServiceSetPropertiesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return ServiceSetPropertiesResponse{}, client.setPropertiesHandleError(resp)
	}
	result, err := client.setPropertiesHandleResponse(resp)
	if err != nil {
		return ServiceSetPropertiesResponse{}, err
	}
	return result, nil
}

// setPropertiesCreateRequest creates the SetProperties request.
func (client serviceClient) setPropertiesCreateRequest(ctx context.Context, storageServiceProperties StorageServiceProperties, options *ServiceSetPropertiesOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("restype", "service")
	query.Set("comp", "properties")
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-12-12")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(storageServiceProperties)
}

// setPropertiesHandleResponse handles the SetProperties response.
func (client serviceClient) setPropertiesHandleResponse(resp *azcore.Response) (ServiceSetPropertiesResponse, error) {
	result := ServiceSetPropertiesResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return result, nil
}

// setPropertiesHandleError handles the SetProperties error response.
func (client serviceClient) setPropertiesHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// SubmitBatch - The Batch operation allows multiple API calls to be embedded into a single HTTP request.
func (client serviceClient) SubmitBatch(ctx context.Context, contentLength int64, multipartContentType string, body azcore.ReadSeekCloser, options *ServiceSubmitBatchOptions) (ServiceSubmitBatchResponse, error) {
	req, err := client.submitBatchCreateRequest(ctx, contentLength, multipartContentType, body, options)
	if err != nil {
		return ServiceSubmitBatchResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ServiceSubmitBatchResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServiceSubmitBatchResponse{}, client.submitBatchHandleError(resp)
	}
	result, err := client.submitBatchHandleResponse(resp)
	if err != nil {
		return ServiceSubmitBatchResponse{}, err
	}
	return result, nil
}

// submitBatchCreateRequest creates the SubmitBatch request.
func (client serviceClient) submitBatchCreateRequest(ctx context.Context, contentLength int64, multipartContentType string, body azcore.ReadSeekCloser, options *ServiceSubmitBatchOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPost, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("comp", "batch")
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.SkipBodyDownload()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	req.Header.Set("Content-Type", multipartContentType)
	req.Header.Set("x-ms-version", "2019-12-12")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(body)
}

// submitBatchHandleResponse handles the SubmitBatch response.
func (client serviceClient) submitBatchHandleResponse(resp *azcore.Response) (ServiceSubmitBatchResponse, error) {
	result := ServiceSubmitBatchResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return result, nil
}

// submitBatchHandleError handles the SubmitBatch error response.
func (client serviceClient) submitBatchHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}