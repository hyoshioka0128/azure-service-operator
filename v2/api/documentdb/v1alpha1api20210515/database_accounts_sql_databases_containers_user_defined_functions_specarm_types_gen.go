// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM struct {
	AzureName string `json:"azureName"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name"`

	//Properties: Properties to create and update Azure Cosmos DB userDefinedFunction.
	Properties SqlUserDefinedFunctionProperties_SpecARM `json:"properties"`
	Tags       map[string]string                        `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (specarm DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/userDefinedFunctions"
func (specarm DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/userDefinedFunctions"
}

type SqlUserDefinedFunctionProperties_SpecARM struct {
	//Options: A key-value pair of options to be applied for the request. This
	//corresponds to the headers sent with the request.
	Options *CreateUpdateOptions_SpecARM `json:"options,omitempty"`

	//Resource: The standard JSON format of a userDefinedFunction
	Resource SqlUserDefinedFunctionResource_SpecARM `json:"resource"`
}

type SqlUserDefinedFunctionResource_SpecARM struct {
	//Body: Body of the User Defined Function
	Body *string `json:"body,omitempty"`

	//Id: Name of the Cosmos DB SQL userDefinedFunction
	Id string `json:"id"`
}
