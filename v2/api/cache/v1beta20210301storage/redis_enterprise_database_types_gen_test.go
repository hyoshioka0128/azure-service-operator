// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210301storage

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

func Test_RedisEnterpriseDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterpriseDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterpriseDatabase, RedisEnterpriseDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterpriseDatabase runs a test to see if a specific instance of RedisEnterpriseDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterpriseDatabase(subject RedisEnterpriseDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterpriseDatabase
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

// Generator of RedisEnterpriseDatabase instances for property testing - lazily instantiated by
// RedisEnterpriseDatabaseGenerator()
var redisEnterpriseDatabaseGenerator gopter.Gen

// RedisEnterpriseDatabaseGenerator returns a generator of RedisEnterpriseDatabase instances for property testing.
func RedisEnterpriseDatabaseGenerator() gopter.Gen {
	if redisEnterpriseDatabaseGenerator != nil {
		return redisEnterpriseDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase(generators)
	redisEnterpriseDatabaseGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabase{}), generators)

	return redisEnterpriseDatabaseGenerator
}

// AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisEnterpriseDatabase_SpecGenerator()
	gens["Status"] = RedisEnterpriseDatabase_STATUSGenerator()
}

func Test_RedisEnterpriseDatabase_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterpriseDatabase_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterpriseDatabase_STATUS, RedisEnterpriseDatabase_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterpriseDatabase_STATUS runs a test to see if a specific instance of RedisEnterpriseDatabase_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterpriseDatabase_STATUS(subject RedisEnterpriseDatabase_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterpriseDatabase_STATUS
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

// Generator of RedisEnterpriseDatabase_STATUS instances for property testing - lazily instantiated by
// RedisEnterpriseDatabase_STATUSGenerator()
var redisEnterpriseDatabase_STATUSGenerator gopter.Gen

// RedisEnterpriseDatabase_STATUSGenerator returns a generator of RedisEnterpriseDatabase_STATUS instances for property testing.
// We first initialize redisEnterpriseDatabase_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterpriseDatabase_STATUSGenerator() gopter.Gen {
	if redisEnterpriseDatabase_STATUSGenerator != nil {
		return redisEnterpriseDatabase_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_STATUS(generators)
	redisEnterpriseDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabase_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_STATUS(generators)
	AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase_STATUS(generators)
	redisEnterpriseDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabase_STATUS{}), generators)

	return redisEnterpriseDatabase_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["ClientProtocol"] = gen.PtrOf(gen.AlphaString())
	gens["ClusteringPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["EvictionPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceState"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Modules"] = gen.SliceOf(Module_STATUSGenerator())
	gens["Persistence"] = gen.PtrOf(Persistence_STATUSGenerator())
}

func Test_RedisEnterpriseDatabase_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterpriseDatabase_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterpriseDatabase_Spec, RedisEnterpriseDatabase_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterpriseDatabase_Spec runs a test to see if a specific instance of RedisEnterpriseDatabase_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterpriseDatabase_Spec(subject RedisEnterpriseDatabase_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterpriseDatabase_Spec
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

// Generator of RedisEnterpriseDatabase_Spec instances for property testing - lazily instantiated by
// RedisEnterpriseDatabase_SpecGenerator()
var redisEnterpriseDatabase_SpecGenerator gopter.Gen

// RedisEnterpriseDatabase_SpecGenerator returns a generator of RedisEnterpriseDatabase_Spec instances for property testing.
// We first initialize redisEnterpriseDatabase_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterpriseDatabase_SpecGenerator() gopter.Gen {
	if redisEnterpriseDatabase_SpecGenerator != nil {
		return redisEnterpriseDatabase_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_Spec(generators)
	redisEnterpriseDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabase_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_Spec(generators)
	AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase_Spec(generators)
	redisEnterpriseDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabase_Spec{}), generators)

	return redisEnterpriseDatabase_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["ClientProtocol"] = gen.PtrOf(gen.AlphaString())
	gens["ClusteringPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["EvictionPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Port"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase_Spec(gens map[string]gopter.Gen) {
	gens["Modules"] = gen.SliceOf(ModuleGenerator())
	gens["Persistence"] = gen.PtrOf(PersistenceGenerator())
}

