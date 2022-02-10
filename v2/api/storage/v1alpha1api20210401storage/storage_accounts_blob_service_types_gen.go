// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=storage.azure.com,resources=storageaccountsblobservices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.azure.com,resources={storageaccountsblobservices/status,storageaccountsblobservices/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210401.StorageAccountsBlobService
type StorageAccountsBlobService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsBlobServices_SPEC `json:"spec,omitempty"`
	Status            BlobServiceProperties_Status     `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsBlobService{}

// GetConditions returns the conditions of the resource
func (service *StorageAccountsBlobService) GetConditions() conditions.Conditions {
	return service.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (service *StorageAccountsBlobService) SetConditions(conditions conditions.Conditions) {
	service.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &StorageAccountsBlobService{}

// AzureName returns the Azure name of the resource
func (service *StorageAccountsBlobService) AzureName() string {
	return service.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (service StorageAccountsBlobService) GetAPIVersion() string {
	return "2021-04-01"
}

// GetResourceKind returns the kind of the resource
func (service *StorageAccountsBlobService) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (service *StorageAccountsBlobService) GetSpec() genruntime.ConvertibleSpec {
	return &service.Spec
}

// GetStatus returns the status of this resource
func (service *StorageAccountsBlobService) GetStatus() genruntime.ConvertibleStatus {
	return &service.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (service *StorageAccountsBlobService) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (service *StorageAccountsBlobService) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &BlobServiceProperties_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (service *StorageAccountsBlobService) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(service.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  service.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (service *StorageAccountsBlobService) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*BlobServiceProperties_Status); ok {
		service.Status = *st
		return nil
	}

	// Convert status to required version
	var st BlobServiceProperties_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	service.Status = st
	return nil
}

// Hub marks that this StorageAccountsBlobService is the hub type for conversion
func (service *StorageAccountsBlobService) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (service *StorageAccountsBlobService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: service.Spec.OriginalVersion,
		Kind:    "StorageAccountsBlobService",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210401.StorageAccountsBlobService
type StorageAccountsBlobServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsBlobService `json:"items"`
}

//Storage version of v1alpha1api20210401.BlobServiceProperties_Status
type BlobServiceProperties_Status struct {
	AutomaticSnapshotPolicyEnabled *bool                                `json:"automaticSnapshotPolicyEnabled,omitempty"`
	ChangeFeed                     *ChangeFeed_Status                   `json:"changeFeed,omitempty"`
	Conditions                     []conditions.Condition               `json:"conditions,omitempty"`
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicy_Status        `json:"containerDeleteRetentionPolicy,omitempty"`
	Cors                           *CorsRules_Status                    `json:"cors,omitempty"`
	DefaultServiceVersion          *string                              `json:"defaultServiceVersion,omitempty"`
	DeleteRetentionPolicy          *DeleteRetentionPolicy_Status        `json:"deleteRetentionPolicy,omitempty"`
	Id                             *string                              `json:"id,omitempty"`
	IsVersioningEnabled            *bool                                `json:"isVersioningEnabled,omitempty"`
	LastAccessTimeTrackingPolicy   *LastAccessTimeTrackingPolicy_Status `json:"lastAccessTimeTrackingPolicy,omitempty"`
	Name                           *string                              `json:"name,omitempty"`
	PropertyBag                    genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	RestorePolicy                  *RestorePolicyProperties_Status      `json:"restorePolicy,omitempty"`
	Sku                            *Sku_Status                          `json:"sku,omitempty"`
	Type                           *string                              `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &BlobServiceProperties_Status{}

