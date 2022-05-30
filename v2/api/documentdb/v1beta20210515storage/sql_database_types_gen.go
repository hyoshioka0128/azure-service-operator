// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=sqldatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={sqldatabases/status,sqldatabases/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210515.SqlDatabase
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
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

// Hub marks that this SqlDatabase is the hub type for conversion
func (database *SqlDatabase) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *SqlDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion,
		Kind:    "SqlDatabase",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210515.SqlDatabase
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
type SqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabase `json:"items"`
}

// Storage version of v1beta20210515.DatabaseAccountsSqlDatabase_STATUS
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
	if source == database {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(database)
}

// ConvertStatusTo populates the provided destination from our DatabaseAccountsSqlDatabase_STATUS
func (database *DatabaseAccountsSqlDatabase_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == database {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(database)
}

// Storage version of v1beta20210515.DatabaseAccountsSqlDatabase_Spec
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
	if source == database {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(database)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabase_Spec
func (database *DatabaseAccountsSqlDatabase_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == database {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(database)
}

// Storage version of v1beta20210515.SqlDatabaseGetProperties_Resource_STATUS
type SqlDatabaseGetProperties_Resource_STATUS struct {
	Colls       *string                `json:"_colls,omitempty"`
	Etag        *string                `json:"_etag,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid         *string                `json:"_rid,omitempty"`
	Ts          *float64               `json:"_ts,omitempty"`
	Users       *string                `json:"_users,omitempty"`
}

// Storage version of v1beta20210515.SqlDatabaseResource
type SqlDatabaseResource struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SqlDatabase{}, &SqlDatabaseList{})
}
