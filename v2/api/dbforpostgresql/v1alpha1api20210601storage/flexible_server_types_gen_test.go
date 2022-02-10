// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601storage

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

func Test_FlexibleServer_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServer via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServer, FlexibleServerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServer runs a test to see if a specific instance of FlexibleServer round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServer(subject FlexibleServer) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServer
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

// Generator of FlexibleServer instances for property testing - lazily instantiated by FlexibleServerGenerator()
var flexibleServerGenerator gopter.Gen

// FlexibleServerGenerator returns a generator of FlexibleServer instances for property testing.
func FlexibleServerGenerator() gopter.Gen {
	if flexibleServerGenerator != nil {
		return flexibleServerGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServer(generators)
	flexibleServerGenerator = gen.Struct(reflect.TypeOf(FlexibleServer{}), generators)

	return flexibleServerGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServer is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServer(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServersSPECGenerator()
	gens["Status"] = ServerStatusGenerator()
}

func Test_FlexibleServers_SPEC_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_SPEC via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersSPEC, FlexibleServersSPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersSPEC runs a test to see if a specific instance of FlexibleServers_SPEC round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersSPEC(subject FlexibleServers_SPEC) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_SPEC
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

// Generator of FlexibleServers_SPEC instances for property testing - lazily instantiated by
//FlexibleServersSPECGenerator()
var flexibleServersSPECGenerator gopter.Gen

// FlexibleServersSPECGenerator returns a generator of FlexibleServers_SPEC instances for property testing.
// We first initialize flexibleServersSPECGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersSPECGenerator() gopter.Gen {
	if flexibleServersSPECGenerator != nil {
		return flexibleServersSPECGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersSPEC(generators)
	flexibleServersSPECGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_SPEC{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersSPEC(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersSPEC(generators)
	flexibleServersSPECGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_SPEC{}), generators)

	return flexibleServersSPECGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersSPEC is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersSPEC(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["AvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["CreateMode"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PointInTimeUTC"] = gen.PtrOf(gen.AlphaString())
	gens["PropertiesTags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersSPEC is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersSPEC(gens map[string]gopter.Gen) {
	gens["Backup"] = gen.PtrOf(BackupSpecGenerator())
	gens["HighAvailability"] = gen.PtrOf(HighAvailabilitySpecGenerator())
	gens["MaintenanceWindow"] = gen.PtrOf(MaintenanceWindowSpecGenerator())
	gens["Network"] = gen.PtrOf(NetworkSpecGenerator())
	gens["Sku"] = gen.PtrOf(SkuSpecGenerator())
	gens["Storage"] = gen.PtrOf(StorageSpecGenerator())
}

func Test_Server_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Server_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerStatus, ServerStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerStatus runs a test to see if a specific instance of Server_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForServerStatus(subject Server_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Server_Status
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

// Generator of Server_Status instances for property testing - lazily instantiated by ServerStatusGenerator()
var serverStatusGenerator gopter.Gen

// ServerStatusGenerator returns a generator of Server_Status instances for property testing.
// We first initialize serverStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerStatusGenerator() gopter.Gen {
	if serverStatusGenerator != nil {
		return serverStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerStatus(generators)
	serverStatusGenerator = gen.Struct(reflect.TypeOf(Server_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerStatus(generators)
	AddRelatedPropertyGeneratorsForServerStatus(generators)
	serverStatusGenerator = gen.Struct(reflect.TypeOf(Server_Status{}), generators)

	return serverStatusGenerator
}

// AddIndependentPropertyGeneratorsForServerStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerStatus(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["AvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["CreateMode"] = gen.PtrOf(gen.AlphaString())
	gens["FullyQualifiedDomainName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PointInTimeUTC"] = gen.PtrOf(gen.AlphaString())
	gens["PropertiesTags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["SourceServerResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServerStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerStatus(gens map[string]gopter.Gen) {
	gens["Backup"] = gen.PtrOf(BackupStatusGenerator())
	gens["HighAvailability"] = gen.PtrOf(HighAvailabilityStatusGenerator())
	gens["MaintenanceWindow"] = gen.PtrOf(MaintenanceWindowStatusGenerator())
	gens["Network"] = gen.PtrOf(NetworkStatusGenerator())
	gens["Sku"] = gen.PtrOf(SkuStatusGenerator())
	gens["Storage"] = gen.PtrOf(StorageStatusGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusGenerator())
}

func Test_Backup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Backup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupSpec, BackupSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupSpec runs a test to see if a specific instance of Backup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupSpec(subject Backup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Backup_Spec
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

// Generator of Backup_Spec instances for property testing - lazily instantiated by BackupSpecGenerator()
var backupSpecGenerator gopter.Gen

// BackupSpecGenerator returns a generator of Backup_Spec instances for property testing.
func BackupSpecGenerator() gopter.Gen {
	if backupSpecGenerator != nil {
		return backupSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupSpec(generators)
	backupSpecGenerator = gen.Struct(reflect.TypeOf(Backup_Spec{}), generators)

	return backupSpecGenerator
}

// AddIndependentPropertyGeneratorsForBackupSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackupSpec(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.AlphaString())
}

func Test_Backup_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Backup_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupStatus, BackupStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupStatus runs a test to see if a specific instance of Backup_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupStatus(subject Backup_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Backup_Status
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

// Generator of Backup_Status instances for property testing - lazily instantiated by BackupStatusGenerator()
var backupStatusGenerator gopter.Gen

// BackupStatusGenerator returns a generator of Backup_Status instances for property testing.
func BackupStatusGenerator() gopter.Gen {
	if backupStatusGenerator != nil {
		return backupStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupStatus(generators)
	backupStatusGenerator = gen.Struct(reflect.TypeOf(Backup_Status{}), generators)

	return backupStatusGenerator
}

// AddIndependentPropertyGeneratorsForBackupStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackupStatus(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["EarliestRestoreDate"] = gen.PtrOf(gen.AlphaString())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.AlphaString())
}

func Test_HighAvailability_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HighAvailability_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHighAvailabilitySpec, HighAvailabilitySpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHighAvailabilitySpec runs a test to see if a specific instance of HighAvailability_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForHighAvailabilitySpec(subject HighAvailability_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HighAvailability_Spec
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

// Generator of HighAvailability_Spec instances for property testing - lazily instantiated by
//HighAvailabilitySpecGenerator()
var highAvailabilitySpecGenerator gopter.Gen

// HighAvailabilitySpecGenerator returns a generator of HighAvailability_Spec instances for property testing.
func HighAvailabilitySpecGenerator() gopter.Gen {
	if highAvailabilitySpecGenerator != nil {
		return highAvailabilitySpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHighAvailabilitySpec(generators)
	highAvailabilitySpecGenerator = gen.Struct(reflect.TypeOf(HighAvailability_Spec{}), generators)

	return highAvailabilitySpecGenerator
}

// AddIndependentPropertyGeneratorsForHighAvailabilitySpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHighAvailabilitySpec(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.AlphaString())
	gens["StandbyAvailabilityZone"] = gen.PtrOf(gen.AlphaString())
}

func Test_HighAvailability_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HighAvailability_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHighAvailabilityStatus, HighAvailabilityStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHighAvailabilityStatus runs a test to see if a specific instance of HighAvailability_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForHighAvailabilityStatus(subject HighAvailability_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HighAvailability_Status
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

// Generator of HighAvailability_Status instances for property testing - lazily instantiated by
//HighAvailabilityStatusGenerator()
var highAvailabilityStatusGenerator gopter.Gen

// HighAvailabilityStatusGenerator returns a generator of HighAvailability_Status instances for property testing.
func HighAvailabilityStatusGenerator() gopter.Gen {
	if highAvailabilityStatusGenerator != nil {
		return highAvailabilityStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHighAvailabilityStatus(generators)
	highAvailabilityStatusGenerator = gen.Struct(reflect.TypeOf(HighAvailability_Status{}), generators)

	return highAvailabilityStatusGenerator
}

// AddIndependentPropertyGeneratorsForHighAvailabilityStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHighAvailabilityStatus(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.AlphaString())
	gens["StandbyAvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.AlphaString())
}

func Test_MaintenanceWindow_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceWindow_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceWindowSpec, MaintenanceWindowSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceWindowSpec runs a test to see if a specific instance of MaintenanceWindow_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceWindowSpec(subject MaintenanceWindow_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceWindow_Spec
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

// Generator of MaintenanceWindow_Spec instances for property testing - lazily instantiated by
//MaintenanceWindowSpecGenerator()
var maintenanceWindowSpecGenerator gopter.Gen

// MaintenanceWindowSpecGenerator returns a generator of MaintenanceWindow_Spec instances for property testing.
func MaintenanceWindowSpecGenerator() gopter.Gen {
	if maintenanceWindowSpecGenerator != nil {
		return maintenanceWindowSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindowSpec(generators)
	maintenanceWindowSpecGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindow_Spec{}), generators)

	return maintenanceWindowSpecGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceWindowSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceWindowSpec(gens map[string]gopter.Gen) {
	gens["CustomWindow"] = gen.PtrOf(gen.AlphaString())
	gens["DayOfWeek"] = gen.PtrOf(gen.Int())
	gens["StartHour"] = gen.PtrOf(gen.Int())
	gens["StartMinute"] = gen.PtrOf(gen.Int())
}

func Test_MaintenanceWindow_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceWindow_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceWindowStatus, MaintenanceWindowStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceWindowStatus runs a test to see if a specific instance of MaintenanceWindow_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceWindowStatus(subject MaintenanceWindow_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceWindow_Status
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

// Generator of MaintenanceWindow_Status instances for property testing - lazily instantiated by
//MaintenanceWindowStatusGenerator()
var maintenanceWindowStatusGenerator gopter.Gen

// MaintenanceWindowStatusGenerator returns a generator of MaintenanceWindow_Status instances for property testing.
func MaintenanceWindowStatusGenerator() gopter.Gen {
	if maintenanceWindowStatusGenerator != nil {
		return maintenanceWindowStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindowStatus(generators)
	maintenanceWindowStatusGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindow_Status{}), generators)

	return maintenanceWindowStatusGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceWindowStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceWindowStatus(gens map[string]gopter.Gen) {
	gens["CustomWindow"] = gen.PtrOf(gen.AlphaString())
	gens["DayOfWeek"] = gen.PtrOf(gen.Int())
	gens["StartHour"] = gen.PtrOf(gen.Int())
	gens["StartMinute"] = gen.PtrOf(gen.Int())
}

func Test_Network_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Network_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSpec, NetworkSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSpec runs a test to see if a specific instance of Network_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSpec(subject Network_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Network_Spec
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

// Generator of Network_Spec instances for property testing - lazily instantiated by NetworkSpecGenerator()
var networkSpecGenerator gopter.Gen

// NetworkSpecGenerator returns a generator of Network_Spec instances for property testing.
func NetworkSpecGenerator() gopter.Gen {
	if networkSpecGenerator != nil {
		return networkSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	networkSpecGenerator = gen.Struct(reflect.TypeOf(Network_Spec{}), generators)

	return networkSpecGenerator
}

func Test_Network_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Network_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkStatus, NetworkStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkStatus runs a test to see if a specific instance of Network_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkStatus(subject Network_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Network_Status
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

// Generator of Network_Status instances for property testing - lazily instantiated by NetworkStatusGenerator()
var networkStatusGenerator gopter.Gen

// NetworkStatusGenerator returns a generator of Network_Status instances for property testing.
func NetworkStatusGenerator() gopter.Gen {
	if networkStatusGenerator != nil {
		return networkStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkStatus(generators)
	networkStatusGenerator = gen.Struct(reflect.TypeOf(Network_Status{}), generators)

	return networkStatusGenerator
}

// AddIndependentPropertyGeneratorsForNetworkStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkStatus(gens map[string]gopter.Gen) {
	gens["DelegatedSubnetResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateDnsZoneArmResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuSpec, SkuSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuSpec runs a test to see if a specific instance of Sku_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuSpec(subject Sku_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_Spec
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

// Generator of Sku_Spec instances for property testing - lazily instantiated by SkuSpecGenerator()
var skuSpecGenerator gopter.Gen

// SkuSpecGenerator returns a generator of Sku_Spec instances for property testing.
func SkuSpecGenerator() gopter.Gen {
	if skuSpecGenerator != nil {
		return skuSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuSpec(generators)
	skuSpecGenerator = gen.Struct(reflect.TypeOf(Sku_Spec{}), generators)

	return skuSpecGenerator
}

// AddIndependentPropertyGeneratorsForSkuSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuSpec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuStatus, SkuStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuStatus runs a test to see if a specific instance of Sku_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuStatus(subject Sku_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_Status
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

// Generator of Sku_Status instances for property testing - lazily instantiated by SkuStatusGenerator()
var skuStatusGenerator gopter.Gen

// SkuStatusGenerator returns a generator of Sku_Status instances for property testing.
func SkuStatusGenerator() gopter.Gen {
	if skuStatusGenerator != nil {
		return skuStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuStatus(generators)
	skuStatusGenerator = gen.Struct(reflect.TypeOf(Sku_Status{}), generators)

	return skuStatusGenerator
}

// AddIndependentPropertyGeneratorsForSkuStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuStatus(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_Storage_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Storage_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageSpec, StorageSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageSpec runs a test to see if a specific instance of Storage_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageSpec(subject Storage_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Storage_Spec
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

// Generator of Storage_Spec instances for property testing - lazily instantiated by StorageSpecGenerator()
var storageSpecGenerator gopter.Gen

// StorageSpecGenerator returns a generator of Storage_Spec instances for property testing.
func StorageSpecGenerator() gopter.Gen {
	if storageSpecGenerator != nil {
		return storageSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageSpec(generators)
	storageSpecGenerator = gen.Struct(reflect.TypeOf(Storage_Spec{}), generators)

	return storageSpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageSpec(gens map[string]gopter.Gen) {
	gens["StorageSizeGB"] = gen.PtrOf(gen.Int())
}

func Test_Storage_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Storage_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageStatus, StorageStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageStatus runs a test to see if a specific instance of Storage_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageStatus(subject Storage_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Storage_Status
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

// Generator of Storage_Status instances for property testing - lazily instantiated by StorageStatusGenerator()
var storageStatusGenerator gopter.Gen

// StorageStatusGenerator returns a generator of Storage_Status instances for property testing.
func StorageStatusGenerator() gopter.Gen {
	if storageStatusGenerator != nil {
		return storageStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageStatus(generators)
	storageStatusGenerator = gen.Struct(reflect.TypeOf(Storage_Status{}), generators)

	return storageStatusGenerator
}

// AddIndependentPropertyGeneratorsForStorageStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageStatus(gens map[string]gopter.Gen) {
	gens["StorageSizeGB"] = gen.PtrOf(gen.Int())
}

func Test_SystemData_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemDataStatus, SystemDataStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemDataStatus runs a test to see if a specific instance of SystemData_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemDataStatus(subject SystemData_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_Status
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

// Generator of SystemData_Status instances for property testing - lazily instantiated by SystemDataStatusGenerator()
var systemDataStatusGenerator gopter.Gen

// SystemDataStatusGenerator returns a generator of SystemData_Status instances for property testing.
func SystemDataStatusGenerator() gopter.Gen {
	if systemDataStatusGenerator != nil {
		return systemDataStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemDataStatus(generators)
	systemDataStatusGenerator = gen.Struct(reflect.TypeOf(SystemData_Status{}), generators)

	return systemDataStatusGenerator
}

// AddIndependentPropertyGeneratorsForSystemDataStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemDataStatus(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.AlphaString())
}
