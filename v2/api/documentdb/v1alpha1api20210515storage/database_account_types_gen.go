// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=databaseaccounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={databaseaccounts/status,databaseaccounts/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210515.DatabaseAccount
//Generator information:
//- Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}
type DatabaseAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_SPEC  `json:"spec,omitempty"`
	Status            DatabaseAccount_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DatabaseAccount{}

// GetConditions returns the conditions of the resource
func (account *DatabaseAccount) GetConditions() conditions.Conditions {
	return account.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (account *DatabaseAccount) SetConditions(conditions conditions.Conditions) {
	account.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &DatabaseAccount{}

// AzureName returns the Azure name of the resource
func (account *DatabaseAccount) AzureName() string {
	return account.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (account DatabaseAccount) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (account *DatabaseAccount) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (account *DatabaseAccount) GetSpec() genruntime.ConvertibleSpec {
	return &account.Spec
}

// GetStatus returns the status of this resource
func (account *DatabaseAccount) GetStatus() genruntime.ConvertibleStatus {
	return &account.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts"
func (account *DatabaseAccount) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts"
}

// NewEmptyStatus returns a new empty (blank) status
func (account *DatabaseAccount) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccount_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (account *DatabaseAccount) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(account.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  account.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (account *DatabaseAccount) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccount_Status); ok {
		account.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccount_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	account.Status = st
	return nil
}

// Hub marks that this DatabaseAccount is the hub type for conversion
func (account *DatabaseAccount) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (account *DatabaseAccount) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: account.Spec.OriginalVersion,
		Kind:    "DatabaseAccount",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.DatabaseAccount
//Generator information:
//- Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}
type DatabaseAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseAccount `json:"items"`
}

// +kubebuilder:validation:Enum={"2021-05-15"}
type APIVersion string

const APIVersionValue = APIVersion("2021-05-15")

//Storage version of v1alpha1api20210515.DatabaseAccount_Status
type DatabaseAccount_Status struct {
	AnalyticalStorageConfiguration     *AnalyticalStorageConfiguration_Status `json:"analyticalStorageConfiguration,omitempty"`
	ApiProperties                      *ApiProperties_Status                  `json:"apiProperties,omitempty"`
	BackupPolicy                       *BackupPolicy_Status                   `json:"backupPolicy,omitempty"`
	Capabilities                       []Capability_Status                    `json:"capabilities,omitempty"`
	Conditions                         []conditions.Condition                 `json:"conditions,omitempty"`
	ConnectorOffer                     *string                                `json:"connectorOffer,omitempty"`
	ConsistencyPolicy                  *ConsistencyPolicy_Status              `json:"consistencyPolicy,omitempty"`
	Cors                               []CorsPolicy_Status                    `json:"cors,omitempty"`
	DatabaseAccountOfferType           *string                                `json:"databaseAccountOfferType,omitempty"`
	DefaultIdentity                    *string                                `json:"defaultIdentity,omitempty"`
	DisableKeyBasedMetadataWriteAccess *bool                                  `json:"disableKeyBasedMetadataWriteAccess,omitempty"`
	EnableAnalyticalStorage            *bool                                  `json:"enableAnalyticalStorage,omitempty"`
	EnableAutomaticFailover            *bool                                  `json:"enableAutomaticFailover,omitempty"`
	EnableCassandraConnector           *bool                                  `json:"enableCassandraConnector,omitempty"`
	EnableFreeTier                     *bool                                  `json:"enableFreeTier,omitempty"`
	EnableMultipleWriteLocations       *bool                                  `json:"enableMultipleWriteLocations,omitempty"`
	Id                                 *string                                `json:"id,omitempty"`
	Identity                           *ManagedServiceIdentity_Status         `json:"identity,omitempty"`
	IpRules                            []IpAddressOrRange_Status              `json:"ipRules,omitempty"`
	IsVirtualNetworkFilterEnabled      *bool                                  `json:"isVirtualNetworkFilterEnabled,omitempty"`
	KeyVaultKeyUri                     *string                                `json:"keyVaultKeyUri,omitempty"`
	Kind                               *string                                `json:"kind,omitempty"`
	Location                           *string                                `json:"location,omitempty"`
	Locations                          []Location_Status                      `json:"locations,omitempty"`
	Name                               *string                                `json:"name,omitempty"`
	NetworkAclBypass                   *string                                `json:"networkAclBypass,omitempty"`
	NetworkAclBypassResourceIds        []string                               `json:"networkAclBypassResourceIds,omitempty"`
	PropertyBag                        genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	PublicNetworkAccess                *string                                `json:"publicNetworkAccess,omitempty"`
	Tags                               map[string]string                      `json:"tags,omitempty"`
	Type                               *string                                `json:"type,omitempty"`
	VirtualNetworkRules                []VirtualNetworkRule_Status            `json:"virtualNetworkRules,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccount_Status{}

// ConvertStatusFrom populates our DatabaseAccount_Status from the provided source
func (account *DatabaseAccount_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(account)
}

// ConvertStatusTo populates the provided destination from our DatabaseAccount_Status
func (account *DatabaseAccount_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(account)
}

//Storage version of v1alpha1api20210515.DatabaseAccounts_SPEC
type DatabaseAccounts_SPEC struct {
	AnalyticalStorageConfiguration *AnalyticalStorageConfiguration_Spec `json:"analyticalStorageConfiguration,omitempty"`
	ApiProperties                  *ApiProperties_Spec                  `json:"apiProperties,omitempty"`

	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName                          string                       `json:"azureName"`
	BackupPolicy                       *BackupPolicy_Spec           `json:"backupPolicy,omitempty"`
	Capabilities                       []Capability_Spec            `json:"capabilities,omitempty"`
	ConnectorOffer                     *string                      `json:"connectorOffer,omitempty"`
	ConsistencyPolicy                  *ConsistencyPolicy_Spec      `json:"consistencyPolicy,omitempty"`
	Cors                               []CorsPolicy_Spec            `json:"cors,omitempty"`
	DatabaseAccountOfferType           *string                      `json:"databaseAccountOfferType,omitempty"`
	DefaultIdentity                    *string                      `json:"defaultIdentity,omitempty"`
	DisableKeyBasedMetadataWriteAccess *bool                        `json:"disableKeyBasedMetadataWriteAccess,omitempty"`
	EnableAnalyticalStorage            *bool                        `json:"enableAnalyticalStorage,omitempty"`
	EnableAutomaticFailover            *bool                        `json:"enableAutomaticFailover,omitempty"`
	EnableCassandraConnector           *bool                        `json:"enableCassandraConnector,omitempty"`
	EnableFreeTier                     *bool                        `json:"enableFreeTier,omitempty"`
	EnableMultipleWriteLocations       *bool                        `json:"enableMultipleWriteLocations,omitempty"`
	Identity                           *ManagedServiceIdentity_Spec `json:"identity,omitempty"`
	IpRules                            []IpAddressOrRange_Spec      `json:"ipRules,omitempty"`
	IsVirtualNetworkFilterEnabled      *bool                        `json:"isVirtualNetworkFilterEnabled,omitempty"`
	KeyVaultKeyUri                     *string                      `json:"keyVaultKeyUri,omitempty"`
	Kind                               *string                      `json:"kind,omitempty"`
	Location                           *string                      `json:"location,omitempty"`
	Locations                          []Location_Spec              `json:"locations,omitempty"`
	NetworkAclBypass                   *string                      `json:"networkAclBypass,omitempty"`
	NetworkAclBypassResourceIds        []string                     `json:"networkAclBypassResourceIds,omitempty"`
	OriginalVersion                    string                       `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner               genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag         genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                           `json:"publicNetworkAccess,omitempty"`
	Tags                map[string]string                 `json:"tags,omitempty"`
	VirtualNetworkRules []VirtualNetworkRule_Spec         `json:"virtualNetworkRules,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_SPEC{}

// ConvertSpecFrom populates our DatabaseAccounts_SPEC from the provided source
func (spec *DatabaseAccounts_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_SPEC
func (spec *DatabaseAccounts_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

//Storage version of v1alpha1api20210515.AnalyticalStorageConfiguration_Spec
type AnalyticalStorageConfiguration_Spec struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SchemaType  *string                `json:"schemaType,omitempty"`
}

//Storage version of v1alpha1api20210515.AnalyticalStorageConfiguration_Status
type AnalyticalStorageConfiguration_Status struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SchemaType  *string                `json:"schemaType,omitempty"`
}

//Storage version of v1alpha1api20210515.ApiProperties_Spec
type ApiProperties_Spec struct {
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServerVersion *string                `json:"serverVersion,omitempty"`
}

//Storage version of v1alpha1api20210515.ApiProperties_Status
type ApiProperties_Status struct {
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServerVersion *string                `json:"serverVersion,omitempty"`
}

//Storage version of v1alpha1api20210515.BackupPolicy_Spec
type BackupPolicy_Spec struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210515.BackupPolicy_Status
type BackupPolicy_Status struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210515.Capability_Spec
type Capability_Spec struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.Capability_Status
type Capability_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ConsistencyPolicy_Spec
type ConsistencyPolicy_Spec struct {
	DefaultConsistencyLevel *string                `json:"defaultConsistencyLevel,omitempty"`
	MaxIntervalInSeconds    *int                   `json:"maxIntervalInSeconds,omitempty"`
	MaxStalenessPrefix      *int                   `json:"maxStalenessPrefix,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ConsistencyPolicy_Status
type ConsistencyPolicy_Status struct {
	DefaultConsistencyLevel *string                `json:"defaultConsistencyLevel,omitempty"`
	MaxIntervalInSeconds    *int                   `json:"maxIntervalInSeconds,omitempty"`
	MaxStalenessPrefix      *int                   `json:"maxStalenessPrefix,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.CorsPolicy_Spec
type CorsPolicy_Spec struct {
	AllowedHeaders  *string                `json:"allowedHeaders,omitempty"`
	AllowedMethods  *string                `json:"allowedMethods,omitempty"`
	AllowedOrigins  *string                `json:"allowedOrigins,omitempty"`
	ExposedHeaders  *string                `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                   `json:"maxAgeInSeconds,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.CorsPolicy_Status
type CorsPolicy_Status struct {
	AllowedHeaders  *string                `json:"allowedHeaders,omitempty"`
	AllowedMethods  *string                `json:"allowedMethods,omitempty"`
	AllowedOrigins  *string                `json:"allowedOrigins,omitempty"`
	ExposedHeaders  *string                `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                   `json:"maxAgeInSeconds,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.IpAddressOrRange_Spec
type IpAddressOrRange_Spec struct {
	IpAddressOrRange *string                `json:"ipAddressOrRange,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.IpAddressOrRange_Status
type IpAddressOrRange_Status struct {
	IpAddressOrRange *string                `json:"ipAddressOrRange,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.Location_Spec
type Location_Spec struct {
	FailoverPriority  *int                   `json:"failoverPriority,omitempty"`
	IsZoneRedundant   *bool                  `json:"isZoneRedundant,omitempty"`
	LocationName      *string                `json:"locationName,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
}

//Storage version of v1alpha1api20210515.Location_Status
type Location_Status struct {
	DocumentEndpoint  *string                `json:"documentEndpoint,omitempty"`
	FailoverPriority  *int                   `json:"failoverPriority,omitempty"`
	Id                *string                `json:"id,omitempty"`
	IsZoneRedundant   *bool                  `json:"isZoneRedundant,omitempty"`
	LocationName      *string                `json:"locationName,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
}

//Storage version of v1alpha1api20210515.ManagedServiceIdentity_Spec
type ManagedServiceIdentity_Spec struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210515.ManagedServiceIdentity_Status
type ManagedServiceIdentity_Status struct {
	PrincipalId            *string                                                         `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                                          `json:"$propertyBag,omitempty"`
	TenantId               *string                                                         `json:"tenantId,omitempty"`
	Type                   *string                                                         `json:"type,omitempty"`
	UserAssignedIdentities map[string]ManagedServiceIdentity_UserAssignedIdentities_Status `json:"userAssignedIdentities,omitempty"`
}

//Storage version of v1alpha1api20210515.VirtualNetworkRule_Spec
type VirtualNetworkRule_Spec struct {
	IgnoreMissingVNetServiceEndpoint *bool                  `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
	PropertyBag                      genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//Reference: Resource ID of a subnet, for example:
	///subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

//Storage version of v1alpha1api20210515.VirtualNetworkRule_Status
type VirtualNetworkRule_Status struct {
	Id                               *string                `json:"id,omitempty"`
	IgnoreMissingVNetServiceEndpoint *bool                  `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
	PropertyBag                      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ManagedServiceIdentity_UserAssignedIdentities_Status
type ManagedServiceIdentity_UserAssignedIdentities_Status struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DatabaseAccount{}, &DatabaseAccountList{})
}
