// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NamespacesEventhubsAuthorizationRules_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: The authorization rule name.
	Name string `json:"name,omitempty"`

	// Properties: Properties supplied to create or update AuthorizationRule
	Properties *AuthorizationRulePropertiesARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesEventhubsAuthorizationRules_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rules NamespacesEventhubsAuthorizationRules_SpecARM) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (rules NamespacesEventhubsAuthorizationRules_SpecARM) GetName() string {
	return rules.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
func (rules NamespacesEventhubsAuthorizationRules_SpecARM) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
}
