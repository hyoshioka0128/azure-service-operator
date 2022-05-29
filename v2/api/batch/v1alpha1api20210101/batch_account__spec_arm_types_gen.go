// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of BatchAccount_Spec. Use v1beta20210101.BatchAccount_Spec instead
type BatchAccount_SpecARM struct {
	AzureName  string                           `json:"azureName,omitempty"`
	Identity   *BatchAccountIdentityARM         `json:"identity,omitempty"`
	Location   *string                          `json:"location,omitempty"`
	Name       string                           `json:"name,omitempty"`
	Properties *BatchAccountCreatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &BatchAccount_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01"
func (account BatchAccount_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (account *BatchAccount_SpecARM) GetName() string {
	return account.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (account *BatchAccount_SpecARM) GetType() string {
	return ""
}

// Deprecated version of BatchAccountCreateProperties. Use v1beta20210101.BatchAccountCreateProperties instead
type BatchAccountCreatePropertiesARM struct {
	AutoStorage         *AutoStorageBasePropertiesARM `json:"autoStorage,omitempty"`
	Encryption          *EncryptionPropertiesARM      `json:"encryption,omitempty"`
	KeyVaultReference   *KeyVaultReferenceARM         `json:"keyVaultReference,omitempty"`
	PoolAllocationMode  *PoolAllocationMode           `json:"poolAllocationMode,omitempty"`
	PublicNetworkAccess *PublicNetworkAccessType      `json:"publicNetworkAccess,omitempty"`
}

// Deprecated version of BatchAccountIdentity. Use v1beta20210101.BatchAccountIdentity instead
type BatchAccountIdentityARM struct {
	PrincipalId *string                    `json:"principalId,omitempty"`
	TenantId    *string                    `json:"tenantId,omitempty"`
	Type        *BatchAccountIdentity_Type `json:"type,omitempty"`
}

// Deprecated version of AutoStorageBaseProperties. Use v1beta20210101.AutoStorageBaseProperties instead
type AutoStorageBasePropertiesARM struct {
	StorageAccountId *string `json:"storageAccountId,omitempty"`
}

// Deprecated version of BatchAccountIdentity_Type. Use v1beta20210101.BatchAccountIdentity_Type instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","UserAssigned"}
type BatchAccountIdentity_Type string

const (
	BatchAccountIdentity_Type_None           = BatchAccountIdentity_Type("None")
	BatchAccountIdentity_Type_SystemAssigned = BatchAccountIdentity_Type("SystemAssigned")
	BatchAccountIdentity_Type_UserAssigned   = BatchAccountIdentity_Type("UserAssigned")
)

// Deprecated version of EncryptionProperties. Use v1beta20210101.EncryptionProperties instead
type EncryptionPropertiesARM struct {
	KeySource          *EncryptionProperties_KeySource `json:"keySource,omitempty"`
	KeyVaultProperties *KeyVaultPropertiesARM          `json:"keyVaultProperties,omitempty"`
}

// Deprecated version of KeyVaultReference. Use v1beta20210101.KeyVaultReference instead
type KeyVaultReferenceARM struct {
	Id  *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
}

// Deprecated version of KeyVaultProperties. Use v1beta20210101.KeyVaultProperties instead
type KeyVaultPropertiesARM struct {
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}
