// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

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

func Test_MongodbDatabaseCollection_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabaseCollection via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabaseCollection, MongodbDatabaseCollectionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabaseCollection runs a test to see if a specific instance of MongodbDatabaseCollection round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabaseCollection(subject MongodbDatabaseCollection) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabaseCollection
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

// Generator of MongodbDatabaseCollection instances for property testing - lazily instantiated by
//MongodbDatabaseCollectionGenerator()
var mongodbDatabaseCollectionGenerator gopter.Gen

// MongodbDatabaseCollectionGenerator returns a generator of MongodbDatabaseCollection instances for property testing.
func MongodbDatabaseCollectionGenerator() gopter.Gen {
	if mongodbDatabaseCollectionGenerator != nil {
		return mongodbDatabaseCollectionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongodbDatabaseCollection(generators)
	mongodbDatabaseCollectionGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseCollection{}), generators)

	return mongodbDatabaseCollectionGenerator
}

// AddRelatedPropertyGeneratorsForMongodbDatabaseCollection is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabaseCollection(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsMongodbDatabasesCollectionsSPECGenerator()
	gens["Status"] = MongoDBCollectionCreateUpdateParametersStatusGenerator()
}

func Test_DatabaseAccountsMongodbDatabasesCollections_SPEC_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabasesCollections_SPEC via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsSPEC, DatabaseAccountsMongodbDatabasesCollectionsSPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsSPEC runs a test to see if a specific instance of DatabaseAccountsMongodbDatabasesCollections_SPEC round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsSPEC(subject DatabaseAccountsMongodbDatabasesCollections_SPEC) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabasesCollections_SPEC
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

