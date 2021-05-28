package media

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/mediaservices/mgmt/2015-10-01/media"

// APIEndpoint the properties for a Media Services REST API endpoint.
type APIEndpoint struct {
	// Endpoint - The Media Services REST endpoint.
	Endpoint *string `json:"endpoint,omitempty"`
	// MajorVersion - The version of Media Services REST API.
	MajorVersion *string `json:"majorVersion,omitempty"`
}

// APIError the error returned from a failed Media Services REST API call.
type APIError struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message.
	Message *string `json:"message,omitempty"`
}

// CheckNameAvailabilityInput the request body for CheckNameAvailability API.
type CheckNameAvailabilityInput struct {
	// Name - The name of the resource. A name must be globally unique.
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource - mediaservices.
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityOutput the response body for CheckNameAvailability API.
type CheckNameAvailabilityOutput struct {
	autorest.Response `json:"-"`
	// NameAvailable - Specifies if the name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - Specifies the reason if the name is not available. Possible values include: 'None', 'Invalid', 'AlreadyExists'
	Reason EntityNameUnavailabilityReason `json:"reason,omitempty"`
	// Message - Specifies the detailed reason if the name is not available.
	Message *string `json:"message,omitempty"`
}

// Operation a Media Services REST API operation
type Operation struct {
	// Name - READ-ONLY; Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// MarshalJSON is the custom marshaler for Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if o.Display != nil {
		objectMap["display"] = o.Display
	}
	return json.Marshal(objectMap)
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - READ-ONLY; Service provider: Microsoft.Media
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; Resource on which the operation is performed: Invoice, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - READ-ONLY; Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// OperationListResult result of the request to list Media Services operations.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of Media Services operations supported by the Microsoft.Media resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationListResult.
func (olr OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// RegenerateKeyInput the request body for a RegenerateKey API.
type RegenerateKeyInput struct {
	// KeyType - The keyType indicating which key you want to regenerate, Primary or Secondary. Possible values include: 'Primary', 'Secondary'
	KeyType KeyType `json:"keyType,omitempty"`
}

// RegenerateKeyOutput the response body for a RegenerateKey API.
type RegenerateKeyOutput struct {
	autorest.Response `json:"-"`
	// Key - The new value of either the primary or secondary key.
	Key *string `json:"key,omitempty"`
}

// Resource the Azure Resource Manager resource.
type Resource struct {
	// ID - READ-ONLY; The id of the resource.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource
	Type *string `json:"type,omitempty"`
	// Location - The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth).
	Location *string `json:"location,omitempty"`
	// Tags - Tags to help categorize the resource in the Azure portal.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// Service the properties of a Media Service resource.
type Service struct {
	autorest.Response `json:"-"`
	// ServiceProperties - The additional properties of a Media Service resource.
	*ServiceProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The id of the resource.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource
	Type *string `json:"type,omitempty"`
	// Location - The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth).
	Location *string `json:"location,omitempty"`
	// Tags - Tags to help categorize the resource in the Azure portal.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Service.
func (s Service) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.ServiceProperties != nil {
		objectMap["properties"] = s.ServiceProperties
	}
	if s.Location != nil {
		objectMap["location"] = s.Location
	}
	if s.Tags != nil {
		objectMap["tags"] = s.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Service struct.
func (s *Service) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var serviceProperties ServiceProperties
				err = json.Unmarshal(*v, &serviceProperties)
				if err != nil {
					return err
				}
				s.ServiceProperties = &serviceProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				s.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				s.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				s.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				s.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				s.Tags = tags
			}
		}
	}

	return nil
}

// ServiceCollection the collection of Media Service resources.
type ServiceCollection struct {
	autorest.Response `json:"-"`
	// Value - The collection of Media Service resources.
	Value *[]Service `json:"value,omitempty"`
}

// ServiceKeys the response body for a ListKeys API.
type ServiceKeys struct {
	autorest.Response `json:"-"`
	// PrimaryAuthEndpoint - The primary authorization endpoint.
	PrimaryAuthEndpoint *string `json:"primaryAuthEndpoint,omitempty"`
	// SecondaryAuthEndpoint - The secondary authorization endpoint.
	SecondaryAuthEndpoint *string `json:"secondaryAuthEndpoint,omitempty"`
	// PrimaryKey - The primary key for the Media Service resource.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - The secondary key for the Media Service resource.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
	// Scope - The authorization scope.
	Scope *string `json:"scope,omitempty"`
}

// ServiceProperties the additional properties of a Media Service resource.
type ServiceProperties struct {
	// APIEndpoints - READ-ONLY; Read-only property that lists the Media Services REST API endpoints for this resource. If supplied on a PUT or PATCH, the value will be ignored.
	APIEndpoints *[]APIEndpoint `json:"apiEndpoints,omitempty"`
	// StorageAccounts - The storage accounts for this resource.
	StorageAccounts *[]StorageAccount `json:"storageAccounts,omitempty"`
}

// MarshalJSON is the custom marshaler for ServiceProperties.
func (sp ServiceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sp.StorageAccounts != nil {
		objectMap["storageAccounts"] = sp.StorageAccounts
	}
	return json.Marshal(objectMap)
}

// StorageAccount the properties of a storage account associated with this resource.
type StorageAccount struct {
	// ID - The id of the storage account resource. Media Services relies on tables and queues as well as blobs, so the primary storage account must be a Standard Storage account (either Microsoft.ClassicStorage or Microsoft.Storage). Blob only storage accounts can be added as secondary storage accounts (isPrimary false).
	ID *string `json:"id,omitempty"`
	// IsPrimary - Is this storage account resource the primary storage account for the Media Service resource. Blob only storage must set this to false.
	IsPrimary *bool `json:"isPrimary,omitempty"`
}

// SyncStorageKeysInput the request  body for a SyncStorageKeys API.
type SyncStorageKeysInput struct {
	// ID - The id of the storage account resource.
	ID *string `json:"id,omitempty"`
}
