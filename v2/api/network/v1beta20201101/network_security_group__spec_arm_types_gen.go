// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NetworkSecurityGroup_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name string `json:"name,omitempty"`

	// Properties: Properties of the network security group.
	Properties *NetworkSecurityGroupPropertiesFormatARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NetworkSecurityGroup_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (group NetworkSecurityGroup_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (group *NetworkSecurityGroup_SpecARM) GetName() string {
	return group.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (group *NetworkSecurityGroup_SpecARM) GetType() string {
	return ""
}

type NetworkSecurityGroupPropertiesFormatARM struct {
	// DefaultSecurityRules: The default security rules of network security group.
	DefaultSecurityRules []SecurityRuleARM `json:"defaultSecurityRules,omitempty"`

	// FlowLogs: A collection of references to flow log resources.
	FlowLogs []FlowLogARM `json:"flowLogs,omitempty"`

	// NetworkInterfaces: A collection of references to network interfaces.
	NetworkInterfaces []NetworkInterfaceSpec_NetworkSecurityGroup_SubResourceEmbeddedARM `json:"networkInterfaces,omitempty"`

	// ProvisioningState: The provisioning state of the network security group resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`

	// ResourceGuid: The resource GUID property of the network security group resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`

	// SecurityRules: A collection of security rules of the network security group.
	SecurityRules []SecurityRuleARM `json:"securityRules,omitempty"`

	// Subnets: A collection of references to subnets.
	Subnets []Subnet_NetworkSecurityGroup_SubResourceEmbeddedARM `json:"subnets,omitempty"`
}

type FlowLogARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type NetworkInterfaceSpec_NetworkSecurityGroup_SubResourceEmbeddedARM struct {
	// ExtendedLocation: The extended location of the network interface.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type SecurityRuleARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type Subnet_NetworkSecurityGroup_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}