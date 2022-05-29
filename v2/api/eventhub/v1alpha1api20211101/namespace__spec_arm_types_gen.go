// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Namespace_Spec. Use v1beta20211101.Namespace_Spec instead
type Namespace_SpecARM struct {
	AzureName  string                        `json:"azureName,omitempty"`
	Id         *string                       `json:"id,omitempty"`
	Identity   *IdentityARM                  `json:"identity,omitempty"`
	Location   *string                       `json:"location,omitempty"`
	Name       string                        `json:"name,omitempty"`
	Properties *Namespace_Spec_PropertiesARM `json:"properties,omitempty"`
	Sku        *SkuARM                       `json:"sku,omitempty"`
	SystemData *SystemDataARM                `json:"systemData,omitempty"`
	Tags       map[string]string             `json:"tags,omitempty"`
	Type       *string                       `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Namespace_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (namespace Namespace_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (namespace *Namespace_SpecARM) GetName() string {
	return namespace.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (namespace *Namespace_SpecARM) GetType() string {
	return ""
}

// Deprecated version of Identity. Use v1beta20211101.Identity instead
type IdentityARM struct {
	PrincipalId *string        `json:"principalId,omitempty"`
	TenantId    *string        `json:"tenantId,omitempty"`
	Type        *Identity_Type `json:"type,omitempty"`
}

// Deprecated version of Namespace_Spec_Properties. Use v1beta20211101.Namespace_Spec_Properties instead
type Namespace_Spec_PropertiesARM struct {
	AlternateName              *string                        `json:"alternateName,omitempty"`
	ClusterArmId               *string                        `json:"clusterArmId,omitempty"`
	CreatedAt                  *string                        `json:"createdAt,omitempty"`
	DisableLocalAuth           *bool                          `json:"disableLocalAuth,omitempty"`
	Encryption                 *EncryptionARM                 `json:"encryption,omitempty"`
	IsAutoInflateEnabled       *bool                          `json:"isAutoInflateEnabled,omitempty"`
	KafkaEnabled               *bool                          `json:"kafkaEnabled,omitempty"`
	MaximumThroughputUnits     *int                           `json:"maximumThroughputUnits,omitempty"`
	MetricId                   *string                        `json:"metricId,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnectionARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *string                        `json:"provisioningState,omitempty"`
	ServiceBusEndpoint         *string                        `json:"serviceBusEndpoint,omitempty"`
	Status                     *string                        `json:"status,omitempty"`
	UpdatedAt                  *string                        `json:"updatedAt,omitempty"`
	ZoneRedundant              *bool                          `json:"zoneRedundant,omitempty"`
}

// Deprecated version of Sku. Use v1beta20211101.Sku instead
type SkuARM struct {
	Capacity *int      `json:"capacity,omitempty"`
	Name     *Sku_Name `json:"name,omitempty"`
	Tier     *Sku_Tier `json:"tier,omitempty"`
}

// Deprecated version of SystemData. Use v1beta20211101.SystemData instead
type SystemDataARM struct {
	CreatedAt          *string                        `json:"createdAt,omitempty"`
	CreatedBy          *string                        `json:"createdBy,omitempty"`
	CreatedByType      *SystemData_CreatedByType      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                        `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                        `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemData_LastModifiedByType `json:"lastModifiedByType,omitempty"`
}

// Deprecated version of Encryption. Use v1beta20211101.Encryption instead
type EncryptionARM struct {
	KeySource                       *Encryption_KeySource   `json:"keySource,omitempty"`
	KeyVaultProperties              []KeyVaultPropertiesARM `json:"keyVaultProperties,omitempty"`
	RequireInfrastructureEncryption *bool                   `json:"requireInfrastructureEncryption,omitempty"`
}

// Deprecated version of Identity_Type. Use v1beta20211101.Identity_Type instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type Identity_Type string

const (
	Identity_Type_None                       = Identity_Type("None")
	Identity_Type_SystemAssigned             = Identity_Type("SystemAssigned")
	Identity_Type_SystemAssignedUserAssigned = Identity_Type("SystemAssigned, UserAssigned")
	Identity_Type_UserAssigned               = Identity_Type("UserAssigned")
)

// Deprecated version of PrivateEndpointConnection. Use v1beta20211101.PrivateEndpointConnection instead
type PrivateEndpointConnectionARM struct {
	Id         *string        `json:"id,omitempty"`
	SystemData *SystemDataARM `json:"systemData,omitempty"`
}

// Deprecated version of Sku_Name. Use v1beta20211101.Sku_Name instead
// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type Sku_Name string

const (
	Sku_Name_Basic    = Sku_Name("Basic")
	Sku_Name_Premium  = Sku_Name("Premium")
	Sku_Name_Standard = Sku_Name("Standard")
)

// Deprecated version of Sku_Tier. Use v1beta20211101.Sku_Tier instead
// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type Sku_Tier string

const (
	Sku_Tier_Basic    = Sku_Tier("Basic")
	Sku_Tier_Premium  = Sku_Tier("Premium")
	Sku_Tier_Standard = Sku_Tier("Standard")
)

// Deprecated version of SystemData_CreatedByType. Use v1beta20211101.SystemData_CreatedByType instead
// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_CreatedByType string

const (
	SystemData_CreatedByType_Application     = SystemData_CreatedByType("Application")
	SystemData_CreatedByType_Key             = SystemData_CreatedByType("Key")
	SystemData_CreatedByType_ManagedIdentity = SystemData_CreatedByType("ManagedIdentity")
	SystemData_CreatedByType_User            = SystemData_CreatedByType("User")
)

// Deprecated version of SystemData_LastModifiedByType. Use v1beta20211101.SystemData_LastModifiedByType instead
// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_LastModifiedByType string

const (
	SystemData_LastModifiedByType_Application     = SystemData_LastModifiedByType("Application")
	SystemData_LastModifiedByType_Key             = SystemData_LastModifiedByType("Key")
	SystemData_LastModifiedByType_ManagedIdentity = SystemData_LastModifiedByType("ManagedIdentity")
	SystemData_LastModifiedByType_User            = SystemData_LastModifiedByType("User")
)

// Deprecated version of KeyVaultProperties. Use v1beta20211101.KeyVaultProperties instead
type KeyVaultPropertiesARM struct {
	Identity    *UserAssignedIdentityPropertiesARM `json:"identity,omitempty"`
	KeyName     *string                            `json:"keyName,omitempty"`
	KeyVaultUri *string                            `json:"keyVaultUri,omitempty"`
	KeyVersion  *string                            `json:"keyVersion,omitempty"`
}

// Deprecated version of UserAssignedIdentityProperties. Use v1beta20211101.UserAssignedIdentityProperties instead
type UserAssignedIdentityPropertiesARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}
