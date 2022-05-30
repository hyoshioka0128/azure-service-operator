// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Domain_Spec. Use v1beta20200601.Domain_Spec instead
type Domain_SpecARM struct {
	AzureName  string               `json:"azureName,omitempty"`
	Id         *string              `json:"id,omitempty"`
	Location   *string              `json:"location,omitempty"`
	Name       string               `json:"name,omitempty"`
	Properties *DomainPropertiesARM `json:"properties,omitempty"`
	SystemData *SystemDataARM       `json:"systemData,omitempty"`
	Tags       map[string]string    `json:"tags,omitempty"`
	Type       *string              `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Domain_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (domain Domain_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (domain *Domain_SpecARM) GetName() string {
	return domain.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (domain *Domain_SpecARM) GetType() string {
	return ""
}

// Deprecated version of DomainProperties. Use v1beta20200601.DomainProperties instead
type DomainPropertiesARM struct {
	Endpoint                   *string                                                   `json:"endpoint,omitempty"`
	InboundIpRules             []InboundIpRuleARM                                        `json:"inboundIpRules,omitempty"`
	InputSchema                *DomainProperties_InputSchema                             `json:"inputSchema,omitempty"`
	InputSchemaMapping         *InputSchemaMappingARM                                    `json:"inputSchemaMapping,omitempty"`
	MetricResourceId           *string                                                   `json:"metricResourceId,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_Domain_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *DomainProperties_ProvisioningState                       `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *DomainProperties_PublicNetworkAccess                     `json:"publicNetworkAccess,omitempty"`
}

// Deprecated version of SystemData. Use v1beta20200601.SystemData instead
type SystemDataARM struct {
	CreatedAt          *string                        `json:"createdAt,omitempty"`
	CreatedBy          *string                        `json:"createdBy,omitempty"`
	CreatedByType      *SystemData_CreatedByType      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                        `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                        `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemData_LastModifiedByType `json:"lastModifiedByType,omitempty"`
}

// Deprecated version of InboundIpRule. Use v1beta20200601.InboundIpRule instead
type InboundIpRuleARM struct {
	Action *InboundIpRule_Action `json:"action,omitempty"`
	IpMask *string               `json:"ipMask,omitempty"`
}

// Deprecated version of InputSchemaMapping. Use v1beta20200601.InputSchemaMapping instead
type InputSchemaMappingARM struct {
	InputSchemaMappingType *InputSchemaMapping_InputSchemaMappingType `json:"inputSchemaMappingType,omitempty"`
}

// Deprecated version of PrivateEndpointConnection_Domain_SubResourceEmbedded. Use v1beta20200601.PrivateEndpointConnection_Domain_SubResourceEmbedded instead
type PrivateEndpointConnection_Domain_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of SystemData_CreatedByType. Use v1beta20200601.SystemData_CreatedByType instead
// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_CreatedByType string

const (
	SystemData_CreatedByType_Application     = SystemData_CreatedByType("Application")
	SystemData_CreatedByType_Key             = SystemData_CreatedByType("Key")
	SystemData_CreatedByType_ManagedIdentity = SystemData_CreatedByType("ManagedIdentity")
	SystemData_CreatedByType_User            = SystemData_CreatedByType("User")
)

// Deprecated version of SystemData_LastModifiedByType. Use v1beta20200601.SystemData_LastModifiedByType instead
// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_LastModifiedByType string

const (
	SystemData_LastModifiedByType_Application     = SystemData_LastModifiedByType("Application")
	SystemData_LastModifiedByType_Key             = SystemData_LastModifiedByType("Key")
	SystemData_LastModifiedByType_ManagedIdentity = SystemData_LastModifiedByType("ManagedIdentity")
	SystemData_LastModifiedByType_User            = SystemData_LastModifiedByType("User")
)