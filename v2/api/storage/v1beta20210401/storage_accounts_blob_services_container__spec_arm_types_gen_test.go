// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

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

func Test_StorageAccountsBlobServicesContainer_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobServicesContainer_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobServicesContainer_SpecARM, StorageAccountsBlobServicesContainer_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobServicesContainer_SpecARM runs a test to see if a specific instance of StorageAccountsBlobServicesContainer_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobServicesContainer_SpecARM(subject StorageAccountsBlobServicesContainer_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobServicesContainer_SpecARM
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

// Generator of StorageAccountsBlobServicesContainer_SpecARM instances for property testing - lazily instantiated by
// StorageAccountsBlobServicesContainer_SpecARMGenerator()
var storageAccountsBlobServicesContainer_SpecARMGenerator gopter.Gen

// StorageAccountsBlobServicesContainer_SpecARMGenerator returns a generator of StorageAccountsBlobServicesContainer_SpecARM instances for property testing.
// We first initialize storageAccountsBlobServicesContainer_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsBlobServicesContainer_SpecARMGenerator() gopter.Gen {
	if storageAccountsBlobServicesContainer_SpecARMGenerator != nil {
		return storageAccountsBlobServicesContainer_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainer_SpecARM(generators)
	storageAccountsBlobServicesContainer_SpecARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServicesContainer_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainer_SpecARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer_SpecARM(generators)
	storageAccountsBlobServicesContainer_SpecARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServicesContainer_SpecARM{}), generators)

	return storageAccountsBlobServicesContainer_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainer_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainer_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ContainerPropertiesARMGenerator())
}

func Test_ContainerPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ContainerPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForContainerPropertiesARM, ContainerPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForContainerPropertiesARM runs a test to see if a specific instance of ContainerPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForContainerPropertiesARM(subject ContainerPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ContainerPropertiesARM
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

// Generator of ContainerPropertiesARM instances for property testing - lazily instantiated by
// ContainerPropertiesARMGenerator()
var containerPropertiesARMGenerator gopter.Gen

// ContainerPropertiesARMGenerator returns a generator of ContainerPropertiesARM instances for property testing.
// We first initialize containerPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ContainerPropertiesARMGenerator() gopter.Gen {
	if containerPropertiesARMGenerator != nil {
		return containerPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerPropertiesARM(generators)
	containerPropertiesARMGenerator = gen.Struct(reflect.TypeOf(ContainerPropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForContainerPropertiesARM(generators)
	containerPropertiesARMGenerator = gen.Struct(reflect.TypeOf(ContainerPropertiesARM{}), generators)

	return containerPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForContainerPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForContainerPropertiesARM(gens map[string]gopter.Gen) {
	gens["DefaultEncryptionScope"] = gen.PtrOf(gen.AlphaString())
	gens["DenyEncryptionScopeOverride"] = gen.PtrOf(gen.Bool())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["PublicAccess"] = gen.PtrOf(gen.OneConstOf(ContainerProperties_PublicAccess_Blob, ContainerProperties_PublicAccess_Container, ContainerProperties_PublicAccess_None))
}

// AddRelatedPropertyGeneratorsForContainerPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForContainerPropertiesARM(gens map[string]gopter.Gen) {
	gens["ImmutableStorageWithVersioning"] = gen.PtrOf(ImmutableStorageWithVersioningARMGenerator())
}

func Test_ImmutableStorageWithVersioningARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutableStorageWithVersioningARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutableStorageWithVersioningARM, ImmutableStorageWithVersioningARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutableStorageWithVersioningARM runs a test to see if a specific instance of ImmutableStorageWithVersioningARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutableStorageWithVersioningARM(subject ImmutableStorageWithVersioningARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutableStorageWithVersioningARM
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

// Generator of ImmutableStorageWithVersioningARM instances for property testing - lazily instantiated by
// ImmutableStorageWithVersioningARMGenerator()
var immutableStorageWithVersioningARMGenerator gopter.Gen

// ImmutableStorageWithVersioningARMGenerator returns a generator of ImmutableStorageWithVersioningARM instances for property testing.
func ImmutableStorageWithVersioningARMGenerator() gopter.Gen {
	if immutableStorageWithVersioningARMGenerator != nil {
		return immutableStorageWithVersioningARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutableStorageWithVersioningARM(generators)
	immutableStorageWithVersioningARMGenerator = gen.Struct(reflect.TypeOf(ImmutableStorageWithVersioningARM{}), generators)

	return immutableStorageWithVersioningARMGenerator
}

// AddIndependentPropertyGeneratorsForImmutableStorageWithVersioningARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutableStorageWithVersioningARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}
