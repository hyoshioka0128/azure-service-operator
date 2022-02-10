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

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=sqldatabasecontainers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={sqldatabasecontainers/status,sqldatabasecontainers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210515.SqlDatabaseContainer
type SqlDatabaseContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainers_SPEC `json:"spec,omitempty"`
	Status            SqlContainerCreateUpdateParameters_Status   `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainer{}

// GetConditions returns the conditions of the resource
func (container *SqlDatabaseContainer) GetConditions() conditions.Conditions {
	return container.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (container *SqlDatabaseContainer) SetConditions(conditions conditions.Conditions) {
	container.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainer{}

// AzureName returns the Azure name of the resource
func (container *SqlDatabaseContainer) AzureName() string {
	return container.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (container SqlDatabaseContainer) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceKind returns the kind of the resource
func (container *SqlDatabaseContainer) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (container *SqlDatabaseContainer) GetSpec() genruntime.ConvertibleSpec {
	return &container.Spec
}

// GetStatus returns the status of this resource
func (container *SqlDatabaseContainer) GetStatus() genruntime.ConvertibleStatus {
	return &container.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (container *SqlDatabaseContainer) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (container *SqlDatabaseContainer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlContainerCreateUpdateParameters_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (container *SqlDatabaseContainer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(container.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  container.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (container *SqlDatabaseContainer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlContainerCreateUpdateParameters_Status); ok {
		container.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlContainerCreateUpdateParameters_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	container.Status = st
	return nil
}

// Hub marks that this SqlDatabaseContainer is the hub type for conversion
func (container *SqlDatabaseContainer) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (container *SqlDatabaseContainer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: container.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainer",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.SqlDatabaseContainer
type SqlDatabaseContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainer `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabasesContainers_SPEC
type DatabaseAccountsSqlDatabasesContainers_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string                    `json:"azureName"`
	Location        *string                   `json:"location,omitempty"`
	Options         *CreateUpdateOptions_Spec `json:"options,omitempty"`
	OriginalVersion string                    `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Resource    *SqlContainerResource_Spec        `json:"resource,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainers_SPEC{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainers_SPEC from the provided source
func (spec *DatabaseAccountsSqlDatabasesContainers_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainers_SPEC
func (spec *DatabaseAccountsSqlDatabasesContainers_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

//Storage version of v1alpha1api20210515.SqlContainerCreateUpdateParameters_Status
type SqlContainerCreateUpdateParameters_Status struct {
	Conditions  []conditions.Condition       `json:"conditions,omitempty"`
	Id          *string                      `json:"id,omitempty"`
	Location    *string                      `json:"location,omitempty"`
	Name        *string                      `json:"name,omitempty"`
	Options     *CreateUpdateOptions_Status  `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Resource    *SqlContainerResource_Status `json:"resource,omitempty"`
	Tags        map[string]string            `json:"tags,omitempty"`
	Type        *string                      `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlContainerCreateUpdateParameters_Status{}

// ConvertStatusFrom populates our SqlContainerCreateUpdateParameters_Status from the provided source
func (parameters *SqlContainerCreateUpdateParameters_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == parameters {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(parameters)
}

// ConvertStatusTo populates the provided destination from our SqlContainerCreateUpdateParameters_Status
func (parameters *SqlContainerCreateUpdateParameters_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == parameters {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(parameters)
}

//Storage version of v1alpha1api20210515.SqlContainerResource_Spec
type SqlContainerResource_Spec struct {
	AnalyticalStorageTtl     *int                           `json:"analyticalStorageTtl,omitempty"`
	ConflictResolutionPolicy *ConflictResolutionPolicy_Spec `json:"conflictResolutionPolicy,omitempty"`
	DefaultTtl               *int                           `json:"defaultTtl,omitempty"`
	Id                       *string                        `json:"id,omitempty"`
	IndexingPolicy           *IndexingPolicy_Spec           `json:"indexingPolicy,omitempty"`
	PartitionKey             *ContainerPartitionKey_Spec    `json:"partitionKey,omitempty"`
	PropertyBag              genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	UniqueKeyPolicy          *UniqueKeyPolicy_Spec          `json:"uniqueKeyPolicy,omitempty"`
}

//Storage version of v1alpha1api20210515.SqlContainerResource_Status
type SqlContainerResource_Status struct {
	AnalyticalStorageTtl     *int                             `json:"analyticalStorageTtl,omitempty"`
	ConflictResolutionPolicy *ConflictResolutionPolicy_Status `json:"conflictResolutionPolicy,omitempty"`
	DefaultTtl               *int                             `json:"defaultTtl,omitempty"`
	Id                       *string                          `json:"id,omitempty"`
	IndexingPolicy           *IndexingPolicy_Status           `json:"indexingPolicy,omitempty"`
	PartitionKey             *ContainerPartitionKey_Status    `json:"partitionKey,omitempty"`
	PropertyBag              genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	UniqueKeyPolicy          *UniqueKeyPolicy_Status          `json:"uniqueKeyPolicy,omitempty"`
}

//Storage version of v1alpha1api20210515.ConflictResolutionPolicy_Spec
type ConflictResolutionPolicy_Spec struct {
	ConflictResolutionPath      *string                `json:"conflictResolutionPath,omitempty"`
	ConflictResolutionProcedure *string                `json:"conflictResolutionProcedure,omitempty"`
	Mode                        *string                `json:"mode,omitempty"`
	PropertyBag                 genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ConflictResolutionPolicy_Status
type ConflictResolutionPolicy_Status struct {
	ConflictResolutionPath      *string                `json:"conflictResolutionPath,omitempty"`
	ConflictResolutionProcedure *string                `json:"conflictResolutionProcedure,omitempty"`
	Mode                        *string                `json:"mode,omitempty"`
	PropertyBag                 genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ContainerPartitionKey_Spec
type ContainerPartitionKey_Spec struct {
	Kind        *string                `json:"kind,omitempty"`
	Paths       []string               `json:"paths,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Version     *int                   `json:"version,omitempty"`
}

//Storage version of v1alpha1api20210515.ContainerPartitionKey_Status
type ContainerPartitionKey_Status struct {
	Kind        *string                `json:"kind,omitempty"`
	Paths       []string               `json:"paths,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemKey   *bool                  `json:"systemKey,omitempty"`
	Version     *int                   `json:"version,omitempty"`
}

//Storage version of v1alpha1api20210515.IndexingPolicy_Spec
type IndexingPolicy_Spec struct {
	Automatic        *bool                  `json:"automatic,omitempty"`
	CompositeIndexes [][]CompositePath_Spec `json:"compositeIndexes,omitempty"`
	ExcludedPaths    []ExcludedPath_Spec    `json:"excludedPaths,omitempty"`
	IncludedPaths    []IncludedPath_Spec    `json:"includedPaths,omitempty"`
	IndexingMode     *string                `json:"indexingMode,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SpatialIndexes   []SpatialSpec_Spec     `json:"spatialIndexes,omitempty"`
}

//Storage version of v1alpha1api20210515.IndexingPolicy_Status
type IndexingPolicy_Status struct {
	Automatic        *bool                    `json:"automatic,omitempty"`
	CompositeIndexes [][]CompositePath_Status `json:"compositeIndexes,omitempty"`
	ExcludedPaths    []ExcludedPath_Status    `json:"excludedPaths,omitempty"`
	IncludedPaths    []IncludedPath_Status    `json:"includedPaths,omitempty"`
	IndexingMode     *string                  `json:"indexingMode,omitempty"`
	PropertyBag      genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	SpatialIndexes   []SpatialSpec_Status     `json:"spatialIndexes,omitempty"`
}

//Storage version of v1alpha1api20210515.UniqueKeyPolicy_Spec
type UniqueKeyPolicy_Spec struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UniqueKeys  []UniqueKey_Spec       `json:"uniqueKeys,omitempty"`
}

//Storage version of v1alpha1api20210515.UniqueKeyPolicy_Status
type UniqueKeyPolicy_Status struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UniqueKeys  []UniqueKey_Status     `json:"uniqueKeys,omitempty"`
}

//Storage version of v1alpha1api20210515.CompositePath_Spec
type CompositePath_Spec struct {
	Order       *string                `json:"order,omitempty"`
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.CompositePath_Status
type CompositePath_Status struct {
	Order       *string                `json:"order,omitempty"`
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ExcludedPath_Spec
type ExcludedPath_Spec struct {
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ExcludedPath_Status
type ExcludedPath_Status struct {
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.IncludedPath_Spec
type IncludedPath_Spec struct {
	Indexes     []Indexes_Spec         `json:"indexes,omitempty"`
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.IncludedPath_Status
type IncludedPath_Status struct {
	Indexes     []Indexes_Status       `json:"indexes,omitempty"`
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.SpatialSpec_Spec
type SpatialSpec_Spec struct {
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Types       []string               `json:"types,omitempty"`
}

//Storage version of v1alpha1api20210515.SpatialSpec_Status
type SpatialSpec_Status struct {
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Types       []string               `json:"types,omitempty"`
}

//Storage version of v1alpha1api20210515.UniqueKey_Spec
type UniqueKey_Spec struct {
	Paths       []string               `json:"paths,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.UniqueKey_Status
type UniqueKey_Status struct {
	Paths       []string               `json:"paths,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.Indexes_Spec
type Indexes_Spec struct {
	DataType    *string                `json:"dataType,omitempty"`
	Kind        *string                `json:"kind,omitempty"`
	Precision   *int                   `json:"precision,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.Indexes_Status
type Indexes_Status struct {
	DataType    *string                `json:"dataType,omitempty"`
	Kind        *string                `json:"kind,omitempty"`
	Precision   *int                   `json:"precision,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainer{}, &SqlDatabaseContainerList{})
}
