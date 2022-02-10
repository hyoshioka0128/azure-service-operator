// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Namespaces_SPECARM struct {
	AzureName string `json:"azureName"`

	//Identity: Properties of BYOK Identity description
	Identity *Identity_SpecARM `json:"identity,omitempty"`

	//Location: The Geo-location where the resource lives
	Location string `json:"location"`
	Name     string `json:"name"`

	//Properties: Properties of the namespace.
	Properties *SBNamespaceProperties_SpecARM `json:"properties,omitempty"`

	//Sku: Properties of SKU
	Sku *SBSku_SpecARM `json:"sku,omitempty"`

	//Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Namespaces_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01"
func (specarm Namespaces_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm Namespaces_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm Namespaces_SPECARM) GetType() string {
	return ""
}

type Identity_SpecARM struct {
	//Type: Type of managed service identity.
	Type *IdentitySpecType `json:"type,omitempty"`
}

type SBNamespaceProperties_SpecARM struct {
	//Encryption: Properties of BYOK Encryption description
	Encryption *Encryption_SpecARM `json:"encryption,omitempty"`

	//PrivateEndpointConnections: List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnection_SpecARM `json:"privateEndpointConnections,omitempty"`

	//ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in
	//regions supported availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

type SBSku_SpecARM struct {
	//Capacity: The specified messaging units for the tier. For Premium tier, capacity
	//are 1,2 and 4.
	Capacity *int `json:"capacity,omitempty"`

	//Name: Name of this SKU.
	Name SBSkuSpecName `json:"name"`

	//Tier: The billing tier of this particular SKU.
	Tier *SBSkuSpecTier `json:"tier,omitempty"`
}

type Encryption_SpecARM struct {
	//KeySource: Enumerates the possible value of keySource for Encryption
	KeySource *EncryptionSpecKeySource `json:"keySource,omitempty"`

	//KeyVaultProperties: Properties of KeyVault
	KeyVaultProperties []KeyVaultProperties_SpecARM `json:"keyVaultProperties,omitempty"`

	//RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double
	//Encryption)
	RequireInfrastructureEncryption *bool `json:"requireInfrastructureEncryption,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type IdentitySpecType string

const (
	IdentitySpecTypeNone                       = IdentitySpecType("None")
	IdentitySpecTypeSystemAssigned             = IdentitySpecType("SystemAssigned")
	IdentitySpecTypeSystemAssignedUserAssigned = IdentitySpecType("SystemAssigned, UserAssigned")
	IdentitySpecTypeUserAssigned               = IdentitySpecType("UserAssigned")
)

type PrivateEndpointConnection_SpecARM struct {
	//Properties: Properties of the PrivateEndpointConnection.
	Properties *PrivateEndpointConnectionProperties_SpecARM `json:"properties,omitempty"`
}

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type SBSkuSpecName string

const (
	SBSkuSpecNameBasic    = SBSkuSpecName("Basic")
	SBSkuSpecNamePremium  = SBSkuSpecName("Premium")
	SBSkuSpecNameStandard = SBSkuSpecName("Standard")
)

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type SBSkuSpecTier string

const (
	SBSkuSpecTierBasic    = SBSkuSpecTier("Basic")
	SBSkuSpecTierPremium  = SBSkuSpecTier("Premium")
	SBSkuSpecTierStandard = SBSkuSpecTier("Standard")
)

type KeyVaultProperties_SpecARM struct {
	Identity *UserAssignedIdentityProperties_SpecARM `json:"identity,omitempty"`

	//KeyName: Name of the Key from KeyVault
	KeyName *string `json:"keyName,omitempty"`

	//KeyVaultUri: Uri of KeyVault
	KeyVaultUri *string `json:"keyVaultUri,omitempty"`

	//KeyVersion: Version of KeyVault
	KeyVersion *string `json:"keyVersion,omitempty"`
}

type PrivateEndpointConnectionProperties_SpecARM struct {
	//PrivateEndpoint: The Private Endpoint resource for this Connection.
	PrivateEndpoint *PrivateEndpoint_SpecARM `json:"privateEndpoint,omitempty"`

	//PrivateLinkServiceConnectionState: Details about the state of the connection.
	PrivateLinkServiceConnectionState *ConnectionState_SpecARM `json:"privateLinkServiceConnectionState,omitempty"`

	//ProvisioningState: Provisioning state of the Private Endpoint Connection.
	ProvisioningState *PrivateEndpointConnectionPropertiesSpecProvisioningState `json:"provisioningState,omitempty"`
}

type ConnectionState_SpecARM struct {
	//Description: Description of the connection state.
	Description *string `json:"description,omitempty"`

	//Status: Status of the connection.
	Status *ConnectionStateSpecStatus `json:"status,omitempty"`
}

type PrivateEndpoint_SpecARM struct {
	Id *string `json:"id,omitempty"`
}

type UserAssignedIdentityProperties_SpecARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}