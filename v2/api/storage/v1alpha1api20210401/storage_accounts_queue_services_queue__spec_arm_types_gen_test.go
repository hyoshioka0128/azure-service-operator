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

func Test_StorageAccountsQueueServicesQueue_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueServicesQueue_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueServicesQueue_SpecARM, StorageAccountsQueueServicesQueue_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueServicesQueue_SpecARM runs a test to see if a specific instance of StorageAccountsQueueServicesQueue_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueServicesQueue_SpecARM(subject StorageAccountsQueueServicesQueue_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueServicesQueue_SpecARM
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

// Generator of StorageAccountsQueueServicesQueue_SpecARM instances for property testing - lazily instantiated by
// StorageAccountsQueueServicesQueue_SpecARMGenerator()
var storageAccountsQueueServicesQueue_SpecARMGenerator gopter.Gen

// StorageAccountsQueueServicesQueue_SpecARMGenerator returns a generator of StorageAccountsQueueServicesQueue_SpecARM instances for property testing.
// We first initialize storageAccountsQueueServicesQueue_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsQueueServicesQueue_SpecARMGenerator() gopter.Gen {
	if storageAccountsQueueServicesQueue_SpecARMGenerator != nil {
		return storageAccountsQueueServicesQueue_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_SpecARM(generators)
	storageAccountsQueueServicesQueue_SpecARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueue_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_SpecARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueue_SpecARM(generators)
	storageAccountsQueueServicesQueue_SpecARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueue_SpecARM{}), generators)

	return storageAccountsQueueServicesQueue_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueue_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueue_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(QueuePropertiesARMGenerator())
}

func Test_QueuePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of QueuePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForQueuePropertiesARM, QueuePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForQueuePropertiesARM runs a test to see if a specific instance of QueuePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForQueuePropertiesARM(subject QueuePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual QueuePropertiesARM
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

// Generator of QueuePropertiesARM instances for property testing - lazily instantiated by QueuePropertiesARMGenerator()
var queuePropertiesARMGenerator gopter.Gen

// QueuePropertiesARMGenerator returns a generator of QueuePropertiesARM instances for property testing.
func QueuePropertiesARMGenerator() gopter.Gen {
	if queuePropertiesARMGenerator != nil {
		return queuePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForQueuePropertiesARM(generators)
	queuePropertiesARMGenerator = gen.Struct(reflect.TypeOf(QueuePropertiesARM{}), generators)

	return queuePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForQueuePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForQueuePropertiesARM(gens map[string]gopter.Gen) {
	gens["ApproximateMessageCount"] = gen.PtrOf(gen.Int())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}