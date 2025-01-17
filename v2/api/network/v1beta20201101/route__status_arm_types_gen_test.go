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

func Test_Route_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Route_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteStatusARM, RouteStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteStatusARM runs a test to see if a specific instance of Route_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteStatusARM(subject Route_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Route_StatusARM
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

// Generator of Route_StatusARM instances for property testing - lazily instantiated by RouteStatusARMGenerator()
var routeStatusARMGenerator gopter.Gen

// RouteStatusARMGenerator returns a generator of Route_StatusARM instances for property testing.
// We first initialize routeStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RouteStatusARMGenerator() gopter.Gen {
	if routeStatusARMGenerator != nil {
		return routeStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteStatusARM(generators)
	routeStatusARMGenerator = gen.Struct(reflect.TypeOf(Route_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteStatusARM(generators)
	AddRelatedPropertyGeneratorsForRouteStatusARM(generators)
	routeStatusARMGenerator = gen.Struct(reflect.TypeOf(Route_StatusARM{}), generators)

	return routeStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRouteStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteStatusARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRouteStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RoutePropertiesFormatStatusARMGenerator())
}

func Test_RoutePropertiesFormat_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoutePropertiesFormat_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoutePropertiesFormatStatusARM, RoutePropertiesFormatStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoutePropertiesFormatStatusARM runs a test to see if a specific instance of RoutePropertiesFormat_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoutePropertiesFormatStatusARM(subject RoutePropertiesFormat_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoutePropertiesFormat_StatusARM
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

// Generator of RoutePropertiesFormat_StatusARM instances for property testing - lazily instantiated by
// RoutePropertiesFormatStatusARMGenerator()
var routePropertiesFormatStatusARMGenerator gopter.Gen

// RoutePropertiesFormatStatusARMGenerator returns a generator of RoutePropertiesFormat_StatusARM instances for property testing.
func RoutePropertiesFormatStatusARMGenerator() gopter.Gen {
	if routePropertiesFormatStatusARMGenerator != nil {
		return routePropertiesFormatStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoutePropertiesFormatStatusARM(generators)
	routePropertiesFormatStatusARMGenerator = gen.Struct(reflect.TypeOf(RoutePropertiesFormat_StatusARM{}), generators)

	return routePropertiesFormatStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRoutePropertiesFormatStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoutePropertiesFormatStatusARM(gens map[string]gopter.Gen) {
	gens["AddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["HasBgpOverride"] = gen.PtrOf(gen.Bool())
	gens["NextHopIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["NextHopType"] = gen.PtrOf(gen.OneConstOf(
		RouteNextHopType_StatusInternet,
		RouteNextHopType_StatusNone,
		RouteNextHopType_StatusVirtualAppliance,
		RouteNextHopType_StatusVirtualNetworkGateway,
		RouteNextHopType_StatusVnetLocal))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_StatusDeleting,
		ProvisioningState_StatusFailed,
		ProvisioningState_StatusSucceeded,
		ProvisioningState_StatusUpdating))
}
