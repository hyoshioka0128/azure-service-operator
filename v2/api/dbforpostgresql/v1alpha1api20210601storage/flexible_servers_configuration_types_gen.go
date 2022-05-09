// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601storage

import (
	"fmt"
	v20210601s "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1beta20210601storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1alpha1api20210601.FlexibleServersConfiguration
// Deprecated version of FlexibleServersConfiguration. Use v1beta20210601.FlexibleServersConfiguration instead
type FlexibleServersConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersConfigurations_Spec `json:"spec,omitempty"`
	Status            Configuration_Status               `json:"status,omitempty"`
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

var _ conversion.Convertible = &FlexibleServersConfiguration{}

// ConvertFrom populates our FlexibleServersConfiguration from the provided hub FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210601s.FlexibleServersConfiguration)
	if !ok {
		return fmt.Errorf("expected dbforpostgresql/v1beta20210601storage/FlexibleServersConfiguration but received %T instead", hub)
	}

	return configuration.AssignPropertiesFromFlexibleServersConfiguration(source)
}

// ConvertTo populates the provided hub FlexibleServersConfiguration from our FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210601s.FlexibleServersConfiguration)
	if !ok {
		return fmt.Errorf("expected dbforpostgresql/v1beta20210601storage/FlexibleServersConfiguration but received %T instead", hub)
	}

	return configuration.AssignPropertiesToFlexibleServersConfiguration(destination)
}

var _ genruntime.KubernetesResource = &FlexibleServersConfiguration{}

// AzureName returns the Azure name of the resource
func (configuration *FlexibleServersConfiguration) AzureName() string {
	return configuration.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (configuration FlexibleServersConfiguration) GetAPIVersion() string {
	return "2021-06-01"
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

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
func (configuration *FlexibleServersConfiguration) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
}

// NewEmptyStatus returns a new empty (blank) status
func (configuration *FlexibleServersConfiguration) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Configuration_Status{}
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
	if st, ok := status.(*Configuration_Status); ok {
		configuration.Status = *st
		return nil
	}

	// Convert status to required version
	var st Configuration_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	configuration.Status = st
	return nil
}

// AssignPropertiesFromFlexibleServersConfiguration populates our FlexibleServersConfiguration from the provided source FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) AssignPropertiesFromFlexibleServersConfiguration(source *v20210601s.FlexibleServersConfiguration) error {

	// ObjectMeta
	configuration.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersConfigurations_Spec
	err := spec.AssignPropertiesFromFlexibleServersConfigurationsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromFlexibleServersConfigurationsSpec() to populate field Spec")
	}
	configuration.Spec = spec

	// Status
	var status Configuration_Status
	err = status.AssignPropertiesFromConfigurationStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromConfigurationStatus() to populate field Status")
	}
	configuration.Status = status

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersConfiguration populates the provided destination FlexibleServersConfiguration from our FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) AssignPropertiesToFlexibleServersConfiguration(destination *v20210601s.FlexibleServersConfiguration) error {

	// ObjectMeta
	destination.ObjectMeta = *configuration.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210601s.FlexibleServersConfigurations_Spec
	err := configuration.Spec.AssignPropertiesToFlexibleServersConfigurationsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToFlexibleServersConfigurationsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210601s.Configuration_Status
	err = configuration.Status.AssignPropertiesToConfigurationStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToConfigurationStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (configuration *FlexibleServersConfiguration) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: configuration.Spec.OriginalVersion,
		Kind:    "FlexibleServersConfiguration",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210601.FlexibleServersConfiguration
// Deprecated version of FlexibleServersConfiguration. Use v1beta20210601.FlexibleServersConfiguration instead
type FlexibleServersConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersConfiguration `json:"items"`
}

