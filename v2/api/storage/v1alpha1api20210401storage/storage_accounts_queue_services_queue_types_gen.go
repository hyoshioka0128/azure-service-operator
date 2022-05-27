// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401storage

import (
	"fmt"
	v20210401s "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401storage"
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
// Storage version of v1alpha1api20210401.StorageAccountsQueueServicesQueue
// Deprecated version of StorageAccountsQueueServicesQueue. Use v1beta20210401.StorageAccountsQueueServicesQueue instead
type StorageAccountsQueueServicesQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsQueueServicesQueue_Spec   `json:"spec,omitempty"`
	Status            StorageAccountsQueueServicesQueue_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsQueueServicesQueue{}

// GetConditions returns the conditions of the resource
func (queue *StorageAccountsQueueServicesQueue) GetConditions() conditions.Conditions {
	return queue.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (queue *StorageAccountsQueueServicesQueue) SetConditions(conditions conditions.Conditions) {
	queue.Status.Conditions = conditions
}

var _ conversion.Convertible = &StorageAccountsQueueServicesQueue{}

// ConvertFrom populates our StorageAccountsQueueServicesQueue from the provided hub StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210401s.StorageAccountsQueueServicesQueue)
	if !ok {
		return fmt.Errorf("expected storage/v1beta20210401storage/StorageAccountsQueueServicesQueue but received %T instead", hub)
	}

	return queue.AssignPropertiesFromStorageAccountsQueueServicesQueue(source)
}

// ConvertTo populates the provided hub StorageAccountsQueueServicesQueue from our StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210401s.StorageAccountsQueueServicesQueue)
	if !ok {
		return fmt.Errorf("expected storage/v1beta20210401storage/StorageAccountsQueueServicesQueue but received %T instead", hub)
	}

	return queue.AssignPropertiesToStorageAccountsQueueServicesQueue(destination)
}

var _ genruntime.KubernetesResource = &StorageAccountsQueueServicesQueue{}

// AzureName returns the Azure name of the resource
func (queue *StorageAccountsQueueServicesQueue) AzureName() string {
	return queue.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "20210401"
func (queue StorageAccountsQueueServicesQueue) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (queue *StorageAccountsQueueServicesQueue) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (queue *StorageAccountsQueueServicesQueue) GetSpec() genruntime.ConvertibleSpec {
	return &queue.Spec
}

// GetStatus returns the status of this resource
func (queue *StorageAccountsQueueServicesQueue) GetStatus() genruntime.ConvertibleStatus {
	return &queue.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (queue *StorageAccountsQueueServicesQueue) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (queue *StorageAccountsQueueServicesQueue) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccountsQueueServicesQueue_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (queue *StorageAccountsQueueServicesQueue) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(queue.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  queue.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (queue *StorageAccountsQueueServicesQueue) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccountsQueueServicesQueue_STATUS); ok {
		queue.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccountsQueueServicesQueue_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	queue.Status = st
	return nil
}

// AssignPropertiesFromStorageAccountsQueueServicesQueue populates our StorageAccountsQueueServicesQueue from the provided source StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) AssignPropertiesFromStorageAccountsQueueServicesQueue(source *v20210401s.StorageAccountsQueueServicesQueue) error {

	// ObjectMeta
	queue.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec StorageAccountsQueueServicesQueue_Spec
	err := spec.AssignPropertiesFromStorageAccountsQueueServicesQueue_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromStorageAccountsQueueServicesQueue_Spec() to populate field Spec")
	}
	queue.Spec = spec

	// Status
	var status StorageAccountsQueueServicesQueue_STATUS
	err = status.AssignPropertiesFromStorageAccountsQueueServicesQueue_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromStorageAccountsQueueServicesQueue_STATUS() to populate field Status")
	}
	queue.Status = status

	// No error
	return nil
}

