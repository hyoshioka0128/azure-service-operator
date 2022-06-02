// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200202

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

func Test_Component_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Component_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForComponent_SpecARM, Component_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForComponent_SpecARM runs a test to see if a specific instance of Component_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForComponent_SpecARM(subject Component_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Component_SpecARM
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

// Generator of Component_SpecARM instances for property testing - lazily instantiated by Component_SpecARMGenerator()
var component_SpecARMGenerator gopter.Gen

// Component_SpecARMGenerator returns a generator of Component_SpecARM instances for property testing.
// We first initialize component_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Component_SpecARMGenerator() gopter.Gen {
	if component_SpecARMGenerator != nil {
		return component_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComponent_SpecARM(generators)
	component_SpecARMGenerator = gen.Struct(reflect.TypeOf(Component_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComponent_SpecARM(generators)
	AddRelatedPropertyGeneratorsForComponent_SpecARM(generators)
	component_SpecARMGenerator = gen.Struct(reflect.TypeOf(Component_SpecARM{}), generators)

	return component_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForComponent_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForComponent_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForComponent_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForComponent_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ApplicationInsightsComponentPropertiesARMGenerator())
}

func Test_ApplicationInsightsComponentPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationInsightsComponentPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationInsightsComponentPropertiesARM, ApplicationInsightsComponentPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationInsightsComponentPropertiesARM runs a test to see if a specific instance of ApplicationInsightsComponentPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationInsightsComponentPropertiesARM(subject ApplicationInsightsComponentPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationInsightsComponentPropertiesARM
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

// Generator of ApplicationInsightsComponentPropertiesARM instances for property testing - lazily instantiated by
// ApplicationInsightsComponentPropertiesARMGenerator()
var applicationInsightsComponentPropertiesARMGenerator gopter.Gen

// ApplicationInsightsComponentPropertiesARMGenerator returns a generator of ApplicationInsightsComponentPropertiesARM instances for property testing.
func ApplicationInsightsComponentPropertiesARMGenerator() gopter.Gen {
	if applicationInsightsComponentPropertiesARMGenerator != nil {
		return applicationInsightsComponentPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationInsightsComponentPropertiesARM(generators)
	applicationInsightsComponentPropertiesARMGenerator = gen.Struct(reflect.TypeOf(ApplicationInsightsComponentPropertiesARM{}), generators)

	return applicationInsightsComponentPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForApplicationInsightsComponentPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationInsightsComponentPropertiesARM(gens map[string]gopter.Gen) {
	gens["Application_Type"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Application_Type_Other, ApplicationInsightsComponentProperties_Application_Type_Web))
	gens["DisableIpMasking"] = gen.PtrOf(gen.Bool())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["Flow_Type"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Flow_Type_Bluefield))
	gens["ForceCustomerStorageForProfiler"] = gen.PtrOf(gen.Bool())
	gens["HockeyAppId"] = gen.PtrOf(gen.AlphaString())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
	gens["IngestionMode"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_IngestionMode_ApplicationInsights, ApplicationInsightsComponentProperties_IngestionMode_ApplicationInsightsWithDiagnosticSettings, ApplicationInsightsComponentProperties_IngestionMode_LogAnalytics))
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_Disabled, PublicNetworkAccessType_Enabled))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_Disabled, PublicNetworkAccessType_Enabled))
	gens["Request_Source"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Request_Source_Rest))
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
	gens["SamplingPercentage"] = gen.PtrOf(gen.Float64())
	gens["WorkspaceResourceId"] = gen.PtrOf(gen.AlphaString())
}