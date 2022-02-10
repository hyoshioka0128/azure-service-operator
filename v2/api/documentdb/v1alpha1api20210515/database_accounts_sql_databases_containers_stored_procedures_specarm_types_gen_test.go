// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

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

func Test_DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARM, DatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARM runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARM(subject DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPECARM
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

// Generator of DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPECARM instances for property testing - lazily
//instantiated by DatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARMGenerator()
var databaseAccountsSqlDatabasesContainersStoredProceduresSPECARMGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARMGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPECARM instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersStoredProceduresSPECARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARMGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersStoredProceduresSPECARMGenerator != nil {
		return databaseAccountsSqlDatabasesContainersStoredProceduresSPECARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARM(generators)
	databaseAccountsSqlDatabasesContainersStoredProceduresSPECARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARM(generators)
	databaseAccountsSqlDatabasesContainersStoredProceduresSPECARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedures_SPECARM{}), generators)

	return databaseAccountsSqlDatabasesContainersStoredProceduresSPECARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = SqlStoredProcedureCreateUpdatePropertiesSpecARMGenerator()
}

func Test_SqlStoredProcedureCreateUpdateProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureCreateUpdateProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureCreateUpdatePropertiesSpecARM, SqlStoredProcedureCreateUpdatePropertiesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureCreateUpdatePropertiesSpecARM runs a test to see if a specific instance of SqlStoredProcedureCreateUpdateProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureCreateUpdatePropertiesSpecARM(subject SqlStoredProcedureCreateUpdateProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureCreateUpdateProperties_SpecARM
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

// Generator of SqlStoredProcedureCreateUpdateProperties_SpecARM instances for property testing - lazily instantiated by
//SqlStoredProcedureCreateUpdatePropertiesSpecARMGenerator()
var sqlStoredProcedureCreateUpdatePropertiesSpecARMGenerator gopter.Gen

// SqlStoredProcedureCreateUpdatePropertiesSpecARMGenerator returns a generator of SqlStoredProcedureCreateUpdateProperties_SpecARM instances for property testing.
func SqlStoredProcedureCreateUpdatePropertiesSpecARMGenerator() gopter.Gen {
	if sqlStoredProcedureCreateUpdatePropertiesSpecARMGenerator != nil {
		return sqlStoredProcedureCreateUpdatePropertiesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdatePropertiesSpecARM(generators)
	sqlStoredProcedureCreateUpdatePropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureCreateUpdateProperties_SpecARM{}), generators)

	return sqlStoredProcedureCreateUpdatePropertiesSpecARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdatePropertiesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdatePropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsSpecARMGenerator())
	gens["Resource"] = SqlStoredProcedureResourceSpecARMGenerator()
}

func Test_SqlStoredProcedureResource_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureResource_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureResourceSpecARM, SqlStoredProcedureResourceSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureResourceSpecARM runs a test to see if a specific instance of SqlStoredProcedureResource_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureResourceSpecARM(subject SqlStoredProcedureResource_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureResource_SpecARM
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

// Generator of SqlStoredProcedureResource_SpecARM instances for property testing - lazily instantiated by
//SqlStoredProcedureResourceSpecARMGenerator()
var sqlStoredProcedureResourceSpecARMGenerator gopter.Gen

// SqlStoredProcedureResourceSpecARMGenerator returns a generator of SqlStoredProcedureResource_SpecARM instances for property testing.
func SqlStoredProcedureResourceSpecARMGenerator() gopter.Gen {
	if sqlStoredProcedureResourceSpecARMGenerator != nil {
		return sqlStoredProcedureResourceSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceSpecARM(generators)
	sqlStoredProcedureResourceSpecARMGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureResource_SpecARM{}), generators)

	return sqlStoredProcedureResourceSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceSpecARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.AlphaString()
}