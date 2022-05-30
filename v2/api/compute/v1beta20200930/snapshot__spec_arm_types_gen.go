// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200930

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Snapshot_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// ExtendedLocation: The extended location where the snapshot will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	// Location: Resource location
	Location   *string                `json:"location,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Properties *SnapshotPropertiesARM `json:"properties,omitempty"`
	Sku        *SnapshotSkuARM        `json:"sku,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`
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

type SnapshotPropertiesARM struct {
	// CreationData: Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData *CreationDataARM `json:"creationData,omitempty"`
	DiskAccessId *string          `json:"diskAccessId,omitempty"`

	// DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
	// create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
	// allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// DiskState: The state of the snapshot.
	DiskState *DiskState `json:"diskState,omitempty"`

	// Encryption: Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *EncryptionARM `json:"encryption,omitempty"`

	// EncryptionSettingsCollection: Encryption settings collection used be Azure Disk Encryption, can contain multiple
	// encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollectionARM `json:"encryptionSettingsCollection,omitempty"`

	// HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *SnapshotProperties_HyperVGeneration `json:"hyperVGeneration,omitempty"`

	// Incremental: Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full
	// snapshots and can be diffed.
	Incremental         *bool                `json:"incremental,omitempty"`
	NetworkAccessPolicy *NetworkAccessPolicy `json:"networkAccessPolicy,omitempty"`

	// OsType: The Operating System type.
	OsType *SnapshotProperties_OsType `json:"osType,omitempty"`

	// PurchasePlan: Purchase plan information for the image from which the source disk for the snapshot was originally created.
	PurchasePlan *PurchasePlanARM `json:"purchasePlan,omitempty"`
}

type SnapshotSkuARM struct {
	// Name: The sku name.
	Name *SnapshotSku_Name `json:"name,omitempty"`
}

// +kubebuilder:validation:Enum={"Premium_LRS","Standard_LRS","Standard_ZRS"}
type SnapshotSku_Name string

const (
	SnapshotSku_Name_Premium_LRS  = SnapshotSku_Name("Premium_LRS")
	SnapshotSku_Name_Standard_LRS = SnapshotSku_Name("Standard_LRS")
	SnapshotSku_Name_Standard_ZRS = SnapshotSku_Name("Standard_ZRS")
)
