// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import (
	"encoding/json"
	alpha20210401s "github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401storage"
	v20210401s "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401storage"
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

func Test_StorageAccountsQueueService_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueService to hub returns original",
		prop.ForAll(RunResourceConversionTestForStorageAccountsQueueService, StorageAccountsQueueServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForStorageAccountsQueueService tests if a specific instance of StorageAccountsQueueService round trips to the hub storage version and back losslessly
func RunResourceConversionTestForStorageAccountsQueueService(subject StorageAccountsQueueService) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210401s.StorageAccountsQueueService
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual StorageAccountsQueueService
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

func Test_StorageAccountsQueueService_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueService to StorageAccountsQueueService via AssignPropertiesToStorageAccountsQueueService & AssignPropertiesFromStorageAccountsQueueService returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsQueueService, StorageAccountsQueueServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsQueueService tests if a specific instance of StorageAccountsQueueService can be assigned to v1alpha1api20210401storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsQueueService(subject StorageAccountsQueueService) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210401s.StorageAccountsQueueService
	err := copied.AssignPropertiesToStorageAccountsQueueService(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsQueueService
	err = actual.AssignPropertiesFromStorageAccountsQueueService(&other)
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

func Test_StorageAccountsQueueService_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueService via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueService, StorageAccountsQueueServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueService runs a test to see if a specific instance of StorageAccountsQueueService round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueService(subject StorageAccountsQueueService) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueService
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

// Generator of StorageAccountsQueueService instances for property testing - lazily instantiated by
// StorageAccountsQueueServiceGenerator()
var storageAccountsQueueServiceGenerator gopter.Gen

// StorageAccountsQueueServiceGenerator returns a generator of StorageAccountsQueueService instances for property testing.
func StorageAccountsQueueServiceGenerator() gopter.Gen {
	if storageAccountsQueueServiceGenerator != nil {
		return storageAccountsQueueServiceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsQueueService(generators)
	storageAccountsQueueServiceGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueService{}), generators)

	return storageAccountsQueueServiceGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsQueueService is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsQueueService(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccountsQueueService_SpecGenerator()
	gens["Status"] = StorageAccountsQueueService_STATUSGenerator()
}

func Test_StorageAccountsQueueService_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueService_STATUS to StorageAccountsQueueService_STATUS via AssignPropertiesToStorageAccountsQueueService_STATUS & AssignPropertiesFromStorageAccountsQueueService_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsQueueService_STATUS, StorageAccountsQueueService_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsQueueService_STATUS tests if a specific instance of StorageAccountsQueueService_STATUS can be assigned to v1alpha1api20210401storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsQueueService_STATUS(subject StorageAccountsQueueService_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210401s.StorageAccountsQueueService_STATUS
	err := copied.AssignPropertiesToStorageAccountsQueueService_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsQueueService_STATUS
	err = actual.AssignPropertiesFromStorageAccountsQueueService_STATUS(&other)
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

func Test_StorageAccountsQueueService_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueService_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueService_STATUS, StorageAccountsQueueService_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueService_STATUS runs a test to see if a specific instance of StorageAccountsQueueService_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueService_STATUS(subject StorageAccountsQueueService_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueService_STATUS
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

// Generator of StorageAccountsQueueService_STATUS instances for property testing - lazily instantiated by
// StorageAccountsQueueService_STATUSGenerator()
var storageAccountsQueueService_STATUSGenerator gopter.Gen

// StorageAccountsQueueService_STATUSGenerator returns a generator of StorageAccountsQueueService_STATUS instances for property testing.
// We first initialize storageAccountsQueueService_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsQueueService_STATUSGenerator() gopter.Gen {
	if storageAccountsQueueService_STATUSGenerator != nil {
		return storageAccountsQueueService_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueService_STATUS(generators)
	storageAccountsQueueService_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueService_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueService_STATUS(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsQueueService_STATUS(generators)
	storageAccountsQueueService_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueService_STATUS{}), generators)

	return storageAccountsQueueService_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsQueueService_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsQueueService_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsQueueService_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsQueueService_STATUS(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRules_STATUSGenerator())
}

func Test_StorageAccountsQueueService_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueService_Spec to StorageAccountsQueueService_Spec via AssignPropertiesToStorageAccountsQueueService_Spec & AssignPropertiesFromStorageAccountsQueueService_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsQueueService_Spec, StorageAccountsQueueService_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsQueueService_Spec tests if a specific instance of StorageAccountsQueueService_Spec can be assigned to v1alpha1api20210401storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsQueueService_Spec(subject StorageAccountsQueueService_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210401s.StorageAccountsQueueService_Spec
	err := copied.AssignPropertiesToStorageAccountsQueueService_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsQueueService_Spec
	err = actual.AssignPropertiesFromStorageAccountsQueueService_Spec(&other)
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

func Test_StorageAccountsQueueService_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueService_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueService_Spec, StorageAccountsQueueService_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueService_Spec runs a test to see if a specific instance of StorageAccountsQueueService_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueService_Spec(subject StorageAccountsQueueService_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueService_Spec
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

// Generator of StorageAccountsQueueService_Spec instances for property testing - lazily instantiated by
// StorageAccountsQueueService_SpecGenerator()
var storageAccountsQueueService_SpecGenerator gopter.Gen

// StorageAccountsQueueService_SpecGenerator returns a generator of StorageAccountsQueueService_Spec instances for property testing.
// We first initialize storageAccountsQueueService_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsQueueService_SpecGenerator() gopter.Gen {
	if storageAccountsQueueService_SpecGenerator != nil {
		return storageAccountsQueueService_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueService_Spec(generators)
	storageAccountsQueueService_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueService_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueService_Spec(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsQueueService_Spec(generators)
	storageAccountsQueueService_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueService_Spec{}), generators)

	return storageAccountsQueueService_SpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsQueueService_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsQueueService_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsQueueService_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsQueueService_Spec(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRulesGenerator())
}
