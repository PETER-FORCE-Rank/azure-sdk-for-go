// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armportalservicescopilot

// CopilotSettingsClientCreateOrUpdateResponse contains the response from method CopilotSettingsClient.CreateOrUpdate.
type CopilotSettingsClientCreateOrUpdateResponse struct {
	// The copilot settings tenant resource definition.
	CopilotSettingsResource
}

// CopilotSettingsClientDeleteResponse contains the response from method CopilotSettingsClient.Delete.
type CopilotSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// CopilotSettingsClientGetResponse contains the response from method CopilotSettingsClient.Get.
type CopilotSettingsClientGetResponse struct {
	// The copilot settings tenant resource definition.
	CopilotSettingsResource
}

// CopilotSettingsClientUpdateResponse contains the response from method CopilotSettingsClient.Update.
type CopilotSettingsClientUpdateResponse struct {
	// The copilot settings tenant resource definition.
	CopilotSettingsResource
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}
