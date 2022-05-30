// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=storage.azure.com,resources=storageaccounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.azure.com,resources={storageaccounts/status,storageaccounts/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210401.StorageAccount
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2021-04-01/storage.json
type StorageAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccount_Spec   `json:"spec,omitempty"`
	Status            StorageAccount_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccount{}

// GetConditions returns the conditions of the resource
func (account *StorageAccount) GetConditions() conditions.Conditions {
	return account.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (account *StorageAccount) SetConditions(conditions conditions.Conditions) {
	account.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &StorageAccount{}

// AzureName returns the Azure name of the resource
func (account *StorageAccount) AzureName() string {
	return account.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (account StorageAccount) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (account *StorageAccount) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (account *StorageAccount) GetSpec() genruntime.ConvertibleSpec {
	return &account.Spec
}

// GetStatus returns the status of this resource
func (account *StorageAccount) GetStatus() genruntime.ConvertibleStatus {
	return &account.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (account *StorageAccount) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (account *StorageAccount) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccount_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (account *StorageAccount) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(account.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  account.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (account *StorageAccount) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccount_STATUS); ok {
		account.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccount_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	account.Status = st
	return nil
}

// Hub marks that this StorageAccount is the hub type for conversion
func (account *StorageAccount) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (account *StorageAccount) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: account.Spec.OriginalVersion,
		Kind:    "StorageAccount",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210401.StorageAccount
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2021-04-01/storage.json
type StorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccount `json:"items"`
}

// Storage version of v1beta20210401.APIVersion
// +kubebuilder:validation:Enum={"2021-04-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-04-01")

// Storage version of v1beta20210401.StorageAccount_STATUS
type StorageAccount_STATUS struct {
	AccessTier                            *string                                       `json:"accessTier,omitempty"`
	AllowBlobPublicAccess                 *bool                                         `json:"allowBlobPublicAccess,omitempty"`
	AllowCrossTenantReplication           *bool                                         `json:"allowCrossTenantReplication,omitempty"`
	AllowSharedKeyAccess                  *bool                                         `json:"allowSharedKeyAccess,omitempty"`
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication_STATUS `json:"azureFilesIdentityBasedAuthentication,omitempty"`
	BlobRestoreStatus                     *BlobRestoreStatus_STATUS                     `json:"blobRestoreStatus,omitempty"`
	Conditions                            []conditions.Condition                        `json:"conditions,omitempty"`
	CreationTime                          *string                                       `json:"creationTime,omitempty"`
	CustomDomain                          *CustomDomain_STATUS                          `json:"customDomain,omitempty"`
	Encryption                            *Encryption_STATUS                            `json:"encryption,omitempty"`
	ExtendedLocation                      *ExtendedLocation_STATUS                      `json:"extendedLocation,omitempty"`
	FailoverInProgress                    *bool                                         `json:"failoverInProgress,omitempty"`
	GeoReplicationStats                   *GeoReplicationStats_STATUS                   `json:"geoReplicationStats,omitempty"`
	Id                                    *string                                       `json:"id,omitempty"`
	Identity                              *Identity_STATUS                              `json:"identity,omitempty"`
	IsHnsEnabled                          *bool                                         `json:"isHnsEnabled,omitempty"`
	IsNfsV3Enabled                        *bool                                         `json:"isNfsV3Enabled,omitempty"`
	KeyCreationTime                       *KeyCreationTime_STATUS                       `json:"keyCreationTime,omitempty"`
	KeyPolicy                             *KeyPolicy_STATUS                             `json:"keyPolicy,omitempty"`
	Kind                                  *string                                       `json:"kind,omitempty"`
	LargeFileSharesState                  *string                                       `json:"largeFileSharesState,omitempty"`
	LastGeoFailoverTime                   *string                                       `json:"lastGeoFailoverTime,omitempty"`
	Location                              *string                                       `json:"location,omitempty"`
	MinimumTlsVersion                     *string                                       `json:"minimumTlsVersion,omitempty"`
	Name                                  *string                                       `json:"name,omitempty"`
	NetworkAcls                           *NetworkRuleSet_STATUS                        `json:"networkAcls,omitempty"`
	PrimaryEndpoints                      *Endpoints_STATUS                             `json:"primaryEndpoints,omitempty"`
	PrimaryLocation                       *string                                       `json:"primaryLocation,omitempty"`
	PrivateEndpointConnections            []PrivateEndpointConnection_STATUS            `json:"privateEndpointConnections,omitempty"`
	PropertyBag                           genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	ProvisioningState                     *string                                       `json:"provisioningState,omitempty"`
	RoutingPreference                     *RoutingPreference_STATUS                     `json:"routingPreference,omitempty"`
	SasPolicy                             *SasPolicy_STATUS                             `json:"sasPolicy,omitempty"`
	SecondaryEndpoints                    *Endpoints_STATUS                             `json:"secondaryEndpoints,omitempty"`
	SecondaryLocation                     *string                                       `json:"secondaryLocation,omitempty"`
	Sku                                   *Sku_STATUS                                   `json:"sku,omitempty"`
	StatusOfPrimary                       *string                                       `json:"statusOfPrimary,omitempty"`
	StatusOfSecondary                     *string                                       `json:"statusOfSecondary,omitempty"`
	SupportsHttpsTrafficOnly              *bool                                         `json:"supportsHttpsTrafficOnly,omitempty"`
	Tags                                  map[string]string                             `json:"tags,omitempty"`
	Type                                  *string                                       `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccount_STATUS{}

// ConvertStatusFrom populates our StorageAccount_STATUS from the provided source
func (account *StorageAccount_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(account)
}

// ConvertStatusTo populates the provided destination from our StorageAccount_STATUS
func (account *StorageAccount_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(account)
}

// Storage version of v1beta20210401.StorageAccount_Spec
type StorageAccount_Spec struct {
	AccessTier                            *string                                `json:"accessTier,omitempty"`
	AllowBlobPublicAccess                 *bool                                  `json:"allowBlobPublicAccess,omitempty"`
	AllowCrossTenantReplication           *bool                                  `json:"allowCrossTenantReplication,omitempty"`
	AllowSharedKeyAccess                  *bool                                  `json:"allowSharedKeyAccess,omitempty"`
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication `json:"azureFilesIdentityBasedAuthentication,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName            string                      `json:"azureName,omitempty"`
	CustomDomain         *CustomDomain               `json:"customDomain,omitempty"`
	Encryption           *Encryption                 `json:"encryption,omitempty"`
	ExtendedLocation     *ExtendedLocation           `json:"extendedLocation,omitempty"`
	Identity             *Identity                   `json:"identity,omitempty"`
	IsHnsEnabled         *bool                       `json:"isHnsEnabled,omitempty"`
	IsNfsV3Enabled       *bool                       `json:"isNfsV3Enabled,omitempty"`
	KeyPolicy            *KeyPolicy                  `json:"keyPolicy,omitempty"`
	Kind                 *string                     `json:"kind,omitempty"`
	LargeFileSharesState *string                     `json:"largeFileSharesState,omitempty"`
	Location             *string                     `json:"location,omitempty"`
	MinimumTlsVersion    *string                     `json:"minimumTlsVersion,omitempty"`
	NetworkAcls          *NetworkRuleSet             `json:"networkAcls,omitempty"`
	OperatorSpec         *StorageAccountOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion      string                      `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                    *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag              genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RoutingPreference        *RoutingPreference                 `json:"routingPreference,omitempty"`
	SasPolicy                *SasPolicy                         `json:"sasPolicy,omitempty"`
	Sku                      *Sku                               `json:"sku,omitempty"`
	SupportsHttpsTrafficOnly *bool                              `json:"supportsHttpsTrafficOnly,omitempty"`
	Tags                     map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccount_Spec{}

// ConvertSpecFrom populates our StorageAccount_Spec from the provided source
func (account *StorageAccount_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(account)
}

// ConvertSpecTo populates the provided destination from our StorageAccount_Spec
func (account *StorageAccount_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(account)
}

// Storage version of v1beta20210401.AzureFilesIdentityBasedAuthentication
type AzureFilesIdentityBasedAuthentication struct {
	ActiveDirectoryProperties *ActiveDirectoryProperties `json:"activeDirectoryProperties,omitempty"`
	DefaultSharePermission    *string                    `json:"defaultSharePermission,omitempty"`
	DirectoryServiceOptions   *string                    `json:"directoryServiceOptions,omitempty"`
	PropertyBag               genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401.AzureFilesIdentityBasedAuthentication_STATUS
type AzureFilesIdentityBasedAuthentication_STATUS struct {
	ActiveDirectoryProperties *ActiveDirectoryProperties_STATUS `json:"activeDirectoryProperties,omitempty"`
	DefaultSharePermission    *string                           `json:"defaultSharePermission,omitempty"`
	DirectoryServiceOptions   *string                           `json:"directoryServiceOptions,omitempty"`
	PropertyBag               genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401.BlobRestoreStatus_STATUS
type BlobRestoreStatus_STATUS struct {
	FailureReason *string                       `json:"failureReason,omitempty"`
	Parameters    *BlobRestoreParameters_STATUS `json:"parameters,omitempty"`
	PropertyBag   genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	RestoreId     *string                       `json:"restoreId,omitempty"`
	Status        *string                       `json:"status,omitempty"`
}

// Storage version of v1beta20210401.CustomDomain
type CustomDomain struct {
	Name             *string                `json:"name,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UseSubDomainName *bool                  `json:"useSubDomainName,omitempty"`
}

// Storage version of v1beta20210401.CustomDomain_STATUS
type CustomDomain_STATUS struct {
	Name             *string                `json:"name,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UseSubDomainName *bool                  `json:"useSubDomainName,omitempty"`
}

// Storage version of v1beta20210401.Encryption
type Encryption struct {
	Identity                        *EncryptionIdentity    `json:"identity,omitempty"`
	KeySource                       *string                `json:"keySource,omitempty"`
	Keyvaultproperties              *KeyVaultProperties    `json:"keyvaultproperties,omitempty"`
	PropertyBag                     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequireInfrastructureEncryption *bool                  `json:"requireInfrastructureEncryption,omitempty"`
	Services                        *EncryptionServices    `json:"services,omitempty"`
}

// Storage version of v1beta20210401.Encryption_STATUS
type Encryption_STATUS struct {
	Identity                        *EncryptionIdentity_STATUS `json:"identity,omitempty"`
	KeySource                       *string                    `json:"keySource,omitempty"`
	Keyvaultproperties              *KeyVaultProperties_STATUS `json:"keyvaultproperties,omitempty"`
	PropertyBag                     genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	RequireInfrastructureEncryption *bool                      `json:"requireInfrastructureEncryption,omitempty"`
	Services                        *EncryptionServices_STATUS `json:"services,omitempty"`
}

// Storage version of v1beta20210401.Endpoints_STATUS
type Endpoints_STATUS struct {
	Blob               *string                                  `json:"blob,omitempty"`
	Dfs                *string                                  `json:"dfs,omitempty"`
	File               *string                                  `json:"file,omitempty"`
	InternetEndpoints  *StorageAccountInternetEndpoints_STATUS  `json:"internetEndpoints,omitempty"`
	MicrosoftEndpoints *StorageAccountMicrosoftEndpoints_STATUS `json:"microsoftEndpoints,omitempty"`
	PropertyBag        genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
	Queue              *string                                  `json:"queue,omitempty"`
	Table              *string                                  `json:"table,omitempty"`
	Web                *string                                  `json:"web,omitempty"`
}

// Storage version of v1beta20210401.ExtendedLocation
type ExtendedLocation struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1beta20210401.ExtendedLocation_STATUS
type ExtendedLocation_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1beta20210401.GeoReplicationStats_STATUS
type GeoReplicationStats_STATUS struct {
	CanFailover  *bool                  `json:"canFailover,omitempty"`
	LastSyncTime *string                `json:"lastSyncTime,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status       *string                `json:"status,omitempty"`
}

// Storage version of v1beta20210401.Identity
type Identity struct {
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId    *string                `json:"tenantId,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1beta20210401.Identity_STATUS
type Identity_STATUS struct {
	PrincipalId            *string                                `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	TenantId               *string                                `json:"tenantId,omitempty"`
	Type                   *string                                `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1beta20210401.KeyCreationTime_STATUS
type KeyCreationTime_STATUS struct {
	Key1        *string                `json:"key1,omitempty"`
	Key2        *string                `json:"key2,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401.KeyPolicy
type KeyPolicy struct {
	KeyExpirationPeriodInDays *int                   `json:"keyExpirationPeriodInDays,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401.KeyPolicy_STATUS
type KeyPolicy_STATUS struct {
	KeyExpirationPeriodInDays *int                   `json:"keyExpirationPeriodInDays,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401.NetworkRuleSet
type NetworkRuleSet struct {
	Bypass              *string                `json:"bypass,omitempty"`
	DefaultAction       *string                `json:"defaultAction,omitempty"`
	IpRules             []IPRule               `json:"ipRules,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceAccessRules []ResourceAccessRule   `json:"resourceAccessRules,omitempty"`
	VirtualNetworkRules []VirtualNetworkRule   `json:"virtualNetworkRules,omitempty"`
}

// Storage version of v1beta20210401.NetworkRuleSet_STATUS
type NetworkRuleSet_STATUS struct {
	Bypass              *string                     `json:"bypass,omitempty"`
	DefaultAction       *string                     `json:"defaultAction,omitempty"`
	IpRules             []IPRule_STATUS             `json:"ipRules,omitempty"`
	PropertyBag         genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	ResourceAccessRules []ResourceAccessRule_STATUS `json:"resourceAccessRules,omitempty"`
	VirtualNetworkRules []VirtualNetworkRule_STATUS `json:"virtualNetworkRules,omitempty"`
}

// Storage version of v1beta20210401.PrivateEndpointConnection_STATUS
type PrivateEndpointConnection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401.RoutingPreference
type RoutingPreference struct {
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublishInternetEndpoints  *bool                  `json:"publishInternetEndpoints,omitempty"`
	PublishMicrosoftEndpoints *bool                  `json:"publishMicrosoftEndpoints,omitempty"`
	RoutingChoice             *string                `json:"routingChoice,omitempty"`
}

// Storage version of v1beta20210401.RoutingPreference_STATUS
type RoutingPreference_STATUS struct {
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublishInternetEndpoints  *bool                  `json:"publishInternetEndpoints,omitempty"`
	PublishMicrosoftEndpoints *bool                  `json:"publishMicrosoftEndpoints,omitempty"`
	RoutingChoice             *string                `json:"routingChoice,omitempty"`
}

// Storage version of v1beta20210401.SasPolicy
type SasPolicy struct {
	ExpirationAction    *string                `json:"expirationAction,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SasExpirationPeriod *string                `json:"sasExpirationPeriod,omitempty"`
}

// Storage version of v1beta20210401.SasPolicy_STATUS
type SasPolicy_STATUS struct {
	ExpirationAction    *string                `json:"expirationAction,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SasExpirationPeriod *string                `json:"sasExpirationPeriod,omitempty"`
}

// Storage version of v1beta20210401.Sku
type Sku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20210401.Sku_STATUS
type Sku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20210401.StorageAccountOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type StorageAccountOperatorSpec struct {
	PropertyBag genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	Secrets     *StorageAccountOperatorSecrets `json:"secrets,omitempty"`
}

// Storage version of v1beta20210401.ActiveDirectoryProperties
type ActiveDirectoryProperties struct {
	AzureStorageSid   *string                `json:"azureStorageSid,omitempty"`
	DomainGuid        *string                `json:"domainGuid,omitempty"`
	DomainName        *string                `json:"domainName,omitempty"`
	DomainSid         *string                `json:"domainSid,omitempty"`
	ForestName        *string                `json:"forestName,omitempty"`
	NetBiosDomainName *string                `json:"netBiosDomainName,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401.ActiveDirectoryProperties_STATUS
type ActiveDirectoryProperties_STATUS struct {
	AzureStorageSid   *string                `json:"azureStorageSid,omitempty"`
	DomainGuid        *string                `json:"domainGuid,omitempty"`
	DomainName        *string                `json:"domainName,omitempty"`
	DomainSid         *string                `json:"domainSid,omitempty"`
	ForestName        *string                `json:"forestName,omitempty"`
	NetBiosDomainName *string                `json:"netBiosDomainName,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401.BlobRestoreParameters_STATUS
type BlobRestoreParameters_STATUS struct {
	BlobRanges    []BlobRestoreRange_STATUS `json:"blobRanges,omitempty"`
	PropertyBag   genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	TimeToRestore *string                   `json:"timeToRestore,omitempty"`
}

// Storage version of v1beta20210401.EncryptionIdentity
type EncryptionIdentity struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// UserAssignedIdentityReference: Resource identifier of the UserAssigned identity to be associated with server-side
	// encryption on the storage account.
	UserAssignedIdentityReference *genruntime.ResourceReference `armReference:"UserAssignedIdentity" json:"userAssignedIdentityReference,omitempty"`
}

// Storage version of v1beta20210401.EncryptionIdentity_STATUS
type EncryptionIdentity_STATUS struct {
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UserAssignedIdentity *string                `json:"userAssignedIdentity,omitempty"`
}

// Storage version of v1beta20210401.EncryptionServices
type EncryptionServices struct {
	Blob        *EncryptionService     `json:"blob,omitempty"`
	File        *EncryptionService     `json:"file,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Queue       *EncryptionService     `json:"queue,omitempty"`
	Table       *EncryptionService     `json:"table,omitempty"`
}

// Storage version of v1beta20210401.EncryptionServices_STATUS
type EncryptionServices_STATUS struct {
	Blob        *EncryptionService_STATUS `json:"blob,omitempty"`
	File        *EncryptionService_STATUS `json:"file,omitempty"`
	PropertyBag genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Queue       *EncryptionService_STATUS `json:"queue,omitempty"`
	Table       *EncryptionService_STATUS `json:"table,omitempty"`
}

// Storage version of v1beta20210401.IPRule
type IPRule struct {
	Action      *string                `json:"action,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1beta20210401.IPRule_STATUS
type IPRule_STATUS struct {
	Action      *string                `json:"action,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1beta20210401.KeyVaultProperties
type KeyVaultProperties struct {
	CurrentVersionedKeyIdentifier *string                `json:"currentVersionedKeyIdentifier,omitempty"`
	Keyname                       *string                `json:"keyname,omitempty"`
	Keyvaulturi                   *string                `json:"keyvaulturi,omitempty"`
	Keyversion                    *string                `json:"keyversion,omitempty"`
	LastKeyRotationTimestamp      *string                `json:"lastKeyRotationTimestamp,omitempty"`
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401.KeyVaultProperties_STATUS
type KeyVaultProperties_STATUS struct {
	CurrentVersionedKeyIdentifier *string                `json:"currentVersionedKeyIdentifier,omitempty"`
	Keyname                       *string                `json:"keyname,omitempty"`
	Keyvaulturi                   *string                `json:"keyvaulturi,omitempty"`
	Keyversion                    *string                `json:"keyversion,omitempty"`
	LastKeyRotationTimestamp      *string                `json:"lastKeyRotationTimestamp,omitempty"`
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401.ResourceAccessRule
type ResourceAccessRule struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ResourceReference: Resource Id
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
	TenantId          *string                       `json:"tenantId,omitempty"`
}

// Storage version of v1beta20210401.ResourceAccessRule_STATUS
type ResourceAccessRule_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId  *string                `json:"resourceId,omitempty"`
	TenantId    *string                `json:"tenantId,omitempty"`
}

// Storage version of v1beta20210401.StorageAccountInternetEndpoints_STATUS
type StorageAccountInternetEndpoints_STATUS struct {
	Blob        *string                `json:"blob,omitempty"`
	Dfs         *string                `json:"dfs,omitempty"`
	File        *string                `json:"file,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Web         *string                `json:"web,omitempty"`
}

// Storage version of v1beta20210401.StorageAccountMicrosoftEndpoints_STATUS
type StorageAccountMicrosoftEndpoints_STATUS struct {
	Blob        *string                `json:"blob,omitempty"`
	Dfs         *string                `json:"dfs,omitempty"`
	File        *string                `json:"file,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Queue       *string                `json:"queue,omitempty"`
	Table       *string                `json:"table,omitempty"`
	Web         *string                `json:"web,omitempty"`
}

// Storage version of v1beta20210401.StorageAccountOperatorSecrets
type StorageAccountOperatorSecrets struct {
	BlobEndpoint  *genruntime.SecretDestination `json:"blobEndpoint,omitempty"`
	DfsEndpoint   *genruntime.SecretDestination `json:"dfsEndpoint,omitempty"`
	FileEndpoint  *genruntime.SecretDestination `json:"fileEndpoint,omitempty"`
	Key1          *genruntime.SecretDestination `json:"key1,omitempty"`
	Key2          *genruntime.SecretDestination `json:"key2,omitempty"`
	PropertyBag   genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	QueueEndpoint *genruntime.SecretDestination `json:"queueEndpoint,omitempty"`
	TableEndpoint *genruntime.SecretDestination `json:"tableEndpoint,omitempty"`
	WebEndpoint   *genruntime.SecretDestination `json:"webEndpoint,omitempty"`
}

// Storage version of v1beta20210401.UserAssignedIdentity_STATUS
type UserAssignedIdentity_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401.VirtualNetworkRule
type VirtualNetworkRule struct {
	Action      *string                `json:"action,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// Reference: Resource ID of a subnet, for example:
	// /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
	State     *string                       `json:"state,omitempty"`
}

// Storage version of v1beta20210401.VirtualNetworkRule_STATUS
type VirtualNetworkRule_STATUS struct {
	Action      *string                `json:"action,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	State       *string                `json:"state,omitempty"`
}

// Storage version of v1beta20210401.BlobRestoreRange_STATUS
type BlobRestoreRange_STATUS struct {
	EndRange    *string                `json:"endRange,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartRange  *string                `json:"startRange,omitempty"`
}

// Storage version of v1beta20210401.EncryptionService
type EncryptionService struct {
	Enabled         *bool                  `json:"enabled,omitempty"`
	KeyType         *string                `json:"keyType,omitempty"`
	LastEnabledTime *string                `json:"lastEnabledTime,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401.EncryptionService_STATUS
type EncryptionService_STATUS struct {
	Enabled         *bool                  `json:"enabled,omitempty"`
	KeyType         *string                `json:"keyType,omitempty"`
	LastEnabledTime *string                `json:"lastEnabledTime,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&StorageAccount{}, &StorageAccountList{})
}