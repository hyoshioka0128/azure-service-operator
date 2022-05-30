// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of RedisPatchSchedule_Spec. Use v1beta20201201.RedisPatchSchedule_Spec instead
type RedisPatchSchedule_SpecARM struct {
	AzureName  string              `json:"azureName,omitempty"`
	Id         *string             `json:"id,omitempty"`
	Location   *string             `json:"location,omitempty"`
	Name       string              `json:"name,omitempty"`
	Properties *ScheduleEntriesARM `json:"properties,omitempty"`
	Type       *string             `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisPatchSchedule_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (schedule RedisPatchSchedule_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (schedule *RedisPatchSchedule_SpecARM) GetName() string {
	return schedule.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (schedule *RedisPatchSchedule_SpecARM) GetType() string {
	return ""
}

// Deprecated version of ScheduleEntries. Use v1beta20201201.ScheduleEntries instead
type ScheduleEntriesARM struct {
	ScheduleEntries []ScheduleEntryARM `json:"scheduleEntries,omitempty"`
}

// Deprecated version of ScheduleEntry. Use v1beta20201201.ScheduleEntry instead
type ScheduleEntryARM struct {
	DayOfWeek         *ScheduleEntry_DayOfWeek `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                  `json:"maintenanceWindow,omitempty"`
	StartHourUtc      *int                     `json:"startHourUtc,omitempty"`
}