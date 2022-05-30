// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Image_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// ExtendedLocation: The extended location of the Image.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	// Location: Resource location
	Location   *string             `json:"location,omitempty"`
	Name       string              `json:"name,omitempty"`
	Properties *ImagePropertiesARM `json:"properties,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Image_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-07-01"
func (image Image_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (image *Image_SpecARM) GetName() string {
	return image.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (image *Image_SpecARM) GetType() string {
	return ""
}

type ExtendedLocationARM struct {
	// Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	// Type: The type of the extended location.
	Type *ExtendedLocationType `json:"type,omitempty"`
}

type ImagePropertiesARM struct {
	// HyperVGeneration: Specifies the HyperVGenerationType of the VirtualMachine created from the image. From API Version
	// 2019-03-01 if the image source is a blob, then we need the user to specify the value, if the source is managed resource
	// like disk or snapshot, we may require the user to specify the property if we cannot deduce it from the source managed
	// resource.
	HyperVGeneration *HyperVGenerationType `json:"hyperVGeneration,omitempty"`

	// SourceVirtualMachine: The source virtual machine from which Image is created.
	SourceVirtualMachine *SubResourceARM `json:"sourceVirtualMachine,omitempty"`

	// StorageProfile: Specifies the storage settings for the virtual machine disks.
	StorageProfile *ImageStorageProfileARM `json:"storageProfile,omitempty"`
}

// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType string

const ExtendedLocationType_EdgeZone = ExtendedLocationType("EdgeZone")

type ImageStorageProfileARM struct {
	// DataDisks: Specifies the parameters that are used to add a data disk to a virtual machine.
	// For more information about disks, see [About disks and VHDs for Azure virtual
	// machines](https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview).
	DataDisks []ImageDataDiskARM `json:"dataDisks,omitempty"`

	// OsDisk: Specifies information about the operating system disk used by the virtual machine.
	// For more information about disks, see [About disks and VHDs for Azure virtual
	// machines](https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview).
	OsDisk *ImageOSDiskARM `json:"osDisk,omitempty"`

	// ZoneResilient: Specifies whether an image is zone resilient or not. Default is false. Zone resilient images can be
	// created only in regions that provide Zone Redundant Storage (ZRS).
	ZoneResilient *bool `json:"zoneResilient,omitempty"`
}

type SubResourceARM struct {
	Id *string `json:"id,omitempty"`
}

type ImageDataDiskARM struct {
	// BlobUri: The Virtual Hard Disk.
	BlobUri *string `json:"blobUri,omitempty"`

	// Caching: Specifies the caching requirements.
	// Possible values are:
	// None
	// ReadOnly
	// ReadWrite
	// Default: None for Standard storage. ReadOnly for Premium storage
	Caching *ImageDataDisk_Caching `json:"caching,omitempty"`

	// DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed image disk.
	DiskEncryptionSet *SubResourceARM `json:"diskEncryptionSet,omitempty"`

	// DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
	// disk in a virtual machine image.
	// This value cannot be larger than 1023 GB
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
	// therefore must be unique for each data disk attached to a VM.
	Lun *int `json:"lun,omitempty"`

	// ManagedDisk: The managedDisk.
	ManagedDisk *SubResourceARM `json:"managedDisk,omitempty"`

	// Snapshot: The snapshot.
	Snapshot *SubResourceARM `json:"snapshot,omitempty"`

	// StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
	// data disks, it cannot be used with OS Disk.
	StorageAccountType *StorageAccountType `json:"storageAccountType,omitempty"`
}

type ImageOSDiskARM struct {
	// BlobUri: The Virtual Hard Disk.
	BlobUri *string `json:"blobUri,omitempty"`

	// Caching: Specifies the caching requirements.
	// Possible values are:
	// None
	// ReadOnly
	// ReadWrite
	// Default: None for Standard storage. ReadOnly for Premium storage
	Caching *ImageOSDisk_Caching `json:"caching,omitempty"`

	// DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed image disk.
	DiskEncryptionSet *SubResourceARM `json:"diskEncryptionSet,omitempty"`

	// DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
	// disk in a virtual machine image.
	// This value cannot be larger than 1023 GB
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// ManagedDisk: The managedDisk.
	ManagedDisk *SubResourceARM `json:"managedDisk,omitempty"`

	// OsState: The OS State.
	OsState *ImageOSDisk_OsState `json:"osState,omitempty"`

	// OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from a
	// custom image.
	// Possible values are:
	// Windows
	// Linux
	OsType *ImageOSDisk_OsType `json:"osType,omitempty"`

	// Snapshot: The snapshot.
	Snapshot *SubResourceARM `json:"snapshot,omitempty"`

	// StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
	// data disks, it cannot be used with OS Disk.
	StorageAccountType *StorageAccountType `json:"storageAccountType,omitempty"`
}
