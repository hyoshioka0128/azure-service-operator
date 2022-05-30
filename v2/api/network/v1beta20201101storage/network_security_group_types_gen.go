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

// +kubebuilder:rbac:groups=network.azure.com,resources=networksecuritygroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={networksecuritygroups/status,networksecuritygroups/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20201101.NetworkSecurityGroup
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/networkSecurityGroup.json
type NetworkSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityGroup_Spec                                            `json:"spec,omitempty"`
	Status            NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NetworkSecurityGroup{}

// GetConditions returns the conditions of the resource
func (group *NetworkSecurityGroup) GetConditions() conditions.Conditions {
	return group.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (group *NetworkSecurityGroup) SetConditions(conditions conditions.Conditions) {
	group.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &NetworkSecurityGroup{}

// AzureName returns the Azure name of the resource
func (group *NetworkSecurityGroup) AzureName() string {
	return group.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (group NetworkSecurityGroup) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (group *NetworkSecurityGroup) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (group *NetworkSecurityGroup) GetSpec() genruntime.ConvertibleSpec {
	return &group.Spec
}

// GetStatus returns the status of this resource
func (group *NetworkSecurityGroup) GetStatus() genruntime.ConvertibleStatus {
	return &group.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (group *NetworkSecurityGroup) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (group *NetworkSecurityGroup) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (group *NetworkSecurityGroup) Owner() *genruntime.ResourceReference {
	ownerGroup, ownerKind := genruntime.LookupOwnerGroupKind(group.Spec)
	return &genruntime.ResourceReference{
		Group: ownerGroup,
		Kind:  ownerKind,
		Name:  group.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (group *NetworkSecurityGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded); ok {
		group.Status = *st
		return nil
	}

	// Convert status to required version
	var st NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	group.Status = st
	return nil
}

// Hub marks that this NetworkSecurityGroup is the hub type for conversion
func (group *NetworkSecurityGroup) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (group *NetworkSecurityGroup) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: group.Spec.OriginalVersion,
		Kind:    "NetworkSecurityGroup",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20201101.NetworkSecurityGroup
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/networkSecurityGroup.json
type NetworkSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityGroup `json:"items"`
}

// Storage version of v1beta20201101.NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded
type NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded struct {
	Conditions           []conditions.Condition                                             `json:"conditions,omitempty"`
	DefaultSecurityRules []SecurityRule_STATUS                                              `json:"defaultSecurityRules,omitempty"`
	Etag                 *string                                                            `json:"etag,omitempty"`
	FlowLogs             []FlowLog_STATUS                                                   `json:"flowLogs,omitempty"`
	Id                   *string                                                            `json:"id,omitempty"`
	Location             *string                                                            `json:"location,omitempty"`
	Name                 *string                                                            `json:"name,omitempty"`
	NetworkInterfaces    []NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded `json:"networkInterfaces,omitempty"`
	PropertyBag          genruntime.PropertyBag                                             `json:"$propertyBag,omitempty"`
	ProvisioningState    *string                                                            `json:"provisioningState,omitempty"`
	ResourceGuid         *string                                                            `json:"resourceGuid,omitempty"`
	SecurityRules        []SecurityRule_STATUS                                              `json:"securityRules,omitempty"`
	Subnets              []Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded           `json:"subnets,omitempty"`
	Tags                 map[string]string                                                  `json:"tags,omitempty"`
	Type                 *string                                                            `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded{}

// ConvertStatusFrom populates our NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded from the provided source
func (embedded *NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == embedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(embedded)
}

// ConvertStatusTo populates the provided destination from our NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded
func (embedded *NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == embedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(embedded)
}

// Storage version of v1beta20201101.NetworkSecurityGroup_Spec
type NetworkSecurityGroup_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference     *genruntime.ResourceReference                           `armReference:"Id" json:"reference,omitempty"`
	SecurityRules []SecurityRule_NetworkSecurityGroup_SubResourceEmbedded `json:"securityRules,omitempty"`
	Tags          map[string]string                                       `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NetworkSecurityGroup_Spec{}

// ConvertSpecFrom populates our NetworkSecurityGroup_Spec from the provided source
func (group *NetworkSecurityGroup_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(group)
}

// ConvertSpecTo populates the provided destination from our NetworkSecurityGroup_Spec
func (group *NetworkSecurityGroup_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(group)
}

// Storage version of v1beta20201101.FlowLog_STATUS
type FlowLog_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded
type NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded struct {
	ExtendedLocation *ExtendedLocation_STATUS `json:"extendedLocation,omitempty"`
	Id               *string                  `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.SecurityRule_NetworkSecurityGroup_SubResourceEmbedded
type SecurityRule_NetworkSecurityGroup_SubResourceEmbedded struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20201101.SecurityRule_STATUS
type SecurityRule_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded
type Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityGroup{}, &NetworkSecurityGroupList{})
}
