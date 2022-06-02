// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccountsQueueServicesQueue_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`
	Name      string `json:"name,omitempty"`

	// Properties: Queue resource properties.
	Properties *QueuePropertiesARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccountsQueueServicesQueue_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (queue StorageAccountsQueueServicesQueue_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (queue *StorageAccountsQueueServicesQueue_SpecARM) GetName() string {
	return queue.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/queueServices/queues"
func (queue *StorageAccountsQueueServicesQueue_SpecARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/queueServices/queues"
}

type QueuePropertiesARM struct {
	// Metadata: A name-value pair that represents queue metadata.
	Metadata map[string]string `json:"metadata,omitempty"`
}