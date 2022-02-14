// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1alpha1api20210601storage"
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
//Generator information:
//- Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/postgresql.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/firewallRules/{firewallRuleName}
type FlexibleServersFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersFirewallRules_SPEC `json:"spec,omitempty"`
	Status            FirewallRule_Status               `json:"status,omitempty"`
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
	source, ok := hub.(*v1alpha1api20210601storage.FlexibleServersFirewallRule)
	if !ok {
		return fmt.Errorf("expected storage:dbforpostgresql/v1alpha1api20210601storage/FlexibleServersFirewallRule but received %T instead", hub)
	}

	return rule.AssignPropertiesFromFlexibleServersFirewallRule(source)
}

// ConvertTo populates the provided hub FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20210601storage.FlexibleServersFirewallRule)
	if !ok {
		return fmt.Errorf("expected storage:dbforpostgresql/v1alpha1api20210601storage/FlexibleServersFirewallRule but received %T instead", hub)
	}

	return rule.AssignPropertiesToFlexibleServersFirewallRule(destination)
}

// +kubebuilder:webhook:path=/mutate-dbforpostgresql-azure-com-v1alpha1api20210601-flexibleserversfirewallrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversfirewallrules,verbs=create;update,versions=v1alpha1api20210601,name=default.v1alpha1api20210601.flexibleserversfirewallrules.dbforpostgresql.azure.com,admissionReviewVersions=v1beta1

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

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (rule FlexibleServersFirewallRule) GetAPIVersion() string {
	return string(APIVersionValue)
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

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"
func (rule *FlexibleServersFirewallRule) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *FlexibleServersFirewallRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FirewallRule_Status{}
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
	if st, ok := status.(*FirewallRule_Status); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st FirewallRule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-dbforpostgresql-azure-com-v1alpha1api20210601-flexibleserversfirewallrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversfirewallrules,verbs=create;update,versions=v1alpha1api20210601,name=validate.v1alpha1api20210601.flexibleserversfirewallrules.dbforpostgresql.azure.com,admissionReviewVersions=v1beta1

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
	}
}

