// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

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

func Test_RedisFirewallRule_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRule_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRule_SpecARM, RedisFirewallRule_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRule_SpecARM runs a test to see if a specific instance of RedisFirewallRule_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRule_SpecARM(subject RedisFirewallRule_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRule_SpecARM
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

// Generator of RedisFirewallRule_SpecARM instances for property testing - lazily instantiated by
// RedisFirewallRule_SpecARMGenerator()
var redisFirewallRule_SpecARMGenerator gopter.Gen

// RedisFirewallRule_SpecARMGenerator returns a generator of RedisFirewallRule_SpecARM instances for property testing.
// We first initialize redisFirewallRule_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisFirewallRule_SpecARMGenerator() gopter.Gen {
	if redisFirewallRule_SpecARMGenerator != nil {
		return redisFirewallRule_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRule_SpecARM(generators)
	redisFirewallRule_SpecARMGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRule_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRule_SpecARM(generators)
	AddRelatedPropertyGeneratorsForRedisFirewallRule_SpecARM(generators)
	redisFirewallRule_SpecARMGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRule_SpecARM{}), generators)

	return redisFirewallRule_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisFirewallRule_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisFirewallRule_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRedisFirewallRule_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisFirewallRule_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisFirewallRulePropertiesARMGenerator())
}

func Test_RedisFirewallRulePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRulePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRulePropertiesARM, RedisFirewallRulePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRulePropertiesARM runs a test to see if a specific instance of RedisFirewallRulePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRulePropertiesARM(subject RedisFirewallRulePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRulePropertiesARM
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

// Generator of RedisFirewallRulePropertiesARM instances for property testing - lazily instantiated by
// RedisFirewallRulePropertiesARMGenerator()
var redisFirewallRulePropertiesARMGenerator gopter.Gen

// RedisFirewallRulePropertiesARMGenerator returns a generator of RedisFirewallRulePropertiesARM instances for property testing.
func RedisFirewallRulePropertiesARMGenerator() gopter.Gen {
	if redisFirewallRulePropertiesARMGenerator != nil {
		return redisFirewallRulePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRulePropertiesARM(generators)
	redisFirewallRulePropertiesARMGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRulePropertiesARM{}), generators)

	return redisFirewallRulePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisFirewallRulePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisFirewallRulePropertiesARM(gens map[string]gopter.Gen) {
	gens["EndIP"] = gen.PtrOf(gen.AlphaString())
	gens["StartIP"] = gen.PtrOf(gen.AlphaString())
}
