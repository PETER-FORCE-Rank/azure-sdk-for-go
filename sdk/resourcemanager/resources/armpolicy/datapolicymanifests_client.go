// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicy

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

// DataPolicyManifestsClient contains the methods for the DataPolicyManifests group.
// Don't use this type directly, use NewDataPolicyManifestsClient() instead.
type DataPolicyManifestsClient struct {
	internal *arm.Client
}

// NewDataPolicyManifestsClient creates a new instance of DataPolicyManifestsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataPolicyManifestsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*DataPolicyManifestsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DataPolicyManifestsClient{
		internal: cl,
	}
	return client, nil
}

// GetByPolicyMode - This operation retrieves the data policy manifest with the given policy mode.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
//   - policyMode - The policy mode of the data policy manifest to get.
//   - options - DataPolicyManifestsClientGetByPolicyModeOptions contains the optional parameters for the DataPolicyManifestsClient.GetByPolicyMode
//     method.
func (client *DataPolicyManifestsClient) GetByPolicyMode(ctx context.Context, policyMode string, options *DataPolicyManifestsClientGetByPolicyModeOptions) (DataPolicyManifestsClientGetByPolicyModeResponse, error) {
	var err error
	const operationName = "DataPolicyManifestsClient.GetByPolicyMode"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByPolicyModeCreateRequest(ctx, policyMode, options)
	if err != nil {
		return DataPolicyManifestsClientGetByPolicyModeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataPolicyManifestsClientGetByPolicyModeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataPolicyManifestsClientGetByPolicyModeResponse{}, err
	}
	resp, err := client.getByPolicyModeHandleResponse(httpResp)
	return resp, err
}

// getByPolicyModeCreateRequest creates the GetByPolicyMode request.
func (client *DataPolicyManifestsClient) getByPolicyModeCreateRequest(ctx context.Context, policyMode string, _ *DataPolicyManifestsClientGetByPolicyModeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/dataPolicyManifests/{policyMode}"
	if policyMode == "" {
		return nil, errors.New("parameter policyMode cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyMode}", url.PathEscape(policyMode))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByPolicyModeHandleResponse handles the GetByPolicyMode response.
func (client *DataPolicyManifestsClient) getByPolicyModeHandleResponse(resp *http.Response) (DataPolicyManifestsClientGetByPolicyModeResponse, error) {
	result := DataPolicyManifestsClientGetByPolicyModeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataPolicyManifest); err != nil {
		return DataPolicyManifestsClientGetByPolicyModeResponse{}, err
	}
	return result, nil
}

// NewListPager - This operation retrieves a list of all the data policy manifests that match the optional given $filter.
// Valid values for $filter are: "$filter=namespace eq '{0}'". If $filter is not provided, the
// unfiltered list includes all data policy manifests for data resource types. If $filter=namespace is provided, the returned
// list only includes all data policy manifests that have a namespace matching
// the provided value.
//
// Generated from API version 2020-09-01
//   - options - DataPolicyManifestsClientListOptions contains the optional parameters for the DataPolicyManifestsClient.NewListPager
//     method.
func (client *DataPolicyManifestsClient) NewListPager(options *DataPolicyManifestsClientListOptions) *runtime.Pager[DataPolicyManifestsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataPolicyManifestsClientListResponse]{
		More: func(page DataPolicyManifestsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataPolicyManifestsClientListResponse) (DataPolicyManifestsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DataPolicyManifestsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DataPolicyManifestsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DataPolicyManifestsClient) listCreateRequest(ctx context.Context, options *DataPolicyManifestsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/dataPolicyManifests"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DataPolicyManifestsClient) listHandleResponse(resp *http.Response) (DataPolicyManifestsClientListResponse, error) {
	result := DataPolicyManifestsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataPolicyManifestListResult); err != nil {
		return DataPolicyManifestsClientListResponse{}, err
	}
	return result, nil
}
