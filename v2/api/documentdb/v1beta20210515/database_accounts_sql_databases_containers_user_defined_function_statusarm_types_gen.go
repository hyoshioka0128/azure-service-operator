// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

type DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSARM struct {
	// Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Name: The name of the ARM resource.
	Name *string `json:"name,omitempty"`

	// Properties: The properties of an Azure Cosmos DB userDefinedFunction
	Properties *SqlUserDefinedFunctionGetProperties_STATUSARM `json:"properties,omitempty"`
	Tags       map[string]string                              `json:"tags,omitempty"`

	// Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

type SqlUserDefinedFunctionGetProperties_STATUSARM struct {
	Resource *SqlUserDefinedFunctionGetProperties_Resource_STATUSARM `json:"resource,omitempty"`
}

type SqlUserDefinedFunctionGetProperties_Resource_STATUSARM struct {
	// Body: Body of the User Defined Function
	Body *string `json:"body,omitempty"`

	// Id: Name of the Cosmos DB SQL userDefinedFunction
	Id *string `json:"id,omitempty"`

	// _Etag: A system generated property representing the resource etag required for optimistic concurrency control.
	_Etag *string `json:"_etag,omitempty"`

	// _Rid: A system generated property. A unique identifier.
	_Rid *string `json:"_rid,omitempty"`

	// _Ts: A system generated property that denotes the last updated timestamp of the resource.
	_Ts *float64 `json:"_ts,omitempty"`
}
