// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Redis_SPECARM struct {
	AzureName string `json:"azureName"`

	//Location: The geo-location where the resource lives
	Location string `json:"location"`
	Name     string `json:"name"`

	//Properties: Redis cache properties.
	Properties RedisProperties_SpecARM `json:"properties"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	//Zones: A list of availability zones denoting where the resource needs to come
	//from.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Redis_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (specarm Redis_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm Redis_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis"
func (specarm Redis_SPECARM) GetType() string {
	return "Microsoft.Cache/redis"
}

type RedisProperties_SpecARM struct {
	//EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is
	//enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`

	//MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or
	//higher) to connect (e,g, '1.0', '1.1', '1.2')
	MinimumTlsVersion *RedisProperties_MinimumTlsVersion_Spec `json:"minimumTlsVersion,omitempty"`

	//PublicNetworkAccess: Whether or not public endpoint access is allowed for this
	//cache.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'. If
	//'Disabled', private endpoints are the exclusive access method. Default value is
	//'Enabled'
	PublicNetworkAccess *RedisProperties_PublicNetworkAccess_Spec `json:"publicNetworkAccess,omitempty"`

	//RedisConfiguration: All Redis Settings. Few possible keys:
	//rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
	//etc.
	RedisConfiguration *RedisProperties_RedisConfiguration_SpecARM `json:"redisConfiguration,omitempty"`

	//RedisVersion: Redis version. Only major version will be used in PUT/PATCH
	//request with current valid values: (4, 6)
	RedisVersion *string `json:"redisVersion,omitempty"`

	//ReplicasPerMaster: The number of replicas to be created per primary.
	ReplicasPerMaster *int `json:"replicasPerMaster,omitempty"`

	//ReplicasPerPrimary: The number of replicas to be created per primary.
	ReplicasPerPrimary *int `json:"replicasPerPrimary,omitempty"`

	//ShardCount: The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int `json:"shardCount,omitempty"`

	//Sku: The SKU of the Redis cache to deploy.
	Sku Sku_SpecARM `json:"sku"`

	//StaticIP: Static IP address. Optionally, may be specified when deploying a Redis
	//cache inside an existing Azure Virtual Network; auto assigned by default.
	StaticIP *string `json:"staticIP,omitempty"`

	//SubnetId: The full resource ID of a subnet in a virtual network to deploy the
	//Redis cache in. Example format:
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetId *string `json:"subnetId,omitempty"`

	//TenantSettings: A dictionary of tenant settings
	TenantSettings map[string]string `json:"tenantSettings,omitempty"`
}

type RedisProperties_RedisConfiguration_SpecARM struct {
	AdditionalProperties map[string]string `json:"additionalProperties"`

	//AofStorageConnectionString0: First storage account connection string
	AofStorageConnectionString0 *string `json:"aof-storage-connection-string-0,omitempty"`

	//AofStorageConnectionString1: Second storage account connection string
	AofStorageConnectionString1 *string `json:"aof-storage-connection-string-1,omitempty"`

	//MaxfragmentationmemoryReserved: Value in megabytes reserved for fragmentation
	//per shard
	MaxfragmentationmemoryReserved *string `json:"maxfragmentationmemory-reserved,omitempty"`

	//MaxmemoryDelta: Value in megabytes reserved for non-cache usage per shard e.g.
	//failover.
	MaxmemoryDelta *string `json:"maxmemory-delta,omitempty"`

	//MaxmemoryPolicy: The eviction strategy used when your data won't fit within its
	//memory limit.
	MaxmemoryPolicy *string `json:"maxmemory-policy,omitempty"`

	//MaxmemoryReserved: Value in megabytes reserved for non-cache usage per shard
	//e.g. failover.
	MaxmemoryReserved *string `json:"maxmemory-reserved,omitempty"`

	//RdbBackupEnabled: Specifies whether the rdb backup is enabled
	RdbBackupEnabled *string `json:"rdb-backup-enabled,omitempty"`

	//RdbBackupFrequency: Specifies the frequency for creating rdb backup
	RdbBackupFrequency *string `json:"rdb-backup-frequency,omitempty"`

	//RdbBackupMaxSnapshotCount: Specifies the maximum number of snapshots for rdb
	//backup
	RdbBackupMaxSnapshotCount *string `json:"rdb-backup-max-snapshot-count,omitempty"`

	//RdbStorageConnectionString: The storage account connection string for storing
	//rdb file
	RdbStorageConnectionString *string `json:"rdb-storage-connection-string,omitempty"`
}

type Sku_SpecARM struct {
	//Capacity: The size of the Redis cache to deploy. Valid values: for C
	//(Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3,
	//4).
	Capacity int `json:"capacity"`

	//Family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P =
	//Premium).
	Family Sku_Family_Spec `json:"family"`

	//Name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
	Name Sku_Name_Spec `json:"name"`
}
