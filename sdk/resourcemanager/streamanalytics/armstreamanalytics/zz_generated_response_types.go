//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

import (
	"context"
	"net/http"
	"time"

	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
)

// ClustersCreateOrUpdatePollerResponse contains the response from method Clusters.CreateOrUpdate.
type ClustersCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClustersCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l ClustersCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClustersCreateOrUpdateResponse, error) {
	respType := ClustersCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Cluster)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ClustersCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *ClustersCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *ClustersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ClustersClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &ClustersCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ClustersCreateOrUpdateResponse contains the response from method Clusters.CreateOrUpdate.
type ClustersCreateOrUpdateResponse struct {
	ClustersCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersCreateOrUpdateResult contains the result from method Clusters.CreateOrUpdate.
type ClustersCreateOrUpdateResult struct {
	Cluster
}

// ClustersDeletePollerResponse contains the response from method Clusters.Delete.
type ClustersDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClustersDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l ClustersDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClustersDeleteResponse, error) {
	respType := ClustersDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ClustersDeletePollerResponse from the provided client and resume token.
func (l *ClustersDeletePollerResponse) Resume(ctx context.Context, client *ClustersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ClustersClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &ClustersDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ClustersDeleteResponse contains the response from method Clusters.Delete.
type ClustersDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersGetResponse contains the response from method Clusters.Get.
type ClustersGetResponse struct {
	ClustersGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersGetResult contains the result from method Clusters.Get.
type ClustersGetResult struct {
	Cluster
}

// ClustersListByResourceGroupResponse contains the response from method Clusters.ListByResourceGroup.
type ClustersListByResourceGroupResponse struct {
	ClustersListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersListByResourceGroupResult contains the result from method Clusters.ListByResourceGroup.
type ClustersListByResourceGroupResult struct {
	ClusterListResult
}

// ClustersListBySubscriptionResponse contains the response from method Clusters.ListBySubscription.
type ClustersListBySubscriptionResponse struct {
	ClustersListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersListBySubscriptionResult contains the result from method Clusters.ListBySubscription.
type ClustersListBySubscriptionResult struct {
	ClusterListResult
}

// ClustersListStreamingJobsResponse contains the response from method Clusters.ListStreamingJobs.
type ClustersListStreamingJobsResponse struct {
	ClustersListStreamingJobsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersListStreamingJobsResult contains the result from method Clusters.ListStreamingJobs.
type ClustersListStreamingJobsResult struct {
	ClusterJobListResult
}

// ClustersUpdatePollerResponse contains the response from method Clusters.Update.
type ClustersUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClustersUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l ClustersUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClustersUpdateResponse, error) {
	respType := ClustersUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Cluster)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ClustersUpdatePollerResponse from the provided client and resume token.
func (l *ClustersUpdatePollerResponse) Resume(ctx context.Context, client *ClustersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ClustersClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &ClustersUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ClustersUpdateResponse contains the response from method Clusters.Update.
type ClustersUpdateResponse struct {
	ClustersUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersUpdateResult contains the result from method Clusters.Update.
type ClustersUpdateResult struct {
	Cluster
}

// FunctionsCreateOrReplaceResponse contains the response from method Functions.CreateOrReplace.
type FunctionsCreateOrReplaceResponse struct {
	FunctionsCreateOrReplaceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FunctionsCreateOrReplaceResult contains the result from method Functions.CreateOrReplace.
type FunctionsCreateOrReplaceResult struct {
	Function
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// FunctionsDeleteResponse contains the response from method Functions.Delete.
type FunctionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FunctionsGetResponse contains the response from method Functions.Get.
type FunctionsGetResponse struct {
	FunctionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FunctionsGetResult contains the result from method Functions.Get.
type FunctionsGetResult struct {
	Function
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// FunctionsListByStreamingJobResponse contains the response from method Functions.ListByStreamingJob.
type FunctionsListByStreamingJobResponse struct {
	FunctionsListByStreamingJobResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FunctionsListByStreamingJobResult contains the result from method Functions.ListByStreamingJob.
type FunctionsListByStreamingJobResult struct {
	FunctionListResult
}

// FunctionsRetrieveDefaultDefinitionResponse contains the response from method Functions.RetrieveDefaultDefinition.
type FunctionsRetrieveDefaultDefinitionResponse struct {
	FunctionsRetrieveDefaultDefinitionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FunctionsRetrieveDefaultDefinitionResult contains the result from method Functions.RetrieveDefaultDefinition.
type FunctionsRetrieveDefaultDefinitionResult struct {
	Function
}

// FunctionsTestPollerResponse contains the response from method Functions.Test.
type FunctionsTestPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *FunctionsTestPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l FunctionsTestPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (FunctionsTestResponse, error) {
	respType := FunctionsTestResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ResourceTestStatus)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a FunctionsTestPollerResponse from the provided client and resume token.
func (l *FunctionsTestPollerResponse) Resume(ctx context.Context, client *FunctionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("FunctionsClient.Test", token, client.pl, client.testHandleError)
	if err != nil {
		return err
	}
	poller := &FunctionsTestPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// FunctionsTestResponse contains the response from method Functions.Test.
type FunctionsTestResponse struct {
	FunctionsTestResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FunctionsTestResult contains the result from method Functions.Test.
type FunctionsTestResult struct {
	ResourceTestStatus
}

// FunctionsUpdateResponse contains the response from method Functions.Update.
type FunctionsUpdateResponse struct {
	FunctionsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FunctionsUpdateResult contains the result from method Functions.Update.
type FunctionsUpdateResult struct {
	Function
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// InputsCreateOrReplaceResponse contains the response from method Inputs.CreateOrReplace.
type InputsCreateOrReplaceResponse struct {
	InputsCreateOrReplaceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// InputsCreateOrReplaceResult contains the result from method Inputs.CreateOrReplace.
type InputsCreateOrReplaceResult struct {
	Input
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// InputsDeleteResponse contains the response from method Inputs.Delete.
type InputsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// InputsGetResponse contains the response from method Inputs.Get.
type InputsGetResponse struct {
	InputsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// InputsGetResult contains the result from method Inputs.Get.
type InputsGetResult struct {
	Input
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// InputsListByStreamingJobResponse contains the response from method Inputs.ListByStreamingJob.
type InputsListByStreamingJobResponse struct {
	InputsListByStreamingJobResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// InputsListByStreamingJobResult contains the result from method Inputs.ListByStreamingJob.
type InputsListByStreamingJobResult struct {
	InputListResult
}

// InputsTestPollerResponse contains the response from method Inputs.Test.
type InputsTestPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *InputsTestPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l InputsTestPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (InputsTestResponse, error) {
	respType := InputsTestResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ResourceTestStatus)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a InputsTestPollerResponse from the provided client and resume token.
func (l *InputsTestPollerResponse) Resume(ctx context.Context, client *InputsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("InputsClient.Test", token, client.pl, client.testHandleError)
	if err != nil {
		return err
	}
	poller := &InputsTestPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// InputsTestResponse contains the response from method Inputs.Test.
type InputsTestResponse struct {
	InputsTestResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// InputsTestResult contains the result from method Inputs.Test.
type InputsTestResult struct {
	ResourceTestStatus
}

// InputsUpdateResponse contains the response from method Inputs.Update.
type InputsUpdateResponse struct {
	InputsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// InputsUpdateResult contains the result from method Inputs.Update.
type InputsUpdateResult struct {
	Input
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// OutputsCreateOrReplaceResponse contains the response from method Outputs.CreateOrReplace.
type OutputsCreateOrReplaceResponse struct {
	OutputsCreateOrReplaceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OutputsCreateOrReplaceResult contains the result from method Outputs.CreateOrReplace.
type OutputsCreateOrReplaceResult struct {
	Output
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// OutputsDeleteResponse contains the response from method Outputs.Delete.
type OutputsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OutputsGetResponse contains the response from method Outputs.Get.
type OutputsGetResponse struct {
	OutputsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OutputsGetResult contains the result from method Outputs.Get.
type OutputsGetResult struct {
	Output
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// OutputsListByStreamingJobResponse contains the response from method Outputs.ListByStreamingJob.
type OutputsListByStreamingJobResponse struct {
	OutputsListByStreamingJobResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OutputsListByStreamingJobResult contains the result from method Outputs.ListByStreamingJob.
type OutputsListByStreamingJobResult struct {
	OutputListResult
}

// OutputsTestPollerResponse contains the response from method Outputs.Test.
type OutputsTestPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *OutputsTestPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l OutputsTestPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (OutputsTestResponse, error) {
	respType := OutputsTestResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ResourceTestStatus)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a OutputsTestPollerResponse from the provided client and resume token.
func (l *OutputsTestPollerResponse) Resume(ctx context.Context, client *OutputsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("OutputsClient.Test", token, client.pl, client.testHandleError)
	if err != nil {
		return err
	}
	poller := &OutputsTestPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// OutputsTestResponse contains the response from method Outputs.Test.
type OutputsTestResponse struct {
	OutputsTestResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OutputsTestResult contains the result from method Outputs.Test.
type OutputsTestResult struct {
	ResourceTestStatus
}

// OutputsUpdateResponse contains the response from method Outputs.Update.
type OutputsUpdateResponse struct {
	OutputsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OutputsUpdateResult contains the result from method Outputs.Update.
type OutputsUpdateResult struct {
	Output
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// PrivateEndpointsCreateOrUpdateResponse contains the response from method PrivateEndpoints.CreateOrUpdate.
type PrivateEndpointsCreateOrUpdateResponse struct {
	PrivateEndpointsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointsCreateOrUpdateResult contains the result from method PrivateEndpoints.CreateOrUpdate.
type PrivateEndpointsCreateOrUpdateResult struct {
	PrivateEndpoint
}

// PrivateEndpointsDeletePollerResponse contains the response from method PrivateEndpoints.Delete.
type PrivateEndpointsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l PrivateEndpointsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointsDeleteResponse, error) {
	respType := PrivateEndpointsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointsDeletePollerResponse from the provided client and resume token.
func (l *PrivateEndpointsDeletePollerResponse) Resume(ctx context.Context, client *PrivateEndpointsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointsDeleteResponse contains the response from method PrivateEndpoints.Delete.
type PrivateEndpointsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointsGetResponse contains the response from method PrivateEndpoints.Get.
type PrivateEndpointsGetResponse struct {
	PrivateEndpointsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointsGetResult contains the result from method PrivateEndpoints.Get.
type PrivateEndpointsGetResult struct {
	PrivateEndpoint
}

// PrivateEndpointsListByClusterResponse contains the response from method PrivateEndpoints.ListByCluster.
type PrivateEndpointsListByClusterResponse struct {
	PrivateEndpointsListByClusterResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointsListByClusterResult contains the result from method PrivateEndpoints.ListByCluster.
type PrivateEndpointsListByClusterResult struct {
	PrivateEndpointListResult
}

// StreamingJobsCreateOrReplacePollerResponse contains the response from method StreamingJobs.CreateOrReplace.
type StreamingJobsCreateOrReplacePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *StreamingJobsCreateOrReplacePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l StreamingJobsCreateOrReplacePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (StreamingJobsCreateOrReplaceResponse, error) {
	respType := StreamingJobsCreateOrReplaceResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.StreamingJob)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a StreamingJobsCreateOrReplacePollerResponse from the provided client and resume token.
func (l *StreamingJobsCreateOrReplacePollerResponse) Resume(ctx context.Context, client *StreamingJobsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("StreamingJobsClient.CreateOrReplace", token, client.pl, client.createOrReplaceHandleError)
	if err != nil {
		return err
	}
	poller := &StreamingJobsCreateOrReplacePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// StreamingJobsCreateOrReplaceResponse contains the response from method StreamingJobs.CreateOrReplace.
type StreamingJobsCreateOrReplaceResponse struct {
	StreamingJobsCreateOrReplaceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StreamingJobsCreateOrReplaceResult contains the result from method StreamingJobs.CreateOrReplace.
type StreamingJobsCreateOrReplaceResult struct {
	StreamingJob
}

// StreamingJobsDeletePollerResponse contains the response from method StreamingJobs.Delete.
type StreamingJobsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *StreamingJobsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l StreamingJobsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (StreamingJobsDeleteResponse, error) {
	respType := StreamingJobsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a StreamingJobsDeletePollerResponse from the provided client and resume token.
func (l *StreamingJobsDeletePollerResponse) Resume(ctx context.Context, client *StreamingJobsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("StreamingJobsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &StreamingJobsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// StreamingJobsDeleteResponse contains the response from method StreamingJobs.Delete.
type StreamingJobsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StreamingJobsGetResponse contains the response from method StreamingJobs.Get.
type StreamingJobsGetResponse struct {
	StreamingJobsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StreamingJobsGetResult contains the result from method StreamingJobs.Get.
type StreamingJobsGetResult struct {
	StreamingJob
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// StreamingJobsListByResourceGroupResponse contains the response from method StreamingJobs.ListByResourceGroup.
type StreamingJobsListByResourceGroupResponse struct {
	StreamingJobsListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StreamingJobsListByResourceGroupResult contains the result from method StreamingJobs.ListByResourceGroup.
type StreamingJobsListByResourceGroupResult struct {
	StreamingJobListResult
}

// StreamingJobsListResponse contains the response from method StreamingJobs.List.
type StreamingJobsListResponse struct {
	StreamingJobsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StreamingJobsListResult contains the result from method StreamingJobs.List.
type StreamingJobsListResult struct {
	StreamingJobListResult
}

// StreamingJobsStartPollerResponse contains the response from method StreamingJobs.Start.
type StreamingJobsStartPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *StreamingJobsStartPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l StreamingJobsStartPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (StreamingJobsStartResponse, error) {
	respType := StreamingJobsStartResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a StreamingJobsStartPollerResponse from the provided client and resume token.
func (l *StreamingJobsStartPollerResponse) Resume(ctx context.Context, client *StreamingJobsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("StreamingJobsClient.Start", token, client.pl, client.startHandleError)
	if err != nil {
		return err
	}
	poller := &StreamingJobsStartPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// StreamingJobsStartResponse contains the response from method StreamingJobs.Start.
type StreamingJobsStartResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StreamingJobsStopPollerResponse contains the response from method StreamingJobs.Stop.
type StreamingJobsStopPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *StreamingJobsStopPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l StreamingJobsStopPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (StreamingJobsStopResponse, error) {
	respType := StreamingJobsStopResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a StreamingJobsStopPollerResponse from the provided client and resume token.
func (l *StreamingJobsStopPollerResponse) Resume(ctx context.Context, client *StreamingJobsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("StreamingJobsClient.Stop", token, client.pl, client.stopHandleError)
	if err != nil {
		return err
	}
	poller := &StreamingJobsStopPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// StreamingJobsStopResponse contains the response from method StreamingJobs.Stop.
type StreamingJobsStopResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StreamingJobsUpdateResponse contains the response from method StreamingJobs.Update.
type StreamingJobsUpdateResponse struct {
	StreamingJobsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StreamingJobsUpdateResult contains the result from method StreamingJobs.Update.
type StreamingJobsUpdateResult struct {
	StreamingJob
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// SubscriptionsCompileQueryResponse contains the response from method Subscriptions.CompileQuery.
type SubscriptionsCompileQueryResponse struct {
	SubscriptionsCompileQueryResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsCompileQueryResult contains the result from method Subscriptions.CompileQuery.
type SubscriptionsCompileQueryResult struct {
	QueryCompilationResult
}

// SubscriptionsListQuotasResponse contains the response from method Subscriptions.ListQuotas.
type SubscriptionsListQuotasResponse struct {
	SubscriptionsListQuotasResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsListQuotasResult contains the result from method Subscriptions.ListQuotas.
type SubscriptionsListQuotasResult struct {
	SubscriptionQuotasListResult
}

// SubscriptionsSampleInputPollerResponse contains the response from method Subscriptions.SampleInput.
type SubscriptionsSampleInputPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SubscriptionsSampleInputPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l SubscriptionsSampleInputPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SubscriptionsSampleInputResponse, error) {
	respType := SubscriptionsSampleInputResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.SampleInputResult)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SubscriptionsSampleInputPollerResponse from the provided client and resume token.
func (l *SubscriptionsSampleInputPollerResponse) Resume(ctx context.Context, client *SubscriptionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SubscriptionsClient.SampleInput", token, client.pl, client.sampleInputHandleError)
	if err != nil {
		return err
	}
	poller := &SubscriptionsSampleInputPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SubscriptionsSampleInputResponse contains the response from method Subscriptions.SampleInput.
type SubscriptionsSampleInputResponse struct {
	SubscriptionsSampleInputResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsSampleInputResult contains the result from method Subscriptions.SampleInput.
type SubscriptionsSampleInputResult struct {
	SampleInputResult
}

// SubscriptionsTestInputPollerResponse contains the response from method Subscriptions.TestInput.
type SubscriptionsTestInputPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SubscriptionsTestInputPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l SubscriptionsTestInputPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SubscriptionsTestInputResponse, error) {
	respType := SubscriptionsTestInputResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.TestDatasourceResult)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SubscriptionsTestInputPollerResponse from the provided client and resume token.
func (l *SubscriptionsTestInputPollerResponse) Resume(ctx context.Context, client *SubscriptionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SubscriptionsClient.TestInput", token, client.pl, client.testInputHandleError)
	if err != nil {
		return err
	}
	poller := &SubscriptionsTestInputPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SubscriptionsTestInputResponse contains the response from method Subscriptions.TestInput.
type SubscriptionsTestInputResponse struct {
	SubscriptionsTestInputResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsTestInputResult contains the result from method Subscriptions.TestInput.
type SubscriptionsTestInputResult struct {
	TestDatasourceResult
}

// SubscriptionsTestOutputPollerResponse contains the response from method Subscriptions.TestOutput.
type SubscriptionsTestOutputPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SubscriptionsTestOutputPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l SubscriptionsTestOutputPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SubscriptionsTestOutputResponse, error) {
	respType := SubscriptionsTestOutputResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.TestDatasourceResult)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SubscriptionsTestOutputPollerResponse from the provided client and resume token.
func (l *SubscriptionsTestOutputPollerResponse) Resume(ctx context.Context, client *SubscriptionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SubscriptionsClient.TestOutput", token, client.pl, client.testOutputHandleError)
	if err != nil {
		return err
	}
	poller := &SubscriptionsTestOutputPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SubscriptionsTestOutputResponse contains the response from method Subscriptions.TestOutput.
type SubscriptionsTestOutputResponse struct {
	SubscriptionsTestOutputResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsTestOutputResult contains the result from method Subscriptions.TestOutput.
type SubscriptionsTestOutputResult struct {
	TestDatasourceResult
}

// SubscriptionsTestQueryPollerResponse contains the response from method Subscriptions.TestQuery.
type SubscriptionsTestQueryPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SubscriptionsTestQueryPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l SubscriptionsTestQueryPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SubscriptionsTestQueryResponse, error) {
	respType := SubscriptionsTestQueryResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.QueryTestingResult)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SubscriptionsTestQueryPollerResponse from the provided client and resume token.
func (l *SubscriptionsTestQueryPollerResponse) Resume(ctx context.Context, client *SubscriptionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SubscriptionsClient.TestQuery", token, client.pl, client.testQueryHandleError)
	if err != nil {
		return err
	}
	poller := &SubscriptionsTestQueryPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SubscriptionsTestQueryResponse contains the response from method Subscriptions.TestQuery.
type SubscriptionsTestQueryResponse struct {
	SubscriptionsTestQueryResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsTestQueryResult contains the result from method Subscriptions.TestQuery.
type SubscriptionsTestQueryResult struct {
	QueryTestingResult
}

// TransformationsCreateOrReplaceResponse contains the response from method Transformations.CreateOrReplace.
type TransformationsCreateOrReplaceResponse struct {
	TransformationsCreateOrReplaceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TransformationsCreateOrReplaceResult contains the result from method Transformations.CreateOrReplace.
type TransformationsCreateOrReplaceResult struct {
	Transformation
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// TransformationsGetResponse contains the response from method Transformations.Get.
type TransformationsGetResponse struct {
	TransformationsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TransformationsGetResult contains the result from method Transformations.Get.
type TransformationsGetResult struct {
	Transformation
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// TransformationsUpdateResponse contains the response from method Transformations.Update.
type TransformationsUpdateResponse struct {
	TransformationsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TransformationsUpdateResult contains the result from method Transformations.Update.
type TransformationsUpdateResult struct {
	Transformation
	// ETag contains the information returned from the ETag header response.
	ETag *string
}
