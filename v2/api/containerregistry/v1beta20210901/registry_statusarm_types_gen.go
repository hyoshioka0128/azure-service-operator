// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210901

type Registry_STATUSARM struct {
	// Id: The resource ID.
	Id *string `json:"id,omitempty"`

	// Identity: The identity of the container registry.
	Identity *IdentityProperties_STATUSARM `json:"identity,omitempty"`

	// Location: The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource.
	Name *string `json:"name,omitempty"`

	// Properties: The properties of the container registry.
	Properties *RegistryProperties_STATUSARM `json:"properties,omitempty"`

	// Sku: The SKU of the container registry.
	Sku *Sku_STATUSARM `json:"sku,omitempty"`

	// SystemData: Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData_STATUSARM `json:"systemData,omitempty"`

	// Tags: The tags of the resource.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource.
	Type *string `json:"type,omitempty"`
}

type IdentityProperties_STATUSARM struct {
	// PrincipalId: The principal ID of resource identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant ID of resource.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The identity type.
	Type *IdentityProperties_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: The list of user identities associated with the resource. The user identity
	// dictionary key references will be ARM resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/
	// providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]UserIdentityProperties_STATUSARM `json:"userAssignedIdentities,omitempty"`
}

type RegistryProperties_STATUSARM struct {
	// AdminUserEnabled: The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `json:"adminUserEnabled,omitempty"`

	// CreationDate: The creation date of the container registry in ISO8601 format.
	CreationDate *string `json:"creationDate,omitempty"`

	// DataEndpointEnabled: Enable a single data endpoint per region for serving data.
	DataEndpointEnabled *bool `json:"dataEndpointEnabled,omitempty"`

	// DataEndpointHostNames: List of host names that will serve data when dataEndpointEnabled is true.
	DataEndpointHostNames []string `json:"dataEndpointHostNames,omitempty"`

	// Encryption: The encryption settings of container registry.
	Encryption *EncryptionProperty_STATUSARM `json:"encryption,omitempty"`

	// LoginServer: The URL that can be used to log into the container registry.
	LoginServer *string `json:"loginServer,omitempty"`

	// NetworkRuleBypassOptions: Whether to allow trusted Azure services to access a network restricted registry.
	NetworkRuleBypassOptions *RegistryProperties_NetworkRuleBypassOptions_STATUS `json:"networkRuleBypassOptions,omitempty"`

	// NetworkRuleSet: The network rule set for a container registry.
	NetworkRuleSet *NetworkRuleSet_STATUSARM `json:"networkRuleSet,omitempty"`

	// Policies: The policies for a container registry.
	Policies *Policies_STATUSARM `json:"policies,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections for a container registry.
	PrivateEndpointConnections []PrivateEndpointConnection_STATUSARM `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: The provisioning state of the container registry at the time the operation was called.
	ProvisioningState *RegistryProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for the container registry.
	PublicNetworkAccess *RegistryProperties_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// Status: The status of the container registry at the time the operation was called.
	Status *Status_STATUSARM `json:"status,omitempty"`

	// ZoneRedundancy: Whether or not zone redundancy is enabled for this container registry
	ZoneRedundancy *RegistryProperties_ZoneRedundancy_STATUS `json:"zoneRedundancy,omitempty"`
}

type Sku_STATUSARM struct {
	// Name: The SKU name of the container registry. Required for registry creation.
	Name *Sku_Name_STATUS `json:"name,omitempty"`

	// Tier: The SKU tier based on the SKU name.
	Tier *Sku_Tier_STATUS `json:"tier,omitempty"`
}

type SystemData_STATUSARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource modification (UTC).
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

type EncryptionProperty_STATUSARM struct {
	// KeyVaultProperties: Key vault properties.
	KeyVaultProperties *KeyVaultProperties_STATUSARM `json:"keyVaultProperties,omitempty"`

	// Status: Indicates whether or not the encryption is enabled for container registry.
	Status *EncryptionProperty_Status_STATUS `json:"status,omitempty"`
}

type IdentityProperties_Type_STATUS string

const (
	IdentityProperties_Type_STATUSNone                       = IdentityProperties_Type_STATUS("None")
	IdentityProperties_Type_STATUSSystemAssigned             = IdentityProperties_Type_STATUS("SystemAssigned")
	IdentityProperties_Type_STATUSSystemAssignedUserAssigned = IdentityProperties_Type_STATUS("SystemAssigned, UserAssigned")
	IdentityProperties_Type_STATUSUserAssigned               = IdentityProperties_Type_STATUS("UserAssigned")
)

type NetworkRuleSet_STATUSARM struct {
	// DefaultAction: The default action of allow or deny when no other rules match.
	DefaultAction *NetworkRuleSet_DefaultAction_STATUS `json:"defaultAction,omitempty"`

	// IpRules: The IP ACL rules.
	IpRules []IPRule_STATUSARM `json:"ipRules,omitempty"`
}

