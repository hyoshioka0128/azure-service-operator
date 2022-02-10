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

func Test_DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARM, DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARM runs a test to see if a specific instance of DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARM(subject DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM
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

// Generator of DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM instances for property testing -
//lazily instantiated by DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARMGenerator()
var databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARMGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARMGenerator returns a generator of DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARMGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARMGenerator != nil {
		return databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARM(generators)
	databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARM(generators)
	databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM{}), generators)

	return databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = ThroughputSettingsUpdatePropertiesSpecARMGenerator()
}

func Test_ThroughputSettingsUpdateProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsUpdateProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsUpdatePropertiesSpecARM, ThroughputSettingsUpdatePropertiesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsUpdatePropertiesSpecARM runs a test to see if a specific instance of ThroughputSettingsUpdateProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsUpdatePropertiesSpecARM(subject ThroughputSettingsUpdateProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsUpdateProperties_SpecARM
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

// Generator of ThroughputSettingsUpdateProperties_SpecARM instances for property testing - lazily instantiated by
//ThroughputSettingsUpdatePropertiesSpecARMGenerator()
var throughputSettingsUpdatePropertiesSpecARMGenerator gopter.Gen

// ThroughputSettingsUpdatePropertiesSpecARMGenerator returns a generator of ThroughputSettingsUpdateProperties_SpecARM instances for property testing.
func ThroughputSettingsUpdatePropertiesSpecARMGenerator() gopter.Gen {
	if throughputSettingsUpdatePropertiesSpecARMGenerator != nil {
		return throughputSettingsUpdatePropertiesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForThroughputSettingsUpdatePropertiesSpecARM(generators)
	throughputSettingsUpdatePropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsUpdateProperties_SpecARM{}), generators)

	return throughputSettingsUpdatePropertiesSpecARMGenerator
}

// AddRelatedPropertyGeneratorsForThroughputSettingsUpdatePropertiesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsUpdatePropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["Resource"] = ThroughputSettingsResourceSpecARMGenerator()
}

func Test_ThroughputSettingsResource_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsResource_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsResourceSpecARM, ThroughputSettingsResourceSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsResourceSpecARM runs a test to see if a specific instance of ThroughputSettingsResource_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsResourceSpecARM(subject ThroughputSettingsResource_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsResource_SpecARM
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

// Generator of ThroughputSettingsResource_SpecARM instances for property testing - lazily instantiated by
//ThroughputSettingsResourceSpecARMGenerator()
var throughputSettingsResourceSpecARMGenerator gopter.Gen

// ThroughputSettingsResourceSpecARMGenerator returns a generator of ThroughputSettingsResource_SpecARM instances for property testing.
// We first initialize throughputSettingsResourceSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ThroughputSettingsResourceSpecARMGenerator() gopter.Gen {
	if throughputSettingsResourceSpecARMGenerator != nil {
		return throughputSettingsResourceSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsResourceSpecARM(generators)
	throughputSettingsResourceSpecARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsResource_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsResourceSpecARM(generators)
	AddRelatedPropertyGeneratorsForThroughputSettingsResourceSpecARM(generators)
	throughputSettingsResourceSpecARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsResource_SpecARM{}), generators)

	return throughputSettingsResourceSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForThroughputSettingsResourceSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputSettingsResourceSpecARM(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForThroughputSettingsResourceSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsResourceSpecARM(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsResourceSpecARMGenerator())
}

