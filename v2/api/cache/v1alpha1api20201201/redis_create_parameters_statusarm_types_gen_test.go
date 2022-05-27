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

func Test_RedisCreateParameters_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisCreateParameters_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisCreateParameters_STATUSARM, RedisCreateParameters_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisCreateParameters_STATUSARM runs a test to see if a specific instance of RedisCreateParameters_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisCreateParameters_STATUSARM(subject RedisCreateParameters_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisCreateParameters_STATUSARM
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

// Generator of RedisCreateParameters_STATUSARM instances for property testing - lazily instantiated by
// RedisCreateParameters_STATUSARMGenerator()
var redisCreateParameters_STATUSARMGenerator gopter.Gen

// RedisCreateParameters_STATUSARMGenerator returns a generator of RedisCreateParameters_STATUSARM instances for property testing.
// We first initialize redisCreateParameters_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisCreateParameters_STATUSARMGenerator() gopter.Gen {
	if redisCreateParameters_STATUSARMGenerator != nil {
		return redisCreateParameters_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreateParameters_STATUSARM(generators)
	redisCreateParameters_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisCreateParameters_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreateParameters_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForRedisCreateParameters_STATUSARM(generators)
	redisCreateParameters_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisCreateParameters_STATUSARM{}), generators)

	return redisCreateParameters_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisCreateParameters_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisCreateParameters_STATUSARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisCreateParameters_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisCreateParameters_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisCreateProperties_STATUSARMGenerator())
}

func Test_RedisCreateProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisCreateProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisCreateProperties_STATUSARM, RedisCreateProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisCreateProperties_STATUSARM runs a test to see if a specific instance of RedisCreateProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisCreateProperties_STATUSARM(subject RedisCreateProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisCreateProperties_STATUSARM
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

// Generator of RedisCreateProperties_STATUSARM instances for property testing - lazily instantiated by
// RedisCreateProperties_STATUSARMGenerator()
var redisCreateProperties_STATUSARMGenerator gopter.Gen

// RedisCreateProperties_STATUSARMGenerator returns a generator of RedisCreateProperties_STATUSARM instances for property testing.
// We first initialize redisCreateProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisCreateProperties_STATUSARMGenerator() gopter.Gen {
	if redisCreateProperties_STATUSARMGenerator != nil {
		return redisCreateProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreateProperties_STATUSARM(generators)
	redisCreateProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisCreateProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreateProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForRedisCreateProperties_STATUSARM(generators)
	redisCreateProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisCreateProperties_STATUSARM{}), generators)

	return redisCreateProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisCreateProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisCreateProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(RedisCreateProperties_MinimumTlsVersion_STATUS10, RedisCreateProperties_MinimumTlsVersion_STATUS11, RedisCreateProperties_MinimumTlsVersion_STATUS12))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(RedisCreateProperties_PublicNetworkAccess_STATUSDisabled, RedisCreateProperties_PublicNetworkAccess_STATUSEnabled))
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicasPerMaster"] = gen.PtrOf(gen.Int())
	gens["ReplicasPerPrimary"] = gen.PtrOf(gen.Int())
	gens["ShardCount"] = gen.PtrOf(gen.Int())
	gens["StaticIP"] = gen.PtrOf(gen.AlphaString())
	gens["SubnetId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantSettings"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisCreateProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisCreateProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["RedisConfiguration"] = gen.PtrOf(RedisCreateProperties_RedisConfiguration_STATUSARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSARMGenerator())
}

func Test_RedisCreateProperties_RedisConfiguration_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisCreateProperties_RedisConfiguration_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisCreateProperties_RedisConfiguration_STATUSARM, RedisCreateProperties_RedisConfiguration_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisCreateProperties_RedisConfiguration_STATUSARM runs a test to see if a specific instance of RedisCreateProperties_RedisConfiguration_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisCreateProperties_RedisConfiguration_STATUSARM(subject RedisCreateProperties_RedisConfiguration_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisCreateProperties_RedisConfiguration_STATUSARM
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

// Generator of RedisCreateProperties_RedisConfiguration_STATUSARM instances for property testing - lazily instantiated
// by RedisCreateProperties_RedisConfiguration_STATUSARMGenerator()
var redisCreateProperties_RedisConfiguration_STATUSARMGenerator gopter.Gen

// RedisCreateProperties_RedisConfiguration_STATUSARMGenerator returns a generator of RedisCreateProperties_RedisConfiguration_STATUSARM instances for property testing.
func RedisCreateProperties_RedisConfiguration_STATUSARMGenerator() gopter.Gen {
	if redisCreateProperties_RedisConfiguration_STATUSARMGenerator != nil {
		return redisCreateProperties_RedisConfiguration_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreateProperties_RedisConfiguration_STATUSARM(generators)
	redisCreateProperties_RedisConfiguration_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisCreateProperties_RedisConfiguration_STATUSARM{}), generators)

	return redisCreateProperties_RedisConfiguration_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisCreateProperties_RedisConfiguration_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisCreateProperties_RedisConfiguration_STATUSARM(gens map[string]gopter.Gen) {
	gens["AdditionalProperties"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["AofStorageConnectionString0"] = gen.PtrOf(gen.AlphaString())
	gens["AofStorageConnectionString1"] = gen.PtrOf(gen.AlphaString())
	gens["Maxclients"] = gen.PtrOf(gen.AlphaString())
	gens["MaxfragmentationmemoryReserved"] = gen.PtrOf(gen.AlphaString())
	gens["MaxmemoryDelta"] = gen.PtrOf(gen.AlphaString())
	gens["MaxmemoryPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["MaxmemoryReserved"] = gen.PtrOf(gen.AlphaString())
	gens["RdbBackupEnabled"] = gen.PtrOf(gen.AlphaString())
	gens["RdbBackupFrequency"] = gen.PtrOf(gen.AlphaString())
	gens["RdbBackupMaxSnapshotCount"] = gen.PtrOf(gen.AlphaString())
	gens["RdbStorageConnectionString"] = gen.PtrOf(gen.AlphaString())
	gens["ZonalConfiguration"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUSARM, Sku_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUSARM runs a test to see if a specific instance of Sku_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUSARM(subject Sku_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUSARM
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

// Generator of Sku_STATUSARM instances for property testing - lazily instantiated by Sku_STATUSARMGenerator()
var sku_STATUSARMGenerator gopter.Gen

// Sku_STATUSARMGenerator returns a generator of Sku_STATUSARM instances for property testing.
func Sku_STATUSARMGenerator() gopter.Gen {
	if sku_STATUSARMGenerator != nil {
		return sku_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUSARM(generators)
	sku_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUSARM{}), generators)

	return sku_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUSARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.OneConstOf(Sku_Family_STATUSC, Sku_Family_STATUSP))
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_Name_STATUSBasic, Sku_Name_STATUSPremium, Sku_Name_STATUSStandard))
}
