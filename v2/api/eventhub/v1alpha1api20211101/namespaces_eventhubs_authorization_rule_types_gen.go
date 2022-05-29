// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import (
	"fmt"
	alpha20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1alpha1api20211101storage"
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
// Deprecated version of NamespacesEventhubsAuthorizationRule. Use v1beta20211101.NamespacesEventhubsAuthorizationRule instead
type NamespacesEventhubsAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesEventhubsAuthorizationRule_Spec   `json:"spec,omitempty"`
	Status            NamespacesEventhubsAuthorizationRule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesEventhubsAuthorizationRule{}

// GetConditions returns the conditions of the resource
func (rule *NamespacesEventhubsAuthorizationRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *NamespacesEventhubsAuthorizationRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesEventhubsAuthorizationRule{}

// ConvertFrom populates our NamespacesEventhubsAuthorizationRule from the provided hub NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20211101s.NamespacesEventhubsAuthorizationRule

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = rule.AssignPropertiesFromNamespacesEventhubsAuthorizationRule(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to rule")
	}

	return nil
}

// ConvertTo populates the provided hub NamespacesEventhubsAuthorizationRule from our NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20211101s.NamespacesEventhubsAuthorizationRule
	err := rule.AssignPropertiesToNamespacesEventhubsAuthorizationRule(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from rule")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-eventhub-azure-com-v1alpha1api20211101-namespaceseventhubsauthorizationrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsauthorizationrules,verbs=create;update,versions=v1alpha1api20211101,name=default.v1alpha1api20211101.namespaceseventhubsauthorizationrules.eventhub.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &NamespacesEventhubsAuthorizationRule{}

// Default applies defaults to the NamespacesEventhubsAuthorizationRule resource
func (rule *NamespacesEventhubsAuthorizationRule) Default() {
	rule.defaultImpl()
	var temp interface{} = rule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (rule *NamespacesEventhubsAuthorizationRule) defaultAzureName() {
	if rule.Spec.AzureName == "" {
		rule.Spec.AzureName = rule.Name
	}
}

// defaultImpl applies the code generated defaults to the NamespacesEventhubsAuthorizationRule resource
func (rule *NamespacesEventhubsAuthorizationRule) defaultImpl() { rule.defaultAzureName() }

var _ genruntime.KubernetesResource = &NamespacesEventhubsAuthorizationRule{}

// AzureName returns the Azure name of the resource
func (rule *NamespacesEventhubsAuthorizationRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule NamespacesEventhubsAuthorizationRule) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (rule *NamespacesEventhubsAuthorizationRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (rule *NamespacesEventhubsAuthorizationRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *NamespacesEventhubsAuthorizationRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (rule *NamespacesEventhubsAuthorizationRule) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *NamespacesEventhubsAuthorizationRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &NamespacesEventhubsAuthorizationRule_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rule *NamespacesEventhubsAuthorizationRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *NamespacesEventhubsAuthorizationRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*NamespacesEventhubsAuthorizationRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st NamespacesEventhubsAuthorizationRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-eventhub-azure-com-v1alpha1api20211101-namespaceseventhubsauthorizationrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsauthorizationrules,verbs=create;update,versions=v1alpha1api20211101,name=validate.v1alpha1api20211101.namespaceseventhubsauthorizationrules.eventhub.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &NamespacesEventhubsAuthorizationRule{}

// ValidateCreate validates the creation of the resource
func (rule *NamespacesEventhubsAuthorizationRule) ValidateCreate() error {
	validations := rule.createValidations()
	var temp interface{} = rule
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
func (rule *NamespacesEventhubsAuthorizationRule) ValidateDelete() error {
	validations := rule.deleteValidations()
	var temp interface{} = rule
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
func (rule *NamespacesEventhubsAuthorizationRule) ValidateUpdate(old runtime.Object) error {
	validations := rule.updateValidations()
	var temp interface{} = rule
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
func (rule *NamespacesEventhubsAuthorizationRule) createValidations() []func() error {
	return []func() error{rule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (rule *NamespacesEventhubsAuthorizationRule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (rule *NamespacesEventhubsAuthorizationRule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return rule.validateResourceReferences()
		},
		rule.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (rule *NamespacesEventhubsAuthorizationRule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *NamespacesEventhubsAuthorizationRule) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*NamespacesEventhubsAuthorizationRule)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, rule)
}

// AssignPropertiesFromNamespacesEventhubsAuthorizationRule populates our NamespacesEventhubsAuthorizationRule from the provided source NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) AssignPropertiesFromNamespacesEventhubsAuthorizationRule(source *alpha20211101s.NamespacesEventhubsAuthorizationRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NamespacesEventhubsAuthorizationRule_Spec
	err := spec.AssignPropertiesFromNamespacesEventhubsAuthorizationRule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromNamespacesEventhubsAuthorizationRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status NamespacesEventhubsAuthorizationRule_STATUS
	err = status.AssignPropertiesFromNamespacesEventhubsAuthorizationRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromNamespacesEventhubsAuthorizationRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignPropertiesToNamespacesEventhubsAuthorizationRule populates the provided destination NamespacesEventhubsAuthorizationRule from our NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) AssignPropertiesToNamespacesEventhubsAuthorizationRule(destination *alpha20211101s.NamespacesEventhubsAuthorizationRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20211101s.NamespacesEventhubsAuthorizationRule_Spec
	err := rule.Spec.AssignPropertiesToNamespacesEventhubsAuthorizationRule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToNamespacesEventhubsAuthorizationRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20211101s.NamespacesEventhubsAuthorizationRule_STATUS
	err = rule.Status.AssignPropertiesToNamespacesEventhubsAuthorizationRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToNamespacesEventhubsAuthorizationRule_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *NamespacesEventhubsAuthorizationRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion(),
		Kind:    "NamespacesEventhubsAuthorizationRule",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of NamespacesEventhubsAuthorizationRule. Use v1beta20211101.NamespacesEventhubsAuthorizationRule instead
type NamespacesEventhubsAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhubsAuthorizationRule `json:"items"`
}

// Deprecated version of NamespacesEventhubsAuthorizationRule_STATUS. Use v1beta20211101.NamespacesEventhubsAuthorizationRule_STATUS instead
type NamespacesEventhubsAuthorizationRule_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition                                          `json:"conditions,omitempty"`
	Id         *string                                                         `json:"id,omitempty"`
	Location   *string                                                         `json:"location,omitempty"`
	Name       *string                                                         `json:"name,omitempty"`
	Rights     []NamespacesEventhubsAuthorizationRule_Properties_Rights_STATUS `json:"rights,omitempty"`
	SystemData *SystemData_STATUS                                              `json:"systemData,omitempty"`
	Type       *string                                                         `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NamespacesEventhubsAuthorizationRule_STATUS{}

// ConvertStatusFrom populates our NamespacesEventhubsAuthorizationRule_STATUS from the provided source
func (rule *NamespacesEventhubsAuthorizationRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20211101s.NamespacesEventhubsAuthorizationRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignPropertiesFromNamespacesEventhubsAuthorizationRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &alpha20211101s.NamespacesEventhubsAuthorizationRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignPropertiesFromNamespacesEventhubsAuthorizationRule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our NamespacesEventhubsAuthorizationRule_STATUS
func (rule *NamespacesEventhubsAuthorizationRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20211101s.NamespacesEventhubsAuthorizationRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignPropertiesToNamespacesEventhubsAuthorizationRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20211101s.NamespacesEventhubsAuthorizationRule_STATUS{}
	err := rule.AssignPropertiesToNamespacesEventhubsAuthorizationRule_STATUS(dst)
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

var _ genruntime.FromARMConverter = &NamespacesEventhubsAuthorizationRule_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *NamespacesEventhubsAuthorizationRule_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &NamespacesEventhubsAuthorizationRule_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *NamespacesEventhubsAuthorizationRule_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(NamespacesEventhubsAuthorizationRule_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected NamespacesEventhubsAuthorizationRule_STATUSARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		rule.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		rule.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		rule.Name = &name
	}

	// Set property ‘Rights’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.Rights {
			rule.Rights = append(rule.Rights, item)
		}
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		rule.SystemData = &systemData
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		rule.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromNamespacesEventhubsAuthorizationRule_STATUS populates our NamespacesEventhubsAuthorizationRule_STATUS from the provided source NamespacesEventhubsAuthorizationRule_STATUS
func (rule *NamespacesEventhubsAuthorizationRule_STATUS) AssignPropertiesFromNamespacesEventhubsAuthorizationRule_STATUS(source *alpha20211101s.NamespacesEventhubsAuthorizationRule_STATUS) error {

	// Conditions
	rule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	rule.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	rule.Name = genruntime.ClonePointerToString(source.Name)

	// Rights
	if source.Rights != nil {
		rightList := make([]NamespacesEventhubsAuthorizationRule_Properties_Rights_STATUS, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = NamespacesEventhubsAuthorizationRule_Properties_Rights_STATUS(rightItem)
		}
		rule.Rights = rightList
	} else {
		rule.Rights = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignPropertiesFromSystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemData_STATUS() to populate field SystemData")
		}
		rule.SystemData = &systemDatum
	} else {
		rule.SystemData = nil
	}

	// Type
	rule.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToNamespacesEventhubsAuthorizationRule_STATUS populates the provided destination NamespacesEventhubsAuthorizationRule_STATUS from our NamespacesEventhubsAuthorizationRule_STATUS
func (rule *NamespacesEventhubsAuthorizationRule_STATUS) AssignPropertiesToNamespacesEventhubsAuthorizationRule_STATUS(destination *alpha20211101s.NamespacesEventhubsAuthorizationRule_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(rule.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(rule.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(rule.Name)

	// Rights
	if rule.Rights != nil {
		rightList := make([]string, len(rule.Rights))
		for rightIndex, rightItem := range rule.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = string(rightItem)
		}
		destination.Rights = rightList
	} else {
		destination.Rights = nil
	}

	// SystemData
	if rule.SystemData != nil {
		var systemDatum alpha20211101s.SystemData_STATUS
		err := rule.SystemData.AssignPropertiesToSystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(rule.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type NamespacesEventhubsAuthorizationRule_Spec struct {
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
	Rights     []NamespacesEventhubsAuthorizationRule_Spec_Properties_Rights `json:"rights,omitempty"`
	SystemData *SystemData                                                   `json:"systemData,omitempty"`
	Type       *string                                                       `json:"type,omitempty"`
}

var _ genruntime.ARMTransformer = &NamespacesEventhubsAuthorizationRule_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rule *NamespacesEventhubsAuthorizationRule_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rule == nil {
		return nil, nil
	}
	result := &NamespacesEventhubsAuthorizationRule_SpecARM{}

	// Set property ‘AzureName’:
	result.AzureName = rule.AzureName

	// Set property ‘Id’:
	if rule.Id != nil {
		id := *rule.Id
		result.Id = &id
	}

	// Set property ‘Location’:
	if rule.Location != nil {
		location := *rule.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if rule.Rights != nil {
		result.Properties = &NamespacesEventhubsAuthorizationRule_Spec_PropertiesARM{}
	}
	for _, item := range rule.Rights {
		result.Properties.Rights = append(result.Properties.Rights, item)
	}

	// Set property ‘SystemData’:
	if rule.SystemData != nil {
		systemDataARM, err := (*rule.SystemData).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		systemData := *systemDataARM.(*SystemDataARM)
		result.SystemData = &systemData
	}

	// Set property ‘Type’:
	if rule.Type != nil {
		typeVar := *rule.Type
		result.Type = &typeVar
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *NamespacesEventhubsAuthorizationRule_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &NamespacesEventhubsAuthorizationRule_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *NamespacesEventhubsAuthorizationRule_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(NamespacesEventhubsAuthorizationRule_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected NamespacesEventhubsAuthorizationRule_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	rule.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		rule.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		rule.Location = &location
	}

	// Set property ‘Owner’:
	rule.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Rights’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.Rights {
			rule.Rights = append(rule.Rights, item)
		}
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		rule.SystemData = &systemData
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		rule.Type = &typeVar
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &NamespacesEventhubsAuthorizationRule_Spec{}

// ConvertSpecFrom populates our NamespacesEventhubsAuthorizationRule_Spec from the provided source
func (rule *NamespacesEventhubsAuthorizationRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20211101s.NamespacesEventhubsAuthorizationRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignPropertiesFromNamespacesEventhubsAuthorizationRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &alpha20211101s.NamespacesEventhubsAuthorizationRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignPropertiesFromNamespacesEventhubsAuthorizationRule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesEventhubsAuthorizationRule_Spec
func (rule *NamespacesEventhubsAuthorizationRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20211101s.NamespacesEventhubsAuthorizationRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignPropertiesToNamespacesEventhubsAuthorizationRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20211101s.NamespacesEventhubsAuthorizationRule_Spec{}
	err := rule.AssignPropertiesToNamespacesEventhubsAuthorizationRule_Spec(dst)
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

// AssignPropertiesFromNamespacesEventhubsAuthorizationRule_Spec populates our NamespacesEventhubsAuthorizationRule_Spec from the provided source NamespacesEventhubsAuthorizationRule_Spec
func (rule *NamespacesEventhubsAuthorizationRule_Spec) AssignPropertiesFromNamespacesEventhubsAuthorizationRule_Spec(source *alpha20211101s.NamespacesEventhubsAuthorizationRule_Spec) error {

	// AzureName
	rule.AzureName = source.AzureName

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	rule.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rule.Owner = &owner
	} else {
		rule.Owner = nil
	}

	// Rights
	if source.Rights != nil {
		rightList := make([]NamespacesEventhubsAuthorizationRule_Spec_Properties_Rights, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = NamespacesEventhubsAuthorizationRule_Spec_Properties_Rights(rightItem)
		}
		rule.Rights = rightList
	} else {
		rule.Rights = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData
		err := systemDatum.AssignPropertiesFromSystemData(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemData() to populate field SystemData")
		}
		rule.SystemData = &systemDatum
	} else {
		rule.SystemData = nil
	}

	// Type
	rule.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToNamespacesEventhubsAuthorizationRule_Spec populates the provided destination NamespacesEventhubsAuthorizationRule_Spec from our NamespacesEventhubsAuthorizationRule_Spec
func (rule *NamespacesEventhubsAuthorizationRule_Spec) AssignPropertiesToNamespacesEventhubsAuthorizationRule_Spec(destination *alpha20211101s.NamespacesEventhubsAuthorizationRule_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rule.AzureName

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(rule.Location)

	// OriginalVersion
	destination.OriginalVersion = rule.OriginalVersion()

	// Owner
	if rule.Owner != nil {
		owner := rule.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Rights
	if rule.Rights != nil {
		rightList := make([]string, len(rule.Rights))
		for rightIndex, rightItem := range rule.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = string(rightItem)
		}
		destination.Rights = rightList
	} else {
		destination.Rights = nil
	}

	// SystemData
	if rule.SystemData != nil {
		var systemDatum alpha20211101s.SystemData
		err := rule.SystemData.AssignPropertiesToSystemData(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemData() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(rule.Type)

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
func (rule *NamespacesEventhubsAuthorizationRule_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rule *NamespacesEventhubsAuthorizationRule_Spec) SetAzureName(azureName string) {
	rule.AzureName = azureName
}

// Deprecated version of NamespacesEventhubsAuthorizationRule_Properties_Rights_STATUS. Use
// v1beta20211101.NamespacesEventhubsAuthorizationRule_Properties_Rights_STATUS instead
type NamespacesEventhubsAuthorizationRule_Properties_Rights_STATUS string

const (
	NamespacesEventhubsAuthorizationRule_Properties_Rights_Listen_STATUS = NamespacesEventhubsAuthorizationRule_Properties_Rights_STATUS("Listen")
	NamespacesEventhubsAuthorizationRule_Properties_Rights_Manage_STATUS = NamespacesEventhubsAuthorizationRule_Properties_Rights_STATUS("Manage")
	NamespacesEventhubsAuthorizationRule_Properties_Rights_Send_STATUS   = NamespacesEventhubsAuthorizationRule_Properties_Rights_STATUS("Send")
)

// Deprecated version of NamespacesEventhubsAuthorizationRule_Spec_Properties_Rights. Use
// v1beta20211101.NamespacesEventhubsAuthorizationRule_Spec_Properties_Rights instead
// +kubebuilder:validation:Enum={"Listen","Manage","Send"}
type NamespacesEventhubsAuthorizationRule_Spec_Properties_Rights string

const (
	NamespacesEventhubsAuthorizationRule_Spec_Properties_Rights_Listen = NamespacesEventhubsAuthorizationRule_Spec_Properties_Rights("Listen")
	NamespacesEventhubsAuthorizationRule_Spec_Properties_Rights_Manage = NamespacesEventhubsAuthorizationRule_Spec_Properties_Rights("Manage")
	NamespacesEventhubsAuthorizationRule_Spec_Properties_Rights_Send   = NamespacesEventhubsAuthorizationRule_Spec_Properties_Rights("Send")
)

func init() {
	SchemeBuilder.Register(&NamespacesEventhubsAuthorizationRule{}, &NamespacesEventhubsAuthorizationRuleList{})
}
