// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

import (
	"fmt"
	v20201201s "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20201201storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2020-12-01/redis.json
type RedisPatchSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisPatchSchedule_Spec   `json:"spec,omitempty"`
	Status            RedisPatchSchedule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisPatchSchedule{}

// GetConditions returns the conditions of the resource
func (schedule *RedisPatchSchedule) GetConditions() conditions.Conditions {
	return schedule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (schedule *RedisPatchSchedule) SetConditions(conditions conditions.Conditions) {
	schedule.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisPatchSchedule{}

// ConvertFrom populates our RedisPatchSchedule from the provided hub RedisPatchSchedule
func (schedule *RedisPatchSchedule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20201201s.RedisPatchSchedule)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20201201storage/RedisPatchSchedule but received %T instead", hub)
	}

	return schedule.AssignPropertiesFromRedisPatchSchedule(source)
}

// ConvertTo populates the provided hub RedisPatchSchedule from our RedisPatchSchedule
func (schedule *RedisPatchSchedule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20201201s.RedisPatchSchedule)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20201201storage/RedisPatchSchedule but received %T instead", hub)
	}

	return schedule.AssignPropertiesToRedisPatchSchedule(destination)
}

// +kubebuilder:webhook:path=/mutate-cache-azure-com-v1beta20201201-redispatchschedule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redispatchschedules,verbs=create;update,versions=v1beta20201201,name=default.v1beta20201201.redispatchschedules.cache.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &RedisPatchSchedule{}

