// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210901

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Registry_Spec. Use v1beta20210901.Registry_Spec instead
type Registry_SpecARM struct {
	AzureName  string                 `json:"azureName,omitempty"`
	Identity   *IdentityPropertiesARM `json:"identity,omitempty"`
	Location   *string                `json:"location,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Properties *RegistryPropertiesARM `json:"properties,omitempty"`
	Sku        *SkuARM                `json:"sku,omitempty"`
	Tags       map[string]string      `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Registry_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-09-01"
func (registry Registry_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (registry *Registry_SpecARM) GetName() string {
	return registry.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerRegistry/registries"
func (registry *Registry_SpecARM) GetType() string {
	return "Microsoft.ContainerRegistry/registries"
}

// Deprecated version of IdentityProperties. Use v1beta20210901.IdentityProperties instead
type IdentityPropertiesARM struct {
	PrincipalId            *string                              `json:"principalId,omitempty"`
	TenantId               *string                              `json:"tenantId,omitempty"`
	Type                   *IdentityProperties_Type             `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserIdentityPropertiesARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of RegistryProperties. Use v1beta20210901.RegistryProperties instead
type RegistryPropertiesARM struct {
	AdminUserEnabled         *bool                                        `json:"adminUserEnabled,omitempty"`
	DataEndpointEnabled      *bool                                        `json:"dataEndpointEnabled,omitempty"`
	Encryption               *EncryptionPropertyARM                       `json:"encryption,omitempty"`
	NetworkRuleBypassOptions *RegistryProperties_NetworkRuleBypassOptions `json:"networkRuleBypassOptions,omitempty"`
	NetworkRuleSet           *NetworkRuleSetARM                           `json:"networkRuleSet,omitempty"`
	Policies                 *PoliciesARM                                 `json:"policies,omitempty"`
	PublicNetworkAccess      *RegistryProperties_PublicNetworkAccess      `json:"publicNetworkAccess,omitempty"`
	ZoneRedundancy           *RegistryProperties_ZoneRedundancy           `json:"zoneRedundancy,omitempty"`
}

// Deprecated version of Sku. Use v1beta20210901.Sku instead
type SkuARM struct {
	Name *Sku_Name `json:"name,omitempty"`
}

// Deprecated version of EncryptionProperty. Use v1beta20210901.EncryptionProperty instead
type EncryptionPropertyARM struct {
	KeyVaultProperties *KeyVaultPropertiesARM     `json:"keyVaultProperties,omitempty"`
	Status             *EncryptionProperty_Status `json:"status,omitempty"`
}

// Deprecated version of IdentityProperties_Type. Use v1beta20210901.IdentityProperties_Type instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type IdentityProperties_Type string

const (
	IdentityProperties_Type_None                       = IdentityProperties_Type("None")
	IdentityProperties_Type_SystemAssigned             = IdentityProperties_Type("SystemAssigned")
	IdentityProperties_Type_SystemAssignedUserAssigned = IdentityProperties_Type("SystemAssigned, UserAssigned")
	IdentityProperties_Type_UserAssigned               = IdentityProperties_Type("UserAssigned")
)

// Deprecated version of NetworkRuleSet. Use v1beta20210901.NetworkRuleSet instead
type NetworkRuleSetARM struct {
	DefaultAction *NetworkRuleSet_DefaultAction `json:"defaultAction,omitempty"`
	IpRules       []IPRuleARM                   `json:"ipRules,omitempty"`
}

// Deprecated version of Policies. Use v1beta20210901.Policies instead
type PoliciesARM struct {
	ExportPolicy     *ExportPolicyARM     `json:"exportPolicy,omitempty"`
	QuarantinePolicy *QuarantinePolicyARM `json:"quarantinePolicy,omitempty"`
	RetentionPolicy  *RetentionPolicyARM  `json:"retentionPolicy,omitempty"`
	TrustPolicy      *TrustPolicyARM      `json:"trustPolicy,omitempty"`
}

// Deprecated version of Sku_Name. Use v1beta20210901.Sku_Name instead
// +kubebuilder:validation:Enum={"Basic","Classic","Premium","Standard"}
type Sku_Name string

const (
	Sku_Name_Basic    = Sku_Name("Basic")
	Sku_Name_Classic  = Sku_Name("Classic")
	Sku_Name_Premium  = Sku_Name("Premium")
	Sku_Name_Standard = Sku_Name("Standard")
)

// Deprecated version of UserIdentityProperties. Use v1beta20210901.UserIdentityProperties instead
type UserIdentityPropertiesARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

// Deprecated version of ExportPolicy. Use v1beta20210901.ExportPolicy instead
type ExportPolicyARM struct {
	Status *ExportPolicy_Status `json:"status,omitempty"`
}

// Deprecated version of IPRule. Use v1beta20210901.IPRule instead
type IPRuleARM struct {
	Action *IPRule_Action `json:"action,omitempty"`
	Value  *string        `json:"value,omitempty"`
}

// Deprecated version of KeyVaultProperties. Use v1beta20210901.KeyVaultProperties instead
type KeyVaultPropertiesARM struct {
	Identity      *string `json:"identity,omitempty"`
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}

// Deprecated version of QuarantinePolicy. Use v1beta20210901.QuarantinePolicy instead
type QuarantinePolicyARM struct {
	Status *QuarantinePolicy_Status `json:"status,omitempty"`
}

// Deprecated version of RetentionPolicy. Use v1beta20210901.RetentionPolicy instead
type RetentionPolicyARM struct {
	Days   *int                    `json:"days,omitempty"`
	Status *RetentionPolicy_Status `json:"status,omitempty"`
}

// Deprecated version of TrustPolicy. Use v1beta20210901.TrustPolicy instead
type TrustPolicyARM struct {
	Status *TrustPolicy_Status `json:"status,omitempty"`
	Type   *TrustPolicy_Type   `json:"type,omitempty"`
}