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

func Test_DatabaseAccountsSqlDatabasesContainers_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainers_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersSPECARM, DatabaseAccountsSqlDatabasesContainersSPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersSPECARM runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainers_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersSPECARM(subject DatabaseAccountsSqlDatabasesContainers_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainers_SPECARM
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

// Generator of DatabaseAccountsSqlDatabasesContainers_SPECARM instances for property testing - lazily instantiated by
//DatabaseAccountsSqlDatabasesContainersSPECARMGenerator()
var databaseAccountsSqlDatabasesContainersSPECARMGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersSPECARMGenerator returns a generator of DatabaseAccountsSqlDatabasesContainers_SPECARM instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersSPECARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersSPECARMGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersSPECARMGenerator != nil {
		return databaseAccountsSqlDatabasesContainersSPECARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSPECARM(generators)
	databaseAccountsSqlDatabasesContainersSPECARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainers_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSPECARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSPECARM(generators)
	databaseAccountsSqlDatabasesContainersSPECARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainers_SPECARM{}), generators)

	return databaseAccountsSqlDatabasesContainersSPECARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersSPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = SqlContainerCreateUpdatePropertiesSpecARMGenerator()
}

func Test_SqlContainerCreateUpdateProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainerCreateUpdateProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainerCreateUpdatePropertiesSpecARM, SqlContainerCreateUpdatePropertiesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainerCreateUpdatePropertiesSpecARM runs a test to see if a specific instance of SqlContainerCreateUpdateProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainerCreateUpdatePropertiesSpecARM(subject SqlContainerCreateUpdateProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainerCreateUpdateProperties_SpecARM
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

// Generator of SqlContainerCreateUpdateProperties_SpecARM instances for property testing - lazily instantiated by
//SqlContainerCreateUpdatePropertiesSpecARMGenerator()
var sqlContainerCreateUpdatePropertiesSpecARMGenerator gopter.Gen

// SqlContainerCreateUpdatePropertiesSpecARMGenerator returns a generator of SqlContainerCreateUpdateProperties_SpecARM instances for property testing.
func SqlContainerCreateUpdatePropertiesSpecARMGenerator() gopter.Gen {
	if sqlContainerCreateUpdatePropertiesSpecARMGenerator != nil {
		return sqlContainerCreateUpdatePropertiesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlContainerCreateUpdatePropertiesSpecARM(generators)
	sqlContainerCreateUpdatePropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerCreateUpdateProperties_SpecARM{}), generators)

	return sqlContainerCreateUpdatePropertiesSpecARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlContainerCreateUpdatePropertiesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainerCreateUpdatePropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsSpecARMGenerator())
	gens["Resource"] = SqlContainerResourceSpecARMGenerator()
}

