// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20201201storage"
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
type RedisLinkedServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisLinkedServers_SPEC                  `json:"spec,omitempty"`
	Status            RedisLinkedServerCreateParameters_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisLinkedServer{}

// GetConditions returns the conditions of the resource
func (server *RedisLinkedServer) GetConditions() conditions.Conditions {
	return server.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (server *RedisLinkedServer) SetConditions(conditions conditions.Conditions) {
	server.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisLinkedServer{}

// ConvertFrom populates our RedisLinkedServer from the provided hub RedisLinkedServer
func (server *RedisLinkedServer) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20201201storage.RedisLinkedServer)
	if !ok {
		return fmt.Errorf("expected storage:cache/v1alpha1api20201201storage/RedisLinkedServer but received %T instead", hub)
	}

	return server.AssignPropertiesFromRedisLinkedServer(source)
}

// ConvertTo populates the provided hub RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20201201storage.RedisLinkedServer)
	if !ok {
		return fmt.Errorf("expected storage:cache/v1alpha1api20201201storage/RedisLinkedServer but received %T instead", hub)
	}

	return server.AssignPropertiesToRedisLinkedServer(destination)
}

// +kubebuilder:webhook:path=/mutate-cache-azure-com-v1alpha1api20201201-redislinkedserver,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redislinkedservers,verbs=create;update,versions=v1alpha1api20201201,name=default.v1alpha1api20201201.redislinkedservers.cache.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &RedisLinkedServer{}

