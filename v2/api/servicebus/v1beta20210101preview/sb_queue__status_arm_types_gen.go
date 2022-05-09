// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

type SBQueue_StatusARM struct {
	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Name: Resource name
	Name *string `json:"name,omitempty"`

	// Properties: Queue Properties
	Properties *SBQueueProperties_StatusARM `json:"properties,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

type SBQueueProperties_StatusARM struct {
	// AccessedAt: Last time a message was sent, or the last time there was a receive request to this queue.
	AccessedAt *string `json:"accessedAt,omitempty"`

	// AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration
	// is 5 minutes.
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// CountDetails: Message Count Details.
	CountDetails *MessageCountDetails_StatusARM `json:"countDetails,omitempty"`

	// CreatedAt: The exact time the message was created.
	CreatedAt *string `json:"createdAt,omitempty"`

	// DeadLetteringOnMessageExpiration: A value that indicates whether this queue has dead letter support when a message
	// expires.
	DeadLetteringOnMessageExpiration *bool `json:"deadLetteringOnMessageExpiration,omitempty"`

	// DefaultMessageTimeToLive: ISO 8601 default message timespan to live value. This is the duration after which the message
	// expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
	// set on a message itself.
	DefaultMessageTimeToLive *string `json:"defaultMessageTimeToLive,omitempty"`

	// DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
	// history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`

	// EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty"`

	// EnableExpress: A value that indicates whether Express Entities are enabled. An express queue holds a message in memory
	// temporarily before writing it to persistent storage.
	EnableExpress *bool `json:"enableExpress,omitempty"`

	// EnablePartitioning: A value that indicates whether the queue is to be partitioned across multiple message brokers.
	EnablePartitioning *bool `json:"enablePartitioning,omitempty"`

	// ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message
	ForwardDeadLetteredMessagesTo *string `json:"forwardDeadLetteredMessagesTo,omitempty"`

	// ForwardTo: Queue/Topic name to forward the messages
	ForwardTo *string `json:"forwardTo,omitempty"`

	// LockDuration: ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for
	// other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.
	LockDuration *string `json:"lockDuration,omitempty"`

	// MaxDeliveryCount: The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
	// default value is 10.
	MaxDeliveryCount *int `json:"maxDeliveryCount,omitempty"`

	// MaxSizeInMegabytes: The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
	// Default is 1024.
	MaxSizeInMegabytes *int `json:"maxSizeInMegabytes,omitempty"`

	// MessageCount: The number of messages in the queue.
	MessageCount *int `json:"messageCount,omitempty"`

	// RequiresDuplicateDetection: A value indicating if this queue requires duplicate detection.
	RequiresDuplicateDetection *bool `json:"requiresDuplicateDetection,omitempty"`

	// RequiresSession: A value that indicates whether the queue supports the concept of sessions.
	RequiresSession *bool `json:"requiresSession,omitempty"`

	// SizeInBytes: The size of the queue, in bytes.
	SizeInBytes *int `json:"sizeInBytes,omitempty"`

	// Status: Enumerates the possible values for the status of a messaging entity.
	Status *EntityStatus_Status `json:"status,omitempty"`

	// UpdatedAt: The exact time the message was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type MessageCountDetails_StatusARM struct {
	// ActiveMessageCount: Number of active messages in the queue, topic, or subscription.
	ActiveMessageCount *int `json:"activeMessageCount,omitempty"`

	// DeadLetterMessageCount: Number of messages that are dead lettered.
	DeadLetterMessageCount *int `json:"deadLetterMessageCount,omitempty"`

	// ScheduledMessageCount: Number of scheduled messages.
	ScheduledMessageCount *int `json:"scheduledMessageCount,omitempty"`

	// TransferDeadLetterMessageCount: Number of messages transferred into dead letters.
	TransferDeadLetterMessageCount *int `json:"transferDeadLetterMessageCount,omitempty"`

	// TransferMessageCount: Number of messages transferred to another queue, topic, or subscription.
	TransferMessageCount *int `json:"transferMessageCount,omitempty"`
}
