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

func Test_StorageAccountsQueueServicesQueues_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueServicesQueues_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueServicesQueuesSPECARM, StorageAccountsQueueServicesQueuesSPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueServicesQueuesSPECARM runs a test to see if a specific instance of StorageAccountsQueueServicesQueues_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueServicesQueuesSPECARM(subject StorageAccountsQueueServicesQueues_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueServicesQueues_SPECARM
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

// Generator of StorageAccountsQueueServicesQueues_SPECARM instances for property testing - lazily instantiated by
//StorageAccountsQueueServicesQueuesSPECARMGenerator()
var storageAccountsQueueServicesQueuesSPECARMGenerator gopter.Gen

// StorageAccountsQueueServicesQueuesSPECARMGenerator returns a generator of StorageAccountsQueueServicesQueues_SPECARM instances for property testing.
// We first initialize storageAccountsQueueServicesQueuesSPECARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsQueueServicesQueuesSPECARMGenerator() gopter.Gen {
	if storageAccountsQueueServicesQueuesSPECARMGenerator != nil {
		return storageAccountsQueueServicesQueuesSPECARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueuesSPECARM(generators)
	storageAccountsQueueServicesQueuesSPECARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueues_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueuesSPECARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueuesSPECARM(generators)
	storageAccountsQueueServicesQueuesSPECARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueues_SPECARM{}), generators)

	return storageAccountsQueueServicesQueuesSPECARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueuesSPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueuesSPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueuesSPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueuesSPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(QueuePropertiesSpecARMGenerator())
}

func Test_QueueProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of QueueProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForQueuePropertiesSpecARM, QueuePropertiesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForQueuePropertiesSpecARM runs a test to see if a specific instance of QueueProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForQueuePropertiesSpecARM(subject QueueProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual QueueProperties_SpecARM
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

// Generator of QueueProperties_SpecARM instances for property testing - lazily instantiated by
//QueuePropertiesSpecARMGenerator()
var queuePropertiesSpecARMGenerator gopter.Gen

// QueuePropertiesSpecARMGenerator returns a generator of QueueProperties_SpecARM instances for property testing.
func QueuePropertiesSpecARMGenerator() gopter.Gen {
	if queuePropertiesSpecARMGenerator != nil {
		return queuePropertiesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForQueuePropertiesSpecARM(generators)
	queuePropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(QueueProperties_SpecARM{}), generators)

	return queuePropertiesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForQueuePropertiesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForQueuePropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}