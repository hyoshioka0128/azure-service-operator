// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

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

func Test_ServersDatabase_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabase_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabase_STATUSARM, ServersDatabase_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabase_STATUSARM runs a test to see if a specific instance of ServersDatabase_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabase_STATUSARM(subject ServersDatabase_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabase_STATUSARM
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

// Generator of ServersDatabase_STATUSARM instances for property testing - lazily instantiated by
// ServersDatabase_STATUSARMGenerator()
var serversDatabase_STATUSARMGenerator gopter.Gen

// ServersDatabase_STATUSARMGenerator returns a generator of ServersDatabase_STATUSARM instances for property testing.
// We first initialize serversDatabase_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersDatabase_STATUSARMGenerator() gopter.Gen {
	if serversDatabase_STATUSARMGenerator != nil {
		return serversDatabase_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabase_STATUSARM(generators)
	serversDatabase_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ServersDatabase_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabase_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForServersDatabase_STATUSARM(generators)
	serversDatabase_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ServersDatabase_STATUSARM{}), generators)

	return serversDatabase_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForServersDatabase_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersDatabase_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServersDatabase_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabase_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabaseProperties_STATUSARMGenerator())
}

func Test_DatabaseProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseProperties_STATUSARM, DatabaseProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseProperties_STATUSARM runs a test to see if a specific instance of DatabaseProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseProperties_STATUSARM(subject DatabaseProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseProperties_STATUSARM
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

// Generator of DatabaseProperties_STATUSARM instances for property testing - lazily instantiated by
// DatabaseProperties_STATUSARMGenerator()
var databaseProperties_STATUSARMGenerator gopter.Gen

// DatabaseProperties_STATUSARMGenerator returns a generator of DatabaseProperties_STATUSARM instances for property testing.
func DatabaseProperties_STATUSARMGenerator() gopter.Gen {
	if databaseProperties_STATUSARMGenerator != nil {
		return databaseProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_STATUSARM(generators)
	databaseProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_STATUSARM{}), generators)

	return databaseProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
}