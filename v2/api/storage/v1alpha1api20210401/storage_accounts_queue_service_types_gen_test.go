// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401storage"
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
	var hub v1alpha1api20210401storage.StorageAccountsQueueService
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
	var other v1alpha1api20210401storage.StorageAccountsQueueService
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

	//Check for a match
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
//StorageAccountsQueueServiceGenerator()
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
	gens["Spec"] = StorageAccountsQueueServicesSPECGenerator()
	gens["Status"] = QueueServicePropertiesStatusGenerator()
}

func Test_QueueServiceProperties_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from QueueServiceProperties_Status to QueueServiceProperties_Status via AssignPropertiesToQueueServicePropertiesStatus & AssignPropertiesFromQueueServicePropertiesStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForQueueServicePropertiesStatus, QueueServicePropertiesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForQueueServicePropertiesStatus tests if a specific instance of QueueServiceProperties_Status can be assigned to v1alpha1api20210401storage and back losslessly
func RunPropertyAssignmentTestForQueueServicePropertiesStatus(subject QueueServiceProperties_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210401storage.QueueServiceProperties_Status
	err := copied.AssignPropertiesToQueueServicePropertiesStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual QueueServiceProperties_Status
	err = actual.AssignPropertiesFromQueueServicePropertiesStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_QueueServiceProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of QueueServiceProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForQueueServicePropertiesStatus, QueueServicePropertiesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForQueueServicePropertiesStatus runs a test to see if a specific instance of QueueServiceProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForQueueServicePropertiesStatus(subject QueueServiceProperties_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual QueueServiceProperties_Status
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

// Generator of QueueServiceProperties_Status instances for property testing - lazily instantiated by
//QueueServicePropertiesStatusGenerator()
var queueServicePropertiesStatusGenerator gopter.Gen

// QueueServicePropertiesStatusGenerator returns a generator of QueueServiceProperties_Status instances for property testing.
// We first initialize queueServicePropertiesStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func QueueServicePropertiesStatusGenerator() gopter.Gen {
	if queueServicePropertiesStatusGenerator != nil {
		return queueServicePropertiesStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForQueueServicePropertiesStatus(generators)
	queueServicePropertiesStatusGenerator = gen.Struct(reflect.TypeOf(QueueServiceProperties_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForQueueServicePropertiesStatus(generators)
	AddRelatedPropertyGeneratorsForQueueServicePropertiesStatus(generators)
	queueServicePropertiesStatusGenerator = gen.Struct(reflect.TypeOf(QueueServiceProperties_Status{}), generators)

	return queueServicePropertiesStatusGenerator
}

// AddIndependentPropertyGeneratorsForQueueServicePropertiesStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForQueueServicePropertiesStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForQueueServicePropertiesStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForQueueServicePropertiesStatus(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRulesStatusGenerator())
}

func Test_StorageAccountsQueueServices_SPEC_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueServices_SPEC to StorageAccountsQueueServices_SPEC via AssignPropertiesToStorageAccountsQueueServicesSPEC & AssignPropertiesFromStorageAccountsQueueServicesSPEC returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsQueueServicesSPEC, StorageAccountsQueueServicesSPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsQueueServicesSPEC tests if a specific instance of StorageAccountsQueueServices_SPEC can be assigned to v1alpha1api20210401storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsQueueServicesSPEC(subject StorageAccountsQueueServices_SPEC) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210401storage.StorageAccountsQueueServices_SPEC
	err := copied.AssignPropertiesToStorageAccountsQueueServicesSPEC(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsQueueServices_SPEC
	err = actual.AssignPropertiesFromStorageAccountsQueueServicesSPEC(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccountsQueueServices_SPEC_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueServices_SPEC via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueServicesSPEC, StorageAccountsQueueServicesSPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueServicesSPEC runs a test to see if a specific instance of StorageAccountsQueueServices_SPEC round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueServicesSPEC(subject StorageAccountsQueueServices_SPEC) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueServices_SPEC
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

// Generator of StorageAccountsQueueServices_SPEC instances for property testing - lazily instantiated by
//StorageAccountsQueueServicesSPECGenerator()
var storageAccountsQueueServicesSPECGenerator gopter.Gen

// StorageAccountsQueueServicesSPECGenerator returns a generator of StorageAccountsQueueServices_SPEC instances for property testing.
// We first initialize storageAccountsQueueServicesSPECGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsQueueServicesSPECGenerator() gopter.Gen {
	if storageAccountsQueueServicesSPECGenerator != nil {
		return storageAccountsQueueServicesSPECGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesSPEC(generators)
	storageAccountsQueueServicesSPECGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServices_SPEC{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesSPEC(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesSPEC(generators)
	storageAccountsQueueServicesSPECGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServices_SPEC{}), generators)

	return storageAccountsQueueServicesSPECGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesSPEC is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesSPEC(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesSPEC is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesSPEC(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRulesSpecGenerator())
}
