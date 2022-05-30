// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersDatabase_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name string `json:"name,omitempty"`

	// Properties: The properties of a database.
	Properties *DatabasePropertiesARM `json:"properties,omitempty"`

	// SystemData: The system metadata relating to this resource.
	SystemData *SystemDataARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersDatabase_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (database FlexibleServersDatabase_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (database *FlexibleServersDatabase_SpecARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (database *FlexibleServersDatabase_SpecARM) GetType() string {
	return ""
}

type DatabasePropertiesARM struct {
	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`
}