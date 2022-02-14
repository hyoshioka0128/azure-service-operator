// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersFirewallRules_SPECARM struct {
	AzureName string `json:"azureName"`
	Name      string `json:"name"`

	//Properties: The properties of a firewall rule.
	Properties FirewallRuleProperties_SpecARM `json:"properties"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersFirewallRules_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (specarm FlexibleServersFirewallRules_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm FlexibleServersFirewallRules_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/firewallRules"
func (specarm FlexibleServersFirewallRules_SPECARM) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/firewallRules"
}

type FirewallRuleProperties_SpecARM struct {
	//EndIpAddress: The end IP address of the server firewall rule. Must be IPv4
	//format.
	EndIpAddress string `json:"endIpAddress"`

	//StartIpAddress: The start IP address of the server firewall rule. Must be IPv4
	//format.
	StartIpAddress string `json:"startIpAddress"`
}