// Generator of DatabaseAccountsMongodbDatabasesCollections_SPEC instances for property testing - lazily instantiated by
//DatabaseAccountsMongodbDatabasesCollectionsSPECGenerator()
var databaseAccountsMongodbDatabasesCollectionsSPECGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesCollectionsSPECGenerator returns a generator of DatabaseAccountsMongodbDatabasesCollections_SPEC instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesCollectionsSPECGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesCollectionsSPECGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesCollectionsSPECGenerator != nil {
		return databaseAccountsMongodbDatabasesCollectionsSPECGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSPEC(generators)
	databaseAccountsMongodbDatabasesCollectionsSPECGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollections_SPEC{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSPEC(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSPEC(generators)
	databaseAccountsMongodbDatabasesCollectionsSPECGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollections_SPEC{}), generators)

	return databaseAccountsMongodbDatabasesCollectionsSPECGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSPEC is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSPEC(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSPEC is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSPEC(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsSpecGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBCollectionResourceSpecGenerator())
}

func Test_MongoDBCollectionCreateUpdateParameters_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionCreateUpdateParameters_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionCreateUpdateParametersStatus, MongoDBCollectionCreateUpdateParametersStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionCreateUpdateParametersStatus runs a test to see if a specific instance of MongoDBCollectionCreateUpdateParameters_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionCreateUpdateParametersStatus(subject MongoDBCollectionCreateUpdateParameters_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionCreateUpdateParameters_Status
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

// Generator of MongoDBCollectionCreateUpdateParameters_Status instances for property testing - lazily instantiated by
//MongoDBCollectionCreateUpdateParametersStatusGenerator()
var mongoDBCollectionCreateUpdateParametersStatusGenerator gopter.Gen

// MongoDBCollectionCreateUpdateParametersStatusGenerator returns a generator of MongoDBCollectionCreateUpdateParameters_Status instances for property testing.
// We first initialize mongoDBCollectionCreateUpdateParametersStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionCreateUpdateParametersStatusGenerator() gopter.Gen {
	if mongoDBCollectionCreateUpdateParametersStatusGenerator != nil {
		return mongoDBCollectionCreateUpdateParametersStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionCreateUpdateParametersStatus(generators)
	mongoDBCollectionCreateUpdateParametersStatusGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionCreateUpdateParameters_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionCreateUpdateParametersStatus(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionCreateUpdateParametersStatus(generators)
	mongoDBCollectionCreateUpdateParametersStatusGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionCreateUpdateParameters_Status{}), generators)

	return mongoDBCollectionCreateUpdateParametersStatusGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionCreateUpdateParametersStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionCreateUpdateParametersStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionCreateUpdateParametersStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionCreateUpdateParametersStatus(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsStatusGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBCollectionResourceStatusGenerator())
}

func Test_MongoDBCollectionResource_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionResource_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionResourceSpec, MongoDBCollectionResourceSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionResourceSpec runs a test to see if a specific instance of MongoDBCollectionResource_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionResourceSpec(subject MongoDBCollectionResource_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionResource_Spec
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

// Generator of MongoDBCollectionResource_Spec instances for property testing - lazily instantiated by
//MongoDBCollectionResourceSpecGenerator()
var mongoDBCollectionResourceSpecGenerator gopter.Gen

// MongoDBCollectionResourceSpecGenerator returns a generator of MongoDBCollectionResource_Spec instances for property testing.
// We first initialize mongoDBCollectionResourceSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionResourceSpecGenerator() gopter.Gen {
	if mongoDBCollectionResourceSpecGenerator != nil {
		return mongoDBCollectionResourceSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResourceSpec(generators)
	mongoDBCollectionResourceSpecGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResourceSpec(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionResourceSpec(generators)
	mongoDBCollectionResourceSpecGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource_Spec{}), generators)

	return mongoDBCollectionResourceSpecGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionResourceSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionResourceSpec(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["ShardKey"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionResourceSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionResourceSpec(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(MongoIndexSpecGenerator())
}

func Test_MongoDBCollectionResource_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionResource_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionResourceStatus, MongoDBCollectionResourceStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionResourceStatus runs a test to see if a specific instance of MongoDBCollectionResource_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionResourceStatus(subject MongoDBCollectionResource_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionResource_Status
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

// Generator of MongoDBCollectionResource_Status instances for property testing - lazily instantiated by
//MongoDBCollectionResourceStatusGenerator()
var mongoDBCollectionResourceStatusGenerator gopter.Gen

// MongoDBCollectionResourceStatusGenerator returns a generator of MongoDBCollectionResource_Status instances for property testing.
// We first initialize mongoDBCollectionResourceStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionResourceStatusGenerator() gopter.Gen {
	if mongoDBCollectionResourceStatusGenerator != nil {
		return mongoDBCollectionResourceStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResourceStatus(generators)
	mongoDBCollectionResourceStatusGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResourceStatus(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionResourceStatus(generators)
	mongoDBCollectionResourceStatusGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource_Status{}), generators)

	return mongoDBCollectionResourceStatusGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionResourceStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionResourceStatus(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["ShardKey"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionResourceStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionResourceStatus(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(MongoIndexStatusGenerator())
}

func Test_MongoIndex_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndex_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexSpec, MongoIndexSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexSpec runs a test to see if a specific instance of MongoIndex_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexSpec(subject MongoIndex_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndex_Spec
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

// Generator of MongoIndex_Spec instances for property testing - lazily instantiated by MongoIndexSpecGenerator()
var mongoIndexSpecGenerator gopter.Gen

// MongoIndexSpecGenerator returns a generator of MongoIndex_Spec instances for property testing.
func MongoIndexSpecGenerator() gopter.Gen {
	if mongoIndexSpecGenerator != nil {
		return mongoIndexSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoIndexSpec(generators)
	mongoIndexSpecGenerator = gen.Struct(reflect.TypeOf(MongoIndex_Spec{}), generators)

	return mongoIndexSpecGenerator
}

// AddRelatedPropertyGeneratorsForMongoIndexSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoIndexSpec(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(MongoIndexKeysSpecGenerator())
	gens["Options"] = gen.PtrOf(MongoIndexOptionsSpecGenerator())
}

func Test_MongoIndex_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndex_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexStatus, MongoIndexStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexStatus runs a test to see if a specific instance of MongoIndex_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexStatus(subject MongoIndex_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndex_Status
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

// Generator of MongoIndex_Status instances for property testing - lazily instantiated by MongoIndexStatusGenerator()
var mongoIndexStatusGenerator gopter.Gen

// MongoIndexStatusGenerator returns a generator of MongoIndex_Status instances for property testing.
func MongoIndexStatusGenerator() gopter.Gen {
	if mongoIndexStatusGenerator != nil {
		return mongoIndexStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoIndexStatus(generators)
	mongoIndexStatusGenerator = gen.Struct(reflect.TypeOf(MongoIndex_Status{}), generators)

	return mongoIndexStatusGenerator
}

// AddRelatedPropertyGeneratorsForMongoIndexStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoIndexStatus(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(MongoIndexKeysStatusGenerator())
	gens["Options"] = gen.PtrOf(MongoIndexOptionsStatusGenerator())
}

func Test_MongoIndexKeys_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexKeys_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexKeysSpec, MongoIndexKeysSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexKeysSpec runs a test to see if a specific instance of MongoIndexKeys_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexKeysSpec(subject MongoIndexKeys_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexKeys_Spec
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

// Generator of MongoIndexKeys_Spec instances for property testing - lazily instantiated by MongoIndexKeysSpecGenerator()
var mongoIndexKeysSpecGenerator gopter.Gen

// MongoIndexKeysSpecGenerator returns a generator of MongoIndexKeys_Spec instances for property testing.
func MongoIndexKeysSpecGenerator() gopter.Gen {
	if mongoIndexKeysSpecGenerator != nil {
		return mongoIndexKeysSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexKeysSpec(generators)
	mongoIndexKeysSpecGenerator = gen.Struct(reflect.TypeOf(MongoIndexKeys_Spec{}), generators)

	return mongoIndexKeysSpecGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexKeysSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexKeysSpec(gens map[string]gopter.Gen) {
	gens["Keys"] = gen.SliceOf(gen.AlphaString())
}

func Test_MongoIndexKeys_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexKeys_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexKeysStatus, MongoIndexKeysStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexKeysStatus runs a test to see if a specific instance of MongoIndexKeys_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexKeysStatus(subject MongoIndexKeys_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexKeys_Status
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

// Generator of MongoIndexKeys_Status instances for property testing - lazily instantiated by
//MongoIndexKeysStatusGenerator()
var mongoIndexKeysStatusGenerator gopter.Gen

// MongoIndexKeysStatusGenerator returns a generator of MongoIndexKeys_Status instances for property testing.
func MongoIndexKeysStatusGenerator() gopter.Gen {
	if mongoIndexKeysStatusGenerator != nil {
		return mongoIndexKeysStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexKeysStatus(generators)
	mongoIndexKeysStatusGenerator = gen.Struct(reflect.TypeOf(MongoIndexKeys_Status{}), generators)

	return mongoIndexKeysStatusGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexKeysStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexKeysStatus(gens map[string]gopter.Gen) {
	gens["Keys"] = gen.SliceOf(gen.AlphaString())
}

func Test_MongoIndexOptions_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexOptions_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexOptionsSpec, MongoIndexOptionsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexOptionsSpec runs a test to see if a specific instance of MongoIndexOptions_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexOptionsSpec(subject MongoIndexOptions_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexOptions_Spec
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

// Generator of MongoIndexOptions_Spec instances for property testing - lazily instantiated by
//MongoIndexOptionsSpecGenerator()
var mongoIndexOptionsSpecGenerator gopter.Gen

// MongoIndexOptionsSpecGenerator returns a generator of MongoIndexOptions_Spec instances for property testing.
func MongoIndexOptionsSpecGenerator() gopter.Gen {
	if mongoIndexOptionsSpecGenerator != nil {
		return mongoIndexOptionsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexOptionsSpec(generators)
	mongoIndexOptionsSpecGenerator = gen.Struct(reflect.TypeOf(MongoIndexOptions_Spec{}), generators)

	return mongoIndexOptionsSpecGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexOptionsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexOptionsSpec(gens map[string]gopter.Gen) {
	gens["ExpireAfterSeconds"] = gen.PtrOf(gen.Int())
	gens["Unique"] = gen.PtrOf(gen.Bool())
}

func Test_MongoIndexOptions_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexOptions_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexOptionsStatus, MongoIndexOptionsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexOptionsStatus runs a test to see if a specific instance of MongoIndexOptions_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexOptionsStatus(subject MongoIndexOptions_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexOptions_Status
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

// Generator of MongoIndexOptions_Status instances for property testing - lazily instantiated by
//MongoIndexOptionsStatusGenerator()
var mongoIndexOptionsStatusGenerator gopter.Gen

// MongoIndexOptionsStatusGenerator returns a generator of MongoIndexOptions_Status instances for property testing.
func MongoIndexOptionsStatusGenerator() gopter.Gen {
	if mongoIndexOptionsStatusGenerator != nil {
		return mongoIndexOptionsStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexOptionsStatus(generators)
	mongoIndexOptionsStatusGenerator = gen.Struct(reflect.TypeOf(MongoIndexOptions_Status{}), generators)

	return mongoIndexOptionsStatusGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexOptionsStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexOptionsStatus(gens map[string]gopter.Gen) {
	gens["ExpireAfterSeconds"] = gen.PtrOf(gen.Int())
	gens["Unique"] = gen.PtrOf(gen.Bool())
}
