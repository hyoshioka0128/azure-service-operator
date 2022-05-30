// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ServersConfiguration_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name string `json:"name,omitempty"`

	// Properties: The properties of a configuration.
	Properties *ConfigurationPropertiesARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ServersConfiguration_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-06-01"
func (configuration ServersConfiguration_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (configuration *ServersConfiguration_SpecARM) GetName() string {
	return configuration.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (configuration *ServersConfiguration_SpecARM) GetType() string {
	return ""
}

type ConfigurationPropertiesARM struct {
	// AllowedValues: Allowed values of the configuration.
	AllowedValues *string `json:"allowedValues,omitempty"`

	// DataType: Data type of the configuration.
	DataType *string `json:"dataType,omitempty"`

	// DefaultValue: Default value of the configuration.
	DefaultValue *string `json:"defaultValue,omitempty"`

	// Description: Description of the configuration.
	Description *string `json:"description,omitempty"`

	// Source: Source of the configuration.
	Source *string `json:"source,omitempty"`

	// Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}