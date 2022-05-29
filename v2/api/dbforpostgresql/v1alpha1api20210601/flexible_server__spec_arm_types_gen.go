// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of FlexibleServer_Spec. Use v1beta20210601.FlexibleServer_Spec instead
type FlexibleServer_SpecARM struct {
	AzureName  string               `json:"azureName,omitempty"`
	Id         *string              `json:"id,omitempty"`
	Location   *string              `json:"location,omitempty"`
	Name       string               `json:"name,omitempty"`
	Properties *ServerPropertiesARM `json:"properties,omitempty"`
	Sku        *SkuARM              `json:"sku,omitempty"`
	SystemData *SystemDataARM       `json:"systemData,omitempty"`
	Tags       map[string]string    `json:"tags,omitempty"`
	Type       *string              `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServer_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
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

// Deprecated version of ServerProperties. Use v1beta20210601.ServerProperties instead
type ServerPropertiesARM struct {
	AdministratorLogin         *string                      `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword *string                      `json:"administratorLoginPassword,omitempty"`
	AvailabilityZone           *string                      `json:"availabilityZone,omitempty"`
	Backup                     *BackupARM                   `json:"backup,omitempty"`
	CreateMode                 *ServerProperties_CreateMode `json:"createMode,omitempty"`
	FullyQualifiedDomainName   *string                      `json:"fullyQualifiedDomainName,omitempty"`
	HighAvailability           *HighAvailabilityARM         `json:"highAvailability,omitempty"`
	MaintenanceWindow          *MaintenanceWindowARM        `json:"maintenanceWindow,omitempty"`
	MinorVersion               *string                      `json:"minorVersion,omitempty"`
	Network                    *NetworkARM                  `json:"network,omitempty"`
	PointInTimeUTC             *string                      `json:"pointInTimeUTC,omitempty"`
	SourceServerResourceId     *string                      `json:"sourceServerResourceId,omitempty"`
	State                      *ServerProperties_State      `json:"state,omitempty"`
	Storage                    *StorageARM                  `json:"storage,omitempty"`
	Version                    *ServerVersion               `json:"version,omitempty"`
}

// Deprecated version of Sku. Use v1beta20210601.Sku instead
type SkuARM struct {
	Name *string   `json:"name,omitempty"`
	Tier *Sku_Tier `json:"tier,omitempty"`
}

// Deprecated version of SystemData. Use v1beta20210601.SystemData instead
type SystemDataARM struct {
	CreatedAt          *string                        `json:"createdAt,omitempty"`
	CreatedBy          *string                        `json:"createdBy,omitempty"`
	CreatedByType      *SystemData_CreatedByType      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                        `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                        `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemData_LastModifiedByType `json:"lastModifiedByType,omitempty"`
}

// Deprecated version of Backup. Use v1beta20210601.Backup instead
type BackupARM struct {
	BackupRetentionDays *int                       `json:"backupRetentionDays,omitempty"`
	EarliestRestoreDate *string                    `json:"earliestRestoreDate,omitempty"`
	GeoRedundantBackup  *Backup_GeoRedundantBackup `json:"geoRedundantBackup,omitempty"`
}

// Deprecated version of HighAvailability. Use v1beta20210601.HighAvailability instead
type HighAvailabilityARM struct {
	Mode                    *HighAvailability_Mode  `json:"mode,omitempty"`
	StandbyAvailabilityZone *string                 `json:"standbyAvailabilityZone,omitempty"`
	State                   *HighAvailability_State `json:"state,omitempty"`
}

// Deprecated version of MaintenanceWindow. Use v1beta20210601.MaintenanceWindow instead
type MaintenanceWindowARM struct {
	CustomWindow *string `json:"customWindow,omitempty"`
	DayOfWeek    *int    `json:"dayOfWeek,omitempty"`
	StartHour    *int    `json:"startHour,omitempty"`
	StartMinute  *int    `json:"startMinute,omitempty"`
}

// Deprecated version of Network. Use v1beta20210601.Network instead
type NetworkARM struct {
	DelegatedSubnetResourceId   *string                      `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneArmResourceId *string                      `json:"privateDnsZoneArmResourceId,omitempty"`
	PublicNetworkAccess         *Network_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
}

// Deprecated version of Sku_Tier. Use v1beta20210601.Sku_Tier instead
// +kubebuilder:validation:Enum={"Burstable","GeneralPurpose","MemoryOptimized"}
type Sku_Tier string

const (
	Sku_Tier_Burstable       = Sku_Tier("Burstable")
	Sku_Tier_GeneralPurpose  = Sku_Tier("GeneralPurpose")
	Sku_Tier_MemoryOptimized = Sku_Tier("MemoryOptimized")
)

// Deprecated version of Storage. Use v1beta20210601.Storage instead
type StorageARM struct {
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`
}

// Deprecated version of SystemData_CreatedByType. Use v1beta20210601.SystemData_CreatedByType instead
// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_CreatedByType string

const (
	SystemData_CreatedByType_Application     = SystemData_CreatedByType("Application")
	SystemData_CreatedByType_Key             = SystemData_CreatedByType("Key")
	SystemData_CreatedByType_ManagedIdentity = SystemData_CreatedByType("ManagedIdentity")
	SystemData_CreatedByType_User            = SystemData_CreatedByType("User")
)

// Deprecated version of SystemData_LastModifiedByType. Use v1beta20210601.SystemData_LastModifiedByType instead
// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_LastModifiedByType string

const (
	SystemData_LastModifiedByType_Application     = SystemData_LastModifiedByType("Application")
	SystemData_LastModifiedByType_Key             = SystemData_LastModifiedByType("Key")
	SystemData_LastModifiedByType_ManagedIdentity = SystemData_LastModifiedByType("ManagedIdentity")
	SystemData_LastModifiedByType_User            = SystemData_LastModifiedByType("User")
)
