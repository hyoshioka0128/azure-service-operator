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

func Test_SqlStoredProcedureCreateUpdateParameters_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureCreateUpdateParameters_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureCreateUpdateParametersStatusARM, SqlStoredProcedureCreateUpdateParametersStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureCreateUpdateParametersStatusARM runs a test to see if a specific instance of SqlStoredProcedureCreateUpdateParameters_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureCreateUpdateParametersStatusARM(subject SqlStoredProcedureCreateUpdateParameters_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureCreateUpdateParameters_StatusARM
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

// Generator of SqlStoredProcedureCreateUpdateParameters_StatusARM instances for property testing - lazily instantiated
//by SqlStoredProcedureCreateUpdateParametersStatusARMGenerator()
var sqlStoredProcedureCreateUpdateParametersStatusARMGenerator gopter.Gen

// SqlStoredProcedureCreateUpdateParametersStatusARMGenerator returns a generator of SqlStoredProcedureCreateUpdateParameters_StatusARM instances for property testing.
// We first initialize sqlStoredProcedureCreateUpdateParametersStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlStoredProcedureCreateUpdateParametersStatusARMGenerator() gopter.Gen {
	if sqlStoredProcedureCreateUpdateParametersStatusARMGenerator != nil {
		return sqlStoredProcedureCreateUpdateParametersStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatusARM(generators)
	sqlStoredProcedureCreateUpdateParametersStatusARMGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureCreateUpdateParameters_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatusARM(generators)
	AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatusARM(generators)
	sqlStoredProcedureCreateUpdateParametersStatusARMGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureCreateUpdateParameters_StatusARM{}), generators)

	return sqlStoredProcedureCreateUpdateParametersStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdateParametersStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlStoredProcedureCreateUpdatePropertiesStatusARMGenerator())
}

func Test_SqlStoredProcedureCreateUpdateProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureCreateUpdateProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureCreateUpdatePropertiesStatusARM, SqlStoredProcedureCreateUpdatePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureCreateUpdatePropertiesStatusARM runs a test to see if a specific instance of SqlStoredProcedureCreateUpdateProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureCreateUpdatePropertiesStatusARM(subject SqlStoredProcedureCreateUpdateProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureCreateUpdateProperties_StatusARM
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

// Generator of SqlStoredProcedureCreateUpdateProperties_StatusARM instances for property testing - lazily instantiated
//by SqlStoredProcedureCreateUpdatePropertiesStatusARMGenerator()
var sqlStoredProcedureCreateUpdatePropertiesStatusARMGenerator gopter.Gen

// SqlStoredProcedureCreateUpdatePropertiesStatusARMGenerator returns a generator of SqlStoredProcedureCreateUpdateProperties_StatusARM instances for property testing.
func SqlStoredProcedureCreateUpdatePropertiesStatusARMGenerator() gopter.Gen {
	if sqlStoredProcedureCreateUpdatePropertiesStatusARMGenerator != nil {
		return sqlStoredProcedureCreateUpdatePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdatePropertiesStatusARM(generators)
	sqlStoredProcedureCreateUpdatePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureCreateUpdateProperties_StatusARM{}), generators)

	return sqlStoredProcedureCreateUpdatePropertiesStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdatePropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdatePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsStatusARMGenerator())
	gens["Resource"] = SqlStoredProcedureResourceStatusARMGenerator()
}

func Test_SqlStoredProcedureResource_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureResource_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureResourceStatusARM, SqlStoredProcedureResourceStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureResourceStatusARM runs a test to see if a specific instance of SqlStoredProcedureResource_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureResourceStatusARM(subject SqlStoredProcedureResource_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureResource_StatusARM
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

// Generator of SqlStoredProcedureResource_StatusARM instances for property testing - lazily instantiated by
//SqlStoredProcedureResourceStatusARMGenerator()
var sqlStoredProcedureResourceStatusARMGenerator gopter.Gen

// SqlStoredProcedureResourceStatusARMGenerator returns a generator of SqlStoredProcedureResource_StatusARM instances for property testing.
func SqlStoredProcedureResourceStatusARMGenerator() gopter.Gen {
	if sqlStoredProcedureResourceStatusARMGenerator != nil {
		return sqlStoredProcedureResourceStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceStatusARM(generators)
	sqlStoredProcedureResourceStatusARMGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureResource_StatusARM{}), generators)

	return sqlStoredProcedureResourceStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceStatusARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.AlphaString()
}