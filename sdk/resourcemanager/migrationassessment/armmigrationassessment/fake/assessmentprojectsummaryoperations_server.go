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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
	"net/http"
	"net/url"
	"regexp"
)

// AssessmentProjectSummaryOperationsServer is a fake server for instances of the armmigrationassessment.AssessmentProjectSummaryOperationsClient type.
type AssessmentProjectSummaryOperationsServer struct {
	// Get is the fake for method AssessmentProjectSummaryOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, projectName string, projectSummaryName string, options *armmigrationassessment.AssessmentProjectSummaryOperationsClientGetOptions) (resp azfake.Responder[armmigrationassessment.AssessmentProjectSummaryOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAssessmentProjectPager is the fake for method AssessmentProjectSummaryOperationsClient.NewListByAssessmentProjectPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAssessmentProjectPager func(resourceGroupName string, projectName string, options *armmigrationassessment.AssessmentProjectSummaryOperationsClientListByAssessmentProjectOptions) (resp azfake.PagerResponder[armmigrationassessment.AssessmentProjectSummaryOperationsClientListByAssessmentProjectResponse])
}

// NewAssessmentProjectSummaryOperationsServerTransport creates a new instance of AssessmentProjectSummaryOperationsServerTransport with the provided implementation.
// The returned AssessmentProjectSummaryOperationsServerTransport instance is connected to an instance of armmigrationassessment.AssessmentProjectSummaryOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAssessmentProjectSummaryOperationsServerTransport(srv *AssessmentProjectSummaryOperationsServer) *AssessmentProjectSummaryOperationsServerTransport {
	return &AssessmentProjectSummaryOperationsServerTransport{
		srv:                             srv,
		newListByAssessmentProjectPager: newTracker[azfake.PagerResponder[armmigrationassessment.AssessmentProjectSummaryOperationsClientListByAssessmentProjectResponse]](),
	}
}

// AssessmentProjectSummaryOperationsServerTransport connects instances of armmigrationassessment.AssessmentProjectSummaryOperationsClient to instances of AssessmentProjectSummaryOperationsServer.
// Don't use this type directly, use NewAssessmentProjectSummaryOperationsServerTransport instead.
type AssessmentProjectSummaryOperationsServerTransport struct {
	srv                             *AssessmentProjectSummaryOperationsServer
	newListByAssessmentProjectPager *tracker[azfake.PagerResponder[armmigrationassessment.AssessmentProjectSummaryOperationsClientListByAssessmentProjectResponse]]
}

// Do implements the policy.Transporter interface for AssessmentProjectSummaryOperationsServerTransport.
func (a *AssessmentProjectSummaryOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AssessmentProjectSummaryOperationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if assessmentProjectSummaryOperationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = assessmentProjectSummaryOperationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "AssessmentProjectSummaryOperationsClient.Get":
				res.resp, res.err = a.dispatchGet(req)
			case "AssessmentProjectSummaryOperationsClient.NewListByAssessmentProjectPager":
				res.resp, res.err = a.dispatchNewListByAssessmentProjectPager(req)
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

func (a *AssessmentProjectSummaryOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/projectSummary/(?P<projectSummaryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	projectSummaryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectSummaryName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, projectNameParam, projectSummaryNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssessmentProjectSummary, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AssessmentProjectSummaryOperationsServerTransport) dispatchNewListByAssessmentProjectPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByAssessmentProjectPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAssessmentProjectPager not implemented")}
	}
	newListByAssessmentProjectPager := a.newListByAssessmentProjectPager.get(req)
	if newListByAssessmentProjectPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/projectSummary`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByAssessmentProjectPager(resourceGroupNameParam, projectNameParam, nil)
		newListByAssessmentProjectPager = &resp
		a.newListByAssessmentProjectPager.add(req, newListByAssessmentProjectPager)
		server.PagerResponderInjectNextLinks(newListByAssessmentProjectPager, req, func(page *armmigrationassessment.AssessmentProjectSummaryOperationsClientListByAssessmentProjectResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAssessmentProjectPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByAssessmentProjectPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAssessmentProjectPager) {
		a.newListByAssessmentProjectPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to AssessmentProjectSummaryOperationsServerTransport
var assessmentProjectSummaryOperationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
