// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

type NamespacesEventhubsAuthorizationRule_STATUSARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties supplied to create or update AuthorizationRule
	Properties *NamespacesEventhubsAuthorizationRule_Properties_STATUSARM `json:"properties,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_STATUSARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string `json:"type,omitempty"`
}

type NamespacesEventhubsAuthorizationRule_Properties_STATUSARM struct {
	// Rights: The rights associated with the rule.
	Rights []NamespacesEventhubsAuthorizationRule_Properties_Rights_STATUS `json:"rights,omitempty"`
}
