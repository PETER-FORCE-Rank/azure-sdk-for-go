//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// InputsClient contains the methods for the Inputs group.
// Don't use this type directly, use NewInputsClient() instead.
type InputsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewInputsClient creates a new instance of InputsClient with the specified values.
func NewInputsClient(con *arm.Connection, subscriptionID string) *InputsClient {
	return &InputsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrReplace - Creates an input or replaces an already existing input under an existing streaming job.
// If the operation fails it returns the *Error error type.
func (client *InputsClient) CreateOrReplace(ctx context.Context, resourceGroupName string, jobName string, inputName string, input Input, options *InputsCreateOrReplaceOptions) (InputsCreateOrReplaceResponse, error) {
	req, err := client.createOrReplaceCreateRequest(ctx, resourceGroupName, jobName, inputName, input, options)
	if err != nil {
		return InputsCreateOrReplaceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InputsCreateOrReplaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return InputsCreateOrReplaceResponse{}, client.createOrReplaceHandleError(resp)
	}
	return client.createOrReplaceHandleResponse(resp)
}

// createOrReplaceCreateRequest creates the CreateOrReplace request.
func (client *InputsClient) createOrReplaceCreateRequest(ctx context.Context, resourceGroupName string, jobName string, inputName string, input Input, options *InputsCreateOrReplaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if inputName == "" {
		return nil, errors.New("parameter inputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inputName}", url.PathEscape(inputName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, input)
}

// createOrReplaceHandleResponse handles the CreateOrReplace response.
func (client *InputsClient) createOrReplaceHandleResponse(resp *http.Response) (InputsCreateOrReplaceResponse, error) {
	result := InputsCreateOrReplaceResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Input); err != nil {
		return InputsCreateOrReplaceResponse{}, err
	}
	return result, nil
}

// createOrReplaceHandleError handles the CreateOrReplace error response.
func (client *InputsClient) createOrReplaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes an input from the streaming job.
// If the operation fails it returns the *Error error type.
func (client *InputsClient) Delete(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsDeleteOptions) (InputsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, jobName, inputName, options)
	if err != nil {
		return InputsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InputsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return InputsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return InputsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *InputsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if inputName == "" {
		return nil, errors.New("parameter inputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inputName}", url.PathEscape(inputName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *InputsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets details about the specified input.
// If the operation fails it returns the *Error error type.
func (client *InputsClient) Get(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsGetOptions) (InputsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, jobName, inputName, options)
	if err != nil {
		return InputsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InputsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InputsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InputsClient) getCreateRequest(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if inputName == "" {
		return nil, errors.New("parameter inputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inputName}", url.PathEscape(inputName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InputsClient) getHandleResponse(resp *http.Response) (InputsGetResponse, error) {
	result := InputsGetResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Input); err != nil {
		return InputsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *InputsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByStreamingJob - Lists all of the inputs under the specified streaming job.
// If the operation fails it returns the *Error error type.
func (client *InputsClient) ListByStreamingJob(resourceGroupName string, jobName string, options *InputsListByStreamingJobOptions) *InputsListByStreamingJobPager {
	return &InputsListByStreamingJobPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByStreamingJobCreateRequest(ctx, resourceGroupName, jobName, options)
		},
		advancer: func(ctx context.Context, resp InputsListByStreamingJobResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.InputListResult.NextLink)
		},
	}
}

// listByStreamingJobCreateRequest creates the ListByStreamingJob request.
func (client *InputsClient) listByStreamingJobCreateRequest(ctx context.Context, resourceGroupName string, jobName string, options *InputsListByStreamingJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByStreamingJobHandleResponse handles the ListByStreamingJob response.
func (client *InputsClient) listByStreamingJobHandleResponse(resp *http.Response) (InputsListByStreamingJobResponse, error) {
	result := InputsListByStreamingJobResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.InputListResult); err != nil {
		return InputsListByStreamingJobResponse{}, err
	}
	return result, nil
}

// listByStreamingJobHandleError handles the ListByStreamingJob error response.
func (client *InputsClient) listByStreamingJobHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginTest - Tests whether an input’s datasource is reachable and usable by the Azure Stream Analytics service.
// If the operation fails it returns the *Error error type.
func (client *InputsClient) BeginTest(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsBeginTestOptions) (InputsTestPollerResponse, error) {
	resp, err := client.test(ctx, resourceGroupName, jobName, inputName, options)
	if err != nil {
		return InputsTestPollerResponse{}, err
	}
	result := InputsTestPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("InputsClient.Test", "", resp, client.pl, client.testHandleError)
	if err != nil {
		return InputsTestPollerResponse{}, err
	}
	result.Poller = &InputsTestPoller{
		pt: pt,
	}
	return result, nil
}

// Test - Tests whether an input’s datasource is reachable and usable by the Azure Stream Analytics service.
// If the operation fails it returns the *Error error type.
func (client *InputsClient) test(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsBeginTestOptions) (*http.Response, error) {
	req, err := client.testCreateRequest(ctx, resourceGroupName, jobName, inputName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.testHandleError(resp)
	}
	return resp, nil
}

// testCreateRequest creates the Test request.
func (client *InputsClient) testCreateRequest(ctx context.Context, resourceGroupName string, jobName string, inputName string, options *InputsBeginTestOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}/test"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if inputName == "" {
		return nil, errors.New("parameter inputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inputName}", url.PathEscape(inputName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Input != nil {
		return req, runtime.MarshalAsJSON(req, *options.Input)
	}
	return req, nil
}

// testHandleError handles the Test error response.
func (client *InputsClient) testHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Updates an existing input under an existing streaming job. This can be used to partially update (ie. update one or two properties) an input
// without affecting the rest the job or input definition.
// If the operation fails it returns the *Error error type.
func (client *InputsClient) Update(ctx context.Context, resourceGroupName string, jobName string, inputName string, input Input, options *InputsUpdateOptions) (InputsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, jobName, inputName, input, options)
	if err != nil {
		return InputsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InputsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InputsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *InputsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, jobName string, inputName string, input Input, options *InputsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if inputName == "" {
		return nil, errors.New("parameter inputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inputName}", url.PathEscape(inputName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, input)
}

// updateHandleResponse handles the Update response.
func (client *InputsClient) updateHandleResponse(resp *http.Response) (InputsUpdateResponse, error) {
	result := InputsUpdateResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Input); err != nil {
		return InputsUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *InputsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
