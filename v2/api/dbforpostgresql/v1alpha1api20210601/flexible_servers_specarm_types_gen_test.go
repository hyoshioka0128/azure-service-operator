// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_FlexibleServers_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersSPECARM, FlexibleServersSPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersSPECARM runs a test to see if a specific instance of FlexibleServers_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersSPECARM(subject FlexibleServers_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_SPECARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of FlexibleServers_SPECARM instances for property testing - lazily instantiated by
//FlexibleServersSPECARMGenerator()
var flexibleServersSPECARMGenerator gopter.Gen

// FlexibleServersSPECARMGenerator returns a generator of FlexibleServers_SPECARM instances for property testing.
// We first initialize flexibleServersSPECARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersSPECARMGenerator() gopter.Gen {
	if flexibleServersSPECARMGenerator != nil {
		return flexibleServersSPECARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersSPECARM(generators)
	flexibleServersSPECARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersSPECARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersSPECARM(generators)
	flexibleServersSPECARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_SPECARM{}), generators)

	return flexibleServersSPECARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersSPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersSPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersSPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersSPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerPropertiesSpecARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuSpecARMGenerator())
}

func Test_ServerProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerPropertiesSpecARM, ServerPropertiesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerPropertiesSpecARM runs a test to see if a specific instance of ServerProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerPropertiesSpecARM(subject ServerProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerProperties_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ServerProperties_SpecARM instances for property testing - lazily instantiated by
//ServerPropertiesSpecARMGenerator()
var serverPropertiesSpecARMGenerator gopter.Gen

// ServerPropertiesSpecARMGenerator returns a generator of ServerProperties_SpecARM instances for property testing.
// We first initialize serverPropertiesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerPropertiesSpecARMGenerator() gopter.Gen {
	if serverPropertiesSpecARMGenerator != nil {
		return serverPropertiesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPropertiesSpecARM(generators)
	serverPropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPropertiesSpecARM(generators)
	AddRelatedPropertyGeneratorsForServerPropertiesSpecARM(generators)
	serverPropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_SpecARM{}), generators)

	return serverPropertiesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForServerPropertiesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerPropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["AdministratorLoginPassword"] = gen.PtrOf(gen.AlphaString())
	gens["AvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["CreateMode"] = gen.PtrOf(gen.OneConstOf(
		ServerPropertiesSpecCreateModeCreate,
		ServerPropertiesSpecCreateModeDefault,
		ServerPropertiesSpecCreateModePointInTimeRestore,
		ServerPropertiesSpecCreateModeUpdate))
	gens["PointInTimeUTC"] = gen.PtrOf(gen.AlphaString())
	gens["SourceServerResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.OneConstOf(ServerVersion_Spec11, ServerVersion_Spec12, ServerVersion_Spec13))
}

// AddRelatedPropertyGeneratorsForServerPropertiesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerPropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["Backup"] = gen.PtrOf(BackupSpecARMGenerator())
	gens["HighAvailability"] = gen.PtrOf(HighAvailabilitySpecARMGenerator())
	gens["MaintenanceWindow"] = gen.PtrOf(MaintenanceWindowSpecARMGenerator())
	gens["Network"] = gen.PtrOf(NetworkSpecARMGenerator())
	gens["Storage"] = gen.PtrOf(StorageSpecARMGenerator())
}

