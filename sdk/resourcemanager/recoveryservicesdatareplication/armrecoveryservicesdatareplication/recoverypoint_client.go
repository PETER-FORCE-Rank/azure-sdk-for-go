// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armrecoveryservicesdatareplication

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

// RecoveryPointClient contains the methods for the RecoveryPoint group.
// Don't use this type directly, use NewRecoveryPointClient() instead.
type RecoveryPointClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRecoveryPointClient creates a new instance of RecoveryPointClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRecoveryPointClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RecoveryPointClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RecoveryPointClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the details of the recovery point of a protected item.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The vault name.
//   - protectedItemName - The protected item name.
//   - recoveryPointName - The recovery point name.
//   - options - RecoveryPointClientGetOptions contains the optional parameters for the RecoveryPointClient.Get method.
func (client *RecoveryPointClient) Get(ctx context.Context, resourceGroupName string, vaultName string, protectedItemName string, recoveryPointName string, options *RecoveryPointClientGetOptions) (RecoveryPointClientGetResponse, error) {
	var err error
	const operationName = "RecoveryPointClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, protectedItemName, recoveryPointName, options)
	if err != nil {
		return RecoveryPointClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecoveryPointClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RecoveryPointClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RecoveryPointClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, protectedItemName string, recoveryPointName string, _ *RecoveryPointClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults/{vaultName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if protectedItemName == "" {
		return nil, errors.New("parameter protectedItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectedItemName}", url.PathEscape(protectedItemName))
	if recoveryPointName == "" {
		return nil, errors.New("parameter recoveryPointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recoveryPointName}", url.PathEscape(recoveryPointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecoveryPointClient) getHandleResponse(resp *http.Response) (RecoveryPointClientGetResponse, error) {
	result := RecoveryPointClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoveryPointModel); err != nil {
		return RecoveryPointClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of recovery points of the given protected item.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The vault name.
//   - protectedItemName - The protected item name.
//   - options - RecoveryPointClientListOptions contains the optional parameters for the RecoveryPointClient.NewListPager method.
func (client *RecoveryPointClient) NewListPager(resourceGroupName string, vaultName string, protectedItemName string, options *RecoveryPointClientListOptions) *runtime.Pager[RecoveryPointClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RecoveryPointClientListResponse]{
		More: func(page RecoveryPointClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RecoveryPointClientListResponse) (RecoveryPointClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RecoveryPointClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, vaultName, protectedItemName, options)
			}, nil)
			if err != nil {
				return RecoveryPointClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RecoveryPointClient) listCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, protectedItemName string, _ *RecoveryPointClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults/{vaultName}/protectedItems/{protectedItemName}/recoveryPoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if protectedItemName == "" {
		return nil, errors.New("parameter protectedItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectedItemName}", url.PathEscape(protectedItemName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RecoveryPointClient) listHandleResponse(resp *http.Response) (RecoveryPointClientListResponse, error) {
	result := RecoveryPointClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoveryPointModelListResult); err != nil {
		return RecoveryPointClientListResponse{}, err
	}
	return result, nil
}