func Test_AutoscaleSettingsResource_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettingsResource_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettingsResourceSpecARM, AutoscaleSettingsResourceSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettingsResourceSpecARM runs a test to see if a specific instance of AutoscaleSettingsResource_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettingsResourceSpecARM(subject AutoscaleSettingsResource_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettingsResource_SpecARM
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

// Generator of AutoscaleSettingsResource_SpecARM instances for property testing - lazily instantiated by
//AutoscaleSettingsResourceSpecARMGenerator()
var autoscaleSettingsResourceSpecARMGenerator gopter.Gen

// AutoscaleSettingsResourceSpecARMGenerator returns a generator of AutoscaleSettingsResource_SpecARM instances for property testing.
// We first initialize autoscaleSettingsResourceSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AutoscaleSettingsResourceSpecARMGenerator() gopter.Gen {
	if autoscaleSettingsResourceSpecARMGenerator != nil {
		return autoscaleSettingsResourceSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceSpecARM(generators)
	autoscaleSettingsResourceSpecARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResource_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceSpecARM(generators)
	AddRelatedPropertyGeneratorsForAutoscaleSettingsResourceSpecARM(generators)
	autoscaleSettingsResourceSpecARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResource_SpecARM{}), generators)

	return autoscaleSettingsResourceSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceSpecARM(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.Int()
}

// AddRelatedPropertyGeneratorsForAutoscaleSettingsResourceSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoscaleSettingsResourceSpecARM(gens map[string]gopter.Gen) {
	gens["AutoUpgradePolicy"] = gen.PtrOf(AutoUpgradePolicyResourceSpecARMGenerator())
}

func Test_AutoUpgradePolicyResource_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoUpgradePolicyResource_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoUpgradePolicyResourceSpecARM, AutoUpgradePolicyResourceSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoUpgradePolicyResourceSpecARM runs a test to see if a specific instance of AutoUpgradePolicyResource_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoUpgradePolicyResourceSpecARM(subject AutoUpgradePolicyResource_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoUpgradePolicyResource_SpecARM
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

// Generator of AutoUpgradePolicyResource_SpecARM instances for property testing - lazily instantiated by
//AutoUpgradePolicyResourceSpecARMGenerator()
var autoUpgradePolicyResourceSpecARMGenerator gopter.Gen

// AutoUpgradePolicyResourceSpecARMGenerator returns a generator of AutoUpgradePolicyResource_SpecARM instances for property testing.
func AutoUpgradePolicyResourceSpecARMGenerator() gopter.Gen {
	if autoUpgradePolicyResourceSpecARMGenerator != nil {
		return autoUpgradePolicyResourceSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAutoUpgradePolicyResourceSpecARM(generators)
	autoUpgradePolicyResourceSpecARMGenerator = gen.Struct(reflect.TypeOf(AutoUpgradePolicyResource_SpecARM{}), generators)

	return autoUpgradePolicyResourceSpecARMGenerator
}

// AddRelatedPropertyGeneratorsForAutoUpgradePolicyResourceSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoUpgradePolicyResourceSpecARM(gens map[string]gopter.Gen) {
	gens["ThroughputPolicy"] = gen.PtrOf(ThroughputPolicyResourceSpecARMGenerator())
}

func Test_ThroughputPolicyResource_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputPolicyResource_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputPolicyResourceSpecARM, ThroughputPolicyResourceSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputPolicyResourceSpecARM runs a test to see if a specific instance of ThroughputPolicyResource_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputPolicyResourceSpecARM(subject ThroughputPolicyResource_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputPolicyResource_SpecARM
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

// Generator of ThroughputPolicyResource_SpecARM instances for property testing - lazily instantiated by
//ThroughputPolicyResourceSpecARMGenerator()
var throughputPolicyResourceSpecARMGenerator gopter.Gen

// ThroughputPolicyResourceSpecARMGenerator returns a generator of ThroughputPolicyResource_SpecARM instances for property testing.
func ThroughputPolicyResourceSpecARMGenerator() gopter.Gen {
	if throughputPolicyResourceSpecARMGenerator != nil {
		return throughputPolicyResourceSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputPolicyResourceSpecARM(generators)
	throughputPolicyResourceSpecARMGenerator = gen.Struct(reflect.TypeOf(ThroughputPolicyResource_SpecARM{}), generators)

	return throughputPolicyResourceSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForThroughputPolicyResourceSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputPolicyResourceSpecARM(gens map[string]gopter.Gen) {
	gens["IncrementPercent"] = gen.PtrOf(gen.Int())
	gens["IsEnabled"] = gen.PtrOf(gen.Bool())
}
