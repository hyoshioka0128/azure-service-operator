// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"fmt"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
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
// Storage version of v1alpha1api20210515.SqlDatabaseContainerUserDefinedFunction
// Deprecated version of SqlDatabaseContainerUserDefinedFunction. Use v1beta20210515.SqlDatabaseContainerUserDefinedFunction instead
type SqlDatabaseContainerUserDefinedFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerUserDefinedFunction{}

// GetConditions returns the conditions of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetConditions() conditions.Conditions {
	return function.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (function *SqlDatabaseContainerUserDefinedFunction) SetConditions(conditions conditions.Conditions) {
	function.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseContainerUserDefinedFunction{}

// ConvertFrom populates our SqlDatabaseContainerUserDefinedFunction from the provided hub SqlDatabaseContainerUserDefinedFunction
func (function *SqlDatabaseContainerUserDefinedFunction) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210515s.SqlDatabaseContainerUserDefinedFunction)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabaseContainerUserDefinedFunction but received %T instead", hub)
	}

	return function.AssignPropertiesFromSqlDatabaseContainerUserDefinedFunction(source)
}

// ConvertTo populates the provided hub SqlDatabaseContainerUserDefinedFunction from our SqlDatabaseContainerUserDefinedFunction
func (function *SqlDatabaseContainerUserDefinedFunction) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.SqlDatabaseContainerUserDefinedFunction)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabaseContainerUserDefinedFunction but received %T instead", hub)
	}

	return function.AssignPropertiesToSqlDatabaseContainerUserDefinedFunction(destination)
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerUserDefinedFunction{}

// AzureName returns the Azure name of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) AzureName() string {
	return function.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (function SqlDatabaseContainerUserDefinedFunction) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetSpec() genruntime.ConvertibleSpec {
	return &function.Spec
}

// GetStatus returns the status of this resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetStatus() genruntime.ConvertibleStatus {
	return &function.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (function *SqlDatabaseContainerUserDefinedFunction) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (function *SqlDatabaseContainerUserDefinedFunction) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (function *SqlDatabaseContainerUserDefinedFunction) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(function.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  function.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (function *SqlDatabaseContainerUserDefinedFunction) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS); ok {
		function.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	function.Status = st
	return nil
}

// AssignPropertiesFromSqlDatabaseContainerUserDefinedFunction populates our SqlDatabaseContainerUserDefinedFunction from the provided source SqlDatabaseContainerUserDefinedFunction
func (function *SqlDatabaseContainerUserDefinedFunction) AssignPropertiesFromSqlDatabaseContainerUserDefinedFunction(source *v20210515s.SqlDatabaseContainerUserDefinedFunction) error {

	// ObjectMeta
	function.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec() to populate field Spec")
	}
	function.Spec = spec

	// Status
	var status DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS
	err = status.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS() to populate field Status")
	}
	function.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseContainerUserDefinedFunction populates the provided destination SqlDatabaseContainerUserDefinedFunction from our SqlDatabaseContainerUserDefinedFunction
func (function *SqlDatabaseContainerUserDefinedFunction) AssignPropertiesToSqlDatabaseContainerUserDefinedFunction(destination *v20210515s.SqlDatabaseContainerUserDefinedFunction) error {

	// ObjectMeta
	destination.ObjectMeta = *function.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec
	err := function.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS
	err = function.Status.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (function *SqlDatabaseContainerUserDefinedFunction) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: function.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainerUserDefinedFunction",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210515.SqlDatabaseContainerUserDefinedFunction
// Deprecated version of SqlDatabaseContainerUserDefinedFunction. Use v1beta20210515.SqlDatabaseContainerUserDefinedFunction instead
type SqlDatabaseContainerUserDefinedFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerUserDefinedFunction `json:"items"`
}

// Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS
// Deprecated version of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS. Use v1beta20210515.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS instead
type DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS struct {
	Conditions  []conditions.Condition                               `json:"conditions,omitempty"`
	Id          *string                                              `json:"id,omitempty"`
	Location    *string                                              `json:"location,omitempty"`
	Name        *string                                              `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                               `json:"$propertyBag,omitempty"`
	Resource    *SqlUserDefinedFunctionGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                                    `json:"tags,omitempty"`
	Type        *string                                              `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS{}

// ConvertStatusFrom populates our DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS from the provided source
func (function *DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS)
	if ok {
		// Populate our instance from source
		return function.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = function.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS
func (function *DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS)
	if ok {
		// Populate destination from our instance
		return function.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS{}
	err := function.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS populates our DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS from the provided source DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS
func (function *DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS) AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(source *v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	function.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	function.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	function.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	function.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource SqlUserDefinedFunctionGetProperties_Resource_STATUS
		err := resource.AssignPropertiesFromSqlUserDefinedFunctionGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlUserDefinedFunctionGetProperties_Resource_STATUS() to populate field Resource")
		}
		function.Resource = &resource
	} else {
		function.Resource = nil
	}

	// Tags
	function.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	function.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		function.PropertyBag = propertyBag
	} else {
		function.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS populates the provided destination DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS from our DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS
func (function *DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS) AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(destination *v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(function.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(function.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(function.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(function.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(function.Name)

	// Resource
	if function.Resource != nil {
		var resource v20210515s.SqlUserDefinedFunctionGetProperties_Resource_STATUS
		err := function.Resource.AssignPropertiesToSqlUserDefinedFunctionGetProperties_Resource_STATUS(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlUserDefinedFunctionGetProperties_Resource_STATUS() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(function.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(function.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec
type DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *SqlUserDefinedFunctionResource    `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec from the provided source
func (function *DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec)
	if ok {
		// Populate our instance from source
		return function.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = function.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec
func (function *DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec)
	if ok {
		// Populate destination from our instance
		return function.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec{}
	err := function.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec populates our DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec from the provided source DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec
func (function *DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(source *v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	function.AzureName = source.AzureName

	// Location
	function.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignPropertiesFromCreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptions() to populate field Options")
		}
		function.Options = &option
	} else {
		function.Options = nil
	}

	// OriginalVersion
	function.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		function.Owner = &owner
	} else {
		function.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlUserDefinedFunctionResource
		err := resource.AssignPropertiesFromSqlUserDefinedFunctionResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlUserDefinedFunctionResource() to populate field Resource")
		}
		function.Resource = &resource
	} else {
		function.Resource = nil
	}

	// Tags
	function.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		function.PropertyBag = propertyBag
	} else {
		function.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec populates the provided destination DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec from our DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec
func (function *DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec) AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(destination *v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(function.PropertyBag)

	// AzureName
	destination.AzureName = function.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(function.Location)

	// Options
	if function.Options != nil {
		var option v20210515s.CreateUpdateOptions
		err := function.Options.AssignPropertiesToCreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptions() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = function.OriginalVersion

	// Owner
	if function.Owner != nil {
		owner := function.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if function.Resource != nil {
		var resource v20210515s.SqlUserDefinedFunctionResource
		err := function.Resource.AssignPropertiesToSqlUserDefinedFunctionResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlUserDefinedFunctionResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(function.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.SqlUserDefinedFunctionGetProperties_Resource_STATUS
// Deprecated version of SqlUserDefinedFunctionGetProperties_Resource_STATUS. Use v1beta20210515.SqlUserDefinedFunctionGetProperties_Resource_STATUS instead
type SqlUserDefinedFunctionGetProperties_Resource_STATUS struct {
	Body        *string                `json:"body,omitempty"`
	Etag        *string                `json:"_etag,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid         *string                `json:"_rid,omitempty"`
	Ts          *float64               `json:"_ts,omitempty"`
}

// AssignPropertiesFromSqlUserDefinedFunctionGetProperties_Resource_STATUS populates our SqlUserDefinedFunctionGetProperties_Resource_STATUS from the provided source SqlUserDefinedFunctionGetProperties_Resource_STATUS
func (resource *SqlUserDefinedFunctionGetProperties_Resource_STATUS) AssignPropertiesFromSqlUserDefinedFunctionGetProperties_Resource_STATUS(source *v20210515s.SqlUserDefinedFunctionGetProperties_Resource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Etag
	resource.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// Rid
	resource.Rid = genruntime.ClonePointerToString(source.Rid)

	// Ts
	if source.Ts != nil {
		t := *source.Ts
		resource.Ts = &t
	} else {
		resource.Ts = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSqlUserDefinedFunctionGetProperties_Resource_STATUS populates the provided destination SqlUserDefinedFunctionGetProperties_Resource_STATUS from our SqlUserDefinedFunctionGetProperties_Resource_STATUS
func (resource *SqlUserDefinedFunctionGetProperties_Resource_STATUS) AssignPropertiesToSqlUserDefinedFunctionGetProperties_Resource_STATUS(destination *v20210515s.SqlUserDefinedFunctionGetProperties_Resource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(resource.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// Rid
	destination.Rid = genruntime.ClonePointerToString(resource.Rid)

	// Ts
	if resource.Ts != nil {
		t := *resource.Ts
		destination.Ts = &t
	} else {
		destination.Ts = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.SqlUserDefinedFunctionResource
// Deprecated version of SqlUserDefinedFunctionResource. Use v1beta20210515.SqlUserDefinedFunctionResource instead
type SqlUserDefinedFunctionResource struct {
	Body        *string                `json:"body,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromSqlUserDefinedFunctionResource populates our SqlUserDefinedFunctionResource from the provided source SqlUserDefinedFunctionResource
func (resource *SqlUserDefinedFunctionResource) AssignPropertiesFromSqlUserDefinedFunctionResource(source *v20210515s.SqlUserDefinedFunctionResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSqlUserDefinedFunctionResource populates the provided destination SqlUserDefinedFunctionResource from our SqlUserDefinedFunctionResource
func (resource *SqlUserDefinedFunctionResource) AssignPropertiesToSqlUserDefinedFunctionResource(destination *v20210515s.SqlUserDefinedFunctionResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

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
	SchemeBuilder.Register(&SqlDatabaseContainerUserDefinedFunction{}, &SqlDatabaseContainerUserDefinedFunctionList{})
}
