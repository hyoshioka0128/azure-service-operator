// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

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

func Test_NamespacesAuthorizationRule_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRule_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRule_STATUSARM, NamespacesAuthorizationRule_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRule_STATUSARM runs a test to see if a specific instance of NamespacesAuthorizationRule_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRule_STATUSARM(subject NamespacesAuthorizationRule_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRule_STATUSARM
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

// Generator of NamespacesAuthorizationRule_STATUSARM instances for property testing - lazily instantiated by
// NamespacesAuthorizationRule_STATUSARMGenerator()
var namespacesAuthorizationRule_STATUSARMGenerator gopter.Gen

// NamespacesAuthorizationRule_STATUSARMGenerator returns a generator of NamespacesAuthorizationRule_STATUSARM instances for property testing.
// We first initialize namespacesAuthorizationRule_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesAuthorizationRule_STATUSARMGenerator() gopter.Gen {
	if namespacesAuthorizationRule_STATUSARMGenerator != nil {
		return namespacesAuthorizationRule_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_STATUSARM(generators)
	namespacesAuthorizationRule_STATUSARMGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_STATUSARM(generators)
	namespacesAuthorizationRule_STATUSARMGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule_STATUSARM{}), generators)

	return namespacesAuthorizationRule_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NamespacesAuthorizationRule_Properties_STATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_NamespacesAuthorizationRule_Properties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRule_Properties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRule_Properties_STATUSARM, NamespacesAuthorizationRule_Properties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRule_Properties_STATUSARM runs a test to see if a specific instance of NamespacesAuthorizationRule_Properties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRule_Properties_STATUSARM(subject NamespacesAuthorizationRule_Properties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRule_Properties_STATUSARM
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

// Generator of NamespacesAuthorizationRule_Properties_STATUSARM instances for property testing - lazily instantiated by
// NamespacesAuthorizationRule_Properties_STATUSARMGenerator()
var namespacesAuthorizationRule_Properties_STATUSARMGenerator gopter.Gen

// NamespacesAuthorizationRule_Properties_STATUSARMGenerator returns a generator of NamespacesAuthorizationRule_Properties_STATUSARM instances for property testing.
func NamespacesAuthorizationRule_Properties_STATUSARMGenerator() gopter.Gen {
	if namespacesAuthorizationRule_Properties_STATUSARMGenerator != nil {
		return namespacesAuthorizationRule_Properties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Properties_STATUSARM(generators)
	namespacesAuthorizationRule_Properties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule_Properties_STATUSARM{}), generators)

	return namespacesAuthorizationRule_Properties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Properties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Properties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(NamespacesAuthorizationRule_Properties_Rights_STATUSListen, NamespacesAuthorizationRule_Properties_Rights_STATUSManage, NamespacesAuthorizationRule_Properties_Rights_STATUSSend))
}
