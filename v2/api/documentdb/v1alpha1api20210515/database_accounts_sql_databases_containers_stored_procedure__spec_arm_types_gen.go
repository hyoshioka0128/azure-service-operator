// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec. Use v1beta20210515.DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec instead
type DatabaseAccountsSqlDatabasesContainersStoredProcedure_SpecARM struct {
	AzureName  string                                       `json:"azureName,omitempty"`
	Location   *string                                      `json:"location,omitempty"`
	Name       string                                       `json:"name,omitempty"`
	Properties *SqlStoredProcedureCreateUpdatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                            `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsSqlDatabasesContainersStoredProcedure_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (procedure DatabaseAccountsSqlDatabasesContainersStoredProcedure_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (procedure *DatabaseAccountsSqlDatabasesContainersStoredProcedure_SpecARM) GetName() string {
	return procedure.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (procedure *DatabaseAccountsSqlDatabasesContainersStoredProcedure_SpecARM) GetType() string {
	return ""
}

// Deprecated version of SqlStoredProcedureCreateUpdateProperties. Use v1beta20210515.SqlStoredProcedureCreateUpdateProperties instead
type SqlStoredProcedureCreateUpdatePropertiesARM struct {
	Options  *CreateUpdateOptionsARM        `json:"options,omitempty"`
	Resource *SqlStoredProcedureResourceARM `json:"resource,omitempty"`
}

// Deprecated version of SqlStoredProcedureResource. Use v1beta20210515.SqlStoredProcedureResource instead
type SqlStoredProcedureResourceARM struct {
	Body *string `json:"body,omitempty"`
	Id   *string `json:"id,omitempty"`
}
