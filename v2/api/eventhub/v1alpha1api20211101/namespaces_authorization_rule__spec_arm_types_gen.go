// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of NamespacesAuthorizationRule_Spec. Use v1beta20211101.NamespacesAuthorizationRule_Spec instead
type NamespacesAuthorizationRule_SpecARM struct {
	AzureName  string                                          `json:"azureName,omitempty"`
	Id         *string                                         `json:"id,omitempty"`
	Location   *string                                         `json:"location,omitempty"`
	Name       string                                          `json:"name,omitempty"`
	Properties *NamespacesAuthorizationRule_Spec_PropertiesARM `json:"properties,omitempty"`
	SystemData *SystemDataARM                                  `json:"systemData,omitempty"`
	Type       *string                                         `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesAuthorizationRule_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule NamespacesAuthorizationRule_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (rule *NamespacesAuthorizationRule_SpecARM) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (rule *NamespacesAuthorizationRule_SpecARM) GetType() string {
	return ""
}

// Deprecated version of NamespacesAuthorizationRule_Spec_Properties. Use v1beta20211101.NamespacesAuthorizationRule_Spec_Properties instead
type NamespacesAuthorizationRule_Spec_PropertiesARM struct {
	Rights []NamespacesAuthorizationRule_Spec_Properties_Rights `json:"rights,omitempty"`
}
