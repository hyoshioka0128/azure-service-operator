// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200801preview

import (
	"fmt"
	alpha20200801ps "github.com/Azure/azure-service-operator/v2/api/authorization/v1alpha1api20200801previewstorage"
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
// Deprecated version of RoleAssignment. Use v1beta20200801preview.RoleAssignment instead
type RoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAssignment_Spec   `json:"spec,omitempty"`
	Status            RoleAssignment_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RoleAssignment{}

// GetConditions returns the conditions of the resource
func (assignment *RoleAssignment) GetConditions() conditions.Conditions {
	return assignment.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (assignment *RoleAssignment) SetConditions(conditions conditions.Conditions) {
	assignment.Status.Conditions = conditions
}

var _ conversion.Convertible = &RoleAssignment{}

// ConvertFrom populates our RoleAssignment from the provided hub RoleAssignment
func (assignment *RoleAssignment) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20200801ps.RoleAssignment

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = assignment.AssignPropertiesFromRoleAssignment(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to assignment")
	}

	return nil
}

// ConvertTo populates the provided hub RoleAssignment from our RoleAssignment
func (assignment *RoleAssignment) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20200801ps.RoleAssignment
	err := assignment.AssignPropertiesToRoleAssignment(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from assignment")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-authorization-azure-com-v1alpha1api20200801preview-roleassignment,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=authorization.azure.com,resources=roleassignments,verbs=create;update,versions=v1alpha1api20200801preview,name=default.v1alpha1api20200801preview.roleassignments.authorization.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &RoleAssignment{}

// Default applies defaults to the RoleAssignment resource
func (assignment *RoleAssignment) Default() {
	assignment.defaultImpl()
	var temp interface{} = assignment
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (assignment *RoleAssignment) defaultAzureName() {
	if assignment.Spec.AzureName == "" {
		assignment.Spec.AzureName = assignment.Name
	}
}

// defaultImpl applies the code generated defaults to the RoleAssignment resource
func (assignment *RoleAssignment) defaultImpl() { assignment.defaultAzureName() }

var _ genruntime.KubernetesResource = &RoleAssignment{}

// AzureName returns the Azure name of the resource
func (assignment *RoleAssignment) AzureName() string {
	return assignment.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "20200801preview"
func (assignment RoleAssignment) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (assignment *RoleAssignment) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (assignment *RoleAssignment) GetSpec() genruntime.ConvertibleSpec {
	return &assignment.Spec
}

// GetStatus returns the status of this resource
func (assignment *RoleAssignment) GetStatus() genruntime.ConvertibleStatus {
	return &assignment.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (assignment *RoleAssignment) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (assignment *RoleAssignment) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RoleAssignment_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (assignment *RoleAssignment) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(assignment.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  assignment.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (assignment *RoleAssignment) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RoleAssignment_STATUS); ok {
		assignment.Status = *st
		return nil
	}

	// Convert status to required version
	var st RoleAssignment_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	assignment.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-authorization-azure-com-v1alpha1api20200801preview-roleassignment,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=authorization.azure.com,resources=roleassignments,verbs=create;update,versions=v1alpha1api20200801preview,name=validate.v1alpha1api20200801preview.roleassignments.authorization.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &RoleAssignment{}

// ValidateCreate validates the creation of the resource
func (assignment *RoleAssignment) ValidateCreate() error {
	validations := assignment.createValidations()
	var temp interface{} = assignment
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
func (assignment *RoleAssignment) ValidateDelete() error {
	validations := assignment.deleteValidations()
	var temp interface{} = assignment
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
func (assignment *RoleAssignment) ValidateUpdate(old runtime.Object) error {
	validations := assignment.updateValidations()
	var temp interface{} = assignment
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
func (assignment *RoleAssignment) createValidations() []func() error {
	return []func() error{assignment.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (assignment *RoleAssignment) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (assignment *RoleAssignment) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return assignment.validateResourceReferences()
		},
		assignment.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (assignment *RoleAssignment) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&assignment.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (assignment *RoleAssignment) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*RoleAssignment)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, assignment)
}

// AssignPropertiesFromRoleAssignment populates our RoleAssignment from the provided source RoleAssignment
func (assignment *RoleAssignment) AssignPropertiesFromRoleAssignment(source *alpha20200801ps.RoleAssignment) error {

	// ObjectMeta
	assignment.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RoleAssignment_Spec
	err := spec.AssignPropertiesFromRoleAssignment_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRoleAssignment_Spec() to populate field Spec")
	}
	assignment.Spec = spec

	// Status
	var status RoleAssignment_STATUS
	err = status.AssignPropertiesFromRoleAssignment_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRoleAssignment_STATUS() to populate field Status")
	}
	assignment.Status = status

	// No error
	return nil
}

// AssignPropertiesToRoleAssignment populates the provided destination RoleAssignment from our RoleAssignment
func (assignment *RoleAssignment) AssignPropertiesToRoleAssignment(destination *alpha20200801ps.RoleAssignment) error {

	// ObjectMeta
	destination.ObjectMeta = *assignment.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20200801ps.RoleAssignment_Spec
	err := assignment.Spec.AssignPropertiesToRoleAssignment_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRoleAssignment_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20200801ps.RoleAssignment_STATUS
	err = assignment.Status.AssignPropertiesToRoleAssignment_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRoleAssignment_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (assignment *RoleAssignment) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: assignment.Spec.OriginalVersion(),
		Kind:    "RoleAssignment",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of RoleAssignment. Use v1beta20200801preview.RoleAssignment instead
type RoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssignment `json:"items"`
}

// Deprecated version of APIVersion. Use v1beta20200801preview.APIVersion instead
// +kubebuilder:validation:Enum={"20200801preview"}
type APIVersion string

const APIVersionValue = APIVersion("20200801preview")

// Deprecated version of RoleAssignment_STATUS. Use v1beta20200801preview.RoleAssignment_STATUS instead
type RoleAssignment_STATUS struct {
	Condition        *string `json:"condition,omitempty"`
	ConditionVersion *string `json:"conditionVersion,omitempty"`

	// Conditions: The observed state of the resource
	Conditions                         []conditions.Condition                         `json:"conditions,omitempty"`
	CreatedBy                          *string                                        `json:"createdBy,omitempty"`
	CreatedOn                          *string                                        `json:"createdOn,omitempty"`
	DelegatedManagedIdentityResourceId *string                                        `json:"delegatedManagedIdentityResourceId,omitempty"`
	Description                        *string                                        `json:"description,omitempty"`
	PrincipalId                        *string                                        `json:"principalId,omitempty"`
	PrincipalType                      *RoleAssignmentProperties_PrincipalType_STATUS `json:"principalType,omitempty"`
	RoleDefinitionId                   *string                                        `json:"roleDefinitionId,omitempty"`
	Scope                              *string                                        `json:"scope,omitempty"`
	UpdatedBy                          *string                                        `json:"updatedBy,omitempty"`
	UpdatedOn                          *string                                        `json:"updatedOn,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RoleAssignment_STATUS{}

// ConvertStatusFrom populates our RoleAssignment_STATUS from the provided source
func (assignment *RoleAssignment_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20200801ps.RoleAssignment_STATUS)
	if ok {
		// Populate our instance from source
		return assignment.AssignPropertiesFromRoleAssignment_STATUS(src)
	}

	// Convert to an intermediate form
	src = &alpha20200801ps.RoleAssignment_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = assignment.AssignPropertiesFromRoleAssignment_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RoleAssignment_STATUS
func (assignment *RoleAssignment_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20200801ps.RoleAssignment_STATUS)
	if ok {
		// Populate destination from our instance
		return assignment.AssignPropertiesToRoleAssignment_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20200801ps.RoleAssignment_STATUS{}
	err := assignment.AssignPropertiesToRoleAssignment_STATUS(dst)
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

var _ genruntime.FromARMConverter = &RoleAssignment_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (assignment *RoleAssignment_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RoleAssignment_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (assignment *RoleAssignment_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RoleAssignment_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RoleAssignment_STATUSARM, got %T", armInput)
	}

	// Set property ‘Condition’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Condition != nil {
			condition := *typedInput.Properties.Condition
			assignment.Condition = &condition
		}
	}

	// Set property ‘ConditionVersion’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ConditionVersion != nil {
			conditionVersion := *typedInput.Properties.ConditionVersion
			assignment.ConditionVersion = &conditionVersion
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘CreatedBy’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreatedBy != nil {
			createdBy := *typedInput.Properties.CreatedBy
			assignment.CreatedBy = &createdBy
		}
	}

	// Set property ‘CreatedOn’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreatedOn != nil {
			createdOn := *typedInput.Properties.CreatedOn
			assignment.CreatedOn = &createdOn
		}
	}

	// Set property ‘DelegatedManagedIdentityResourceId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DelegatedManagedIdentityResourceId != nil {
			delegatedManagedIdentityResourceId := *typedInput.Properties.DelegatedManagedIdentityResourceId
			assignment.DelegatedManagedIdentityResourceId = &delegatedManagedIdentityResourceId
		}
	}

	// Set property ‘Description’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			assignment.Description = &description
		}
	}

	// Set property ‘PrincipalId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalId != nil {
			principalId := *typedInput.Properties.PrincipalId
			assignment.PrincipalId = &principalId
		}
	}

	// Set property ‘PrincipalType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalType != nil {
			principalType := *typedInput.Properties.PrincipalType
			assignment.PrincipalType = &principalType
		}
	}

	// Set property ‘RoleDefinitionId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.RoleDefinitionId != nil {
			roleDefinitionId := *typedInput.Properties.RoleDefinitionId
			assignment.RoleDefinitionId = &roleDefinitionId
		}
	}

	// Set property ‘Scope’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Scope != nil {
			scope := *typedInput.Properties.Scope
			assignment.Scope = &scope
		}
	}

	// Set property ‘UpdatedBy’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UpdatedBy != nil {
			updatedBy := *typedInput.Properties.UpdatedBy
			assignment.UpdatedBy = &updatedBy
		}
	}

	// Set property ‘UpdatedOn’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UpdatedOn != nil {
			updatedOn := *typedInput.Properties.UpdatedOn
			assignment.UpdatedOn = &updatedOn
		}
	}

	// No error
	return nil
}

// AssignPropertiesFromRoleAssignment_STATUS populates our RoleAssignment_STATUS from the provided source RoleAssignment_STATUS
func (assignment *RoleAssignment_STATUS) AssignPropertiesFromRoleAssignment_STATUS(source *alpha20200801ps.RoleAssignment_STATUS) error {

	// Condition
	assignment.Condition = genruntime.ClonePointerToString(source.Condition)

	// ConditionVersion
	assignment.ConditionVersion = genruntime.ClonePointerToString(source.ConditionVersion)

	// Conditions
	assignment.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CreatedBy
	assignment.CreatedBy = genruntime.ClonePointerToString(source.CreatedBy)

	// CreatedOn
	assignment.CreatedOn = genruntime.ClonePointerToString(source.CreatedOn)

	// DelegatedManagedIdentityResourceId
	assignment.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(source.DelegatedManagedIdentityResourceId)

	// Description
	assignment.Description = genruntime.ClonePointerToString(source.Description)

	// PrincipalId
	assignment.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// PrincipalType
	if source.PrincipalType != nil {
		principalType := RoleAssignmentProperties_PrincipalType_STATUS(*source.PrincipalType)
		assignment.PrincipalType = &principalType
	} else {
		assignment.PrincipalType = nil
	}

	// RoleDefinitionId
	assignment.RoleDefinitionId = genruntime.ClonePointerToString(source.RoleDefinitionId)

	// Scope
	assignment.Scope = genruntime.ClonePointerToString(source.Scope)

	// UpdatedBy
	assignment.UpdatedBy = genruntime.ClonePointerToString(source.UpdatedBy)

	// UpdatedOn
	assignment.UpdatedOn = genruntime.ClonePointerToString(source.UpdatedOn)

	// No error
	return nil
}

// AssignPropertiesToRoleAssignment_STATUS populates the provided destination RoleAssignment_STATUS from our RoleAssignment_STATUS
func (assignment *RoleAssignment_STATUS) AssignPropertiesToRoleAssignment_STATUS(destination *alpha20200801ps.RoleAssignment_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Condition
	destination.Condition = genruntime.ClonePointerToString(assignment.Condition)

	// ConditionVersion
	destination.ConditionVersion = genruntime.ClonePointerToString(assignment.ConditionVersion)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(assignment.Conditions)

	// CreatedBy
	destination.CreatedBy = genruntime.ClonePointerToString(assignment.CreatedBy)

	// CreatedOn
	destination.CreatedOn = genruntime.ClonePointerToString(assignment.CreatedOn)

	// DelegatedManagedIdentityResourceId
	destination.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(assignment.DelegatedManagedIdentityResourceId)

	// Description
	destination.Description = genruntime.ClonePointerToString(assignment.Description)

	// PrincipalId
	destination.PrincipalId = genruntime.ClonePointerToString(assignment.PrincipalId)

	// PrincipalType
	if assignment.PrincipalType != nil {
		principalType := string(*assignment.PrincipalType)
		destination.PrincipalType = &principalType
	} else {
		destination.PrincipalType = nil
	}

	// RoleDefinitionId
	destination.RoleDefinitionId = genruntime.ClonePointerToString(assignment.RoleDefinitionId)

	// Scope
	destination.Scope = genruntime.ClonePointerToString(assignment.Scope)

	// UpdatedBy
	destination.UpdatedBy = genruntime.ClonePointerToString(assignment.UpdatedBy)

	// UpdatedOn
	destination.UpdatedOn = genruntime.ClonePointerToString(assignment.UpdatedOn)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type RoleAssignment_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                          string  `json:"azureName,omitempty"`
	Condition                          *string `json:"condition,omitempty"`
	ConditionVersion                   *string `json:"conditionVersion,omitempty"`
	CreatedBy                          *string `json:"createdBy,omitempty"`
	CreatedOn                          *string `json:"createdOn,omitempty"`
	DelegatedManagedIdentityResourceId *string `json:"delegatedManagedIdentityResourceId,omitempty"`
	Description                        *string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// +kubebuilder:validation:Required
	PrincipalId   *string                                 `json:"principalId,omitempty"`
	PrincipalType *RoleAssignmentProperties_PrincipalType `json:"principalType,omitempty"`

	// +kubebuilder:validation:Required
	RoleDefinitionReference *genruntime.ResourceReference `armReference:"RoleDefinitionId" json:"roleDefinitionReference,omitempty"`
	Scope                   *string                       `json:"scope,omitempty"`
	UpdatedBy               *string                       `json:"updatedBy,omitempty"`
	UpdatedOn               *string                       `json:"updatedOn,omitempty"`
}

var _ genruntime.ARMTransformer = &RoleAssignment_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (assignment *RoleAssignment_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if assignment == nil {
		return nil, nil
	}
	result := &RoleAssignment_SpecARM{}

	// Set property ‘AzureName’:
	result.AzureName = assignment.AzureName

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if assignment.Condition != nil ||
		assignment.ConditionVersion != nil ||
		assignment.CreatedBy != nil ||
		assignment.CreatedOn != nil ||
		assignment.DelegatedManagedIdentityResourceId != nil ||
		assignment.Description != nil ||
		assignment.PrincipalId != nil ||
		assignment.PrincipalType != nil ||
		assignment.RoleDefinitionReference != nil ||
		assignment.Scope != nil ||
		assignment.UpdatedBy != nil ||
		assignment.UpdatedOn != nil {
		result.Properties = &RoleAssignmentPropertiesARM{}
	}
	if assignment.Condition != nil {
		condition := *assignment.Condition
		result.Properties.Condition = &condition
	}
	if assignment.ConditionVersion != nil {
		conditionVersion := *assignment.ConditionVersion
		result.Properties.ConditionVersion = &conditionVersion
	}
	if assignment.CreatedBy != nil {
		createdBy := *assignment.CreatedBy
		result.Properties.CreatedBy = &createdBy
	}
	if assignment.CreatedOn != nil {
		createdOn := *assignment.CreatedOn
		result.Properties.CreatedOn = &createdOn
	}
	if assignment.DelegatedManagedIdentityResourceId != nil {
		delegatedManagedIdentityResourceId := *assignment.DelegatedManagedIdentityResourceId
		result.Properties.DelegatedManagedIdentityResourceId = &delegatedManagedIdentityResourceId
	}
	if assignment.Description != nil {
		description := *assignment.Description
		result.Properties.Description = &description
	}
	if assignment.PrincipalId != nil {
		principalId := *assignment.PrincipalId
		result.Properties.PrincipalId = &principalId
	}
	if assignment.PrincipalType != nil {
		principalType := *assignment.PrincipalType
		result.Properties.PrincipalType = &principalType
	}
	if assignment.RoleDefinitionReference != nil {
		roleDefinitionIdARMID, err := resolved.ResolvedReferences.ARMIDOrErr(*assignment.RoleDefinitionReference)
		if err != nil {
			return nil, err
		}
		roleDefinitionId := roleDefinitionIdARMID
		result.Properties.RoleDefinitionId = &roleDefinitionId
	}
	if assignment.Scope != nil {
		scope := *assignment.Scope
		result.Properties.Scope = &scope
	}
	if assignment.UpdatedBy != nil {
		updatedBy := *assignment.UpdatedBy
		result.Properties.UpdatedBy = &updatedBy
	}
	if assignment.UpdatedOn != nil {
		updatedOn := *assignment.UpdatedOn
		result.Properties.UpdatedOn = &updatedOn
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (assignment *RoleAssignment_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RoleAssignment_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (assignment *RoleAssignment_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RoleAssignment_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RoleAssignment_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	assignment.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Condition’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Condition != nil {
			condition := *typedInput.Properties.Condition
			assignment.Condition = &condition
		}
	}

	// Set property ‘ConditionVersion’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ConditionVersion != nil {
			conditionVersion := *typedInput.Properties.ConditionVersion
			assignment.ConditionVersion = &conditionVersion
		}
	}

	// Set property ‘CreatedBy’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreatedBy != nil {
			createdBy := *typedInput.Properties.CreatedBy
			assignment.CreatedBy = &createdBy
		}
	}

	// Set property ‘CreatedOn’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreatedOn != nil {
			createdOn := *typedInput.Properties.CreatedOn
			assignment.CreatedOn = &createdOn
		}
	}

	// Set property ‘DelegatedManagedIdentityResourceId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DelegatedManagedIdentityResourceId != nil {
			delegatedManagedIdentityResourceId := *typedInput.Properties.DelegatedManagedIdentityResourceId
			assignment.DelegatedManagedIdentityResourceId = &delegatedManagedIdentityResourceId
		}
	}

	// Set property ‘Description’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			assignment.Description = &description
		}
	}

	// Set property ‘Owner’:
	assignment.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘PrincipalId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalId != nil {
			principalId := *typedInput.Properties.PrincipalId
			assignment.PrincipalId = &principalId
		}
	}

	// Set property ‘PrincipalType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalType != nil {
			principalType := *typedInput.Properties.PrincipalType
			assignment.PrincipalType = &principalType
		}
	}

	// no assignment for property ‘RoleDefinitionReference’

	// Set property ‘Scope’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Scope != nil {
			scope := *typedInput.Properties.Scope
			assignment.Scope = &scope
		}
	}

	// Set property ‘UpdatedBy’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UpdatedBy != nil {
			updatedBy := *typedInput.Properties.UpdatedBy
			assignment.UpdatedBy = &updatedBy
		}
	}

	// Set property ‘UpdatedOn’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UpdatedOn != nil {
			updatedOn := *typedInput.Properties.UpdatedOn
			assignment.UpdatedOn = &updatedOn
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RoleAssignment_Spec{}

// ConvertSpecFrom populates our RoleAssignment_Spec from the provided source
func (assignment *RoleAssignment_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20200801ps.RoleAssignment_Spec)
	if ok {
		// Populate our instance from source
		return assignment.AssignPropertiesFromRoleAssignment_Spec(src)
	}

	// Convert to an intermediate form
	src = &alpha20200801ps.RoleAssignment_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = assignment.AssignPropertiesFromRoleAssignment_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RoleAssignment_Spec
func (assignment *RoleAssignment_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20200801ps.RoleAssignment_Spec)
	if ok {
		// Populate destination from our instance
		return assignment.AssignPropertiesToRoleAssignment_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20200801ps.RoleAssignment_Spec{}
	err := assignment.AssignPropertiesToRoleAssignment_Spec(dst)
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

// AssignPropertiesFromRoleAssignment_Spec populates our RoleAssignment_Spec from the provided source RoleAssignment_Spec
func (assignment *RoleAssignment_Spec) AssignPropertiesFromRoleAssignment_Spec(source *alpha20200801ps.RoleAssignment_Spec) error {

	// AzureName
	assignment.AzureName = source.AzureName

	// Condition
	assignment.Condition = genruntime.ClonePointerToString(source.Condition)

	// ConditionVersion
	assignment.ConditionVersion = genruntime.ClonePointerToString(source.ConditionVersion)

	// CreatedBy
	assignment.CreatedBy = genruntime.ClonePointerToString(source.CreatedBy)

	// CreatedOn
	if source.CreatedOn != nil {
		createdOn := *source.CreatedOn
		assignment.CreatedOn = &createdOn
	} else {
		assignment.CreatedOn = nil
	}

	// DelegatedManagedIdentityResourceId
	assignment.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(source.DelegatedManagedIdentityResourceId)

	// Description
	assignment.Description = genruntime.ClonePointerToString(source.Description)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		assignment.Owner = &owner
	} else {
		assignment.Owner = nil
	}

	// PrincipalId
	assignment.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// PrincipalType
	if source.PrincipalType != nil {
		principalType := RoleAssignmentProperties_PrincipalType(*source.PrincipalType)
		assignment.PrincipalType = &principalType
	} else {
		assignment.PrincipalType = nil
	}

	// RoleDefinitionReference
	if source.RoleDefinitionReference != nil {
		roleDefinitionReference := source.RoleDefinitionReference.Copy()
		assignment.RoleDefinitionReference = &roleDefinitionReference
	} else {
		assignment.RoleDefinitionReference = nil
	}

	// Scope
	assignment.Scope = genruntime.ClonePointerToString(source.Scope)

	// UpdatedBy
	assignment.UpdatedBy = genruntime.ClonePointerToString(source.UpdatedBy)

	// UpdatedOn
	if source.UpdatedOn != nil {
		updatedOn := *source.UpdatedOn
		assignment.UpdatedOn = &updatedOn
	} else {
		assignment.UpdatedOn = nil
	}

	// No error
	return nil
}

// AssignPropertiesToRoleAssignment_Spec populates the provided destination RoleAssignment_Spec from our RoleAssignment_Spec
func (assignment *RoleAssignment_Spec) AssignPropertiesToRoleAssignment_Spec(destination *alpha20200801ps.RoleAssignment_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = assignment.AzureName

	// Condition
	destination.Condition = genruntime.ClonePointerToString(assignment.Condition)

	// ConditionVersion
	destination.ConditionVersion = genruntime.ClonePointerToString(assignment.ConditionVersion)

	// CreatedBy
	destination.CreatedBy = genruntime.ClonePointerToString(assignment.CreatedBy)

	// CreatedOn
	if assignment.CreatedOn != nil {
		createdOn := *assignment.CreatedOn
		destination.CreatedOn = &createdOn
	} else {
		destination.CreatedOn = nil
	}

	// DelegatedManagedIdentityResourceId
	destination.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(assignment.DelegatedManagedIdentityResourceId)

	// Description
	destination.Description = genruntime.ClonePointerToString(assignment.Description)

	// OriginalVersion
	destination.OriginalVersion = assignment.OriginalVersion()

	// Owner
	if assignment.Owner != nil {
		owner := assignment.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// PrincipalId
	destination.PrincipalId = genruntime.ClonePointerToString(assignment.PrincipalId)

	// PrincipalType
	if assignment.PrincipalType != nil {
		principalType := string(*assignment.PrincipalType)
		destination.PrincipalType = &principalType
	} else {
		destination.PrincipalType = nil
	}

	// RoleDefinitionReference
	if assignment.RoleDefinitionReference != nil {
		roleDefinitionReference := assignment.RoleDefinitionReference.Copy()
		destination.RoleDefinitionReference = &roleDefinitionReference
	} else {
		destination.RoleDefinitionReference = nil
	}

	// Scope
	destination.Scope = genruntime.ClonePointerToString(assignment.Scope)

	// UpdatedBy
	destination.UpdatedBy = genruntime.ClonePointerToString(assignment.UpdatedBy)

	// UpdatedOn
	if assignment.UpdatedOn != nil {
		updatedOn := *assignment.UpdatedOn
		destination.UpdatedOn = &updatedOn
	} else {
		destination.UpdatedOn = nil
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
func (assignment *RoleAssignment_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (assignment *RoleAssignment_Spec) SetAzureName(azureName string) {
	assignment.AzureName = azureName
}

// Deprecated version of RoleAssignmentProperties_PrincipalType. Use
// v1beta20200801preview.RoleAssignmentProperties_PrincipalType instead
// +kubebuilder:validation:Enum={"ForeignGroup","Group","ServicePrincipal","User"}
type RoleAssignmentProperties_PrincipalType string

const (
	RoleAssignmentProperties_PrincipalTypeForeignGroup     = RoleAssignmentProperties_PrincipalType("ForeignGroup")
	RoleAssignmentProperties_PrincipalTypeGroup            = RoleAssignmentProperties_PrincipalType("Group")
	RoleAssignmentProperties_PrincipalTypeServicePrincipal = RoleAssignmentProperties_PrincipalType("ServicePrincipal")
	RoleAssignmentProperties_PrincipalTypeUser             = RoleAssignmentProperties_PrincipalType("User")
)

// Deprecated version of RoleAssignmentProperties_PrincipalType_STATUS. Use
// v1beta20200801preview.RoleAssignmentProperties_PrincipalType_STATUS instead
type RoleAssignmentProperties_PrincipalType_STATUS string

const (
	RoleAssignmentProperties_PrincipalType_STATUSForeignGroup     = RoleAssignmentProperties_PrincipalType_STATUS("ForeignGroup")
	RoleAssignmentProperties_PrincipalType_STATUSGroup            = RoleAssignmentProperties_PrincipalType_STATUS("Group")
	RoleAssignmentProperties_PrincipalType_STATUSServicePrincipal = RoleAssignmentProperties_PrincipalType_STATUS("ServicePrincipal")
	RoleAssignmentProperties_PrincipalType_STATUSUser             = RoleAssignmentProperties_PrincipalType_STATUS("User")
)

func init() {
	SchemeBuilder.Register(&RoleAssignment{}, &RoleAssignmentList{})
}
