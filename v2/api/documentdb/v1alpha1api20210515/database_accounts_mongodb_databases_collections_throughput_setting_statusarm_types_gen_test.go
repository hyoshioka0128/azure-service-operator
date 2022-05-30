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

func Test_DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM, DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM runs a test to see if a specific instance of DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM(subject DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM
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

// Generator of DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM instances for property testing -
// lazily instantiated by DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARMGenerator()
var databaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARMGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARMGenerator returns a generator of DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARMGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARMGenerator != nil {
		return databaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM(generators)
	databaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM(generators)
	databaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM{}), generators)

	return databaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ThroughputSettingsGetProperties_STATUSARMGenerator())
}

func Test_ThroughputSettingsGetProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsGetProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsGetProperties_STATUSARM, ThroughputSettingsGetProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsGetProperties_STATUSARM runs a test to see if a specific instance of ThroughputSettingsGetProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsGetProperties_STATUSARM(subject ThroughputSettingsGetProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsGetProperties_STATUSARM
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

// Generator of ThroughputSettingsGetProperties_STATUSARM instances for property testing - lazily instantiated by
// ThroughputSettingsGetProperties_STATUSARMGenerator()
var throughputSettingsGetProperties_STATUSARMGenerator gopter.Gen

// ThroughputSettingsGetProperties_STATUSARMGenerator returns a generator of ThroughputSettingsGetProperties_STATUSARM instances for property testing.
func ThroughputSettingsGetProperties_STATUSARMGenerator() gopter.Gen {
	if throughputSettingsGetProperties_STATUSARMGenerator != nil {
		return throughputSettingsGetProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForThroughputSettingsGetProperties_STATUSARM(generators)
	throughputSettingsGetProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsGetProperties_STATUSARM{}), generators)

	return throughputSettingsGetProperties_STATUSARMGenerator
}

// AddRelatedPropertyGeneratorsForThroughputSettingsGetProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsGetProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsGetProperties_Resource_STATUSARMGenerator())
}

func Test_ThroughputSettingsGetProperties_Resource_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsGetProperties_Resource_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsGetProperties_Resource_STATUSARM, ThroughputSettingsGetProperties_Resource_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsGetProperties_Resource_STATUSARM runs a test to see if a specific instance of ThroughputSettingsGetProperties_Resource_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsGetProperties_Resource_STATUSARM(subject ThroughputSettingsGetProperties_Resource_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsGetProperties_Resource_STATUSARM
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

// Generator of ThroughputSettingsGetProperties_Resource_STATUSARM instances for property testing - lazily instantiated
// by ThroughputSettingsGetProperties_Resource_STATUSARMGenerator()
var throughputSettingsGetProperties_Resource_STATUSARMGenerator gopter.Gen

// ThroughputSettingsGetProperties_Resource_STATUSARMGenerator returns a generator of ThroughputSettingsGetProperties_Resource_STATUSARM instances for property testing.
// We first initialize throughputSettingsGetProperties_Resource_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ThroughputSettingsGetProperties_Resource_STATUSARMGenerator() gopter.Gen {
	if throughputSettingsGetProperties_Resource_STATUSARMGenerator != nil {
		return throughputSettingsGetProperties_Resource_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUSARM(generators)
	throughputSettingsGetProperties_Resource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsGetProperties_Resource_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUSARM(generators)
	throughputSettingsGetProperties_Resource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsGetProperties_Resource_STATUSARM{}), generators)

	return throughputSettingsGetProperties_Resource_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUSARM(gens map[string]gopter.Gen) {
	gens["MinimumThroughput"] = gen.PtrOf(gen.AlphaString())
	gens["OfferReplacePending"] = gen.PtrOf(gen.AlphaString())
	gens["Throughput"] = gen.PtrOf(gen.Int())
	gens["_Etag"] = gen.PtrOf(gen.AlphaString())
	gens["_Rid"] = gen.PtrOf(gen.AlphaString())
	gens["_Ts"] = gen.PtrOf(gen.Float64())
}

// AddRelatedPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsGetProperties_Resource_STATUSARM(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsResource_STATUSARMGenerator())
}

