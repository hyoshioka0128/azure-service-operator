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
//Deprecated version of NamespacesEventhubsAuthorizationRule. Use v1beta20211101.NamespacesEventhubsAuthorizationRule instead
type NamespacesEventhubsAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesEventhubsAuthorizationRules_Spec `json:"spec,omitempty"`
	Status            AuthorizationRule_Status                   `json:"status,omitempty"`
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
	return "2021-11-01"
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

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
func (rule *NamespacesEventhubsAuthorizationRule) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *NamespacesEventhubsAuthorizationRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &AuthorizationRule_Status{}
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
	if st, ok := status.(*AuthorizationRule_Status); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st AuthorizationRule_Status
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
		rule.validateImmutableProperties}
}

// validateImmutableProperties validates all immutable properties
func (rule *NamespacesEventhubsAuthorizationRule) validateImmutableProperties(old runtime.Object) error {

	resourceID := genruntime.GetResourceIDOrDefault(rule)
	if resourceID == "" {
		return nil
	}

	oldObj, ok := old.(*NamespacesEventhubsAuthorizationRule)
	if !ok {
		return nil
	}

	if oldObj.AzureName() != rule.AzureName() {
		return errors.New("update for 'AzureName()' is not allowed")
	}

	if oldObj.Owner().Name != rule.Owner().Name {
		return errors.New("update for 'Owner().Name' is not allowed")
	}

	// No error
	return nil
}

// validateResourceReferences validates all resource references
func (rule *NamespacesEventhubsAuthorizationRule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromNamespacesEventhubsAuthorizationRule populates our NamespacesEventhubsAuthorizationRule from the provided source NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) AssignPropertiesFromNamespacesEventhubsAuthorizationRule(source *alpha20211101s.NamespacesEventhubsAuthorizationRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NamespacesEventhubsAuthorizationRules_Spec
	err := spec.AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status AuthorizationRule_Status
	err = status.AssignPropertiesFromAuthorizationRuleStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromAuthorizationRuleStatus() to populate field Status")
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
	var spec alpha20211101s.NamespacesEventhubsAuthorizationRules_Spec
	err := rule.Spec.AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20211101s.AuthorizationRule_Status
	err = rule.Status.AssignPropertiesToAuthorizationRuleStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToAuthorizationRuleStatus() to populate field Status")
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
//Deprecated version of NamespacesEventhubsAuthorizationRule. Use v1beta20211101.NamespacesEventhubsAuthorizationRule instead
type NamespacesEventhubsAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhubsAuthorizationRule `json:"items"`
}

type NamespacesEventhubsAuthorizationRules_Spec struct {
	// +kubebuilder:validation:MinLength=1
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName string  `json:"azureName,omitempty"`
	Location  *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a eventhub.azure.com/NamespacesEventhub resource
	Owner *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"NamespacesEventhub"`

	// +kubebuilder:validation:Required
	Rights []AuthorizationRulePropertiesRights `json:"rights,omitempty"`
	Tags   map[string]string                   `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &NamespacesEventhubsAuthorizationRules_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rules *NamespacesEventhubsAuthorizationRules_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rules == nil {
		return nil, nil
	}
	var result NamespacesEventhubsAuthorizationRules_SpecARM

	// Set property ‘Location’:
	if rules.Location != nil {
		location := *rules.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if rules.Rights != nil {
		result.Properties = &AuthorizationRulePropertiesARM{}
	}
	for _, item := range rules.Rights {
		result.Properties.Rights = append(result.Properties.Rights, item)
	}

	// Set property ‘Tags’:
	if rules.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range rules.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rules *NamespacesEventhubsAuthorizationRules_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &NamespacesEventhubsAuthorizationRules_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rules *NamespacesEventhubsAuthorizationRules_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(NamespacesEventhubsAuthorizationRules_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected NamespacesEventhubsAuthorizationRules_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	rules.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		rules.Location = &location
	}

	// Set property ‘Owner’:
	rules.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Rights’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.Rights {
			rules.Rights = append(rules.Rights, item)
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		rules.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			rules.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &NamespacesEventhubsAuthorizationRules_Spec{}

// ConvertSpecFrom populates our NamespacesEventhubsAuthorizationRules_Spec from the provided source
func (rules *NamespacesEventhubsAuthorizationRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20211101s.NamespacesEventhubsAuthorizationRules_Spec)
	if ok {
		// Populate our instance from source
		return rules.AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec(src)
	}

	// Convert to an intermediate form
	src = &alpha20211101s.NamespacesEventhubsAuthorizationRules_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rules.AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesEventhubsAuthorizationRules_Spec
func (rules *NamespacesEventhubsAuthorizationRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20211101s.NamespacesEventhubsAuthorizationRules_Spec)
	if ok {
		// Populate destination from our instance
		return rules.AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20211101s.NamespacesEventhubsAuthorizationRules_Spec{}
	err := rules.AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec(dst)
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

// AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec populates our NamespacesEventhubsAuthorizationRules_Spec from the provided source NamespacesEventhubsAuthorizationRules_Spec
func (rules *NamespacesEventhubsAuthorizationRules_Spec) AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec(source *alpha20211101s.NamespacesEventhubsAuthorizationRules_Spec) error {

	// AzureName
	rules.AzureName = source.AzureName

	// Location
	rules.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rules.Owner = &owner
	} else {
		rules.Owner = nil
	}

	// Rights
	if source.Rights != nil {
		rightList := make([]AuthorizationRulePropertiesRights, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = AuthorizationRulePropertiesRights(rightItem)
		}
		rules.Rights = rightList
	} else {
		rules.Rights = nil
	}

	// Tags
	rules.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec populates the provided destination NamespacesEventhubsAuthorizationRules_Spec from our NamespacesEventhubsAuthorizationRules_Spec
func (rules *NamespacesEventhubsAuthorizationRules_Spec) AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec(destination *alpha20211101s.NamespacesEventhubsAuthorizationRules_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rules.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(rules.Location)

	// OriginalVersion
	destination.OriginalVersion = rules.OriginalVersion()

	// Owner
	if rules.Owner != nil {
		owner := rules.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Rights
	if rules.Rights != nil {
		rightList := make([]string, len(rules.Rights))
		for rightIndex, rightItem := range rules.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = string(rightItem)
		}
		destination.Rights = rightList
	} else {
		destination.Rights = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(rules.Tags)

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
func (rules *NamespacesEventhubsAuthorizationRules_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rules *NamespacesEventhubsAuthorizationRules_Spec) SetAzureName(azureName string) {
	rules.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&NamespacesEventhubsAuthorizationRule{}, &NamespacesEventhubsAuthorizationRuleList{})
}
