//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchaos

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CapabilityTypesClient contains the methods for the CapabilityTypes group.
// Don't use this type directly, use NewCapabilityTypesClient() instead.
type CapabilityTypesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCapabilityTypesClient creates a new instance of CapabilityTypesClient with the specified values.
//   - subscriptionID - GUID that represents an Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCapabilityTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CapabilityTypesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CapabilityTypesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a Capability Type resource for given Target Type and location.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - locationName - String that represents a Location resource name.
//   - targetTypeName - String that represents a Target Type resource name.
//   - capabilityTypeName - String that represents a Capability Type resource name.
//   - options - CapabilityTypesClientGetOptions contains the optional parameters for the CapabilityTypesClient.Get method.
func (client *CapabilityTypesClient) Get(ctx context.Context, locationName string, targetTypeName string, capabilityTypeName string, options *CapabilityTypesClientGetOptions) (CapabilityTypesClientGetResponse, error) {
	var err error
	const operationName = "CapabilityTypesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, locationName, targetTypeName, capabilityTypeName, options)
	if err != nil {
		return CapabilityTypesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapabilityTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CapabilityTypesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CapabilityTypesClient) getCreateRequest(ctx context.Context, locationName string, targetTypeName string, capabilityTypeName string, options *CapabilityTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Chaos/locations/{locationName}/targetTypes/{targetTypeName}/capabilityTypes/{capabilityTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if targetTypeName == "" {
		return nil, errors.New("parameter targetTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetTypeName}", url.PathEscape(targetTypeName))
	if capabilityTypeName == "" {
		return nil, errors.New("parameter capabilityTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capabilityTypeName}", url.PathEscape(capabilityTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CapabilityTypesClient) getHandleResponse(resp *http.Response) (CapabilityTypesClientGetResponse, error) {
	result := CapabilityTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilityType); err != nil {
		return CapabilityTypesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get a list of Capability Type resources for given Target Type and location.
//
// Generated from API version 2024-01-01
//   - locationName - String that represents a Location resource name.
//   - targetTypeName - String that represents a Target Type resource name.
//   - options - CapabilityTypesClientListOptions contains the optional parameters for the CapabilityTypesClient.NewListPager
//     method.
func (client *CapabilityTypesClient) NewListPager(locationName string, targetTypeName string, options *CapabilityTypesClientListOptions) *runtime.Pager[CapabilityTypesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CapabilityTypesClientListResponse]{
		More: func(page CapabilityTypesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CapabilityTypesClientListResponse) (CapabilityTypesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CapabilityTypesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, locationName, targetTypeName, options)
			}, nil)
			if err != nil {
				return CapabilityTypesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *CapabilityTypesClient) listCreateRequest(ctx context.Context, locationName string, targetTypeName string, options *CapabilityTypesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Chaos/locations/{locationName}/targetTypes/{targetTypeName}/capabilityTypes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if targetTypeName == "" {
		return nil, errors.New("parameter targetTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetTypeName}", url.PathEscape(targetTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CapabilityTypesClient) listHandleResponse(resp *http.Response) (CapabilityTypesClientListResponse, error) {
	result := CapabilityTypesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilityTypeListResult); err != nil {
		return CapabilityTypesClientListResponse{}, err
	}
	return result, nil
}
