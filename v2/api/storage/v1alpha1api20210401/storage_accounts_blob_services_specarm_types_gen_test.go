// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

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

func Test_StorageAccountsBlobServices_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobServices_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobServicesSPECARM, StorageAccountsBlobServicesSPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobServicesSPECARM runs a test to see if a specific instance of StorageAccountsBlobServices_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobServicesSPECARM(subject StorageAccountsBlobServices_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobServices_SPECARM
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

// Generator of StorageAccountsBlobServices_SPECARM instances for property testing - lazily instantiated by
//StorageAccountsBlobServicesSPECARMGenerator()
var storageAccountsBlobServicesSPECARMGenerator gopter.Gen

// StorageAccountsBlobServicesSPECARMGenerator returns a generator of StorageAccountsBlobServices_SPECARM instances for property testing.
// We first initialize storageAccountsBlobServicesSPECARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsBlobServicesSPECARMGenerator() gopter.Gen {
	if storageAccountsBlobServicesSPECARMGenerator != nil {
		return storageAccountsBlobServicesSPECARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesSPECARM(generators)
	storageAccountsBlobServicesSPECARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServices_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesSPECARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesSPECARM(generators)
	storageAccountsBlobServicesSPECARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServices_SPECARM{}), generators)

	return storageAccountsBlobServicesSPECARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesSPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesSPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesSPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesSPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(StorageAccountsBlobServicesSPECPropertiesARMGenerator())
}

