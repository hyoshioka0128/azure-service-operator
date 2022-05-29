// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=eventgrid.azure.com,resources=domains,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventgrid.azure.com,resources={domains/status,domains/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20200601.Domain
// Generator information:
// - Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Domain_Spec   `json:"spec,omitempty"`
	Status            Domain_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Domain{}

// GetConditions returns the conditions of the resource
func (domain *Domain) GetConditions() conditions.Conditions {
	return domain.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (domain *Domain) SetConditions(conditions conditions.Conditions) {
	domain.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Domain{}

// AzureName returns the Azure name of the resource
func (domain *Domain) AzureName() string {
	return domain.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (domain Domain) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (domain *Domain) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (domain *Domain) GetSpec() genruntime.ConvertibleSpec {
	return &domain.Spec
}

// GetStatus returns the status of this resource
func (domain *Domain) GetStatus() genruntime.ConvertibleStatus {
	return &domain.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (domain *Domain) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (domain *Domain) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Domain_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (domain *Domain) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(domain.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  domain.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (domain *Domain) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Domain_STATUS); ok {
		domain.Status = *st
		return nil
	}

	// Convert status to required version
	var st Domain_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	domain.Status = st
	return nil
}

// Hub marks that this Domain is the hub type for conversion
func (domain *Domain) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (domain *Domain) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: domain.Spec.OriginalVersion,
		Kind:    "Domain",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20200601.Domain
// Generator information:
// - Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Storage version of v1beta20200601.APIVersion
// +kubebuilder:validation:Enum={"2020-06-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-06-01")

// Storage version of v1beta20200601.Domain_STATUS
type Domain_STATUS struct {
	Conditions                 []conditions.Condition                                        `json:"conditions,omitempty"`
	Endpoint                   *string                                                       `json:"endpoint,omitempty"`
	Id                         *string                                                       `json:"id,omitempty"`
	InboundIpRules             []InboundIpRule_STATUS                                        `json:"inboundIpRules,omitempty"`
	InputSchema                *string                                                       `json:"inputSchema,omitempty"`
	InputSchemaMapping         *InputSchemaMapping_STATUS                                    `json:"inputSchemaMapping,omitempty"`
	Location                   *string                                                       `json:"location,omitempty"`
	MetricResourceId           *string                                                       `json:"metricResourceId,omitempty"`
	Name                       *string                                                       `json:"name,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_Domain_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                        `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                                       `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                                                       `json:"publicNetworkAccess,omitempty"`
	SystemData                 *SystemData_STATUS                                            `json:"systemData,omitempty"`
	Tags                       map[string]string                                             `json:"tags,omitempty"`
	Type                       *string                                                       `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Domain_STATUS{}

// ConvertStatusFrom populates our Domain_STATUS from the provided source
func (domain *Domain_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == domain {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(domain)
}

// ConvertStatusTo populates the provided destination from our Domain_STATUS
func (domain *Domain_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == domain {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(domain)
}

// Storage version of v1beta20200601.Domain_Spec
type Domain_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName          string              `json:"azureName,omitempty"`
	Endpoint           *string             `json:"endpoint,omitempty"`
	Id                 *string             `json:"id,omitempty"`
	InboundIpRules     []InboundIpRule     `json:"inboundIpRules,omitempty"`
	InputSchema        *string             `json:"inputSchema,omitempty"`
	InputSchemaMapping *InputSchemaMapping `json:"inputSchemaMapping,omitempty"`
	Location           *string             `json:"location,omitempty"`
	MetricResourceId   *string             `json:"metricResourceId,omitempty"`
	OriginalVersion    string              `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                      *genruntime.KnownResourceReference                     `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PrivateEndpointConnections []PrivateEndpointConnection_Domain_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                 `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                                `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                                                `json:"publicNetworkAccess,omitempty"`
	SystemData                 *SystemData                                            `json:"systemData,omitempty"`
	Tags                       map[string]string                                      `json:"tags,omitempty"`
	Type                       *string                                                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Domain_Spec{}

// ConvertSpecFrom populates our Domain_Spec from the provided source
func (domain *Domain_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == domain {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(domain)
}

// ConvertSpecTo populates the provided destination from our Domain_Spec
func (domain *Domain_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == domain {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(domain)
}

// Storage version of v1beta20200601.InboundIpRule
type InboundIpRule struct {
	Action      *string                `json:"action,omitempty"`
	IpMask      *string                `json:"ipMask,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.InboundIpRule_STATUS
type InboundIpRule_STATUS struct {
	Action      *string                `json:"action,omitempty"`
	IpMask      *string                `json:"ipMask,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.InputSchemaMapping
type InputSchemaMapping struct {
	InputSchemaMappingType *string                `json:"inputSchemaMappingType,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.InputSchemaMapping_STATUS
type InputSchemaMapping_STATUS struct {
	InputSchemaMappingType *string                `json:"inputSchemaMappingType,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.PrivateEndpointConnection_Domain_SubResourceEmbedded
type PrivateEndpointConnection_Domain_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.PrivateEndpointConnection_STATUS_Domain_SubResourceEmbedded
type PrivateEndpointConnection_STATUS_Domain_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.SystemData
type SystemData struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.SystemData_STATUS
type SystemData_STATUS struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