func Test_SqlContainerResource_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainerResource_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainerResourceSpecARM, SqlContainerResourceSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainerResourceSpecARM runs a test to see if a specific instance of SqlContainerResource_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainerResourceSpecARM(subject SqlContainerResource_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainerResource_SpecARM
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

// Generator of SqlContainerResource_SpecARM instances for property testing - lazily instantiated by
//SqlContainerResourceSpecARMGenerator()
var sqlContainerResourceSpecARMGenerator gopter.Gen

// SqlContainerResourceSpecARMGenerator returns a generator of SqlContainerResource_SpecARM instances for property testing.
// We first initialize sqlContainerResourceSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlContainerResourceSpecARMGenerator() gopter.Gen {
	if sqlContainerResourceSpecARMGenerator != nil {
		return sqlContainerResourceSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerResourceSpecARM(generators)
	sqlContainerResourceSpecARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerResource_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerResourceSpecARM(generators)
	AddRelatedPropertyGeneratorsForSqlContainerResourceSpecARM(generators)
	sqlContainerResourceSpecARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerResource_SpecARM{}), generators)

	return sqlContainerResourceSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlContainerResourceSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlContainerResourceSpecARM(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["DefaultTtl"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForSqlContainerResourceSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainerResourceSpecARM(gens map[string]gopter.Gen) {
	gens["ConflictResolutionPolicy"] = gen.PtrOf(ConflictResolutionPolicySpecARMGenerator())
	gens["IndexingPolicy"] = gen.PtrOf(IndexingPolicySpecARMGenerator())
	gens["PartitionKey"] = gen.PtrOf(ContainerPartitionKeySpecARMGenerator())
	gens["UniqueKeyPolicy"] = gen.PtrOf(UniqueKeyPolicySpecARMGenerator())
}

func Test_ConflictResolutionPolicy_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConflictResolutionPolicy_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConflictResolutionPolicySpecARM, ConflictResolutionPolicySpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConflictResolutionPolicySpecARM runs a test to see if a specific instance of ConflictResolutionPolicy_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConflictResolutionPolicySpecARM(subject ConflictResolutionPolicy_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConflictResolutionPolicy_SpecARM
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

// Generator of ConflictResolutionPolicy_SpecARM instances for property testing - lazily instantiated by
//ConflictResolutionPolicySpecARMGenerator()
var conflictResolutionPolicySpecARMGenerator gopter.Gen

// ConflictResolutionPolicySpecARMGenerator returns a generator of ConflictResolutionPolicy_SpecARM instances for property testing.
func ConflictResolutionPolicySpecARMGenerator() gopter.Gen {
	if conflictResolutionPolicySpecARMGenerator != nil {
		return conflictResolutionPolicySpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConflictResolutionPolicySpecARM(generators)
	conflictResolutionPolicySpecARMGenerator = gen.Struct(reflect.TypeOf(ConflictResolutionPolicy_SpecARM{}), generators)

	return conflictResolutionPolicySpecARMGenerator
}

// AddIndependentPropertyGeneratorsForConflictResolutionPolicySpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConflictResolutionPolicySpecARM(gens map[string]gopter.Gen) {
	gens["ConflictResolutionPath"] = gen.PtrOf(gen.AlphaString())
	gens["ConflictResolutionProcedure"] = gen.PtrOf(gen.AlphaString())
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(ConflictResolutionPolicySpecModeCustom, ConflictResolutionPolicySpecModeLastWriterWins))
}

func Test_ContainerPartitionKey_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ContainerPartitionKey_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForContainerPartitionKeySpecARM, ContainerPartitionKeySpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForContainerPartitionKeySpecARM runs a test to see if a specific instance of ContainerPartitionKey_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForContainerPartitionKeySpecARM(subject ContainerPartitionKey_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ContainerPartitionKey_SpecARM
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

// Generator of ContainerPartitionKey_SpecARM instances for property testing - lazily instantiated by
//ContainerPartitionKeySpecARMGenerator()
var containerPartitionKeySpecARMGenerator gopter.Gen

// ContainerPartitionKeySpecARMGenerator returns a generator of ContainerPartitionKey_SpecARM instances for property testing.
func ContainerPartitionKeySpecARMGenerator() gopter.Gen {
	if containerPartitionKeySpecARMGenerator != nil {
		return containerPartitionKeySpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerPartitionKeySpecARM(generators)
	containerPartitionKeySpecARMGenerator = gen.Struct(reflect.TypeOf(ContainerPartitionKey_SpecARM{}), generators)

	return containerPartitionKeySpecARMGenerator
}

// AddIndependentPropertyGeneratorsForContainerPartitionKeySpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForContainerPartitionKeySpecARM(gens map[string]gopter.Gen) {
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(ContainerPartitionKeySpecKindHash, ContainerPartitionKeySpecKindMultiHash, ContainerPartitionKeySpecKindRange))
	gens["Paths"] = gen.SliceOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.Int())
}

func Test_IndexingPolicy_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IndexingPolicy_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIndexingPolicySpecARM, IndexingPolicySpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIndexingPolicySpecARM runs a test to see if a specific instance of IndexingPolicy_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIndexingPolicySpecARM(subject IndexingPolicy_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IndexingPolicy_SpecARM
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

// Generator of IndexingPolicy_SpecARM instances for property testing - lazily instantiated by
//IndexingPolicySpecARMGenerator()
var indexingPolicySpecARMGenerator gopter.Gen

// IndexingPolicySpecARMGenerator returns a generator of IndexingPolicy_SpecARM instances for property testing.
// We first initialize indexingPolicySpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IndexingPolicySpecARMGenerator() gopter.Gen {
	if indexingPolicySpecARMGenerator != nil {
		return indexingPolicySpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexingPolicySpecARM(generators)
	indexingPolicySpecARMGenerator = gen.Struct(reflect.TypeOf(IndexingPolicy_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexingPolicySpecARM(generators)
	AddRelatedPropertyGeneratorsForIndexingPolicySpecARM(generators)
	indexingPolicySpecARMGenerator = gen.Struct(reflect.TypeOf(IndexingPolicy_SpecARM{}), generators)

	return indexingPolicySpecARMGenerator
}

// AddIndependentPropertyGeneratorsForIndexingPolicySpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIndexingPolicySpecARM(gens map[string]gopter.Gen) {
	gens["Automatic"] = gen.PtrOf(gen.Bool())
	gens["IndexingMode"] = gen.PtrOf(gen.OneConstOf(IndexingPolicySpecIndexingModeConsistent, IndexingPolicySpecIndexingModeLazy, IndexingPolicySpecIndexingModeNone))
}

// AddRelatedPropertyGeneratorsForIndexingPolicySpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIndexingPolicySpecARM(gens map[string]gopter.Gen) {
	gens["CompositeIndexes"] = gen.SliceOf(gen.SliceOf(CompositePathSpecARMGenerator()))
	gens["ExcludedPaths"] = gen.SliceOf(ExcludedPathSpecARMGenerator())
	gens["IncludedPaths"] = gen.SliceOf(IncludedPathSpecARMGenerator())
	gens["SpatialIndexes"] = gen.SliceOf(SpatialSpecSpecARMGenerator())
}

func Test_UniqueKeyPolicy_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UniqueKeyPolicy_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUniqueKeyPolicySpecARM, UniqueKeyPolicySpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUniqueKeyPolicySpecARM runs a test to see if a specific instance of UniqueKeyPolicy_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUniqueKeyPolicySpecARM(subject UniqueKeyPolicy_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UniqueKeyPolicy_SpecARM
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

// Generator of UniqueKeyPolicy_SpecARM instances for property testing - lazily instantiated by
//UniqueKeyPolicySpecARMGenerator()
var uniqueKeyPolicySpecARMGenerator gopter.Gen

// UniqueKeyPolicySpecARMGenerator returns a generator of UniqueKeyPolicy_SpecARM instances for property testing.
func UniqueKeyPolicySpecARMGenerator() gopter.Gen {
	if uniqueKeyPolicySpecARMGenerator != nil {
		return uniqueKeyPolicySpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForUniqueKeyPolicySpecARM(generators)
	uniqueKeyPolicySpecARMGenerator = gen.Struct(reflect.TypeOf(UniqueKeyPolicy_SpecARM{}), generators)

	return uniqueKeyPolicySpecARMGenerator
}

// AddRelatedPropertyGeneratorsForUniqueKeyPolicySpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUniqueKeyPolicySpecARM(gens map[string]gopter.Gen) {
	gens["UniqueKeys"] = gen.SliceOf(UniqueKeySpecARMGenerator())
}

func Test_CompositePath_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CompositePath_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCompositePathSpecARM, CompositePathSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCompositePathSpecARM runs a test to see if a specific instance of CompositePath_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCompositePathSpecARM(subject CompositePath_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CompositePath_SpecARM
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

// Generator of CompositePath_SpecARM instances for property testing - lazily instantiated by
//CompositePathSpecARMGenerator()
var compositePathSpecARMGenerator gopter.Gen

// CompositePathSpecARMGenerator returns a generator of CompositePath_SpecARM instances for property testing.
func CompositePathSpecARMGenerator() gopter.Gen {
	if compositePathSpecARMGenerator != nil {
		return compositePathSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCompositePathSpecARM(generators)
	compositePathSpecARMGenerator = gen.Struct(reflect.TypeOf(CompositePath_SpecARM{}), generators)

	return compositePathSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForCompositePathSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCompositePathSpecARM(gens map[string]gopter.Gen) {
	gens["Order"] = gen.PtrOf(gen.OneConstOf(CompositePathSpecOrderAscending, CompositePathSpecOrderDescending))
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

func Test_ExcludedPath_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExcludedPath_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExcludedPathSpecARM, ExcludedPathSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExcludedPathSpecARM runs a test to see if a specific instance of ExcludedPath_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExcludedPathSpecARM(subject ExcludedPath_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExcludedPath_SpecARM
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

// Generator of ExcludedPath_SpecARM instances for property testing - lazily instantiated by
//ExcludedPathSpecARMGenerator()
var excludedPathSpecARMGenerator gopter.Gen

// ExcludedPathSpecARMGenerator returns a generator of ExcludedPath_SpecARM instances for property testing.
func ExcludedPathSpecARMGenerator() gopter.Gen {
	if excludedPathSpecARMGenerator != nil {
		return excludedPathSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExcludedPathSpecARM(generators)
	excludedPathSpecARMGenerator = gen.Struct(reflect.TypeOf(ExcludedPath_SpecARM{}), generators)

	return excludedPathSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForExcludedPathSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExcludedPathSpecARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

func Test_IncludedPath_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IncludedPath_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIncludedPathSpecARM, IncludedPathSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIncludedPathSpecARM runs a test to see if a specific instance of IncludedPath_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIncludedPathSpecARM(subject IncludedPath_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IncludedPath_SpecARM
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

// Generator of IncludedPath_SpecARM instances for property testing - lazily instantiated by
//IncludedPathSpecARMGenerator()
var includedPathSpecARMGenerator gopter.Gen

// IncludedPathSpecARMGenerator returns a generator of IncludedPath_SpecARM instances for property testing.
// We first initialize includedPathSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IncludedPathSpecARMGenerator() gopter.Gen {
	if includedPathSpecARMGenerator != nil {
		return includedPathSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIncludedPathSpecARM(generators)
	includedPathSpecARMGenerator = gen.Struct(reflect.TypeOf(IncludedPath_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIncludedPathSpecARM(generators)
	AddRelatedPropertyGeneratorsForIncludedPathSpecARM(generators)
	includedPathSpecARMGenerator = gen.Struct(reflect.TypeOf(IncludedPath_SpecARM{}), generators)

	return includedPathSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForIncludedPathSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIncludedPathSpecARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForIncludedPathSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIncludedPathSpecARM(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(IndexesSpecARMGenerator())
}

func Test_SpatialSpec_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SpatialSpec_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSpatialSpecSpecARM, SpatialSpecSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSpatialSpecSpecARM runs a test to see if a specific instance of SpatialSpec_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSpatialSpecSpecARM(subject SpatialSpec_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SpatialSpec_SpecARM
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

// Generator of SpatialSpec_SpecARM instances for property testing - lazily instantiated by SpatialSpecSpecARMGenerator()
var spatialSpecSpecARMGenerator gopter.Gen

// SpatialSpecSpecARMGenerator returns a generator of SpatialSpec_SpecARM instances for property testing.
func SpatialSpecSpecARMGenerator() gopter.Gen {
	if spatialSpecSpecARMGenerator != nil {
		return spatialSpecSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSpatialSpecSpecARM(generators)
	spatialSpecSpecARMGenerator = gen.Struct(reflect.TypeOf(SpatialSpec_SpecARM{}), generators)

	return spatialSpecSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForSpatialSpecSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSpatialSpecSpecARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
	gens["Types"] = gen.SliceOf(gen.OneConstOf(
		SpatialType_SpecLineString,
		SpatialType_SpecMultiPolygon,
		SpatialType_SpecPoint,
		SpatialType_SpecPolygon))
}

func Test_UniqueKey_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UniqueKey_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUniqueKeySpecARM, UniqueKeySpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUniqueKeySpecARM runs a test to see if a specific instance of UniqueKey_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUniqueKeySpecARM(subject UniqueKey_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UniqueKey_SpecARM
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

// Generator of UniqueKey_SpecARM instances for property testing - lazily instantiated by UniqueKeySpecARMGenerator()
var uniqueKeySpecARMGenerator gopter.Gen

// UniqueKeySpecARMGenerator returns a generator of UniqueKey_SpecARM instances for property testing.
func UniqueKeySpecARMGenerator() gopter.Gen {
	if uniqueKeySpecARMGenerator != nil {
		return uniqueKeySpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUniqueKeySpecARM(generators)
	uniqueKeySpecARMGenerator = gen.Struct(reflect.TypeOf(UniqueKey_SpecARM{}), generators)

	return uniqueKeySpecARMGenerator
}

// AddIndependentPropertyGeneratorsForUniqueKeySpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUniqueKeySpecARM(gens map[string]gopter.Gen) {
	gens["Paths"] = gen.SliceOf(gen.AlphaString())
}

func Test_Indexes_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Indexes_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIndexesSpecARM, IndexesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIndexesSpecARM runs a test to see if a specific instance of Indexes_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIndexesSpecARM(subject Indexes_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Indexes_SpecARM
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

// Generator of Indexes_SpecARM instances for property testing - lazily instantiated by IndexesSpecARMGenerator()
var indexesSpecARMGenerator gopter.Gen

// IndexesSpecARMGenerator returns a generator of Indexes_SpecARM instances for property testing.
func IndexesSpecARMGenerator() gopter.Gen {
	if indexesSpecARMGenerator != nil {
		return indexesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexesSpecARM(generators)
	indexesSpecARMGenerator = gen.Struct(reflect.TypeOf(Indexes_SpecARM{}), generators)

	return indexesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForIndexesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIndexesSpecARM(gens map[string]gopter.Gen) {
	gens["DataType"] = gen.PtrOf(gen.OneConstOf(
		IndexesSpecDataTypeLineString,
		IndexesSpecDataTypeMultiPolygon,
		IndexesSpecDataTypeNumber,
		IndexesSpecDataTypePoint,
		IndexesSpecDataTypePolygon,
		IndexesSpecDataTypeString))
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(IndexesSpecKindHash, IndexesSpecKindRange, IndexesSpecKindSpatial))
	gens["Precision"] = gen.PtrOf(gen.Int())
}
