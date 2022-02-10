// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Domains_SPECARM struct {
	AzureName string `json:"azureName"`

	//Location: Location of the resource.
	Location string `json:"location"`
	Name     string `json:"name"`

	//Properties: Properties of the domain.
	Properties *DomainProperties_SpecARM `json:"properties,omitempty"`

	//Tags: Tags of the resource.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Domains_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (specarm Domains_SPECARM) GetAPIVersion() string {
	return "2020-06-01"
}

// GetName returns the Name of the resource
func (specarm Domains_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm Domains_SPECARM) GetType() string {
	return ""
}

type DomainProperties_SpecARM struct {
	//InboundIpRules: This can be used to restrict traffic from specific IPs instead
	//of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRule_SpecARM `json:"inboundIpRules,omitempty"`

	//InputSchema: This determines the format that Event Grid should expect for
	//incoming events published to the domain.
	InputSchema *DomainPropertiesSpecInputSchema `json:"inputSchema,omitempty"`

	//InputSchemaMapping: Information about the InputSchemaMapping which specified the
	//info about mapping event payload.
	InputSchemaMapping *InputSchemaMapping_SpecARM `json:"inputSchemaMapping,omitempty"`

	//PublicNetworkAccess: This determines if traffic is allowed over public network.
	//By default it is enabled.
	//You can further restrict to specific IPs by configuring <seealso
	//cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules"
	///>
	PublicNetworkAccess *DomainPropertiesSpecPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
}

type InboundIpRule_SpecARM struct {
	//Action: Action to perform based on the match or no match of the IpMask.
	Action *InboundIpRuleSpecAction `json:"action,omitempty"`

	//IpMask: IP Address in CIDR notation e.g., 10.0.0.0/8.
	IpMask *string `json:"ipMask,omitempty"`
}

type InputSchemaMapping_SpecARM struct {
	//InputSchemaMappingType: Type of the custom mapping
	InputSchemaMappingType InputSchemaMappingSpecInputSchemaMappingType `json:"inputSchemaMappingType"`
}
