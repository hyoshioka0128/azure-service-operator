// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type BatchAccounts_SPECARM struct {
	AzureName string `json:"azureName"`

	//Identity: The identity of the Batch account.
	Identity *BatchAccountIdentity_SpecARM `json:"identity,omitempty"`

	//Location: The region in which to create the account.
	Location string `json:"location"`
	Name     string `json:"name"`

	//Properties: The properties of the Batch account.
	Properties *BatchAccountProperties_SpecARM `json:"properties,omitempty"`

	//Tags: The user-specified tags associated with the account.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &BatchAccounts_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01"
func (specarm BatchAccounts_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm BatchAccounts_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Batch/batchAccounts"
func (specarm BatchAccounts_SPECARM) GetType() string {
	return "Microsoft.Batch/batchAccounts"
}

type BatchAccountIdentity_SpecARM struct {
	//Type: The type of identity used for the Batch account.
	Type BatchAccountIdentity_Type_Spec `json:"type"`
}

type BatchAccountProperties_SpecARM struct {
	//AutoStorage: The properties related to the auto-storage account.
	AutoStorage *AutoStorageBaseProperties_SpecARM `json:"autoStorage,omitempty"`

	//Encryption: Configures how customer data is encrypted inside the Batch account.
	//By default, accounts are encrypted using a Microsoft managed key. For additional
	//control, a customer-managed key can be used instead.
	Encryption *EncryptionProperties_SpecARM `json:"encryption,omitempty"`

	//KeyVaultReference: A reference to the Azure key vault associated with the Batch
	//account.
	KeyVaultReference *KeyVaultReference_SpecARM `json:"keyVaultReference,omitempty"`

	//PoolAllocationMode: The pool allocation mode also affects how clients may
	//authenticate to the Batch Service API. If the mode is BatchService, clients may
	//authenticate using access keys or Azure Active Directory. If the mode is
	//UserSubscription, clients must use Azure Active Directory. The default is
	//BatchService.
	PoolAllocationMode *PoolAllocationMode_Spec `json:"poolAllocationMode,omitempty"`

	//PublicNetworkAccess: If not specified, the default value is 'enabled'.
	PublicNetworkAccess *PublicNetworkAccessType_Spec `json:"publicNetworkAccess,omitempty"`
}

type AutoStorageBaseProperties_SpecARM struct {
	StorageAccountId string `json:"storageAccountId"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","UserAssigned"}
type BatchAccountIdentity_Type_Spec string

const (
	BatchAccountIdentity_Type_SpecNone           = BatchAccountIdentity_Type_Spec("None")
	BatchAccountIdentity_Type_SpecSystemAssigned = BatchAccountIdentity_Type_Spec("SystemAssigned")
	BatchAccountIdentity_Type_SpecUserAssigned   = BatchAccountIdentity_Type_Spec("UserAssigned")
)

type EncryptionProperties_SpecARM struct {
	//KeySource: Type of the key source.
	KeySource *EncryptionProperties_KeySource_Spec `json:"keySource,omitempty"`

	//KeyVaultProperties: Additional details when using Microsoft.KeyVault
	KeyVaultProperties *KeyVaultProperties_SpecARM `json:"keyVaultProperties,omitempty"`
}

type KeyVaultReference_SpecARM struct {
	Id string `json:"id"`

	//Url: The URL of the Azure key vault associated with the Batch account.
	Url string `json:"url"`
}

type KeyVaultProperties_SpecARM struct {
	//KeyIdentifier: Full path to the versioned secret. Example
	//https://mykeyvault.vault.azure.net/keys/testkey/6e34a81fef704045975661e297a4c053.
	//To be usable the following prerequisites must be met:
	//The Batch Account has a System Assigned identity
	//The account identity has been granted Key/Get, Key/Unwrap and Key/Wrap
	//permissions
	//The KeyVault has soft-delete and purge protection enabled
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}
