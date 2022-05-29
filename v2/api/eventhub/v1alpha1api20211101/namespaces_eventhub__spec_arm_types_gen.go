// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of NamespacesEventhub_Spec. Use v1beta20211101.NamespacesEventhub_Spec instead
type NamespacesEventhub_SpecARM struct {
	AzureName  string                                 `json:"azureName,omitempty"`
	Id         *string                                `json:"id,omitempty"`
	Location   *string                                `json:"location,omitempty"`
	Name       string                                 `json:"name,omitempty"`
	Properties *NamespacesEventhub_Spec_PropertiesARM `json:"properties,omitempty"`
	SystemData *SystemDataARM                         `json:"systemData,omitempty"`
	Type       *string                                `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesEventhub_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (eventhub NamespacesEventhub_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (eventhub *NamespacesEventhub_SpecARM) GetName() string {
	return eventhub.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (eventhub *NamespacesEventhub_SpecARM) GetType() string {
	return ""
}

// Deprecated version of NamespacesEventhub_Spec_Properties. Use v1beta20211101.NamespacesEventhub_Spec_Properties instead
type NamespacesEventhub_Spec_PropertiesARM struct {
	CaptureDescription     *CaptureDescriptionARM                     `json:"captureDescription,omitempty"`
	CreatedAt              *string                                    `json:"createdAt,omitempty"`
	MessageRetentionInDays *int                                       `json:"messageRetentionInDays,omitempty"`
	PartitionCount         *int                                       `json:"partitionCount,omitempty"`
	PartitionIds           []string                                   `json:"partitionIds,omitempty"`
	Status                 *NamespacesEventhub_Spec_Properties_Status `json:"status,omitempty"`
	UpdatedAt              *string                                    `json:"updatedAt,omitempty"`
}

// Deprecated version of CaptureDescription. Use v1beta20211101.CaptureDescription instead
type CaptureDescriptionARM struct {
	Destination       *DestinationARM              `json:"destination,omitempty"`
	Enabled           *bool                        `json:"enabled,omitempty"`
	Encoding          *CaptureDescription_Encoding `json:"encoding,omitempty"`
	IntervalInSeconds *int                         `json:"intervalInSeconds,omitempty"`
	SizeLimitInBytes  *int                         `json:"sizeLimitInBytes,omitempty"`
	SkipEmptyArchives *bool                        `json:"skipEmptyArchives,omitempty"`
}

// Deprecated version of Destination. Use v1beta20211101.Destination instead
type DestinationARM struct {
	Name       *string                    `json:"name,omitempty"`
	Properties *Destination_PropertiesARM `json:"properties,omitempty"`
}

// Deprecated version of Destination_Properties. Use v1beta20211101.Destination_Properties instead
type Destination_PropertiesARM struct {
	ArchiveNameFormat        *string `json:"archiveNameFormat,omitempty"`
	BlobContainer            *string `json:"blobContainer,omitempty"`
	DataLakeAccountName      *string `json:"dataLakeAccountName,omitempty"`
	DataLakeFolderPath       *string `json:"dataLakeFolderPath,omitempty"`
	DataLakeSubscriptionId   *string `json:"dataLakeSubscriptionId,omitempty"`
	StorageAccountResourceId *string `json:"storageAccountResourceId,omitempty"`
}
