// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

import (
	"encoding/json"
	v20180601s "github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1beta20180601storage"
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

func Test_Configuration_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Configuration to hub returns original",
		prop.ForAll(RunResourceConversionTestForConfiguration, ConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForConfiguration tests if a specific instance of Configuration round trips to the hub storage version and back losslessly
func RunResourceConversionTestForConfiguration(subject Configuration) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20180601s.Configuration
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Configuration
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

func Test_Configuration_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Configuration to Configuration via AssignPropertiesToConfiguration & AssignPropertiesFromConfiguration returns original",
		prop.ForAll(RunPropertyAssignmentTestForConfiguration, ConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForConfiguration tests if a specific instance of Configuration can be assigned to v1beta20180601storage and back losslessly
func RunPropertyAssignmentTestForConfiguration(subject Configuration) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180601s.Configuration
	err := copied.AssignPropertiesToConfiguration(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Configuration
	err = actual.AssignPropertiesFromConfiguration(&other)
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

func Test_Configuration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Configuration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfiguration, ConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfiguration runs a test to see if a specific instance of Configuration round trips to JSON and back losslessly
func RunJSONSerializationTestForConfiguration(subject Configuration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Configuration
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

// Generator of Configuration instances for property testing - lazily instantiated by ConfigurationGenerator()
var configurationGenerator gopter.Gen

// ConfigurationGenerator returns a generator of Configuration instances for property testing.
func ConfigurationGenerator() gopter.Gen {
	if configurationGenerator != nil {
		return configurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForConfiguration(generators)
	configurationGenerator = gen.Struct(reflect.TypeOf(Configuration{}), generators)

	return configurationGenerator
}

// AddRelatedPropertyGeneratorsForConfiguration is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConfiguration(gens map[string]gopter.Gen) {
	gens["Spec"] = ServersConfigurationsSpecGenerator()
	gens["Status"] = ConfigurationStatusGenerator()
}

func Test_Configuration_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Configuration_Status to Configuration_Status via AssignPropertiesToConfigurationStatus & AssignPropertiesFromConfigurationStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForConfigurationStatus, ConfigurationStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForConfigurationStatus tests if a specific instance of Configuration_Status can be assigned to v1beta20180601storage and back losslessly
func RunPropertyAssignmentTestForConfigurationStatus(subject Configuration_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180601s.Configuration_Status
	err := copied.AssignPropertiesToConfigurationStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Configuration_Status
	err = actual.AssignPropertiesFromConfigurationStatus(&other)
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

func Test_Configuration_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
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
// ConfigurationStatusGenerator()
var configurationStatusGenerator gopter.Gen

// ConfigurationStatusGenerator returns a generator of Configuration_Status instances for property testing.
func ConfigurationStatusGenerator() gopter.Gen {
	if configurationStatusGenerator != nil {
		return configurationStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationStatus(generators)
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

func Test_ServersConfigurations_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersConfigurations_Spec to ServersConfigurations_Spec via AssignPropertiesToServersConfigurationsSpec & AssignPropertiesFromServersConfigurationsSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersConfigurationsSpec, ServersConfigurationsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersConfigurationsSpec tests if a specific instance of ServersConfigurations_Spec can be assigned to v1beta20180601storage and back losslessly
func RunPropertyAssignmentTestForServersConfigurationsSpec(subject ServersConfigurations_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180601s.ServersConfigurations_Spec
	err := copied.AssignPropertiesToServersConfigurationsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersConfigurations_Spec
	err = actual.AssignPropertiesFromServersConfigurationsSpec(&other)
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

func Test_ServersConfigurations_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersConfigurations_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersConfigurationsSpec, ServersConfigurationsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersConfigurationsSpec runs a test to see if a specific instance of ServersConfigurations_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersConfigurationsSpec(subject ServersConfigurations_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersConfigurations_Spec
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

// Generator of ServersConfigurations_Spec instances for property testing - lazily instantiated by
// ServersConfigurationsSpecGenerator()
var serversConfigurationsSpecGenerator gopter.Gen

// ServersConfigurationsSpecGenerator returns a generator of ServersConfigurations_Spec instances for property testing.
func ServersConfigurationsSpecGenerator() gopter.Gen {
	if serversConfigurationsSpecGenerator != nil {
		return serversConfigurationsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersConfigurationsSpec(generators)
	serversConfigurationsSpecGenerator = gen.Struct(reflect.TypeOf(ServersConfigurations_Spec{}), generators)

	return serversConfigurationsSpecGenerator
}

// AddIndependentPropertyGeneratorsForServersConfigurationsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersConfigurationsSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}
