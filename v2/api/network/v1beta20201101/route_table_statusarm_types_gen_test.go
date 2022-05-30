// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

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

func Test_RouteTable_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTable_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTable_STATUSARM, RouteTable_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTable_STATUSARM runs a test to see if a specific instance of RouteTable_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTable_STATUSARM(subject RouteTable_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTable_STATUSARM
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

// Generator of RouteTable_STATUSARM instances for property testing - lazily instantiated by
// RouteTable_STATUSARMGenerator()
var routeTable_STATUSARMGenerator gopter.Gen

// RouteTable_STATUSARMGenerator returns a generator of RouteTable_STATUSARM instances for property testing.
// We first initialize routeTable_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RouteTable_STATUSARMGenerator() gopter.Gen {
	if routeTable_STATUSARMGenerator != nil {
		return routeTable_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_STATUSARM(generators)
	routeTable_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RouteTable_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForRouteTable_STATUSARM(generators)
	routeTable_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RouteTable_STATUSARM{}), generators)

	return routeTable_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRouteTable_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTable_STATUSARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRouteTable_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTable_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RouteTablePropertiesFormat_STATUSARMGenerator())
}

func Test_RouteTablePropertiesFormat_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTablePropertiesFormat_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTablePropertiesFormat_STATUSARM, RouteTablePropertiesFormat_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTablePropertiesFormat_STATUSARM runs a test to see if a specific instance of RouteTablePropertiesFormat_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTablePropertiesFormat_STATUSARM(subject RouteTablePropertiesFormat_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTablePropertiesFormat_STATUSARM
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

// Generator of RouteTablePropertiesFormat_STATUSARM instances for property testing - lazily instantiated by
// RouteTablePropertiesFormat_STATUSARMGenerator()
var routeTablePropertiesFormat_STATUSARMGenerator gopter.Gen

// RouteTablePropertiesFormat_STATUSARMGenerator returns a generator of RouteTablePropertiesFormat_STATUSARM instances for property testing.
func RouteTablePropertiesFormat_STATUSARMGenerator() gopter.Gen {
	if routeTablePropertiesFormat_STATUSARMGenerator != nil {
		return routeTablePropertiesFormat_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTablePropertiesFormat_STATUSARM(generators)
	routeTablePropertiesFormat_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RouteTablePropertiesFormat_STATUSARM{}), generators)

	return routeTablePropertiesFormat_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRouteTablePropertiesFormat_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTablePropertiesFormat_STATUSARM(gens map[string]gopter.Gen) {
	gens["DisableBgpRoutePropagation"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_Deleting_STATUS,
		ProvisioningState_Failed_STATUS,
		ProvisioningState_Succeeded_STATUS,
		ProvisioningState_Updating_STATUS))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}