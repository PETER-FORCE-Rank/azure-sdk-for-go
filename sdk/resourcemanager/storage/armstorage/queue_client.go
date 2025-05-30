// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

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

// QueueClient contains the methods for the Queue group.
// Don't use this type directly, use NewQueueClient() instead.
type QueueClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewQueueClient creates a new instance of QueueClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewQueueClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*QueueClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &QueueClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates a new queue with the specified queue name, under the specified account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - queueName - A queue name must be unique within a storage account and must be between 3 and 63 characters.The name must
//     comprise of lowercase alphanumeric and dash(-) characters only, it should begin and end with
//     an alphanumeric character and it cannot have two consecutive dash(-) characters.
//   - queue - Queue properties and metadata to be created with
//   - options - QueueClientCreateOptions contains the optional parameters for the QueueClient.Create method.
func (client *QueueClient) Create(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue Queue, options *QueueClientCreateOptions) (QueueClientCreateResponse, error) {
	var err error
	const operationName = "QueueClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, queueName, queue, options)
	if err != nil {
		return QueueClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueueClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QueueClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *QueueClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue Queue, _ *QueueClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, queue); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *QueueClient) createHandleResponse(resp *http.Response) (QueueClientCreateResponse, error) {
	result := QueueClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Queue); err != nil {
		return QueueClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the queue with the specified queue name, under the specified account if it exists.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - queueName - A queue name must be unique within a storage account and must be between 3 and 63 characters.The name must
//     comprise of lowercase alphanumeric and dash(-) characters only, it should begin and end with
//     an alphanumeric character and it cannot have two consecutive dash(-) characters.
//   - options - QueueClientDeleteOptions contains the optional parameters for the QueueClient.Delete method.
func (client *QueueClient) Delete(ctx context.Context, resourceGroupName string, accountName string, queueName string, options *QueueClientDeleteOptions) (QueueClientDeleteResponse, error) {
	var err error
	const operationName = "QueueClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, queueName, options)
	if err != nil {
		return QueueClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueueClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueueClientDeleteResponse{}, err
	}
	return QueueClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *QueueClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, queueName string, _ *QueueClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the queue with the specified queue name, under the specified account if it exists.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - queueName - A queue name must be unique within a storage account and must be between 3 and 63 characters.The name must
//     comprise of lowercase alphanumeric and dash(-) characters only, it should begin and end with
//     an alphanumeric character and it cannot have two consecutive dash(-) characters.
//   - options - QueueClientGetOptions contains the optional parameters for the QueueClient.Get method.
func (client *QueueClient) Get(ctx context.Context, resourceGroupName string, accountName string, queueName string, options *QueueClientGetOptions) (QueueClientGetResponse, error) {
	var err error
	const operationName = "QueueClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, queueName, options)
	if err != nil {
		return QueueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QueueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *QueueClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, queueName string, _ *QueueClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
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
func (client *QueueClient) getHandleResponse(resp *http.Response) (QueueClientGetResponse, error) {
	result := QueueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Queue); err != nil {
		return QueueClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of all the queues under the specified storage account
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - options - QueueClientListOptions contains the optional parameters for the QueueClient.NewListPager method.
func (client *QueueClient) NewListPager(resourceGroupName string, accountName string, options *QueueClientListOptions) *runtime.Pager[QueueClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[QueueClientListResponse]{
		More: func(page QueueClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *QueueClientListResponse) (QueueClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "QueueClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return QueueClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *QueueClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *QueueClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Maxpagesize != nil {
		reqQP.Set("$maxpagesize", *options.Maxpagesize)
	}
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *QueueClient) listHandleResponse(resp *http.Response) (QueueClientListResponse, error) {
	result := QueueClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListQueueResource); err != nil {
		return QueueClientListResponse{}, err
	}
	return result, nil
}

// Update - Creates a new queue with the specified queue name, under the specified account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - queueName - A queue name must be unique within a storage account and must be between 3 and 63 characters.The name must
//     comprise of lowercase alphanumeric and dash(-) characters only, it should begin and end with
//     an alphanumeric character and it cannot have two consecutive dash(-) characters.
//   - queue - Queue properties and metadata to be created with
//   - options - QueueClientUpdateOptions contains the optional parameters for the QueueClient.Update method.
func (client *QueueClient) Update(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue Queue, options *QueueClientUpdateOptions) (QueueClientUpdateResponse, error) {
	var err error
	const operationName = "QueueClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, queueName, queue, options)
	if err != nil {
		return QueueClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueueClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QueueClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *QueueClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue Queue, _ *QueueClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, queue); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *QueueClient) updateHandleResponse(resp *http.Response) (QueueClientUpdateResponse, error) {
	result := QueueClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Queue); err != nil {
		return QueueClientUpdateResponse{}, err
	}
	return result, nil
}
