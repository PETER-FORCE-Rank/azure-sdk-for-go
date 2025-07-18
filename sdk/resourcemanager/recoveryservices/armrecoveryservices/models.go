// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservices

import "time"

// AssociatedIdentity - Identity details to be used for an operation
type AssociatedIdentity struct {
	// Identity type that should be used for an operation.
	OperationIdentityType *IdentityType

	// User assigned identity to be used for an operation if operationIdentityType is UserAssigned.
	UserAssignedIdentity *string
}

// AzureMonitorAlertSettings - Settings for Azure Monitor based alerts
type AzureMonitorAlertSettings struct {
	AlertsForAllFailoverIssues    *AlertsState
	AlertsForAllJobFailures       *AlertsState
	AlertsForAllReplicationIssues *AlertsState
}

// CapabilitiesProperties - Capabilities information
type CapabilitiesProperties struct {
	DNSZones []*DNSZone
}

// CapabilitiesResponse - Capabilities response for Microsoft.RecoveryServices
type CapabilitiesResponse struct {
	// REQUIRED; Describes the Resource type: Microsoft.RecoveryServices/Vaults
	Type *string

	// Capabilities properties in response
	Properties *CapabilitiesResponseProperties
}

// CapabilitiesResponseProperties - Capabilities properties in response
type CapabilitiesResponseProperties struct {
	DNSZones []*DNSZoneResponse
}

// CertificateRequest - Details of the certificate to be uploaded to the vault.
type CertificateRequest struct {
	// Raw certificate data.
	Properties *RawCertificateData
}

// CheckNameAvailabilityParameters - Resource Name availability input parameters - Resource type and resource name
type CheckNameAvailabilityParameters struct {
	// Resource name for which availability needs to be checked
	Name *string

	// Describes the Resource type: Microsoft.RecoveryServices/Vaults
	Type *string
}

// CheckNameAvailabilityResult - Response for check name availability API. Resource provider will set availability as true
// | false.
type CheckNameAvailabilityResult struct {
	Message       *string
	NameAvailable *bool
	Reason        *string
}

// ClassicAlertSettings - Settings for classic alerts
type ClassicAlertSettings struct {
	AlertsForCriticalOperations       *AlertsState
	EmailNotificationsForSiteRecovery *AlertsState
}

// ClientDiscoveryDisplay - Localized display information of an operation.
type ClientDiscoveryDisplay struct {
	// Description of the operation having details of what operation is about.
	Description *string

	// Operations Name itself.
	Operation *string

	// Name of the provider for display purposes
	Provider *string

	// ResourceType for which this Operation can be performed.
	Resource *string
}

// ClientDiscoveryForLogSpecification - Class to represent shoebox log specification in json client discovery.
type ClientDiscoveryForLogSpecification struct {
	// Blobs created in customer storage account per hour
	BlobDuration *string

	// Localized display name
	DisplayName *string

	// Name of the log.
	Name *string
}

// ClientDiscoveryForProperties - Class to represent shoebox properties in json client discovery.
type ClientDiscoveryForProperties struct {
	// Operation properties.
	ServiceSpecification *ClientDiscoveryForServiceSpecification
}

// ClientDiscoveryForServiceSpecification - Class to represent shoebox service specification in json client discovery.
type ClientDiscoveryForServiceSpecification struct {
	// List of log specifications of this operation.
	LogSpecifications []*ClientDiscoveryForLogSpecification
}

// ClientDiscoveryResponse - Operations List response which contains list of available APIs.
type ClientDiscoveryResponse struct {
	// Link to the next chunk of the response
	NextLink *string

	// List of available operations.
	Value []*ClientDiscoveryValueForSingleAPI
}

// ClientDiscoveryValueForSingleAPI - Available operation details.
type ClientDiscoveryValueForSingleAPI struct {
	// Contains the localized display information for this particular operation
	Display *ClientDiscoveryDisplay

	// Name of the Operation.
	Name *string

	// The intended executor of the operation;governs the display of the operation in the RBAC UX and the audit logs UX
	Origin *string

	// ShoeBox properties for the given operation.
	Properties *ClientDiscoveryForProperties
}

