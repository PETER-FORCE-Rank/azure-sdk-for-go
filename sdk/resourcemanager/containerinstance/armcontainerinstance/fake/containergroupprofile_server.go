//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ContainerGroupProfileServer is a fake server for instances of the armcontainerinstance.ContainerGroupProfileClient type.
type ContainerGroupProfileServer struct {
	// GetByRevisionNumber is the fake for method ContainerGroupProfileClient.GetByRevisionNumber
	// HTTP status codes to indicate success: http.StatusOK
	GetByRevisionNumber func(ctx context.Context, resourceGroupName string, containerGroupProfileName string, revisionNumber string, options *armcontainerinstance.ContainerGroupProfileClientGetByRevisionNumberOptions) (resp azfake.Responder[armcontainerinstance.ContainerGroupProfileClientGetByRevisionNumberResponse], errResp azfake.ErrorResponder)

	// NewListAllRevisionsPager is the fake for method ContainerGroupProfileClient.NewListAllRevisionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllRevisionsPager func(resourceGroupName string, containerGroupProfileName string, options *armcontainerinstance.ContainerGroupProfileClientListAllRevisionsOptions) (resp azfake.PagerResponder[armcontainerinstance.ContainerGroupProfileClientListAllRevisionsResponse])
}

// NewContainerGroupProfileServerTransport creates a new instance of ContainerGroupProfileServerTransport with the provided implementation.
// The returned ContainerGroupProfileServerTransport instance is connected to an instance of armcontainerinstance.ContainerGroupProfileClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewContainerGroupProfileServerTransport(srv *ContainerGroupProfileServer) *ContainerGroupProfileServerTransport {
	return &ContainerGroupProfileServerTransport{
		srv:                      srv,
		newListAllRevisionsPager: newTracker[azfake.PagerResponder[armcontainerinstance.ContainerGroupProfileClientListAllRevisionsResponse]](),
	}
}

// ContainerGroupProfileServerTransport connects instances of armcontainerinstance.ContainerGroupProfileClient to instances of ContainerGroupProfileServer.
// Don't use this type directly, use NewContainerGroupProfileServerTransport instead.
type ContainerGroupProfileServerTransport struct {
	srv                      *ContainerGroupProfileServer
	newListAllRevisionsPager *tracker[azfake.PagerResponder[armcontainerinstance.ContainerGroupProfileClientListAllRevisionsResponse]]
}

// Do implements the policy.Transporter interface for ContainerGroupProfileServerTransport.
func (c *ContainerGroupProfileServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ContainerGroupProfileClient.GetByRevisionNumber":
		resp, err = c.dispatchGetByRevisionNumber(req)
	case "ContainerGroupProfileClient.NewListAllRevisionsPager":
		resp, err = c.dispatchNewListAllRevisionsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ContainerGroupProfileServerTransport) dispatchGetByRevisionNumber(req *http.Request) (*http.Response, error) {
	if c.srv.GetByRevisionNumber == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByRevisionNumber not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroupProfiles/(?P<containerGroupProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revisions/(?P<revisionNumber>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerGroupProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupProfileName")])
	if err != nil {
		return nil, err
	}
	revisionNumberParam, err := url.PathUnescape(matches[regex.SubexpIndex("revisionNumber")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetByRevisionNumber(req.Context(), resourceGroupNameParam, containerGroupProfileNameParam, revisionNumberParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerGroupProfile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerGroupProfileServerTransport) dispatchNewListAllRevisionsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListAllRevisionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllRevisionsPager not implemented")}
	}
	newListAllRevisionsPager := c.newListAllRevisionsPager.get(req)
	if newListAllRevisionsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroupProfiles/(?P<containerGroupProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revisions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerGroupProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupProfileName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListAllRevisionsPager(resourceGroupNameParam, containerGroupProfileNameParam, nil)
		newListAllRevisionsPager = &resp
		c.newListAllRevisionsPager.add(req, newListAllRevisionsPager)
		server.PagerResponderInjectNextLinks(newListAllRevisionsPager, req, func(page *armcontainerinstance.ContainerGroupProfileClientListAllRevisionsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAllRevisionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListAllRevisionsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAllRevisionsPager) {
		c.newListAllRevisionsPager.remove(req)
	}
	return resp, nil
}
