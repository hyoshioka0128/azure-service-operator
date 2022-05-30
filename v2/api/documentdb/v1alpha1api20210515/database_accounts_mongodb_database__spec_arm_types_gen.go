// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of DatabaseAccountsMongodbDatabase_Spec. Use v1beta20210515.DatabaseAccountsMongodbDatabase_Spec instead
type DatabaseAccountsMongodbDatabase_SpecARM struct {
	AzureName  string                                    `json:"azureName,omitempty"`
	Id         *string                                   `json:"id,omitempty"`
	Location   *string                                   `json:"location,omitempty"`
	Name       string                                    `json:"name,omitempty"`
	Properties *MongoDBDatabaseCreateUpdatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                         `json:"tags,omitempty"`
	Type       *string                                   `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsMongodbDatabase_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (database DatabaseAccountsMongodbDatabase_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (database *DatabaseAccountsMongodbDatabase_SpecARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (database *DatabaseAccountsMongodbDatabase_SpecARM) GetType() string {
	return ""
}

// Deprecated version of MongoDBDatabaseCreateUpdateProperties. Use v1beta20210515.MongoDBDatabaseCreateUpdateProperties instead
type MongoDBDatabaseCreateUpdatePropertiesARM struct {
	Options  *CreateUpdateOptionsARM     `json:"options,omitempty"`
	Resource *MongoDBDatabaseResourceARM `json:"resource,omitempty"`
}

// Deprecated version of CreateUpdateOptions. Use v1beta20210515.CreateUpdateOptions instead
type CreateUpdateOptionsARM struct {
	AutoscaleSettings *AutoscaleSettingsARM `json:"autoscaleSettings,omitempty"`
	Throughput        *int                  `json:"throughput,omitempty"`
}

// Deprecated version of MongoDBDatabaseResource. Use v1beta20210515.MongoDBDatabaseResource instead
type MongoDBDatabaseResourceARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of AutoscaleSettings. Use v1beta20210515.AutoscaleSettings instead
type AutoscaleSettingsARM struct {
	MaxThroughput *int `json:"maxThroughput,omitempty"`
}