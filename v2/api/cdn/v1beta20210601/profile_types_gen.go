// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

import (
	"fmt"
	v20210601s "github.com/Azure/azure-service-operator/v2/api/cdn/v1beta20210601storage"
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
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/cdn.json
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Profile_Spec   `json:"spec,omitempty"`
	Status            Profile_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Profile{}

// GetConditions returns the conditions of the resource
func (profile *Profile) GetConditions() conditions.Conditions {
	return profile.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (profile *Profile) SetConditions(conditions conditions.Conditions) {
	profile.Status.Conditions = conditions
}

var _ conversion.Convertible = &Profile{}

// ConvertFrom populates our Profile from the provided hub Profile
func (profile *Profile) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210601s.Profile)
	if !ok {
		return fmt.Errorf("expected cdn/v1beta20210601storage/Profile but received %T instead", hub)
	}

	return profile.AssignPropertiesFromProfile(source)
}

// ConvertTo populates the provided hub Profile from our Profile
func (profile *Profile) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210601s.Profile)
	if !ok {
		return fmt.Errorf("expected cdn/v1beta20210601storage/Profile but received %T instead", hub)
	}

	return profile.AssignPropertiesToProfile(destination)
}

// +kubebuilder:webhook:path=/mutate-cdn-azure-com-v1beta20210601-profile,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cdn.azure.com,resources=profiles,verbs=create;update,versions=v1beta20210601,name=default.v1beta20210601.profiles.cdn.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &Profile{}

