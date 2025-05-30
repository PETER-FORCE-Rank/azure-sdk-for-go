//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v10"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ApproveRejectPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "connection", armdatafactory.PrivateLinkConnectionApprovalRequestResource{
		Properties: &armdatafactory.PrivateLinkConnectionApprovalRequest{
			PrivateEndpoint: &armdatafactory.PrivateEndpoint{
				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/privateEndpoints/myPrivateEndpoint"),
			},
			PrivateLinkServiceConnectionState: &armdatafactory.PrivateLinkConnectionState{
				Description:     to.Ptr("Approved by admin."),
				ActionsRequired: to.Ptr(""),
				Status:          to.Ptr("Approved"),
			},
		},
	}, &armdatafactory.PrivateEndpointConnectionClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnectionResource = armdatafactory.PrivateEndpointConnectionResource{
	// 	Name: to.Ptr("exampleFactoryName"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
	// 	Properties: &armdatafactory.RemotePrivateEndpointConnection{
	// 		PrivateEndpoint: &armdatafactory.ArmIDWrapper{
	// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/privateEndpoints/myPrivateEndpoint"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armdatafactory.PrivateLinkConnectionState{
	// 			Description: to.Ptr("Approved by admin."),
	// 			ActionsRequired: to.Ptr(""),
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/GetPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "connection", &armdatafactory.PrivateEndpointConnectionClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnectionResource = armdatafactory.PrivateEndpointConnectionResource{
	// 	Name: to.Ptr("exampleFactoryName"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
	// 	Properties: &armdatafactory.RemotePrivateEndpointConnection{
	// 		PrivateEndpoint: &armdatafactory.ArmIDWrapper{
	// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/privateEndpoints/myPrivateEndpoint"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armdatafactory.PrivateLinkConnectionState{
	// 			Description: to.Ptr("Approved by admin."),
	// 			ActionsRequired: to.Ptr(""),
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DeletePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPrivateEndpointConnectionClient().Delete(ctx, "exampleResourceGroup", "exampleFactoryName", "connection", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
