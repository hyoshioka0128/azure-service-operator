// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200202

import (
	"encoding/json"
	v20200202s "github.com/Azure/azure-service-operator/v2/api/insights/v1beta20200202storage"
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

func Test_Component_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Component to hub returns original",
		prop.ForAll(RunResourceConversionTestForComponent, ComponentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForComponent tests if a specific instance of Component round trips to the hub storage version and back losslessly
func RunResourceConversionTestForComponent(subject Component) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20200202s.Component
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Component
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

func Test_Component_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Component to Component via AssignPropertiesToComponent & AssignPropertiesFromComponent returns original",
		prop.ForAll(RunPropertyAssignmentTestForComponent, ComponentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForComponent tests if a specific instance of Component can be assigned to v1beta20200202storage and back losslessly
func RunPropertyAssignmentTestForComponent(subject Component) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200202s.Component
	err := copied.AssignPropertiesToComponent(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Component
	err = actual.AssignPropertiesFromComponent(&other)
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

func Test_Component_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Component via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForComponent, ComponentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForComponent runs a test to see if a specific instance of Component round trips to JSON and back losslessly
func RunJSONSerializationTestForComponent(subject Component) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Component
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

// Generator of Component instances for property testing - lazily instantiated by ComponentGenerator()
var componentGenerator gopter.Gen

// ComponentGenerator returns a generator of Component instances for property testing.
func ComponentGenerator() gopter.Gen {
	if componentGenerator != nil {
		return componentGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForComponent(generators)
	componentGenerator = gen.Struct(reflect.TypeOf(Component{}), generators)

	return componentGenerator
}

// AddRelatedPropertyGeneratorsForComponent is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForComponent(gens map[string]gopter.Gen) {
	gens["Spec"] = Component_SpecGenerator()
	gens["Status"] = Component_STATUSGenerator()
}

func Test_Component_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Component_STATUS to Component_STATUS via AssignPropertiesToComponent_STATUS & AssignPropertiesFromComponent_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForComponent_STATUS, Component_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForComponent_STATUS tests if a specific instance of Component_STATUS can be assigned to v1beta20200202storage and back losslessly
func RunPropertyAssignmentTestForComponent_STATUS(subject Component_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200202s.Component_STATUS
	err := copied.AssignPropertiesToComponent_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Component_STATUS
	err = actual.AssignPropertiesFromComponent_STATUS(&other)
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

func Test_Component_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Component_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForComponent_STATUS, Component_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForComponent_STATUS runs a test to see if a specific instance of Component_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForComponent_STATUS(subject Component_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Component_STATUS
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

// Generator of Component_STATUS instances for property testing - lazily instantiated by Component_STATUSGenerator()
var component_STATUSGenerator gopter.Gen

// Component_STATUSGenerator returns a generator of Component_STATUS instances for property testing.
// We first initialize component_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Component_STATUSGenerator() gopter.Gen {
	if component_STATUSGenerator != nil {
		return component_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComponent_STATUS(generators)
	component_STATUSGenerator = gen.Struct(reflect.TypeOf(Component_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComponent_STATUS(generators)
	AddRelatedPropertyGeneratorsForComponent_STATUS(generators)
	component_STATUSGenerator = gen.Struct(reflect.TypeOf(Component_STATUS{}), generators)

	return component_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForComponent_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForComponent_STATUS(gens map[string]gopter.Gen) {
	gens["AppId"] = gen.PtrOf(gen.AlphaString())
	gens["ApplicationId"] = gen.PtrOf(gen.AlphaString())
	gens["Application_Type"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Application_Type_Other_STATUS, ApplicationInsightsComponentProperties_Application_Type_Web_STATUS))
	gens["ConnectionString"] = gen.PtrOf(gen.AlphaString())
	gens["CreationDate"] = gen.PtrOf(gen.AlphaString())
	gens["DisableIpMasking"] = gen.PtrOf(gen.Bool())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Flow_Type"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Flow_Type_Bluefield_STATUS))
	gens["ForceCustomerStorageForProfiler"] = gen.PtrOf(gen.Bool())
	gens["HockeyAppId"] = gen.PtrOf(gen.AlphaString())
	gens["HockeyAppToken"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
	gens["IngestionMode"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_IngestionMode_ApplicationInsights_STATUS, ApplicationInsightsComponentProperties_IngestionMode_ApplicationInsightsWithDiagnosticSettings_STATUS, ApplicationInsightsComponentProperties_IngestionMode_LogAnalytics_STATUS))
	gens["InstrumentationKey"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["LaMigrationDate"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PropertiesName"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_Disabled_STATUS, PublicNetworkAccessType_Enabled_STATUS))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_Disabled_STATUS, PublicNetworkAccessType_Enabled_STATUS))
	gens["Request_Source"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Request_Source_Rest_STATUS))
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
	gens["SamplingPercentage"] = gen.PtrOf(gen.Float64())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["WorkspaceResourceId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForComponent_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForComponent_STATUS(gens map[string]gopter.Gen) {
	gens["PrivateLinkScopedResources"] = gen.SliceOf(PrivateLinkScopedResource_STATUSGenerator())
}

func Test_Component_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Component_Spec to Component_Spec via AssignPropertiesToComponent_Spec & AssignPropertiesFromComponent_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForComponent_Spec, Component_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForComponent_Spec tests if a specific instance of Component_Spec can be assigned to v1beta20200202storage and back losslessly
func RunPropertyAssignmentTestForComponent_Spec(subject Component_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200202s.Component_Spec
	err := copied.AssignPropertiesToComponent_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Component_Spec
	err = actual.AssignPropertiesFromComponent_Spec(&other)
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

func Test_Component_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Component_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForComponent_Spec, Component_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForComponent_Spec runs a test to see if a specific instance of Component_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForComponent_Spec(subject Component_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Component_Spec
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

// Generator of Component_Spec instances for property testing - lazily instantiated by Component_SpecGenerator()
var component_SpecGenerator gopter.Gen

// Component_SpecGenerator returns a generator of Component_Spec instances for property testing.
func Component_SpecGenerator() gopter.Gen {
	if component_SpecGenerator != nil {
		return component_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComponent_Spec(generators)
	component_SpecGenerator = gen.Struct(reflect.TypeOf(Component_Spec{}), generators)

	return component_SpecGenerator
}

// AddIndependentPropertyGeneratorsForComponent_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForComponent_Spec(gens map[string]gopter.Gen) {
	gens["Application_Type"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Application_Type_Other, ApplicationInsightsComponentProperties_Application_Type_Web))
	gens["AzureName"] = gen.AlphaString()
	gens["DisableIpMasking"] = gen.PtrOf(gen.Bool())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Flow_Type"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Flow_Type_Bluefield))
	gens["ForceCustomerStorageForProfiler"] = gen.PtrOf(gen.Bool())
	gens["HockeyAppId"] = gen.PtrOf(gen.AlphaString())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
	gens["IngestionMode"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_IngestionMode_ApplicationInsights, ApplicationInsightsComponentProperties_IngestionMode_ApplicationInsightsWithDiagnosticSettings, ApplicationInsightsComponentProperties_IngestionMode_LogAnalytics))
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_Disabled, PublicNetworkAccessType_Enabled))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_Disabled, PublicNetworkAccessType_Enabled))
	gens["Request_Source"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Request_Source_Rest))
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
	gens["SamplingPercentage"] = gen.PtrOf(gen.Float64())
}

func Test_PrivateLinkScopedResource_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateLinkScopedResource_STATUS to PrivateLinkScopedResource_STATUS via AssignPropertiesToPrivateLinkScopedResource_STATUS & AssignPropertiesFromPrivateLinkScopedResource_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateLinkScopedResource_STATUS, PrivateLinkScopedResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateLinkScopedResource_STATUS tests if a specific instance of PrivateLinkScopedResource_STATUS can be assigned to v1beta20200202storage and back losslessly
func RunPropertyAssignmentTestForPrivateLinkScopedResource_STATUS(subject PrivateLinkScopedResource_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200202s.PrivateLinkScopedResource_STATUS
	err := copied.AssignPropertiesToPrivateLinkScopedResource_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateLinkScopedResource_STATUS
	err = actual.AssignPropertiesFromPrivateLinkScopedResource_STATUS(&other)
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

func Test_PrivateLinkScopedResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkScopedResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkScopedResource_STATUS, PrivateLinkScopedResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkScopedResource_STATUS runs a test to see if a specific instance of PrivateLinkScopedResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkScopedResource_STATUS(subject PrivateLinkScopedResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkScopedResource_STATUS
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

// Generator of PrivateLinkScopedResource_STATUS instances for property testing - lazily instantiated by
// PrivateLinkScopedResource_STATUSGenerator()
var privateLinkScopedResource_STATUSGenerator gopter.Gen

// PrivateLinkScopedResource_STATUSGenerator returns a generator of PrivateLinkScopedResource_STATUS instances for property testing.
func PrivateLinkScopedResource_STATUSGenerator() gopter.Gen {
	if privateLinkScopedResource_STATUSGenerator != nil {
		return privateLinkScopedResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUS(generators)
	privateLinkScopedResource_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateLinkScopedResource_STATUS{}), generators)

	return privateLinkScopedResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUS(gens map[string]gopter.Gen) {
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ScopeId"] = gen.PtrOf(gen.AlphaString())
}
