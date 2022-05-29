// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200930

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Snapshot_Spec. Use v1beta20200930.Snapshot_Spec instead
type Snapshot_SpecARM struct {
	AzureName        string                 `json:"azureName,omitempty"`
	ExtendedLocation *ExtendedLocationARM   `json:"extendedLocation,omitempty"`
	Id               *string                `json:"id,omitempty"`
	Location         *string                `json:"location,omitempty"`
	ManagedBy        *string                `json:"managedBy,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Properties       *SnapshotPropertiesARM `json:"properties,omitempty"`
	Sku              *SnapshotSkuARM        `json:"sku,omitempty"`
	Tags             map[string]string      `json:"tags,omitempty"`
	Type             *string                `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Snapshot_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-09-30"
func (snapshot Snapshot_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (snapshot *Snapshot_SpecARM) GetName() string {
	return snapshot.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (snapshot *Snapshot_SpecARM) GetType() string {
	return ""
}

// Deprecated version of SnapshotProperties. Use v1beta20200930.SnapshotProperties instead
type SnapshotPropertiesARM struct {
	CreationData                 *CreationDataARM                     `json:"creationData,omitempty"`
	DiskAccessId                 *string                              `json:"diskAccessId,omitempty"`
	DiskSizeBytes                *int                                 `json:"diskSizeBytes,omitempty"`
	DiskSizeGB                   *int                                 `json:"diskSizeGB,omitempty"`
	DiskState                    *DiskState                           `json:"diskState,omitempty"`
	Encryption                   *EncryptionARM                       `json:"encryption,omitempty"`
	EncryptionSettingsCollection *EncryptionSettingsCollectionARM     `json:"encryptionSettingsCollection,omitempty"`
	HyperVGeneration             *SnapshotProperties_HyperVGeneration `json:"hyperVGeneration,omitempty"`
	Incremental                  *bool                                `json:"incremental,omitempty"`
	NetworkAccessPolicy          *NetworkAccessPolicy                 `json:"networkAccessPolicy,omitempty"`
	OsType                       *SnapshotProperties_OsType           `json:"osType,omitempty"`
	ProvisioningState            *string                              `json:"provisioningState,omitempty"`
	PurchasePlan                 *PurchasePlanARM                     `json:"purchasePlan,omitempty"`
	TimeCreated                  *string                              `json:"timeCreated,omitempty"`
	UniqueId                     *string                              `json:"uniqueId,omitempty"`
}

// Deprecated version of SnapshotSku. Use v1beta20200930.SnapshotSku instead
type SnapshotSkuARM struct {
	Name *SnapshotSku_Name `json:"name,omitempty"`
	Tier *string           `json:"tier,omitempty"`
}

// Deprecated version of SnapshotSku_Name. Use v1beta20200930.SnapshotSku_Name instead
// +kubebuilder:validation:Enum={"Premium_LRS","Standard_LRS","Standard_ZRS"}
type SnapshotSku_Name string

const (
	SnapshotSku_Name_Premium_LRS  = SnapshotSku_Name("Premium_LRS")
	SnapshotSku_Name_Standard_LRS = SnapshotSku_Name("Standard_LRS")
	SnapshotSku_Name_Standard_ZRS = SnapshotSku_Name("Standard_ZRS")
)
