// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of NamespacesQueue_Spec. Use v1beta20210101preview.NamespacesQueue_Spec instead
type NamespacesQueue_SpecARM struct {
	AzureName  string                `json:"azureName,omitempty"`
	Id         *string               `json:"id,omitempty"`
	Name       string                `json:"name,omitempty"`
	Properties *SBQueuePropertiesARM `json:"properties,omitempty"`
	SystemData *SystemDataARM        `json:"systemData,omitempty"`
	Type       *string               `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesQueue_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (queue NamespacesQueue_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (queue *NamespacesQueue_SpecARM) GetName() string {
	return queue.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (queue *NamespacesQueue_SpecARM) GetType() string {
	return ""
}

// Deprecated version of SBQueueProperties. Use v1beta20210101preview.SBQueueProperties instead
type SBQueuePropertiesARM struct {
	AccessedAt                          *string                 `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string                 `json:"autoDeleteOnIdle,omitempty"`
	CountDetails                        *MessageCountDetailsARM `json:"countDetails,omitempty"`
	CreatedAt                           *string                 `json:"createdAt,omitempty"`
	DeadLetteringOnMessageExpiration    *bool                   `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive            *string                 `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                 `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                   `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                   `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                   `json:"enablePartitioning,omitempty"`
	ForwardDeadLetteredMessagesTo       *string                 `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                           *string                 `json:"forwardTo,omitempty"`
	LockDuration                        *string                 `json:"lockDuration,omitempty"`
	MaxDeliveryCount                    *int                    `json:"maxDeliveryCount,omitempty"`
	MaxSizeInMegabytes                  *int                    `json:"maxSizeInMegabytes,omitempty"`
	MessageCount                        *int                    `json:"messageCount,omitempty"`
	RequiresDuplicateDetection          *bool                   `json:"requiresDuplicateDetection,omitempty"`
	RequiresSession                     *bool                   `json:"requiresSession,omitempty"`
	SizeInBytes                         *int                    `json:"sizeInBytes,omitempty"`
	UpdatedAt                           *string                 `json:"updatedAt,omitempty"`
}

// Deprecated version of MessageCountDetails. Use v1beta20210101preview.MessageCountDetails instead
type MessageCountDetailsARM struct {
	ActiveMessageCount             *int `json:"activeMessageCount,omitempty"`
	DeadLetterMessageCount         *int `json:"deadLetterMessageCount,omitempty"`
	ScheduledMessageCount          *int `json:"scheduledMessageCount,omitempty"`
	TransferDeadLetterMessageCount *int `json:"transferDeadLetterMessageCount,omitempty"`
	TransferMessageCount           *int `json:"transferMessageCount,omitempty"`
}
