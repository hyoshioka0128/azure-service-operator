// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200202storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=insights.azure.com,resources=components,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=insights.azure.com,resources={components/status,components/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20200202.Component
//Generator information:
//- Generated from: /applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/components_API.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}
type Component struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Components_SPEC                     `json:"spec,omitempty"`
	Status            ApplicationInsightsComponent_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Component{}

// GetConditions returns the conditions of the resource
func (component *Component) GetConditions() conditions.Conditions {
	return component.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (component *Component) SetConditions(conditions conditions.Conditions) {
	component.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Component{}

// AzureName returns the Azure name of the resource
func (component *Component) AzureName() string {
	return component.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-02-02"
func (component Component) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (component *Component) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (component *Component) GetSpec() genruntime.ConvertibleSpec {
	return &component.Spec
}

// GetStatus returns the status of this resource
func (component *Component) GetStatus() genruntime.ConvertibleStatus {
	return &component.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Insights/components"
func (component *Component) GetType() string {
	return "Microsoft.Insights/components"
}

// NewEmptyStatus returns a new empty (blank) status
func (component *Component) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ApplicationInsightsComponent_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (component *Component) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(component.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  component.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (component *Component) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ApplicationInsightsComponent_Status); ok {
		component.Status = *st
		return nil
	}

	// Convert status to required version
	var st ApplicationInsightsComponent_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	component.Status = st
	return nil
}

// Hub marks that this Component is the hub type for conversion
func (component *Component) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (component *Component) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: component.Spec.OriginalVersion,
		Kind:    "Component",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20200202.Component
//Generator information:
//- Generated from: /applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/components_API.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}
type ComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Component `json:"items"`
}

// +kubebuilder:validation:Enum={"2020-02-02"}
type APIVersion string

const APIVersionValue = APIVersion("2020-02-02")

//Storage version of v1alpha1api20200202.ApplicationInsightsComponent_Status
type ApplicationInsightsComponent_Status struct {
	AppId                           *string                            `json:"AppId,omitempty"`
	ApplicationId                   *string                            `json:"ApplicationId,omitempty"`
	Application_Type                *string                            `json:"Application_Type,omitempty"`
	Conditions                      []conditions.Condition             `json:"conditions,omitempty"`
	ConnectionString                *string                            `json:"ConnectionString,omitempty"`
	CreationDate                    *string                            `json:"CreationDate,omitempty"`
	DisableIpMasking                *bool                              `json:"DisableIpMasking,omitempty"`
	DisableLocalAuth                *bool                              `json:"DisableLocalAuth,omitempty"`
	Etag                            *string                            `json:"etag,omitempty"`
	Flow_Type                       *string                            `json:"Flow_Type,omitempty"`
	ForceCustomerStorageForProfiler *bool                              `json:"ForceCustomerStorageForProfiler,omitempty"`
	HockeyAppId                     *string                            `json:"HockeyAppId,omitempty"`
	HockeyAppToken                  *string                            `json:"HockeyAppToken,omitempty"`
	Id                              *string                            `json:"id,omitempty"`
	ImmediatePurgeDataOn30Days      *bool                              `json:"ImmediatePurgeDataOn30Days,omitempty"`
	IngestionMode                   *string                            `json:"IngestionMode,omitempty"`
	InstrumentationKey              *string                            `json:"InstrumentationKey,omitempty"`
	Kind                            *string                            `json:"kind,omitempty"`
	LaMigrationDate                 *string                            `json:"LaMigrationDate,omitempty"`
	Location                        *string                            `json:"location,omitempty"`
	Name                            *string                            `json:"name,omitempty"`
	PrivateLinkScopedResources      []PrivateLinkScopedResource_Status `json:"PrivateLinkScopedResources,omitempty"`
	PropertiesName                  *string                            `json:"properties_name,omitempty"`
	PropertyBag                     genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ProvisioningState               *string                            `json:"provisioningState,omitempty"`
	PublicNetworkAccessForIngestion *string                            `json:"publicNetworkAccessForIngestion,omitempty"`
	PublicNetworkAccessForQuery     *string                            `json:"publicNetworkAccessForQuery,omitempty"`
	Request_Source                  *string                            `json:"Request_Source,omitempty"`
	RetentionInDays                 *int                               `json:"RetentionInDays,omitempty"`
	SamplingPercentage              *float64                           `json:"SamplingPercentage,omitempty"`
	Tags                            *v1.JSON                           `json:"tags,omitempty"`
	TenantId                        *string                            `json:"TenantId,omitempty"`
	Type                            *string                            `json:"type,omitempty"`
	WorkspaceResourceId             *string                            `json:"WorkspaceResourceId,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ApplicationInsightsComponent_Status{}

// ConvertStatusFrom populates our ApplicationInsightsComponent_Status from the provided source
func (component *ApplicationInsightsComponent_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == component {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(component)
}

// ConvertStatusTo populates the provided destination from our ApplicationInsightsComponent_Status
func (component *ApplicationInsightsComponent_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == component {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(component)
}

//Storage version of v1alpha1api20200202.Components_SPEC
type Components_SPEC struct {
	Application_Type *string `json:"Application_Type,omitempty"`

	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName                       string  `json:"azureName"`
	DisableIpMasking                *bool   `json:"DisableIpMasking,omitempty"`
	DisableLocalAuth                *bool   `json:"DisableLocalAuth,omitempty"`
	Etag                            *string `json:"etag,omitempty"`
	Flow_Type                       *string `json:"Flow_Type,omitempty"`
	ForceCustomerStorageForProfiler *bool   `json:"ForceCustomerStorageForProfiler,omitempty"`
	HockeyAppId                     *string `json:"HockeyAppId,omitempty"`
	ImmediatePurgeDataOn30Days      *bool   `json:"ImmediatePurgeDataOn30Days,omitempty"`
	IngestionMode                   *string `json:"IngestionMode,omitempty"`
	Kind                            *string `json:"kind,omitempty"`
	Location                        *string `json:"location,omitempty"`
	OriginalVersion                 string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner                           genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag                     genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	PublicNetworkAccessForIngestion *string                           `json:"publicNetworkAccessForIngestion,omitempty"`
	PublicNetworkAccessForQuery     *string                           `json:"publicNetworkAccessForQuery,omitempty"`
	Request_Source                  *string                           `json:"Request_Source,omitempty"`
	RetentionInDays                 *int                              `json:"RetentionInDays,omitempty"`
	SamplingPercentage              *float64                          `json:"SamplingPercentage,omitempty"`
	Tags                            *v1.JSON                          `json:"tags,omitempty"`

	//WorkspaceResourceReference: Resource Id of the log analytics workspace which the
	//data will be ingested to. This property is required to create an application
	//with this API version. Applications from older versions will not have this
	//property.
	WorkspaceResourceReference *genruntime.ResourceReference `armReference:"WorkspaceResourceId" json:"workspaceResourceReference,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Components_SPEC{}

// ConvertSpecFrom populates our Components_SPEC from the provided source
func (spec *Components_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our Components_SPEC
func (spec *Components_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

//Storage version of v1alpha1api20200202.PrivateLinkScopedResource_Status
type PrivateLinkScopedResource_Status struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId  *string                `json:"ResourceId,omitempty"`
	ScopeId     *string                `json:"ScopeId,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Component{}, &ComponentList{})
}
