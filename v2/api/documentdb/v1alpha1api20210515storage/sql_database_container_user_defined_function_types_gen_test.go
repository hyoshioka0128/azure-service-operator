// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"encoding/json"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
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

func Test_SqlDatabaseContainerUserDefinedFunction_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerUserDefinedFunction to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlDatabaseContainerUserDefinedFunction, SqlDatabaseContainerUserDefinedFunctionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlDatabaseContainerUserDefinedFunction tests if a specific instance of SqlDatabaseContainerUserDefinedFunction round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlDatabaseContainerUserDefinedFunction(subject SqlDatabaseContainerUserDefinedFunction) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210515s.SqlDatabaseContainerUserDefinedFunction
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlDatabaseContainerUserDefinedFunction
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseContainerUserDefinedFunction_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerUserDefinedFunction to SqlDatabaseContainerUserDefinedFunction via AssignPropertiesToSqlDatabaseContainerUserDefinedFunction & AssignPropertiesFromSqlDatabaseContainerUserDefinedFunction returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerUserDefinedFunction, SqlDatabaseContainerUserDefinedFunctionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerUserDefinedFunction tests if a specific instance of SqlDatabaseContainerUserDefinedFunction can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerUserDefinedFunction(subject SqlDatabaseContainerUserDefinedFunction) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlDatabaseContainerUserDefinedFunction
	err := copied.AssignPropertiesToSqlDatabaseContainerUserDefinedFunction(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerUserDefinedFunction
	err = actual.AssignPropertiesFromSqlDatabaseContainerUserDefinedFunction(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseContainerUserDefinedFunction_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerUserDefinedFunction via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerUserDefinedFunction, SqlDatabaseContainerUserDefinedFunctionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerUserDefinedFunction runs a test to see if a specific instance of SqlDatabaseContainerUserDefinedFunction round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerUserDefinedFunction(subject SqlDatabaseContainerUserDefinedFunction) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerUserDefinedFunction
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

// Generator of SqlDatabaseContainerUserDefinedFunction instances for property testing - lazily instantiated by
// SqlDatabaseContainerUserDefinedFunctionGenerator()
var sqlDatabaseContainerUserDefinedFunctionGenerator gopter.Gen

// SqlDatabaseContainerUserDefinedFunctionGenerator returns a generator of SqlDatabaseContainerUserDefinedFunction instances for property testing.
func SqlDatabaseContainerUserDefinedFunctionGenerator() gopter.Gen {
	if sqlDatabaseContainerUserDefinedFunctionGenerator != nil {
		return sqlDatabaseContainerUserDefinedFunctionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerUserDefinedFunction(generators)
	sqlDatabaseContainerUserDefinedFunctionGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerUserDefinedFunction{}), generators)

	return sqlDatabaseContainerUserDefinedFunctionGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerUserDefinedFunction is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerUserDefinedFunction(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecGenerator()
	gens["Status"] = DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSGenerator()
}

func Test_DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS to DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS via AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS & AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS, DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS tests if a specific instance of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(subject DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS
	err := copied.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS
	err = actual.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS, DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(subject DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS
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

// Generator of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS instances for property testing - lazily
// instantiated by DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSGenerator()
var databaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSGenerator != nil {
		return databaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(generators)
	databaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(generators)
	databaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS{}), generators)

	return databaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator())
}

func Test_DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec to DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec via AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec & AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec, DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec tests if a specific instance of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(subject DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec
	err := copied.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec
	err = actual.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec, DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(subject DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec
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

// Generator of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec instances for property testing - lazily
// instantiated by DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecGenerator()
var databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecGenerator != nil {
		return databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(generators)
	databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(generators)
	databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec{}), generators)

	return databaseAccountsSqlDatabasesContainersUserDefinedFunction_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunction_Spec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlUserDefinedFunctionResourceGenerator())
}

func Test_SqlUserDefinedFunctionGetProperties_Resource_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlUserDefinedFunctionGetProperties_Resource_STATUS to SqlUserDefinedFunctionGetProperties_Resource_STATUS via AssignPropertiesToSqlUserDefinedFunctionGetProperties_Resource_STATUS & AssignPropertiesFromSqlUserDefinedFunctionGetProperties_Resource_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlUserDefinedFunctionGetProperties_Resource_STATUS, SqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlUserDefinedFunctionGetProperties_Resource_STATUS tests if a specific instance of SqlUserDefinedFunctionGetProperties_Resource_STATUS can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlUserDefinedFunctionGetProperties_Resource_STATUS(subject SqlUserDefinedFunctionGetProperties_Resource_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlUserDefinedFunctionGetProperties_Resource_STATUS
	err := copied.AssignPropertiesToSqlUserDefinedFunctionGetProperties_Resource_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlUserDefinedFunctionGetProperties_Resource_STATUS
	err = actual.AssignPropertiesFromSqlUserDefinedFunctionGetProperties_Resource_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlUserDefinedFunctionGetProperties_Resource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionGetProperties_Resource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionGetProperties_Resource_STATUS, SqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionGetProperties_Resource_STATUS runs a test to see if a specific instance of SqlUserDefinedFunctionGetProperties_Resource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionGetProperties_Resource_STATUS(subject SqlUserDefinedFunctionGetProperties_Resource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionGetProperties_Resource_STATUS
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

// Generator of SqlUserDefinedFunctionGetProperties_Resource_STATUS instances for property testing - lazily instantiated
// by SqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator()
var sqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator gopter.Gen

// SqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator returns a generator of SqlUserDefinedFunctionGetProperties_Resource_STATUS instances for property testing.
func SqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator() gopter.Gen {
	if sqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator != nil {
		return sqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetProperties_Resource_STATUS(generators)
	sqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionGetProperties_Resource_STATUS{}), generators)

	return sqlUserDefinedFunctionGetProperties_Resource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

func Test_SqlUserDefinedFunctionResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlUserDefinedFunctionResource to SqlUserDefinedFunctionResource via AssignPropertiesToSqlUserDefinedFunctionResource & AssignPropertiesFromSqlUserDefinedFunctionResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlUserDefinedFunctionResource, SqlUserDefinedFunctionResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlUserDefinedFunctionResource tests if a specific instance of SqlUserDefinedFunctionResource can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlUserDefinedFunctionResource(subject SqlUserDefinedFunctionResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlUserDefinedFunctionResource
	err := copied.AssignPropertiesToSqlUserDefinedFunctionResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlUserDefinedFunctionResource
	err = actual.AssignPropertiesFromSqlUserDefinedFunctionResource(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlUserDefinedFunctionResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionResource, SqlUserDefinedFunctionResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionResource runs a test to see if a specific instance of SqlUserDefinedFunctionResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionResource(subject SqlUserDefinedFunctionResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionResource
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

// Generator of SqlUserDefinedFunctionResource instances for property testing - lazily instantiated by
// SqlUserDefinedFunctionResourceGenerator()
var sqlUserDefinedFunctionResourceGenerator gopter.Gen

// SqlUserDefinedFunctionResourceGenerator returns a generator of SqlUserDefinedFunctionResource instances for property testing.
func SqlUserDefinedFunctionResourceGenerator() gopter.Gen {
	if sqlUserDefinedFunctionResourceGenerator != nil {
		return sqlUserDefinedFunctionResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource(generators)
	sqlUserDefinedFunctionResourceGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionResource{}), generators)

	return sqlUserDefinedFunctionResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
