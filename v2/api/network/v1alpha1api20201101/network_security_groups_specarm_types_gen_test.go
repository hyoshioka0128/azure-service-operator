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

func Test_NetworkSecurityGroups_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroups_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSPECARM, NetworkSecurityGroupsSPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSPECARM runs a test to see if a specific instance of NetworkSecurityGroups_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSPECARM(subject NetworkSecurityGroups_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroups_SPECARM
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

// Generator of NetworkSecurityGroups_SPECARM instances for property testing - lazily instantiated by
//NetworkSecurityGroupsSPECARMGenerator()
var networkSecurityGroupsSPECARMGenerator gopter.Gen

// NetworkSecurityGroupsSPECARMGenerator returns a generator of NetworkSecurityGroups_SPECARM instances for property testing.
// We first initialize networkSecurityGroupsSPECARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupsSPECARMGenerator() gopter.Gen {
	if networkSecurityGroupsSPECARMGenerator != nil {
		return networkSecurityGroupsSPECARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSPECARM(generators)
	networkSecurityGroupsSPECARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSPECARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSPECARM(generators)
	networkSecurityGroupsSPECARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_SPECARM{}), generators)

	return networkSecurityGroupsSPECARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NetworkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator())
}

func Test_NetworkSecurityGroupPropertiesFormat_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupPropertiesFormat_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARM, NetworkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARM runs a test to see if a specific instance of NetworkSecurityGroupPropertiesFormat_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARM(subject NetworkSecurityGroupPropertiesFormat_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupPropertiesFormat_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM
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

// Generator of NetworkSecurityGroupPropertiesFormat_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM instances for
//property testing - lazily instantiated by
//NetworkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator()
var networkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator gopter.Gen

// NetworkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator returns a generator of NetworkSecurityGroupPropertiesFormat_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
func NetworkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator() gopter.Gen {
	if networkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator != nil {
		return networkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	networkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupPropertiesFormat_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return networkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatSpecNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["SecurityRules"] = gen.SliceOf(SecurityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator())
}

func Test_SecurityRule_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRule_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARM, SecurityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARM runs a test to see if a specific instance of SecurityRule_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARM(subject SecurityRule_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRule_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM
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

// Generator of SecurityRule_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing - lazily
//instantiated by SecurityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator()
var securityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator gopter.Gen

// SecurityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator returns a generator of SecurityRule_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
func SecurityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator() gopter.Gen {
	if securityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator != nil {
		return securityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	securityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(SecurityRule_Spec_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return securityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRuleSpecNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}