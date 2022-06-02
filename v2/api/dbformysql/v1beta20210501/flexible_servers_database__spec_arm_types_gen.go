// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersDatabase_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`
	Name      string `json:"name,omitempty"`

	// Properties: The properties of a database.
	Properties *DatabasePropertiesARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersDatabase_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (database FlexibleServersDatabase_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (database *FlexibleServersDatabase_SpecARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/databases"
func (database *FlexibleServersDatabase_SpecARM) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/databases"
}

type DatabasePropertiesARM struct {
	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`
}