// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:rbac:groups=cache.azure.com,resources=redis,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.azure.com,resources={redis/status,redis/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.Cache.json#/resourceDefinitions/redis
type Redis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Redis_Spec           `json:"spec,omitempty"`
	Status            RedisResource_Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.Cache.json#/resourceDefinitions/redis
type RedisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Redis `json:"items"`
}

type RedisResource_Status struct {
	v1alpha1.ResourceStatus `json:",inline,omitempty"`
	AtProvider              RedisResourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:validation:Enum={"2020-06-01"}
type RedisSpecAPIVersion string

const RedisSpecAPIVersion20200601 = RedisSpecAPIVersion("2020-06-01")

type Redis_Spec struct {
	v1alpha1.ResourceSpec `json:",inline,omitempty"`
	ForProvider           RedisParameters `json:"forProvider,omitempty"`
}

type RedisParameters struct {
	// EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1',
	// '1.2').
	MinimumTlsVersion *RedisCreatePropertiesMinimumTlsVersion `json:"minimumTlsVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Name: The name of the Redis cache.
	Name string `json:"name,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method. Default value is
	// 'Enabled'.
	PublicNetworkAccess *RedisCreatePropertiesPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// RedisConfiguration: All Redis Settings. Few possible keys:
	// rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
	// etc.
	RedisConfiguration *RedisCommonPropertiesRedisConfiguration `json:"redisConfiguration,omitempty"`

	// ReplicasPerMaster: The number of replicas to be created per master.
	ReplicasPerMaster         *int                `json:"replicasPerMaster,omitempty"`
	ResourceGroupName         string              `json:"resourceGroupName,omitempty"`
	ResourceGroupNameRef      *v1alpha1.Reference `json:"resourceGroupNameRef,omitempty"`
	ResourceGroupNameSelector *v1alpha1.Selector  `json:"resourceGroupNameSelector,omitempty"`

	// ShardCount: The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int `json:"shardCount,omitempty"`

	// +kubebuilder:validation:Required
	// Sku: SKU parameters supplied to the create Redis operation.
	Sku *Sku `json:"sku,omitempty"`

	// +kubebuilder:validation:Pattern="^\\d+\\.\\d+\\.\\d+\\.\\d+$"
	// StaticIP: Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual
	// Network; auto assigned by default.
	StaticIP *string `json:"staticIP,omitempty"`

	// +kubebuilder:validation:Pattern="^/subscriptions/[^/]*/resourceGroups/[^/]*/providers/Microsoft.(ClassicNetwork|Network)/virtualNetworks/[^/]*/subnets/[^/]*$"
	// SubnetId: The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetId *string `json:"subnetId,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	// TenantSettings: A dictionary of tenant settings
	TenantSettings map[string]string `json:"tenantSettings,omitempty"`

	// Zones: A list of availability zones denoting where the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type RedisResourceObservation struct {
	// AccessKeys: The keys of the Redis cache - not set if this object is not the response to Create or Update redis cache
	AccessKeys *RedisAccessKeys_Status `json:"accessKeys,omitempty"`

	// EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`

	// HostName: Redis host name.
	HostName *string `json:"hostName,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Instances: List of the Redis instances associated with the cache
	Instances []RedisInstanceDetails_Status `json:"instances,omitempty"`

	// LinkedServers: List of the linked servers associated with the cache
	LinkedServers []RedisLinkedServer_Status `json:"linkedServers,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1',
	// '1.2')
	MinimumTlsVersion *RedisPropertiesStatusMinimumTlsVersion `json:"minimumTlsVersion,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Port: Redis non-SSL port.
	Port *int `json:"port,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connection associated with the specified redis cache
	PrivateEndpointConnections []PrivateEndpointConnection_Status `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: Redis instance provisioning status.
	ProvisioningState *RedisPropertiesStatusProvisioningState `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method. Default value is
	// 'Enabled'
	PublicNetworkAccess *RedisPropertiesStatusPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// RedisConfiguration: All Redis Settings. Few possible keys:
	// rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
	// etc.
	RedisConfiguration *RedisProperties_Status_RedisConfiguration `json:"redisConfiguration,omitempty"`

	// RedisVersion: Redis version.
	RedisVersion *string `json:"redisVersion,omitempty"`

	// ReplicasPerMaster: The number of replicas to be created per master.
	ReplicasPerMaster *int `json:"replicasPerMaster,omitempty"`

	// ShardCount: The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int `json:"shardCount,omitempty"`

	// Sku: The SKU of the Redis cache to deploy.
	Sku *Sku_Status `json:"sku,omitempty"`

	// SslPort: Redis SSL port.
	SslPort *int `json:"sslPort,omitempty"`

	// StaticIP: Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual
	// Network; auto assigned by default.
	StaticIP *string `json:"staticIP,omitempty"`

	// SubnetId: The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetId *string `json:"subnetId,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// TenantSettings: A dictionary of tenant settings
	TenantSettings map[string]string `json:"tenantSettings,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// Zones: A list of availability zones denoting where the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type PrivateEndpointConnection_Status struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// PrivateEndpoint: The resource of private end point.
	PrivateEndpoint *PrivateEndpoint_Status `json:"privateEndpoint,omitempty"`

	// PrivateLinkServiceConnectionState: A collection of information about the state of the connection between service
	// consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState_Status `json:"privateLinkServiceConnectionState,omitempty"`

	// ProvisioningState: The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState_Status `json:"provisioningState,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type RedisAccessKeys_Status struct {
	// PrimaryKey: The current primary key that clients can use to authenticate with Redis cache.
	PrimaryKey *string `json:"primaryKey,omitempty"`

	// SecondaryKey: The current secondary key that clients can use to authenticate with Redis cache.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.Cache.json#/definitions/RedisCommonPropertiesRedisConfiguration
type RedisCommonPropertiesRedisConfiguration struct {
	// AdditionalProperties: Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty"`

	// AofStorageConnectionString0: First storage account connection string
	AofStorageConnectionString0 *string `json:"aof-storage-connection-string-0,omitempty"`

	// AofStorageConnectionString1: First storage account connection string
	AofStorageConnectionString1 *string `json:"aof-storage-connection-string-1,omitempty"`

	// MaxfragmentationmemoryReserved: Value in megabytes reserved for fragmentation per shard
	MaxfragmentationmemoryReserved *string `json:"maxfragmentationmemory-reserved,omitempty"`

	// MaxmemoryDelta: Value in megabytes reserved for non-cache usage per shard e.g. failover.
	MaxmemoryDelta *string `json:"maxmemory-delta,omitempty"`

	// MaxmemoryPolicy: The eviction strategy used when your data won't fit within its memory limit.
	MaxmemoryPolicy *string `json:"maxmemory-policy,omitempty"`

	// MaxmemoryReserved: Value in megabytes reserved for non-cache usage per shard e.g. failover.
	MaxmemoryReserved *string `json:"maxmemory-reserved,omitempty"`

	// RdbBackupEnabled: Specifies whether the rdb backup is enabled
	RdbBackupEnabled *string `json:"rdb-backup-enabled,omitempty"`

	// RdbBackupFrequency: Specifies the frequency for creating rdb backup
	RdbBackupFrequency *string `json:"rdb-backup-frequency,omitempty"`

	// RdbBackupMaxSnapshotCount: Specifies the maximum number of snapshots for rdb backup
	RdbBackupMaxSnapshotCount *string `json:"rdb-backup-max-snapshot-count,omitempty"`

	// RdbStorageConnectionString: The storage account connection string for storing rdb file
	RdbStorageConnectionString *string `json:"rdb-storage-connection-string,omitempty"`
}

// +kubebuilder:validation:Enum={"1.0","1.1","1.2"}
type RedisCreatePropertiesMinimumTlsVersion string

const (
	RedisCreatePropertiesMinimumTlsVersion10 = RedisCreatePropertiesMinimumTlsVersion("1.0")
	RedisCreatePropertiesMinimumTlsVersion11 = RedisCreatePropertiesMinimumTlsVersion("1.1")
	RedisCreatePropertiesMinimumTlsVersion12 = RedisCreatePropertiesMinimumTlsVersion("1.2")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type RedisCreatePropertiesPublicNetworkAccess string

const (
	RedisCreatePropertiesPublicNetworkAccessDisabled = RedisCreatePropertiesPublicNetworkAccess("Disabled")
	RedisCreatePropertiesPublicNetworkAccessEnabled  = RedisCreatePropertiesPublicNetworkAccess("Enabled")
)

type RedisInstanceDetails_Status struct {
	// IsMaster: Specifies whether the instance is a master node.
	IsMaster *bool `json:"isMaster,omitempty"`

	// NonSslPort: If enableNonSslPort is true, provides Redis instance Non-SSL port.
	NonSslPort *int `json:"nonSslPort,omitempty"`

	// ShardId: If clustering is enabled, the Shard ID of Redis Instance
	ShardId *int `json:"shardId,omitempty"`

	// SslPort: Redis instance SSL port.
	SslPort *int `json:"sslPort,omitempty"`

	// Zone: If the Cache uses availability zones, specifies availability zone where this instance is located.
	Zone *string `json:"zone,omitempty"`
}

type RedisLinkedServer_Status struct {
	// Id: Linked server Id.
	Id *string `json:"id,omitempty"`
}

type RedisPropertiesStatusMinimumTlsVersion string

const (
	RedisPropertiesStatusMinimumTlsVersion10 = RedisPropertiesStatusMinimumTlsVersion("1.0")
	RedisPropertiesStatusMinimumTlsVersion11 = RedisPropertiesStatusMinimumTlsVersion("1.1")
	RedisPropertiesStatusMinimumTlsVersion12 = RedisPropertiesStatusMinimumTlsVersion("1.2")
)

type RedisPropertiesStatusProvisioningState string

const (
	RedisPropertiesStatusProvisioningStateCreating               = RedisPropertiesStatusProvisioningState("Creating")
	RedisPropertiesStatusProvisioningStateDeleting               = RedisPropertiesStatusProvisioningState("Deleting")
	RedisPropertiesStatusProvisioningStateDisabled               = RedisPropertiesStatusProvisioningState("Disabled")
	RedisPropertiesStatusProvisioningStateFailed                 = RedisPropertiesStatusProvisioningState("Failed")
	RedisPropertiesStatusProvisioningStateLinking                = RedisPropertiesStatusProvisioningState("Linking")
	RedisPropertiesStatusProvisioningStateProvisioning           = RedisPropertiesStatusProvisioningState("Provisioning")
	RedisPropertiesStatusProvisioningStateRecoveringScaleFailure = RedisPropertiesStatusProvisioningState("RecoveringScaleFailure")
	RedisPropertiesStatusProvisioningStateScaling                = RedisPropertiesStatusProvisioningState("Scaling")
	RedisPropertiesStatusProvisioningStateSucceeded              = RedisPropertiesStatusProvisioningState("Succeeded")
	RedisPropertiesStatusProvisioningStateUnlinking              = RedisPropertiesStatusProvisioningState("Unlinking")
	RedisPropertiesStatusProvisioningStateUnprovisioning         = RedisPropertiesStatusProvisioningState("Unprovisioning")
	RedisPropertiesStatusProvisioningStateUpdating               = RedisPropertiesStatusProvisioningState("Updating")
)

type RedisPropertiesStatusPublicNetworkAccess string

const (
	RedisPropertiesStatusPublicNetworkAccessDisabled = RedisPropertiesStatusPublicNetworkAccess("Disabled")
	RedisPropertiesStatusPublicNetworkAccessEnabled  = RedisPropertiesStatusPublicNetworkAccess("Enabled")
)

type RedisProperties_Status_RedisConfiguration struct {
	// AofStorageConnectionString0: First storage account connection string
	AofStorageConnectionString0 *string `json:"aof-storage-connection-string-0,omitempty"`

	// AofStorageConnectionString1: Second storage account connection string
	AofStorageConnectionString1 *string `json:"aof-storage-connection-string-1,omitempty"`

	// Maxclients: The max clients config
	Maxclients *string `json:"maxclients,omitempty"`

	// MaxfragmentationmemoryReserved: Value in megabytes reserved for fragmentation per shard
	MaxfragmentationmemoryReserved *string `json:"maxfragmentationmemory-reserved,omitempty"`

	// MaxmemoryDelta: Value in megabytes reserved for non-cache usage per shard e.g. failover.
	MaxmemoryDelta *string `json:"maxmemory-delta,omitempty"`

	// MaxmemoryPolicy: The eviction strategy used when your data won't fit within its memory limit.
	MaxmemoryPolicy *string `json:"maxmemory-policy,omitempty"`

	// MaxmemoryReserved: Value in megabytes reserved for non-cache usage per shard e.g. failover.
	MaxmemoryReserved *string `json:"maxmemory-reserved,omitempty"`

	// RdbBackupEnabled: Specifies whether the rdb backup is enabled
	RdbBackupEnabled *string `json:"rdb-backup-enabled,omitempty"`

	// RdbBackupFrequency: Specifies the frequency for creating rdb backup
	RdbBackupFrequency *string `json:"rdb-backup-frequency,omitempty"`

	// RdbBackupMaxSnapshotCount: Specifies the maximum number of snapshots for rdb backup
	RdbBackupMaxSnapshotCount *string `json:"rdb-backup-max-snapshot-count,omitempty"`

	// RdbStorageConnectionString: The storage account connection string for storing rdb file
	RdbStorageConnectionString *string           `json:"rdb-storage-connection-string,omitempty"`
	additionalProperties       map[string]string `json:"additionalProperties,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.Cache.json#/definitions/Sku
type Sku struct {
	// +kubebuilder:validation:Required
	// Capacity: The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for
	// P (Premium) family (1, 2, 3, 4).
	Capacity *int `json:"capacity,omitempty"`

	// +kubebuilder:validation:Required
	// Family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
	Family *SkuFamily `json:"family,omitempty"`

	// +kubebuilder:validation:Required
	// Name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium).
	Name *SkuName `json:"name,omitempty"`
}

type Sku_Status struct {
	// Capacity: The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for
	// P (Premium) family (1, 2, 3, 4).
	Capacity *int `json:"capacity,omitempty"`

	// Family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
	Family *SkuStatusFamily `json:"family,omitempty"`

	// Name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
	Name *SkuStatusName `json:"name,omitempty"`
}

type PrivateEndpointConnectionProvisioningState_Status string

const (
	PrivateEndpointConnectionProvisioningState_StatusCreating  = PrivateEndpointConnectionProvisioningState_Status("Creating")
	PrivateEndpointConnectionProvisioningState_StatusDeleting  = PrivateEndpointConnectionProvisioningState_Status("Deleting")
	PrivateEndpointConnectionProvisioningState_StatusFailed    = PrivateEndpointConnectionProvisioningState_Status("Failed")
	PrivateEndpointConnectionProvisioningState_StatusSucceeded = PrivateEndpointConnectionProvisioningState_Status("Succeeded")
)

type PrivateEndpoint_Status struct {
	// Id: The ARM identifier for Private Endpoint
	Id *string `json:"id,omitempty"`
}

type PrivateLinkServiceConnectionState_Status struct {
	// ActionsRequired: A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string `json:"actionsRequired,omitempty"`

	// Description: The reason for approval/rejection of the connection.
	Description *string `json:"description,omitempty"`

	// Status: Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus_Status `json:"status,omitempty"`
}

// +kubebuilder:validation:Enum={"C","P"}
type SkuFamily string

const (
	SkuFamilyC = SkuFamily("C")
	SkuFamilyP = SkuFamily("P")
)

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type SkuName string

const (
	SkuNameBasic    = SkuName("Basic")
	SkuNamePremium  = SkuName("Premium")
	SkuNameStandard = SkuName("Standard")
)

type SkuStatusFamily string

const (
	SkuStatusFamilyC = SkuStatusFamily("C")
	SkuStatusFamilyP = SkuStatusFamily("P")
)

type SkuStatusName string

const (
	SkuStatusNameBasic    = SkuStatusName("Basic")
	SkuStatusNamePremium  = SkuStatusName("Premium")
	SkuStatusNameStandard = SkuStatusName("Standard")
)

type PrivateEndpointServiceConnectionStatus_Status string

const (
	PrivateEndpointServiceConnectionStatus_StatusApproved = PrivateEndpointServiceConnectionStatus_Status("Approved")
	PrivateEndpointServiceConnectionStatus_StatusPending  = PrivateEndpointServiceConnectionStatus_Status("Pending")
	PrivateEndpointServiceConnectionStatus_StatusRejected = PrivateEndpointServiceConnectionStatus_Status("Rejected")
)

func init() {
	SchemeBuilder.Register(&Redis{}, &RedisList{})
}
