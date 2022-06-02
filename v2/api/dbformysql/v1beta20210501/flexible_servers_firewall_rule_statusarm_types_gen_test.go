// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

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

func Test_FlexibleServersFirewallRule_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersFirewallRule_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersFirewallRule_STATUSARM, FlexibleServersFirewallRule_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersFirewallRule_STATUSARM runs a test to see if a specific instance of FlexibleServersFirewallRule_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersFirewallRule_STATUSARM(subject FlexibleServersFirewallRule_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersFirewallRule_STATUSARM
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

// Generator of FlexibleServersFirewallRule_STATUSARM instances for property testing - lazily instantiated by
// FlexibleServersFirewallRule_STATUSARMGenerator()
var flexibleServersFirewallRule_STATUSARMGenerator gopter.Gen

// FlexibleServersFirewallRule_STATUSARMGenerator returns a generator of FlexibleServersFirewallRule_STATUSARM instances for property testing.
// We first initialize flexibleServersFirewallRule_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersFirewallRule_STATUSARMGenerator() gopter.Gen {
	if flexibleServersFirewallRule_STATUSARMGenerator != nil {
		return flexibleServersFirewallRule_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_STATUSARM(generators)
	flexibleServersFirewallRule_STATUSARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRule_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule_STATUSARM(generators)
	flexibleServersFirewallRule_STATUSARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRule_STATUSARM{}), generators)

	return flexibleServersFirewallRule_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FirewallRuleProperties_STATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_FirewallRuleProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FirewallRuleProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFirewallRuleProperties_STATUSARM, FirewallRuleProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFirewallRuleProperties_STATUSARM runs a test to see if a specific instance of FirewallRuleProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFirewallRuleProperties_STATUSARM(subject FirewallRuleProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FirewallRuleProperties_STATUSARM
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

// Generator of FirewallRuleProperties_STATUSARM instances for property testing - lazily instantiated by
// FirewallRuleProperties_STATUSARMGenerator()
var firewallRuleProperties_STATUSARMGenerator gopter.Gen

// FirewallRuleProperties_STATUSARMGenerator returns a generator of FirewallRuleProperties_STATUSARM instances for property testing.
func FirewallRuleProperties_STATUSARMGenerator() gopter.Gen {
	if firewallRuleProperties_STATUSARMGenerator != nil {
		return firewallRuleProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRuleProperties_STATUSARM(generators)
	firewallRuleProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(FirewallRuleProperties_STATUSARM{}), generators)

	return firewallRuleProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForFirewallRuleProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFirewallRuleProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
}