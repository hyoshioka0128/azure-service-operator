// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101

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

func Test_BatchAccountCreateParameters_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccountCreateParameters_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccountCreateParametersStatusARM, BatchAccountCreateParametersStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccountCreateParametersStatusARM runs a test to see if a specific instance of BatchAccountCreateParameters_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccountCreateParametersStatusARM(subject BatchAccountCreateParameters_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccountCreateParameters_StatusARM
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

// Generator of BatchAccountCreateParameters_StatusARM instances for property testing - lazily instantiated by
//BatchAccountCreateParametersStatusARMGenerator()
var batchAccountCreateParametersStatusARMGenerator gopter.Gen

// BatchAccountCreateParametersStatusARMGenerator returns a generator of BatchAccountCreateParameters_StatusARM instances for property testing.
// We first initialize batchAccountCreateParametersStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BatchAccountCreateParametersStatusARMGenerator() gopter.Gen {
	if batchAccountCreateParametersStatusARMGenerator != nil {
		return batchAccountCreateParametersStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountCreateParametersStatusARM(generators)
	batchAccountCreateParametersStatusARMGenerator = gen.Struct(reflect.TypeOf(BatchAccountCreateParameters_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountCreateParametersStatusARM(generators)
	AddRelatedPropertyGeneratorsForBatchAccountCreateParametersStatusARM(generators)
	batchAccountCreateParametersStatusARMGenerator = gen.Struct(reflect.TypeOf(BatchAccountCreateParameters_StatusARM{}), generators)

	return batchAccountCreateParametersStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForBatchAccountCreateParametersStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBatchAccountCreateParametersStatusARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBatchAccountCreateParametersStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBatchAccountCreateParametersStatusARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(BatchAccountIdentityStatusARMGenerator())
	gens["Properties"] = gen.PtrOf(BatchAccountCreatePropertiesStatusARMGenerator())
}

func Test_BatchAccountCreateProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccountCreateProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccountCreatePropertiesStatusARM, BatchAccountCreatePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccountCreatePropertiesStatusARM runs a test to see if a specific instance of BatchAccountCreateProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccountCreatePropertiesStatusARM(subject BatchAccountCreateProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccountCreateProperties_StatusARM
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

// Generator of BatchAccountCreateProperties_StatusARM instances for property testing - lazily instantiated by
//BatchAccountCreatePropertiesStatusARMGenerator()
var batchAccountCreatePropertiesStatusARMGenerator gopter.Gen

// BatchAccountCreatePropertiesStatusARMGenerator returns a generator of BatchAccountCreateProperties_StatusARM instances for property testing.
// We first initialize batchAccountCreatePropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BatchAccountCreatePropertiesStatusARMGenerator() gopter.Gen {
	if batchAccountCreatePropertiesStatusARMGenerator != nil {
		return batchAccountCreatePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountCreatePropertiesStatusARM(generators)
	batchAccountCreatePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(BatchAccountCreateProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountCreatePropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForBatchAccountCreatePropertiesStatusARM(generators)
	batchAccountCreatePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(BatchAccountCreateProperties_StatusARM{}), generators)

	return batchAccountCreatePropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForBatchAccountCreatePropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBatchAccountCreatePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["PoolAllocationMode"] = gen.PtrOf(gen.OneConstOf(PoolAllocationMode_StatusBatchService, PoolAllocationMode_StatusUserSubscription))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_StatusDisabled, PublicNetworkAccessType_StatusEnabled))
}

// AddRelatedPropertyGeneratorsForBatchAccountCreatePropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBatchAccountCreatePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["AutoStorage"] = gen.PtrOf(AutoStorageBasePropertiesStatusARMGenerator())
	gens["Encryption"] = gen.PtrOf(EncryptionPropertiesStatusARMGenerator())
	gens["KeyVaultReference"] = gen.PtrOf(KeyVaultReferenceStatusARMGenerator())
}

func Test_BatchAccountIdentity_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccountIdentity_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccountIdentityStatusARM, BatchAccountIdentityStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccountIdentityStatusARM runs a test to see if a specific instance of BatchAccountIdentity_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccountIdentityStatusARM(subject BatchAccountIdentity_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccountIdentity_StatusARM
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

// Generator of BatchAccountIdentity_StatusARM instances for property testing - lazily instantiated by
//BatchAccountIdentityStatusARMGenerator()
var batchAccountIdentityStatusARMGenerator gopter.Gen

// BatchAccountIdentityStatusARMGenerator returns a generator of BatchAccountIdentity_StatusARM instances for property testing.
// We first initialize batchAccountIdentityStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BatchAccountIdentityStatusARMGenerator() gopter.Gen {
	if batchAccountIdentityStatusARMGenerator != nil {
		return batchAccountIdentityStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountIdentityStatusARM(generators)
	batchAccountIdentityStatusARMGenerator = gen.Struct(reflect.TypeOf(BatchAccountIdentity_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountIdentityStatusARM(generators)
	AddRelatedPropertyGeneratorsForBatchAccountIdentityStatusARM(generators)
	batchAccountIdentityStatusARMGenerator = gen.Struct(reflect.TypeOf(BatchAccountIdentity_StatusARM{}), generators)

	return batchAccountIdentityStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForBatchAccountIdentityStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBatchAccountIdentityStatusARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(BatchAccountIdentityStatusTypeNone, BatchAccountIdentityStatusTypeSystemAssigned, BatchAccountIdentityStatusTypeUserAssigned)
}

// AddRelatedPropertyGeneratorsForBatchAccountIdentityStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBatchAccountIdentityStatusARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(gen.AlphaString(), BatchAccountIdentityStatusUserAssignedIdentitiesARMGenerator())
}

func Test_AutoStorageBaseProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoStorageBaseProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoStorageBasePropertiesStatusARM, AutoStorageBasePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoStorageBasePropertiesStatusARM runs a test to see if a specific instance of AutoStorageBaseProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoStorageBasePropertiesStatusARM(subject AutoStorageBaseProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoStorageBaseProperties_StatusARM
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

// Generator of AutoStorageBaseProperties_StatusARM instances for property testing - lazily instantiated by
//AutoStorageBasePropertiesStatusARMGenerator()
var autoStorageBasePropertiesStatusARMGenerator gopter.Gen

// AutoStorageBasePropertiesStatusARMGenerator returns a generator of AutoStorageBaseProperties_StatusARM instances for property testing.
func AutoStorageBasePropertiesStatusARMGenerator() gopter.Gen {
	if autoStorageBasePropertiesStatusARMGenerator != nil {
		return autoStorageBasePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoStorageBasePropertiesStatusARM(generators)
	autoStorageBasePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(AutoStorageBaseProperties_StatusARM{}), generators)

	return autoStorageBasePropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoStorageBasePropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoStorageBasePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["StorageAccountId"] = gen.AlphaString()
}

func Test_BatchAccountIdentity_Status_UserAssignedIdentitiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccountIdentity_Status_UserAssignedIdentitiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccountIdentityStatusUserAssignedIdentitiesARM, BatchAccountIdentityStatusUserAssignedIdentitiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccountIdentityStatusUserAssignedIdentitiesARM runs a test to see if a specific instance of BatchAccountIdentity_Status_UserAssignedIdentitiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccountIdentityStatusUserAssignedIdentitiesARM(subject BatchAccountIdentity_Status_UserAssignedIdentitiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccountIdentity_Status_UserAssignedIdentitiesARM
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

// Generator of BatchAccountIdentity_Status_UserAssignedIdentitiesARM instances for property testing - lazily
//instantiated by BatchAccountIdentityStatusUserAssignedIdentitiesARMGenerator()
var batchAccountIdentityStatusUserAssignedIdentitiesARMGenerator gopter.Gen

// BatchAccountIdentityStatusUserAssignedIdentitiesARMGenerator returns a generator of BatchAccountIdentity_Status_UserAssignedIdentitiesARM instances for property testing.
func BatchAccountIdentityStatusUserAssignedIdentitiesARMGenerator() gopter.Gen {
	if batchAccountIdentityStatusUserAssignedIdentitiesARMGenerator != nil {
		return batchAccountIdentityStatusUserAssignedIdentitiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountIdentityStatusUserAssignedIdentitiesARM(generators)
	batchAccountIdentityStatusUserAssignedIdentitiesARMGenerator = gen.Struct(reflect.TypeOf(BatchAccountIdentity_Status_UserAssignedIdentitiesARM{}), generators)

	return batchAccountIdentityStatusUserAssignedIdentitiesARMGenerator
}

// AddIndependentPropertyGeneratorsForBatchAccountIdentityStatusUserAssignedIdentitiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBatchAccountIdentityStatusUserAssignedIdentitiesARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
}

func Test_EncryptionProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionPropertiesStatusARM, EncryptionPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionPropertiesStatusARM runs a test to see if a specific instance of EncryptionProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionPropertiesStatusARM(subject EncryptionProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionProperties_StatusARM
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

// Generator of EncryptionProperties_StatusARM instances for property testing - lazily instantiated by
//EncryptionPropertiesStatusARMGenerator()
var encryptionPropertiesStatusARMGenerator gopter.Gen

// EncryptionPropertiesStatusARMGenerator returns a generator of EncryptionProperties_StatusARM instances for property testing.
// We first initialize encryptionPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionPropertiesStatusARMGenerator() gopter.Gen {
	if encryptionPropertiesStatusARMGenerator != nil {
		return encryptionPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionPropertiesStatusARM(generators)
	encryptionPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(EncryptionProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionPropertiesStatusARM(generators)
	encryptionPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(EncryptionProperties_StatusARM{}), generators)

	return encryptionPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["KeySource"] = gen.PtrOf(gen.OneConstOf(EncryptionPropertiesStatusKeySourceMicrosoftBatch, EncryptionPropertiesStatusKeySourceMicrosoftKeyVault))
}

// AddRelatedPropertyGeneratorsForEncryptionPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultPropertiesStatusARMGenerator())
}

func Test_KeyVaultReference_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultReference_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultReferenceStatusARM, KeyVaultReferenceStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultReferenceStatusARM runs a test to see if a specific instance of KeyVaultReference_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultReferenceStatusARM(subject KeyVaultReference_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultReference_StatusARM
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

// Generator of KeyVaultReference_StatusARM instances for property testing - lazily instantiated by
//KeyVaultReferenceStatusARMGenerator()
var keyVaultReferenceStatusARMGenerator gopter.Gen

// KeyVaultReferenceStatusARMGenerator returns a generator of KeyVaultReference_StatusARM instances for property testing.
func KeyVaultReferenceStatusARMGenerator() gopter.Gen {
	if keyVaultReferenceStatusARMGenerator != nil {
		return keyVaultReferenceStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultReferenceStatusARM(generators)
	keyVaultReferenceStatusARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultReference_StatusARM{}), generators)

	return keyVaultReferenceStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultReferenceStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultReferenceStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.AlphaString()
	gens["Url"] = gen.AlphaString()
}

func Test_KeyVaultProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultPropertiesStatusARM, KeyVaultPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultPropertiesStatusARM runs a test to see if a specific instance of KeyVaultProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultPropertiesStatusARM(subject KeyVaultProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultProperties_StatusARM
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

// Generator of KeyVaultProperties_StatusARM instances for property testing - lazily instantiated by
//KeyVaultPropertiesStatusARMGenerator()
var keyVaultPropertiesStatusARMGenerator gopter.Gen

// KeyVaultPropertiesStatusARMGenerator returns a generator of KeyVaultProperties_StatusARM instances for property testing.
func KeyVaultPropertiesStatusARMGenerator() gopter.Gen {
	if keyVaultPropertiesStatusARMGenerator != nil {
		return keyVaultPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultPropertiesStatusARM(generators)
	keyVaultPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_StatusARM{}), generators)

	return keyVaultPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["KeyIdentifier"] = gen.PtrOf(gen.AlphaString())
}
