// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of NetworkSecurityGroup_Spec. Use v1beta20201101.NetworkSecurityGroup_Spec instead
type NetworkSecurityGroup_SpecARM struct {
	AzureName  string                                   `json:"azureName,omitempty"`
	Etag       *string                                  `json:"etag,omitempty"`
	Id         *string                                  `json:"id,omitempty"`
	Location   *string                                  `json:"location,omitempty"`
	Name       string                                   `json:"name,omitempty"`
	Properties *NetworkSecurityGroupPropertiesFormatARM `json:"properties,omitempty"`
	Tags       map[string]string                        `json:"tags,omitempty"`
	Type       *string                                  `json:"type,omitempty"`
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

// Deprecated version of NetworkSecurityGroupPropertiesFormat. Use v1beta20201101.NetworkSecurityGroupPropertiesFormat instead
type NetworkSecurityGroupPropertiesFormatARM struct {
	DefaultSecurityRules []SecurityRuleARM                                                  `json:"defaultSecurityRules,omitempty"`
	FlowLogs             []FlowLogARM                                                       `json:"flowLogs,omitempty"`
	NetworkInterfaces    []NetworkInterfaceSpec_NetworkSecurityGroup_SubResourceEmbeddedARM `json:"networkInterfaces,omitempty"`
	ProvisioningState    *ProvisioningState                                                 `json:"provisioningState,omitempty"`
	ResourceGuid         *string                                                            `json:"resourceGuid,omitempty"`
	SecurityRules        []SecurityRuleARM                                                  `json:"securityRules,omitempty"`
	Subnets              []Subnet_NetworkSecurityGroup_SubResourceEmbeddedARM               `json:"subnets,omitempty"`
}

// Deprecated version of FlowLog. Use v1beta20201101.FlowLog instead
type FlowLogARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of NetworkInterfaceSpec_NetworkSecurityGroup_SubResourceEmbedded. Use v1beta20201101.NetworkInterfaceSpec_NetworkSecurityGroup_SubResourceEmbedded instead
type NetworkInterfaceSpec_NetworkSecurityGroup_SubResourceEmbeddedARM struct {
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`
	Id               *string              `json:"id,omitempty"`
}

// Deprecated version of SecurityRule. Use v1beta20201101.SecurityRule instead
type SecurityRuleARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of Subnet_NetworkSecurityGroup_SubResourceEmbedded. Use v1beta20201101.Subnet_NetworkSecurityGroup_SubResourceEmbedded instead
type Subnet_NetworkSecurityGroup_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}