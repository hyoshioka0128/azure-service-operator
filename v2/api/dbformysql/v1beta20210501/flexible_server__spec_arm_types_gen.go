// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type FlexibleServer_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Identity: The cmk identity for the server.
	Identity *IdentityARM `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name string `json:"name,omitempty"`

	// Properties: Properties of the server.
	Properties *ServerPropertiesARM `json:"properties,omitempty"`

	// Sku: The SKU (pricing tier) of the server.
	Sku *SkuARM `json:"sku,omitempty"`

	// SystemData: The system metadata relating to this resource.
	SystemData *SystemDataARM `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServer_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (server FlexibleServer_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (server *FlexibleServer_SpecARM) GetName() string {
	return server.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (server *FlexibleServer_SpecARM) GetType() string {
	return ""
}

type IdentityARM struct {
	// PrincipalId: ObjectId from the KeyVault
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: TenantId from the KeyVault
	TenantId *string `json:"tenantId,omitempty"`

	// Type: Type of managed service identity.
	Type *Identity_Type `json:"type,omitempty"`

	// UserAssignedIdentities: Metadata of user assigned identity.
	UserAssignedIdentities map[string]v1.JSON `json:"userAssignedIdentities,omitempty"`
}

type ServerPropertiesARM struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The password of the administrator login (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// AvailabilityZone: availability Zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Backup: Backup related properties of a server.
	Backup *BackupARM `json:"backup,omitempty"`

	// CreateMode: The mode to create a new MySQL server.
	CreateMode *ServerProperties_CreateMode `json:"createMode,omitempty"`

	// DataEncryption: The Data Encryption for CMK.
	DataEncryption *DataEncryptionARM `json:"dataEncryption,omitempty"`

	// FullyQualifiedDomainName: The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	// HighAvailability: High availability related properties of a server.
	HighAvailability *HighAvailabilityARM `json:"highAvailability,omitempty"`

	// MaintenanceWindow: Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindowARM `json:"maintenanceWindow,omitempty"`

	// Network: Network related properties of a server.
	Network *NetworkARM `json:"network,omitempty"`

	// ReplicaCapacity: The maximum number of replicas that a primary server can have.
	ReplicaCapacity *int `json:"replicaCapacity,omitempty"`

	// ReplicationRole: The replication role.
	ReplicationRole *ReplicationRole `json:"replicationRole,omitempty"`

	// RestorePointInTime: Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime *string `json:"restorePointInTime,omitempty"`

	// SourceServerResourceId: The source MySQL server id.
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	// State: The state of a server.
	State *ServerProperties_State `json:"state,omitempty"`

	// Storage: Storage related properties of a server.
	Storage *StorageARM `json:"storage,omitempty"`

	// Version: Server version.
	Version *ServerVersion `json:"version,omitempty"`
}

type SkuARM struct {
	// Name: The name of the sku, e.g. Standard_D32s_v3.
	Name *string `json:"name,omitempty"`

	// Tier: The tier of the particular SKU, e.g. GeneralPurpose.
	Tier *Sku_Tier `json:"tier,omitempty"`
}

type SystemDataARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType `json:"lastModifiedByType,omitempty"`
}

type BackupARM struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// EarliestRestoreDate: Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate *string `json:"earliestRestoreDate,omitempty"`

	// GeoRedundantBackup: Whether or not geo redundant backup is enabled.
	GeoRedundantBackup *EnableStatusEnum `json:"geoRedundantBackup,omitempty"`
}

type DataEncryptionARM struct {
	// GeoBackupKeyUri: Geo backup key uri as key vault can't cross region, need cmk in same region as geo backup
	GeoBackupKeyUri *string `json:"geoBackupKeyUri,omitempty"`

	// GeoBackupUserAssignedIdentityId: Geo backup user identity resource id as identity can't cross region, need identity in
	// same region as geo backup
	GeoBackupUserAssignedIdentityId *string `json:"geoBackupUserAssignedIdentityId,omitempty"`

	// PrimaryKeyUri: Primary key uri
	PrimaryKeyUri *string `json:"primaryKeyUri,omitempty"`

	// PrimaryUserAssignedIdentityId: Primary user identity resource id
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// Type: The key type, AzureKeyVault for enable cmk, SystemManaged for disable cmk.
	Type *DataEncryption_Type `json:"type,omitempty"`
}

type HighAvailabilityARM struct {
	// Mode: High availability mode for a server.
	Mode *HighAvailability_Mode `json:"mode,omitempty"`

	// StandbyAvailabilityZone: Availability zone of the standby server.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`

	// State: The state of server high availability.
	State *HighAvailability_State `json:"state,omitempty"`
}

// +kubebuilder:validation:Enum={"UserAssigned"}
type Identity_Type string

const Identity_Type_UserAssigned = Identity_Type("UserAssigned")

type MaintenanceWindowARM struct {
	// CustomWindow: indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	// DayOfWeek: day of week for maintenance window
	DayOfWeek *int `json:"dayOfWeek,omitempty"`

	// StartHour: start hour for maintenance window
	StartHour *int `json:"startHour,omitempty"`

	// StartMinute: start minute for maintenance window
	StartMinute *int `json:"startMinute,omitempty"`
}

type NetworkARM struct {
	DelegatedSubnetResourceId *string `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneResourceId  *string `json:"privateDnsZoneResourceId,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is 'Disabled' when server
	// has VNet integration.
	PublicNetworkAccess *EnableStatusEnum `json:"publicNetworkAccess,omitempty"`
}

// +kubebuilder:validation:Enum={"Burstable","GeneralPurpose","MemoryOptimized"}
type Sku_Tier string

const (
	Sku_Tier_Burstable       = Sku_Tier("Burstable")
	Sku_Tier_GeneralPurpose  = Sku_Tier("GeneralPurpose")
	Sku_Tier_MemoryOptimized = Sku_Tier("MemoryOptimized")
)

type StorageARM struct {
	// AutoGrow: Enable Storage Auto Grow or not.
	AutoGrow *EnableStatusEnum `json:"autoGrow,omitempty"`

	// Iops: Storage IOPS for a server.
	Iops *int `json:"iops,omitempty"`

	// StorageSizeGB: Max storage size allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`

	// StorageSku: The sku name of the server storage.
	StorageSku *string `json:"storageSku,omitempty"`
}

// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_CreatedByType string

const (
	SystemData_CreatedByType_Application     = SystemData_CreatedByType("Application")
	SystemData_CreatedByType_Key             = SystemData_CreatedByType("Key")
	SystemData_CreatedByType_ManagedIdentity = SystemData_CreatedByType("ManagedIdentity")
	SystemData_CreatedByType_User            = SystemData_CreatedByType("User")
)

// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_LastModifiedByType string

const (
	SystemData_LastModifiedByType_Application     = SystemData_LastModifiedByType("Application")
	SystemData_LastModifiedByType_Key             = SystemData_LastModifiedByType("Key")
	SystemData_LastModifiedByType_ManagedIdentity = SystemData_LastModifiedByType("ManagedIdentity")
	SystemData_LastModifiedByType_User            = SystemData_LastModifiedByType("User")
)