// CmkKekIdentity - The details of the identity used for CMK
type CmkKekIdentity struct {
	// Indicate that system assigned identity should be used. Mutually exclusive with 'userAssignedIdentity' field
	UseSystemAssignedIdentity *bool

	// The user assigned identity to be used to grant permissions in case the type of identity used is UserAssigned
	UserAssignedIdentity *string
}

// CmkKeyVaultProperties - The properties of the Key Vault which hosts CMK
type CmkKeyVaultProperties struct {
	// The key uri of the Customer Managed Key
	KeyURI *string
}

// CrossSubscriptionRestoreSettings - Settings for Cross Subscription Restore Settings
type CrossSubscriptionRestoreSettings struct {
	CrossSubscriptionRestoreState *CrossSubscriptionRestoreState
}

// DNSZone information
type DNSZone struct {
	// Subresource type for vault AzureBackup, AzureBackup_secondary or AzureSiteRecovery
	SubResource *VaultSubResourceType
}

// DNSZoneResponse - DNSZone information for Microsoft.RecoveryServices
type DNSZoneResponse struct {
	// The private link resource Private link DNS zone names.
	RequiredZoneNames []*string

	// Subresource type for vault AzureBackup, AzureBackup_secondary or AzureSiteRecovery
	SubResource *VaultSubResourceType
}

// Error - The resource management error response.
type Error struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*Error

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// IdentityData - Identity for the resource.
type IdentityData struct {
	// REQUIRED; The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created
	// identity and a set of user-assigned identities. The type 'None' will remove any
	// identities.
	Type *ResourceIdentityType

	// The list of user-assigned identities associated with the resource. The user-assigned identity dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]*UserIdentity

	// READ-ONLY; The principal ID of resource identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of resource.
	TenantID *string
}

// ImmutabilitySettings - Immutability Settings of vault
type ImmutabilitySettings struct {
	State *ImmutabilityState
}

// JobsSummary - Summary of the replication job data for this vault.
type JobsSummary struct {
	// Count of failed jobs.
	FailedJobs *int32

	// Count of in-progress jobs.
	InProgressJobs *int32

	// Count of suspended jobs.
	SuspendedJobs *int32
}

// MonitoringSettings - Monitoring Settings of the vault
type MonitoringSettings struct {
	// Settings for Azure Monitor based alerts
	AzureMonitorAlertSettings *AzureMonitorAlertSettings

	// Settings for classic alerts
	ClassicAlertSettings *ClassicAlertSettings
}

// MonitoringSummary - Summary of the replication monitoring data for this vault.
type MonitoringSummary struct {
	// Count of all deprecated recovery service providers.
	DeprecatedProviderCount *int32

	// Count of all critical warnings.
	EventsCount *int32

	// Count of all the supported recovery service providers.
	SupportedProviderCount *int32

	// Count of unhealthy replication providers.
	UnHealthyProviderCount *int32

	// Count of unhealthy VMs.
	UnHealthyVMCount *int32

	// Count of all the unsupported recovery service providers.
	UnsupportedProviderCount *int32
}

// NameInfo - The name of usage.
type NameInfo struct {
	// Localized value of usage.
	LocalizedValue *string

	// Value of usage.
	Value *string
}

// OperationResource - Operation Resource
type OperationResource struct {
	// End time of the operation
	EndTime *time.Time

	// Required if status == failed or status == canceled. This is the OData v4 error format, used by the RPC and will go into
	// the v2.2 Azure REST API guidelines.
	Error *Error

	// It should match what is used to GET the operation result
	ID *string

	// It must match the last segment of the "id" field, and will typically be a GUID / system generated value
	Name *string

	// Start time of the operation
	StartTime *time.Time

	// The status of the operation. (InProgress/Success/Failed/Cancelled)
	Status *string
}

