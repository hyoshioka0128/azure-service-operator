// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisFirewallRules_SPECARM struct {
	AzureName string `json:"azureName"`
	Name      string `json:"name"`

	//Properties: redis cache firewall rule properties
	Properties RedisFirewallRuleProperties_SpecARM `json:"properties"`
}

var _ genruntime.ARMResourceSpec = &RedisFirewallRules_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (specarm RedisFirewallRules_SPECARM) GetAPIVersion() string {
	return "2020-12-01"
}

// GetName returns the Name of the resource
func (specarm RedisFirewallRules_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm RedisFirewallRules_SPECARM) GetType() string {
	return ""
}

type RedisFirewallRuleProperties_SpecARM struct {
	//EndIP: highest IP address included in the range
	EndIP string `json:"endIP"`

	//StartIP: lowest IP address included in the range
	StartIP string `json:"startIP"`
}
