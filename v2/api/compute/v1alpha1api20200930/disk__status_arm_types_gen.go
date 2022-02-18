// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200930

type Disk_StatusARM struct {
	//ExtendedLocation: The extended location where the disk will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation_StatusARM `json:"extendedLocation,omitempty"`

	//Id: Resource Id
	Id *string `json:"id,omitempty"`

	//Location: Resource location
	Location *string `json:"location,omitempty"`

	//ManagedBy: A relative URI containing the ID of the VM that has the disk attached.
	ManagedBy *string `json:"managedBy,omitempty"`

	//ManagedByExtended: List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be
	//set to a value greater than one for disks to allow attaching them to multiple VMs.
	ManagedByExtended []string `json:"managedByExtended,omitempty"`

	//Name: Resource name
	Name       *string                   `json:"name,omitempty"`
	Properties *DiskProperties_StatusARM `json:"properties,omitempty"`
	Sku        *DiskSku_StatusARM        `json:"sku,omitempty"`

	//Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Resource type
	Type *string `json:"type,omitempty"`

	//Zones: The Logical zone list for Disk.
	Zones []string `json:"zones,omitempty"`
}

type DiskProperties_StatusARM struct {
	//BurstingEnabled: Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is
	//disabled by default. Does not apply to Ultra disks.
	BurstingEnabled *bool `json:"burstingEnabled,omitempty"`

	//CreationData: Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationData_StatusARM `json:"creationData"`

	//DiskAccessId: ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId *string `json:"diskAccessId,omitempty"`

	//DiskIOPSReadOnly: The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One
	//operation can transfer between 4k and 256k bytes.
	DiskIOPSReadOnly *int `json:"diskIOPSReadOnly,omitempty"`

	//DiskIOPSReadWrite: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can
	//transfer between 4k and 256k bytes.
	DiskIOPSReadWrite *int `json:"diskIOPSReadWrite,omitempty"`

	//DiskMBpsReadOnly: The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly.
	//MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadOnly *int `json:"diskMBpsReadOnly,omitempty"`

	//DiskMBpsReadWrite: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes
	//per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadWrite *int `json:"diskMBpsReadWrite,omitempty"`

	//DiskSizeBytes: The size of the disk in bytes. This field is read only.
	DiskSizeBytes *int `json:"diskSizeBytes,omitempty"`

	//DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
	//create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
	//allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	//DiskState: The state of the disk.
	DiskState *DiskState_Status `json:"diskState,omitempty"`

	//Encryption: Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *Encryption_StatusARM `json:"encryption,omitempty"`

	//EncryptionSettingsCollection: Encryption settings collection used for Azure Disk Encryption, can contain multiple
	//encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollection_StatusARM `json:"encryptionSettingsCollection,omitempty"`

	//HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *DiskPropertiesStatusHyperVGeneration `json:"hyperVGeneration,omitempty"`

	//MaxShares: The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a
	//disk that can be mounted on multiple VMs at the same time.
	MaxShares           *int                        `json:"maxShares,omitempty"`
	NetworkAccessPolicy *NetworkAccessPolicy_Status `json:"networkAccessPolicy,omitempty"`

	//OsType: The Operating System type.
	OsType *DiskPropertiesStatusOsType `json:"osType,omitempty"`

	//ProvisioningState: The disk provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	//PurchasePlan: Purchase plan information for the the image from which the OS disk was created. E.g. - {name:
	//2019-Datacenter, publisher: MicrosoftWindowsServer, product: WindowsServer}
	PurchasePlan *PurchasePlan_StatusARM `json:"purchasePlan,omitempty"`

	//ShareInfo: Details of the list of all VMs that have the disk attached. maxShares should be set to a value greater than
	//one for disks to allow attaching them to multiple VMs.
	ShareInfo []ShareInfoElement_StatusARM `json:"shareInfo,omitempty"`

	//Tier: Performance tier of the disk (e.g, P4, S10) as described here:
	//https://azure.microsoft.com/en-us/pricing/details/managed-disks/. Does not apply to Ultra disks.
	Tier *string `json:"tier,omitempty"`

	//TimeCreated: The time when the disk was created.
	TimeCreated *string `json:"timeCreated,omitempty"`

	//UniqueId: Unique Guid identifying the resource.
	UniqueId *string `json:"uniqueId,omitempty"`
}

type DiskSku_StatusARM struct {
	//Name: The sku name.
	Name *DiskSkuStatusName `json:"name,omitempty"`

	//Tier: The sku tier.
	Tier *string `json:"tier,omitempty"`
}

type ExtendedLocation_StatusARM struct {
	//Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	//Type: The type of the extended location.
	Type *ExtendedLocationType_Status `json:"type,omitempty"`
}