// ConvertStatusFrom populates our BlobServiceProperties_Status from the provided source
func (properties *BlobServiceProperties_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == properties {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(properties)
}

// ConvertStatusTo populates the provided destination from our BlobServiceProperties_Status
func (properties *BlobServiceProperties_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == properties {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(properties)
}

//Storage version of v1alpha1api20210401.StorageAccountsBlobServices_SPEC
type StorageAccountsBlobServices_SPEC struct {
	AutomaticSnapshotPolicyEnabled *bool `json:"automaticSnapshotPolicyEnabled,omitempty"`

	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName                      string                             `json:"azureName"`
	ChangeFeed                     *ChangeFeed_Spec                   `json:"changeFeed,omitempty"`
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicy_Spec        `json:"containerDeleteRetentionPolicy,omitempty"`
	Cors                           *CorsRules_Spec                    `json:"cors,omitempty"`
	DefaultServiceVersion          *string                            `json:"defaultServiceVersion,omitempty"`
	DeleteRetentionPolicy          *DeleteRetentionPolicy_Spec        `json:"deleteRetentionPolicy,omitempty"`
	IsVersioningEnabled            *bool                              `json:"isVersioningEnabled,omitempty"`
	LastAccessTimeTrackingPolicy   *LastAccessTimeTrackingPolicy_Spec `json:"lastAccessTimeTrackingPolicy,omitempty"`
	OriginalVersion                string                             `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner         genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag   genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	RestorePolicy *RestorePolicyProperties_Spec     `json:"restorePolicy,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccountsBlobServices_SPEC{}

// ConvertSpecFrom populates our StorageAccountsBlobServices_SPEC from the provided source
func (spec *StorageAccountsBlobServices_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our StorageAccountsBlobServices_SPEC
func (spec *StorageAccountsBlobServices_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

//Storage version of v1alpha1api20210401.ChangeFeed_Spec
type ChangeFeed_Spec struct {
	Enabled         *bool                  `json:"enabled,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RetentionInDays *int                   `json:"retentionInDays,omitempty"`
}

//Storage version of v1alpha1api20210401.ChangeFeed_Status
type ChangeFeed_Status struct {
	Enabled         *bool                  `json:"enabled,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RetentionInDays *int                   `json:"retentionInDays,omitempty"`
}

//Storage version of v1alpha1api20210401.CorsRules_Spec
type CorsRules_Spec struct {
	CorsRules   []CorsRule_Spec        `json:"corsRules,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.CorsRules_Status
type CorsRules_Status struct {
	CorsRules   []CorsRule_Status      `json:"corsRules,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.DeleteRetentionPolicy_Spec
type DeleteRetentionPolicy_Spec struct {
	Days        *int                   `json:"days,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.DeleteRetentionPolicy_Status
type DeleteRetentionPolicy_Status struct {
	Days        *int                   `json:"days,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.LastAccessTimeTrackingPolicy_Spec
type LastAccessTimeTrackingPolicy_Spec struct {
	BlobType                  []string               `json:"blobType,omitempty"`
	Enable                    *bool                  `json:"enable,omitempty"`
	Name                      *string                `json:"name,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TrackingGranularityInDays *int                   `json:"trackingGranularityInDays,omitempty"`
}

//Storage version of v1alpha1api20210401.LastAccessTimeTrackingPolicy_Status
type LastAccessTimeTrackingPolicy_Status struct {
	BlobType                  []string               `json:"blobType,omitempty"`
	Enable                    *bool                  `json:"enable,omitempty"`
	Name                      *string                `json:"name,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TrackingGranularityInDays *int                   `json:"trackingGranularityInDays,omitempty"`
}

//Storage version of v1alpha1api20210401.RestorePolicyProperties_Spec
type RestorePolicyProperties_Spec struct {
	Days        *int                   `json:"days,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.RestorePolicyProperties_Status
type RestorePolicyProperties_Status struct {
	Days            *int                   `json:"days,omitempty"`
	Enabled         *bool                  `json:"enabled,omitempty"`
	LastEnabledTime *string                `json:"lastEnabledTime,omitempty"`
	MinRestoreTime  *string                `json:"minRestoreTime,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.CorsRule_Spec
type CorsRule_Spec struct {
	AllowedHeaders  []string               `json:"allowedHeaders,omitempty"`
	AllowedMethods  []string               `json:"allowedMethods,omitempty"`
	AllowedOrigins  []string               `json:"allowedOrigins,omitempty"`
	ExposedHeaders  []string               `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                   `json:"maxAgeInSeconds,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.CorsRule_Status
type CorsRule_Status struct {
	AllowedHeaders  []string               `json:"allowedHeaders,omitempty"`
	AllowedMethods  []string               `json:"allowedMethods,omitempty"`
	AllowedOrigins  []string               `json:"allowedOrigins,omitempty"`
	ExposedHeaders  []string               `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                   `json:"maxAgeInSeconds,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&StorageAccountsBlobService{}, &StorageAccountsBlobServiceList{})
}