// Default applies defaults to the RedisLinkedServer resource
func (server *RedisLinkedServer) Default() {
	server.defaultImpl()
	var temp interface{} = server
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (server *RedisLinkedServer) defaultAzureName() {
	if server.Spec.AzureName == "" {
		server.Spec.AzureName = server.Name
	}
}

// defaultImpl applies the code generated defaults to the RedisLinkedServer resource
func (server *RedisLinkedServer) defaultImpl() { server.defaultAzureName() }

var _ genruntime.KubernetesResource = &RedisLinkedServer{}

// AzureName returns the Azure name of the resource
func (server *RedisLinkedServer) AzureName() string {
	return server.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (server RedisLinkedServer) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (server *RedisLinkedServer) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (server *RedisLinkedServer) GetSpec() genruntime.ConvertibleSpec {
	return &server.Spec
}

// GetStatus returns the status of this resource
func (server *RedisLinkedServer) GetStatus() genruntime.ConvertibleStatus {
	return &server.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (server *RedisLinkedServer) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (server *RedisLinkedServer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisLinkedServerCreateParameters_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (server *RedisLinkedServer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(server.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  server.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (server *RedisLinkedServer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisLinkedServerCreateParameters_Status); ok {
		server.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisLinkedServerCreateParameters_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	server.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-cache-azure-com-v1alpha1api20201201-redislinkedserver,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redislinkedservers,verbs=create;update,versions=v1alpha1api20201201,name=validate.v1alpha1api20201201.redislinkedservers.cache.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &RedisLinkedServer{}

// ValidateCreate validates the creation of the resource
func (server *RedisLinkedServer) ValidateCreate() error {
	validations := server.createValidations()
	var temp interface{} = server
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
func (server *RedisLinkedServer) ValidateDelete() error {
	validations := server.deleteValidations()
	var temp interface{} = server
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
func (server *RedisLinkedServer) ValidateUpdate(old runtime.Object) error {
	validations := server.updateValidations()
	var temp interface{} = server
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
func (server *RedisLinkedServer) createValidations() []func() error {
	return []func() error{server.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (server *RedisLinkedServer) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (server *RedisLinkedServer) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return server.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (server *RedisLinkedServer) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&server.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromRedisLinkedServer populates our RedisLinkedServer from the provided source RedisLinkedServer
func (server *RedisLinkedServer) AssignPropertiesFromRedisLinkedServer(source *v1alpha1api20201201storage.RedisLinkedServer) error {

	// ObjectMeta
	server.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisLinkedServers_SPEC
	err := spec.AssignPropertiesFromRedisLinkedServersSPEC(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisLinkedServersSPEC() to populate field Spec")
	}
	server.Spec = spec

	// Status
	var status RedisLinkedServerCreateParameters_Status
	err = status.AssignPropertiesFromRedisLinkedServerCreateParametersStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisLinkedServerCreateParametersStatus() to populate field Status")
	}
	server.Status = status

	// No error
	return nil
}

// AssignPropertiesToRedisLinkedServer populates the provided destination RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) AssignPropertiesToRedisLinkedServer(destination *v1alpha1api20201201storage.RedisLinkedServer) error {

	// ObjectMeta
	destination.ObjectMeta = *server.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20201201storage.RedisLinkedServers_SPEC
	err := server.Spec.AssignPropertiesToRedisLinkedServersSPEC(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisLinkedServersSPEC() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20201201storage.RedisLinkedServerCreateParameters_Status
	err = server.Status.AssignPropertiesToRedisLinkedServerCreateParametersStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisLinkedServerCreateParametersStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (server *RedisLinkedServer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: server.Spec.OriginalVersion(),
		Kind:    "RedisLinkedServer",
	}
}

// +kubebuilder:object:root=true
type RedisLinkedServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisLinkedServer `json:"items"`
}

type RedisLinkedServerCreateParameters_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//LinkedRedisCacheId: Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheId *string `json:"linkedRedisCacheId,omitempty"`

	//LinkedRedisCacheLocation: Location of the linked redis cache.
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	//ServerRole: Role of the linked server.
	ServerRole *RedisLinkedServerCreatePropertiesStatusServerRole `json:"serverRole,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisLinkedServerCreateParameters_Status{}

// ConvertStatusFrom populates our RedisLinkedServerCreateParameters_Status from the provided source
func (parameters *RedisLinkedServerCreateParameters_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20201201storage.RedisLinkedServerCreateParameters_Status)
	if ok {
		// Populate our instance from source
		return parameters.AssignPropertiesFromRedisLinkedServerCreateParametersStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20201201storage.RedisLinkedServerCreateParameters_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = parameters.AssignPropertiesFromRedisLinkedServerCreateParametersStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisLinkedServerCreateParameters_Status
func (parameters *RedisLinkedServerCreateParameters_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20201201storage.RedisLinkedServerCreateParameters_Status)
	if ok {
		// Populate destination from our instance
		return parameters.AssignPropertiesToRedisLinkedServerCreateParametersStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20201201storage.RedisLinkedServerCreateParameters_Status{}
	err := parameters.AssignPropertiesToRedisLinkedServerCreateParametersStatus(dst)
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

var _ genruntime.FromARMConverter = &RedisLinkedServerCreateParameters_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (parameters *RedisLinkedServerCreateParameters_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisLinkedServerCreateParameters_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (parameters *RedisLinkedServerCreateParameters_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisLinkedServerCreateParameters_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisLinkedServerCreateParameters_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘LinkedRedisCacheId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		parameters.LinkedRedisCacheId = &typedInput.Properties.LinkedRedisCacheId
	}

	// Set property ‘LinkedRedisCacheLocation’:
	// copying flattened property:
	if typedInput.Properties != nil {
		parameters.LinkedRedisCacheLocation = &typedInput.Properties.LinkedRedisCacheLocation
	}

	// Set property ‘ServerRole’:
	// copying flattened property:
	if typedInput.Properties != nil {
		parameters.ServerRole = &typedInput.Properties.ServerRole
	}

	// No error
	return nil
}

// AssignPropertiesFromRedisLinkedServerCreateParametersStatus populates our RedisLinkedServerCreateParameters_Status from the provided source RedisLinkedServerCreateParameters_Status
func (parameters *RedisLinkedServerCreateParameters_Status) AssignPropertiesFromRedisLinkedServerCreateParametersStatus(source *v1alpha1api20201201storage.RedisLinkedServerCreateParameters_Status) error {

	// Conditions
	parameters.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// LinkedRedisCacheId
	parameters.LinkedRedisCacheId = genruntime.ClonePointerToString(source.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	parameters.LinkedRedisCacheLocation = genruntime.ClonePointerToString(source.LinkedRedisCacheLocation)

	// ServerRole
	if source.ServerRole != nil {
		serverRole := RedisLinkedServerCreatePropertiesStatusServerRole(*source.ServerRole)
		parameters.ServerRole = &serverRole
	} else {
		parameters.ServerRole = nil
	}

	// No error
	return nil
}

// AssignPropertiesToRedisLinkedServerCreateParametersStatus populates the provided destination RedisLinkedServerCreateParameters_Status from our RedisLinkedServerCreateParameters_Status
func (parameters *RedisLinkedServerCreateParameters_Status) AssignPropertiesToRedisLinkedServerCreateParametersStatus(destination *v1alpha1api20201201storage.RedisLinkedServerCreateParameters_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(parameters.Conditions)

	// LinkedRedisCacheId
	destination.LinkedRedisCacheId = genruntime.ClonePointerToString(parameters.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = genruntime.ClonePointerToString(parameters.LinkedRedisCacheLocation)

	// ServerRole
	if parameters.ServerRole != nil {
		serverRole := string(*parameters.ServerRole)
		destination.ServerRole = &serverRole
	} else {
		destination.ServerRole = nil
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

type RedisLinkedServers_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	// +kubebuilder:validation:Required
	//LinkedRedisCacheLocation: Location of the linked redis cache.
	LinkedRedisCacheLocation string `json:"linkedRedisCacheLocation"`

	// +kubebuilder:validation:Required
	//LinkedRedisCacheReference: Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheReference genruntime.ResourceReference `armReference:"LinkedRedisCacheId" json:"linkedRedisCacheReference"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`

	// +kubebuilder:validation:Required
	//ServerRole: Role of the linked server.
	ServerRole RedisLinkedServerCreatePropertiesSpecServerRole `json:"serverRole"`
}

var _ genruntime.ARMTransformer = &RedisLinkedServers_SPEC{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (spec *RedisLinkedServers_SPEC) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if spec == nil {
		return nil, nil
	}
	var result RedisLinkedServers_SPECARM

	// Set property ‘AzureName’:
	result.AzureName = spec.AzureName

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	result.Properties.LinkedRedisCacheLocation = spec.LinkedRedisCacheLocation
	linkedRedisCacheIdARMID, err := resolved.ResolvedReferences.ARMIDOrErr(spec.LinkedRedisCacheReference)
	if err != nil {
		return nil, err
	}
	result.Properties.LinkedRedisCacheId = linkedRedisCacheIdARMID
	result.Properties.ServerRole = spec.ServerRole
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (spec *RedisLinkedServers_SPEC) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisLinkedServers_SPECARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (spec *RedisLinkedServers_SPEC) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisLinkedServers_SPECARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisLinkedServers_SPECARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	spec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘LinkedRedisCacheLocation’:
	// copying flattened property:
	spec.LinkedRedisCacheLocation = typedInput.Properties.LinkedRedisCacheLocation

	// no assignment for property ‘LinkedRedisCacheReference’

	// Set property ‘Owner’:
	spec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘ServerRole’:
	// copying flattened property:
	spec.ServerRole = typedInput.Properties.ServerRole

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RedisLinkedServers_SPEC{}

// ConvertSpecFrom populates our RedisLinkedServers_SPEC from the provided source
func (spec *RedisLinkedServers_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20201201storage.RedisLinkedServers_SPEC)
	if ok {
		// Populate our instance from source
		return spec.AssignPropertiesFromRedisLinkedServersSPEC(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20201201storage.RedisLinkedServers_SPEC{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = spec.AssignPropertiesFromRedisLinkedServersSPEC(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisLinkedServers_SPEC
func (spec *RedisLinkedServers_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20201201storage.RedisLinkedServers_SPEC)
	if ok {
		// Populate destination from our instance
		return spec.AssignPropertiesToRedisLinkedServersSPEC(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20201201storage.RedisLinkedServers_SPEC{}
	err := spec.AssignPropertiesToRedisLinkedServersSPEC(dst)
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

// AssignPropertiesFromRedisLinkedServersSPEC populates our RedisLinkedServers_SPEC from the provided source RedisLinkedServers_SPEC
func (spec *RedisLinkedServers_SPEC) AssignPropertiesFromRedisLinkedServersSPEC(source *v1alpha1api20201201storage.RedisLinkedServers_SPEC) error {

	// AzureName
	spec.AzureName = source.AzureName

	// LinkedRedisCacheLocation
	spec.LinkedRedisCacheLocation = genruntime.GetOptionalStringValue(source.LinkedRedisCacheLocation)

	// LinkedRedisCacheReference
	spec.LinkedRedisCacheReference = source.LinkedRedisCacheReference.Copy()

	// Owner
	spec.Owner = source.Owner.Copy()

	// ServerRole
	if source.ServerRole != nil {
		spec.ServerRole = RedisLinkedServerCreatePropertiesSpecServerRole(*source.ServerRole)
	} else {
		spec.ServerRole = ""
	}

	// No error
	return nil
}

// AssignPropertiesToRedisLinkedServersSPEC populates the provided destination RedisLinkedServers_SPEC from our RedisLinkedServers_SPEC
func (spec *RedisLinkedServers_SPEC) AssignPropertiesToRedisLinkedServersSPEC(destination *v1alpha1api20201201storage.RedisLinkedServers_SPEC) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = spec.AzureName

	// LinkedRedisCacheLocation
	linkedRedisCacheLocation := spec.LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = &linkedRedisCacheLocation

	// LinkedRedisCacheReference
	destination.LinkedRedisCacheReference = spec.LinkedRedisCacheReference.Copy()

	// OriginalVersion
	destination.OriginalVersion = spec.OriginalVersion()

	// Owner
	destination.Owner = spec.Owner.Copy()

	// ServerRole
	serverRole := string(spec.ServerRole)
	destination.ServerRole = &serverRole

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
func (spec *RedisLinkedServers_SPEC) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (spec *RedisLinkedServers_SPEC) SetAzureName(azureName string) { spec.AzureName = azureName }

// +kubebuilder:validation:Enum={"Primary","Secondary"}
type RedisLinkedServerCreatePropertiesSpecServerRole string

const (
	RedisLinkedServerCreatePropertiesSpecServerRolePrimary   = RedisLinkedServerCreatePropertiesSpecServerRole("Primary")
	RedisLinkedServerCreatePropertiesSpecServerRoleSecondary = RedisLinkedServerCreatePropertiesSpecServerRole("Secondary")
)

type RedisLinkedServerCreatePropertiesStatusServerRole string

const (
	RedisLinkedServerCreatePropertiesStatusServerRolePrimary   = RedisLinkedServerCreatePropertiesStatusServerRole("Primary")
	RedisLinkedServerCreatePropertiesStatusServerRoleSecondary = RedisLinkedServerCreatePropertiesStatusServerRole("Secondary")
)

func init() {
	SchemeBuilder.Register(&RedisLinkedServer{}, &RedisLinkedServerList{})
}