// PatchTrackedResource - Tracked resource with location.
type PatchTrackedResource struct {
	// Optional ETag.
	Etag *string

	// Resource location.
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Resource Id represents the complete path to the resource.
	ID *string

	// READ-ONLY; Resource name associated with the resource.
	Name *string

	// READ-ONLY; Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/…
	Type *string
}

// PatchVault - Patch Resource information, as returned by the resource provider.
type PatchVault struct {
	// Optional ETag.
	Etag *string

	// Identity for the resource.
	Identity *IdentityData

	// Resource location.
	Location *string

	// Properties of the vault.
	Properties *VaultProperties

	// Identifies the unique system identifier for each Azure resource.
	SKU *SKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Resource Id represents the complete path to the resource.
	ID *string

	// READ-ONLY; Resource name associated with the resource.
	Name *string

	// READ-ONLY; Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/…
	Type *string
}

// PrivateEndpoint - The Private Endpoint network resource that is linked to the Private Endpoint connection.
type PrivateEndpoint struct {
	// READ-ONLY; Gets or sets id.
	ID *string
}

// PrivateEndpointConnection - Private Endpoint Connection Response Properties.
type PrivateEndpointConnection struct {
	// Group Ids for the Private Endpoint
	GroupIDs []*VaultSubResourceType

	// READ-ONLY; The Private Endpoint network resource that is linked to the Private Endpoint connection.
	PrivateEndpoint *PrivateEndpoint

	// READ-ONLY; Gets or sets private link service connection state.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// READ-ONLY; Gets or sets provisioning state of the private endpoint connection.
	ProvisioningState *ProvisioningState
}

// PrivateEndpointConnectionVaultProperties - Information to be stored in Vault properties as an element of privateEndpointConnections
// List.
type PrivateEndpointConnectionVaultProperties struct {
	// READ-ONLY; Format of id subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.[Service]/{resource}/{resourceName}/privateEndpointConnections/{connectionName}.
	ID *string

	// READ-ONLY; The location of the private Endpoint connection
	Location *string

	// READ-ONLY; The name of the private Endpoint Connection
	Name *string

	// READ-ONLY; Private Endpoint Connection Response Properties.
	Properties *PrivateEndpointConnection

	// READ-ONLY; The type, which will be of the format, Microsoft.RecoveryServices/vaults/privateEndpointConnections
	Type *string
}

// PrivateLinkResource - Information of the private link resource.
type PrivateLinkResource struct {
	// Resource properties
	Properties *PrivateLinkResourceProperties

	// READ-ONLY; Fully qualified identifier of the resource.
	ID *string

	// READ-ONLY; Name of the resource.
	Name *string

	// READ-ONLY; e.g. Microsoft.RecoveryServices/vaults/privateLinkResources
	Type *string
}

// PrivateLinkResourceProperties - Properties of the private link resource.
type PrivateLinkResourceProperties struct {
	// READ-ONLY; e.g. f9ad6492-33d4-4690-9999-6bfd52a0d081 (Backup) or f9ad6492-33d4-4690-9999-6bfd52a0d082 (SiteRecovery)
	GroupID *string

	// READ-ONLY; [backup-ecs1, backup-prot1, backup-prot1b, backup-prot1c, backup-id1]
	RequiredMembers []*string

	// READ-ONLY; The private link resource Private link DNS zone name.
	RequiredZoneNames []*string
}

// PrivateLinkResources - Class which represent the stamps associated with the vault.
type PrivateLinkResources struct {
	// Link to the next chunk of the response
	NextLink *string

	// A collection of private link resources
	Value []*PrivateLinkResource
}

// PrivateLinkServiceConnectionState - Gets or sets private link service connection state.
type PrivateLinkServiceConnectionState struct {
	// READ-ONLY; Gets or sets actions required.
	ActionsRequired *string

	// READ-ONLY; Gets or sets description.
	Description *string

	// READ-ONLY; Gets or sets the status.
	Status *PrivateEndpointConnectionStatus
}

