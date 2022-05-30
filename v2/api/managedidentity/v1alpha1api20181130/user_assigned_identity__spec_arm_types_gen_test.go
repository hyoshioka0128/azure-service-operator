// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20181130

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

func Test_UserAssignedIdentity_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_SpecARM, UserAssignedIdentity_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_SpecARM runs a test to see if a specific instance of UserAssignedIdentity_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_SpecARM(subject UserAssignedIdentity_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_SpecARM
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

// Generator of UserAssignedIdentity_SpecARM instances for property testing - lazily instantiated by
// UserAssignedIdentity_SpecARMGenerator()
var userAssignedIdentity_SpecARMGenerator gopter.Gen

// UserAssignedIdentity_SpecARMGenerator returns a generator of UserAssignedIdentity_SpecARM instances for property testing.
func UserAssignedIdentity_SpecARMGenerator() gopter.Gen {
	if userAssignedIdentity_SpecARMGenerator != nil {
		return userAssignedIdentity_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_SpecARM(generators)
	userAssignedIdentity_SpecARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_SpecARM{}), generators)

	return userAssignedIdentity_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
