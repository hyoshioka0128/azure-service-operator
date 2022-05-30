// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccountsQueueService_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`
	Name      string `json:"name,omitempty"`

	// Properties: The properties of a storage account’s Queue service.
	Properties *StorageAccountsQueueService_Spec_PropertiesARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccountsQueueService_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (service StorageAccountsQueueService_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (service *StorageAccountsQueueService_SpecARM) GetName() string {
	return service.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (service *StorageAccountsQueueService_SpecARM) GetType() string {
	return ""
}

type StorageAccountsQueueService_Spec_PropertiesARM struct {
	// Cors: Specifies CORS rules for the Queue service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// Queue service.
	Cors *CorsRulesARM `json:"cors,omitempty"`
}