// RawCertificateData - Raw certificate data.
type RawCertificateData struct {
	// Specifies the authentication type.
	AuthType *AuthType

	// The base64 encoded certificate raw data string
	Certificate []byte
}

// ReplicationUsage - Replication usages of a vault.
type ReplicationUsage struct {
	// Summary of the replication jobs data for this vault.
	JobsSummary *JobsSummary

	// Summary of the replication monitoring data for this vault.
	MonitoringSummary *MonitoringSummary

	// Number of replication protected items for this vault.
	ProtectedItemCount *int32

	// Number of replication recovery plans for this vault.
	RecoveryPlanCount *int32

	// The authentication type of recovery service providers in the vault.
	RecoveryServicesProviderAuthType *int32

	// Number of servers registered to this vault.
	RegisteredServersCount *int32
}

// ReplicationUsageList - Replication usages for vault.
type ReplicationUsageList struct {
	// The list of replication usages for the given vault.
	Value []*ReplicationUsage
}

// Resource - ARM Resource.
type Resource struct {
	// Optional ETag.
	Etag *string

	// READ-ONLY; Resource Id represents the complete path to the resource.
	ID *string

	// READ-ONLY; Resource name associated with the resource.
	Name *string

	// READ-ONLY; Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/…
	Type *string
}

// ResourceCapabilities - Input to get capabilities information for Microsoft.RecoveryServices
type ResourceCapabilities struct {
	// REQUIRED; Describes the Resource type: Microsoft.RecoveryServices/Vaults
	Type *string

	// Capabilities information
	Properties *CapabilitiesProperties
}

// ResourceCapabilitiesBase - Base class for request and response capabilities information for Microsoft.RecoveryServices
type ResourceCapabilitiesBase struct {
	// REQUIRED; Describes the Resource type: Microsoft.RecoveryServices/Vaults
	Type *string
}

// ResourceCertificateAndAADDetails - Certificate details representing the Vault credentials for AAD.
type ResourceCertificateAndAADDetails struct {
	// REQUIRED; AAD tenant authority.
	AADAuthority *string

	// REQUIRED; AAD tenant Id.
	AADTenantID *string

	// REQUIRED; This property will be used as the discriminator for deciding the specific types in the polymorphic chain of types.
	AuthType *string

	// REQUIRED; Azure Management Endpoint Audience.
	AzureManagementEndpointAudience *string

	// REQUIRED; AAD service principal clientId.
	ServicePrincipalClientID *string

	// REQUIRED; AAD service principal ObjectId.
	ServicePrincipalObjectID *string

	// AAD audience for the resource
	AADAudience *string

	// The base64 encoded certificate raw data string.
	Certificate []byte

	// Certificate friendly name.
	FriendlyName *string

	// Certificate issuer.
	Issuer *string

	// Resource ID of the vault.
	ResourceID *int64

	// Service Resource Id.
	ServiceResourceID *string

	// Certificate Subject Name.
	Subject *string

	// Certificate thumbprint.
	Thumbprint *string

	// Certificate Validity start Date time.
	ValidFrom *time.Time

	// Certificate Validity End Date time.
	ValidTo *time.Time
}

// GetResourceCertificateDetails implements the ResourceCertificateDetailsClassification interface for type ResourceCertificateAndAADDetails.
func (r *ResourceCertificateAndAADDetails) GetResourceCertificateDetails() *ResourceCertificateDetails {
	return &ResourceCertificateDetails{
		AuthType:     r.AuthType,
		Certificate:  r.Certificate,
		FriendlyName: r.FriendlyName,
		Issuer:       r.Issuer,
		ResourceID:   r.ResourceID,
		Subject:      r.Subject,
		Thumbprint:   r.Thumbprint,
		ValidFrom:    r.ValidFrom,
		ValidTo:      r.ValidTo,
	}
}

