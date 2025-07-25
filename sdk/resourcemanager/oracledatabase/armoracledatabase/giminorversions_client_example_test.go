// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armoracledatabase_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
	"log"
)

// Generated from example definition: 2025-03-01/GiMinorVersions_Get_MaximumSet_Gen.json
func ExampleGiMinorVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGiMinorVersionsClient().Get(ctx, "eastus", "giVersionName", "giMinorVersionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armoracledatabase.GiMinorVersionsClientGetResponse{
	// 	GiMinorVersion: &armoracledatabase.GiMinorVersion{
	// 		Properties: &armoracledatabase.GiMinorVersionProperties{
	// 			Version: to.Ptr("oznlupakllqglkgrepsttjqwcocrdmacagcrdloaikfhhwepottsdoezcoiyjmemxgbxstdjmjcmgjdoozvefmsungmipnnuhixcipskafmyczxzffyqoinjdbnlkthcsbqobeddkuyxqcrveozao"),
	// 			GridImageOcid: to.Ptr("ocid1.dbpatch.oc1..aaaaa3klq"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/giVersions/19.0.0.0/giMinorVersions/minorVersion"),
	// 		Name: to.Ptr("bthzelohzaolpeytksswrnkbcnen"),
	// 		Type: to.Ptr("acesitmovkyb"),
	// 		SystemData: &armoracledatabase.SystemData{
	// 			CreatedBy: to.Ptr("ilrpjodjmvzhybazxipoplnql"),
	// 			CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:12.592Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("lhjbxchqkaia"),
	// 			LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:12.592Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-03-01/GiMinorVersions_ListByParent_MaximumSet_Gen.json
func ExampleGiMinorVersionsClient_NewListByParentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGiMinorVersionsClient().NewListByParentPager("eastus", "giVersionName", &armoracledatabase.GiMinorVersionsClientListByParentOptions{
		ShapeFamily: to.Ptr(armoracledatabase.ShapeFamily("rtfcosvtlpeeqoicsjqggtgc"))})
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
		// page = armoracledatabase.GiMinorVersionsClientListByParentResponse{
		// 	GiMinorVersionListResult: armoracledatabase.GiMinorVersionListResult{
		// 		Value: []*armoracledatabase.GiMinorVersion{
		// 			{
		// 				Properties: &armoracledatabase.GiMinorVersionProperties{
		// 					Version: to.Ptr("oznlupakllqglkgrepsttjqwcocrdmacagcrdloaikfhhwepottsdoezcoiyjmemxgbxstdjmjcmgjdoozvefmsungmipnnuhixcipskafmyczxzffyqoinjdbnlkthcsbqobeddkuyxqcrveozao"),
		// 					GridImageOcid: to.Ptr("ocid1.dbpatch.oc1..aaaaa3klq"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/giVersions/19.0.0.0/giMinorVersions/minorVersion"),
		// 				Name: to.Ptr("bthzelohzaolpeytksswrnkbcnen"),
		// 				Type: to.Ptr("acesitmovkyb"),
		// 				SystemData: &armoracledatabase.SystemData{
		// 					CreatedBy: to.Ptr("ilrpjodjmvzhybazxipoplnql"),
		// 					CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:12.592Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("lhjbxchqkaia"),
		// 					LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:12.592Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
