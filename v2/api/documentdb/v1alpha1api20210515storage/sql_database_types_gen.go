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
// Storage version of v1alpha1api20210515.SqlDatabase
// Deprecated version of SqlDatabase. Use v1beta20210515.SqlDatabase instead
type SqlDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabase_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccountsSqlDatabase_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabase{}

// GetConditions returns the conditions of the resource
func (database *SqlDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *SqlDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabase{}

// ConvertFrom populates our SqlDatabase from the provided hub SqlDatabase
func (database *SqlDatabase) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210515s.SqlDatabase)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesFromSqlDatabase(source)
}

// ConvertTo populates the provided hub SqlDatabase from our SqlDatabase
func (database *SqlDatabase) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.SqlDatabase)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesToSqlDatabase(destination)
}

var _ genruntime.KubernetesResource = &SqlDatabase{}

// AzureName returns the Azure name of the resource
func (database *SqlDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (database SqlDatabase) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (database *SqlDatabase) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (database *SqlDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *SqlDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (database *SqlDatabase) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (database *SqlDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccountsSqlDatabase_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (database *SqlDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *SqlDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccountsSqlDatabase_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccountsSqlDatabase_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// AssignPropertiesFromSqlDatabase populates our SqlDatabase from the provided source SqlDatabase
func (database *SqlDatabase) AssignPropertiesFromSqlDatabase(source *v20210515s.SqlDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabase_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabase_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabase_Spec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status DatabaseAccountsSqlDatabase_STATUS
	err = status.AssignPropertiesFromDatabaseAccountsSqlDatabase_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabase_STATUS() to populate field Status")
	}
	database.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabase populates the provided destination SqlDatabase from our SqlDatabase
func (database *SqlDatabase) AssignPropertiesToSqlDatabase(destination *v20210515s.SqlDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccountsSqlDatabase_Spec
	err := database.Spec.AssignPropertiesToDatabaseAccountsSqlDatabase_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabase_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.DatabaseAccountsSqlDatabase_STATUS
	err = database.Status.AssignPropertiesToDatabaseAccountsSqlDatabase_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabase_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *SqlDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion,
		Kind:    "SqlDatabase",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210515.SqlDatabase
// Deprecated version of SqlDatabase. Use v1beta20210515.SqlDatabase instead
type SqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabase `json:"items"`
}

// Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabase_STATUS
// Deprecated version of DatabaseAccountsSqlDatabase_STATUS. Use v1beta20210515.DatabaseAccountsSqlDatabase_STATUS instead
type DatabaseAccountsSqlDatabase_STATUS struct {
	Conditions  []conditions.Condition                    `json:"conditions,omitempty"`
	Id          *string                                   `json:"id,omitempty"`
	Location    *string                                   `json:"location,omitempty"`
	Name        *string                                   `json:"name,omitempty"`
	Options     *OptionsResource_STATUS                   `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	Resource    *SqlDatabaseGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                         `json:"tags,omitempty"`
	Type        *string                                   `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccountsSqlDatabase_STATUS{}

// ConvertStatusFrom populates our DatabaseAccountsSqlDatabase_STATUS from the provided source
func (database *DatabaseAccountsSqlDatabase_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210515s.DatabaseAccountsSqlDatabase_STATUS)
	if ok {
		// Populate our instance from source
		return database.AssignPropertiesFromDatabaseAccountsSqlDatabase_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccountsSqlDatabase_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = database.AssignPropertiesFromDatabaseAccountsSqlDatabase_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DatabaseAccountsSqlDatabase_STATUS
func (database *DatabaseAccountsSqlDatabase_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210515s.DatabaseAccountsSqlDatabase_STATUS)
	if ok {
		// Populate destination from our instance
		return database.AssignPropertiesToDatabaseAccountsSqlDatabase_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccountsSqlDatabase_STATUS{}
	err := database.AssignPropertiesToDatabaseAccountsSqlDatabase_STATUS(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabase_STATUS populates our DatabaseAccountsSqlDatabase_STATUS from the provided source DatabaseAccountsSqlDatabase_STATUS
func (database *DatabaseAccountsSqlDatabase_STATUS) AssignPropertiesFromDatabaseAccountsSqlDatabase_STATUS(source *v20210515s.DatabaseAccountsSqlDatabase_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	database.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	database.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	database.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	database.Name = genruntime.ClonePointerToString(source.Name)

	// Options
	if source.Options != nil {
		var option OptionsResource_STATUS
		err := option.AssignPropertiesFromOptionsResource_STATUS(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromOptionsResource_STATUS() to populate field Options")
		}
		database.Options = &option
	} else {
		database.Options = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlDatabaseGetProperties_Resource_STATUS
		err := resource.AssignPropertiesFromSqlDatabaseGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlDatabaseGetProperties_Resource_STATUS() to populate field Resource")
		}
		database.Resource = &resource
	} else {
		database.Resource = nil
	}

	// Tags
	database.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	database.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		database.PropertyBag = propertyBag
	} else {
		database.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabase_STATUS populates the provided destination DatabaseAccountsSqlDatabase_STATUS from our DatabaseAccountsSqlDatabase_STATUS
func (database *DatabaseAccountsSqlDatabase_STATUS) AssignPropertiesToDatabaseAccountsSqlDatabase_STATUS(destination *v20210515s.DatabaseAccountsSqlDatabase_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(database.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(database.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(database.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(database.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(database.Name)

	// Options
	if database.Options != nil {
		var option v20210515s.OptionsResource_STATUS
		err := database.Options.AssignPropertiesToOptionsResource_STATUS(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToOptionsResource_STATUS() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// Resource
	if database.Resource != nil {
		var resource v20210515s.SqlDatabaseGetProperties_Resource_STATUS
		err := database.Resource.AssignPropertiesToSqlDatabaseGetProperties_Resource_STATUS(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlDatabaseGetProperties_Resource_STATUS() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(database.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(database.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabase_Spec
type DatabaseAccountsSqlDatabase_Spec struct {
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
	Resource    *SqlDatabaseResource               `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabase_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabase_Spec from the provided source
func (database *DatabaseAccountsSqlDatabase_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccountsSqlDatabase_Spec)
	if ok {
		// Populate our instance from source
		return database.AssignPropertiesFromDatabaseAccountsSqlDatabase_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccountsSqlDatabase_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = database.AssignPropertiesFromDatabaseAccountsSqlDatabase_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabase_Spec
func (database *DatabaseAccountsSqlDatabase_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccountsSqlDatabase_Spec)
	if ok {
		// Populate destination from our instance
		return database.AssignPropertiesToDatabaseAccountsSqlDatabase_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccountsSqlDatabase_Spec{}
	err := database.AssignPropertiesToDatabaseAccountsSqlDatabase_Spec(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabase_Spec populates our DatabaseAccountsSqlDatabase_Spec from the provided source DatabaseAccountsSqlDatabase_Spec
func (database *DatabaseAccountsSqlDatabase_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabase_Spec(source *v20210515s.DatabaseAccountsSqlDatabase_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	database.AzureName = source.AzureName

	// Location
	database.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignPropertiesFromCreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptions() to populate field Options")
		}
		database.Options = &option
	} else {
		database.Options = nil
	}

	// OriginalVersion
	database.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		database.Owner = &owner
	} else {
		database.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlDatabaseResource
		err := resource.AssignPropertiesFromSqlDatabaseResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlDatabaseResource() to populate field Resource")
		}
		database.Resource = &resource
	} else {
		database.Resource = nil
	}

	// Tags
	database.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		database.PropertyBag = propertyBag
	} else {
		database.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabase_Spec populates the provided destination DatabaseAccountsSqlDatabase_Spec from our DatabaseAccountsSqlDatabase_Spec
func (database *DatabaseAccountsSqlDatabase_Spec) AssignPropertiesToDatabaseAccountsSqlDatabase_Spec(destination *v20210515s.DatabaseAccountsSqlDatabase_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(database.PropertyBag)

	// AzureName
	destination.AzureName = database.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(database.Location)

	// Options
	if database.Options != nil {
		var option v20210515s.CreateUpdateOptions
		err := database.Options.AssignPropertiesToCreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptions() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = database.OriginalVersion

	// Owner
	if database.Owner != nil {
		owner := database.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if database.Resource != nil {
		var resource v20210515s.SqlDatabaseResource
		err := database.Resource.AssignPropertiesToSqlDatabaseResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlDatabaseResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(database.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.SqlDatabaseGetProperties_Resource_STATUS
// Deprecated version of SqlDatabaseGetProperties_Resource_STATUS. Use v1beta20210515.SqlDatabaseGetProperties_Resource_STATUS instead
type SqlDatabaseGetProperties_Resource_STATUS struct {
	Colls       *string                `json:"_colls,omitempty"`
	Etag        *string                `json:"_etag,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid         *string                `json:"_rid,omitempty"`
	Ts          *float64               `json:"_ts,omitempty"`
	Users       *string                `json:"_users,omitempty"`
}

// AssignPropertiesFromSqlDatabaseGetProperties_Resource_STATUS populates our SqlDatabaseGetProperties_Resource_STATUS from the provided source SqlDatabaseGetProperties_Resource_STATUS
func (resource *SqlDatabaseGetProperties_Resource_STATUS) AssignPropertiesFromSqlDatabaseGetProperties_Resource_STATUS(source *v20210515s.SqlDatabaseGetProperties_Resource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Colls
	resource.Colls = genruntime.ClonePointerToString(source.Colls)

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

	// Users
	resource.Users = genruntime.ClonePointerToString(source.Users)

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseGetProperties_Resource_STATUS populates the provided destination SqlDatabaseGetProperties_Resource_STATUS from our SqlDatabaseGetProperties_Resource_STATUS
func (resource *SqlDatabaseGetProperties_Resource_STATUS) AssignPropertiesToSqlDatabaseGetProperties_Resource_STATUS(destination *v20210515s.SqlDatabaseGetProperties_Resource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Colls
	destination.Colls = genruntime.ClonePointerToString(resource.Colls)

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

	// Users
	destination.Users = genruntime.ClonePointerToString(resource.Users)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.SqlDatabaseResource
// Deprecated version of SqlDatabaseResource. Use v1beta20210515.SqlDatabaseResource instead
type SqlDatabaseResource struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromSqlDatabaseResource populates our SqlDatabaseResource from the provided source SqlDatabaseResource
func (resource *SqlDatabaseResource) AssignPropertiesFromSqlDatabaseResource(source *v20210515s.SqlDatabaseResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

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

// AssignPropertiesToSqlDatabaseResource populates the provided destination SqlDatabaseResource from our SqlDatabaseResource
func (resource *SqlDatabaseResource) AssignPropertiesToSqlDatabaseResource(destination *v20210515s.SqlDatabaseResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

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
	SchemeBuilder.Register(&SqlDatabase{}, &SqlDatabaseList{})
}
