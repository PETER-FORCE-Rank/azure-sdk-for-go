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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v10"
	"net/http"
	"net/url"
	"regexp"
)

// CredentialOperationsServer is a fake server for instances of the armdatafactory.CredentialOperationsClient type.
type CredentialOperationsServer struct {
	// CreateOrUpdate is the fake for method CredentialOperationsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, factoryName string, credentialName string, credential armdatafactory.CredentialResource, options *armdatafactory.CredentialOperationsClientCreateOrUpdateOptions) (resp azfake.Responder[armdatafactory.CredentialOperationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method CredentialOperationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, factoryName string, credentialName string, options *armdatafactory.CredentialOperationsClientDeleteOptions) (resp azfake.Responder[armdatafactory.CredentialOperationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CredentialOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotModified
	Get func(ctx context.Context, resourceGroupName string, factoryName string, credentialName string, options *armdatafactory.CredentialOperationsClientGetOptions) (resp azfake.Responder[armdatafactory.CredentialOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByFactoryPager is the fake for method CredentialOperationsClient.NewListByFactoryPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFactoryPager func(resourceGroupName string, factoryName string, options *armdatafactory.CredentialOperationsClientListByFactoryOptions) (resp azfake.PagerResponder[armdatafactory.CredentialOperationsClientListByFactoryResponse])
}

// NewCredentialOperationsServerTransport creates a new instance of CredentialOperationsServerTransport with the provided implementation.
// The returned CredentialOperationsServerTransport instance is connected to an instance of armdatafactory.CredentialOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCredentialOperationsServerTransport(srv *CredentialOperationsServer) *CredentialOperationsServerTransport {
	return &CredentialOperationsServerTransport{
		srv:                   srv,
		newListByFactoryPager: newTracker[azfake.PagerResponder[armdatafactory.CredentialOperationsClientListByFactoryResponse]](),
	}
}

// CredentialOperationsServerTransport connects instances of armdatafactory.CredentialOperationsClient to instances of CredentialOperationsServer.
// Don't use this type directly, use NewCredentialOperationsServerTransport instead.
type CredentialOperationsServerTransport struct {
	srv                   *CredentialOperationsServer
	newListByFactoryPager *tracker[azfake.PagerResponder[armdatafactory.CredentialOperationsClientListByFactoryResponse]]
}

// Do implements the policy.Transporter interface for CredentialOperationsServerTransport.
func (c *CredentialOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CredentialOperationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if credentialOperationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = credentialOperationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "CredentialOperationsClient.CreateOrUpdate":
				res.resp, res.err = c.dispatchCreateOrUpdate(req)
			case "CredentialOperationsClient.Delete":
				res.resp, res.err = c.dispatchDelete(req)
			case "CredentialOperationsClient.Get":
				res.resp, res.err = c.dispatchGet(req)
			case "CredentialOperationsClient.NewListByFactoryPager":
				res.resp, res.err = c.dispatchNewListByFactoryPager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (c *CredentialOperationsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/credentials/(?P<credentialName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.CredentialResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	credentialNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("credentialName")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armdatafactory.CredentialOperationsClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armdatafactory.CredentialOperationsClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := c.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, factoryNameParam, credentialNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CredentialResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CredentialOperationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if c.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/credentials/(?P<credentialName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	credentialNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("credentialName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Delete(req.Context(), resourceGroupNameParam, factoryNameParam, credentialNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CredentialOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/credentials/(?P<credentialName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	credentialNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("credentialName")])
	if err != nil {
		return nil, err
	}
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *armdatafactory.CredentialOperationsClientGetOptions
	if ifNoneMatchParam != nil {
		options = &armdatafactory.CredentialOperationsClientGetOptions{
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, factoryNameParam, credentialNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotModified}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotModified", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CredentialResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CredentialOperationsServerTransport) dispatchNewListByFactoryPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByFactoryPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFactoryPager not implemented")}
	}
	newListByFactoryPager := c.newListByFactoryPager.get(req)
	if newListByFactoryPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/credentials`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByFactoryPager(resourceGroupNameParam, factoryNameParam, nil)
		newListByFactoryPager = &resp
		c.newListByFactoryPager.add(req, newListByFactoryPager)
		server.PagerResponderInjectNextLinks(newListByFactoryPager, req, func(page *armdatafactory.CredentialOperationsClientListByFactoryResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFactoryPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByFactoryPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFactoryPager) {
		c.newListByFactoryPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to CredentialOperationsServerTransport
var credentialOperationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
