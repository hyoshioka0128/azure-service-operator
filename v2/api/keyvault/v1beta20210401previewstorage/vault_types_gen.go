// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401previewstorage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=keyvault.azure.com,resources=vaults,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=keyvault.azure.com,resources={vaults/status,vaults/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210401preview.Vault
// Generator information:
// - Generated from: /keyvault/resource-manager/Microsoft.KeyVault/preview/2021-04-01-preview/keyvault.json
type Vault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Vault_Spec                           `json:"spec,omitempty"`
	Status            VaultCreateOrUpdateParameters_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Vault{}

// GetConditions returns the conditions of the resource
func (vault *Vault) GetConditions() conditions.Conditions {
	return vault.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (vault *Vault) SetConditions(conditions conditions.Conditions) {
	vault.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Vault{}

// AzureName returns the Azure name of the resource
func (vault *Vault) AzureName() string {
	return vault.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "20210401preview"
func (vault Vault) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (vault *Vault) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (vault *Vault) GetSpec() genruntime.ConvertibleSpec {
	return &vault.Spec
}

// GetStatus returns the status of this resource
func (vault *Vault) GetStatus() genruntime.ConvertibleStatus {
	return &vault.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (vault *Vault) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (vault *Vault) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &VaultCreateOrUpdateParameters_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (vault *Vault) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(vault.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  vault.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (vault *Vault) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*VaultCreateOrUpdateParameters_STATUS); ok {
		vault.Status = *st
		return nil
	}

	// Convert status to required version
	var st VaultCreateOrUpdateParameters_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	vault.Status = st
	return nil
}

// Hub marks that this Vault is the hub type for conversion
func (vault *Vault) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (vault *Vault) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: vault.Spec.OriginalVersion,
		Kind:    "Vault",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210401preview.Vault
// Generator information:
// - Generated from: /keyvault/resource-manager/Microsoft.KeyVault/preview/2021-04-01-preview/keyvault.json
type VaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vault `json:"items"`
}

// Storage version of v1beta20210401preview.APIVersion
// +kubebuilder:validation:Enum={"20210401preview"}
type APIVersion string

const APIVersionValue = APIVersion("20210401preview")

// Storage version of v1beta20210401preview.VaultCreateOrUpdateParameters_STATUS
type VaultCreateOrUpdateParameters_STATUS struct {
	Conditions  []conditions.Condition  `json:"conditions,omitempty"`
	Location    *string                 `json:"location,omitempty"`
	Properties  *VaultProperties_STATUS `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	Tags        map[string]string       `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleStatus = &VaultCreateOrUpdateParameters_STATUS{}

// ConvertStatusFrom populates our VaultCreateOrUpdateParameters_STATUS from the provided source
func (parameters *VaultCreateOrUpdateParameters_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == parameters {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(parameters)
}

// ConvertStatusTo populates the provided destination from our VaultCreateOrUpdateParameters_STATUS
func (parameters *VaultCreateOrUpdateParameters_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == parameters {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(parameters)
}

// Storage version of v1beta20210401preview.Vault_Spec
type Vault_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	Properties  *VaultProperties                   `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Vault_Spec{}

// ConvertSpecFrom populates our Vault_Spec from the provided source
func (vault *Vault_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == vault {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(vault)
}

// ConvertSpecTo populates the provided destination from our Vault_Spec
func (vault *Vault_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == vault {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(vault)
}

// Storage version of v1beta20210401preview.VaultProperties
type VaultProperties struct {
	AccessPolicies               []AccessPolicyEntry             `json:"accessPolicies,omitempty"`
	CreateMode                   *string                         `json:"createMode,omitempty"`
	EnablePurgeProtection        *bool                           `json:"enablePurgeProtection,omitempty"`
	EnableRbacAuthorization      *bool                           `json:"enableRbacAuthorization,omitempty"`
	EnableSoftDelete             *bool                           `json:"enableSoftDelete,omitempty"`
	EnabledForDeployment         *bool                           `json:"enabledForDeployment,omitempty"`
	EnabledForDiskEncryption     *bool                           `json:"enabledForDiskEncryption,omitempty"`
	EnabledForTemplateDeployment *bool                           `json:"enabledForTemplateDeployment,omitempty"`
	HsmPoolResourceId            *string                         `json:"hsmPoolResourceId,omitempty"`
	NetworkAcls                  *NetworkRuleSet                 `json:"networkAcls,omitempty"`
	PrivateEndpointConnections   []PrivateEndpointConnectionItem `json:"privateEndpointConnections,omitempty"`
	PropertyBag                  genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
	ProvisioningState            *string                         `json:"provisioningState,omitempty"`
	Sku                          *Sku                            `json:"sku,omitempty"`
	SoftDeleteRetentionInDays    *int                            `json:"softDeleteRetentionInDays,omitempty"`
	TenantId                     *string                         `json:"tenantId,omitempty"`
	VaultUri                     *string                         `json:"vaultUri,omitempty"`
}

// Storage version of v1beta20210401preview.VaultProperties_STATUS
type VaultProperties_STATUS struct {
	AccessPolicies               []AccessPolicyEntry_STATUS             `json:"accessPolicies,omitempty"`
	CreateMode                   *string                                `json:"createMode,omitempty"`
	EnablePurgeProtection        *bool                                  `json:"enablePurgeProtection,omitempty"`
	EnableRbacAuthorization      *bool                                  `json:"enableRbacAuthorization,omitempty"`
	EnableSoftDelete             *bool                                  `json:"enableSoftDelete,omitempty"`
	EnabledForDeployment         *bool                                  `json:"enabledForDeployment,omitempty"`
	EnabledForDiskEncryption     *bool                                  `json:"enabledForDiskEncryption,omitempty"`
	EnabledForTemplateDeployment *bool                                  `json:"enabledForTemplateDeployment,omitempty"`
	HsmPoolResourceId            *string                                `json:"hsmPoolResourceId,omitempty"`
	NetworkAcls                  *NetworkRuleSet_STATUS                 `json:"networkAcls,omitempty"`
	PrivateEndpointConnections   []PrivateEndpointConnectionItem_STATUS `json:"privateEndpointConnections,omitempty"`
	PropertyBag                  genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	ProvisioningState            *string                                `json:"provisioningState,omitempty"`
	Sku                          *Sku_STATUS                            `json:"sku,omitempty"`
	SoftDeleteRetentionInDays    *int                                   `json:"softDeleteRetentionInDays,omitempty"`
	TenantId                     *string                                `json:"tenantId,omitempty"`
	VaultUri                     *string                                `json:"vaultUri,omitempty"`
}

// Storage version of v1beta20210401preview.AccessPolicyEntry
type AccessPolicyEntry struct {
	ApplicationId *string                `json:"applicationId,omitempty"`
	ObjectId      *string                `json:"objectId,omitempty"`
	Permissions   *Permissions           `json:"permissions,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId      *string                `json:"tenantId,omitempty"`
}

// Storage version of v1beta20210401preview.AccessPolicyEntry_STATUS
type AccessPolicyEntry_STATUS struct {
	ApplicationId *string                `json:"applicationId,omitempty"`
	ObjectId      *string                `json:"objectId,omitempty"`
	Permissions   *Permissions_STATUS    `json:"permissions,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId      *string                `json:"tenantId,omitempty"`
}

// Storage version of v1beta20210401preview.NetworkRuleSet
type NetworkRuleSet struct {
	Bypass              *string                `json:"bypass,omitempty"`
	DefaultAction       *string                `json:"defaultAction,omitempty"`
	IpRules             []IPRule               `json:"ipRules,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VirtualNetworkRules []VirtualNetworkRule   `json:"virtualNetworkRules,omitempty"`
}

// Storage version of v1beta20210401preview.NetworkRuleSet_STATUS
type NetworkRuleSet_STATUS struct {
	Bypass              *string                     `json:"bypass,omitempty"`
	DefaultAction       *string                     `json:"defaultAction,omitempty"`
	IpRules             []IPRule_STATUS             `json:"ipRules,omitempty"`
	PropertyBag         genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	VirtualNetworkRules []VirtualNetworkRule_STATUS `json:"virtualNetworkRules,omitempty"`
}

// Storage version of v1beta20210401preview.PrivateEndpointConnectionItem
type PrivateEndpointConnectionItem struct {
	Etag                              *string                            `json:"etag,omitempty"`
	Id                                *string                            `json:"id,omitempty"`
	PrivateEndpoint                   *PrivateEndpoint                   `json:"privateEndpoint,omitempty"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`
	PropertyBag                       genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ProvisioningState                 *string                            `json:"provisioningState,omitempty"`
}

// Storage version of v1beta20210401preview.PrivateEndpointConnectionItem_STATUS
type PrivateEndpointConnectionItem_STATUS struct {
	Etag                              *string                                   `json:"etag,omitempty"`
	Id                                *string                                   `json:"id,omitempty"`
	PrivateEndpoint                   *PrivateEndpoint_STATUS                   `json:"privateEndpoint,omitempty"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState_STATUS `json:"privateLinkServiceConnectionState,omitempty"`
	PropertyBag                       genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	ProvisioningState                 *string                                   `json:"provisioningState,omitempty"`
}

// Storage version of v1beta20210401preview.Sku
type Sku struct {
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401preview.Sku_STATUS
type Sku_STATUS struct {
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401preview.IPRule
type IPRule struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1beta20210401preview.IPRule_STATUS
type IPRule_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1beta20210401preview.Permissions
type Permissions struct {
	Certificates []string               `json:"certificates,omitempty"`
	Keys         []string               `json:"keys,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Secrets      []string               `json:"secrets,omitempty"`
	Storage      []string               `json:"storage,omitempty"`
}

// Storage version of v1beta20210401preview.Permissions_STATUS
type Permissions_STATUS struct {
	Certificates []string               `json:"certificates,omitempty"`
	Keys         []string               `json:"keys,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Secrets      []string               `json:"secrets,omitempty"`
	Storage      []string               `json:"storage,omitempty"`
}

// Storage version of v1beta20210401preview.PrivateEndpoint
type PrivateEndpoint struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401preview.PrivateEndpoint_STATUS
type PrivateEndpoint_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401preview.PrivateLinkServiceConnectionState
type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string                `json:"actionsRequired,omitempty"`
	Description     *string                `json:"description,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status          *string                `json:"status,omitempty"`
}

// Storage version of v1beta20210401preview.PrivateLinkServiceConnectionState_STATUS
type PrivateLinkServiceConnectionState_STATUS struct {
	ActionsRequired *string                `json:"actionsRequired,omitempty"`
	Description     *string                `json:"description,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status          *string                `json:"status,omitempty"`
}

// Storage version of v1beta20210401preview.VirtualNetworkRule
type VirtualNetworkRule struct {
	IgnoreMissingVnetServiceEndpoint *bool                  `json:"ignoreMissingVnetServiceEndpoint,omitempty"`
	PropertyBag                      genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// Reference: Full resource id of a vnet subnet, such as
	// '/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1'.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20210401preview.VirtualNetworkRule_STATUS
type VirtualNetworkRule_STATUS struct {
	Id                               *string                `json:"id,omitempty"`
	IgnoreMissingVnetServiceEndpoint *bool                  `json:"ignoreMissingVnetServiceEndpoint,omitempty"`
	PropertyBag                      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Vault{}, &VaultList{})
}
