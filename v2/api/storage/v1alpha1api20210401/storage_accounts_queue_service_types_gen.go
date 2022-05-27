// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import (
	"fmt"
	alpha20210401s "github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401storage"
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
// Deprecated version of StorageAccountsQueueService. Use v1beta20210401.StorageAccountsQueueService instead
type StorageAccountsQueueService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsQueueService_Spec   `json:"spec,omitempty"`
	Status            StorageAccountsQueueService_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsQueueService{}

// GetConditions returns the conditions of the resource
func (service *StorageAccountsQueueService) GetConditions() conditions.Conditions {
	return service.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (service *StorageAccountsQueueService) SetConditions(conditions conditions.Conditions) {
	service.Status.Conditions = conditions
}

var _ conversion.Convertible = &StorageAccountsQueueService{}

// ConvertFrom populates our StorageAccountsQueueService from the provided hub StorageAccountsQueueService
func (service *StorageAccountsQueueService) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20210401s.StorageAccountsQueueService

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = service.AssignPropertiesFromStorageAccountsQueueService(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to service")
	}

	return nil
}

// ConvertTo populates the provided hub StorageAccountsQueueService from our StorageAccountsQueueService
func (service *StorageAccountsQueueService) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20210401s.StorageAccountsQueueService
	err := service.AssignPropertiesToStorageAccountsQueueService(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from service")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-storage-azure-com-v1alpha1api20210401-storageaccountsqueueservice,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsqueueservices,verbs=create;update,versions=v1alpha1api20210401,name=default.v1alpha1api20210401.storageaccountsqueueservices.storage.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &StorageAccountsQueueService{}

// Default applies defaults to the StorageAccountsQueueService resource
func (service *StorageAccountsQueueService) Default() {
	service.defaultImpl()
	var temp interface{} = service
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (service *StorageAccountsQueueService) defaultAzureName() {
	if service.Spec.AzureName == "" {
		service.Spec.AzureName = service.Name
	}
}

// defaultImpl applies the code generated defaults to the StorageAccountsQueueService resource
func (service *StorageAccountsQueueService) defaultImpl() { service.defaultAzureName() }

var _ genruntime.KubernetesResource = &StorageAccountsQueueService{}

// AzureName returns the Azure name of the resource
func (service *StorageAccountsQueueService) AzureName() string {
	return service.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "20210401"
func (service StorageAccountsQueueService) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (service *StorageAccountsQueueService) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (service *StorageAccountsQueueService) GetSpec() genruntime.ConvertibleSpec {
	return &service.Spec
}

// GetStatus returns the status of this resource
func (service *StorageAccountsQueueService) GetStatus() genruntime.ConvertibleStatus {
	return &service.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (service *StorageAccountsQueueService) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (service *StorageAccountsQueueService) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccountsQueueService_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (service *StorageAccountsQueueService) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(service.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  service.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (service *StorageAccountsQueueService) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccountsQueueService_STATUS); ok {
		service.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccountsQueueService_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	service.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-storage-azure-com-v1alpha1api20210401-storageaccountsqueueservice,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsqueueservices,verbs=create;update,versions=v1alpha1api20210401,name=validate.v1alpha1api20210401.storageaccountsqueueservices.storage.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &StorageAccountsQueueService{}

// ValidateCreate validates the creation of the resource
func (service *StorageAccountsQueueService) ValidateCreate() error {
	validations := service.createValidations()
	var temp interface{} = service
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
func (service *StorageAccountsQueueService) ValidateDelete() error {
	validations := service.deleteValidations()
	var temp interface{} = service
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
func (service *StorageAccountsQueueService) ValidateUpdate(old runtime.Object) error {
	validations := service.updateValidations()
	var temp interface{} = service
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
func (service *StorageAccountsQueueService) createValidations() []func() error {
	return []func() error{service.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (service *StorageAccountsQueueService) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (service *StorageAccountsQueueService) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return service.validateResourceReferences()
		},
		service.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (service *StorageAccountsQueueService) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&service.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (service *StorageAccountsQueueService) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*StorageAccountsQueueService)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, service)
}

// AssignPropertiesFromStorageAccountsQueueService populates our StorageAccountsQueueService from the provided source StorageAccountsQueueService
func (service *StorageAccountsQueueService) AssignPropertiesFromStorageAccountsQueueService(source *alpha20210401s.StorageAccountsQueueService) error {

	// ObjectMeta
	service.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec StorageAccountsQueueService_Spec
	err := spec.AssignPropertiesFromStorageAccountsQueueService_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromStorageAccountsQueueService_Spec() to populate field Spec")
	}
	service.Spec = spec

	// Status
	var status StorageAccountsQueueService_STATUS
	err = status.AssignPropertiesFromStorageAccountsQueueService_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromStorageAccountsQueueService_STATUS() to populate field Status")
	}
	service.Status = status

	// No error
	return nil
}

// AssignPropertiesToStorageAccountsQueueService populates the provided destination StorageAccountsQueueService from our StorageAccountsQueueService
func (service *StorageAccountsQueueService) AssignPropertiesToStorageAccountsQueueService(destination *alpha20210401s.StorageAccountsQueueService) error {

	// ObjectMeta
	destination.ObjectMeta = *service.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20210401s.StorageAccountsQueueService_Spec
	err := service.Spec.AssignPropertiesToStorageAccountsQueueService_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToStorageAccountsQueueService_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20210401s.StorageAccountsQueueService_STATUS
	err = service.Status.AssignPropertiesToStorageAccountsQueueService_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToStorageAccountsQueueService_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (service *StorageAccountsQueueService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: service.Spec.OriginalVersion(),
		Kind:    "StorageAccountsQueueService",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of StorageAccountsQueueService. Use v1beta20210401.StorageAccountsQueueService instead
type StorageAccountsQueueServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsQueueService `json:"items"`
}

// Deprecated version of StorageAccountsQueueService_STATUS. Use v1beta20210401.StorageAccountsQueueService_STATUS instead
type StorageAccountsQueueService_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`
	Cors       *CorsRules_STATUS      `json:"cors,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccountsQueueService_STATUS{}

// ConvertStatusFrom populates our StorageAccountsQueueService_STATUS from the provided source
func (service *StorageAccountsQueueService_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20210401s.StorageAccountsQueueService_STATUS)
	if ok {
		// Populate our instance from source
		return service.AssignPropertiesFromStorageAccountsQueueService_STATUS(src)
	}

	// Convert to an intermediate form
	src = &alpha20210401s.StorageAccountsQueueService_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = service.AssignPropertiesFromStorageAccountsQueueService_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our StorageAccountsQueueService_STATUS
func (service *StorageAccountsQueueService_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20210401s.StorageAccountsQueueService_STATUS)
	if ok {
		// Populate destination from our instance
		return service.AssignPropertiesToStorageAccountsQueueService_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210401s.StorageAccountsQueueService_STATUS{}
	err := service.AssignPropertiesToStorageAccountsQueueService_STATUS(dst)
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

var _ genruntime.FromARMConverter = &StorageAccountsQueueService_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (service *StorageAccountsQueueService_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &StorageAccountsQueueService_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (service *StorageAccountsQueueService_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(StorageAccountsQueueService_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected StorageAccountsQueueService_STATUSARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Cors’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Cors != nil {
			var cors1 CorsRules_STATUS
			err := cors1.PopulateFromARM(owner, *typedInput.Properties.Cors)
			if err != nil {
				return err
			}
			cors := cors1
			service.Cors = &cors
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		service.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		service.Name = &name
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		service.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromStorageAccountsQueueService_STATUS populates our StorageAccountsQueueService_STATUS from the provided source StorageAccountsQueueService_STATUS
func (service *StorageAccountsQueueService_STATUS) AssignPropertiesFromStorageAccountsQueueService_STATUS(source *alpha20210401s.StorageAccountsQueueService_STATUS) error {

	// Conditions
	service.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Cors
	if source.Cors != nil {
		var cor CorsRules_STATUS
		err := cor.AssignPropertiesFromCorsRules_STATUS(source.Cors)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCorsRules_STATUS() to populate field Cors")
		}
		service.Cors = &cor
	} else {
		service.Cors = nil
	}

	// Id
	service.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	service.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	service.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToStorageAccountsQueueService_STATUS populates the provided destination StorageAccountsQueueService_STATUS from our StorageAccountsQueueService_STATUS
func (service *StorageAccountsQueueService_STATUS) AssignPropertiesToStorageAccountsQueueService_STATUS(destination *alpha20210401s.StorageAccountsQueueService_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(service.Conditions)

	// Cors
	if service.Cors != nil {
		var cor alpha20210401s.CorsRules_STATUS
		err := service.Cors.AssignPropertiesToCorsRules_STATUS(&cor)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCorsRules_STATUS() to populate field Cors")
		}
		destination.Cors = &cor
	} else {
		destination.Cors = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(service.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(service.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(service.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type StorageAccountsQueueService_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string     `json:"azureName,omitempty"`
	Cors      *CorsRules `json:"cors,omitempty"`
	Id        *string    `json:"id,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	Type  *string                            `json:"type,omitempty"`
}

var _ genruntime.ARMTransformer = &StorageAccountsQueueService_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (service *StorageAccountsQueueService_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if service == nil {
		return nil, nil
	}
	result := &StorageAccountsQueueService_SpecARM{}

	// Set property ‘AzureName’:
	result.AzureName = service.AzureName

	// Set property ‘Id’:
	if service.Id != nil {
		id := *service.Id
		result.Id = &id
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if service.Cors != nil {
		result.Properties = &StorageAccountsQueueService_Spec_PropertiesARM{}
	}
	if service.Cors != nil {
		corsARM, err := (*service.Cors).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		cors := *corsARM.(*CorsRulesARM)
		result.Properties.Cors = &cors
	}

	// Set property ‘Type’:
	if service.Type != nil {
		typeVar := *service.Type
		result.Type = &typeVar
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (service *StorageAccountsQueueService_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &StorageAccountsQueueService_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (service *StorageAccountsQueueService_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(StorageAccountsQueueService_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected StorageAccountsQueueService_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	service.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Cors’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Cors != nil {
			var cors1 CorsRules
			err := cors1.PopulateFromARM(owner, *typedInput.Properties.Cors)
			if err != nil {
				return err
			}
			cors := cors1
			service.Cors = &cors
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		service.Id = &id
	}

	// Set property ‘Owner’:
	service.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		service.Type = &typeVar
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &StorageAccountsQueueService_Spec{}

// ConvertSpecFrom populates our StorageAccountsQueueService_Spec from the provided source
func (service *StorageAccountsQueueService_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20210401s.StorageAccountsQueueService_Spec)
	if ok {
		// Populate our instance from source
		return service.AssignPropertiesFromStorageAccountsQueueService_Spec(src)
	}

	// Convert to an intermediate form
	src = &alpha20210401s.StorageAccountsQueueService_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = service.AssignPropertiesFromStorageAccountsQueueService_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our StorageAccountsQueueService_Spec
func (service *StorageAccountsQueueService_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20210401s.StorageAccountsQueueService_Spec)
	if ok {
		// Populate destination from our instance
		return service.AssignPropertiesToStorageAccountsQueueService_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210401s.StorageAccountsQueueService_Spec{}
	err := service.AssignPropertiesToStorageAccountsQueueService_Spec(dst)
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

// AssignPropertiesFromStorageAccountsQueueService_Spec populates our StorageAccountsQueueService_Spec from the provided source StorageAccountsQueueService_Spec
func (service *StorageAccountsQueueService_Spec) AssignPropertiesFromStorageAccountsQueueService_Spec(source *alpha20210401s.StorageAccountsQueueService_Spec) error {

	// AzureName
	service.AzureName = source.AzureName

	// Cors
	if source.Cors != nil {
		var cor CorsRules
		err := cor.AssignPropertiesFromCorsRules(source.Cors)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCorsRules() to populate field Cors")
		}
		service.Cors = &cor
	} else {
		service.Cors = nil
	}

	// Id
	service.Id = genruntime.ClonePointerToString(source.Id)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		service.Owner = &owner
	} else {
		service.Owner = nil
	}

	// Type
	service.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToStorageAccountsQueueService_Spec populates the provided destination StorageAccountsQueueService_Spec from our StorageAccountsQueueService_Spec
func (service *StorageAccountsQueueService_Spec) AssignPropertiesToStorageAccountsQueueService_Spec(destination *alpha20210401s.StorageAccountsQueueService_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = service.AzureName

	// Cors
	if service.Cors != nil {
		var cor alpha20210401s.CorsRules
		err := service.Cors.AssignPropertiesToCorsRules(&cor)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCorsRules() to populate field Cors")
		}
		destination.Cors = &cor
	} else {
		destination.Cors = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(service.Id)

	// OriginalVersion
	destination.OriginalVersion = service.OriginalVersion()

	// Owner
	if service.Owner != nil {
		owner := service.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(service.Type)

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
func (service *StorageAccountsQueueService_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (service *StorageAccountsQueueService_Spec) SetAzureName(azureName string) {
	service.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&StorageAccountsQueueService{}, &StorageAccountsQueueServiceList{})
}
