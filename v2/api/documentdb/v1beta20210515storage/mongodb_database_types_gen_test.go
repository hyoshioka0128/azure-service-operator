// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515storage

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

func Test_MongodbDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabase, MongodbDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabase runs a test to see if a specific instance of MongodbDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabase(subject MongodbDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabase
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

// Generator of MongodbDatabase instances for property testing - lazily instantiated by MongodbDatabaseGenerator()
var mongodbDatabaseGenerator gopter.Gen

// MongodbDatabaseGenerator returns a generator of MongodbDatabase instances for property testing.
func MongodbDatabaseGenerator() gopter.Gen {
	if mongodbDatabaseGenerator != nil {
		return mongodbDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongodbDatabase(generators)
	mongodbDatabaseGenerator = gen.Struct(reflect.TypeOf(MongodbDatabase{}), generators)

	return mongodbDatabaseGenerator
}

// AddRelatedPropertyGeneratorsForMongodbDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsMongodbDatabase_SpecGenerator()
	gens["Status"] = DatabaseAccountsMongodbDatabase_STATUSGenerator()
}

func Test_DatabaseAccountsMongodbDatabase_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabase_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabase_STATUS, DatabaseAccountsMongodbDatabase_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabase_STATUS runs a test to see if a specific instance of DatabaseAccountsMongodbDatabase_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabase_STATUS(subject DatabaseAccountsMongodbDatabase_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabase_STATUS
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

// Generator of DatabaseAccountsMongodbDatabase_STATUS instances for property testing - lazily instantiated by
// DatabaseAccountsMongodbDatabase_STATUSGenerator()
var databaseAccountsMongodbDatabase_STATUSGenerator gopter.Gen

// DatabaseAccountsMongodbDatabase_STATUSGenerator returns a generator of DatabaseAccountsMongodbDatabase_STATUS instances for property testing.
// We first initialize databaseAccountsMongodbDatabase_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabase_STATUSGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabase_STATUSGenerator != nil {
		return databaseAccountsMongodbDatabase_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUS(generators)
	databaseAccountsMongodbDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabase_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUS(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUS(generators)
	databaseAccountsMongodbDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabase_STATUS{}), generators)

	return databaseAccountsMongodbDatabase_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptions_STATUSGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBDatabaseResource_STATUSGenerator())
}

func Test_DatabaseAccountsMongodbDatabase_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabase_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabase_Spec, DatabaseAccountsMongodbDatabase_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabase_Spec runs a test to see if a specific instance of DatabaseAccountsMongodbDatabase_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabase_Spec(subject DatabaseAccountsMongodbDatabase_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabase_Spec
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

// Generator of DatabaseAccountsMongodbDatabase_Spec instances for property testing - lazily instantiated by
// DatabaseAccountsMongodbDatabase_SpecGenerator()
var databaseAccountsMongodbDatabase_SpecGenerator gopter.Gen

// DatabaseAccountsMongodbDatabase_SpecGenerator returns a generator of DatabaseAccountsMongodbDatabase_Spec instances for property testing.
// We first initialize databaseAccountsMongodbDatabase_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabase_SpecGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabase_SpecGenerator != nil {
		return databaseAccountsMongodbDatabase_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabase_Spec(generators)
	databaseAccountsMongodbDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabase_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabase_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabase_Spec(generators)
	databaseAccountsMongodbDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabase_Spec{}), generators)

	return databaseAccountsMongodbDatabase_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabase_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabase_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabase_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabase_Spec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBDatabaseResourceGenerator())
}

func Test_CreateUpdateOptions_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CreateUpdateOptions via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCreateUpdateOptions, CreateUpdateOptionsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCreateUpdateOptions runs a test to see if a specific instance of CreateUpdateOptions round trips to JSON and back losslessly
func RunJSONSerializationTestForCreateUpdateOptions(subject CreateUpdateOptions) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CreateUpdateOptions
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

// Generator of CreateUpdateOptions instances for property testing - lazily instantiated by
// CreateUpdateOptionsGenerator()
var createUpdateOptionsGenerator gopter.Gen

// CreateUpdateOptionsGenerator returns a generator of CreateUpdateOptions instances for property testing.
// We first initialize createUpdateOptionsGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CreateUpdateOptionsGenerator() gopter.Gen {
	if createUpdateOptionsGenerator != nil {
		return createUpdateOptionsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptions(generators)
	createUpdateOptionsGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptions{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptions(generators)
	AddRelatedPropertyGeneratorsForCreateUpdateOptions(generators)
	createUpdateOptionsGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptions{}), generators)

	return createUpdateOptionsGenerator
}

// AddIndependentPropertyGeneratorsForCreateUpdateOptions is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCreateUpdateOptions(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForCreateUpdateOptions is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCreateUpdateOptions(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsGenerator())
}

