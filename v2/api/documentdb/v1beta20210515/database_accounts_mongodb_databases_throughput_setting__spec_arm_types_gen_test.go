// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

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

func Test_DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM, DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM runs a test to see if a specific instance of DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM(subject DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM
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

// Generator of DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM instances for property testing - lazily
// instantiated by DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARMGenerator()
var databaseAccountsMongodbDatabasesThroughputSetting_SpecARMGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARMGenerator returns a generator of DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesThroughputSetting_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARMGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesThroughputSetting_SpecARMGenerator != nil {
		return databaseAccountsMongodbDatabasesThroughputSetting_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM(generators)
	databaseAccountsMongodbDatabasesThroughputSetting_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM(generators)
	databaseAccountsMongodbDatabasesThroughputSetting_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM{}), generators)

	return databaseAccountsMongodbDatabasesThroughputSetting_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesThroughputSetting_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ThroughputSettingsUpdatePropertiesARMGenerator())
}
