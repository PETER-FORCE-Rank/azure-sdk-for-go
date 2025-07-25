// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armoracledatabase

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

// CloudExadataInfrastructuresClient contains the methods for the CloudExadataInfrastructures group.
// Don't use this type directly, use NewCloudExadataInfrastructuresClient() instead.
type CloudExadataInfrastructuresClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCloudExadataInfrastructuresClient creates a new instance of CloudExadataInfrastructuresClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCloudExadataInfrastructuresClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CloudExadataInfrastructuresClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CloudExadataInfrastructuresClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginAddStorageCapacity - Perform add storage capacity on exadata infra
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudexadatainfrastructurename - CloudExadataInfrastructure name
//   - options - CloudExadataInfrastructuresClientBeginAddStorageCapacityOptions contains the optional parameters for the CloudExadataInfrastructuresClient.BeginAddStorageCapacity
//     method.
func (client *CloudExadataInfrastructuresClient) BeginAddStorageCapacity(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, options *CloudExadataInfrastructuresClientBeginAddStorageCapacityOptions) (*runtime.Poller[CloudExadataInfrastructuresClientAddStorageCapacityResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.addStorageCapacity(ctx, resourceGroupName, cloudexadatainfrastructurename, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CloudExadataInfrastructuresClientAddStorageCapacityResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CloudExadataInfrastructuresClientAddStorageCapacityResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// AddStorageCapacity - Perform add storage capacity on exadata infra
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *CloudExadataInfrastructuresClient) addStorageCapacity(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, options *CloudExadataInfrastructuresClientBeginAddStorageCapacityOptions) (*http.Response, error) {
	var err error
	const operationName = "CloudExadataInfrastructuresClient.BeginAddStorageCapacity"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.addStorageCapacityCreateRequest(ctx, resourceGroupName, cloudexadatainfrastructurename, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// addStorageCapacityCreateRequest creates the AddStorageCapacity request.
func (client *CloudExadataInfrastructuresClient) addStorageCapacityCreateRequest(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, _ *CloudExadataInfrastructuresClientBeginAddStorageCapacityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/cloudExadataInfrastructures/{cloudexadatainfrastructurename}/addStorageCapacity"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudexadatainfrastructurename == "" {
		return nil, errors.New("parameter cloudexadatainfrastructurename cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudexadatainfrastructurename}", url.PathEscape(cloudexadatainfrastructurename))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginCreateOrUpdate - Create a CloudExadataInfrastructure
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudexadatainfrastructurename - CloudExadataInfrastructure name
//   - resource - Resource create parameters.
//   - options - CloudExadataInfrastructuresClientBeginCreateOrUpdateOptions contains the optional parameters for the CloudExadataInfrastructuresClient.BeginCreateOrUpdate
//     method.
func (client *CloudExadataInfrastructuresClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, resource CloudExadataInfrastructure, options *CloudExadataInfrastructuresClientBeginCreateOrUpdateOptions) (*runtime.Poller[CloudExadataInfrastructuresClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, cloudexadatainfrastructurename, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CloudExadataInfrastructuresClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CloudExadataInfrastructuresClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a CloudExadataInfrastructure
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *CloudExadataInfrastructuresClient) createOrUpdate(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, resource CloudExadataInfrastructure, options *CloudExadataInfrastructuresClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CloudExadataInfrastructuresClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, cloudexadatainfrastructurename, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CloudExadataInfrastructuresClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, resource CloudExadataInfrastructure, _ *CloudExadataInfrastructuresClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/cloudExadataInfrastructures/{cloudexadatainfrastructurename}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudexadatainfrastructurename == "" {
		return nil, errors.New("parameter cloudexadatainfrastructurename cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudexadatainfrastructurename}", url.PathEscape(cloudexadatainfrastructurename))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a CloudExadataInfrastructure
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudexadatainfrastructurename - CloudExadataInfrastructure name
//   - options - CloudExadataInfrastructuresClientBeginDeleteOptions contains the optional parameters for the CloudExadataInfrastructuresClient.BeginDelete
//     method.
func (client *CloudExadataInfrastructuresClient) BeginDelete(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, options *CloudExadataInfrastructuresClientBeginDeleteOptions) (*runtime.Poller[CloudExadataInfrastructuresClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, cloudexadatainfrastructurename, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CloudExadataInfrastructuresClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CloudExadataInfrastructuresClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a CloudExadataInfrastructure
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *CloudExadataInfrastructuresClient) deleteOperation(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, options *CloudExadataInfrastructuresClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "CloudExadataInfrastructuresClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, cloudexadatainfrastructurename, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CloudExadataInfrastructuresClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, _ *CloudExadataInfrastructuresClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/cloudExadataInfrastructures/{cloudexadatainfrastructurename}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudexadatainfrastructurename == "" {
		return nil, errors.New("parameter cloudexadatainfrastructurename cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudexadatainfrastructurename}", url.PathEscape(cloudexadatainfrastructurename))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a CloudExadataInfrastructure
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudexadatainfrastructurename - CloudExadataInfrastructure name
//   - options - CloudExadataInfrastructuresClientGetOptions contains the optional parameters for the CloudExadataInfrastructuresClient.Get
//     method.
func (client *CloudExadataInfrastructuresClient) Get(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, options *CloudExadataInfrastructuresClientGetOptions) (CloudExadataInfrastructuresClientGetResponse, error) {
	var err error
	const operationName = "CloudExadataInfrastructuresClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, cloudexadatainfrastructurename, options)
	if err != nil {
		return CloudExadataInfrastructuresClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CloudExadataInfrastructuresClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CloudExadataInfrastructuresClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CloudExadataInfrastructuresClient) getCreateRequest(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, _ *CloudExadataInfrastructuresClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/cloudExadataInfrastructures/{cloudexadatainfrastructurename}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudexadatainfrastructurename == "" {
		return nil, errors.New("parameter cloudexadatainfrastructurename cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudexadatainfrastructurename}", url.PathEscape(cloudexadatainfrastructurename))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CloudExadataInfrastructuresClient) getHandleResponse(resp *http.Response) (CloudExadataInfrastructuresClientGetResponse, error) {
	result := CloudExadataInfrastructuresClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudExadataInfrastructure); err != nil {
		return CloudExadataInfrastructuresClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List CloudExadataInfrastructure resources by resource group
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - CloudExadataInfrastructuresClientListByResourceGroupOptions contains the optional parameters for the CloudExadataInfrastructuresClient.NewListByResourceGroupPager
//     method.
func (client *CloudExadataInfrastructuresClient) NewListByResourceGroupPager(resourceGroupName string, options *CloudExadataInfrastructuresClientListByResourceGroupOptions) *runtime.Pager[CloudExadataInfrastructuresClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[CloudExadataInfrastructuresClientListByResourceGroupResponse]{
		More: func(page CloudExadataInfrastructuresClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CloudExadataInfrastructuresClientListByResourceGroupResponse) (CloudExadataInfrastructuresClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CloudExadataInfrastructuresClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return CloudExadataInfrastructuresClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CloudExadataInfrastructuresClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *CloudExadataInfrastructuresClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/cloudExadataInfrastructures"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CloudExadataInfrastructuresClient) listByResourceGroupHandleResponse(resp *http.Response) (CloudExadataInfrastructuresClientListByResourceGroupResponse, error) {
	result := CloudExadataInfrastructuresClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudExadataInfrastructureListResult); err != nil {
		return CloudExadataInfrastructuresClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List CloudExadataInfrastructure resources by subscription ID
//
// Generated from API version 2025-03-01
//   - options - CloudExadataInfrastructuresClientListBySubscriptionOptions contains the optional parameters for the CloudExadataInfrastructuresClient.NewListBySubscriptionPager
//     method.
func (client *CloudExadataInfrastructuresClient) NewListBySubscriptionPager(options *CloudExadataInfrastructuresClientListBySubscriptionOptions) *runtime.Pager[CloudExadataInfrastructuresClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[CloudExadataInfrastructuresClientListBySubscriptionResponse]{
		More: func(page CloudExadataInfrastructuresClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CloudExadataInfrastructuresClientListBySubscriptionResponse) (CloudExadataInfrastructuresClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CloudExadataInfrastructuresClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return CloudExadataInfrastructuresClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CloudExadataInfrastructuresClient) listBySubscriptionCreateRequest(ctx context.Context, _ *CloudExadataInfrastructuresClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/cloudExadataInfrastructures"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CloudExadataInfrastructuresClient) listBySubscriptionHandleResponse(resp *http.Response) (CloudExadataInfrastructuresClientListBySubscriptionResponse, error) {
	result := CloudExadataInfrastructuresClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudExadataInfrastructureListResult); err != nil {
		return CloudExadataInfrastructuresClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a CloudExadataInfrastructure
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudexadatainfrastructurename - CloudExadataInfrastructure name
//   - properties - The resource properties to be updated.
//   - options - CloudExadataInfrastructuresClientBeginUpdateOptions contains the optional parameters for the CloudExadataInfrastructuresClient.BeginUpdate
//     method.
func (client *CloudExadataInfrastructuresClient) BeginUpdate(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, properties CloudExadataInfrastructureUpdate, options *CloudExadataInfrastructuresClientBeginUpdateOptions) (*runtime.Poller[CloudExadataInfrastructuresClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, cloudexadatainfrastructurename, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CloudExadataInfrastructuresClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CloudExadataInfrastructuresClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a CloudExadataInfrastructure
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *CloudExadataInfrastructuresClient) update(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, properties CloudExadataInfrastructureUpdate, options *CloudExadataInfrastructuresClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CloudExadataInfrastructuresClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, cloudexadatainfrastructurename, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *CloudExadataInfrastructuresClient) updateCreateRequest(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, properties CloudExadataInfrastructureUpdate, _ *CloudExadataInfrastructuresClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/cloudExadataInfrastructures/{cloudexadatainfrastructurename}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudexadatainfrastructurename == "" {
		return nil, errors.New("parameter cloudexadatainfrastructurename cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudexadatainfrastructurename}", url.PathEscape(cloudexadatainfrastructurename))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