type CreationData_StatusARM struct {
	//CreateOption: This enumerates the possible sources of a disk's creation.
	CreateOption CreationDataStatusCreateOption `json:"createOption"`

	//GalleryImageReference: Required if creating from a Gallery Image. The id of the ImageDiskReference will be the ARM id of
	//the shared galley image version from which to create a disk.
	GalleryImageReference *ImageDiskReference_StatusARM `json:"galleryImageReference,omitempty"`

	//ImageReference: Disk source information.
	ImageReference *ImageDiskReference_StatusARM `json:"imageReference,omitempty"`

	//LogicalSectorSize: Logical sector size in bytes for Ultra disks. Supported values are 512 ad 4096. 4096 is the default.
	LogicalSectorSize *int `json:"logicalSectorSize,omitempty"`

	//SourceResourceId: If createOption is Copy, this is the ARM id of the source snapshot or disk.
	SourceResourceId *string `json:"sourceResourceId,omitempty"`

	//SourceUniqueId: If this field is set, this is the unique id identifying the source of this resource.
	SourceUniqueId *string `json:"sourceUniqueId,omitempty"`

	//SourceUri: If createOption is Import, this is the URI of a blob to be imported into a managed disk.
	SourceUri *string `json:"sourceUri,omitempty"`

	//StorageAccountId: Required if createOption is Import. The Azure Resource Manager identifier of the storage account
	//containing the blob to import as a disk.
	StorageAccountId *string `json:"storageAccountId,omitempty"`

	//UploadSizeBytes: If createOption is Upload, this is the size of the contents of the upload including the VHD footer.
	//This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512
	//bytes for the VHD footer).
	UploadSizeBytes *int `json:"uploadSizeBytes,omitempty"`
}

type DiskSkuStatusName string

const (
	DiskSkuStatusNamePremiumLRS     = DiskSkuStatusName("Premium_LRS")
	DiskSkuStatusNameStandardLRS    = DiskSkuStatusName("Standard_LRS")
	DiskSkuStatusNameStandardSSDLRS = DiskSkuStatusName("StandardSSD_LRS")
	DiskSkuStatusNameUltraSSDLRS    = DiskSkuStatusName("UltraSSD_LRS")
)

type EncryptionSettingsCollection_StatusARM struct {
	//Enabled: Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set
	//this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is
	//null in the request object, the existing settings remain unchanged.
	Enabled bool `json:"enabled"`

	//EncryptionSettings: A collection of encryption settings, one for each disk volume.
	EncryptionSettings []EncryptionSettingsElement_StatusARM `json:"encryptionSettings,omitempty"`

	//EncryptionSettingsVersion: Describes what type of encryption is used for the disks. Once this field is set, it cannot be
	//overwritten. '1.0' corresponds to Azure Disk Encryption with AAD app.'1.1' corresponds to Azure Disk Encryption.
	EncryptionSettingsVersion *string `json:"encryptionSettingsVersion,omitempty"`
}

type Encryption_StatusARM struct {
	//DiskEncryptionSetId: ResourceId of the disk encryption set to use for enabling encryption at rest.
	DiskEncryptionSetId *string                `json:"diskEncryptionSetId,omitempty"`
	Type                *EncryptionType_Status `json:"type,omitempty"`
}

type ExtendedLocationType_Status string

const ExtendedLocationType_StatusEdgeZone = ExtendedLocationType_Status("EdgeZone")

type PurchasePlan_StatusARM struct {
	//Name: The plan ID.
	Name string `json:"name"`

	//Product: Specifies the product of the image from the marketplace. This is the same value as Offer under the
	//imageReference element.
	Product string `json:"product"`

	//PromotionCode: The Offer Promotion Code.
	PromotionCode *string `json:"promotionCode,omitempty"`

	//Publisher: The publisher ID.
	Publisher string `json:"publisher"`
}

type ShareInfoElement_StatusARM struct {
	//VmUri: A relative URI containing the ID of the VM that has the disk attached.
	VmUri *string `json:"vmUri,omitempty"`
}

type EncryptionSettingsElement_StatusARM struct {
	//DiskEncryptionKey: Key Vault Secret Url and vault id of the disk encryption key
	DiskEncryptionKey *KeyVaultAndSecretReference_StatusARM `json:"diskEncryptionKey,omitempty"`

	//KeyEncryptionKey: Key Vault Key Url and vault id of the key encryption key. KeyEncryptionKey is optional and when
	//provided is used to unwrap the disk encryption key.
	KeyEncryptionKey *KeyVaultAndKeyReference_StatusARM `json:"keyEncryptionKey,omitempty"`
}

type ImageDiskReference_StatusARM struct {
	//Id: A relative uri containing either a Platform Image Repository or user image reference.
	Id string `json:"id"`

	//Lun: If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the
	//image to use. For OS disks, this field is null.
	Lun *int `json:"lun,omitempty"`
}

type KeyVaultAndKeyReference_StatusARM struct {
	//KeyUrl: Url pointing to a key or secret in KeyVault
	KeyUrl string `json:"keyUrl"`

	//SourceVault: Resource id of the KeyVault containing the key or secret
	SourceVault SourceVault_StatusARM `json:"sourceVault"`
}

type KeyVaultAndSecretReference_StatusARM struct {
	//SecretUrl: Url pointing to a key or secret in KeyVault
	SecretUrl string `json:"secretUrl"`

	//SourceVault: Resource id of the KeyVault containing the key or secret
	SourceVault SourceVault_StatusARM `json:"sourceVault"`
}

type SourceVault_StatusARM struct {
	//Id: Resource Id
	Id *string `json:"id,omitempty"`
}