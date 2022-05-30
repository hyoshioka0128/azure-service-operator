// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20181130

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

func Test_UserAssignedIdentity_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_STATUSARM, UserAssignedIdentity_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_STATUSARM runs a test to see if a specific instance of UserAssignedIdentity_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_STATUSARM(subject UserAssignedIdentity_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_STATUSARM
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

// Generator of UserAssignedIdentity_STATUSARM instances for property testing - lazily instantiated by
// UserAssignedIdentity_STATUSARMGenerator()
var userAssignedIdentity_STATUSARMGenerator gopter.Gen

// UserAssignedIdentity_STATUSARMGenerator returns a generator of UserAssignedIdentity_STATUSARM instances for property testing.
// We first initialize userAssignedIdentity_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func UserAssignedIdentity_STATUSARMGenerator() gopter.Gen {
	if userAssignedIdentity_STATUSARMGenerator != nil {
		return userAssignedIdentity_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUSARM(generators)
	userAssignedIdentity_STATUSARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForUserAssignedIdentity_STATUSARM(generators)
	userAssignedIdentity_STATUSARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_STATUSARM{}), generators)

	return userAssignedIdentity_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForUserAssignedIdentity_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUserAssignedIdentity_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(UserAssignedIdentityProperties_STATUSARMGenerator())
}

func Test_UserAssignedIdentityProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityProperties_STATUSARM, UserAssignedIdentityProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityProperties_STATUSARM runs a test to see if a specific instance of UserAssignedIdentityProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityProperties_STATUSARM(subject UserAssignedIdentityProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityProperties_STATUSARM
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

// Generator of UserAssignedIdentityProperties_STATUSARM instances for property testing - lazily instantiated by
// UserAssignedIdentityProperties_STATUSARMGenerator()
var userAssignedIdentityProperties_STATUSARMGenerator gopter.Gen

// UserAssignedIdentityProperties_STATUSARMGenerator returns a generator of UserAssignedIdentityProperties_STATUSARM instances for property testing.
func UserAssignedIdentityProperties_STATUSARMGenerator() gopter.Gen {
	if userAssignedIdentityProperties_STATUSARMGenerator != nil {
		return userAssignedIdentityProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_STATUSARM(generators)
	userAssignedIdentityProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityProperties_STATUSARM{}), generators)

	return userAssignedIdentityProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}