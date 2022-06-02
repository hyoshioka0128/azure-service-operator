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

func Test_NetworkSecurityGroup_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroup_SpecARM, NetworkSecurityGroup_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroup_SpecARM runs a test to see if a specific instance of NetworkSecurityGroup_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroup_SpecARM(subject NetworkSecurityGroup_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup_SpecARM
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

// Generator of NetworkSecurityGroup_SpecARM instances for property testing - lazily instantiated by
// NetworkSecurityGroup_SpecARMGenerator()
var networkSecurityGroup_SpecARMGenerator gopter.Gen

// NetworkSecurityGroup_SpecARMGenerator returns a generator of NetworkSecurityGroup_SpecARM instances for property testing.
// We first initialize networkSecurityGroup_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroup_SpecARMGenerator() gopter.Gen {
	if networkSecurityGroup_SpecARMGenerator != nil {
		return networkSecurityGroup_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_SpecARM(generators)
	networkSecurityGroup_SpecARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_SpecARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroup_SpecARM(generators)
	networkSecurityGroup_SpecARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_SpecARM{}), generators)

	return networkSecurityGroup_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroup_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroup_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroup_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroup_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator())
}

func Test_NetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM, NetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM runs a test to see if a specific instance of NetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM(subject NetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM
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

// Generator of NetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property
// testing - lazily instantiated by
// NetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator()
var networkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator gopter.Gen

// NetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator returns a generator of NetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
func NetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator() gopter.Gen {
	if networkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator != nil {
		return networkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM(generators)
	networkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return networkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_NetworkSecurityGroup_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["SecurityRules"] = gen.SliceOf(SecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator())
}

func Test_SecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM, SecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM runs a test to see if a specific instance of SecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM(subject SecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM
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

// Generator of SecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by SecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator()
var securityRule_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator gopter.Gen

// SecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator returns a generator of SecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
func SecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator() gopter.Gen {
	if securityRule_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator != nil {
		return securityRule_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM(generators)
	securityRule_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(SecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return securityRule_NetworkSecurityGroup_SubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRule_NetworkSecurityGroup_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}