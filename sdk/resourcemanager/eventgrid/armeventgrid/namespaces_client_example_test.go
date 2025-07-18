//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Namespaces_Get.json
func ExampleNamespacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().Get(ctx, "examplerg", "exampleNamespaceName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Namespace = armeventgrid.Namespace{
	// 	Name: to.Ptr("exampleNamespaceName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/namespaces"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 		"key3": to.Ptr("value3"),
	// 	},
	// 	Properties: &armeventgrid.NamespaceProperties{
	// 		ProvisioningState: to.Ptr(armeventgrid.NamespaceProvisioningStateSucceeded),
	// 		TopicSpacesConfiguration: &armeventgrid.TopicSpacesConfiguration{
	// 			Hostname: to.Ptr("exampleNamespaceName1.westus-1.mqtt.eventgrid-int.azure.net"),
	// 			RouteTopicResourceID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampleTopic1"),
	// 			State: to.Ptr(armeventgrid.TopicSpacesConfigurationStateEnabled),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Namespaces_CreateOrUpdate.json
func ExampleNamespacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespacesClient().BeginCreateOrUpdate(ctx, "examplerg", "exampleNamespaceName1", armeventgrid.Namespace{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"tag1": to.Ptr("value11"),
			"tag2": to.Ptr("value22"),
		},
		Properties: &armeventgrid.NamespaceProperties{
			TopicSpacesConfiguration: &armeventgrid.TopicSpacesConfiguration{
				RouteTopicResourceID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampleTopic1"),
				State:                to.Ptr(armeventgrid.TopicSpacesConfigurationStateEnabled),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Namespace = armeventgrid.Namespace{
	// 	Name: to.Ptr("exampleNamespaceName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/Namespaces"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value11"),
	// 		"key2": to.Ptr("value22"),
	// 	},
	// 	Properties: &armeventgrid.NamespaceProperties{
	// 		ProvisioningState: to.Ptr(armeventgrid.NamespaceProvisioningStateSucceeded),
	// 		TopicSpacesConfiguration: &armeventgrid.TopicSpacesConfiguration{
	// 			RouteTopicResourceID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampleTopic1"),
	// 			State: to.Ptr(armeventgrid.TopicSpacesConfigurationStateEnabled),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Namespaces_Delete.json
func ExampleNamespacesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespacesClient().BeginDelete(ctx, "examplerg", "exampleNamespaceName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Namespaces_Update.json
func ExampleNamespacesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespacesClient().BeginUpdate(ctx, "examplerg", "exampleNamespaceName1", armeventgrid.NamespaceUpdateParameters{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1Updated"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Namespace = armeventgrid.Namespace{
	// 	Name: to.Ptr("exampleNamespaceName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/namespaces"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armeventgrid.NamespaceProperties{
	// 		ProvisioningState: to.Ptr(armeventgrid.NamespaceProvisioningStateSucceeded),
	// 		TopicSpacesConfiguration: &armeventgrid.TopicSpacesConfiguration{
	// 			Hostname: to.Ptr("exampleNamespaceName1.westus-1.mqtt.eventgrid-int.azure.net"),
	// 			RouteTopicResourceID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampleTopic1"),
	// 			State: to.Ptr(armeventgrid.TopicSpacesConfigurationStateEnabled),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Namespaces_ListBySubscription.json
func ExampleNamespacesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespacesClient().NewListBySubscriptionPager(&armeventgrid.NamespacesClientListBySubscriptionOptions{Filter: nil,
		Top: nil,
	})
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
		// page.NamespacesListResult = armeventgrid.NamespacesListResult{
		// 	Value: []*armeventgrid.Namespace{
		// 		{
		// 			Name: to.Ptr("exampleNamespaceName1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/namespaces"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 				"key3": to.Ptr("value3"),
		// 			},
		// 			Properties: &armeventgrid.NamespaceProperties{
		// 				ProvisioningState: to.Ptr(armeventgrid.NamespaceProvisioningStateSucceeded),
		// 				TopicSpacesConfiguration: &armeventgrid.TopicSpacesConfiguration{
		// 					Hostname: to.Ptr("exampleNamespaceName1.westus-1.mqtt.eventgrid-int.azure.net"),
		// 					RouteTopicResourceID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampleTopic1"),
		// 					State: to.Ptr(armeventgrid.TopicSpacesConfigurationStateEnabled),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Namespaces_ListByResourceGroup.json
func ExampleNamespacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespacesClient().NewListByResourceGroupPager("examplerg", &armeventgrid.NamespacesClientListByResourceGroupOptions{Filter: nil,
		Top: nil,
	})
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
		// page.NamespacesListResult = armeventgrid.NamespacesListResult{
		// 	Value: []*armeventgrid.Namespace{
		// 		{
		// 			Name: to.Ptr("exampleNamespaceName1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/namespaces"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 				"key3": to.Ptr("value3"),
		// 			},
		// 			Properties: &armeventgrid.NamespaceProperties{
		// 				ProvisioningState: to.Ptr(armeventgrid.NamespaceProvisioningStateSucceeded),
		// 				TopicSpacesConfiguration: &armeventgrid.TopicSpacesConfiguration{
		// 					Hostname: to.Ptr("exampleNamespaceName1.westus-1.mqtt.eventgrid-int.azure.net"),
		// 					RouteTopicResourceID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampleTopic1"),
		// 					State: to.Ptr(armeventgrid.TopicSpacesConfigurationStateEnabled),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Namespaces_ListSharedAccessKeys.json
func ExampleNamespacesClient_ListSharedAccessKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().ListSharedAccessKeys(ctx, "examplerg", "exampleNamespaceName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NamespaceSharedAccessKeys = armeventgrid.NamespaceSharedAccessKeys{
	// 	Key1: to.Ptr("testKey1Value"),
	// 	Key2: to.Ptr("testKey2Value"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Namespaces_RegenerateKey.json
func ExampleNamespacesClient_BeginRegenerateKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespacesClient().BeginRegenerateKey(ctx, "examplerg", "exampleNamespaceName1", armeventgrid.NamespaceRegenerateKeyRequest{
		KeyName: to.Ptr("key1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NamespaceSharedAccessKeys = armeventgrid.NamespaceSharedAccessKeys{
	// 	Key1: to.Ptr("testKey1Value"),
	// 	Key2: to.Ptr("testKey2Value"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Namespaces_ValidateCustomDomainOwnership.json
func ExampleNamespacesClient_BeginValidateCustomDomainOwnership() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespacesClient().BeginValidateCustomDomainOwnership(ctx, "examplerg", "exampleNamespaceName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomDomainOwnershipValidationResult = armeventgrid.CustomDomainOwnershipValidationResult{
	// 	CustomDomainsForTopicSpacesConfiguration: []*armeventgrid.CustomDomainConfiguration{
	// 		{
	// 			CertificateURL: to.Ptr("string"),
	// 			ExpectedTxtRecordName: to.Ptr("txt-record-name-2"),
	// 			ExpectedTxtRecordValue: to.Ptr("txt-record-value-2"),
	// 			FullyQualifiedDomainName: to.Ptr("custom-domain-name-2"),
	// 			Identity: &armeventgrid.CustomDomainIdentity{
	// 				Type: to.Ptr(armeventgrid.CustomDomainIdentityTypeSystemAssigned),
	// 			},
	// 			ValidationState: to.Ptr(armeventgrid.CustomDomainValidationStatePending),
	// 	}},
	// 	CustomDomainsForTopicsConfiguration: []*armeventgrid.CustomDomainConfiguration{
	// 		{
	// 			CertificateURL: to.Ptr("string"),
	// 			ExpectedTxtRecordName: to.Ptr("txt-record-name-1"),
	// 			ExpectedTxtRecordValue: to.Ptr("txt-record-value-1"),
	// 			FullyQualifiedDomainName: to.Ptr("custom-domain-name-1"),
	// 			Identity: &armeventgrid.CustomDomainIdentity{
	// 				Type: to.Ptr(armeventgrid.CustomDomainIdentityTypeSystemAssigned),
	// 			},
	// 			ValidationState: to.Ptr(armeventgrid.CustomDomainValidationStatePending),
	// 	}},
	// }
}
