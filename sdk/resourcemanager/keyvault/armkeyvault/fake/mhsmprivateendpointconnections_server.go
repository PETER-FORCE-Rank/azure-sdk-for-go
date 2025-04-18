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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// MHSMPrivateEndpointConnectionsServer is a fake server for instances of the armkeyvault.MHSMPrivateEndpointConnectionsClient type.
type MHSMPrivateEndpointConnectionsServer struct {
	// BeginDelete is the fake for method MHSMPrivateEndpointConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, options *armkeyvault.MHSMPrivateEndpointConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armkeyvault.MHSMPrivateEndpointConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method MHSMPrivateEndpointConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, options *armkeyvault.MHSMPrivateEndpointConnectionsClientGetOptions) (resp azfake.Responder[armkeyvault.MHSMPrivateEndpointConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourcePager is the fake for method MHSMPrivateEndpointConnectionsClient.NewListByResourcePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourcePager func(resourceGroupName string, name string, options *armkeyvault.MHSMPrivateEndpointConnectionsClientListByResourceOptions) (resp azfake.PagerResponder[armkeyvault.MHSMPrivateEndpointConnectionsClientListByResourceResponse])

	// Put is the fake for method MHSMPrivateEndpointConnectionsClient.Put
	// HTTP status codes to indicate success: http.StatusOK
	Put func(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, properties armkeyvault.MHSMPrivateEndpointConnection, options *armkeyvault.MHSMPrivateEndpointConnectionsClientPutOptions) (resp azfake.Responder[armkeyvault.MHSMPrivateEndpointConnectionsClientPutResponse], errResp azfake.ErrorResponder)
}

// NewMHSMPrivateEndpointConnectionsServerTransport creates a new instance of MHSMPrivateEndpointConnectionsServerTransport with the provided implementation.
// The returned MHSMPrivateEndpointConnectionsServerTransport instance is connected to an instance of armkeyvault.MHSMPrivateEndpointConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMHSMPrivateEndpointConnectionsServerTransport(srv *MHSMPrivateEndpointConnectionsServer) *MHSMPrivateEndpointConnectionsServerTransport {
	return &MHSMPrivateEndpointConnectionsServerTransport{
		srv:                    srv,
		beginDelete:            newTracker[azfake.PollerResponder[armkeyvault.MHSMPrivateEndpointConnectionsClientDeleteResponse]](),
		newListByResourcePager: newTracker[azfake.PagerResponder[armkeyvault.MHSMPrivateEndpointConnectionsClientListByResourceResponse]](),
	}
}

// MHSMPrivateEndpointConnectionsServerTransport connects instances of armkeyvault.MHSMPrivateEndpointConnectionsClient to instances of MHSMPrivateEndpointConnectionsServer.
// Don't use this type directly, use NewMHSMPrivateEndpointConnectionsServerTransport instead.
type MHSMPrivateEndpointConnectionsServerTransport struct {
	srv                    *MHSMPrivateEndpointConnectionsServer
	beginDelete            *tracker[azfake.PollerResponder[armkeyvault.MHSMPrivateEndpointConnectionsClientDeleteResponse]]
	newListByResourcePager *tracker[azfake.PagerResponder[armkeyvault.MHSMPrivateEndpointConnectionsClientListByResourceResponse]]
}

// Do implements the policy.Transporter interface for MHSMPrivateEndpointConnectionsServerTransport.
func (m *MHSMPrivateEndpointConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *MHSMPrivateEndpointConnectionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if mhsmPrivateEndpointConnectionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = mhsmPrivateEndpointConnectionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "MHSMPrivateEndpointConnectionsClient.BeginDelete":
				res.resp, res.err = m.dispatchBeginDelete(req)
			case "MHSMPrivateEndpointConnectionsClient.Get":
				res.resp, res.err = m.dispatchGet(req)
			case "MHSMPrivateEndpointConnectionsClient.NewListByResourcePager":
				res.resp, res.err = m.dispatchNewListByResourcePager(req)
			case "MHSMPrivateEndpointConnectionsClient.Put":
				res.resp, res.err = m.dispatchPut(req)
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

func (m *MHSMPrivateEndpointConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if m.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := m.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginDelete(req.Context(), resourceGroupNameParam, nameParam, privateEndpointConnectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		m.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		m.beginDelete.remove(req)
	}

	return resp, nil
}

func (m *MHSMPrivateEndpointConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, nameParam, privateEndpointConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MHSMPrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MHSMPrivateEndpointConnectionsServerTransport) dispatchNewListByResourcePager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByResourcePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourcePager not implemented")}
	}
	newListByResourcePager := m.newListByResourcePager.get(req)
	if newListByResourcePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListByResourcePager(resourceGroupNameParam, nameParam, nil)
		newListByResourcePager = &resp
		m.newListByResourcePager.add(req, newListByResourcePager)
		server.PagerResponderInjectNextLinks(newListByResourcePager, req, func(page *armkeyvault.MHSMPrivateEndpointConnectionsClientListByResourceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourcePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByResourcePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourcePager) {
		m.newListByResourcePager.remove(req)
	}
	return resp, nil
}

func (m *MHSMPrivateEndpointConnectionsServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if m.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armkeyvault.MHSMPrivateEndpointConnection](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Put(req.Context(), resourceGroupNameParam, nameParam, privateEndpointConnectionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MHSMPrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).AzureAsyncOperation; val != nil {
		resp.Header.Set("Azure-AsyncOperation", *val)
	}
	if val := server.GetResponse(respr).RetryAfter; val != nil {
		resp.Header.Set("Retry-After", strconv.FormatInt(int64(*val), 10))
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to MHSMPrivateEndpointConnectionsServerTransport
var mhsmPrivateEndpointConnectionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
