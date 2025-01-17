// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RouteTables_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: Name of the resource
	Name string `json:"name,omitempty"`

	// Properties: Properties of the route table.
	Properties *RouteTables_Spec_PropertiesARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RouteTables_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (tables RouteTables_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (tables *RouteTables_SpecARM) GetName() string {
	return tables.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/routeTables"
func (tables *RouteTables_SpecARM) GetType() string {
	return "Microsoft.Network/routeTables"
}

type RouteTables_Spec_PropertiesARM struct {
	// DisableBgpRoutePropagation: Whether to disable the routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation *bool `json:"disableBgpRoutePropagation,omitempty"`

	// Routes: Collection of routes contained within a route table.
	Routes []RouteTables_Spec_Properties_RoutesARM `json:"routes,omitempty"`
}

type RouteTables_Spec_Properties_RoutesARM struct {
	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the route.
	Properties *RoutePropertiesFormatARM `json:"properties,omitempty"`
}
