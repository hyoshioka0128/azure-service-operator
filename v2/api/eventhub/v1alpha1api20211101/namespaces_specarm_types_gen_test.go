// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

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

func Test_Namespaces_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesSPECARM, NamespacesSPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesSPECARM runs a test to see if a specific instance of Namespaces_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesSPECARM(subject Namespaces_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_SPECARM
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

// Generator of Namespaces_SPECARM instances for property testing - lazily instantiated by NamespacesSPECARMGenerator()
var namespacesSPECARMGenerator gopter.Gen

// NamespacesSPECARMGenerator returns a generator of Namespaces_SPECARM instances for property testing.
// We first initialize namespacesSPECARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesSPECARMGenerator() gopter.Gen {
	if namespacesSPECARMGenerator != nil {
		return namespacesSPECARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesSPECARM(generators)
	namespacesSPECARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesSPECARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesSPECARM(generators)
	namespacesSPECARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_SPECARM{}), generators)

	return namespacesSPECARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesSPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesSPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesSPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesSPECARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentitySpecARMGenerator())
	gens["Properties"] = gen.PtrOf(NamespacesSPECPropertiesARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuSpecARMGenerator())
}

func Test_Identity_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentitySpecARM, IdentitySpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentitySpecARM runs a test to see if a specific instance of Identity_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentitySpecARM(subject Identity_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_SpecARM
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

// Generator of Identity_SpecARM instances for property testing - lazily instantiated by IdentitySpecARMGenerator()
var identitySpecARMGenerator gopter.Gen

// IdentitySpecARMGenerator returns a generator of Identity_SpecARM instances for property testing.
func IdentitySpecARMGenerator() gopter.Gen {
	if identitySpecARMGenerator != nil {
		return identitySpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentitySpecARM(generators)
	identitySpecARMGenerator = gen.Struct(reflect.TypeOf(Identity_SpecARM{}), generators)

	return identitySpecARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentitySpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentitySpecARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		IdentitySpecTypeNone,
		IdentitySpecTypeSystemAssigned,
		IdentitySpecTypeSystemAssignedUserAssigned,
		IdentitySpecTypeUserAssigned))
}