// Default applies defaults to the Profile resource
func (profile *Profile) Default() {
	profile.defaultImpl()
	var temp interface{} = profile
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (profile *Profile) defaultAzureName() {
	if profile.Spec.AzureName == "" {
		profile.Spec.AzureName = profile.Name
	}
}

// defaultImpl applies the code generated defaults to the Profile resource
func (profile *Profile) defaultImpl() { profile.defaultAzureName() }

var _ genruntime.KubernetesResource = &Profile{}

// AzureName returns the Azure name of the resource
func (profile *Profile) AzureName() string {
	return profile.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (profile Profile) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (profile *Profile) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (profile *Profile) GetSpec() genruntime.ConvertibleSpec {
	return &profile.Spec
}

// GetStatus returns the status of this resource
func (profile *Profile) GetStatus() genruntime.ConvertibleStatus {
	return &profile.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (profile *Profile) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (profile *Profile) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Profile_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (profile *Profile) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(profile.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  profile.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (profile *Profile) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Profile_STATUS); ok {
		profile.Status = *st
		return nil
	}

	// Convert status to required version
	var st Profile_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	profile.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-cdn-azure-com-v1beta20210601-profile,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cdn.azure.com,resources=profiles,verbs=create;update,versions=v1beta20210601,name=validate.v1beta20210601.profiles.cdn.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &Profile{}

// ValidateCreate validates the creation of the resource
func (profile *Profile) ValidateCreate() error {
	validations := profile.createValidations()
	var temp interface{} = profile
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
func (profile *Profile) ValidateDelete() error {
	validations := profile.deleteValidations()
	var temp interface{} = profile
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
func (profile *Profile) ValidateUpdate(old runtime.Object) error {
	validations := profile.updateValidations()
	var temp interface{} = profile
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
func (profile *Profile) createValidations() []func() error {
	return []func() error{profile.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (profile *Profile) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (profile *Profile) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return profile.validateResourceReferences()
		},
		profile.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (profile *Profile) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&profile.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (profile *Profile) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*Profile)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, profile)
}

// AssignPropertiesFromProfile populates our Profile from the provided source Profile
func (profile *Profile) AssignPropertiesFromProfile(source *v20210601s.Profile) error {

	// ObjectMeta
	profile.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Profile_Spec
	err := spec.AssignPropertiesFromProfile_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromProfile_Spec() to populate field Spec")
	}
	profile.Spec = spec

	// Status
	var status Profile_STATUS
	err = status.AssignPropertiesFromProfile_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromProfile_STATUS() to populate field Status")
	}
	profile.Status = status

	// No error
	return nil
}

// AssignPropertiesToProfile populates the provided destination Profile from our Profile
func (profile *Profile) AssignPropertiesToProfile(destination *v20210601s.Profile) error {

	// ObjectMeta
	destination.ObjectMeta = *profile.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210601s.Profile_Spec
	err := profile.Spec.AssignPropertiesToProfile_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToProfile_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210601s.Profile_STATUS
	err = profile.Status.AssignPropertiesToProfile_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToProfile_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (profile *Profile) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: profile.Spec.OriginalVersion(),
		Kind:    "Profile",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/cdn.json
type ProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Profile `json:"items"`
}

// +kubebuilder:validation:Enum={"2021-06-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-06-01")

type Profile_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// FrontDoorId: The Id of the frontdoor.
	FrontDoorId *string `json:"frontDoorId,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Kind: Kind of the profile. Used by portal to differentiate traditional CDN profile and new AFD profile.
	Kind *string `json:"kind,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// OriginResponseTimeoutSeconds: Send and receive timeout on forwarding request to the origin. When timeout is reached, the
	// request fails and returns.
	OriginResponseTimeoutSeconds *int `json:"originResponseTimeoutSeconds,omitempty"`

	// ProvisioningState: Provisioning status of the profile.
	ProvisioningState *ProfileProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResourceState: Resource status of the profile.
	ResourceState *ProfileProperties_ResourceState_STATUS `json:"resourceState,omitempty"`

	// Sku: The pricing tier (defines Azure Front Door Standard or Premium or a CDN provider, feature list and rate) of the
	// profile.
	Sku        *Sku_STATUS        `json:"sku,omitempty"`
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Profile_STATUS{}

// ConvertStatusFrom populates our Profile_STATUS from the provided source
func (profile *Profile_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210601s.Profile_STATUS)
	if ok {
		// Populate our instance from source
		return profile.AssignPropertiesFromProfile_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210601s.Profile_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = profile.AssignPropertiesFromProfile_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Profile_STATUS
func (profile *Profile_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210601s.Profile_STATUS)
	if ok {
		// Populate destination from our instance
		return profile.AssignPropertiesToProfile_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210601s.Profile_STATUS{}
	err := profile.AssignPropertiesToProfile_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Profile_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (profile *Profile_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Profile_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (profile *Profile_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Profile_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Profile_STATUSARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘FrontDoorId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.FrontDoorId != nil {
			frontDoorId := *typedInput.Properties.FrontDoorId
			profile.FrontDoorId = &frontDoorId
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		profile.Id = &id
	}

	// Set property ‘Kind’:
	if typedInput.Kind != nil {
		kind := *typedInput.Kind
		profile.Kind = &kind
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		profile.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		profile.Name = &name
	}

	// Set property ‘OriginResponseTimeoutSeconds’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.OriginResponseTimeoutSeconds != nil {
			originResponseTimeoutSeconds := *typedInput.Properties.OriginResponseTimeoutSeconds
			profile.OriginResponseTimeoutSeconds = &originResponseTimeoutSeconds
		}
	}

	// Set property ‘ProvisioningState’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			profile.ProvisioningState = &provisioningState
		}
	}

	// Set property ‘ResourceState’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ResourceState != nil {
			resourceState := *typedInput.Properties.ResourceState
			profile.ResourceState = &resourceState
		}
	}

	// Set property ‘Sku’:
	if typedInput.Sku != nil {
		var sku1 Sku_STATUS
		err := sku1.PopulateFromARM(owner, *typedInput.Sku)
		if err != nil {
			return err
		}
		sku := sku1
		profile.Sku = &sku
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		profile.SystemData = &systemData
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		profile.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			profile.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		profile.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromProfile_STATUS populates our Profile_STATUS from the provided source Profile_STATUS
func (profile *Profile_STATUS) AssignPropertiesFromProfile_STATUS(source *v20210601s.Profile_STATUS) error {

	// Conditions
	profile.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// FrontDoorId
	profile.FrontDoorId = genruntime.ClonePointerToString(source.FrontDoorId)

	// Id
	profile.Id = genruntime.ClonePointerToString(source.Id)

	// Kind
	profile.Kind = genruntime.ClonePointerToString(source.Kind)

	// Location
	profile.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	profile.Name = genruntime.ClonePointerToString(source.Name)

	// OriginResponseTimeoutSeconds
	profile.OriginResponseTimeoutSeconds = genruntime.ClonePointerToInt(source.OriginResponseTimeoutSeconds)

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := ProfileProperties_ProvisioningState_STATUS(*source.ProvisioningState)
		profile.ProvisioningState = &provisioningState
	} else {
		profile.ProvisioningState = nil
	}

	// ResourceState
	if source.ResourceState != nil {
		resourceState := ProfileProperties_ResourceState_STATUS(*source.ResourceState)
		profile.ResourceState = &resourceState
	} else {
		profile.ResourceState = nil
	}

	// Sku
	if source.Sku != nil {
		var sku Sku_STATUS
		err := sku.AssignPropertiesFromSku_STATUS(source.Sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSku_STATUS() to populate field Sku")
		}
		profile.Sku = &sku
	} else {
		profile.Sku = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignPropertiesFromSystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemData_STATUS() to populate field SystemData")
		}
		profile.SystemData = &systemDatum
	} else {
		profile.SystemData = nil
	}

	// Tags
	profile.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	profile.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToProfile_STATUS populates the provided destination Profile_STATUS from our Profile_STATUS