// AssignPropertiesToStorageAccountsQueueServicesQueue populates the provided destination StorageAccountsQueueServicesQueue from our StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) AssignPropertiesToStorageAccountsQueueServicesQueue(destination *v20210401s.StorageAccountsQueueServicesQueue) error {

	// ObjectMeta
	destination.ObjectMeta = *queue.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210401s.StorageAccountsQueueServicesQueue_Spec
	err := queue.Spec.AssignPropertiesToStorageAccountsQueueServicesQueue_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToStorageAccountsQueueServicesQueue_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210401s.StorageAccountsQueueServicesQueue_STATUS
	err = queue.Status.AssignPropertiesToStorageAccountsQueueServicesQueue_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToStorageAccountsQueueServicesQueue_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (queue *StorageAccountsQueueServicesQueue) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: queue.Spec.OriginalVersion,
		Kind:    "StorageAccountsQueueServicesQueue",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210401.StorageAccountsQueueServicesQueue
// Deprecated version of StorageAccountsQueueServicesQueue. Use v1beta20210401.StorageAccountsQueueServicesQueue instead
type StorageAccountsQueueServicesQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsQueueServicesQueue `json:"items"`
}

// Storage version of v1alpha1api20210401.StorageAccountsQueueServicesQueue_STATUS
// Deprecated version of StorageAccountsQueueServicesQueue_STATUS. Use v1beta20210401.StorageAccountsQueueServicesQueue_STATUS instead
type StorageAccountsQueueServicesQueue_STATUS struct {
	ApproximateMessageCount *int                   `json:"approximateMessageCount,omitempty"`
	Conditions              []conditions.Condition `json:"conditions,omitempty"`
	Id                      *string                `json:"id,omitempty"`
	Metadata                map[string]string      `json:"metadata,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type                    *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccountsQueueServicesQueue_STATUS{}

// ConvertStatusFrom populates our StorageAccountsQueueServicesQueue_STATUS from the provided source
func (queue *StorageAccountsQueueServicesQueue_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210401s.StorageAccountsQueueServicesQueue_STATUS)
	if ok {
		// Populate our instance from source
		return queue.AssignPropertiesFromStorageAccountsQueueServicesQueue_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210401s.StorageAccountsQueueServicesQueue_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = queue.AssignPropertiesFromStorageAccountsQueueServicesQueue_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our StorageAccountsQueueServicesQueue_STATUS
func (queue *StorageAccountsQueueServicesQueue_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210401s.StorageAccountsQueueServicesQueue_STATUS)
	if ok {
		// Populate destination from our instance
		return queue.AssignPropertiesToStorageAccountsQueueServicesQueue_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210401s.StorageAccountsQueueServicesQueue_STATUS{}
	err := queue.AssignPropertiesToStorageAccountsQueueServicesQueue_STATUS(dst)
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

// AssignPropertiesFromStorageAccountsQueueServicesQueue_STATUS populates our StorageAccountsQueueServicesQueue_STATUS from the provided source StorageAccountsQueueServicesQueue_STATUS
func (queue *StorageAccountsQueueServicesQueue_STATUS) AssignPropertiesFromStorageAccountsQueueServicesQueue_STATUS(source *v20210401s.StorageAccountsQueueServicesQueue_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ApproximateMessageCount
	queue.ApproximateMessageCount = genruntime.ClonePointerToInt(source.ApproximateMessageCount)

	// Conditions
	queue.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	queue.Id = genruntime.ClonePointerToString(source.Id)

	// Metadata
	queue.Metadata = genruntime.CloneMapOfStringToString(source.Metadata)

	// Name
	queue.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	queue.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		queue.PropertyBag = propertyBag
	} else {
		queue.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToStorageAccountsQueueServicesQueue_STATUS populates the provided destination StorageAccountsQueueServicesQueue_STATUS from our StorageAccountsQueueServicesQueue_STATUS
func (queue *StorageAccountsQueueServicesQueue_STATUS) AssignPropertiesToStorageAccountsQueueServicesQueue_STATUS(destination *v20210401s.StorageAccountsQueueServicesQueue_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(queue.PropertyBag)

	// ApproximateMessageCount
	destination.ApproximateMessageCount = genruntime.ClonePointerToInt(queue.ApproximateMessageCount)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(queue.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(queue.Id)

	// Metadata
	destination.Metadata = genruntime.CloneMapOfStringToString(queue.Metadata)

	// Name
	destination.Name = genruntime.ClonePointerToString(queue.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(queue.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210401.StorageAccountsQueueServicesQueue_Spec
type StorageAccountsQueueServicesQueue_Spec struct {
	ApproximateMessageCount *int `json:"approximateMessageCount,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string            `json:"azureName,omitempty"`
	Id              *string           `json:"id,omitempty"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	OriginalVersion string            `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Type        *string                            `json:"type,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccountsQueueServicesQueue_Spec{}

// ConvertSpecFrom populates our StorageAccountsQueueServicesQueue_Spec from the provided source
func (queue *StorageAccountsQueueServicesQueue_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210401s.StorageAccountsQueueServicesQueue_Spec)
	if ok {
		// Populate our instance from source
		return queue.AssignPropertiesFromStorageAccountsQueueServicesQueue_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210401s.StorageAccountsQueueServicesQueue_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = queue.AssignPropertiesFromStorageAccountsQueueServicesQueue_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our StorageAccountsQueueServicesQueue_Spec
func (queue *StorageAccountsQueueServicesQueue_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210401s.StorageAccountsQueueServicesQueue_Spec)
	if ok {
		// Populate destination from our instance
		return queue.AssignPropertiesToStorageAccountsQueueServicesQueue_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210401s.StorageAccountsQueueServicesQueue_Spec{}
	err := queue.AssignPropertiesToStorageAccountsQueueServicesQueue_Spec(dst)
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

// AssignPropertiesFromStorageAccountsQueueServicesQueue_Spec populates our StorageAccountsQueueServicesQueue_Spec from the provided source StorageAccountsQueueServicesQueue_Spec
func (queue *StorageAccountsQueueServicesQueue_Spec) AssignPropertiesFromStorageAccountsQueueServicesQueue_Spec(source *v20210401s.StorageAccountsQueueServicesQueue_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ApproximateMessageCount
	queue.ApproximateMessageCount = genruntime.ClonePointerToInt(source.ApproximateMessageCount)

	// AzureName
	queue.AzureName = source.AzureName

	// Id
	queue.Id = genruntime.ClonePointerToString(source.Id)

	// Metadata
	queue.Metadata = genruntime.CloneMapOfStringToString(source.Metadata)

	// OriginalVersion
	queue.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		queue.Owner = &owner
	} else {
		queue.Owner = nil
	}

	// Type
	queue.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		queue.PropertyBag = propertyBag
	} else {
		queue.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToStorageAccountsQueueServicesQueue_Spec populates the provided destination StorageAccountsQueueServicesQueue_Spec from our StorageAccountsQueueServicesQueue_Spec
func (queue *StorageAccountsQueueServicesQueue_Spec) AssignPropertiesToStorageAccountsQueueServicesQueue_Spec(destination *v20210401s.StorageAccountsQueueServicesQueue_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(queue.PropertyBag)

	// ApproximateMessageCount
	destination.ApproximateMessageCount = genruntime.ClonePointerToInt(queue.ApproximateMessageCount)

	// AzureName
	destination.AzureName = queue.AzureName

	// Id
	destination.Id = genruntime.ClonePointerToString(queue.Id)

	// Metadata
	destination.Metadata = genruntime.CloneMapOfStringToString(queue.Metadata)

	// OriginalVersion
	destination.OriginalVersion = queue.OriginalVersion

	// Owner
	if queue.Owner != nil {
		owner := queue.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(queue.Type)

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
	SchemeBuilder.Register(&StorageAccountsQueueServicesQueue{}, &StorageAccountsQueueServicesQueueList{})
}