// validateResourceReferences validates all resource references
func (rule *FlexibleServersFirewallRule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromFlexibleServersFirewallRule populates our FlexibleServersFirewallRule from the provided source FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) AssignPropertiesFromFlexibleServersFirewallRule(source *v1alpha1api20210601storage.FlexibleServersFirewallRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersFirewallRules_SPEC
	err := spec.AssignPropertiesFromFlexibleServersFirewallRules_SPEC(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromFlexibleServersFirewallRules_SPEC() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status FirewallRule_Status
	err = status.AssignPropertiesFromFirewallRule_Status(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromFirewallRule_Status() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersFirewallRule populates the provided destination FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) AssignPropertiesToFlexibleServersFirewallRule(destination *v1alpha1api20210601storage.FlexibleServersFirewallRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20210601storage.FlexibleServersFirewallRules_SPEC
	err := rule.Spec.AssignPropertiesToFlexibleServersFirewallRules_SPEC(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToFlexibleServersFirewallRules_SPEC() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210601storage.FirewallRule_Status
	err = rule.Status.AssignPropertiesToFirewallRule_Status(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToFirewallRule_Status() to populate field Status")
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
//Generator information:
//- Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/postgresql.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/firewallRules/{firewallRuleName}
type FlexibleServersFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersFirewallRule `json:"items"`
}

type FirewallRule_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//EndIpAddress: The end IP address of the server firewall rule. Must be IPv4
	//format.
	EndIpAddress *string `json:"endIpAddress,omitempty"`

	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//StartIpAddress: The start IP address of the server firewall rule. Must be IPv4
	//format.
	StartIpAddress *string `json:"startIpAddress,omitempty"`

	//SystemData: The system metadata relating to this resource.
	SystemData *SystemData_Status `json:"systemData,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
	//"Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FirewallRule_Status{}

// ConvertStatusFrom populates our FirewallRule_Status from the provided source
func (rule *FirewallRule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20210601storage.FirewallRule_Status)
	if ok {
		// Populate our instance from source
		return rule.AssignPropertiesFromFirewallRule_Status(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210601storage.FirewallRule_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignPropertiesFromFirewallRule_Status(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FirewallRule_Status
func (rule *FirewallRule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20210601storage.FirewallRule_Status)
	if ok {
		// Populate destination from our instance
		return rule.AssignPropertiesToFirewallRule_Status(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210601storage.FirewallRule_Status{}
	err := rule.AssignPropertiesToFirewallRule_Status(dst)
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

var _ genruntime.FromARMConverter = &FirewallRule_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *FirewallRule_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FirewallRule_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *FirewallRule_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FirewallRule_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FirewallRule_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘EndIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		rule.EndIpAddress = &typedInput.Properties.EndIpAddress
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
		rule.StartIpAddress = &typedInput.Properties.StartIpAddress
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

// AssignPropertiesFromFirewallRule_Status populates our FirewallRule_Status from the provided source FirewallRule_Status
func (rule *FirewallRule_Status) AssignPropertiesFromFirewallRule_Status(source *v1alpha1api20210601storage.FirewallRule_Status) error {

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
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemData_Status(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemData_Status() to populate field SystemData")
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

// AssignPropertiesToFirewallRule_Status populates the provided destination FirewallRule_Status from our FirewallRule_Status
func (rule *FirewallRule_Status) AssignPropertiesToFirewallRule_Status(destination *v1alpha1api20210601storage.FirewallRule_Status) error {
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
		var systemDatum v1alpha1api20210601storage.SystemData_Status
		err := rule.SystemData.AssignPropertiesToSystemData_Status(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemData_Status() to populate field SystemData")
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

type FlexibleServersFirewallRules_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	//EndIpAddress: The end IP address of the server firewall rule. Must be IPv4
	//format.
	EndIpAddress string `json:"endIpAddress"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	//StartIpAddress: The start IP address of the server firewall rule. Must be IPv4
	//format.
	StartIpAddress string `json:"startIpAddress"`
}

var _ genruntime.ARMTransformer = &FlexibleServersFirewallRules_SPEC{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (spec *FlexibleServersFirewallRules_SPEC) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if spec == nil {
		return nil, nil
	}
	var result FlexibleServersFirewallRules_SPECARM

	// Set property ‘AzureName’:
	result.AzureName = spec.AzureName

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	result.Properties.EndIpAddress = spec.EndIpAddress
	result.Properties.StartIpAddress = spec.StartIpAddress
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (spec *FlexibleServersFirewallRules_SPEC) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FlexibleServersFirewallRules_SPECARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (spec *FlexibleServersFirewallRules_SPEC) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FlexibleServersFirewallRules_SPECARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FlexibleServersFirewallRules_SPECARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	spec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘EndIpAddress’:
	// copying flattened property:
	spec.EndIpAddress = typedInput.Properties.EndIpAddress

	// Set property ‘Owner’:
	spec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘StartIpAddress’:
	// copying flattened property:
	spec.StartIpAddress = typedInput.Properties.StartIpAddress

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &FlexibleServersFirewallRules_SPEC{}

// ConvertSpecFrom populates our FlexibleServersFirewallRules_SPEC from the provided source
func (spec *FlexibleServersFirewallRules_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210601storage.FlexibleServersFirewallRules_SPEC)
	if ok {
		// Populate our instance from source
		return spec.AssignPropertiesFromFlexibleServersFirewallRules_SPEC(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210601storage.FlexibleServersFirewallRules_SPEC{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = spec.AssignPropertiesFromFlexibleServersFirewallRules_SPEC(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersFirewallRules_SPEC
func (spec *FlexibleServersFirewallRules_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210601storage.FlexibleServersFirewallRules_SPEC)
	if ok {
		// Populate destination from our instance
		return spec.AssignPropertiesToFlexibleServersFirewallRules_SPEC(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210601storage.FlexibleServersFirewallRules_SPEC{}
	err := spec.AssignPropertiesToFlexibleServersFirewallRules_SPEC(dst)
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

// AssignPropertiesFromFlexibleServersFirewallRules_SPEC populates our FlexibleServersFirewallRules_SPEC from the provided source FlexibleServersFirewallRules_SPEC
func (spec *FlexibleServersFirewallRules_SPEC) AssignPropertiesFromFlexibleServersFirewallRules_SPEC(source *v1alpha1api20210601storage.FlexibleServersFirewallRules_SPEC) error {

	// AzureName
	spec.AzureName = source.AzureName

	// EndIpAddress
	if source.EndIpAddress != nil {
		spec.EndIpAddress = *source.EndIpAddress
	} else {
		spec.EndIpAddress = ""
	}

	// Owner
	spec.Owner = source.Owner.Copy()

	// StartIpAddress
	if source.StartIpAddress != nil {
		spec.StartIpAddress = *source.StartIpAddress
	} else {
		spec.StartIpAddress = ""
	}

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersFirewallRules_SPEC populates the provided destination FlexibleServersFirewallRules_SPEC from our FlexibleServersFirewallRules_SPEC
func (spec *FlexibleServersFirewallRules_SPEC) AssignPropertiesToFlexibleServersFirewallRules_SPEC(destination *v1alpha1api20210601storage.FlexibleServersFirewallRules_SPEC) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = spec.AzureName

	// EndIpAddress
	endIpAddress := spec.EndIpAddress
	destination.EndIpAddress = &endIpAddress

	// OriginalVersion
	destination.OriginalVersion = spec.OriginalVersion()

	// Owner
	destination.Owner = spec.Owner.Copy()

	// StartIpAddress
	startIpAddress := spec.StartIpAddress
	destination.StartIpAddress = &startIpAddress

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
func (spec *FlexibleServersFirewallRules_SPEC) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (spec *FlexibleServersFirewallRules_SPEC) SetAzureName(azureName string) {
	spec.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&FlexibleServersFirewallRule{}, &FlexibleServersFirewallRuleList{})
}
