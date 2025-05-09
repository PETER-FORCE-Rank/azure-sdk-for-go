// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetworkcloud

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewAgentPoolsClient creates a new instance of AgentPoolsClient.
func (c *ClientFactory) NewAgentPoolsClient() *AgentPoolsClient {
	return &AgentPoolsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBareMetalMachineKeySetsClient creates a new instance of BareMetalMachineKeySetsClient.
func (c *ClientFactory) NewBareMetalMachineKeySetsClient() *BareMetalMachineKeySetsClient {
	return &BareMetalMachineKeySetsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBareMetalMachinesClient creates a new instance of BareMetalMachinesClient.
func (c *ClientFactory) NewBareMetalMachinesClient() *BareMetalMachinesClient {
	return &BareMetalMachinesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBmcKeySetsClient creates a new instance of BmcKeySetsClient.
func (c *ClientFactory) NewBmcKeySetsClient() *BmcKeySetsClient {
	return &BmcKeySetsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCloudServicesNetworksClient creates a new instance of CloudServicesNetworksClient.
func (c *ClientFactory) NewCloudServicesNetworksClient() *CloudServicesNetworksClient {
	return &CloudServicesNetworksClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewClusterManagersClient creates a new instance of ClusterManagersClient.
func (c *ClientFactory) NewClusterManagersClient() *ClusterManagersClient {
	return &ClusterManagersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewClustersClient creates a new instance of ClustersClient.
func (c *ClientFactory) NewClustersClient() *ClustersClient {
	return &ClustersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConsolesClient creates a new instance of ConsolesClient.
func (c *ClientFactory) NewConsolesClient() *ConsolesClient {
	return &ConsolesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewKubernetesClusterFeaturesClient creates a new instance of KubernetesClusterFeaturesClient.
func (c *ClientFactory) NewKubernetesClusterFeaturesClient() *KubernetesClusterFeaturesClient {
	return &KubernetesClusterFeaturesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewKubernetesClustersClient creates a new instance of KubernetesClustersClient.
func (c *ClientFactory) NewKubernetesClustersClient() *KubernetesClustersClient {
	return &KubernetesClustersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewL2NetworksClient creates a new instance of L2NetworksClient.
func (c *ClientFactory) NewL2NetworksClient() *L2NetworksClient {
	return &L2NetworksClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewL3NetworksClient creates a new instance of L3NetworksClient.
func (c *ClientFactory) NewL3NetworksClient() *L3NetworksClient {
	return &L3NetworksClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMetricsConfigurationsClient creates a new instance of MetricsConfigurationsClient.
func (c *ClientFactory) NewMetricsConfigurationsClient() *MetricsConfigurationsClient {
	return &MetricsConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewRackSKUsClient creates a new instance of RackSKUsClient.
func (c *ClientFactory) NewRackSKUsClient() *RackSKUsClient {
	return &RackSKUsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRacksClient creates a new instance of RacksClient.
func (c *ClientFactory) NewRacksClient() *RacksClient {
	return &RacksClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewStorageAppliancesClient creates a new instance of StorageAppliancesClient.
func (c *ClientFactory) NewStorageAppliancesClient() *StorageAppliancesClient {
	return &StorageAppliancesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTrunkedNetworksClient creates a new instance of TrunkedNetworksClient.
func (c *ClientFactory) NewTrunkedNetworksClient() *TrunkedNetworksClient {
	return &TrunkedNetworksClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachinesClient creates a new instance of VirtualMachinesClient.
func (c *ClientFactory) NewVirtualMachinesClient() *VirtualMachinesClient {
	return &VirtualMachinesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVolumesClient creates a new instance of VolumesClient.
func (c *ClientFactory) NewVolumesClient() *VolumesClient {
	return &VolumesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
