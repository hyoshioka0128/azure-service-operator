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

func Test_RedisCreateParameters_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisCreateParameters_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisCreateParametersStatusARM, RedisCreateParametersStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisCreateParametersStatusARM runs a test to see if a specific instance of RedisCreateParameters_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisCreateParametersStatusARM(subject RedisCreateParameters_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisCreateParameters_StatusARM
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

// Generator of RedisCreateParameters_StatusARM instances for property testing - lazily instantiated by
//RedisCreateParametersStatusARMGenerator()
var redisCreateParametersStatusARMGenerator gopter.Gen

// RedisCreateParametersStatusARMGenerator returns a generator of RedisCreateParameters_StatusARM instances for property testing.
// We first initialize redisCreateParametersStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisCreateParametersStatusARMGenerator() gopter.Gen {
	if redisCreateParametersStatusARMGenerator != nil {
		return redisCreateParametersStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreateParametersStatusARM(generators)
	redisCreateParametersStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisCreateParameters_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreateParametersStatusARM(generators)
	AddRelatedPropertyGeneratorsForRedisCreateParametersStatusARM(generators)
	redisCreateParametersStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisCreateParameters_StatusARM{}), generators)

	return redisCreateParametersStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisCreateParametersStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisCreateParametersStatusARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisCreateParametersStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisCreateParametersStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisCreatePropertiesStatusARMGenerator())
}

func Test_RedisCreateProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisCreateProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisCreatePropertiesStatusARM, RedisCreatePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisCreatePropertiesStatusARM runs a test to see if a specific instance of RedisCreateProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisCreatePropertiesStatusARM(subject RedisCreateProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisCreateProperties_StatusARM
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

// Generator of RedisCreateProperties_StatusARM instances for property testing - lazily instantiated by
//RedisCreatePropertiesStatusARMGenerator()
var redisCreatePropertiesStatusARMGenerator gopter.Gen

// RedisCreatePropertiesStatusARMGenerator returns a generator of RedisCreateProperties_StatusARM instances for property testing.
// We first initialize redisCreatePropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisCreatePropertiesStatusARMGenerator() gopter.Gen {
	if redisCreatePropertiesStatusARMGenerator != nil {
		return redisCreatePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreatePropertiesStatusARM(generators)
	redisCreatePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisCreateProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreatePropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForRedisCreatePropertiesStatusARM(generators)
	redisCreatePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisCreateProperties_StatusARM{}), generators)

	return redisCreatePropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisCreatePropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisCreatePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(RedisCreatePropertiesStatusMinimumTlsVersion10, RedisCreatePropertiesStatusMinimumTlsVersion11, RedisCreatePropertiesStatusMinimumTlsVersion12))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(RedisCreatePropertiesStatusPublicNetworkAccessDisabled, RedisCreatePropertiesStatusPublicNetworkAccessEnabled))
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicasPerMaster"] = gen.PtrOf(gen.Int())
	gens["ReplicasPerPrimary"] = gen.PtrOf(gen.Int())
	gens["ShardCount"] = gen.PtrOf(gen.Int())
	gens["StaticIP"] = gen.PtrOf(gen.AlphaString())
	gens["SubnetId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantSettings"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisCreatePropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisCreatePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["RedisConfiguration"] = gen.PtrOf(RedisCreatePropertiesStatusRedisConfigurationARMGenerator())
	gens["Sku"] = SkuStatusARMGenerator()
}

func Test_RedisCreateProperties_Status_RedisConfigurationARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisCreateProperties_Status_RedisConfigurationARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisCreatePropertiesStatusRedisConfigurationARM, RedisCreatePropertiesStatusRedisConfigurationARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisCreatePropertiesStatusRedisConfigurationARM runs a test to see if a specific instance of RedisCreateProperties_Status_RedisConfigurationARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisCreatePropertiesStatusRedisConfigurationARM(subject RedisCreateProperties_Status_RedisConfigurationARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisCreateProperties_Status_RedisConfigurationARM
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

// Generator of RedisCreateProperties_Status_RedisConfigurationARM instances for property testing - lazily instantiated
//by RedisCreatePropertiesStatusRedisConfigurationARMGenerator()
var redisCreatePropertiesStatusRedisConfigurationARMGenerator gopter.Gen

// RedisCreatePropertiesStatusRedisConfigurationARMGenerator returns a generator of RedisCreateProperties_Status_RedisConfigurationARM instances for property testing.
func RedisCreatePropertiesStatusRedisConfigurationARMGenerator() gopter.Gen {
	if redisCreatePropertiesStatusRedisConfigurationARMGenerator != nil {
		return redisCreatePropertiesStatusRedisConfigurationARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreatePropertiesStatusRedisConfigurationARM(generators)
	redisCreatePropertiesStatusRedisConfigurationARMGenerator = gen.Struct(reflect.TypeOf(RedisCreateProperties_Status_RedisConfigurationARM{}), generators)

	return redisCreatePropertiesStatusRedisConfigurationARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisCreatePropertiesStatusRedisConfigurationARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisCreatePropertiesStatusRedisConfigurationARM(gens map[string]gopter.Gen) {
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
	gens["additionalProperties"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_Sku_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuStatusARM, SkuStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuStatusARM runs a test to see if a specific instance of Sku_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuStatusARM(subject Sku_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_StatusARM
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

// Generator of Sku_StatusARM instances for property testing - lazily instantiated by SkuStatusARMGenerator()
var skuStatusARMGenerator gopter.Gen

// SkuStatusARMGenerator returns a generator of Sku_StatusARM instances for property testing.
func SkuStatusARMGenerator() gopter.Gen {
	if skuStatusARMGenerator != nil {
		return skuStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuStatusARM(generators)
	skuStatusARMGenerator = gen.Struct(reflect.TypeOf(Sku_StatusARM{}), generators)

	return skuStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuStatusARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.Int()
	gens["Family"] = gen.OneConstOf(SkuStatusFamilyC, SkuStatusFamilyP)
	gens["Name"] = gen.OneConstOf(SkuStatusNameBasic, SkuStatusNamePremium, SkuStatusNameStandard)
}