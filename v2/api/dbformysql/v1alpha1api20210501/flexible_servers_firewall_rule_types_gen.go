// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

import (
	"fmt"
	alpha20210501s "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1alpha1api20210501storage"
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
// Deprecated version of FlexibleServersFirewallRule. Use v1beta20210501.FlexibleServersFirewallRule instead
type FlexibleServersFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersFirewallRule_Spec   `json:"spec,omitempty"`
	Status            FlexibleServersFirewallRule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersFirewallRule{}

// GetConditions returns the conditions of the resource
func (rule *FlexibleServersFirewallRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *FlexibleServersFirewallRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ conversion.Convertible = &FlexibleServersFirewallRule{}

// ConvertFrom populates our FlexibleServersFirewallRule from the provided hub FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20210501s.FlexibleServersFirewallRule

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = rule.AssignPropertiesFromFlexibleServersFirewallRule(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to rule")
	}

	return nil
}

// ConvertTo populates the provided hub FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20210501s.FlexibleServersFirewallRule
	err := rule.AssignPropertiesToFlexibleServersFirewallRule(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from rule")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-dbformysql-azure-com-v1alpha1api20210501-flexibleserversfirewallrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformysql.azure.com,resources=flexibleserversfirewallrules,verbs=create;update,versions=v1alpha1api20210501,name=default.v1alpha1api20210501.flexibleserversfirewallrules.dbformysql.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &FlexibleServersFirewallRule{}

// Default applies defaults to the FlexibleServersFirewallRule resource
func (rule *FlexibleServersFirewallRule) Default() {
	rule.defaultImpl()
	var temp interface{} = rule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (rule *FlexibleServersFirewallRule) defaultAzureName() {
	if rule.Spec.AzureName == "" {
		rule.Spec.AzureName = rule.Name
	}
}

// defaultImpl applies the code generated defaults to the FlexibleServersFirewallRule resource
func (rule *FlexibleServersFirewallRule) defaultImpl() { rule.defaultAzureName() }

var _ genruntime.KubernetesResource = &FlexibleServersFirewallRule{}

// AzureName returns the Azure name of the resource
func (rule *FlexibleServersFirewallRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (rule FlexibleServersFirewallRule) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (rule *FlexibleServersFirewallRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (rule *FlexibleServersFirewallRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *FlexibleServersFirewallRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (rule *FlexibleServersFirewallRule) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *FlexibleServersFirewallRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServersFirewallRule_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rule *FlexibleServersFirewallRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *FlexibleServersFirewallRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServersFirewallRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServersFirewallRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-dbformysql-azure-com-v1alpha1api20210501-flexibleserversfirewallrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformysql.azure.com,resources=flexibleserversfirewallrules,verbs=create;update,versions=v1alpha1api20210501,name=validate.v1alpha1api20210501.flexibleserversfirewallrules.dbformysql.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &FlexibleServersFirewallRule{}

// ValidateCreate validates the creation of the resource
func (rule *FlexibleServersFirewallRule) ValidateCreate() error {
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
func (rule *FlexibleServersFirewallRule) ValidateDelete() error {
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
func (rule *FlexibleServersFirewallRule) ValidateUpdate(old runtime.Object) error {
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
func (rule *FlexibleServersFirewallRule) createValidations() []func() error {
	return []func() error{rule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (rule *FlexibleServersFirewallRule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (rule *FlexibleServersFirewallRule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return rule.validateResourceReferences()
		},
		rule.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (rule *FlexibleServersFirewallRule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *FlexibleServersFirewallRule) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*FlexibleServersFirewallRule)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, rule)
}

// AssignPropertiesFromFlexibleServersFirewallRule populates our FlexibleServersFirewallRule from the provided source FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) AssignPropertiesFromFlexibleServersFirewallRule(source *alpha20210501s.FlexibleServersFirewallRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersFirewallRule_Spec
	err := spec.AssignPropertiesFromFlexibleServersFirewallRule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromFlexibleServersFirewallRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status FlexibleServersFirewallRule_STATUS
	err = status.AssignPropertiesFromFlexibleServersFirewallRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromFlexibleServersFirewallRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersFirewallRule populates the provided destination FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) AssignPropertiesToFlexibleServersFirewallRule(destination *alpha20210501s.FlexibleServersFirewallRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20210501s.FlexibleServersFirewallRule_Spec
	err := rule.Spec.AssignPropertiesToFlexibleServersFirewallRule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToFlexibleServersFirewallRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20210501s.FlexibleServersFirewallRule_STATUS
	err = rule.Status.AssignPropertiesToFlexibleServersFirewallRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToFlexibleServersFirewallRule_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *FlexibleServersFirewallRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion(),
		Kind:    "FlexibleServersFirewallRule",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of FlexibleServersFirewallRule. Use v1beta20210501.FlexibleServersFirewallRule instead
type FlexibleServersFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersFirewallRule `json:"items"`
}

// Deprecated version of FlexibleServersFirewallRule_STATUS. Use v1beta20210501.FlexibleServersFirewallRule_STATUS instead
type FlexibleServersFirewallRule_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions     []conditions.Condition `json:"conditions,omitempty"`
	EndIpAddress   *string                `json:"endIpAddress,omitempty"`
	Id             *string                `json:"id,omitempty"`
	Name           *string                `json:"name,omitempty"`
	StartIpAddress *string                `json:"startIpAddress,omitempty"`
	SystemData     *SystemData_STATUS     `json:"systemData,omitempty"`
	Type           *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServersFirewallRule_STATUS{}

// ConvertStatusFrom populates our FlexibleServersFirewallRule_STATUS from the provided source
func (rule *FlexibleServersFirewallRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20210501s.FlexibleServersFirewallRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignPropertiesFromFlexibleServersFirewallRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &alpha20210501s.FlexibleServersFirewallRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignPropertiesFromFlexibleServersFirewallRule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FlexibleServersFirewallRule_STATUS
func (rule *FlexibleServersFirewallRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20210501s.FlexibleServersFirewallRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignPropertiesToFlexibleServersFirewallRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210501s.FlexibleServersFirewallRule_STATUS{}
	err := rule.AssignPropertiesToFlexibleServersFirewallRule_STATUS(dst)
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

var _ genruntime.FromARMConverter = &FlexibleServersFirewallRule_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *FlexibleServersFirewallRule_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FlexibleServersFirewallRule_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *FlexibleServersFirewallRule_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FlexibleServersFirewallRule_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FlexibleServersFirewallRule_STATUSARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘EndIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIpAddress != nil {
			endIpAddress := *typedInput.Properties.EndIpAddress
			rule.EndIpAddress = &endIpAddress
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		rule.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		rule.Name = &name
	}

	// Set property ‘StartIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIpAddress != nil {
			startIpAddress := *typedInput.Properties.StartIpAddress
			rule.StartIpAddress = &startIpAddress
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

// AssignPropertiesFromFlexibleServersFirewallRule_STATUS populates our FlexibleServersFirewallRule_STATUS from the provided source FlexibleServersFirewallRule_STATUS
func (rule *FlexibleServersFirewallRule_STATUS) AssignPropertiesFromFlexibleServersFirewallRule_STATUS(source *alpha20210501s.FlexibleServersFirewallRule_STATUS) error {

	// Conditions
	rule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// EndIpAddress
	rule.EndIpAddress = genruntime.ClonePointerToString(source.EndIpAddress)

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	rule.Name = genruntime.ClonePointerToString(source.Name)

	// StartIpAddress
	rule.StartIpAddress = genruntime.ClonePointerToString(source.StartIpAddress)

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

// AssignPropertiesToFlexibleServersFirewallRule_STATUS populates the provided destination FlexibleServersFirewallRule_STATUS from our FlexibleServersFirewallRule_STATUS
func (rule *FlexibleServersFirewallRule_STATUS) AssignPropertiesToFlexibleServersFirewallRule_STATUS(destination *alpha20210501s.FlexibleServersFirewallRule_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(rule.Conditions)

	// EndIpAddress
	destination.EndIpAddress = genruntime.ClonePointerToString(rule.EndIpAddress)

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(rule.Name)

	// StartIpAddress
	destination.StartIpAddress = genruntime.ClonePointerToString(rule.StartIpAddress)

	// SystemData
	if rule.SystemData != nil {
		var systemDatum alpha20210501s.SystemData_STATUS
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

type FlexibleServersFirewallRule_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	EndIpAddress *string `json:"endIpAddress,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	StartIpAddress *string `json:"startIpAddress,omitempty"`
}

var _ genruntime.ARMTransformer = &FlexibleServersFirewallRule_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rule *FlexibleServersFirewallRule_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rule == nil {
		return nil, nil
	}
	result := &FlexibleServersFirewallRule_SpecARM{}

	// Set property ‘AzureName’:
	result.AzureName = rule.AzureName

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if rule.EndIpAddress != nil || rule.StartIpAddress != nil {
		result.Properties = &FirewallRulePropertiesARM{}
	}
	if rule.EndIpAddress != nil {
		endIpAddress := *rule.EndIpAddress
		result.Properties.EndIpAddress = &endIpAddress
	}
	if rule.StartIpAddress != nil {
		startIpAddress := *rule.StartIpAddress
		result.Properties.StartIpAddress = &startIpAddress
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *FlexibleServersFirewallRule_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FlexibleServersFirewallRule_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *FlexibleServersFirewallRule_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FlexibleServersFirewallRule_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FlexibleServersFirewallRule_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	rule.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘EndIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIpAddress != nil {
			endIpAddress := *typedInput.Properties.EndIpAddress
			rule.EndIpAddress = &endIpAddress
		}
	}

	// Set property ‘Owner’:
	rule.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘StartIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIpAddress != nil {
			startIpAddress := *typedInput.Properties.StartIpAddress
			rule.StartIpAddress = &startIpAddress
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &FlexibleServersFirewallRule_Spec{}

// ConvertSpecFrom populates our FlexibleServersFirewallRule_Spec from the provided source
func (rule *FlexibleServersFirewallRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20210501s.FlexibleServersFirewallRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignPropertiesFromFlexibleServersFirewallRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &alpha20210501s.FlexibleServersFirewallRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignPropertiesFromFlexibleServersFirewallRule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersFirewallRule_Spec
func (rule *FlexibleServersFirewallRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20210501s.FlexibleServersFirewallRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignPropertiesToFlexibleServersFirewallRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210501s.FlexibleServersFirewallRule_Spec{}
	err := rule.AssignPropertiesToFlexibleServersFirewallRule_Spec(dst)
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

// AssignPropertiesFromFlexibleServersFirewallRule_Spec populates our FlexibleServersFirewallRule_Spec from the provided source FlexibleServersFirewallRule_Spec
func (rule *FlexibleServersFirewallRule_Spec) AssignPropertiesFromFlexibleServersFirewallRule_Spec(source *alpha20210501s.FlexibleServersFirewallRule_Spec) error {

	// AzureName
	rule.AzureName = source.AzureName

	// EndIpAddress
	if source.EndIpAddress != nil {
		endIpAddress := *source.EndIpAddress
		rule.EndIpAddress = &endIpAddress
	} else {
		rule.EndIpAddress = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rule.Owner = &owner
	} else {
		rule.Owner = nil
	}

	// StartIpAddress
	if source.StartIpAddress != nil {
		startIpAddress := *source.StartIpAddress
		rule.StartIpAddress = &startIpAddress
	} else {
		rule.StartIpAddress = nil
	}

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersFirewallRule_Spec populates the provided destination FlexibleServersFirewallRule_Spec from our FlexibleServersFirewallRule_Spec
func (rule *FlexibleServersFirewallRule_Spec) AssignPropertiesToFlexibleServersFirewallRule_Spec(destination *alpha20210501s.FlexibleServersFirewallRule_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rule.AzureName

	// EndIpAddress
	if rule.EndIpAddress != nil {
		endIpAddress := *rule.EndIpAddress
		destination.EndIpAddress = &endIpAddress
	} else {
		destination.EndIpAddress = nil
	}

	// OriginalVersion
	destination.OriginalVersion = rule.OriginalVersion()

	// Owner
	if rule.Owner != nil {
		owner := rule.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// StartIpAddress
	if rule.StartIpAddress != nil {
		startIpAddress := *rule.StartIpAddress
		destination.StartIpAddress = &startIpAddress
	} else {
		destination.StartIpAddress = nil
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
func (rule *FlexibleServersFirewallRule_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rule *FlexibleServersFirewallRule_Spec) SetAzureName(azureName string) {
	rule.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&FlexibleServersFirewallRule{}, &FlexibleServersFirewallRuleList{})
}