// Storage version of v1alpha1api20210601.Configuration_Status
// Deprecated version of Configuration_Status. Use v1beta20210601.Configuration_Status instead
type Configuration_Status struct {
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
	SystemData             *SystemData_Status     `json:"systemData,omitempty"`
	Type                   *string                `json:"type,omitempty"`
	Unit                   *string                `json:"unit,omitempty"`
	Value                  *string                `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Configuration_Status{}

// ConvertStatusFrom populates our Configuration_Status from the provided source
func (configuration *Configuration_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210601s.Configuration_Status)
	if ok {
		// Populate our instance from source
		return configuration.AssignPropertiesFromConfigurationStatus(src)
	}

	// Convert to an intermediate form
	src = &v20210601s.Configuration_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = configuration.AssignPropertiesFromConfigurationStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Configuration_Status
func (configuration *Configuration_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210601s.Configuration_Status)
	if ok {
		// Populate destination from our instance
		return configuration.AssignPropertiesToConfigurationStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20210601s.Configuration_Status{}
	err := configuration.AssignPropertiesToConfigurationStatus(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignPropertiesFromConfigurationStatus populates our Configuration_Status from the provided source Configuration_Status
func (configuration *Configuration_Status) AssignPropertiesFromConfigurationStatus(source *v20210601s.Configuration_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AllowedValues
	configuration.AllowedValues = genruntime.ClonePointerToString(source.AllowedValues)

	// Conditions
	configuration.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DataType
	configuration.DataType = genruntime.ClonePointerToString(source.DataType)

	// DefaultValue
	configuration.DefaultValue = genruntime.ClonePointerToString(source.DefaultValue)

	// Description
	configuration.Description = genruntime.ClonePointerToString(source.Description)

	// DocumentationLink
	configuration.DocumentationLink = genruntime.ClonePointerToString(source.DocumentationLink)

	// Id
	configuration.Id = genruntime.ClonePointerToString(source.Id)

	// IsConfigPendingRestart
	if source.IsConfigPendingRestart != nil {
		isConfigPendingRestart := *source.IsConfigPendingRestart
		configuration.IsConfigPendingRestart = &isConfigPendingRestart
	} else {
		configuration.IsConfigPendingRestart = nil
	}

	// IsDynamicConfig
	if source.IsDynamicConfig != nil {
		isDynamicConfig := *source.IsDynamicConfig
		configuration.IsDynamicConfig = &isDynamicConfig
	} else {
		configuration.IsDynamicConfig = nil
	}

	// IsReadOnly
	if source.IsReadOnly != nil {
		isReadOnly := *source.IsReadOnly
		configuration.IsReadOnly = &isReadOnly
	} else {
		configuration.IsReadOnly = nil
	}

	// Name
	configuration.Name = genruntime.ClonePointerToString(source.Name)

	// Source
	configuration.Source = genruntime.ClonePointerToString(source.Source)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemDataStatus() to populate field SystemData")
		}
		configuration.SystemData = &systemDatum
	} else {
		configuration.SystemData = nil
	}

	// Type
	configuration.Type = genruntime.ClonePointerToString(source.Type)

	// Unit
	configuration.Unit = genruntime.ClonePointerToString(source.Unit)

	// Value
	configuration.Value = genruntime.ClonePointerToString(source.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		configuration.PropertyBag = propertyBag
	} else {
		configuration.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToConfigurationStatus populates the provided destination Configuration_Status from our Configuration_Status
func (configuration *Configuration_Status) AssignPropertiesToConfigurationStatus(destination *v20210601s.Configuration_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(configuration.PropertyBag)

	// AllowedValues
	destination.AllowedValues = genruntime.ClonePointerToString(configuration.AllowedValues)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(configuration.Conditions)

	// DataType
	destination.DataType = genruntime.ClonePointerToString(configuration.DataType)

	// DefaultValue
	destination.DefaultValue = genruntime.ClonePointerToString(configuration.DefaultValue)

	// Description
	destination.Description = genruntime.ClonePointerToString(configuration.Description)

	// DocumentationLink
	destination.DocumentationLink = genruntime.ClonePointerToString(configuration.DocumentationLink)

	// Id
	destination.Id = genruntime.ClonePointerToString(configuration.Id)

	// IsConfigPendingRestart
	if configuration.IsConfigPendingRestart != nil {
		isConfigPendingRestart := *configuration.IsConfigPendingRestart
		destination.IsConfigPendingRestart = &isConfigPendingRestart
	} else {
		destination.IsConfigPendingRestart = nil
	}

	// IsDynamicConfig
	if configuration.IsDynamicConfig != nil {
		isDynamicConfig := *configuration.IsDynamicConfig
		destination.IsDynamicConfig = &isDynamicConfig
	} else {
		destination.IsDynamicConfig = nil
	}

	// IsReadOnly
	if configuration.IsReadOnly != nil {
		isReadOnly := *configuration.IsReadOnly
		destination.IsReadOnly = &isReadOnly
	} else {
		destination.IsReadOnly = nil
	}

	// Name
	destination.Name = genruntime.ClonePointerToString(configuration.Name)

	// Source
	destination.Source = genruntime.ClonePointerToString(configuration.Source)

	// SystemData
	if configuration.SystemData != nil {
		var systemDatum v20210601s.SystemData_Status
		err := configuration.SystemData.AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemDataStatus() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(configuration.Type)

	// Unit
	destination.Unit = genruntime.ClonePointerToString(configuration.Unit)

	// Value
	destination.Value = genruntime.ClonePointerToString(configuration.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210601.FlexibleServersConfigurations_Spec
type FlexibleServersConfigurations_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbforpostgresql.azure.com/FlexibleServer resource
	Owner       *genruntime.KnownResourceReference `group:"dbforpostgresql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Source      *string                            `json:"source,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
	Value       *string                            `json:"value,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServersConfigurations_Spec{}

// ConvertSpecFrom populates our FlexibleServersConfigurations_Spec from the provided source
func (configurations *FlexibleServersConfigurations_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210601s.FlexibleServersConfigurations_Spec)
	if ok {
		// Populate our instance from source
		return configurations.AssignPropertiesFromFlexibleServersConfigurationsSpec(src)
	}

	// Convert to an intermediate form
	src = &v20210601s.FlexibleServersConfigurations_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = configurations.AssignPropertiesFromFlexibleServersConfigurationsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersConfigurations_Spec
func (configurations *FlexibleServersConfigurations_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210601s.FlexibleServersConfigurations_Spec)
	if ok {
		// Populate destination from our instance
		return configurations.AssignPropertiesToFlexibleServersConfigurationsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210601s.FlexibleServersConfigurations_Spec{}
	err := configurations.AssignPropertiesToFlexibleServersConfigurationsSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromFlexibleServersConfigurationsSpec populates our FlexibleServersConfigurations_Spec from the provided source FlexibleServersConfigurations_Spec
func (configurations *FlexibleServersConfigurations_Spec) AssignPropertiesFromFlexibleServersConfigurationsSpec(source *v20210601s.FlexibleServersConfigurations_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	configurations.AzureName = source.AzureName

	// Location
	configurations.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	configurations.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		configurations.Owner = &owner
	} else {
		configurations.Owner = nil
	}

	// Source
	configurations.Source = genruntime.ClonePointerToString(source.Source)

	// Tags
	configurations.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Value
	configurations.Value = genruntime.ClonePointerToString(source.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		configurations.PropertyBag = propertyBag
	} else {
		configurations.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersConfigurationsSpec populates the provided destination FlexibleServersConfigurations_Spec from our FlexibleServersConfigurations_Spec
func (configurations *FlexibleServersConfigurations_Spec) AssignPropertiesToFlexibleServersConfigurationsSpec(destination *v20210601s.FlexibleServersConfigurations_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(configurations.PropertyBag)

	// AzureName
	destination.AzureName = configurations.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(configurations.Location)

	// OriginalVersion
	destination.OriginalVersion = configurations.OriginalVersion

	// Owner
	if configurations.Owner != nil {
		owner := configurations.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Source
	destination.Source = genruntime.ClonePointerToString(configurations.Source)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(configurations.Tags)

	// Value
	destination.Value = genruntime.ClonePointerToString(configurations.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&FlexibleServersConfiguration{}, &FlexibleServersConfigurationList{})
}
