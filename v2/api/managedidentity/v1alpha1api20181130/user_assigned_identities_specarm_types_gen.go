// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20181130

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type UserAssignedIdentities_SPECARM struct {
	AzureName string `json:"azureName"`

	//Location: The geo-location where the resource lives
	Location string `json:"location"`
	Name     string `json:"name"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &UserAssignedIdentities_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-11-30"
func (specarm UserAssignedIdentities_SPECARM) GetAPIVersion() string {
	return "2018-11-30"
}

// GetName returns the Name of the resource
func (specarm UserAssignedIdentities_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm UserAssignedIdentities_SPECARM) GetType() string {
	return ""
}
