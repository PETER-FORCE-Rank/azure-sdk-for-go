// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armhardwaresecuritymodules_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/v2"
	"log"
)

// Generated from example definition: 2025-03-31/CloudHsmClusterPrivateLinkResource_ListByCloudHsmCluster_MaximumSet_Gen.json
func ExampleCloudHsmClusterPrivateLinkResourcesClient_NewListByCloudHsmClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudHsmClusterPrivateLinkResourcesClient().NewListByCloudHsmClusterPager("rgcloudhsm", "chsm1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armhardwaresecuritymodules.CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterResponse{
		// 	PrivateLinkResourceListResult: armhardwaresecuritymodules.PrivateLinkResourceListResult{
		// 		Value: []*armhardwaresecuritymodules.PrivateLinkResource{
		// 			{
		// 				Name: to.Ptr("sample-pls"),
		// 				Type: to.Ptr("Microsoft.HardwareSecurityModules/cloudHsmClusters/privateLinkResources"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgcloudhsm/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/chsm1/privateLinkResources/sample-pls"),
		// 				Properties: &armhardwaresecuritymodules.PrivateLinkResourceProperties{
		// 					GroupID: to.Ptr("cloudHsm"),
		// 					RequiredMembers: []*string{
		// 						to.Ptr("hsm1"),
		// 						to.Ptr("hsm2"),
		// 						to.Ptr("hsm3"),
		// 					},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.cloudhsm.azure-int.net"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
