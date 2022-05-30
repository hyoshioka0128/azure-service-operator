// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

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

func Test_DomainsTopic_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainsTopic_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopic_STATUSARM, DomainsTopic_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopic_STATUSARM runs a test to see if a specific instance of DomainsTopic_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopic_STATUSARM(subject DomainsTopic_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainsTopic_STATUSARM
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

// Generator of DomainsTopic_STATUSARM instances for property testing - lazily instantiated by
// DomainsTopic_STATUSARMGenerator()
var domainsTopic_STATUSARMGenerator gopter.Gen

// DomainsTopic_STATUSARMGenerator returns a generator of DomainsTopic_STATUSARM instances for property testing.
// We first initialize domainsTopic_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainsTopic_STATUSARMGenerator() gopter.Gen {
	if domainsTopic_STATUSARMGenerator != nil {
		return domainsTopic_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopic_STATUSARM(generators)
	domainsTopic_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DomainsTopic_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopic_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForDomainsTopic_STATUSARM(generators)
	domainsTopic_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DomainsTopic_STATUSARM{}), generators)

	return domainsTopic_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDomainsTopic_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainsTopic_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDomainsTopic_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainsTopic_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DomainTopicProperties_STATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_DomainTopicProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainTopicProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainTopicProperties_STATUSARM, DomainTopicProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainTopicProperties_STATUSARM runs a test to see if a specific instance of DomainTopicProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainTopicProperties_STATUSARM(subject DomainTopicProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainTopicProperties_STATUSARM
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

// Generator of DomainTopicProperties_STATUSARM instances for property testing - lazily instantiated by
// DomainTopicProperties_STATUSARMGenerator()
var domainTopicProperties_STATUSARMGenerator gopter.Gen

// DomainTopicProperties_STATUSARMGenerator returns a generator of DomainTopicProperties_STATUSARM instances for property testing.
func DomainTopicProperties_STATUSARMGenerator() gopter.Gen {
	if domainTopicProperties_STATUSARMGenerator != nil {
		return domainTopicProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainTopicProperties_STATUSARM(generators)
	domainTopicProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DomainTopicProperties_STATUSARM{}), generators)

	return domainTopicProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDomainTopicProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainTopicProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DomainTopicProperties_ProvisioningState_Canceled_STATUS,
		DomainTopicProperties_ProvisioningState_Creating_STATUS,
		DomainTopicProperties_ProvisioningState_Deleting_STATUS,
		DomainTopicProperties_ProvisioningState_Failed_STATUS,
		DomainTopicProperties_ProvisioningState_Succeeded_STATUS,
		DomainTopicProperties_ProvisioningState_Updating_STATUS))
}
