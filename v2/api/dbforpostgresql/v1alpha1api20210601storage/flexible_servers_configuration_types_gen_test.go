// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601storage

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

func Test_FlexibleServersConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersConfiguration, FlexibleServersConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersConfiguration runs a test to see if a specific instance of FlexibleServersConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersConfiguration(subject FlexibleServersConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersConfiguration
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

// Generator of FlexibleServersConfiguration instances for property testing - lazily instantiated by
//FlexibleServersConfigurationGenerator()
var flexibleServersConfigurationGenerator gopter.Gen

// FlexibleServersConfigurationGenerator returns a generator of FlexibleServersConfiguration instances for property testing.
func FlexibleServersConfigurationGenerator() gopter.Gen {
	if flexibleServersConfigurationGenerator != nil {
		return flexibleServersConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersConfiguration(generators)
	flexibleServersConfigurationGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfiguration{}), generators)

	return flexibleServersConfigurationGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersConfiguration is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersConfiguration(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServersConfigurationsSPECGenerator()
	gens["Status"] = ConfigurationStatusGenerator()
}

func Test_Configuration_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Configuration_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfigurationStatus, ConfigurationStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfigurationStatus runs a test to see if a specific instance of Configuration_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForConfigurationStatus(subject Configuration_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Configuration_Status
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

// Generator of Configuration_Status instances for property testing - lazily instantiated by
//ConfigurationStatusGenerator()
var configurationStatusGenerator gopter.Gen

// ConfigurationStatusGenerator returns a generator of Configuration_Status instances for property testing.
// We first initialize configurationStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ConfigurationStatusGenerator() gopter.Gen {
	if configurationStatusGenerator != nil {
		return configurationStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationStatus(generators)
	configurationStatusGenerator = gen.Struct(reflect.TypeOf(Configuration_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationStatus(generators)
	AddRelatedPropertyGeneratorsForConfigurationStatus(generators)
	configurationStatusGenerator = gen.Struct(reflect.TypeOf(Configuration_Status{}), generators)

	return configurationStatusGenerator
}

// AddIndependentPropertyGeneratorsForConfigurationStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfigurationStatus(gens map[string]gopter.Gen) {
	gens["AllowedValues"] = gen.PtrOf(gen.AlphaString())
	gens["DataType"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultValue"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForConfigurationStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConfigurationStatus(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemDataStatusGenerator())
}

func Test_FlexibleServersConfigurations_SPEC_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersConfigurations_SPEC via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersConfigurationsSPEC, FlexibleServersConfigurationsSPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersConfigurationsSPEC runs a test to see if a specific instance of FlexibleServersConfigurations_SPEC round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersConfigurationsSPEC(subject FlexibleServersConfigurations_SPEC) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersConfigurations_SPEC
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

// Generator of FlexibleServersConfigurations_SPEC instances for property testing - lazily instantiated by
//FlexibleServersConfigurationsSPECGenerator()
var flexibleServersConfigurationsSPECGenerator gopter.Gen

// FlexibleServersConfigurationsSPECGenerator returns a generator of FlexibleServersConfigurations_SPEC instances for property testing.
func FlexibleServersConfigurationsSPECGenerator() gopter.Gen {
	if flexibleServersConfigurationsSPECGenerator != nil {
		return flexibleServersConfigurationsSPECGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersConfigurationsSPEC(generators)
	flexibleServersConfigurationsSPECGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfigurations_SPEC{}), generators)

	return flexibleServersConfigurationsSPECGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersConfigurationsSPEC is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersConfigurationsSPEC(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}
