// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of StorageAccountsBlobServicesContainer_Spec. Use v1beta20210401.StorageAccountsBlobServicesContainer_Spec instead
type StorageAccountsBlobServicesContainer_SpecARM struct {
	AzureName  string                  `json:"azureName,omitempty"`
	Etag       *string                 `json:"etag,omitempty"`
	Id         *string                 `json:"id,omitempty"`
	Name       string                  `json:"name,omitempty"`
	Properties *ContainerPropertiesARM `json:"properties,omitempty"`
	Type       *string                 `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccountsBlobServicesContainer_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (container StorageAccountsBlobServicesContainer_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (container *StorageAccountsBlobServicesContainer_SpecARM) GetName() string {
	return container.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (container *StorageAccountsBlobServicesContainer_SpecARM) GetType() string {
	return ""
}

// Deprecated version of ContainerProperties. Use v1beta20210401.ContainerProperties instead
type ContainerPropertiesARM struct {
	DefaultEncryptionScope         *string                            `json:"defaultEncryptionScope,omitempty"`
	Deleted                        *bool                              `json:"deleted,omitempty"`
	DeletedTime                    *string                            `json:"deletedTime,omitempty"`
	DenyEncryptionScopeOverride    *bool                              `json:"denyEncryptionScopeOverride,omitempty"`
	HasImmutabilityPolicy          *bool                              `json:"hasImmutabilityPolicy,omitempty"`
	HasLegalHold                   *bool                              `json:"hasLegalHold,omitempty"`
	ImmutabilityPolicy             *ImmutabilityPolicyPropertiesARM   `json:"immutabilityPolicy,omitempty"`
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioningARM `json:"immutableStorageWithVersioning,omitempty"`
	LastModifiedTime               *string                            `json:"lastModifiedTime,omitempty"`
	LeaseDuration                  *ContainerProperties_LeaseDuration `json:"leaseDuration,omitempty"`
	LeaseState                     *ContainerProperties_LeaseState    `json:"leaseState,omitempty"`
	LeaseStatus                    *ContainerProperties_LeaseStatus   `json:"leaseStatus,omitempty"`
	LegalHold                      *LegalHoldPropertiesARM            `json:"legalHold,omitempty"`
	Metadata                       map[string]string                  `json:"metadata,omitempty"`
	PublicAccess                   *ContainerProperties_PublicAccess  `json:"publicAccess,omitempty"`
	RemainingRetentionDays         *int                               `json:"remainingRetentionDays,omitempty"`
	Version                        *string                            `json:"version,omitempty"`
}

// Deprecated version of ImmutabilityPolicyProperties. Use v1beta20210401.ImmutabilityPolicyProperties instead
type ImmutabilityPolicyPropertiesARM struct {
	Etag          *string                        `json:"etag,omitempty"`
	Properties    *ImmutabilityPolicyPropertyARM `json:"properties,omitempty"`
	UpdateHistory []UpdateHistoryPropertyARM     `json:"updateHistory,omitempty"`
}

// Deprecated version of ImmutableStorageWithVersioning. Use v1beta20210401.ImmutableStorageWithVersioning instead
type ImmutableStorageWithVersioningARM struct {
	Enabled        *bool                                          `json:"enabled,omitempty"`
	MigrationState *ImmutableStorageWithVersioning_MigrationState `json:"migrationState,omitempty"`
	TimeStamp      *string                                        `json:"timeStamp,omitempty"`
}

// Deprecated version of LegalHoldProperties. Use v1beta20210401.LegalHoldProperties instead
type LegalHoldPropertiesARM struct {
	HasLegalHold *bool            `json:"hasLegalHold,omitempty"`
	Tags         []TagPropertyARM `json:"tags,omitempty"`
}

// Deprecated version of ImmutabilityPolicyProperty. Use v1beta20210401.ImmutabilityPolicyProperty instead
type ImmutabilityPolicyPropertyARM struct {
	AllowProtectedAppendWrites            *bool                             `json:"allowProtectedAppendWrites,omitempty"`
	ImmutabilityPeriodSinceCreationInDays *int                              `json:"immutabilityPeriodSinceCreationInDays,omitempty"`
	State                                 *ImmutabilityPolicyProperty_State `json:"state,omitempty"`
}

// Deprecated version of TagProperty. Use v1beta20210401.TagProperty instead
type TagPropertyARM struct {
	ObjectIdentifier *string `json:"objectIdentifier,omitempty"`
	Tag              *string `json:"tag,omitempty"`
	TenantId         *string `json:"tenantId,omitempty"`
	Timestamp        *string `json:"timestamp,omitempty"`
	Upn              *string `json:"upn,omitempty"`
}

// Deprecated version of UpdateHistoryProperty. Use v1beta20210401.UpdateHistoryProperty instead
type UpdateHistoryPropertyARM struct {
	ImmutabilityPeriodSinceCreationInDays *int                          `json:"immutabilityPeriodSinceCreationInDays,omitempty"`
	ObjectIdentifier                      *string                       `json:"objectIdentifier,omitempty"`
	TenantId                              *string                       `json:"tenantId,omitempty"`
	Timestamp                             *string                       `json:"timestamp,omitempty"`
	Update                                *UpdateHistoryProperty_Update `json:"update,omitempty"`
	Upn                                   *string                       `json:"upn,omitempty"`
}