// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601storage

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

func Test_Configuration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
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
	gens["Spec"] = ServersConfiguration_SpecGenerator()
	gens["Status"] = ServersConfiguration_STATUSGenerator()
}

func Test_ServersConfiguration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersConfiguration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersConfiguration_STATUS, ServersConfiguration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersConfiguration_STATUS runs a test to see if a specific instance of ServersConfiguration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServersConfiguration_STATUS(subject ServersConfiguration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersConfiguration_STATUS
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

// Generator of ServersConfiguration_STATUS instances for property testing - lazily instantiated by
// ServersConfiguration_STATUSGenerator()
var serversConfiguration_STATUSGenerator gopter.Gen

// ServersConfiguration_STATUSGenerator returns a generator of ServersConfiguration_STATUS instances for property testing.
func ServersConfiguration_STATUSGenerator() gopter.Gen {
	if serversConfiguration_STATUSGenerator != nil {
		return serversConfiguration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersConfiguration_STATUS(generators)
	serversConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersConfiguration_STATUS{}), generators)

	return serversConfiguration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServersConfiguration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersConfiguration_STATUS(gens map[string]gopter.Gen) {
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

func Test_ServersConfiguration_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersConfiguration_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersConfiguration_Spec, ServersConfiguration_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersConfiguration_Spec runs a test to see if a specific instance of ServersConfiguration_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersConfiguration_Spec(subject ServersConfiguration_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersConfiguration_Spec
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

// Generator of ServersConfiguration_Spec instances for property testing - lazily instantiated by
// ServersConfiguration_SpecGenerator()
var serversConfiguration_SpecGenerator gopter.Gen

// ServersConfiguration_SpecGenerator returns a generator of ServersConfiguration_Spec instances for property testing.
func ServersConfiguration_SpecGenerator() gopter.Gen {
	if serversConfiguration_SpecGenerator != nil {
		return serversConfiguration_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersConfiguration_Spec(generators)
	serversConfiguration_SpecGenerator = gen.Struct(reflect.TypeOf(ServersConfiguration_Spec{}), generators)

	return serversConfiguration_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServersConfiguration_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersConfiguration_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}
