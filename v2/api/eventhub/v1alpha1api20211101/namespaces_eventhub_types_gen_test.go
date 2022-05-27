// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import (
	"encoding/json"
	alpha20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1alpha1api20211101storage"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1beta20211101storage"
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

func Test_NamespacesEventhub_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhub to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesEventhub, NamespacesEventhubGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesEventhub tests if a specific instance of NamespacesEventhub round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesEventhub(subject NamespacesEventhub) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.NamespacesEventhub
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesEventhub
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

func Test_NamespacesEventhub_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhub to NamespacesEventhub via AssignPropertiesToNamespacesEventhub & AssignPropertiesFromNamespacesEventhub returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhub, NamespacesEventhubGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhub tests if a specific instance of NamespacesEventhub can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhub(subject NamespacesEventhub) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.NamespacesEventhub
	err := copied.AssignPropertiesToNamespacesEventhub(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhub
	err = actual.AssignPropertiesFromNamespacesEventhub(&other)
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

func Test_NamespacesEventhub_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhub via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhub, NamespacesEventhubGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhub runs a test to see if a specific instance of NamespacesEventhub round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhub(subject NamespacesEventhub) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhub
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

// Generator of NamespacesEventhub instances for property testing - lazily instantiated by NamespacesEventhubGenerator()
var namespacesEventhubGenerator gopter.Gen

// NamespacesEventhubGenerator returns a generator of NamespacesEventhub instances for property testing.
func NamespacesEventhubGenerator() gopter.Gen {
	if namespacesEventhubGenerator != nil {
		return namespacesEventhubGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesEventhub(generators)
	namespacesEventhubGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhub{}), generators)

	return namespacesEventhubGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesEventhub is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhub(gens map[string]gopter.Gen) {
	gens["Spec"] = NamespacesEventhub_SpecGenerator()
	gens["Status"] = NamespacesEventhub_STATUSGenerator()
}

func Test_NamespacesEventhub_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhub_STATUS to NamespacesEventhub_STATUS via AssignPropertiesToNamespacesEventhub_STATUS & AssignPropertiesFromNamespacesEventhub_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhub_STATUS, NamespacesEventhub_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhub_STATUS tests if a specific instance of NamespacesEventhub_STATUS can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhub_STATUS(subject NamespacesEventhub_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.NamespacesEventhub_STATUS
	err := copied.AssignPropertiesToNamespacesEventhub_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhub_STATUS
	err = actual.AssignPropertiesFromNamespacesEventhub_STATUS(&other)
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

func Test_NamespacesEventhub_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhub_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhub_STATUS, NamespacesEventhub_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhub_STATUS runs a test to see if a specific instance of NamespacesEventhub_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhub_STATUS(subject NamespacesEventhub_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhub_STATUS
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

// Generator of NamespacesEventhub_STATUS instances for property testing - lazily instantiated by
// NamespacesEventhub_STATUSGenerator()
var namespacesEventhub_STATUSGenerator gopter.Gen

// NamespacesEventhub_STATUSGenerator returns a generator of NamespacesEventhub_STATUS instances for property testing.
// We first initialize namespacesEventhub_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesEventhub_STATUSGenerator() gopter.Gen {
	if namespacesEventhub_STATUSGenerator != nil {
		return namespacesEventhub_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhub_STATUS(generators)
	namespacesEventhub_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhub_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhub_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespacesEventhub_STATUS(generators)
	namespacesEventhub_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhub_STATUS{}), generators)

	return namespacesEventhub_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhub_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhub_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MessageRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
	gens["PartitionIds"] = gen.SliceOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		NamespacesEventhub_Properties_Status_STATUSActive,
		NamespacesEventhub_Properties_Status_STATUSCreating,
		NamespacesEventhub_Properties_Status_STATUSDeleting,
		NamespacesEventhub_Properties_Status_STATUSDisabled,
		NamespacesEventhub_Properties_Status_STATUSReceiveDisabled,
		NamespacesEventhub_Properties_Status_STATUSRenaming,
		NamespacesEventhub_Properties_Status_STATUSRestoring,
		NamespacesEventhub_Properties_Status_STATUSSendDisabled,
		NamespacesEventhub_Properties_Status_STATUSUnknown))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesEventhub_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhub_STATUS(gens map[string]gopter.Gen) {
	gens["CaptureDescription"] = gen.PtrOf(CaptureDescription_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_NamespacesEventhub_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhub_Spec to NamespacesEventhub_Spec via AssignPropertiesToNamespacesEventhub_Spec & AssignPropertiesFromNamespacesEventhub_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhub_Spec, NamespacesEventhub_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhub_Spec tests if a specific instance of NamespacesEventhub_Spec can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhub_Spec(subject NamespacesEventhub_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.NamespacesEventhub_Spec
	err := copied.AssignPropertiesToNamespacesEventhub_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhub_Spec
	err = actual.AssignPropertiesFromNamespacesEventhub_Spec(&other)
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

func Test_NamespacesEventhub_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhub_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhub_Spec, NamespacesEventhub_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhub_Spec runs a test to see if a specific instance of NamespacesEventhub_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhub_Spec(subject NamespacesEventhub_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhub_Spec
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

// Generator of NamespacesEventhub_Spec instances for property testing - lazily instantiated by
// NamespacesEventhub_SpecGenerator()
var namespacesEventhub_SpecGenerator gopter.Gen

// NamespacesEventhub_SpecGenerator returns a generator of NamespacesEventhub_Spec instances for property testing.
// We first initialize namespacesEventhub_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesEventhub_SpecGenerator() gopter.Gen {
	if namespacesEventhub_SpecGenerator != nil {
		return namespacesEventhub_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhub_Spec(generators)
	namespacesEventhub_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhub_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhub_Spec(generators)
	AddRelatedPropertyGeneratorsForNamespacesEventhub_Spec(generators)
	namespacesEventhub_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhub_Spec{}), generators)

	return namespacesEventhub_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhub_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhub_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MessageRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
	gens["PartitionIds"] = gen.SliceOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		NamespacesEventhub_Spec_Properties_StatusActive,
		NamespacesEventhub_Spec_Properties_StatusCreating,
		NamespacesEventhub_Spec_Properties_StatusDeleting,
		NamespacesEventhub_Spec_Properties_StatusDisabled,
		NamespacesEventhub_Spec_Properties_StatusReceiveDisabled,
		NamespacesEventhub_Spec_Properties_StatusRenaming,
		NamespacesEventhub_Spec_Properties_StatusRestoring,
		NamespacesEventhub_Spec_Properties_StatusSendDisabled,
		NamespacesEventhub_Spec_Properties_StatusUnknown))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesEventhub_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhub_Spec(gens map[string]gopter.Gen) {
	gens["CaptureDescription"] = gen.PtrOf(CaptureDescriptionGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataGenerator())
}

func Test_CaptureDescription_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from CaptureDescription to CaptureDescription via AssignPropertiesToCaptureDescription & AssignPropertiesFromCaptureDescription returns original",
		prop.ForAll(RunPropertyAssignmentTestForCaptureDescription, CaptureDescriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForCaptureDescription tests if a specific instance of CaptureDescription can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForCaptureDescription(subject CaptureDescription) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.CaptureDescription
	err := copied.AssignPropertiesToCaptureDescription(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual CaptureDescription
	err = actual.AssignPropertiesFromCaptureDescription(&other)
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

func Test_CaptureDescription_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CaptureDescription via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCaptureDescription, CaptureDescriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCaptureDescription runs a test to see if a specific instance of CaptureDescription round trips to JSON and back losslessly
func RunJSONSerializationTestForCaptureDescription(subject CaptureDescription) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CaptureDescription
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

// Generator of CaptureDescription instances for property testing - lazily instantiated by CaptureDescriptionGenerator()
var captureDescriptionGenerator gopter.Gen

// CaptureDescriptionGenerator returns a generator of CaptureDescription instances for property testing.
// We first initialize captureDescriptionGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CaptureDescriptionGenerator() gopter.Gen {
	if captureDescriptionGenerator != nil {
		return captureDescriptionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription(generators)
	captureDescriptionGenerator = gen.Struct(reflect.TypeOf(CaptureDescription{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription(generators)
	AddRelatedPropertyGeneratorsForCaptureDescription(generators)
	captureDescriptionGenerator = gen.Struct(reflect.TypeOf(CaptureDescription{}), generators)

	return captureDescriptionGenerator
}

// AddIndependentPropertyGeneratorsForCaptureDescription is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCaptureDescription(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Encoding"] = gen.PtrOf(gen.OneConstOf(CaptureDescription_EncodingAvro, CaptureDescription_EncodingAvroDeflate))
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SizeLimitInBytes"] = gen.PtrOf(gen.Int())
	gens["SkipEmptyArchives"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForCaptureDescription is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCaptureDescription(gens map[string]gopter.Gen) {
	gens["Destination"] = gen.PtrOf(DestinationGenerator())
}

func Test_CaptureDescription_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from CaptureDescription_STATUS to CaptureDescription_STATUS via AssignPropertiesToCaptureDescription_STATUS & AssignPropertiesFromCaptureDescription_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForCaptureDescription_STATUS, CaptureDescription_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForCaptureDescription_STATUS tests if a specific instance of CaptureDescription_STATUS can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForCaptureDescription_STATUS(subject CaptureDescription_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.CaptureDescription_STATUS
	err := copied.AssignPropertiesToCaptureDescription_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual CaptureDescription_STATUS
	err = actual.AssignPropertiesFromCaptureDescription_STATUS(&other)
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

func Test_CaptureDescription_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CaptureDescription_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCaptureDescription_STATUS, CaptureDescription_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCaptureDescription_STATUS runs a test to see if a specific instance of CaptureDescription_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCaptureDescription_STATUS(subject CaptureDescription_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CaptureDescription_STATUS
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

// Generator of CaptureDescription_STATUS instances for property testing - lazily instantiated by
// CaptureDescription_STATUSGenerator()
var captureDescription_STATUSGenerator gopter.Gen

// CaptureDescription_STATUSGenerator returns a generator of CaptureDescription_STATUS instances for property testing.
// We first initialize captureDescription_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CaptureDescription_STATUSGenerator() gopter.Gen {
	if captureDescription_STATUSGenerator != nil {
		return captureDescription_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription_STATUS(generators)
	captureDescription_STATUSGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription_STATUS(generators)
	AddRelatedPropertyGeneratorsForCaptureDescription_STATUS(generators)
	captureDescription_STATUSGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_STATUS{}), generators)

	return captureDescription_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCaptureDescription_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCaptureDescription_STATUS(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Encoding"] = gen.PtrOf(gen.OneConstOf(CaptureDescription_Encoding_STATUSAvro, CaptureDescription_Encoding_STATUSAvroDeflate))
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SizeLimitInBytes"] = gen.PtrOf(gen.Int())
	gens["SkipEmptyArchives"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForCaptureDescription_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCaptureDescription_STATUS(gens map[string]gopter.Gen) {
	gens["Destination"] = gen.PtrOf(Destination_STATUSGenerator())
}

func Test_Destination_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Destination to Destination via AssignPropertiesToDestination & AssignPropertiesFromDestination returns original",
		prop.ForAll(RunPropertyAssignmentTestForDestination, DestinationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDestination tests if a specific instance of Destination can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForDestination(subject Destination) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.Destination
	err := copied.AssignPropertiesToDestination(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Destination
	err = actual.AssignPropertiesFromDestination(&other)
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

func Test_Destination_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Destination via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestination, DestinationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestination runs a test to see if a specific instance of Destination round trips to JSON and back losslessly
func RunJSONSerializationTestForDestination(subject Destination) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Destination
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

// Generator of Destination instances for property testing - lazily instantiated by DestinationGenerator()
var destinationGenerator gopter.Gen

// DestinationGenerator returns a generator of Destination instances for property testing.
func DestinationGenerator() gopter.Gen {
	if destinationGenerator != nil {
		return destinationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestination(generators)
	destinationGenerator = gen.Struct(reflect.TypeOf(Destination{}), generators)

	return destinationGenerator
}

// AddIndependentPropertyGeneratorsForDestination is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestination(gens map[string]gopter.Gen) {
	gens["ArchiveNameFormat"] = gen.PtrOf(gen.AlphaString())
	gens["BlobContainer"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeAccountName"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeFolderPath"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_Destination_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Destination_STATUS to Destination_STATUS via AssignPropertiesToDestination_STATUS & AssignPropertiesFromDestination_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDestination_STATUS, Destination_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDestination_STATUS tests if a specific instance of Destination_STATUS can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForDestination_STATUS(subject Destination_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20211101s.Destination_STATUS
	err := copied.AssignPropertiesToDestination_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Destination_STATUS
	err = actual.AssignPropertiesFromDestination_STATUS(&other)
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

func Test_Destination_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Destination_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestination_STATUS, Destination_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestination_STATUS runs a test to see if a specific instance of Destination_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDestination_STATUS(subject Destination_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Destination_STATUS
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

// Generator of Destination_STATUS instances for property testing - lazily instantiated by Destination_STATUSGenerator()
var destination_STATUSGenerator gopter.Gen

// Destination_STATUSGenerator returns a generator of Destination_STATUS instances for property testing.
func Destination_STATUSGenerator() gopter.Gen {
	if destination_STATUSGenerator != nil {
		return destination_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestination_STATUS(generators)
	destination_STATUSGenerator = gen.Struct(reflect.TypeOf(Destination_STATUS{}), generators)

	return destination_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDestination_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestination_STATUS(gens map[string]gopter.Gen) {
	gens["ArchiveNameFormat"] = gen.PtrOf(gen.AlphaString())
	gens["BlobContainer"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeAccountName"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeFolderPath"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountResourceId"] = gen.PtrOf(gen.AlphaString())
}
