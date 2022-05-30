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
// Storage version of v1alpha1api20210515.MongodbDatabase
// Deprecated version of MongodbDatabase. Use v1beta20210515.MongodbDatabase instead
type MongodbDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsMongodbDatabase_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccountsMongodbDatabase_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabase{}

// GetConditions returns the conditions of the resource
func (database *MongodbDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *MongodbDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ conversion.Convertible = &MongodbDatabase{}

// ConvertFrom populates our MongodbDatabase from the provided hub MongodbDatabase
func (database *MongodbDatabase) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210515s.MongodbDatabase)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/MongodbDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesFromMongodbDatabase(source)
}

// ConvertTo populates the provided hub MongodbDatabase from our MongodbDatabase
func (database *MongodbDatabase) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.MongodbDatabase)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/MongodbDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesToMongodbDatabase(destination)
}

var _ genruntime.KubernetesResource = &MongodbDatabase{}

// AzureName returns the Azure name of the resource
func (database *MongodbDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (database MongodbDatabase) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (database *MongodbDatabase) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (database *MongodbDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *MongodbDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (database *MongodbDatabase) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (database *MongodbDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccountsMongodbDatabase_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (database *MongodbDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *MongodbDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccountsMongodbDatabase_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccountsMongodbDatabase_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// AssignPropertiesFromMongodbDatabase populates our MongodbDatabase from the provided source MongodbDatabase
func (database *MongodbDatabase) AssignPropertiesFromMongodbDatabase(source *v20210515s.MongodbDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsMongodbDatabase_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsMongodbDatabase_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsMongodbDatabase_Spec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status DatabaseAccountsMongodbDatabase_STATUS
	err = status.AssignPropertiesFromDatabaseAccountsMongodbDatabase_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsMongodbDatabase_STATUS() to populate field Status")
	}
	database.Status = status

	// No error
	return nil
}

// AssignPropertiesToMongodbDatabase populates the provided destination MongodbDatabase from our MongodbDatabase
func (database *MongodbDatabase) AssignPropertiesToMongodbDatabase(destination *v20210515s.MongodbDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccountsMongodbDatabase_Spec
	err := database.Spec.AssignPropertiesToDatabaseAccountsMongodbDatabase_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsMongodbDatabase_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.DatabaseAccountsMongodbDatabase_STATUS
	err = database.Status.AssignPropertiesToDatabaseAccountsMongodbDatabase_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsMongodbDatabase_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *MongodbDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion,
		Kind:    "MongodbDatabase",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210515.MongodbDatabase
// Deprecated version of MongodbDatabase. Use v1beta20210515.MongodbDatabase instead
type MongodbDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabase `json:"items"`
}

// Storage version of v1alpha1api20210515.DatabaseAccountsMongodbDatabase_STATUS
// Deprecated version of DatabaseAccountsMongodbDatabase_STATUS. Use v1beta20210515.DatabaseAccountsMongodbDatabase_STATUS instead
type DatabaseAccountsMongodbDatabase_STATUS struct {
	Conditions  []conditions.Condition                        `json:"conditions,omitempty"`
	Id          *string                                       `json:"id,omitempty"`
	Location    *string                                       `json:"location,omitempty"`
	Name        *string                                       `json:"name,omitempty"`
	Options     *OptionsResource_STATUS                       `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	Resource    *MongoDBDatabaseGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                             `json:"tags,omitempty"`
	Type        *string                                       `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccountsMongodbDatabase_STATUS{}

// ConvertStatusFrom populates our DatabaseAccountsMongodbDatabase_STATUS from the provided source
func (database *DatabaseAccountsMongodbDatabase_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210515s.DatabaseAccountsMongodbDatabase_STATUS)
	if ok {
		// Populate our instance from source
		return database.AssignPropertiesFromDatabaseAccountsMongodbDatabase_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccountsMongodbDatabase_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = database.AssignPropertiesFromDatabaseAccountsMongodbDatabase_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DatabaseAccountsMongodbDatabase_STATUS
func (database *DatabaseAccountsMongodbDatabase_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210515s.DatabaseAccountsMongodbDatabase_STATUS)
	if ok {
		// Populate destination from our instance
		return database.AssignPropertiesToDatabaseAccountsMongodbDatabase_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccountsMongodbDatabase_STATUS{}
	err := database.AssignPropertiesToDatabaseAccountsMongodbDatabase_STATUS(dst)
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

// AssignPropertiesFromDatabaseAccountsMongodbDatabase_STATUS populates our DatabaseAccountsMongodbDatabase_STATUS from the provided source DatabaseAccountsMongodbDatabase_STATUS
func (database *DatabaseAccountsMongodbDatabase_STATUS) AssignPropertiesFromDatabaseAccountsMongodbDatabase_STATUS(source *v20210515s.DatabaseAccountsMongodbDatabase_STATUS) error {
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
		var resource MongoDBDatabaseGetProperties_Resource_STATUS
		err := resource.AssignPropertiesFromMongoDBDatabaseGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromMongoDBDatabaseGetProperties_Resource_STATUS() to populate field Resource")
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

// AssignPropertiesToDatabaseAccountsMongodbDatabase_STATUS populates the provided destination DatabaseAccountsMongodbDatabase_STATUS from our DatabaseAccountsMongodbDatabase_STATUS
func (database *DatabaseAccountsMongodbDatabase_STATUS) AssignPropertiesToDatabaseAccountsMongodbDatabase_STATUS(destination *v20210515s.DatabaseAccountsMongodbDatabase_STATUS) error {
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
		var resource v20210515s.MongoDBDatabaseGetProperties_Resource_STATUS
		err := database.Resource.AssignPropertiesToMongoDBDatabaseGetProperties_Resource_STATUS(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToMongoDBDatabaseGetProperties_Resource_STATUS() to populate field Resource")
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

// Storage version of v1alpha1api20210515.DatabaseAccountsMongodbDatabase_Spec
type DatabaseAccountsMongodbDatabase_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	Id              *string              `json:"id,omitempty"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *MongoDBDatabaseResource           `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
	Type        *string                            `json:"type,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsMongodbDatabase_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsMongodbDatabase_Spec from the provided source
func (database *DatabaseAccountsMongodbDatabase_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccountsMongodbDatabase_Spec)
	if ok {
		// Populate our instance from source
		return database.AssignPropertiesFromDatabaseAccountsMongodbDatabase_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccountsMongodbDatabase_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = database.AssignPropertiesFromDatabaseAccountsMongodbDatabase_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsMongodbDatabase_Spec
func (database *DatabaseAccountsMongodbDatabase_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccountsMongodbDatabase_Spec)
	if ok {
		// Populate destination from our instance
		return database.AssignPropertiesToDatabaseAccountsMongodbDatabase_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccountsMongodbDatabase_Spec{}
	err := database.AssignPropertiesToDatabaseAccountsMongodbDatabase_Spec(dst)
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

// AssignPropertiesFromDatabaseAccountsMongodbDatabase_Spec populates our DatabaseAccountsMongodbDatabase_Spec from the provided source DatabaseAccountsMongodbDatabase_Spec
func (database *DatabaseAccountsMongodbDatabase_Spec) AssignPropertiesFromDatabaseAccountsMongodbDatabase_Spec(source *v20210515s.DatabaseAccountsMongodbDatabase_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	database.AzureName = source.AzureName

	// Id
	database.Id = genruntime.ClonePointerToString(source.Id)

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
		var resource MongoDBDatabaseResource
		err := resource.AssignPropertiesFromMongoDBDatabaseResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromMongoDBDatabaseResource() to populate field Resource")
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

// AssignPropertiesToDatabaseAccountsMongodbDatabase_Spec populates the provided destination DatabaseAccountsMongodbDatabase_Spec from our DatabaseAccountsMongodbDatabase_Spec
func (database *DatabaseAccountsMongodbDatabase_Spec) AssignPropertiesToDatabaseAccountsMongodbDatabase_Spec(destination *v20210515s.DatabaseAccountsMongodbDatabase_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(database.PropertyBag)

	// AzureName
	destination.AzureName = database.AzureName

	// Id
	destination.Id = genruntime.ClonePointerToString(database.Id)

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
		var resource v20210515s.MongoDBDatabaseResource
		err := database.Resource.AssignPropertiesToMongoDBDatabaseResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToMongoDBDatabaseResource() to populate field Resource")
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

// Storage version of v1alpha1api20210515.CreateUpdateOptions
// Deprecated version of CreateUpdateOptions. Use v1beta20210515.CreateUpdateOptions instead
type CreateUpdateOptions struct {
	AutoscaleSettings *AutoscaleSettings     `json:"autoscaleSettings,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Throughput        *int                   `json:"throughput,omitempty"`
}

// AssignPropertiesFromCreateUpdateOptions populates our CreateUpdateOptions from the provided source CreateUpdateOptions
func (options *CreateUpdateOptions) AssignPropertiesFromCreateUpdateOptions(source *v20210515s.CreateUpdateOptions) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AutoscaleSettings
	if source.AutoscaleSettings != nil {
		var autoscaleSetting AutoscaleSettings
		err := autoscaleSetting.AssignPropertiesFromAutoscaleSettings(source.AutoscaleSettings)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromAutoscaleSettings() to populate field AutoscaleSettings")
		}
		options.AutoscaleSettings = &autoscaleSetting
	} else {
		options.AutoscaleSettings = nil
	}

	// Throughput
	options.Throughput = genruntime.ClonePointerToInt(source.Throughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		options.PropertyBag = propertyBag
	} else {
		options.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToCreateUpdateOptions populates the provided destination CreateUpdateOptions from our CreateUpdateOptions
func (options *CreateUpdateOptions) AssignPropertiesToCreateUpdateOptions(destination *v20210515s.CreateUpdateOptions) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(options.PropertyBag)

	// AutoscaleSettings
	if options.AutoscaleSettings != nil {
		var autoscaleSetting v20210515s.AutoscaleSettings
		err := options.AutoscaleSettings.AssignPropertiesToAutoscaleSettings(&autoscaleSetting)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToAutoscaleSettings() to populate field AutoscaleSettings")
		}
		destination.AutoscaleSettings = &autoscaleSetting
	} else {
		destination.AutoscaleSettings = nil
	}

	// Throughput
	destination.Throughput = genruntime.ClonePointerToInt(options.Throughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.MongoDBDatabaseGetProperties_Resource_STATUS
// Deprecated version of MongoDBDatabaseGetProperties_Resource_STATUS. Use v1beta20210515.MongoDBDatabaseGetProperties_Resource_STATUS instead
type MongoDBDatabaseGetProperties_Resource_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	_Etag       *string                `json:"_etag,omitempty"`
	_Rid        *string                `json:"_rid,omitempty"`
	_Ts         *float64               `json:"_ts,omitempty"`
}

// AssignPropertiesFromMongoDBDatabaseGetProperties_Resource_STATUS populates our MongoDBDatabaseGetProperties_Resource_STATUS from the provided source MongoDBDatabaseGetProperties_Resource_STATUS
func (resource *MongoDBDatabaseGetProperties_Resource_STATUS) AssignPropertiesFromMongoDBDatabaseGetProperties_Resource_STATUS(source *v20210515s.MongoDBDatabaseGetProperties_Resource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// _Etag
	resource._Etag = genruntime.ClonePointerToString(source._Etag)

	// _Rid
	resource._Rid = genruntime.ClonePointerToString(source._Rid)

	// _Ts
	if source._Ts != nil {
		_T := *source._Ts
		resource._Ts = &_T
	} else {
		resource._Ts = nil
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

// AssignPropertiesToMongoDBDatabaseGetProperties_Resource_STATUS populates the provided destination MongoDBDatabaseGetProperties_Resource_STATUS from our MongoDBDatabaseGetProperties_Resource_STATUS
func (resource *MongoDBDatabaseGetProperties_Resource_STATUS) AssignPropertiesToMongoDBDatabaseGetProperties_Resource_STATUS(destination *v20210515s.MongoDBDatabaseGetProperties_Resource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// _Etag
	destination._Etag = genruntime.ClonePointerToString(resource._Etag)

	// _Rid
	destination._Rid = genruntime.ClonePointerToString(resource._Rid)

	// _Ts
	if resource._Ts != nil {
		_T := *resource._Ts
		destination._Ts = &_T
	} else {
		destination._Ts = nil
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

// Storage version of v1alpha1api20210515.MongoDBDatabaseResource
// Deprecated version of MongoDBDatabaseResource. Use v1beta20210515.MongoDBDatabaseResource instead
type MongoDBDatabaseResource struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromMongoDBDatabaseResource populates our MongoDBDatabaseResource from the provided source MongoDBDatabaseResource
func (resource *MongoDBDatabaseResource) AssignPropertiesFromMongoDBDatabaseResource(source *v20210515s.MongoDBDatabaseResource) error {
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

// AssignPropertiesToMongoDBDatabaseResource populates the provided destination MongoDBDatabaseResource from our MongoDBDatabaseResource
func (resource *MongoDBDatabaseResource) AssignPropertiesToMongoDBDatabaseResource(destination *v20210515s.MongoDBDatabaseResource) error {
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

// Storage version of v1alpha1api20210515.OptionsResource_STATUS
// Deprecated version of OptionsResource_STATUS. Use v1beta20210515.OptionsResource_STATUS instead
type OptionsResource_STATUS struct {
	AutoscaleSettings *AutoscaleSettings_STATUS `json:"autoscaleSettings,omitempty"`
	PropertyBag       genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Throughput        *int                      `json:"throughput,omitempty"`
}

// AssignPropertiesFromOptionsResource_STATUS populates our OptionsResource_STATUS from the provided source OptionsResource_STATUS
func (resource *OptionsResource_STATUS) AssignPropertiesFromOptionsResource_STATUS(source *v20210515s.OptionsResource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AutoscaleSettings
	if source.AutoscaleSettings != nil {
		var autoscaleSetting AutoscaleSettings_STATUS
		err := autoscaleSetting.AssignPropertiesFromAutoscaleSettings_STATUS(source.AutoscaleSettings)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromAutoscaleSettings_STATUS() to populate field AutoscaleSettings")
		}
		resource.AutoscaleSettings = &autoscaleSetting
	} else {
		resource.AutoscaleSettings = nil
	}

	// Throughput
	resource.Throughput = genruntime.ClonePointerToInt(source.Throughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToOptionsResource_STATUS populates the provided destination OptionsResource_STATUS from our OptionsResource_STATUS
func (resource *OptionsResource_STATUS) AssignPropertiesToOptionsResource_STATUS(destination *v20210515s.OptionsResource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// AutoscaleSettings
	if resource.AutoscaleSettings != nil {
		var autoscaleSetting v20210515s.AutoscaleSettings_STATUS
		err := resource.AutoscaleSettings.AssignPropertiesToAutoscaleSettings_STATUS(&autoscaleSetting)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToAutoscaleSettings_STATUS() to populate field AutoscaleSettings")
		}
		destination.AutoscaleSettings = &autoscaleSetting
	} else {
		destination.AutoscaleSettings = nil
	}

	// Throughput
	destination.Throughput = genruntime.ClonePointerToInt(resource.Throughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.AutoscaleSettings
// Deprecated version of AutoscaleSettings. Use v1beta20210515.AutoscaleSettings instead
type AutoscaleSettings struct {
	MaxThroughput *int                   `json:"maxThroughput,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromAutoscaleSettings populates our AutoscaleSettings from the provided source AutoscaleSettings
func (settings *AutoscaleSettings) AssignPropertiesFromAutoscaleSettings(source *v20210515s.AutoscaleSettings) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// MaxThroughput
	settings.MaxThroughput = genruntime.ClonePointerToInt(source.MaxThroughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		settings.PropertyBag = propertyBag
	} else {
		settings.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToAutoscaleSettings populates the provided destination AutoscaleSettings from our AutoscaleSettings
func (settings *AutoscaleSettings) AssignPropertiesToAutoscaleSettings(destination *v20210515s.AutoscaleSettings) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(settings.PropertyBag)

	// MaxThroughput
	destination.MaxThroughput = genruntime.ClonePointerToInt(settings.MaxThroughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.AutoscaleSettings_STATUS
// Deprecated version of AutoscaleSettings_STATUS. Use v1beta20210515.AutoscaleSettings_STATUS instead
type AutoscaleSettings_STATUS struct {
	MaxThroughput *int                   `json:"maxThroughput,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromAutoscaleSettings_STATUS populates our AutoscaleSettings_STATUS from the provided source AutoscaleSettings_STATUS
func (settings *AutoscaleSettings_STATUS) AssignPropertiesFromAutoscaleSettings_STATUS(source *v20210515s.AutoscaleSettings_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// MaxThroughput
	settings.MaxThroughput = genruntime.ClonePointerToInt(source.MaxThroughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		settings.PropertyBag = propertyBag
	} else {
		settings.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToAutoscaleSettings_STATUS populates the provided destination AutoscaleSettings_STATUS from our AutoscaleSettings_STATUS
func (settings *AutoscaleSettings_STATUS) AssignPropertiesToAutoscaleSettings_STATUS(destination *v20210515s.AutoscaleSettings_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(settings.PropertyBag)

	// MaxThroughput
	destination.MaxThroughput = genruntime.ClonePointerToInt(settings.MaxThroughput)

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
	SchemeBuilder.Register(&MongodbDatabase{}, &MongodbDatabaseList{})
}
