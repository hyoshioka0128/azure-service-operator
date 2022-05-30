// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201storage

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

func Test_RedisLinkedServer_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServer, RedisLinkedServerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServer runs a test to see if a specific instance of RedisLinkedServer round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServer(subject RedisLinkedServer) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer
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

// Generator of RedisLinkedServer instances for property testing - lazily instantiated by RedisLinkedServerGenerator()
var redisLinkedServerGenerator gopter.Gen

// RedisLinkedServerGenerator returns a generator of RedisLinkedServer instances for property testing.
func RedisLinkedServerGenerator() gopter.Gen {
	if redisLinkedServerGenerator != nil {
		return redisLinkedServerGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisLinkedServer(generators)
	redisLinkedServerGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer{}), generators)

	return redisLinkedServerGenerator
}

// AddRelatedPropertyGeneratorsForRedisLinkedServer is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisLinkedServer(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisLinkedServer_SpecGenerator()
	gens["Status"] = RedisLinkedServerWithProperties_STATUSGenerator()
}

func Test_RedisLinkedServerWithProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServerWithProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerWithProperties_STATUS, RedisLinkedServerWithProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerWithProperties_STATUS runs a test to see if a specific instance of RedisLinkedServerWithProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerWithProperties_STATUS(subject RedisLinkedServerWithProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServerWithProperties_STATUS
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

// Generator of RedisLinkedServerWithProperties_STATUS instances for property testing - lazily instantiated by
// RedisLinkedServerWithProperties_STATUSGenerator()
var redisLinkedServerWithProperties_STATUSGenerator gopter.Gen

// RedisLinkedServerWithProperties_STATUSGenerator returns a generator of RedisLinkedServerWithProperties_STATUS instances for property testing.
func RedisLinkedServerWithProperties_STATUSGenerator() gopter.Gen {
	if redisLinkedServerWithProperties_STATUSGenerator != nil {
		return redisLinkedServerWithProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerWithProperties_STATUS(generators)
	redisLinkedServerWithProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServerWithProperties_STATUS{}), generators)

	return redisLinkedServerWithProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServerWithProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServerWithProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheId"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ServerRole"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisLinkedServer_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServer_Spec, RedisLinkedServer_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServer_Spec runs a test to see if a specific instance of RedisLinkedServer_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServer_Spec(subject RedisLinkedServer_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer_Spec
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

// Generator of RedisLinkedServer_Spec instances for property testing - lazily instantiated by
// RedisLinkedServer_SpecGenerator()
var redisLinkedServer_SpecGenerator gopter.Gen

// RedisLinkedServer_SpecGenerator returns a generator of RedisLinkedServer_Spec instances for property testing.
func RedisLinkedServer_SpecGenerator() gopter.Gen {
	if redisLinkedServer_SpecGenerator != nil {
		return redisLinkedServer_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec(generators)
	redisLinkedServer_SpecGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_Spec{}), generators)

	return redisLinkedServer_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["ServerRole"] = gen.PtrOf(gen.AlphaString())
}
