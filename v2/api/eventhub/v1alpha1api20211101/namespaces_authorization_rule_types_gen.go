// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/eventhub/v1alpha1api20211101storage"
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
type NamespacesAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesAuthorizationRules_SPEC `json:"spec,omitempty"`
	Status            AuthorizationRule_Status          `json:"status,omitempty"`
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
	source, ok := hub.(*v1alpha1api20211101storage.NamespacesAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected storage:eventhub/v1alpha1api20211101storage/NamespacesAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignPropertiesFromNamespacesAuthorizationRule(source)
}

// ConvertTo populates the provided hub NamespacesAuthorizationRule from our NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20211101storage.NamespacesAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected storage:eventhub/v1alpha1api20211101storage/NamespacesAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignPropertiesToNamespacesAuthorizationRule(destination)
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

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule NamespacesAuthorizationRule) GetAPIVersion() string {
	return "2021-11-01"
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
	return &AuthorizationRule_Status{}
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
	}
}

// validateResourceReferences validates all resource references
func (rule *NamespacesAuthorizationRule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromNamespacesAuthorizationRule populates our NamespacesAuthorizationRule from the provided source NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) AssignPropertiesFromNamespacesAuthorizationRule(source *v1alpha1api20211101storage.NamespacesAuthorizationRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NamespacesAuthorizationRules_SPEC
	err := spec.AssignPropertiesFromNamespacesAuthorizationRulesSPEC(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromNamespacesAuthorizationRulesSPEC() to populate field Spec")
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

// AssignPropertiesToNamespacesAuthorizationRule populates the provided destination NamespacesAuthorizationRule from our NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) AssignPropertiesToNamespacesAuthorizationRule(destination *v1alpha1api20211101storage.NamespacesAuthorizationRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20211101storage.NamespacesAuthorizationRules_SPEC
	err := rule.Spec.AssignPropertiesToNamespacesAuthorizationRulesSPEC(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToNamespacesAuthorizationRulesSPEC() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20211101storage.AuthorizationRule_Status
	err = rule.Status.AssignPropertiesToAuthorizationRuleStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToAuthorizationRuleStatus() to populate field Status")
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
type NamespacesAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesAuthorizationRule `json:"items"`
}

type AuthorizationRule_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//Rights: The rights associated with the rule.
	Rights []AuthorizationRuleStatusPropertiesRights `json:"rights,omitempty"`

	//SystemData: The system meta data relating to this resource.
	SystemData *SystemData_Status `json:"systemData,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or
	//"Microsoft.EventHub/Namespaces/EventHubs"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &AuthorizationRule_Status{}

// ConvertStatusFrom populates our AuthorizationRule_Status from the provided source
func (rule *AuthorizationRule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20211101storage.AuthorizationRule_Status)
	if ok {
		// Populate our instance from source
		return rule.AssignPropertiesFromAuthorizationRuleStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20211101storage.AuthorizationRule_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignPropertiesFromAuthorizationRuleStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our AuthorizationRule_Status
func (rule *AuthorizationRule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20211101storage.AuthorizationRule_Status)
	if ok {
		// Populate destination from our instance
		return rule.AssignPropertiesToAuthorizationRuleStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20211101storage.AuthorizationRule_Status{}
	err := rule.AssignPropertiesToAuthorizationRuleStatus(dst)
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

var _ genruntime.FromARMConverter = &AuthorizationRule_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *AuthorizationRule_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &AuthorizationRule_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *AuthorizationRule_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(AuthorizationRule_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected AuthorizationRule_StatusARM, got %T", armInput)
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
		var systemData1 SystemData_Status
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

// AssignPropertiesFromAuthorizationRuleStatus populates our AuthorizationRule_Status from the provided source AuthorizationRule_Status
func (rule *AuthorizationRule_Status) AssignPropertiesFromAuthorizationRuleStatus(source *v1alpha1api20211101storage.AuthorizationRule_Status) error {

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
		rightList := make([]AuthorizationRuleStatusPropertiesRights, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = AuthorizationRuleStatusPropertiesRights(rightItem)
		}
		rule.Rights = rightList
	} else {
		rule.Rights = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemDataStatus() to populate field SystemData")
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

// AssignPropertiesToAuthorizationRuleStatus populates the provided destination AuthorizationRule_Status from our AuthorizationRule_Status
func (rule *AuthorizationRule_Status) AssignPropertiesToAuthorizationRuleStatus(destination *v1alpha1api20211101storage.AuthorizationRule_Status) error {
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
		var systemDatum v1alpha1api20211101storage.SystemData_Status
		err := rule.SystemData.AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemDataStatus() to populate field SystemData")
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

type NamespacesAuthorizationRules_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`

	//Rights: The rights associated with the rule.
	Rights []NamespacesAuthorizationRulesSPECPropertiesRights `json:"rights,omitempty"`
}

var _ genruntime.ARMTransformer = &NamespacesAuthorizationRules_SPEC{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (spec *NamespacesAuthorizationRules_SPEC) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if spec == nil {
		return nil, nil
	}
	var result NamespacesAuthorizationRules_SPECARM

	// Set property ‘AzureName’:
	result.AzureName = spec.AzureName

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if spec.Rights != nil {
		result.Properties = &NamespacesAuthorizationRules_SPEC_PropertiesARM{}
	}
	for _, item := range spec.Rights {
		result.Properties.Rights = append(result.Properties.Rights, item)
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (spec *NamespacesAuthorizationRules_SPEC) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &NamespacesAuthorizationRules_SPECARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (spec *NamespacesAuthorizationRules_SPEC) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(NamespacesAuthorizationRules_SPECARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected NamespacesAuthorizationRules_SPECARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	spec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Owner’:
	spec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Rights’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.Rights {
			spec.Rights = append(spec.Rights, item)
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &NamespacesAuthorizationRules_SPEC{}

// ConvertSpecFrom populates our NamespacesAuthorizationRules_SPEC from the provided source
func (spec *NamespacesAuthorizationRules_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20211101storage.NamespacesAuthorizationRules_SPEC)
	if ok {
		// Populate our instance from source
		return spec.AssignPropertiesFromNamespacesAuthorizationRulesSPEC(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20211101storage.NamespacesAuthorizationRules_SPEC{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = spec.AssignPropertiesFromNamespacesAuthorizationRulesSPEC(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesAuthorizationRules_SPEC
func (spec *NamespacesAuthorizationRules_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20211101storage.NamespacesAuthorizationRules_SPEC)
	if ok {
		// Populate destination from our instance
		return spec.AssignPropertiesToNamespacesAuthorizationRulesSPEC(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20211101storage.NamespacesAuthorizationRules_SPEC{}
	err := spec.AssignPropertiesToNamespacesAuthorizationRulesSPEC(dst)
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

// AssignPropertiesFromNamespacesAuthorizationRulesSPEC populates our NamespacesAuthorizationRules_SPEC from the provided source NamespacesAuthorizationRules_SPEC
func (spec *NamespacesAuthorizationRules_SPEC) AssignPropertiesFromNamespacesAuthorizationRulesSPEC(source *v1alpha1api20211101storage.NamespacesAuthorizationRules_SPEC) error {

	// AzureName
	spec.AzureName = source.AzureName

	// Owner
	spec.Owner = source.Owner.Copy()

	// Rights
	if source.Rights != nil {
		rightList := make([]NamespacesAuthorizationRulesSPECPropertiesRights, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = NamespacesAuthorizationRulesSPECPropertiesRights(rightItem)
		}
		spec.Rights = rightList
	} else {
		spec.Rights = nil
	}

	// No error
	return nil
}

// AssignPropertiesToNamespacesAuthorizationRulesSPEC populates the provided destination NamespacesAuthorizationRules_SPEC from our NamespacesAuthorizationRules_SPEC
func (spec *NamespacesAuthorizationRules_SPEC) AssignPropertiesToNamespacesAuthorizationRulesSPEC(destination *v1alpha1api20211101storage.NamespacesAuthorizationRules_SPEC) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = spec.AzureName

	// OriginalVersion
	destination.OriginalVersion = spec.OriginalVersion()

	// Owner
	destination.Owner = spec.Owner.Copy()

	// Rights
	if spec.Rights != nil {
		rightList := make([]string, len(spec.Rights))
		for rightIndex, rightItem := range spec.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = string(rightItem)
		}
		destination.Rights = rightList
	} else {
		destination.Rights = nil
	}

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
func (spec *NamespacesAuthorizationRules_SPEC) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (spec *NamespacesAuthorizationRules_SPEC) SetAzureName(azureName string) {
	spec.AzureName = azureName
}

// +kubebuilder:validation:Enum={"Listen","Manage","Send"}
type NamespacesAuthorizationRulesSPECPropertiesRights string

const (
	NamespacesAuthorizationRulesSPECPropertiesRightsListen = NamespacesAuthorizationRulesSPECPropertiesRights("Listen")
	NamespacesAuthorizationRulesSPECPropertiesRightsManage = NamespacesAuthorizationRulesSPECPropertiesRights("Manage")
	NamespacesAuthorizationRulesSPECPropertiesRightsSend   = NamespacesAuthorizationRulesSPECPropertiesRights("Send")
)

func init() {
	SchemeBuilder.Register(&NamespacesAuthorizationRule{}, &NamespacesAuthorizationRuleList{})
}