type Policies_STATUSARM struct {
	// ExportPolicy: The export policy for a container registry.
	ExportPolicy *ExportPolicy_STATUSARM `json:"exportPolicy,omitempty"`

	// QuarantinePolicy: The quarantine policy for a container registry.
	QuarantinePolicy *QuarantinePolicy_STATUSARM `json:"quarantinePolicy,omitempty"`

	// RetentionPolicy: The retention policy for a container registry.
	RetentionPolicy *RetentionPolicy_STATUSARM `json:"retentionPolicy,omitempty"`

	// TrustPolicy: The content trust policy for a container registry.
	TrustPolicy *TrustPolicy_STATUSARM `json:"trustPolicy,omitempty"`
}

type PrivateEndpointConnection_STATUSARM struct {
	// Id: The resource ID.
	Id *string `json:"id,omitempty"`

	// SystemData: Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData_STATUSARM `json:"systemData,omitempty"`
}

type Sku_Name_STATUS string

const (
	Sku_Name_STATUSBasic    = Sku_Name_STATUS("Basic")
	Sku_Name_STATUSClassic  = Sku_Name_STATUS("Classic")
	Sku_Name_STATUSPremium  = Sku_Name_STATUS("Premium")
	Sku_Name_STATUSStandard = Sku_Name_STATUS("Standard")
)

type Sku_Tier_STATUS string

const (
	Sku_Tier_STATUSBasic    = Sku_Tier_STATUS("Basic")
	Sku_Tier_STATUSClassic  = Sku_Tier_STATUS("Classic")
	Sku_Tier_STATUSPremium  = Sku_Tier_STATUS("Premium")
	Sku_Tier_STATUSStandard = Sku_Tier_STATUS("Standard")
)

type Status_STATUSARM struct {
	// DisplayStatus: The short label for the status.
	DisplayStatus *string `json:"displayStatus,omitempty"`

	// Message: The detailed message for the status, including alerts and error messages.
	Message *string `json:"message,omitempty"`

	// Timestamp: The timestamp when the status was changed to the current value.
	Timestamp *string `json:"timestamp,omitempty"`
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUSApplication     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUSKey             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUSManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUSUser            = SystemData_CreatedByType_STATUS("User")
)

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUSApplication     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUSKey             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUSManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUSUser            = SystemData_LastModifiedByType_STATUS("User")
)

type UserIdentityProperties_STATUSARM struct {
	// ClientId: The client id of user assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal id of user assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

type ExportPolicy_STATUSARM struct {
	// Status: The value that indicates whether the policy is enabled or not.
	Status *ExportPolicy_Status_STATUS `json:"status,omitempty"`
}

type IPRule_STATUSARM struct {
	// Action: The action of IP ACL rule.
	Action *IPRule_Action_STATUS `json:"action,omitempty"`

	// Value: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
	Value *string `json:"value,omitempty"`
}

type KeyVaultProperties_STATUSARM struct {
	// Identity: The client id of the identity which will be used to access key vault.
	Identity *string `json:"identity,omitempty"`

	// KeyIdentifier: Key vault uri to access the encryption key.
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`

	// KeyRotationEnabled: Auto key rotation status for a CMK enabled registry.
	KeyRotationEnabled *bool `json:"keyRotationEnabled,omitempty"`

	// LastKeyRotationTimestamp: Timestamp of the last successful key rotation.
	LastKeyRotationTimestamp *string `json:"lastKeyRotationTimestamp,omitempty"`

	// VersionedKeyIdentifier: The fully qualified key identifier that includes the version of the key that is actually used
	// for encryption.
	VersionedKeyIdentifier *string `json:"versionedKeyIdentifier,omitempty"`
}

type QuarantinePolicy_STATUSARM struct {
	// Status: The value that indicates whether the policy is enabled or not.
	Status *QuarantinePolicy_Status_STATUS `json:"status,omitempty"`
}

type RetentionPolicy_STATUSARM struct {
	// Days: The number of days to retain an untagged manifest after which it gets purged.
	Days *int `json:"days,omitempty"`

	// LastUpdatedTime: The timestamp when the policy was last updated.
	LastUpdatedTime *string `json:"lastUpdatedTime,omitempty"`

	// Status: The value that indicates whether the policy is enabled or not.
	Status *RetentionPolicy_Status_STATUS `json:"status,omitempty"`
}

type TrustPolicy_STATUSARM struct {
	// Status: The value that indicates whether the policy is enabled or not.
	Status *TrustPolicy_Status_STATUS `json:"status,omitempty"`

	// Type: The type of trust policy.
	Type *TrustPolicy_Type_STATUS `json:"type,omitempty"`
}