func Test_Sku_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuSpecARM, SkuSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuSpecARM runs a test to see if a specific instance of Sku_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuSpecARM(subject Sku_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Sku_SpecARM instances for property testing - lazily instantiated by SkuSpecARMGenerator()
var skuSpecARMGenerator gopter.Gen

// SkuSpecARMGenerator returns a generator of Sku_SpecARM instances for property testing.
func SkuSpecARMGenerator() gopter.Gen {
	if skuSpecARMGenerator != nil {
		return skuSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuSpecARM(generators)
	skuSpecARMGenerator = gen.Struct(reflect.TypeOf(Sku_SpecARM{}), generators)

	return skuSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuSpecARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
	gens["Tier"] = gen.OneConstOf(SkuSpecTierBurstable, SkuSpecTierGeneralPurpose, SkuSpecTierMemoryOptimized)
}

func Test_Backup_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Backup_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupSpecARM, BackupSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupSpecARM runs a test to see if a specific instance of Backup_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupSpecARM(subject Backup_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Backup_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Backup_SpecARM instances for property testing - lazily instantiated by BackupSpecARMGenerator()
var backupSpecARMGenerator gopter.Gen

// BackupSpecARMGenerator returns a generator of Backup_SpecARM instances for property testing.
func BackupSpecARMGenerator() gopter.Gen {
	if backupSpecARMGenerator != nil {
		return backupSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupSpecARM(generators)
	backupSpecARMGenerator = gen.Struct(reflect.TypeOf(Backup_SpecARM{}), generators)

	return backupSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForBackupSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackupSpecARM(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.OneConstOf(BackupSpecGeoRedundantBackupDisabled, BackupSpecGeoRedundantBackupEnabled))
}

func Test_HighAvailability_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HighAvailability_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHighAvailabilitySpecARM, HighAvailabilitySpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHighAvailabilitySpecARM runs a test to see if a specific instance of HighAvailability_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForHighAvailabilitySpecARM(subject HighAvailability_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HighAvailability_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of HighAvailability_SpecARM instances for property testing - lazily instantiated by
//HighAvailabilitySpecARMGenerator()
var highAvailabilitySpecARMGenerator gopter.Gen

// HighAvailabilitySpecARMGenerator returns a generator of HighAvailability_SpecARM instances for property testing.
func HighAvailabilitySpecARMGenerator() gopter.Gen {
	if highAvailabilitySpecARMGenerator != nil {
		return highAvailabilitySpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHighAvailabilitySpecARM(generators)
	highAvailabilitySpecARMGenerator = gen.Struct(reflect.TypeOf(HighAvailability_SpecARM{}), generators)

	return highAvailabilitySpecARMGenerator
}

// AddIndependentPropertyGeneratorsForHighAvailabilitySpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHighAvailabilitySpecARM(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(HighAvailabilitySpecModeDisabled, HighAvailabilitySpecModeZoneRedundant))
	gens["StandbyAvailabilityZone"] = gen.PtrOf(gen.AlphaString())
}

func Test_MaintenanceWindow_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceWindow_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceWindowSpecARM, MaintenanceWindowSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceWindowSpecARM runs a test to see if a specific instance of MaintenanceWindow_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceWindowSpecARM(subject MaintenanceWindow_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceWindow_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of MaintenanceWindow_SpecARM instances for property testing - lazily instantiated by
//MaintenanceWindowSpecARMGenerator()
var maintenanceWindowSpecARMGenerator gopter.Gen

// MaintenanceWindowSpecARMGenerator returns a generator of MaintenanceWindow_SpecARM instances for property testing.
func MaintenanceWindowSpecARMGenerator() gopter.Gen {
	if maintenanceWindowSpecARMGenerator != nil {
		return maintenanceWindowSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindowSpecARM(generators)
	maintenanceWindowSpecARMGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindow_SpecARM{}), generators)

	return maintenanceWindowSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceWindowSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceWindowSpecARM(gens map[string]gopter.Gen) {
	gens["CustomWindow"] = gen.PtrOf(gen.AlphaString())
	gens["DayOfWeek"] = gen.PtrOf(gen.Int())
	gens["StartHour"] = gen.PtrOf(gen.Int())
	gens["StartMinute"] = gen.PtrOf(gen.Int())
}

func Test_Network_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Network_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSpecARM, NetworkSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSpecARM runs a test to see if a specific instance of Network_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSpecARM(subject Network_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Network_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Network_SpecARM instances for property testing - lazily instantiated by NetworkSpecARMGenerator()
var networkSpecARMGenerator gopter.Gen

// NetworkSpecARMGenerator returns a generator of Network_SpecARM instances for property testing.
func NetworkSpecARMGenerator() gopter.Gen {
	if networkSpecARMGenerator != nil {
		return networkSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSpecARM(generators)
	networkSpecARMGenerator = gen.Struct(reflect.TypeOf(Network_SpecARM{}), generators)

	return networkSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSpecARM(gens map[string]gopter.Gen) {
	gens["DelegatedSubnetResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateDnsZoneArmResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_Storage_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Storage_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageSpecARM, StorageSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageSpecARM runs a test to see if a specific instance of Storage_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageSpecARM(subject Storage_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Storage_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Storage_SpecARM instances for property testing - lazily instantiated by StorageSpecARMGenerator()
var storageSpecARMGenerator gopter.Gen

// StorageSpecARMGenerator returns a generator of Storage_SpecARM instances for property testing.
func StorageSpecARMGenerator() gopter.Gen {
	if storageSpecARMGenerator != nil {
		return storageSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageSpecARM(generators)
	storageSpecARMGenerator = gen.Struct(reflect.TypeOf(Storage_SpecARM{}), generators)

	return storageSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageSpecARM(gens map[string]gopter.Gen) {
	gens["StorageSizeGB"] = gen.PtrOf(gen.Int())
}