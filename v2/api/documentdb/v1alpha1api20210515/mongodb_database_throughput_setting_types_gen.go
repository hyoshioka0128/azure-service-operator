// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"fmt"
	alpha20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515storage"
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
// Deprecated version of MongodbDatabaseThroughputSetting. Use v1beta20210515.MongodbDatabaseThroughputSetting instead
type MongodbDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsMongodbDatabasesThroughputSetting_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabaseThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *MongodbDatabaseThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *MongodbDatabaseThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ conversion.Convertible = &MongodbDatabaseThroughputSetting{}

// ConvertFrom populates our MongodbDatabaseThroughputSetting from the provided hub MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20210515s.MongodbDatabaseThroughputSetting

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = setting.AssignPropertiesFromMongodbDatabaseThroughputSetting(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to setting")
	}

	return nil
}

// ConvertTo populates the provided hub MongodbDatabaseThroughputSetting from our MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20210515s.MongodbDatabaseThroughputSetting
	err := setting.AssignPropertiesToMongodbDatabaseThroughputSetting(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from setting")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1alpha1api20210515-mongodbdatabasethroughputsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=mongodbdatabasethroughputsettings,verbs=create;update,versions=v1alpha1api20210515,name=default.v1alpha1api20210515.mongodbdatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &MongodbDatabaseThroughputSetting{}

// Default applies defaults to the MongodbDatabaseThroughputSetting resource
func (setting *MongodbDatabaseThroughputSetting) Default() {
	setting.defaultImpl()
	var temp interface{} = setting
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (setting *MongodbDatabaseThroughputSetting) defaultAzureName() {
	if setting.Spec.AzureName == "" {
		setting.Spec.AzureName = setting.Name
	}
}

// defaultImpl applies the code generated defaults to the MongodbDatabaseThroughputSetting resource
func (setting *MongodbDatabaseThroughputSetting) defaultImpl() { setting.defaultAzureName() }

var _ genruntime.KubernetesResource = &MongodbDatabaseThroughputSetting{}

// AzureName returns the Azure name of the resource
func (setting *MongodbDatabaseThroughputSetting) AzureName() string {
	return setting.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "20210515"
func (setting MongodbDatabaseThroughputSetting) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (setting *MongodbDatabaseThroughputSetting) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (setting *MongodbDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *MongodbDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (setting *MongodbDatabaseThroughputSetting) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *MongodbDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (setting *MongodbDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  setting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (setting *MongodbDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1alpha1api20210515-mongodbdatabasethroughputsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=mongodbdatabasethroughputsettings,verbs=create;update,versions=v1alpha1api20210515,name=validate.v1alpha1api20210515.mongodbdatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &MongodbDatabaseThroughputSetting{}

// ValidateCreate validates the creation of the resource
func (setting *MongodbDatabaseThroughputSetting) ValidateCreate() error {
	validations := setting.createValidations()
	var temp interface{} = setting
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
func (setting *MongodbDatabaseThroughputSetting) ValidateDelete() error {
	validations := setting.deleteValidations()
	var temp interface{} = setting
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
func (setting *MongodbDatabaseThroughputSetting) ValidateUpdate(old runtime.Object) error {
	validations := setting.updateValidations()
	var temp interface{} = setting
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
func (setting *MongodbDatabaseThroughputSetting) createValidations() []func() error {
	return []func() error{setting.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (setting *MongodbDatabaseThroughputSetting) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (setting *MongodbDatabaseThroughputSetting) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return setting.validateResourceReferences()
		},
		setting.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (setting *MongodbDatabaseThroughputSetting) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&setting.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (setting *MongodbDatabaseThroughputSetting) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*MongodbDatabaseThroughputSetting)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, setting)
}

// AssignPropertiesFromMongodbDatabaseThroughputSetting populates our MongodbDatabaseThroughputSetting from the provided source MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) AssignPropertiesFromMongodbDatabaseThroughputSetting(source *alpha20210515s.MongodbDatabaseThroughputSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsMongodbDatabasesThroughputSetting_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSetting_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSetting_Spec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS
	err = status.AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSetting_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSetting_STATUS() to populate field Status")
	}
	setting.Status = status

	// No error
	return nil
}

// AssignPropertiesToMongodbDatabaseThroughputSetting populates the provided destination MongodbDatabaseThroughputSetting from our MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) AssignPropertiesToMongodbDatabaseThroughputSetting(destination *alpha20210515s.MongodbDatabaseThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_Spec
	err := setting.Spec.AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSetting_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSetting_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS
	err = setting.Status.AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSetting_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSetting_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *MongodbDatabaseThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion(),
		Kind:    "MongodbDatabaseThroughputSetting",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of MongodbDatabaseThroughputSetting. Use v1beta20210515.MongodbDatabaseThroughputSetting instead
type MongodbDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseThroughputSetting `json:"items"`
}

// Deprecated version of DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS. Use v1beta20210515.DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS instead
type DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition             `json:"conditions,omitempty"`
	Id         *string                            `json:"id,omitempty"`
	Location   *string                            `json:"location,omitempty"`
	Name       *string                            `json:"name,omitempty"`
	Resource   *ThroughputSettingsResource_STATUS `json:"resource,omitempty"`
	Tags       map[string]string                  `json:"tags,omitempty"`
	Type       *string                            `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS{}

// ConvertStatusFrom populates our DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS from the provided source
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS)
	if ok {
		// Populate our instance from source
		return setting.AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSetting_STATUS(src)
	}

	// Convert to an intermediate form
	src = &alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = setting.AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSetting_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS)
	if ok {
		// Populate destination from our instance
		return setting.AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSetting_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS{}
	err := setting.AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSetting_STATUS(dst)
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

var _ genruntime.FromARMConverter = &DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsMongodbDatabasesThroughputSetting_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsMongodbDatabasesThroughputSetting_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsMongodbDatabasesThroughputSetting_STATUSARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		setting.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		setting.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		setting.Name = &name
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 ThroughputSettingsResource_STATUS
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			setting.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		setting.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			setting.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		setting.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSetting_STATUS populates our DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS from the provided source DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS) AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSetting_STATUS(source *alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS) error {

	// Conditions
	setting.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	setting.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	setting.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource_STATUS
		err := resource.AssignPropertiesFromThroughputSettingsResource_STATUS(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromThroughputSettingsResource_STATUS() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	setting.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSetting_STATUS populates the provided destination DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS from our DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS) AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSetting_STATUS(destination *alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(setting.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(setting.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(setting.Name)

	// Resource
	if setting.Resource != nil {
		var resource alpha20210515s.ThroughputSettingsResource_STATUS
		err := setting.Resource.AssignPropertiesToThroughputSettingsResource_STATUS(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToThroughputSettingsResource_STATUS() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(setting.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(setting.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type DatabaseAccountsMongodbDatabasesThroughputSetting_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string  `json:"azureName,omitempty"`
	Id        *string `json:"id,omitempty"`
	Location  *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// +kubebuilder:validation:Required
	Resource *ThroughputSettingsResource `json:"resource,omitempty"`
	Tags     map[string]string           `json:"tags,omitempty"`
	Type     *string                     `json:"type,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsMongodbDatabasesThroughputSetting_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if setting == nil {
		return nil, nil
	}
	result := &DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM{}

	// Set property ‘AzureName’:
	result.AzureName = setting.AzureName

	// Set property ‘Id’:
	if setting.Id != nil {
		id := *setting.Id
		result.Id = &id
	}

	// Set property ‘Location’:
	if setting.Location != nil {
		location := *setting.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if setting.Resource != nil {
		result.Properties = &ThroughputSettingsUpdatePropertiesARM{}
	}
	if setting.Resource != nil {
		resourceARM, err := (*setting.Resource).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		resource := *resourceARM.(*ThroughputSettingsResourceARM)
		result.Properties.Resource = &resource
	}

	// Set property ‘Tags’:
	if setting.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range setting.Tags {
			result.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if setting.Type != nil {
		typeVar := *setting.Type
		result.Type = &typeVar
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	setting.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		setting.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		setting.Location = &location
	}

	// Set property ‘Owner’:
	setting.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 ThroughputSettingsResource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			setting.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		setting.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			setting.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		setting.Type = &typeVar
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsMongodbDatabasesThroughputSetting_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsMongodbDatabasesThroughputSetting_Spec from the provided source
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_Spec)
	if ok {
		// Populate our instance from source
		return setting.AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSetting_Spec(src)
	}

	// Convert to an intermediate form
	src = &alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = setting.AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSetting_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsMongodbDatabasesThroughputSetting_Spec
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_Spec)
	if ok {
		// Populate destination from our instance
		return setting.AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSetting_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_Spec{}
	err := setting.AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSetting_Spec(dst)
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

// AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSetting_Spec populates our DatabaseAccountsMongodbDatabasesThroughputSetting_Spec from the provided source DatabaseAccountsMongodbDatabasesThroughputSetting_Spec
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_Spec) AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSetting_Spec(source *alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_Spec) error {

	// AzureName
	setting.AzureName = source.AzureName

	// Id
	setting.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		setting.Owner = &owner
	} else {
		setting.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.AssignPropertiesFromThroughputSettingsResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromThroughputSettingsResource() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	setting.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSetting_Spec populates the provided destination DatabaseAccountsMongodbDatabasesThroughputSetting_Spec from our DatabaseAccountsMongodbDatabasesThroughputSetting_Spec
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_Spec) AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSetting_Spec(destination *alpha20210515s.DatabaseAccountsMongodbDatabasesThroughputSetting_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = setting.AzureName

	// Id
	destination.Id = genruntime.ClonePointerToString(setting.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// OriginalVersion
	destination.OriginalVersion = setting.OriginalVersion()

	// Owner
	if setting.Owner != nil {
		owner := setting.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if setting.Resource != nil {
		var resource alpha20210515s.ThroughputSettingsResource
		err := setting.Resource.AssignPropertiesToThroughputSettingsResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToThroughputSettingsResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(setting.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(setting.Type)

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
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (setting *DatabaseAccountsMongodbDatabasesThroughputSetting_Spec) SetAzureName(azureName string) {
	setting.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseThroughputSetting{}, &MongodbDatabaseThroughputSettingList{})
}
