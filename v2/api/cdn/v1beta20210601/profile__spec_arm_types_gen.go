// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Profile_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Kind: Kind of the profile. Used by portal to differentiate traditional CDN profile and new AFD profile.
	Kind *string `json:"kind,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name       string                `json:"name,omitempty"`
	Properties *ProfilePropertiesARM `json:"properties,omitempty"`

	// Sku: The pricing tier (defines Azure Front Door Standard or Premium or a CDN provider, feature list and rate) of the
	// profile.
	Sku        *SkuARM        `json:"sku,omitempty"`
	SystemData *SystemDataARM `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Profile_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (profile Profile_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (profile *Profile_SpecARM) GetName() string {
	return profile.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (profile *Profile_SpecARM) GetType() string {
	return ""
}

type ProfilePropertiesARM struct {
	// FrontDoorId: The Id of the frontdoor.
	FrontDoorId *string `json:"frontDoorId,omitempty"`

	// OriginResponseTimeoutSeconds: Send and receive timeout on forwarding request to the origin. When timeout is reached, the
	// request fails and returns.
	OriginResponseTimeoutSeconds *int `json:"originResponseTimeoutSeconds,omitempty"`

	// ProvisioningState: Provisioning status of the profile.
	ProvisioningState *ProfileProperties_ProvisioningState `json:"provisioningState,omitempty"`

	// ResourceState: Resource status of the profile.
	ResourceState *ProfileProperties_ResourceState `json:"resourceState,omitempty"`
}

type SkuARM struct {
	// Name: Name of the pricing tier.
	Name *Sku_Name `json:"name,omitempty"`
}

type SystemDataARM struct {
	// CreatedAt: The timestamp of resource creation (UTC)
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: An identifier for the identity that created the resource
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource
	CreatedByType *IdentityType `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: An identifier for the identity that last modified the resource
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource
	LastModifiedByType *IdentityType `json:"lastModifiedByType,omitempty"`
}

// +kubebuilder:validation:Enum={"application","key","managedIdentity","user"}
type IdentityType string

const (
	IdentityType_Application     = IdentityType("application")
	IdentityType_Key             = IdentityType("key")
	IdentityType_ManagedIdentity = IdentityType("managedIdentity")
	IdentityType_User            = IdentityType("user")
)

// +kubebuilder:validation:Enum={"Custom_Verizon","Premium_AzureFrontDoor","Premium_Verizon","StandardPlus_955BandWidth_ChinaCdn","StandardPlus_AvgBandWidth_ChinaCdn","StandardPlus_ChinaCdn","Standard_955BandWidth_ChinaCdn","Standard_Akamai","Standard_AvgBandWidth_ChinaCdn","Standard_AzureFrontDoor","Standard_ChinaCdn","Standard_Microsoft","Standard_Verizon"}
type Sku_Name string

const (
	Sku_Name_Custom_Verizon                     = Sku_Name("Custom_Verizon")
	Sku_Name_Premium_AzureFrontDoor             = Sku_Name("Premium_AzureFrontDoor")
	Sku_Name_Premium_Verizon                    = Sku_Name("Premium_Verizon")
	Sku_Name_StandardPlus_955BandWidth_ChinaCdn = Sku_Name("StandardPlus_955BandWidth_ChinaCdn")
	Sku_Name_StandardPlus_AvgBandWidth_ChinaCdn = Sku_Name("StandardPlus_AvgBandWidth_ChinaCdn")
	Sku_Name_StandardPlus_ChinaCdn              = Sku_Name("StandardPlus_ChinaCdn")
	Sku_Name_Standard_955BandWidth_ChinaCdn     = Sku_Name("Standard_955BandWidth_ChinaCdn")
	Sku_Name_Standard_Akamai                    = Sku_Name("Standard_Akamai")
	Sku_Name_Standard_AvgBandWidth_ChinaCdn     = Sku_Name("Standard_AvgBandWidth_ChinaCdn")
	Sku_Name_Standard_AzureFrontDoor            = Sku_Name("Standard_AzureFrontDoor")
	Sku_Name_Standard_ChinaCdn                  = Sku_Name("Standard_ChinaCdn")
	Sku_Name_Standard_Microsoft                 = Sku_Name("Standard_Microsoft")
	Sku_Name_Standard_Verizon                   = Sku_Name("Standard_Verizon")
)
