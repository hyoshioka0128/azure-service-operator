// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

import (
	"encoding/json"
	v20201201s "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20201201storage"
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

func Test_Redis_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Redis to hub returns original",
		prop.ForAll(RunResourceConversionTestForRedis, RedisGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRedis tests if a specific instance of Redis round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRedis(subject Redis) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20201201s.Redis
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Redis
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Redis_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Redis to Redis via AssignPropertiesToRedis & AssignPropertiesFromRedis returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedis, RedisGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedis tests if a specific instance of Redis can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedis(subject Redis) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.Redis
	err := copied.AssignPropertiesToRedis(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Redis
	err = actual.AssignPropertiesFromRedis(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Redis_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis, RedisGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis runs a test to see if a specific instance of Redis round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis(subject Redis) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis
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

// Generator of Redis instances for property testing - lazily instantiated by RedisGenerator()
var redisGenerator gopter.Gen

// RedisGenerator returns a generator of Redis instances for property testing.
func RedisGenerator() gopter.Gen {
	if redisGenerator != nil {
		return redisGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedis(generators)
	redisGenerator = gen.Struct(reflect.TypeOf(Redis{}), generators)

	return redisGenerator
}

// AddRelatedPropertyGeneratorsForRedis is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis(gens map[string]gopter.Gen) {
	gens["Spec"] = Redis_SpecGenerator()
	gens["Status"] = Redis_STATUSGenerator()
}

func Test_Redis_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Redis_STATUS to Redis_STATUS via AssignPropertiesToRedis_STATUS & AssignPropertiesFromRedis_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedis_STATUS, Redis_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedis_STATUS tests if a specific instance of Redis_STATUS can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedis_STATUS(subject Redis_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.Redis_STATUS
	err := copied.AssignPropertiesToRedis_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Redis_STATUS
	err = actual.AssignPropertiesFromRedis_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Redis_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_STATUS, Redis_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_STATUS runs a test to see if a specific instance of Redis_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_STATUS(subject Redis_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_STATUS
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

// Generator of Redis_STATUS instances for property testing - lazily instantiated by Redis_STATUSGenerator()
var redis_STATUSGenerator gopter.Gen

// Redis_STATUSGenerator returns a generator of Redis_STATUS instances for property testing.
// We first initialize redis_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Redis_STATUSGenerator() gopter.Gen {
	if redis_STATUSGenerator != nil {
		return redis_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_STATUS(generators)
	redis_STATUSGenerator = gen.Struct(reflect.TypeOf(Redis_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_STATUS(generators)
	AddRelatedPropertyGeneratorsForRedis_STATUS(generators)
	redis_STATUSGenerator = gen.Struct(reflect.TypeOf(Redis_STATUS{}), generators)

	return redis_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedis_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_STATUS(gens map[string]gopter.Gen) {
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(RedisCreateProperties_MinimumTlsVersion_STATUS10, RedisCreateProperties_MinimumTlsVersion_STATUS11, RedisCreateProperties_MinimumTlsVersion_STATUS12))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(RedisCreateProperties_PublicNetworkAccess_STATUSDisabled, RedisCreateProperties_PublicNetworkAccess_STATUSEnabled))
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicasPerMaster"] = gen.PtrOf(gen.Int())
	gens["ReplicasPerPrimary"] = gen.PtrOf(gen.Int())
	gens["ShardCount"] = gen.PtrOf(gen.Int())
	gens["StaticIP"] = gen.PtrOf(gen.AlphaString())
	gens["SubnetId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TenantSettings"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedis_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis_STATUS(gens map[string]gopter.Gen) {
	gens["RedisConfiguration"] = gen.PtrOf(RedisCreateProperties_RedisConfiguration_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSGenerator())
}

func Test_Redis_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Redis_Spec to Redis_Spec via AssignPropertiesToRedis_Spec & AssignPropertiesFromRedis_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedis_Spec, Redis_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedis_Spec tests if a specific instance of Redis_Spec can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedis_Spec(subject Redis_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.Redis_Spec
	err := copied.AssignPropertiesToRedis_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Redis_Spec
	err = actual.AssignPropertiesFromRedis_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Redis_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_Spec, Redis_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_Spec runs a test to see if a specific instance of Redis_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_Spec(subject Redis_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_Spec
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

// Generator of Redis_Spec instances for property testing - lazily instantiated by Redis_SpecGenerator()
var redis_SpecGenerator gopter.Gen

// Redis_SpecGenerator returns a generator of Redis_Spec instances for property testing.
// We first initialize redis_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Redis_SpecGenerator() gopter.Gen {
	if redis_SpecGenerator != nil {
		return redis_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_Spec(generators)
	redis_SpecGenerator = gen.Struct(reflect.TypeOf(Redis_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_Spec(generators)
	AddRelatedPropertyGeneratorsForRedis_Spec(generators)
	redis_SpecGenerator = gen.Struct(reflect.TypeOf(Redis_Spec{}), generators)

	return redis_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedis_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(RedisCreateProperties_MinimumTlsVersion10, RedisCreateProperties_MinimumTlsVersion11, RedisCreateProperties_MinimumTlsVersion12))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(RedisCreateProperties_PublicNetworkAccessDisabled, RedisCreateProperties_PublicNetworkAccessEnabled))
	gens["RedisConfiguration"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicasPerMaster"] = gen.PtrOf(gen.Int())
	gens["ReplicasPerPrimary"] = gen.PtrOf(gen.Int())
	gens["ShardCount"] = gen.PtrOf(gen.Int())
	gens["StaticIP"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TenantSettings"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedis_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(RedisOperatorSpecGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_RedisCreateProperties_RedisConfiguration_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisCreateProperties_RedisConfiguration_STATUS to RedisCreateProperties_RedisConfiguration_STATUS via AssignPropertiesToRedisCreateProperties_RedisConfiguration_STATUS & AssignPropertiesFromRedisCreateProperties_RedisConfiguration_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisCreateProperties_RedisConfiguration_STATUS, RedisCreateProperties_RedisConfiguration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisCreateProperties_RedisConfiguration_STATUS tests if a specific instance of RedisCreateProperties_RedisConfiguration_STATUS can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedisCreateProperties_RedisConfiguration_STATUS(subject RedisCreateProperties_RedisConfiguration_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.RedisCreateProperties_RedisConfiguration_STATUS
	err := copied.AssignPropertiesToRedisCreateProperties_RedisConfiguration_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisCreateProperties_RedisConfiguration_STATUS
	err = actual.AssignPropertiesFromRedisCreateProperties_RedisConfiguration_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RedisCreateProperties_RedisConfiguration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisCreateProperties_RedisConfiguration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisCreateProperties_RedisConfiguration_STATUS, RedisCreateProperties_RedisConfiguration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisCreateProperties_RedisConfiguration_STATUS runs a test to see if a specific instance of RedisCreateProperties_RedisConfiguration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisCreateProperties_RedisConfiguration_STATUS(subject RedisCreateProperties_RedisConfiguration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisCreateProperties_RedisConfiguration_STATUS
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

// Generator of RedisCreateProperties_RedisConfiguration_STATUS instances for property testing - lazily instantiated by
// RedisCreateProperties_RedisConfiguration_STATUSGenerator()
var redisCreateProperties_RedisConfiguration_STATUSGenerator gopter.Gen

// RedisCreateProperties_RedisConfiguration_STATUSGenerator returns a generator of RedisCreateProperties_RedisConfiguration_STATUS instances for property testing.
func RedisCreateProperties_RedisConfiguration_STATUSGenerator() gopter.Gen {
	if redisCreateProperties_RedisConfiguration_STATUSGenerator != nil {
		return redisCreateProperties_RedisConfiguration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreateProperties_RedisConfiguration_STATUS(generators)
	redisCreateProperties_RedisConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisCreateProperties_RedisConfiguration_STATUS{}), generators)

	return redisCreateProperties_RedisConfiguration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisCreateProperties_RedisConfiguration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisCreateProperties_RedisConfiguration_STATUS(gens map[string]gopter.Gen) {
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

func Test_RedisOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisOperatorSpec to RedisOperatorSpec via AssignPropertiesToRedisOperatorSpec & AssignPropertiesFromRedisOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisOperatorSpec, RedisOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisOperatorSpec tests if a specific instance of RedisOperatorSpec can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedisOperatorSpec(subject RedisOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.RedisOperatorSpec
	err := copied.AssignPropertiesToRedisOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisOperatorSpec
	err = actual.AssignPropertiesFromRedisOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RedisOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisOperatorSpec, RedisOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisOperatorSpec runs a test to see if a specific instance of RedisOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisOperatorSpec(subject RedisOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisOperatorSpec
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

// Generator of RedisOperatorSpec instances for property testing - lazily instantiated by RedisOperatorSpecGenerator()
var redisOperatorSpecGenerator gopter.Gen

// RedisOperatorSpecGenerator returns a generator of RedisOperatorSpec instances for property testing.
func RedisOperatorSpecGenerator() gopter.Gen {
	if redisOperatorSpecGenerator != nil {
		return redisOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisOperatorSpec(generators)
	redisOperatorSpecGenerator = gen.Struct(reflect.TypeOf(RedisOperatorSpec{}), generators)

	return redisOperatorSpecGenerator
}

// AddRelatedPropertyGeneratorsForRedisOperatorSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisOperatorSpec(gens map[string]gopter.Gen) {
	gens["Secrets"] = gen.PtrOf(RedisOperatorSecretsGenerator())
}

func Test_Sku_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Sku to Sku via AssignPropertiesToSku & AssignPropertiesFromSku returns original",
		prop.ForAll(RunPropertyAssignmentTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSku tests if a specific instance of Sku can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForSku(subject Sku) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.Sku
	err := copied.AssignPropertiesToSku(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Sku
	err = actual.AssignPropertiesFromSku(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Sku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku runs a test to see if a specific instance of Sku round trips to JSON and back losslessly
func RunJSONSerializationTestForSku(subject Sku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku
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

// Generator of Sku instances for property testing - lazily instantiated by SkuGenerator()
var skuGenerator gopter.Gen

// SkuGenerator returns a generator of Sku instances for property testing.
func SkuGenerator() gopter.Gen {
	if skuGenerator != nil {
		return skuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku(generators)
	skuGenerator = gen.Struct(reflect.TypeOf(Sku{}), generators)

	return skuGenerator
}

// AddIndependentPropertyGeneratorsForSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.OneConstOf(Sku_FamilyC, Sku_FamilyP))
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_NameBasic, Sku_NamePremium, Sku_NameStandard))
}

func Test_Sku_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Sku_STATUS to Sku_STATUS via AssignPropertiesToSku_STATUS & AssignPropertiesFromSku_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSku_STATUS, Sku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSku_STATUS tests if a specific instance of Sku_STATUS can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForSku_STATUS(subject Sku_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.Sku_STATUS
	err := copied.AssignPropertiesToSku_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Sku_STATUS
	err = actual.AssignPropertiesFromSku_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Sku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUS, Sku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUS runs a test to see if a specific instance of Sku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUS(subject Sku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS
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

// Generator of Sku_STATUS instances for property testing - lazily instantiated by Sku_STATUSGenerator()
var sku_STATUSGenerator gopter.Gen

// Sku_STATUSGenerator returns a generator of Sku_STATUS instances for property testing.
func Sku_STATUSGenerator() gopter.Gen {
	if sku_STATUSGenerator != nil {
		return sku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUS(generators)
	sku_STATUSGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS{}), generators)

	return sku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUS(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.OneConstOf(Sku_Family_STATUSC, Sku_Family_STATUSP))
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_Name_STATUSBasic, Sku_Name_STATUSPremium, Sku_Name_STATUSStandard))
}

func Test_RedisOperatorSecrets_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisOperatorSecrets to RedisOperatorSecrets via AssignPropertiesToRedisOperatorSecrets & AssignPropertiesFromRedisOperatorSecrets returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisOperatorSecrets, RedisOperatorSecretsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisOperatorSecrets tests if a specific instance of RedisOperatorSecrets can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedisOperatorSecrets(subject RedisOperatorSecrets) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.RedisOperatorSecrets
	err := copied.AssignPropertiesToRedisOperatorSecrets(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisOperatorSecrets
	err = actual.AssignPropertiesFromRedisOperatorSecrets(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RedisOperatorSecrets_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisOperatorSecrets via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisOperatorSecrets, RedisOperatorSecretsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisOperatorSecrets runs a test to see if a specific instance of RedisOperatorSecrets round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisOperatorSecrets(subject RedisOperatorSecrets) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisOperatorSecrets
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

// Generator of RedisOperatorSecrets instances for property testing - lazily instantiated by
// RedisOperatorSecretsGenerator()
var redisOperatorSecretsGenerator gopter.Gen

// RedisOperatorSecretsGenerator returns a generator of RedisOperatorSecrets instances for property testing.
func RedisOperatorSecretsGenerator() gopter.Gen {
	if redisOperatorSecretsGenerator != nil {
		return redisOperatorSecretsGenerator
	}

	generators := make(map[string]gopter.Gen)
	redisOperatorSecretsGenerator = gen.Struct(reflect.TypeOf(RedisOperatorSecrets{}), generators)

	return redisOperatorSecretsGenerator
}