func Test_AutoscaleSettingsResource_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettingsResource_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettingsResource_STATUSARM, AutoscaleSettingsResource_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettingsResource_STATUSARM runs a test to see if a specific instance of AutoscaleSettingsResource_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettingsResource_STATUSARM(subject AutoscaleSettingsResource_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettingsResource_STATUSARM
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

// Generator of AutoscaleSettingsResource_STATUSARM instances for property testing - lazily instantiated by
// AutoscaleSettingsResource_STATUSARMGenerator()
var autoscaleSettingsResource_STATUSARMGenerator gopter.Gen

// AutoscaleSettingsResource_STATUSARMGenerator returns a generator of AutoscaleSettingsResource_STATUSARM instances for property testing.
// We first initialize autoscaleSettingsResource_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AutoscaleSettingsResource_STATUSARMGenerator() gopter.Gen {
	if autoscaleSettingsResource_STATUSARMGenerator != nil {
		return autoscaleSettingsResource_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResource_STATUSARM(generators)
	autoscaleSettingsResource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResource_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResource_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForAutoscaleSettingsResource_STATUSARM(generators)
	autoscaleSettingsResource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResource_STATUSARM{}), generators)

	return autoscaleSettingsResource_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettingsResource_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettingsResource_STATUSARM(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
	gens["TargetMaxThroughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForAutoscaleSettingsResource_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoscaleSettingsResource_STATUSARM(gens map[string]gopter.Gen) {
	gens["AutoUpgradePolicy"] = gen.PtrOf(AutoUpgradePolicyResource_STATUSARMGenerator())
}

func Test_AutoUpgradePolicyResource_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoUpgradePolicyResource_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoUpgradePolicyResource_STATUSARM, AutoUpgradePolicyResource_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoUpgradePolicyResource_STATUSARM runs a test to see if a specific instance of AutoUpgradePolicyResource_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoUpgradePolicyResource_STATUSARM(subject AutoUpgradePolicyResource_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoUpgradePolicyResource_STATUSARM
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

// Generator of AutoUpgradePolicyResource_STATUSARM instances for property testing - lazily instantiated by
// AutoUpgradePolicyResource_STATUSARMGenerator()
var autoUpgradePolicyResource_STATUSARMGenerator gopter.Gen

// AutoUpgradePolicyResource_STATUSARMGenerator returns a generator of AutoUpgradePolicyResource_STATUSARM instances for property testing.
func AutoUpgradePolicyResource_STATUSARMGenerator() gopter.Gen {
	if autoUpgradePolicyResource_STATUSARMGenerator != nil {
		return autoUpgradePolicyResource_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAutoUpgradePolicyResource_STATUSARM(generators)
	autoUpgradePolicyResource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(AutoUpgradePolicyResource_STATUSARM{}), generators)

	return autoUpgradePolicyResource_STATUSARMGenerator
}

// AddRelatedPropertyGeneratorsForAutoUpgradePolicyResource_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoUpgradePolicyResource_STATUSARM(gens map[string]gopter.Gen) {
	gens["ThroughputPolicy"] = gen.PtrOf(ThroughputPolicyResource_STATUSARMGenerator())
}

func Test_ThroughputPolicyResource_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputPolicyResource_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputPolicyResource_STATUSARM, ThroughputPolicyResource_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputPolicyResource_STATUSARM runs a test to see if a specific instance of ThroughputPolicyResource_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputPolicyResource_STATUSARM(subject ThroughputPolicyResource_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputPolicyResource_STATUSARM
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

// Generator of ThroughputPolicyResource_STATUSARM instances for property testing - lazily instantiated by
// ThroughputPolicyResource_STATUSARMGenerator()
var throughputPolicyResource_STATUSARMGenerator gopter.Gen

// ThroughputPolicyResource_STATUSARMGenerator returns a generator of ThroughputPolicyResource_STATUSARM instances for property testing.
func ThroughputPolicyResource_STATUSARMGenerator() gopter.Gen {
	if throughputPolicyResource_STATUSARMGenerator != nil {
		return throughputPolicyResource_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputPolicyResource_STATUSARM(generators)
	throughputPolicyResource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ThroughputPolicyResource_STATUSARM{}), generators)

	return throughputPolicyResource_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForThroughputPolicyResource_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputPolicyResource_STATUSARM(gens map[string]gopter.Gen) {
	gens["IncrementPercent"] = gen.PtrOf(gen.Int())
	gens["IsEnabled"] = gen.PtrOf(gen.Bool())
}
