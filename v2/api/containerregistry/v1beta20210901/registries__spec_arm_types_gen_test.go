// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210901

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

func Test_Registries_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Registries_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRegistriesSpecARM, RegistriesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRegistriesSpecARM runs a test to see if a specific instance of Registries_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRegistriesSpecARM(subject Registries_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Registries_SpecARM
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

// Generator of Registries_SpecARM instances for property testing - lazily instantiated by RegistriesSpecARMGenerator()
var registriesSpecARMGenerator gopter.Gen

// RegistriesSpecARMGenerator returns a generator of Registries_SpecARM instances for property testing.
// We first initialize registriesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RegistriesSpecARMGenerator() gopter.Gen {
	if registriesSpecARMGenerator != nil {
		return registriesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRegistriesSpecARM(generators)
	registriesSpecARMGenerator = gen.Struct(reflect.TypeOf(Registries_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRegistriesSpecARM(generators)
	AddRelatedPropertyGeneratorsForRegistriesSpecARM(generators)
	registriesSpecARMGenerator = gen.Struct(reflect.TypeOf(Registries_SpecARM{}), generators)

	return registriesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForRegistriesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRegistriesSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRegistriesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRegistriesSpecARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentityPropertiesARMGenerator())
	gens["Properties"] = gen.PtrOf(RegistryPropertiesARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuARMGenerator())
}

func Test_IdentityPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IdentityPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentityPropertiesARM, IdentityPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentityPropertiesARM runs a test to see if a specific instance of IdentityPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentityPropertiesARM(subject IdentityPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IdentityPropertiesARM
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

// Generator of IdentityPropertiesARM instances for property testing - lazily instantiated by
// IdentityPropertiesARMGenerator()
var identityPropertiesARMGenerator gopter.Gen

// IdentityPropertiesARMGenerator returns a generator of IdentityPropertiesARM instances for property testing.
// We first initialize identityPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IdentityPropertiesARMGenerator() gopter.Gen {
	if identityPropertiesARMGenerator != nil {
		return identityPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityPropertiesARM(generators)
	identityPropertiesARMGenerator = gen.Struct(reflect.TypeOf(IdentityPropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForIdentityPropertiesARM(generators)
	identityPropertiesARMGenerator = gen.Struct(reflect.TypeOf(IdentityPropertiesARM{}), generators)

	return identityPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentityPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentityPropertiesARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		IdentityPropertiesTypeNone,
		IdentityPropertiesTypeSystemAssigned,
		IdentityPropertiesTypeSystemAssignedUserAssigned,
		IdentityPropertiesTypeUserAssigned))
}

// AddRelatedPropertyGeneratorsForIdentityPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIdentityPropertiesARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(gen.AlphaString(), UserIdentityPropertiesARMGenerator())
}

func Test_RegistryPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RegistryPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRegistryPropertiesARM, RegistryPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRegistryPropertiesARM runs a test to see if a specific instance of RegistryPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRegistryPropertiesARM(subject RegistryPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RegistryPropertiesARM
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

// Generator of RegistryPropertiesARM instances for property testing - lazily instantiated by
// RegistryPropertiesARMGenerator()
var registryPropertiesARMGenerator gopter.Gen

// RegistryPropertiesARMGenerator returns a generator of RegistryPropertiesARM instances for property testing.
// We first initialize registryPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RegistryPropertiesARMGenerator() gopter.Gen {
	if registryPropertiesARMGenerator != nil {
		return registryPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRegistryPropertiesARM(generators)
	registryPropertiesARMGenerator = gen.Struct(reflect.TypeOf(RegistryPropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRegistryPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForRegistryPropertiesARM(generators)
	registryPropertiesARMGenerator = gen.Struct(reflect.TypeOf(RegistryPropertiesARM{}), generators)

	return registryPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForRegistryPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRegistryPropertiesARM(gens map[string]gopter.Gen) {
	gens["AdminUserEnabled"] = gen.PtrOf(gen.Bool())
	gens["DataEndpointEnabled"] = gen.PtrOf(gen.Bool())
	gens["NetworkRuleBypassOptions"] = gen.PtrOf(gen.OneConstOf(RegistryPropertiesNetworkRuleBypassOptionsAzureServices, RegistryPropertiesNetworkRuleBypassOptionsNone))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(RegistryPropertiesPublicNetworkAccessDisabled, RegistryPropertiesPublicNetworkAccessEnabled))
	gens["ZoneRedundancy"] = gen.PtrOf(gen.OneConstOf(RegistryPropertiesZoneRedundancyDisabled, RegistryPropertiesZoneRedundancyEnabled))
}

// AddRelatedPropertyGeneratorsForRegistryPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRegistryPropertiesARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionPropertyARMGenerator())
	gens["NetworkRuleSet"] = gen.PtrOf(NetworkRuleSetARMGenerator())
	gens["Policies"] = gen.PtrOf(PoliciesARMGenerator())
}

func Test_SkuARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SkuARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuARM, SkuARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuARM runs a test to see if a specific instance of SkuARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuARM(subject SkuARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SkuARM
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

// Generator of SkuARM instances for property testing - lazily instantiated by SkuARMGenerator()
var skuARMGenerator gopter.Gen

// SkuARMGenerator returns a generator of SkuARM instances for property testing.
func SkuARMGenerator() gopter.Gen {
	if skuARMGenerator != nil {
		return skuARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuARM(generators)
	skuARMGenerator = gen.Struct(reflect.TypeOf(SkuARM{}), generators)

	return skuARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		SkuNameBasic,
		SkuNameClassic,
		SkuNamePremium,
		SkuNameStandard))
}

func Test_EncryptionPropertyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionPropertyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionPropertyARM, EncryptionPropertyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionPropertyARM runs a test to see if a specific instance of EncryptionPropertyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionPropertyARM(subject EncryptionPropertyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionPropertyARM
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

// Generator of EncryptionPropertyARM instances for property testing - lazily instantiated by
// EncryptionPropertyARMGenerator()
var encryptionPropertyARMGenerator gopter.Gen

// EncryptionPropertyARMGenerator returns a generator of EncryptionPropertyARM instances for property testing.
// We first initialize encryptionPropertyARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionPropertyARMGenerator() gopter.Gen {
	if encryptionPropertyARMGenerator != nil {
		return encryptionPropertyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionPropertyARM(generators)
	encryptionPropertyARMGenerator = gen.Struct(reflect.TypeOf(EncryptionPropertyARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionPropertyARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionPropertyARM(generators)
	encryptionPropertyARMGenerator = gen.Struct(reflect.TypeOf(EncryptionPropertyARM{}), generators)

	return encryptionPropertyARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionPropertyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionPropertyARM(gens map[string]gopter.Gen) {
	gens["Status"] = gen.PtrOf(gen.OneConstOf(EncryptionPropertyStatusDisabled, EncryptionPropertyStatusEnabled))
}

// AddRelatedPropertyGeneratorsForEncryptionPropertyARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionPropertyARM(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultPropertiesARMGenerator())
}

func Test_NetworkRuleSetARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkRuleSetARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkRuleSetARM, NetworkRuleSetARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkRuleSetARM runs a test to see if a specific instance of NetworkRuleSetARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkRuleSetARM(subject NetworkRuleSetARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkRuleSetARM
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

// Generator of NetworkRuleSetARM instances for property testing - lazily instantiated by NetworkRuleSetARMGenerator()
var networkRuleSetARMGenerator gopter.Gen

// NetworkRuleSetARMGenerator returns a generator of NetworkRuleSetARM instances for property testing.
// We first initialize networkRuleSetARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkRuleSetARMGenerator() gopter.Gen {
	if networkRuleSetARMGenerator != nil {
		return networkRuleSetARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkRuleSetARM(generators)
	networkRuleSetARMGenerator = gen.Struct(reflect.TypeOf(NetworkRuleSetARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkRuleSetARM(generators)
	AddRelatedPropertyGeneratorsForNetworkRuleSetARM(generators)
	networkRuleSetARMGenerator = gen.Struct(reflect.TypeOf(NetworkRuleSetARM{}), generators)

	return networkRuleSetARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkRuleSetARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkRuleSetARM(gens map[string]gopter.Gen) {
	gens["DefaultAction"] = gen.PtrOf(gen.OneConstOf(NetworkRuleSetDefaultActionAllow, NetworkRuleSetDefaultActionDeny))
}

// AddRelatedPropertyGeneratorsForNetworkRuleSetARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkRuleSetARM(gens map[string]gopter.Gen) {
	gens["IpRules"] = gen.SliceOf(IPRuleARMGenerator())
}

func Test_PoliciesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PoliciesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPoliciesARM, PoliciesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPoliciesARM runs a test to see if a specific instance of PoliciesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPoliciesARM(subject PoliciesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PoliciesARM
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

// Generator of PoliciesARM instances for property testing - lazily instantiated by PoliciesARMGenerator()
var policiesARMGenerator gopter.Gen

// PoliciesARMGenerator returns a generator of PoliciesARM instances for property testing.
func PoliciesARMGenerator() gopter.Gen {
	if policiesARMGenerator != nil {
		return policiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPoliciesARM(generators)
	policiesARMGenerator = gen.Struct(reflect.TypeOf(PoliciesARM{}), generators)

	return policiesARMGenerator
}

// AddRelatedPropertyGeneratorsForPoliciesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPoliciesARM(gens map[string]gopter.Gen) {
	gens["ExportPolicy"] = gen.PtrOf(ExportPolicyARMGenerator())
	gens["QuarantinePolicy"] = gen.PtrOf(QuarantinePolicyARMGenerator())
	gens["RetentionPolicy"] = gen.PtrOf(RetentionPolicyARMGenerator())
	gens["TrustPolicy"] = gen.PtrOf(TrustPolicyARMGenerator())
}

func Test_UserIdentityPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserIdentityPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserIdentityPropertiesARM, UserIdentityPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserIdentityPropertiesARM runs a test to see if a specific instance of UserIdentityPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserIdentityPropertiesARM(subject UserIdentityPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserIdentityPropertiesARM
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

// Generator of UserIdentityPropertiesARM instances for property testing - lazily instantiated by
// UserIdentityPropertiesARMGenerator()
var userIdentityPropertiesARMGenerator gopter.Gen

// UserIdentityPropertiesARMGenerator returns a generator of UserIdentityPropertiesARM instances for property testing.
func UserIdentityPropertiesARMGenerator() gopter.Gen {
	if userIdentityPropertiesARMGenerator != nil {
		return userIdentityPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserIdentityPropertiesARM(generators)
	userIdentityPropertiesARMGenerator = gen.Struct(reflect.TypeOf(UserIdentityPropertiesARM{}), generators)

	return userIdentityPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForUserIdentityPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserIdentityPropertiesARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
}

func Test_ExportPolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExportPolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExportPolicyARM, ExportPolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExportPolicyARM runs a test to see if a specific instance of ExportPolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExportPolicyARM(subject ExportPolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExportPolicyARM
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

// Generator of ExportPolicyARM instances for property testing - lazily instantiated by ExportPolicyARMGenerator()
var exportPolicyARMGenerator gopter.Gen

// ExportPolicyARMGenerator returns a generator of ExportPolicyARM instances for property testing.
func ExportPolicyARMGenerator() gopter.Gen {
	if exportPolicyARMGenerator != nil {
		return exportPolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExportPolicyARM(generators)
	exportPolicyARMGenerator = gen.Struct(reflect.TypeOf(ExportPolicyARM{}), generators)

	return exportPolicyARMGenerator
}

// AddIndependentPropertyGeneratorsForExportPolicyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExportPolicyARM(gens map[string]gopter.Gen) {
	gens["Status"] = gen.PtrOf(gen.OneConstOf(ExportPolicyStatusDisabled, ExportPolicyStatusEnabled))
}

func Test_IPRuleARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IPRuleARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIPRuleARM, IPRuleARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIPRuleARM runs a test to see if a specific instance of IPRuleARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIPRuleARM(subject IPRuleARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IPRuleARM
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

// Generator of IPRuleARM instances for property testing - lazily instantiated by IPRuleARMGenerator()
var ipRuleARMGenerator gopter.Gen

// IPRuleARMGenerator returns a generator of IPRuleARM instances for property testing.
func IPRuleARMGenerator() gopter.Gen {
	if ipRuleARMGenerator != nil {
		return ipRuleARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIPRuleARM(generators)
	ipRuleARMGenerator = gen.Struct(reflect.TypeOf(IPRuleARM{}), generators)

	return ipRuleARMGenerator
}

// AddIndependentPropertyGeneratorsForIPRuleARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIPRuleARM(gens map[string]gopter.Gen) {
	gens["Action"] = gen.PtrOf(gen.OneConstOf(IPRuleActionAllow))
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultPropertiesARM, KeyVaultPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultPropertiesARM runs a test to see if a specific instance of KeyVaultPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultPropertiesARM(subject KeyVaultPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultPropertiesARM
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

// Generator of KeyVaultPropertiesARM instances for property testing - lazily instantiated by
// KeyVaultPropertiesARMGenerator()
var keyVaultPropertiesARMGenerator gopter.Gen

// KeyVaultPropertiesARMGenerator returns a generator of KeyVaultPropertiesARM instances for property testing.
func KeyVaultPropertiesARMGenerator() gopter.Gen {
	if keyVaultPropertiesARMGenerator != nil {
		return keyVaultPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM(generators)
	keyVaultPropertiesARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultPropertiesARM{}), generators)

	return keyVaultPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(gen.AlphaString())
	gens["KeyIdentifier"] = gen.PtrOf(gen.AlphaString())
}

func Test_QuarantinePolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of QuarantinePolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForQuarantinePolicyARM, QuarantinePolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForQuarantinePolicyARM runs a test to see if a specific instance of QuarantinePolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForQuarantinePolicyARM(subject QuarantinePolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual QuarantinePolicyARM
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

// Generator of QuarantinePolicyARM instances for property testing - lazily instantiated by
// QuarantinePolicyARMGenerator()
var quarantinePolicyARMGenerator gopter.Gen

// QuarantinePolicyARMGenerator returns a generator of QuarantinePolicyARM instances for property testing.
func QuarantinePolicyARMGenerator() gopter.Gen {
	if quarantinePolicyARMGenerator != nil {
		return quarantinePolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForQuarantinePolicyARM(generators)
	quarantinePolicyARMGenerator = gen.Struct(reflect.TypeOf(QuarantinePolicyARM{}), generators)

	return quarantinePolicyARMGenerator
}

// AddIndependentPropertyGeneratorsForQuarantinePolicyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForQuarantinePolicyARM(gens map[string]gopter.Gen) {
	gens["Status"] = gen.PtrOf(gen.OneConstOf(QuarantinePolicyStatusDisabled, QuarantinePolicyStatusEnabled))
}

func Test_RetentionPolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RetentionPolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRetentionPolicyARM, RetentionPolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRetentionPolicyARM runs a test to see if a specific instance of RetentionPolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRetentionPolicyARM(subject RetentionPolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RetentionPolicyARM
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

// Generator of RetentionPolicyARM instances for property testing - lazily instantiated by RetentionPolicyARMGenerator()
var retentionPolicyARMGenerator gopter.Gen

// RetentionPolicyARMGenerator returns a generator of RetentionPolicyARM instances for property testing.
func RetentionPolicyARMGenerator() gopter.Gen {
	if retentionPolicyARMGenerator != nil {
		return retentionPolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRetentionPolicyARM(generators)
	retentionPolicyARMGenerator = gen.Struct(reflect.TypeOf(RetentionPolicyARM{}), generators)

	return retentionPolicyARMGenerator
}

// AddIndependentPropertyGeneratorsForRetentionPolicyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRetentionPolicyARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(RetentionPolicyStatusDisabled, RetentionPolicyStatusEnabled))
}

func Test_TrustPolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrustPolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrustPolicyARM, TrustPolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrustPolicyARM runs a test to see if a specific instance of TrustPolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTrustPolicyARM(subject TrustPolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrustPolicyARM
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

// Generator of TrustPolicyARM instances for property testing - lazily instantiated by TrustPolicyARMGenerator()
var trustPolicyARMGenerator gopter.Gen

// TrustPolicyARMGenerator returns a generator of TrustPolicyARM instances for property testing.
func TrustPolicyARMGenerator() gopter.Gen {
	if trustPolicyARMGenerator != nil {
		return trustPolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrustPolicyARM(generators)
	trustPolicyARMGenerator = gen.Struct(reflect.TypeOf(TrustPolicyARM{}), generators)

	return trustPolicyARMGenerator
}

// AddIndependentPropertyGeneratorsForTrustPolicyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrustPolicyARM(gens map[string]gopter.Gen) {
	gens["Status"] = gen.PtrOf(gen.OneConstOf(TrustPolicyStatusDisabled, TrustPolicyStatusEnabled))
	gens["Type"] = gen.PtrOf(gen.OneConstOf(TrustPolicyTypeNotary))
}