// ResourceCertificateAndAcsDetails - Certificate details representing the Vault credentials for ACS.
type ResourceCertificateAndAcsDetails struct {
	// REQUIRED; This property will be used as the discriminator for deciding the specific types in the polymorphic chain of types.
	AuthType *string

	// REQUIRED; Acs mgmt host name to connect to.
	GlobalAcsHostName *string

	// REQUIRED; ACS namespace name - tenant for our service.
	GlobalAcsNamespace *string

	// REQUIRED; Global ACS namespace RP realm.
	GlobalAcsRPRealm *string

	// The base64 encoded certificate raw data string.
	Certificate []byte

	// Certificate friendly name.
	FriendlyName *string

	// Certificate issuer.
	Issuer *string

	// Resource ID of the vault.
	ResourceID *int64

	// Certificate Subject Name.
	Subject *string

	// Certificate thumbprint.
	Thumbprint *string

	// Certificate Validity start Date time.
	ValidFrom *time.Time

	// Certificate Validity End Date time.
	ValidTo *time.Time
}

// GetResourceCertificateDetails implements the ResourceCertificateDetailsClassification interface for type ResourceCertificateAndAcsDetails.
func (r *ResourceCertificateAndAcsDetails) GetResourceCertificateDetails() *ResourceCertificateDetails {
	return &ResourceCertificateDetails{
		AuthType:     r.AuthType,
		Certificate:  r.Certificate,
		FriendlyName: r.FriendlyName,
		Issuer:       r.Issuer,
		ResourceID:   r.ResourceID,
		Subject:      r.Subject,
		Thumbprint:   r.Thumbprint,
		ValidFrom:    r.ValidFrom,
		ValidTo:      r.ValidTo,
	}
}

// ResourceCertificateDetails - Certificate details representing the Vault credentials.
type ResourceCertificateDetails struct {
	// REQUIRED; This property will be used as the discriminator for deciding the specific types in the polymorphic chain of types.
	AuthType *string

	// The base64 encoded certificate raw data string.
	Certificate []byte

	// Certificate friendly name.
	FriendlyName *string

	// Certificate issuer.
	Issuer *string

	// Resource ID of the vault.
	ResourceID *int64

	// Certificate Subject Name.
	Subject *string

	// Certificate thumbprint.
	Thumbprint *string

	// Certificate Validity start Date time.
	ValidFrom *time.Time

	// Certificate Validity End Date time.
	ValidTo *time.Time
}

// GetResourceCertificateDetails implements the ResourceCertificateDetailsClassification interface for type ResourceCertificateDetails.
func (r *ResourceCertificateDetails) GetResourceCertificateDetails() *ResourceCertificateDetails {
	return r
}

// RestoreSettings - Restore Settings of the vault
type RestoreSettings struct {
	// Settings for CrossSubscriptionRestore
	CrossSubscriptionRestoreSettings *CrossSubscriptionRestoreSettings
}

// SKU - Identifies the unique system identifier for each Azure resource.
type SKU struct {
	// REQUIRED; Name of SKU is RS0 (Recovery Services 0th version) and the tier is standard tier. They do not have affect on
	// backend storage redundancy or any other vault settings. To manage storage redundancy, use
	// the backupstorageconfig
	Name *SKUName

	// The sku capacity
	Capacity *string

	// The sku family
	Family *string

	// The sku size
	Size *string

	// The Sku tier.
	Tier *string
}

// SecuritySettings - Security Settings of the vault
type SecuritySettings struct {
	// Immutability Settings of a vault
	ImmutabilitySettings *ImmutabilitySettings

	// Soft delete Settings of a vault
	SoftDeleteSettings *SoftDeleteSettings

	// Source scan configuration of vault
	SourceScanConfiguration *SourceScanConfiguration

	// READ-ONLY; MUA Settings of a vault
	MultiUserAuthorization *MultiUserAuthorization
}

