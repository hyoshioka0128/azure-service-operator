// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101storage

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

func Test_NetworkSecurityGroupsSecurityRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule, NetworkSecurityGroupsSecurityRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRule round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule(subject NetworkSecurityGroupsSecurityRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupsSecurityRule
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

// Generator of NetworkSecurityGroupsSecurityRule instances for property testing - lazily instantiated by
//NetworkSecurityGroupsSecurityRuleGenerator()
var networkSecurityGroupsSecurityRuleGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRuleGenerator returns a generator of NetworkSecurityGroupsSecurityRule instances for property testing.
func NetworkSecurityGroupsSecurityRuleGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRuleGenerator != nil {
		return networkSecurityGroupsSecurityRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule(generators)
	networkSecurityGroupsSecurityRuleGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule{}), generators)

	return networkSecurityGroupsSecurityRuleGenerator
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule(gens map[string]gopter.Gen) {
	gens["Spec"] = NetworkSecurityGroupsSecurityRule_SpecGenerator()
	gens["Status"] = SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator()
}

func Test_NetworkSecurityGroupsSecurityRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_Spec, NetworkSecurityGroupsSecurityRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_Spec runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_Spec(subject NetworkSecurityGroupsSecurityRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupsSecurityRule_Spec
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

// Generator of NetworkSecurityGroupsSecurityRule_Spec instances for property testing - lazily instantiated by
//NetworkSecurityGroupsSecurityRule_SpecGenerator()
var networkSecurityGroupsSecurityRule_specGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRule_SpecGenerator returns a generator of NetworkSecurityGroupsSecurityRule_Spec instances for property testing.
// We first initialize networkSecurityGroupsSecurityRule_specGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupsSecurityRule_SpecGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRule_specGenerator != nil {
		return networkSecurityGroupsSecurityRule_specGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(generators)
	networkSecurityGroupsSecurityRule_specGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(generators)
	networkSecurityGroupsSecurityRule_specGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule_Spec{}), generators)

	return networkSecurityGroupsSecurityRule_specGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpecGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpecGenerator())
}

func Test_SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded, SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded runs a test to see if a specific instance of SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(subject SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
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

// Generator of SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for property testing
//- lazily instantiated by SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator()
var securityRule_status_networkSecurityGroupsSecurityRule_subResourceEmbeddedGenerator gopter.Gen

// SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator returns a generator of SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for property testing.
// We first initialize securityRule_status_networkSecurityGroupsSecurityRule_subResourceEmbeddedGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator() gopter.Gen {
	if securityRule_status_networkSecurityGroupsSecurityRule_subResourceEmbeddedGenerator != nil {
		return securityRule_status_networkSecurityGroupsSecurityRule_subResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(generators)
	securityRule_status_networkSecurityGroupsSecurityRule_subResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(generators)
	AddRelatedPropertyGeneratorsForSecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(generators)
	securityRule_status_networkSecurityGroupsSecurityRule_subResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded{}), generators)

	return securityRule_status_networkSecurityGroupsSecurityRule_subResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator())
}

func Test_ApplicationSecurityGroupSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroupSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroupSpec, ApplicationSecurityGroupSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroupSpec runs a test to see if a specific instance of ApplicationSecurityGroupSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroupSpec(subject ApplicationSecurityGroupSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroupSpec
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

// Generator of ApplicationSecurityGroupSpec instances for property testing - lazily instantiated by
//ApplicationSecurityGroupSpecGenerator()
var applicationSecurityGroupSpecGenerator gopter.Gen

// ApplicationSecurityGroupSpecGenerator returns a generator of ApplicationSecurityGroupSpec instances for property testing.
func ApplicationSecurityGroupSpecGenerator() gopter.Gen {
	if applicationSecurityGroupSpecGenerator != nil {
		return applicationSecurityGroupSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec(generators)
	applicationSecurityGroupSpecGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroupSpec{}), generators)

	return applicationSecurityGroupSpecGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded, ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded runs a test to see if a specific instance of ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(subject ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
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

// Generator of ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for
//property testing - lazily instantiated by
//ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator()
var applicationSecurityGroup_status_networkSecurityGroupsSecurityRule_subResourceEmbeddedGenerator gopter.Gen

// ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator returns a generator of ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for property testing.
func ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator() gopter.Gen {
	if applicationSecurityGroup_status_networkSecurityGroupsSecurityRule_subResourceEmbeddedGenerator != nil {
		return applicationSecurityGroup_status_networkSecurityGroupsSecurityRule_subResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(generators)
	applicationSecurityGroup_status_networkSecurityGroupsSecurityRule_subResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded{}), generators)

	return applicationSecurityGroup_status_networkSecurityGroupsSecurityRule_subResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