func Test_CreateUpdateOptions_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CreateUpdateOptions_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCreateUpdateOptions_STATUS, CreateUpdateOptions_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCreateUpdateOptions_STATUS runs a test to see if a specific instance of CreateUpdateOptions_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCreateUpdateOptions_STATUS(subject CreateUpdateOptions_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CreateUpdateOptions_STATUS
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

// Generator of CreateUpdateOptions_STATUS instances for property testing - lazily instantiated by
// CreateUpdateOptions_STATUSGenerator()
var createUpdateOptions_STATUSGenerator gopter.Gen

// CreateUpdateOptions_STATUSGenerator returns a generator of CreateUpdateOptions_STATUS instances for property testing.
// We first initialize createUpdateOptions_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CreateUpdateOptions_STATUSGenerator() gopter.Gen {
	if createUpdateOptions_STATUSGenerator != nil {
		return createUpdateOptions_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptions_STATUS(generators)
	createUpdateOptions_STATUSGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptions_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptions_STATUS(generators)
	AddRelatedPropertyGeneratorsForCreateUpdateOptions_STATUS(generators)
	createUpdateOptions_STATUSGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptions_STATUS{}), generators)

	return createUpdateOptions_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCreateUpdateOptions_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCreateUpdateOptions_STATUS(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForCreateUpdateOptions_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCreateUpdateOptions_STATUS(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettings_STATUSGenerator())
}

func Test_MongoDBDatabaseResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseResource, MongoDBDatabaseResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseResource runs a test to see if a specific instance of MongoDBDatabaseResource round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseResource(subject MongoDBDatabaseResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseResource
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

// Generator of MongoDBDatabaseResource instances for property testing - lazily instantiated by
// MongoDBDatabaseResourceGenerator()
var mongoDBDatabaseResourceGenerator gopter.Gen

// MongoDBDatabaseResourceGenerator returns a generator of MongoDBDatabaseResource instances for property testing.
func MongoDBDatabaseResourceGenerator() gopter.Gen {
	if mongoDBDatabaseResourceGenerator != nil {
		return mongoDBDatabaseResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseResource(generators)
	mongoDBDatabaseResourceGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseResource{}), generators)

	return mongoDBDatabaseResourceGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseResource(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_MongoDBDatabaseResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseResource_STATUS, MongoDBDatabaseResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseResource_STATUS runs a test to see if a specific instance of MongoDBDatabaseResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseResource_STATUS(subject MongoDBDatabaseResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseResource_STATUS
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

// Generator of MongoDBDatabaseResource_STATUS instances for property testing - lazily instantiated by
// MongoDBDatabaseResource_STATUSGenerator()
var mongoDBDatabaseResource_STATUSGenerator gopter.Gen

// MongoDBDatabaseResource_STATUSGenerator returns a generator of MongoDBDatabaseResource_STATUS instances for property testing.
func MongoDBDatabaseResource_STATUSGenerator() gopter.Gen {
	if mongoDBDatabaseResource_STATUSGenerator != nil {
		return mongoDBDatabaseResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseResource_STATUS(generators)
	mongoDBDatabaseResource_STATUSGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseResource_STATUS{}), generators)

	return mongoDBDatabaseResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseResource_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_AutoscaleSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettings, AutoscaleSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettings runs a test to see if a specific instance of AutoscaleSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettings(subject AutoscaleSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettings
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

// Generator of AutoscaleSettings instances for property testing - lazily instantiated by AutoscaleSettingsGenerator()
var autoscaleSettingsGenerator gopter.Gen

// AutoscaleSettingsGenerator returns a generator of AutoscaleSettings instances for property testing.
func AutoscaleSettingsGenerator() gopter.Gen {
	if autoscaleSettingsGenerator != nil {
		return autoscaleSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettings(generators)
	autoscaleSettingsGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettings{}), generators)

	return autoscaleSettingsGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettings(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}

func Test_AutoscaleSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettings_STATUS, AutoscaleSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettings_STATUS runs a test to see if a specific instance of AutoscaleSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettings_STATUS(subject AutoscaleSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettings_STATUS
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

// Generator of AutoscaleSettings_STATUS instances for property testing - lazily instantiated by
// AutoscaleSettings_STATUSGenerator()
var autoscaleSettings_STATUSGenerator gopter.Gen

// AutoscaleSettings_STATUSGenerator returns a generator of AutoscaleSettings_STATUS instances for property testing.
func AutoscaleSettings_STATUSGenerator() gopter.Gen {
	if autoscaleSettings_STATUSGenerator != nil {
		return autoscaleSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettings_STATUS(generators)
	autoscaleSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettings_STATUS{}), generators)

	return autoscaleSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettings_STATUS(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}
