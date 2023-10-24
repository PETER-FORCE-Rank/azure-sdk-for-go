//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevcenter

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

// ImageVersionsClient contains the methods for the ImageVersions group.
// Don't use this type directly, use NewImageVersionsClient() instead.
type ImageVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewImageVersionsClient creates a new instance of ImageVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewImageVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ImageVersionsClient, error) {
	cl, err := arm.NewClient(moduleName+".ImageVersionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ImageVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets an image version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - galleryName - The name of the gallery.
//   - imageName - The name of the image.
//   - versionName - The version of the image.
//   - options - ImageVersionsClientGetOptions contains the optional parameters for the ImageVersionsClient.Get method.
func (client *ImageVersionsClient) Get(ctx context.Context, resourceGroupName string, devCenterName string, galleryName string, imageName string, versionName string, options *ImageVersionsClientGetOptions) (ImageVersionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, devCenterName, galleryName, imageName, versionName, options)
	if err != nil {
		return ImageVersionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImageVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImageVersionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ImageVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, galleryName string, imageName string, versionName string, options *ImageVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/galleries/{galleryName}/images/{imageName}/versions/{versionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if imageName == "" {
		return nil, errors.New("parameter imageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ImageVersionsClient) getHandleResponse(resp *http.Response) (ImageVersionsClientGetResponse, error) {
	result := ImageVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageVersion); err != nil {
		return ImageVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByImagePager - Lists versions for an image.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - galleryName - The name of the gallery.
//   - imageName - The name of the image.
//   - options - ImageVersionsClientListByImageOptions contains the optional parameters for the ImageVersionsClient.NewListByImagePager
//     method.
func (client *ImageVersionsClient) NewListByImagePager(resourceGroupName string, devCenterName string, galleryName string, imageName string, options *ImageVersionsClientListByImageOptions) *runtime.Pager[ImageVersionsClientListByImageResponse] {
	return runtime.NewPager(runtime.PagingHandler[ImageVersionsClientListByImageResponse]{
		More: func(page ImageVersionsClientListByImageResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ImageVersionsClientListByImageResponse) (ImageVersionsClientListByImageResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByImageCreateRequest(ctx, resourceGroupName, devCenterName, galleryName, imageName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ImageVersionsClientListByImageResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ImageVersionsClientListByImageResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ImageVersionsClientListByImageResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByImageHandleResponse(resp)
		},
	})
}

// listByImageCreateRequest creates the ListByImage request.
func (client *ImageVersionsClient) listByImageCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, galleryName string, imageName string, options *ImageVersionsClientListByImageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/galleries/{galleryName}/images/{imageName}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if imageName == "" {
		return nil, errors.New("parameter imageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByImageHandleResponse handles the ListByImage response.
func (client *ImageVersionsClient) listByImageHandleResponse(resp *http.Response) (ImageVersionsClientListByImageResponse, error) {
	result := ImageVersionsClientListByImageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageVersionListResult); err != nil {
		return ImageVersionsClientListByImageResponse{}, err
	}
	return result, nil
}
