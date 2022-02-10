// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

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

func Test_SqlDatabaseContainerStoredProcedure_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerStoredProcedure via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure, SqlDatabaseContainerStoredProcedureGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure runs a test to see if a specific instance of SqlDatabaseContainerStoredProcedure round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure(subject SqlDatabaseContainerStoredProcedure) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerStoredProcedure
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

// Generator of SqlDatabaseContainerStoredProcedure instances for property testing - lazily instantiated by
//SqlDatabaseContainerStoredProcedureGenerator()
var sqlDatabaseContainerStoredProcedureGenerator gopter.Gen

// SqlDatabaseContainerStoredProcedureGenerator returns a generator of SqlDatabaseContainerStoredProcedure instances for property testing.
func SqlDatabaseContainerStoredProcedureGenerator() gopter.Gen {
	if sqlDatabaseContainerStoredProcedureGenerator != nil {
		return sqlDatabaseContainerStoredProcedureGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure(generators)
	sqlDatabaseContainerStoredProcedureGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerStoredProcedure{}), generators)

	return sqlDatabaseContainerStoredProcedureGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsSqlDatabasesContainersStoredProceduresSPECGenerator()
	gens["Status"] = SqlStoredProcedureCreateUpdateParametersStatusGenerator()
}

func Test_DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC, DatabaseAccountsSqlDatabasesContainersStoredProceduresSPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(subject DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC
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

// Generator of DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC instances for property testing - lazily
//instantiated by DatabaseAccountsSqlDatabasesContainersStoredProceduresSPECGenerator()
var databaseAccountsSqlDatabasesContainersStoredProceduresSPECGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersStoredProceduresSPECGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersStoredProceduresSPECGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersStoredProceduresSPECGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersStoredProceduresSPECGenerator != nil {
		return databaseAccountsSqlDatabasesContainersStoredProceduresSPECGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(generators)
	databaseAccountsSqlDatabasesContainersStoredProceduresSPECGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(generators)
	databaseAccountsSqlDatabasesContainersStoredProceduresSPECGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPEC{}), generators)

	return databaseAccountsSqlDatabasesContainersStoredProceduresSPECGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPEC(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsSpecGenerator())
	gens["Resource"] = gen.PtrOf(SqlStoredProcedureResourceSpecGenerator())
}

func Test_SqlStoredProcedureCreateUpdateParameters_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureCreateUpdateParameters_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureCreateUpdateParametersStatus, SqlStoredProcedureCreateUpdateParametersStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureCreateUpdateParametersStatus runs a test to see if a specific instance of SqlStoredProcedureCreateUpdateParameters_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureCreateUpdateParametersStatus(subject SqlStoredProcedureCreateUpdateParameters_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureCreateUpdateParameters_Status
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

// Generator of SqlStoredProcedureCreateUpdateParameters_Status instances for property testing - lazily instantiated by
//SqlStoredProcedureCreateUpdateParametersStatusGenerator()
var sqlStoredProcedureCreateUpdateParametersStatusGenerator gopter.Gen

// SqlStoredProcedureCreateUpdateParametersStatusGenerator returns a generator of SqlStoredProcedureCreateUpdateParameters_Status instances for property testing.
// We first initialize sqlStoredProcedureCreateUpdateParametersStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlStoredProcedureCreateUpdateParametersStatusGenerator() gopter.Gen {
	if sqlStoredProcedureCreateUpdateParametersStatusGenerator != nil {
		return sqlStoredProcedureCreateUpdateParametersStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatus(generators)
	sqlStoredProcedureCreateUpdateParametersStatusGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureCreateUpdateParameters_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatus(generators)
	AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatus(generators)
	sqlStoredProcedureCreateUpdateParametersStatusGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureCreateUpdateParameters_Status{}), generators)

	return sqlStoredProcedureCreateUpdateParametersStatusGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatus(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsStatusGenerator())
	gens["Resource"] = gen.PtrOf(SqlStoredProcedureResourceStatusGenerator())
}

func Test_SqlStoredProcedureResource_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureResource_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureResourceSpec, SqlStoredProcedureResourceSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureResourceSpec runs a test to see if a specific instance of SqlStoredProcedureResource_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureResourceSpec(subject SqlStoredProcedureResource_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureResource_Spec
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

// Generator of SqlStoredProcedureResource_Spec instances for property testing - lazily instantiated by
//SqlStoredProcedureResourceSpecGenerator()
var sqlStoredProcedureResourceSpecGenerator gopter.Gen

// SqlStoredProcedureResourceSpecGenerator returns a generator of SqlStoredProcedureResource_Spec instances for property testing.
func SqlStoredProcedureResourceSpecGenerator() gopter.Gen {
	if sqlStoredProcedureResourceSpecGenerator != nil {
		return sqlStoredProcedureResourceSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceSpec(generators)
	sqlStoredProcedureResourceSpecGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureResource_Spec{}), generators)

	return sqlStoredProcedureResourceSpecGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceSpec(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_SqlStoredProcedureResource_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureResource_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureResourceStatus, SqlStoredProcedureResourceStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureResourceStatus runs a test to see if a specific instance of SqlStoredProcedureResource_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureResourceStatus(subject SqlStoredProcedureResource_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureResource_Status
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

// Generator of SqlStoredProcedureResource_Status instances for property testing - lazily instantiated by
//SqlStoredProcedureResourceStatusGenerator()
var sqlStoredProcedureResourceStatusGenerator gopter.Gen

// SqlStoredProcedureResourceStatusGenerator returns a generator of SqlStoredProcedureResource_Status instances for property testing.
func SqlStoredProcedureResourceStatusGenerator() gopter.Gen {
	if sqlStoredProcedureResourceStatusGenerator != nil {
		return sqlStoredProcedureResourceStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceStatus(generators)
	sqlStoredProcedureResourceStatusGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureResource_Status{}), generators)

	return sqlStoredProcedureResourceStatusGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceStatus(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}