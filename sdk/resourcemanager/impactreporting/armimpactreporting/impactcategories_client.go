// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armimpactreporting

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

// ImpactCategoriesClient contains the methods for the ImpactCategories group.
// Don't use this type directly, use NewImpactCategoriesClient() instead.
type ImpactCategoriesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewImpactCategoriesClient creates a new instance of ImpactCategoriesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewImpactCategoriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ImpactCategoriesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ImpactCategoriesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a ImpactCategory
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - impactCategoryName - Name of the impact category
//   - options - ImpactCategoriesClientGetOptions contains the optional parameters for the ImpactCategoriesClient.Get method.
func (client *ImpactCategoriesClient) Get(ctx context.Context, impactCategoryName string, options *ImpactCategoriesClientGetOptions) (ImpactCategoriesClientGetResponse, error) {
	var err error
	const operationName = "ImpactCategoriesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, impactCategoryName, options)
	if err != nil {
		return ImpactCategoriesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImpactCategoriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImpactCategoriesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ImpactCategoriesClient) getCreateRequest(ctx context.Context, impactCategoryName string, _ *ImpactCategoriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Impact/impactCategories/{impactCategoryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if impactCategoryName == "" {
		return nil, errors.New("parameter impactCategoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{impactCategoryName}", url.PathEscape(impactCategoryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ImpactCategoriesClient) getHandleResponse(resp *http.Response) (ImpactCategoriesClientGetResponse, error) {
	result := ImpactCategoriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImpactCategory); err != nil {
		return ImpactCategoriesClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List ImpactCategory resources by subscription
//
// Generated from API version 2024-05-01-preview
//   - resourceType - Filter by resource type
//   - options - ImpactCategoriesClientListBySubscriptionOptions contains the optional parameters for the ImpactCategoriesClient.NewListBySubscriptionPager
//     method.
func (client *ImpactCategoriesClient) NewListBySubscriptionPager(resourceType string, options *ImpactCategoriesClientListBySubscriptionOptions) *runtime.Pager[ImpactCategoriesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ImpactCategoriesClientListBySubscriptionResponse]{
		More: func(page ImpactCategoriesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ImpactCategoriesClientListBySubscriptionResponse) (ImpactCategoriesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ImpactCategoriesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, resourceType, options)
			}, nil)
			if err != nil {
				return ImpactCategoriesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ImpactCategoriesClient) listBySubscriptionCreateRequest(ctx context.Context, resourceType string, options *ImpactCategoriesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Impact/impactCategories"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	if options != nil && options.CategoryName != nil {
		reqQP.Set("categoryName", *options.CategoryName)
	}
	reqQP.Set("resourceType", resourceType)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ImpactCategoriesClient) listBySubscriptionHandleResponse(resp *http.Response) (ImpactCategoriesClientListBySubscriptionResponse, error) {
	result := ImpactCategoriesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImpactCategoryListResult); err != nil {
		return ImpactCategoriesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
