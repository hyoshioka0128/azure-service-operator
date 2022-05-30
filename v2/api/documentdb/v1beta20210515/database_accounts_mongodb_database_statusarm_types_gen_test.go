// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

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

func Test_DatabaseAccountsMongodbDatabase_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabase_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabase_STATUSARM, DatabaseAccountsMongodbDatabase_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabase_STATUSARM runs a test to see if a specific instance of DatabaseAccountsMongodbDatabase_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabase_STATUSARM(subject DatabaseAccountsMongodbDatabase_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabase_STATUSARM
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

// Generator of DatabaseAccountsMongodbDatabase_STATUSARM instances for property testing - lazily instantiated by
// DatabaseAccountsMongodbDatabase_STATUSARMGenerator()
var databaseAccountsMongodbDatabase_STATUSARMGenerator gopter.Gen

// DatabaseAccountsMongodbDatabase_STATUSARMGenerator returns a generator of DatabaseAccountsMongodbDatabase_STATUSARM instances for property testing.
// We first initialize databaseAccountsMongodbDatabase_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabase_STATUSARMGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabase_STATUSARMGenerator != nil {
		return databaseAccountsMongodbDatabase_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUSARM(generators)
	databaseAccountsMongodbDatabase_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabase_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUSARM(generators)
	databaseAccountsMongodbDatabase_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabase_STATUSARM{}), generators)

	return databaseAccountsMongodbDatabase_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabase_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MongoDBDatabaseGetProperties_STATUSARMGenerator())
}

func Test_MongoDBDatabaseGetProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseGetProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseGetProperties_STATUSARM, MongoDBDatabaseGetProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseGetProperties_STATUSARM runs a test to see if a specific instance of MongoDBDatabaseGetProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseGetProperties_STATUSARM(subject MongoDBDatabaseGetProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseGetProperties_STATUSARM
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

// Generator of MongoDBDatabaseGetProperties_STATUSARM instances for property testing - lazily instantiated by
// MongoDBDatabaseGetProperties_STATUSARMGenerator()
var mongoDBDatabaseGetProperties_STATUSARMGenerator gopter.Gen

// MongoDBDatabaseGetProperties_STATUSARMGenerator returns a generator of MongoDBDatabaseGetProperties_STATUSARM instances for property testing.
func MongoDBDatabaseGetProperties_STATUSARMGenerator() gopter.Gen {
	if mongoDBDatabaseGetProperties_STATUSARMGenerator != nil {
		return mongoDBDatabaseGetProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_STATUSARM(generators)
	mongoDBDatabaseGetProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetProperties_STATUSARM{}), generators)

	return mongoDBDatabaseGetProperties_STATUSARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResource_STATUSARMGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBDatabaseGetProperties_Resource_STATUSARMGenerator())
}

func Test_MongoDBDatabaseGetProperties_Resource_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseGetProperties_Resource_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseGetProperties_Resource_STATUSARM, MongoDBDatabaseGetProperties_Resource_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseGetProperties_Resource_STATUSARM runs a test to see if a specific instance of MongoDBDatabaseGetProperties_Resource_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseGetProperties_Resource_STATUSARM(subject MongoDBDatabaseGetProperties_Resource_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseGetProperties_Resource_STATUSARM
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

// Generator of MongoDBDatabaseGetProperties_Resource_STATUSARM instances for property testing - lazily instantiated by
// MongoDBDatabaseGetProperties_Resource_STATUSARMGenerator()
var mongoDBDatabaseGetProperties_Resource_STATUSARMGenerator gopter.Gen

// MongoDBDatabaseGetProperties_Resource_STATUSARMGenerator returns a generator of MongoDBDatabaseGetProperties_Resource_STATUSARM instances for property testing.
func MongoDBDatabaseGetProperties_Resource_STATUSARMGenerator() gopter.Gen {
	if mongoDBDatabaseGetProperties_Resource_STATUSARMGenerator != nil {
		return mongoDBDatabaseGetProperties_Resource_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUSARM(generators)
	mongoDBDatabaseGetProperties_Resource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetProperties_Resource_STATUSARM{}), generators)

	return mongoDBDatabaseGetProperties_Resource_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["_Etag"] = gen.PtrOf(gen.AlphaString())
	gens["_Rid"] = gen.PtrOf(gen.AlphaString())
	gens["_Ts"] = gen.PtrOf(gen.Float64())
}

func Test_OptionsResource_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OptionsResource_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOptionsResource_STATUSARM, OptionsResource_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOptionsResource_STATUSARM runs a test to see if a specific instance of OptionsResource_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForOptionsResource_STATUSARM(subject OptionsResource_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OptionsResource_STATUSARM
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

// Generator of OptionsResource_STATUSARM instances for property testing - lazily instantiated by
// OptionsResource_STATUSARMGenerator()
var optionsResource_STATUSARMGenerator gopter.Gen

// OptionsResource_STATUSARMGenerator returns a generator of OptionsResource_STATUSARM instances for property testing.
// We first initialize optionsResource_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func OptionsResource_STATUSARMGenerator() gopter.Gen {
	if optionsResource_STATUSARMGenerator != nil {
		return optionsResource_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOptionsResource_STATUSARM(generators)
	optionsResource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(OptionsResource_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOptionsResource_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForOptionsResource_STATUSARM(generators)
	optionsResource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(OptionsResource_STATUSARM{}), generators)

	return optionsResource_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForOptionsResource_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForOptionsResource_STATUSARM(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForOptionsResource_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForOptionsResource_STATUSARM(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettings_STATUSARMGenerator())
}

func Test_AutoscaleSettings_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettings_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettings_STATUSARM, AutoscaleSettings_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettings_STATUSARM runs a test to see if a specific instance of AutoscaleSettings_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettings_STATUSARM(subject AutoscaleSettings_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettings_STATUSARM
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

// Generator of AutoscaleSettings_STATUSARM instances for property testing - lazily instantiated by
// AutoscaleSettings_STATUSARMGenerator()
var autoscaleSettings_STATUSARMGenerator gopter.Gen

// AutoscaleSettings_STATUSARMGenerator returns a generator of AutoscaleSettings_STATUSARM instances for property testing.
func AutoscaleSettings_STATUSARMGenerator() gopter.Gen {
	if autoscaleSettings_STATUSARMGenerator != nil {
		return autoscaleSettings_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettings_STATUSARM(generators)
	autoscaleSettings_STATUSARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettings_STATUSARM{}), generators)

	return autoscaleSettings_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettings_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettings_STATUSARM(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}
