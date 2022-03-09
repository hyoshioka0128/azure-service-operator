// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NetworkSecurityGroup_SpecARM struct {
	AzureName string `json:"azureName"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name"`

	//Properties: Properties of the network security group.
	Properties *NetworkSecurityGroupPropertiesFormatARM `json:"properties,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NetworkSecurityGroup_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (group NetworkSecurityGroup_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (group NetworkSecurityGroup_SpecARM) GetName() string {
	return group.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/networkSecurityGroups"
func (group NetworkSecurityGroup_SpecARM) GetType() string {
	return "Microsoft.Network/networkSecurityGroups"
}

type NetworkSecurityGroupPropertiesFormatARM struct {
	//SecurityRules: A collection of security rules of the network security group.
	SecurityRules []genruntime.ResourceReference `json:"securityRules,omitempty"`
}
