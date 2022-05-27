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

func Test_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM, NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM runs a test to see if a specific instance of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(subject NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM
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

// Generator of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing -
// lazily instantiated by NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator()
var networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator gopter.Gen

// NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator returns a generator of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
// We first initialize networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator() gopter.Gen {
	if networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator != nil {
		return networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(generators)
	networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(generators)
	networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NetworkSecurityGroupPropertiesFormat_STATUSARMGenerator())
}

func Test_NetworkSecurityGroupPropertiesFormat_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupPropertiesFormat_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormat_STATUSARM, NetworkSecurityGroupPropertiesFormat_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormat_STATUSARM runs a test to see if a specific instance of NetworkSecurityGroupPropertiesFormat_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormat_STATUSARM(subject NetworkSecurityGroupPropertiesFormat_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupPropertiesFormat_STATUSARM
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

// Generator of NetworkSecurityGroupPropertiesFormat_STATUSARM instances for property testing - lazily instantiated by
// NetworkSecurityGroupPropertiesFormat_STATUSARMGenerator()
var networkSecurityGroupPropertiesFormat_STATUSARMGenerator gopter.Gen

// NetworkSecurityGroupPropertiesFormat_STATUSARMGenerator returns a generator of NetworkSecurityGroupPropertiesFormat_STATUSARM instances for property testing.
// We first initialize networkSecurityGroupPropertiesFormat_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupPropertiesFormat_STATUSARMGenerator() gopter.Gen {
	if networkSecurityGroupPropertiesFormat_STATUSARMGenerator != nil {
		return networkSecurityGroupPropertiesFormat_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUSARM(generators)
	networkSecurityGroupPropertiesFormat_STATUSARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupPropertiesFormat_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUSARM(generators)
	networkSecurityGroupPropertiesFormat_STATUSARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupPropertiesFormat_STATUSARM{}), generators)

	return networkSecurityGroupPropertiesFormat_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUSARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUSDeleting,
		ProvisioningState_STATUSFailed,
		ProvisioningState_STATUSSucceeded,
		ProvisioningState_STATUSUpdating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUSARM(gens map[string]gopter.Gen) {
	gens["DefaultSecurityRules"] = gen.SliceOf(SecurityRule_STATUSARMGenerator())
	gens["FlowLogs"] = gen.SliceOf(FlowLog_STATUSARMGenerator())
	gens["NetworkInterfaces"] = gen.SliceOf(NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator())
	gens["SecurityRules"] = gen.SliceOf(SecurityRule_STATUSARMGenerator())
	gens["Subnets"] = gen.SliceOf(Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator())
}

func Test_FlowLog_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlowLog_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlowLog_STATUSARM, FlowLog_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlowLog_STATUSARM runs a test to see if a specific instance of FlowLog_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlowLog_STATUSARM(subject FlowLog_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlowLog_STATUSARM
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

// Generator of FlowLog_STATUSARM instances for property testing - lazily instantiated by FlowLog_STATUSARMGenerator()
var flowLog_STATUSARMGenerator gopter.Gen

// FlowLog_STATUSARMGenerator returns a generator of FlowLog_STATUSARM instances for property testing.
func FlowLog_STATUSARMGenerator() gopter.Gen {
	if flowLog_STATUSARMGenerator != nil {
		return flowLog_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlowLog_STATUSARM(generators)
	flowLog_STATUSARMGenerator = gen.Struct(reflect.TypeOf(FlowLog_STATUSARM{}), generators)

	return flowLog_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForFlowLog_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlowLog_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM, NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM runs a test to see if a specific instance of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(subject NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM
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

// Generator of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing -
// lazily instantiated by NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator()
var networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator gopter.Gen

// NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator returns a generator of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
// We first initialize networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator() gopter.Gen {
	if networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator != nil {
		return networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(generators)
	networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(generators)
	AddRelatedPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(generators)
	networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUSARMGenerator())
}

func Test_SecurityRule_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRule_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRule_STATUSARM, SecurityRule_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRule_STATUSARM runs a test to see if a specific instance of SecurityRule_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRule_STATUSARM(subject SecurityRule_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRule_STATUSARM
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

// Generator of SecurityRule_STATUSARM instances for property testing - lazily instantiated by
// SecurityRule_STATUSARMGenerator()
var securityRule_STATUSARMGenerator gopter.Gen

// SecurityRule_STATUSARMGenerator returns a generator of SecurityRule_STATUSARM instances for property testing.
func SecurityRule_STATUSARMGenerator() gopter.Gen {
	if securityRule_STATUSARMGenerator != nil {
		return securityRule_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRule_STATUSARM(generators)
	securityRule_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SecurityRule_STATUSARM{}), generators)

	return securityRule_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRule_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRule_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM, Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM runs a test to see if a specific instance of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(subject Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM
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

// Generator of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator()
var subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator gopter.Gen

// Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator returns a generator of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
func Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator() gopter.Gen {
	if subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator != nil {
		return subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(generators)
	subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
