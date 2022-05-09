// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101

// Deprecated version of BatchAccount_Status. Use v1beta20210101.BatchAccount_Status instead
type BatchAccount_StatusARM struct {
	Id         *string                           `json:"id,omitempty"`
	Identity   *BatchAccountIdentity_StatusARM   `json:"identity,omitempty"`
	Location   *string                           `json:"location,omitempty"`
	Name       *string                           `json:"name,omitempty"`
	Properties *BatchAccountProperties_StatusARM `json:"properties,omitempty"`
	Tags       map[string]string                 `json:"tags,omitempty"`
	Type       *string                           `json:"type,omitempty"`
}

// Deprecated version of BatchAccountIdentity_Status. Use v1beta20210101.BatchAccountIdentity_Status instead
type BatchAccountIdentity_StatusARM struct {
	PrincipalId            *string                                                          `json:"principalId,omitempty"`
	TenantId               *string                                                          `json:"tenantId,omitempty"`
	Type                   *BatchAccountIdentityStatusType                                  `json:"type,omitempty"`
	UserAssignedIdentities map[string]BatchAccountIdentity_Status_UserAssignedIdentitiesARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of BatchAccountProperties_Status. Use v1beta20210101.BatchAccountProperties_Status instead
type BatchAccountProperties_StatusARM struct {
	AccountEndpoint                       *string                                        `json:"accountEndpoint,omitempty"`
	ActiveJobAndJobScheduleQuota          *int                                           `json:"activeJobAndJobScheduleQuota,omitempty"`
	AutoStorage                           *AutoStorageProperties_StatusARM               `json:"autoStorage,omitempty"`
	DedicatedCoreQuota                    *int                                           `json:"dedicatedCoreQuota,omitempty"`
	DedicatedCoreQuotaPerVMFamily         []VirtualMachineFamilyCoreQuota_StatusARM      `json:"dedicatedCoreQuotaPerVMFamily,omitempty"`
	DedicatedCoreQuotaPerVMFamilyEnforced *bool                                          `json:"dedicatedCoreQuotaPerVMFamilyEnforced,omitempty"`
	Encryption                            *EncryptionProperties_StatusARM                `json:"encryption,omitempty"`
	KeyVaultReference                     *KeyVaultReference_StatusARM                   `json:"keyVaultReference,omitempty"`
	LowPriorityCoreQuota                  *int                                           `json:"lowPriorityCoreQuota,omitempty"`
	PoolAllocationMode                    *PoolAllocationMode_Status                     `json:"poolAllocationMode,omitempty"`
	PoolQuota                             *int                                           `json:"poolQuota,omitempty"`
	PrivateEndpointConnections            []PrivateEndpointConnection_StatusARM          `json:"privateEndpointConnections,omitempty"`
	ProvisioningState                     *BatchAccountPropertiesStatusProvisioningState `json:"provisioningState,omitempty"`
	PublicNetworkAccess                   *PublicNetworkAccessType_Status                `json:"publicNetworkAccess,omitempty"`
}

// Deprecated version of AutoStorageProperties_Status. Use v1beta20210101.AutoStorageProperties_Status instead
type AutoStorageProperties_StatusARM struct {
	LastKeySync      *string `json:"lastKeySync,omitempty"`
	StorageAccountId *string `json:"storageAccountId,omitempty"`
}

// Deprecated version of BatchAccountIdentityStatusType. Use v1beta20210101.BatchAccountIdentityStatusType instead
type BatchAccountIdentityStatusType string

const (
	BatchAccountIdentityStatusTypeNone           = BatchAccountIdentityStatusType("None")
	BatchAccountIdentityStatusTypeSystemAssigned = BatchAccountIdentityStatusType("SystemAssigned")
	BatchAccountIdentityStatusTypeUserAssigned   = BatchAccountIdentityStatusType("UserAssigned")
)

// Deprecated version of BatchAccountIdentity_Status_UserAssignedIdentities. Use v1beta20210101.BatchAccountIdentity_Status_UserAssignedIdentities instead
type BatchAccountIdentity_Status_UserAssignedIdentitiesARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

// Deprecated version of EncryptionProperties_Status. Use v1beta20210101.EncryptionProperties_Status instead
type EncryptionProperties_StatusARM struct {
	KeySource          *EncryptionPropertiesStatusKeySource `json:"keySource,omitempty"`
	KeyVaultProperties *KeyVaultProperties_StatusARM        `json:"keyVaultProperties,omitempty"`
}

// Deprecated version of KeyVaultReference_Status. Use v1beta20210101.KeyVaultReference_Status instead
type KeyVaultReference_StatusARM struct {
	Id  *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
}

// Deprecated version of PrivateEndpointConnection_Status. Use v1beta20210101.PrivateEndpointConnection_Status instead
type PrivateEndpointConnection_StatusARM struct {
	Etag       *string                                        `json:"etag,omitempty"`
	Id         *string                                        `json:"id,omitempty"`
	Name       *string                                        `json:"name,omitempty"`
	Properties *PrivateEndpointConnectionProperties_StatusARM `json:"properties,omitempty"`
	Type       *string                                        `json:"type,omitempty"`
}

// Deprecated version of VirtualMachineFamilyCoreQuota_Status. Use v1beta20210101.VirtualMachineFamilyCoreQuota_Status instead
type VirtualMachineFamilyCoreQuota_StatusARM struct {
	CoreQuota *int    `json:"coreQuota,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// Deprecated version of KeyVaultProperties_Status. Use v1beta20210101.KeyVaultProperties_Status instead
type KeyVaultProperties_StatusARM struct {
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}

// Deprecated version of PrivateEndpointConnectionProperties_Status. Use v1beta20210101.PrivateEndpointConnectionProperties_Status instead
type PrivateEndpointConnectionProperties_StatusARM struct {
	PrivateEndpoint                   *PrivateEndpoint_StatusARM                                  `json:"privateEndpoint,omitempty"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState_StatusARM                `json:"privateLinkServiceConnectionState,omitempty"`
	ProvisioningState                 *PrivateEndpointConnectionPropertiesStatusProvisioningState `json:"provisioningState,omitempty"`
}

// Deprecated version of PrivateEndpoint_Status. Use v1beta20210101.PrivateEndpoint_Status instead
type PrivateEndpoint_StatusARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of PrivateLinkServiceConnectionState_Status. Use v1beta20210101.PrivateLinkServiceConnectionState_Status instead
type PrivateLinkServiceConnectionState_StatusARM struct {
	ActionRequired *string                                    `json:"actionRequired,omitempty"`
	Description    *string                                    `json:"description,omitempty"`
	Status         *PrivateLinkServiceConnectionStatus_Status `json:"status,omitempty"`
}
