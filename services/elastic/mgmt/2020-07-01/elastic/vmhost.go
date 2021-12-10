package elastic

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VMHostClient is the client for the VMHost methods of the Elastic service.
type VMHostClient struct {
	BaseClient
}

// NewVMHostClient creates an instance of the VMHostClient client.
func NewVMHostClient(subscriptionID string) VMHostClient {
	return NewVMHostClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVMHostClientWithBaseURI creates an instance of the VMHostClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVMHostClientWithBaseURI(baseURI string, subscriptionID string) VMHostClient {
	return VMHostClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List sends the list request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Elastic resource belongs.
// monitorName - monitor resource name
func (client VMHostClient) List(ctx context.Context, resourceGroupName string, monitorName string) (result VMHostListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VMHostClient.List")
		defer func() {
			sc := -1
			if result.vhlr.Response.Response != nil {
				sc = result.vhlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, monitorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.VMHostClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.vhlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "elastic.VMHostClient", "List", resp, "Failure sending request")
		return
	}

	result.vhlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.VMHostClient", "List", resp, "Failure responding to request")
		return
	}
	if result.vhlr.hasNextLink() && result.vhlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VMHostClient) ListPreparer(ctx context.Context, resourceGroupName string, monitorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/listVMHost", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VMHostClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VMHostClient) ListResponder(resp *http.Response) (result VMHostListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VMHostClient) listNextResults(ctx context.Context, lastResults VMHostListResponse) (result VMHostListResponse, err error) {
	req, err := lastResults.vMHostListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "elastic.VMHostClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "elastic.VMHostClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.VMHostClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VMHostClient) ListComplete(ctx context.Context, resourceGroupName string, monitorName string) (result VMHostListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VMHostClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, monitorName)
	return
}
