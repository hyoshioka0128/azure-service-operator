// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources=flexibleserversconfigurations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources={flexibleserversconfigurations/status,flexibleserversconfigurations/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210601.FlexibleServersConfiguration
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/postgresql.json
type FlexibleServersConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersConfiguration_Spec   `json:"spec,omitempty"`
	Status            FlexibleServersConfiguration_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersConfiguration{}

// GetConditions returns the conditions of the resource
func (configuration *FlexibleServersConfiguration) GetConditions() conditions.Conditions {
	return configuration.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (configuration *FlexibleServersConfiguration) SetConditions(conditions conditions.Conditions) {
	configuration.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &FlexibleServersConfiguration{}

// AzureName returns the Azure name of the resource
func (configuration *FlexibleServersConfiguration) AzureName() string {
	return configuration.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (configuration FlexibleServersConfiguration) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (configuration *FlexibleServersConfiguration) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (configuration *FlexibleServersConfiguration) GetSpec() genruntime.ConvertibleSpec {
	return &configuration.Spec
}

// GetStatus returns the status of this resource
func (configuration *FlexibleServersConfiguration) GetStatus() genruntime.ConvertibleStatus {
	return &configuration.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (configuration *FlexibleServersConfiguration) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (configuration *FlexibleServersConfiguration) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServersConfiguration_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (configuration *FlexibleServersConfiguration) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(configuration.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  configuration.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (configuration *FlexibleServersConfiguration) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServersConfiguration_STATUS); ok {
		configuration.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServersConfiguration_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	configuration.Status = st
	return nil
}

// Hub marks that this FlexibleServersConfiguration is the hub type for conversion
func (configuration *FlexibleServersConfiguration) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (configuration *FlexibleServersConfiguration) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: configuration.Spec.OriginalVersion,
		Kind:    "FlexibleServersConfiguration",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210601.FlexibleServersConfiguration
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/postgresql.json
type FlexibleServersConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersConfiguration `json:"items"`
}

// Storage version of v1beta20210601.FlexibleServersConfiguration_STATUS
type FlexibleServersConfiguration_STATUS struct {
	AllowedValues          *string                `json:"allowedValues,omitempty"`
	Conditions             []conditions.Condition `json:"conditions,omitempty"`
	DataType               *string                `json:"dataType,omitempty"`
	DefaultValue           *string                `json:"defaultValue,omitempty"`
	Description            *string                `json:"description,omitempty"`
	DocumentationLink      *string                `json:"documentationLink,omitempty"`
	Id                     *string                `json:"id,omitempty"`
	IsConfigPendingRestart *bool                  `json:"isConfigPendingRestart,omitempty"`
	IsDynamicConfig        *bool                  `json:"isDynamicConfig,omitempty"`
	IsReadOnly             *bool                  `json:"isReadOnly,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Source                 *string                `json:"source,omitempty"`
	SystemData             *SystemData_STATUS     `json:"systemData,omitempty"`
	Type                   *string                `json:"type,omitempty"`
	Unit                   *string                `json:"unit,omitempty"`
	Value                  *string                `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServersConfiguration_STATUS{}

// ConvertStatusFrom populates our FlexibleServersConfiguration_STATUS from the provided source
func (configuration *FlexibleServersConfiguration_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == configuration {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(configuration)
}

// ConvertStatusTo populates the provided destination from our FlexibleServersConfiguration_STATUS
func (configuration *FlexibleServersConfiguration_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == configuration {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(configuration)
}

// Storage version of v1beta20210601.FlexibleServersConfiguration_Spec
type FlexibleServersConfiguration_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string `json:"azureName,omitempty"`
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Source      *string                            `json:"source,omitempty"`
	Value       *string                            `json:"value,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServersConfiguration_Spec{}

// ConvertSpecFrom populates our FlexibleServersConfiguration_Spec from the provided source
func (configuration *FlexibleServersConfiguration_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == configuration {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(configuration)
}

// ConvertSpecTo populates the provided destination from our FlexibleServersConfiguration_Spec
func (configuration *FlexibleServersConfiguration_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == configuration {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(configuration)
}

func init() {
	SchemeBuilder.Register(&FlexibleServersConfiguration{}, &FlexibleServersConfigurationList{})
}
