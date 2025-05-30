// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// ConnectedEnvironmentsStoragesClient contains the methods for the ConnectedEnvironmentsStorages group.
// Don't use this type directly, use NewConnectedEnvironmentsStoragesClient() instead.
type ConnectedEnvironmentsStoragesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConnectedEnvironmentsStoragesClient creates a new instance of ConnectedEnvironmentsStoragesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectedEnvironmentsStoragesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectedEnvironmentsStoragesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectedEnvironmentsStoragesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update storage for a connectedEnvironment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the Environment.
//   - storageName - Name of the storage.
//   - storageEnvelope - Configuration details of storage.
//   - options - ConnectedEnvironmentsStoragesClientCreateOrUpdateOptions contains the optional parameters for the ConnectedEnvironmentsStoragesClient.CreateOrUpdate
//     method.
func (client *ConnectedEnvironmentsStoragesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, storageName string, storageEnvelope ConnectedEnvironmentStorage, options *ConnectedEnvironmentsStoragesClientCreateOrUpdateOptions) (ConnectedEnvironmentsStoragesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ConnectedEnvironmentsStoragesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, storageName, storageEnvelope, options)
	if err != nil {
		return ConnectedEnvironmentsStoragesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedEnvironmentsStoragesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedEnvironmentsStoragesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectedEnvironmentsStoragesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, storageName string, storageEnvelope ConnectedEnvironmentStorage, _ *ConnectedEnvironmentsStoragesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/storages/{storageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if storageName == "" {
		return nil, errors.New("parameter storageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageName}", url.PathEscape(storageName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, storageEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectedEnvironmentsStoragesClient) createOrUpdateHandleResponse(resp *http.Response) (ConnectedEnvironmentsStoragesClientCreateOrUpdateResponse, error) {
	result := ConnectedEnvironmentsStoragesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedEnvironmentStorage); err != nil {
		return ConnectedEnvironmentsStoragesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete storage for a connectedEnvironment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the Environment.
//   - storageName - Name of the storage.
//   - options - ConnectedEnvironmentsStoragesClientDeleteOptions contains the optional parameters for the ConnectedEnvironmentsStoragesClient.Delete
//     method.
func (client *ConnectedEnvironmentsStoragesClient) Delete(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, storageName string, options *ConnectedEnvironmentsStoragesClientDeleteOptions) (ConnectedEnvironmentsStoragesClientDeleteResponse, error) {
	var err error
	const operationName = "ConnectedEnvironmentsStoragesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, storageName, options)
	if err != nil {
		return ConnectedEnvironmentsStoragesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedEnvironmentsStoragesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedEnvironmentsStoragesClientDeleteResponse{}, err
	}
	return ConnectedEnvironmentsStoragesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectedEnvironmentsStoragesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, storageName string, _ *ConnectedEnvironmentsStoragesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/storages/{storageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if storageName == "" {
		return nil, errors.New("parameter storageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageName}", url.PathEscape(storageName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get storage for a connectedEnvironment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the Environment.
//   - storageName - Name of the storage.
//   - options - ConnectedEnvironmentsStoragesClientGetOptions contains the optional parameters for the ConnectedEnvironmentsStoragesClient.Get
//     method.
func (client *ConnectedEnvironmentsStoragesClient) Get(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, storageName string, options *ConnectedEnvironmentsStoragesClientGetOptions) (ConnectedEnvironmentsStoragesClientGetResponse, error) {
	var err error
	const operationName = "ConnectedEnvironmentsStoragesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, storageName, options)
	if err != nil {
		return ConnectedEnvironmentsStoragesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedEnvironmentsStoragesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedEnvironmentsStoragesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConnectedEnvironmentsStoragesClient) getCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, storageName string, _ *ConnectedEnvironmentsStoragesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/storages/{storageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if storageName == "" {
		return nil, errors.New("parameter storageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageName}", url.PathEscape(storageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectedEnvironmentsStoragesClient) getHandleResponse(resp *http.Response) (ConnectedEnvironmentsStoragesClientGetResponse, error) {
	result := ConnectedEnvironmentsStoragesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedEnvironmentStorage); err != nil {
		return ConnectedEnvironmentsStoragesClientGetResponse{}, err
	}
	return result, nil
}

// List - Get all storages for a connectedEnvironment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the Environment.
//   - options - ConnectedEnvironmentsStoragesClientListOptions contains the optional parameters for the ConnectedEnvironmentsStoragesClient.List
//     method.
func (client *ConnectedEnvironmentsStoragesClient) List(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsStoragesClientListOptions) (ConnectedEnvironmentsStoragesClientListResponse, error) {
	var err error
	const operationName = "ConnectedEnvironmentsStoragesClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, options)
	if err != nil {
		return ConnectedEnvironmentsStoragesClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedEnvironmentsStoragesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedEnvironmentsStoragesClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *ConnectedEnvironmentsStoragesClient) listCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, _ *ConnectedEnvironmentsStoragesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/storages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConnectedEnvironmentsStoragesClient) listHandleResponse(resp *http.Response) (ConnectedEnvironmentsStoragesClientListResponse, error) {
	result := ConnectedEnvironmentsStoragesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedEnvironmentStoragesCollection); err != nil {
		return ConnectedEnvironmentsStoragesClientListResponse{}, err
	}
	return result, nil
}