func Test_Module_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Module via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForModule, ModuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForModule runs a test to see if a specific instance of Module round trips to JSON and back losslessly
func RunJSONSerializationTestForModule(subject Module) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Module
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

// Generator of Module instances for property testing - lazily instantiated by ModuleGenerator()
var moduleGenerator gopter.Gen

// ModuleGenerator returns a generator of Module instances for property testing.
func ModuleGenerator() gopter.Gen {
	if moduleGenerator != nil {
		return moduleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForModule(generators)
	moduleGenerator = gen.Struct(reflect.TypeOf(Module{}), generators)

	return moduleGenerator
}

// AddIndependentPropertyGeneratorsForModule is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForModule(gens map[string]gopter.Gen) {
	gens["Args"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_Module_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Module_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForModule_STATUS, Module_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForModule_STATUS runs a test to see if a specific instance of Module_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForModule_STATUS(subject Module_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Module_STATUS
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

// Generator of Module_STATUS instances for property testing - lazily instantiated by Module_STATUSGenerator()
var module_STATUSGenerator gopter.Gen

// Module_STATUSGenerator returns a generator of Module_STATUS instances for property testing.
func Module_STATUSGenerator() gopter.Gen {
	if module_STATUSGenerator != nil {
		return module_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForModule_STATUS(generators)
	module_STATUSGenerator = gen.Struct(reflect.TypeOf(Module_STATUS{}), generators)

	return module_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForModule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForModule_STATUS(gens map[string]gopter.Gen) {
	gens["Args"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

func Test_Persistence_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Persistence via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPersistence, PersistenceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPersistence runs a test to see if a specific instance of Persistence round trips to JSON and back losslessly
func RunJSONSerializationTestForPersistence(subject Persistence) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Persistence
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

// Generator of Persistence instances for property testing - lazily instantiated by PersistenceGenerator()
var persistenceGenerator gopter.Gen

// PersistenceGenerator returns a generator of Persistence instances for property testing.
func PersistenceGenerator() gopter.Gen {
	if persistenceGenerator != nil {
		return persistenceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPersistence(generators)
	persistenceGenerator = gen.Struct(reflect.TypeOf(Persistence{}), generators)

	return persistenceGenerator
}

// AddIndependentPropertyGeneratorsForPersistence is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPersistence(gens map[string]gopter.Gen) {
	gens["AofEnabled"] = gen.PtrOf(gen.Bool())
	gens["AofFrequency"] = gen.PtrOf(gen.AlphaString())
	gens["RdbEnabled"] = gen.PtrOf(gen.Bool())
	gens["RdbFrequency"] = gen.PtrOf(gen.AlphaString())
}

func Test_Persistence_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Persistence_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPersistence_STATUS, Persistence_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPersistence_STATUS runs a test to see if a specific instance of Persistence_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPersistence_STATUS(subject Persistence_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Persistence_STATUS
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

// Generator of Persistence_STATUS instances for property testing - lazily instantiated by Persistence_STATUSGenerator()
var persistence_STATUSGenerator gopter.Gen

// Persistence_STATUSGenerator returns a generator of Persistence_STATUS instances for property testing.
func Persistence_STATUSGenerator() gopter.Gen {
	if persistence_STATUSGenerator != nil {
		return persistence_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPersistence_STATUS(generators)
	persistence_STATUSGenerator = gen.Struct(reflect.TypeOf(Persistence_STATUS{}), generators)

	return persistence_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPersistence_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPersistence_STATUS(gens map[string]gopter.Gen) {
	gens["AofEnabled"] = gen.PtrOf(gen.Bool())
	gens["AofFrequency"] = gen.PtrOf(gen.AlphaString())
	gens["RdbEnabled"] = gen.PtrOf(gen.Bool())
	gens["RdbFrequency"] = gen.PtrOf(gen.AlphaString())
}
