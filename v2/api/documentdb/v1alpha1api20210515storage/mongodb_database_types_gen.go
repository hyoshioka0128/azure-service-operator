// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=mongodbdatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={mongodbdatabases/status,mongodbdatabases/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210515.MongodbDatabase
type MongodbDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsMongodbDatabases_SPEC        `json:"spec,omitempty"`
	Status            MongoDBDatabaseCreateUpdateParameters_Status `json:"status,omitempty"`
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

var _ genruntime.KubernetesResource = &MongodbDatabase{}

// AzureName returns the Azure name of the resource
func (database *MongodbDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (database MongodbDatabase) GetAPIVersion() string {
	return string(APIVersionValue)
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
	return &MongoDBDatabaseCreateUpdateParameters_Status{}
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
	if st, ok := status.(*MongoDBDatabaseCreateUpdateParameters_Status); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st MongoDBDatabaseCreateUpdateParameters_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// Hub marks that this MongodbDatabase is the hub type for conversion
func (database *MongodbDatabase) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *MongodbDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion,
		Kind:    "MongodbDatabase",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.MongodbDatabase
type MongodbDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabase `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountsMongodbDatabases_SPEC
type DatabaseAccountsMongodbDatabases_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string                    `json:"azureName"`
	Location        *string                   `json:"location,omitempty"`
	Options         *CreateUpdateOptions_Spec `json:"options,omitempty"`
	OriginalVersion string                    `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Resource    *MongoDBDatabaseResource_Spec     `json:"resource,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsMongodbDatabases_SPEC{}

// ConvertSpecFrom populates our DatabaseAccountsMongodbDatabases_SPEC from the provided source
func (spec *DatabaseAccountsMongodbDatabases_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsMongodbDatabases_SPEC
func (spec *DatabaseAccountsMongodbDatabases_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

//Storage version of v1alpha1api20210515.MongoDBDatabaseCreateUpdateParameters_Status
type MongoDBDatabaseCreateUpdateParameters_Status struct {
	Conditions  []conditions.Condition          `json:"conditions,omitempty"`
	Id          *string                         `json:"id,omitempty"`
	Location    *string                         `json:"location,omitempty"`
	Name        *string                         `json:"name,omitempty"`
	Options     *CreateUpdateOptions_Status     `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
	Resource    *MongoDBDatabaseResource_Status `json:"resource,omitempty"`
	Tags        map[string]string               `json:"tags,omitempty"`
	Type        *string                         `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &MongoDBDatabaseCreateUpdateParameters_Status{}

// ConvertStatusFrom populates our MongoDBDatabaseCreateUpdateParameters_Status from the provided source
func (parameters *MongoDBDatabaseCreateUpdateParameters_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == parameters {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(parameters)
}

// ConvertStatusTo populates the provided destination from our MongoDBDatabaseCreateUpdateParameters_Status
func (parameters *MongoDBDatabaseCreateUpdateParameters_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == parameters {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(parameters)
}

//Storage version of v1alpha1api20210515.CreateUpdateOptions_Spec
type CreateUpdateOptions_Spec struct {
	AutoscaleSettings *AutoscaleSettings_Spec `json:"autoscaleSettings,omitempty"`
	PropertyBag       genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	Throughput        *int                    `json:"throughput,omitempty"`
}

//Storage version of v1alpha1api20210515.CreateUpdateOptions_Status
type CreateUpdateOptions_Status struct {
	AutoscaleSettings *AutoscaleSettings_Status `json:"autoscaleSettings,omitempty"`
	PropertyBag       genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Throughput        *int                      `json:"throughput,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoDBDatabaseResource_Spec
type MongoDBDatabaseResource_Spec struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoDBDatabaseResource_Status
type MongoDBDatabaseResource_Status struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.AutoscaleSettings_Spec
type AutoscaleSettings_Spec struct {
	MaxThroughput *int                   `json:"maxThroughput,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.AutoscaleSettings_Status
type AutoscaleSettings_Status struct {
	MaxThroughput *int                   `json:"maxThroughput,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&MongodbDatabase{}, &MongodbDatabaseList{})
}