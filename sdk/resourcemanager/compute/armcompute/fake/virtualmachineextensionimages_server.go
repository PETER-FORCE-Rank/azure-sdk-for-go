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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// VirtualMachineExtensionImagesServer is a fake server for instances of the armcompute.VirtualMachineExtensionImagesClient type.
type VirtualMachineExtensionImagesServer struct {
	// Get is the fake for method VirtualMachineExtensionImagesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, publisherName string, typeParam string, version string, options *armcompute.VirtualMachineExtensionImagesClientGetOptions) (resp azfake.Responder[armcompute.VirtualMachineExtensionImagesClientGetResponse], errResp azfake.ErrorResponder)

	// ListTypes is the fake for method VirtualMachineExtensionImagesClient.ListTypes
	// HTTP status codes to indicate success: http.StatusOK
	ListTypes func(ctx context.Context, location string, publisherName string, options *armcompute.VirtualMachineExtensionImagesClientListTypesOptions) (resp azfake.Responder[armcompute.VirtualMachineExtensionImagesClientListTypesResponse], errResp azfake.ErrorResponder)

	// ListVersions is the fake for method VirtualMachineExtensionImagesClient.ListVersions
	// HTTP status codes to indicate success: http.StatusOK
	ListVersions func(ctx context.Context, location string, publisherName string, typeParam string, options *armcompute.VirtualMachineExtensionImagesClientListVersionsOptions) (resp azfake.Responder[armcompute.VirtualMachineExtensionImagesClientListVersionsResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineExtensionImagesServerTransport creates a new instance of VirtualMachineExtensionImagesServerTransport with the provided implementation.
// The returned VirtualMachineExtensionImagesServerTransport instance is connected to an instance of armcompute.VirtualMachineExtensionImagesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachineExtensionImagesServerTransport(srv *VirtualMachineExtensionImagesServer) *VirtualMachineExtensionImagesServerTransport {
	return &VirtualMachineExtensionImagesServerTransport{srv: srv}
}

// VirtualMachineExtensionImagesServerTransport connects instances of armcompute.VirtualMachineExtensionImagesClient to instances of VirtualMachineExtensionImagesServer.
// Don't use this type directly, use NewVirtualMachineExtensionImagesServerTransport instead.
type VirtualMachineExtensionImagesServerTransport struct {
	srv *VirtualMachineExtensionImagesServer
}

// Do implements the policy.Transporter interface for VirtualMachineExtensionImagesServerTransport.
func (v *VirtualMachineExtensionImagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VirtualMachineExtensionImagesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if virtualMachineExtensionImagesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = virtualMachineExtensionImagesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "VirtualMachineExtensionImagesClient.Get":
				res.resp, res.err = v.dispatchGet(req)
			case "VirtualMachineExtensionImagesClient.ListTypes":
				res.resp, res.err = v.dispatchListTypes(req)
			case "VirtualMachineExtensionImagesClient.ListVersions":
				res.resp, res.err = v.dispatchListVersions(req)
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

func (v *VirtualMachineExtensionImagesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifacttypes/vmextension/types/(?P<type>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	typeParamParam, err := url.PathUnescape(matches[regex.SubexpIndex("type")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), locationParam, publisherNameParam, typeParamParam, versionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineExtensionImage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineExtensionImagesServerTransport) dispatchListTypes(req *http.Request) (*http.Response, error) {
	if v.srv.ListTypes == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListTypes not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifacttypes/vmextension/types`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.ListTypes(req.Context(), locationParam, publisherNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineExtensionImageArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineExtensionImagesServerTransport) dispatchListVersions(req *http.Request) (*http.Response, error) {
	if v.srv.ListVersions == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListVersions not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifacttypes/vmextension/types/(?P<type>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	typeParamParam, err := url.PathUnescape(matches[regex.SubexpIndex("type")])
	if err != nil {
		return nil, err
	}
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
	if err != nil {
		return nil, err
	}
	topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
	if err != nil {
		return nil, err
	}
	orderbyParam := getOptional(orderbyUnescaped)
	var options *armcompute.VirtualMachineExtensionImagesClientListVersionsOptions
	if filterParam != nil || topParam != nil || orderbyParam != nil {
		options = &armcompute.VirtualMachineExtensionImagesClientListVersionsOptions{
			Filter:  filterParam,
			Top:     topParam,
			Orderby: orderbyParam,
		}
	}
	respr, errRespr := v.srv.ListVersions(req.Context(), locationParam, publisherNameParam, typeParamParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineExtensionImageArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to VirtualMachineExtensionImagesServerTransport
var virtualMachineExtensionImagesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
