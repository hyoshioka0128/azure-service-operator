// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

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

func Test_NamespacesQueue_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesQueue_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesQueue_SpecARM, NamespacesQueue_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesQueue_SpecARM runs a test to see if a specific instance of NamespacesQueue_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesQueue_SpecARM(subject NamespacesQueue_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesQueue_SpecARM
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

// Generator of NamespacesQueue_SpecARM instances for property testing - lazily instantiated by
// NamespacesQueue_SpecARMGenerator()
var namespacesQueue_SpecARMGenerator gopter.Gen

// NamespacesQueue_SpecARMGenerator returns a generator of NamespacesQueue_SpecARM instances for property testing.
// We first initialize namespacesQueue_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesQueue_SpecARMGenerator() gopter.Gen {
	if namespacesQueue_SpecARMGenerator != nil {
		return namespacesQueue_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesQueue_SpecARM(generators)
	namespacesQueue_SpecARMGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesQueue_SpecARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesQueue_SpecARM(generators)
	namespacesQueue_SpecARMGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue_SpecARM{}), generators)

	return namespacesQueue_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesQueue_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesQueue_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForNamespacesQueue_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesQueue_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SBQueuePropertiesARMGenerator())
}

func Test_SBQueuePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBQueuePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBQueuePropertiesARM, SBQueuePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBQueuePropertiesARM runs a test to see if a specific instance of SBQueuePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBQueuePropertiesARM(subject SBQueuePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBQueuePropertiesARM
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

// Generator of SBQueuePropertiesARM instances for property testing - lazily instantiated by
// SBQueuePropertiesARMGenerator()
var sbQueuePropertiesARMGenerator gopter.Gen

// SBQueuePropertiesARMGenerator returns a generator of SBQueuePropertiesARM instances for property testing.
func SBQueuePropertiesARMGenerator() gopter.Gen {
	if sbQueuePropertiesARMGenerator != nil {
		return sbQueuePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueuePropertiesARM(generators)
	sbQueuePropertiesARMGenerator = gen.Struct(reflect.TypeOf(SBQueuePropertiesARM{}), generators)

	return sbQueuePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForSBQueuePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBQueuePropertiesARM(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["DeadLetteringOnMessageExpiration"] = gen.PtrOf(gen.Bool())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["ForwardDeadLetteredMessagesTo"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardTo"] = gen.PtrOf(gen.AlphaString())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
}
