// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccountsSqlDatabasesContainersThroughputSetting_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties to update Azure Cosmos DB resource throughput.
	Properties *ThroughputSettingsUpdatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                      `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsSqlDatabasesContainersThroughputSetting_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting DatabaseAccountsSqlDatabasesContainersThroughputSetting_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (setting *DatabaseAccountsSqlDatabasesContainersThroughputSetting_SpecARM) GetName() string {
	return setting.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (setting *DatabaseAccountsSqlDatabasesContainersThroughputSetting_SpecARM) GetType() string {
	return ""
}
