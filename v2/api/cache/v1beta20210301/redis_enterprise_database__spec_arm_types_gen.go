// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210301

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisEnterpriseDatabase_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name string `json:"name,omitempty"`

	// Properties: Other properties of the database.
	Properties *DatabasePropertiesARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisEnterpriseDatabase_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-03-01"
func (database RedisEnterpriseDatabase_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (database *RedisEnterpriseDatabase_SpecARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (database *RedisEnterpriseDatabase_SpecARM) GetType() string {
	return ""
}

type DatabasePropertiesARM struct {
	// ClientProtocol: Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is
	// TLS-encrypted.
	ClientProtocol *DatabaseProperties_ClientProtocol `json:"clientProtocol,omitempty"`

	// ClusteringPolicy: Clustering policy - default is OSSCluster. Specified at create time.
	ClusteringPolicy *DatabaseProperties_ClusteringPolicy `json:"clusteringPolicy,omitempty"`

	// EvictionPolicy: Redis eviction policy - default is VolatileLRU
	EvictionPolicy *DatabaseProperties_EvictionPolicy `json:"evictionPolicy,omitempty"`

	// Modules: Optional set of redis modules to enable in this database - modules can only be added at creation time.
	Modules []ModuleARM `json:"modules,omitempty"`

	// Persistence: Persistence settings
	Persistence *PersistenceARM `json:"persistence,omitempty"`

	// Port: TCP port of the database endpoint. Specified at create time. Defaults to an available port.
	Port *int `json:"port,omitempty"`

	// ProvisioningState: Current provisioning status of the database
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`

	// ResourceState: Current resource status of the database
	ResourceState *ResourceState `json:"resourceState,omitempty"`
}

type ModuleARM struct {
	// Args: Configuration options for the module, e.g. 'ERROR_RATE 0.00 INITIAL_SIZE 400'.
	Args *string `json:"args,omitempty"`

	// Name: The name of the module, e.g. 'RedisBloom', 'RediSearch', 'RedisTimeSeries'
	Name *string `json:"name,omitempty"`

	// Version: The version of the module, e.g. '1.0'.
	Version *string `json:"version,omitempty"`
}

type PersistenceARM struct {
	// AofEnabled: Sets whether AOF is enabled.
	AofEnabled *bool `json:"aofEnabled,omitempty"`

	// AofFrequency: Sets the frequency at which data is written to disk.
	AofFrequency *Persistence_AofFrequency `json:"aofFrequency,omitempty"`

	// RdbEnabled: Sets whether RDB is enabled.
	RdbEnabled *bool `json:"rdbEnabled,omitempty"`

	// RdbFrequency: Sets the frequency at which a snapshot of the database is created.
	RdbFrequency *Persistence_RdbFrequency `json:"rdbFrequency,omitempty"`
}