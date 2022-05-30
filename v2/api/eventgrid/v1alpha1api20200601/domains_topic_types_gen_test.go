// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import (
	"encoding/json"
	alpha20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601storage"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601storage"
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

func Test_DomainsTopic_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DomainsTopic to hub returns original",
		prop.ForAll(RunResourceConversionTestForDomainsTopic, DomainsTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForDomainsTopic tests if a specific instance of DomainsTopic round trips to the hub storage version and back losslessly
func RunResourceConversionTestForDomainsTopic(subject DomainsTopic) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20200601s.DomainsTopic
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual DomainsTopic
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

func Test_DomainsTopic_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DomainsTopic to DomainsTopic via AssignPropertiesToDomainsTopic & AssignPropertiesFromDomainsTopic returns original",
		prop.ForAll(RunPropertyAssignmentTestForDomainsTopic, DomainsTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDomainsTopic tests if a specific instance of DomainsTopic can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForDomainsTopic(subject DomainsTopic) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20200601s.DomainsTopic
	err := copied.AssignPropertiesToDomainsTopic(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DomainsTopic
	err = actual.AssignPropertiesFromDomainsTopic(&other)
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

func Test_DomainsTopic_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainsTopic via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopic, DomainsTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopic runs a test to see if a specific instance of DomainsTopic round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopic(subject DomainsTopic) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainsTopic
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

// Generator of DomainsTopic instances for property testing - lazily instantiated by DomainsTopicGenerator()
var domainsTopicGenerator gopter.Gen

// DomainsTopicGenerator returns a generator of DomainsTopic instances for property testing.
func DomainsTopicGenerator() gopter.Gen {
	if domainsTopicGenerator != nil {
		return domainsTopicGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDomainsTopic(generators)
	domainsTopicGenerator = gen.Struct(reflect.TypeOf(DomainsTopic{}), generators)

	return domainsTopicGenerator
}

// AddRelatedPropertyGeneratorsForDomainsTopic is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainsTopic(gens map[string]gopter.Gen) {
	gens["Spec"] = DomainsTopic_SpecGenerator()
	gens["Status"] = DomainsTopic_STATUSGenerator()
}

func Test_DomainsTopic_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DomainsTopic_STATUS to DomainsTopic_STATUS via AssignPropertiesToDomainsTopic_STATUS & AssignPropertiesFromDomainsTopic_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDomainsTopic_STATUS, DomainsTopic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDomainsTopic_STATUS tests if a specific instance of DomainsTopic_STATUS can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForDomainsTopic_STATUS(subject DomainsTopic_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20200601s.DomainsTopic_STATUS
	err := copied.AssignPropertiesToDomainsTopic_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DomainsTopic_STATUS
	err = actual.AssignPropertiesFromDomainsTopic_STATUS(&other)
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

func Test_DomainsTopic_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainsTopic_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopic_STATUS, DomainsTopic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopic_STATUS runs a test to see if a specific instance of DomainsTopic_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopic_STATUS(subject DomainsTopic_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainsTopic_STATUS
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

// Generator of DomainsTopic_STATUS instances for property testing - lazily instantiated by
// DomainsTopic_STATUSGenerator()
var domainsTopic_STATUSGenerator gopter.Gen

// DomainsTopic_STATUSGenerator returns a generator of DomainsTopic_STATUS instances for property testing.
// We first initialize domainsTopic_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainsTopic_STATUSGenerator() gopter.Gen {
	if domainsTopic_STATUSGenerator != nil {
		return domainsTopic_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopic_STATUS(generators)
	domainsTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(DomainsTopic_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopic_STATUS(generators)
	AddRelatedPropertyGeneratorsForDomainsTopic_STATUS(generators)
	domainsTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(DomainsTopic_STATUS{}), generators)

	return domainsTopic_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDomainsTopic_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainsTopic_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DomainTopicProperties_ProvisioningState_Canceled_STATUS,
		DomainTopicProperties_ProvisioningState_Creating_STATUS,
		DomainTopicProperties_ProvisioningState_Deleting_STATUS,
		DomainTopicProperties_ProvisioningState_Failed_STATUS,
		DomainTopicProperties_ProvisioningState_Succeeded_STATUS,
		DomainTopicProperties_ProvisioningState_Updating_STATUS))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDomainsTopic_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainsTopic_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_DomainsTopic_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DomainsTopic_Spec to DomainsTopic_Spec via AssignPropertiesToDomainsTopic_Spec & AssignPropertiesFromDomainsTopic_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDomainsTopic_Spec, DomainsTopic_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDomainsTopic_Spec tests if a specific instance of DomainsTopic_Spec can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForDomainsTopic_Spec(subject DomainsTopic_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20200601s.DomainsTopic_Spec
	err := copied.AssignPropertiesToDomainsTopic_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DomainsTopic_Spec
	err = actual.AssignPropertiesFromDomainsTopic_Spec(&other)
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

func Test_DomainsTopic_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainsTopic_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopic_Spec, DomainsTopic_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopic_Spec runs a test to see if a specific instance of DomainsTopic_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopic_Spec(subject DomainsTopic_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainsTopic_Spec
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

// Generator of DomainsTopic_Spec instances for property testing - lazily instantiated by DomainsTopic_SpecGenerator()
var domainsTopic_SpecGenerator gopter.Gen

// DomainsTopic_SpecGenerator returns a generator of DomainsTopic_Spec instances for property testing.
func DomainsTopic_SpecGenerator() gopter.Gen {
	if domainsTopic_SpecGenerator != nil {
		return domainsTopic_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopic_Spec(generators)
	domainsTopic_SpecGenerator = gen.Struct(reflect.TypeOf(DomainsTopic_Spec{}), generators)

	return domainsTopic_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDomainsTopic_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainsTopic_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
}
