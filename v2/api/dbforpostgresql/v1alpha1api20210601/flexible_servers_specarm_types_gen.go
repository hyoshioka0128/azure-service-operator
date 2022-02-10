// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServers_SPECARM struct {
	AzureName string `json:"azureName"`

	//Location: The geo-location where the resource lives
	Location string `json:"location"`
	Name     string `json:"name"`

	//Properties: Properties of the server.
	Properties *ServerProperties_SpecARM `json:"properties,omitempty"`

	//Sku: The SKU (pricing tier) of the server.
	Sku *Sku_SpecARM `json:"sku,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServers_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (specarm FlexibleServers_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm FlexibleServers_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm FlexibleServers_SPECARM) GetType() string {
	return ""
}

type ServerProperties_SpecARM struct {
	//AdministratorLogin: The administrator's login name of a server. Can only be
	//specified when the server is being created (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	//AdministratorLoginPassword: The administrator login password (required for
	//server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	//AvailabilityZone: availability zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	//Backup: Backup properties of a server.
	Backup *Backup_SpecARM `json:"backup,omitempty"`

	//CreateMode: The mode to create a new PostgreSQL server.
	CreateMode *ServerPropertiesSpecCreateMode `json:"createMode,omitempty"`

	//HighAvailability: High availability properties of a server.
	HighAvailability *HighAvailability_SpecARM `json:"highAvailability,omitempty"`

	//MaintenanceWindow: Maintenance window properties of a server.
	MaintenanceWindow *MaintenanceWindow_SpecARM `json:"maintenanceWindow,omitempty"`

	//Network: Network properties of a server.
	Network *Network_SpecARM `json:"network,omitempty"`

	//PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the
	//time to restore from. It's required when 'createMode' is 'PointInTimeRestore'.
	PointInTimeUTC         *string `json:"pointInTimeUTC,omitempty"`
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	//Storage: Storage properties of a server.
	Storage *Storage_SpecARM `json:"storage,omitempty"`

	//Tags: Application-specific metadata in the form of key-value pairs.
	Tags map[string]string `json:"tags,omitempty"`

	//Version: PostgreSQL Server version.
	Version *ServerVersion_Spec `json:"version,omitempty"`
}

type Sku_SpecARM struct {
	//Name: The name of the sku, typically, tier + family + cores, e.g.
	//Standard_D4s_v3.
	Name string `json:"name"`

	//Tier: The tier of the particular SKU, e.g. Burstable.
	Tier SkuSpecTier `json:"tier"`
}

type Backup_SpecARM struct {
	//BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	//GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled
	//on the server.
	GeoRedundantBackup *BackupSpecGeoRedundantBackup `json:"geoRedundantBackup,omitempty"`
}

type HighAvailability_SpecARM struct {
	//Mode: The HA mode for the server.
	Mode *HighAvailabilitySpecMode `json:"mode,omitempty"`

	//StandbyAvailabilityZone: availability zone information of the standby.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`
}

type MaintenanceWindow_SpecARM struct {
	//CustomWindow: indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	//DayOfWeek: day of week for maintenance window
	DayOfWeek *int `json:"dayOfWeek,omitempty"`

	//StartHour: start hour for maintenance window
	StartHour *int `json:"startHour,omitempty"`

	//StartMinute: start minute for maintenance window
	StartMinute *int `json:"startMinute,omitempty"`
}

type Network_SpecARM struct {
	DelegatedSubnetResourceId   *string `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneArmResourceId *string `json:"privateDnsZoneArmResourceId,omitempty"`
}

// +kubebuilder:validation:Enum={"Burstable","GeneralPurpose","MemoryOptimized"}
type SkuSpecTier string

const (
	SkuSpecTierBurstable       = SkuSpecTier("Burstable")
	SkuSpecTierGeneralPurpose  = SkuSpecTier("GeneralPurpose")
	SkuSpecTierMemoryOptimized = SkuSpecTier("MemoryOptimized")
)

type Storage_SpecARM struct {
	//StorageSizeGB: Max storage allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`
}