// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisPatchSchedules_SPECARM struct {
	AzureName string `json:"azureName"`
	Name      string `json:"name"`

	//Properties: List of patch schedules for a Redis cache.
	Properties ScheduleEntries_SpecARM `json:"properties"`
}

var _ genruntime.ARMResourceSpec = &RedisPatchSchedules_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (specarm RedisPatchSchedules_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm RedisPatchSchedules_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/patchSchedules"
func (specarm RedisPatchSchedules_SPECARM) GetType() string {
	return "Microsoft.Cache/redis/patchSchedules"
}

type ScheduleEntries_SpecARM struct {
	//ScheduleEntries: List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntry_SpecARM `json:"scheduleEntries"`
}

type ScheduleEntry_SpecARM struct {
	//DayOfWeek: Day of the week when a cache can be patched.
	DayOfWeek ScheduleEntry_DayOfWeek_Spec `json:"dayOfWeek"`

	//MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can
	//take.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty"`

	//StartHourUtc: Start hour after which cache patching can start.
	StartHourUtc int `json:"startHourUtc"`
}