// SoftDeleteSettings - Soft delete Settings of vault
type SoftDeleteSettings struct {
	EnhancedSecurityState *EnhancedSecurityState

	// Soft delete retention period in days
	SoftDeleteRetentionPeriodInDays *int32
	SoftDeleteState                 *SoftDeleteState
}

// SourceScanConfiguration - Source scan configuration of vault
type SourceScanConfiguration struct {
	// Identity details to be used for an operation
	SourceScanIdentity *AssociatedIdentity
	State              *State
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The type of identity that last modified the resource.
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TrackedResource - Tracked resource with location.
type TrackedResource struct {
	// REQUIRED; Resource location.
	Location *string

	// Optional ETag.
	Etag *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Resource Id represents the complete path to the resource.
	ID *string

	// READ-ONLY; Resource name associated with the resource.
	Name *string

	// READ-ONLY; Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/…
	Type *string
}

// UpgradeDetails - Details for upgrading vault.
type UpgradeDetails struct {
	// READ-ONLY; UTC time at which the upgrade operation has ended.
	EndTimeUTC *time.Time

	// READ-ONLY; UTC time at which the upgrade operation status was last updated.
	LastUpdatedTimeUTC *time.Time

	// READ-ONLY; Message to the user containing information about the upgrade operation.
	Message *string

	// READ-ONLY; ID of the vault upgrade operation.
	OperationID *string

	// READ-ONLY; Resource ID of the vault before the upgrade.
	PreviousResourceID *string

	// READ-ONLY; UTC time at which the upgrade operation has started.
	StartTimeUTC *time.Time

	// READ-ONLY; Status of the vault upgrade operation.
	Status *VaultUpgradeState

	// READ-ONLY; The way the vault upgrade was triggered.
	TriggerType *TriggerType

	// READ-ONLY; Resource ID of the upgraded vault.
	UpgradedResourceID *string
}

// UserIdentity - A resource identity that is managed by the user of the service.
type UserIdentity struct {
	// READ-ONLY; The client ID of the user-assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the user-assigned identity.
	PrincipalID *string
}

// Vault - Resource information, as returned by the resource provider.
type Vault struct {
	// REQUIRED; Resource location.
	Location *string

	// Optional ETag.
	Etag *string

	// Identity for the resource.
	Identity *IdentityData

	// Properties of the vault.
	Properties *VaultProperties

	// Identifies the unique system identifier for each Azure resource.
	SKU *SKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Resource Id represents the complete path to the resource.
	ID *string

	// READ-ONLY; Resource name associated with the resource.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/…
	Type *string
}

// VaultCertificateResponse - Certificate corresponding to a vault that can be used by clients to register themselves with
// the vault.
type VaultCertificateResponse struct {
	// Certificate details representing the Vault credentials.
	Properties ResourceCertificateDetailsClassification

	// READ-ONLY; Resource Id represents the complete path to the resource.
	ID *string

	// READ-ONLY; Resource name associated with the resource.
	Name *string

	// READ-ONLY; Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/…
	Type *string
}

// VaultExtendedInfo - Vault extended information.
type VaultExtendedInfo struct {
	// Algorithm for Vault ExtendedInfo
	Algorithm *string

	// Encryption key.
	EncryptionKey *string

	// Encryption key thumbprint.
	EncryptionKeyThumbprint *string

	// Integrity key.
	IntegrityKey *string
}

// VaultExtendedInfoResource - Vault extended information.
type VaultExtendedInfoResource struct {
	// Optional ETag.
	Etag *string

	// Vault extended information.
	Properties *VaultExtendedInfo

	// READ-ONLY; Resource Id represents the complete path to the resource.
	ID *string

	// READ-ONLY; Resource name associated with the resource.
	Name *string

	// READ-ONLY; Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/…
	Type *string
}

// VaultList - The response model for a list of Vaults.
type VaultList struct {
	Value []*Vault

	// READ-ONLY
	NextLink *string
}

// VaultProperties - Properties of the vault.
type VaultProperties struct {
	// Customer Managed Key details of the resource.
	Encryption *VaultPropertiesEncryption

	// Monitoring Settings of the vault
	MonitoringSettings *MonitoringSettings

	// The details of the latest move operation performed on the Azure Resource
	MoveDetails *VaultPropertiesMoveDetails

	// property to enable or disable resource provider inbound network traffic from public clients
	PublicNetworkAccess *PublicNetworkAccess

	// The redundancy Settings of a Vault
	RedundancySettings *VaultPropertiesRedundancySettings

	// ResourceGuardOperationRequests on which LAC check will be performed
	ResourceGuardOperationRequests []*string

	// Restore Settings of the vault
	RestoreSettings *RestoreSettings

	// Security Settings of the vault
	SecuritySettings *SecuritySettings

	// Details for upgrading vault.
	UpgradeDetails *UpgradeDetails

	// READ-ONLY; Backup storage version
	BackupStorageVersion *BackupStorageVersion

	// READ-ONLY; Security levels of Recovery Services Vault for business continuity and disaster recovery
	BcdrSecurityLevel *BCDRSecurityLevel

	// READ-ONLY; The State of the Resource after the move operation
	MoveState *ResourceMoveState

	// READ-ONLY; List of private endpoint connection.
	PrivateEndpointConnections []*PrivateEndpointConnectionVaultProperties

	// READ-ONLY; Private endpoint state for backup.
	PrivateEndpointStateForBackup *VaultPrivateEndpointState

	// READ-ONLY; Private endpoint state for site recovery.
	PrivateEndpointStateForSiteRecovery *VaultPrivateEndpointState

	// READ-ONLY; Provisioning State.
	ProvisioningState *string

	// READ-ONLY; Secure Score of Recovery Services Vault
	SecureScore *SecureScoreLevel
}

// VaultPropertiesEncryption - Customer Managed Key details of the resource.
type VaultPropertiesEncryption struct {
	// Enabling/Disabling the Double Encryption state
	InfrastructureEncryption *InfrastructureEncryptionState

	// The details of the identity used for CMK
	KekIdentity *CmkKekIdentity

	// The properties of the Key Vault which hosts CMK
	KeyVaultProperties *CmkKeyVaultProperties
}

// VaultPropertiesMoveDetails - The details of the latest move operation performed on the Azure Resource
type VaultPropertiesMoveDetails struct {
	// READ-ONLY; End Time of the Resource Move Operation
	CompletionTimeUTC *time.Time

	// READ-ONLY; OperationId of the Resource Move Operation
	OperationID *string

	// READ-ONLY; Source Resource of the Resource Move Operation
	SourceResourceID *string

	// READ-ONLY; Start Time of the Resource Move Operation
	StartTimeUTC *time.Time

	// READ-ONLY; Target Resource of the Resource Move Operation
	TargetResourceID *string
}

// VaultPropertiesRedundancySettings - The redundancy Settings of a Vault
type VaultPropertiesRedundancySettings struct {
	// Flag to show if Cross Region Restore is enabled on the Vault or not
	CrossRegionRestore *CrossRegionRestore

	// The storage redundancy setting of a vault
	StandardTierStorageRedundancy *StandardTierStorageRedundancy
}

// VaultUsage - Usages of a vault.
type VaultUsage struct {
	// Current value of usage.
	CurrentValue *int64

	// Limit of usage.
	Limit *int64

	// Name of usage.
	Name *NameInfo

	// Next reset time of usage.
	NextResetTime *time.Time

	// Quota period of usage.
	QuotaPeriod *string

	// Unit of the usage.
	Unit *UsagesUnit
}

// VaultUsageList - Usage for vault.
type VaultUsageList struct {
	// The list of usages for the given vault.
	Value []*VaultUsage
}
