// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200202

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/insights/v1alpha1api20200202storage"
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
	var hub v1alpha1api20200202storage.Component
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

// RunPropertyAssignmentTestForComponent tests if a specific instance of Component can be assigned to v1alpha1api20200202storage and back losslessly
func RunPropertyAssignmentTestForComponent(subject Component) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200202storage.Component
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
	gens["Spec"] = ComponentsSPECGenerator()
	gens["Status"] = ApplicationInsightsComponentStatusGenerator()
}

func Test_ApplicationInsightsComponent_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ApplicationInsightsComponent_Status to ApplicationInsightsComponent_Status via AssignPropertiesToApplicationInsightsComponentStatus & AssignPropertiesFromApplicationInsightsComponentStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForApplicationInsightsComponentStatus, ApplicationInsightsComponentStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForApplicationInsightsComponentStatus tests if a specific instance of ApplicationInsightsComponent_Status can be assigned to v1alpha1api20200202storage and back losslessly
func RunPropertyAssignmentTestForApplicationInsightsComponentStatus(subject ApplicationInsightsComponent_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200202storage.ApplicationInsightsComponent_Status
	err := copied.AssignPropertiesToApplicationInsightsComponentStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ApplicationInsightsComponent_Status
	err = actual.AssignPropertiesFromApplicationInsightsComponentStatus(&other)
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

func Test_ApplicationInsightsComponent_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationInsightsComponent_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationInsightsComponentStatus, ApplicationInsightsComponentStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationInsightsComponentStatus runs a test to see if a specific instance of ApplicationInsightsComponent_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationInsightsComponentStatus(subject ApplicationInsightsComponent_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationInsightsComponent_Status
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

// Generator of ApplicationInsightsComponent_Status instances for property testing - lazily instantiated by
//ApplicationInsightsComponentStatusGenerator()
var applicationInsightsComponentStatusGenerator gopter.Gen

// ApplicationInsightsComponentStatusGenerator returns a generator of ApplicationInsightsComponent_Status instances for property testing.
// We first initialize applicationInsightsComponentStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ApplicationInsightsComponentStatusGenerator() gopter.Gen {
	if applicationInsightsComponentStatusGenerator != nil {
		return applicationInsightsComponentStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationInsightsComponentStatus(generators)
	applicationInsightsComponentStatusGenerator = gen.Struct(reflect.TypeOf(ApplicationInsightsComponent_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationInsightsComponentStatus(generators)
	AddRelatedPropertyGeneratorsForApplicationInsightsComponentStatus(generators)
	applicationInsightsComponentStatusGenerator = gen.Struct(reflect.TypeOf(ApplicationInsightsComponent_Status{}), generators)

	return applicationInsightsComponentStatusGenerator
}

// AddIndependentPropertyGeneratorsForApplicationInsightsComponentStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationInsightsComponentStatus(gens map[string]gopter.Gen) {
	gens["AppId"] = gen.PtrOf(gen.AlphaString())
	gens["ApplicationId"] = gen.PtrOf(gen.AlphaString())
	gens["ApplicationType"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesStatusApplicationTypeOther, ApplicationInsightsComponentPropertiesStatusApplicationTypeWeb))
	gens["ConnectionString"] = gen.PtrOf(gen.AlphaString())
	gens["CreationDate"] = gen.PtrOf(gen.AlphaString())
	gens["DisableIpMasking"] = gen.PtrOf(gen.Bool())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["FlowType"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesStatusFlowTypeBluefield))
	gens["ForceCustomerStorageForProfiler"] = gen.PtrOf(gen.Bool())
	gens["HockeyAppId"] = gen.PtrOf(gen.AlphaString())
	gens["HockeyAppToken"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
	gens["IngestionMode"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesStatusIngestionModeApplicationInsights, ApplicationInsightsComponentPropertiesStatusIngestionModeApplicationInsightsWithDiagnosticSettings, ApplicationInsightsComponentPropertiesStatusIngestionModeLogAnalytics))
	gens["InstrumentationKey"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["LaMigrationDate"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PropertiesName"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_StatusDisabled, PublicNetworkAccessType_StatusEnabled))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_StatusDisabled, PublicNetworkAccessType_StatusEnabled))
	gens["RequestSource"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesStatusRequestSourceRest))
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
	gens["SamplingPercentage"] = gen.PtrOf(gen.Float64())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["WorkspaceResourceId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForApplicationInsightsComponentStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApplicationInsightsComponentStatus(gens map[string]gopter.Gen) {
	gens["PrivateLinkScopedResources"] = gen.SliceOf(PrivateLinkScopedResourceStatusGenerator())
}

func Test_Components_SPEC_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Components_SPEC to Components_SPEC via AssignPropertiesToComponentsSPEC & AssignPropertiesFromComponentsSPEC returns original",
		prop.ForAll(RunPropertyAssignmentTestForComponentsSPEC, ComponentsSPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForComponentsSPEC tests if a specific instance of Components_SPEC can be assigned to v1alpha1api20200202storage and back losslessly
func RunPropertyAssignmentTestForComponentsSPEC(subject Components_SPEC) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200202storage.Components_SPEC
	err := copied.AssignPropertiesToComponentsSPEC(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Components_SPEC
	err = actual.AssignPropertiesFromComponentsSPEC(&other)
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

func Test_Components_SPEC_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Components_SPEC via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForComponentsSPEC, ComponentsSPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForComponentsSPEC runs a test to see if a specific instance of Components_SPEC round trips to JSON and back losslessly
func RunJSONSerializationTestForComponentsSPEC(subject Components_SPEC) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Components_SPEC
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

// Generator of Components_SPEC instances for property testing - lazily instantiated by ComponentsSPECGenerator()
var componentsSPECGenerator gopter.Gen

// ComponentsSPECGenerator returns a generator of Components_SPEC instances for property testing.
func ComponentsSPECGenerator() gopter.Gen {
	if componentsSPECGenerator != nil {
		return componentsSPECGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComponentsSPEC(generators)
	componentsSPECGenerator = gen.Struct(reflect.TypeOf(Components_SPEC{}), generators)

	return componentsSPECGenerator
}

// AddIndependentPropertyGeneratorsForComponentsSPEC is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForComponentsSPEC(gens map[string]gopter.Gen) {
	gens["ApplicationType"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesSpecApplicationTypeOther, ApplicationInsightsComponentPropertiesSpecApplicationTypeWeb))
	gens["AzureName"] = gen.AlphaString()
	gens["DisableIpMasking"] = gen.PtrOf(gen.Bool())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["FlowType"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesSpecFlowTypeBluefield))
	gens["ForceCustomerStorageForProfiler"] = gen.PtrOf(gen.Bool())
	gens["HockeyAppId"] = gen.PtrOf(gen.AlphaString())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
	gens["IngestionMode"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesSpecIngestionModeApplicationInsights, ApplicationInsightsComponentPropertiesSpecIngestionModeApplicationInsightsWithDiagnosticSettings, ApplicationInsightsComponentPropertiesSpecIngestionModeLogAnalytics))
	gens["Kind"] = gen.AlphaString()
	gens["Location"] = gen.AlphaString()
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_SpecDisabled, PublicNetworkAccessType_SpecEnabled))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_SpecDisabled, PublicNetworkAccessType_SpecEnabled))
	gens["RequestSource"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentPropertiesSpecRequestSourceRest))
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
	gens["SamplingPercentage"] = gen.PtrOf(gen.Float64())
}

