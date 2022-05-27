// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=routetablesroutes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={routetablesroutes/status,routetablesroutes/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20201101.RouteTablesRoute
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/routeTable.json
type RouteTablesRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTablesRoute_Spec   `json:"spec,omitempty"`
	Status            RouteTablesRoute_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RouteTablesRoute{}

// GetConditions returns the conditions of the resource
func (route *RouteTablesRoute) GetConditions() conditions.Conditions {
	return route.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (route *RouteTablesRoute) SetConditions(conditions conditions.Conditions) {
	route.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &RouteTablesRoute{}

// AzureName returns the Azure name of the resource
func (route *RouteTablesRoute) AzureName() string {
	return route.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "20201101"
func (route RouteTablesRoute) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (route *RouteTablesRoute) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (route *RouteTablesRoute) GetSpec() genruntime.ConvertibleSpec {
	return &route.Spec
}

// GetStatus returns the status of this resource
func (route *RouteTablesRoute) GetStatus() genruntime.ConvertibleStatus {
	return &route.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (route *RouteTablesRoute) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (route *RouteTablesRoute) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RouteTablesRoute_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (route *RouteTablesRoute) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(route.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  route.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (route *RouteTablesRoute) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RouteTablesRoute_STATUS); ok {
		route.Status = *st
		return nil
	}

	// Convert status to required version
	var st RouteTablesRoute_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	route.Status = st
	return nil
}

// Hub marks that this RouteTablesRoute is the hub type for conversion
func (route *RouteTablesRoute) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (route *RouteTablesRoute) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: route.Spec.OriginalVersion,
		Kind:    "RouteTablesRoute",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20201101.RouteTablesRoute
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/routeTable.json
type RouteTablesRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTablesRoute `json:"items"`
}

// Storage version of v1beta20201101.RouteTablesRoute_STATUS
type RouteTablesRoute_STATUS struct {
	AddressPrefix     *string                `json:"addressPrefix,omitempty"`
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Etag              *string                `json:"etag,omitempty"`
	HasBgpOverride    *bool                  `json:"hasBgpOverride,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	NextHopIpAddress  *string                `json:"nextHopIpAddress,omitempty"`
	NextHopType       *string                `json:"nextHopType,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RouteTablesRoute_STATUS{}

// ConvertStatusFrom populates our RouteTablesRoute_STATUS from the provided source
func (route *RouteTablesRoute_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == route {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(route)
}

// ConvertStatusTo populates the provided destination from our RouteTablesRoute_STATUS
func (route *RouteTablesRoute_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == route {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(route)
}

// Storage version of v1beta20201101.RouteTablesRoute_Spec
type RouteTablesRoute_Spec struct {
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string  `json:"azureName,omitempty"`
	Etag             *string `json:"etag,omitempty"`
	HasBgpOverride   *bool   `json:"hasBgpOverride,omitempty"`
	Id               *string `json:"id,omitempty"`
	NextHopIpAddress *string `json:"nextHopIpAddress,omitempty"`
	NextHopType      *string `json:"nextHopType,omitempty"`
	OriginalVersion  string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner             *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag       genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ProvisioningState *string                            `json:"provisioningState,omitempty"`
	Type              *string                            `json:"type,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RouteTablesRoute_Spec{}

// ConvertSpecFrom populates our RouteTablesRoute_Spec from the provided source
func (route *RouteTablesRoute_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == route {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(route)
}

// ConvertSpecTo populates the provided destination from our RouteTablesRoute_Spec
func (route *RouteTablesRoute_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == route {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(route)
}

func init() {
	SchemeBuilder.Register(&RouteTablesRoute{}, &RouteTablesRouteList{})
}
