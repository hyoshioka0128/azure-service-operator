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

func Test_SqlDatabaseContainerStoredProcedure_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerStoredProcedure to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlDatabaseContainerStoredProcedure, SqlDatabaseContainerStoredProcedureGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlDatabaseContainerStoredProcedure tests if a specific instance of SqlDatabaseContainerStoredProcedure round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlDatabaseContainerStoredProcedure(subject SqlDatabaseContainerStoredProcedure) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210515s.SqlDatabaseContainerStoredProcedure
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlDatabaseContainerStoredProcedure
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

func Test_SqlDatabaseContainerStoredProcedure_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerStoredProcedure to SqlDatabaseContainerStoredProcedure via AssignPropertiesToSqlDatabaseContainerStoredProcedure & AssignPropertiesFromSqlDatabaseContainerStoredProcedure returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure, SqlDatabaseContainerStoredProcedureGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure tests if a specific instance of SqlDatabaseContainerStoredProcedure can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure(subject SqlDatabaseContainerStoredProcedure) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlDatabaseContainerStoredProcedure
	err := copied.AssignPropertiesToSqlDatabaseContainerStoredProcedure(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerStoredProcedure
	err = actual.AssignPropertiesFromSqlDatabaseContainerStoredProcedure(&other)
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
// SqlDatabaseContainerStoredProcedureGenerator()
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
	gens["Spec"] = DatabaseAccountsSqlDatabasesContainersStoredProcedure_SpecGenerator()
	gens["Status"] = DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUSGenerator()
}

func Test_DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS to DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS via AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS & AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS, DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS tests if a specific instance of DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(subject DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS
	err := copied.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS
	err = actual.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(&other)
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

func Test_DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS, DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(subject DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS
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

// Generator of DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS instances for property testing - lazily
// instantiated by DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUSGenerator()
var databaseAccountsSqlDatabasesContainersStoredProcedure_STATUSGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUSGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersStoredProcedure_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUSGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersStoredProcedure_STATUSGenerator != nil {
		return databaseAccountsSqlDatabasesContainersStoredProcedure_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(generators)
	databaseAccountsSqlDatabasesContainersStoredProcedure_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(generators)
	databaseAccountsSqlDatabasesContainersStoredProcedure_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS{}), generators)

	return databaseAccountsSqlDatabasesContainersStoredProcedure_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlStoredProcedureGetProperties_Resource_STATUSGenerator())
}

func Test_DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec to DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec via AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec & AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec, DatabaseAccountsSqlDatabasesContainersStoredProcedure_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec tests if a specific instance of DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(subject DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec
	err := copied.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec
	err = actual.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(&other)
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

func Test_DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec, DatabaseAccountsSqlDatabasesContainersStoredProcedure_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(subject DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec
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

// Generator of DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec instances for property testing - lazily
// instantiated by DatabaseAccountsSqlDatabasesContainersStoredProcedure_SpecGenerator()
var databaseAccountsSqlDatabasesContainersStoredProcedure_SpecGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersStoredProcedure_SpecGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersStoredProcedure_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersStoredProcedure_SpecGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersStoredProcedure_SpecGenerator != nil {
		return databaseAccountsSqlDatabasesContainersStoredProcedure_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(generators)
	databaseAccountsSqlDatabasesContainersStoredProcedure_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(generators)
	databaseAccountsSqlDatabasesContainersStoredProcedure_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec{}), generators)

	return databaseAccountsSqlDatabasesContainersStoredProcedure_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlStoredProcedureResourceGenerator())
}

func Test_SqlStoredProcedureGetProperties_Resource_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlStoredProcedureGetProperties_Resource_STATUS to SqlStoredProcedureGetProperties_Resource_STATUS via AssignPropertiesToSqlStoredProcedureGetProperties_Resource_STATUS & AssignPropertiesFromSqlStoredProcedureGetProperties_Resource_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlStoredProcedureGetProperties_Resource_STATUS, SqlStoredProcedureGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlStoredProcedureGetProperties_Resource_STATUS tests if a specific instance of SqlStoredProcedureGetProperties_Resource_STATUS can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlStoredProcedureGetProperties_Resource_STATUS(subject SqlStoredProcedureGetProperties_Resource_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlStoredProcedureGetProperties_Resource_STATUS
	err := copied.AssignPropertiesToSqlStoredProcedureGetProperties_Resource_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlStoredProcedureGetProperties_Resource_STATUS
	err = actual.AssignPropertiesFromSqlStoredProcedureGetProperties_Resource_STATUS(&other)
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

func Test_SqlStoredProcedureGetProperties_Resource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureGetProperties_Resource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureGetProperties_Resource_STATUS, SqlStoredProcedureGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureGetProperties_Resource_STATUS runs a test to see if a specific instance of SqlStoredProcedureGetProperties_Resource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureGetProperties_Resource_STATUS(subject SqlStoredProcedureGetProperties_Resource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureGetProperties_Resource_STATUS
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

// Generator of SqlStoredProcedureGetProperties_Resource_STATUS instances for property testing - lazily instantiated by
// SqlStoredProcedureGetProperties_Resource_STATUSGenerator()
var sqlStoredProcedureGetProperties_Resource_STATUSGenerator gopter.Gen

// SqlStoredProcedureGetProperties_Resource_STATUSGenerator returns a generator of SqlStoredProcedureGetProperties_Resource_STATUS instances for property testing.
func SqlStoredProcedureGetProperties_Resource_STATUSGenerator() gopter.Gen {
	if sqlStoredProcedureGetProperties_Resource_STATUSGenerator != nil {
		return sqlStoredProcedureGetProperties_Resource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureGetProperties_Resource_STATUS(generators)
	sqlStoredProcedureGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureGetProperties_Resource_STATUS{}), generators)

	return sqlStoredProcedureGetProperties_Resource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["_Etag"] = gen.PtrOf(gen.AlphaString())
	gens["_Rid"] = gen.PtrOf(gen.AlphaString())
	gens["_Ts"] = gen.PtrOf(gen.Float64())
}

func Test_SqlStoredProcedureResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlStoredProcedureResource to SqlStoredProcedureResource via AssignPropertiesToSqlStoredProcedureResource & AssignPropertiesFromSqlStoredProcedureResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlStoredProcedureResource, SqlStoredProcedureResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlStoredProcedureResource tests if a specific instance of SqlStoredProcedureResource can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlStoredProcedureResource(subject SqlStoredProcedureResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlStoredProcedureResource
	err := copied.AssignPropertiesToSqlStoredProcedureResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlStoredProcedureResource
	err = actual.AssignPropertiesFromSqlStoredProcedureResource(&other)
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

func Test_SqlStoredProcedureResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureResource, SqlStoredProcedureResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureResource runs a test to see if a specific instance of SqlStoredProcedureResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureResource(subject SqlStoredProcedureResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureResource
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

// Generator of SqlStoredProcedureResource instances for property testing - lazily instantiated by
// SqlStoredProcedureResourceGenerator()
var sqlStoredProcedureResourceGenerator gopter.Gen

// SqlStoredProcedureResourceGenerator returns a generator of SqlStoredProcedureResource instances for property testing.
func SqlStoredProcedureResourceGenerator() gopter.Gen {
	if sqlStoredProcedureResourceGenerator != nil {
		return sqlStoredProcedureResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureResource(generators)
	sqlStoredProcedureResourceGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureResource{}), generators)

	return sqlStoredProcedureResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
