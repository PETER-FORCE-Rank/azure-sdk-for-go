// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armelasticsan

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

// VolumeGroupsClient contains the methods for the VolumeGroups group.
// Don't use this type directly, use NewVolumeGroupsClient() instead.
type VolumeGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVolumeGroupsClient creates a new instance of VolumeGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVolumeGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VolumeGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VolumeGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a Volume Group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - elasticSanName - The name of the ElasticSan.
//   - volumeGroupName - The name of the VolumeGroup.
//   - parameters - Volume Group object.
//   - options - VolumeGroupsClientBeginCreateOptions contains the optional parameters for the VolumeGroupsClient.BeginCreate
//     method.
func (client *VolumeGroupsClient) BeginCreate(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, parameters VolumeGroup, options *VolumeGroupsClientBeginCreateOptions) (*runtime.Poller[VolumeGroupsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, elasticSanName, volumeGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VolumeGroupsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VolumeGroupsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a Volume Group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01-preview
func (client *VolumeGroupsClient) create(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, parameters VolumeGroup, options *VolumeGroupsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "VolumeGroupsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, elasticSanName, volumeGroupName, parameters, options)
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

// createCreateRequest creates the Create request.
func (client *VolumeGroupsClient) createCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, parameters VolumeGroup, _ *VolumeGroupsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete an VolumeGroup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - elasticSanName - The name of the ElasticSan.
//   - volumeGroupName - The name of the VolumeGroup.
//   - options - VolumeGroupsClientBeginDeleteOptions contains the optional parameters for the VolumeGroupsClient.BeginDelete
//     method.
func (client *VolumeGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, options *VolumeGroupsClientBeginDeleteOptions) (*runtime.Poller[VolumeGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, elasticSanName, volumeGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VolumeGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VolumeGroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete an VolumeGroup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01-preview
func (client *VolumeGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, options *VolumeGroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VolumeGroupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, elasticSanName, volumeGroupName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VolumeGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, _ *VolumeGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an VolumeGroups.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - elasticSanName - The name of the ElasticSan.
//   - volumeGroupName - The name of the VolumeGroup.
//   - options - VolumeGroupsClientGetOptions contains the optional parameters for the VolumeGroupsClient.Get method.
func (client *VolumeGroupsClient) Get(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, options *VolumeGroupsClientGetOptions) (VolumeGroupsClientGetResponse, error) {
	var err error
	const operationName = "VolumeGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, elasticSanName, volumeGroupName, options)
	if err != nil {
		return VolumeGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VolumeGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VolumeGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VolumeGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, _ *VolumeGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VolumeGroupsClient) getHandleResponse(resp *http.Response) (VolumeGroupsClientGetResponse, error) {
	result := VolumeGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeGroup); err != nil {
		return VolumeGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByElasticSanPager - List VolumeGroups.
//
// Generated from API version 2024-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - elasticSanName - The name of the ElasticSan.
//   - options - VolumeGroupsClientListByElasticSanOptions contains the optional parameters for the VolumeGroupsClient.NewListByElasticSanPager
//     method.
func (client *VolumeGroupsClient) NewListByElasticSanPager(resourceGroupName string, elasticSanName string, options *VolumeGroupsClientListByElasticSanOptions) *runtime.Pager[VolumeGroupsClientListByElasticSanResponse] {
	return runtime.NewPager(runtime.PagingHandler[VolumeGroupsClientListByElasticSanResponse]{
		More: func(page VolumeGroupsClientListByElasticSanResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VolumeGroupsClientListByElasticSanResponse) (VolumeGroupsClientListByElasticSanResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VolumeGroupsClient.NewListByElasticSanPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByElasticSanCreateRequest(ctx, resourceGroupName, elasticSanName, options)
			}, nil)
			if err != nil {
				return VolumeGroupsClientListByElasticSanResponse{}, err
			}
			return client.listByElasticSanHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByElasticSanCreateRequest creates the ListByElasticSan request.
func (client *VolumeGroupsClient) listByElasticSanCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, options *VolumeGroupsClientListByElasticSanOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumeGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.XMSAccessSoftDeletedResources != nil {
		req.Raw().Header["x-ms-access-soft-deleted-resources"] = []string{string(*options.XMSAccessSoftDeletedResources)}
	}
	return req, nil
}

// listByElasticSanHandleResponse handles the ListByElasticSan response.
func (client *VolumeGroupsClient) listByElasticSanHandleResponse(resp *http.Response) (VolumeGroupsClientListByElasticSanResponse, error) {
	result := VolumeGroupsClientListByElasticSanResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeGroupList); err != nil {
		return VolumeGroupsClientListByElasticSanResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update an VolumeGroup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - elasticSanName - The name of the ElasticSan.
//   - volumeGroupName - The name of the VolumeGroup.
//   - parameters - Volume Group object.
//   - options - VolumeGroupsClientBeginUpdateOptions contains the optional parameters for the VolumeGroupsClient.BeginUpdate
//     method.
func (client *VolumeGroupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, parameters VolumeGroupUpdate, options *VolumeGroupsClientBeginUpdateOptions) (*runtime.Poller[VolumeGroupsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, elasticSanName, volumeGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VolumeGroupsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VolumeGroupsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update an VolumeGroup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01-preview
func (client *VolumeGroupsClient) update(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, parameters VolumeGroupUpdate, options *VolumeGroupsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VolumeGroupsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, elasticSanName, volumeGroupName, parameters, options)
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
func (client *VolumeGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, parameters VolumeGroupUpdate, _ *VolumeGroupsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
