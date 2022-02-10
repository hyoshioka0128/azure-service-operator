// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515storage"
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
type SqlDatabaseContainerStoredProcedure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC `json:"spec,omitempty"`
	Status            SqlStoredProcedureCreateUpdateParameters_Status             `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerStoredProcedure{}

// GetConditions returns the conditions of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetConditions() conditions.Conditions {
	return procedure.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (procedure *SqlDatabaseContainerStoredProcedure) SetConditions(conditions conditions.Conditions) {
	procedure.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseContainerStoredProcedure{}

// ConvertFrom populates our SqlDatabaseContainerStoredProcedure from the provided hub SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20210515storage.SqlDatabaseContainerStoredProcedure)
	if !ok {
		return fmt.Errorf("expected storage:documentdb/v1alpha1api20210515storage/SqlDatabaseContainerStoredProcedure but received %T instead", hub)
	}

	return procedure.AssignPropertiesFromSqlDatabaseContainerStoredProcedure(source)
}

// ConvertTo populates the provided hub SqlDatabaseContainerStoredProcedure from our SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20210515storage.SqlDatabaseContainerStoredProcedure)
	if !ok {
		return fmt.Errorf("expected storage:documentdb/v1alpha1api20210515storage/SqlDatabaseContainerStoredProcedure but received %T instead", hub)
	}

	return procedure.AssignPropertiesToSqlDatabaseContainerStoredProcedure(destination)
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1alpha1api20210515-sqldatabasecontainerstoredprocedure,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontainerstoredprocedures,verbs=create;update,versions=v1alpha1api20210515,name=default.v1alpha1api20210515.sqldatabasecontainerstoredprocedures.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &SqlDatabaseContainerStoredProcedure{}

// Default applies defaults to the SqlDatabaseContainerStoredProcedure resource
func (procedure *SqlDatabaseContainerStoredProcedure) Default() {
	procedure.defaultImpl()
	var temp interface{} = procedure
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (procedure *SqlDatabaseContainerStoredProcedure) defaultAzureName() {
	if procedure.Spec.AzureName == "" {
		procedure.Spec.AzureName = procedure.Name
	}
}

// defaultImpl applies the code generated defaults to the SqlDatabaseContainerStoredProcedure resource
func (procedure *SqlDatabaseContainerStoredProcedure) defaultImpl() { procedure.defaultAzureName() }

var _ genruntime.KubernetesResource = &SqlDatabaseContainerStoredProcedure{}

// AzureName returns the Azure name of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) AzureName() string {
	return procedure.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (procedure SqlDatabaseContainerStoredProcedure) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetSpec() genruntime.ConvertibleSpec {
	return &procedure.Spec
}

// GetStatus returns the status of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetStatus() genruntime.ConvertibleStatus {
	return &procedure.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (procedure *SqlDatabaseContainerStoredProcedure) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (procedure *SqlDatabaseContainerStoredProcedure) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlStoredProcedureCreateUpdateParameters_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (procedure *SqlDatabaseContainerStoredProcedure) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(procedure.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  procedure.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlStoredProcedureCreateUpdateParameters_Status); ok {
		procedure.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlStoredProcedureCreateUpdateParameters_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	procedure.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1alpha1api20210515-sqldatabasecontainerstoredprocedure,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontainerstoredprocedures,verbs=create;update,versions=v1alpha1api20210515,name=validate.v1alpha1api20210515.sqldatabasecontainerstoredprocedures.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &SqlDatabaseContainerStoredProcedure{}

// ValidateCreate validates the creation of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) ValidateCreate() error {
	validations := procedure.createValidations()
	var temp interface{} = procedure
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
func (procedure *SqlDatabaseContainerStoredProcedure) ValidateDelete() error {
	validations := procedure.deleteValidations()
	var temp interface{} = procedure
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
func (procedure *SqlDatabaseContainerStoredProcedure) ValidateUpdate(old runtime.Object) error {
	validations := procedure.updateValidations()
	var temp interface{} = procedure
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
func (procedure *SqlDatabaseContainerStoredProcedure) createValidations() []func() error {
	return []func() error{procedure.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return procedure.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (procedure *SqlDatabaseContainerStoredProcedure) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&procedure.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromSqlDatabaseContainerStoredProcedure populates our SqlDatabaseContainerStoredProcedure from the provided source SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) AssignPropertiesFromSqlDatabaseContainerStoredProcedure(source *v1alpha1api20210515storage.SqlDatabaseContainerStoredProcedure) error {

	// ObjectMeta
	procedure.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC() to populate field Spec")
	}
	procedure.Spec = spec

	// Status
	var status SqlStoredProcedureCreateUpdateParameters_Status
	err = status.AssignPropertiesFromSqlStoredProcedureCreateUpdateParametersStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromSqlStoredProcedureCreateUpdateParametersStatus() to populate field Status")
	}
	procedure.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseContainerStoredProcedure populates the provided destination SqlDatabaseContainerStoredProcedure from our SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) AssignPropertiesToSqlDatabaseContainerStoredProcedure(destination *v1alpha1api20210515storage.SqlDatabaseContainerStoredProcedure) error {

	// ObjectMeta
	destination.ObjectMeta = *procedure.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC
	err := procedure.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210515storage.SqlStoredProcedureCreateUpdateParameters_Status
	err = procedure.Status.AssignPropertiesToSqlStoredProcedureCreateUpdateParametersStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSqlStoredProcedureCreateUpdateParametersStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (procedure *SqlDatabaseContainerStoredProcedure) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: procedure.Spec.OriginalVersion(),
		Kind:    "SqlDatabaseContainerStoredProcedure",
	}
}

// +kubebuilder:object:root=true
type SqlDatabaseContainerStoredProcedureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerStoredProcedure `json:"items"`
}

type DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Options: A key-value pair of options to be applied for the request. This
	//corresponds to the headers sent with the request.
	Options *CreateUpdateOptions_Spec `json:"options,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`

	// +kubebuilder:validation:Required
	//Resource: The standard JSON format of a storedProcedure
	Resource SqlStoredProcedureResource_Spec `json:"resource"`
	Tags     map[string]string               `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (spec *DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if spec == nil {
		return nil, nil
	}
	var result DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPECARM

	// Set property ‘AzureName’:
	result.AzureName = spec.AzureName

	// Set property ‘Location’:
	if spec.Location != nil {
		location := *spec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if spec.Options != nil {
		optionsARM, err := (*spec.Options).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		options := optionsARM.(CreateUpdateOptions_SpecARM)
		result.Properties.Options = &options
	}
	resourceARM, err := spec.Resource.ConvertToARM(resolved)
	if err != nil {
		return nil, err
	}
	result.Properties.Resource = resourceARM.(SqlStoredProcedureResource_SpecARM)

	// Set property ‘Tags’:
	if spec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range spec.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (spec *DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPECARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (spec *DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPECARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPECARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	spec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		spec.Location = &location
	}

	// Set property ‘Options’:
	// copying flattened property:
	if typedInput.Properties.Options != nil {
		var options1 CreateUpdateOptions_Spec
		err := options1.PopulateFromARM(owner, *typedInput.Properties.Options)
		if err != nil {
			return err
		}
		options := options1
		spec.Options = &options
	}

	// Set property ‘Owner’:
	spec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	var resource SqlStoredProcedureResource_Spec
	err := resource.PopulateFromARM(owner, typedInput.Properties.Resource)
	if err != nil {
		return err
	}
	spec.Resource = resource

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		spec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			spec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC from the provided source
func (spec *DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC)
	if ok {
		// Populate our instance from source
		return spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC
func (spec *DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC)
	if ok {
		// Populate destination from our instance
		return spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC{}
	err := spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC populates our DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC from the provided source DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC
func (spec *DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC) AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(source *v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC) error {

	// AzureName
	spec.AzureName = source.AzureName

	// Location
	spec.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions_Spec
		err := option.AssignPropertiesFromCreateUpdateOptionsSpec(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptionsSpec() to populate field Options")
		}
		spec.Options = &option
	} else {
		spec.Options = nil
	}

	// Owner
	spec.Owner = source.Owner.Copy()

	// Resource
	if source.Resource != nil {
		var resource SqlStoredProcedureResource_Spec
		err := resource.AssignPropertiesFromSqlStoredProcedureResourceSpec(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlStoredProcedureResourceSpec() to populate field Resource")
		}
		spec.Resource = resource
	} else {
		spec.Resource = SqlStoredProcedureResource_Spec{}
	}

	// Tags
	spec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC populates the provided destination DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC from our DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC
func (spec *DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC) AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(destination *v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = spec.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(spec.Location)

	// Options
	if spec.Options != nil {
		var option v1alpha1api20210515storage.CreateUpdateOptions_Spec
		err := spec.Options.AssignPropertiesToCreateUpdateOptionsSpec(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptionsSpec() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = spec.OriginalVersion()

	// Owner
	destination.Owner = spec.Owner.Copy()

	// Resource
	var resource v1alpha1api20210515storage.SqlStoredProcedureResource_Spec
	err := spec.Resource.AssignPropertiesToSqlStoredProcedureResourceSpec(&resource)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSqlStoredProcedureResourceSpec() to populate field Resource")
	}
	destination.Resource = &resource

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(spec.Tags)

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
func (spec *DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (spec *DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC) SetAzureName(azureName string) {
	spec.AzureName = azureName
}

type SqlStoredProcedureCreateUpdateParameters_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: The name of the ARM resource.
	Name *string `json:"name,omitempty"`

	//Options: A key-value pair of options to be applied for the request. This
	//corresponds to the headers sent with the request.
	Options *CreateUpdateOptions_Status `json:"options,omitempty"`

	//Resource: The standard JSON format of a storedProcedure
	Resource *SqlStoredProcedureResource_Status `json:"resource,omitempty"`
	Tags     map[string]string                  `json:"tags,omitempty"`

	//Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlStoredProcedureCreateUpdateParameters_Status{}

// ConvertStatusFrom populates our SqlStoredProcedureCreateUpdateParameters_Status from the provided source
func (parameters *SqlStoredProcedureCreateUpdateParameters_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20210515storage.SqlStoredProcedureCreateUpdateParameters_Status)
	if ok {
		// Populate our instance from source
		return parameters.AssignPropertiesFromSqlStoredProcedureCreateUpdateParametersStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210515storage.SqlStoredProcedureCreateUpdateParameters_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = parameters.AssignPropertiesFromSqlStoredProcedureCreateUpdateParametersStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlStoredProcedureCreateUpdateParameters_Status
func (parameters *SqlStoredProcedureCreateUpdateParameters_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20210515storage.SqlStoredProcedureCreateUpdateParameters_Status)
	if ok {
		// Populate destination from our instance
		return parameters.AssignPropertiesToSqlStoredProcedureCreateUpdateParametersStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210515storage.SqlStoredProcedureCreateUpdateParameters_Status{}
	err := parameters.AssignPropertiesToSqlStoredProcedureCreateUpdateParametersStatus(dst)
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

var _ genruntime.FromARMConverter = &SqlStoredProcedureCreateUpdateParameters_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (parameters *SqlStoredProcedureCreateUpdateParameters_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlStoredProcedureCreateUpdateParameters_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (parameters *SqlStoredProcedureCreateUpdateParameters_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlStoredProcedureCreateUpdateParameters_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlStoredProcedureCreateUpdateParameters_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		parameters.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		parameters.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		parameters.Name = &name
	}

	// Set property ‘Options’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Options != nil {
			var options1 CreateUpdateOptions_Status
			err := options1.PopulateFromARM(owner, *typedInput.Properties.Options)
			if err != nil {
				return err
			}
			options := options1
			parameters.Options = &options
		}
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		var temp SqlStoredProcedureResource_Status
		var temp1 SqlStoredProcedureResource_Status
		err := temp1.PopulateFromARM(owner, typedInput.Properties.Resource)
		if err != nil {
			return err
		}
		temp = temp1
		parameters.Resource = &temp
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		parameters.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			parameters.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		parameters.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlStoredProcedureCreateUpdateParametersStatus populates our SqlStoredProcedureCreateUpdateParameters_Status from the provided source SqlStoredProcedureCreateUpdateParameters_Status
func (parameters *SqlStoredProcedureCreateUpdateParameters_Status) AssignPropertiesFromSqlStoredProcedureCreateUpdateParametersStatus(source *v1alpha1api20210515storage.SqlStoredProcedureCreateUpdateParameters_Status) error {

	// Conditions
	parameters.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	parameters.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	parameters.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	parameters.Name = genruntime.ClonePointerToString(source.Name)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions_Status
		err := option.AssignPropertiesFromCreateUpdateOptionsStatus(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptionsStatus() to populate field Options")
		}
		parameters.Options = &option
	} else {
		parameters.Options = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlStoredProcedureResource_Status
		err := resource.AssignPropertiesFromSqlStoredProcedureResourceStatus(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlStoredProcedureResourceStatus() to populate field Resource")
		}
		parameters.Resource = &resource
	} else {
		parameters.Resource = nil
	}

	// Tags
	parameters.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	parameters.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToSqlStoredProcedureCreateUpdateParametersStatus populates the provided destination SqlStoredProcedureCreateUpdateParameters_Status from our SqlStoredProcedureCreateUpdateParameters_Status
func (parameters *SqlStoredProcedureCreateUpdateParameters_Status) AssignPropertiesToSqlStoredProcedureCreateUpdateParametersStatus(destination *v1alpha1api20210515storage.SqlStoredProcedureCreateUpdateParameters_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(parameters.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(parameters.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(parameters.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(parameters.Name)

	// Options
	if parameters.Options != nil {
		var option v1alpha1api20210515storage.CreateUpdateOptions_Status
		err := parameters.Options.AssignPropertiesToCreateUpdateOptionsStatus(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptionsStatus() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// Resource
	if parameters.Resource != nil {
		var resource v1alpha1api20210515storage.SqlStoredProcedureResource_Status
		err := parameters.Resource.AssignPropertiesToSqlStoredProcedureResourceStatus(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlStoredProcedureResourceStatus() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(parameters.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(parameters.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type SqlStoredProcedureResource_Spec struct {
	//Body: Body of the Stored Procedure
	Body *string `json:"body,omitempty"`

	// +kubebuilder:validation:Required
	//Id: Name of the Cosmos DB SQL storedProcedure
	Id string `json:"id"`
}

var _ genruntime.ARMTransformer = &SqlStoredProcedureResource_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (resource *SqlStoredProcedureResource_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if resource == nil {
		return nil, nil
	}
	var result SqlStoredProcedureResource_SpecARM

	// Set property ‘Body’:
	if resource.Body != nil {
		body := *resource.Body
		result.Body = &body
	}

	// Set property ‘Id’:
	result.Id = resource.Id
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlStoredProcedureResource_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlStoredProcedureResource_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlStoredProcedureResource_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlStoredProcedureResource_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlStoredProcedureResource_SpecARM, got %T", armInput)
	}

	// Set property ‘Body’:
	if typedInput.Body != nil {
		body := *typedInput.Body
		resource.Body = &body
	}

	// Set property ‘Id’:
	resource.Id = typedInput.Id

	// No error
	return nil
}

// AssignPropertiesFromSqlStoredProcedureResourceSpec populates our SqlStoredProcedureResource_Spec from the provided source SqlStoredProcedureResource_Spec
func (resource *SqlStoredProcedureResource_Spec) AssignPropertiesFromSqlStoredProcedureResourceSpec(source *v1alpha1api20210515storage.SqlStoredProcedureResource_Spec) error {

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Id
	resource.Id = genruntime.GetOptionalStringValue(source.Id)

	// No error
	return nil
}

// AssignPropertiesToSqlStoredProcedureResourceSpec populates the provided destination SqlStoredProcedureResource_Spec from our SqlStoredProcedureResource_Spec
func (resource *SqlStoredProcedureResource_Spec) AssignPropertiesToSqlStoredProcedureResourceSpec(destination *v1alpha1api20210515storage.SqlStoredProcedureResource_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Id
	id := resource.Id
	destination.Id = &id

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type SqlStoredProcedureResource_Status struct {
	//Body: Body of the Stored Procedure
	Body *string `json:"body,omitempty"`

	// +kubebuilder:validation:Required
	//Id: Name of the Cosmos DB SQL storedProcedure
	Id string `json:"id"`
}

var _ genruntime.FromARMConverter = &SqlStoredProcedureResource_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlStoredProcedureResource_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlStoredProcedureResource_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlStoredProcedureResource_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlStoredProcedureResource_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlStoredProcedureResource_StatusARM, got %T", armInput)
	}

	// Set property ‘Body’:
	if typedInput.Body != nil {
		body := *typedInput.Body
		resource.Body = &body
	}

	// Set property ‘Id’:
	resource.Id = typedInput.Id

	// No error
	return nil
}

// AssignPropertiesFromSqlStoredProcedureResourceStatus populates our SqlStoredProcedureResource_Status from the provided source SqlStoredProcedureResource_Status
func (resource *SqlStoredProcedureResource_Status) AssignPropertiesFromSqlStoredProcedureResourceStatus(source *v1alpha1api20210515storage.SqlStoredProcedureResource_Status) error {

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Id
	resource.Id = genruntime.GetOptionalStringValue(source.Id)

	// No error
	return nil
}

// AssignPropertiesToSqlStoredProcedureResourceStatus populates the provided destination SqlStoredProcedureResource_Status from our SqlStoredProcedureResource_Status
func (resource *SqlStoredProcedureResource_Status) AssignPropertiesToSqlStoredProcedureResourceStatus(destination *v1alpha1api20210515storage.SqlStoredProcedureResource_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Id
	id := resource.Id
	destination.Id = &id

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerStoredProcedure{}, &SqlDatabaseContainerStoredProcedureList{})
}