func Test_StorageAccountsBlobServices_SPEC_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobServices_SPEC_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobServicesSPECPropertiesARM, StorageAccountsBlobServicesSPECPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobServicesSPECPropertiesARM runs a test to see if a specific instance of StorageAccountsBlobServices_SPEC_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobServicesSPECPropertiesARM(subject StorageAccountsBlobServices_SPEC_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobServices_SPEC_PropertiesARM
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

// Generator of StorageAccountsBlobServices_SPEC_PropertiesARM instances for property testing - lazily instantiated by
//StorageAccountsBlobServicesSPECPropertiesARMGenerator()
var storageAccountsBlobServicesSPECPropertiesARMGenerator gopter.Gen

// StorageAccountsBlobServicesSPECPropertiesARMGenerator returns a generator of StorageAccountsBlobServices_SPEC_PropertiesARM instances for property testing.
// We first initialize storageAccountsBlobServicesSPECPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsBlobServicesSPECPropertiesARMGenerator() gopter.Gen {
	if storageAccountsBlobServicesSPECPropertiesARMGenerator != nil {
		return storageAccountsBlobServicesSPECPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesSPECPropertiesARM(generators)
	storageAccountsBlobServicesSPECPropertiesARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServices_SPEC_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesSPECPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesSPECPropertiesARM(generators)
	storageAccountsBlobServicesSPECPropertiesARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServices_SPEC_PropertiesARM{}), generators)

	return storageAccountsBlobServicesSPECPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesSPECPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesSPECPropertiesARM(gens map[string]gopter.Gen) {
	gens["AutomaticSnapshotPolicyEnabled"] = gen.PtrOf(gen.Bool())
	gens["DefaultServiceVersion"] = gen.PtrOf(gen.AlphaString())
	gens["IsVersioningEnabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesSPECPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesSPECPropertiesARM(gens map[string]gopter.Gen) {
	gens["ChangeFeed"] = gen.PtrOf(ChangeFeedSpecARMGenerator())
	gens["ContainerDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicySpecARMGenerator())
	gens["Cors"] = gen.PtrOf(CorsRulesSpecARMGenerator())
	gens["DeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicySpecARMGenerator())
	gens["LastAccessTimeTrackingPolicy"] = gen.PtrOf(LastAccessTimeTrackingPolicySpecARMGenerator())
	gens["RestorePolicy"] = gen.PtrOf(RestorePolicyPropertiesSpecARMGenerator())
}

func Test_ChangeFeed_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ChangeFeed_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForChangeFeedSpecARM, ChangeFeedSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForChangeFeedSpecARM runs a test to see if a specific instance of ChangeFeed_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForChangeFeedSpecARM(subject ChangeFeed_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ChangeFeed_SpecARM
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

// Generator of ChangeFeed_SpecARM instances for property testing - lazily instantiated by ChangeFeedSpecARMGenerator()
var changeFeedSpecARMGenerator gopter.Gen

// ChangeFeedSpecARMGenerator returns a generator of ChangeFeed_SpecARM instances for property testing.
func ChangeFeedSpecARMGenerator() gopter.Gen {
	if changeFeedSpecARMGenerator != nil {
		return changeFeedSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForChangeFeedSpecARM(generators)
	changeFeedSpecARMGenerator = gen.Struct(reflect.TypeOf(ChangeFeed_SpecARM{}), generators)

	return changeFeedSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForChangeFeedSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForChangeFeedSpecARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

func Test_CorsRules_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRules_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRulesSpecARM, CorsRulesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRulesSpecARM runs a test to see if a specific instance of CorsRules_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRulesSpecARM(subject CorsRules_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRules_SpecARM
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

// Generator of CorsRules_SpecARM instances for property testing - lazily instantiated by CorsRulesSpecARMGenerator()
var corsRulesSpecARMGenerator gopter.Gen

// CorsRulesSpecARMGenerator returns a generator of CorsRules_SpecARM instances for property testing.
func CorsRulesSpecARMGenerator() gopter.Gen {
	if corsRulesSpecARMGenerator != nil {
		return corsRulesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForCorsRulesSpecARM(generators)
	corsRulesSpecARMGenerator = gen.Struct(reflect.TypeOf(CorsRules_SpecARM{}), generators)

	return corsRulesSpecARMGenerator
}

// AddRelatedPropertyGeneratorsForCorsRulesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCorsRulesSpecARM(gens map[string]gopter.Gen) {
	gens["CorsRules"] = gen.SliceOf(CorsRuleSpecARMGenerator())
}

func Test_DeleteRetentionPolicy_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeleteRetentionPolicy_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeleteRetentionPolicySpecARM, DeleteRetentionPolicySpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeleteRetentionPolicySpecARM runs a test to see if a specific instance of DeleteRetentionPolicy_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDeleteRetentionPolicySpecARM(subject DeleteRetentionPolicy_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DeleteRetentionPolicy_SpecARM
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

// Generator of DeleteRetentionPolicy_SpecARM instances for property testing - lazily instantiated by
//DeleteRetentionPolicySpecARMGenerator()
var deleteRetentionPolicySpecARMGenerator gopter.Gen

// DeleteRetentionPolicySpecARMGenerator returns a generator of DeleteRetentionPolicy_SpecARM instances for property testing.
func DeleteRetentionPolicySpecARMGenerator() gopter.Gen {
	if deleteRetentionPolicySpecARMGenerator != nil {
		return deleteRetentionPolicySpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeleteRetentionPolicySpecARM(generators)
	deleteRetentionPolicySpecARMGenerator = gen.Struct(reflect.TypeOf(DeleteRetentionPolicy_SpecARM{}), generators)

	return deleteRetentionPolicySpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDeleteRetentionPolicySpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeleteRetentionPolicySpecARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_LastAccessTimeTrackingPolicy_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LastAccessTimeTrackingPolicy_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLastAccessTimeTrackingPolicySpecARM, LastAccessTimeTrackingPolicySpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLastAccessTimeTrackingPolicySpecARM runs a test to see if a specific instance of LastAccessTimeTrackingPolicy_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLastAccessTimeTrackingPolicySpecARM(subject LastAccessTimeTrackingPolicy_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LastAccessTimeTrackingPolicy_SpecARM
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

// Generator of LastAccessTimeTrackingPolicy_SpecARM instances for property testing - lazily instantiated by
//LastAccessTimeTrackingPolicySpecARMGenerator()
var lastAccessTimeTrackingPolicySpecARMGenerator gopter.Gen

// LastAccessTimeTrackingPolicySpecARMGenerator returns a generator of LastAccessTimeTrackingPolicy_SpecARM instances for property testing.
func LastAccessTimeTrackingPolicySpecARMGenerator() gopter.Gen {
	if lastAccessTimeTrackingPolicySpecARMGenerator != nil {
		return lastAccessTimeTrackingPolicySpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicySpecARM(generators)
	lastAccessTimeTrackingPolicySpecARMGenerator = gen.Struct(reflect.TypeOf(LastAccessTimeTrackingPolicy_SpecARM{}), generators)

	return lastAccessTimeTrackingPolicySpecARMGenerator
}

// AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicySpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicySpecARM(gens map[string]gopter.Gen) {
	gens["BlobType"] = gen.SliceOf(gen.AlphaString())
	gens["Enable"] = gen.Bool()
	gens["Name"] = gen.PtrOf(gen.OneConstOf(LastAccessTimeTrackingPolicySpecNameAccessTimeTracking))
	gens["TrackingGranularityInDays"] = gen.PtrOf(gen.Int())
}

func Test_RestorePolicyProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestorePolicyProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestorePolicyPropertiesSpecARM, RestorePolicyPropertiesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestorePolicyPropertiesSpecARM runs a test to see if a specific instance of RestorePolicyProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRestorePolicyPropertiesSpecARM(subject RestorePolicyProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RestorePolicyProperties_SpecARM
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

// Generator of RestorePolicyProperties_SpecARM instances for property testing - lazily instantiated by
//RestorePolicyPropertiesSpecARMGenerator()
var restorePolicyPropertiesSpecARMGenerator gopter.Gen

// RestorePolicyPropertiesSpecARMGenerator returns a generator of RestorePolicyProperties_SpecARM instances for property testing.
func RestorePolicyPropertiesSpecARMGenerator() gopter.Gen {
	if restorePolicyPropertiesSpecARMGenerator != nil {
		return restorePolicyPropertiesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestorePolicyPropertiesSpecARM(generators)
	restorePolicyPropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(RestorePolicyProperties_SpecARM{}), generators)

	return restorePolicyPropertiesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForRestorePolicyPropertiesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestorePolicyPropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.Bool()
}

func Test_CorsRule_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRule_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRuleSpecARM, CorsRuleSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRuleSpecARM runs a test to see if a specific instance of CorsRule_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRuleSpecARM(subject CorsRule_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRule_SpecARM
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

// Generator of CorsRule_SpecARM instances for property testing - lazily instantiated by CorsRuleSpecARMGenerator()
var corsRuleSpecARMGenerator gopter.Gen

// CorsRuleSpecARMGenerator returns a generator of CorsRule_SpecARM instances for property testing.
func CorsRuleSpecARMGenerator() gopter.Gen {
	if corsRuleSpecARMGenerator != nil {
		return corsRuleSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorsRuleSpecARM(generators)
	corsRuleSpecARMGenerator = gen.Struct(reflect.TypeOf(CorsRule_SpecARM{}), generators)

	return corsRuleSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForCorsRuleSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorsRuleSpecARM(gens map[string]gopter.Gen) {
	gens["AllowedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedMethods"] = gen.SliceOf(gen.OneConstOf(
		CorsRuleSpecAllowedMethodsDELETE,
		CorsRuleSpecAllowedMethodsGET,
		CorsRuleSpecAllowedMethodsHEAD,
		CorsRuleSpecAllowedMethodsMERGE,
		CorsRuleSpecAllowedMethodsOPTIONS,
		CorsRuleSpecAllowedMethodsPOST,
		CorsRuleSpecAllowedMethodsPUT))
	gens["AllowedOrigins"] = gen.SliceOf(gen.AlphaString())
	gens["ExposedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["MaxAgeInSeconds"] = gen.Int()
}