func Test_PrivateLinkScopedResource_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateLinkScopedResource_Status to PrivateLinkScopedResource_Status via AssignPropertiesToPrivateLinkScopedResourceStatus & AssignPropertiesFromPrivateLinkScopedResourceStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateLinkScopedResourceStatus, PrivateLinkScopedResourceStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateLinkScopedResourceStatus tests if a specific instance of PrivateLinkScopedResource_Status can be assigned to v1alpha1api20200202storage and back losslessly
func RunPropertyAssignmentTestForPrivateLinkScopedResourceStatus(subject PrivateLinkScopedResource_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200202storage.PrivateLinkScopedResource_Status
	err := copied.AssignPropertiesToPrivateLinkScopedResourceStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateLinkScopedResource_Status
	err = actual.AssignPropertiesFromPrivateLinkScopedResourceStatus(&other)
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

func Test_PrivateLinkScopedResource_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkScopedResource_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkScopedResourceStatus, PrivateLinkScopedResourceStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkScopedResourceStatus runs a test to see if a specific instance of PrivateLinkScopedResource_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkScopedResourceStatus(subject PrivateLinkScopedResource_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkScopedResource_Status
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

// Generator of PrivateLinkScopedResource_Status instances for property testing - lazily instantiated by
//PrivateLinkScopedResourceStatusGenerator()
var privateLinkScopedResourceStatusGenerator gopter.Gen

// PrivateLinkScopedResourceStatusGenerator returns a generator of PrivateLinkScopedResource_Status instances for property testing.
func PrivateLinkScopedResourceStatusGenerator() gopter.Gen {
	if privateLinkScopedResourceStatusGenerator != nil {
		return privateLinkScopedResourceStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkScopedResourceStatus(generators)
	privateLinkScopedResourceStatusGenerator = gen.Struct(reflect.TypeOf(PrivateLinkScopedResource_Status{}), generators)

	return privateLinkScopedResourceStatusGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkScopedResourceStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkScopedResourceStatus(gens map[string]gopter.Gen) {
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ScopeId"] = gen.PtrOf(gen.AlphaString())
}