func Test_Namespaces_SPEC_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_SPEC_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesSPECPropertiesARM, NamespacesSPECPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesSPECPropertiesARM runs a test to see if a specific instance of Namespaces_SPEC_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesSPECPropertiesARM(subject Namespaces_SPEC_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_SPEC_PropertiesARM
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

// Generator of Namespaces_SPEC_PropertiesARM instances for property testing - lazily instantiated by
//NamespacesSPECPropertiesARMGenerator()
var namespacesSPECPropertiesARMGenerator gopter.Gen

// NamespacesSPECPropertiesARMGenerator returns a generator of Namespaces_SPEC_PropertiesARM instances for property testing.
// We first initialize namespacesSPECPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesSPECPropertiesARMGenerator() gopter.Gen {
	if namespacesSPECPropertiesARMGenerator != nil {
		return namespacesSPECPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesSPECPropertiesARM(generators)
	namespacesSPECPropertiesARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_SPEC_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesSPECPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesSPECPropertiesARM(generators)
	namespacesSPECPropertiesARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_SPEC_PropertiesARM{}), generators)

	return namespacesSPECPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesSPECPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesSPECPropertiesARM(gens map[string]gopter.Gen) {
	gens["AlternateName"] = gen.PtrOf(gen.AlphaString())
	gens["ClusterArmId"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["IsAutoInflateEnabled"] = gen.PtrOf(gen.Bool())
	gens["KafkaEnabled"] = gen.PtrOf(gen.Bool())
	gens["MaximumThroughputUnits"] = gen.PtrOf(gen.Int())
	gens["ZoneRedundant"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForNamespacesSPECPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesSPECPropertiesARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionSpecARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionSpecARMGenerator())
}

func Test_Sku_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuSpecARM, SkuSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuSpecARM runs a test to see if a specific instance of Sku_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuSpecARM(subject Sku_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_SpecARM
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

// Generator of Sku_SpecARM instances for property testing - lazily instantiated by SkuSpecARMGenerator()
var skuSpecARMGenerator gopter.Gen

// SkuSpecARMGenerator returns a generator of Sku_SpecARM instances for property testing.
func SkuSpecARMGenerator() gopter.Gen {
	if skuSpecARMGenerator != nil {
		return skuSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuSpecARM(generators)
	skuSpecARMGenerator = gen.Struct(reflect.TypeOf(Sku_SpecARM{}), generators)

	return skuSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuSpecARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.OneConstOf(SkuSpecNameBasic, SkuSpecNamePremium, SkuSpecNameStandard)
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(SkuSpecTierBasic, SkuSpecTierPremium, SkuSpecTierStandard))
}

func Test_Encryption_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Encryption_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionSpecARM, EncryptionSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionSpecARM runs a test to see if a specific instance of Encryption_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionSpecARM(subject Encryption_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Encryption_SpecARM
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

// Generator of Encryption_SpecARM instances for property testing - lazily instantiated by EncryptionSpecARMGenerator()
var encryptionSpecARMGenerator gopter.Gen

// EncryptionSpecARMGenerator returns a generator of Encryption_SpecARM instances for property testing.
// We first initialize encryptionSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionSpecARMGenerator() gopter.Gen {
	if encryptionSpecARMGenerator != nil {
		return encryptionSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSpecARM(generators)
	encryptionSpecARMGenerator = gen.Struct(reflect.TypeOf(Encryption_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSpecARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionSpecARM(generators)
	encryptionSpecARMGenerator = gen.Struct(reflect.TypeOf(Encryption_SpecARM{}), generators)

	return encryptionSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionSpecARM(gens map[string]gopter.Gen) {
	gens["KeySource"] = gen.PtrOf(gen.OneConstOf(EncryptionSpecKeySourceMicrosoftKeyVault))
	gens["RequireInfrastructureEncryption"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForEncryptionSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionSpecARM(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.SliceOf(KeyVaultPropertiesSpecARMGenerator())
}

func Test_PrivateEndpointConnection_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionSpecARM, PrivateEndpointConnectionSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionSpecARM runs a test to see if a specific instance of PrivateEndpointConnection_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionSpecARM(subject PrivateEndpointConnection_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_SpecARM
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

// Generator of PrivateEndpointConnection_SpecARM instances for property testing - lazily instantiated by
//PrivateEndpointConnectionSpecARMGenerator()
var privateEndpointConnectionSpecARMGenerator gopter.Gen

// PrivateEndpointConnectionSpecARMGenerator returns a generator of PrivateEndpointConnection_SpecARM instances for property testing.
func PrivateEndpointConnectionSpecARMGenerator() gopter.Gen {
	if privateEndpointConnectionSpecARMGenerator != nil {
		return privateEndpointConnectionSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateEndpointConnectionSpecARM(generators)
	privateEndpointConnectionSpecARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_SpecARM{}), generators)

	return privateEndpointConnectionSpecARMGenerator
}

// AddRelatedPropertyGeneratorsForPrivateEndpointConnectionSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointConnectionSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateEndpointConnectionPropertiesSpecARMGenerator())
}

func Test_KeyVaultProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultPropertiesSpecARM, KeyVaultPropertiesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultPropertiesSpecARM runs a test to see if a specific instance of KeyVaultProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultPropertiesSpecARM(subject KeyVaultProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultProperties_SpecARM
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

// Generator of KeyVaultProperties_SpecARM instances for property testing - lazily instantiated by
//KeyVaultPropertiesSpecARMGenerator()
var keyVaultPropertiesSpecARMGenerator gopter.Gen

// KeyVaultPropertiesSpecARMGenerator returns a generator of KeyVaultProperties_SpecARM instances for property testing.
// We first initialize keyVaultPropertiesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultPropertiesSpecARMGenerator() gopter.Gen {
	if keyVaultPropertiesSpecARMGenerator != nil {
		return keyVaultPropertiesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultPropertiesSpecARM(generators)
	keyVaultPropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultPropertiesSpecARM(generators)
	AddRelatedPropertyGeneratorsForKeyVaultPropertiesSpecARM(generators)
	keyVaultPropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_SpecARM{}), generators)

	return keyVaultPropertiesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultPropertiesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultPropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["KeyName"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVaultUri"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultPropertiesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultPropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(UserAssignedIdentityPropertiesSpecARMGenerator())
}

func Test_PrivateEndpointConnectionProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnectionProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionPropertiesSpecARM, PrivateEndpointConnectionPropertiesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionPropertiesSpecARM runs a test to see if a specific instance of PrivateEndpointConnectionProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionPropertiesSpecARM(subject PrivateEndpointConnectionProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnectionProperties_SpecARM
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

// Generator of PrivateEndpointConnectionProperties_SpecARM instances for property testing - lazily instantiated by
//PrivateEndpointConnectionPropertiesSpecARMGenerator()
var privateEndpointConnectionPropertiesSpecARMGenerator gopter.Gen

// PrivateEndpointConnectionPropertiesSpecARMGenerator returns a generator of PrivateEndpointConnectionProperties_SpecARM instances for property testing.
// We first initialize privateEndpointConnectionPropertiesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateEndpointConnectionPropertiesSpecARMGenerator() gopter.Gen {
	if privateEndpointConnectionPropertiesSpecARMGenerator != nil {
		return privateEndpointConnectionPropertiesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionPropertiesSpecARM(generators)
	privateEndpointConnectionPropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnectionProperties_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionPropertiesSpecARM(generators)
	AddRelatedPropertyGeneratorsForPrivateEndpointConnectionPropertiesSpecARM(generators)
	privateEndpointConnectionPropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnectionProperties_SpecARM{}), generators)

	return privateEndpointConnectionPropertiesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnectionPropertiesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnectionPropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		PrivateEndpointConnectionPropertiesSpecProvisioningStateCanceled,
		PrivateEndpointConnectionPropertiesSpecProvisioningStateCreating,
		PrivateEndpointConnectionPropertiesSpecProvisioningStateDeleting,
		PrivateEndpointConnectionPropertiesSpecProvisioningStateFailed,
		PrivateEndpointConnectionPropertiesSpecProvisioningStateSucceeded,
		PrivateEndpointConnectionPropertiesSpecProvisioningStateUpdating))
}

// AddRelatedPropertyGeneratorsForPrivateEndpointConnectionPropertiesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointConnectionPropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["PrivateEndpoint"] = gen.PtrOf(PrivateEndpointSpecARMGenerator())
	gens["PrivateLinkServiceConnectionState"] = gen.PtrOf(ConnectionStateSpecARMGenerator())
}

func Test_ConnectionState_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConnectionState_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConnectionStateSpecARM, ConnectionStateSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConnectionStateSpecARM runs a test to see if a specific instance of ConnectionState_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConnectionStateSpecARM(subject ConnectionState_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConnectionState_SpecARM
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

// Generator of ConnectionState_SpecARM instances for property testing - lazily instantiated by
//ConnectionStateSpecARMGenerator()
var connectionStateSpecARMGenerator gopter.Gen

// ConnectionStateSpecARMGenerator returns a generator of ConnectionState_SpecARM instances for property testing.
func ConnectionStateSpecARMGenerator() gopter.Gen {
	if connectionStateSpecARMGenerator != nil {
		return connectionStateSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConnectionStateSpecARM(generators)
	connectionStateSpecARMGenerator = gen.Struct(reflect.TypeOf(ConnectionState_SpecARM{}), generators)

	return connectionStateSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForConnectionStateSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConnectionStateSpecARM(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		ConnectionStateSpecStatusApproved,
		ConnectionStateSpecStatusDisconnected,
		ConnectionStateSpecStatusPending,
		ConnectionStateSpecStatusRejected))
}

func Test_PrivateEndpoint_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpoint_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointSpecARM, PrivateEndpointSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointSpecARM runs a test to see if a specific instance of PrivateEndpoint_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointSpecARM(subject PrivateEndpoint_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpoint_SpecARM
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

// Generator of PrivateEndpoint_SpecARM instances for property testing - lazily instantiated by
//PrivateEndpointSpecARMGenerator()
var privateEndpointSpecARMGenerator gopter.Gen

// PrivateEndpointSpecARMGenerator returns a generator of PrivateEndpoint_SpecARM instances for property testing.
func PrivateEndpointSpecARMGenerator() gopter.Gen {
	if privateEndpointSpecARMGenerator != nil {
		return privateEndpointSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointSpecARM(generators)
	privateEndpointSpecARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpoint_SpecARM{}), generators)

	return privateEndpointSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointSpecARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentityProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityPropertiesSpecARM, UserAssignedIdentityPropertiesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityPropertiesSpecARM runs a test to see if a specific instance of UserAssignedIdentityProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityPropertiesSpecARM(subject UserAssignedIdentityProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityProperties_SpecARM
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

// Generator of UserAssignedIdentityProperties_SpecARM instances for property testing - lazily instantiated by
//UserAssignedIdentityPropertiesSpecARMGenerator()
var userAssignedIdentityPropertiesSpecARMGenerator gopter.Gen

// UserAssignedIdentityPropertiesSpecARMGenerator returns a generator of UserAssignedIdentityProperties_SpecARM instances for property testing.
func UserAssignedIdentityPropertiesSpecARMGenerator() gopter.Gen {
	if userAssignedIdentityPropertiesSpecARMGenerator != nil {
		return userAssignedIdentityPropertiesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesSpecARM(generators)
	userAssignedIdentityPropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityProperties_SpecARM{}), generators)

	return userAssignedIdentityPropertiesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
}
