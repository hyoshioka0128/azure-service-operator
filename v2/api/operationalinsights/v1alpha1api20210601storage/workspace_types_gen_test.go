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

func Test_Workspace_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspace via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspace, WorkspaceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspace runs a test to see if a specific instance of Workspace round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspace(subject Workspace) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspace
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

// Generator of Workspace instances for property testing - lazily instantiated by WorkspaceGenerator()
var workspaceGenerator gopter.Gen

// WorkspaceGenerator returns a generator of Workspace instances for property testing.
func WorkspaceGenerator() gopter.Gen {
	if workspaceGenerator != nil {
		return workspaceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForWorkspace(generators)
	workspaceGenerator = gen.Struct(reflect.TypeOf(Workspace{}), generators)

	return workspaceGenerator
}

// AddRelatedPropertyGeneratorsForWorkspace is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspace(gens map[string]gopter.Gen) {
	gens["Spec"] = WorkspacesSPECGenerator()
	gens["Status"] = WorkspaceStatusGenerator()
}

func Test_Workspace_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspace_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceStatus, WorkspaceStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceStatus runs a test to see if a specific instance of Workspace_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceStatus(subject Workspace_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspace_Status
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

// Generator of Workspace_Status instances for property testing - lazily instantiated by WorkspaceStatusGenerator()
var workspaceStatusGenerator gopter.Gen

// WorkspaceStatusGenerator returns a generator of Workspace_Status instances for property testing.
// We first initialize workspaceStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspaceStatusGenerator() gopter.Gen {
	if workspaceStatusGenerator != nil {
		return workspaceStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceStatus(generators)
	workspaceStatusGenerator = gen.Struct(reflect.TypeOf(Workspace_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceStatus(generators)
	AddRelatedPropertyGeneratorsForWorkspaceStatus(generators)
	workspaceStatusGenerator = gen.Struct(reflect.TypeOf(Workspace_Status{}), generators)

	return workspaceStatusGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceStatus(gens map[string]gopter.Gen) {
	gens["CreatedDate"] = gen.PtrOf(gen.AlphaString())
	gens["CustomerId"] = gen.PtrOf(gen.AlphaString())
	gens["ETag"] = gen.PtrOf(gen.AlphaString())
	gens["ForceCmkForQuery"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["ModifiedDate"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.AlphaString())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspaceStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaceStatus(gens map[string]gopter.Gen) {
	gens["Features"] = gen.PtrOf(WorkspaceFeaturesStatusGenerator())
	gens["PrivateLinkScopedResources"] = gen.SliceOf(PrivateLinkScopedResourceStatusGenerator())
	gens["Sku"] = gen.PtrOf(WorkspaceSkuStatusGenerator())
	gens["WorkspaceCapping"] = gen.PtrOf(WorkspaceCappingStatusGenerator())
}

func Test_Workspaces_SPEC_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspaces_SPEC via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspacesSPEC, WorkspacesSPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspacesSPEC runs a test to see if a specific instance of Workspaces_SPEC round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspacesSPEC(subject Workspaces_SPEC) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspaces_SPEC
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

// Generator of Workspaces_SPEC instances for property testing - lazily instantiated by WorkspacesSPECGenerator()
var workspacesSPECGenerator gopter.Gen

// WorkspacesSPECGenerator returns a generator of Workspaces_SPEC instances for property testing.
// We first initialize workspacesSPECGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspacesSPECGenerator() gopter.Gen {
	if workspacesSPECGenerator != nil {
		return workspacesSPECGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacesSPEC(generators)
	workspacesSPECGenerator = gen.Struct(reflect.TypeOf(Workspaces_SPEC{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacesSPEC(generators)
	AddRelatedPropertyGeneratorsForWorkspacesSPEC(generators)
	workspacesSPECGenerator = gen.Struct(reflect.TypeOf(Workspaces_SPEC{}), generators)

	return workspacesSPECGenerator
}

// AddIndependentPropertyGeneratorsForWorkspacesSPEC is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspacesSPEC(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["ETag"] = gen.PtrOf(gen.AlphaString())
	gens["ForceCmkForQuery"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.AlphaString())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspacesSPEC is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspacesSPEC(gens map[string]gopter.Gen) {
	gens["Features"] = gen.PtrOf(WorkspaceFeaturesSpecGenerator())
	gens["Sku"] = gen.PtrOf(WorkspaceSkuSpecGenerator())
	gens["WorkspaceCapping"] = gen.PtrOf(WorkspaceCappingSpecGenerator())
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

func Test_WorkspaceCapping_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceCapping_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceCappingSpec, WorkspaceCappingSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceCappingSpec runs a test to see if a specific instance of WorkspaceCapping_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceCappingSpec(subject WorkspaceCapping_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceCapping_Spec
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

// Generator of WorkspaceCapping_Spec instances for property testing - lazily instantiated by
//WorkspaceCappingSpecGenerator()
var workspaceCappingSpecGenerator gopter.Gen

// WorkspaceCappingSpecGenerator returns a generator of WorkspaceCapping_Spec instances for property testing.
func WorkspaceCappingSpecGenerator() gopter.Gen {
	if workspaceCappingSpecGenerator != nil {
		return workspaceCappingSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceCappingSpec(generators)
	workspaceCappingSpecGenerator = gen.Struct(reflect.TypeOf(WorkspaceCapping_Spec{}), generators)

	return workspaceCappingSpecGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceCappingSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceCappingSpec(gens map[string]gopter.Gen) {
	gens["DailyQuotaGb"] = gen.PtrOf(gen.Float64())
}

func Test_WorkspaceCapping_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceCapping_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceCappingStatus, WorkspaceCappingStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceCappingStatus runs a test to see if a specific instance of WorkspaceCapping_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceCappingStatus(subject WorkspaceCapping_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceCapping_Status
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

// Generator of WorkspaceCapping_Status instances for property testing - lazily instantiated by
//WorkspaceCappingStatusGenerator()
var workspaceCappingStatusGenerator gopter.Gen

// WorkspaceCappingStatusGenerator returns a generator of WorkspaceCapping_Status instances for property testing.
func WorkspaceCappingStatusGenerator() gopter.Gen {
	if workspaceCappingStatusGenerator != nil {
		return workspaceCappingStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceCappingStatus(generators)
	workspaceCappingStatusGenerator = gen.Struct(reflect.TypeOf(WorkspaceCapping_Status{}), generators)

	return workspaceCappingStatusGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceCappingStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceCappingStatus(gens map[string]gopter.Gen) {
	gens["DailyQuotaGb"] = gen.PtrOf(gen.Float64())
	gens["DataIngestionStatus"] = gen.PtrOf(gen.AlphaString())
	gens["QuotaNextResetTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_WorkspaceFeatures_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceFeatures_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceFeaturesSpec, WorkspaceFeaturesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceFeaturesSpec runs a test to see if a specific instance of WorkspaceFeatures_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceFeaturesSpec(subject WorkspaceFeatures_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceFeatures_Spec
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

// Generator of WorkspaceFeatures_Spec instances for property testing - lazily instantiated by
//WorkspaceFeaturesSpecGenerator()
var workspaceFeaturesSpecGenerator gopter.Gen

// WorkspaceFeaturesSpecGenerator returns a generator of WorkspaceFeatures_Spec instances for property testing.
func WorkspaceFeaturesSpecGenerator() gopter.Gen {
	if workspaceFeaturesSpecGenerator != nil {
		return workspaceFeaturesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceFeaturesSpec(generators)
	workspaceFeaturesSpecGenerator = gen.Struct(reflect.TypeOf(WorkspaceFeatures_Spec{}), generators)

	return workspaceFeaturesSpecGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceFeaturesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceFeaturesSpec(gens map[string]gopter.Gen) {
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["EnableDataExport"] = gen.PtrOf(gen.Bool())
	gens["EnableLogAccessUsingOnlyResourcePermissions"] = gen.PtrOf(gen.Bool())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
}

func Test_WorkspaceFeatures_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceFeatures_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceFeaturesStatus, WorkspaceFeaturesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceFeaturesStatus runs a test to see if a specific instance of WorkspaceFeatures_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceFeaturesStatus(subject WorkspaceFeatures_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceFeatures_Status
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

// Generator of WorkspaceFeatures_Status instances for property testing - lazily instantiated by
//WorkspaceFeaturesStatusGenerator()
var workspaceFeaturesStatusGenerator gopter.Gen

// WorkspaceFeaturesStatusGenerator returns a generator of WorkspaceFeatures_Status instances for property testing.
func WorkspaceFeaturesStatusGenerator() gopter.Gen {
	if workspaceFeaturesStatusGenerator != nil {
		return workspaceFeaturesStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceFeaturesStatus(generators)
	workspaceFeaturesStatusGenerator = gen.Struct(reflect.TypeOf(WorkspaceFeatures_Status{}), generators)

	return workspaceFeaturesStatusGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceFeaturesStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceFeaturesStatus(gens map[string]gopter.Gen) {
	gens["ClusterResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["EnableDataExport"] = gen.PtrOf(gen.Bool())
	gens["EnableLogAccessUsingOnlyResourcePermissions"] = gen.PtrOf(gen.Bool())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
}

func Test_WorkspaceSku_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceSku_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceSkuSpec, WorkspaceSkuSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceSkuSpec runs a test to see if a specific instance of WorkspaceSku_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceSkuSpec(subject WorkspaceSku_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceSku_Spec
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

// Generator of WorkspaceSku_Spec instances for property testing - lazily instantiated by WorkspaceSkuSpecGenerator()
var workspaceSkuSpecGenerator gopter.Gen

// WorkspaceSkuSpecGenerator returns a generator of WorkspaceSku_Spec instances for property testing.
func WorkspaceSkuSpecGenerator() gopter.Gen {
	if workspaceSkuSpecGenerator != nil {
		return workspaceSkuSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceSkuSpec(generators)
	workspaceSkuSpecGenerator = gen.Struct(reflect.TypeOf(WorkspaceSku_Spec{}), generators)

	return workspaceSkuSpecGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceSkuSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceSkuSpec(gens map[string]gopter.Gen) {
	gens["CapacityReservationLevel"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_WorkspaceSku_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceSku_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceSkuStatus, WorkspaceSkuStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceSkuStatus runs a test to see if a specific instance of WorkspaceSku_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceSkuStatus(subject WorkspaceSku_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceSku_Status
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

// Generator of WorkspaceSku_Status instances for property testing - lazily instantiated by WorkspaceSkuStatusGenerator()
var workspaceSkuStatusGenerator gopter.Gen

// WorkspaceSkuStatusGenerator returns a generator of WorkspaceSku_Status instances for property testing.
func WorkspaceSkuStatusGenerator() gopter.Gen {
	if workspaceSkuStatusGenerator != nil {
		return workspaceSkuStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceSkuStatus(generators)
	workspaceSkuStatusGenerator = gen.Struct(reflect.TypeOf(WorkspaceSku_Status{}), generators)

	return workspaceSkuStatusGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceSkuStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceSkuStatus(gens map[string]gopter.Gen) {
	gens["CapacityReservationLevel"] = gen.PtrOf(gen.Int())
	gens["LastSkuUpdate"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}
