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
// Deprecated version of NamespacesAuthorizationRule. Use v1beta20211101.NamespacesAuthorizationRule instead
type NamespacesAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesAuthorizationRule_Spec   `json:"spec,omitempty"`
	Status            NamespacesAuthorizationRule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesAuthorizationRule{}

// GetConditions returns the conditions of the resource
func (rule *NamespacesAuthorizationRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *NamespacesAuthorizationRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesAuthorizationRule{}

// ConvertFrom populates our NamespacesAuthorizationRule from the provided hub NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20211101s.NamespacesAuthorizationRule

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = rule.AssignPropertiesFromNamespacesAuthorizationRule(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to rule")
	}

	return nil
}

// ConvertTo populates the provided hub NamespacesAuthorizationRule from our NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20211101s.NamespacesAuthorizationRule
	err := rule.AssignPropertiesToNamespacesAuthorizationRule(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from rule")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-eventhub-azure-com-v1alpha1api20211101-namespacesauthorizationrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespacesauthorizationrules,verbs=create;update,versions=v1alpha1api20211101,name=default.v1alpha1api20211101.namespacesauthorizationrules.eventhub.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &NamespacesAuthorizationRule{}

// Default applies defaults to the NamespacesAuthorizationRule resource
func (rule *NamespacesAuthorizationRule) Default() {
	rule.defaultImpl()
	var temp interface{} = rule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (rule *NamespacesAuthorizationRule) defaultAzureName() {
	if rule.Spec.AzureName == "" {
		rule.Spec.AzureName = rule.Name
	}
}

// defaultImpl applies the code generated defaults to the NamespacesAuthorizationRule resource
func (rule *NamespacesAuthorizationRule) defaultImpl() { rule.defaultAzureName() }

var _ genruntime.KubernetesResource = &NamespacesAuthorizationRule{}

// AzureName returns the Azure name of the resource
func (rule *NamespacesAuthorizationRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "20211101"
func (rule NamespacesAuthorizationRule) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (rule *NamespacesAuthorizationRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (rule *NamespacesAuthorizationRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *NamespacesAuthorizationRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (rule *NamespacesAuthorizationRule) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *NamespacesAuthorizationRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &NamespacesAuthorizationRule_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rule *NamespacesAuthorizationRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *NamespacesAuthorizationRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*NamespacesAuthorizationRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st NamespacesAuthorizationRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-eventhub-azure-com-v1alpha1api20211101-namespacesauthorizationrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespacesauthorizationrules,verbs=create;update,versions=v1alpha1api20211101,name=validate.v1alpha1api20211101.namespacesauthorizationrules.eventhub.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &NamespacesAuthorizationRule{}

// ValidateCreate validates the creation of the resource
func (rule *NamespacesAuthorizationRule) ValidateCreate() error {
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
func (rule *NamespacesAuthorizationRule) ValidateDelete() error {
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
func (rule *NamespacesAuthorizationRule) ValidateUpdate(old runtime.Object) error {
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
func (rule *NamespacesAuthorizationRule) createValidations() []func() error {
	return []func() error{rule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (rule *NamespacesAuthorizationRule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (rule *NamespacesAuthorizationRule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return rule.validateResourceReferences()
		},
		rule.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (rule *NamespacesAuthorizationRule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *NamespacesAuthorizationRule) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*NamespacesAuthorizationRule)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, rule)
}

// AssignPropertiesFromNamespacesAuthorizationRule populates our NamespacesAuthorizationRule from the provided source NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) AssignPropertiesFromNamespacesAuthorizationRule(source *alpha20211101s.NamespacesAuthorizationRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NamespacesAuthorizationRule_Spec
	err := spec.AssignPropertiesFromNamespacesAuthorizationRule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromNamespacesAuthorizationRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status NamespacesAuthorizationRule_STATUS
	err = status.AssignPropertiesFromNamespacesAuthorizationRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromNamespacesAuthorizationRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignPropertiesToNamespacesAuthorizationRule populates the provided destination NamespacesAuthorizationRule from our NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) AssignPropertiesToNamespacesAuthorizationRule(destination *alpha20211101s.NamespacesAuthorizationRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20211101s.NamespacesAuthorizationRule_Spec
	err := rule.Spec.AssignPropertiesToNamespacesAuthorizationRule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToNamespacesAuthorizationRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20211101s.NamespacesAuthorizationRule_STATUS
	err = rule.Status.AssignPropertiesToNamespacesAuthorizationRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToNamespacesAuthorizationRule_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *NamespacesAuthorizationRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion(),
		Kind:    "NamespacesAuthorizationRule",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of NamespacesAuthorizationRule. Use v1beta20211101.NamespacesAuthorizationRule instead
type NamespacesAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesAuthorizationRule `json:"items"`
}

// Deprecated version of NamespacesAuthorizationRule_STATUS. Use v1beta20211101.NamespacesAuthorizationRule_STATUS instead
type NamespacesAuthorizationRule_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition                                 `json:"conditions,omitempty"`
	Id         *string                                                `json:"id,omitempty"`
	Location   *string                                                `json:"location,omitempty"`
	Name       *string                                                `json:"name,omitempty"`
	Rights     []NamespacesAuthorizationRule_Properties_Rights_STATUS `json:"rights,omitempty"`
	SystemData *SystemData_STATUS                                     `json:"systemData,omitempty"`
	Type       *string                                                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NamespacesAuthorizationRule_STATUS{}

// ConvertStatusFrom populates our NamespacesAuthorizationRule_STATUS from the provided source
func (rule *NamespacesAuthorizationRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20211101s.NamespacesAuthorizationRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignPropertiesFromNamespacesAuthorizationRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &alpha20211101s.NamespacesAuthorizationRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignPropertiesFromNamespacesAuthorizationRule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our NamespacesAuthorizationRule_STATUS
func (rule *NamespacesAuthorizationRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20211101s.NamespacesAuthorizationRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignPropertiesToNamespacesAuthorizationRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20211101s.NamespacesAuthorizationRule_STATUS{}
	err := rule.AssignPropertiesToNamespacesAuthorizationRule_STATUS(dst)
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

var _ genruntime.FromARMConverter = &NamespacesAuthorizationRule_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *NamespacesAuthorizationRule_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &NamespacesAuthorizationRule_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *NamespacesAuthorizationRule_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(NamespacesAuthorizationRule_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected NamespacesAuthorizationRule_STATUSARM, got %T", armInput)
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

// AssignPropertiesFromNamespacesAuthorizationRule_STATUS populates our NamespacesAuthorizationRule_STATUS from the provided source NamespacesAuthorizationRule_STATUS
func (rule *NamespacesAuthorizationRule_STATUS) AssignPropertiesFromNamespacesAuthorizationRule_STATUS(source *alpha20211101s.NamespacesAuthorizationRule_STATUS) error {

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
		rightList := make([]NamespacesAuthorizationRule_Properties_Rights_STATUS, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = NamespacesAuthorizationRule_Properties_Rights_STATUS(rightItem)
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

// AssignPropertiesToNamespacesAuthorizationRule_STATUS populates the provided destination NamespacesAuthorizationRule_STATUS from our NamespacesAuthorizationRule_STATUS
func (rule *NamespacesAuthorizationRule_STATUS) AssignPropertiesToNamespacesAuthorizationRule_STATUS(destination *alpha20211101s.NamespacesAuthorizationRule_STATUS) error {
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

type NamespacesAuthorizationRule_Spec struct {
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
	Rights     []NamespacesAuthorizationRule_Spec_Properties_Rights `json:"rights,omitempty"`
	SystemData *SystemData                                          `json:"systemData,omitempty"`
	Type       *string                                              `json:"type,omitempty"`
}

var _ genruntime.ARMTransformer = &NamespacesAuthorizationRule_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rule *NamespacesAuthorizationRule_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rule == nil {
		return nil, nil
	}
	result := &NamespacesAuthorizationRule_SpecARM{}

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
		result.Properties = &NamespacesAuthorizationRule_Spec_PropertiesARM{}
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
func (rule *NamespacesAuthorizationRule_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &NamespacesAuthorizationRule_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *NamespacesAuthorizationRule_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(NamespacesAuthorizationRule_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected NamespacesAuthorizationRule_SpecARM, got %T", armInput)
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

var _ genruntime.ConvertibleSpec = &NamespacesAuthorizationRule_Spec{}

// ConvertSpecFrom populates our NamespacesAuthorizationRule_Spec from the provided source
func (rule *NamespacesAuthorizationRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20211101s.NamespacesAuthorizationRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignPropertiesFromNamespacesAuthorizationRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &alpha20211101s.NamespacesAuthorizationRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignPropertiesFromNamespacesAuthorizationRule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesAuthorizationRule_Spec
func (rule *NamespacesAuthorizationRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20211101s.NamespacesAuthorizationRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignPropertiesToNamespacesAuthorizationRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20211101s.NamespacesAuthorizationRule_Spec{}
	err := rule.AssignPropertiesToNamespacesAuthorizationRule_Spec(dst)
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

// AssignPropertiesFromNamespacesAuthorizationRule_Spec populates our NamespacesAuthorizationRule_Spec from the provided source NamespacesAuthorizationRule_Spec
func (rule *NamespacesAuthorizationRule_Spec) AssignPropertiesFromNamespacesAuthorizationRule_Spec(source *alpha20211101s.NamespacesAuthorizationRule_Spec) error {

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
		rightList := make([]NamespacesAuthorizationRule_Spec_Properties_Rights, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = NamespacesAuthorizationRule_Spec_Properties_Rights(rightItem)
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

// AssignPropertiesToNamespacesAuthorizationRule_Spec populates the provided destination NamespacesAuthorizationRule_Spec from our NamespacesAuthorizationRule_Spec
func (rule *NamespacesAuthorizationRule_Spec) AssignPropertiesToNamespacesAuthorizationRule_Spec(destination *alpha20211101s.NamespacesAuthorizationRule_Spec) error {
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
func (rule *NamespacesAuthorizationRule_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rule *NamespacesAuthorizationRule_Spec) SetAzureName(azureName string) {
	rule.AzureName = azureName
}

// Deprecated version of NamespacesAuthorizationRule_Properties_Rights_STATUS. Use
// v1beta20211101.NamespacesAuthorizationRule_Properties_Rights_STATUS instead
type NamespacesAuthorizationRule_Properties_Rights_STATUS string

const (
	NamespacesAuthorizationRule_Properties_Rights_STATUSListen = NamespacesAuthorizationRule_Properties_Rights_STATUS("Listen")
	NamespacesAuthorizationRule_Properties_Rights_STATUSManage = NamespacesAuthorizationRule_Properties_Rights_STATUS("Manage")
	NamespacesAuthorizationRule_Properties_Rights_STATUSSend   = NamespacesAuthorizationRule_Properties_Rights_STATUS("Send")
)

// Deprecated version of NamespacesAuthorizationRule_Spec_Properties_Rights. Use
// v1beta20211101.NamespacesAuthorizationRule_Spec_Properties_Rights instead
// +kubebuilder:validation:Enum={"Listen","Manage","Send"}
type NamespacesAuthorizationRule_Spec_Properties_Rights string

const (
	NamespacesAuthorizationRule_Spec_Properties_RightsListen = NamespacesAuthorizationRule_Spec_Properties_Rights("Listen")
	NamespacesAuthorizationRule_Spec_Properties_RightsManage = NamespacesAuthorizationRule_Spec_Properties_Rights("Manage")
	NamespacesAuthorizationRule_Spec_Properties_RightsSend   = NamespacesAuthorizationRule_Spec_Properties_Rights("Send")
)

func init() {
	SchemeBuilder.Register(&NamespacesAuthorizationRule{}, &NamespacesAuthorizationRuleList{})
}