func (profile *Profile_STATUS) AssignPropertiesToProfile_STATUS(destination *v20210601s.Profile_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(profile.Conditions)

	// FrontDoorId
	destination.FrontDoorId = genruntime.ClonePointerToString(profile.FrontDoorId)

	// Id
	destination.Id = genruntime.ClonePointerToString(profile.Id)

	// Kind
	destination.Kind = genruntime.ClonePointerToString(profile.Kind)

	// Location
	destination.Location = genruntime.ClonePointerToString(profile.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(profile.Name)

	// OriginResponseTimeoutSeconds
	destination.OriginResponseTimeoutSeconds = genruntime.ClonePointerToInt(profile.OriginResponseTimeoutSeconds)

	// ProvisioningState
	if profile.ProvisioningState != nil {
		provisioningState := string(*profile.ProvisioningState)
		destination.ProvisioningState = &provisioningState
	} else {
		destination.ProvisioningState = nil
	}

	// ResourceState
	if profile.ResourceState != nil {
		resourceState := string(*profile.ResourceState)
		destination.ResourceState = &resourceState
	} else {
		destination.ResourceState = nil
	}

	// Sku
	if profile.Sku != nil {
		var sku v20210601s.Sku_STATUS
		err := profile.Sku.AssignPropertiesToSku_STATUS(&sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSku_STATUS() to populate field Sku")
		}
		destination.Sku = &sku
	} else {
		destination.Sku = nil
	}

	// SystemData
	if profile.SystemData != nil {
		var systemDatum v20210601s.SystemData_STATUS
		err := profile.SystemData.AssignPropertiesToSystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(profile.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(profile.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type Profile_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Minimum=16
	// OriginResponseTimeoutSeconds: Send and receive timeout on forwarding request to the origin. When timeout is reached, the
	// request fails and returns.
	OriginResponseTimeoutSeconds *int `json:"originResponseTimeoutSeconds,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// +kubebuilder:validation:Required
	// Sku: The pricing tier (defines Azure Front Door Standard or Premium or a CDN provider, feature list and rate) of the
	// profile.
	Sku *Sku `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &Profile_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (profile *Profile_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if profile == nil {
		return nil, nil
	}
	result := &Profile_SpecARM{}

	// Set property ‘AzureName’:
	result.AzureName = profile.AzureName

	// Set property ‘Location’:
	if profile.Location != nil {
		location := *profile.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if profile.OriginResponseTimeoutSeconds != nil {
		result.Properties = &ProfilePropertiesARM{}
	}
	if profile.OriginResponseTimeoutSeconds != nil {
		originResponseTimeoutSeconds := *profile.OriginResponseTimeoutSeconds
		result.Properties.OriginResponseTimeoutSeconds = &originResponseTimeoutSeconds
	}

	// Set property ‘Sku’:
	if profile.Sku != nil {
		skuARM, err := (*profile.Sku).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		sku := *skuARM.(*SkuARM)
		result.Sku = &sku
	}

	// Set property ‘Tags’:
	if profile.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range profile.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (profile *Profile_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Profile_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (profile *Profile_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Profile_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Profile_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	profile.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		profile.Location = &location
	}

	// Set property ‘OriginResponseTimeoutSeconds’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.OriginResponseTimeoutSeconds != nil {
			originResponseTimeoutSeconds := *typedInput.Properties.OriginResponseTimeoutSeconds
			profile.OriginResponseTimeoutSeconds = &originResponseTimeoutSeconds
		}
	}

	// Set property ‘Owner’:
	profile.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Sku’:
	if typedInput.Sku != nil {
		var sku1 Sku
		err := sku1.PopulateFromARM(owner, *typedInput.Sku)
		if err != nil {
			return err
		}
		sku := sku1
		profile.Sku = &sku
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		profile.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			profile.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Profile_Spec{}

// ConvertSpecFrom populates our Profile_Spec from the provided source
func (profile *Profile_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210601s.Profile_Spec)
	if ok {
		// Populate our instance from source
		return profile.AssignPropertiesFromProfile_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210601s.Profile_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = profile.AssignPropertiesFromProfile_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Profile_Spec
func (profile *Profile_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210601s.Profile_Spec)
	if ok {
		// Populate destination from our instance
		return profile.AssignPropertiesToProfile_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210601s.Profile_Spec{}
	err := profile.AssignPropertiesToProfile_Spec(dst)
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

// AssignPropertiesFromProfile_Spec populates our Profile_Spec from the provided source Profile_Spec
func (profile *Profile_Spec) AssignPropertiesFromProfile_Spec(source *v20210601s.Profile_Spec) error {

	// AzureName
	profile.AzureName = source.AzureName

	// Location
	profile.Location = genruntime.ClonePointerToString(source.Location)

	// OriginResponseTimeoutSeconds
	if source.OriginResponseTimeoutSeconds != nil {
		originResponseTimeoutSecond := *source.OriginResponseTimeoutSeconds
		profile.OriginResponseTimeoutSeconds = &originResponseTimeoutSecond
	} else {
		profile.OriginResponseTimeoutSeconds = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		profile.Owner = &owner
	} else {
		profile.Owner = nil
	}

	// Sku
	if source.Sku != nil {
		var sku Sku
		err := sku.AssignPropertiesFromSku(source.Sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSku() to populate field Sku")
		}
		profile.Sku = &sku
	} else {
		profile.Sku = nil
	}

	// Tags
	profile.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToProfile_Spec populates the provided destination Profile_Spec from our Profile_Spec
func (profile *Profile_Spec) AssignPropertiesToProfile_Spec(destination *v20210601s.Profile_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = profile.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(profile.Location)

	// OriginResponseTimeoutSeconds
	if profile.OriginResponseTimeoutSeconds != nil {
		originResponseTimeoutSecond := *profile.OriginResponseTimeoutSeconds
		destination.OriginResponseTimeoutSeconds = &originResponseTimeoutSecond
	} else {
		destination.OriginResponseTimeoutSeconds = nil
	}

	// OriginalVersion
	destination.OriginalVersion = profile.OriginalVersion()

	// Owner
	if profile.Owner != nil {
		owner := profile.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Sku
	if profile.Sku != nil {
		var sku v20210601s.Sku
		err := profile.Sku.AssignPropertiesToSku(&sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSku() to populate field Sku")
		}
		destination.Sku = &sku
	} else {
		destination.Sku = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(profile.Tags)

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
func (profile *Profile_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (profile *Profile_Spec) SetAzureName(azureName string) { profile.AzureName = azureName }

type ProfileProperties_ProvisioningState_STATUS string

const (
	ProfileProperties_ProvisioningState_Creating_STATUS  = ProfileProperties_ProvisioningState_STATUS("Creating")
	ProfileProperties_ProvisioningState_Deleting_STATUS  = ProfileProperties_ProvisioningState_STATUS("Deleting")
	ProfileProperties_ProvisioningState_Failed_STATUS    = ProfileProperties_ProvisioningState_STATUS("Failed")
	ProfileProperties_ProvisioningState_Succeeded_STATUS = ProfileProperties_ProvisioningState_STATUS("Succeeded")
	ProfileProperties_ProvisioningState_Updating_STATUS  = ProfileProperties_ProvisioningState_STATUS("Updating")
)

type ProfileProperties_ResourceState_STATUS string

const (
	ProfileProperties_ResourceState_Active_STATUS   = ProfileProperties_ResourceState_STATUS("Active")
	ProfileProperties_ResourceState_Creating_STATUS = ProfileProperties_ResourceState_STATUS("Creating")
	ProfileProperties_ResourceState_Deleting_STATUS = ProfileProperties_ResourceState_STATUS("Deleting")
	ProfileProperties_ResourceState_Disabled_STATUS = ProfileProperties_ResourceState_STATUS("Disabled")
)

type Sku struct {
	// Name: Name of the pricing tier.
	Name *Sku_Name `json:"name,omitempty"`
}

var _ genruntime.ARMTransformer = &Sku{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (sku *Sku) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if sku == nil {
		return nil, nil
	}
	result := &SkuARM{}

	// Set property ‘Name’:
	if sku.Name != nil {
		name := *sku.Name
		result.Name = &name
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (sku *Sku) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SkuARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (sku *Sku) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SkuARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SkuARM, got %T", armInput)
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		sku.Name = &name
	}

	// No error
	return nil
}

// AssignPropertiesFromSku populates our Sku from the provided source Sku
func (sku *Sku) AssignPropertiesFromSku(source *v20210601s.Sku) error {

	// Name
	if source.Name != nil {
		name := Sku_Name(*source.Name)
		sku.Name = &name
	} else {
		sku.Name = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSku populates the provided destination Sku from our Sku
func (sku *Sku) AssignPropertiesToSku(destination *v20210601s.Sku) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Name
	if sku.Name != nil {
		name := string(*sku.Name)
		destination.Name = &name
	} else {
		destination.Name = nil
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

type Sku_STATUS struct {
	// Name: Name of the pricing tier.
	Name *Sku_Name_STATUS `json:"name,omitempty"`
}

var _ genruntime.FromARMConverter = &Sku_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (sku *Sku_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Sku_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (sku *Sku_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Sku_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Sku_STATUSARM, got %T", armInput)
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		sku.Name = &name
	}

	// No error
	return nil
}

// AssignPropertiesFromSku_STATUS populates our Sku_STATUS from the provided source Sku_STATUS
func (sku *Sku_STATUS) AssignPropertiesFromSku_STATUS(source *v20210601s.Sku_STATUS) error {

	// Name
	if source.Name != nil {
		name := Sku_Name_STATUS(*source.Name)
		sku.Name = &name
	} else {
		sku.Name = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSku_STATUS populates the provided destination Sku_STATUS from our Sku_STATUS
func (sku *Sku_STATUS) AssignPropertiesToSku_STATUS(destination *v20210601s.Sku_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Name
	if sku.Name != nil {
		name := string(*sku.Name)
		destination.Name = &name
	} else {
		destination.Name = nil
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

type SystemData_STATUS struct {
	// CreatedAt: The timestamp of resource creation (UTC)
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: An identifier for the identity that created the resource
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource
	CreatedByType *IdentityType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: An identifier for the identity that last modified the resource
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource
	LastModifiedByType *IdentityType_STATUS `json:"lastModifiedByType,omitempty"`
}

var _ genruntime.FromARMConverter = &SystemData_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (data *SystemData_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SystemData_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (data *SystemData_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SystemData_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SystemData_STATUSARM, got %T", armInput)
	}

	// Set property ‘CreatedAt’:
	if typedInput.CreatedAt != nil {
		createdAt := *typedInput.CreatedAt
		data.CreatedAt = &createdAt
	}

	// Set property ‘CreatedBy’:
	if typedInput.CreatedBy != nil {
		createdBy := *typedInput.CreatedBy
		data.CreatedBy = &createdBy
	}

	// Set property ‘CreatedByType’:
	if typedInput.CreatedByType != nil {
		createdByType := *typedInput.CreatedByType
		data.CreatedByType = &createdByType
	}

	// Set property ‘LastModifiedAt’:
	if typedInput.LastModifiedAt != nil {
		lastModifiedAt := *typedInput.LastModifiedAt
		data.LastModifiedAt = &lastModifiedAt
	}

	// Set property ‘LastModifiedBy’:
	if typedInput.LastModifiedBy != nil {
		lastModifiedBy := *typedInput.LastModifiedBy
		data.LastModifiedBy = &lastModifiedBy
	}

	// Set property ‘LastModifiedByType’:
	if typedInput.LastModifiedByType != nil {
		lastModifiedByType := *typedInput.LastModifiedByType
		data.LastModifiedByType = &lastModifiedByType
	}

	// No error
	return nil
}

// AssignPropertiesFromSystemData_STATUS populates our SystemData_STATUS from the provided source SystemData_STATUS
func (data *SystemData_STATUS) AssignPropertiesFromSystemData_STATUS(source *v20210601s.SystemData_STATUS) error {

	// CreatedAt
	data.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// CreatedBy
	data.CreatedBy = genruntime.ClonePointerToString(source.CreatedBy)

	// CreatedByType
	if source.CreatedByType != nil {
		createdByType := IdentityType_STATUS(*source.CreatedByType)
		data.CreatedByType = &createdByType
	} else {
		data.CreatedByType = nil
	}

	// LastModifiedAt
	data.LastModifiedAt = genruntime.ClonePointerToString(source.LastModifiedAt)

	// LastModifiedBy
	data.LastModifiedBy = genruntime.ClonePointerToString(source.LastModifiedBy)

	// LastModifiedByType
	if source.LastModifiedByType != nil {
		lastModifiedByType := IdentityType_STATUS(*source.LastModifiedByType)
		data.LastModifiedByType = &lastModifiedByType
	} else {
		data.LastModifiedByType = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSystemData_STATUS populates the provided destination SystemData_STATUS from our SystemData_STATUS
func (data *SystemData_STATUS) AssignPropertiesToSystemData_STATUS(destination *v20210601s.SystemData_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(data.CreatedAt)

	// CreatedBy
	destination.CreatedBy = genruntime.ClonePointerToString(data.CreatedBy)

	// CreatedByType
	if data.CreatedByType != nil {
		createdByType := string(*data.CreatedByType)
		destination.CreatedByType = &createdByType
	} else {
		destination.CreatedByType = nil
	}

	// LastModifiedAt
	destination.LastModifiedAt = genruntime.ClonePointerToString(data.LastModifiedAt)

	// LastModifiedBy
	destination.LastModifiedBy = genruntime.ClonePointerToString(data.LastModifiedBy)

	// LastModifiedByType
	if data.LastModifiedByType != nil {
		lastModifiedByType := string(*data.LastModifiedByType)
		destination.LastModifiedByType = &lastModifiedByType
	} else {
		destination.LastModifiedByType = nil
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

func init() {
	SchemeBuilder.Register(&Profile{}, &ProfileList{})
}
