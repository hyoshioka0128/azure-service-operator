// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccounts_SPECARM struct {
	AzureName string                          `json:"azureName"`
	Identity  *ManagedServiceIdentity_SpecARM `json:"identity,omitempty"`

	//Kind: Indicates the type of database account. This can only be set at database
	//account creation.
	Kind *DatabaseAccountsSPECKind `json:"kind,omitempty"`

	//Location: The location of the resource group to which the resource belongs.
	Location   *string                                       `json:"location,omitempty"`
	Name       string                                        `json:"name"`
	Properties DatabaseAccountCreateUpdateProperties_SpecARM `json:"properties"`
	Tags       map[string]string                             `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccounts_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (specarm DatabaseAccounts_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm DatabaseAccounts_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm DatabaseAccounts_SPECARM) GetType() string {
	return ""
}

type DatabaseAccountCreateUpdateProperties_SpecARM struct {
	//AnalyticalStorageConfiguration: Analytical storage specific properties.
	AnalyticalStorageConfiguration *AnalyticalStorageConfiguration_SpecARM `json:"analyticalStorageConfiguration,omitempty"`

	//ApiProperties: API specific properties. Currently, supported only for MongoDB
	//API.
	ApiProperties *ApiProperties_SpecARM `json:"apiProperties,omitempty"`

	//BackupPolicy: The object representing the policy for taking backups on an
	//account.
	BackupPolicy *BackupPolicy_SpecARM `json:"backupPolicy,omitempty"`

	//Capabilities: List of Cosmos DB capabilities for the account
	Capabilities []Capability_SpecARM `json:"capabilities,omitempty"`

	//ConnectorOffer: The cassandra connector offer type for the Cosmos DB database C*
	//account.
	ConnectorOffer *ConnectorOffer_Spec `json:"connectorOffer,omitempty"`

	//ConsistencyPolicy: The consistency policy for the Cosmos DB account.
	ConsistencyPolicy *ConsistencyPolicy_SpecARM `json:"consistencyPolicy,omitempty"`

	//Cors: The CORS policy for the Cosmos DB database account.
	Cors []CorsPolicy_SpecARM `json:"cors,omitempty"`

	//DatabaseAccountOfferType: The offer type for the database
	DatabaseAccountOfferType DatabaseAccountOfferType_Spec `json:"databaseAccountOfferType"`

	//DefaultIdentity: The default identity for accessing key vault used in features
	//like customer managed keys. The default identity needs to be explicitly set by
	//the users. It can be "FirstPartyIdentity", "SystemAssignedIdentity" and more.
	DefaultIdentity *string `json:"defaultIdentity,omitempty"`

	//DisableKeyBasedMetadataWriteAccess: Disable write operations on metadata
	//resources (databases, containers, throughput) via account keys
	DisableKeyBasedMetadataWriteAccess *bool `json:"disableKeyBasedMetadataWriteAccess,omitempty"`

	//EnableAnalyticalStorage: Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage *bool `json:"enableAnalyticalStorage,omitempty"`

	//EnableAutomaticFailover: Enables automatic failover of the write region in the
	//rare event that the region is unavailable due to an outage. Automatic failover
	//will result in a new write region for the account and is chosen based on the
	//failover priorities configured for the account.
	EnableAutomaticFailover *bool `json:"enableAutomaticFailover,omitempty"`

	//EnableCassandraConnector: Enables the cassandra connector on the Cosmos DB C*
	//account
	EnableCassandraConnector *bool `json:"enableCassandraConnector,omitempty"`

	//EnableFreeTier: Flag to indicate whether Free Tier is enabled.
	EnableFreeTier *bool `json:"enableFreeTier,omitempty"`

	//EnableMultipleWriteLocations: Enables the account to write in multiple locations
	EnableMultipleWriteLocations *bool `json:"enableMultipleWriteLocations,omitempty"`

	//IpRules: List of IpRules.
	IpRules []IpAddressOrRange_SpecARM `json:"ipRules,omitempty"`

	//IsVirtualNetworkFilterEnabled: Flag to indicate whether to enable/disable
	//Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled *bool `json:"isVirtualNetworkFilterEnabled,omitempty"`

	//KeyVaultKeyUri: The URI of the key vault
	KeyVaultKeyUri *string `json:"keyVaultKeyUri,omitempty"`

	//Locations: An array that contains the georeplication locations enabled for the
	//Cosmos DB account.
	Locations []Location_SpecARM `json:"locations"`

	//NetworkAclBypass: Indicates what services are allowed to bypass firewall checks.
	NetworkAclBypass *NetworkAclBypass_Spec `json:"networkAclBypass,omitempty"`

	//NetworkAclBypassResourceIds: An array that contains the Resource Ids for Network
	//Acl Bypass for the Cosmos DB account.
	NetworkAclBypassResourceIds []string `json:"networkAclBypassResourceIds,omitempty"`

	//PublicNetworkAccess: Whether requests from Public Network are allowed
	PublicNetworkAccess *PublicNetworkAccess_Spec `json:"publicNetworkAccess,omitempty"`

	//VirtualNetworkRules: List of Virtual Network ACL rules configured for the Cosmos
	//DB account.
	VirtualNetworkRules []VirtualNetworkRule_SpecARM `json:"virtualNetworkRules,omitempty"`
}

// +kubebuilder:validation:Enum={"GlobalDocumentDB","MongoDB","Parse"}
type DatabaseAccountsSPECKind string

const (
	DatabaseAccountsSPECKindGlobalDocumentDB = DatabaseAccountsSPECKind("GlobalDocumentDB")
	DatabaseAccountsSPECKindMongoDB          = DatabaseAccountsSPECKind("MongoDB")
	DatabaseAccountsSPECKindParse            = DatabaseAccountsSPECKind("Parse")
)

type ManagedServiceIdentity_SpecARM struct {
	//Type: The type of identity used for the resource. The type
	//'SystemAssigned,UserAssigned' includes both an implicitly created identity and a
	//set of user assigned identities. The type 'None' will remove any identities from
	//the service.
	Type *ManagedServiceIdentitySpecType `json:"type,omitempty"`
}

type AnalyticalStorageConfiguration_SpecARM struct {
	SchemaType *AnalyticalStorageSchemaType_Spec `json:"schemaType,omitempty"`
}

type ApiProperties_SpecARM struct {
	//ServerVersion: Describes the ServerVersion of an a MongoDB account.
	ServerVersion *ApiPropertiesSpecServerVersion `json:"serverVersion,omitempty"`
}

type BackupPolicy_SpecARM struct {
	Type BackupPolicyType_Spec `json:"type"`
}

type Capability_SpecARM struct {
	//Name: Name of the Cosmos DB capability. For example, "name": "EnableCassandra".
	//Current values also include "EnableTable" and "EnableGremlin".
	Name *string `json:"name,omitempty"`
}

type ConsistencyPolicy_SpecARM struct {
	//DefaultConsistencyLevel: The default consistency level and configuration
	//settings of the Cosmos DB account.
	DefaultConsistencyLevel ConsistencyPolicySpecDefaultConsistencyLevel `json:"defaultConsistencyLevel"`

	//MaxIntervalInSeconds: When used with the Bounded Staleness consistency level,
	//this value represents the time amount of staleness (in seconds) tolerated.
	//Accepted range for this value is 5 - 86400. Required when
	//defaultConsistencyPolicy is set to 'BoundedStaleness'.
	MaxIntervalInSeconds *int `json:"maxIntervalInSeconds,omitempty"`

	//MaxStalenessPrefix: When used with the Bounded Staleness consistency level, this
	//value represents the number of stale requests tolerated. Accepted range for this
	//value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set to
	//'BoundedStaleness'.
	MaxStalenessPrefix *int `json:"maxStalenessPrefix,omitempty"`
}

type CorsPolicy_SpecARM struct {
	//AllowedHeaders: The request headers that the origin domain may specify on the
	//CORS request.
	AllowedHeaders *string `json:"allowedHeaders,omitempty"`

	//AllowedMethods: The methods (HTTP request verbs) that the origin domain may use
	//for a CORS request.
	AllowedMethods *string `json:"allowedMethods,omitempty"`

	//AllowedOrigins: The origin domains that are permitted to make a request against
	//the service via CORS.
	AllowedOrigins string `json:"allowedOrigins"`

	//ExposedHeaders: The response headers that may be sent in the response to the
	//CORS request and exposed by the browser to the request issuer.
	ExposedHeaders *string `json:"exposedHeaders,omitempty"`

	//MaxAgeInSeconds: The maximum amount time that a browser should cache the
	//preflight OPTIONS request.
	MaxAgeInSeconds *int `json:"maxAgeInSeconds,omitempty"`
}

type IpAddressOrRange_SpecARM struct {
	//IpAddressOrRange: A single IPv4 address or a single IPv4 address range in CIDR
	//format. Provided IPs must be well-formatted and cannot be contained in one of
	//the following ranges: 10.0.0.0/8, 100.64.0.0/10, 172.16.0.0/12, 192.168.0.0/16,
	//since these are not enforceable by the IP address filter. Example of valid
	//inputs: “23.40.210.245” or “23.40.210.0/8”.
	IpAddressOrRange *string `json:"ipAddressOrRange,omitempty"`
}

type Location_SpecARM struct {
	//FailoverPriority: The failover priority of the region. A failover priority of 0
	//indicates a write region. The maximum value for a failover priority = (total
	//number of regions - 1). Failover priority values must be unique for each of the
	//regions in which the database account exists.
	FailoverPriority *int `json:"failoverPriority,omitempty"`

	//IsZoneRedundant: Flag to indicate whether or not this region is an
	//AvailabilityZone region
	IsZoneRedundant *bool `json:"isZoneRedundant,omitempty"`

	//LocationName: The name of the region.
	LocationName      *string `json:"locationName,omitempty"`
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type ManagedServiceIdentitySpecType string

const (
	ManagedServiceIdentitySpecTypeNone                       = ManagedServiceIdentitySpecType("None")
	ManagedServiceIdentitySpecTypeSystemAssigned             = ManagedServiceIdentitySpecType("SystemAssigned")
	ManagedServiceIdentitySpecTypeSystemAssignedUserAssigned = ManagedServiceIdentitySpecType("SystemAssigned,UserAssigned")
	ManagedServiceIdentitySpecTypeUserAssigned               = ManagedServiceIdentitySpecType("UserAssigned")
)

type VirtualNetworkRule_SpecARM struct {
	Id *string `json:"id,omitempty"`

	//IgnoreMissingVNetServiceEndpoint: Create firewall rule before the virtual
	//network has vnet service endpoint enabled.
	IgnoreMissingVNetServiceEndpoint *bool `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
}