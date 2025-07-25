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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices/v2"
	"net/http"
	"net/url"
	"regexp"
)

// VaultCertificatesServer is a fake server for instances of the armrecoveryservices.VaultCertificatesClient type.
type VaultCertificatesServer struct {
	// Create is the fake for method VaultCertificatesClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, vaultName string, certificateName string, certificateRequest armrecoveryservices.CertificateRequest, options *armrecoveryservices.VaultCertificatesClientCreateOptions) (resp azfake.Responder[armrecoveryservices.VaultCertificatesClientCreateResponse], errResp azfake.ErrorResponder)
}

// NewVaultCertificatesServerTransport creates a new instance of VaultCertificatesServerTransport with the provided implementation.
// The returned VaultCertificatesServerTransport instance is connected to an instance of armrecoveryservices.VaultCertificatesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVaultCertificatesServerTransport(srv *VaultCertificatesServer) *VaultCertificatesServerTransport {
	return &VaultCertificatesServerTransport{srv: srv}
}

// VaultCertificatesServerTransport connects instances of armrecoveryservices.VaultCertificatesClient to instances of VaultCertificatesServer.
// Don't use this type directly, use NewVaultCertificatesServerTransport instead.
type VaultCertificatesServerTransport struct {
	srv *VaultCertificatesServer
}

// Do implements the policy.Transporter interface for VaultCertificatesServerTransport.
func (v *VaultCertificatesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VaultCertificatesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if vaultCertificatesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = vaultCertificatesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "VaultCertificatesClient.Create":
				res.resp, res.err = v.dispatchCreate(req)
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

func (v *VaultCertificatesServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if v.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/Subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/certificates/(?P<certificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armrecoveryservices.CertificateRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	certificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("certificateName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Create(req.Context(), resourceGroupNameParam, vaultNameParam, certificateNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VaultCertificateResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to VaultCertificatesServerTransport
var vaultCertificatesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
