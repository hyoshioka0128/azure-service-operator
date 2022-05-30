// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

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

func Test_DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM, DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM runs a test to see if a specific instance of DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM(subject DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM
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

// Generator of DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM instances for property testing -
// lazily instantiated by DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARMGenerator()
var databaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARMGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARMGenerator returns a generator of DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARMGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARMGenerator != nil {
		return databaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM(generators)
	databaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM(generators)
	databaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM{}), generators)

	return databaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ThroughputSettingsUpdatePropertiesARMGenerator())
}

func Test_ThroughputSettingsUpdatePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsUpdatePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsUpdatePropertiesARM, ThroughputSettingsUpdatePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsUpdatePropertiesARM runs a test to see if a specific instance of ThroughputSettingsUpdatePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsUpdatePropertiesARM(subject ThroughputSettingsUpdatePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsUpdatePropertiesARM
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

// Generator of ThroughputSettingsUpdatePropertiesARM instances for property testing - lazily instantiated by
// ThroughputSettingsUpdatePropertiesARMGenerator()
var throughputSettingsUpdatePropertiesARMGenerator gopter.Gen

// ThroughputSettingsUpdatePropertiesARMGenerator returns a generator of ThroughputSettingsUpdatePropertiesARM instances for property testing.
func ThroughputSettingsUpdatePropertiesARMGenerator() gopter.Gen {
	if throughputSettingsUpdatePropertiesARMGenerator != nil {
		return throughputSettingsUpdatePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForThroughputSettingsUpdatePropertiesARM(generators)
	throughputSettingsUpdatePropertiesARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsUpdatePropertiesARM{}), generators)

	return throughputSettingsUpdatePropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForThroughputSettingsUpdatePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsUpdatePropertiesARM(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceARMGenerator())
}

func Test_ThroughputSettingsResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsResourceARM, ThroughputSettingsResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsResourceARM runs a test to see if a specific instance of ThroughputSettingsResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsResourceARM(subject ThroughputSettingsResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsResourceARM
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

// Generator of ThroughputSettingsResourceARM instances for property testing - lazily instantiated by
// ThroughputSettingsResourceARMGenerator()
var throughputSettingsResourceARMGenerator gopter.Gen

// ThroughputSettingsResourceARMGenerator returns a generator of ThroughputSettingsResourceARM instances for property testing.
// We first initialize throughputSettingsResourceARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ThroughputSettingsResourceARMGenerator() gopter.Gen {
	if throughputSettingsResourceARMGenerator != nil {
		return throughputSettingsResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsResourceARM(generators)
	throughputSettingsResourceARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsResourceARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsResourceARM(generators)
	AddRelatedPropertyGeneratorsForThroughputSettingsResourceARM(generators)
	throughputSettingsResourceARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsResourceARM{}), generators)

	return throughputSettingsResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForThroughputSettingsResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputSettingsResourceARM(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForThroughputSettingsResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsResourceARM(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsResourceARMGenerator())
}

func Test_AutoscaleSettingsResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettingsResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettingsResourceARM, AutoscaleSettingsResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettingsResourceARM runs a test to see if a specific instance of AutoscaleSettingsResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettingsResourceARM(subject AutoscaleSettingsResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettingsResourceARM
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

// Generator of AutoscaleSettingsResourceARM instances for property testing - lazily instantiated by
// AutoscaleSettingsResourceARMGenerator()
var autoscaleSettingsResourceARMGenerator gopter.Gen

// AutoscaleSettingsResourceARMGenerator returns a generator of AutoscaleSettingsResourceARM instances for property testing.
// We first initialize autoscaleSettingsResourceARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AutoscaleSettingsResourceARMGenerator() gopter.Gen {
	if autoscaleSettingsResourceARMGenerator != nil {
		return autoscaleSettingsResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceARM(generators)
	autoscaleSettingsResourceARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResourceARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceARM(generators)
	AddRelatedPropertyGeneratorsForAutoscaleSettingsResourceARM(generators)
	autoscaleSettingsResourceARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResourceARM{}), generators)

	return autoscaleSettingsResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceARM(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForAutoscaleSettingsResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoscaleSettingsResourceARM(gens map[string]gopter.Gen) {
	gens["AutoUpgradePolicy"] = gen.PtrOf(AutoUpgradePolicyResourceARMGenerator())
}

func Test_AutoUpgradePolicyResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoUpgradePolicyResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoUpgradePolicyResourceARM, AutoUpgradePolicyResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoUpgradePolicyResourceARM runs a test to see if a specific instance of AutoUpgradePolicyResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoUpgradePolicyResourceARM(subject AutoUpgradePolicyResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoUpgradePolicyResourceARM
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

// Generator of AutoUpgradePolicyResourceARM instances for property testing - lazily instantiated by
// AutoUpgradePolicyResourceARMGenerator()
var autoUpgradePolicyResourceARMGenerator gopter.Gen

// AutoUpgradePolicyResourceARMGenerator returns a generator of AutoUpgradePolicyResourceARM instances for property testing.
func AutoUpgradePolicyResourceARMGenerator() gopter.Gen {
	if autoUpgradePolicyResourceARMGenerator != nil {
		return autoUpgradePolicyResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAutoUpgradePolicyResourceARM(generators)
	autoUpgradePolicyResourceARMGenerator = gen.Struct(reflect.TypeOf(AutoUpgradePolicyResourceARM{}), generators)

	return autoUpgradePolicyResourceARMGenerator
}

// AddRelatedPropertyGeneratorsForAutoUpgradePolicyResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoUpgradePolicyResourceARM(gens map[string]gopter.Gen) {
	gens["ThroughputPolicy"] = gen.PtrOf(ThroughputPolicyResourceARMGenerator())
}

func Test_ThroughputPolicyResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputPolicyResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputPolicyResourceARM, ThroughputPolicyResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputPolicyResourceARM runs a test to see if a specific instance of ThroughputPolicyResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputPolicyResourceARM(subject ThroughputPolicyResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputPolicyResourceARM
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

// Generator of ThroughputPolicyResourceARM instances for property testing - lazily instantiated by
// ThroughputPolicyResourceARMGenerator()
var throughputPolicyResourceARMGenerator gopter.Gen

// ThroughputPolicyResourceARMGenerator returns a generator of ThroughputPolicyResourceARM instances for property testing.
func ThroughputPolicyResourceARMGenerator() gopter.Gen {
	if throughputPolicyResourceARMGenerator != nil {
		return throughputPolicyResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputPolicyResourceARM(generators)
	throughputPolicyResourceARMGenerator = gen.Struct(reflect.TypeOf(ThroughputPolicyResourceARM{}), generators)

	return throughputPolicyResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForThroughputPolicyResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputPolicyResourceARM(gens map[string]gopter.Gen) {
	gens["IncrementPercent"] = gen.PtrOf(gen.Int())
	gens["IsEnabled"] = gen.PtrOf(gen.Bool())
}
