// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v7"
	"net/http"
	"net/url"
	"regexp"
)

// Server is a fake server for instances of the armcontainerservice.Client type.
type Server struct {
	// NewListNodeImageVersionsPager is the fake for method Client.NewListNodeImageVersionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListNodeImageVersionsPager func(location string, options *armcontainerservice.ClientListNodeImageVersionsOptions) (resp azfake.PagerResponder[armcontainerservice.ClientListNodeImageVersionsResponse])
}

// NewServerTransport creates a new instance of ServerTransport with the provided implementation.
// The returned ServerTransport instance is connected to an instance of armcontainerservice.Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerTransport(srv *Server) *ServerTransport {
	return &ServerTransport{
		srv:                           srv,
		newListNodeImageVersionsPager: newTracker[azfake.PagerResponder[armcontainerservice.ClientListNodeImageVersionsResponse]](),
	}
}

// ServerTransport connects instances of armcontainerservice.Client to instances of Server.
// Don't use this type directly, use NewServerTransport instead.
type ServerTransport struct {
	srv                           *Server
	newListNodeImageVersionsPager *tracker[azfake.PagerResponder[armcontainerservice.ClientListNodeImageVersionsResponse]]
}

// Do implements the policy.Transporter interface for ServerTransport.
func (s *ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serverTransportInterceptor != nil {
			res.resp, res.err, intercepted = serverTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "Client.NewListNodeImageVersionsPager":
				res.resp, res.err = s.dispatchNewListNodeImageVersionsPager(req)
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

func (s *ServerTransport) dispatchNewListNodeImageVersionsPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListNodeImageVersionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListNodeImageVersionsPager not implemented")}
	}
	newListNodeImageVersionsPager := s.newListNodeImageVersionsPager.get(req)
	if newListNodeImageVersionsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/nodeImageVersions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListNodeImageVersionsPager(locationParam, nil)
		newListNodeImageVersionsPager = &resp
		s.newListNodeImageVersionsPager.add(req, newListNodeImageVersionsPager)
		server.PagerResponderInjectNextLinks(newListNodeImageVersionsPager, req, func(page *armcontainerservice.ClientListNodeImageVersionsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListNodeImageVersionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListNodeImageVersionsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListNodeImageVersionsPager) {
		s.newListNodeImageVersionsPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ServerTransport
var serverTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
