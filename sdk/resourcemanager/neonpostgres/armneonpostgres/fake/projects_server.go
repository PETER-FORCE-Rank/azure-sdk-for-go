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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/neonpostgres/armneonpostgres"
	"net/http"
	"net/url"
	"regexp"
)

// ProjectsServer is a fake server for instances of the armneonpostgres.ProjectsClient type.
type ProjectsServer struct {
	// BeginCreateOrUpdate is the fake for method ProjectsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, organizationName string, projectName string, resource armneonpostgres.Project, options *armneonpostgres.ProjectsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armneonpostgres.ProjectsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ProjectsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, organizationName string, projectName string, options *armneonpostgres.ProjectsClientDeleteOptions) (resp azfake.Responder[armneonpostgres.ProjectsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ProjectsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, organizationName string, projectName string, options *armneonpostgres.ProjectsClientGetOptions) (resp azfake.Responder[armneonpostgres.ProjectsClientGetResponse], errResp azfake.ErrorResponder)

	// GetConnectionURI is the fake for method ProjectsClient.GetConnectionURI
	// HTTP status codes to indicate success: http.StatusOK
	GetConnectionURI func(ctx context.Context, resourceGroupName string, organizationName string, projectName string, connectionURIParameters armneonpostgres.ConnectionURIProperties, options *armneonpostgres.ProjectsClientGetConnectionURIOptions) (resp azfake.Responder[armneonpostgres.ProjectsClientGetConnectionURIResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ProjectsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, organizationName string, options *armneonpostgres.ProjectsClientListOptions) (resp azfake.PagerResponder[armneonpostgres.ProjectsClientListResponse])

	// BeginUpdate is the fake for method ProjectsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, organizationName string, projectName string, properties armneonpostgres.Project, options *armneonpostgres.ProjectsClientBeginUpdateOptions) (resp azfake.PollerResponder[armneonpostgres.ProjectsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewProjectsServerTransport creates a new instance of ProjectsServerTransport with the provided implementation.
// The returned ProjectsServerTransport instance is connected to an instance of armneonpostgres.ProjectsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProjectsServerTransport(srv *ProjectsServer) *ProjectsServerTransport {
	return &ProjectsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armneonpostgres.ProjectsClientCreateOrUpdateResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armneonpostgres.ProjectsClientListResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armneonpostgres.ProjectsClientUpdateResponse]](),
	}
}

// ProjectsServerTransport connects instances of armneonpostgres.ProjectsClient to instances of ProjectsServer.
// Don't use this type directly, use NewProjectsServerTransport instead.
type ProjectsServerTransport struct {
	srv                 *ProjectsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armneonpostgres.ProjectsClientCreateOrUpdateResponse]]
	newListPager        *tracker[azfake.PagerResponder[armneonpostgres.ProjectsClientListResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armneonpostgres.ProjectsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ProjectsServerTransport.
func (p *ProjectsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *ProjectsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if projectsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = projectsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ProjectsClient.BeginCreateOrUpdate":
				res.resp, res.err = p.dispatchBeginCreateOrUpdate(req)
			case "ProjectsClient.Delete":
				res.resp, res.err = p.dispatchDelete(req)
			case "ProjectsClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "ProjectsClient.GetConnectionURI":
				res.resp, res.err = p.dispatchGetConnectionURI(req)
			case "ProjectsClient.NewListPager":
				res.resp, res.err = p.dispatchNewListPager(req)
			case "ProjectsClient.BeginUpdate":
				res.resp, res.err = p.dispatchBeginUpdate(req)
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

func (p *ProjectsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Neon\.Postgres/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/projects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armneonpostgres.Project](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		organizationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("organizationName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, organizationNameParam, projectNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *ProjectsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Neon\.Postgres/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/projects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	organizationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("organizationName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), resourceGroupNameParam, organizationNameParam, projectNameParam, nil)
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

func (p *ProjectsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Neon\.Postgres/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/projects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	organizationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("organizationName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, organizationNameParam, projectNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Project, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProjectsServerTransport) dispatchGetConnectionURI(req *http.Request) (*http.Response, error) {
	if p.srv.GetConnectionURI == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetConnectionURI not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Neon\.Postgres/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/projects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getConnectionUri`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armneonpostgres.ConnectionURIProperties](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	organizationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("organizationName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetConnectionURI(req.Context(), resourceGroupNameParam, organizationNameParam, projectNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectionURIProperties, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProjectsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Neon\.Postgres/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/projects`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		organizationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("organizationName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, organizationNameParam, nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armneonpostgres.ProjectsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}

func (p *ProjectsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := p.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Neon\.Postgres/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/projects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armneonpostgres.Project](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		organizationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("organizationName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginUpdate(req.Context(), resourceGroupNameParam, organizationNameParam, projectNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		p.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		p.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ProjectsServerTransport
var projectsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
