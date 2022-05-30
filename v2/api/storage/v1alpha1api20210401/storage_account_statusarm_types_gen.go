// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

// Deprecated version of StorageAccount_STATUS. Use v1beta20210401.StorageAccount_STATUS instead
type StorageAccount_STATUSARM struct {
	ExtendedLocation *ExtendedLocation_STATUSARM         `json:"extendedLocation,omitempty"`
	Id               *string                             `json:"id,omitempty"`
	Identity         *Identity_STATUSARM                 `json:"identity,omitempty"`
	Kind             *StorageAccount_Kind_STATUS         `json:"kind,omitempty"`
	Location         *string                             `json:"location,omitempty"`
	Name             *string                             `json:"name,omitempty"`
	Properties       *StorageAccountProperties_STATUSARM `json:"properties,omitempty"`
	Sku              *Sku_STATUSARM                      `json:"sku,omitempty"`
	Tags             map[string]string                   `json:"tags,omitempty"`
	Type             *string                             `json:"type,omitempty"`
}

// Deprecated version of ExtendedLocation_STATUS. Use v1beta20210401.ExtendedLocation_STATUS instead
type ExtendedLocation_STATUSARM struct {
	Name *string                      `json:"name,omitempty"`
	Type *ExtendedLocationType_STATUS `json:"type,omitempty"`
}

