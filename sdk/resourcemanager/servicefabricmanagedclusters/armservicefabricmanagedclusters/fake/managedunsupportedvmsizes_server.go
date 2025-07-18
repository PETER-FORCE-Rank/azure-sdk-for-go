// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
	"net/http"
	"net/url"
	"regexp"
)

// ManagedUnsupportedVMSizesServer is a fake server for instances of the armservicefabricmanagedclusters.ManagedUnsupportedVMSizesClient type.
type ManagedUnsupportedVMSizesServer struct {
	// Get is the fake for method ManagedUnsupportedVMSizesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, vmSize string, options *armservicefabricmanagedclusters.ManagedUnsupportedVMSizesClientGetOptions) (resp azfake.Responder[armservicefabricmanagedclusters.ManagedUnsupportedVMSizesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ManagedUnsupportedVMSizesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armservicefabricmanagedclusters.ManagedUnsupportedVMSizesClientListOptions) (resp azfake.PagerResponder[armservicefabricmanagedclusters.ManagedUnsupportedVMSizesClientListResponse])
}

// NewManagedUnsupportedVMSizesServerTransport creates a new instance of ManagedUnsupportedVMSizesServerTransport with the provided implementation.
// The returned ManagedUnsupportedVMSizesServerTransport instance is connected to an instance of armservicefabricmanagedclusters.ManagedUnsupportedVMSizesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedUnsupportedVMSizesServerTransport(srv *ManagedUnsupportedVMSizesServer) *ManagedUnsupportedVMSizesServerTransport {
	return &ManagedUnsupportedVMSizesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armservicefabricmanagedclusters.ManagedUnsupportedVMSizesClientListResponse]](),
	}
}

// ManagedUnsupportedVMSizesServerTransport connects instances of armservicefabricmanagedclusters.ManagedUnsupportedVMSizesClient to instances of ManagedUnsupportedVMSizesServer.
// Don't use this type directly, use NewManagedUnsupportedVMSizesServerTransport instead.
type ManagedUnsupportedVMSizesServerTransport struct {
	srv          *ManagedUnsupportedVMSizesServer
	newListPager *tracker[azfake.PagerResponder[armservicefabricmanagedclusters.ManagedUnsupportedVMSizesClientListResponse]]
}

// Do implements the policy.Transporter interface for ManagedUnsupportedVMSizesServerTransport.
func (m *ManagedUnsupportedVMSizesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *ManagedUnsupportedVMSizesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if managedUnsupportedVMSizesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = managedUnsupportedVMSizesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ManagedUnsupportedVMSizesClient.Get":
				res.resp, res.err = m.dispatchGet(req)
			case "ManagedUnsupportedVMSizesClient.NewListPager":
				res.resp, res.err = m.dispatchNewListPager(req)
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

func (m *ManagedUnsupportedVMSizesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabric/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedUnsupportedVMSizes/(?P<vmSize>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	vmSizeParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmSize")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), locationParam, vmSizeParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedVMSize, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedUnsupportedVMSizesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabric/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedUnsupportedVMSizes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListPager(locationParam, nil)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armservicefabricmanagedclusters.ManagedUnsupportedVMSizesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ManagedUnsupportedVMSizesServerTransport
var managedUnsupportedVMSizesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
