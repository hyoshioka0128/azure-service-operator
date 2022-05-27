// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

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

func Test_VirtualNetwork_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetwork_STATUSARM, VirtualNetwork_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetwork_STATUSARM runs a test to see if a specific instance of VirtualNetwork_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetwork_STATUSARM(subject VirtualNetwork_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_STATUSARM
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

// Generator of VirtualNetwork_STATUSARM instances for property testing - lazily instantiated by
// VirtualNetwork_STATUSARMGenerator()
var virtualNetwork_STATUSARMGenerator gopter.Gen

// VirtualNetwork_STATUSARMGenerator returns a generator of VirtualNetwork_STATUSARM instances for property testing.
// We first initialize virtualNetwork_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetwork_STATUSARMGenerator() gopter.Gen {
	if virtualNetwork_STATUSARMGenerator != nil {
		return virtualNetwork_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_STATUSARM(generators)
	virtualNetwork_STATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetwork_STATUSARM(generators)
	virtualNetwork_STATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_STATUSARM{}), generators)

	return virtualNetwork_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetwork_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetwork_STATUSARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetwork_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetwork_STATUSARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUSARMGenerator())
	gens["Properties"] = gen.PtrOf(VirtualNetworkPropertiesFormat_STATUSARMGenerator())
}

func Test_VirtualNetworkPropertiesFormat_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPropertiesFormat_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPropertiesFormat_STATUSARM, VirtualNetworkPropertiesFormat_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPropertiesFormat_STATUSARM runs a test to see if a specific instance of VirtualNetworkPropertiesFormat_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPropertiesFormat_STATUSARM(subject VirtualNetworkPropertiesFormat_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPropertiesFormat_STATUSARM
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

// Generator of VirtualNetworkPropertiesFormat_STATUSARM instances for property testing - lazily instantiated by
// VirtualNetworkPropertiesFormat_STATUSARMGenerator()
var virtualNetworkPropertiesFormat_STATUSARMGenerator gopter.Gen

// VirtualNetworkPropertiesFormat_STATUSARMGenerator returns a generator of VirtualNetworkPropertiesFormat_STATUSARM instances for property testing.
// We first initialize virtualNetworkPropertiesFormat_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPropertiesFormat_STATUSARMGenerator() gopter.Gen {
	if virtualNetworkPropertiesFormat_STATUSARMGenerator != nil {
		return virtualNetworkPropertiesFormat_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUSARM(generators)
	virtualNetworkPropertiesFormat_STATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPropertiesFormat_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUSARM(generators)
	virtualNetworkPropertiesFormat_STATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPropertiesFormat_STATUSARM{}), generators)

	return virtualNetworkPropertiesFormat_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUSARM(gens map[string]gopter.Gen) {
	gens["EnableDdosProtection"] = gen.PtrOf(gen.Bool())
	gens["EnableVmProtection"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUSDeleting,
		ProvisioningState_STATUSFailed,
		ProvisioningState_STATUSSucceeded,
		ProvisioningState_STATUSUpdating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUSARM(gens map[string]gopter.Gen) {
	gens["AddressSpace"] = gen.PtrOf(AddressSpace_STATUSARMGenerator())
	gens["BgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunities_STATUSARMGenerator())
	gens["DdosProtectionPlan"] = gen.PtrOf(SubResource_STATUSARMGenerator())
	gens["DhcpOptions"] = gen.PtrOf(DhcpOptions_STATUSARMGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResource_STATUSARMGenerator())
}

func Test_DhcpOptions_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DhcpOptions_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDhcpOptions_STATUSARM, DhcpOptions_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDhcpOptions_STATUSARM runs a test to see if a specific instance of DhcpOptions_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDhcpOptions_STATUSARM(subject DhcpOptions_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DhcpOptions_STATUSARM
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

// Generator of DhcpOptions_STATUSARM instances for property testing - lazily instantiated by
// DhcpOptions_STATUSARMGenerator()
var dhcpOptions_STATUSARMGenerator gopter.Gen

// DhcpOptions_STATUSARMGenerator returns a generator of DhcpOptions_STATUSARM instances for property testing.
func DhcpOptions_STATUSARMGenerator() gopter.Gen {
	if dhcpOptions_STATUSARMGenerator != nil {
		return dhcpOptions_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDhcpOptions_STATUSARM(generators)
	dhcpOptions_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DhcpOptions_STATUSARM{}), generators)

	return dhcpOptions_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDhcpOptions_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDhcpOptions_STATUSARM(gens map[string]gopter.Gen) {
	gens["DnsServers"] = gen.SliceOf(gen.AlphaString())
}

func Test_VirtualNetworkBgpCommunities_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkBgpCommunities_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUSARM, VirtualNetworkBgpCommunities_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUSARM runs a test to see if a specific instance of VirtualNetworkBgpCommunities_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUSARM(subject VirtualNetworkBgpCommunities_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkBgpCommunities_STATUSARM
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

// Generator of VirtualNetworkBgpCommunities_STATUSARM instances for property testing - lazily instantiated by
// VirtualNetworkBgpCommunities_STATUSARMGenerator()
var virtualNetworkBgpCommunities_STATUSARMGenerator gopter.Gen

// VirtualNetworkBgpCommunities_STATUSARMGenerator returns a generator of VirtualNetworkBgpCommunities_STATUSARM instances for property testing.
func VirtualNetworkBgpCommunities_STATUSARMGenerator() gopter.Gen {
	if virtualNetworkBgpCommunities_STATUSARMGenerator != nil {
		return virtualNetworkBgpCommunities_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUSARM(generators)
	virtualNetworkBgpCommunities_STATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkBgpCommunities_STATUSARM{}), generators)

	return virtualNetworkBgpCommunities_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUSARM(gens map[string]gopter.Gen) {
	gens["RegionalCommunity"] = gen.PtrOf(gen.AlphaString())
	gens["VirtualNetworkCommunity"] = gen.PtrOf(gen.AlphaString())
}