// Deprecated version of Identity_STATUS. Use v1beta20210401.Identity_STATUS instead
type Identity_STATUSARM struct {
	PrincipalId            *string                                   `json:"principalId,omitempty"`
	TenantId               *string                                   `json:"tenantId,omitempty"`
	Type                   *Identity_Type_STATUS                     `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUSARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of Sku_STATUS. Use v1beta20210401.Sku_STATUS instead
type Sku_STATUSARM struct {
	Name *SkuName_STATUS `json:"name,omitempty"`
	Tier *Tier_STATUS    `json:"tier,omitempty"`
}

// Deprecated version of StorageAccountProperties_STATUS. Use v1beta20210401.StorageAccountProperties_STATUS instead
type StorageAccountProperties_STATUSARM struct {
	AccessTier                            *StorageAccountProperties_AccessTier_STATUS           `json:"accessTier,omitempty"`
	AllowBlobPublicAccess                 *bool                                                 `json:"allowBlobPublicAccess,omitempty"`
	AllowCrossTenantReplication           *bool                                                 `json:"allowCrossTenantReplication,omitempty"`
	AllowSharedKeyAccess                  *bool                                                 `json:"allowSharedKeyAccess,omitempty"`
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication_STATUSARM      `json:"azureFilesIdentityBasedAuthentication,omitempty"`
	BlobRestoreStatus                     *BlobRestoreStatus_STATUSARM                          `json:"blobRestoreStatus,omitempty"`
	CreationTime                          *string                                               `json:"creationTime,omitempty"`
	CustomDomain                          *CustomDomain_STATUSARM                               `json:"customDomain,omitempty"`
	Encryption                            *Encryption_STATUSARM                                 `json:"encryption,omitempty"`
	FailoverInProgress                    *bool                                                 `json:"failoverInProgress,omitempty"`
	GeoReplicationStats                   *GeoReplicationStats_STATUSARM                        `json:"geoReplicationStats,omitempty"`
	IsHnsEnabled                          *bool                                                 `json:"isHnsEnabled,omitempty"`
	IsNfsV3Enabled                        *bool                                                 `json:"isNfsV3Enabled,omitempty"`
	KeyCreationTime                       *KeyCreationTime_STATUSARM                            `json:"keyCreationTime,omitempty"`
	KeyPolicy                             *KeyPolicy_STATUSARM                                  `json:"keyPolicy,omitempty"`
	LargeFileSharesState                  *StorageAccountProperties_LargeFileSharesState_STATUS `json:"largeFileSharesState,omitempty"`
	LastGeoFailoverTime                   *string                                               `json:"lastGeoFailoverTime,omitempty"`
	MinimumTlsVersion                     *StorageAccountProperties_MinimumTlsVersion_STATUS    `json:"minimumTlsVersion,omitempty"`
	NetworkAcls                           *NetworkRuleSet_STATUSARM                             `json:"networkAcls,omitempty"`
	PrimaryEndpoints                      *Endpoints_STATUSARM                                  `json:"primaryEndpoints,omitempty"`
	PrimaryLocation                       *string                                               `json:"primaryLocation,omitempty"`
	PrivateEndpointConnections            []PrivateEndpointConnection_STATUSARM                 `json:"privateEndpointConnections,omitempty"`
	ProvisioningState                     *StorageAccountProperties_ProvisioningState_STATUS    `json:"provisioningState,omitempty"`
	RoutingPreference                     *RoutingPreference_STATUSARM                          `json:"routingPreference,omitempty"`
	SasPolicy                             *SasPolicy_STATUSARM                                  `json:"sasPolicy,omitempty"`
	SecondaryEndpoints                    *Endpoints_STATUSARM                                  `json:"secondaryEndpoints,omitempty"`
	SecondaryLocation                     *string                                               `json:"secondaryLocation,omitempty"`
	StatusOfPrimary                       *StorageAccountProperties_StatusOfPrimary_STATUS      `json:"statusOfPrimary,omitempty"`
	StatusOfSecondary                     *StorageAccountProperties_StatusOfSecondary_STATUS    `json:"statusOfSecondary,omitempty"`
	SupportsHttpsTrafficOnly              *bool                                                 `json:"supportsHttpsTrafficOnly,omitempty"`
}

// Deprecated version of StorageAccount_Kind_STATUS. Use v1beta20210401.StorageAccount_Kind_STATUS instead
type StorageAccount_Kind_STATUS string

const (
	StorageAccount_Kind_BlobStorage_STATUS      = StorageAccount_Kind_STATUS("BlobStorage")
	StorageAccount_Kind_BlockBlobStorage_STATUS = StorageAccount_Kind_STATUS("BlockBlobStorage")
	StorageAccount_Kind_FileStorage_STATUS      = StorageAccount_Kind_STATUS("FileStorage")
	StorageAccount_Kind_Storage_STATUS          = StorageAccount_Kind_STATUS("Storage")
	StorageAccount_Kind_StorageV2_STATUS        = StorageAccount_Kind_STATUS("StorageV2")
)

// Deprecated version of AzureFilesIdentityBasedAuthentication_STATUS. Use v1beta20210401.AzureFilesIdentityBasedAuthentication_STATUS instead
type AzureFilesIdentityBasedAuthentication_STATUSARM struct {
	ActiveDirectoryProperties *ActiveDirectoryProperties_STATUSARM                                  `json:"activeDirectoryProperties,omitempty"`
	DefaultSharePermission    *AzureFilesIdentityBasedAuthentication_DefaultSharePermission_STATUS  `json:"defaultSharePermission,omitempty"`
	DirectoryServiceOptions   *AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_STATUS `json:"directoryServiceOptions,omitempty"`
}

// Deprecated version of BlobRestoreStatus_STATUS. Use v1beta20210401.BlobRestoreStatus_STATUS instead
type BlobRestoreStatus_STATUSARM struct {
	FailureReason *string                          `json:"failureReason,omitempty"`
	Parameters    *BlobRestoreParameters_STATUSARM `json:"parameters,omitempty"`
	RestoreId     *string                          `json:"restoreId,omitempty"`
	Status        *BlobRestoreStatus_Status_STATUS `json:"status,omitempty"`
}

// Deprecated version of CustomDomain_STATUS. Use v1beta20210401.CustomDomain_STATUS instead
type CustomDomain_STATUSARM struct {
	Name             *string `json:"name,omitempty"`
	UseSubDomainName *bool   `json:"useSubDomainName,omitempty"`
}

// Deprecated version of Encryption_STATUS. Use v1beta20210401.Encryption_STATUS instead
type Encryption_STATUSARM struct {
	Identity                        *EncryptionIdentity_STATUSARM `json:"identity,omitempty"`
	KeySource                       *Encryption_KeySource_STATUS  `json:"keySource,omitempty"`
	Keyvaultproperties              *KeyVaultProperties_STATUSARM `json:"keyvaultproperties,omitempty"`
	RequireInfrastructureEncryption *bool                         `json:"requireInfrastructureEncryption,omitempty"`
	Services                        *EncryptionServices_STATUSARM `json:"services,omitempty"`
}

// Deprecated version of Endpoints_STATUS. Use v1beta20210401.Endpoints_STATUS instead
type Endpoints_STATUSARM struct {
	Blob               *string                                     `json:"blob,omitempty"`
	Dfs                *string                                     `json:"dfs,omitempty"`
	File               *string                                     `json:"file,omitempty"`
	InternetEndpoints  *StorageAccountInternetEndpoints_STATUSARM  `json:"internetEndpoints,omitempty"`
	MicrosoftEndpoints *StorageAccountMicrosoftEndpoints_STATUSARM `json:"microsoftEndpoints,omitempty"`
	Queue              *string                                     `json:"queue,omitempty"`
	Table              *string                                     `json:"table,omitempty"`
	Web                *string                                     `json:"web,omitempty"`
}

// Deprecated version of ExtendedLocationType_STATUS. Use v1beta20210401.ExtendedLocationType_STATUS instead
type ExtendedLocationType_STATUS string

const ExtendedLocationType_EdgeZone_STATUS = ExtendedLocationType_STATUS("EdgeZone")

// Deprecated version of GeoReplicationStats_STATUS. Use v1beta20210401.GeoReplicationStats_STATUS instead
type GeoReplicationStats_STATUSARM struct {
	CanFailover  *bool                              `json:"canFailover,omitempty"`
	LastSyncTime *string                            `json:"lastSyncTime,omitempty"`
	Status       *GeoReplicationStats_Status_STATUS `json:"status,omitempty"`
}

// Deprecated version of Identity_Type_STATUS. Use v1beta20210401.Identity_Type_STATUS instead
type Identity_Type_STATUS string

const (
	Identity_Type_None_STATUS                       = Identity_Type_STATUS("None")
	Identity_Type_SystemAssigned_STATUS             = Identity_Type_STATUS("SystemAssigned")
	Identity_Type_SystemAssignedUserAssigned_STATUS = Identity_Type_STATUS("SystemAssigned,UserAssigned")
	Identity_Type_UserAssigned_STATUS               = Identity_Type_STATUS("UserAssigned")
)

// Deprecated version of KeyCreationTime_STATUS. Use v1beta20210401.KeyCreationTime_STATUS instead
type KeyCreationTime_STATUSARM struct {
	Key1 *string `json:"key1,omitempty"`
	Key2 *string `json:"key2,omitempty"`
}

// Deprecated version of KeyPolicy_STATUS. Use v1beta20210401.KeyPolicy_STATUS instead
type KeyPolicy_STATUSARM struct {
	KeyExpirationPeriodInDays *int `json:"keyExpirationPeriodInDays,omitempty"`
}

// Deprecated version of NetworkRuleSet_STATUS. Use v1beta20210401.NetworkRuleSet_STATUS instead
type NetworkRuleSet_STATUSARM struct {
	Bypass              *NetworkRuleSet_Bypass_STATUS        `json:"bypass,omitempty"`
	DefaultAction       *NetworkRuleSet_DefaultAction_STATUS `json:"defaultAction,omitempty"`
	IpRules             []IPRule_STATUSARM                   `json:"ipRules,omitempty"`
	ResourceAccessRules []ResourceAccessRule_STATUSARM       `json:"resourceAccessRules,omitempty"`
	VirtualNetworkRules []VirtualNetworkRule_STATUSARM       `json:"virtualNetworkRules,omitempty"`
}

// Deprecated version of PrivateEndpointConnection_STATUS. Use v1beta20210401.PrivateEndpointConnection_STATUS instead
type PrivateEndpointConnection_STATUSARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of RoutingPreference_STATUS. Use v1beta20210401.RoutingPreference_STATUS instead
type RoutingPreference_STATUSARM struct {
	PublishInternetEndpoints  *bool                                   `json:"publishInternetEndpoints,omitempty"`
	PublishMicrosoftEndpoints *bool                                   `json:"publishMicrosoftEndpoints,omitempty"`
	RoutingChoice             *RoutingPreference_RoutingChoice_STATUS `json:"routingChoice,omitempty"`
}

// Deprecated version of SasPolicy_STATUS. Use v1beta20210401.SasPolicy_STATUS instead
type SasPolicy_STATUSARM struct {
	ExpirationAction    *SasPolicy_ExpirationAction_STATUS `json:"expirationAction,omitempty"`
	SasExpirationPeriod *string                            `json:"sasExpirationPeriod,omitempty"`
}

// Deprecated version of SkuName_STATUS. Use v1beta20210401.SkuName_STATUS instead
type SkuName_STATUS string

const (
	SkuName_Premium_LRS_STATUS     = SkuName_STATUS("Premium_LRS")
	SkuName_Premium_ZRS_STATUS     = SkuName_STATUS("Premium_ZRS")
	SkuName_Standard_GRS_STATUS    = SkuName_STATUS("Standard_GRS")
	SkuName_Standard_GZRS_STATUS   = SkuName_STATUS("Standard_GZRS")
	SkuName_Standard_LRS_STATUS    = SkuName_STATUS("Standard_LRS")
	SkuName_Standard_RAGRS_STATUS  = SkuName_STATUS("Standard_RAGRS")
	SkuName_Standard_RAGZRS_STATUS = SkuName_STATUS("Standard_RAGZRS")
	SkuName_Standard_ZRS_STATUS    = SkuName_STATUS("Standard_ZRS")
)

// Deprecated version of Tier_STATUS. Use v1beta20210401.Tier_STATUS instead
type Tier_STATUS string

const (
	Tier_Premium_STATUS  = Tier_STATUS("Premium")
	Tier_Standard_STATUS = Tier_STATUS("Standard")
)

// Deprecated version of UserAssignedIdentity_STATUS. Use v1beta20210401.UserAssignedIdentity_STATUS instead
type UserAssignedIdentity_STATUSARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

// Deprecated version of ActiveDirectoryProperties_STATUS. Use v1beta20210401.ActiveDirectoryProperties_STATUS instead
type ActiveDirectoryProperties_STATUSARM struct {
	AzureStorageSid   *string `json:"azureStorageSid,omitempty"`
	DomainGuid        *string `json:"domainGuid,omitempty"`
	DomainName        *string `json:"domainName,omitempty"`
	DomainSid         *string `json:"domainSid,omitempty"`
	ForestName        *string `json:"forestName,omitempty"`
	NetBiosDomainName *string `json:"netBiosDomainName,omitempty"`
}

// Deprecated version of BlobRestoreParameters_STATUS. Use v1beta20210401.BlobRestoreParameters_STATUS instead
type BlobRestoreParameters_STATUSARM struct {
	BlobRanges    []BlobRestoreRange_STATUSARM `json:"blobRanges,omitempty"`
	TimeToRestore *string                      `json:"timeToRestore,omitempty"`
}

// Deprecated version of EncryptionIdentity_STATUS. Use v1beta20210401.EncryptionIdentity_STATUS instead
type EncryptionIdentity_STATUSARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

// Deprecated version of EncryptionServices_STATUS. Use v1beta20210401.EncryptionServices_STATUS instead
type EncryptionServices_STATUSARM struct {
	Blob  *EncryptionService_STATUSARM `json:"blob,omitempty"`
	File  *EncryptionService_STATUSARM `json:"file,omitempty"`
	Queue *EncryptionService_STATUSARM `json:"queue,omitempty"`
	Table *EncryptionService_STATUSARM `json:"table,omitempty"`
}

// Deprecated version of IPRule_STATUS. Use v1beta20210401.IPRule_STATUS instead
type IPRule_STATUSARM struct {
	Action *IPRule_Action_STATUS `json:"action,omitempty"`
	Value  *string               `json:"value,omitempty"`
}

// Deprecated version of KeyVaultProperties_STATUS. Use v1beta20210401.KeyVaultProperties_STATUS instead
type KeyVaultProperties_STATUSARM struct {
	CurrentVersionedKeyIdentifier *string `json:"currentVersionedKeyIdentifier,omitempty"`
	Keyname                       *string `json:"keyname,omitempty"`
	Keyvaulturi                   *string `json:"keyvaulturi,omitempty"`
	Keyversion                    *string `json:"keyversion,omitempty"`
	LastKeyRotationTimestamp      *string `json:"lastKeyRotationTimestamp,omitempty"`
}

// Deprecated version of ResourceAccessRule_STATUS. Use v1beta20210401.ResourceAccessRule_STATUS instead
type ResourceAccessRule_STATUSARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
	TenantId   *string `json:"tenantId,omitempty"`
}

// Deprecated version of StorageAccountInternetEndpoints_STATUS. Use v1beta20210401.StorageAccountInternetEndpoints_STATUS instead
type StorageAccountInternetEndpoints_STATUSARM struct {
	Blob *string `json:"blob,omitempty"`
	Dfs  *string `json:"dfs,omitempty"`
	File *string `json:"file,omitempty"`
	Web  *string `json:"web,omitempty"`
}

// Deprecated version of StorageAccountMicrosoftEndpoints_STATUS. Use v1beta20210401.StorageAccountMicrosoftEndpoints_STATUS instead
type StorageAccountMicrosoftEndpoints_STATUSARM struct {
	Blob  *string `json:"blob,omitempty"`
	Dfs   *string `json:"dfs,omitempty"`
	File  *string `json:"file,omitempty"`
	Queue *string `json:"queue,omitempty"`
	Table *string `json:"table,omitempty"`
	Web   *string `json:"web,omitempty"`
}

// Deprecated version of VirtualNetworkRule_STATUS. Use v1beta20210401.VirtualNetworkRule_STATUS instead
type VirtualNetworkRule_STATUSARM struct {
	Action *VirtualNetworkRule_Action_STATUS `json:"action,omitempty"`
	Id     *string                           `json:"id,omitempty"`
	State  *VirtualNetworkRule_State_STATUS  `json:"state,omitempty"`
}

// Deprecated version of BlobRestoreRange_STATUS. Use v1beta20210401.BlobRestoreRange_STATUS instead
type BlobRestoreRange_STATUSARM struct {
	EndRange   *string `json:"endRange,omitempty"`
	StartRange *string `json:"startRange,omitempty"`
}

// Deprecated version of EncryptionService_STATUS. Use v1beta20210401.EncryptionService_STATUS instead
type EncryptionService_STATUSARM struct {
	Enabled         *bool                             `json:"enabled,omitempty"`
	KeyType         *EncryptionService_KeyType_STATUS `json:"keyType,omitempty"`
	LastEnabledTime *string                           `json:"lastEnabledTime,omitempty"`
}
