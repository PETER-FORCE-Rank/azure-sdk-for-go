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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool/v2"
	"net/http"
	"net/url"
	"regexp"
)

// StandbyContainerGroupPoolRuntimeViewsServer is a fake server for instances of the armstandbypool.StandbyContainerGroupPoolRuntimeViewsClient type.
type StandbyContainerGroupPoolRuntimeViewsServer struct {
	// Get is the fake for method StandbyContainerGroupPoolRuntimeViewsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, standbyContainerGroupPoolName string, runtimeView string, options *armstandbypool.StandbyContainerGroupPoolRuntimeViewsClientGetOptions) (resp azfake.Responder[armstandbypool.StandbyContainerGroupPoolRuntimeViewsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByStandbyPoolPager is the fake for method StandbyContainerGroupPoolRuntimeViewsClient.NewListByStandbyPoolPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByStandbyPoolPager func(resourceGroupName string, standbyContainerGroupPoolName string, options *armstandbypool.StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolOptions) (resp azfake.PagerResponder[armstandbypool.StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse])
}

// NewStandbyContainerGroupPoolRuntimeViewsServerTransport creates a new instance of StandbyContainerGroupPoolRuntimeViewsServerTransport with the provided implementation.
// The returned StandbyContainerGroupPoolRuntimeViewsServerTransport instance is connected to an instance of armstandbypool.StandbyContainerGroupPoolRuntimeViewsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewStandbyContainerGroupPoolRuntimeViewsServerTransport(srv *StandbyContainerGroupPoolRuntimeViewsServer) *StandbyContainerGroupPoolRuntimeViewsServerTransport {
	return &StandbyContainerGroupPoolRuntimeViewsServerTransport{
		srv:                       srv,
		newListByStandbyPoolPager: newTracker[azfake.PagerResponder[armstandbypool.StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse]](),
	}
}

// StandbyContainerGroupPoolRuntimeViewsServerTransport connects instances of armstandbypool.StandbyContainerGroupPoolRuntimeViewsClient to instances of StandbyContainerGroupPoolRuntimeViewsServer.
// Don't use this type directly, use NewStandbyContainerGroupPoolRuntimeViewsServerTransport instead.
type StandbyContainerGroupPoolRuntimeViewsServerTransport struct {
	srv                       *StandbyContainerGroupPoolRuntimeViewsServer
	newListByStandbyPoolPager *tracker[azfake.PagerResponder[armstandbypool.StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse]]
}

// Do implements the policy.Transporter interface for StandbyContainerGroupPoolRuntimeViewsServerTransport.
func (s *StandbyContainerGroupPoolRuntimeViewsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *StandbyContainerGroupPoolRuntimeViewsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if standbyContainerGroupPoolRuntimeViewsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = standbyContainerGroupPoolRuntimeViewsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "StandbyContainerGroupPoolRuntimeViewsClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "StandbyContainerGroupPoolRuntimeViewsClient.NewListByStandbyPoolPager":
				res.resp, res.err = s.dispatchNewListByStandbyPoolPager(req)
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

func (s *StandbyContainerGroupPoolRuntimeViewsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StandbyPool/standbyContainerGroupPools/(?P<standbyContainerGroupPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runtimeViews/(?P<runtimeView>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	standbyContainerGroupPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("standbyContainerGroupPoolName")])
	if err != nil {
		return nil, err
	}
	runtimeViewParam, err := url.PathUnescape(matches[regex.SubexpIndex("runtimeView")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, standbyContainerGroupPoolNameParam, runtimeViewParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StandbyContainerGroupPoolRuntimeViewResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StandbyContainerGroupPoolRuntimeViewsServerTransport) dispatchNewListByStandbyPoolPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByStandbyPoolPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByStandbyPoolPager not implemented")}
	}
	newListByStandbyPoolPager := s.newListByStandbyPoolPager.get(req)
	if newListByStandbyPoolPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StandbyPool/standbyContainerGroupPools/(?P<standbyContainerGroupPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runtimeViews`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		standbyContainerGroupPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("standbyContainerGroupPoolName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByStandbyPoolPager(resourceGroupNameParam, standbyContainerGroupPoolNameParam, nil)
		newListByStandbyPoolPager = &resp
		s.newListByStandbyPoolPager.add(req, newListByStandbyPoolPager)
		server.PagerResponderInjectNextLinks(newListByStandbyPoolPager, req, func(page *armstandbypool.StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByStandbyPoolPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByStandbyPoolPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByStandbyPoolPager) {
		s.newListByStandbyPoolPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to StandbyContainerGroupPoolRuntimeViewsServerTransport
var standbyContainerGroupPoolRuntimeViewsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
