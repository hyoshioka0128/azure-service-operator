// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of DatabaseAccount_Spec. Use v1beta20210515.DatabaseAccount_Spec instead
type DatabaseAccount_SpecARM struct {
	AzureName  string                                    `json:"azureName,omitempty"`
	Identity   *ManagedServiceIdentityARM                `json:"identity,omitempty"`
	Kind       *DatabaseAccount_Spec_Kind                `json:"kind,omitempty"`
	Location   *string                                   `json:"location,omitempty"`
	Name       string                                    `json:"name,omitempty"`
	Properties *DatabaseAccountCreateUpdatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                         `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccount_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (account DatabaseAccount_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (account *DatabaseAccount_SpecARM) GetName() string {
	return account.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (account *DatabaseAccount_SpecARM) GetType() string {
	return ""
}

// Deprecated version of DatabaseAccountCreateUpdateProperties. Use v1beta20210515.DatabaseAccountCreateUpdateProperties instead
type DatabaseAccountCreateUpdatePropertiesARM struct {
	AnalyticalStorageConfiguration     *AnalyticalStorageConfigurationARM `json:"analyticalStorageConfiguration,omitempty"`
	ApiProperties                      *ApiPropertiesARM                  `json:"apiProperties,omitempty"`
	BackupPolicy                       *BackupPolicyARM                   `json:"backupPolicy,omitempty"`
	Capabilities                       []CapabilityARM                    `json:"capabilities,omitempty"`
	ConnectorOffer                     *ConnectorOffer                    `json:"connectorOffer,omitempty"`
	ConsistencyPolicy                  *ConsistencyPolicyARM              `json:"consistencyPolicy,omitempty"`
	Cors                               []CorsPolicyARM                    `json:"cors,omitempty"`
	DatabaseAccountOfferType           *DatabaseAccountOfferType          `json:"databaseAccountOfferType,omitempty"`
	DefaultIdentity                    *string                            `json:"defaultIdentity,omitempty"`
	DisableKeyBasedMetadataWriteAccess *bool                              `json:"disableKeyBasedMetadataWriteAccess,omitempty"`
	EnableAnalyticalStorage            *bool                              `json:"enableAnalyticalStorage,omitempty"`
	EnableAutomaticFailover            *bool                              `json:"enableAutomaticFailover,omitempty"`
	EnableCassandraConnector           *bool                              `json:"enableCassandraConnector,omitempty"`
	EnableFreeTier                     *bool                              `json:"enableFreeTier,omitempty"`
	EnableMultipleWriteLocations       *bool                              `json:"enableMultipleWriteLocations,omitempty"`
	IpRules                            []IpAddressOrRangeARM              `json:"ipRules,omitempty"`
	IsVirtualNetworkFilterEnabled      *bool                              `json:"isVirtualNetworkFilterEnabled,omitempty"`
	KeyVaultKeyUri                     *string                            `json:"keyVaultKeyUri,omitempty"`
	Locations                          []LocationARM                      `json:"locations,omitempty"`
	NetworkAclBypass                   *NetworkAclBypass                  `json:"networkAclBypass,omitempty"`
	NetworkAclBypassResourceIds        []string                           `json:"networkAclBypassResourceIds,omitempty"`
	PublicNetworkAccess                *PublicNetworkAccess               `json:"publicNetworkAccess,omitempty"`
	VirtualNetworkRules                []VirtualNetworkRuleARM            `json:"virtualNetworkRules,omitempty"`
}

// Deprecated version of DatabaseAccount_Spec_Kind. Use v1beta20210515.DatabaseAccount_Spec_Kind instead
// +kubebuilder:validation:Enum={"GlobalDocumentDB","MongoDB","Parse"}
type DatabaseAccount_Spec_Kind string

const (
	DatabaseAccount_Spec_Kind_GlobalDocumentDB = DatabaseAccount_Spec_Kind("GlobalDocumentDB")
	DatabaseAccount_Spec_Kind_MongoDB          = DatabaseAccount_Spec_Kind("MongoDB")
	DatabaseAccount_Spec_Kind_Parse            = DatabaseAccount_Spec_Kind("Parse")
)

// Deprecated version of ManagedServiceIdentity. Use v1beta20210515.ManagedServiceIdentity instead
type ManagedServiceIdentityARM struct {
	Type *ManagedServiceIdentity_Type `json:"type,omitempty"`
}

// Deprecated version of AnalyticalStorageConfiguration. Use v1beta20210515.AnalyticalStorageConfiguration instead
type AnalyticalStorageConfigurationARM struct {
	SchemaType *AnalyticalStorageSchemaType `json:"schemaType,omitempty"`
}

// Deprecated version of ApiProperties. Use v1beta20210515.ApiProperties instead
type ApiPropertiesARM struct {
	ServerVersion *ApiProperties_ServerVersion `json:"serverVersion,omitempty"`
}

// Deprecated version of BackupPolicy. Use v1beta20210515.BackupPolicy instead
type BackupPolicyARM struct {
	Type *BackupPolicyType `json:"type,omitempty"`
}

// Deprecated version of Capability. Use v1beta20210515.Capability instead
type CapabilityARM struct {
	Name *string `json:"name,omitempty"`
}

// Deprecated version of ConsistencyPolicy. Use v1beta20210515.ConsistencyPolicy instead
type ConsistencyPolicyARM struct {
	DefaultConsistencyLevel *ConsistencyPolicy_DefaultConsistencyLevel `json:"defaultConsistencyLevel,omitempty"`
	MaxIntervalInSeconds    *int                                       `json:"maxIntervalInSeconds,omitempty"`
	MaxStalenessPrefix      *int                                       `json:"maxStalenessPrefix,omitempty"`
}

// Deprecated version of CorsPolicy. Use v1beta20210515.CorsPolicy instead
type CorsPolicyARM struct {
	AllowedHeaders  *string `json:"allowedHeaders,omitempty"`
	AllowedMethods  *string `json:"allowedMethods,omitempty"`
	AllowedOrigins  *string `json:"allowedOrigins,omitempty"`
	ExposedHeaders  *string `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int    `json:"maxAgeInSeconds,omitempty"`
}

// Deprecated version of IpAddressOrRange. Use v1beta20210515.IpAddressOrRange instead
type IpAddressOrRangeARM struct {
	IpAddressOrRange *string `json:"ipAddressOrRange,omitempty"`
}

// Deprecated version of Location. Use v1beta20210515.Location instead
type LocationARM struct {
	FailoverPriority *int    `json:"failoverPriority,omitempty"`
	IsZoneRedundant  *bool   `json:"isZoneRedundant,omitempty"`
	LocationName     *string `json:"locationName,omitempty"`
}

// Deprecated version of ManagedServiceIdentity_Type. Use v1beta20210515.ManagedServiceIdentity_Type instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type ManagedServiceIdentity_Type string

const (
	ManagedServiceIdentity_Type_None                       = ManagedServiceIdentity_Type("None")
	ManagedServiceIdentity_Type_SystemAssigned             = ManagedServiceIdentity_Type("SystemAssigned")
	ManagedServiceIdentity_Type_SystemAssignedUserAssigned = ManagedServiceIdentity_Type("SystemAssigned,UserAssigned")
	ManagedServiceIdentity_Type_UserAssigned               = ManagedServiceIdentity_Type("UserAssigned")
)

// Deprecated version of VirtualNetworkRule. Use v1beta20210515.VirtualNetworkRule instead
type VirtualNetworkRuleARM struct {
	Id                               *string `json:"id,omitempty"`
	IgnoreMissingVNetServiceEndpoint *bool   `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
}