// Default applies defaults to the RedisPatchSchedule resource
func (schedule *RedisPatchSchedule) Default() {
	schedule.defaultImpl()
	var temp interface{} = schedule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (schedule *RedisPatchSchedule) defaultAzureName() {
	if schedule.Spec.AzureName == "" {
		schedule.Spec.AzureName = schedule.Name
	}
}

// defaultImpl applies the code generated defaults to the RedisPatchSchedule resource
func (schedule *RedisPatchSchedule) defaultImpl() { schedule.defaultAzureName() }

var _ genruntime.KubernetesResource = &RedisPatchSchedule{}

// AzureName returns the Azure name of the resource
func (schedule *RedisPatchSchedule) AzureName() string {
	return schedule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (schedule RedisPatchSchedule) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (schedule *RedisPatchSchedule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (schedule *RedisPatchSchedule) GetSpec() genruntime.ConvertibleSpec {
	return &schedule.Spec
}

// GetStatus returns the status of this resource
func (schedule *RedisPatchSchedule) GetStatus() genruntime.ConvertibleStatus {
	return &schedule.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (schedule *RedisPatchSchedule) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (schedule *RedisPatchSchedule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisPatchSchedule_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (schedule *RedisPatchSchedule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(schedule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  schedule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (schedule *RedisPatchSchedule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisPatchSchedule_STATUS); ok {
		schedule.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisPatchSchedule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	schedule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-cache-azure-com-v1beta20201201-redispatchschedule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redispatchschedules,verbs=create;update,versions=v1beta20201201,name=validate.v1beta20201201.redispatchschedules.cache.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &RedisPatchSchedule{}

// ValidateCreate validates the creation of the resource
func (schedule *RedisPatchSchedule) ValidateCreate() error {
	validations := schedule.createValidations()
	var temp interface{} = schedule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (schedule *RedisPatchSchedule) ValidateDelete() error {
	validations := schedule.deleteValidations()
	var temp interface{} = schedule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (schedule *RedisPatchSchedule) ValidateUpdate(old runtime.Object) error {
	validations := schedule.updateValidations()
	var temp interface{} = schedule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (schedule *RedisPatchSchedule) createValidations() []func() error {
	return []func() error{schedule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (schedule *RedisPatchSchedule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (schedule *RedisPatchSchedule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return schedule.validateResourceReferences()
		},
		schedule.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (schedule *RedisPatchSchedule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&schedule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (schedule *RedisPatchSchedule) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*RedisPatchSchedule)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, schedule)
}

// AssignPropertiesFromRedisPatchSchedule populates our RedisPatchSchedule from the provided source RedisPatchSchedule
func (schedule *RedisPatchSchedule) AssignPropertiesFromRedisPatchSchedule(source *v20201201s.RedisPatchSchedule) error {

	// ObjectMeta
	schedule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisPatchSchedule_Spec
	err := spec.AssignPropertiesFromRedisPatchSchedule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisPatchSchedule_Spec() to populate field Spec")
	}
	schedule.Spec = spec

	// Status
	var status RedisPatchSchedule_STATUS
	err = status.AssignPropertiesFromRedisPatchSchedule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisPatchSchedule_STATUS() to populate field Status")
	}
	schedule.Status = status

	// No error
	return nil
}

// AssignPropertiesToRedisPatchSchedule populates the provided destination RedisPatchSchedule from our RedisPatchSchedule
func (schedule *RedisPatchSchedule) AssignPropertiesToRedisPatchSchedule(destination *v20201201s.RedisPatchSchedule) error {

	// ObjectMeta
	destination.ObjectMeta = *schedule.ObjectMeta.DeepCopy()

	// Spec
	var spec v20201201s.RedisPatchSchedule_Spec
	err := schedule.Spec.AssignPropertiesToRedisPatchSchedule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisPatchSchedule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20201201s.RedisPatchSchedule_STATUS
	err = schedule.Status.AssignPropertiesToRedisPatchSchedule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisPatchSchedule_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (schedule *RedisPatchSchedule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: schedule.Spec.OriginalVersion(),
		Kind:    "RedisPatchSchedule",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2020-12-01/redis.json
type RedisPatchScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisPatchSchedule `json:"items"`
}

type RedisPatchSchedule_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// ScheduleEntries: List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntry_STATUS `json:"scheduleEntries,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisPatchSchedule_STATUS{}

// ConvertStatusFrom populates our RedisPatchSchedule_STATUS from the provided source
func (schedule *RedisPatchSchedule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20201201s.RedisPatchSchedule_STATUS)
	if ok {
		// Populate our instance from source
		return schedule.AssignPropertiesFromRedisPatchSchedule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20201201s.RedisPatchSchedule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = schedule.AssignPropertiesFromRedisPatchSchedule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisPatchSchedule_STATUS
func (schedule *RedisPatchSchedule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20201201s.RedisPatchSchedule_STATUS)
	if ok {
		// Populate destination from our instance
		return schedule.AssignPropertiesToRedisPatchSchedule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20201201s.RedisPatchSchedule_STATUS{}
	err := schedule.AssignPropertiesToRedisPatchSchedule_STATUS(dst)
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

var _ genruntime.FromARMConverter = &RedisPatchSchedule_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (schedule *RedisPatchSchedule_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisPatchSchedule_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (schedule *RedisPatchSchedule_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisPatchSchedule_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisPatchSchedule_STATUSARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		schedule.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		schedule.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		schedule.Name = &name
	}

	// Set property ‘ScheduleEntries’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.ScheduleEntries {
			var item1 ScheduleEntry_STATUS
			err := item1.PopulateFromARM(owner, item)
			if err != nil {
				return err
			}
			schedule.ScheduleEntries = append(schedule.ScheduleEntries, item1)
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		schedule.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromRedisPatchSchedule_STATUS populates our RedisPatchSchedule_STATUS from the provided source RedisPatchSchedule_STATUS
func (schedule *RedisPatchSchedule_STATUS) AssignPropertiesFromRedisPatchSchedule_STATUS(source *v20201201s.RedisPatchSchedule_STATUS) error {

	// Conditions
	schedule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	schedule.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	schedule.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	schedule.Name = genruntime.ClonePointerToString(source.Name)

	// ScheduleEntries
	if source.ScheduleEntries != nil {
		scheduleEntryList := make([]ScheduleEntry_STATUS, len(source.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range source.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry ScheduleEntry_STATUS
			err := scheduleEntry.AssignPropertiesFromScheduleEntry_STATUS(&scheduleEntryItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromScheduleEntry_STATUS() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		schedule.ScheduleEntries = scheduleEntryList
	} else {
		schedule.ScheduleEntries = nil
	}

	// Type
	schedule.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToRedisPatchSchedule_STATUS populates the provided destination RedisPatchSchedule_STATUS from our RedisPatchSchedule_STATUS
func (schedule *RedisPatchSchedule_STATUS) AssignPropertiesToRedisPatchSchedule_STATUS(destination *v20201201s.RedisPatchSchedule_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(schedule.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(schedule.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(schedule.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(schedule.Name)

	// ScheduleEntries
	if schedule.ScheduleEntries != nil {
		scheduleEntryList := make([]v20201201s.ScheduleEntry_STATUS, len(schedule.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range schedule.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry v20201201s.ScheduleEntry_STATUS
			err := scheduleEntryItem.AssignPropertiesToScheduleEntry_STATUS(&scheduleEntry)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToScheduleEntry_STATUS() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		destination.ScheduleEntries = scheduleEntryList
	} else {
		destination.ScheduleEntries = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(schedule.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type RedisPatchSchedule_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// +kubebuilder:validation:Required
	// ScheduleEntries: List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntry `json:"scheduleEntries,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMTransformer = &RedisPatchSchedule_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (schedule *RedisPatchSchedule_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if schedule == nil {
		return nil, nil
	}
	result := &RedisPatchSchedule_SpecARM{}

	// Set property ‘AzureName’:
	result.AzureName = schedule.AzureName

	// Set property ‘Id’:
	if schedule.Id != nil {
		id := *schedule.Id
		result.Id = &id
	}

	// Set property ‘Location’:
	if schedule.Location != nil {
		location := *schedule.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if schedule.ScheduleEntries != nil {
		result.Properties = &ScheduleEntriesARM{}
	}
	for _, item := range schedule.ScheduleEntries {
		itemARM, err := item.ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		result.Properties.ScheduleEntries = append(result.Properties.ScheduleEntries, *itemARM.(*ScheduleEntryARM))
	}

	// Set property ‘Type’:
	if schedule.Type != nil {
		typeVar := *schedule.Type
		result.Type = &typeVar
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (schedule *RedisPatchSchedule_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisPatchSchedule_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (schedule *RedisPatchSchedule_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisPatchSchedule_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisPatchSchedule_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	schedule.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		schedule.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		schedule.Location = &location
	}

	// Set property ‘Owner’:
	schedule.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘ScheduleEntries’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.ScheduleEntries {
			var item1 ScheduleEntry
			err := item1.PopulateFromARM(owner, item)
			if err != nil {
				return err
			}
			schedule.ScheduleEntries = append(schedule.ScheduleEntries, item1)
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		schedule.Type = &typeVar
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RedisPatchSchedule_Spec{}

// ConvertSpecFrom populates our RedisPatchSchedule_Spec from the provided source
func (schedule *RedisPatchSchedule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20201201s.RedisPatchSchedule_Spec)
	if ok {
		// Populate our instance from source
		return schedule.AssignPropertiesFromRedisPatchSchedule_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20201201s.RedisPatchSchedule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = schedule.AssignPropertiesFromRedisPatchSchedule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisPatchSchedule_Spec
func (schedule *RedisPatchSchedule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20201201s.RedisPatchSchedule_Spec)
	if ok {
		// Populate destination from our instance
		return schedule.AssignPropertiesToRedisPatchSchedule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20201201s.RedisPatchSchedule_Spec{}
	err := schedule.AssignPropertiesToRedisPatchSchedule_Spec(dst)
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

// AssignPropertiesFromRedisPatchSchedule_Spec populates our RedisPatchSchedule_Spec from the provided source RedisPatchSchedule_Spec
func (schedule *RedisPatchSchedule_Spec) AssignPropertiesFromRedisPatchSchedule_Spec(source *v20201201s.RedisPatchSchedule_Spec) error {

	// AzureName
	schedule.AzureName = source.AzureName

	// Id
	schedule.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	schedule.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		schedule.Owner = &owner
	} else {
		schedule.Owner = nil
	}

	// ScheduleEntries
	if source.ScheduleEntries != nil {
		scheduleEntryList := make([]ScheduleEntry, len(source.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range source.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry ScheduleEntry
			err := scheduleEntry.AssignPropertiesFromScheduleEntry(&scheduleEntryItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromScheduleEntry() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		schedule.ScheduleEntries = scheduleEntryList
	} else {
		schedule.ScheduleEntries = nil
	}

	// Type
	schedule.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToRedisPatchSchedule_Spec populates the provided destination RedisPatchSchedule_Spec from our RedisPatchSchedule_Spec
func (schedule *RedisPatchSchedule_Spec) AssignPropertiesToRedisPatchSchedule_Spec(destination *v20201201s.RedisPatchSchedule_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = schedule.AzureName

	// Id
	destination.Id = genruntime.ClonePointerToString(schedule.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(schedule.Location)

	// OriginalVersion
	destination.OriginalVersion = schedule.OriginalVersion()

	// Owner
	if schedule.Owner != nil {
		owner := schedule.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// ScheduleEntries
	if schedule.ScheduleEntries != nil {
		scheduleEntryList := make([]v20201201s.ScheduleEntry, len(schedule.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range schedule.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry v20201201s.ScheduleEntry
			err := scheduleEntryItem.AssignPropertiesToScheduleEntry(&scheduleEntry)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToScheduleEntry() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		destination.ScheduleEntries = scheduleEntryList
	} else {
		destination.ScheduleEntries = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(schedule.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (schedule *RedisPatchSchedule_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (schedule *RedisPatchSchedule_Spec) SetAzureName(azureName string) {
	schedule.AzureName = azureName
}

type ScheduleEntry struct {
	// +kubebuilder:validation:Required
	// DayOfWeek: Day of the week when a cache can be patched.
	DayOfWeek *ScheduleEntry_DayOfWeek `json:"dayOfWeek,omitempty"`

	// MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can take.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty"`

	// +kubebuilder:validation:Required
	// StartHourUtc: Start hour after which cache patching can start.
	StartHourUtc *int `json:"startHourUtc,omitempty"`
}

var _ genruntime.ARMTransformer = &ScheduleEntry{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (entry *ScheduleEntry) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if entry == nil {
		return nil, nil
	}
	result := &ScheduleEntryARM{}

	// Set property ‘DayOfWeek’:
	if entry.DayOfWeek != nil {
		dayOfWeek := *entry.DayOfWeek
		result.DayOfWeek = &dayOfWeek
	}

	// Set property ‘MaintenanceWindow’:
	if entry.MaintenanceWindow != nil {
		maintenanceWindow := *entry.MaintenanceWindow
		result.MaintenanceWindow = &maintenanceWindow
	}

	// Set property ‘StartHourUtc’:
	if entry.StartHourUtc != nil {
		startHourUtc := *entry.StartHourUtc
		result.StartHourUtc = &startHourUtc
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (entry *ScheduleEntry) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ScheduleEntryARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (entry *ScheduleEntry) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ScheduleEntryARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ScheduleEntryARM, got %T", armInput)
	}

	// Set property ‘DayOfWeek’:
	if typedInput.DayOfWeek != nil {
		dayOfWeek := *typedInput.DayOfWeek
		entry.DayOfWeek = &dayOfWeek
	}

	// Set property ‘MaintenanceWindow’:
	if typedInput.MaintenanceWindow != nil {
		maintenanceWindow := *typedInput.MaintenanceWindow
		entry.MaintenanceWindow = &maintenanceWindow
	}

	// Set property ‘StartHourUtc’:
	if typedInput.StartHourUtc != nil {
		startHourUtc := *typedInput.StartHourUtc
		entry.StartHourUtc = &startHourUtc
	}

	// No error
	return nil
}

// AssignPropertiesFromScheduleEntry populates our ScheduleEntry from the provided source ScheduleEntry
func (entry *ScheduleEntry) AssignPropertiesFromScheduleEntry(source *v20201201s.ScheduleEntry) error {

	// DayOfWeek
	if source.DayOfWeek != nil {
		dayOfWeek := ScheduleEntry_DayOfWeek(*source.DayOfWeek)
		entry.DayOfWeek = &dayOfWeek
	} else {
		entry.DayOfWeek = nil
	}

	// MaintenanceWindow
	if source.MaintenanceWindow != nil {
		maintenanceWindow := *source.MaintenanceWindow
		entry.MaintenanceWindow = &maintenanceWindow
	} else {
		entry.MaintenanceWindow = nil
	}

	// StartHourUtc
	entry.StartHourUtc = genruntime.ClonePointerToInt(source.StartHourUtc)

	// No error
	return nil
}

// AssignPropertiesToScheduleEntry populates the provided destination ScheduleEntry from our ScheduleEntry
func (entry *ScheduleEntry) AssignPropertiesToScheduleEntry(destination *v20201201s.ScheduleEntry) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// DayOfWeek
	if entry.DayOfWeek != nil {
		dayOfWeek := string(*entry.DayOfWeek)
		destination.DayOfWeek = &dayOfWeek
	} else {
		destination.DayOfWeek = nil
	}

	// MaintenanceWindow
	if entry.MaintenanceWindow != nil {
		maintenanceWindow := *entry.MaintenanceWindow
		destination.MaintenanceWindow = &maintenanceWindow
	} else {
		destination.MaintenanceWindow = nil
	}

	// StartHourUtc
	destination.StartHourUtc = genruntime.ClonePointerToInt(entry.StartHourUtc)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type ScheduleEntry_STATUS struct {
	// DayOfWeek: Day of the week when a cache can be patched.
	DayOfWeek *ScheduleEntry_DayOfWeek_STATUS `json:"dayOfWeek,omitempty"`

	// MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can take.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty"`

	// StartHourUtc: Start hour after which cache patching can start.
	StartHourUtc *int `json:"startHourUtc,omitempty"`
}

var _ genruntime.FromARMConverter = &ScheduleEntry_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (entry *ScheduleEntry_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ScheduleEntry_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (entry *ScheduleEntry_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ScheduleEntry_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ScheduleEntry_STATUSARM, got %T", armInput)
	}

	// Set property ‘DayOfWeek’:
	if typedInput.DayOfWeek != nil {
		dayOfWeek := *typedInput.DayOfWeek
		entry.DayOfWeek = &dayOfWeek
	}

	// Set property ‘MaintenanceWindow’:
	if typedInput.MaintenanceWindow != nil {
		maintenanceWindow := *typedInput.MaintenanceWindow
		entry.MaintenanceWindow = &maintenanceWindow
	}

	// Set property ‘StartHourUtc’:
	if typedInput.StartHourUtc != nil {
		startHourUtc := *typedInput.StartHourUtc
		entry.StartHourUtc = &startHourUtc
	}

	// No error
	return nil
}

// AssignPropertiesFromScheduleEntry_STATUS populates our ScheduleEntry_STATUS from the provided source ScheduleEntry_STATUS
func (entry *ScheduleEntry_STATUS) AssignPropertiesFromScheduleEntry_STATUS(source *v20201201s.ScheduleEntry_STATUS) error {

	// DayOfWeek
	if source.DayOfWeek != nil {
		dayOfWeek := ScheduleEntry_DayOfWeek_STATUS(*source.DayOfWeek)
		entry.DayOfWeek = &dayOfWeek
	} else {
		entry.DayOfWeek = nil
	}

	// MaintenanceWindow
	entry.MaintenanceWindow = genruntime.ClonePointerToString(source.MaintenanceWindow)

	// StartHourUtc
	entry.StartHourUtc = genruntime.ClonePointerToInt(source.StartHourUtc)

	// No error
	return nil
}

// AssignPropertiesToScheduleEntry_STATUS populates the provided destination ScheduleEntry_STATUS from our ScheduleEntry_STATUS
func (entry *ScheduleEntry_STATUS) AssignPropertiesToScheduleEntry_STATUS(destination *v20201201s.ScheduleEntry_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// DayOfWeek
	if entry.DayOfWeek != nil {
		dayOfWeek := string(*entry.DayOfWeek)
		destination.DayOfWeek = &dayOfWeek
	} else {
		destination.DayOfWeek = nil
	}

	// MaintenanceWindow
	destination.MaintenanceWindow = genruntime.ClonePointerToString(entry.MaintenanceWindow)

	// StartHourUtc
	destination.StartHourUtc = genruntime.ClonePointerToInt(entry.StartHourUtc)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"Everyday","Friday","Monday","Saturday","Sunday","Thursday","Tuesday","Wednesday","Weekend"}
type ScheduleEntry_DayOfWeek string

const (
	ScheduleEntry_DayOfWeek_Everyday  = ScheduleEntry_DayOfWeek("Everyday")
	ScheduleEntry_DayOfWeek_Friday    = ScheduleEntry_DayOfWeek("Friday")
	ScheduleEntry_DayOfWeek_Monday    = ScheduleEntry_DayOfWeek("Monday")
	ScheduleEntry_DayOfWeek_Saturday  = ScheduleEntry_DayOfWeek("Saturday")
	ScheduleEntry_DayOfWeek_Sunday    = ScheduleEntry_DayOfWeek("Sunday")
	ScheduleEntry_DayOfWeek_Thursday  = ScheduleEntry_DayOfWeek("Thursday")
	ScheduleEntry_DayOfWeek_Tuesday   = ScheduleEntry_DayOfWeek("Tuesday")
	ScheduleEntry_DayOfWeek_Wednesday = ScheduleEntry_DayOfWeek("Wednesday")
	ScheduleEntry_DayOfWeek_Weekend   = ScheduleEntry_DayOfWeek("Weekend")
)

type ScheduleEntry_DayOfWeek_STATUS string

const (
	ScheduleEntry_DayOfWeek_Everyday_STATUS  = ScheduleEntry_DayOfWeek_STATUS("Everyday")
	ScheduleEntry_DayOfWeek_Friday_STATUS    = ScheduleEntry_DayOfWeek_STATUS("Friday")
	ScheduleEntry_DayOfWeek_Monday_STATUS    = ScheduleEntry_DayOfWeek_STATUS("Monday")
	ScheduleEntry_DayOfWeek_Saturday_STATUS  = ScheduleEntry_DayOfWeek_STATUS("Saturday")
	ScheduleEntry_DayOfWeek_Sunday_STATUS    = ScheduleEntry_DayOfWeek_STATUS("Sunday")
	ScheduleEntry_DayOfWeek_Thursday_STATUS  = ScheduleEntry_DayOfWeek_STATUS("Thursday")
	ScheduleEntry_DayOfWeek_Tuesday_STATUS   = ScheduleEntry_DayOfWeek_STATUS("Tuesday")
	ScheduleEntry_DayOfWeek_Wednesday_STATUS = ScheduleEntry_DayOfWeek_STATUS("Wednesday")
	ScheduleEntry_DayOfWeek_Weekend_STATUS   = ScheduleEntry_DayOfWeek_STATUS("Weekend")
)

func init() {
	SchemeBuilder.Register(&RedisPatchSchedule{}, &RedisPatchScheduleList{})
}
